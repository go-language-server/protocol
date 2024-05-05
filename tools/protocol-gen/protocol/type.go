// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/flect"
)

// Type is the LSP type.
type Type interface {
	fmt.Stringer

	isType()
	HasSubTypes() bool
	SubTypes() []Type
}

type BaseType interface {
	fmt.Stringer

	isBaseType()
}

func RecursiveTypesOf(t Type) []Type {
	stack := []Type{t}
	seen := map[Type]struct{}{}
	types := []Type{}
	for len(stack) > 0 {
		stack, t = stack[:len(stack)-1], stack[len(stack)-1]
		if _, existing := seen[t]; !existing {
			seen[t] = struct{}{}
			types = append(types, t)
			stack = append(stack, t.SubTypes()...)
		}
	}
	return types
}

type URIType struct{}

var (
	_ Type     = URIType{}
	_ BaseType = URIType{}
)

func (URIType) isType()           {}
func (URIType) HasSubTypes() bool { return false }
func (URIType) SubTypes() []Type  { return nil }
func (t URIType) String() string  { return "uri.URI" }
func (URIType) isBaseType()       {}

type DocumentUriType struct{}

var (
	_ Type     = DocumentUriType{}
	_ BaseType = DocumentUriType{}
)

func (DocumentUriType) isType()           {}
func (DocumentUriType) HasSubTypes() bool { return false }
func (DocumentUriType) SubTypes() []Type  { return nil }
func (t DocumentUriType) String() string  { return flect.Pascalize("DocumentUri") }
func (DocumentUriType) isBaseType()       {}

type IntegerType struct{}

var (
	_ Type     = IntegerType{}
	_ BaseType = IntegerType{}
)

func (IntegerType) isType()           {}
func (IntegerType) HasSubTypes() bool { return false }
func (IntegerType) SubTypes() []Type  { return nil }
func (t IntegerType) String() string  { return "int32" }
func (IntegerType) isBaseType()       {}

type UintegerType struct{}

var (
	_ Type     = UintegerType{}
	_ BaseType = UintegerType{}
)

func (UintegerType) isType()           {}
func (UintegerType) HasSubTypes() bool { return false }
func (UintegerType) SubTypes() []Type  { return nil }
func (t UintegerType) String() string  { return "uint32" }
func (UintegerType) isBaseType()       {}

type DecimalType struct{}

var (
	_ Type     = DecimalType{}
	_ BaseType = DecimalType{}
)

func (DecimalType) isType()           {}
func (DecimalType) HasSubTypes() bool { return false }
func (DecimalType) SubTypes() []Type  { return nil }
func (t DecimalType) String() string  { return "float64" }
func (DecimalType) isBaseType()       {}

type StringType struct{}

var (
	_ Type     = StringType{}
	_ BaseType = StringType{}
)

func (StringType) isType()           {}
func (StringType) HasSubTypes() bool { return false }
func (StringType) SubTypes() []Type  { return nil }
func (t StringType) String() string  { return "string" }
func (StringType) isBaseType()       {}

type BooleanType struct{}

var (
	_ Type     = BooleanType{}
	_ BaseType = BooleanType{}
)

func (BooleanType) isType()           {}
func (BooleanType) HasSubTypes() bool { return false }
func (BooleanType) SubTypes() []Type  { return nil }
func (t BooleanType) String() string  { return "bool" }
func (BooleanType) isBaseType()       {}

type NullType struct{}

var _ Type = NullType{}

func (NullType) isType()           {}
func (NullType) HasSubTypes() bool { return false }
func (NullType) SubTypes() []Type  { return nil }
func (t NullType) String() string  { return "nil" }

type RegExpType struct{}

var _ Type = RegExpType{}

func (RegExpType) isType()           {}
func (RegExpType) HasSubTypes() bool { return false }
func (RegExpType) SubTypes() []Type  { return nil }
func (t RegExpType) String() string  { return "*regexp.Regexp" }

// AndType represents an and type (e.g. TextDocumentParams & WorkDoneProgressParams).
type AndType struct{ Items []Type }

var _ Type = (*AndType)(nil)

func (AndType) isType()             {}
func (AndType) HasSubTypes() bool   { return true }
func (t *AndType) SubTypes() []Type { return t.Items }
func (t *AndType) String() string   { return "TODO(zchee): AndType" } // TODO(zchee): implements

