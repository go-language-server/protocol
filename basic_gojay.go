// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"github.com/francoispqt/gojay"
	"github.com/go-language-server/uri"
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Position) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key(keyLine, v.Line)
	enc.Float64Key(keyCharacter, v.Character)
}

// IsNil returns wether the structure is nil value or not.
func (v *Position) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Position) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyLine:
		return dec.Float64(&v.Line)
	case keyCharacter:
		return dec.Float64(&v.Character)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Position) NKeys() int { return 2 }

// Reset reset fields.
func (v *Position) Reset() {
	v.Line = 0.0
	v.Character = 0.0
}

// compile time check whether the Position implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*Position)(nil)
var _ gojay.UnmarshalerJSONObject = (*Position)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Range) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyStart, &v.Start)
	enc.ObjectKey(keyEnd, &v.End)
}

// IsNil returns wether the structure is nil value or not.
func (v *Range) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Range) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyStart:
		return dec.Object(&v.Start)
	case keyEnd:
		return dec.Object(&v.End)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Range) NKeys() int { return 2 }

// Reset reset fields.
func (v *Range) Reset() {
	(&v.Start).Reset()
	LocationPool.Put(&v.Start)
	(&v.End).Reset()
	LocationPool.Put(&v.End)
}

// compile time check whether the Range implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*Range)(nil)
var _ gojay.UnmarshalerJSONObject = (*Range)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Location) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, string(v.URI))
	enc.ObjectKey(keyRange, &v.Range)
}

// IsNil returns wether the structure is nil value or not.
func (v *Location) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Location) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyRange:
		value := RangePool.Get().(*Range)
		err := dec.Object(value)
		if err == nil {
			v.Range = *value
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Location) NKeys() int { return 2 }

// Reset reset fields.
func (v *Location) Reset() {
	(&v.Range).Reset()
	RangePool.Put(&v.Range)
}

// compile time check whether the Location implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*Location)(nil)
var _ gojay.UnmarshalerJSONObject = (*Location)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *LocationLink) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyOriginSelectionRange, v.OriginSelectionRange)
	enc.StringKey(keyTargetURI, string(v.TargetURI))
	enc.ObjectKey(keyTargetRange, &v.TargetRange)
	enc.ObjectKey(keyTargetSelectionRange, &v.TargetSelectionRange)
}

// IsNil returns wether the structure is nil value or not.
func (v *LocationLink) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *LocationLink) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyOriginSelectionRange:
		value := RangePool.Get().(*Range)
		err := dec.Object(value)
		if err == nil {
			v.OriginSelectionRange = value
		}
		return err
	case keyTargetURI:
		return dec.String((*string)(&v.TargetURI))
	case keyTargetRange:
		value := RangePool.Get().(*Range)
		err := dec.Object(value)
		if err == nil {
			v.TargetRange = *value
		}
		return err
	case keyTargetSelectionRange:
		value := RangePool.Get().(*Range)
		err := dec.Object(value)
		if err == nil {
			v.TargetSelectionRange = *value
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *LocationLink) NKeys() int { return 4 }

// Reset reset fields.
func (v *LocationLink) Reset() {
	RangePool.Put(v.OriginSelectionRange)
	v.OriginSelectionRange = nil
	v.TargetURI = ""
	(&v.TargetRange).Reset()
	RangePool.Put(&v.TargetRange)
	(&v.TargetSelectionRange).Reset()
	RangePool.Put(&v.TargetSelectionRange)
}

// compile time check whether the LocationLink implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*LocationLink)(nil)
var _ gojay.UnmarshalerJSONObject = (*LocationLink)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Diagnostic) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyRange, &v.Range)
	enc.Float64KeyOmitEmpty(keySeverity, float64(v.Severity))
	enc.AddInterfaceKeyOmitEmpty(keyCode, v.Code)
	enc.StringKeyOmitEmpty(keySource, v.Source)
	enc.StringKey(keyMessage, v.Message)
	enc.ArrayKeyOmitEmpty(keyRelatedInformation, DiagnosticRelatedInformations(v.RelatedInformation))
}

