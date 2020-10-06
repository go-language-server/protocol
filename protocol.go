// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
)

// NewServer returns the new Server, Client and jsonrpc2.Conn.
func NewServer(pctx context.Context, server ServerInterface, stream jsonrpc2.Stream, logger *zap.Logger) (ctx context.Context, conn jsonrpc2.Conn, clientInterface ClientInterface) {
	conn = jsonrpc2.NewConn(stream)
	client := &client{
		Conn:   conn,
		logger: logger.Named("client"),
	}
	ctx = WithClient(pctx, client)

	// conn.AddHandler(&serverHandler{server: server})
	// conn.AddHandler(&canceller{logger: logger.Named("canceller")})

	return ctx, conn, client
}

// NewClient returns the new context, jsonrpc2.Conn and ServerInterface.
func NewClient(pctx context.Context, client ClientInterface, stream jsonrpc2.Stream, logger *zap.Logger) (ctx context.Context, conn jsonrpc2.Conn, serverInterface ServerInterface) {
	ctx = WithClient(pctx, client)

	conn = jsonrpc2.NewConn(stream)
	// conn.AddHandler(&clientHandler{client: client})
	// conn.AddHandler(&canceller{logger: logger.Named("canceller")})

	s := &server{
		Conn:   conn,
		logger: logger.Named("server"),
	}

	return ctx, conn, s
}
