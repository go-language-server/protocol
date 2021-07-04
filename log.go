// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"go.lsp.dev/jsonrpc2"
)

type writeLock struct {
	sync.Mutex
	w io.Writer
}

// loggingReader represents a logging of jsonrpc2.Stream.
type loggingReader struct {
	reader    jsonrpc2.Reader
	writeLock *writeLock
}

// LoggingReader returns a reader that does LSP protocol logging.
func LoggingReader(reader jsonrpc2.Reader, w io.Writer) jsonrpc2.Reader {
	return &loggingReader{
		reader:    reader,
		writeLock: &writeLock{w: w},
	}
}

// Read implements jsonrpc2.Reader.Read.
func (s *loggingReader) Read(ctx context.Context) (jsonrpc2.Message, int64, error) {
	msg, n, err := s.reader.Read(ctx)
	if err == nil {
		log(msg, s.writeLock, true)
	}

	return msg, n, err
}

// loggingReader represents a logging of jsonrpc2.Stream.
type loggingWriter struct {
	writer    jsonrpc2.Writer
	writeLock *writeLock
}

// LoggingWriter returns a writer that does LSP protocol logging.
func LoggingWriter(writer jsonrpc2.Writer, w io.Writer) jsonrpc2.Writer {
	return &loggingWriter{
		writer:    writer,
		writeLock: &writeLock{w: w},
	}
}

// Write implements jsonrpc2.Writer.Write.
func (s *loggingWriter) Write(ctx context.Context, msg jsonrpc2.Message) (int64, error) {
	log(msg, s.writeLock, false)
	n, err := s.writer.Write(ctx, msg)

	return n, err
}

type req struct {
	method string
	start  time.Time
}

type mapped struct {
	mu          sync.Mutex
	clientCalls map[string]req
	serverCalls map[string]req
}

var maps = &mapped{
	mu:          sync.Mutex{},
	clientCalls: make(map[string]req),
	serverCalls: make(map[string]req),
}

// these 4 methods are each used exactly once, but it seemed
// better to have the encapsulation rather than ad hoc mutex
// code in 4 places.
func (m *mapped) client(id string) req {
	m.mu.Lock()
	v := m.clientCalls[id]
	delete(m.clientCalls, id)
	m.mu.Unlock()

	return v
}

func (m *mapped) server(id string) req {
	m.mu.Lock()
	v := m.serverCalls[id]
	delete(m.serverCalls, id)
	m.mu.Unlock()

	return v
}

func (m *mapped) setClient(id string, r req) {
	m.mu.Lock()
	m.clientCalls[id] = r
	m.mu.Unlock()
}

func (m *mapped) setServer(id string, r req) {
	m.mu.Lock()
	m.serverCalls[id] = r
	m.mu.Unlock()
}

const eor = "\r\n\r\n\r\n"

func log(msg jsonrpc2.Message, wl *writeLock, isRead bool) {
	if msg == nil || wl.w == nil {
		return
	}

	wl.Lock()

	direction, pastTense := "Received", "Received"
	get, set := maps.client, maps.setServer
	if isRead {
		direction, pastTense = "Sending", "Sent"
		get, set = maps.server, maps.setClient
	}

	tm := time.Now()
	tmfmt := tm.Format("15:04:05.000 PM")

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "[Trace - %s] ", tmfmt) // common beginning

	switch msg := msg.(type) {
	case *jsonrpc2.Request:
		if msg.IsCall() {
			id := fmt.Sprint(msg.ID)
			fmt.Fprintf(&buf, "%s request '%s - (%s)'.\n", direction, msg.Method, id)
			fmt.Fprintf(&buf, "Params: %s%s", msg.Params, eor)
			set(id, req{method: msg.Method, start: tm})
		} else {
			fmt.Fprintf(&buf, "%s notification '%s'.\n", direction, msg.Method)
			fmt.Fprintf(&buf, "Params: %s%s", msg.Params, eor)
		}

	case *jsonrpc2.Response:
		id := fmt.Sprint(msg.ID)
		if err := msg.Error; err != nil {
			fmt.Fprintf(wl.w, "[Error - %s] %s #%s %s%s", pastTense, tmfmt, id, err, eor)

			return
		}

		cc := get(id)
		elapsed := tm.Sub(cc.start)
		fmt.Fprintf(&buf, "%s response '%s - (%s)' in %dms.\n",
			direction, cc.method, id, elapsed/time.Millisecond)
		fmt.Fprintf(&buf, "Result: %s%s", msg.Result, eor)
	}

	wl.w.Write(buf.Bytes())

	wl.Unlock()
}
