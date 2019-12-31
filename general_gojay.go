// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import "github.com/francoispqt/gojay"

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v WorkspaceFolders) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Object(&v[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v WorkspaceFolders) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *WorkspaceFolders) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value WorkspaceFolder
	if err := dec.Object(&value); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// compile time check whether the WorkspaceFolders implements a gojay.MarshalerJSONArray, gojay.UnmarshalerJSONArray and Pooler interface.
var (
	_ gojay.MarshalerJSONArray   = (*WorkspaceFolders)(nil)
	_ gojay.UnmarshalerJSONArray = (*WorkspaceFolders)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *InitializeParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key(keyProcessID, v.ProcessID)
	enc.StringKeyOmitEmpty(keyRootPath, v.RootPath)
	enc.StringKey(keyRootURI, string(v.RootURI))
	enc.AddInterfaceKey(keyInitializationOptions, v.InitializationOptions)
	enc.ObjectKey(keyCapabilities, &v.Capabilities)
	enc.StringKeyOmitEmpty(keyTrace, string(v.Trace))
	enc.ArrayKeyOmitEmpty(keyWorkspaceFolders, WorkspaceFolders(v.WorkspaceFolders))
}

// IsNil returns wether the structure is nil value or not.
func (v *InitializeParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializeParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyProcessID:
		return dec.Float64(&v.ProcessID)
	case keyRootPath:
		return dec.String(&v.RootPath)
	case keyRootURI:
		return dec.String((*string)(&v.RootURI))
	case keyInitializationOptions:
		return dec.Interface(&v.InitializationOptions)
	case keyCapabilities:
		return dec.Object(&v.Capabilities)
	case keyTrace:
		return dec.String((*string)(&v.Trace))
	case keyWorkspaceFolders:
		var values WorkspaceFolders
		err := dec.Array(&values)
		if err == nil && len(values) > 0 {
			v.WorkspaceFolders = []WorkspaceFolder(values)
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializeParams) NKeys() int { return 7 }

// compile time check whether the InitializeParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*InitializeParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*InitializeParams)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDocumentChanges, v.DocumentChanges)
	enc.StringKeyOmitEmpty(keyFailureHandling, v.FailureHandling)
	enc.ArrayKeyOmitEmpty(keyResourceOperations, (*Strings)(&v.ResourceOperations))
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentChanges:
		return dec.Bool(&v.DocumentChanges)
	case keyFailureHandling:
		return dec.String(&v.FailureHandling)
	case keyResourceOperations:
		var values Strings
		err := dec.Array(&values)
		if err == nil && len(values) > 0 {
			v.ResourceOperations = []string(values)
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) NKeys() int { return 3 }

// compile time check whether the WorkspaceClientCapabilitiesWorkspaceEdit implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesWorkspaceEdit)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesWorkspaceEdit)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesDidChangeConfiguration implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesDidChangeConfiguration)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesDidChangeConfiguration)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesDidChangeWatchedFiles implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesDidChangeWatchedFiles)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesDidChangeWatchedFiles)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keySymbolKind, v.SymbolKind)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesSymbol) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbol) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keySymbolKind:
		if v.SymbolKind == nil {
			v.SymbolKind = &WorkspaceClientCapabilitiesSymbolKind{}
		}
		return dec.Object(v.SymbolKind)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesSymbol) NKeys() int { return 2 }

// compile time check whether the WorkspaceClientCapabilitiesSymbol implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesSymbol)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesSymbol)(nil)
)

