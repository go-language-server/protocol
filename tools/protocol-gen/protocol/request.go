// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Request represents a LSP request.
type Request struct {
	// Whether the request is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string

	// An optional documentation.
	Documentation string

	// An optional error data type.
	ErrorData Type

	// The direction in which this request is sent in the protocol.
	MessageDirection MessageDirection

	// The request's method name.
	Method string

	// The parameter type(s) if any.
	Params []Type

	// Optional partial result type if the request supports partial result reporting.
	PartialResult Type

	// Whether this is a proposed feature. If omitted the feature is final.
	Proposed bool

	// Optional a dynamic registration method if it different from the request's method.
	RegistrationMethod string

	// Optional registration options if the request supports dynamic registration.
	RegistrationOptions Type

	// The result type.
	Result Type

	// Since when (release number) this request is available. Is undefined if not known.
	Since string
}