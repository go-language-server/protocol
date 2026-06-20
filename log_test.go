// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"go.lsp.dev/jsonrpc2"
)

// fakeStream is a recording jsonrpc2.Stream. Read replays a queued message and
// Write captures the message it was handed, so a loggingStream wrapped around it
// can be driven without a real transport.
type fakeStream struct {
	readMsg   jsonrpc2.Message
	readErr   error
	wroteMsg  jsonrpc2.Message
	writeErr  error
	closed    bool
	closeErr  error
	readCount int64
}

// compile-time assertion that *fakeStream satisfies jsonrpc2.Stream.
var _ jsonrpc2.Stream = (*fakeStream)(nil)

func (s *fakeStream) Read(context.Context) (jsonrpc2.Message, int64, error) {
	return s.readMsg, s.readCount, s.readErr
}

func (s *fakeStream) Write(_ context.Context, msg jsonrpc2.Message) (int64, error) {
	s.wroteMsg = msg
	if s.writeErr != nil {
		return 0, s.writeErr
	}

	return int64(0), nil
}

func (s *fakeStream) Close() error {
	s.closed = true

	return s.closeErr
}

// TestLoggingStreamReadLogsAndPropagates asserts Read forwards the underlying
// message/count/err and, on success, traces the message as a "Sending"
// (server-origin) record.
func TestLoggingStreamReadLogsAndPropagates(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	want := jsonrpc2.NewNotification("textDocument/didOpen", jsonrpc2.RawMessage(`{"x":1}`))
	inner := &fakeStream{readMsg: want, readCount: 42}
	ls := LoggingStream(inner, &buf)

	got, count, err := ls.Read(t.Context())
	if err != nil {
		t.Fatalf("Read err = %v, want nil", err)
	}
	if got != want {
		t.Errorf("Read msg = %v, want %v", got, want)
	}
	if count != 42 {
		t.Errorf("Read count = %d, want 42", count)
	}
	// Read is the server->client emit path, traced as "Sending".
	if out := buf.String(); !strings.Contains(out, "Sending notification 'textDocument/didOpen'") {
		t.Errorf("trace = %q, want Sending notification record", out)
	}
}

// TestLoggingStreamReadErrorSkipsLog asserts a failed Read forwards the error and
// does not emit a trace record.
func TestLoggingStreamReadErrorSkipsLog(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	wantErr := errors.New("boom")
	inner := &fakeStream{readErr: wantErr}
	ls := LoggingStream(inner, &buf)

	_, _, err := ls.Read(t.Context())
	if !errors.Is(err, wantErr) {
		t.Fatalf("Read err = %v, want %v", err, wantErr)
	}
	if buf.Len() != 0 {
		t.Errorf("trace = %q, want empty on read error", buf.String())
	}
}

// TestLoggingStreamWriteLogsAndPropagates asserts Write traces the message as a
// "Received" (client-origin) record, forwards it to the inner stream, and
// returns the inner write's result.
func TestLoggingStreamWriteLogsAndPropagates(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	msg := jsonrpc2.NewNotification("window/logMessage", jsonrpc2.RawMessage(`{"type":3}`))
	inner := &fakeStream{}
	ls := LoggingStream(inner, &buf)

	if _, err := ls.Write(t.Context(), msg); err != nil {
		t.Fatalf("Write err = %v, want nil", err)
	}
	if inner.wroteMsg != msg {
		t.Errorf("inner.wroteMsg = %v, want forwarded %v", inner.wroteMsg, msg)
	}
	// Write is the client->server receive path, traced as "Received".
	if out := buf.String(); !strings.Contains(out, "Received notification 'window/logMessage'") {
		t.Errorf("trace = %q, want Received notification record", out)
	}
}

// TestLoggingStreamWriteErrorPropagates asserts Write traces then returns the
// inner stream's write error.
func TestLoggingStreamWriteErrorPropagates(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	wantErr := errors.New("write failed")
	inner := &fakeStream{writeErr: wantErr}
	ls := LoggingStream(inner, &buf)

	_, err := ls.Write(t.Context(), jsonrpc2.NewNotification("$/progress", nil))
	if !errors.Is(err, wantErr) {
		t.Errorf("Write err = %v, want %v", err, wantErr)
	}
}

// TestLoggingStreamClose asserts Close forwards to the inner stream and returns
// its error.
func TestLoggingStreamClose(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		inner := &fakeStream{}
		ls := LoggingStream(inner, io.Discard)
		if err := ls.Close(); err != nil {
			t.Fatalf("Close err = %v, want nil", err)
		}
		if !inner.closed {
			t.Error("inner stream was not closed")
		}
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		wantErr := errors.New("close failed")
		inner := &fakeStream{closeErr: wantErr}
		ls := LoggingStream(inner, io.Discard)
		if err := ls.Close(); !errors.Is(err, wantErr) {
			t.Errorf("Close err = %v, want %v", err, wantErr)
		}
	})
}

