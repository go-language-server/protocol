// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import "github.com/francoispqt/gojay"

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *Registration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyID, v.ID)
	enc.StringKey(keyMethod, v.Method)
	enc.AddInterfaceKeyOmitEmpty(keyRegisterOptions, &v.RegisterOptions)
}

// IsNil returns wether the structure is nil value or not.
func (v *Registration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Registration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyID:
		return dec.String(&v.ID)
	case keyMethod:
		return dec.String(&v.Method)
	case keyRegisterOptions:
		return dec.Interface(&v.RegisterOptions)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Registration) NKeys() int { return 3 }

// compile time check whether the Registration implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*Registration)(nil)
	_ gojay.UnmarshalerJSONObject = (*Registration)(nil)
)

type registrations []Registration

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v registrations) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.ObjectOmitEmpty(&v[i])
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v registrations) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *registrations) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := Registration{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// compile time check whether the registrations implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*registrations)(nil)
	_ gojay.UnmarshalerJSONArray = (*registrations)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *RegistrationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyRegistrations, (*registrations)(&v.Registrations))
}

// IsNil returns wether the structure is nil value or not.
func (v *RegistrationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RegistrationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyRegistrations {
		return dec.Array((*registrations)(&v.Registrations))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RegistrationParams) NKeys() int { return 1 }

// compile time check whether the RegistrationParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*RegistrationParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*RegistrationParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDocumentSelector {
		return dec.Array(&v.DocumentSelector)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentRegistrationOptions) NKeys() int { return 1 }

// compile time check whether the TextDocumentRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *Unregistration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyID, v.ID)
	enc.StringKey(keyMethod, v.Method)
}

// IsNil returns wether the structure is nil value or not.
func (v *Unregistration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Unregistration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyID:
		return dec.String(&v.ID)
	case keyMethod:
		return dec.String(&v.Method)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Unregistration) NKeys() int { return 2 }

// compile time check whether the Unregistration implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*Unregistration)(nil)
	_ gojay.UnmarshalerJSONObject = (*Unregistration)(nil)
)

type unregisterations []Unregistration

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v unregisterations) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.ObjectOmitEmpty(&v[i])
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v unregisterations) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *unregisterations) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := Unregistration{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// compile time check whether the unregisterations implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*unregisterations)(nil)
	_ gojay.UnmarshalerJSONArray = (*unregisterations)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *UnregistrationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyUnregisterations, (*unregisterations)(&v.Unregisterations))
}

// IsNil returns wether the structure is nil value or not.
func (v *UnregistrationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *UnregistrationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyUnregisterations {
		return dec.Array((*unregisterations)(&v.Unregisterations))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *UnregistrationParams) NKeys() int { return 1 }

// compile time check whether the UnregistrationParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*UnregistrationParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*UnregistrationParams)(nil)
)
