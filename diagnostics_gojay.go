// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gojay
// +build gojay

package protocol

import (
	"github.com/francoispqt/gojay"
)

// Diagnostics represents a slice of Diagnostics.
type Diagnostics []Diagnostic

// compile time check whether the Diagnostics implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*Diagnostics)(nil)
	_ gojay.UnmarshalerJSONArray = (*Diagnostics)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v Diagnostics) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *Diagnostics) UnmarshalJSONArray(dec *gojay.Decoder) error {
	value := Diagnostic{}
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// IsNil implements gojay.MarshalerJSONArray.
func (v Diagnostics) IsNil() bool { return len(v) == 0 }

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *PublishDiagnosticsParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, string(v.URI))
	enc.Uint32KeyOmitEmpty(keyVersion, v.Version)
	enc.ArrayKey(keyDiagnostics, Diagnostics(v.Diagnostics))
}

// IsNil returns wether the structure is nil value or not.
func (v *PublishDiagnosticsParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *PublishDiagnosticsParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyVersion:
		return dec.Uint32(&v.Version)
	case keyDiagnostics:
		value := Diagnostics{}
		err := dec.Array(&value)
		if err == nil && len(value) > 0 {
			v.Diagnostics = []Diagnostic(value)
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *PublishDiagnosticsParams) NKeys() int { return 3 }

// compile time check whether the PublishDiagnosticsParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*PublishDiagnosticsParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*PublishDiagnosticsParams)(nil)
)