// ArrayType represents an array type (e.g. TextDocument[]).
type ArrayType struct{ Element Type }

var _ Type = (*ArrayType)(nil)

func (ArrayType) isType()             {}
func (ArrayType) HasSubTypes() bool   { return true }
func (t *ArrayType) SubTypes() []Type { return []Type{t.Element} }
func (t *ArrayType) String() string   { return "TODO(zchee): ArrayType" } // TODO(zchee): implements

// BooleanLiteralType represents a boolean literal type (e.g. kind: true).
// kind: booleanLiteral
type BooleanLiteralType struct{ Value bool }

var _ Type = BooleanLiteralType{}

func (BooleanLiteralType) isType()           {}
func (BooleanLiteralType) HasSubTypes() bool { return false }
func (BooleanLiteralType) SubTypes() []Type  { return nil }
func (t BooleanLiteralType) String() string  { return strconv.FormatBool(t.Value) } // TODO(zchee): implements

// IntegerLiteralType represents an integer literal type (e.g. kind: 1).
type IntegerLiteralType struct{ Value int }

var _ Type = IntegerLiteralType{}

func (IntegerLiteralType) isType()           {}
func (IntegerLiteralType) HasSubTypes() bool { return false }
func (IntegerLiteralType) SubTypes() []Type  { return nil }
func (t IntegerLiteralType) String() string  { return strconv.FormatInt(int64(t.Value), 10) } // TODO(zchee): implements

// MapType represents a JSON object map (e.g. interface Map<K extends string | integer, V> { [key: K] => V; }).
type MapType struct {
	Key   Type
	Value Type
}

var _ Type = (*MapType)(nil)

func (MapType) isType()             {}
func (MapType) HasSubTypes() bool   { return true }
func (t *MapType) SubTypes() []Type { return []Type{t.Key, t.Value} }
func (t *MapType) String() string   { return "TODO(zchee): MapType" } // TODO(zchee): implements

// OrType represents an or type (e.g. Location | LocationLink)
type OrType struct{ Items []Type }

var _ Type = (*OrType)(nil)

func (OrType) isType()             {}
func (OrType) HasSubTypes() bool   { return true }
func (t *OrType) SubTypes() []Type { return t.Items }
func (t *OrType) String() string   { return "TODO(zchee): OrType" } // TODO(zchee): implements

// StringLiteralType represents a string literal type (e.g. kind: 'rename').
type StringLiteralType struct{ Value string }

var _ Type = (*StringLiteralType)(nil)

func (StringLiteralType) isType()           {}
func (StringLiteralType) HasSubTypes() bool { return false }
func (StringLiteralType) SubTypes() []Type  { return nil }
func (t *StringLiteralType) String() string { return t.Value } // TODO(zchee): implements

// StructureLiteralType represents a literal structure (e.g. property: { start: uinteger; end: uinteger; }).
type StructureLiteralType struct{ Value *StructureLiteral }

var _ Type = (*StructureLiteralType)(nil)

func (StructureLiteralType) isType()           {}
func (StructureLiteralType) HasSubTypes() bool { return false }
func (StructureLiteralType) SubTypes() []Type  { return nil }
func (t *StructureLiteralType) String() string { return "TODO(zchee): StructureLiteralType" } // TODO(zchee): implements

// ReferenceType represents a reference to another type (e.g. TextDocument).
// This is either a Structure, a Enumeration or a TypeAlias in the same meta model.
type ReferenceType struct {
	Name     string
	TypeDecl TypeDecl
}

var _ Type = (*ReferenceType)(nil)

func (ReferenceType) isType()           {}
func (ReferenceType) HasSubTypes() bool { return false }
func (ReferenceType) SubTypes() []Type  { return nil }
func (t ReferenceType) String() string  { return t.Name } // TODO(zchee): correct?

// TupleType represents a tuple type (e.g. [integer, integer]).
type TupleType struct{ Items []Type }

var _ Type = (*TupleType)(nil)

func (TupleType) isType()             {}
func (TupleType) HasSubTypes() bool   { return true }
func (t *TupleType) SubTypes() []Type { return t.Items }
func (t *TupleType) String() string   { return "TODO(zchee): TupleType" } // TODO(zchee): implements
