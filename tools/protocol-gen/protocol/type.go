// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Type is an LSP type
type Type interface {
	isType()
	SubTypes() []Type
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

func (URIType) isType()          {}
func (URIType) SubTypes() []Type { return nil }

type DocumentUriType struct{}

func (DocumentUriType) isType()           {}
func (*DocumentUriType) SubTypes() []Type { return nil }

type IntegerType struct{}

func (IntegerType) isType()          {}
func (IntegerType) SubTypes() []Type { return nil }

type UintegerType struct{}

func (UintegerType) isType()          {}
func (UintegerType) SubTypes() []Type { return nil }

type DecimalType struct{}

func (DecimalType) isType()          {}
func (DecimalType) SubTypes() []Type { return nil }

type RegExpType struct{}

func (RegExpType) isType()          {}
func (RegExpType) SubTypes() []Type { return nil }

type StringType struct{}

func (StringType) isType()          {}
func (StringType) SubTypes() []Type { return nil }

type BooleanType struct{}

func (BooleanType) isType()          {}
func (BooleanType) SubTypes() []Type { return nil }

type NullType struct{}

func (NullType) isType()          {}
func (NullType) SubTypes() []Type { return nil }

// AndType represents an and type (e.g. TextDocumentParams & WorkDoneProgressParams).
type AndType struct{ Items []Type }

func (AndType) isType()             {}
func (t *AndType) SubTypes() []Type { return t.Items }

// ArrayType represents an array type (e.g. TextDocument[]).
type ArrayType struct{ Element Type }

func (ArrayType) isType()             {}
func (t *ArrayType) SubTypes() []Type { return []Type{t.Element} }

// BooleanLiteralType represents a boolean literal type (e.g. kind: true).
// kind: booleanLiteral
type BooleanLiteralType struct{ Value bool }

func (BooleanLiteralType) isType()          {}
func (BooleanLiteralType) SubTypes() []Type { return nil }

// IntegerLiteralType represents an integer literal type (e.g. kind: 1).
type IntegerLiteralType struct{ Value bool }

func (IntegerLiteralType) isType()          {}
func (IntegerLiteralType) SubTypes() []Type { return nil }

// MapType represents a JSON object map (e.g. interface Map<K extends string | integer, V> { [key: K] => V; }).
type MapType struct {
	Key   Type
	Value Type
}

func (MapType) isType()             {}
func (t *MapType) SubTypes() []Type { return []Type{t.Key, t.Value} }

// OrType represents an or type (e.g. Location | LocationLink)
type OrType struct{ Items []Type }

func (OrType) isType()             {}
func (t *OrType) SubTypes() []Type { return t.Items }

// StringLiteralType represents a string literal type (e.g. kind: 'rename')
type StringLiteralType struct{ Value string }

func (StringLiteralType) isType()          {}
func (StringLiteralType) SubTypes() []Type { return nil }

// StructureLiteralType represents a literal structure (e.g. property: { start: uinteger; end: uinteger; })
type StructureLiteralType struct{ Value *StructureLiteral }

func (StructureLiteralType) isType()          {}
func (StructureLiteralType) SubTypes() []Type { return nil }

// ReferenceType represents a reference to another type (e.g. TextDocument).
// This is either a Structure, a Enumeration or a TypeAlias in the same meta model.
type ReferenceType struct {
	Name     string
	TypeDecl TypeDecl
}

func (ReferenceType) isType()          {}
func (ReferenceType) SubTypes() []Type { return nil }

// TupleType represents a tuple type (e.g. [integer, integer]).
type TupleType struct{ Items []Type }

func (TupleType) isType()             {}
func (t *TupleType) SubTypes() []Type { return t.Items }
