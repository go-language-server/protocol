// Copyright 2020 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"bytes"
	"context"
	"encoding/json"

	"go.lsp.dev/jsonrpc2"
	"go.uber.org/zap"
)

// RequestCancelledError should be used when a request is canceled early.
const RequestCancelledError jsonrpc2.Code = -32800

type canceller struct {
	logger *zap.Logger
}

// compile time check whether the canceller implements jsonrpc2.Handler interface.
var _ jsonrpc2.Handler = (*canceller)(nil)

// Deliver implements Handler interface.
func (canceller) Deliver(ctx context.Context, r *jsonrpc2.Request, delivered bool) bool {
	// Hide cancellations from downstream handlers.
	return r.Method == MethodCancelRequest
}

// Cancel implements Handler interface.
func (canceller) Cancel(ctx context.Context, conn *jsonrpc2.Conn, id jsonrpc2.ID, canceled bool) bool {
	if canceled {
		return false
	}

	conn.Notify(ctx, MethodCancelRequest, &CancelParams{ID: id})
	return true
}

// Request implements Handler interface.
func (c *canceller) Request(ctx context.Context, conn *jsonrpc2.Conn, direction jsonrpc2.Direction, r *jsonrpc2.WireRequest) context.Context {
	if direction == jsonrpc2.Receive && r.Method == MethodCancelRequest {
		dec := json.NewDecoder(bytes.NewReader(*r.Params))
		var params CancelParams
		if err := dec.Decode(&params); err != nil {
			c.logger.Error("Request", zap.Error(err))
			return ctx
		}

		conn.Cancel(params.ID)
	}

	return ctx
}

// Response implements Handler interface.
func (canceller) Response(ctx context.Context, conn *jsonrpc2.Conn, direction jsonrpc2.Direction, r *jsonrpc2.WireResponse) context.Context {
	return ctx
}

// Done implements Handler interface.
func (canceller) Done(ctx context.Context, err error) {}

// Read implements Handler interface.
func (canceller) Read(ctx context.Context, n int64) context.Context { return ctx }

// Write implements Handler interface.
func (canceller) Write(ctx context.Context, n int64) context.Context { return ctx }

// Error implements Handler interface.
func (canceller) Error(ctx context.Context, err error) {}
