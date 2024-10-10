// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

// TextDocumentSyncKind defines how the host (editor) should sync document changes to the language server.
type TextDocumentSyncKind uint32

const (
	// NoneTextDocumentSyncKind documents should not be synced at all.
	NoneTextDocumentSyncKind TextDocumentSyncKind = 0

	// FullTextDocumentSyncKind documents are synced by always sending the full content of the document.
	FullTextDocumentSyncKind TextDocumentSyncKind = 1

	// IncrementalTextDocumentSyncKind documents are synced by sending the full content on open. After that only incremental updates to the
	// document are send.
	IncrementalTextDocumentSyncKind TextDocumentSyncKind = 2
)

// TextDocumentRegistrationOptions general text document registration options.
type TextDocumentRegistrationOptions struct {
	// DocumentSelector a document selector to identify the scope of the registration. If set to null the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector,omitempty"`
}

// Registration general parameters to register for a notification or to register a provider.
type Registration struct {
	// ID the id used to register the request. The id can be used to deregister the request again.
	ID string `json:"id"`

	// Method the method / capability to register for.
	Method string `json:"method"`

	// RegisterOptions options necessary for the registration.
	RegisterOptions any `json:"registerOptions,omitempty"`
}

type RegistrationParams struct {
	Registrations []Registration `json:"registrations"`
}

// Unregistration general parameters to unregister a request or notification.
type Unregistration struct {
	// ID the id used to unregister the request or notification. Usually an id provided during the register request.
	ID string `json:"id"`

	// Method the method to unregister for.
	Method string `json:"method"`
}

type UnregistrationParams struct {
	Unregisterations []Unregistration `json:"unregisterations"`
}

// ClientInfo information about the client  3.15.0  3.18.0 ClientInfo type name added.
//
// @since 3.18.0 ClientInfo type name added.
type ClientInfo struct {
	// Name the name of the client as defined by the client.
	//
	// @since 3.18.0 ClientInfo type name added.
	Name string `json:"name"`

	// Version the client's version as defined by the client.
	//
	// @since 3.18.0 ClientInfo type name added.
	Version string `json:"version,omitempty"`
}

// SemanticTokensWorkspaceClientCapabilities.
//
// @since 3.16.0
type SemanticTokensWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all semantic tokens currently shown. It should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation.
	//
	// @since 3.16.0
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// CodeLensWorkspaceClientCapabilities.
//
// @since 3.16.0
type CodeLensWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all code lenses currently shown. It
	// should be used with absolute care and is useful for situation where a server for example detect a project wide change that requires such a calculation.
	//
	// @since 3.16.0
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// FileOperationClientCapabilities capabilities relating to events from file operations by the user in the client. These events do not come from the file system, they come from user operations like renaming a file in the UI.
//
// @since 3.16.0
type FileOperationClientCapabilities struct {
	// DynamicRegistration whether the client supports dynamic registration for file requests/notifications.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// DidCreate the client has support for sending didCreateFiles notifications.
	//
	// @since 3.16.0
	DidCreate bool `json:"didCreate,omitempty"`

	// WillCreate the client has support for sending willCreateFiles requests.
	//
	// @since 3.16.0
	WillCreate bool `json:"willCreate,omitempty"`

	// DidRename the client has support for sending didRenameFiles notifications.
	//
	// @since 3.16.0
	DidRename bool `json:"didRename,omitempty"`

	// WillRename the client has support for sending willRenameFiles requests.
	//
	// @since 3.16.0
	WillRename bool `json:"willRename,omitempty"`

	// DidDelete the client has support for sending didDeleteFiles notifications.
	//
	// @since 3.16.0
	DidDelete bool `json:"didDelete,omitempty"`

	// WillDelete the client has support for sending willDeleteFiles requests.
	//
	// @since 3.16.0
	WillDelete bool `json:"willDelete,omitempty"`
}

// InlineValueWorkspaceClientCapabilities client workspace capabilities specific to inline values.
//
// @since 3.17.0
type InlineValueWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all inline values currently shown. It should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation.
	//
	// @since 3.17.0
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// InlayHintWorkspaceClientCapabilities client workspace capabilities specific to inlay hints.
//
// @since 3.17.0
type InlayHintWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all inlay hints currently shown. It
	// should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation.
	//
	// @since 3.17.0
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// DiagnosticWorkspaceClientCapabilities workspace client capabilities specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all pulled diagnostics currently shown. It should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation.
	//
	// @since 3.17.0
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// FoldingRangeWorkspaceClientCapabilities client workspace capabilities specific to folding ranges  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type FoldingRangeWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all folding ranges currently shown.
	// It should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation. 3.18.0 @proposed.
	// @since 3.18.0 proposed
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// TextDocumentContentClientCapabilities client capabilities for a text document content provider.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type TextDocumentContentClientCapabilities struct {
	// DynamicRegistration text document content provider supports dynamic registration.
	//
	// @since 3.18.0 proposed
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// WorkspaceClientCapabilities workspace specific client capabilities.
type WorkspaceClientCapabilities struct {
	// ApplyEdit the client supports applying batch edits to the workspace by supporting the request 'workspace/applyEdit'.
	ApplyEdit bool `json:"applyEdit,omitempty"`

	// WorkspaceEdit capabilities specific to `WorkspaceEdit`s.
	WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspaceEdit,omitempty"`

	// DidChangeConfiguration capabilities specific to the `workspace/didChangeConfiguration` notification.
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitempty"`

	// DidChangeWatchedFiles capabilities specific to the `workspace/didChangeWatchedFiles` notification.
	DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"didChangeWatchedFiles,omitempty"`

	// Symbol capabilities specific to the `workspace/symbol` request.
	Symbol *WorkspaceSymbolClientCapabilities `json:"symbol,omitempty"`

	// ExecuteCommand capabilities specific to the `workspace/executeCommand` request.
	ExecuteCommand *ExecuteCommandClientCapabilities `json:"executeCommand,omitempty"`

	// WorkspaceFolders the client has support for workspace folders.
	WorkspaceFolders bool `json:"workspaceFolders,omitempty"`

	// Configuration the client supports `workspace/configuration` requests.
	Configuration bool `json:"configuration,omitempty"`

	// SemanticTokens capabilities specific to the semantic token requests scoped to the workspace.
	SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitempty"`

	// CodeLens capabilities specific to the code lens requests scoped to the workspace.
	CodeLens *CodeLensWorkspaceClientCapabilities `json:"codeLens,omitempty"`

	// FileOperations the client has support for file notifications/requests for user operations on files. Since .
	FileOperations *FileOperationClientCapabilities `json:"fileOperations,omitempty"`

	// InlineValue capabilities specific to the inline values requests scoped to the workspace.
	InlineValue *InlineValueWorkspaceClientCapabilities `json:"inlineValue,omitempty"`

	// InlayHint capabilities specific to the inlay hint requests scoped to the workspace.
	InlayHint *InlayHintWorkspaceClientCapabilities `json:"inlayHint,omitempty"`

	// Diagnostics capabilities specific to the diagnostic requests scoped to the workspace.
	Diagnostics *DiagnosticWorkspaceClientCapabilities `json:"diagnostics,omitempty"`

	// FoldingRange capabilities specific to the folding range requests scoped to the workspace.  3.18.0 @proposed.
	FoldingRange *FoldingRangeWorkspaceClientCapabilities `json:"foldingRange,omitempty"`

	// TextDocumentContent capabilities specific to the `workspace/textDocumentContent` request.  3.18.0 @proposed.
	TextDocumentContent *TextDocumentContentClientCapabilities `json:"textDocumentContent,omitempty"`
}

