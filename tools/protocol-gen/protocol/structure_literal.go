// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// StructureLiteral defines an unnamed structure of an object literal.
type StructureLiteral struct {
	// The properties.
	Properties []*Property

	// An optional documentation.
	Documentation string

	// Since when (release number) this structure is available. Is undefined if not known.
	Since string

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string

	// Whether this is a proposed structure. If omitted, the structure is final.
	Proposed bool

	// Whether the literal is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string
}
