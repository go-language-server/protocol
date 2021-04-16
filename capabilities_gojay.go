// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"github.com/francoispqt/gojay"
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *RegularExpressionsClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyEngine, v.Engine)
	enc.StringKeyOmitEmpty(keyVersion, v.Version)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *RegularExpressionsClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *RegularExpressionsClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyEngine:
		return dec.String(&v.Engine)
	case keyVersion:
		return dec.String(&v.Version)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *RegularExpressionsClientCapabilities) NKeys() int { return 2 }

// compile time check whether the RegularExpressionsClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*RegularExpressionsClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*RegularExpressionsClientCapabilities)(nil)
)
