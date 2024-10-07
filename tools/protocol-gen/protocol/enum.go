// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Enumeration defines an enumeration.
type Enumeration struct {
	// The name of the enumeration.
	Name string

	// The type of the elements.
	Type EnumerationType

	// The enum values.
	Values []*EnumerationEntry

	// Whether the enumeration supports custom values (e.g. values which are not part of the set defined in values). If omitted no custom values are supported.
	SupportsCustomValues bool

	// An optional documentation.
	Documentation string

	// Since when (release number) this enumeration is available. Is empty if not known.
	Since string

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string

	// Whether this is a proposed enumeration. If omitted, the enumeration is final.
	Proposed bool

	// Whether the enumeration is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string
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
	// The name of the enum item.
	Name string

	// The value (string or number).
	Value any

	// An optional documentation.
	Documentation string

	// Since when (release number) this enumeration entry is available. Is undefined if not known.
	Since string

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string

	// Whether this is a proposed enumeration entry. If omitted, the enumeration entry is final.
	Proposed bool

	// Whether the enum entry is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string

	// The type name of the request if any.
	TypeName string
}
