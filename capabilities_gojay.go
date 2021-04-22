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

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyGroupsOnLabel, v.GroupsOnLabel)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport) IsNil() bool {
	return v == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyGroupsOnLabel {
		return dec.Bool(&v.GroupsOnLabel)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDocumentChanges, v.DocumentChanges)
	enc.StringKeyOmitEmpty(keyFailureHandling, v.FailureHandling)
	enc.ArrayKeyOmitEmpty(keyResourceOperations, (*Strings)(&v.ResourceOperations))
	enc.BoolKeyOmitEmpty(keyNormalizesLineEndings, v.NormalizesLineEndings)
	enc.ObjectKeyOmitEmpty(keyChangeAnnotationSupport, v.ChangeAnnotationSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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
	case keyNormalizesLineEndings:
		return dec.Bool(&v.NormalizesLineEndings)
	case keyChangeAnnotationSupport:
		if v.ChangeAnnotationSupport == nil {
			v.ChangeAnnotationSupport = &WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport{}
		}
		return dec.Object(v.ChangeAnnotationSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) NKeys() int { return 5 }

// compile time check whether the WorkspaceClientCapabilitiesWorkspaceEdit implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesWorkspaceEdit)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesWorkspaceEdit)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesDidChangeConfiguration implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesDidChangeConfiguration)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesDidChangeConfiguration)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesDidChangeWatchedFiles implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesDidChangeWatchedFiles)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesDidChangeWatchedFiles)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keySymbolKind, v.SymbolKind)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbol) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the WorkspaceClientCapabilitiesSymbol implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesSymbol)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesSymbol)(nil)
)

// SymbolKinds represents a slice of SymbolKind.
type SymbolKinds []SymbolKind

// compile time check whether the SymbolKinds implements a gojay.MarshalerJSONArray and gojay.UnmarshalerJSONArray interfaces.
var (
	_ gojay.MarshalerJSONArray   = (*SymbolKinds)(nil)
	_ gojay.UnmarshalerJSONArray = (*SymbolKinds)(nil)
)

// MarshalJSONArray implements gojay.MarshalerJSONArray.
func (v SymbolKinds) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range v {
		enc.Float64(float64(v[i]))
	}
}

// IsNil implements gojay.MarshalerJSONObject.
func (v SymbolKinds) IsNil() bool { return len(v) == 0 }

// UnmarshalJSONArray decodes JSON array elements into slice.
func (v *SymbolKinds) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value SymbolKind
	if err := dec.Float64((*float64)(&value)); err != nil {
		return err
	}
	*v = append(*v, value)
	return nil
}

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbolKind) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyValueSet, (*SymbolKinds)(&v.ValueSet))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbolKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSymbolKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		return dec.Array((*SymbolKinds)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesSymbolKind) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesSymbolKind implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesSymbolKind)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesSymbolKind)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesExecuteCommand) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesExecuteCommand) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesExecuteCommand) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesExecuteCommand) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesExecuteCommand implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesExecuteCommand)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesExecuteCommand)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSemanticTokens) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyRefreshSupport, v.RefreshSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSemanticTokens) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSemanticTokens) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyRefreshSupport {
		return dec.Bool(&v.RefreshSupport)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSemanticTokens) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesSemanticTokens implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesSemanticTokens)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesSemanticTokens)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesCodeLens) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyRefreshSupport, v.RefreshSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesCodeLens) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesCodeLens) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyRefreshSupport {
		return dec.Bool(&v.RefreshSupport)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesCodeLens) NKeys() int { return 1 }

// compile time check whether the WorkspaceClientCapabilitiesCodeLens implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesCodeLens)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesCodeLens)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DocumentSymbolOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
	enc.StringKeyOmitEmpty(keyLabel, v.Label)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DocumentSymbolOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DocumentSymbolOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	case keyLabel:
		return dec.String(&v.Label)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DocumentSymbolOptions) NKeys() int { return 2 }

// compile time check whether the DocumentSymbolOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentSymbolOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentSymbolOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesFileOperations) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyDidCreate, v.DidCreate)
	enc.BoolKeyOmitEmpty(keyWillCreate, v.WillCreate)
	enc.BoolKeyOmitEmpty(keyDidRename, v.DidRename)
	enc.BoolKeyOmitEmpty(keyWillRename, v.WillRename)
	enc.BoolKeyOmitEmpty(keyDidDelete, v.DidDelete)
	enc.BoolKeyOmitEmpty(keyWillDelete, v.WillDelete)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesFileOperations) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesFileOperations) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyDidCreate:
		return dec.Bool(&v.DidCreate)
	case keyWillCreate:
		return dec.Bool(&v.WillCreate)
	case keyDidRename:
		return dec.Bool(&v.DidRename)
	case keyWillRename:
		return dec.Bool(&v.WillRename)
	case keyDidDelete:
		return dec.Bool(&v.DidDelete)
	case keyWillDelete:
		return dec.Bool(&v.WillDelete)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilitiesFileOperations) NKeys() int { return 7 }

