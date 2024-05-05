// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// TypeAlias defines a type alias. (e.g. type Definition = Location | LocationLink).
type TypeAlias struct {
	// Whether the type alias is deprecated or not.
	// If deprecated the property contains the deprecation message.
	Deprecated string

	// An optional documentation.
	Documentation string

	// The name of the type alias.
	Name string

	// Whether this is a proposed type alias. If omitted, the type alias is final.
	Proposed bool

	// Since when (release number) this structure is available. Is undefined if not known.
	Since string

	// The aliased type.
	Type Type
}

func (TypeAlias) isTypeDecl() {}
