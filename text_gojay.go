// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import "github.com/francoispqt/gojay"

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidOpenTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DidOpenTextDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidOpenTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTextDocument {
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidOpenTextDocumentParams) NKeys() int { return 1 }

// compile time check whether the DidOpenTextDocumentParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*DidOpenTextDocumentParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*DidOpenTextDocumentParams)(nil)
	// _ Pooler                      = (*DidOpenTextDocumentParams)(nil)
)

type textDocumentContentChangeEvents []TextDocumentContentChangeEvent

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v textDocumentContentChangeEvents) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.ObjectOmitEmpty(&v[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v textDocumentContentChangeEvents) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *textDocumentContentChangeEvents) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := TextDocumentContentChangeEvent{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// compile time check whether the textDocumentContentChangeEvents implements a gojay.MarshalerJSONArray, gojay.UnmarshalerJSONArray and Pooler interface.
var (
	_ gojay.MarshalerJSONArray   = (*textDocumentContentChangeEvents)(nil)
	_ gojay.UnmarshalerJSONArray = (*textDocumentContentChangeEvents)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidChangeTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ArrayKey(keyContentChanges, (*textDocumentContentChangeEvents)(&v.ContentChanges))
}

// IsNil returns wether the structure is nil value or not.
func (v *DidChangeTextDocumentParams) IsNil() bool { return v == nil }

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

// compile time check whether the DidChangeTextDocumentParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*DidChangeTextDocumentParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*DidChangeTextDocumentParams)(nil)
	// _ Pooler                      = (*DidChangeTextDocumentParams)(nil)
)

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

// compile time check whether the TextDocument implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocument)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocument)(nil)
	// _ Pooler                      = (*TextDocument)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentChangeEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyDocument, &v.Document)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentChangeEvent) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentChangeEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDocument {
		return dec.Object(&v.Document)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentChangeEvent) NKeys() int { return 1 }

// compile time check whether the TextDocumentChangeEvent implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentChangeEvent)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentChangeEvent)(nil)
	// _ Pooler                      = (*TextDocumentChangeEvent)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentWillSaveEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyDocument, &v.Document)
	enc.Float64Key(keyReason, float64(v.Reason))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentWillSaveEvent) IsNil() bool { return v == nil }

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

// compile time check whether the TextDocumentWillSaveEvent implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentWillSaveEvent)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentWillSaveEvent)(nil)
	// _ Pooler                      = (*TextDocumentWillSaveEvent)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentContentChangeEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyRange, v.Range)
	enc.Float64KeyOmitEmpty(keyRangeLength, v.RangeLength)
	enc.StringKey(keyText, v.Text)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentContentChangeEvent) IsNil() bool { return v == nil }

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

// compile time check whether the TextDocumentContentChangeEvent implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentContentChangeEvent)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentContentChangeEvent)(nil)
	// _ Pooler                      = (*TextDocumentContentChangeEvent)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentChangeRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.Float64Key(keySyncKind, v.SyncKind)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentChangeRegistrationOptions) IsNil() bool { return v == nil }

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

// compile time check whether the TextDocumentChangeRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentChangeRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentChangeRegistrationOptions)(nil)
	// _ Pooler                      = (*TextDocumentChangeRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WillSaveTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.Float64Key(keyReason, float64(v.Reason))
}

// IsNil returns wether the structure is nil value or not.
func (v *WillSaveTextDocumentParams) IsNil() bool { return v == nil }

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

// compile time check whether the WillSaveTextDocumentParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*WillSaveTextDocumentParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*WillSaveTextDocumentParams)(nil)
	// _ Pooler                      = (*WillSaveTextDocumentParams)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidSaveTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty(keyText, v.Text)
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DidSaveTextDocumentParams) IsNil() bool { return v == nil }

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

// compile time check whether the DidSaveTextDocumentParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*DidSaveTextDocumentParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*DidSaveTextDocumentParams)(nil)
	// _ Pooler                      = (*DidSaveTextDocumentParams)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentSaveRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyIncludeText, v.IncludeText)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentSaveRegistrationOptions) IsNil() bool { return v == nil }

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

// compile time check whether the TextDocumentSaveRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentSaveRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentSaveRegistrationOptions)(nil)
	// _ Pooler                      = (*TextDocumentSaveRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidCloseTextDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DidCloseTextDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidCloseTextDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTextDocument {
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidCloseTextDocumentParams) NKeys() int { return 1 }

// compile time check whether the DidCloseTextDocumentParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*DidCloseTextDocumentParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*DidCloseTextDocumentParams)(nil)
	// _ Pooler                      = (*DidCloseTextDocumentParams)(nil)
)
