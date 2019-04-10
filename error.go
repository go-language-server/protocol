// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"
)

// ErrorInvalidParams reports InvalidParams error.
func ErrorInvalidParams(format string, args ...interface{}) error {
	return jsonrpc2.Errorf(jsonrpc2.CodeInvalidParams, format, args...)
}

// ReplyError reply errors.
func ReplyError(ctx context.Context, err error, conn *jsonrpc2.Conn, req *jsonrpc2.Request, logger *zap.Logger) {
	if _, ok := err.(*jsonrpc2.Error); !ok {
		err = jsonrpc2.Errorf(jsonrpc2.CodeParseError, "%v", err)
	}

	if err := conn.Reply(ctx, req, nil, err); err != nil {
		logger.Error("sendParseError", zap.Error(err))
	}
}
