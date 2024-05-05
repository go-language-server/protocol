// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"fmt"

	"go.lsp.dev/jsonrpc2"
)

// Handlers default jsonrpc2.Handler.
func Handlers(handler jsonrpc2.Handler) jsonrpc2.Handler {
	return jsonrpc2.AsyncHandler(
		jsonrpc2.ReplyHandler(handler),
	)
}

// Call calls method to params and result.
func Call(ctx context.Context, conn jsonrpc2.Conn, method string, params, result interface{}) error {
	_, err := conn.Call(ctx, method, params, result)
	if ctx.Err() != nil {
	}

	return err
}

func replyParseError(ctx context.Context, reply jsonrpc2.Replier, err error) error {
	return reply(ctx, nil, fmt.Errorf("%w: %w", jsonrpc2.ErrParse, err))
}