// SymbolKinds represents a slice of SymbolKind.
type SymbolKinds []SymbolKind

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v SymbolKinds) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Float64(float64(v[i]))
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v SymbolKinds) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray decodes JSON array elements into slice
func (v *SymbolKinds) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value SymbolKind
	if err := dec.Float64((*float64)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// compile time check whether the SymbolKinds implements a gojay.MarshalerJSONArray, gojay.UnmarshalerJSONArray and Pooler interface.
var (
	_ gojay.MarshalerJSONArray   = (*SymbolKinds)(nil)
	_ gojay.UnmarshalerJSONArray = (*SymbolKinds)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbolKind) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyValueSet, (*SymbolKinds)(&v.ValueSet))
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesSymbolKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbolKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		return dec.Array((*SymbolKinds)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesSymbolKind) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesSymbolKind implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesSymbolKind)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesSymbolKind)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesExecuteCommand) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesExecuteCommand) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesExecuteCommand) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesExecuteCommand) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesExecuteCommand implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesExecuteCommand)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesExecuteCommand)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyApplyEdit, v.ApplyEdit)
	enc.ObjectKeyOmitEmpty(keyWorkspaceEdit, v.WorkspaceEdit)
	enc.ObjectKeyOmitEmpty(keyDidChangeConfiguration, v.DidChangeConfiguration)
	enc.ObjectKeyOmitEmpty(keyDidChangeWatchedFiles, v.DidChangeWatchedFiles)
	enc.ObjectKeyOmitEmpty(keySymbol, v.Symbol)
	enc.ObjectKeyOmitEmpty(keyExecuteCommand, v.ExecuteCommand)
	enc.BoolKeyOmitEmpty(keyWorkspaceFolders, v.WorkspaceFolders)
	enc.BoolKeyOmitEmpty(keyConfiguration, v.Configuration)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyApplyEdit:
		return dec.Bool(&v.ApplyEdit)
	case keyWorkspaceEdit:
		if v.WorkspaceEdit == nil {
			v.WorkspaceEdit = &WorkspaceClientCapabilitiesWorkspaceEdit{}
		}
		return dec.Object(v.WorkspaceEdit)
	case keyDidChangeConfiguration:
		if v.DidChangeConfiguration == nil {
			v.DidChangeConfiguration = &WorkspaceClientCapabilitiesDidChangeConfiguration{}
		}
		return dec.Object(v.DidChangeConfiguration)
	case keyDidChangeWatchedFiles:
		if v.DidChangeWatchedFiles == nil {
			v.DidChangeWatchedFiles = &WorkspaceClientCapabilitiesDidChangeWatchedFiles{}
		}
		return dec.Object(v.DidChangeWatchedFiles)
	case keySymbol:
		if v.Symbol == nil {
			v.Symbol = &WorkspaceClientCapabilitiesSymbol{}
		}
		return dec.Object(v.Symbol)
	case keyExecuteCommand:
		if v.ExecuteCommand == nil {
			v.ExecuteCommand = &WorkspaceClientCapabilitiesExecuteCommand{}
		}
		return dec.Object(v.ExecuteCommand)
	case keyWorkspaceFolders:
		return dec.Bool(&v.WorkspaceFolders)
	case keyConfiguration:
		return dec.Bool(&v.Configuration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilities) NKeys() int { return 8 }

// compile time check whether the WorkspaceClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilities)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSynchronization) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDidSave, v.DidSave)
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyWillSave, v.WillSave)
	enc.BoolKeyOmitEmpty(keyWillSaveWaitUntil, v.WillSaveWaitUntil)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesSynchronization) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSynchronization) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDidSave:
		return dec.Bool(&v.DidSave)
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyWillSave:
		return dec.Bool(&v.WillSave)
	case keyWillSaveWaitUntil:
		return dec.Bool(&v.WillSaveWaitUntil)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSynchronization) NKeys() int { return 4 }

// compile time check whether the TextDocumentClientCapabilitiesSynchronization implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesSynchronization)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesSynchronization)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletion) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keyCompletionItem, v.CompletionItem)
	enc.ObjectKeyOmitEmpty(keyCompletionItemKind, v.CompletionItemKind)
	enc.BoolKeyOmitEmpty(keyContextSupport, v.ContextSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCompletion) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletion) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyCompletionItem:
		if v.CompletionItem == nil {
			v.CompletionItem = &TextDocumentClientCapabilitiesCompletionItem{}
		}
		return dec.Object(v.CompletionItem)
	case keyCompletionItemKind:
		if v.CompletionItemKind == nil {
			v.CompletionItemKind = &TextDocumentClientCapabilitiesCompletionItemKind{}
		}
		return dec.Object(v.CompletionItemKind)
	case keyContextSupport:
		return dec.Bool(&v.ContextSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCompletion) NKeys() int { return 4 }

// compile time check whether the TextDocumentClientCapabilitiesCompletion implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCompletion)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCompletion)(nil)
)

// CompletionItemKinds represents a slice of CompletionItemKind.
type CompletionItemKinds []CompletionItemKind

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v CompletionItemKinds) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Int(int(v[i]))
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v CompletionItemKinds) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *CompletionItemKinds) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value CompletionItemKind
	if err := dec.Int((*int)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// compile time check whether the CompletionItemKinds implements a gojay.MarshalerJSONArray, gojay.UnmarshalerJSONArray and Pooler interface.
var (
	_ gojay.MarshalerJSONArray   = (*CompletionItemKinds)(nil)
	_ gojay.UnmarshalerJSONArray = (*CompletionItemKinds)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemKind) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyValueSet, (*CompletionItemKinds)(&v.ValueSet))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCompletionItemKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		return dec.Array((*CompletionItemKinds)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCompletionItemKind) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCompletion implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCompletionItemKind)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCompletionItemKind)(nil)
)