// TestLogCommonCallRecordsAndResponseElapsed drives a request and its matching
// response through the same loggingStream, covering the request-timing
// bookkeeping: the Call records a start time and the Response looks it up to emit
// an elapsed-ms record with the recovered method name.
//
// A request and its response travel in opposite directions on the wire, so the
// Call is traced on the Read path (isRead=true → setClient writes clientCalls)
// and the Response on the Write path (isRead=false → get=client reads
// clientCalls). Pairing both on the same direction would consult the wrong map
// and recover a zero req (empty method, garbage elapsed).
func TestLogCommonCallRecordsAndResponseElapsed(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	inner := &fakeStream{}
	ls := LoggingStream(inner, &buf).(*loggingStream)

	id := jsonrpc2.NewNumberID(9001)
	call := jsonrpc2.NewCall(id, "textDocument/hover", jsonrpc2.RawMessage(`{"q":1}`))
	resp := jsonrpc2.NewResponse(id, jsonrpc2.RawMessage(`{"ok":true}`), nil)

	// Read path: trace the Call as "Sending request" and record it under
	// setClient (clientCalls).
	ls.logCommon(call, true)
	if out := buf.String(); !strings.Contains(out, "Sending request 'textDocument/hover - (9001)'") {
		t.Fatalf("call trace = %q, want Sending request record", out)
	}
	buf.Reset()

	// Write path: the response arrives in the opposite direction, traced as
	// "Received response" with the method recovered from clientCalls and an
	// elapsed-ms figure.
	ls.logCommon(resp, false)
	out := buf.String()
	if !strings.Contains(out, "Received response 'textDocument/hover - (9001)' in ") {
		t.Errorf("response trace = %q, want Received response with recovered method", out)
	}
	if !strings.Contains(out, "ms.") {
		t.Errorf("response trace = %q, want elapsed-ms figure", out)
	}
	if !strings.Contains(out, `Result: {"ok":true}`) {
		t.Errorf("response trace = %q, want Result line", out)
	}
}

// TestLogCommonClientRequestElapsed covers the mirror-direction timing pair: a
// client-originated request travels out on the Write path (isRead=false →
// setServer records serverCalls) and its response arrives on the Read path
// (isRead=true → get=server reads serverCalls). This exercises setServer and
// server, the counterparts of the setClient/client pair.
func TestLogCommonClientRequestElapsed(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	ls := LoggingStream(&fakeStream{}, &buf).(*loggingStream)

	id := jsonrpc2.NewNumberID(4242)
	call := jsonrpc2.NewCall(id, "workspace/symbol", jsonrpc2.RawMessage(`{"query":"x"}`))
	resp := jsonrpc2.NewResponse(id, jsonrpc2.RawMessage(`[]`), nil)

	// Write path: trace the Call as "Received request" and record it under
	// setServer (serverCalls).
	ls.logCommon(call, false)
	if out := buf.String(); !strings.Contains(out, "Received request 'workspace/symbol - (4242)'") {
		t.Fatalf("call trace = %q, want Received request record", out)
	}
	buf.Reset()

	// Read path: the response is emitted in the opposite direction, traced as
	// "Sending response" with the method recovered from serverCalls.
	ls.logCommon(resp, true)
	out := buf.String()
	if !strings.Contains(out, "Sending response 'workspace/symbol - (4242)' in ") {
		t.Errorf("response trace = %q, want Sending response with recovered method", out)
	}
	if !strings.Contains(out, `Result: []`) {
		t.Errorf("response trace = %q, want Result line", out)
	}
}

// TestLogCommonResponseError covers the Response error branch, which writes a
// "[Error - ...]" line directly to the log and returns early.
func TestLogCommonResponseError(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	inner := &fakeStream{}
	ls := LoggingStream(inner, &buf).(*loggingStream)

	id := jsonrpc2.NewNumberID(7)
	resp := jsonrpc2.NewResponse(id, nil, jsonrpc2.NewError(jsonrpc2.ErrMethodNotFound.Code, "nope"))

	// Write path: pastTense is "Received".
	ls.logCommon(resp, false)
	out := buf.String()
	if !strings.Contains(out, "[Error - Received]") {
		t.Errorf("trace = %q, want [Error - Received] line", out)
	}
	if !strings.Contains(out, "#7") {
		t.Errorf("trace = %q, want id #7", out)
	}
	if !strings.Contains(out, "nope") {
		t.Errorf("trace = %q, want error message", out)
	}
}

// TestLogCommonNilMessageNoop asserts a nil message produces no output and does
// not panic.
func TestLogCommonNilMessageNoop(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	ls := LoggingStream(&fakeStream{}, &buf).(*loggingStream)
	ls.logCommon(nil, true)
	if buf.Len() != 0 {
		t.Errorf("trace = %q, want empty for nil message", buf.String())
	}
}

// TestLogCommonNilLogNoop asserts a nil log writer short-circuits before any
// formatting and does not panic.
func TestLogCommonNilLogNoop(t *testing.T) {
	t.Parallel()

	ls := LoggingStream(&fakeStream{}, nil).(*loggingStream)
	// Must not panic despite a non-nil message and nil writer.
	ls.logCommon(jsonrpc2.NewNotification("$/cancelRequest", nil), false)
}

// TestLogCommonNotificationReceived covers the Write-path notification branch
// (direction "Received") to complement the Read-path coverage above.
func TestLogCommonNotificationReceived(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	ls := LoggingStream(&fakeStream{}, &buf).(*loggingStream)
	ls.logCommon(jsonrpc2.NewNotification("$/setTrace", jsonrpc2.RawMessage(`{"value":"verbose"}`)), false)
	out := buf.String()
	if !strings.Contains(out, "Received notification '$/setTrace'") {
		t.Errorf("trace = %q, want Received notification record", out)
	}
	if !strings.Contains(out, `Params: {"value":"verbose"}`) {
		t.Errorf("trace = %q, want Params line", out)
	}
}