// compile time check whether the WorkspaceClientCapabilitiesFileOperations implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesFileOperations)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesFileOperations)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyApplyEdit, v.ApplyEdit)
	enc.ObjectKeyOmitEmpty(keyWorkspaceEdit, v.WorkspaceEdit)
	enc.ObjectKeyOmitEmpty(keyDidChangeConfiguration, v.DidChangeConfiguration)
	enc.ObjectKeyOmitEmpty(keyDidChangeWatchedFiles, v.DidChangeWatchedFiles)
	enc.ObjectKeyOmitEmpty(keySymbol, v.Symbol)
	enc.ObjectKeyOmitEmpty(keyExecuteCommand, v.ExecuteCommand)
	enc.BoolKeyOmitEmpty(keyWorkspaceFolders, v.WorkspaceFolders)
	enc.BoolKeyOmitEmpty(keyConfiguration, v.Configuration)
	enc.ObjectKeyOmitEmpty(keySemanticTokens, v.SemanticTokens)
	enc.ObjectKeyOmitEmpty(keyCodeLens, v.CodeLens)
	enc.ObjectKeyOmitEmpty(keyFileOperations, v.FileOperations)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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
	case keySemanticTokens:
		if v.SemanticTokens == nil {
			v.SemanticTokens = &WorkspaceClientCapabilitiesSemanticTokens{}
		}
		return dec.Object(v.SemanticTokens)
	case keyCodeLens:
		if v.CodeLens == nil {
			v.CodeLens = &WorkspaceClientCapabilitiesCodeLens{}
		}
		return dec.Object(v.CodeLens)
	case keyFileOperations:
		if v.FileOperations == nil {
			v.FileOperations = &WorkspaceClientCapabilitiesFileOperations{}
		}
		return dec.Object(v.FileOperations)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceClientCapabilities) NKeys() int { return 11 }

// compile time check whether the WorkspaceClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilities)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSynchronization) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDidSave, v.DidSave)
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyWillSave, v.WillSave)
	enc.BoolKeyOmitEmpty(keyWillSaveWaitUntil, v.WillSaveWaitUntil)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSynchronization) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the TextDocumentClientCapabilitiesSynchronization implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesSynchronization)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesSynchronization)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletion) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keyCompletionItem, v.CompletionItem)
	enc.ObjectKeyOmitEmpty(keyCompletionItemKind, v.CompletionItemKind)
	enc.BoolKeyOmitEmpty(keyContextSupport, v.ContextSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletion) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the TextDocumentClientCapabilitiesCompletion implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCompletion)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCompletion)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemKind) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyValueSet, (*CompletionItemKinds)(&v.ValueSet))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		return dec.Array((*CompletionItemKinds)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCompletionItemKind) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCompletion implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCompletionItemKind)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCompletionItemKind)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemTagSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyValueSet, (*CompletionItemTags)(&v.ValueSet))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemTagSupport) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemTagSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		return dec.Array((*CompletionItemTags)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCompletionItemTagSupport) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCompletionItemTagSupport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCompletionItemTagSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCompletionItemTagSupport)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemResolveSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyProperties, (*Strings)(&v.Properties))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemResolveSupport) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemResolveSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyProperties {
		return dec.Array((*Strings)(&v.Properties))
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemResolveSupport) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCompletionItemResolveSupport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCompletionItemResolveSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCompletionItemResolveSupport)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyValueSet, (*InsertTextModes)(&v.ValueSet))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport) IsNil() bool {
	return v == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		return dec.Array((*InsertTextModes)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keySnippetSupport, v.SnippetSupport)
	enc.BoolKeyOmitEmpty(keyCommitCharactersSupport, v.CommitCharactersSupport)
	enc.ArrayKeyOmitEmpty(keyDocumentationFormat, (*MarkupKinds)(&v.DocumentationFormat))
	enc.BoolKeyOmitEmpty(keyDeprecatedSupport, v.DeprecatedSupport)
	enc.BoolKeyOmitEmpty(keyPreselectSupport, v.PreselectSupport)
	enc.ObjectKeyOmitEmpty(keyTagSupport, v.TagSupport)
	enc.BoolKeyOmitEmpty(keyInsertReplaceSupport, v.InsertReplaceSupport)
	enc.ObjectKeyOmitEmpty(keyResolveSupport, v.ResolveSupport)
	enc.ObjectKeyOmitEmpty(keyInsertTextModeSupport, v.InsertTextModeSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCompletionItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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
	case keyTagSupport:
		if v.TagSupport == nil {
			v.TagSupport = &TextDocumentClientCapabilitiesCompletionItemTagSupport{}
		}
		return dec.Object(v.TagSupport)
	case keyInsertReplaceSupport:
		return dec.Bool(&v.InsertReplaceSupport)
	case keyResolveSupport:
		if v.ResolveSupport == nil {
			v.ResolveSupport = &TextDocumentClientCapabilitiesCompletionItemResolveSupport{}
		}
		return dec.Object(v.ResolveSupport)
	case keyInsertTextModeSupport:
		if v.InsertTextModeSupport == nil {
			v.InsertTextModeSupport = &TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport{}
		}
		return dec.Object(v.InsertTextModeSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCompletionItem) NKeys() int { return 9 }

// compile time check whether the TextDocumentClientCapabilitiesCompletionItem implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCompletionItem)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCompletionItem)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesHover) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ArrayKeyOmitEmpty(keyContentFormat, (*MarkupKinds)(&v.ContentFormat))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesHover) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the TextDocumentClientCapabilitiesHover implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesHover)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesHover)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureHelp) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keySignatureInformation, v.SignatureInformation)
	enc.BoolKeyOmitEmpty(keyContextSupport, v.ContextSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureHelp) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureHelp) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keySignatureInformation:
		if v.SignatureInformation == nil {
			v.SignatureInformation = &TextDocumentClientCapabilitiesSignatureInformation{}
		}
		return dec.Object(v.SignatureInformation)
	case keyContextSupport:
		return dec.Bool(&v.ContextSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSignatureHelp) NKeys() int { return 3 }

// compile time check whether the TextDocumentClientCapabilitiesSignatureHelp implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesSignatureHelp)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesSignatureHelp)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyDocumentationFormat, (*MarkupKinds)(&v.DocumentationFormat))
	enc.ObjectKeyOmitEmpty(keyParameterInformation, v.ParameterInformation)
	enc.BoolKeyOmitEmpty(keyActiveParameterSupport, v.ActiveParameterSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSignatureInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentationFormat:
		return dec.Array((*MarkupKinds)(&v.DocumentationFormat))
	case keyParameterInformation:
		if v.ParameterInformation == nil {
			v.ParameterInformation = &TextDocumentClientCapabilitiesParameterInformation{}
		}
		return dec.Object(v.ParameterInformation)
	case keyActiveParameterSupport:
		return dec.Bool(&v.ActiveParameterSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSignatureInformation) NKeys() int { return 3 }

// compile time check whether the TextDocumentClientCapabilitiesSignatureInformation implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesSignatureInformation)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesSignatureInformation)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesParameterInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyLabelOffsetSupport, v.LabelOffsetSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesParameterInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesParameterInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyLabelOffsetSupport {
		return dec.Bool(&v.LabelOffsetSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesParameterInformation) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesSignatureInformation implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesParameterInformation)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesParameterInformation)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesReferences) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesReferences) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesReferences) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesReferences) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesReferences implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesReferences)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesReferences)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentHighlight) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesDocumentHighlight implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDocumentHighlight)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDocumentHighlight)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keySymbolKind, v.SymbolKind)
	enc.BoolKeyOmitEmpty(keyHierarchicalDocumentSymbolSupport, v.HierarchicalDocumentSymbolSupport)
	enc.ObjectKeyOmitEmpty(keyTagSupport, v.TagSupport)
	enc.BoolKeyOmitEmpty(keyLabelSupport, v.LabelSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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
	case keyTagSupport:
		if v.TagSupport == nil {
			v.TagSupport = &TextDocumentClientCapabilitiesDocumentSymbolTagSupport{}
		}
		return dec.Object(v.TagSupport)
	case keyLabelSupport:
		return dec.Bool(&v.LabelSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentSymbol) NKeys() int { return 5 }

// compile time check whether the TextDocumentClientCapabilitiesDocumentSymbol implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDocumentSymbol)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDocumentSymbol)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentSymbolTagSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyValueSet, (*SymbolTags)(&v.ValueSet))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentSymbolTagSupport) IsNil() bool {
	return v == nil
}

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentSymbolTagSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		return dec.Array((*SymbolTags)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentSymbolTagSupport) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesDocumentSymbolTagSupport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDocumentSymbolTagSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDocumentSymbolTagSupport)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesFormatting) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesFormatting implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesFormatting)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesFormatting)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRangeFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRangeFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRangeFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesRangeFormatting) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesFormatting implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesRangeFormatting)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesRangeFormatting)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesOnTypeFormatting implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesOnTypeFormatting)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesOnTypeFormatting)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDeclaration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyLinkSupport, v.LinkSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDeclaration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the TextDocumentClientCapabilitiesDeclaration implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDeclaration)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDeclaration)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDefinition) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyLinkSupport, v.LinkSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDefinition) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the TextDocumentClientCapabilitiesDefinition implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDefinition)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDefinition)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesTypeDefinition) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyLinkSupport, v.LinkSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesTypeDefinition) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the TextDocumentClientCapabilitiesTypeDefinition implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesTypeDefinition)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesTypeDefinition)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesImplementation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyLinkSupport, v.LinkSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesImplementation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the TextDocumentClientCapabilitiesImplementation implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesImplementation)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesImplementation)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeAction) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKeyOmitEmpty(keyCodeActionLiteralSupport, v.CodeActionLiteralSupport)
	enc.BoolKeyOmitEmpty(keyIsPreferredSupport, v.IsPreferredSupport)
	enc.BoolKeyOmitEmpty(keyDisabledSupport, v.DisabledSupport)
	enc.BoolKeyOmitEmpty(keyDataSupport, v.DataSupport)
	enc.ObjectKeyOmitEmpty(keyResolveSupport, v.ResolveSupport)
	enc.BoolKeyOmitEmpty(keyHonorsChangeAnnotations, v.HonorsChangeAnnotations)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeAction) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeAction) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyCodeActionLiteralSupport:
		if v.CodeActionLiteralSupport == nil {
			v.CodeActionLiteralSupport = &TextDocumentClientCapabilitiesCodeActionLiteralSupport{}
		}
		return dec.Object(v.CodeActionLiteralSupport)
	case keyIsPreferredSupport:
		return dec.Bool(&v.IsPreferredSupport)
	case keyDisabledSupport:
		return dec.Bool(&v.DisabledSupport)
	case keyDataSupport:
		return dec.Bool(&v.DataSupport)
	case keyResolveSupport:
		if v.ResolveSupport == nil {
			v.ResolveSupport = &TextDocumentClientCapabilitiesCodeActionResolveSupport{}
		}
		return dec.Object(v.ResolveSupport)
	case keyHonorsChangeAnnotations:
		return dec.Bool(&v.HonorsChangeAnnotations)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeAction) NKeys() int { return 7 }