// MarkupKinds represents a slice of MarkupKind.
type MarkupKinds []MarkupKind

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v MarkupKinds) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.String(string(v[i]))
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v MarkupKinds) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray decodes JSON array elements into slice
func (v *MarkupKinds) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value MarkupKind
	if err := dec.String((*string)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// compile time check whether the MarkupKinds implements a gojay.MarshalerJSONArray, gojay.UnmarshalerJSONArray and Pooler interface.
var (
	_ gojay.MarshalerJSONArray   = (*MarkupKinds)(nil)
	_ gojay.UnmarshalerJSONArray = (*MarkupKinds)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keySnippetSupport, v.SnippetSupport)
	enc.BoolKeyOmitEmpty(keyCommitCharactersSupport, v.CommitCharactersSupport)
	enc.ArrayKeyOmitEmpty(keyDocumentationFormat, (*MarkupKinds)(&v.DocumentationFormat))
	enc.BoolKeyOmitEmpty(keyDeprecatedSupport, v.DeprecatedSupport)
	enc.BoolKeyOmitEmpty(keyPreselectSupport, v.PreselectSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCompletionItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keySnippetSupport:
		return dec.Bool(&v.SnippetSupport)
	case keyCommitCharactersSupport:
		return dec.Bool(&v.CommitCharactersSupport)
	case keyDocumentationFormat:
		return dec.Array((*MarkupKinds)(&v.DocumentationFormat))
	case keyDeprecatedSupport:
		return dec.Bool(&v.DeprecatedSupport)
	case keyPreselectSupport:
		return dec.Bool(&v.PreselectSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCompletionItem) NKeys() int { return 5 }

// compile time check whether the TextDocumentClientCapabilitiesCompletionItem implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCompletionItem)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCompletionItem)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesHover) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ArrayKeyOmitEmpty(keyContentFormat, (*MarkupKinds)(&v.ContentFormat))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesHover) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesHover) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyContentFormat:
		return dec.Array((*MarkupKinds)(&v.ContentFormat))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesHover) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesHover implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesHover)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesHover)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureHelp) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keySignatureInformation, v.SignatureInformation)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesSignatureHelp) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureHelp) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keySignatureInformation:
		if v.SignatureInformation == nil {
			v.SignatureInformation = &TextDocumentClientCapabilitiesSignatureInformation{}
		}
		return dec.Object(v.SignatureInformation)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSignatureHelp) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesSignatureHelp implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesSignatureHelp)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesSignatureHelp)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyDocumentationFormat, (*MarkupKinds)(&v.DocumentationFormat))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesSignatureInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDocumentationFormat {
		return dec.Array((*MarkupKinds)(&v.DocumentationFormat))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSignatureInformation) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesSignatureInformation implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesSignatureInformation)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesSignatureInformation)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesReferences) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesReferences) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesReferences) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesReferences) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesReferences implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesReferences)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesReferences)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesDocumentHighlight implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDocumentHighlight)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDocumentHighlight)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keySymbolKind, v.SymbolKind)
	enc.BoolKeyOmitEmpty(keyHierarchicalDocumentSymbolSupport, v.HierarchicalDocumentSymbolSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keySymbolKind:
		if v.SymbolKind == nil {
			v.SymbolKind = &WorkspaceClientCapabilitiesSymbolKind{}
		}
		return dec.Object(v.SymbolKind)
	case keyHierarchicalDocumentSymbolSupport:
		return dec.Bool(&v.HierarchicalDocumentSymbolSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) NKeys() int { return 3 }

// compile time check whether the TextDocumentClientCapabilitiesDocumentSymbol implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDocumentSymbol)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDocumentSymbol)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesFormatting) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesFormatting implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesFormatting)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesFormatting)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRangeFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesRangeFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRangeFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesRangeFormatting) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesFormatting implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesRangeFormatting)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesRangeFormatting)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesOnTypeFormatting implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesOnTypeFormatting)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesOnTypeFormatting)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDeclaration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyLinkSupport, v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDeclaration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDeclaration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyLinkSupport:
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDeclaration) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesDeclaration implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDeclaration)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDeclaration)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDefinition) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyLinkSupport, v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDefinition) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDefinition) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyLinkSupport:
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDefinition) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesDefinition implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDefinition)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDefinition)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesTypeDefinition) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyLinkSupport, v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesTypeDefinition) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesTypeDefinition) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyLinkSupport:
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesTypeDefinition) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesTypeDefinition implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesTypeDefinition)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesTypeDefinition)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesImplementation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyLinkSupport, v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesImplementation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesImplementation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyLinkSupport:
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesImplementation) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesImplementation implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesImplementation)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesImplementation)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeAction) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keyCodeActionLiteralSupport, v.CodeActionLiteralSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCodeAction) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeAction) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyCodeActionLiteralSupport:
		if v.CodeActionLiteralSupport == nil {
			v.CodeActionLiteralSupport = &TextDocumentClientCapabilitiesCodeActionLiteralSupport{}
		}
		return dec.Object(v.CodeActionLiteralSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeAction) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesCodeAction implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCodeAction)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCodeAction)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyCodeActionKind, v.CodeActionKind)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyCodeActionKind {
		if v.CodeActionKind == nil {
			v.CodeActionKind = &TextDocumentClientCapabilitiesCodeActionKind{}
		}
		return dec.Object(v.CodeActionKind)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCodeActionLiteralSupport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCodeActionLiteralSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCodeActionLiteralSupport)(nil)
)

