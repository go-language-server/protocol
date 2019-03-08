// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Registration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "id":
		return dec.String(&v.ID)
	case "method":
		return dec.String(&v.Method)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Registration) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Registration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("id", v.ID)
	enc.StringKey("method", v.Method)
}

// IsNil returns wether the structure is nil value or not.
func (v *Registration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RegistrationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RegistrationParams) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *RegistrationParams) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not.
func (v *RegistrationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentRegistrationOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("documentSelector", v.DocumentSelector)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Unregistration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "id":
		return dec.String(&v.ID)
	case "method":
		return dec.String(&v.Method)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Unregistration) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Unregistration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("id", v.ID)
	enc.StringKey("method", v.Method)
}

// IsNil returns wether the structure is nil value or not.
func (v *Unregistration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *UnregistrationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *UnregistrationParams) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *UnregistrationParams) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not.
func (v *UnregistrationParams) IsNil() bool { return v == nil }