type TextDocumentSyncClientCapabilities struct {
	// DynamicRegistration whether text document synchronization supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// WillSave the client supports sending will save notifications.
	WillSave bool `json:"willSave,omitempty"`

	// WillSaveWaitUntil the client supports sending a will save request and waits for a response providing text edits which will be applied to the document before it is saved.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

	// DidSave the client supports did save notifications.
	DidSave bool `json:"didSave,omitempty"`
}

// CompletionItemTagOptions.
//
// @since 3.18.0
type CompletionItemTagOptions struct {
	// ValueSet the tags supported by the client.
	//
	// @since 3.18.0
	ValueSet []CompletionItemTag `json:"valueSet"`
}

// ClientCompletionItemResolveOptions.
//
// @since 3.18.0
type ClientCompletionItemResolveOptions struct {
	// Properties the properties that a client can resolve lazily.
	//
	// @since 3.18.0
	Properties []string `json:"properties"`
}

// ClientCompletionItemInsertTextModeOptions.
//
// @since 3.18.0
type ClientCompletionItemInsertTextModeOptions struct {
	// @since 3.18.0
	ValueSet []InsertTextMode `json:"valueSet"`
}

// ClientCompletionItemOptions.
//
// @since 3.18.0
type ClientCompletionItemOptions struct {
	// SnippetSupport client supports snippets as insert text. A snippet can define tab stops and placeholders with `$1`, `$2` and `${3:foo}`. `$0` defines the final tab stop, it defaults to the end of the snippet. Placeholders with equal identifiers are linked, that is typing in one will update others too.
	//
	// @since 3.18.0
	SnippetSupport bool `json:"snippetSupport,omitempty"`

	// CommitCharactersSupport client supports commit characters on a completion item.
	//
	// @since 3.18.0
	CommitCharactersSupport bool `json:"commitCharactersSupport,omitempty"`

	// DocumentationFormat client supports the following content formats for the documentation property. The order describes the preferred format of the client.
	//
	// @since 3.18.0
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

	// DeprecatedSupport client supports the deprecated property on a completion item.
	//
	// @since 3.18.0
	DeprecatedSupport bool `json:"deprecatedSupport,omitempty"`

	// PreselectSupport client supports the preselect property on a completion item.
	//
	// @since 3.18.0
	PreselectSupport bool `json:"preselectSupport,omitempty"`

	// TagSupport client supports the tag property on a completion item. Clients supporting tags have to handle unknown tags gracefully. Clients especially need to preserve unknown tags when sending a completion item back to the server in a resolve call.
	// @since 3.18.0
	TagSupport *CompletionItemTagOptions `json:"tagSupport,omitempty"`

	// InsertReplaceSupport client support insert replace edit to control different behavior if a completion item is inserted in
	// the text or should replace text.
	// @since 3.18.0
	InsertReplaceSupport bool `json:"insertReplaceSupport,omitempty"`

	// ResolveSupport indicates which properties a client can resolve lazily on a completion item. Before version 3.16.0 only the predefined properties `documentation` and `details` could be resolved lazily.
	// @since 3.18.0
	ResolveSupport *ClientCompletionItemResolveOptions `json:"resolveSupport,omitempty"`

	// InsertTextModeSupport the client supports the `insertTextMode` property on a completion item to override the whitespace handling mode as defined by the client (see `insertTextMode`).
	// @since 3.18.0
	InsertTextModeSupport *ClientCompletionItemInsertTextModeOptions `json:"insertTextModeSupport,omitempty"`

	// LabelDetailsSupport the client has support for completion item label details (see also `CompletionItemLabelDetails`).
	// @since 3.18.0
	LabelDetailsSupport bool `json:"labelDetailsSupport,omitempty"`
}

// ClientCompletionItemOptionsKind.
//
// @since 3.18.0
type ClientCompletionItemOptionsKind struct {
	// ValueSet the completion item kind values the client supports. When this property exists the client also guarantees that it will handle values outside its set gracefully and falls back to a default value when unknown. If this property is not present the client only supports the completion items kinds from `Text` to `Reference` as defined in the initial version of the protocol.
	//
	// @since 3.18.0
	ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
}

// CompletionListCapabilities the client supports the following `CompletionList` specific capabilities.
//
// @since 3.17.0
type CompletionListCapabilities struct {
	// ItemDefaults the client supports the following itemDefaults on a completion list. The value lists the supported property names of the `CompletionList.itemDefaults` object. If omitted no properties are supported.
	// @since 3.17.0
	ItemDefaults []string `json:"itemDefaults,omitempty"`

	// ApplyKindSupport specifies whether the client supports `CompletionList.applyKind` to indicate how supported values from `completionList.itemDefaults` and `completion` will be combined. If a client supports `applyKind`
	// it must support it for all fields that it supports that are listed in `CompletionList.applyKind`. This means when clients add support for new/future fields in completion items the MUST also support merge for them if those fields are defined in `CompletionList.applyKind`.
	// @since 3.17.0
	ApplyKindSupport bool `json:"applyKindSupport,omitempty"`
}

// CompletionClientCapabilities completion client capabilities.
type CompletionClientCapabilities struct {
	// DynamicRegistration whether completion supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// CompletionItem the client supports the following `CompletionItem` specific capabilities.
	CompletionItem *ClientCompletionItemOptions `json:"completionItem,omitempty"`

	CompletionItemKind *ClientCompletionItemOptionsKind `json:"completionItemKind,omitempty"`

	// InsertTextMode defines how the client handles whitespace and indentation when accepting a completion item that uses
	// multi line text in either `insertText` or `textEdit`.
	InsertTextMode InsertTextMode `json:"insertTextMode,omitempty"`

	// ContextSupport the client supports to send additional context information for a `textDocument/completion` request.
	ContextSupport bool `json:"contextSupport,omitempty"`

	// CompletionList the client supports the following `CompletionList` specific capabilities.
	CompletionList *CompletionListCapabilities `json:"completionList,omitempty"`
}