// CodeActionKinds represents a slice of CodeActionKind.
type CodeActionKinds []CodeActionKind

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v CodeActionKinds) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.String(string(v[i]))
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v CodeActionKinds) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *CodeActionKinds) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value CodeActionKind
	if err := dec.String((*string)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// compile time check whether the CodeActionKinds implements a gojay.MarshalerJSONArray, gojay.UnmarshalerJSONArray and Pooler interface.
var (
	_ gojay.MarshalerJSONArray   = (*CodeActionKinds)(nil)
	_ gojay.UnmarshalerJSONArray = (*CodeActionKinds)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionKind) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyValueSet, (*CodeActionKinds)(&v.ValueSet))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCodeActionKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		return dec.Array((*CodeActionKinds)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeActionKind) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCodeActionKind implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCodeActionKind)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCodeActionKind)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeLens) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCodeLens) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeLens) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeLens) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCodeLens implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCodeLens)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCodeLens)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentLink) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDocumentLink) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentLink) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentLink) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesDocumentLink implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDocumentLink)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDocumentLink)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesColorProvider) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesColorProvider) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesColorProvider) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesColorProvider) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesColorProvider implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesColorProvider)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesColorProvider)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRename) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyPrepareSupport, v.PrepareSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesRename) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRename) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyPrepareSupport:
		return dec.Bool(&v.PrepareSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesRename) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesRename implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesRename)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesRename)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyRelatedInformation, v.RelatedInformation)
	enc.BoolKeyOmitEmpty(keyTagSupport, v.TagSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRelatedInformation:
		return dec.Bool(&v.RelatedInformation)
	case keyTagSupport:
		return dec.Bool(&v.TagSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesPublishDiagnostics implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesPublishDiagnostics)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesPublishDiagnostics)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFoldingRange) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.Float64KeyOmitEmpty(keyRangeLimit, v.RangeLimit)
	enc.BoolKeyOmitEmpty(keyLineFoldingOnly, v.LineFoldingOnly)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesFoldingRange) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFoldingRange) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyRangeLimit:
		return dec.Float64(&v.RangeLimit)
	case keyLineFoldingOnly:
		return dec.Bool(&v.LineFoldingOnly)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesFoldingRange) NKeys() int { return 3 }

