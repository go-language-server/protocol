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
func NewClient(ctx context.Context, conn *jsonrpc2.Conn, logger *zap.Logger) (ClientInterface, ServerInterface) {
	c := &Client{Conn: conn}

	return c, &Server{Conn: conn}
}

// NewServer returns the new Server, Client and jsonrpc2.Conn.
func NewServer(ctx context.Context, conn *jsonrpc2.Conn, logger *zap.Logger) (ServerInterface, ClientInterface) {
	s := &Server{Conn: conn}

	return s, &Client{Conn: conn}
}
