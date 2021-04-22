// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gojay
// +build gojay

package protocol

import (
	"github.com/francoispqt/gojay"
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v WorkspaceFolders) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v WorkspaceFolders) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *WorkspaceFolders) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value WorkspaceFolder
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// compile time check whether the WorkspaceFolders implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*WorkspaceFolders)(nil)
	_ gojay.UnmarshalerJSONArray = (*WorkspaceFolders)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ClientInfo) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyName, v.Name)
	enc.StringKeyOmitEmpty(keyVersion, v.Version)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ClientInfo) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ClientInfo) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyName:
		return dec.String(&v.Name)
	case keyVersion:
		return dec.String(&v.Version)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ClientInfo) NKeys() int { return 2 }

// compile time check whether the ClientInfo implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ClientInfo)(nil)
	_ gojay.UnmarshalerJSONObject = (*ClientInfo)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *InitializeParams) MarshalJSONObject(enc *gojay.Encoder) {
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
	enc.Int32KeyNullEmpty(keyProcessID, v.ProcessID)
	enc.ObjectKeyOmitEmpty(keyClientInfo, v.ClientInfo)
	enc.StringKeyOmitEmpty(keyLocale, v.Locale)
	enc.StringKeyOmitEmpty(keyRootPath, v.RootPath)
	enc.StringKeyNullEmpty(keyRootURI, string(v.RootURI))
	enc.AddInterfaceKeyOmitEmpty(keyInitializationOptions, v.InitializationOptions)
	enc.ObjectKey(keyCapabilities, &v.Capabilities)
	enc.StringKeyOmitEmpty(keyTrace, string(v.Trace))
	enc.ArrayKeyOmitEmpty(keyWorkspaceFolders, WorkspaceFolders(v.WorkspaceFolders))
}

// IsNil returns wether the structure is nil value or not.
func (v *InitializeParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializeParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyWorkDoneToken:
		return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
	case keyProcessID:
		processID := &v.ProcessID
		return dec.Int32Null(&processID)
	case keyClientInfo:
		if v.ClientInfo == nil {
			v.ClientInfo = &ClientInfo{}
		}
		return dec.Object(v.ClientInfo)
	case keyLocale:
		return dec.String(&v.Locale)
	case keyRootPath:
		return dec.String(&v.RootPath)
	case keyRootURI:
		s := (*string)(&v.RootURI)
		return dec.StringNull(&s)
	case keyInitializationOptions:
		return dec.Interface(&v.InitializationOptions)
	case keyCapabilities:
		return dec.Object(&v.Capabilities)
	case keyTrace:
		return dec.String((*string)(&v.Trace))
	case keyWorkspaceFolders:
		return dec.Array((*WorkspaceFolders)(&v.WorkspaceFolders))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializeParams) NKeys() int { return 10 }

// compile time check whether the InitializeParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*InitializeParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*InitializeParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *LogTraceParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyMessage, v.Message)
	enc.StringKeyOmitEmpty(keyVerbose, string(v.Verbose))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *LogTraceParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *LogTraceParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyMessage:
		return dec.String(&v.Message)
	case keyVerbose:
		return dec.String((*string)(&v.Verbose))
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *LogTraceParams) NKeys() int { return 2 }

// compile time check whether the LogTraceParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*LogTraceParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*LogTraceParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *SetTraceParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyValue, string(v.Value))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *SetTraceParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *SetTraceParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValue {
		return dec.String((*string)(&v.Value))
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *SetTraceParams) NKeys() int { return 1 }

// compile time check whether the SetTraceParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*SetTraceParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*SetTraceParams)(nil)
)

// FileOperationFilters represents a slice of FileOperationFilter.
type FileOperationFilters []FileOperationFilter

// compile time check whether the FileOperationFilters implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*FileOperationFilters)(nil)
	_ gojay.UnmarshalerJSONArray = (*FileOperationFilters)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v FileOperationFilters) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v FileOperationFilters) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *FileOperationFilters) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value FileOperationFilter
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *FileOperationRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyFilters, FileOperationFilters(v.Filters))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *FileOperationRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *FileOperationRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyFilters {
		return dec.Array((*FileOperationFilters)(&v.Filters))
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *FileOperationRegistrationOptions) NKeys() int { return 1 }