// compile time check whether the TextDocumentClientCapabilitiesFoldingRange implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesFoldingRange)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesFoldingRange)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSelectionRange) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesSelectionRange) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSelectionRange) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSelectionRange) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesSelectionRange implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesSelectionRange)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesSelectionRange)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keySynchronization, v.Synchronization)
	enc.ObjectKeyOmitEmpty(keyCompletion, v.Completion)
	enc.ObjectKeyOmitEmpty(keyHover, v.Hover)
	enc.ObjectKeyOmitEmpty(keySignatureHelp, v.SignatureHelp)
	enc.ObjectKeyOmitEmpty(keyReferences, v.References)
	enc.ObjectKeyOmitEmpty(keyDocumentHighlight, v.DocumentHighlight)
	enc.ObjectKeyOmitEmpty(keyDocumentSymbol, v.DocumentSymbol)
	enc.ObjectKeyOmitEmpty(keyFormatting, v.Formatting)
	enc.ObjectKeyOmitEmpty(keyRangeFormatting, v.RangeFormatting)
	enc.ObjectKeyOmitEmpty(keyOnTypeFormatting, v.OnTypeFormatting)
	enc.ObjectKeyOmitEmpty(keyDeclaration, v.Declaration)
	enc.ObjectKeyOmitEmpty(keyDefinition, v.Definition)
	enc.ObjectKeyOmitEmpty(keyTypeDefinition, v.TypeDefinition)
	enc.ObjectKeyOmitEmpty(keyImplementation, v.Implementation)
	enc.ObjectKeyOmitEmpty(keyCodeAction, v.CodeAction)
	enc.ObjectKeyOmitEmpty(keyCodeLens, v.CodeLens)
	enc.ObjectKeyOmitEmpty(keyDocumentLink, v.DocumentLink)
	enc.ObjectKeyOmitEmpty(keyColorProvider, v.ColorProvider)
	enc.ObjectKeyOmitEmpty(keyRename, v.Rename)
	enc.ObjectKeyOmitEmpty(keyPublishDiagnostics, v.PublishDiagnostics)
	enc.ObjectKeyOmitEmpty(keyFoldingRange, v.FoldingRange)
	enc.ObjectKeyOmitEmpty(keySelectionRange, v.SelectionRange)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keySynchronization:
		if v.Synchronization == nil {
			v.Synchronization = &TextDocumentClientCapabilitiesSynchronization{}
		}
		return dec.Object(v.Synchronization)
	case keyCompletion:
		if v.Completion == nil {
			v.Completion = &TextDocumentClientCapabilitiesCompletion{}
		}
		return dec.Object(v.Completion)
	case keyHover:
		if v.Hover == nil {
			v.Hover = &TextDocumentClientCapabilitiesHover{}
		}
		return dec.Object(v.Hover)
	case keySignatureHelp:
		if v.SignatureHelp == nil {
			v.SignatureHelp = &TextDocumentClientCapabilitiesSignatureHelp{}
		}
		return dec.Object(v.SignatureHelp)
	case keyReferences:
		if v.References == nil {
			v.References = &TextDocumentClientCapabilitiesReferences{}
		}
		return dec.Object(v.References)
	case keyDocumentHighlight:
		if v.DocumentHighlight == nil {
			v.DocumentHighlight = &TextDocumentClientCapabilitiesDocumentHighlight{}
		}
		return dec.Object(v.DocumentHighlight)
	case keyDocumentSymbol:
		if v.DocumentSymbol == nil {
			v.DocumentSymbol = &TextDocumentClientCapabilitiesDocumentSymbol{}
		}
		return dec.Object(v.DocumentSymbol)
	case keyFormatting:
		if v.Formatting == nil {
			v.Formatting = &TextDocumentClientCapabilitiesFormatting{}
		}
		return dec.Object(v.Formatting)
	case keyRangeFormatting:
		if v.RangeFormatting == nil {
			v.RangeFormatting = &TextDocumentClientCapabilitiesRangeFormatting{}
		}
		return dec.Object(v.RangeFormatting)
	case keyOnTypeFormatting:
		if v.OnTypeFormatting == nil {
			v.OnTypeFormatting = &TextDocumentClientCapabilitiesOnTypeFormatting{}
		}
		return dec.Object(v.OnTypeFormatting)
	case keyDeclaration:
		if v.Declaration == nil {
			v.Declaration = &TextDocumentClientCapabilitiesDeclaration{}
		}
		return dec.Object(v.Declaration)
	case keyDefinition:
		if v.Definition == nil {
			v.Definition = &TextDocumentClientCapabilitiesDefinition{}
		}
		return dec.Object(v.Definition)
	case keyTypeDefinition:
		if v.TypeDefinition == nil {
			v.TypeDefinition = &TextDocumentClientCapabilitiesTypeDefinition{}
		}
		return dec.Object(v.TypeDefinition)
	case keyImplementation:
		if v.Implementation == nil {
			v.Implementation = &TextDocumentClientCapabilitiesImplementation{}
		}
		return dec.Object(v.Implementation)
	case keyCodeAction:
		if v.CodeAction == nil {
			v.CodeAction = &TextDocumentClientCapabilitiesCodeAction{}
		}
		return dec.Object(v.CodeAction)
	case keyCodeLens:
		if v.CodeLens == nil {
			v.CodeLens = &TextDocumentClientCapabilitiesCodeLens{}
		}
		return dec.Object(v.CodeLens)
	case keyDocumentLink:
		if v.DocumentLink == nil {
			v.DocumentLink = &TextDocumentClientCapabilitiesDocumentLink{}
		}
		return dec.Object(v.DocumentLink)
	case keyColorProvider:
		if v.ColorProvider == nil {
			v.ColorProvider = &TextDocumentClientCapabilitiesColorProvider{}
		}
		return dec.Object(v.ColorProvider)
	case keyRename:
		if v.Rename == nil {
			v.Rename = &TextDocumentClientCapabilitiesRename{}
		}
		return dec.Object(v.Rename)
	case keyPublishDiagnostics:
		if v.PublishDiagnostics == nil {
			v.PublishDiagnostics = &TextDocumentClientCapabilitiesPublishDiagnostics{}
		}
		return dec.Object(v.PublishDiagnostics)
	case keyFoldingRange:
		if v.FoldingRange == nil {
			v.FoldingRange = &TextDocumentClientCapabilitiesFoldingRange{}
		}
		return dec.Object(v.FoldingRange)
	case keySelectionRange:
		if v.SelectionRange == nil {
			v.SelectionRange = &TextDocumentClientCapabilitiesSelectionRange{}
		}
		return dec.Object(v.SelectionRange)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilities) NKeys() int { return 22 }

