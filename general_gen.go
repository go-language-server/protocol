// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *InitializeParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "processID":
		return dec.Float64(v.ProcessID)
	case "rootPath":
		return dec.String(v.RootPath)
	case "rootURI":
		return dec.String((*string)(v.RootURI))
	case "capabilities":
		if v.Capabilities == nil {
			v.Capabilities = &ClientCapabilities{}
		}
		return dec.Object(v.Capabilities)
	case "trace":
		return dec.String(&v.Trace)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *InitializeParams) NKeys() int { return 5 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *InitializeParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key("processID", *v.ProcessID)
	enc.StringKey("rootPath", *v.RootPath)
	enc.StringKey("rootURI", string(*v.RootURI))
	enc.ObjectKey("capabilities", v.Capabilities)
	enc.StringKey("trace", v.Trace)
}

// IsNil returns wether the structure is nil value or not
func (v *InitializeParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentChanges":
		return dec.Bool(&v.DocumentChanges)
	case "failureHandling":
		return dec.String(&v.FailureHandling)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("documentChanges", v.DocumentChanges)
	enc.StringKey("failureHandling", v.FailureHandling)
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceClientCapabilitiesWorkspaceEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceClientCapabilitiesDidChangeConfiguration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceClientCapabilitiesDidChangeWatchedFiles) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceClientCapabilitiesSymbol) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceClientCapabilitiesSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKey("symbolKind", v.SymbolKind)
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceClientCapabilitiesSymbol) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WorkspaceClientCapabilitiesSymbolKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceClientCapabilitiesSymbolKind) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceClientCapabilitiesSymbolKind) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceClientCapabilitiesSymbolKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WorkspaceClientCapabilitiesExecuteCommand) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceClientCapabilitiesExecuteCommand) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceClientCapabilitiesExecuteCommand) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceClientCapabilitiesExecuteCommand) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceClientCapabilities) NKeys() int { return 8 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("applyEdit", v.ApplyEdit)
	enc.ObjectKey("workspaceEdit", v.WorkspaceEdit)
	enc.ObjectKey("didChangeConfiguration", v.DidChangeConfiguration)
	enc.ObjectKey("didChangeWatchedFiles", v.DidChangeWatchedFiles)
	enc.ObjectKey("symbol", v.Symbol)
	enc.ObjectKey("executeCommand", v.ExecuteCommand)
	enc.BoolKey("workspaceFolders", v.WorkspaceFolders)
	enc.BoolKey("configuration", v.Configuration)
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesSynchronization) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesSynchronization) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("didSave", v.DidSave)
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKey("willSave", v.WillSave)
	enc.BoolKey("willSaveWaitUntil", v.WillSaveWaitUntil)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesSynchronization) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesCompletion) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesCompletion) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKey("completionItem", v.CompletionItem)
	enc.IntKey("completionItemKind", int(*v.CompletionItemKind))
	enc.BoolKey("contextSupport", v.ContextSupport)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesCompletion) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesCompletionItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "snippetSupport":
		return dec.Bool(&v.SnippetSupport)
	case "commitCharactersSupport":
		return dec.Bool(&v.CommitCharactersSupport)
	case "deprecatedSupport":
		return dec.Bool(&v.DeprecatedSupport)
	case "preselectSupport":
		return dec.Bool(&v.PreselectSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesCompletionItem) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesCompletionItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("snippetSupport", v.SnippetSupport)
	enc.BoolKey("commitCharactersSupport", v.CommitCharactersSupport)
	enc.BoolKey("deprecatedSupport", v.DeprecatedSupport)
	enc.BoolKey("preselectSupport", v.PreselectSupport)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesCompletionItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesHover) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesHover) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesHover) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesHover) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesSignatureHelp) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesSignatureHelp) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKey("signatureInformation", v.SignatureInformation)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesSignatureHelp) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesSignatureInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesSignatureInformation) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesSignatureInformation) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesSignatureInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesReferences) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesReferences) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesReferences) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesReferences) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesDocumentHighlight) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesDocumentHighlight) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesDocumentHighlight) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesDocumentHighlight) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesDocumentSymbol) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesDocumentSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKey("symbolKind", v.SymbolKind)
	enc.BoolKey("hierarchicalDocumentSymbolSupport", v.HierarchicalDocumentSymbolSupport)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesDocumentSymbol) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesFormatting) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesRangeFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesRangeFormatting) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesRangeFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesRangeFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesOnTypeFormatting) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesDeclaration) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "linkSupport":
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesDeclaration) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesDeclaration) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKey("linkSupport", v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesDeclaration) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesDefinition) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "linkSupport":
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesDefinition) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesDefinition) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKey("linkSupport", v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesDefinition) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesTypeDefinition) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "linkSupport":
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesTypeDefinition) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesTypeDefinition) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKey("linkSupport", v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesTypeDefinition) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesImplementation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "linkSupport":
		return dec.Bool(&v.LinkSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesImplementation) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesImplementation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKey("linkSupport", v.LinkSupport)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesImplementation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesCodeAction) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesCodeAction) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.ObjectKey("codeActionLiteralSupport", v.CodeActionLiteralSupport)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesCodeAction) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("codeActionKind", v.CodeActionKind)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesCodeActionLiteralSupport) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesCodeActionKind) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesCodeActionKind) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesCodeActionKind) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesCodeActionKind) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesCodeLens) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesCodeLens) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesCodeLens) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesCodeLens) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesDocumentLink) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesDocumentLink) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesDocumentLink) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesDocumentLink) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesColorProvider) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesColorProvider) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesColorProvider) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesColorProvider) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesRename) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "dynamicRegistration":
		return dec.Bool(&v.DynamicRegistration)
	case "prepareSupport":
		return dec.Bool(&v.PrepareSupport)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesRename) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesRename) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.BoolKey("prepareSupport", v.PrepareSupport)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesRename) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "relatedInformation":
		return dec.Bool(&v.RelatedInformation)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("relatedInformation", v.RelatedInformation)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesPublishDiagnostics) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilitiesFoldingRange) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilitiesFoldingRange) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("dynamicRegistration", v.DynamicRegistration)
	enc.Float64Key("rangeLimit", v.RangeLimit)
	enc.BoolKey("lineFoldingOnly", v.LineFoldingOnly)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilitiesFoldingRange) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentClientCapabilities) NKeys() int { return 21 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("synchronization", v.Synchronization)
	enc.ObjectKey("completion", v.Completion)
	enc.ObjectKey("hover", v.Hover)
	enc.ObjectKey("signatureHelp", v.SignatureHelp)
	enc.ObjectKey("references", v.References)
	enc.ObjectKey("documentHighlight", v.DocumentHighlight)
	enc.ObjectKey("documentSymbol", v.DocumentSymbol)
	enc.ObjectKey("formatting", v.Formatting)
	enc.ObjectKey("rangeFormatting", v.RangeFormatting)
	enc.ObjectKey("onTypeFormatting", v.OnTypeFormatting)
	enc.ObjectKey("declaration", v.Declaration)
	enc.ObjectKey("definition", v.Definition)
	enc.ObjectKey("typeDefinition", v.TypeDefinition)
	enc.ObjectKey("implementation", v.Implementation)
	enc.ObjectKey("codeAction", v.CodeAction)
	enc.ObjectKey("codeLens", v.CodeLens)
	enc.ObjectKey("documentLink", v.DocumentLink)
	enc.ObjectKey("colorProvider", v.ColorProvider)
	enc.ObjectKey("rename", v.Rename)
	enc.ObjectKey("publishDiagnostics", v.PublishDiagnostics)
	enc.ObjectKey("foldingRange", v.FoldingRange)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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
		return dec.DecodeInterface(&v.Experimental)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ClientCapabilities) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ClientCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("workspace", v.Workspace)
	enc.ObjectKey("textDocument", v.TextDocument)
	enc.AddInterfaceKey("experimental", v.Experimental)
}

