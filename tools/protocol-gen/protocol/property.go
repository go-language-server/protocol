// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Property represents an object property.
type Property struct {
	// Whether the property is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string

	// An optional documentation.
	Documentation string

	// The property JSON name.
	JSONName string

	// The property name.
	Name string

	// Whether the property is optional. If omitted, the property is mandatory.
	Optional bool

	// Whether this is a proposed property. If omitted, the structure is final.
	Proposed bool

	// Since when (release number) this property is available. Is undefined if not known.
	Since string

	// The type of the property.
	Type Type
}
