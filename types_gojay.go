// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import "github.com/francoispqt/gojay"

// Interfaces represents a slice of interface.
type Interfaces []interface{}

// compile time check whether the Interfaces implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*Interfaces)(nil)
	_ gojay.UnmarshalerJSONArray = (*Interfaces)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v Interfaces) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.AddInterface(v[i])
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v Interfaces) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *Interfaces) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var t interface{}
	if err := dec.Interface(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// Strings represents a slice of string.
type Strings []string

// compile time check whether the Strings implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*Strings)(nil)
	_ gojay.UnmarshalerJSONArray = (*Strings)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v Strings) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.String(v[i])
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v Strings) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *Strings) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := ""
	if err := dec.String(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// StringInterfaceMap represents a string key and interface value map.
type StringInterfaceMap map[string]interface{}

// compile time check whether the Interfaces implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var (
	_ gojay.MarshalerJSONObject   = (*StringInterfaceMap)(nil)
	_ gojay.UnmarshalerJSONObject = (*StringInterfaceMap)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (m StringInterfaceMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range m {
		enc.AddInterfaceKey(k, v)
	}
}

// IsNil implements gojay.MarshalerJSONObject.
func (m StringInterfaceMap) IsNil() bool {
	return m == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (m StringInterfaceMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var iface interface{}
	err := dec.Interface(&iface)
	if err != nil {
		return err
	}
	m[k] = iface
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (m StringInterfaceMap) NKeys() int {
	return 0
}

// StringStringMap represents a string key and string value map.
type StringStringMap map[string]string

// compile time check whether the Interfaces implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var (
	_ gojay.MarshalerJSONObject   = (*StringStringMap)(nil)
	_ gojay.UnmarshalerJSONObject = (*StringStringMap)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (m StringStringMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range m {
		enc.StringKey(k, v)
	}
}

// IsNil implements gojay.MarshalerJSONObject.
func (m StringStringMap) IsNil() bool {
	return m == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (m StringStringMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	s := ""
	err := dec.String(&s)
	if err != nil {
		return err
	}
	m[k] = s
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (m StringStringMap) NKeys() int {
	return 0
}
