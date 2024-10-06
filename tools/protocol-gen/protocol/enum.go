// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Enumeration defines an enumeration.
type Enumeration struct {
	// Whether the enumeration is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string

	// An optional documentation.
	Documentation string

	// The name of the enumeration.
	Name string

	// Whether this is a proposed enumeration. If omitted, the enumeration is final.
	Proposed bool

	// Since when (release number) this enumeration is available. Is empty if not known.
	Since string

	// Whether the enumeration supports custom values (e.g. values which are not part of the set defined in values). If omitted no custom values are supported.
	SupportsCustomValues bool

	// The type of the elements.
	Type EnumerationType

	// The enum values.
	Values []*EnumerationEntry
}

func (Enumeration) isTypeDecl() {}

type EnumerationType struct {
	// Kind corresponds to the JSON schema field "kind".
	Kind string

	// Name corresponds to the JSON schema field "name".
	Name EnumerationTypeName
}

type EnumerationTypeName string

const (
	EnumerationNameInteger  EnumerationTypeName = "integer"
	EnumerationNameString   EnumerationTypeName = "string"
	EnumerationNameUinteger EnumerationTypeName = "uinteger"
)

// EnumerationEntry defines an enumeration entry.
type EnumerationEntry struct {
	// Whether the enum entry is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string

	// An optional documentation.
	Documentation string

	// The name of the enum item.
	Name string

	// Whether this is a proposed enumeration entry. If omitted, the enumeration entry is final.
	Proposed bool

	// Since when (release number) this enumeration entry is available. Is undefined if not known.
	Since string

	// The value (string or number).
	Value any
}