// IsNil returns wether the structure is nil value or not
func (v *ClientCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *InitializeResult) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "capabilities":
		if v.Capabilities == nil {
			v.Capabilities = &ServerCapabilities{}
		}
		return dec.Object(v.Capabilities)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *InitializeResult) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *InitializeResult) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("capabilities", v.Capabilities)
}

// IsNil returns wether the structure is nil value or not
func (v *InitializeResult) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *InitializeError) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "retry":
		return dec.Bool(&v.Retry)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *InitializeError) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *InitializeError) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("retry", v.Retry)
}

// IsNil returns wether the structure is nil value or not
func (v *InitializeError) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *CompletionOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "resolveProvider":
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *CompletionOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *CompletionOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("resolveProvider", v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not
func (v *CompletionOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *SignatureHelpOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *SignatureHelpOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *SignatureHelpOptions) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *SignatureHelpOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *CodeActionOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *CodeActionOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *CodeActionOptions) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *CodeActionOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *CodeLensOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "resolveProvider":
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *CodeLensOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *CodeLensOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("resolveProvider", v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not
func (v *CodeLensOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DocumentOnTypeFormattingOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "firstTriggerCharacter":
		return dec.String(&v.FirstTriggerCharacter)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DocumentOnTypeFormattingOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DocumentOnTypeFormattingOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("firstTriggerCharacter", v.FirstTriggerCharacter)
}

// IsNil returns wether the structure is nil value or not
func (v *DocumentOnTypeFormattingOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *RenameOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "prepareProvider":
		return dec.Bool(&v.PrepareProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *RenameOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *RenameOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("prepareProvider", v.PrepareProvider)
}

// IsNil returns wether the structure is nil value or not
func (v *RenameOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DocumentLinkOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "resolveProvider":
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DocumentLinkOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DocumentLinkOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("resolveProvider", v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not
func (v *DocumentLinkOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ExecuteCommandOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ExecuteCommandOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ExecuteCommandOptions) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *ExecuteCommandOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *SaveOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "includeText":
		return dec.Bool(&v.IncludeText)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *SaveOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *SaveOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("includeText", v.IncludeText)
}

// IsNil returns wether the structure is nil value or not
func (v *SaveOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ColorProviderOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ColorProviderOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ColorProviderOptions) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *ColorProviderOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *FoldingRangeProviderOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *FoldingRangeProviderOptions) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *FoldingRangeProviderOptions) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *FoldingRangeProviderOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentSyncOptions) NKeys() int { return 5 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentSyncOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("openClose", v.OpenClose)
	enc.Float64Key("change", v.Change)
	enc.BoolKey("willSave", v.WillSave)
	enc.BoolKey("willSaveWaitUntil", v.WillSaveWaitUntil)
	enc.ObjectKey("save", v.Save)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentSyncOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *StaticRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "id":
		return dec.String(&v.ID)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *StaticRegistrationOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *StaticRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("id", v.ID)
}

