// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import "github.com/francoispqt/gojay"

// Interfaces represents a slice of interface.
type Interfaces []interface{}

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *Interfaces) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.AddInterface(t)
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *Interfaces) IsNil() bool {
	return len(*v) == 0
}

// UnmarshalJSONArray decodes JSON array elements into slice
func (v *Interfaces) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var t interface{}
	if err := dec.Interface(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// compile time check whether the Interfaces implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interface.
var _ gojay.MarshalerJSONArray = (*Interfaces)(nil)
var _ gojay.UnmarshalerJSONArray = (*Interfaces)(nil)

// Strings represents a slice of string.
type Strings []string

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *Strings) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.String(t)
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *Strings) IsNil() bool {
	return len(*v) == 0
}

// UnmarshalJSONArray decodes JSON array elements into slice
func (v *Strings) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := ""
	if err := dec.String(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// compile time check whether the Strings implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interface.
var _ gojay.MarshalerJSONArray = (*Strings)(nil)
var _ gojay.UnmarshalerJSONArray = (*Strings)(nil)
