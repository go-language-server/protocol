// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CancelParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "id" {
		return dec.Object(&v.ID)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CancelParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CancelParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("id", &v.ID)
}

// IsNil returns wether the structure is nil value or not.
func (v *CancelParams) IsNil() bool { return v == nil }

type workspaceFolders []WorkspaceFolder

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *workspaceFolders) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := WorkspaceFolder{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *workspaceFolders) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *workspaceFolders) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.ObjectOmitEmpty(&t)
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *workspaceFolders) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializeParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "processId":
		return dec.Float64(&v.ProcessID)
	case "rootPath":
		return dec.String(&v.RootPath)
	case "rootUri":
		return dec.String((*string)(&v.RootURI))
	case "initializationOptions":
		return dec.Interface(&v.InitializationOptions)
	case "capabilities":
		return dec.Object(&v.Capabilities)
	case "trace":
		return dec.String(&v.Trace)
	case "workspaceFolders":
		return dec.Array((*workspaceFolders)(&v.WorkspaceFolders))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializeParams) NKeys() int { return 7 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *InitializeParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key("processId", v.ProcessID)
	enc.StringKeyOmitEmpty("rootPath", v.RootPath)
	enc.StringKey("rootUri", string(v.RootURI))
	enc.AddInterfaceKey("initializationOptions", &v.InitializationOptions)
	enc.ObjectKey("capabilities", &v.Capabilities)
	enc.StringKeyOmitEmpty("trace", v.Trace)
	enc.ArrayKeyOmitEmpty("workspaceFolders", (*workspaceFolders)(&v.WorkspaceFolders))
}

// IsNil returns wether the structure is nil value or not.
func (v *InitializeParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentChanges":
		return dec.Bool(&v.DocumentChanges)
	case "failureHandling":
		return dec.String(&v.FailureHandling)
	case "resourceOperations":
		return dec.Array((*stringSlice)(&v.ResourceOperations))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("documentChanges", v.DocumentChanges)
	enc.StringKeyOmitEmpty("failureHandling", v.FailureHandling)
	enc.ArrayKeyOmitEmpty("resourceOperations", (*stringSlice)(&v.ResourceOperations))
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbol) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "symbolKind":
		if v.SymbolKind == nil {
			v.SymbolKind = &WorkspaceClientCapabilitiesSymbolKind{}
		}
		return dec.Object(v.SymbolKind)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesSymbol) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty("symbolKind", v.SymbolKind)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesSymbol) IsNil() bool { return v == nil }

