// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// Interfaces represents a slice of interface.
type Interfaces []interface{}

// UnmarshalJSONArray decodes JSON array elements into slice
func (v *Interfaces) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var t interface{}
	if err := dec.Interface(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Interfaces) NKeys() int { return 1 }

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

// Strings represents a slice of string.
type Strings []string

// UnmarshalJSONArray decodes JSON array elements into slice
func (v *Strings) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := ""
	if err := dec.String(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

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