type HoverClientCapabilities struct {
	// DynamicRegistration whether hover supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// ContentFormat client supports the following content formats for the content property. The order describes the preferred format of the client.
	ContentFormat []MarkupKind `json:"contentFormat,omitempty"`
}

// ClientSignatureParameterInformationOptions.
//
// @since 3.18.0
type ClientSignatureParameterInformationOptions struct {
	// LabelOffsetSupport the client supports processing label offsets instead of a simple label string.
	// @since 3.18.0
	LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
}

// ClientSignatureInformationOptions.
//
// @since 3.18.0
type ClientSignatureInformationOptions struct {
	// DocumentationFormat client supports the following content formats for the documentation property. The order describes the preferred format of the client.
	//
	// @since 3.18.0
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

	// ParameterInformation client capabilities specific to parameter information.
	//
	// @since 3.18.0
	ParameterInformation *ClientSignatureParameterInformationOptions `json:"parameterInformation,omitempty"`

	// ActiveParameterSupport the client supports the `activeParameter` property on `SignatureInformation` literal.
	// @since 3.18.0
	ActiveParameterSupport bool `json:"activeParameterSupport,omitempty"`

	// NoActiveParameterSupport the client supports the `activeParameter` property on `SignatureHelp`/`SignatureInformation` being set to `null` to indicate that no parameter should be active. 3.18.0 @proposed.
	// @since 3.18.0
	NoActiveParameterSupport bool `json:"noActiveParameterSupport,omitempty"`
}

// SignatureHelpClientCapabilities client Capabilities for a SignatureHelpRequest.
type SignatureHelpClientCapabilities struct {
	// DynamicRegistration whether signature help supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// SignatureInformation the client supports the following `SignatureInformation` specific properties.
	SignatureInformation *ClientSignatureInformationOptions `json:"signatureInformation,omitempty"`

	// ContextSupport the client supports to send additional context information for a `textDocument/signatureHelp` request. A client that opts into contextSupport will also support the `retriggerCharacters` on `SignatureHelpOptions`.
	ContextSupport bool `json:"contextSupport,omitempty"`
}

