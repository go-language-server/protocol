// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"go.lsp.dev/jsonrpc2"
)

// spikeCodec is the spike's union-aware [jsonrpc2.Codec]: it routes message
// payloads through the generated union-aware [Marshal]/[Unmarshal] while passing
// already-encoded [jsonrpc2.RawMessage] values through verbatim (mirroring the
// jsonrpc2 JSONCodec passthrough contract). The migration's real codec.go will
// replace this; the spike proves the seam works end-to-end.
type spikeCodec struct{}

func (spikeCodec) Marshal(v any) ([]byte, error) {
	switch m := v.(type) {
	case jsonrpc2.RawMessage:
		if m == nil {
			return []byte("null"), nil
		}
		return m, nil
	case *jsonrpc2.RawMessage:
		if m == nil || *m == nil {
			return []byte("null"), nil
		}
		return *m, nil
	}
	return Marshal(v)
}

func (spikeCodec) Unmarshal(data []byte, v any) error {
	if p, ok := v.(*jsonrpc2.RawMessage); ok {
		b := make(jsonrpc2.RawMessage, len(data))
		copy(b, data)
		*p = b
		return nil
	}
	return Unmarshal(data, v)
}

// TestSpikeTransportRoundTrip is the Phase-0 GATE of the LSP 3.18 root migration
// (see .omc/plans/lsp-3.18-root-migration.md). It proves the three load-bearing
// behaviors before the bulk dispatch port:
//
//  1. a client->server request with a union result ([]CommandOrCodeAction)
//     round-trips through the WithCodec seam and discriminates to the correct
//     dynamic arms (*Command, *CodeAction);
//  2. a client->server notification reaches the server handler;
//  3. a server->client request issued from within a handler (released with
//     jsonrpc2.Async) round-trips without deadlocking the read loop.
func TestSpikeTransportRoundTrip(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	cliEnd, srvEnd := net.Pipe()

	clientConn := jsonrpc2.NewConn(jsonrpc2.NewStream(cliEnd), jsonrpc2.WithCodec(spikeCodec{}))
	serverConn := jsonrpc2.NewConn(jsonrpc2.NewStream(srvEnd), jsonrpc2.WithCodec(spikeCodec{}))

	var (
		mu         sync.Mutex
		gotDidOpen DocumentURI
		gotApplied bool
	)

	// Client handler services the server-initiated workspace/applyEdit request.
	clientHandler := func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		if req.Method() == MethodWorkspaceApplyEdit {
			var params ApplyWorkspaceEditParams
			if err := Unmarshal(req.Params(), &params); err != nil {
				return reply(ctx, nil, err)
			}
			return reply(ctx, &ApplyWorkspaceEditResult{Applied: true}, nil)
		}
		return jsonrpc2.MethodNotFoundHandler(ctx, reply, req)
	}

	// Server handler services the codeAction request (with an s2c callback) and
	// the didOpen notification.
	serverHandler := func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		switch req.Method() {
		case MethodTextDocumentDidOpen:
			var params DidOpenTextDocumentParams
			if err := Unmarshal(req.Params(), &params); err != nil {
				return reply(ctx, nil, err)
			}
			mu.Lock()
			gotDidOpen = params.TextDocument.URI
			mu.Unlock()
			return reply(ctx, nil, nil) // no-op for a notification

		case MethodTextDocumentCodeAction:
			var params CodeActionParams
			if err := Unmarshal(req.Params(), &params); err != nil {
				return reply(ctx, nil, err)
			}
			// Release the read loop so the s2c callback's response can be read.
			jsonrpc2.Async(ctx)

			var applied ApplyWorkspaceEditResult
			if _, err := serverConn.Call(ctx, MethodWorkspaceApplyEdit,
				&ApplyWorkspaceEditParams{Edit: WorkspaceEdit{}}, &applied); err != nil {
				return reply(ctx, nil, err)
			}
			mu.Lock()
			gotApplied = applied.Applied
			mu.Unlock()

			result := []CommandOrCodeAction{
				&Command{Title: "do it", Command: "spike.run"},
				&CodeAction{Title: "fix it"},
			}
			return reply(ctx, result, nil)

		default:
			return jsonrpc2.MethodNotFoundHandler(ctx, reply, req)
		}
	}

	clientConn.Go(ctx, clientHandler)
	serverConn.Go(ctx, serverHandler)
	defer func() { _ = clientConn.Close() }()
	defer func() { _ = serverConn.Close() }()

	// (2) notification — ordered before the call on the same stream.
	if err := clientConn.Notify(ctx, MethodTextDocumentDidOpen, &DidOpenTextDocumentParams{
		TextDocument: TextDocumentItem{URI: "file:///spike.go", LanguageID: "go", Version: 1, Text: "package spike"},
	}); err != nil {
		t.Fatalf("didOpen notify: %v", err)
	}

	// (1)+(3) request with union result + server->client callback.
	var result []CommandOrCodeAction
	if _, err := clientConn.Call(ctx, MethodTextDocumentCodeAction, &CodeActionParams{}, &result); err != nil {
		t.Fatalf("codeAction call: %v", err)
	}

	// (1) union discrimination.
	if len(result) != 2 {
		t.Fatalf("want 2 code-action arms, got %d", len(result))
	}
	cmd, ok := result[0].(*Command)
	if !ok {
		t.Fatalf("arm 0: want *Command, got %T", result[0])
	}
	if cmd.Command != "spike.run" {
		t.Errorf("arm 0 Command = %q, want %q", cmd.Command, "spike.run")
	}
	ca, ok := result[1].(*CodeAction)
	if !ok {
		t.Fatalf("arm 1: want *CodeAction, got %T", result[1])
	}
	if ca.Title != "fix it" {
		t.Errorf("arm 1 Title = %q, want %q", ca.Title, "fix it")
	}

	mu.Lock()
	defer mu.Unlock()
	// (3) server->client reentrancy.
	if !gotApplied {
		t.Error("server->client applyEdit did not round-trip (Applied=false)")
	}
	// (2) notification arrival.
	if gotDidOpen != "file:///spike.go" {
		t.Errorf("didOpen URI = %q, want %q", gotDidOpen, "file:///spike.go")
	}
}
