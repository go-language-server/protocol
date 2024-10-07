// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Property represents an object property.
type Property struct {
	// The property name.
	Name string

	// The type of the property.
	Type Type

	// Whether the property is optional. If omitted, the property is mandatory.
	Optional bool

	// An optional documentation.
	Documentation string

	// Since when (release number) this property is available. Is undefined if not known.
	Since string

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string

	// Whether this is a proposed property. If omitted, the structure is final.
	Proposed bool

	// Whether the property is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string

	// The property JSON name.
	JSONName string
}
