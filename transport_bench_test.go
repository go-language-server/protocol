// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"net"
	"testing"
	"time"

	"go.lsp.dev/jsonrpc2"
)

// benchTransportMethod is the private method the transport benchmark serves;
// it never collides with a real LSP method.
const benchTransportMethod = "bench/publishDiagnostics"

// BenchmarkTransportCall measures a full client->server->client round trip
// over the real jsonrpc2 header stream with the production lspCodec: request
// encode, frame, decode, dispatch, response encode with a corpus-sized
// publishDiagnostics payload, frame, and decode. It is the Phase-3 M5
// transport baseline: allocs/op spans both peers and the wire.
func BenchmarkTransportCall(b *testing.B) {
	payload := benchCorpus(b, "publish_diagnostics")
	var diag PublishDiagnosticsParams
	if err := Unmarshal(payload, &diag); err != nil {
		b.Fatalf("decode corpus: %v", err)
	}

	ctx, cancel := context.WithTimeout(b.Context(), time.Minute)
	defer cancel()

	cliEnd, srvEnd := net.Pipe()
	clientConn := jsonrpc2.NewConn(jsonrpc2.NewStream(cliEnd), jsonrpc2.WithCodec(lspCodec{}))
	serverConn := jsonrpc2.NewConn(jsonrpc2.NewStream(srvEnd), jsonrpc2.WithCodec(lspCodec{}))

	serverHandler := func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		if req.Method() == benchTransportMethod {
			return reply(ctx, &diag, nil)
		}
		return jsonrpc2.MethodNotFoundHandler(ctx, reply, req)
	}
	clientConn.Go(ctx, jsonrpc2.MethodNotFoundHandler)
	serverConn.Go(ctx, serverHandler)
	defer func() { _ = clientConn.Close() }()
	defer func() { _ = serverConn.Close() }()

	b.ReportAllocs()
	b.SetBytes(int64(len(payload)))
	for b.Loop() {
		var out PublishDiagnosticsParams
		if _, err := clientConn.Call(ctx, benchTransportMethod, nil, &out); err != nil {
			b.Fatalf("call: %v", err)
		}
	}
}