// compile time check whether the FileOperationRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*FileOperationRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*FileOperationRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *FileOperationPatternOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyIgnoreCase, v.IgnoreCase)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *FileOperationPatternOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *FileOperationPatternOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyIgnoreCase {
		return dec.Bool(&v.IgnoreCase)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *FileOperationPatternOptions) NKeys() int { return 1 }

// compile time check whether the FileOperationPatternOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*FileOperationPatternOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*FileOperationPatternOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *FileOperationPattern) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyGlob, v.Glob)
	enc.StringKeyOmitEmpty(keyMatches, string(v.Matches))
	enc.ObjectKeyOmitEmpty(keyOptions, &v.Options)
}

// IsNil returns wether the structure is nil value or not.
func (v *FileOperationPattern) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FileOperationPattern) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyGlob:
		return dec.String(&v.Glob)
	case keyMatches:
		return dec.String((*string)(&v.Matches))
	case keyOptions:
		return dec.Object(&v.Options)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FileOperationPattern) NKeys() int { return 3 }

// compile time check whether the FileOperationPattern implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*FileOperationPattern)(nil)
	_ gojay.UnmarshalerJSONObject = (*FileOperationPattern)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *FileOperationFilter) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty(keyScheme, v.Scheme)
	enc.ObjectKey(keyPattern, &v.Pattern)
}

// IsNil returns wether the structure is nil value or not.
func (v *FileOperationFilter) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FileOperationFilter) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyScheme:
		return dec.String(&v.Scheme)
	case keyPattern:
		return dec.Object(&v.Pattern)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FileOperationFilter) NKeys() int { return 2 }

// compile time check whether the FileOperationFilter implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*FileOperationFilter)(nil)
	_ gojay.UnmarshalerJSONObject = (*FileOperationFilter)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *CreateFilesParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyFiles, FileCreates(v.Files))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *CreateFilesParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *CreateFilesParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyFiles {
		return dec.Array((*FileCreates)(&v.Files))
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *CreateFilesParams) NKeys() int { return 1 }

// compile time check whether the CreateFilesParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*CreateFilesParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*CreateFilesParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *FileCreate) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, v.URI)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *FileCreate) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *FileCreate) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyURI {
		return dec.String(&v.URI)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *FileCreate) NKeys() int { return 1 }

// compile time check whether the FileCreate implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*FileCreate)(nil)
	_ gojay.UnmarshalerJSONObject = (*FileCreate)(nil)
)

// FileCreates represents a slice of FileCreate.
type FileCreates []FileCreate

// compile time check whether the FileCreates implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*FileCreates)(nil)
	_ gojay.UnmarshalerJSONArray = (*FileCreates)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v FileCreates) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v FileCreates) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *FileCreates) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value FileCreate
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *RenameFilesParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyFiles, FileRenames(v.Files))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *RenameFilesParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *RenameFilesParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyFiles {
		return dec.Array((*FileRenames)(&v.Files))
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *RenameFilesParams) NKeys() int { return 1 }

// compile time check whether the RenameFilesParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*RenameFilesParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*RenameFilesParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *FileRename) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyOldURI, v.OldURI)
	enc.StringKey(keyNewURI, v.NewURI)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *FileRename) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *FileRename) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyOldURI:
		return dec.String(&v.OldURI)
	case keyNewURI:
		return dec.String(&v.NewURI)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *FileRename) NKeys() int { return 2 }

// compile time check whether the FileRename implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*FileRename)(nil)
	_ gojay.UnmarshalerJSONObject = (*FileRename)(nil)
)

// FileRenames represents a slice of FileRename.
type FileRenames []FileRename

// compile time check whether the FileRenames implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*FileRenames)(nil)
	_ gojay.UnmarshalerJSONArray = (*FileRenames)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v FileRenames) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v FileRenames) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *FileRenames) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value FileRename
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DeleteFilesParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyFiles, FileDeletes(v.Files))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DeleteFilesParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DeleteFilesParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyFiles {
		return dec.Array((*FileDeletes)(&v.Files))
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DeleteFilesParams) NKeys() int { return 1 }

