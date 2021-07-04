// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

// // Handlers default jsonrpc2.Handler.
// func Handlers(handler jsonrpc2.Handler) jsonrpc2.Handler {
// 	return CancelHandler(
// 		jsonrpc2.AsyncHandler(
// 			jsonrpc2.ReplyHandler(handler),
// 		),
// 	)
// }

// Request calls method to params and result.
// func Request(ctx context.Context, conn *jsonrpc2.Connection, method string, params interface{}) error {
// 	req, err := conn.Request(ctx, method, params)
// 	if ctx.Err() != nil {
// 		notifyCancel(ctx, conn, id)
// 	}
//
// 	var result interface{}
// 	err := req.Await(ctx, &result)
//
// 	return err
// }

// func notifyCancel(ctx context.Context, conn jsonrpc2.Conn, id jsonrpc2.ID) {
// 	ctx = xcontext.Detach(ctx)
// 	// Note that only *jsonrpc2.ID implements json.Marshaler.
// 	conn.Notify(ctx, MethodCancelRequest, &CancelParams{ID: &id})
// }
//
// func replyParseError(ctx context.Context, reply jsonrpc2.Replier, err error) error {
// 	return reply(ctx, nil, fmt.Errorf("%s: %w", jsonrpc2.ErrParse, err))
// }
