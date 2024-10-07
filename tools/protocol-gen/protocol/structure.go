// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Structure defines the structure of an object literal.
type Structure struct {
	// The name of the structure.
	Name string

	// Structures extended from. This structures form a polymorphic type hierarchy.
	Extends []Type

	// Structures to mix in. The properties of these structures are `copied` into this structure. Mixins don't form a polymorphic type hierarchy in LSP.
	Mixins []Type

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

	// Whether the structure is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string

	// The 'kind' property, used to identify the structure type.
	Kind string

	// Child structures of this structure.
	NestedStructures []*Structure

	// The list of structure names (outermost first) to get to this structure.
	NestedNames []string
}

func (Structure) isTypeDecl() {}
