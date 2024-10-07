// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"encoding/json"
	"fmt"
	"io"
)

// Decode decodes the LSP JSON to a MetaModel
func Decode(r io.Reader) (MetaModel, error) {
	out := MetaModel{}
	d := json.NewDecoder(r)
	d.DisallowUnknownFields()
	if err := d.Decode(&out); err != nil {
		return MetaModel{}, err
	}
	return out, nil
}

// MetaModel represents the actual meta model
type MetaModel struct {
	// Additional meta data.
	MetaData MetaData `json:"metaData" yaml:"metaData"`

	// The requests.
	Requests []Request `json:"requests" yaml:"requests"`

	// The notifications.
	Notifications []Notification `json:"notifications" yaml:"notifications"`

	// The structures.
	Structures []Structure `json:"structures" yaml:"structures"`

	// The enumerations.
	Enumerations []Enumeration `json:"enumerations" yaml:"enumerations"`

	// The type aliases.
	TypeAliases []TypeAlias `json:"typeAliases" yaml:"typeAliases"`
}

type MetaData struct {
	// The protocol version
	Version string `json:"version" yaml:"version"`
}

type BaseTypes string

const (
	URI         BaseTypes = "URI"
	DocumentUri BaseTypes = "DocumentUri"
	Integer     BaseTypes = "integer"
	Uinteger    BaseTypes = "uinteger"
	Decimal     BaseTypes = "decimal"
	RegExp      BaseTypes = "RegExp"
	String      BaseTypes = "string"
	Boolean     BaseTypes = "boolean"
	Null        BaseTypes = "null"
)

type TypeKind string

const (
	TypeKindBase           TypeKind = "base"
	TypeKindReference      TypeKind = "reference"
	TypeKindArray          TypeKind = "array"
	TypeKindMap            TypeKind = "map"
	TypeKindAnd            TypeKind = "and"
	TypeKindOr             TypeKind = "or"
	TypeKindTuple          TypeKind = "tuple"
	TypeKindLiteral        TypeKind = "literal"
	TypeKindStringLiteral  TypeKind = "stringLiteral"
	TypeKindIntegerLiteral TypeKind = "integerLiteral"
	TypeKindBooleanLiteral TypeKind = "booleanLiteral"
)

// MessageDirection indicates in which direction a message is sent in the protocol
type MessageDirection string

const (
	MessageDirectionClientToServer MessageDirection = "clientToServer"
	MessageDirectionServerToClient MessageDirection = "serverToClient"
	MessageDirectionBoth           MessageDirection = "both"
)

type Node interface {
	isType()
}

// Type represents a metaModel type.
type Type struct {
	Node Node
}

func (t *Type) UnmarshalJSON(data []byte) error {
	s := struct {
		Kind TypeKind
	}{}
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s.Kind {
	case TypeKindBase:
		t.Node = &BaseType{}
	case TypeKindReference:
		t.Node = &ReferenceType{}
	case TypeKindArray:
		t.Node = &ArrayType{}
	case TypeKindMap:
		t.Node = &MapType{}
	case TypeKindAnd:
		t.Node = &AndType{}
	case TypeKindOr:
		t.Node = &OrType{}
	case TypeKindTuple:
		t.Node = &TupleType{}
	case TypeKindLiteral:
		t.Node = &StructureLiteralType{}
	case TypeKindStringLiteral:
		t.Node = &StringLiteralType{}
	case TypeKindIntegerLiteral:
		t.Node = &IntegerLiteralType{}
	case TypeKindBooleanLiteral:
		t.Node = &BooleanLiteralType{}
	default:
		return fmt.Errorf("unhandled Type kind '%v'", s.Kind)
	}

	return json.Unmarshal(data, t.Node)
}

// Nodes represents a slice of [Type].
type Nodes struct {
	Nodes []Node
}

func (t *Nodes) UnmarshalJSON(data []byte) error {
	single := Type{}
	if err := json.Unmarshal(data, &single); err == nil {
		t.Nodes = []Node{single.Node}
		return nil
	}

	multi := []Type{}
	if err := json.Unmarshal(data, &multi); err != nil {
		return err
	}
	for _, e := range multi {
		t.Nodes = append(t.Nodes, e.Node)
	}

	return nil
}

