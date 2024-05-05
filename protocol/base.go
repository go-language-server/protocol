// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

// ErrorCodes predefined error codes.
type ErrorCodes int32

const (
	ParseErrorErrorCodes ErrorCodes = -32700

	InvalidRequestErrorCodes ErrorCodes = -32600

	MethodNotFoundErrorCodes ErrorCodes = -32601

	InvalidParamsErrorCodes ErrorCodes = -32602

	InternalErrorErrorCodes ErrorCodes = -32603

	// ServerNotInitializedErrorCodes error code indicating that a server received a notification or request before the server has received the `initialize` request.
	ServerNotInitializedErrorCodes ErrorCodes = -32002

	UnknownErrorCodeErrorCodes ErrorCodes = -32001
)

type LSPErrorCodes int32

const (
	// RequestFailedLSPErrorCodes a request failed but it was syntactically correct, e.g the method name was known and the parameters were valid. The error message should contain human readable information about why the request failed.
	//
	// @since 3.17.0
	RequestFailedLSPErrorCodes LSPErrorCodes = -32803

	// ServerCancelledLSPErrorCodes the server cancelled the request. This error code should only be used for requests that explicitly support being server cancellable.
	//
	// @since 3.17.0
	ServerCancelledLSPErrorCodes LSPErrorCodes = -32802

	// ContentModifiedLSPErrorCodes the server detected that the content of a document got modified outside normal conditions. A server should NOT send this error code if it detects a content change in it unprocessed messages. The result even computed on an older state might still be useful for the client. If a client decides that a result is not of any use anymore the client should cancel the request.
	ContentModifiedLSPErrorCodes LSPErrorCodes = -32801

	// RequestCancelledLSPErrorCodes the client has canceled a request and a server as detected the cancel.
	RequestCancelledLSPErrorCodes LSPErrorCodes = -32800
)

type CancelParams struct {
	// ID the request id to cancel.
	ID CancelParamsID `json:"id"`
}

type ProgressParams struct {
	// Token the progress token provided by the client or server.
	Token ProgressToken `json:"token"`

	// Value the progress data.
	Value any `json:"value"`
}