type symbolKindValueSet []SymbolKind

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *symbolKindValueSet) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var t SymbolKind
	if err := dec.Float64((*float64)(&t)); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *symbolKindValueSet) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *symbolKindValueSet) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.Float64KeyOmitEmpty("valueSet", float64(t))
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *symbolKindValueSet) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbolKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "valueSet" {
		dec.Array((*symbolKindValueSet)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesSymbolKind) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbolKind) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty("valueSet", (*symbolKindValueSet)(&v.ValueSet))
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesSymbolKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesExecuteCommand) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesExecuteCommand) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesExecuteCommand) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilitiesExecuteCommand) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "applyEdit":
		return dec.Bool(&v.ApplyEdit)
	case "workspaceEdit":
		if v.WorkspaceEdit == nil {
			v.WorkspaceEdit = &WorkspaceClientCapabilitiesWorkspaceEdit{}
		}
		return dec.Object(v.WorkspaceEdit)
	case "didChangeConfiguration":
		if v.DidChangeConfiguration == nil {
			v.DidChangeConfiguration = &WorkspaceClientCapabilitiesDidChangeConfiguration{}
		}
		return dec.Object(v.DidChangeConfiguration)
	case "didChangeWatchedFiles":
		if v.DidChangeWatchedFiles == nil {
			v.DidChangeWatchedFiles = &WorkspaceClientCapabilitiesDidChangeWatchedFiles{}
		}
		return dec.Object(v.DidChangeWatchedFiles)
	case "symbol":
		if v.Symbol == nil {
			v.Symbol = &WorkspaceClientCapabilitiesSymbol{}
		}
		return dec.Object(v.Symbol)
	case "executeCommand":
		if v.ExecuteCommand == nil {
			v.ExecuteCommand = &WorkspaceClientCapabilitiesExecuteCommand{}
		}
		return dec.Object(v.ExecuteCommand)
	case "workspaceFolders":
		return dec.Bool(&v.WorkspaceFolders)
	case "configuration":
		return dec.Bool(&v.Configuration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilities) NKeys() int { return 8 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("applyEdit", v.ApplyEdit)
	enc.ObjectKeyOmitEmpty("workspaceEdit", v.WorkspaceEdit)
	enc.ObjectKeyOmitEmpty("didChangeConfiguration", v.DidChangeConfiguration)
	enc.ObjectKeyOmitEmpty("didChangeWatchedFiles", v.DidChangeWatchedFiles)
	enc.ObjectKeyOmitEmpty("symbol", v.Symbol)
	enc.ObjectKeyOmitEmpty("executeCommand", v.ExecuteCommand)
	enc.BoolKeyOmitEmpty("workspaceFolders", v.WorkspaceFolders)
	enc.BoolKeyOmitEmpty("configuration", v.Configuration)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSynchronization) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "didSave":
		return dec.Bool(&v.DidSave)
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "willSave":
		return dec.Bool(&v.WillSave)
	case "willSaveWaitUntil":
		return dec.Bool(&v.WillSaveWaitUntil)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSynchronization) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSynchronization) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("didSave", v.DidSave)
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKeyOmitEmpty("willSave", v.WillSave)
	enc.BoolKeyOmitEmpty("willSaveWaitUntil", v.WillSaveWaitUntil)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesSynchronization) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletion) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "completionItem":
		if v.CompletionItem == nil {
			v.CompletionItem = &TextDocumentClientCapabilitiesCompletionItem{}
		}
		return dec.Object(v.CompletionItem)
	case "completionItemKind":
		return dec.Int((*int)(v.CompletionItemKind))
	case "contextSupport":
		return dec.Bool(&v.ContextSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCompletion) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletion) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty("completionItem", v.CompletionItem)
	enc.IntKeyOmitEmpty("completionItemKind", int(*v.CompletionItemKind))
	enc.BoolKeyOmitEmpty("contextSupport", v.ContextSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCompletion) IsNil() bool { return v == nil }

type documentationFormat []MarkupKind

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *documentationFormat) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var t MarkupKind
	if err := dec.String((*string)(&t)); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *documentationFormat) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *documentationFormat) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.StringKeyOmitEmpty("valueSet", string(t))
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *documentationFormat) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "snippetSupport":
		return dec.Bool(&v.SnippetSupport)
	case "commitCharactersSupport":
		return dec.Bool(&v.CommitCharactersSupport)
	case "documentationFormat":
		return dec.Array((*documentationFormat)(&v.DocumentationFormat))
	case "deprecatedSupport":
		return dec.Bool(&v.DeprecatedSupport)
	case "preselectSupport":
		return dec.Bool(&v.PreselectSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCompletionItem) NKeys() int { return 5 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("snippetSupport", v.SnippetSupport)
	enc.BoolKeyOmitEmpty("commitCharactersSupport", v.CommitCharactersSupport)
	enc.ArrayKeyOmitEmpty("documentationFormat", (*documentationFormat)(&v.DocumentationFormat))
	enc.BoolKeyOmitEmpty("deprecatedSupport", v.DeprecatedSupport)
	enc.BoolKeyOmitEmpty("preselectSupport", v.PreselectSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCompletionItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesHover) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "contentFormat":
		return dec.Array((*documentationFormat)(&v.ContentFormat))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesHover) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesHover) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.ArrayKeyOmitEmpty("contentFormat", (*documentationFormat)(&v.ContentFormat))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesHover) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureHelp) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "signatureInformation":
		if v.SignatureInformation == nil {
			v.SignatureInformation = &TextDocumentClientCapabilitiesSignatureInformation{}
		}
		return dec.Object(v.SignatureInformation)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSignatureHelp) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureHelp) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty("signatureInformation", v.SignatureInformation)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesSignatureHelp) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "documentationFormat" {
		return dec.Array((*documentationFormat)(&v.DocumentationFormat))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSignatureInformation) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty("documentationFormat", (*documentationFormat)(&v.DocumentationFormat))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesSignatureInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesReferences) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesReferences) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesReferences) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesReferences) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "symbolKind":
		if v.SymbolKind == nil {
			v.SymbolKind = &WorkspaceClientCapabilitiesSymbolKind{}
		}
		return dec.Object(v.SymbolKind)
	case "hierarchicalDocumentSymbolSupport":
		return dec.Bool(&v.HierarchicalDocumentSymbolSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty("symbolKind", v.SymbolKind)
	enc.BoolKeyOmitEmpty("hierarchicalDocumentSymbolSupport", v.HierarchicalDocumentSymbolSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesFormatting) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRangeFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesRangeFormatting) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRangeFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesRangeFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDeclaration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "linkSupport":
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDeclaration) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDeclaration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKeyOmitEmpty("linkSupport", v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDeclaration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDefinition) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "linkSupport":
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDefinition) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDefinition) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKeyOmitEmpty("linkSupport", v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDefinition) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesTypeDefinition) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "linkSupport":
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesTypeDefinition) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesTypeDefinition) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKeyOmitEmpty("linkSupport", v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesTypeDefinition) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesImplementation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "linkSupport":
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesImplementation) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesImplementation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKeyOmitEmpty("linkSupport", v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesImplementation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeAction) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "codeActionLiteralSupport":
		if v.CodeActionLiteralSupport == nil {
			v.CodeActionLiteralSupport = &TextDocumentClientCapabilitiesCodeActionLiteralSupport{}
		}
		return dec.Object(v.CodeActionLiteralSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeAction) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeAction) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty("codeActionLiteralSupport", v.CodeActionLiteralSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCodeAction) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "codeActionKind":
		if v.CodeActionKind == nil {
			v.CodeActionKind = &TextDocumentClientCapabilitiesCodeActionKind{}
		}
		return dec.Object(v.CodeActionKind)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty("codeActionKind", v.CodeActionKind)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) IsNil() bool { return v == nil }

