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

// NewClient returns the new Client, Server and jsonrpc2.Conn.
func NewClient(ctx context.Context, client ClientInterface, stream jsonrpc2.Stream, logger *zap.Logger) (*jsonrpc2.Conn, ServerInterface) {
	logger = logger.Named("jsonrpc2")
	opts := []jsonrpc2.Options{
		jsonrpc2.WithCanceler(jsonrpc2.Canceler(Canceller)),
		jsonrpc2.WithCapacity(DefaultBufferSize),
		jsonrpc2.WithOverloaded(true),
		jsonrpc2.WithLogger(logger.Named("client")),
	}
	conn := jsonrpc2.NewConn(ctx, stream, opts...)
	conn.Handler = ClientHandler(client, logger.Named("handler"))

	return conn, &Server{Conn: conn, logger: logger.Named("server")}
}

// NewServer returns the new Server, Client and jsonrpc2.Conn.
func NewServer(ctx context.Context, server ServerInterface, stream jsonrpc2.Stream, logger *zap.Logger) (*jsonrpc2.Conn, ClientInterface) {
	logger = logger.Named("jsonrpc2")
	opts := []jsonrpc2.Options{
		jsonrpc2.WithCanceler(jsonrpc2.Canceler(Canceller)),
		jsonrpc2.WithCapacity(DefaultBufferSize),
		jsonrpc2.WithOverloaded(true),
		jsonrpc2.WithLogger(logger.Named("server")),
	}
	conn := jsonrpc2.NewConn(ctx, stream, opts...)

	client := &Client{Conn: conn, logger: logger.Named("client")}
	conn.Handler = ServerHandler(server, logger.Named("handler"))

	return conn, client
}
