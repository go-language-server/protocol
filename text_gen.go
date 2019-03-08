// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DidOpenTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DidOpenTextDocumentParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DidOpenTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not
func (v *DidOpenTextDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DidChangeTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DidChangeTextDocumentParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DidChangeTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not
func (v *DidChangeTextDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentContentChangeEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "range":
		return dec.Object(v.Range)
	case "rangeLength":
		return dec.Float64(&v.RangeLength)
	case "text":
		return dec.String(&v.Text)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentContentChangeEvent) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentContentChangeEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("range", v.Range)
	enc.Float64Key("rangeLength", v.RangeLength)
	enc.StringKey("text", v.Text)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentContentChangeEvent) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentChangeRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	case "syncKind":
		return dec.Float64(&v.SyncKind)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentChangeRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentChangeRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("documentSelector", v.DocumentSelector)
	enc.Float64Key("syncKind", v.SyncKind)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentChangeRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WillSaveTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(&v.TextDocument)
	case "reason":
		return dec.Float64((*float64)(&v.Reason))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WillSaveTextDocumentParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WillSaveTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", &v.TextDocument)
	enc.Float64Key("reason", float64(v.Reason))
}

// IsNil returns wether the structure is nil value or not
func (v *WillSaveTextDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DidSaveTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "text":
		return dec.String(&v.Text)
	case "textDocument":
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DidSaveTextDocumentParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DidSaveTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("text", v.Text)
	enc.ObjectKey("textDocument", &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not
func (v *DidSaveTextDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentSaveRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	case "includeText":
		return dec.Bool(&v.IncludeText)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentSaveRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentSaveRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("documentSelector", v.DocumentSelector)
	enc.BoolKey("includeText", v.IncludeText)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentSaveRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DidCloseTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DidCloseTextDocumentParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DidCloseTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not
func (v *DidCloseTextDocumentParams) IsNil() bool { return v == nil }
