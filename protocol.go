// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"
)

const (
	// Version is the version of the language-server-protocol specification being implemented.
	Version = "3.14.0"
)

const defaultMessageBufferSize = 20

func canceller(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	conn.Notify(context.Background(), "$/cancelRequest", &CancelParams{ID: *req.ID})
}

func NewServer(ctx context.Context, server Server, stream jsonrpc2.Stream, logger *zap.Logger) (*jsonrpc2.Conn, Client, *zap.Logger) {
	conn := jsonrpc2.NewConn(ctx, stream, jsonrpc2.WithHandler(serverHandler(server, logger)), jsonrpc2.WithCanceler(jsonrpc2.Canceler(canceller)))
	client := &clientDispatcher{Conn: conn}
	// conn.Capacity = defaultMessageBufferSize
	// conn.RejectIfOverloaded = true
	return conn, client, logger
}

func NewClient(ctx context.Context, client Client, stream jsonrpc2.Stream, logger *zap.Logger) (*jsonrpc2.Conn, Server, *zap.Logger) {
	conn := jsonrpc2.NewConn(ctx, stream, jsonrpc2.WithHandler(clientHandler(client, logger)), jsonrpc2.WithCanceler(jsonrpc2.Canceler(canceller)))
	// conn.Capacity = defaultMessageBufferSize
	// conn.RejectIfOverloaded = true
	return conn, &serverDispatcher{Conn: conn}, logger
}
