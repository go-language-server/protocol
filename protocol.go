// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"log/slog"

	"go.lsp.dev/jsonrpc2"
)

// NewServer returns the context in which the [Client] dispatcher is embedded, the
// jsonrpc2 connection, and that [Client]. The connection serves the supplied
// [Server] and is wired with the union-aware [lspCodec].
func NewServer(ctx context.Context, server Server, stream jsonrpc2.Stream, logger *slog.Logger) (context.Context, jsonrpc2.Conn, Client) {
	if logger == nil {
		logger = nopLogger
	}
	conn := jsonrpc2.NewConn(stream, jsonrpc2.WithCodec(lspCodec{}))
	client := ClientDispatcher(conn, logger.With("name", "client"))
	ctx = WithClient(ctx, client)

	conn.Go(ctx, Handlers(ServerHandler(server, jsonrpc2.MethodNotFoundHandler)))

	return ctx, conn, client
}

// NewClient returns the context in which the [Client] is embedded, the jsonrpc2
// connection, and the [Server] dispatcher. The connection serves the supplied
// [Client] and is wired with the union-aware [lspCodec].
func NewClient(ctx context.Context, client Client, stream jsonrpc2.Stream, logger *slog.Logger) (context.Context, jsonrpc2.Conn, Server) {
	if logger == nil {
		logger = nopLogger
	}
	ctx = WithClient(ctx, client)

	conn := jsonrpc2.NewConn(stream, jsonrpc2.WithCodec(lspCodec{}))
	conn.Go(ctx, Handlers(ClientHandler(client, jsonrpc2.MethodNotFoundHandler)))
	server := ServerDispatcher(conn, logger.With("name", "server"))

	return ctx, conn, server
}