// compile time check whether the DeleteFilesParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DeleteFilesParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*DeleteFilesParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *FileDelete) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, v.URI)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *FileDelete) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *FileDelete) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyURI {
		return dec.String(&v.URI)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *FileDelete) NKeys() int { return 1 }

// compile time check whether the FileDelete implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*FileDelete)(nil)
	_ gojay.UnmarshalerJSONObject = (*FileDelete)(nil)
)

// FileDeletes represents a slice of FileDelete.
type FileDeletes []FileDelete

// compile time check whether the FileDeletes implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*FileDeletes)(nil)
	_ gojay.UnmarshalerJSONArray = (*FileDeletes)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v FileDeletes) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v FileDeletes) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *FileDeletes) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value FileDelete
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// CompletionItemKinds represents a slice of CompletionItemKind.
type CompletionItemKinds []CompletionItemKind

// compile time check whether the CompletionItemKinds implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*CompletionItemKinds)(nil)
	_ gojay.UnmarshalerJSONArray = (*CompletionItemKinds)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v CompletionItemKinds) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Float64(float64(v[i]))
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v CompletionItemKinds) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *CompletionItemKinds) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value CompletionItemKind
	if err := dec.Float64((*float64)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// InsertTextModes represents a slice of InsertTextMode.
type InsertTextModes []InsertTextMode

// compile time check whether the InsertTextModes implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*InsertTextModes)(nil)
	_ gojay.UnmarshalerJSONArray = (*InsertTextModes)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v InsertTextModes) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Float64(float64(v[i]))
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v InsertTextModes) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *InsertTextModes) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value InsertTextMode
	if err := dec.Float64((*float64)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarkupKinds represents a slice of MarkupKind.
type MarkupKinds []MarkupKind

// compile time check whether the MarkupKinds implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*MarkupKinds)(nil)
	_ gojay.UnmarshalerJSONArray = (*MarkupKinds)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v MarkupKinds) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.String(string(v[i]))
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v MarkupKinds) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray decodes JSON array elements into slice.
func (v *MarkupKinds) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value MarkupKind
	if err := dec.String((*string)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ReferencesParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyTextDocument, &v.TextDocument)
	enc.ObjectKeyOmitEmpty(keyPosition, &v.Position)
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
	encodeProgressToken(enc, keyPartialResultToken, v.PartialResultToken)
	enc.ObjectKeyOmitEmpty(keyContext, &v.Context)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ReferencesParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ReferencesParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyWorkDoneToken:
		return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
	case keyPartialResultToken:
		return decodeProgressToken(dec, k, keyPartialResultToken, v.PartialResultToken)
	case keyContext:
		return dec.Object(&v.Context)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ReferencesParams) NKeys() int { return 5 }

// compile time check whether the ReferencesParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ReferencesParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*ReferencesParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DocumentHighlightOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentHighlightOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentHighlightOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentHighlightOptions) NKeys() int { return 1 }

// compile time check whether the DocumentHighlightOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentHighlightOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentHighlightOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DocumentHighlightParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyTextDocument, &v.TextDocument)
	enc.ObjectKeyOmitEmpty(keyPosition, &v.Position)
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
	encodeProgressToken(enc, keyPartialResultToken, v.PartialResultToken)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DocumentHighlightParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DocumentHighlightParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyWorkDoneToken:
		return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
	case keyPartialResultToken:
		return decodeProgressToken(dec, k, keyPartialResultToken, v.PartialResultToken)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DocumentHighlightParams) NKeys() int { return 4 }

// compile time check whether the DocumentHighlightParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentHighlightParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentHighlightParams)(nil)
)

// SymbolTags represents a slice of SymbolTag.
type SymbolTags []SymbolTag

