// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"
)

// NewServer returns the new Server, Client and jsonrpc2.Conn.
func NewServer(pctx context.Context, server ServerInterface, stream jsonrpc2.Stream, logger *zap.Logger, options ...jsonrpc2.Options) (ctx context.Context, conn *jsonrpc2.Conn, clientInterface ClientInterface) {
	conn = jsonrpc2.NewConn(stream, options...)
	client := &client{
		Conn:   conn,
		logger: logger.Named("client"),
	}
	ctx = WithClient(pctx, client)

	conn.AddHandler(&serverHandler{server: server})
	conn.AddHandler(&canceller{logger: logger.Named("canceller")})

	return ctx, conn, client
}

// NewClient returns the new context, jsonrpc2.Conn and ServerInterface.
func NewClient(pctx context.Context, client ClientInterface, stream jsonrpc2.Stream, logger *zap.Logger, options ...jsonrpc2.Options) (ctx context.Context, conn *jsonrpc2.Conn, serverInterface ServerInterface) {
	ctx = WithClient(pctx, client)

	conn = jsonrpc2.NewConn(stream, options...)
	conn.AddHandler(&clientHandler{client: client})
	conn.AddHandler(&canceller{logger: logger.Named("canceller")})

	s := &server{
		Conn:   conn,
		logger: logger.Named("server"),
	}

	return ctx, conn, s
}