// compile time check whether the TextDocumentClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilities)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyWorkspace, v.Workspace)
	enc.ObjectKeyOmitEmpty(keyTextDocument, v.TextDocument)
	enc.AddInterfaceKeyOmitEmpty(keyExperimental, v.Experimental)
}

// IsNil returns wether the structure is nil value or not.
func (v *ClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyWorkspace:
		if v.Workspace == nil {
			v.Workspace = &WorkspaceClientCapabilities{}
		}
		return dec.Object(v.Workspace)
	case keyTextDocument:
		if v.TextDocument == nil {
			v.TextDocument = &TextDocumentClientCapabilities{}
		}
		return dec.Object(v.TextDocument)
	case keyExperimental:
		return dec.Interface(&v.Experimental)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ClientCapabilities) NKeys() int { return 3 }

// compile time check whether the ClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*ClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*ClientCapabilities)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *InitializeResult) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyCapabilities, &v.Capabilities)
}

// IsNil returns wether the structure is nil value or not.
func (v *InitializeResult) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializeResult) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyCapabilities {
		return dec.Object(&v.Capabilities)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializeResult) NKeys() int { return 1 }

// compile time check whether the InitializeResult implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*InitializeResult)(nil)
	_ gojay.UnmarshalerJSONObject = (*InitializeResult)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the InitializeError implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*InitializeError)(nil)
	_ gojay.UnmarshalerJSONObject = (*InitializeError)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the CompletionOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*CompletionOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*CompletionOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SignatureHelpOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyTriggerCharacters, (*Strings)(&v.TriggerCharacters))
}

// IsNil returns wether the structure is nil value or not.
func (v *SignatureHelpOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SignatureHelpOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTriggerCharacters {
		var values Strings
		err := dec.Array(&values)
		if err == nil && len(values) > 0 {
			v.TriggerCharacters = []string(values)
		}
		return err
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SignatureHelpOptions) NKeys() int { return 1 }

// compile time check whether the SignatureHelpOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*SignatureHelpOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*SignatureHelpOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeActionOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyCodeActionKinds, (*CodeActionKinds)(&v.CodeActionKinds))
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeActionOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeActionOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyCodeActionKinds {
		return dec.Array((*CodeActionKinds)(&v.CodeActionKinds))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionOptions) NKeys() int { return 1 }

// compile time check whether the CodeActionOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*CodeActionOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*CodeActionOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the CodeLensOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*CodeLensOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*CodeLensOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the DocumentOnTypeFormattingOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentOnTypeFormattingOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentOnTypeFormattingOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the RenameOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*RenameOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*RenameOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the DocumentLinkOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentLinkOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentLinkOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the ExecuteCommandOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*ExecuteCommandOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*ExecuteCommandOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the SaveOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*SaveOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*SaveOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ColorProviderOptions) MarshalJSONObject(enc *gojay.Encoder) {}

// IsNil returns wether the structure is nil value or not.
func (v *ColorProviderOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ColorProviderOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ColorProviderOptions) NKeys() int { return 0 }

// compile time check whether the ColorProviderOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*ColorProviderOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*ColorProviderOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FoldingRangeProviderOptions) MarshalJSONObject(enc *gojay.Encoder) {}

// IsNil returns wether the structure is nil value or not.
func (v *FoldingRangeProviderOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FoldingRangeProviderOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FoldingRangeProviderOptions) NKeys() int { return 0 }

// compile time check whether the FoldingRangeProviderOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*FoldingRangeProviderOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*FoldingRangeProviderOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentSyncOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyOpenClose, v.OpenClose)
	enc.Float64KeyOmitEmpty(keyChange, v.Change)
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
		return dec.Float64(&v.Change)
	case keyWillSave:
		return dec.Bool(&v.WillSave)
	case keyWillSaveWaitUntil:
		return dec.Bool(&v.WillSaveWaitUntil)
	case keySave:
		return dec.Object(v.Save)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentSyncOptions) NKeys() int { return 5 }