// compile time check whether the TextDocumentClientCapabilitiesCodeAction implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCodeAction)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCodeAction)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionResolveSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyProperties, (*Strings)(&v.Properties))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionResolveSupport) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionResolveSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyProperties {
		return dec.Array((*Strings)(&v.Properties))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeActionResolveSupport) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCodeActionResolveSupport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCodeActionResolveSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCodeActionResolveSupport)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyCodeActionKind, v.CodeActionKind)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the TextDocumentClientCapabilitiesCodeActionLiteralSupport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCodeActionLiteralSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCodeActionLiteralSupport)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionKind) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyValueSet, (*CodeActionKinds)(&v.ValueSet))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeActionKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		return dec.Array((*CodeActionKinds)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeActionKind) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCodeActionKind implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCodeActionKind)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCodeActionKind)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyValueSet, (*DiagnosticTags)(&v.ValueSet))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyValueSet {
		if v.ValueSet == nil {
			v.ValueSet = []DiagnosticTag{}
		}
		return dec.Array((*DiagnosticTags)(&v.ValueSet))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCodeActionKind implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeLens) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyTooltipSupport, v.TooltipSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeLens) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCodeLens) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyTooltipSupport:
		return dec.Bool(&v.TooltipSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesCodeLens) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesCodeLens implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCodeLens)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCodeLens)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentLink) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyTooltipSupport, v.TooltipSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentLink) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesDocumentLink) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyTooltipSupport:
		return dec.Bool(&v.TooltipSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesDocumentLink) NKeys() int { return 2 }