// IsNil returns wether the structure is nil value or not
func (v *StaticRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ServerCapabilitiesWorkspace) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "workspaceFolders":
		if v.WorkspaceFolders == nil {
			v.WorkspaceFolders = &ServerCapabilitiesWorkspaceFolders{}
		}
		return dec.Object(v.WorkspaceFolders)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ServerCapabilitiesWorkspace) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ServerCapabilitiesWorkspace) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("workspaceFolders", v.WorkspaceFolders)
}

// IsNil returns wether the structure is nil value or not
func (v *ServerCapabilitiesWorkspace) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ServerCapabilitiesWorkspaceFolders) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "supported":
		return dec.Bool(&v.Supported)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ServerCapabilitiesWorkspaceFolders) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ServerCapabilitiesWorkspaceFolders) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("supported", v.Supported)
}

// IsNil returns wether the structure is nil value or not
func (v *ServerCapabilitiesWorkspaceFolders) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
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

// NKeys returns the number of keys to unmarshal
func (v *ServerCapabilities) NKeys() int { return 23 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ServerCapabilities) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddInterfaceKey("textDocumentSync", v.TextDocumentSync)
	enc.BoolKey("hoverProvider", v.HoverProvider)
	enc.ObjectKey("completionProvider", v.CompletionProvider)
	enc.ObjectKey("signatureHelpProvider", v.SignatureHelpProvider)
	enc.BoolKey("definitionProvider", v.DefinitionProvider)
	enc.AddInterfaceKey("typeDefinitionProvider", v.TypeDefinitionProvider)
	enc.AddInterfaceKey("implementationProvider", v.ImplementationProvider)
	enc.BoolKey("referencesProvider", v.ReferencesProvider)
	enc.BoolKey("documentHighlightProvider", v.DocumentHighlightProvider)
	enc.BoolKey("documentSymbolProvider", v.DocumentSymbolProvider)
	enc.BoolKey("workspaceSymbolProvider", v.WorkspaceSymbolProvider)
	enc.BoolKey("codeActionProvider", v.CodeActionProvider)
	enc.ObjectKey("codeLensProvider", v.CodeLensProvider)
	enc.BoolKey("documentFormattingProvider", v.DocumentFormattingProvider)
	enc.BoolKey("documentRangeFormattingProvider", v.DocumentRangeFormattingProvider)
	enc.ObjectKey("documentOnTypeFormattingProvider", v.DocumentOnTypeFormattingProvider)
	enc.BoolKey("renameProvider", v.RenameProvider)
	enc.ObjectKey("documentLinkProvider", v.DocumentLinkProvider)
	enc.AddInterfaceKey("colorProvider", v.ColorProvider)
	enc.AddInterfaceKey("foldingRangeProvider", v.FoldingRangeProvider)
	enc.ObjectKey("executeCommandProvider", v.ExecuteCommandProvider)
	enc.ObjectKey("workspace", v.Workspace)
	enc.AddInterfaceKey("experimental", v.Experimental)
}

// IsNil returns wether the structure is nil value or not
func (v *ServerCapabilities) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DocumentLinkRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "resolveProvider":
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DocumentLinkRegistrationOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DocumentLinkRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("resolveProvider", v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not
func (v *DocumentLinkRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *InitializedParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *InitializedParams) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *InitializedParams) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *InitializedParams) IsNil() bool { return v == nil }
