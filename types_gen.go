// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

type interfaces []interface{}

func (v *interfaces) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var t interface{}
	if err := dec.Interface(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *interfaces) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *interfaces) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.AddInterface(t)
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *interfaces) IsNil() bool {
	return v == nil || len(*v) == 0
}

type stringSlice []string

func (v *stringSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := ""
	if err := dec.String(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *stringSlice) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *stringSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.String(t)
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *stringSlice) IsNil() bool {
	return v == nil || len(*v) == 0
}