// DeclarationClientCapabilities.
//
// @since 3.14.0
type DeclarationClientCapabilities struct {
	// DynamicRegistration whether declaration supports dynamic registration. If this is set to `true` the client supports the new `DeclarationRegistrationOptions` return value for the corresponding server capability as well.
	//
	// @since 3.14.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// LinkSupport the client supports additional metadata in the form of declaration links.
	//
	// @since 3.14.0
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// DefinitionClientCapabilities client Capabilities for a DefinitionRequest.
type DefinitionClientCapabilities struct {
	// DynamicRegistration whether definition supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// LinkSupport the client supports additional metadata in the form of definition links.
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// TypeDefinitionClientCapabilities since .
type TypeDefinitionClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `TypeDefinitionRegistrationOptions` return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// LinkSupport the client supports additional metadata in the form of definition links. Since .
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// ImplementationClientCapabilities.
//
// @since 3.6.0
type ImplementationClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `ImplementationRegistrationOptions` return value for the corresponding server capability as well.
	//
	// @since 3.6.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// LinkSupport the client supports additional metadata in the form of definition links.
	// @since 3.6.0
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// ReferenceClientCapabilities client Capabilities for a ReferencesRequest.
type ReferenceClientCapabilities struct {
	// DynamicRegistration whether references supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentHighlightClientCapabilities client Capabilities for a DocumentHighlightRequest.
type DocumentHighlightClientCapabilities struct {
	// DynamicRegistration whether document highlight supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentSymbolClientCapabilities client Capabilities for a DocumentSymbolRequest.
type DocumentSymbolClientCapabilities struct {
	// DynamicRegistration whether document symbol supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// SymbolKind specific capabilities for the `SymbolKind` in the `textDocument/documentSymbol` request.
	SymbolKind *ClientSymbolKindOptions `json:"symbolKind,omitempty"`

	// HierarchicalDocumentSymbolSupport the client supports hierarchical document symbols.
	HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`

	// TagSupport the client supports tags on `SymbolInformation`. Tags are supported on `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true. Clients supporting tags have to handle unknown tags gracefully.
	TagSupport *ClientSymbolTagOptions `json:"tagSupport,omitempty"`

	// LabelSupport the client supports an additional label presented in the UI when registering a document symbol provider.
	LabelSupport bool `json:"labelSupport,omitempty"`
}

// ClientCodeActionKindOptions.
//
// @since 3.18.0
type ClientCodeActionKindOptions struct {
	// ValueSet the code action kind values the client supports. When this property exists the client also guarantees that it will handle values outside its set gracefully and falls back to a default value when unknown.
	//
	// @since 3.18.0
	ValueSet []CodeActionKind `json:"valueSet"`
}

// ClientCodeActionLiteralOptions.
//
// @since 3.18.0
type ClientCodeActionLiteralOptions struct {
	// CodeActionKind the code action kind is support with the following value set.
	//
	// @since 3.18.0
	CodeActionKind ClientCodeActionKindOptions `json:"codeActionKind"`
}

// ClientCodeActionResolveOptions.
//
// @since 3.18.0
type ClientCodeActionResolveOptions struct {
	// Properties the properties that a client can resolve lazily.
	//
	// @since 3.18.0
	Properties []string `json:"properties"`
}

// CodeActionClientCapabilities the Client Capabilities of a CodeActionRequest.
type CodeActionClientCapabilities struct {
	// DynamicRegistration whether code action supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// CodeActionLiteralSupport the client support code action literals of type `CodeAction` as a valid response of the `textDocument/codeAction` request. If the property is not set the request can only return `Command` literals.
	CodeActionLiteralSupport *ClientCodeActionLiteralOptions `json:"codeActionLiteralSupport,omitempty"`

	// IsPreferredSupport whether code action supports the `isPreferred` property.
	IsPreferredSupport bool `json:"isPreferredSupport,omitempty"`

	// DisabledSupport whether code action supports the `disabled` property.
	DisabledSupport bool `json:"disabledSupport,omitempty"`

	// DataSupport whether code action supports the `data` property which is preserved between a `textDocument/codeAction` and a `codeAction/resolve` request.
	DataSupport bool `json:"dataSupport,omitempty"`

	// ResolveSupport whether the client supports resolving additional code action properties via a separate `codeAction/resolve` request.
	ResolveSupport *ClientCodeActionResolveOptions `json:"resolveSupport,omitempty"`

	// HonorsChangeAnnotations whether the client honors the change annotations in text edits and resource operations returned via the `CodeAction#edit` property by for example presenting the workspace edit in the user interface and asking for confirmation.
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`

	// DocumentationSupport whether the client supports documentation for a class of code actions.  3.18.0 @proposed.
	DocumentationSupport bool `json:"documentationSupport,omitempty"`

	// TagSupport client supports the tag property on a code action. Clients supporting tags have to handle unknown tags gracefully. 3.18.0 - proposed.
	TagSupport *CodeActionTagOptions `json:"tagSupport,omitempty"`
}

// ClientCodeLensResolveOptions.
//
// @since 3.18.0
type ClientCodeLensResolveOptions struct {
	// Properties the properties that a client can resolve lazily.
	//
	// @since 3.18.0
	Properties []string `json:"properties"`
}

// CodeLensClientCapabilities the client capabilities of a CodeLensRequest.
type CodeLensClientCapabilities struct {
	// DynamicRegistration whether code lens supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// ResolveSupport whether the client supports resolving additional code lens properties via a separate `codeLens/resolve` request.
	ResolveSupport *ClientCodeLensResolveOptions `json:"resolveSupport,omitempty"`
}

// DocumentLinkClientCapabilities the client capabilities of a DocumentLinkRequest.
type DocumentLinkClientCapabilities struct {
	// DynamicRegistration whether document link supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// TooltipSupport whether the client supports the `tooltip` property on `DocumentLink`.
	TooltipSupport bool `json:"tooltipSupport,omitempty"`
}

type DocumentColorClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `DocumentColorRegistrationOptions` return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentFormattingClientCapabilities client capabilities of a DocumentFormattingRequest.
type DocumentFormattingClientCapabilities struct {
	// DynamicRegistration whether formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentRangeFormattingClientCapabilities client capabilities of a DocumentRangeFormattingRequest.
type DocumentRangeFormattingClientCapabilities struct {
	// DynamicRegistration whether range formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// RangesSupport whether the client supports formatting multiple ranges at once.  3.18.0 @proposed.
	RangesSupport bool `json:"rangesSupport,omitempty"`
}

// DocumentOnTypeFormattingClientCapabilities client capabilities of a DocumentOnTypeFormattingRequest.
type DocumentOnTypeFormattingClientCapabilities struct {
	// DynamicRegistration whether on type formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type RenameClientCapabilities struct {
	// DynamicRegistration whether rename supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// PrepareSupport client supports testing for validity of rename operations before execution.
	PrepareSupport bool `json:"prepareSupport,omitempty"`

	// PrepareSupportDefaultBehavior client supports the default behavior result. The value indicates the default behavior used by the client.
	PrepareSupportDefaultBehavior PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior,omitempty"`

	// HonorsChangeAnnotations whether the client honors the change annotations in text edits and resource operations returned via the rename request's workspace edit by for example presenting the workspace edit in the user interface and asking for confirmation.
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`
}

// ClientFoldingRangeKindOptions.
//
// @since 3.18.0
type ClientFoldingRangeKindOptions struct {
	// ValueSet the folding range kind values the client supports. When this property exists the client also guarantees that it will handle values outside its set gracefully and falls back to a default value when unknown.
	//
	// @since 3.18.0
	ValueSet []FoldingRangeKind `json:"valueSet,omitempty"`
}

// ClientFoldingRangeOptions.
//
// @since 3.18.0
type ClientFoldingRangeOptions struct {
	// CollapsedText if set, the client signals that it supports setting collapsedText on folding ranges to display custom labels instead of the default text.
	// @since 3.18.0
	CollapsedText bool `json:"collapsedText,omitempty"`
}

type FoldingRangeClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration for folding range providers. If this is set to `true` the client supports the new `FoldingRangeRegistrationOptions` return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// RangeLimit the maximum number of folding ranges that the client prefers to receive per document. The value serves as a hint, servers are free to follow the limit.
	RangeLimit uint32 `json:"rangeLimit,omitempty"`

	// LineFoldingOnly if set, the client signals that it only supports folding complete lines. If set, client will ignore specified `startCharacter` and `endCharacter` properties in a FoldingRange.
	LineFoldingOnly bool `json:"lineFoldingOnly,omitempty"`

	// FoldingRangeKind specific options for the folding range kind.
	FoldingRangeKind *ClientFoldingRangeKindOptions `json:"foldingRangeKind,omitempty"`

	// FoldingRange specific options for the folding range.
	FoldingRange *ClientFoldingRangeOptions `json:"foldingRange,omitempty"`
}

type SelectionRangeClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration for selection range providers. If this is set to `true` the client supports the new `SelectionRangeRegistrationOptions` return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// ClientDiagnosticsTagOptions.
//
// @since 3.18.0
type ClientDiagnosticsTagOptions struct {
	// ValueSet the tags supported by the client.
	//
	// @since 3.18.0
	ValueSet []DiagnosticTag `json:"valueSet"`
}

// DiagnosticsCapabilities general diagnostics capabilities for pull and push model.
type DiagnosticsCapabilities struct {
	// RelatedInformation whether the clients accepts diagnostics with related information.
	RelatedInformation bool `json:"relatedInformation,omitempty"`

	// TagSupport client supports the tag property to provide meta data about a diagnostic. Clients supporting tags have to handle unknown tags gracefully.
	TagSupport *ClientDiagnosticsTagOptions `json:"tagSupport,omitempty"`

	// CodeDescriptionSupport client supports a codeDescription property
	CodeDescriptionSupport bool `json:"codeDescriptionSupport,omitempty"`

	// DataSupport whether code action supports the `data` property which is preserved between a `textDocument/publishDiagnostics` and `textDocument/codeAction` request.
	DataSupport bool `json:"dataSupport,omitempty"`
}

// PublishDiagnosticsClientCapabilities the publish diagnostic client capabilities.
type PublishDiagnosticsClientCapabilities struct {
	// extends
	DiagnosticsCapabilities

	// VersionSupport whether the client interprets the version property of the `textDocument/publishDiagnostics` notification's parameter.
	VersionSupport bool `json:"versionSupport,omitempty"`
}

// CallHierarchyClientCapabilities.
//
// @since 3.16.0
type CallHierarchyClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// ClientSemanticTokensRequestFullDelta.
//
// @since 3.18.0
type ClientSemanticTokensRequestFullDelta struct {
	// Delta the client will send the `textDocument/semanticTokens/full/delta` request if the server provides a corresponding handler.
	//
	// @since 3.18.0
	Delta bool `json:"delta,omitempty"`
}

// ClientSemanticTokensRequestOptions.
//
// @since 3.18.0
type ClientSemanticTokensRequestOptions struct {
	// Range the client will send the `textDocument/semanticTokens/range` request if the server provides a corresponding handler.
	//
	// @since 3.18.0
	Range *ClientSemanticTokensRequestOptionsRange `json:"range,omitempty"`

	// Full the client will send the `textDocument/semanticTokens/full` request if the server provides a corresponding handler.
	//
	// @since 3.18.0
	Full *ClientSemanticTokensRequestOptionsFull `json:"full,omitempty"`
}

// SemanticTokensClientCapabilities.
//
// @since 3.16.0
type SemanticTokensClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Requests which requests the client supports and might send to the server depending on the server's capability. Please note that clients might not show semantic tokens or degrade some of the user experience if a range or full request is advertised by the client but not provided by the server. If for example the client capability `requests.full` and `request.range` are both set to true but the server only provides a range provider the client might not render a minimap correctly or might even decide to not show any semantic tokens at all.
	//
	// @since 3.16.0
	Requests ClientSemanticTokensRequestOptions `json:"requests"`

	// TokenTypes the token types that the client supports.
	//
	// @since 3.16.0
	TokenTypes []string `json:"tokenTypes"`

	// TokenModifiers the token modifiers that the client supports.
	//
	// @since 3.16.0
	TokenModifiers []string `json:"tokenModifiers"`

	// Formats the token formats the clients supports.
	//
	// @since 3.16.0
	Formats []TokenFormat `json:"formats"`

	// OverlappingTokenSupport whether the client supports tokens that can overlap each other.
	//
	// @since 3.16.0
	OverlappingTokenSupport bool `json:"overlappingTokenSupport,omitempty"`

	// MultilineTokenSupport whether the client supports tokens that can span multiple lines.
	//
	// @since 3.16.0
	MultilineTokenSupport bool `json:"multilineTokenSupport,omitempty"`

	// ServerCancelSupport whether the client allows the server to actively cancel a semantic token request, e.g. supports returning LSPErrorCodes.ServerCancelled. If a server does the client needs to retrigger the request.
	// @since 3.16.0
	ServerCancelSupport bool `json:"serverCancelSupport,omitempty"`

	// AugmentsSyntaxTokens whether the client uses semantic tokens to augment existing syntax tokens. If set to `true` client side created syntax tokens and semantic tokens are both used for colorization. If set to `false` the client only uses the returned semantic tokens for colorization. If the value is `undefined` then the
	// client behavior is not specified.
	// @since 3.16.0
	AugmentsSyntaxTokens bool `json:"augmentsSyntaxTokens,omitempty"`
}

// LinkedEditingRangeClientCapabilities client capabilities for the linked editing range request.
//
// @since 3.16.0
type LinkedEditingRangeClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// MonikerClientCapabilities client capabilities specific to the moniker request.
//
// @since 3.16.0
type MonikerClientCapabilities struct {
	// DynamicRegistration whether moniker supports dynamic registration. If this is set to `true` the client supports the new `MonikerRegistrationOptions` return value for the corresponding server capability as well.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TypeHierarchyClientCapabilities.
//
// @since 3.17.0
type TypeHierarchyClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// InlineValueClientCapabilities client capabilities specific to inline values.
//
// @since 3.17.0
type InlineValueClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration for inline value providers.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// ClientInlayHintResolveOptions.
//
// @since 3.18.0
type ClientInlayHintResolveOptions struct {
	// Properties the properties that a client can resolve lazily.
	//
	// @since 3.18.0
	Properties []string `json:"properties"`
}

// InlayHintClientCapabilities inlay hint client capabilities.
//
// @since 3.17.0
type InlayHintClientCapabilities struct {
	// DynamicRegistration whether inlay hints support dynamic registration.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// ResolveSupport indicates which properties a client can resolve lazily on an inlay hint.
	//
	// @since 3.17.0
	ResolveSupport *ClientInlayHintResolveOptions `json:"resolveSupport,omitempty"`
}

// DiagnosticClientCapabilities client capabilities specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticClientCapabilities struct {
	// extends
	DiagnosticsCapabilities

	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// RelatedDocumentSupport whether the clients supports related documents for document diagnostic pulls.
	//
	// @since 3.17.0
	RelatedDocumentSupport bool `json:"relatedDocumentSupport,omitempty"`
}

// InlineCompletionClientCapabilities client capabilities specific to inline completions.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration for inline completion providers.
	//
	// @since 3.18.0 proposed
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilities text document specific client capabilities.
type TextDocumentClientCapabilities struct {
	// Synchronization defines which synchronization capabilities the client supports.
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitempty"`

	// Completion capabilities specific to the `textDocument/completion` request.
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`

	// Hover capabilities specific to the `textDocument/hover` request.
	Hover *HoverClientCapabilities `json:"hover,omitempty"`

	// SignatureHelp capabilities specific to the `textDocument/signatureHelp` request.
	SignatureHelp *SignatureHelpClientCapabilities `json:"signatureHelp,omitempty"`

	// Declaration capabilities specific to the `textDocument/declaration` request.
	Declaration *DeclarationClientCapabilities `json:"declaration,omitempty"`

	// Definition capabilities specific to the `textDocument/definition` request.
	Definition *DefinitionClientCapabilities `json:"definition,omitempty"`

	// TypeDefinition capabilities specific to the `textDocument/typeDefinition` request.
	TypeDefinition *TypeDefinitionClientCapabilities `json:"typeDefinition,omitempty"`

	// Implementation capabilities specific to the `textDocument/implementation` request.
	Implementation *ImplementationClientCapabilities `json:"implementation,omitempty"`

	// References capabilities specific to the `textDocument/references` request.
	References *ReferenceClientCapabilities `json:"references,omitempty"`

	// DocumentHighlight capabilities specific to the `textDocument/documentHighlight` request.
	DocumentHighlight *DocumentHighlightClientCapabilities `json:"documentHighlight,omitempty"`

	// DocumentSymbol capabilities specific to the `textDocument/documentSymbol` request.
	DocumentSymbol *DocumentSymbolClientCapabilities `json:"documentSymbol,omitempty"`

	// CodeAction capabilities specific to the `textDocument/codeAction` request.
	CodeAction *CodeActionClientCapabilities `json:"codeAction,omitempty"`

	// CodeLens capabilities specific to the `textDocument/codeLens` request.
	CodeLens *CodeLensClientCapabilities `json:"codeLens,omitempty"`

	// DocumentLink capabilities specific to the `textDocument/documentLink` request.
	DocumentLink *DocumentLinkClientCapabilities `json:"documentLink,omitempty"`

	// ColorProvider capabilities specific to the `textDocument/documentColor` and the `textDocument/colorPresentation` request.
	ColorProvider *DocumentColorClientCapabilities `json:"colorProvider,omitempty"`

	// Formatting capabilities specific to the `textDocument/formatting` request.
	Formatting *DocumentFormattingClientCapabilities `json:"formatting,omitempty"`

	// RangeFormatting capabilities specific to the `textDocument/rangeFormatting` request.
	RangeFormatting *DocumentRangeFormattingClientCapabilities `json:"rangeFormatting,omitempty"`

	// OnTypeFormatting capabilities specific to the `textDocument/onTypeFormatting` request.
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitempty"`

	// Rename capabilities specific to the `textDocument/rename` request.
	Rename *RenameClientCapabilities `json:"rename,omitempty"`

	// FoldingRange capabilities specific to the `textDocument/foldingRange` request.
	FoldingRange *FoldingRangeClientCapabilities `json:"foldingRange,omitempty"`

	// SelectionRange capabilities specific to the `textDocument/selectionRange` request.
	SelectionRange *SelectionRangeClientCapabilities `json:"selectionRange,omitempty"`

	// PublishDiagnostics capabilities specific to the `textDocument/publishDiagnostics` notification.
	PublishDiagnostics *PublishDiagnosticsClientCapabilities `json:"publishDiagnostics,omitempty"`

	// CallHierarchy capabilities specific to the various call hierarchy requests.
	CallHierarchy *CallHierarchyClientCapabilities `json:"callHierarchy,omitempty"`

	// SemanticTokens capabilities specific to the various semantic token request.
	SemanticTokens *SemanticTokensClientCapabilities `json:"semanticTokens,omitempty"`

	// LinkedEditingRange capabilities specific to the `textDocument/linkedEditingRange` request.
	LinkedEditingRange *LinkedEditingRangeClientCapabilities `json:"linkedEditingRange,omitempty"`

	// Moniker client capabilities specific to the `textDocument/moniker` request.
	Moniker *MonikerClientCapabilities `json:"moniker,omitempty"`

	// TypeHierarchy capabilities specific to the various type hierarchy requests.
	TypeHierarchy *TypeHierarchyClientCapabilities `json:"typeHierarchy,omitempty"`

	// InlineValue capabilities specific to the `textDocument/inlineValue` request.
	InlineValue *InlineValueClientCapabilities `json:"inlineValue,omitempty"`

	// InlayHint capabilities specific to the `textDocument/inlayHint` request.
	InlayHint *InlayHintClientCapabilities `json:"inlayHint,omitempty"`

	// Diagnostic capabilities specific to the diagnostic pull model.
	Diagnostic *DiagnosticClientCapabilities `json:"diagnostic,omitempty"`

	// InlineCompletion client capabilities specific to inline completions.  3.18.0 @proposed.
	InlineCompletion *InlineCompletionClientCapabilities `json:"inlineCompletion,omitempty"`
}

// NotebookDocumentSyncClientCapabilities notebook specific client capabilities.
//
// @since 3.17.0
type NotebookDocumentSyncClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// ExecutionSummarySupport the client supports sending execution summary data per cell.
	//
	// @since 3.17.0
	ExecutionSummarySupport bool `json:"executionSummarySupport,omitempty"`
}

// NotebookDocumentClientCapabilities capabilities specific to the notebook document support.
//
// @since 3.17.0
type NotebookDocumentClientCapabilities struct {
	// Synchronization capabilities specific to notebook document synchronization
	// @since 3.17.0
	Synchronization NotebookDocumentSyncClientCapabilities `json:"synchronization"`
}

// ClientShowMessageActionItemOptions.
//
// @since 3.18.0
type ClientShowMessageActionItemOptions struct {
	// AdditionalPropertiesSupport whether the client supports additional attributes which are preserved and send back to the server in
	// the request's response.
	//
	// @since 3.18.0
	AdditionalPropertiesSupport bool `json:"additionalPropertiesSupport,omitempty"`
}

// ShowMessageRequestClientCapabilities show message request client capabilities.
type ShowMessageRequestClientCapabilities struct {
	// MessageActionItem capabilities specific to the `MessageActionItem` type.
	MessageActionItem *ClientShowMessageActionItemOptions `json:"messageActionItem,omitempty"`
}

// ShowDocumentClientCapabilities client capabilities for the showDocument request.
//
// @since 3.16.0
type ShowDocumentClientCapabilities struct {
	// Support the client has support for the showDocument request.
	//
	// @since 3.16.0
	Support bool `json:"support"`
}

type WindowClientCapabilities struct {
	// WorkDoneProgress it indicates whether the client supports server initiated progress using the `window/workDoneProgress/create` request. The capability also controls Whether client supports handling of progress notifications. If set servers are allowed to report a `workDoneProgress` property in the request specific server capabilities.
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`

	// ShowMessage capabilities specific to the showMessage request.
	ShowMessage *ShowMessageRequestClientCapabilities `json:"showMessage,omitempty"`

	// ShowDocument capabilities specific to the showDocument request.
	ShowDocument *ShowDocumentClientCapabilities `json:"showDocument,omitempty"`
}

// StaleRequestSupportOptions.
//
// @since 3.18.0
type StaleRequestSupportOptions struct {
	// Cancel the client will actively cancel the request.
	//
	// @since 3.18.0
	Cancel bool `json:"cancel"`

	// RetryOnContentModified the list of requests for which the client will retry the request if it receives a response with error code `ContentModified`.
	//
	// @since 3.18.0
	RetryOnContentModified []string `json:"retryOnContentModified"`
}

// RegularExpressionsClientCapabilities client capabilities specific to regular expressions.
//
// @since 3.16.0
type RegularExpressionsClientCapabilities struct {
	// Engine the engine's name.
	//
	// @since 3.16.0
	Engine RegularExpressionEngineKind `json:"engine"`

	// Version the engine's version.
	//
	// @since 3.16.0
	Version string `json:"version,omitempty"`
}

// MarkdownClientCapabilities client capabilities specific to the used markdown parser.
//
// @since 3.16.0
type MarkdownClientCapabilities struct {
	// Parser the name of the parser.
	//
	// @since 3.16.0
	Parser string `json:"parser"`

	// Version the version of the parser.
	//
	// @since 3.16.0
	Version string `json:"version,omitempty"`

	// AllowedTags a list of HTML tags that the client allows / supports in Markdown.
	// @since 3.16.0
	AllowedTags []string `json:"allowedTags,omitempty"`
}

// GeneralClientCapabilities general client capabilities.
//
// @since 3.16.0
type GeneralClientCapabilities struct {
	// StaleRequestSupport client capability that signals how the client handles stale requests (e.g. a request for which the client will not process the response anymore since the information is outdated).
	// @since 3.16.0
	StaleRequestSupport *StaleRequestSupportOptions `json:"staleRequestSupport,omitempty"`

	// RegularExpressions client capabilities specific to regular expressions.
	// @since 3.16.0
	RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions,omitempty"`

	// Markdown client capabilities specific to the client's markdown parser.
	// @since 3.16.0
	Markdown *MarkdownClientCapabilities `json:"markdown,omitempty"`

	// PositionEncodings the position encodings supported by the client. Client and server have to agree on the same position
	// encoding to ensure that offsets (e.g. character position in a line) are interpreted the same on both sides. To keep the protocol backwards compatible the following applies: if the value 'utf-16'
	// is missing from the array of position encodings servers can assume that the client supports UTF-16. UTF-16 is therefore a mandatory encoding. If omitted it defaults to ['utf-16']. Implementation
	// considerations: since the conversion from one encoding into another requires the content of the file / line the conversion is best done where the file is read which is usually on the server side.
	// @since 3.16.0
	PositionEncodings []PositionEncodingKind `json:"positionEncodings,omitempty"`
}

// ClientCapabilities defines the capabilities provided by the client.
type ClientCapabilities struct {
	// Workspace workspace specific client capabilities.
	Workspace *WorkspaceClientCapabilities `json:"workspace,omitempty"`

	// TextDocument text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`

	// NotebookDocument capabilities specific to the notebook document support.
	NotebookDocument *NotebookDocumentClientCapabilities `json:"notebookDocument,omitempty"`

	// Window window specific client capabilities.
	Window *WindowClientCapabilities `json:"window,omitempty"`

	// General general client capabilities.
	General *GeneralClientCapabilities `json:"general,omitempty"`

	// Experimental experimental client capabilities.
	Experimental any `json:"experimental,omitempty"`
}

// InitializeParamsBase the initialize parameters.
type InitializeParamsBase struct {
	// mixins
	WorkDoneProgressParams

	// ProcessID the process Id of the parent process that started the server. Is `null` if the process has not been started by another process. If the parent process is not alive then the server should exit.
	ProcessID int32 `json:"processId,omitempty"`

	// ClientInfo information about the client
	ClientInfo *ClientInfo `json:"clientInfo,omitempty"`

	// Locale the locale the client is currently showing the user interface in. This must not necessarily be the locale of the operating system. Uses IETF language tags as the value's syntax (See https://en.wikipedia.org/wiki/IETF_language_tag)
	Locale string `json:"locale,omitempty"`

	// RootPath the rootPath of the workspace. Is null if no folder is open. // // Deprecated: in favour of rootUri.
	RootPath string `json:"rootPath,omitempty"`

	// RootURI the rootUri of the workspace. Is null if no folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins. // // Deprecated: in favour of workspaceFolders.
	RootURI DocumentURI `json:"rootUri,omitempty"`

	// Capabilities the capabilities provided by the client (editor or tool).
	Capabilities ClientCapabilities `json:"capabilities"`

	// InitializationOptions user provided initialization options.
	InitializationOptions any `json:"initializationOptions,omitempty"`

	// Trace the initial trace setting. If omitted trace is disabled ('off').
	Trace TraceValue `json:"trace,omitempty"`
}

type WorkspaceFoldersInitializeParams struct {
	// WorkspaceFolders the workspace folders configured in the client when the server starts. This property is only available if the client supports workspace folders. It can be `null` if the client supports workspace folders but none are configured.
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

type InitializeParams struct {
	// extends
	InitializeParamsBase
	WorkspaceFoldersInitializeParams
}

// FileOperationOptions options for notifications/requests for user operations on files.
//
// @since 3.16.0
type FileOperationOptions struct {
	// DidCreate the server is interested in receiving didCreateFiles notifications.
	//
	// @since 3.16.0
	DidCreate *FileOperationRegistrationOptions `json:"didCreate,omitempty"`

	// WillCreate the server is interested in receiving willCreateFiles requests.
	//
	// @since 3.16.0
	WillCreate *FileOperationRegistrationOptions `json:"willCreate,omitempty"`

	// DidRename the server is interested in receiving didRenameFiles notifications.
	//
	// @since 3.16.0
	DidRename *FileOperationRegistrationOptions `json:"didRename,omitempty"`

	// WillRename the server is interested in receiving willRenameFiles requests.
	//
	// @since 3.16.0
	WillRename *FileOperationRegistrationOptions `json:"willRename,omitempty"`

	// DidDelete the server is interested in receiving didDeleteFiles file notifications.
	//
	// @since 3.16.0
	DidDelete *FileOperationRegistrationOptions `json:"didDelete,omitempty"`

	// WillDelete the server is interested in receiving willDeleteFiles file requests.
	//
	// @since 3.16.0
	WillDelete *FileOperationRegistrationOptions `json:"willDelete,omitempty"`
}

// WorkspaceOptions defines workspace specific capabilities of the server.
//
// @since 3.18.0
type WorkspaceOptions struct {
	// WorkspaceFolders the server supports workspace folder.
	// @since 3.18.0
	WorkspaceFolders *WorkspaceFoldersServerCapabilities `json:"workspaceFolders,omitempty"`

	// FileOperations the server is interested in notifications/requests for operations on files.
	// @since 3.18.0
	FileOperations *FileOperationOptions `json:"fileOperations,omitempty"`

	// TextDocumentContent the server supports the `workspace/textDocumentContent` request.  3.18.0 @proposed.
	// @since 3.18.0
	TextDocumentContent *WorkspaceOptionsTextDocumentContent `json:"textDocumentContent,omitempty"`
}

// ServerCapabilities defines the capabilities provided by a language server.
type ServerCapabilities struct {
	// PositionEncoding the position encoding the server picked from the encodings offered by the client via the client capability `general.positionEncodings`. If the client didn't provide any position encodings the only valid value that a server can return is 'utf-16'. If omitted it defaults to 'utf-16'.
	PositionEncoding PositionEncodingKind `json:"positionEncoding,omitempty"`

	// TextDocumentSync defines how text documents are synced. Is either a detailed structure defining each notification or for backwards compatibility the TextDocumentSyncKind number.
	TextDocumentSync *ServerCapabilitiesTextDocumentSync `json:"textDocumentSync,omitempty"`

	// NotebookDocumentSync defines how notebook documents are synced.
	NotebookDocumentSync *ServerCapabilitiesNotebookDocumentSync `json:"notebookDocumentSync,omitempty"`

	// CompletionProvider the server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	// HoverProvider the server provides hover support.
	HoverProvider *ServerCapabilitiesHoverProvider `json:"hoverProvider,omitempty"`

	// SignatureHelpProvider the server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	// DeclarationProvider the server provides Goto Declaration support.
	DeclarationProvider *ServerCapabilitiesDeclarationProvider `json:"declarationProvider,omitempty"`

	// DefinitionProvider the server provides goto definition support.
	DefinitionProvider *ServerCapabilitiesDefinitionProvider `json:"definitionProvider,omitempty"`

	// TypeDefinitionProvider the server provides Goto Type Definition support.
	TypeDefinitionProvider *ServerCapabilitiesTypeDefinitionProvider `json:"typeDefinitionProvider,omitempty"`

	// ImplementationProvider the server provides Goto Implementation support.
	ImplementationProvider *ServerCapabilitiesImplementationProvider `json:"implementationProvider,omitempty"`

	// ReferencesProvider the server provides find references support.
	ReferencesProvider *ServerCapabilitiesReferencesProvider `json:"referencesProvider,omitempty"`

	// DocumentHighlightProvider the server provides document highlight support.
	DocumentHighlightProvider *ServerCapabilitiesDocumentHighlightProvider `json:"documentHighlightProvider,omitempty"`

	// DocumentSymbolProvider the server provides document symbol support.
	DocumentSymbolProvider *ServerCapabilitiesDocumentSymbolProvider `json:"documentSymbolProvider,omitempty"`

	// CodeActionProvider the server provides code actions. CodeActionOptions may only be specified if the client states that it supports `codeActionLiteralSupport` in its initial `initialize` request.
	CodeActionProvider *ServerCapabilitiesCodeActionProvider `json:"codeActionProvider,omitempty"`

	// CodeLensProvider the server provides code lens.
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider,omitempty"`

	// DocumentLinkProvider the server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider,omitempty"`

	// ColorProvider the server provides color provider support.
	ColorProvider *ServerCapabilitiesColorProvider `json:"colorProvider,omitempty"`

	// WorkspaceSymbolProvider the server provides workspace symbol support.
	WorkspaceSymbolProvider *ServerCapabilitiesWorkspaceSymbolProvider `json:"workspaceSymbolProvider,omitempty"`

	// DocumentFormattingProvider the server provides document formatting.
	DocumentFormattingProvider *ServerCapabilitiesDocumentFormattingProvider `json:"documentFormattingProvider,omitempty"`

	// DocumentRangeFormattingProvider the server provides document range formatting.
	DocumentRangeFormattingProvider *ServerCapabilitiesDocumentRangeFormattingProvider `json:"documentRangeFormattingProvider,omitempty"`

	// DocumentOnTypeFormattingProvider the server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`

	// RenameProvider the server provides rename support. RenameOptions may only be specified if the client states that it
	// supports `prepareSupport` in its initial `initialize` request.
	RenameProvider *ServerCapabilitiesRenameProvider `json:"renameProvider,omitempty"`

	// FoldingRangeProvider the server provides folding provider support.
	FoldingRangeProvider *ServerCapabilitiesFoldingRangeProvider `json:"foldingRangeProvider,omitempty"`

	// SelectionRangeProvider the server provides selection range support.
	SelectionRangeProvider *ServerCapabilitiesSelectionRangeProvider `json:"selectionRangeProvider,omitempty"`

	// ExecuteCommandProvider the server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider,omitempty"`

	// CallHierarchyProvider the server provides call hierarchy support.
	CallHierarchyProvider *ServerCapabilitiesCallHierarchyProvider `json:"callHierarchyProvider,omitempty"`

	// LinkedEditingRangeProvider the server provides linked editing range support.
	LinkedEditingRangeProvider *ServerCapabilitiesLinkedEditingRangeProvider `json:"linkedEditingRangeProvider,omitempty"`

	// SemanticTokensProvider the server provides semantic tokens support.
	SemanticTokensProvider *ServerCapabilitiesSemanticTokensProvider `json:"semanticTokensProvider,omitempty"`

	// MonikerProvider the server provides moniker support.
	MonikerProvider *ServerCapabilitiesMonikerProvider `json:"monikerProvider,omitempty"`

	// TypeHierarchyProvider the server provides type hierarchy support.
	TypeHierarchyProvider *ServerCapabilitiesTypeHierarchyProvider `json:"typeHierarchyProvider,omitempty"`

	// InlineValueProvider the server provides inline values.
	InlineValueProvider *ServerCapabilitiesInlineValueProvider `json:"inlineValueProvider,omitempty"`

	// InlayHintProvider the server provides inlay hints.
	InlayHintProvider *ServerCapabilitiesInlayHintProvider `json:"inlayHintProvider,omitempty"`

	// DiagnosticProvider the server has support for pull model diagnostics.
	DiagnosticProvider *ServerCapabilitiesDiagnosticProvider `json:"diagnosticProvider,omitempty"`

	// InlineCompletionProvider inline completion options used during static registration.  3.18.0 @proposed.
	InlineCompletionProvider *ServerCapabilitiesInlineCompletionProvider `json:"inlineCompletionProvider,omitempty"`

	// Workspace workspace specific server capabilities.
	Workspace *WorkspaceOptions `json:"workspace,omitempty"`

	// Experimental experimental server capabilities.
	Experimental any `json:"experimental,omitempty"`
}

// ServerInfo information about the server  3.15.0  3.18.0 ServerInfo type name added.
//
// @since 3.18.0 ServerInfo type name added.
type ServerInfo struct {
	// Name the name of the server as defined by the server.
	//
	// @since 3.18.0 ServerInfo type name added.
	Name string `json:"name"`

	// Version the server's version as defined by the server.
	//
	// @since 3.18.0 ServerInfo type name added.
	Version string `json:"version,omitempty"`
}

// InitializeResult the result returned from an initialize request.
type InitializeResult struct {
	// Capabilities the capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`

	// ServerInfo information about the server.
	ServerInfo *ServerInfo `json:"serverInfo,omitempty"`
}

// InitializeError the data type of the ResponseError if the initialize request fails.
type InitializeError struct {
	// Retry indicates whether the client execute the following retry logic: (1) show the message provided by the
	// ResponseError to the user (2) user selects retry or cancel (3) if user selected retry the initialize method is sent again.
	Retry bool `json:"retry"`
}

type InitializedParams struct{}

type SetTraceParams struct {
	Value TraceValue `json:"value"`
}

type LogTraceParams struct {
	Message string `json:"message"`

	Verbose string `json:"verbose,omitempty"`
}

// StaticRegistrationOptions static registration options to be returned in the initialize request.
type StaticRegistrationOptions struct {
	// ID the id used to register the request. The id can be used to deregister the request again. See also Registration#id.
	ID string `json:"id,omitempty"`
}