// IsNil returns wether the structure is nil value or not.
func (v *Diagnostic) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Diagnostic) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRange:
		value := RangePool.Get().(*Range)
		err := dec.Object(value)
		if err == nil {
			v.Range = *value
		}
		return err
	case keySeverity:
		return dec.Float64((*float64)(&v.Severity))
	case keyCode:
		return dec.Interface(&v.Code)
	case keySource:
		return dec.String(&v.Source)
	case keyMessage:
		return dec.String(&v.Message)
	case keyRelatedInformation:
		values := DiagnosticRelatedInformations{}
		err := dec.Array(&values)
		if err == nil && len(values) > 0 {
			v.RelatedInformation = []DiagnosticRelatedInformation(values)
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Diagnostic) NKeys() int { return 6 }

// Reset reset fields.
func (v *Diagnostic) Reset() {
	(&v.Range).Reset()
	RangePool.Put(&v.Range)
	v.Severity = 0.0
	v.Code = nil
	v.Source = ""
	v.Message = ""
	for i := range v.RelatedInformation {
		v.RelatedInformation[i].Reset()
		DiagnosticRelatedInformationPool.Put(&v.RelatedInformation[i])
	}
	v.RelatedInformation = nil
}

// compile time check whether the Diagnostic implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*Diagnostic)(nil)
var _ gojay.UnmarshalerJSONObject = (*Diagnostic)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DiagnosticRelatedInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyLocation, &v.Location)
	enc.StringKey(keyMessage, v.Message)
}

// IsNil returns wether the structure is nil value or not.
func (v *DiagnosticRelatedInformation) IsNil() bool { return v == nil }

// Reset reset fields.
func (v *DiagnosticRelatedInformation) Reset() {
	(&v.Location).Reset()
	LocationPool.Put(&v.Location)
	v.Message = ""
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DiagnosticRelatedInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyLocation:
		return dec.Object(&v.Location)
	case keyMessage:
		return dec.String(&v.Message)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DiagnosticRelatedInformation) NKeys() int { return 2 }

// compile time check whether the DiagnosticRelatedInformation implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*DiagnosticRelatedInformation)(nil)
var _ gojay.UnmarshalerJSONObject = (*DiagnosticRelatedInformation)(nil)

// DiagnosticRelatedInformations represents a slice of DiagnosticRelatedInformation.
type DiagnosticRelatedInformations []DiagnosticRelatedInformation

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v DiagnosticRelatedInformations) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v DiagnosticRelatedInformations) IsNil() bool {
	return len(v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *DiagnosticRelatedInformations) UnmarshalJSONArray(dec *gojay.Decoder) error {
	value := DiagnosticRelatedInformation{}
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// compile time check whether the DiagnosticRelatedInformation implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interface.
var _ gojay.MarshalerJSONArray = (*DiagnosticRelatedInformations)(nil)
var _ gojay.UnmarshalerJSONArray = (*DiagnosticRelatedInformations)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Command) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyTitle, v.Title)
	enc.StringKey(keyCommand, v.Command)
	enc.ArrayKeyOmitEmpty(keyArguments, (*Interfaces)(&v.Arguments))
}

// IsNil returns wether the structure is nil value or not.
func (v *Command) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Command) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTitle:
		return dec.String(&v.Title)
	case keyCommand:
		return dec.String(&v.Command)
	case keyArguments:
		return dec.Array((*Interfaces)(&v.Arguments))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Command) NKeys() int { return 3 }

// Reset reset fields.
func (v *Command) Reset() {
	v.Title = ""
	v.Command = ""
	v.Arguments = nil
}

// compile time check whether the Command implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*Command)(nil)
var _ gojay.UnmarshalerJSONObject = (*Command)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextEdit) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyRange, &v.Range)
	enc.StringKey(keyNewText, v.NewText)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextEdit) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRange:
		value := RangePool.Get().(*Range)
		err := dec.Object(value)
		if err == nil {
			v.Range = *value
		}
		return err
	case keyNewText:
		return dec.String(&v.NewText)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextEdit) NKeys() int { return 2 }

// Reset reset fields.
func (v *TextEdit) Reset() {
	(&v.Range).Reset()
	RangePool.Put(&v.Range)
	v.NewText = ""
}

// compile time check whether the TextEdit implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*TextEdit)(nil)
var _ gojay.UnmarshalerJSONObject = (*TextEdit)(nil)

// TextEdits represents a slice of TextEdit.
type TextEdits []TextEdit

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v TextEdits) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// IsNil returns wether the structure is nil value or not.
func (v TextEdits) IsNil() bool {
	return len(v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *TextEdits) UnmarshalJSONArray(dec *gojay.Decoder) error {
	value := TextEdit{}
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// compile time check whether the TextEdits implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interface.
var _ gojay.MarshalerJSONArray = (*TextEdits)(nil)
var _ gojay.UnmarshalerJSONArray = (*TextEdits)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentEdit) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ArrayKey(keyEdits, (*TextEdits)(&v.Edits))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentEdit) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyEdits:
		return dec.Array((*TextEdits)(&v.Edits))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentEdit) NKeys() int { return 2 }

