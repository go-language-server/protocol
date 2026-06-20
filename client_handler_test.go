// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"errors"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"go.lsp.dev/jsonrpc2"
)

type recordingClient struct {
	UnimplementedClient

	mu sync.Mutex

	showMessageRequestParams *ShowMessageRequestParams
	showMessageRequestResult *MessageActionItem
	logMessageParams         *LogMessageParams

	showMessageRequestCalled chan struct{}
	logMessageCalled         chan struct{}
}

func newRecordingClient() *recordingClient {
	return &recordingClient{
		showMessageRequestCalled: make(chan struct{}, 1),
		logMessageCalled:         make(chan struct{}, 1),
	}
}

func (c *recordingClient) ShowMessageRequest(_ context.Context, params *ShowMessageRequestParams) (*MessageActionItem, error) {
	c.mu.Lock()
	c.showMessageRequestParams = params
	result := c.showMessageRequestResult
	c.mu.Unlock()
	c.signal(c.showMessageRequestCalled)

	return result, nil
}

func (c *recordingClient) LogMessage(_ context.Context, params *LogMessageParams) error {
	c.mu.Lock()
	c.logMessageParams = params
	c.mu.Unlock()
	c.signal(c.logMessageCalled)

	return nil
}

func (c *recordingClient) signal(ch chan struct{}) {
	select {
	case ch <- struct{}{}:
	default:
	}
}

func (c *recordingClient) snapshotShowMessageRequestParams() *ShowMessageRequestParams {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.showMessageRequestParams
}

func (c *recordingClient) snapshotLogMessageParams() *LogMessageParams {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.logMessageParams
}

func (c *recordingClient) setShowMessageRequestResult(result *MessageActionItem) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.showMessageRequestResult = result
}

func TestClientHandlerShowMessageRequestReturnsDirectResult(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	client := newRecordingClient()
	wantResult := &MessageActionItem{Title: "Apply"}
	client.setShowMessageRequestResult(wantResult)
	caller, served := newClientHandlerConnPair(ctx, t, client, jsonrpc2.MethodNotFoundHandler)
	defer closeJSONRPCConns(t, caller, served)

	params := ShowMessageRequestParams{
		Type:    MessageTypeInfo,
		Message: "Proceed?",
		Actions: []MessageActionItem{{Title: "Apply"}},
	}
	var got MessageActionItem
	if _, err := caller.Call(ctx, MethodWindowShowMessageRequest, params, &got); err != nil {
		t.Fatalf("Call(%s): %v", MethodWindowShowMessageRequest, err)
	}
	if diff := cmp.Diff(*wantResult, got); diff != "" {
		t.Fatalf("ShowMessageRequest result mismatch (-want +got):\n%s", diff)
	}
	select {
	case <-client.showMessageRequestCalled:
	case <-ctx.Done():
		t.Fatalf("ShowMessageRequest was not called: %v", ctx.Err())
	}
	if diff := cmp.Diff(&params, client.snapshotShowMessageRequestParams()); diff != "" {
		t.Fatalf("ShowMessageRequest params mismatch (-want +got):\n%s", diff)
	}
}

func TestClientHandlerLogMessageNotificationReturnsNilResult(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	client := newRecordingClient()
	caller, served := newClientHandlerConnPair(ctx, t, client, jsonrpc2.MethodNotFoundHandler)
	defer closeJSONRPCConns(t, caller, served)

	params := LogMessageParams{Type: MessageTypeLog, Message: "ready"}
	if err := caller.Notify(ctx, MethodWindowLogMessage, params); err != nil {
		t.Fatalf("Notify(%s): %v", MethodWindowLogMessage, err)
	}
	select {
	case <-client.logMessageCalled:
	case <-ctx.Done():
		t.Fatalf("LogMessage was not called: %v", ctx.Err())
	}
	if diff := cmp.Diff(&params, client.snapshotLogMessageParams()); diff != "" {
		t.Fatalf("LogMessage params mismatch (-want +got):\n%s", diff)
	}
}

func TestClientHandlerMalformedParamsReturnsParseError(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	client := newRecordingClient()
	caller, served := newClientHandlerConnPair(ctx, t, client, jsonrpc2.MethodNotFoundHandler)
	defer closeJSONRPCConns(t, caller, served)

	_, err := caller.Call(ctx, MethodWindowLogMessage, jsonrpc2.RawMessage(`{"type":"bad"}`), nil)
	if !errors.Is(err, jsonrpc2.ErrParse) {
		t.Fatalf("Call malformed params error = %v, want wrapping %v", err, jsonrpc2.ErrParse)
	}
	if got := client.snapshotLogMessageParams(); got != nil {
		t.Fatalf("LogMessage was invoked after malformed params: %#v", got)
	}
}

