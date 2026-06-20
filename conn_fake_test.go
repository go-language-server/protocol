// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"sync"

	"go.lsp.dev/jsonrpc2"
)

// fakeConn is a recording [jsonrpc2.Conn] used to drive the *client and *server
// dispatchers without real I/O. It captures the method and params of the most
// recent Call/Notify and replays a canned result (and error) into the caller's
// result pointer via the union-aware [lspCodec], so the test can assert that a
// sending method targeted the right method constant and decoded the right
// result type.
//
// It is concurrency-safe so that tests issuing reentrant calls (or running with
// -race) observe a consistent snapshot.
type fakeConn struct {
	mu sync.Mutex

	// recorded request.
	lastMethod string
	lastParams any
	callCount  int
	notifyCnt  int

	// canned response for the next Call: result is marshaled with lspCodec and
	// unmarshaled into the caller's result pointer; err is returned verbatim.
	result any
	err    error

	// id handed back from Call. Zero value is a number id 0.
	id jsonrpc2.ID
}

// compile-time assertion that *fakeConn satisfies jsonrpc2.Conn.
var _ jsonrpc2.Conn = (*fakeConn)(nil)

func (c *fakeConn) Call(_ context.Context, method string, params, result any) (jsonrpc2.ID, error) {
	c.mu.Lock()
	c.lastMethod = method
	c.lastParams = params
	c.callCount++
	canned := c.result
	err := c.err
	id := c.id
	c.mu.Unlock()

	if err != nil {
		return id, err
	}
	if result != nil && canned != nil {
		// Round-trip the canned result through the production codec so the
		// caller's typed result pointer is populated exactly as the real
		// transport would populate it.
		codec := lspCodec{}
		raw, mErr := codec.Marshal(canned)
		if mErr != nil {
			return id, mErr
		}
		if uErr := codec.Unmarshal(raw, result); uErr != nil {
			return id, uErr
		}
	}

	return id, nil
}

func (c *fakeConn) Notify(_ context.Context, method string, params any) error {
	c.mu.Lock()
	c.lastMethod = method
	c.lastParams = params
	c.notifyCnt++
	err := c.err
	c.mu.Unlock()

	return err
}

// Go, Close, Done, and Err satisfy the interface; the dispatchers under test
// never invoke them.
func (*fakeConn) Go(context.Context, jsonrpc2.Handler) {}

func (*fakeConn) Close() error { return nil }

func (c *fakeConn) Done() <-chan struct{} {
	ch := make(chan struct{})
	close(ch)

	return ch
}

func (*fakeConn) Err() error { return nil }

func (c *fakeConn) snapshot() (method string, params any, calls, notifies int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.lastMethod, c.lastParams, c.callCount, c.notifyCnt
}

func (c *fakeConn) setResult(result any, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.result = result
	c.err = err
}