// compile time check whether the TextDocumentClientCapabilitiesDocumentLink implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesDocumentLink)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesDocumentLink)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DocumentColorOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DocumentColorOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DocumentColorOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWorkDoneProgress {
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DocumentColorOptions) NKeys() int { return 1 }

// compile time check whether the DocumentColorOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentColorOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentColorOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *DocumentColorRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.StringKeyOmitEmpty(keyID, v.ID)
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *DocumentColorRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *DocumentColorRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		if v.DocumentSelector == nil {
			v.DocumentSelector = DocumentSelector{}
		}
		return dec.Array(&v.DocumentSelector)
	case keyID:
		return dec.String(&v.ID)
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *DocumentColorRegistrationOptions) NKeys() int { return 4 }

// compile time check whether the DocumentColorRegistrationOptions implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*DocumentColorRegistrationOptions)(nil)
	_ gojay.UnmarshalerJSONObject = (*DocumentColorRegistrationOptions)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesColorProvider) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesColorProvider) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesColorProvider) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesColorProvider) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesColorProvider implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesColorProvider)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesColorProvider)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRename) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.BoolKeyOmitEmpty(keyPrepareSupport, v.PrepareSupport)
	enc.Float64KeyOmitEmpty(keyPrepareSupportDefaultBehavior, float64(v.PrepareSupportDefaultBehavior))
	enc.BoolKeyOmitEmpty(keyHonorsChangeAnnotations, v.HonorsChangeAnnotations)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRename) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesRename) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyPrepareSupport:
		return dec.Bool(&v.PrepareSupport)
	case keyPrepareSupportDefaultBehavior:
		return dec.Float64((*float64)(&v.PrepareSupportDefaultBehavior))
	case keyHonorsChangeAnnotations:
		return dec.Bool(&v.HonorsChangeAnnotations)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesRename) NKeys() int { return 4 }

// compile time check whether the TextDocumentClientCapabilitiesRename implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesRename)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesRename)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyRelatedInformation, v.RelatedInformation)
	enc.ObjectKeyOmitEmpty(keyTagSupport, v.TagSupport)
	enc.BoolKeyOmitEmpty(keyVersionSupport, v.VersionSupport)
	enc.BoolKeyOmitEmpty(keyCodeDescriptionSupport, v.CodeDescriptionSupport)
	enc.BoolKeyOmitEmpty(keyDataSupport, v.DataSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRelatedInformation:
		return dec.Bool(&v.RelatedInformation)
	case keyTagSupport:
		if v.TagSupport == nil {
			v.TagSupport = &TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport{}
		}
		return dec.Object(v.TagSupport)
	case keyVersionSupport:
		return dec.Bool(&v.VersionSupport)
	case keyCodeDescriptionSupport:
		return dec.Bool(&v.CodeDescriptionSupport)
	case keyDataSupport:
		return dec.Bool(&v.DataSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) NKeys() int { return 5 }

// compile time check whether the TextDocumentClientCapabilitiesPublishDiagnostics implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesPublishDiagnostics)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesPublishDiagnostics)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFoldingRange) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.Uint32KeyOmitEmpty(keyRangeLimit, v.RangeLimit)
	enc.BoolKeyOmitEmpty(keyLineFoldingOnly, v.LineFoldingOnly)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFoldingRange) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesFoldingRange) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyRangeLimit:
		return dec.Uint32(&v.RangeLimit)
	case keyLineFoldingOnly:
		return dec.Bool(&v.LineFoldingOnly)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesFoldingRange) NKeys() int { return 3 }

// compile time check whether the TextDocumentClientCapabilitiesFoldingRange implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesFoldingRange)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesFoldingRange)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSelectionRange) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSelectionRange) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSelectionRange) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilitiesSelectionRange) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesSelectionRange implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesSelectionRange)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesSelectionRange)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesLinkedEditingRange) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesLinkedEditingRange) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesLinkedEditingRange) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesLinkedEditingRange) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesLinkedEditingRange implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesLinkedEditingRange)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesLinkedEditingRange)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCallHierarchy) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCallHierarchy) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCallHierarchy) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesCallHierarchy) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesCallHierarchy implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesCallHierarchy)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesCallHierarchy)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSemanticTokensRequests) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyRange, v.Range)
	enc.BoolKeyOmitEmpty(keyFull, v.Full)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSemanticTokensRequests) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSemanticTokensRequests) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRange:
		return dec.Bool(&v.Range)
	case keyFull:
		return dec.Bool(&v.Full)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WorkspaceClientCapabilitiesSemanticTokensRequests) NKeys() int { return 2 }

// compile time check whether the WorkspaceClientCapabilitiesSemanticTokensRequests implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WorkspaceClientCapabilitiesSemanticTokensRequests)(nil)
	_ gojay.UnmarshalerJSONObject = (*WorkspaceClientCapabilitiesSemanticTokensRequests)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSemanticTokens) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
	enc.ObjectKey(keyRequests, &v.Requests)
	enc.ArrayKey(keyTokenTypes, (*Strings)(&v.TokenTypes))
	enc.ArrayKey(keyTokenModifiers, (*Strings)(&v.TokenModifiers))
	enc.ArrayKey(keyFormats, (*TokenFormats)(&v.Formats))
	enc.BoolKeyOmitEmpty(keyOverlappingTokenSupport, v.OverlappingTokenSupport)
	enc.BoolKeyOmitEmpty(keyMultilineTokenSupport, v.MultilineTokenSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSemanticTokens) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSemanticTokens) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDynamicRegistration:
		return dec.Bool(&v.DynamicRegistration)
	case keyRequests:
		return dec.Object(&v.Requests)
	case keyTokenTypes:
		return dec.Array((*Strings)(&v.TokenTypes))
	case keyTokenModifiers:
		return dec.Array((*Strings)(&v.TokenModifiers))
	case keyFormats:
		return dec.Array((*TokenFormats)(&v.Formats))
	case keyOverlappingTokenSupport:
		return dec.Bool(&v.OverlappingTokenSupport)
	case keyMultilineTokenSupport:
		return dec.Bool(&v.MultilineTokenSupport)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesSemanticTokens) NKeys() int { return 7 }

