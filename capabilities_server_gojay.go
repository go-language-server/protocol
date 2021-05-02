// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"github.com/francoispqt/gojay"
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
func (v *CodeActionClientCapabilitiesResolveSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyProperties, (*Strings)(&v.Properties))
}

// IsNil implements gojay.MarshalerJSONObject.
func (v *CodeActionClientCapabilitiesResolveSupport) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay.UnmarshalerJSONObject.
func (v *CodeActionClientCapabilitiesResolveSupport) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyProperties {
		return dec.Array((*Strings)(&v.Properties))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionClientCapabilitiesResolveSupport) NKeys() int { return 1 }

// compile time check whether the CodeActionClientCapabilitiesResolveSupport implements a gojay.MarshalerJSONObject and gojay.UnmarshalerJSONObject interfaces.
var (
	_ gojay.MarshalerJSONObject   = (*CodeActionClientCapabilitiesResolveSupport)(nil)
	_ gojay.UnmarshalerJSONObject = (*CodeActionClientCapabilitiesResolveSupport)(nil)
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

	enc.ObjectKeyOmitEmpty(keyCompletionProvider, v.CompletionProvider)

	switch v.HoverProvider.(type) {
	case bool:
		enc.BoolKey(keyHoverProvider, v.HoverProvider.(bool))
	case *HoverOptions:
		enc.ObjectKey(keyHoverProvider, v.HoverProvider.(*HoverOptions))
	}

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

	switch v.CodeActionProvider.(type) {
	case bool:
		enc.BoolKey(keyCodeActionProvider, v.CodeActionProvider.(bool))
	case *CodeActionOptions:
		enc.ObjectKey(keyCodeActionProvider, v.CodeActionProvider.(*CodeActionOptions))
	}

	enc.ObjectKeyOmitEmpty(keyCodeLensProvider, v.CodeLensProvider)
	enc.ObjectKeyOmitEmpty(keyDocumentLinkProvider, v.DocumentLinkProvider)

	switch v.ColorProvider.(type) {
	case bool:
		enc.BoolKey(keyColorProvider, v.ColorProvider.(bool))
	case *DocumentColorOptions:
		enc.ObjectKey(keyColorProvider, v.ColorProvider.(*DocumentColorOptions))
	case *DocumentColorRegistrationOptions:
		enc.ObjectKey(keyColorProvider, v.ColorProvider.(*DocumentColorRegistrationOptions))
	}

	switch v.WorkspaceSymbolProvider.(type) {
	case bool:
		enc.BoolKey(keyWorkspaceSymbolProvider, v.WorkspaceSymbolProvider.(bool))
	case *WorkspaceSymbolOptions:
		enc.ObjectKey(keyWorkspaceSymbolProvider, v.WorkspaceSymbolProvider.(*WorkspaceSymbolOptions))
	}

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

	switch v.CallHierarchyProvider.(type) {
	case bool:
		enc.BoolKey(keyCallHierarchyProvider, v.CallHierarchyProvider.(bool))
	case *CallHierarchyOptions:
		enc.ObjectKey(keyCallHierarchyProvider, v.CallHierarchyProvider.(*CallHierarchyOptions))
	case *CallHierarchyRegistrationOptions:
		enc.ObjectKey(keyCallHierarchyProvider, v.CallHierarchyProvider.(*CallHierarchyRegistrationOptions))
	}

	switch v.LinkedEditingRangeProvider.(type) {
	case bool:
		enc.BoolKey(keyLinkedEditingRangeProvider, v.LinkedEditingRangeProvider.(bool))
	case *LinkedEditingRangeOptions:
		enc.ObjectKey(keyLinkedEditingRangeProvider, v.LinkedEditingRangeProvider.(*LinkedEditingRangeOptions))
	case *LinkedEditingRangeRegistrationOptions:
		enc.ObjectKey(keyLinkedEditingRangeProvider, v.LinkedEditingRangeProvider.(*LinkedEditingRangeRegistrationOptions))
	}

	switch v.SemanticTokensProvider.(type) {
	case *SemanticTokensOptions:
		enc.ObjectKey(keySemanticTokensProvider, v.SemanticTokensProvider.(*SemanticTokensOptions))
	case *SemanticTokensRegistrationOptions:
		enc.ObjectKey(keySemanticTokensProvider, v.SemanticTokensProvider.(*SemanticTokensRegistrationOptions))
	}

	enc.ObjectKeyOmitEmpty(keyWorkspace, v.Workspace)

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
	case keyCompletionProvider:
		if v.CompletionProvider == nil {
			v.CompletionProvider = &CompletionOptions{}
		}
		return dec.Object(v.CompletionProvider)
	case keyHoverProvider:
		return dec.Interface(&v.HoverProvider)
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
	case keyCodeActionProvider:
		return dec.Interface(&v.CodeActionProvider)
	case keyCodeLensProvider:
		if v.CodeLensProvider == nil {
			v.CodeLensProvider = &CodeLensOptions{}
		}
		return dec.Object(v.CodeLensProvider)
	case keyDocumentLinkProvider:
		if v.DocumentLinkProvider == nil {
			v.DocumentLinkProvider = &DocumentLinkOptions{}
		}
		return dec.Object(v.DocumentLinkProvider)
	case keyColorProvider:
		return dec.Interface(&v.ColorProvider)
	case keyWorkspaceSymbolProvider:
		return dec.Interface(&v.WorkspaceSymbolProvider)
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