// compile time check whether the SymbolTags implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*SymbolTags)(nil)
	_ gojay.UnmarshalerJSONArray = (*SymbolTags)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v SymbolTags) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Float64(float64(v[i]))
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v SymbolTags) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray decodes JSON array elements into slice.
func (v *SymbolTags) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value SymbolTag
	if err := dec.Float64((*float64)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceSymbolOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceSymbolOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceSymbolOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceSymbolOptions) NKeys() int { return 1 }

// compile time check whether the WorkspaceSymbolOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceSymbolOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceSymbolOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DocumentFormattingOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentFormattingOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentFormattingOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentFormattingOptions) NKeys() int { return 1 }

// compile time check whether the DocumentFormattingOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentFormattingOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentFormattingOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DeclarationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DeclarationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DeclarationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DeclarationOptions) NKeys() int { return 1 }

// compile time check whether the DeclarationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DeclarationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DeclarationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DeclarationRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
	enc.AddArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.StringKeyOmitEmpty(keyID, v.ID)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DeclarationRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DeclarationRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	case keyDocumentSelector:
		if v.DocumentSelector == nil {
			v.DocumentSelector = DocumentSelector{}
		}
		return dec.Array(&v.DocumentSelector)
	case keyID:
		return dec.String(&v.ID)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DeclarationRegistrationOptions) NKeys() int { return 3 }

// compile time check whether the DeclarationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DeclarationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DeclarationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DeclarationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyTextDocument, &v.TextDocument)
	enc.ObjectKeyOmitEmpty(keyPosition, &v.Position)
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
	encodeProgressToken(enc, keyPartialResultToken, v.PartialResultToken)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DeclarationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DeclarationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyWorkDoneToken:
		return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
	case keyPartialResultToken:
		return decodeProgressToken(dec, k, keyPartialResultToken, v.PartialResultToken)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DeclarationParams) NKeys() int { return 4 }

// compile time check whether the DeclarationParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DeclarationParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*DeclarationParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DocumentRangeFormattingOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentRangeFormattingOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentRangeFormattingOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentRangeFormattingOptions) NKeys() int { return 1 }

// compile time check whether the DocumentRangeFormattingOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentRangeFormattingOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentRangeFormattingOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DefinitionOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DefinitionOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DefinitionOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DefinitionOptions) NKeys() int { return 1 }

// compile time check whether the DefinitionOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DefinitionOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DefinitionOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DefinitionParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyTextDocument, &v.TextDocument)
	enc.ObjectKeyOmitEmpty(keyPosition, &v.Position)
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
	encodeProgressToken(enc, keyPartialResultToken, v.PartialResultToken)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DefinitionParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DefinitionParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyWorkDoneToken:
		return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
	case keyPartialResultToken:
		return decodeProgressToken(dec, k, keyPartialResultToken, v.PartialResultToken)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DefinitionParams) NKeys() int { return 4 }

// compile time check whether the DefinitionParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DefinitionParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*DefinitionParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TypeDefinitionOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TypeDefinitionOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TypeDefinitionOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *TypeDefinitionOptions) NKeys() int { return 1 }

// compile time check whether the TypeDefinitionOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TypeDefinitionOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*TypeDefinitionOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TypeDefinitionRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
	enc.StringKeyOmitEmpty(keyID, v.ID)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TypeDefinitionRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TypeDefinitionRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	case keyDocumentSelector:
		if v.DocumentSelector == nil {
			v.DocumentSelector = DocumentSelector{}
		}
		return dec.Array(&v.DocumentSelector)
	case keyID:
		return dec.String(&v.ID)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *TypeDefinitionRegistrationOptions) NKeys() int { return 3 }

// compile time check whether the TypeDefinitionRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TypeDefinitionRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*TypeDefinitionRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TypeDefinitionParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyTextDocument, &v.TextDocument)
	enc.ObjectKeyOmitEmpty(keyPosition, &v.Position)
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
	encodeProgressToken(enc, keyPartialResultToken, v.PartialResultToken)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TypeDefinitionParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TypeDefinitionParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyWorkDoneToken:
		return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
	case keyPartialResultToken:
		return decodeProgressToken(dec, k, keyPartialResultToken, v.PartialResultToken)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *TypeDefinitionParams) NKeys() int { return 4 }