type codeActionKindValueSet []CodeActionKind

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *codeActionKindValueSet) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var t CodeActionKind
	if err := dec.String((*string)(&t)); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *codeActionKindValueSet) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *codeActionKindValueSet) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.StringKey("valueSet", string(t))
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *codeActionKindValueSet) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "valueSet" {
		dec.Array((*codeActionKindValueSet)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeActionKind) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionKind) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("valueSet", (*codeActionKindValueSet)(&v.ValueSet))
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCodeActionKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeLens) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeLens) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeLens) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesCodeLens) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentLink) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentLink) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentLink) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesDocumentLink) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesColorProvider) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "dynamicRegistration" {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesColorProvider) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesColorProvider) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesColorProvider) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRename) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "prepareSupport":
		return dec.Bool(&v.PrepareSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesRename) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRename) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKeyOmitEmpty("prepareSupport", v.PrepareSupport)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesRename) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "relatedInformation" {
		return dec.Bool(&v.RelatedInformation)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("relatedInformation", v.RelatedInformation)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFoldingRange) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "rangeLimit":
		return dec.Float64(&v.RangeLimit)
	case "lineFoldingOnly":
		return dec.Bool(&v.LineFoldingOnly)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesFoldingRange) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFoldingRange) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("dynamicRegistration", v.DynamicRegistration)
	enc.Float64KeyOmitEmpty("rangeLimit", v.RangeLimit)
	enc.BoolKeyOmitEmpty("lineFoldingOnly", v.LineFoldingOnly)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilitiesFoldingRange) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "synchronization":
		if v.Synchronization == nil {
			v.Synchronization = &TextDocumentClientCapabilitiesSynchronization{}
		}
		return dec.Object(v.Synchronization)
	case "completion":
		if v.Completion == nil {
			v.Completion = &TextDocumentClientCapabilitiesCompletion{}
		}
		return dec.Object(v.Completion)
	case "hover":
		if v.Hover == nil {
			v.Hover = &TextDocumentClientCapabilitiesHover{}
		}
		return dec.Object(v.Hover)
	case "signatureHelp":
		if v.SignatureHelp == nil {
			v.SignatureHelp = &TextDocumentClientCapabilitiesSignatureHelp{}
		}
		return dec.Object(v.SignatureHelp)
	case "references":
		if v.References == nil {
			v.References = &TextDocumentClientCapabilitiesReferences{}
		}
		return dec.Object(v.References)
	case "documentHighlight":
		if v.DocumentHighlight == nil {
			v.DocumentHighlight = &TextDocumentClientCapabilitiesDocumentHighlight{}
		}
		return dec.Object(v.DocumentHighlight)
	case "documentSymbol":
		if v.DocumentSymbol == nil {
			v.DocumentSymbol = &TextDocumentClientCapabilitiesDocumentSymbol{}
		}
		return dec.Object(v.DocumentSymbol)
	case "formatting":
		if v.Formatting == nil {
			v.Formatting = &TextDocumentClientCapabilitiesFormatting{}
		}
		return dec.Object(v.Formatting)
	case "rangeFormatting":
		if v.RangeFormatting == nil {
			v.RangeFormatting = &TextDocumentClientCapabilitiesRangeFormatting{}
		}
		return dec.Object(v.RangeFormatting)
	case "onTypeFormatting":
		if v.OnTypeFormatting == nil {
			v.OnTypeFormatting = &TextDocumentClientCapabilitiesOnTypeFormatting{}
		}
		return dec.Object(v.OnTypeFormatting)
	case "declaration":
		if v.Declaration == nil {
			v.Declaration = &TextDocumentClientCapabilitiesDeclaration{}
		}
		return dec.Object(v.Declaration)
	case "definition":
		if v.Definition == nil {
			v.Definition = &TextDocumentClientCapabilitiesDefinition{}
		}
		return dec.Object(v.Definition)
	case "typeDefinition":
		if v.TypeDefinition == nil {
			v.TypeDefinition = &TextDocumentClientCapabilitiesTypeDefinition{}
		}
		return dec.Object(v.TypeDefinition)
	case "implementation":
		if v.Implementation == nil {
			v.Implementation = &TextDocumentClientCapabilitiesImplementation{}
		}
		return dec.Object(v.Implementation)
	case "codeAction":
		if v.CodeAction == nil {
			v.CodeAction = &TextDocumentClientCapabilitiesCodeAction{}
		}
		return dec.Object(v.CodeAction)
	case "codeLens":
		if v.CodeLens == nil {
			v.CodeLens = &TextDocumentClientCapabilitiesCodeLens{}
		}
		return dec.Object(v.CodeLens)
	case "documentLink":
		if v.DocumentLink == nil {
			v.DocumentLink = &TextDocumentClientCapabilitiesDocumentLink{}
		}
		return dec.Object(v.DocumentLink)
	case "colorProvider":
		if v.ColorProvider == nil {
			v.ColorProvider = &TextDocumentClientCapabilitiesColorProvider{}
		}
		return dec.Object(v.ColorProvider)
	case "rename":
		if v.Rename == nil {
			v.Rename = &TextDocumentClientCapabilitiesRename{}
		}
		return dec.Object(v.Rename)
	case "publishDiagnostics":
		if v.PublishDiagnostics == nil {
			v.PublishDiagnostics = &TextDocumentClientCapabilitiesPublishDiagnostics{}
		}
		return dec.Object(v.PublishDiagnostics)
	case "foldingRange":
		if v.FoldingRange == nil {
			v.FoldingRange = &TextDocumentClientCapabilitiesFoldingRange{}
		}
		return dec.Object(v.FoldingRange)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilities) NKeys() int { return 21 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty("synchronization", v.Synchronization)
	enc.ObjectKeyOmitEmpty("completion", v.Completion)
	enc.ObjectKeyOmitEmpty("hover", v.Hover)
	enc.ObjectKeyOmitEmpty("signatureHelp", v.SignatureHelp)
	enc.ObjectKeyOmitEmpty("references", v.References)
	enc.ObjectKeyOmitEmpty("documentHighlight", v.DocumentHighlight)
	enc.ObjectKeyOmitEmpty("documentSymbol", v.DocumentSymbol)
	enc.ObjectKeyOmitEmpty("formatting", v.Formatting)
	enc.ObjectKeyOmitEmpty("rangeFormatting", v.RangeFormatting)
	enc.ObjectKeyOmitEmpty("onTypeFormatting", v.OnTypeFormatting)
	enc.ObjectKeyOmitEmpty("declaration", v.Declaration)
	enc.ObjectKeyOmitEmpty("definition", v.Definition)
	enc.ObjectKeyOmitEmpty("typeDefinition", v.TypeDefinition)
	enc.ObjectKeyOmitEmpty("implementation", v.Implementation)
	enc.ObjectKeyOmitEmpty("codeAction", v.CodeAction)
	enc.ObjectKeyOmitEmpty("codeLens", v.CodeLens)
	enc.ObjectKeyOmitEmpty("documentLink", v.DocumentLink)
	enc.ObjectKeyOmitEmpty("colorProvider", v.ColorProvider)
	enc.ObjectKeyOmitEmpty("rename", v.Rename)
	enc.ObjectKeyOmitEmpty("publishDiagnostics", v.PublishDiagnostics)
	enc.ObjectKeyOmitEmpty("foldingRange", v.FoldingRange)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "workspace":
		if v.Workspace == nil {
			v.Workspace = &WorkspaceClientCapabilities{}
		}
		return dec.Object(v.Workspace)
	case "textDocument":
		if v.TextDocument == nil {
			v.TextDocument = &TextDocumentClientCapabilities{}
		}
		return dec.Object(v.TextDocument)
	case "experimental":
		return dec.Interface(&v.Experimental)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ClientCapabilities) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("workspace", v.Workspace)
	enc.ObjectKey("textDocument", v.TextDocument)
	enc.AddInterfaceKey("experimental", v.Experimental)
}

