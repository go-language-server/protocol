// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
)

// ReplyError replies error message.
func ReplyError(ctx context.Context, err error, req *jsonrpc2.Request) {
	var jrpcErr *jsonrpc2.Error
	if !errors.As(err, &jrpcErr) {
		err = jsonrpc2.Errorf(jsonrpc2.ParseError, "%v", err)
	}

	if err := req.Reply(ctx, nil, err); err != nil {
		LoggerFromContext(ctx).Error("ReplyError", zap.Error(err))
	}
}
