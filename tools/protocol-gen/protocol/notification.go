// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Notification represents a LSP notification.
type Notification struct {
	// The request's method name.
	Method string

	// The parameter type(s) if any.
	Params []Type

	// Optional a dynamic registration method if it different from the request's method.
	RegistrationMethod string

	// Optional registration options if the notification supports dynamic registration.
	RegistrationOptions Type

	// The direction in which this notification is sent in the protocol.
	MessageDirection MessageDirection

	// An optional documentation.
	Documentation string

	// Since when (release number) this notification is available. Is undefined if not knownz.
	Since string

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string

	// Whether this is a proposed notification. If omitted the notification is final.
	Proposed bool

	// Whether the notification is deprecated or not.
	// If deprecated the property contains the deprecation message.
	Deprecated string

	// The type name of the request if any.
	TypeName string
}