// compile time check whether the TextDocumentClientCapabilitiesSemanticTokens implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesSemanticTokens)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesSemanticTokens)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesMoniker) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyDynamicRegistration, v.DynamicRegistration)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesMoniker) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesMoniker) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyDynamicRegistration {
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *TextDocumentClientCapabilitiesMoniker) NKeys() int { return 1 }

// compile time check whether the TextDocumentClientCapabilitiesMoniker implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilitiesMoniker)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilitiesMoniker)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
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
	enc.ObjectKeyOmitEmpty(keyLinkedEditingRange, v.LinkedEditingRange)
	enc.ObjectKeyOmitEmpty(keyCallHierarchy, v.CallHierarchy)
	enc.ObjectKeyOmitEmpty(keySemanticTokens, v.SemanticTokens)
	enc.ObjectKeyOmitEmpty(keyMoniker, v.Moniker)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *TextDocumentClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
//nolint:funlen,gocognit
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
	case keyLinkedEditingRange:
		if v.LinkedEditingRange == nil {
			v.LinkedEditingRange = &TextDocumentClientCapabilitiesLinkedEditingRange{}
		}
		return dec.Object(v.LinkedEditingRange)
	case keyCallHierarchy:
		if v.CallHierarchy == nil {
			v.CallHierarchy = &TextDocumentClientCapabilitiesCallHierarchy{}
		}
		return dec.Object(v.CallHierarchy)
	case keySemanticTokens:
		if v.SemanticTokens == nil {
			v.SemanticTokens = &TextDocumentClientCapabilitiesSemanticTokens{}
		}
		return dec.Object(v.SemanticTokens)
	case keyMoniker:
		if v.Moniker == nil {
			v.Moniker = &TextDocumentClientCapabilitiesMoniker{}
		}
		return dec.Object(v.Moniker)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *TextDocumentClientCapabilities) NKeys() int { return 26 }

// compile time check whether the TextDocumentClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*TextDocumentClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*TextDocumentClientCapabilities)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ClientCapabilitiesShowMessageRequest) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyMessageActionItem, v.MessageActionItem)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ClientCapabilitiesShowMessageRequest) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ClientCapabilitiesShowMessageRequest) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyMessageActionItem {
		if v.MessageActionItem == nil {
			v.MessageActionItem = &ClientCapabilitiesShowMessageRequestMessageActionItem{}
		}
		return dec.Object(v.MessageActionItem)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ClientCapabilitiesShowMessageRequest) NKeys() int { return 1 }

// compile time check whether the ClientCapabilitiesShowMessageRequest implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ClientCapabilitiesShowMessageRequest)(nil)
	_ gojay.UnmarshalerJSONObject = (*ClientCapabilitiesShowMessageRequest)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ClientCapabilitiesShowMessageRequestMessageActionItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyAdditionalPropertiesSupport, v.AdditionalPropertiesSupport)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ClientCapabilitiesShowMessageRequestMessageActionItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ClientCapabilitiesShowMessageRequestMessageActionItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyAdditionalPropertiesSupport {
		return dec.Bool(&v.AdditionalPropertiesSupport)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ClientCapabilitiesShowMessageRequestMessageActionItem) NKeys() int { return 1 }

// compile time check whether the ClientCapabilitiesShowMessageRequestMessageActionItem implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ClientCapabilitiesShowMessageRequestMessageActionItem)(nil)
	_ gojay.UnmarshalerJSONObject = (*ClientCapabilitiesShowMessageRequestMessageActionItem)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ClientCapabilitiesShowDocument) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey(keySupport, v.Support)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ClientCapabilitiesShowDocument) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ClientCapabilitiesShowDocument) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keySupport {
		return dec.Bool(&v.Support)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *ClientCapabilitiesShowDocument) NKeys() int { return 1 }

// compile time check whether the ClientCapabilitiesShowDocument implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ClientCapabilitiesShowDocument)(nil)
	_ gojay.UnmarshalerJSONObject = (*ClientCapabilitiesShowDocument)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *WindowClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKeyOmitEmpty(keyWorkDoneProgress, v.WorkDoneProgress)
	enc.ObjectKeyOmitEmpty(keyShowMessage, v.ShowMessage)
	enc.ObjectKeyOmitEmpty(keyShowDocument, v.ShowDocument)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *WindowClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *WindowClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyWorkDoneProgress:
		return dec.Bool(&v.WorkDoneProgress)
	case keyShowMessage:
		if v.ShowMessage == nil {
			v.ShowMessage = &ClientCapabilitiesShowMessageRequest{}
		}
		return dec.Object(v.ShowMessage)
	case keyShowDocument:
		if v.ShowDocument == nil {
			v.ShowDocument = &ClientCapabilitiesShowDocument{}
		}
		return dec.Object(v.ShowDocument)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *WindowClientCapabilities) NKeys() int { return 3 }

// compile time check whether the WindowClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*WindowClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*WindowClientCapabilities)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *MarkdownClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyParser, v.Parser)
	enc.StringKeyOmitEmpty(keyVersion, v.Version)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *MarkdownClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *MarkdownClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyParser:
		return dec.String(&v.Parser)
	case keyVersion:
		return dec.String(&v.Version)
	}
	return nil
}