func TestClientHandlerUnknownMethodFallsBack(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	client := newRecordingClient()
	fallbackCalled := make(chan string, 1)
	fallback := func(_ context.Context, req *jsonrpc2.Request) (any, error) {
		fallbackCalled <- req.Method()
		return jsonrpc2.RawMessage(`{"fallback":true}`), nil
	}
	caller, served := newClientHandlerConnPair(ctx, t, client, fallback)
	defer closeJSONRPCConns(t, caller, served)

	var got jsonrpc2.RawMessage
	const method = "workspace/customClientMethod"
	if _, err := caller.Call(ctx, method, jsonrpc2.RawMessage(`{"x":1}`), &got); err != nil {
		t.Fatalf("Call fallback: %v", err)
	}
	select {
	case gotMethod := <-fallbackCalled:
		if gotMethod != method {
			t.Fatalf("fallback method = %q, want %q", gotMethod, method)
		}
	case <-ctx.Done():
		t.Fatalf("fallback handler was not called: %v", ctx.Err())
	}
	if diff := cmp.Diff(jsonrpc2.RawMessage(`{"fallback":true}`), got); diff != "" {
		t.Fatalf("fallback result mismatch (-want +got):\n%s", diff)
	}
}

func TestClientHandlerFallbackCloneAllowsRetention(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	client := newRecordingClient()
	retained := make(chan *jsonrpc2.Request, 1)
	fallback := func(_ context.Context, req *jsonrpc2.Request) (any, error) {
		retained <- req.Clone()
		return jsonrpc2.RawMessage(`{"fallback":true}`), nil
	}
	caller, served := newClientHandlerConnPair(ctx, t, client, fallback)
	defer closeJSONRPCConns(t, caller, served)

	const firstMethod = "workspace/customClientMethod"
	firstParams := jsonrpc2.RawMessage(`{"marker":1,"nested":{"ok":true}}`)
	var got jsonrpc2.RawMessage
	if _, err := caller.Call(ctx, firstMethod, firstParams, &got); err != nil {
		t.Fatalf("Call first fallback: %v", err)
	}
	var cloned *jsonrpc2.Request
	select {
	case cloned = <-retained:
	case <-ctx.Done():
		t.Fatalf("fallback clone was not retained: %v", ctx.Err())
	}

	// Drive more traffic through the same connection after the fallback returns. If
	// the fallback had retained borrowed request data instead of Clone's owned
	// copy, this is the kind of reuse that can invalidate it.
	if _, err := caller.Call(ctx, "workspace/customClientMethodAfterClone", jsonrpc2.RawMessage(`{"marker":2}`), &got); err != nil {
		t.Fatalf("Call second fallback: %v", err)
	}

	if cloned.Method() != firstMethod {
		t.Fatalf("cloned fallback method = %q, want %q", cloned.Method(), firstMethod)
	}
	if diff := cmp.Diff(firstParams, cloned.Params()); diff != "" {
		t.Fatalf("cloned fallback params mismatch (-want +got):\n%s", diff)
	}
}

func TestClientHandlerCanceledContextShortCircuits(t *testing.T) {
	t.Parallel()

	client := newRecordingClient()
	fallbackCalled := false
	fallback := func(context.Context, *jsonrpc2.Request) (any, error) {
		fallbackCalled = true
		return nil, nil
	}
	handler := ClientHandler(client, fallback)
	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	_, err := handler(ctx, &jsonrpc2.Request{})
	if !errors.Is(err, ErrRequestCancelled) {
		t.Fatalf("ClientHandler canceled error = %v, want %v", err, ErrRequestCancelled)
	}
	if fallbackCalled {
		t.Fatal("fallback handler was called for canceled context")
	}
	if client.snapshotLogMessageParams() != nil || client.snapshotShowMessageRequestParams() != nil {
		t.Fatal("client handler was invoked for canceled context")
	}
}

func newClientHandlerConnPair(ctx context.Context, t *testing.T, client Client, fallback jsonrpc2.Handler) (caller, served jsonrpc2.Conn) {
	t.Helper()

	callerEnd, servedEnd := net.Pipe()
	caller = jsonrpc2.NewConn(jsonrpc2.NewStream(callerEnd), jsonrpc2.WithCodec(lspCodec{}))
	served = jsonrpc2.NewConn(jsonrpc2.NewStream(servedEnd), jsonrpc2.WithCodec(lspCodec{}))
	caller.Go(ctx, jsonrpc2.MethodNotFoundHandler)
	served.Go(ctx, ClientHandler(client, fallback))

	return caller, served
}

func closeJSONRPCConns(t *testing.T, conns ...jsonrpc2.Conn) {
	t.Helper()

	for _, conn := range conns {
		if err := conn.Close(); err != nil {
			t.Logf("close jsonrpc2 conn: %v", err)
		}
	}
	for _, conn := range conns {
		<-conn.Done()
	}
}
