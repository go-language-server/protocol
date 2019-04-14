// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"
)

const (
	// Version is the version of the language-server-protocol specification being implemented.
	Version = "3.14.0"
)

// NewClient returns the new Client, Server and jsonrpc2.Conn.
func NewClient(client ClientInterface, stream jsonrpc2.Stream, logger *zap.Logger, options ...jsonrpc2.Options) (*jsonrpc2.Conn, ServerInterface) {
	conn := jsonrpc2.NewConn(stream, options...)
	conn.Handler = ClientHandler(client, logger.Named("handler"))
	// conn.Logger = logger.Named("jsonrpc2")

	// serverConn := jsonrpc2.NewConn(stream, options...)
	s := &Server{Conn: conn, logger: logger.Named("server")}

	return conn, s
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