// NKeys implements gojay.UnmarshalerJSONObject.
func (v *MarkdownClientCapabilities) NKeys() int { return 2 }

// compile time check whether the MarkdownClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*MarkdownClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*MarkdownClientCapabilities)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *GeneralClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyRegularExpressions, v.RegularExpressions)
	enc.ObjectKeyOmitEmpty(keyMarkdown, v.Markdown)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *GeneralClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *GeneralClientCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRegularExpressions:
		if v.RegularExpressions == nil {
			v.RegularExpressions = &RegularExpressionsClientCapabilities{}
		}
		return dec.Object(v.RegularExpressions)
	case keyMarkdown:
		if v.Markdown == nil {
			v.Markdown = &MarkdownClientCapabilities{}
		}
		return dec.Object(v.Markdown)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *GeneralClientCapabilities) NKeys() int { return 2 }

// compile time check whether the GeneralClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*GeneralClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*GeneralClientCapabilities)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyWorkspace, v.Workspace)
	enc.ObjectKeyOmitEmpty(keyTextDocument, v.TextDocument)
	enc.ObjectKeyOmitEmpty(keyWindow, v.Window)
	enc.ObjectKeyOmitEmpty(keyGeneral, v.General)
	enc.AddInterfaceKeyOmitEmpty(keyExperimental, v.Experimental)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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
	case keyWindow:
		if v.Window == nil {
			v.Window = &WindowClientCapabilities{}
		}
		return dec.Object(v.Window)
	case keyGeneral:
		if v.General == nil {
			v.General = &GeneralClientCapabilities{}
		}
		return dec.Object(v.General)
	case keyExperimental:
		return dec.Interface(&v.Experimental)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ClientCapabilities) NKeys() int { return 5 }

// compile time check whether the ClientCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ClientCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*ClientCapabilities)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspace) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyWorkspaceFolders, v.WorkspaceFolders)
	enc.ObjectKeyOmitEmpty(keyFileOperations, v.FileOperations)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspace) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ServerCapabilitiesWorkspace) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyWorkspaceFolders:
		if v.WorkspaceFolders == nil {
			v.WorkspaceFolders = &ServerCapabilitiesWorkspaceFolders{}
		}
		return dec.Object(v.WorkspaceFolders)
	case keyFileOperations:
		if v.FileOperations == nil {
			v.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		return dec.Object(v.FileOperations)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ServerCapabilitiesWorkspace) NKeys() int { return 2 }

// compile time check whether the ServerCapabilitiesWorkspace implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ServerCapabilitiesWorkspace)(nil)
	_ gojay.UnmarshalerJSONObject = (*ServerCapabilitiesWorkspace)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspaceFolders) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey(keySupported, v.Supported)
	enc.AddInterfaceKeyOmitEmpty(keyChangeNotifications, v.ChangeNotifications)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspaceFolders) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
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

// compile time check whether the ServerCapabilitiesWorkspaceFolders implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ServerCapabilitiesWorkspaceFolders)(nil)
	_ gojay.UnmarshalerJSONObject = (*ServerCapabilitiesWorkspaceFolders)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspaceFileOperations) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKeyOmitEmpty(keyDidCreate, v.DidCreate)
	enc.ObjectKeyOmitEmpty(keyWillCreate, v.WillCreate)
	enc.ObjectKeyOmitEmpty(keyDidRename, v.DidRename)
	enc.ObjectKeyOmitEmpty(keyWillRename, v.WillRename)
	enc.ObjectKeyOmitEmpty(keyDidDelete, v.DidDelete)
	enc.ObjectKeyOmitEmpty(keyWillDelete, v.WillDelete)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ServerCapabilitiesWorkspaceFileOperations) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *ServerCapabilitiesWorkspaceFileOperations) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDidCreate:
		if v.DidCreate == nil {
			v.DidCreate = &FileOperationRegistrationOptions{}
		}
		return dec.Object(v.DidCreate)
	case keyWillCreate:
		if v.WillCreate == nil {
			v.WillCreate = &FileOperationRegistrationOptions{}
		}
		return dec.Object(v.WillCreate)
	case keyDidRename:
		if v.DidRename == nil {
			v.DidRename = &FileOperationRegistrationOptions{}
		}
		return dec.Object(v.DidRename)
	case keyWillRename:
		if v.WillRename == nil {
			v.WillRename = &FileOperationRegistrationOptions{}
		}
		return dec.Object(v.WillRename)
	case keyDidDelete:
		if v.DidDelete == nil {
			v.DidDelete = &FileOperationRegistrationOptions{}
		}
		return dec.Object(v.DidDelete)
	case keyWillDelete:
		if v.WillDelete == nil {
			v.WillDelete = &FileOperationRegistrationOptions{}
		}
		return dec.Object(v.WillDelete)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ServerCapabilitiesWorkspaceFileOperations) NKeys() int { return 6 }

// compile time check whether the ServerCapabilitiesWorkspaceFileOperations implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ServerCapabilitiesWorkspaceFileOperations)(nil)
	_ gojay.UnmarshalerJSONObject = (*ServerCapabilitiesWorkspaceFileOperations)(nil)
)

