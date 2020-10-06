// Copyright 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"encoding/json"
	"fmt"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/pkg/xcontext"
)

const RequestCancelledCode jsonrpc2.Code = -32800

// RequestCancelledError should be used when a request is canceled early.
var RequestCancelledError = jsonrpc2.NewError(RequestCancelledCode, "JSON RPC cancelled")

func Handlers(handler jsonrpc2.Handler) jsonrpc2.Handler {
	return CancelHandler(
		jsonrpc2.AsyncHandler(
			jsonrpc2.ReplyHandler(handler),
		),
	)
}

func CancelHandler(handler jsonrpc2.Handler) jsonrpc2.Handler {
	handler, canceller := jsonrpc2.CancelHandler(handler)

	h := func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Requester) error {
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
					err = RequestCancelledError
				}
				ctx = xcontext.Detach(ctx)
				return reply(ctx, resp, err)
			}
			return handler(ctx, reply, req)
		}

		var params CancelParams
		if err := json.Unmarshal(req.Params(), &params); err != nil {
			return sendParseError(ctx, reply, err)
		}

		switch id := params.ID.(type) {
		case float64:
			canceller(jsonrpc2.NewNumberID(int64(id)))
		case string:
			canceller(jsonrpc2.NewStringID(id))
		default:
			return sendParseError(ctx, reply, fmt.Errorf("request ID %v malformed", id))
		}

		return reply(ctx, nil, nil)
	}

	return h
}

func Call(ctx context.Context, conn jsonrpc2.Conn, method string, params, result interface{}) error {
	id, err := conn.Call(ctx, method, params, result)
	if ctx.Err() != nil {
		cancelCall(ctx, conn, id)
	}
	return err
}

func cancelCall(ctx context.Context, conn jsonrpc2.Conn, id jsonrpc2.ID) {
	ctx = xcontext.Detach(ctx)
	// Note that only *jsonrpc2.ID implements json.Marshaler.
	conn.Notify(ctx, "$/cancelRequest", &CancelParams{ID: &id})
}

func sendParseError(ctx context.Context, reply jsonrpc2.Replier, err error) error {
	return reply(ctx, nil, fmt.Errorf("%s: %w", jsonrpc2.ErrParse, err))
}