// compile time check whether the TypeDefinitionParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TypeDefinitionParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*TypeDefinitionParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ImplementationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ImplementationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ImplementationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ImplementationOptions) NKeys() int { return 1 }

// compile time check whether the ImplementationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ImplementationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*ImplementationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ImplementationRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
	enc.StringKeyOmitEmpty(keyID, v.ID)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ImplementationRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ImplementationRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		if v.DocumentSelector == nil {
			v.DocumentSelector = DocumentSelector{}
		}
		return dec.Array(&v.DocumentSelector)
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	case keyID:
		return dec.String(&v.ID)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ImplementationRegistrationOptions) NKeys() int { return 3 }

// compile time check whether the ImplementationRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ImplementationRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*ImplementationRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ImplementationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyTextDocument, &v.TextDocument)
	enc.ObjectKeyOmitEmpty(keyPosition, &v.Position)
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
	encodeProgressToken(enc, keyPartialResultToken, v.PartialResultToken)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ImplementationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ImplementationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyWorkDoneToken:
		return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
	case keyPartialResultToken:
		return decodeProgressToken(dec, k, keyPartialResultToken, v.PartialResultToken)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ImplementationParams) NKeys() int { return 4 }

// compile time check whether the ImplementationParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ImplementationParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*ImplementationParams)(nil)
)

// CodeActionKinds represents a slice of CodeActionKind.
type CodeActionKinds []CodeActionKind

// compile time check whether the CodeActionKinds implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*CodeActionKinds)(nil)
	_ gojay.UnmarshalerJSONArray = (*CodeActionKinds)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v CodeActionKinds) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.String(string(v[i]))
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v CodeActionKinds) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *CodeActionKinds) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value CodeActionKind
	if err := dec.String((*string)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *FoldingRangeOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *FoldingRangeOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *FoldingRangeOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *FoldingRangeOptions) NKeys() int { return 1 }

// compile time check whether the FoldingRangeOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*FoldingRangeOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*FoldingRangeOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *FoldingRangeRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
	enc.StringKeyOmitEmpty(keyID, v.ID)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *FoldingRangeRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *FoldingRangeRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		if v.DocumentSelector == nil {
			v.DocumentSelector = DocumentSelector{}
		}
		return dec.Array(&v.DocumentSelector)
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	case keyID:
		return dec.String(&v.ID)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *FoldingRangeRegistrationOptions) NKeys() int { return 4 }

// compile time check whether the FoldingRangeRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*FoldingRangeRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*FoldingRangeRegistrationOptions)(nil)
)

// TokenFormats represents a slice of TokenFormat.
type TokenFormats []TokenFormat

// compile time check whether the CodeActionKinds implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*TokenFormats)(nil)
	_ gojay.UnmarshalerJSONArray = (*TokenFormats)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v TokenFormats) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.String(string(v[i]))
	}
}

// IsNil implements gojay.MarshalerJSONArray.
func (v TokenFormats) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray.
func (v *TokenFormats) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value TokenFormat
	if err := dec.String((*string)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ShowDocumentParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, string(v.URI))
	enc.BoolKeyOmitEmpty(keyExternal, v.External)
	enc.BoolKeyOmitEmpty(keyTakeFocus, v.TakeFocus)
	enc.ObjectKeyOmitEmpty(keySelection, v.Selection)
}

// IsNil returns wether the structure is nil value or not.
func (v *ShowDocumentParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ShowDocumentParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyURI:
		return dec.String((*string)(&v.URI))
	case keyExternal:
		return dec.Bool(&v.External)
	case keyTakeFocus:
		return dec.Bool(&v.TakeFocus)
	case keySelection:
		if v.Selection == nil {
			v.Selection = &Range{}
		}
		return dec.Object(v.Selection)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ShowDocumentParams) NKeys() int { return 4 }

// compile time check whether the ShowDocumentParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ShowDocumentParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*ShowDocumentParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ShowDocumentResult) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey(keySuccess, v.Success)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ShowDocumentResult) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ShowDocumentResult) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keySuccess {
		return dec.Bool(&v.Success)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ShowDocumentResult) NKeys() int { return 1 }