// MarshalJSONObject implements gojay.MarshalerJSONObject.
//nolint:funlen,gocritic // TODO(zchee): fix gocritic:typeSwitchVar
func (v *ServerCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	switch v.TextDocumentSync.(type) {
	case float64: // TextDocumentSyncKind
		enc.Uint32Key(keyTextDocumentSync, uint32(v.TextDocumentSync.(float64)))
	case TextDocumentSyncKind: // TextDocumentSyncKind
		enc.Uint32Key(keyTextDocumentSync, uint32(v.TextDocumentSync.(TextDocumentSyncKind)))
	case *TextDocumentSyncOptions:
		enc.ObjectKey(keyTextDocumentSync, v.TextDocumentSync.(*TextDocumentSyncOptions))
	}

	switch v.HoverProvider.(type) {
	case bool:
		enc.BoolKey(keyHoverProvider, v.HoverProvider.(bool))
	case *HoverOptions:
		enc.ObjectKey(keyHoverProvider, v.HoverProvider.(*HoverOptions))
	}

	enc.ObjectKeyOmitEmpty(keyCompletionProvider, v.CompletionProvider)
	enc.ObjectKeyOmitEmpty(keySignatureHelpProvider, v.SignatureHelpProvider)

	switch v.DeclarationProvider.(type) {
	case bool:
		enc.BoolKey(keyDeclarationProvider, v.DeclarationProvider.(bool))
	case *DeclarationOptions:
		enc.ObjectKey(keyDeclarationProvider, v.DeclarationProvider.(*DeclarationOptions))
	case *DeclarationRegistrationOptions:
		enc.ObjectKey(keyDeclarationProvider, v.DeclarationProvider.(*DeclarationRegistrationOptions))
	}

	switch v.DefinitionProvider.(type) {
	case bool:
		enc.BoolKey(keyDefinitionProvider, v.DefinitionProvider.(bool))
	case *DefinitionOptions:
		enc.ObjectKey(keyDefinitionProvider, v.DefinitionProvider.(*DefinitionOptions))
	}

	switch v.TypeDefinitionProvider.(type) {
	case bool:
		enc.BoolKey(keyTypeDefinitionProvider, v.TypeDefinitionProvider.(bool))
	case *TypeDefinitionOptions:
		enc.ObjectKey(keyTypeDefinitionProvider, v.TypeDefinitionProvider.(*TypeDefinitionOptions))
	case *TypeDefinitionRegistrationOptions:
		enc.ObjectKey(keyTypeDefinitionProvider, v.TypeDefinitionProvider.(*TypeDefinitionRegistrationOptions))
	}

	switch v.ImplementationProvider.(type) {
	case bool:
		enc.BoolKey(keyImplementationProvider, v.ImplementationProvider.(bool))
	case *ImplementationOptions:
		enc.ObjectKey(keyImplementationProvider, v.ImplementationProvider.(*ImplementationOptions))
	case *ImplementationRegistrationOptions:
		enc.ObjectKey(keyImplementationProvider, v.ImplementationProvider.(*ImplementationRegistrationOptions))
	}

	switch v.ReferencesProvider.(type) {
	case bool:
		enc.BoolKey(keyReferencesProvider, v.ReferencesProvider.(bool))
	case *ReferencesOptions:
		enc.ObjectKey(keyReferencesProvider, v.ReferencesProvider.(*ReferencesOptions))
	}

	switch v.DocumentHighlightProvider.(type) {
	case bool:
		enc.BoolKey(keyDocumentHighlightProvider, v.DocumentHighlightProvider.(bool))
	case *DocumentHighlightOptions:
		enc.ObjectKey(keyDocumentHighlightProvider, v.DocumentHighlightProvider.(*DocumentHighlightOptions))
	}

	switch v.DocumentSymbolProvider.(type) {
	case bool:
		enc.BoolKey(keyDocumentSymbolProvider, v.DocumentSymbolProvider.(bool))
	case *DocumentSymbolOptions:
		enc.ObjectKey(keyDocumentSymbolProvider, v.DocumentSymbolProvider.(*DocumentSymbolOptions))
	}

	switch v.WorkspaceSymbolProvider.(type) {
	case bool:
		enc.BoolKey(keyWorkspaceSymbolProvider, v.WorkspaceSymbolProvider.(bool))
	case *WorkspaceSymbolOptions:
		enc.ObjectKey(keyWorkspaceSymbolProvider, v.WorkspaceSymbolProvider.(*WorkspaceSymbolOptions))
	}

	switch v.CodeActionProvider.(type) {
	case bool:
		enc.BoolKey(keyCodeActionProvider, v.CodeActionProvider.(bool))
	case *CodeActionOptions:
		enc.ObjectKey(keyCodeActionProvider, v.CodeActionProvider.(*CodeActionOptions))
	}

	enc.ObjectKeyOmitEmpty(keyCodeLensProvider, v.CodeLensProvider)

	switch v.DocumentFormattingProvider.(type) {
	case bool:
		enc.BoolKey(keyDocumentFormattingProvider, v.DocumentFormattingProvider.(bool))
	case *DocumentFormattingOptions:
		enc.ObjectKey(keyDocumentFormattingProvider, v.DocumentFormattingProvider.(*DocumentFormattingOptions))
	}

	switch v.DocumentRangeFormattingProvider.(type) {
	case bool:
		enc.BoolKey(keyDocumentRangeFormattingProvider, v.DocumentRangeFormattingProvider.(bool))
	case *DocumentRangeFormattingOptions:
		enc.ObjectKey(keyDocumentRangeFormattingProvider, v.DocumentRangeFormattingProvider.(*DocumentRangeFormattingOptions))
	}

	enc.ObjectKeyOmitEmpty(keyDocumentOnTypeFormattingProvider, v.DocumentOnTypeFormattingProvider)

	switch v.RenameProvider.(type) {
	case bool:
		enc.BoolKey(keyRenameProvider, v.RenameProvider.(bool))
	case *RenameOptions:
		enc.ObjectKey(keyRenameProvider, v.RenameProvider.(*RenameOptions))
	}

	enc.ObjectKeyOmitEmpty(keyDocumentLinkProvider, v.DocumentLinkProvider)

	switch v.ColorProvider.(type) {
	case bool:
		enc.BoolKey(keyColorProvider, v.ColorProvider.(bool))
	case *DocumentColorOptions:
		enc.ObjectKey(keyColorProvider, v.ColorProvider.(*DocumentColorOptions))
	case *DocumentColorRegistrationOptions:
		enc.ObjectKey(keyColorProvider, v.ColorProvider.(*DocumentColorRegistrationOptions))
	}

	switch v.FoldingRangeProvider.(type) {
	case bool:
		enc.BoolKey(keyFoldingRangeProvider, v.FoldingRangeProvider.(bool))
	case *FoldingRangeOptions:
		enc.ObjectKey(keyFoldingRangeProvider, v.FoldingRangeProvider.(*FoldingRangeOptions))
	case *FoldingRangeRegistrationOptions:
		enc.ObjectKey(keyFoldingRangeProvider, v.FoldingRangeProvider.(*FoldingRangeRegistrationOptions))
	}

	switch v.SelectionRangeProvider.(type) {
	case bool:
		enc.BoolKey(keySelectionRangeProvider, v.SelectionRangeProvider.(bool))
	case *EnableSelectionRange:
		enc.BoolKey(keySelectionRangeProvider, bool(*v.SelectionRangeProvider.(*EnableSelectionRange)))
	case *SelectionRangeOptions:
		enc.ObjectKey(keySelectionRangeProvider, v.SelectionRangeProvider.(*SelectionRangeOptions))
	case *SelectionRangeRegistrationOptions:
		enc.ObjectKey(keySelectionRangeProvider, v.SelectionRangeProvider.(*SelectionRangeRegistrationOptions))
	}

	enc.ObjectKeyOmitEmpty(keyExecuteCommandProvider, v.ExecuteCommandProvider)
	enc.ObjectKeyOmitEmpty(keyWorkspace, v.Workspace)

	switch v.LinkedEditingRangeProvider.(type) {
	case bool:
		enc.BoolKey(keyLinkedEditingRangeProvider, v.LinkedEditingRangeProvider.(bool))
	case *LinkedEditingRangeOptions:
		enc.ObjectKey(keyLinkedEditingRangeProvider, v.LinkedEditingRangeProvider.(*LinkedEditingRangeOptions))
	case *LinkedEditingRangeRegistrationOptions:
		enc.ObjectKey(keyLinkedEditingRangeProvider, v.LinkedEditingRangeProvider.(*LinkedEditingRangeRegistrationOptions))
	}

	switch v.CallHierarchyProvider.(type) {
	case bool:
		enc.BoolKey(keyCallHierarchyProvider, v.CallHierarchyProvider.(bool))
	case *CallHierarchyOptions:
		enc.ObjectKey(keyCallHierarchyProvider, v.CallHierarchyProvider.(*CallHierarchyOptions))
	case *CallHierarchyRegistrationOptions:
		enc.ObjectKey(keyCallHierarchyProvider, v.CallHierarchyProvider.(*CallHierarchyRegistrationOptions))
	}

	switch v.SemanticTokensProvider.(type) {
	case *SemanticTokensOptions:
		enc.ObjectKey(keySemanticTokensProvider, v.SemanticTokensProvider.(*SemanticTokensOptions))
	case *SemanticTokensRegistrationOptions:
		enc.ObjectKey(keySemanticTokensProvider, v.SemanticTokensProvider.(*SemanticTokensRegistrationOptions))
	}

	switch v.MonikerProvider.(type) {
	case bool:
		enc.BoolKey(keyMonikerProvider, v.MonikerProvider.(bool))
	case *MonikerOptions:
		enc.ObjectKey(keyMonikerProvider, v.MonikerProvider.(*MonikerOptions))
	case *MonikerRegistrationOptions:
		enc.ObjectKey(keyMonikerProvider, v.MonikerProvider.(*MonikerRegistrationOptions))
	}

	enc.AddInterfaceKeyOmitEmpty(keyExperimental, v.Experimental)
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *ServerCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
//nolint:funlen
func (v *ServerCapabilities) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocumentSync:
		return dec.Interface(&v.TextDocumentSync)
	case keyHoverProvider:
		return dec.Interface(&v.HoverProvider)
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
	case keyDeclarationProvider:
		return dec.Interface(&v.DeclarationProvider)
	case keyDefinitionProvider:
		return dec.Interface(&v.DefinitionProvider)
	case keyTypeDefinitionProvider:
		return dec.Interface(&v.TypeDefinitionProvider)
	case keyImplementationProvider:
		return dec.Interface(&v.ImplementationProvider)
	case keyReferencesProvider:
		return dec.Interface(&v.ReferencesProvider)
	case keyDocumentHighlightProvider:
		return dec.Interface(&v.DocumentHighlightProvider)
	case keyDocumentSymbolProvider:
		return dec.Interface(&v.DocumentSymbolProvider)
	case keyWorkspaceSymbolProvider:
		return dec.Interface(&v.WorkspaceSymbolProvider)
	case keyCodeActionProvider:
		return dec.Interface(&v.CodeActionProvider)
	case keyCodeLensProvider:
		if v.CodeLensProvider == nil {
			v.CodeLensProvider = &CodeLensOptions{}
		}
		return dec.Object(v.CodeLensProvider)
	case keyDocumentFormattingProvider:
		return dec.Interface(&v.DocumentFormattingProvider)
	case keyDocumentRangeFormattingProvider:
		return dec.Interface(&v.DocumentRangeFormattingProvider)
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
	case keyLinkedEditingRangeProvider:
		return dec.Interface(&v.LinkedEditingRangeProvider)
	case keyCallHierarchyProvider:
		return dec.Interface(&v.CallHierarchyProvider)
	case keySemanticTokensProvider:
		return dec.Interface(&v.SemanticTokensProvider)
	case keyMonikerProvider:
		return dec.Interface(&v.MonikerProvider)
	case keyExperimental:
		return dec.Interface(&v.Experimental)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ServerCapabilities) NKeys() int { return 29 }

// compile time check whether the ServerCapabilities implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*ServerCapabilities)(nil)
	_ gojay.UnmarshalerJSONObject = (*ServerCapabilities)(nil)
)
