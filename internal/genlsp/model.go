// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

// Package genlsp implements a generator that lowers the LSP meta-model
// (metaModel.json) into Go source for the go.lsp.dev/protocol package.
//
// The data model in this file mirrors metaModel.schema.json (JSON Schema
// draft-07). The central construct is [Type], a discriminated union keyed by
// "kind"; it is decoded by a hand-written [Type.UnmarshalJSON] that routes the
// payload to the fields relevant for that kind.
package genlsp

import (
	"fmt"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// MetaModel is the root of metaModel.json.
type MetaModel struct {
	MetaData      MetaData        `json:"metaData"`
	Requests      []*Request      `json:"requests"`
	Notifications []*Notification `json:"notifications"`
	Structures    []*Structure    `json:"structures"`
	Enumerations  []*Enumeration  `json:"enumerations"`
	TypeAliases   []*TypeAlias    `json:"typeAliases"`
}

// MetaData holds additional meta data of the model.
//
// Note: Version is unreliable (the 3.18 model reports "3.17.0"); never key
// generation decisions on it.
type MetaData struct {
	Version string `json:"version"`
}

// TypeKind enumerates the discriminator values of [Type].
type TypeKind string

// The set of [Type] kinds defined by the meta-schema.
const (
	KindBase           TypeKind = "base"
	KindReference      TypeKind = "reference"
	KindArray          TypeKind = "array"
	KindMap            TypeKind = "map"
	KindAnd            TypeKind = "and"
	KindOr             TypeKind = "or"
	KindTuple          TypeKind = "tuple"
	KindLiteral        TypeKind = "literal"
	KindStringLiteral  TypeKind = "stringLiteral"
	KindIntegerLiteral TypeKind = "integerLiteral"
	KindBooleanLiteral TypeKind = "booleanLiteral"
)

// BaseTypeName enumerates the names a "base" [Type] may carry.
type BaseTypeName string

// The set of base type names defined by the meta-schema.
const (
	BaseURI         BaseTypeName = "URI"
	BaseDocumentURI BaseTypeName = "DocumentUri"
	BaseInteger     BaseTypeName = "integer"
	BaseUinteger    BaseTypeName = "uinteger"
	BaseDecimal     BaseTypeName = "decimal"
	BaseRegExp      BaseTypeName = "RegExp"
	BaseString      BaseTypeName = "string"
	BaseBoolean     BaseTypeName = "boolean"
	BaseNull        BaseTypeName = "null"
)

// Type is a meta-model type. The active fields depend on Kind:
//
//   - base:           Name
//   - reference:      Name
//   - array:          Element
//   - map:            Key, Value
//   - and/or/tuple:   Items
//   - literal:        Literal
//   - stringLiteral:  StringValue
//   - integerLiteral: IntegerValue
//   - booleanLiteral: BooleanValue
type Type struct {
	Kind TypeKind

	Name         string            // base, reference
	Element      *Type             // array
	Key          *Type             // map
	Value        *Type             // map
	Items        []*Type           // and, or, tuple
	Literal      *StructureLiteral // literal
	StringValue  string            // stringLiteral
	IntegerValue int64             // integerLiteral
	BooleanValue bool              // booleanLiteral
}

// UnmarshalJSON decodes a meta-model type, dispatching on its "kind".
func (t *Type) UnmarshalJSON(data []byte) error {
	var probe struct {
		Kind TypeKind `json:"kind"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return fmt.Errorf("type kind: %w", err)
	}
	t.Kind = probe.Kind

	switch probe.Kind {
	case KindBase, KindReference:
		var v struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		t.Name = v.Name
	case KindArray:
		var v struct {
			Element *Type `json:"element"`
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		t.Element = v.Element
	case KindMap:
		var v struct {
			Key   *Type `json:"key"`
			Value *Type `json:"value"`
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		t.Key, t.Value = v.Key, v.Value
	case KindAnd, KindOr, KindTuple:
		var v struct {
			Items []*Type `json:"items"`
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		t.Items = v.Items
	case KindLiteral:
		var v struct {
			Value *StructureLiteral `json:"value"`
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		t.Literal = v.Value
	case KindStringLiteral:
		var v struct {
			Value string `json:"value"`
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		t.StringValue = v.Value
	case KindIntegerLiteral:
		var v struct {
			Value int64 `json:"value"`
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		t.IntegerValue = v.Value
	case KindBooleanLiteral:
		var v struct {
			Value bool `json:"value"`
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		t.BooleanValue = v.Value
	default:
		return fmt.Errorf("unknown type kind %q", probe.Kind)
	}
	return nil
}

// Structure defines the structure of an object literal.
type Structure struct {
	Name          string      `json:"name"`
	Properties    []*Property `json:"properties"`
	Extends       []*Type     `json:"extends"`
	Mixins        []*Type     `json:"mixins"`
	Documentation string      `json:"documentation"`
	Since         string      `json:"since"`
	SinceTags     []string    `json:"sinceTags"`
	Deprecated    string      `json:"deprecated"`
	Proposed      bool        `json:"proposed"`
}

// Property is an object property.
type Property struct {
	Name          string   `json:"name"`
	Type          *Type    `json:"type"`
	Optional      bool     `json:"optional"`
	Documentation string   `json:"documentation"`
	Since         string   `json:"since"`
	SinceTags     []string `json:"sinceTags"`
	Deprecated    string   `json:"deprecated"`
	Proposed      bool     `json:"proposed"`
}

// StructureLiteral is an unnamed object literal (a "literal" [Type]'s value).
type StructureLiteral struct {
	Properties    []*Property `json:"properties"`
	Documentation string      `json:"documentation"`
	Since         string      `json:"since"`
	SinceTags     []string    `json:"sinceTags"`
	Deprecated    string      `json:"deprecated"`
	Proposed      bool        `json:"proposed"`
}

// Enumeration defines an enumeration.
type Enumeration struct {
	Name                 string              `json:"name"`
	Type                 EnumerationBaseType `json:"type"`
	Values               []*EnumerationEntry `json:"values"`
	SupportsCustomValues bool                `json:"supportsCustomValues"`
	Documentation        string              `json:"documentation"`
	Since                string              `json:"since"`
	SinceTags            []string            `json:"sinceTags"`
	Deprecated           string              `json:"deprecated"`
	Proposed             bool                `json:"proposed"`
}

// EnumerationBaseType is the element type of an [Enumeration]: string, integer
// or uinteger.
type EnumerationBaseType struct {
	Kind TypeKind     `json:"kind"`
	Name BaseTypeName `json:"name"`
}

// EnumerationEntry is a single enumeration value. Value holds the raw JSON
// token (a string or a number) and is interpreted using the owning
// [Enumeration]'s base type.
type EnumerationEntry struct {
	Name          string         `json:"name"`
	Value         jsontext.Value `json:"value"`
	Documentation string         `json:"documentation"`
	Since         string         `json:"since"`
	SinceTags     []string       `json:"sinceTags"`
	Deprecated    string         `json:"deprecated"`
	Proposed      bool           `json:"proposed"`
}

// TypeAlias defines a type alias (e.g. Definition = Location | LocationLink).
type TypeAlias struct {
	Name          string   `json:"name"`
	Type          *Type    `json:"type"`
	Documentation string   `json:"documentation"`
	Since         string   `json:"since"`
	SinceTags     []string `json:"sinceTags"`
	Deprecated    string   `json:"deprecated"`
	Proposed      bool     `json:"proposed"`
}

// MessageDirection indicates in which direction a message is sent.
type MessageDirection string

// The set of message directions defined by the meta-schema.
const (
	DirectionClientToServer MessageDirection = "clientToServer"
	DirectionServerToClient MessageDirection = "serverToClient"
	DirectionBoth           MessageDirection = "both"
)

// Request represents an LSP request.
type Request struct {
	Method              string           `json:"method"`
	Result              *Type            `json:"result"`
	MessageDirection    MessageDirection `json:"messageDirection"`
	Params              Params           `json:"params"`
	PartialResult       *Type            `json:"partialResult"`
	RegistrationOptions *Type            `json:"registrationOptions"`
	RegistrationMethod  string           `json:"registrationMethod"`
	ErrorData           *Type            `json:"errorData"`
	Documentation       string           `json:"documentation"`
	Since               string           `json:"since"`
	Deprecated          string           `json:"deprecated"`
	Proposed            bool             `json:"proposed"`
	TypeName            string           `json:"typeName"`
}

// Notification represents an LSP notification.
type Notification struct {
	Method              string           `json:"method"`
	MessageDirection    MessageDirection `json:"messageDirection"`
	Params              Params           `json:"params"`
	RegistrationOptions *Type            `json:"registrationOptions"`
	RegistrationMethod  string           `json:"registrationMethod"`
	Documentation       string           `json:"documentation"`
	Since               string           `json:"since"`
	Deprecated          string           `json:"deprecated"`
	Proposed            bool             `json:"proposed"`
	TypeName            string           `json:"typeName"`
}

// Params is the parameter type(s) of a request or notification. The schema
// allows either a single Type or an array of Type; both decode to this slice.
type Params []*Type

// UnmarshalJSON accepts either a single type object or an array of them.
func (p *Params) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '[' {
		var arr []*Type
		if err := json.Unmarshal(data, &arr); err != nil {
			return err
		}
		*p = arr
		return nil
	}
	var single *Type
	if err := json.Unmarshal(data, &single); err != nil {
		return err
	}
	if single == nil {
		*p = nil
		return nil
	}
	*p = Params{single}
	return nil
}