// BaseType represents a base type like string or DocumentUri.
type BaseType struct {
	Name BaseTypes `json:"name" yaml:"name"`
}

var _ Node = BaseType{}

func (BaseType) isType() {}

// ReferenceType represents a reference to another type (e.g. TextDocument).
// This is either a Structure, a Enumeration or a TypeAlias in the same meta model.
type ReferenceType struct {
	Name string `json:"name" yaml:"name"`
}

var _ Node = ReferenceType{}

func (ReferenceType) isType()       {}
func (ReferenceType) isMapKeyType() {}

// ArrayType represents an array type (e.g. TextDocument[]).
type ArrayType struct {
	Element Type `json:"element" yaml:"element"`
}

var _ Node = ArrayType{}

func (ArrayType) isType() {}

// MapKeyType represents a type that can be used as a key in a map type.
// If a reference type is used then the type must either resolve to a string or integer type.
// (e.g. type ChangeAnnotationIdentifier === string).
type MapKeyType interface {
	isMapKeyType()
}

type MapKeyTypeBase string

func (MapKeyTypeBase) isMapKeyType() {}

const (
	MapKeyType_URI         MapKeyTypeBase = "URI"
	MapKeyType_DocumentUri MapKeyTypeBase = "DocumentUri"
	MapKeyType_String      MapKeyTypeBase = "string"
	MapKeyType_Integer     MapKeyTypeBase = "integer"
)

// MapType represents a JSON object map (e.g. interface Map<K extends string | integer, V> { [key: K] => V; }).
type MapType struct {
	Key   Type `json:"key" yaml:"key"`
	Value Type `json:"value" yaml:"value"`
}

var _ Node = MapType{}

func (MapType) isType() {}

// AndType represents an and type (e.g. TextDocumentParams & WorkDoneProgressParams).
type AndType struct {
	Items []Type `json:"items" yaml:"items"`
}

var _ Node = AndType{}

func (AndType) isType() {}

// OrType represents an or type (e.g. Location | LocationLink)
type OrType struct {
	Items []Type `json:"items" yaml:"items"`
}

var _ Node = OrType{}

func (OrType) isType() {}

// TupleType represents a tuple type (e.g. [integer, integer]).
type TupleType struct {
	Items []Type `json:"items" yaml:"items"`
}

var _ Node = TupleType{}

func (TupleType) isType() {}

// StructureLiteralType represents a literal structure (e.g. property: { start: uinteger; end: uinteger; })
type StructureLiteralType struct {
	Value StructureLiteral `json:"value" yaml:"value"`
}

var _ Node = StructureLiteralType{}

func (StructureLiteralType) isType() {}

