// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors

package protocol

import (
	"context"
	"fmt"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/pkg/event"
	"go.lsp.dev/pkg/xcontext"
)

// Handlers default jsonrpc2.Handler.
func Handlers(handler jsonrpc2.Handler) jsonrpc2.Handler {
	return CancelHandler(
		jsonrpc2.AsyncHandler(
			jsonrpc2.ReplyHandler(handler),
		),
	)
}

// Call calls method to params and result.
func Call(ctx context.Context, conn jsonrpc2.Connection, method string, params, result interface{}) error {
	id, err := conn.Call(ctx, method, params, result)
	if ctx.Err() != nil {
		cancelCall(ctx, clientConn{conn}, id)
	}
	return err
}

func notifyCancel(ctx context.Context, sender connSender, id jsonrpc2.ID) {
	ctx = xcontext.Detach(ctx)
	ctx, done := event.Start(ctx, "protocol.canceller")
	defer done()
	// Note that only *jsonrpc2.ID implements json.Marshaler.
	sender.Notify(ctx, MethodCancelRequest, &CancelParams{ID: &id})
}

func replyParseError(ctx context.Context, reply jsonrpc2.Replier, err error) error {
	return reply(ctx, nil, fmt.Errorf("%s: %w", jsonrpc2.ErrParse, err))
}
