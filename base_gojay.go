// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"github.com/francoispqt/gojay"
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *CancelParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddInterfaceKey(keyID, v.ID)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *CancelParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *CancelParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyID {
		return dec.Interface(&v.ID)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *CancelParams) NKeys() int { return 1 }

// compile time check whether the CancelParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*CancelParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*CancelParams)(nil)
)