// StructureLiteral defines an unnamed structure of an object literal
type StructureLiteral struct {
	// The properties.
	Properties []Property `json:"properties" yaml:"properties"`

	// An optional documentation.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Since when (release number) this structure is available. Is undefined if not known.
	Since string `json:"since,omitempty" yaml:"since,omitempty"`

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string `json:"sinceTags,omitempty" yaml:"sinceTags,omitempty"`

	// Whether this is a proposed structure. If omitted, the structure is final.
	Proposed bool `json:"proposed,omitempty" yaml:"proposed,omitempty"`

	// Whether the literal is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}

// StringLiteralType represents a string literal type (e.g. kind: 'rename')
type StringLiteralType struct {
	Value string `json:"value" yaml:"value"`
}

var _ Node = StringLiteralType{}

func (StringLiteralType) isType() {}

// IntegerLiteralType represents an integer literal type (e.g. kind: 1).
type IntegerLiteralType struct {
	Value bool `json:"value" yaml:"value"`
}

var _ Node = IntegerLiteralType{}

func (IntegerLiteralType) isType() {}

// BooleanLiteralType represents a boolean literal type (e.g. kind: true).
// kind: booleanLiteral
type BooleanLiteralType struct {
	Value bool `json:"value" yaml:"value"`
}

var _ Node = BooleanLiteralType{}

func (BooleanLiteralType) isType() {}

// Request represents a LSP request
type Request struct {
	// The request's method name.
	Method string `json:"method" yaml:"method"`

	// The parameter type(s) if any.
	Params Nodes `json:"params,omitempty" yaml:"params,omitempty"`

	// The result type.
	Result Type `json:"result" yaml:"result"`

	// Optional partial result type if the request supports partial result reporting.
	PartialResult Type `json:"partialResult,omitempty" yaml:"partialResult,omitempty"`

	// An optional error data type.
	ErrorData Type `json:"errorData,omitempty" yaml:"errorData,omitempty"`

	// Optional a dynamic registration method if it different from the request's method.
	RegistrationMethod string `json:"registrationMethod,omitempty" yaml:"registrationMethod,omitempty"`

	// Optional registration options if the request supports dynamic registration.
	RegistrationOptions Type `json:"registrationOptions,omitempty" yaml:"registrationOptions,omitempty"`

	// The direction in which this request is sent in the protocol.
	MessageDirection MessageDirection `json:"messageDirection" yaml:"messageDirection"`

	// An optional documentation.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Since when (release number) this request is available. Is undefined if not known.
	Since string `json:"since,omitempty" yaml:"since,omitempty"`

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string `json:"sinceTags,omitempty" yaml:"sinceTags,omitempty"`

	// Whether this is a proposed feature. If omitted the feature is final.
	Proposed bool `json:"proposed,omitempty" yaml:"proposed,omitempty"`

	// Whether the request is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// The type name of the request if any.
	TypeName string `json:"typeName,omitempty" yaml:"typeName,omitempty"`
}

// Notification represents a LSP notification
type Notification struct {
	// The request's method name.
	Method string `json:"method" yaml:"method"`

	// The parameter type(s) if any.
	Params Nodes `json:"params,omitempty" yaml:"params,omitempty"`

	// Optional a dynamic registration method if it different from the request's method.
	RegistrationMethod string `json:"registrationMethod,omitempty" yaml:"registrationMethod,omitempty"`

	// Optional registration options if the notification supports dynamic registration.
	RegistrationOptions Type `json:"registrationOptions,omitempty" yaml:"registrationOptions,omitempty"`

	// The direction in which this notification is sent in the protocol.
	MessageDirection MessageDirection `json:"messageDirection" yaml:"messageDirection"`

	// An optional documentation.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Since when (release number) this notification is available. Is undefined if not known.
	Since string `json:"since,omitempty" yaml:"since,omitempty"`

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string `json:"sinceTags,omitempty" yaml:"sinceTags,omitempty"`

	// Whether this is a proposed notification. If omitted the notification is final.
	Proposed bool `json:"proposed,omitempty" yaml:"proposed,omitempty"`

	// Whether the notification is deprecated or not.
	// If deprecated the property contains the deprecation message.
	Deprecated string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// The type name of the request if any.
	TypeName string `json:"typeName,omitempty" yaml:"typeName,omitempty"`
}

// Structure defines the structure of an object literal
type Structure struct {
	// The name of the structure.
	Name string `json:"name" yaml:"name"`

	// Structures extended from. This structures form a polymorphic type hierarchy.
	Extends []Type `json:"extends,omitempty" yaml:"extends,omitempty"`

	// Structures to mix in. The properties of these structures are `copied` into this structure. Mixins don't form a polymorphic type hierarchy in LSP.
	Mixins []Type `json:"mixins,omitempty" yaml:"mixins,omitempty"`

	// The properties.
	Properties []Property `json:"properties" yaml:"properties"`

	// An optional documentation.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Since when (release number) this structure is available. Is undefined if not known.
	Since string `json:"since,omitempty" yaml:"since,omitempty"`

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string `json:"sinceTags,omitempty" yaml:"sinceTags,omitempty"`

	// Whether this is a proposed structure. If omitted, the structure is final.
	Proposed bool `json:"proposed,omitempty" yaml:"proposed,omitempty"`

	// Whether the structure is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}

// Enumeration defines an enumeration.
type Enumeration struct {
	// The name of the enumeration.
	Name string `json:"name" yaml:"name"`

	// The type of the elements.
	Type EnumerationType `json:"type" yaml:"type"`

	// The enum values.
	Values []EnumerationEntry `json:"values" yaml:"values"`

	// Whether the enumeration supports custom values (e.g. values which are not part of the set defined in values). If omitted no custom values are supported.
	SupportsCustomValues bool `json:"supportsCustomValues,omitempty" yaml:"supportsCustomValues,omitempty"`

	// An optional documentation.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Since when (release number) this enumeration is available. Is empty if not known.
	Since string `json:"since,omitempty" yaml:"since,omitempty"`

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string `json:"sinceTags,omitempty" yaml:"sinceTags,omitempty"`

	// Whether this is a proposed enumeration. If omitted, the enumeration is final.
	Proposed bool `json:"proposed,omitempty" yaml:"proposed,omitempty"`

	// Whether the enumeration is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}

type EnumerationType struct {
	// Kind corresponds to the JSON schema field "kind".
	Kind string `json:"kind" yaml:"kind"`

	// Name corresponds to the JSON schema field "name".
	Name EnumerationName `json:"name" yaml:"name"`
}

type EnumerationName string

const (
	EnumerationNameInteger  EnumerationName = "integer"
	EnumerationNameString   EnumerationName = "string"
	EnumerationNameUinteger EnumerationName = "uinteger"
)

// EnumerationEntry defines an enumeration entry
type EnumerationEntry struct {
	// The name of the enum item.
	Name string `json:"name,strictcase" yaml:"name"`

	// The value (string or number).
	Value any `json:"value" yaml:"value"`

	// An optional documentation.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Since when (release number) this enumeration entry is available. Is undefined if not known.
	Since string `json:"since,omitempty" yaml:"since,omitempty"`

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string `json:"sinceTags,omitempty" yaml:"sinceTags,omitempty"`

	// Whether this is a proposed enumeration entry. If omitted, the enumeration entry is final.
	Proposed bool `json:"proposed,omitempty" yaml:"proposed,omitempty"`

	// Whether the enum entry is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// The type name of the request if any.
	TypeName string `json:"typeName,omitempty" yaml:"typeName,omitempty"`
}

// TypeAlias defines a type alias. (e.g. type Definition = Location | LocationLink)
type TypeAlias struct {
	// The name of the type alias.
	Name string `json:"name" yaml:"name"`

	// The aliased type.
	Type Type `json:"type" yaml:"type"`

	// An optional documentation.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Since when (release number) this structure is available. Is undefined if not known.
	Since string `json:"since,omitempty" yaml:"since,omitempty"`

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string `json:"sinceTags,omitempty" yaml:"sinceTags,omitempty"`

	// Whether this is a proposed type alias. If omitted, the type alias is final.
	Proposed bool `json:"proposed,omitempty" yaml:"proposed,omitempty"`

	// Whether the type alias is deprecated or not.
	// If deprecated the property contains the deprecation message.
	Deprecated string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}

// Property represents an object property
type Property struct {
	// The property name.
	Name string `json:"name" yaml:"name"`

	// The type of the property.
	Type Type `json:"type" yaml:"type"`

	// Whether the property is optional. If omitted, the property is mandatory.
	Optional bool `json:"optional,omitempty" yaml:"optional,omitempty"`

	// An optional documentation.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Since when (release number) this property is available. Is undefined if not known.
	Since string `json:"since,omitempty" yaml:"since,omitempty"`

	// All since tags in case there was more than one tag. Is undefined if not known.
	SinceTags []string `json:"sinceTags,omitempty" yaml:"sinceTags,omitempty"`

	// Whether this is a proposed property. If omitted, the structure is final.
	Proposed bool `json:"proposed,omitempty" yaml:"proposed,omitempty"`

	// Whether the property is deprecated or not. If deprecated the property contains the deprecation message.
	Deprecated string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}
