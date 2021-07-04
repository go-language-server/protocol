// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"go.lsp.dev/jsonrpc2"
)

const (
	// ContentModified is the state change that invalidates the result of a request in execution.
	//
	// Defined by the protocol.
	CodeContentModified = jsonrpc2.Code(-32801)

	// RequestCancelled is the cancellation error.
	//
	// Defined by the protocol.
	CodeRequestCancelled = jsonrpc2.Code(-32800)
)

var (
	// ErrContentModified should be used when a request is canceled early.
	ErrContentModified = jsonrpc2.NewError(CodeContentModified, "JSON RPC modified")

	// ErrRequestCancelled should be used when a request is canceled early.
	ErrRequestCancelled = jsonrpc2.NewError(CodeRequestCancelled, "JSON RPC cancelled")
)
