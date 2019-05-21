// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"github.com/francoispqt/gojay"
)

// Diagnostics represents a slice of Diagnostics.
type Diagnostics []Diagnostic

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *Diagnostics) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value = Diagnostic{}
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v Diagnostics) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v Diagnostics) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *PublishDiagnosticsParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyDiagnostics:
		var value = Diagnostics{}
		err := dec.Array(&value)
		if err == nil && len(value) > 0 {
			v.Diagnostics = []Diagnostic(value)
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *PublishDiagnosticsParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *PublishDiagnosticsParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, string(v.URI))
	enc.ArrayKey(keyDiagnostics, Diagnostics(v.Diagnostics))
}

// IsNil returns wether the structure is nil value or not.
func (v *PublishDiagnosticsParams) IsNil() bool { return v == nil }

// Reset reset fields.
func (v *PublishDiagnosticsParams) Reset() {
	v.URI = DocumentURI("")
	for i := range v.Diagnostics {
		v.Diagnostics[i].Reset()
		DiagnosticPool.Put(&v.Diagnostics[i])
	}
	v.Diagnostics = nil
}
