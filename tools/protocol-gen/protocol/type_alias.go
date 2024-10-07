// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// TypeAlias defines a type alias. (e.g. type Definition = Location | LocationLink).
type TypeAlias struct {
	// The name of the type alias.
	Name string

	// The aliased type.
	Type Type

	// An optional documentation.
	Documentation string

	// Since when (release number) this structure is available. Is undefined if not known.
	Since string

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string

	// Whether this is a proposed type alias. If omitted, the type alias is final.
	Proposed bool

	// Whether the type alias is deprecated or not.
	// If deprecated the property contains the deprecation message.
	Deprecated string
}

func (TypeAlias) isTypeDecl() {}
