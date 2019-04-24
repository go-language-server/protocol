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

// DefaultBufferSize default message buffer size.
const DefaultBufferSize = 20

// Canceller returns the default canceler function.
func Canceller(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	conn.Notify(ctx, MethodCancelRequest, &CancelParams{ID: *req.ID})
}

// NewClient returns the new Client, Server and jsonrpc2.Conn.
func NewClient(ctx context.Context, client ClientInterface, stream jsonrpc2.Stream, logger *zap.Logger, options ...jsonrpc2.Options) (*jsonrpc2.Conn, ServerInterface) {
	clientLogger := logger.Named("client")
	serverLogger := logger.Named("server")

	conn := jsonrpc2.NewConn(stream, options...)
	conn.Handler = ClientHandler(ctx, client, clientLogger.Named("handler"))

	return conn, &Server{Conn: conn, logger: serverLogger}
}

// NewServer returns the new Server, Client and jsonrpc2.Conn.
func NewServer(server ServerInterface, stream jsonrpc2.Stream, logger *zap.Logger, options ...jsonrpc2.Options) (*jsonrpc2.Conn, ClientInterface) {
	conn := jsonrpc2.NewConn(stream, options...)
	conn.Handler = ServerHandler(server, logger.Named("handler"))
	conn.Logger = logger.Named("jsonrpc2")

	clientConn := jsonrpc2.NewConn(stream, options...)
	c := &Client{Conn: clientConn, logger: logger.Named("client")}
	clientConn.Handler = ClientHandler(c, c.Logger.Named("handler"))

	return conn, c
}
