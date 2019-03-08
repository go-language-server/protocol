// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *PublishDiagnosticsParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "uri":
		return dec.String((*string)(&v.URI))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *PublishDiagnosticsParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *PublishDiagnosticsParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("uri", string(v.URI))
}

// IsNil returns wether the structure is nil value or not
func (v *PublishDiagnosticsParams) IsNil() bool { return v == nil }