// compile time check whether the TextDocumentSyncOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentSyncOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentSyncOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the StaticRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*StaticRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*StaticRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspace) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyWorkspaceFolders, v.WorkspaceFolders)
}

// IsNil returns wether the structure is nil value or not.
func (v *ServerCapabilitiesWorkspace) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ServerCapabilitiesWorkspace) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkspaceFolders {
		if v.WorkspaceFolders == nil {
			v.WorkspaceFolders = &ServerCapabilitiesWorkspaceFolders{}
		}
		return dec.Object(v.WorkspaceFolders)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ServerCapabilitiesWorkspace) NKeys() int { return 1 }

// compile time check whether the ServerCapabilitiesWorkspace implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*ServerCapabilitiesWorkspace)(nil)
	_ gojay.UnmarshalerJSONObject = (*ServerCapabilitiesWorkspace)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspaceFolders) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey(keySupported, v.Supported)
	enc.AddInterfaceKeyOmitEmpty(keyChangeNotifications, v.ChangeNotifications)
}

// IsNil returns wether the structure is nil value or not.
func (v *ServerCapabilitiesWorkspaceFolders) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ServerCapabilitiesWorkspaceFolders) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keySupported:
		return dec.Bool(&v.Supported)
	case keyChangeNotifications:
		return dec.Interface(&v.ChangeNotifications)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ServerCapabilitiesWorkspaceFolders) NKeys() int { return 2 }

// compile time check whether the ServerCapabilitiesWorkspaceFolders implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*ServerCapabilitiesWorkspaceFolders)(nil)
	_ gojay.UnmarshalerJSONObject = (*ServerCapabilitiesWorkspaceFolders)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ServerCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddInterfaceKeyOmitEmpty(keyTextDocumentSync, v.TextDocumentSync)
	enc.BoolKeyOmitEmpty(keyHoverProvider, v.HoverProvider)
	enc.ObjectKeyOmitEmpty(keyCompletionProvider, v.CompletionProvider)
	enc.ObjectKeyOmitEmpty(keySignatureHelpProvider, v.SignatureHelpProvider)
	enc.BoolKeyOmitEmpty(keyDefinitionProvider, v.DefinitionProvider)
	enc.AddInterfaceKeyOmitEmpty(keyTypeDefinitionProvider, v.TypeDefinitionProvider)
	enc.AddInterfaceKeyOmitEmpty(keyImplementationProvider, v.ImplementationProvider)
	enc.BoolKeyOmitEmpty(keyReferencesProvider, v.ReferencesProvider)
	enc.BoolKeyOmitEmpty(keyDocumentHighlightProvider, v.DocumentHighlightProvider)
	enc.BoolKeyOmitEmpty(keyDocumentSymbolProvider, v.DocumentSymbolProvider)
	enc.BoolKeyOmitEmpty(keyWorkspaceSymbolProvider, v.WorkspaceSymbolProvider)
	enc.BoolKeyOmitEmpty(keyCodeActionProvider, v.CodeActionProvider)
	enc.ObjectKeyOmitEmpty(keyCodeLensProvider, v.CodeLensProvider)
	enc.BoolKeyOmitEmpty(keyDocumentFormattingProvider, v.DocumentFormattingProvider)
	enc.BoolKeyOmitEmpty(keyDocumentRangeFormattingProvider, v.DocumentRangeFormattingProvider)
	enc.ObjectKeyOmitEmpty(keyDocumentOnTypeFormattingProvider, v.DocumentOnTypeFormattingProvider)
	enc.AddInterfaceKeyOmitEmpty(keyRenameProvider, v.RenameProvider)
	enc.ObjectKeyOmitEmpty(keyDocumentLinkProvider, v.DocumentLinkProvider)
	enc.AddInterfaceKeyOmitEmpty(keyColorProvider, v.ColorProvider)
	enc.AddInterfaceKeyOmitEmpty(keyFoldingRangeProvider, v.FoldingRangeProvider)
	enc.AddInterfaceKeyOmitEmpty(keySelectionRangeProvider, v.SelectionRangeProvider)
	enc.ObjectKeyOmitEmpty(keyExecuteCommandProvider, v.ExecuteCommandProvider)
	enc.ObjectKeyOmitEmpty(keyWorkspace, v.Workspace)
	enc.AddInterfaceKeyOmitEmpty(keyExperimental, v.Experimental)
}

