// Copyright 2021 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gojay
// +build gojay

package protocol

import (
	"github.com/francoispqt/gojay"
)

// ProgressToken is the progress token provided by the client or server.
//
// @since 3.15.0.
type ProgressToken interface{}

// NewProgressToken returns a new ProgressToken.
//nolint:gocritic
func NewProgressToken(s string) *ProgressToken {
	var iface interface{} = s
	return (*ProgressToken)(&iface)
}

// NewNumberProgressToken returns a new number ProgressToken.
//nolint:gocritic
func NewNumberProgressToken(n int64) *ProgressToken {
	var iface interface{} = n
	return (*ProgressToken)(&iface)
}

//nolint:gocritic
func encodeProgressToken(enc *gojay.Encoder, key string, v *ProgressToken) {
	if v == nil {
		return
	}

	enc.AddInterfaceKey(key, (interface{})(*v))
}

//nolint:gocritic
func decodeProgressToken(dec *gojay.Decoder, k, key string, v *ProgressToken) error {
	if v == nil || k != key {
		return nil
	}

	vv := (interface{})(*v)
	return dec.Interface(&vv)
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ProgressParams) MarshalJSONObject(enc *gojay.Encoder) {
	encodeProgressToken(enc, keyToken, &v.Token)
	enc.AddInterfaceKey(keyValue, v.Value)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ProgressParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ProgressParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyToken:
		return decodeProgressToken(dec, k, keyToken, &v.Token)
	case keyValue:
		return dec.Interface(&v.Value)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ProgressParams) NKeys() int { return 2 }

// compile time check whether the ProgressParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ProgressParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*ProgressParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkDoneProgressBegin) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyKind, string(v.Kind))
	enc.StringKey(keyTitle, v.Title)
	enc.BoolKeyOmitEmpty(keyCancellable, v.Cancellable)
	enc.StringKeyOmitEmpty(keyMessage, v.Message)
	enc.Float64KeyOmitEmpty(keyPercentage, v.Percentage)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkDoneProgressBegin) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkDoneProgressBegin) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyKind:
		return dec.String((*string)(&v.Kind))
	case keyTitle:
		return dec.String(&v.Title)
	case keyCancellable:
		return dec.Bool(&v.Cancellable)
	case keyMessage:
		return dec.String(&v.Message)
	case keyPercentage:
		return dec.Float(&v.Percentage)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WorkDoneProgressBegin) NKeys() int { return 5 }

// compile time check whether the WorkDoneProgressBegin implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkDoneProgressBegin)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkDoneProgressBegin)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkDoneProgressReport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyKind, string(v.Kind))
	enc.BoolKeyOmitEmpty(keyCancellable, v.Cancellable)
	enc.StringKeyOmitEmpty(keyMessage, v.Message)
	enc.Float64KeyOmitEmpty(keyPercentage, v.Percentage)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkDoneProgressReport) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkDoneProgressReport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyKind:
		return dec.String((*string)(&v.Kind))
	case keyCancellable:
		return dec.Bool(&v.Cancellable)
	case keyMessage:
		return dec.String(&v.Message)
	case keyPercentage:
		return dec.Float(&v.Percentage)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WorkDoneProgressReport) NKeys() int { return 4 }

// compile time check whether the WorkDoneProgressReport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkDoneProgressReport)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkDoneProgressReport)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkDoneProgressEnd) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyKind, string(v.Kind))
	enc.StringKeyOmitEmpty(keyMessage, v.Message)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkDoneProgressEnd) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkDoneProgressEnd) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyKind:
		return dec.String((*string)(&v.Kind))
	case keyMessage:
		return dec.String(&v.Message)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WorkDoneProgressEnd) NKeys() int { return 2 }

// compile time check whether the WorkDoneProgressEnd implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkDoneProgressEnd)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkDoneProgressEnd)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkDoneProgressParams) MarshalJSONObject(enc *gojay.Encoder) {
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkDoneProgressParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkDoneProgressParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WorkDoneProgressParams) NKeys() int { return 1 }

// compile time check whether the WorkDoneProgressParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkDoneProgressParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkDoneProgressParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *PartialResultParams) MarshalJSONObject(enc *gojay.Encoder) {
	encodeProgressToken(enc, keyPartialResultToken, v.PartialResultToken)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *PartialResultParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *PartialResultParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return decodeProgressToken(dec, k, keyPartialResultToken, v.PartialResultToken)
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *PartialResultParams) NKeys() int { return 1 }

// compile time check whether the PartialResultParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*PartialResultParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*PartialResultParams)(nil)
)
