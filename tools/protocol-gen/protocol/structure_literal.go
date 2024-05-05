// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// StructureLiteral defines an unnamed structure of an object literal.
type StructureLiteral struct {
	// Whether the literal is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string

	// An optional documentation.
	Documentation string

	// The properties.
	Properties []*Property

	// Whether this is a proposed structure. If omitted, the structure is final.
	Proposed bool

	// Since when (release number) this structure is available. Is undefined if not known.
	Since string
}
