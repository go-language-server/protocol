// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Notification represents a LSP notification.
type Notification struct {
	// Whether the notification is deprecated or not.
	// If deprecated the property contains the deprecation message.
	Deprecated string

	// An optional documentation.
	Documentation string

	// The direction in which this notification is sent in the protocol.
	MessageDirection MessageDirection

	// The request's method name.
	Method string

	// The parameter type(s) if any.
	Params []Type

	// Whether this is a proposed notification. If omitted the notification is final.
	Proposed bool

	// Optional a dynamic registration method if it different from the request's method.
	RegistrationMethod string

	// Optional registration options if the notification supports dynamic registration.
	RegistrationOptions Type

	// Since when (release number) this notification is available. Is undefined if not knownz.
	Since string
}
