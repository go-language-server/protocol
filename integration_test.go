// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/uri"
)

// testServer is a partial [Server] that exercises the production dispatch path.
// It embeds [UnimplementedServer] so only the methods under test are overridden,
// and holds the [Client] dispatcher returned by [NewServer] to drive the
// server->client callback from inside a request handler.
type testServer struct {
	UnimplementedServer

	mu         sync.Mutex
	client     Client  // dispatcher used for the server->client applyEdit.
	didOpenURI uri.URI // recorded by DidOpen.
	applied    bool    // recorded from the s2c ApplyEdit response.
}

// compile-time assertion that *testServer satisfies Server.
var _ Server = (*testServer)(nil)

// Hover returns a fixed result and, before replying, issues a server->client
// workspace/applyEdit request through the stored client dispatcher to prove the
// reentrant s2c path round-trips. The handler is dispatched through Handlers(),
// which wraps jsonrpc2.AsyncHandler; AsyncHandler releases the read loop before
// the handler body runs, so the callback does not need an explicit
// jsonrpc2.Async and does not deadlock the connection's read loop.
func (s *testServer) Hover(ctx context.Context, _ *HoverParams) (*Hover, error) {
	s.mu.Lock()
	client := s.client
	s.mu.Unlock()

	res, err := client.ApplyEdit(ctx, &ApplyWorkspaceEditParams{Edit: WorkspaceEdit{}})
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	s.applied = res.Applied
	s.mu.Unlock()

	return &Hover{Contents: String("integration hover")}, nil
}

// DidOpen records the opened document URI under the mutex.
func (s *testServer) DidOpen(_ context.Context, params *DidOpenTextDocumentParams) error {
	s.mu.Lock()
	s.didOpenURI = params.TextDocument.URI
	s.mu.Unlock()

	return nil
}

// testClient is a partial [Client] that services the server-initiated
// workspace/applyEdit request.
type testClient struct {
	UnimplementedClient
}

// compile-time assertion that *testClient satisfies Client.
var _ Client = (*testClient)(nil)

// ApplyEdit reports the edit as applied so the s2c round-trip can be asserted.
func (testClient) ApplyEdit(context.Context, *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error) {
	return &ApplyWorkspaceEditResult{Applied: true}, nil
}

// TestIntegrationRoundTrip drives the real NewServer/NewClient transport over an
// in-memory net.Pipe and proves three end-to-end behaviors through the
// production dispatch path:
//
//  1. a client->server request with a result round-trips (textDocument/hover);
//  2. a client->server notification reaches the server (textDocument/didOpen);
//  3. a server->client request issued from within a handler round-trips
//     (workspace/applyEdit).
//
// TestIntegrationRoundTrip drives the production Handlers() chain through
// NewServer/NewClient over an in-memory pipe: a client->server request, a
// notification, and a server->client callback. It is the regression guard for
// the reply-clobbering defect once present in protocol.CancelHandler.
func TestIntegrationRoundTrip(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	a, b := net.Pipe()

	ts := &testServer{}
	tc := &testClient{}

	// Endpoint A speaks the server role: it serves *testServer requests and
	// returns clientDispatcher, the Client used for server->client calls.
	_, connA, clientDispatcher := NewServer(ctx, ts, jsonrpc2.NewStream(a))
	defer func() { _ = connA.Close() }()

	ts.mu.Lock()
	ts.client = clientDispatcher
	ts.mu.Unlock()

	// Endpoint B speaks the client role: it serves *testClient requests and
	// returns serverDispatcher, the Server used to drive client->server calls.
	_, connB, serverDispatcher := NewClient(ctx, tc, jsonrpc2.NewStream(b))
	defer func() { _ = connB.Close() }()

	// (2) notification reaches the server.
	const wantURI uri.URI = "file:///integration.go"
	if err := serverDispatcher.DidOpen(ctx, &DidOpenTextDocumentParams{
		TextDocument: TextDocumentItem{
			URI:        wantURI,
			LanguageID: "go",
			Version:    1,
			Text:       "package integration",
		},
	}); err != nil {
		t.Fatalf("didOpen notify: %v", err)
	}

	// (1) request with a result + (3) server->client callback from the handler.
	hover, err := serverDispatcher.Hover(ctx, &HoverParams{})
	if err != nil {
		t.Fatalf("hover call: %v", err)
	}
	if hover == nil {
		t.Fatal("hover result is nil")
	}
	contents, ok := hover.Contents.(String)
	if !ok {
		t.Fatalf("hover contents: want String, got %T", hover.Contents)
	}
	if contents != "integration hover" {
		t.Errorf("hover contents = %q, want %q", contents, "integration hover")
	}

	ts.mu.Lock()
	gotURI := ts.didOpenURI
	gotApplied := ts.applied
	ts.mu.Unlock()

	// (3) server->client reentrancy.
	if !gotApplied {
		t.Error("server->client applyEdit did not round-trip (Applied=false)")
	}

	// (2) notification arrival. net.Pipe preserves ordering on a single stream,
	// and the hover request that follows it has already completed, so the
	// notification has been dispatched by now.
	if gotURI != wantURI {
		t.Errorf("didOpen URI = %q, want %q", gotURI, wantURI)
	}
}
