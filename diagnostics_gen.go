// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

type diagnostics []Diagnostic

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *diagnostics) UnmarshalJSONArray(dec *gojay.Decoder) error {
	s := Diagnostic{}
	if err := dec.Object(&s); err != nil {
		return err
	}
	*v = append(*v, s)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *diagnostics) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *diagnostics) MarshalJSONArray(enc *gojay.Encoder) {
	vv := *v
	for i := range vv {
		enc.Object(&vv[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *diagnostics) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *PublishDiagnosticsParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyDiagnostics:
		return dec.Array((*diagnostics)(&v.Diagnostics))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *PublishDiagnosticsParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *PublishDiagnosticsParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, string(v.URI))
	enc.ArrayKey(keyDiagnostics, (*diagnostics)(&v.Diagnostics))
}

// IsNil returns wether the structure is nil value or not.
func (v *PublishDiagnosticsParams) IsNil() bool { return v == nil }