// Reset reset fields.
func (v *TextDocumentEdit) Reset() {
	for i := range v.Edits {
		v.Edits[i].Reset()
		TextEditPool.Put(&v.Edits[i])
	}
	v.Edits = nil
}

// compile time check whether the TextDocumentEdit implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*TextDocumentEdit)(nil)
var _ gojay.UnmarshalerJSONObject = (*TextDocumentEdit)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CreateFileOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyOverwrite, v.Overwrite)
	enc.BoolKeyOmitEmpty(keyIgnoreIfExists, v.IgnoreIfExists)
}

// IsNil returns wether the structure is nil value or not.
func (v *CreateFileOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CreateFileOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyOverwrite:
		return dec.Bool(&v.Overwrite)
	case keyIgnoreIfExists:
		return dec.Bool(&v.IgnoreIfExists)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CreateFileOptions) NKeys() int { return 2 }

// Reset reset fields.
func (v *CreateFileOptions) Reset() {
	v.Overwrite = false
	v.IgnoreIfExists = false
}

// compile time check whether the CreateFileOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*CreateFileOptions)(nil)
var _ gojay.UnmarshalerJSONObject = (*CreateFileOptions)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CreateFile) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyKind, string(v.Kind))
	enc.StringKey(keyURI, string(v.URI))
	enc.ObjectKeyOmitEmpty(keyOptions, v.Options)
}

// IsNil returns wether the structure is nil value or not.
func (v *CreateFile) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CreateFile) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyKind:
		return dec.String((*string)(&v.Kind))
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyOptions:
		value := CreateFileOptionsPool.Get().(*CreateFileOptions)
		err := dec.Object(value)
		if err == nil {
			v.Options = value
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CreateFile) NKeys() int { return 3 }

// Reset reset fields.
func (v *CreateFile) Reset() {
	v.Kind = ""
	v.URI = ""
	CreateFileOptionsPool.Put(v.Options)
	v.Options = nil
}

// compile time check whether the CreateFile implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*CreateFile)(nil)
var _ gojay.UnmarshalerJSONObject = (*CreateFile)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *RenameFileOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyOverwrite, v.Overwrite)
	enc.BoolKeyOmitEmpty(keyIgnoreIfExists, v.IgnoreIfExists)
}

// IsNil returns wether the structure is nil value or not.
func (v *RenameFileOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RenameFileOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyOverwrite:
		return dec.Bool(&v.Overwrite)
	case keyIgnoreIfExists:
		return dec.Bool(&v.IgnoreIfExists)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RenameFileOptions) NKeys() int { return 2 }

// Reset reset fields.
func (v *RenameFileOptions) Reset() {
	v.Overwrite = false
	v.IgnoreIfExists = false
}

// compile time check whether the RenameFileOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*RenameFileOptions)(nil)
var _ gojay.UnmarshalerJSONObject = (*RenameFileOptions)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *RenameFile) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyKind, string(v.Kind))
	enc.StringKey(keyOldURI, string(v.OldURI))
	enc.StringKey(keyNewURI, string(v.NewURI))
	enc.ObjectKeyOmitEmpty(keyOptions, v.Options)
}

// IsNil returns wether the structure is nil value or not.
func (v *RenameFile) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RenameFile) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyKind:
		return dec.String((*string)(&v.Kind))
	case keyOldURI:
		return dec.String((*string)(&v.OldURI))
	case keyNewURI:
		return dec.String((*string)(&v.NewURI))
	case keyOptions:
		value := RenameFileOptionsPool.Get().(*RenameFileOptions)
		err := dec.Object(value)
		if err == nil {
			v.Options = value
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RenameFile) NKeys() int { return 4 }

// Reset reset fields.
func (v *RenameFile) Reset() {
	v.Kind = ""
	v.OldURI = ""
	v.NewURI = ""
	RenameFileOptionsPool.Put(v.Options)
	v.Options = nil
}

// compile time check whether the RenameFile implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*RenameFile)(nil)
var _ gojay.UnmarshalerJSONObject = (*RenameFile)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DeleteFileOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyRecursive, v.Recursive)
	enc.BoolKeyOmitEmpty(keyIgnoreIfNotExists, v.IgnoreIfNotExists)
}

