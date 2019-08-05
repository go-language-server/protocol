// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidOpenTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTextDocument {
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidOpenTextDocumentParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidOpenTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DidOpenTextDocumentParams) IsNil() bool { return v == nil }

type textDocumentContentChangeEvents []TextDocumentContentChangeEvent

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *textDocumentContentChangeEvents) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := TextDocumentContentChangeEvent{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *textDocumentContentChangeEvents) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *textDocumentContentChangeEvents) MarshalJSONArray(enc *gojay.Encoder) {
	vv := *v
	for i := range vv {
		enc.ObjectOmitEmpty(&vv[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *textDocumentContentChangeEvents) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidChangeTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyContentChanges:
		return dec.Array((*textDocumentContentChangeEvents)(&v.ContentChanges))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidChangeTextDocumentParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidChangeTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ArrayKey(keyContentChanges, (*textDocumentContentChangeEvents)(&v.ContentChanges))
}

// IsNil returns wether the structure is nil value or not.
func (v *DidChangeTextDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocument) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyLanguageID:
		return dec.String(&v.LanguageID)
	case keyVersion:
		return dec.Float64(&v.Version)
	case keyLineCount:
		return dec.Float64(&v.LineCount)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocument) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocument) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, string(v.URI))
	enc.StringKey(keyLanguageID, v.LanguageID)
	enc.Float64Key(keyVersion, v.Version)
	enc.Float64Key(keyLineCount, v.LineCount)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocument) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentChangeEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDocument {
		return dec.Object(&v.Document)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentChangeEvent) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentChangeEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyDocument, &v.Document)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentChangeEvent) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentWillSaveEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocument:
		return dec.Object(&v.Document)
	case keyReason:
		return dec.Float64((*float64)(&v.Reason))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentWillSaveEvent) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentWillSaveEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyDocument, &v.Document)
	enc.Float64Key(keyReason, float64(v.Reason))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentWillSaveEvent) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentContentChangeEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRange:
		if v.Range == nil {
			v.Range = &Range{}
		}
		return dec.Object(v.Range)
	case keyRangeLength:
		return dec.Float64(&v.RangeLength)
	case keyText:
		return dec.String(&v.Text)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentContentChangeEvent) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentContentChangeEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyRange, v.Range)
	enc.Float64KeyOmitEmpty(keyRangeLength, v.RangeLength)
	enc.StringKey(keyText, v.Text)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentContentChangeEvent) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentChangeRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		return dec.Array(&v.DocumentSelector)
	case keySyncKind:
		return dec.Float64(&v.SyncKind)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentChangeRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentChangeRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.Float64Key(keySyncKind, v.SyncKind)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentChangeRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WillSaveTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyReason:
		return dec.Float64((*float64)(&v.Reason))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WillSaveTextDocumentParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WillSaveTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.Float64Key(keyReason, float64(v.Reason))
}

// IsNil returns wether the structure is nil value or not.
func (v *WillSaveTextDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidSaveTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyText:
		return dec.String(&v.Text)
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidSaveTextDocumentParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidSaveTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty(keyText, v.Text)
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DidSaveTextDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentSaveRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		return dec.Array(&v.DocumentSelector)
	case keyIncludeText:
		return dec.Bool(&v.IncludeText)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentSaveRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentSaveRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyIncludeText, v.IncludeText)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentSaveRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidCloseTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTextDocument {
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidCloseTextDocumentParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidCloseTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DidCloseTextDocumentParams) IsNil() bool { return v == nil }