// IsNil returns wether the structure is nil value or not.
func (v *ServerCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ServerCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocumentSync:
		return dec.Interface(&v.TextDocumentSync)
	case keyHoverProvider:
		return dec.Bool(&v.HoverProvider)
	case keyCompletionProvider:
		if v.CompletionProvider == nil {
			v.CompletionProvider = &CompletionOptions{}
		}
		return dec.Object(v.CompletionProvider)
	case keySignatureHelpProvider:
		if v.SignatureHelpProvider == nil {
			v.SignatureHelpProvider = &SignatureHelpOptions{}
		}
		return dec.Object(v.SignatureHelpProvider)
	case keyDefinitionProvider:
		return dec.Bool(&v.DefinitionProvider)
	case keyTypeDefinitionProvider:
		return dec.Interface(&v.TypeDefinitionProvider)
	case keyImplementationProvider:
		return dec.Interface(&v.ImplementationProvider)
	case keyReferencesProvider:
		return dec.Bool(&v.ReferencesProvider)
	case keyDocumentHighlightProvider:
		return dec.Bool(&v.DocumentHighlightProvider)
	case keyDocumentSymbolProvider:
		return dec.Bool(&v.DocumentSymbolProvider)
	case keyWorkspaceSymbolProvider:
		return dec.Bool(&v.WorkspaceSymbolProvider)
	case keyCodeActionProvider:
		return dec.Bool(&v.CodeActionProvider)
	case keyCodeLensProvider:
		if v.CodeLensProvider == nil {
			v.CodeLensProvider = &CodeLensOptions{}
		}
		return dec.Object(v.CodeLensProvider)
	case keyDocumentFormattingProvider:
		return dec.Bool(&v.DocumentFormattingProvider)
	case keyDocumentRangeFormattingProvider:
		return dec.Bool(&v.DocumentRangeFormattingProvider)
	case keyDocumentOnTypeFormattingProvider:
		if v.DocumentOnTypeFormattingProvider == nil {
			v.DocumentOnTypeFormattingProvider = &DocumentOnTypeFormattingOptions{}
		}
		return dec.Object(v.DocumentOnTypeFormattingProvider)
	case keyRenameProvider:
		return dec.Interface(&v.RenameProvider)
	case keyDocumentLinkProvider:
		if v.DocumentLinkProvider == nil {
			v.DocumentLinkProvider = &DocumentLinkOptions{}
		}
		return dec.Object(v.DocumentLinkProvider)
	case keyColorProvider:
		return dec.Interface(&v.ColorProvider)
	case keyFoldingRangeProvider:
		return dec.Interface(&v.FoldingRangeProvider)
	case keySelectionRangeProvider:
		return dec.Interface(&v.SelectionRangeProvider)
	case keyExecuteCommandProvider:
		if v.ExecuteCommandProvider == nil {
			v.ExecuteCommandProvider = &ExecuteCommandOptions{}
		}
		return dec.Object(v.ExecuteCommandProvider)
	case keyWorkspace:
		if v.Workspace == nil {
			v.Workspace = &ServerCapabilitiesWorkspace{}
		}
		return dec.Object(v.Workspace)
	case keyExperimental:
		return dec.Interface(&v.Experimental)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ServerCapabilities) NKeys() int { return 24 }

// compile time check whether the ServerCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*ServerCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*ServerCapabilities)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
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

// compile time check whether the DocumentLinkRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentLinkRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentLinkRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *InitializedParams) MarshalJSONObject(enc *gojay.Encoder) {}

// IsNil returns wether the structure is nil value or not.
func (v *InitializedParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializedParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializedParams) NKeys() int { return 0 }

// compile time check whether the InitializedParams implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject and Pooler interface.
var (
	_ gojay.MarshalerJSONObject   = (*InitializedParams)(nil)
	_ gojay.UnmarshalerJSONObject = (*InitializedParams)(nil)
)
