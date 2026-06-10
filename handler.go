// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"fmt"
	"time"

	"go.lsp.dev/jsonrpc2"
)

// detachedContext wraps a parent [context.Context], preserving its values while
// shedding its cancellation and deadline. It is used to issue the
// reply/notification of a request whose own context has already been canceled,
// mirroring x/tools xcontext.Detach.
type detachedContext struct{ parent context.Context } //nolint:containedctx // intentional: detaches cancellation/deadline while preserving values, mirroring x/tools xcontext.Detach

func (detachedContext) Deadline() (time.Time, bool) { return time.Time{}, false }

func (detachedContext) Done() <-chan struct{} { return nil }

func (detachedContext) Err() error { return nil }

func (c detachedContext) Value(k any) any { return c.parent.Value(k) }

// detach returns a context that keeps ctx's values but drops its cancellation
// and deadline.
func detach(ctx context.Context) context.Context { return detachedContext{ctx} }

// CancelHandler returns a [jsonrpc2.Handler] that observes "$/cancelRequest"
// notifications and cancels the in-flight request they name. Replies are issued
// over a detached context so they are written even after the request's own
// context is canceled.
func CancelHandler(handler jsonrpc2.Handler) jsonrpc2.Handler {
	handler, canceller := jsonrpc2.CancelHandler(handler)

	h := func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		if req.Method() != MethodCancelRequest {
			// TODO(iancottrell): See if we can generate a reply for the request to be
			// cancelled at the point of cancellation rather than waiting for gopls to
			// naturally reply. To do that, we need to keep track of whether a reply has
			// been sent already and be careful about racing between the two paths.
			// The vendored jsonrpc2.CancelHandler cancels each request's context
			// at reply time as cleanup, then forwards that already-canceled
			// context here. A reply-time ctx.Err() check therefore cannot tell a
			// genuine $/cancelRequest from normal completion, so deriving
			// ErrRequestCancelled here would clobber every successful reply.
			// Genuine cancellations are surfaced by the dispatchers' pre-dispatch
			// ctx.Err() gate; here we only detach so the response is still written
			// after that self-cancel.
			wrapped := func(ctx context.Context, resp any, err error) error {
				return reply(detach(ctx), resp, err)
			}

			return handler(ctx, wrapped, req)
		}

		var params CancelParams
		if err := Unmarshal(req.Params(), &params); err != nil {
			return replyParseError(ctx, reply, err)
		}

		switch id := params.ID.(type) {
		case Integer:
			canceller(jsonrpc2.NewNumberID(int64(id)))
		case String:
			canceller(jsonrpc2.NewStringID(string(id)))
		default:
			return replyParseError(ctx, reply, fmt.Errorf("malformed cancel id %v", params.ID))
		}

		return reply(ctx, nil, nil)
	}

	return h
}

// Handlers wraps handler with the standard LSP middleware chain: cancellation,
// asynchronous dispatch, and reply accounting.
func Handlers(handler jsonrpc2.Handler) jsonrpc2.Handler {
	return CancelHandler(
		jsonrpc2.AsyncHandler(
			jsonrpc2.ReplyHandler(handler),
		),
	)
}

// Call invokes method on conn with params, decoding the response into result. If
// ctx is canceled while the call is outstanding, a "$/cancelRequest" notification
// is sent for the call's id.
func Call(ctx context.Context, conn jsonrpc2.Conn, method string, params, result any) error {
	id, err := conn.Call(ctx, method, params, result)
	if ctx.Err() != nil {
		notifyCancel(ctx, conn, id)
	}

	return err
}

// notifyCancel sends a "$/cancelRequest" notification for id over a detached
// context so the cancellation is delivered even though the caller's context is
// already done.
func notifyCancel(ctx context.Context, conn jsonrpc2.Conn, id jsonrpc2.ID) {
	ctx = detach(ctx)
	// The notification is best-effort: the request may already have completed.
	_ = conn.Notify(ctx, MethodCancelRequest, &CancelParams{ID: idToProgressToken(id)})
}

// idToProgressToken converts a jsonrpc2 request id into the [ProgressToken] union
// carried by [CancelParams].
func idToProgressToken(id jsonrpc2.ID) ProgressToken {
	if n, ok := id.Number(); ok {
		return Integer(n) //nolint:gosec // LSP request IDs are within the int32 range
	}
	s, _ := id.StringValue()

	return String(s)
}

// replyParseError replies with a parse error wrapping err.
func replyParseError(ctx context.Context, reply jsonrpc2.Replier, err error) error {
	return reply(ctx, nil, fmt.Errorf("%w: %w", jsonrpc2.ErrParse, err))
}
