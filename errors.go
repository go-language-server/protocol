// Copyright 2021 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import "go.lsp.dev/jsonrpc2"

const (
	// ContentModified is the state change that invalidates the result of a request in execution.
	//
	// Defined by the protocol.
	CodeContentModified jsonrpc2.Code = -32801

	// RequestCancelled is the cancellation error.
	//
	// Defined by the protocol.
	CodeRequestCancelled jsonrpc2.Code = -32800
)

var (
	// ErrContentModified should be used when a request is canceled early.
	ErrContentModified = jsonrpc2.NewError(CodeContentModified, "cancelled JSON-RPC")

	// ErrRequestCancelled should be used when a request is canceled early.
	ErrRequestCancelled = jsonrpc2.NewError(CodeRequestCancelled, "cancelled JSON-RPC")
)
