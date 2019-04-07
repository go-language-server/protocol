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

// DefaultCanceller returns the default canceler function.
func DefaultCanceller(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	conn.Notify(context.Background(), cancelRequest, &CancelParams{ID: *req.ID})
}

// NewServer returns the new jsonrpc2.Conn for Server and Client.
func NewServer(ctx context.Context, server Server, stream jsonrpc2.Stream, logger *zap.Logger, opts ...jsonrpc2.Options) (*jsonrpc2.Conn, Client) {
	opts = append(opts, jsonrpc2.WithHandler(serverHandler(server, logger)))

	conn := jsonrpc2.NewConn(ctx, stream, opts...)
	c := &client{Conn: conn}

	return conn, c
}

// NewClient returns the new jsonrpc2.Conn for Client and Server.
func NewClient(ctx context.Context, client Client, stream jsonrpc2.Stream, logger *zap.Logger, opts ...jsonrpc2.Options) (*jsonrpc2.Conn, Server) {
	opts = append(opts, jsonrpc2.WithHandler(clientHandler(client, logger)))

	conn := jsonrpc2.NewConn(ctx, stream, opts...)

	return conn, &server{Conn: conn}
}