// compile time check whether the ShowDocumentResult implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ShowDocumentResult)(nil)
	_ gojay.UnmarshalerJSONObject = (*ShowDocumentResult)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *InitializeResult) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyCapabilities, &v.Capabilities)
	enc.ObjectKeyOmitEmpty(keyServerInfo, v.ServerInfo)
}

// IsNil returns wether the structure is nil value or not.
func (v *InitializeResult) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializeResult) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyCapabilities:
		return dec.Object(&v.Capabilities)
	case keyServerInfo:
		if v.ServerInfo == nil {
			v.ServerInfo = &ServerInfo{}
		}
		return dec.Object(v.ServerInfo)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializeResult) NKeys() int { return 2 }

// compile time check whether the InitializeResult implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*InitializeResult)(nil)
	_ gojay.UnmarshalerJSONObject = (*InitializeResult)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ServerInfo) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyName, v.Name)
	enc.StringKeyOmitEmpty(keyVersion, v.Version)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ServerInfo) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ServerInfo) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyName:
		return dec.String(&v.Name)
	case keyVersion:
		return dec.String(&v.Version)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ServerInfo) NKeys() int { return 2 }

// compile time check whether the ServerInfo implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ServerInfo)(nil)
	_ gojay.UnmarshalerJSONObject = (*ServerInfo)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *InitializeError) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyRetry, v.Retry)
}

// IsNil returns wether the structure is nil value or not.
func (v *InitializeError) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializeError) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyRetry {
		return dec.Bool(&v.Retry)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializeError) NKeys() int { return 1 }

// compile time check whether the InitializeError implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*InitializeError)(nil)
	_ gojay.UnmarshalerJSONObject = (*InitializeError)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *CompletionOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyResolveProvider, v.ResolveProvider)
	enc.ArrayKeyOmitEmpty(keyTriggerCharacters, (*Strings)(&v.TriggerCharacters))
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyResolveProvider:
		return dec.Bool(&v.ResolveProvider)
	case keyTriggerCharacters:
		if v.TriggerCharacters == nil {
			v.TriggerCharacters = Strings{}
		}
		return dec.Array((*Strings)(&v.TriggerCharacters))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionOptions) NKeys() int { return 2 }

// compile time check whether the CompletionOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*CompletionOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*CompletionOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *SignatureHelpOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyTriggerCharacters, (*Strings)(&v.TriggerCharacters))
	enc.ArrayKeyOmitEmpty(keyRetriggerCharacters, (*Strings)(&v.RetriggerCharacters))
}

// IsNil returns wether the structure is nil value or not.
func (v *SignatureHelpOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SignatureHelpOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTriggerCharacters:
		var values Strings
		err := dec.Array(&values)
		if err == nil && len(values) > 0 {
			v.TriggerCharacters = []string(values)
		}
		return err
	case keyRetriggerCharacters:
		return dec.Array((*Strings)(&v.RetriggerCharacters))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SignatureHelpOptions) NKeys() int { return 2 }

// compile time check whether the SignatureHelpOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*SignatureHelpOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*SignatureHelpOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ReferencesOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ReferencesOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ReferencesOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ReferencesOptions) NKeys() int { return 1 }

// compile time check whether the ReferencesOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ReferencesOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*ReferencesOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *CodeActionOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyCodeActionKinds, (*CodeActionKinds)(&v.CodeActionKinds))
	enc.BoolKeyOmitEmpty(keyResolveProvider, v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeActionOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeActionOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyCodeActionKinds:
		return dec.Array((*CodeActionKinds)(&v.CodeActionKinds))
	case keyResolveProvider:
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionOptions) NKeys() int { return 2 }

// compile time check whether the CodeActionOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*CodeActionOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*CodeActionOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *CodeLensOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyResolveProvider, v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeLensOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeLensOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyResolveProvider {
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeLensOptions) NKeys() int { return 1 }

