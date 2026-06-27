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

// notifyingServer pushes a server->client window/logMessage notification from
// inside its Hover handler. It models the common case of a server that emits
// log notifications early in a session, which is what exposed the connection
// teardown bug fixed in this change.
type notifyingServer struct {
	UnimplementedServer

	mu     sync.Mutex
	client Client
}

var _ Server = (*notifyingServer)(nil)

func (s *notifyingServer) Hover(ctx context.Context, _ *HoverParams) (*Hover, error) {
	s.mu.Lock()
	client := s.client
	s.mu.Unlock()

	// A server->client notification: with the bug present, the client's
	// un-overridden UnimplementedClient.LogMessage returned errNotImplemented,
	// which the dispatcher escalated to conn.fail and tore the connection down.
	if err := client.LogMessage(ctx, &LogMessageParams{Type: MessageTypeInfo, Message: "hello"}); err != nil {
		return nil, err
	}

	return &Hover{Contents: String("ok")}, nil
}

// TestIntegrationUnimplementedClientSurvivesNotification is the regression guard
// for the bug where an un-overridden UnimplementedClient notification method
// returned errNotImplemented, causing the jsonrpc2 dispatcher to fail the whole
// connection. A bare UnimplementedClient (overriding nothing) must absorb a
// server-sent window/logMessage notification and the connection must remain
// usable for a subsequent request.
func TestIntegrationUnimplementedClientSurvivesNotification(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	a, b := net.Pipe()

	ns := &notifyingServer{}

	_, connA, clientDispatcher := NewServer(ctx, ns, jsonrpc2.NewStream(a))
	defer func() { _ = connA.Close() }()

	ns.mu.Lock()
	ns.client = clientDispatcher
	ns.mu.Unlock()

	// The client overrides nothing: every notification falls through to
	// UnimplementedClient, the exact configuration that previously killed the
	// connection on the first server notification.
	_, connB, serverDispatcher := NewClient(ctx, UnimplementedClient{}, jsonrpc2.NewStream(b))
	defer func() { _ = connB.Close() }()

	// The handler pushes window/logMessage to the client mid-request. If the
	// notification tore the connection down, this Hover call would fail (or the
	// LogMessage inside the handler would error and surface here).
	hover, err := serverDispatcher.Hover(ctx, &HoverParams{})
	if err != nil {
		t.Fatalf("hover call after server notification: %v", err)
	}
	if contents, ok := hover.Contents.(String); !ok || contents != "ok" {
		t.Fatalf("hover contents = %#v, want String(%q)", hover.Contents, "ok")
	}

	// Prove the connection is still alive after the notification by issuing a
	// second request that must also round-trip.
	if _, err := serverDispatcher.Hover(ctx, &HoverParams{}); err != nil {
		t.Fatalf("second hover call: connection did not survive the notification: %v", err)
	}
}
