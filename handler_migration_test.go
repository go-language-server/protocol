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

	gocmp "github.com/google/go-cmp/cmp"

	"go.lsp.dev/jsonrpc2"
)

func newJSONRPCTestStreams(t *testing.T) (jsonrpc2.Conn, jsonrpc2.Stream) {
	t.Helper()

	serverEnd, clientEnd := net.Pipe()
	serverConn := jsonrpc2.NewConn(jsonrpc2.NewStream(serverEnd), jsonrpc2.WithCodec(lspCodec{}))
	clientStream := jsonrpc2.NewStream(clientEnd)
	t.Cleanup(func() {
		if err := clientStream.Close(); err != nil {
			t.Logf("close client stream: %v", err)
		}
		if err := serverConn.Close(); err != nil {
			t.Logf("close server conn: %v", err)
		}
	})

	return serverConn, clientStream
}

func TestCancelHandlerCancelsInFlightRequests(t *testing.T) {
	tests := map[string]struct {
		id           jsonrpc2.ID
		cancelParams jsonrpc2.RawMessage
	}{
		"success: numeric id": {
			id:           jsonrpc2.NewNumberID(7),
			cancelParams: jsonrpc2.RawMessage(`{"id":7}`),
		},
		"success: string id": {
			id:           jsonrpc2.NewStringID("request-7"),
			cancelParams: jsonrpc2.RawMessage(`{"id":"request-7"}`),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
			defer cancel()

			serverConn, clientStream := newJSONRPCTestStreams(t)

			started := make(chan struct{})
			canceled := make(chan struct{})
			serverConn.Go(ctx, Handlers(func(ctx context.Context, req *jsonrpc2.Request) (any, error) {
				if req.Method() != "test/block" {
					return nil, jsonrpc2.ErrMethodNotFound
				}
				close(started)
				select {
				case <-ctx.Done():
					close(canceled)
					return nil, ErrRequestCancelled
				case <-time.After(2 * time.Second):
					return nil, errors.New("request was not canceled")
				}
			}))

			if _, err := clientStream.Write(ctx, jsonrpc2.NewCall(tt.id, "test/block", nil)); err != nil {
				t.Fatalf("write blocking call: %v", err)
			}
			select {
			case <-started:
			case <-ctx.Done():
				t.Fatalf("blocking call did not start: %v", ctx.Err())
			}

			if _, err := clientStream.Write(ctx, jsonrpc2.NewNotification(MethodCancelRequest, tt.cancelParams)); err != nil {
				t.Fatalf("write cancel notification: %v", err)
			}
			select {
			case <-canceled:
			case <-ctx.Done():
				t.Fatalf("request was not canceled: %v", ctx.Err())
			}

			msg, _, err := clientStream.Read(ctx)
			if err != nil {
				t.Fatalf("read canceled response: %v", err)
			}
			resp, ok := msg.(*jsonrpc2.Response)
			if !ok {
				t.Fatalf("canceled response type = %T, want *jsonrpc2.Response", msg)
			}
			if resp.ID() != tt.id {
				t.Fatalf("canceled response id = %v, want %v", resp.ID(), tt.id)
			}
			if err := resp.Err(); !errors.Is(err, ErrRequestCancelled) {
				t.Fatalf("canceled response error = %v, want wrapping %v", err, ErrRequestCancelled)
			}
		})
	}
}

func TestCancelHandlerMalformedCancelIDCallReturnsParseError(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	serverConn, clientStream := newJSONRPCTestStreams(t)
	serverConn.Go(ctx, Handlers(jsonrpc2.MethodNotFoundHandler))

	const callID int64 = 99
	if _, err := clientStream.Write(ctx, jsonrpc2.NewCall(
		jsonrpc2.NewNumberID(callID),
		MethodCancelRequest,
		jsonrpc2.RawMessage(`{"id":{}}`),
	)); err != nil {
		t.Fatalf("write malformed cancel call: %v", err)
	}

	msg, _, err := clientStream.Read(ctx)
	if err != nil {
		t.Fatalf("read malformed cancel response: %v", err)
	}
	resp, ok := msg.(*jsonrpc2.Response)
	if !ok {
		t.Fatalf("malformed cancel response type = %T, want *jsonrpc2.Response", msg)
	}
	if resp.ID() != jsonrpc2.NewNumberID(callID) {
		t.Fatalf("malformed cancel response id = %v, want %v", resp.ID(), jsonrpc2.NewNumberID(callID))
	}
	if err := resp.Err(); !errors.Is(err, jsonrpc2.ErrParse) {
		t.Fatalf("malformed cancel response error = %v, want wrapping %v", err, jsonrpc2.ErrParse)
	}
}