// IsNil returns wether the structure is nil value or not.
func (v *DeleteFileOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DeleteFileOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRecursive:
		return dec.Bool(&v.Recursive)
	case keyIgnoreIfNotExists:
		return dec.Bool(&v.IgnoreIfNotExists)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DeleteFileOptions) NKeys() int { return 2 }

// Reset reset fields.
func (v *DeleteFileOptions) Reset() {
	v.Recursive = false
	v.IgnoreIfNotExists = false
}

// compile time check whether the DeleteFileOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*DeleteFileOptions)(nil)
var _ gojay.UnmarshalerJSONObject = (*DeleteFileOptions)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DeleteFile) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyKind, string(v.Kind))
	enc.StringKey(keyURI, string(v.URI))
	enc.ObjectKeyOmitEmpty(keyOptions, v.Options)
}

// IsNil returns wether the structure is nil value or not.
func (v *DeleteFile) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DeleteFile) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyKind:
		return dec.String((*string)(&v.Kind))
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyOptions:
		value := DeleteFileOptionsPool.Get().(*DeleteFileOptions)
		err := dec.Object(value)
		if err == nil {
			v.Options = value
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DeleteFile) NKeys() int { return 3 }

// Reset reset fields.
func (v *DeleteFile) Reset() {
	v.Kind = ""
	v.URI = ""
	DeleteFileOptionsPool.Put(v.Options)
	v.Options = nil
}

// compile time check whether the DeleteFile implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*DeleteFile)(nil)
var _ gojay.UnmarshalerJSONObject = (*DeleteFile)(nil)

// TextEditsMap represents a map of WorkspaceEdit.Changes.
type TextEditsMap map[uri.URI][]TextEdit

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v TextEditsMap) MarshalJSONObject(enc *gojay.Encoder) {
	for key, value := range v {
		enc.ArrayKeyOmitEmpty(string(key), (*TextEdits)(&value))
	}
}

// IsNil returns wether the structure is nil value or not.
func (v TextEditsMap) IsNil() bool {
	return v == nil
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v TextEditsMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	edits := []TextEdit{}
	err := dec.Array((*TextEdits)(&edits))
	if err != nil {
		return err
	}
	v[uri.URI(k)] = TextEdits(edits)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v TextEditsMap) NKeys() int { return 0 }

type documentChanges []TextDocumentEdit

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v documentChanges) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.ObjectOmitEmpty(&v[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v documentChanges) IsNil() bool {
	return v == nil || len(v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *documentChanges) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := TextDocumentEdit{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// compile time check whether the documentChanges implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interface.
var _ gojay.MarshalerJSONArray = (*documentChanges)(nil)
var _ gojay.UnmarshalerJSONArray = (*documentChanges)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceEdit) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyChanges, (*TextEditsMap)(&v.Changes))
	enc.ArrayKeyOmitEmpty(keyDocumentChanges, (*documentChanges)(&v.DocumentChanges))
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceEdit) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyChanges:
		if v.Changes == nil {
			v.Changes = make(map[uri.URI][]TextEdit)
		}
		return dec.Object(TextEditsMap(v.Changes))
	case keyDocumentChanges:
		if v.DocumentChanges == nil {
			v.DocumentChanges = []TextDocumentEdit{}
		}
		return dec.Array((*documentChanges)(&v.DocumentChanges))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceEdit) NKeys() int { return 2 }

// Reset reset fields.
func (v *WorkspaceEdit) Reset() {
	for k := range v.Changes {
		for i := range v.Changes[k] {
			v.Changes[k][i].Reset()
			TextEditPool.Put(&v.Changes[k][i])
		}
	}
	for i := range v.DocumentChanges {
		v.DocumentChanges[i].Reset()
		TextDocumentEditPool.Put(&v.DocumentChanges[i])
	}
}

// compile time check whether the WorkspaceEdit implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*WorkspaceEdit)(nil)
var _ gojay.UnmarshalerJSONObject = (*WorkspaceEdit)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentIdentifier) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, string(v.URI))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentIdentifier) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentIdentifier) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyURI {
		return dec.String((*string)(&v.URI))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentIdentifier) NKeys() int { return 1 }

// Reset reset fields.
func (v *TextDocumentIdentifier) Reset() {
	v.URI = uri.URI("")
}