// compile time check whether the CodeLensOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*CodeLensOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*CodeLensOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v DocumentOnTypeFormattingOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyFirstTriggerCharacter, v.FirstTriggerCharacter)
	enc.ArrayKeyOmitEmpty(keyMoreTriggerCharacter, (*Strings)(&v.MoreTriggerCharacter))
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentOnTypeFormattingOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentOnTypeFormattingOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyFirstTriggerCharacter:
		return dec.String(&v.FirstTriggerCharacter)
	case keyMoreTriggerCharacter:
		var values Strings
		err := dec.Array(&values)
		if err == nil && len(values) > 0 {
			v.MoreTriggerCharacter = []string(values)
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentOnTypeFormattingOptions) NKeys() int { return 2 }

// compile time check whether the DocumentOnTypeFormattingOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentOnTypeFormattingOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentOnTypeFormattingOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *RenameOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyPrepareProvider, v.PrepareProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *RenameOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RenameOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyPrepareProvider {
		return dec.Bool(&v.PrepareProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RenameOptions) NKeys() int { return 1 }

// compile time check whether the RenameOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*RenameOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*RenameOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DocumentLinkOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyResolveProvider, v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentLinkOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentLinkOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyResolveProvider {
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentLinkOptions) NKeys() int { return 1 }

// compile time check whether the DocumentLinkOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentLinkOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentLinkOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ExecuteCommandOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyCommands, (*Strings)(&v.Commands))
}

// IsNil returns wether the structure is nil value or not.
func (v *ExecuteCommandOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ExecuteCommandOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyCommands {
		return dec.Array((*Strings)(&v.Commands))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ExecuteCommandOptions) NKeys() int { return 1 }

// compile time check whether the ExecuteCommandOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ExecuteCommandOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*ExecuteCommandOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *SaveOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyIncludeText, v.IncludeText)
}

// IsNil returns wether the structure is nil value or not.
func (v *SaveOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SaveOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyIncludeText {
		return dec.Bool(&v.IncludeText)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SaveOptions) NKeys() int { return 1 }

// compile time check whether the SaveOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*SaveOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*SaveOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentSyncOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyOpenClose, v.OpenClose)
	enc.Float64KeyOmitEmpty(keyChange, float64(v.Change))
	enc.BoolKeyOmitEmpty(keyWillSave, v.WillSave)
	enc.BoolKeyOmitEmpty(keyWillSaveWaitUntil, v.WillSaveWaitUntil)
	enc.ObjectKeyOmitEmpty(keySave, v.Save)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentSyncOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentSyncOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyOpenClose:
		return dec.Bool(&v.OpenClose)
	case keyChange:
		return dec.Float64((*float64)(&v.Change))
	case keyWillSave:
		return dec.Bool(&v.WillSave)
	case keyWillSaveWaitUntil:
		return dec.Bool(&v.WillSaveWaitUntil)
	case keySave:
		if v.Save == nil {
			v.Save = &SaveOptions{}
		}
		return dec.Object(v.Save)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentSyncOptions) NKeys() int { return 5 }

// compile time check whether the TextDocumentSyncOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentSyncOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentSyncOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *HoverOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *HoverOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *HoverOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *HoverOptions) NKeys() int { return 1 }

// compile time check whether the HoverOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*HoverOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*HoverOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *SemanticTokensOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil returns wether the structure is nil value or not.
func (v *SemanticTokensOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SemanticTokensOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SemanticTokensOptions) NKeys() int { return 1 }

// compile time check whether the SemanticTokensOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*SemanticTokensOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*SemanticTokensOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *SemanticTokensRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
	enc.StringKeyOmitEmpty(keyID, v.ID)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *SemanticTokensRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *SemanticTokensRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		if v.DocumentSelector == nil {
			v.DocumentSelector = DocumentSelector{}
		}
		return dec.Array(&v.DocumentSelector)
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	case keyID:
		return dec.String(&v.ID)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *SemanticTokensRegistrationOptions) NKeys() int { return 3 }

// compile time check whether the SemanticTokensRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*SemanticTokensRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*SemanticTokensRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *LinkedEditingRangeOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil returns wether the structure is nil value or not.
func (v *LinkedEditingRangeOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *LinkedEditingRangeOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *LinkedEditingRangeOptions) NKeys() int { return 1 }

