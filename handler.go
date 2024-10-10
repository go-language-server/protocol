// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"fmt"

	"go.lsp.dev/jsonrpc2"
)

// CancelHandler handler of cancelling.
func CancelHandler(handler jsonrpc2.Handler) jsonrpc2.Handler {
	handler, canceller := jsonrpc2.CancelHandler(handler)

	h := func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		if req.Method() != "$/cancelRequest" {
			// TODO(iancottrell): See if we can generate a reply for the request to be cancelled
			// at the point of cancellation rather than waiting for gopls to naturally reply.
			// To do that, we need to keep track of whether a reply has been sent already and
			// be careful about racing between the two paths.
			// TODO(iancottrell): Add a test that watches the stream and verifies the response
			// for the cancelled request flows.
			reply := func(ctx context.Context, resp interface{}, err error) error {
				// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#cancelRequest
				if ctx.Err() != nil && err == nil {
					err = ErrRequestCancelled
				}
				ctx = context.WithoutCancel(ctx)

				return reply(ctx, resp, err)
			}

			return handler(ctx, reply, req)
		}

		var params CancelParams
		if err := unmarshal(req.Params(), &params); err != nil {
			return replyParseError(ctx, reply, err)
		}

		switch id := params.ID.value.(type) {
		case int32:
			canceller(jsonrpc2.NewNumberID(id))
		case string:
			canceller(jsonrpc2.NewStringID(id))
		default:
			return replyParseError(ctx, reply, fmt.Errorf("request ID %v malformed", id))
		}

		return reply(ctx, nil, nil)
	}

	return h
}

// Handlers default jsonrpc2.Handler.
func Handlers(handler jsonrpc2.Handler) jsonrpc2.Handler {
	return CancelHandler(
		jsonrpc2.AsyncHandler(
			jsonrpc2.ReplyHandler(handler),
		),
	)
}

// Call calls method to params and result.
func Call(ctx context.Context, conn jsonrpc2.Conn, method string, params, result interface{}) error {
	_, err := conn.Call(ctx, method, params, result)
	if ctx.Err() != nil {
	}

	return err
}

func notifyCancel(ctx context.Context, conn jsonrpc2.Conn, id jsonrpc2.ID) {
	ctx = context.WithoutCancel(ctx)
	// Note that only *jsonrpc2.ID implements json.Marshaler.
	cid := NewCancelParamsID(fmt.Sprint(id))
	conn.Notify(ctx, "$/cancelRequest", &CancelParams{ID: cid})
}

func replyParseError(ctx context.Context, reply jsonrpc2.Replier, err error) error {
	return reply(ctx, nil, fmt.Errorf("%w: %w", jsonrpc2.ErrParse, err))
}