// compile time check whether the TextDocumentIdentifier implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*TextDocumentIdentifier)(nil)
var _ gojay.UnmarshalerJSONObject = (*TextDocumentIdentifier)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, string(v.URI))
	enc.StringKey(keyLanguageID, string(v.LanguageID))
	enc.Float64Key(keyVersion, v.Version)
	enc.StringKey(keyText, v.Text)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyLanguageID:
		return dec.String((*string)(&v.LanguageID))
	case keyVersion:
		return dec.Float64(&v.Version)
	case keyText:
		return dec.String(&v.Text)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentItem) NKeys() int { return 4 }

// Reset reset fields.
func (v *TextDocumentItem) Reset() {
	v.URI = uri.URI("")
	v.LanguageID = LanguageIdentifier("")
	v.Version = 0.0
	v.Text = ""
}

// compile time check whether the TextDocumentItem implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*TextDocumentItem)(nil)
var _ gojay.UnmarshalerJSONObject = (*TextDocumentItem)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *VersionedTextDocumentIdentifier) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, string(v.URI))
	enc.Uint64KeyNullEmpty(keyVersion, *v.Version)
}

// IsNil returns wether the structure is nil value or not.
func (v *VersionedTextDocumentIdentifier) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *VersionedTextDocumentIdentifier) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyVersion:
		version := &v.Version
		return dec.Uint64Null(version)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *VersionedTextDocumentIdentifier) NKeys() int { return 2 }

// Reset reset fields.
func (v *VersionedTextDocumentIdentifier) Reset() {
	v.URI = uri.URI("")
	v.Version = nil
}

// compile time check whether the VersionedTextDocumentIdentifier implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*VersionedTextDocumentIdentifier)(nil)
var _ gojay.UnmarshalerJSONObject = (*VersionedTextDocumentIdentifier)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentPositionParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ObjectKey(keyPosition, &v.Position)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentPositionParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentPositionParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentPositionParams) NKeys() int { return 2 }

// Reset reset fields.
func (v *TextDocumentPositionParams) Reset() {
	v.TextDocument.Reset()
	(&v.Position).Reset()
	PositionPool.Put(&v.Position)
}

// compile time check whether the TextDocumentPositionParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*TextDocumentPositionParams)(nil)
var _ gojay.UnmarshalerJSONObject = (*TextDocumentPositionParams)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentFilter) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty(keyLanguage, v.Language)
	enc.StringKeyOmitEmpty(keyScheme, v.Scheme)
	enc.StringKeyOmitEmpty(keyPattern, v.Pattern)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentFilter) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentFilter) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyLanguage:
		return dec.String(&v.Language)
	case keyScheme:
		return dec.String(&v.Scheme)
	case keyPattern:
		return dec.String(&v.Pattern)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentFilter) NKeys() int { return 3 }

// Reset reset fields.
func (v *DocumentFilter) Reset() {
	v.Language = ""
	v.Scheme = ""
	v.Pattern = ""
}

// compile time check whether the DocumentFilter implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*DocumentFilter)(nil)
var _ gojay.UnmarshalerJSONObject = (*DocumentFilter)(nil)

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v DocumentSelector) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(v[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v DocumentSelector) IsNil() bool {
	return v == nil || len(v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *DocumentSelector) UnmarshalJSONArray(dec *gojay.Decoder) error {
	value := &DocumentFilter{}
	if err := dec.Object(value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// Reset reset fields.
func (v *DocumentSelector) Reset() {
	values := *v
	for i := range values {
		values[i].Reset()
		DocumentFilterPool.Put(&values[i])
	}
}

// compile time check whether the DocumentSelector implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interface.
var _ gojay.MarshalerJSONArray = (*DocumentSelector)(nil)
var _ gojay.UnmarshalerJSONArray = (*DocumentSelector)(nil)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *MarkupContent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyKind, string(v.Kind))
	enc.StringKey(keyValue, v.Value)
}

// IsNil returns wether the structure is nil value or not.
func (v *MarkupContent) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *MarkupContent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyKind:
		return dec.String((*string)(&v.Kind))
	case keyValue:
		return dec.String(&v.Value)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *MarkupContent) NKeys() int { return 2 }

// Reset reset fields.
func (v *MarkupContent) Reset() {
	v.Kind = MarkupKind("")
	v.Value = ""
}

// compile time check whether the MarkupContent implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interface.
var _ gojay.MarshalerJSONObject = (*MarkupContent)(nil)
var _ gojay.UnmarshalerJSONObject = (*MarkupContent)(nil)