// compile time check whether the LinkedEditingRangeOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*LinkedEditingRangeOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*LinkedEditingRangeOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *LinkedEditingRangeRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
	enc.StringKeyOmitEmpty(keyID, v.ID)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *LinkedEditingRangeRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *LinkedEditingRangeRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		if v.DocumentSelector == nil {
			v.DocumentSelector = DocumentSelector{}
		}
		return dec.Array(&v.DocumentSelector)
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	case keyID:
		return dec.String(&v.ID)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *LinkedEditingRangeRegistrationOptions) NKeys() int { return 3 }

// compile time check whether the LinkedEditingRangeRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*LinkedEditingRangeRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*LinkedEditingRangeRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *LinkedEditingRangeParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyTextDocument, &v.TextDocument)
	enc.ObjectKeyOmitEmpty(keyPosition, &v.Position)
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *LinkedEditingRangeParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *LinkedEditingRangeParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyWorkDoneToken:
		return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *LinkedEditingRangeParams) NKeys() int { return 3 }

// compile time check whether the LinkedEditingRangeParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*LinkedEditingRangeParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*LinkedEditingRangeParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *LinkedEditingRanges) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyRanges, Ranges(v.Ranges))
	enc.StringKeyOmitEmpty(keyWordPattern, v.WordPattern)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *LinkedEditingRanges) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *LinkedEditingRanges) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRanges:
		return dec.Array((*Ranges)(&v.Ranges))
	case keyWordPattern:
		return dec.String(&v.WordPattern)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *LinkedEditingRanges) NKeys() int { return 2 }

// compile time check whether the LinkedEditingRanges implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*LinkedEditingRanges)(nil)
	_ gojay.UnmarshalerJSONObject = (*LinkedEditingRanges)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *MonikerOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil returns wether the structure is nil value or not.
func (v *MonikerOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *MonikerOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *MonikerOptions) NKeys() int { return 1 }

// compile time check whether the MonikerOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*MonikerOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*MonikerOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *MonikerRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *MonikerRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *MonikerRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		if v.DocumentSelector == nil {
			v.DocumentSelector = DocumentSelector{}
		}
		return dec.Array(&v.DocumentSelector)
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *MonikerRegistrationOptions) NKeys() int { return 2 }

// compile time check whether the MonikerRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*MonikerRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*MonikerRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *MonikerParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyTextDocument, &v.TextDocument)
	enc.ObjectKeyOmitEmpty(keyPosition, &v.Position)
	encodeProgressToken(enc, keyWorkDoneToken, v.WorkDoneToken)
	encodeProgressToken(enc, keyPartialResultToken, v.PartialResultToken)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *MonikerParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *MonikerParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyWorkDoneToken:
		return decodeProgressToken(dec, k, keyWorkDoneToken, v.WorkDoneToken)
	case keyPartialResultToken:
		return decodeProgressToken(dec, k, keyPartialResultToken, v.PartialResultToken)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *MonikerParams) NKeys() int { return 3 }

// compile time check whether the MonikerParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*MonikerParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*MonikerParams)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *StaticRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty(keyID, v.ID)
}

// IsNil returns wether the structure is nil value or not.
func (v *StaticRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *StaticRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyID {
		return dec.String(&v.ID)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *StaticRegistrationOptions) NKeys() int { return 1 }

// compile time check whether the StaticRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*StaticRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*StaticRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DocumentLinkRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyResolveProvider, v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentLinkRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentLinkRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		if v.DocumentSelector == nil {
			v.DocumentSelector = DocumentSelector{}
		}
		return dec.Array(&v.DocumentSelector)
	case keyResolveProvider:
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentLinkRegistrationOptions) NKeys() int { return 2 }

// compile time check whether the DocumentLinkRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentLinkRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentLinkRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *InitializedParams) MarshalJSONObject(enc *gojay.Encoder) {}

// IsNil returns wether the structure is nil value or not.
func (v *InitializedParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializedParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializedParams) NKeys() int { return 0 }

// compile time check whether the InitializedParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*InitializedParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*InitializedParams)(nil)
)
