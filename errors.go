// Copyright 2021 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import "go.lsp.dev/jsonrpc2"

const (
	// LSPReservedErrorRangeStart is the start range of LSP reserved error codes.
	//
	// It doesn't denote a real error code.
	//
	// @since 3.16.0.
	LSPReservedErrorRangeStart jsonrpc2.Code = -32899

	// ContentModified is the state change that invalidates the result of a request in execution.
	//
	// Defined by the protocol.
	CodeContentModified jsonrpc2.Code = -32801

	// RequestCancelled is the cancellation error.
	//
	// Defined by the protocol.
	CodeRequestCancelled jsonrpc2.Code = -32800

	// LSPReservedErrorRangeEnd is the end range of LSP reserved error codes.
	//
	// It doesn't denote a real error code.
	//
	// @since 3.16.0.
	LSPReservedErrorRangeEnd jsonrpc2.Code = -32800
)

var (
	// ErrContentModified should be used when a request is canceled early.
	ErrContentModified = jsonrpc2.NewError(CodeContentModified, "cancelled JSON-RPC")

	// ErrRequestCancelled should be used when a request is canceled early.
	ErrRequestCancelled = jsonrpc2.NewError(CodeRequestCancelled, "JSON RPC cancelled")
)
