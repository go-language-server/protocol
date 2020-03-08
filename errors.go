// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"go.lsp.dev/jsonrpc2"
	"go.uber.org/zap"
)

// ReplyError replies error message.
func ReplyError(ctx context.Context, err error, req *jsonrpc2.Request) {
	if _, ok := err.(*jsonrpc2.Error); !ok {
		err = jsonrpc2.Errorf(jsonrpc2.ParseError, "%v", err)
	}

	if err := req.Reply(ctx, nil, err); err != nil {
		LoggerFromContext(ctx).Error("ReplyError", zap.Error(err))
	}
}