func TestCancelHandlerMalformedCancelIDNotificationFailsConnection(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	serverConn, clientStream := newJSONRPCTestStreams(t)
	serverConn.Go(ctx, Handlers(jsonrpc2.MethodNotFoundHandler))

	if _, err := clientStream.Write(ctx, jsonrpc2.NewNotification(
		MethodCancelRequest,
		jsonrpc2.RawMessage(`{"id":{}}`),
	)); err != nil {
		t.Fatalf("write malformed cancel notification: %v", err)
	}

	select {
	case <-serverConn.Done():
	case <-ctx.Done():
		t.Fatalf("server connection did not fail for malformed cancel notification: %v", ctx.Err())
	}
	if err := serverConn.Err(); !errors.Is(err, jsonrpc2.ErrParse) {
		t.Fatalf("server connection error = %v, want wrapping %v", err, jsonrpc2.ErrParse)
	}
}

type fallbackRecord struct {
	method string
	params LSPAny
}

type fallbackRecordingServer struct {
	UnimplementedServer

	mu           sync.Mutex
	records      []fallbackRecord
	standardHits int
}

func (s *fallbackRecordingServer) Request(_ context.Context, method string, params any) (any, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	record := fallbackRecord{method: method}
	if p, ok := params.(LSPAny); ok {
		// Store the LSPAny as received from production ServerHandler. This avoids
		// hiding aliasing problems behind a test-double copy.
		record.params = p
	}
	s.records = append(s.records, record)

	return nil, nil //nolint:nilnil // Server.Request contract: (nil, nil) is the valid "no result, no error" fallback return
}

func (s *fallbackRecordingServer) DidOpen(_ context.Context, _ *DidOpenTextDocumentParams) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.standardHits++

	return nil
}

func TestServerHandlerCustomFallbackRetainsOwnedMethodAndParams(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	serverEnd, clientEnd := net.Pipe()
	server := &fallbackRecordingServer{}
	_, serverConn, _ := NewServer(ctx, server, jsonrpc2.NewStream(serverEnd))
	clientStream := jsonrpc2.NewStream(clientEnd)
	t.Cleanup(func() {
		if err := clientStream.Close(); err != nil {
			t.Logf("close client stream: %v", err)
		}
		if err := serverConn.Close(); err != nil {
			t.Logf("close server conn: %v", err)
		}
	})

	firstMethod := "workspace/x-custom-first"
	firstParams := jsonrpc2.RawMessage(`{"marker":1,"nested":{"ok":true}}`)
	wantFirstParams := append(LSPAny(nil), firstParams...)
	if _, err := clientStream.Write(ctx, jsonrpc2.NewCall(jsonrpc2.NewNumberID(1), firstMethod, firstParams)); err != nil {
		t.Fatalf("write first custom call: %v", err)
	}
	if _, _, err := clientStream.Read(ctx); err != nil {
		t.Fatalf("read first custom response: %v", err)
	}

	// The custom fallback may retain method/params beyond handler return. Mutate
	// the caller-owned request bytes and then process more traffic to pin that
	// ServerHandler cloned/unmarshaled borrowed jsonrpc2 data before user handoff.
	for i := range firstParams {
		firstParams[i] = 'x'
	}

	secondMethod := "workspace/x-custom-second-after-reuse"
	secondParams := jsonrpc2.RawMessage(`{"marker":2,"nested":{"ok":false},"items":[1,2,3]}`)
	wantSecondParams := append(LSPAny(nil), secondParams...)
	if _, err := clientStream.Write(ctx, jsonrpc2.NewCall(jsonrpc2.NewNumberID(2), secondMethod, secondParams)); err != nil {
		t.Fatalf("write second custom call: %v", err)
	}
	if _, _, err := clientStream.Read(ctx); err != nil {
		t.Fatalf("read second custom response: %v", err)
	}

	standardParams, err := Marshal(&DidOpenTextDocumentParams{
		TextDocument: TextDocumentItem{
			URI:        "file:///fallback.go",
			LanguageID: "go",
			Version:    1,
			Text:       "package fallback",
		},
	})
	if err != nil {
		t.Fatalf("marshal didOpen params: %v", err)
	}
	if _, err := clientStream.Write(ctx, jsonrpc2.NewNotification(MethodTextDocumentDidOpen, standardParams)); err != nil {
		t.Fatalf("write standard notification: %v", err)
	}

	deadline := time.After(2 * time.Second)
	for {
		server.mu.Lock()
		records := make([]fallbackRecord, len(server.records))
		copy(records, server.records)
		standardHits := server.standardHits
		server.mu.Unlock()

		if len(records) == 2 && standardHits == 1 {
			wantRecords := []fallbackRecord{
				{method: firstMethod, params: wantFirstParams},
				{method: secondMethod, params: wantSecondParams},
			}
			if diff := gocmp.Diff(wantRecords, records, gocmp.AllowUnexported(fallbackRecord{})); diff != "" {
				t.Fatalf("fallback records mismatch (-want +got):\n%s", diff)
			}

			return
		}

		select {
		case <-deadline:
			t.Fatalf("server hits = fallback:%d standard:%d, want fallback:2 standard:1", len(records), standardHits)
		case <-time.After(10 * time.Millisecond):
		}
	}
}