// IsNil returns wether the structure is nil value or not.
func (v *ClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializeResult) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "capabilities" {
		return dec.Object(&v.Capabilities)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializeResult) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *InitializeResult) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("capabilities", &v.Capabilities)
}

// IsNil returns wether the structure is nil value or not.
func (v *InitializeResult) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializeError) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "retry" {
		return dec.Bool(&v.Retry)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializeError) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *InitializeError) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("retry", v.Retry)
}

// IsNil returns wether the structure is nil value or not.
func (v *InitializeError) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "resolveProvider":
		return dec.Bool(&v.ResolveProvider)
	case "triggerCharacters":
		return dec.Array((*stringSlice)(&v.TriggerCharacters))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("resolveProvider", v.ResolveProvider)
	enc.ArrayKeyOmitEmpty("triggerCharacters", (*stringSlice)(&v.TriggerCharacters))
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SignatureHelpOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "codeActionKinds" {
		dec.Array((*stringSlice)(&v.TriggerCharacters))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SignatureHelpOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SignatureHelpOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty("codeActionKinds", (*stringSlice)(&v.TriggerCharacters))
}

// IsNil returns wether the structure is nil value or not.
func (v *SignatureHelpOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeActionOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "codeActionKinds" {
		dec.Array((*codeActionKindValueSet)(&v.CodeActionKinds))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeActionOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty("codeActionKinds", (*codeActionKindValueSet)(&v.CodeActionKinds))
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeActionOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeLensOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "resolveProvider" {
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeLensOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeLensOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("resolveProvider", v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeLensOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentOnTypeFormattingOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "firstTriggerCharacter":
		return dec.String(&v.FirstTriggerCharacter)
	case "moreTriggerCharacter":
		return dec.Array((*stringSlice)(&v.MoreTriggerCharacter))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentOnTypeFormattingOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentOnTypeFormattingOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("firstTriggerCharacter", v.FirstTriggerCharacter)
	enc.ArrayKeyOmitEmpty("moreTriggerCharacter", (*stringSlice)(&v.MoreTriggerCharacter))
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentOnTypeFormattingOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RenameOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "prepareProvider" {
		return dec.Bool(&v.PrepareProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RenameOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *RenameOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("prepareProvider", v.PrepareProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *RenameOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentLinkOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "resolveProvider" {
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentLinkOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentLinkOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("resolveProvider", v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentLinkOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ExecuteCommandOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "commands" {
		return dec.Array((*stringSlice)(&v.Commands))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ExecuteCommandOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ExecuteCommandOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("commands", (*stringSlice)(&v.Commands))
}

// IsNil returns wether the structure is nil value or not.
func (v *ExecuteCommandOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SaveOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "includeText" {
		return dec.Bool(&v.IncludeText)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SaveOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SaveOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("includeText", v.IncludeText)
}

// IsNil returns wether the structure is nil value or not.
func (v *SaveOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ColorProviderOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ColorProviderOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ColorProviderOptions) MarshalJSONObject(enc *gojay.Encoder) {}

// IsNil returns wether the structure is nil value or not.
func (v *ColorProviderOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FoldingRangeProviderOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FoldingRangeProviderOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FoldingRangeProviderOptions) MarshalJSONObject(enc *gojay.Encoder) {}

// IsNil returns wether the structure is nil value or not.
func (v *FoldingRangeProviderOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *TextDocumentSyncOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "openClose":
		return dec.Bool(&v.OpenClose)
	case "change":
		return dec.Float64(&v.Change)
	case "willSave":
		return dec.Bool(&v.WillSave)
	case "willSaveWaitUntil":
		return dec.Bool(&v.WillSaveWaitUntil)
	case "save":
		if v.Save == nil {
			v.Save = &SaveOptions{}
		}
		return dec.Object(v.Save)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentSyncOptions) NKeys() int { return 5 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *TextDocumentSyncOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty("openClose", v.OpenClose)
	enc.Float64KeyOmitEmpty("change", v.Change)
	enc.BoolKeyOmitEmpty("willSave", v.WillSave)
	enc.BoolKeyOmitEmpty("willSaveWaitUntil", v.WillSaveWaitUntil)
	enc.ObjectKeyOmitEmpty("save", v.Save)
}

// IsNil returns wether the structure is nil value or not.
func (v *TextDocumentSyncOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *StaticRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "id" {
		return dec.String(&v.ID)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *StaticRegistrationOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *StaticRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty("id", v.ID)
}

// IsNil returns wether the structure is nil value or not.
func (v *StaticRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ServerCapabilitiesWorkspace) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "workspaceFolders" {
		if v.WorkspaceFolders == nil {
			v.WorkspaceFolders = &ServerCapabilitiesWorkspaceFolders{}
		}
		return dec.Object(v.WorkspaceFolders)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ServerCapabilitiesWorkspace) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspace) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty("workspaceFolders", v.WorkspaceFolders)
}

// IsNil returns wether the structure is nil value or not.
func (v *ServerCapabilitiesWorkspace) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ServerCapabilitiesWorkspaceFolders) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "supported":
		return dec.Bool(&v.Supported)
	case "changeNotifications":
		return dec.Interface(&v.ChangeNotifications)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ServerCapabilitiesWorkspaceFolders) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspaceFolders) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("supported", v.Supported)
	enc.AddInterfaceKeyOmitEmpty("changeNotifications", &v.ChangeNotifications)
}

// IsNil returns wether the structure is nil value or not.
func (v *ServerCapabilitiesWorkspaceFolders) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ServerCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocumentSync":
		return dec.DecodeInterface(&v.TextDocumentSync)
	case "hoverProvider":
		return dec.Bool(&v.HoverProvider)
	case "completionProvider":
		if v.CompletionProvider == nil {
			v.CompletionProvider = &CompletionOptions{}
		}
		return dec.Object(v.CompletionProvider)
	case "signatureHelpProvider":
		if v.SignatureHelpProvider == nil {
			v.SignatureHelpProvider = &SignatureHelpOptions{}
		}
		return dec.Object(v.SignatureHelpProvider)
	case "definitionProvider":
		return dec.Bool(&v.DefinitionProvider)
	case "typeDefinitionProvider":
		return dec.DecodeInterface(&v.TypeDefinitionProvider)
	case "implementationProvider":
		return dec.DecodeInterface(&v.ImplementationProvider)
	case "referencesProvider":
		return dec.Bool(&v.ReferencesProvider)
	case "documentHighlightProvider":
		return dec.Bool(&v.DocumentHighlightProvider)
	case "documentSymbolProvider":
		return dec.Bool(&v.DocumentSymbolProvider)
	case "workspaceSymbolProvider":
		return dec.Bool(&v.WorkspaceSymbolProvider)
	case "codeActionProvider":
		return dec.Bool(&v.CodeActionProvider)
	case "codeLensProvider":
		if v.CodeLensProvider == nil {
			v.CodeLensProvider = &CodeLensOptions{}
		}
		return dec.Object(v.CodeLensProvider)
	case "documentFormattingProvider":
		return dec.Bool(&v.DocumentFormattingProvider)
	case "documentRangeFormattingProvider":
		return dec.Bool(&v.DocumentRangeFormattingProvider)
	case "documentOnTypeFormattingProvider":
		if v.DocumentOnTypeFormattingProvider == nil {
			v.DocumentOnTypeFormattingProvider = &DocumentOnTypeFormattingOptions{}
		}
		return dec.Object(v.DocumentOnTypeFormattingProvider)
	case "renameProvider":
		return dec.Bool(&v.RenameProvider)
	case "documentLinkProvider":
		if v.DocumentLinkProvider == nil {
			v.DocumentLinkProvider = &DocumentLinkOptions{}
		}
		return dec.Object(v.DocumentLinkProvider)
	case "colorProvider":
		return dec.DecodeInterface(&v.ColorProvider)
	case "foldingRangeProvider":
		return dec.DecodeInterface(&v.FoldingRangeProvider)
	case "executeCommandProvider":
		if v.ExecuteCommandProvider == nil {
			v.ExecuteCommandProvider = &ExecuteCommandOptions{}
		}
		return dec.Object(v.ExecuteCommandProvider)
	case "workspace":
		if v.Workspace == nil {
			v.Workspace = &ServerCapabilitiesWorkspace{}
		}
		return dec.Object(v.Workspace)
	case "experimental":
		return dec.DecodeInterface(&v.Experimental)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ServerCapabilities) NKeys() int { return 23 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ServerCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddInterfaceKeyOmitEmpty("textDocumentSync", v.TextDocumentSync)
	enc.BoolKeyOmitEmpty("hoverProvider", v.HoverProvider)
	enc.ObjectKeyOmitEmpty("completionProvider", v.CompletionProvider)
	enc.ObjectKeyOmitEmpty("signatureHelpProvider", v.SignatureHelpProvider)
	enc.BoolKeyOmitEmpty("definitionProvider", v.DefinitionProvider)
	enc.AddInterfaceKeyOmitEmpty("typeDefinitionProvider", v.TypeDefinitionProvider)
	enc.AddInterfaceKeyOmitEmpty("implementationProvider", v.ImplementationProvider)
	enc.BoolKeyOmitEmpty("referencesProvider", v.ReferencesProvider)
	enc.BoolKeyOmitEmpty("documentHighlightProvider", v.DocumentHighlightProvider)
	enc.BoolKeyOmitEmpty("documentSymbolProvider", v.DocumentSymbolProvider)
	enc.BoolKeyOmitEmpty("workspaceSymbolProvider", v.WorkspaceSymbolProvider)
	enc.BoolKeyOmitEmpty("codeActionProvider", v.CodeActionProvider)
	enc.ObjectKeyOmitEmpty("codeLensProvider", v.CodeLensProvider)
	enc.BoolKeyOmitEmpty("documentFormattingProvider", v.DocumentFormattingProvider)
	enc.BoolKeyOmitEmpty("documentRangeFormattingProvider", v.DocumentRangeFormattingProvider)
	enc.ObjectKeyOmitEmpty("documentOnTypeFormattingProvider", v.DocumentOnTypeFormattingProvider)
	enc.BoolKeyOmitEmpty("renameProvider", v.RenameProvider)
	enc.ObjectKeyOmitEmpty("documentLinkProvider", v.DocumentLinkProvider)
	enc.AddInterfaceKey("colorProvider", v.ColorProvider)
	enc.AddInterfaceKey("foldingRangeProvider", v.FoldingRangeProvider)
	enc.ObjectKeyOmitEmpty("executeCommandProvider", v.ExecuteCommandProvider)
	enc.ObjectKeyOmitEmpty("workspace", v.Workspace)
	enc.AddInterfaceKeyOmitEmpty("experimental", v.Experimental)
}

// IsNil returns wether the structure is nil value or not.
func (v *ServerCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentLinkRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	case "resolveProvider":
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentLinkRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentLinkRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey("documentSelector", v.DocumentSelector)
	enc.BoolKeyOmitEmpty("resolveProvider", v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentLinkRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *InitializedParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *InitializedParams) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *InitializedParams) MarshalJSONObject(enc *gojay.Encoder) {}

// IsNil returns wether the structure is nil value or not.
func (v *InitializedParams) IsNil() bool { return v == nil }
