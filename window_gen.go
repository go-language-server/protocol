// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ShowMessageParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "message":
		return dec.String(&v.Message)
	case "type":
		return dec.Float64((*float64)(&v.Type))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ShowMessageParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ShowMessageParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("message", v.Message)
	enc.Float64Key("type", float64(v.Type))
}

// IsNil returns wether the structure is nil value or not.
func (v *ShowMessageParams) IsNil() bool { return v == nil }

type actions []MessageActionItem

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *actions) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := MessageActionItem{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *actions) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *actions) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.ObjectOmitEmpty(&t)
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *actions) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ShowMessageRequestParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "actions":
		return dec.Array((*actions)(&v.Actions))
	case "message":
		return dec.String(&v.Message)
	case "type":
		return dec.Float64((*float64)(&v.Type))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ShowMessageRequestParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ShowMessageRequestParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("actions", (*actions)(&v.Actions))
	enc.StringKey("message", v.Message)
	enc.Float64Key("type", float64(v.Type))
}

// IsNil returns wether the structure is nil value or not.
func (v *ShowMessageRequestParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *MessageActionItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "title":
		return dec.String(&v.Title)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *MessageActionItem) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *MessageActionItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("title", v.Title)
}

// IsNil returns wether the structure is nil value or not.
func (v *MessageActionItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *LogMessageParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "message":
		return dec.String(&v.Message)
	case "type":
		return dec.Float64((*float64)(&v.Type))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *LogMessageParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *LogMessageParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("message", v.Message)
	enc.Float64Key("type", float64(v.Type))
}

// IsNil returns wether the structure is nil value or not.
func (v *LogMessageParams) IsNil() bool { return v == nil }
