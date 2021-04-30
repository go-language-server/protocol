// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

// ServerCapabilitiesWorkspaceFolders is the server supports workspace folder.
//
// @since 3.6.0.
type ServerCapabilitiesWorkspaceFolders struct {
	// Supported is the server has support for workspace folders
	Supported bool `json:"supported,omitempty"`

	// ChangeNotifications whether the server wants to receive workspace folder
	// change notifications.
	//
	// If a strings is provided the string is treated as a ID
	// under which the notification is registered on the client
	// side. The ID can be used to unregister for these events
	// using the `client/unregisterCapability` request.
	ChangeNotifications interface{} `json:"changeNotifications,omitempty"` // string | boolean
}

// CompletionTextDocumentClientCapabilitiesItemTagSupport specific capabilities for the "TagSupport" in the "textDocument/completion" request.
//
// @since 3.15.0.
type CompletionTextDocumentClientCapabilitiesItemTagSupport struct {
	// ValueSet is the tags supported by the client.
	//
	// @since 3.15.0.
	ValueSet []CompletionItemTag `json:"valueSet,omitempty"`
}

// CompletionTextDocumentClientCapabilitiesItemResolveSupport specific capabilities for the ResolveSupport in the CompletionTextDocumentClientCapabilitiesItem.
//
// @since 3.16.0.
type CompletionTextDocumentClientCapabilitiesItemResolveSupport struct {
	// Properties is the properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}

// CompletionTextDocumentClientCapabilitiesItemInsertTextModeSupport specific capabilities for the InsertTextModeSupport in the CompletionTextDocumentClientCapabilitiesItem.
//
// @since 3.16.0.
type CompletionTextDocumentClientCapabilitiesItemInsertTextModeSupport struct {
	// ValueSet is the tags supported by the client.
	//
	// @since 3.16.0.
	ValueSet []InsertTextMode `json:"valueSet,omitempty"`
}

// TextDocumentClientCapabilitiesSignatureInformation is the client supports the following "SignatureInformation"
// specific properties.
type TextDocumentClientCapabilitiesSignatureInformation struct {
	// DocumentationFormat is the client supports the follow content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

	// ParameterInformation is the Client capabilities specific to parameter information.
	ParameterInformation *TextDocumentClientCapabilitiesParameterInformation `json:"parameterInformation,omitempty"`

	// ActiveParameterSupport is the client supports the `activeParameter` property on
	// `SignatureInformation` literal.
	//
	// @since 3.16.0.
	ActiveParameterSupport bool `json:"activeParameterSupport,omitempty"`
}

// TextDocumentClientCapabilitiesParameterInformation is the client capabilities specific to parameter information.
type TextDocumentClientCapabilitiesParameterInformation struct {
	// LabelOffsetSupport is the client supports processing label offsets instead of a
	// simple label string.
	//
	// @since 3.14.0.
	LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
}

// DocumentSymbolClientCapabilitiesTagSupport TagSupport in the DocumentSymbolClientCapabilities.
//
// @since 3.16.0.
type DocumentSymbolClientCapabilitiesTagSupport struct {
	// ValueSet is the tags supported by the client.
	ValueSet []SymbolTag `json:"valueSet"`
}

// TextDocumentClientCapabilitiesFormatting capabilities specific to the textDocument/formatting.
type TextDocumentClientCapabilitiesFormatting struct {
	// DynamicRegistration whether formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilitiesRangeFormatting capabilities specific to the "textDocument/rangeFormatting".
type TextDocumentClientCapabilitiesRangeFormatting struct {
	// DynamicRegistration whether range formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilitiesOnTypeFormatting Capabilities specific to the "textDocument/onTypeFormatting".
type TextDocumentClientCapabilitiesOnTypeFormatting struct {
	// DynamicRegistration whether on type formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// CodeActionClientCapabilitiesResolveSupport ResolveSupport in the CodeActionClientCapabilities.
//
// @since 3.16.0.
type CodeActionClientCapabilitiesResolveSupport struct {
	// Properties is the properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}

// CodeActionClientCapabilitiesLiteralSupport is the client support code action literals as a valid response of the "textDocument/codeAction" request.
type CodeActionClientCapabilitiesLiteralSupport struct {
	// CodeActionKind is the code action kind is support with the following value
	// set.
	CodeActionKind *CodeActionClientCapabilitiesKind `json:"codeActionKind"`
}

// CodeActionClientCapabilitiesKind is the code action kind is support with the following value set.
type CodeActionClientCapabilitiesKind struct {
	// ValueSet is the code action kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	ValueSet []CodeActionKind `json:"valueSet"`
}

// PublishDiagnosticsClientCapabilitiesTagSupport is the client capacity of TagSupport.
//
// @since 3.15.0.
type PublishDiagnosticsClientCapabilitiesTagSupport struct {
	// ValueSet is the tags supported by the client.
	ValueSet []DiagnosticTag `json:"valueSet"`
}

// SemanticTokensWorkspaceClientCapabilitiesRequests capabilities specific to the "textDocument/semanticTokens/xxx" request.
//
// @since 3.16.0.
type SemanticTokensWorkspaceClientCapabilitiesRequests struct {
	// Range is the client will send the "textDocument/semanticTokens/range" request
	// if the server provides a corresponding handler.
	Range bool `json:"range,omitempty"`

	// Full is the client will send the "textDocument/semanticTokens/full" request
	// if the server provides a corresponding handler.
	Full bool `json:"full,omitempty"` // delta?: boolean
}

// ShowMessageRequestClientCapabilitiesMessageActionItem capabilities specific to the "MessageActionItem" type.
//
// @since 3.16.0.
type ShowMessageRequestClientCapabilitiesMessageActionItem struct {
	// AdditionalPropertiesSupport whether the client supports additional attributes which
	// are preserved and sent back to the server in the
	// request's response.
	AdditionalPropertiesSupport bool `json:"additionalPropertiesSupport,omitempty"`
}

// InitializeResult result of ClientCapabilities.
type InitializeResult struct {
	// Capabilities is the capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`

	// ServerInfo Information about the server.
	//
	// @since 3.15.0.
	ServerInfo *ServerInfo `json:"serverInfo,omitempty"`
}

// ServerCapabilitiesWorkspace specific server capabilities.
type ServerCapabilitiesWorkspace struct {
	// WorkspaceFolders is the server supports workspace folder.
	//
	// @since 3.6.0.
	WorkspaceFolders *ServerCapabilitiesWorkspaceFolders `json:"workspaceFolders,omitempty"`

	// FileOperations is the server is interested in file notifications/requests.
	//
	// @since 3.16.0.
	FileOperations *ServerCapabilitiesWorkspaceFileOperations `json:"fileOperations,omitempty"`
}

// FileOperationRegistrationOptions is the options to register for file operations.
//
// @since 3.16.0.
type FileOperationRegistrationOptions struct {
	// filters is the actual filters.
	Filters []FileOperationFilter `json:"filters"`
}

// ServerCapabilitiesWorkspaceFileOperations is the server is interested in file notifications/requests.
//
// @since 3.16.0.
type ServerCapabilitiesWorkspaceFileOperations struct {
	// DidCreate is the server is interested in receiving didCreateFiles
	// notifications.
	DidCreate *FileOperationRegistrationOptions `json:"didCreate,omitempty"`

	// WillCreate is the server is interested in receiving willCreateFiles requests.
	WillCreate *FileOperationRegistrationOptions `json:"willCreate,omitempty"`

	// DidRename is the server is interested in receiving didRenameFiles
	// notifications.
	DidRename *FileOperationRegistrationOptions `json:"didRename,omitempty"`

	// WillRename is the server is interested in receiving willRenameFiles requests.
	WillRename *FileOperationRegistrationOptions `json:"willRename,omitempty"`

	// DidDelete is the server is interested in receiving didDeleteFiles file
	// notifications.
	DidDelete *FileOperationRegistrationOptions `json:"didDelete,omitempty"`

	// WillDelete is the server is interested in receiving willDeleteFiles file
	// requests.
	WillDelete *FileOperationRegistrationOptions `json:"willDelete,omitempty"`
}

// ServerCapabilities efines the capabilities provided by a language server.
type ServerCapabilities struct {
	// TextDocumentSync defines how text documents are synced. Is either a detailed structure defining each notification
	// or for backwards compatibility the TextDocumentSyncKind number.
	//
	// If omitted it defaults to TextDocumentSyncKind.None`
	TextDocumentSync interface{} `json:"textDocumentSync,omitempty"` // *TextDocumentSyncOptions | TextDocumentSyncKind

	// CompletionProvider is The server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	// HoverProvider is the server provides hover support.
	HoverProvider interface{} `json:"hoverProvider,omitempty"` // TODO(zchee): bool | *HoverOptions

	// SignatureHelpProvider is the server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	// DeclarationProvider is the server provides Goto Declaration support.
	//
	// @since 3.14.0.
	DeclarationProvider interface{} `json:"declarationProvider,omitempty"` // TODO(zchee): bool | *DeclarationOptions | *DeclarationRegistrationOptions

	// DefinitionProvider is the server provides Goto definition support.
	DefinitionProvider interface{} `json:"definitionProvider,omitempty"` // TODO(zchee): bool | *DefinitionOptions

	// TypeDefinitionProvider is the provides Goto Type Definition support.
	//
	// @since 3.6.0.
	TypeDefinitionProvider interface{} `json:"typeDefinitionProvider,omitempty"` // TODO(zchee): bool | *TypeDefinitionOptions | *TypeDefinitionRegistrationOptions

	// ImplementationProvider is the provides Goto Implementation support.
	//
	// @since 3.6.0.
	ImplementationProvider interface{} `json:"implementationProvider,omitempty"` // TODO(zchee): bool | *ImplementationOptions | *ImplementationRegistrationOptions

	// ReferencesProvider is the server provides find references support.
	ReferencesProvider interface{} `json:"referencesProvider,omitempty"` // TODO(zchee): bool | *ReferenceOptions

	// DocumentHighlightProvider is the server provides document highlight support.
	DocumentHighlightProvider interface{} `json:"documentHighlightProvider,omitempty"` // TODO(zchee): bool | *DocumentHighlightOptions

	// DocumentSymbolProvider is the server provides document symbol support.
	DocumentSymbolProvider interface{} `json:"documentSymbolProvider,omitempty"` // TODO(zchee): bool | *DocumentSymbolOptions

	// CodeActionProvider is the server provides code actions.
	//
	// CodeActionOptions may only be specified if the client states that it supports CodeActionLiteralSupport in its
	// initial Initialize request.
	CodeActionProvider interface{} `json:"codeActionProvider,omitempty"` // TODO(zchee): bool | *CodeActionOptions

	// CodeLensProvider is the server provides code lens.
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider,omitempty"`

	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider,omitempty"`

	// ColorProvider is the server provides color provider support.
	//
	// @since 3.6.0.
	ColorProvider interface{} `json:"colorProvider,omitempty"` // TODO(zchee): bool | *DocumentColorOptions | *DocumentColorRegistrationOptions

	// WorkspaceSymbolProvider is the server provides workspace symbol support.
	WorkspaceSymbolProvider interface{} `json:"workspaceSymbolProvider,omitempty"` // TODO(zchee): bool | *WorkspaceSymbolOptions

	// DocumentFormattingProvider is the server provides document formatting.
	DocumentFormattingProvider interface{} `json:"documentFormattingProvider,omitempty"` // TODO(zchee): bool | *DocumentFormattingOptions

	// DocumentRangeFormattingProvider is the server provides document range formatting.
	DocumentRangeFormattingProvider interface{} `json:"documentRangeFormattingProvider,omitempty"` // TODO(zchee): bool | *DocumentRangeFormattingOptions

	// DocumentOnTypeFormattingProvider is the server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`

	// RenameProvider is the server provides rename support.
	//
	// RenameOptions may only be specified if the client states that it supports PrepareSupport in its
	// initial Initialize request.
	RenameProvider interface{} `json:"renameProvider,omitempty"` // TODO(zchee): bool | *RenameOptions

	// FoldingRangeProvider is the server provides folding provider support.
	//
	// @since 3.10.0.
	FoldingRangeProvider interface{} `json:"foldingRangeProvider,omitempty"` // TODO(zchee): bool | *FoldingRangeOptions | *FoldingRangeRegistrationOptions

	// SelectionRangeProvider is the server provides selection range support.
	//
	// @since 3.15.0.
	SelectionRangeProvider interface{} `json:"selectionRangeProvider,omitempty"` // TODO(zchee): bool | *SelectionRangeOptions | *SelectionRangeRegistrationOptions

	// ExecuteCommandProvider is the server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider,omitempty"`

	// CallHierarchyProvider is the server provides call hierarchy support.
	//
	// @since 3.16.0.
	CallHierarchyProvider interface{} `json:"callHierarchyProvider,omitempty"` // TODO(zchee): bool | *CallHierarchyOptions | *CallHierarchyRegistrationOptions

	// LinkedEditingRangeProvider is the server provides linked editing range support.
	//
	// @since 3.16.0.
	LinkedEditingRangeProvider interface{} `json:"linkedEditingRangeProvider,omitempty"` // TODO(zchee): bool | *LinkedEditingRangeOptions | *LinkedEditingRangeRegistrationOptions

	// SemanticTokensProvider is the server provides semantic tokens support.
	//
	// @since 3.16.0.
	SemanticTokensProvider interface{} `json:"semanticTokensProvider,omitempty"` // TODO(zchee): *SemanticTokensOptions | *SemanticTokensRegistrationOptions

	// Workspace is the window specific server capabilities.
	Workspace *ServerCapabilitiesWorkspace `json:"workspace,omitempty"`

	// MonikerProvider is the server provides moniker support.
	//
	// @since 3.16.0.
	MonikerProvider interface{} `json:"monikerProvider,omitempty"` // TODO(zchee): bool | *MonikerOptions | *MonikerRegistrationOptions

	// Experimental server capabilities.
	Experimental interface{} `json:"experimental,omitempty"`
}
