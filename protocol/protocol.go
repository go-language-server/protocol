// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"encoding/json"
	"io"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
)

// MarshalFunc function type of marshal JSON data.
//
// Default is used [json.Marshal].
type MarshalFunc func(v any) ([]byte, error)

var marshal MarshalFunc = json.Marshal

func RegiserMarshaler(fn MarshalFunc) {
	marshal = fn
}

// UnmarshalFunc function type of unmarshal JSON data.
//
// Default is used [json.Unmarshal].
type UnmarshalFunc func(data []byte, v any) error

var unmarshal UnmarshalFunc = json.Unmarshal

func RegiserUnmarshaler(fn UnmarshalFunc) {
	unmarshal = fn
}

// JSONEncoder encodes and writes to the underlying data stream.
type JSONEncoder interface {
	Encode(any) error
}

// EncoderFunc function type of JSONEncoder.
//
// Default is used [json.NewEncoder] with SetEscapeHTML to false.
type EncoderFunc func(io.Writer) JSONEncoder

var newEncoder EncoderFunc = defaultEncoder

func defaultEncoder(w io.Writer) JSONEncoder {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc
}

func RegiserEncoder(fn EncoderFunc) {
	newEncoder = fn
}

// JSONDecoder decodes and reads to the underlying data stream.
type JSONDecoder interface {
	Decode(v any) error
}

// DecoderFunc function type of JSONDecoder.
//
// Default is used [json.NewDecoder].
type DecoderFunc func(io.Reader) JSONDecoder

var newDecoder DecoderFunc = defaultDecoder

func defaultDecoder(r io.Reader) JSONDecoder {
	dec := json.NewDecoder(r)
	return dec
}

func RegiserDecoder(fn DecoderFunc) {
	newDecoder = fn
}

// NewServer returns the context in which client is embedded, jsonrpc2.Conn, and the Client.
func NewServer(ctx context.Context, server Server, stream jsonrpc2.Stream, logger *zap.Logger) (context.Context, jsonrpc2.Conn, Client) {
	conn := jsonrpc2.NewConn(stream)
	cliint := ClientDispatcher(conn, logger.Named("client"))
	ctx = WithClient(ctx, cliint)

	conn.Go(ctx,
		Handlers(
			ServerHandler(server, jsonrpc2.MethodNotFoundHandler),
		),
	)

	return ctx, conn, cliint
}

// NewClient returns the context in which Client is embedded, jsonrpc2.Conn, and the Server.
func NewClient(ctx context.Context, client Client, stream jsonrpc2.Stream, logger *zap.Logger) (context.Context, jsonrpc2.Conn, Server) {
	ctx = WithClient(ctx, client)

	conn := jsonrpc2.NewConn(stream)
	conn.Go(ctx,
		Handlers(
			ClientHandler(client, jsonrpc2.MethodNotFoundHandler),
		),
	)
	server := ServerDispatcher(conn, logger.Named("server"))

	return ctx, conn, server
}
