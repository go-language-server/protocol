// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

// RegularExpressionsClientCapabilities represents a client capabilities specific to regular expressions.
//
// The following features from the ECMAScript 2020 regular expression specification are NOT mandatory for a client:
//
//  Assertions
// Lookahead assertion, Negative lookahead assertion, lookbehind assertion, negative lookbehind assertion.
//  Character classes
// Matching control characters using caret notation (e.g. "\cX") and matching UTF-16 code units (e.g. "\uhhhh").
//  Group and ranges
// Named capturing groups.
//  Unicode property escapes
// None of the features needs to be supported.
//
// The only regular expression flag that a client needs to support is "i" to specify a case insensitive search.
//
// @since 3.16.0.
type RegularExpressionsClientCapabilities struct {
	// Engine is the engine's name.
	//
	// Well known engine name is "ECMAScript".
	//  https://tc39.es/ecma262/#sec-regexp-regular-expression-objects
	//  https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions
	Engine string `json:"engine"`

	// Version is the engine's version.
	//
	// Well known engine version is "ES2020".
	//  https://tc39.es/ecma262/#sec-regexp-regular-expression-objects
	//  https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions
	Version string `json:"version,omitempty"`
}

// WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport is the ChangeAnnotationSupport of WorkspaceClientCapabilitiesWorkspaceEdit.
//
// @since 3.16.0.
type WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport struct {
	// GroupsOnLabel whether the client groups edits with equal labels into tree nodes,
	// for instance all edits labeled with "Changes in Strings" would
	// be a tree node.
	GroupsOnLabel bool `json:"groupsOnLabel,omitempty"`
}

// WorkspaceClientCapabilitiesWorkspaceEdit capabilities specific to "WorkspaceEdit"s.
type WorkspaceClientCapabilitiesWorkspaceEdit struct {
	// DocumentChanges is the client supports versioned document changes in `WorkspaceEdit`s
	DocumentChanges bool `json:"documentChanges,omitempty"`

	// FailureHandling is the failure handling strategy of a client if applying the workspace edit
	// fails.
	FailureHandling string `json:"failureHandling,omitempty"`

	// ResourceOperations is the resource operations the client supports. Clients should at least
	// support "create", "rename" and "delete" files and folders.
	ResourceOperations []string `json:"resourceOperations,omitempty"`

	// NormalizesLineEndings whether the client normalizes line endings to the client specific
	// setting.
	// If set to `true` the client will normalize line ending characters
	// in a workspace edit to the client specific new line character(s).
	//
	// @since 3.16.0.
	NormalizesLineEndings bool `json:"normalizesLineEndings,omitempty"`

	// ChangeAnnotationSupport whether the client in general supports change annotations on text edits,
	// create file, rename file and delete file changes.
	//
	// @since 3.16.0.
	ChangeAnnotationSupport *WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport `json:"changeAnnotationSupport,omitempty"`
}

// WorkspaceClientCapabilitiesDidChangeConfiguration capabilities specific to the "workspace/didChangeConfiguration" notification.
//
// @since 3.16.0.
type WorkspaceClientCapabilitiesDidChangeConfiguration struct {
	// Did change configuration notification supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// WorkspaceClientCapabilitiesDidChangeWatchedFiles capabilities specific to the "workspace/didChangeWatchedFiles" notification.
//
// @since 3.16.0.
type WorkspaceClientCapabilitiesDidChangeWatchedFiles struct {
	// Did change watched files notification supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// WorkspaceClientCapabilitiesSymbol capabilities specific to the `workspace/symbol` request.
type WorkspaceClientCapabilitiesSymbol struct {
	// Symbol request supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
	SymbolKind *WorkspaceClientCapabilitiesSymbolKind `json:"symbolKind,omitempty"`
}

// WorkspaceClientCapabilitiesSymbolKind specific capabilities for the "SymbolKind" in the "workspace/symbol" request.
type WorkspaceClientCapabilitiesSymbolKind struct {
	// ValueSet is the symbol kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	//
	// If this property is not present the client only supports
	// the symbol kinds from `File` to `Array` as defined in
	// the initial version of the protocol.
	ValueSet []SymbolKind `json:"valueSet,omitempty"`
}

// WorkspaceClientCapabilitiesExecuteCommand capabilities specific to the "workspace/executeCommand" request.
type WorkspaceClientCapabilitiesExecuteCommand struct {
	// DynamicRegistration Execute command supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// WorkspaceClientCapabilitiesSemanticTokens capabilities specific to the "workspace/semanticToken" request.
//
// @since 3.16.0.
type WorkspaceClientCapabilitiesSemanticTokens struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from
	// the server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// semantic tokens currently shown. It should be used with absolute care
	// and is useful for situation where a server for example detect a project
	// wide change that requires such a calculation.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// WorkspaceClientCapabilitiesCodeLens capabilities specific to the "workspace/codeLens" request.
//
// @since 3.16.0.
type WorkspaceClientCapabilitiesCodeLens struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the
	// server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// code lenses currently shown. It should be used with absolute care and is
	// useful for situation where a server for example detect a project wide
	// change that requires such a calculation.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// WorkspaceClientCapabilitiesFileOperations capabilities specific to the fileOperations.
//
// @since 3.16.0.
type WorkspaceClientCapabilitiesFileOperations struct {
	// DynamicRegistration whether the client supports dynamic registration for file
	// requests/notifications.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// DidCreate is the client has support for sending didCreateFiles notifications.
	DidCreate bool `json:"didCreate,omitempty"`

	// WillCreate is the client has support for sending willCreateFiles requests.
	WillCreate bool `json:"willCreate,omitempty"`

	// DidRename is the client has support for sending didRenameFiles notifications.
	DidRename bool `json:"didRename,omitempty"`

	// WillRename is the client has support for sending willRenameFiles requests.
	WillRename bool `json:"willRename,omitempty"`

	// DidDelete is the client has support for sending didDeleteFiles notifications.
	DidDelete bool `json:"didDelete,omitempty"`

	// WillDelete is the client has support for sending willDeleteFiles requests.
	WillDelete bool `json:"willDelete,omitempty"`
}

// WorkspaceClientCapabilities Workspace specific client capabilities.
type WorkspaceClientCapabilities struct {
	// The client supports applying batch edits to the workspace by supporting
	// the request "workspace/applyEdit".
	ApplyEdit bool `json:"applyEdit,omitempty"`

	// WorkspaceEdit capabilities specific to `WorkspaceEdit`s.
	WorkspaceEdit *WorkspaceClientCapabilitiesWorkspaceEdit `json:"workspaceEdit,omitempty"`

	// DidChangeConfiguration capabilities specific to the `workspace/didChangeConfiguration` notification.
	DidChangeConfiguration *WorkspaceClientCapabilitiesDidChangeConfiguration `json:"didChangeConfiguration,omitempty"`

	// DidChangeWatchedFiles capabilities specific to the `workspace/didChangeWatchedFiles` notification.
	DidChangeWatchedFiles *WorkspaceClientCapabilitiesDidChangeWatchedFiles `json:"didChangeWatchedFiles,omitempty"`

	// Symbol capabilities specific to the "workspace/symbol" request.
	Symbol *WorkspaceClientCapabilitiesSymbol `json:"symbol,omitempty"`

	// ExecuteCommand capabilities specific to the "workspace/executeCommand" request.
	ExecuteCommand *WorkspaceClientCapabilitiesExecuteCommand `json:"executeCommand,omitempty"`

	// WorkspaceFolders is the client has support for workspace folders.
	//
	// @since 3.6.0.
	WorkspaceFolders bool `json:"workspaceFolders,omitempty"`

	// Configuration is the client supports "workspace/configuration" requests.
	//
	// @since 3.6.0.
	Configuration bool `json:"configuration,omitempty"`

	// SemanticTokens is the capabilities specific to the semantic token requests scoped to the
	// workspace.
	//
	// @since 3.16.0.
	SemanticTokens *WorkspaceClientCapabilitiesSemanticTokens `json:"semanticTokens,omitempty"`

	// CodeLens is the Capabilities specific to the code lens requests scoped to the
	// workspace.
	//
	// @since 3.16.0.
	CodeLens *WorkspaceClientCapabilitiesCodeLens `json:"codeLens,omitempty"`

	// FileOperations is the client has support for file requests/notifications.
	//
	// @since 3.16.0.
	FileOperations *WorkspaceClientCapabilitiesFileOperations `json:"fileOperations,omitempty"`
}

// TextDocumentClientCapabilitiesSynchronization defines which synchronization capabilities the client supports.
type TextDocumentClientCapabilitiesSynchronization struct {
	// DidSave is the client supports did save notifications.
	DidSave bool `json:"didSave,omitempty"`

	// DynamicRegistration whether text document synchronization supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// WillSave is the client supports sending will save notifications.
	WillSave bool `json:"willSave,omitempty"`

	// WillSaveWaitUntil is the client supports sending a will save request and
	// waits for a response providing text edits which will
	// be applied to the document before it is saved.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`
}

// TextDocumentClientCapabilitiesCompletion Capabilities specific to the "textDocument/completion".
type TextDocumentClientCapabilitiesCompletion struct {
	// Whether completion supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The client supports the following `CompletionItem` specific
	// capabilities.
	CompletionItem *TextDocumentClientCapabilitiesCompletionItem `json:"completionItem,omitempty"`

	CompletionItemKind *TextDocumentClientCapabilitiesCompletionItemKind `json:"completionItemKind,omitempty"`

	// ContextSupport is the client supports to send additional context information for a
	// `textDocument/completion` request.
	ContextSupport bool `json:"contextSupport,omitempty"`
}

// TextDocumentClientCapabilitiesCompletionItemKind specific capabilities for the "CompletionItemKind" in the "textDocument/completion" request.
type TextDocumentClientCapabilitiesCompletionItemKind struct {
	// The completion item kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	//
	// If this property is not present the client only supports
	// the completion items kinds from `Text` to `Reference` as defined in
	// the initial version of the protocol.
	//
	ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
}

// TextDocumentClientCapabilitiesCompletionItemTagSupport specific capabilities for the "TagSupport" in the "textDocument/completion" request.
//
// @since 3.15.0.
type TextDocumentClientCapabilitiesCompletionItemTagSupport struct {
	// ValueSet is the tags supported by the client.
	//
	// @since 3.15.0.
	ValueSet []CompletionItemTag `json:"valueSet,omitempty"`
}

// TextDocumentClientCapabilitiesCompletionItemResolveSupport specific capabilities for the ResolveSupport in the TextDocumentClientCapabilitiesCompletionItem.
//
// @since 3.16.0.
type TextDocumentClientCapabilitiesCompletionItemResolveSupport struct {
	// Properties is the properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}

// TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport specific capabilities for the InsertTextModeSupport in the TextDocumentClientCapabilitiesCompletionItem.
//
// @since 3.16.0.
type TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport struct {
	// ValueSet is the tags supported by the client.
	//
	// @since 3.16.0.
	ValueSet []InsertTextMode `json:"valueSet,omitempty"`
}

// TextDocumentClientCapabilitiesCompletionItem is the client supports the following "CompletionItem" specific
// capabilities.
type TextDocumentClientCapabilitiesCompletionItem struct {
	// SnippetSupport client supports snippets as insert text.
	//
	// A snippet can define tab stops and placeholders with `$1`, `$2`
	// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	// the end of the snippet. Placeholders with equal identifiers are linked,
	// that is typing in one will update others too.
	SnippetSupport bool `json:"snippetSupport,omitempty"`

	// CommitCharactersSupport client supports commit characters on a completion item.
	CommitCharactersSupport bool `json:"commitCharactersSupport,omitempty"`

	// DocumentationFormat client supports the follow content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

	// DeprecatedSupport client supports the deprecated property on a completion item.
	DeprecatedSupport bool `json:"deprecatedSupport,omitempty"`

	// PreselectSupport client supports the preselect property on a completion item.
	PreselectSupport bool `json:"preselectSupport,omitempty"`

	// TagSupport is the client supports the tag property on a completion item.
	//
	// Clients supporting tags have to handle unknown tags gracefully.
	// Clients especially need to preserve unknown tags when sending
	// a completion item back to the server in a resolve call.
	//
	// @since 3.15.0.
	TagSupport *TextDocumentClientCapabilitiesCompletionItemTagSupport `json:"tagSupport,omitempty"`

	// InsertReplaceSupport client supports insert replace edit to control different behavior if
	// a completion item is inserted in the text or should replace text.
	//
	// @since 3.16.0.
	InsertReplaceSupport bool `json:"insertReplaceSupport,omitempty"`

	// ResolveSupport indicates which properties a client can resolve lazily on a
	// completion item. Before version 3.16.0 only the predefined properties
	// `documentation` and `details` could be resolved lazily.
	//
	// @since 3.16.0.
	ResolveSupport *TextDocumentClientCapabilitiesCompletionItemResolveSupport `json:"resolveSupport,omitempty"`

	// InsertTextModeSupport is the client supports the `insertTextMode` property on
	// a completion item to override the whitespace handling mode
	// as defined by the client (see `insertTextMode`).
	//
	// @since 3.16.0.
	InsertTextModeSupport *TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport `json:"insertTextModeSupport,omitempty"`
}

// TextDocumentClientCapabilitiesHover capabilities specific to the "textDocument/hover".
type TextDocumentClientCapabilitiesHover struct {
	// DynamicRegistration whether hover supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// ContentFormat is the client supports the follow content formats for the content
	// property. The order describes the preferred format of the client.
	ContentFormat []MarkupKind `json:"contentFormat,omitempty"`
}

// TextDocumentClientCapabilitiesSignatureHelp capabilities specific to the "textDocument/signatureHelp".
type TextDocumentClientCapabilitiesSignatureHelp struct {
	// DynamicRegistration whether signature help supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// SignatureInformation is the client supports the following "SignatureInformation"
	// specific properties.
	SignatureInformation *TextDocumentClientCapabilitiesSignatureInformation `json:"signatureInformation,omitempty"`

	// ContextSupport is the client supports to send additional context information for a "textDocument/signatureHelp" request.
	//
	// A client that opts into contextSupport will also support the "retriggerCharacters" on "SignatureHelpOptions".
	//
	// @since 3.15.0.
	ContextSupport bool `json:"contextSupport,omitempty"`
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

// TextDocumentClientCapabilitiesReferences capabilities specific to the "textDocument/references".
type TextDocumentClientCapabilitiesReferences struct {
	// DynamicRegistration whether references supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilitiesDocumentHighlight capabilities specific to the "textDocument/documentHighlight".
type TextDocumentClientCapabilitiesDocumentHighlight struct {
	// DynamicRegistration Whether document highlight supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilitiesDocumentSymbol capabilities specific to the "textDocument/documentSymbol".
type TextDocumentClientCapabilitiesDocumentSymbol struct {
	// DynamicRegistration whether document symbol supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// SymbolKind specific capabilities for the "SymbolKind".
	SymbolKind *WorkspaceClientCapabilitiesSymbolKind `json:"symbolKind,omitempty"`

	// HierarchicalDocumentSymbolSupport is the client support hierarchical document symbols.
	HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`

	// TagSupport is the client supports tags on "SymbolInformation". Tags are supported on
	// "DocumentSymbol" if "HierarchicalDocumentSymbolSupport" is set to true.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.16.0.
	TagSupport *TextDocumentClientCapabilitiesDocumentSymbolTagSupport `json:"tagSupport,omitempty"`

	// LabelSupport is the client supports an additional label presented in the UI when
	// registering a document symbol provider.
	//
	// @since 3.16.0.
	LabelSupport bool `json:"labelSupport,omitempty"`
}

// TextDocumentClientCapabilitiesDocumentSymbolTagSupport TagSupport in the TextDocumentClientCapabilitiesDocumentSymbol.
//
// @since 3.16.0.
type TextDocumentClientCapabilitiesDocumentSymbolTagSupport struct {
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

// TextDocumentClientCapabilitiesDeclaration capabilities specific to the "textDocument/declaration".
type TextDocumentClientCapabilitiesDeclaration struct {
	// DynamicRegistration whether declaration supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// LinkSupport is the client supports additional metadata in the form of declaration links.
	//
	// @since 3.14.0.
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// TextDocumentClientCapabilitiesDefinition capabilities specific to the "textDocument/definition".
//
// @since 3.14.0.
type TextDocumentClientCapabilitiesDefinition struct {
	// DynamicRegistration whether definition supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// LinkSupport is the client supports additional metadata in the form of definition links.
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// TextDocumentClientCapabilitiesTypeDefinition capabilities specific to the "textDocument/typeDefinition".
//
// @since 3.6.0.
type TextDocumentClientCapabilitiesTypeDefinition struct {
	// DynamicRegistration whether typeDefinition supports dynamic registration. If this is set to `true`
	// the client supports the new "(TextDocumentRegistrationOptions & StaticRegistrationOptions)"
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// LinkSupport is the client supports additional metadata in the form of definition links.
	//
	// @since 3.14.0
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// TextDocumentClientCapabilitiesImplementation capabilities specific to the "textDocument/implementation".
//
// @since 3.6.0.
type TextDocumentClientCapabilitiesImplementation struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new "(TextDocumentRegistrationOptions & StaticRegistrationOptions)"
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// LinkSupport is the client supports additional metadata in the form of definition links.
	//
	// @since 3.14.0
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// TextDocumentClientCapabilitiesCodeAction capabilities specific to the "textDocument/codeAction".
type TextDocumentClientCapabilitiesCodeAction struct {
	// DynamicRegistration whether code action supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// CodeActionLiteralSupport is the client support code action literals as a valid
	// response of the "textDocument/codeAction" request.
	//
	// @since 3.8.0
	CodeActionLiteralSupport *TextDocumentClientCapabilitiesCodeActionLiteralSupport `json:"codeActionLiteralSupport,omitempty"`

	// IsPreferredSupport whether code action supports the "isPreferred" property.
	//
	// @since 3.15.0.
	IsPreferredSupport bool `json:"isPreferredSupport,omitempty"`

	// DisabledSupport whether code action supports the `disabled` property.
	//
	// @since 3.16.0.
	DisabledSupport bool `json:"disabledSupport,omitempty"`

	// DataSupport whether code action supports the `data` property which is
	// preserved between a `textDocument/codeAction` and a
	// `codeAction/resolve` request.
	//
	// @since 3.16.0.
	DataSupport bool `json:"dataSupport,omitempty"`

	// ResolveSupport whether the client supports resolving additional code action
	// properties via a separate `codeAction/resolve` request.
	//
	// @since 3.16.0.
	ResolveSupport *TextDocumentClientCapabilitiesCodeActionResolveSupport `json:"resolveSupport,omitempty"`

	// HonorsChangeAnnotations whether the client honors the change annotations in
	// text edits and resource operations returned via the
	// `CodeAction#edit` property by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	//
	// @since 3.16.0.
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`
}

// TextDocumentClientCapabilitiesCodeActionResolveSupport ResolveSupport in the TextDocumentClientCapabilitiesCodeAction.
//
// @since 3.16.0.
type TextDocumentClientCapabilitiesCodeActionResolveSupport struct {
	// Properties is the properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}

// TextDocumentClientCapabilitiesCodeActionLiteralSupport is the client support code action literals as a valid response of the "textDocument/codeAction" request.
type TextDocumentClientCapabilitiesCodeActionLiteralSupport struct {
	// CodeActionKind is the code action kind is support with the following value
	// set.
	CodeActionKind *TextDocumentClientCapabilitiesCodeActionKind `json:"codeActionKind"`
}

// TextDocumentClientCapabilitiesCodeActionKind is the code action kind is support with the following value set.
type TextDocumentClientCapabilitiesCodeActionKind struct {
	// ValueSet is the code action kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	ValueSet []CodeActionKind `json:"valueSet"`
}

// TextDocumentClientCapabilitiesCodeLens capabilities specific to the "textDocument/codeLens".
type TextDocumentClientCapabilitiesCodeLens struct {
	// DynamicRegistration Whether code lens supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// TooltipSupport whether the client support the "tooltip" property on "DocumentLink".
	//
	// @since 3.15.0.
	TooltipSupport bool `json:"tooltipSupport,omitempty"`
}

// TextDocumentClientCapabilitiesDocumentLink capabilities specific to the "textDocument/documentLink".
type TextDocumentClientCapabilitiesDocumentLink struct {
	// DynamicRegistration whether document link supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// TooltipSupport whether the client supports the "tooltip" property on "DocumentLink".
	//
	// @since 3.15.0.
	TooltipSupport bool `json:"tooltipSupport,omitempty"`
}

// TextDocumentClientCapabilitiesColorProvider capabilities specific to the "textDocument/documentColor" and the
// "textDocument/colorPresentation" request.
//
// @since 3.6.0.
type TextDocumentClientCapabilitiesColorProvider struct {
	// DynamicRegistration whether colorProvider supports dynamic registration. If this is set to `true`
	// the client supports the new "(ColorProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)"
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilitiesRename capabilities specific to the "textDocument/rename".
type TextDocumentClientCapabilitiesRename struct {
	// DynamicRegistration whether rename supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// PrepareSupport is the client supports testing for validity of rename operations
	// before execution.
	PrepareSupport bool `json:"prepareSupport,omitempty"`

	// PrepareSupportDefaultBehavior client supports the default behavior result
	// (`{ defaultBehavior: boolean }`).
	//
	// The value indicates the default behavior used by the
	// client.
	//
	// @since 3.16.0.
	PrepareSupportDefaultBehavior PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior,omitempty"`

	// HonorsChangeAnnotations whether th client honors the change annotations in
	// text edits and resource operations returned via the
	// rename request's workspace edit by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	//
	// @since 3.16.0.
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`
}

// TextDocumentClientCapabilitiesPublishDiagnostics capabilities specific to "textDocument/publishDiagnostics".
type TextDocumentClientCapabilitiesPublishDiagnostics struct {
	// RelatedInformation whether the clients accepts diagnostics with related information.
	RelatedInformation bool `json:"relatedInformation,omitempty"`

	// TagSupport clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.15.0.
	TagSupport *TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport `json:"tagSupport,omitempty"`

	// VersionSupport whether the client interprets the version property of the
	// "textDocument/publishDiagnostics" notification`s parameter.
	//
	// @since 3.15.0.
	VersionSupport bool `json:"versionSupport,omitempty"`

	// CodeDescriptionSupport client supports a codeDescription property
	//
	// @since 3.16.0.
	CodeDescriptionSupport bool `json:"codeDescriptionSupport,omitempty"`

	// DataSupport whether code action supports the `data` property which is
	// preserved between a `textDocument/publishDiagnostics` and
	// `textDocument/codeAction` request.
	//
	// @since 3.16.0.
	DataSupport bool `json:"dataSupport,omitempty"`
}

// TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport is the client capacity of TagSupport.
//
// @since 3.15.0.
type TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport struct {
	// ValueSet is the tags supported by the client.
	ValueSet []DiagnosticTag `json:"valueSet"`
}

// TextDocumentClientCapabilitiesFoldingRange capabilities specific to "textDocument/foldingRange" requests.
//
// @since 3.10.0.
type TextDocumentClientCapabilitiesFoldingRange struct {
	// DynamicRegistration whether implementation supports dynamic registration for folding range providers. If this is set to `true`
	// the client supports the new "(FoldingRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)"
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// RangeLimit is the maximum number of folding ranges that the client prefers to receive per document. The value serves as a
	// hint, servers are free to follow the limit.
	RangeLimit uint32 `json:"rangeLimit,omitempty"`

	// LineFoldingOnly if set, the client signals that it only supports folding complete lines. If set, client will
	// ignore specified "startCharacter" and "endCharacter" properties in a FoldingRange.
	LineFoldingOnly bool `json:"lineFoldingOnly,omitempty"`
}

// TextDocumentClientCapabilitiesSelectionRange capabilities specific to "textDocument/selectionRange" requests.
//
// @since 3.16.0.
type TextDocumentClientCapabilitiesSelectionRange struct {
	// DynamicRegistration whether implementation supports dynamic registration for selection range providers. If this is set to `true`
	// the client supports the new "(SelectionRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)"
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilitiesLinkedEditingRange capabilities specific to "textDocument/linkedEditingRange" requests.
//
// @since 3.16.0.
type TextDocumentClientCapabilitiesLinkedEditingRange struct {
	// DynamicRegistration whether implementation supports dynamic registration.
	// If this is set to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilitiesCallHierarchy capabilities specific to "textDocument/callHierarchy" requests.
//
// @since 3.16.0.
type TextDocumentClientCapabilitiesCallHierarchy struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.}
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// WorkspaceClientCapabilitiesSemanticTokensRequests capabilities specific to the "textDocument/semanticTokens/xxx" request.
//
// @since 3.16.0.
type WorkspaceClientCapabilitiesSemanticTokensRequests struct {
	// Range is the client will send the "textDocument/semanticTokens/range" request
	// if the server provides a corresponding handler.
	Range bool `json:"range,omitempty"`

	// Full is the client will send the "textDocument/semanticTokens/full" request
	// if the server provides a corresponding handler.
	Full bool `json:"full,omitempty"` // delta?: boolean
}

// TextDocumentClientCapabilitiesSemanticTokens capabilities specific to the "textDocument.semanticTokens" request.
//
// @since 3.16.0.
type TextDocumentClientCapabilitiesSemanticTokens struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Requests which requests the client supports and might send to the server
	// depending on the server's capability. Please note that clients might not
	// show semantic tokens or degrade some of the user experience if a range
	// or full request is advertised by the client but not provided by the
	// server. If for example the client capability `requests.full` and
	// `request.range` are both set to true but the server only provides a
	// range provider the client might not render a minimap correctly or might
	// even decide to not show any semantic tokens at all.
	Requests WorkspaceClientCapabilitiesSemanticTokensRequests `json:"requests"`

	// TokenTypes is the token types that the client supports.
	TokenTypes []string `json:"tokenTypes"`

	// TokenModifiers is the token modifiers that the client supports.
	TokenModifiers []string `json:"tokenModifiers"`

	// Formats is the formats the clients supports.
	Formats []TokenFormat `json:"formats"`

	// OverlappingTokenSupport whether the client supports tokens that can overlap each other.
	OverlappingTokenSupport bool `json:"overlappingTokenSupport,omitempty"`

	// MultilineTokenSupport whether the client supports tokens that can span multiple lines.
	MultilineTokenSupport bool `json:"multilineTokenSupport,omitempty"`
}

// TextDocumentClientCapabilitiesMoniker capabilities specific to the "textDocument/moniker" request.
//
// @since 3.16.0.
type TextDocumentClientCapabilitiesMoniker struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.// DynamicRegistration whether implementation supports dynamic registration. If this is set to
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilities Text document specific client capabilities.
type TextDocumentClientCapabilities struct {
	// Synchronization defines which synchronization capabilities the client supports.
	Synchronization *TextDocumentClientCapabilitiesSynchronization `json:"synchronization,omitempty"`

	// Completion Capabilities specific to the "textDocument/completion".
	Completion *TextDocumentClientCapabilitiesCompletion `json:"completion,omitempty"`

	// Hover capabilities specific to the "textDocument/hover".
	Hover *TextDocumentClientCapabilitiesHover `json:"hover,omitempty"`

	// SignatureHelp capabilities specific to the "textDocument/signatureHelp".
	SignatureHelp *TextDocumentClientCapabilitiesSignatureHelp `json:"signatureHelp,omitempty"`

	// References capabilities specific to the "textDocument/references".
	References *TextDocumentClientCapabilitiesReferences `json:"references,omitempty"`

	// DocumentHighlight capabilities specific to the "textDocument/documentHighlight".
	DocumentHighlight *TextDocumentClientCapabilitiesDocumentHighlight `json:"documentHighlight,omitempty"`

	// DocumentSymbol capabilities specific to the "textDocument/documentSymbol".
	DocumentSymbol *TextDocumentClientCapabilitiesDocumentSymbol `json:"documentSymbol,omitempty"`

	// Formatting capabilities specific to the "textDocument/formatting".
	Formatting *TextDocumentClientCapabilitiesFormatting `json:"formatting,omitempty"`

	// RangeFormatting capabilities specific to the "textDocument/rangeFormatting".
	RangeFormatting *TextDocumentClientCapabilitiesRangeFormatting `json:"rangeFormatting,omitempty"`

	// OnTypeFormatting Capabilities specific to the "textDocument/onTypeFormatting".
	OnTypeFormatting *TextDocumentClientCapabilitiesOnTypeFormatting `json:"onTypeFormatting,omitempty"`

	// Declaration capabilities specific to the "textDocument/declaration".
	Declaration *TextDocumentClientCapabilitiesDeclaration `json:"declaration,omitempty"`

	// Definition capabilities specific to the "textDocument/definition".
	//
	// @since 3.14.0.
	Definition *TextDocumentClientCapabilitiesDefinition `json:"definition,omitempty"`

	// TypeDefinition capabilities specific to the "textDocument/typeDefinition".
	//
	// @since 3.6.0.
	TypeDefinition *TextDocumentClientCapabilitiesTypeDefinition `json:"typeDefinition,omitempty"`

	// Implementation capabilities specific to the "textDocument/implementation".
	//
	// @since 3.6.0.
	Implementation *TextDocumentClientCapabilitiesImplementation `json:"implementation,omitempty"`

	// CodeAction capabilities specific to the "textDocument/codeAction".
	CodeAction *TextDocumentClientCapabilitiesCodeAction `json:"codeAction,omitempty"`

	// CodeLens capabilities specific to the "textDocument/codeLens".
	CodeLens *TextDocumentClientCapabilitiesCodeLens `json:"codeLens,omitempty"`

	// DocumentLink capabilities specific to the "textDocument/documentLink".
	DocumentLink *TextDocumentClientCapabilitiesDocumentLink `json:"documentLink,omitempty"`

	// ColorProvider capabilities specific to the "textDocument/documentColor" and the
	// "textDocument/colorPresentation" request.
	//
	// @since 3.6.0.
	ColorProvider *TextDocumentClientCapabilitiesColorProvider `json:"colorProvider,omitempty"`

	// Rename capabilities specific to the "textDocument/rename".
	Rename *TextDocumentClientCapabilitiesRename `json:"rename,omitempty"`

	// PublishDiagnostics capabilities specific to "textDocument/publishDiagnostics".
	PublishDiagnostics *TextDocumentClientCapabilitiesPublishDiagnostics `json:"publishDiagnostics,omitempty"`

	// FoldingRange capabilities specific to "textDocument/foldingRange" requests.
	//
	// @since 3.10.0.
	FoldingRange *TextDocumentClientCapabilitiesFoldingRange `json:"foldingRange,omitempty"`

	// SelectionRange capabilities specific to "textDocument/selectionRange" requests.
	//
	// @since 3.15.0.
	SelectionRange *TextDocumentClientCapabilitiesSelectionRange `json:"selectionRange,omitempty"`

	// LinkedEditingRange capabilities specific to the "textDocument/linkedEditingRange" request.
	//
	// @since 3.16.0.
	LinkedEditingRange *TextDocumentClientCapabilitiesLinkedEditingRange `json:"linkedEditingRange,omitempty"`

	// CallHierarchy capabilities specific to the various call hierarchy requests.
	//
	// @since 3.16.0.
	CallHierarchy *TextDocumentClientCapabilitiesCallHierarchy `json:"callHierarchy,omitempty"`

	// SemanticTokens capabilities specific to the various semantic token requests.
	//
	// @since 3.16.0.
	SemanticTokens *TextDocumentClientCapabilitiesSemanticTokens `json:"semanticTokens,omitempty"`

	// Moniker capabilities specific to the "textDocument/moniker" request.
	//
	// @since 3.16.0.
	Moniker *TextDocumentClientCapabilitiesMoniker `json:"moniker,omitempty"`
}

// ClientCapabilitiesShowMessageRequest show message request client capabilities.
//
// @since 3.16.0.
type ClientCapabilitiesShowMessageRequest struct {
	// MessageActionItem capabilities specific to the "MessageActionItem" type.
	MessageActionItem *ClientCapabilitiesShowMessageRequestMessageActionItem `json:"messageActionItem,omitempty"`
}

// ClientCapabilitiesShowMessageRequestMessageActionItem capabilities specific to the "MessageActionItem" type.
//
// @since 3.16.0.
type ClientCapabilitiesShowMessageRequestMessageActionItem struct {
	// AdditionalPropertiesSupport whether the client supports additional attributes which
	// are preserved and sent back to the server in the
	// request's response.
	AdditionalPropertiesSupport bool `json:"additionalPropertiesSupport,omitempty"`
}

// ClientCapabilitiesShowDocument client capabilities for the show document request.
//
// @since 3.16.0.
type ClientCapabilitiesShowDocument struct {
	// Support is the client has support for the show document
	// request.
	Support bool `json:"support"`
}

// WindowClientCapabilities represents a WindowClientCapabilities specific client capabilities.
//
// @since 3.15.0.
type WindowClientCapabilities struct {
	// WorkDoneProgress whether client supports handling progress notifications. If set servers are allowed to
	// report in "workDoneProgress" property in the request specific server capabilities.
	//
	// @since 3.15.0.
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`

	// ShowMessage capabilities specific to the showMessage request.
	//
	// @since 3.16.0.
	ShowMessage *ClientCapabilitiesShowMessageRequest `json:"showMessage,omitempty"`

	// ShowDocument client capabilities for the show document request.
	//
	// @since 3.16.0.
	ShowDocument *ClientCapabilitiesShowDocument `json:"showDocument,omitempty"`
}

// MarkdownClientCapabilities represents a client capabilities specific to the used markdown parser.
//
// @since 3.16.0.
type MarkdownClientCapabilities struct {
	// Parser is the name of the parser.
	Parser string `json:"parser"`

	// version is the version of the parser.
	Version string `json:"version,omitempty"`
}

// GeneralClientCapabilities represents a General specific client capabilities.
//
// @since 3.16.0.
type GeneralClientCapabilities struct {
	// RegularExpressions is the client capabilities specific to regular expressions.
	//
	// @since 3.16.0.
	RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions,omitempty"`

	// Markdown client capabilities specific to the client's markdown parser.
	//
	// @since 3.16.0.
	Markdown *MarkdownClientCapabilities `json:"markdown,omitempty"`
}

// ClientCapabilities now define capabilities for dynamic registration, workspace and text document features the client supports.
// The experimental can be used to pass experimental capabilities under development. For future compatibility a ClientCapabilities object literal can have more properties set than currently defined.
// Servers receiving a ClientCapabilities object literal with unknown properties should ignore these properties. A missing property should be interpreted as an absence of the capability.
// If a missing property normally defines sub properties, all missing sub properties should be interpreted as an absence of the corresponding capability.
type ClientCapabilities struct {
	// Workspace specific client capabilities.
	Workspace *WorkspaceClientCapabilities `json:"workspace,omitempty"`

	// TextDocument specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`

	// Window specific client capabilities.
	Window *WindowClientCapabilities `json:"window,omitempty"`

	// General client capabilities.
	//
	// @since 3.16.0.
	General *GeneralClientCapabilities `json:"general,omitempty"`

	// Experimental client capabilities.
	Experimental interface{} `json:"experimental,omitempty"`
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

// ServerCapabilities server capabilities.
type ServerCapabilities struct {
	// TextDocumentSync defines how text documents are synced. Is either a detailed structure defining each notification or
	// for backwards compatibility the TextDocumentSyncKind number. If omitted it defaults to "TextDocumentSyncKind.None".
	TextDocumentSync interface{} `json:"textDocumentSync,omitempty"` // *TextDocumentSyncOptions | TextDocumentSyncKind

	// HoverProvider is the server provides hover support.
	HoverProvider interface{} `json:"hoverProvider,omitempty"` // TODO(zchee): bool | *HoverOptions

	// CompletionProvider is the server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	// SignatureHelpProvider is the server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	// DeclarationProvider is the server provides go to declaration support.
	//
	// @since 3.14.0.
	DeclarationProvider interface{} `json:"declarationProvider,omitempty"` // TODO(zchee): bool | *DeclarationOptions | *DeclarationRegistrationOptions

	// DefinitionProvider is the server provides goto definition support.
	DefinitionProvider interface{} `json:"definitionProvider,omitempty"` // TODO(zchee): bool | *DefinitionOptions

	// TypeDefinitionProvider is the server provides Goto Type Definition support.
	//
	// @since 3.6.0.
	TypeDefinitionProvider interface{} `json:"typeDefinitionProvider,omitempty"` // TODO(zchee): bool | *TypeDefinitionOptions | *TypeDefinitionRegistrationOptions

	// ImplementationProvider is the server provides Goto Implementation support.
	//
	// @since 3.6.0.
	ImplementationProvider interface{} `json:"implementationProvider,omitempty"` // TODO(zchee): bool | *ImplementationOptions | *ImplementationRegistrationOptions

	// ReferencesProvider is the server provides find references support.
	ReferencesProvider interface{} `json:"referencesProvider,omitempty"` // TODO(zchee): bool | *ReferencesOptions

	// DocumentHighlightProvider is the server provides document highlight support.
	DocumentHighlightProvider interface{} `json:"documentHighlightProvider,omitempty"` // TODO(zchee): bool | *DocumentHighlightOptions

	// DocumentSymbolProvider is the server provides document symbol support.
	DocumentSymbolProvider interface{} `json:"documentSymbolProvider,omitempty"` // TODO(zchee): bool | *DocumentSymbolOptions

	// WorkspaceSymbolProvider is the server provides workspace symbol support.
	WorkspaceSymbolProvider interface{} `json:"workspaceSymbolProvider,omitempty"` // TODO(zchee): bool | *WorkspaceSymbolOptions

	// CodeActionProvider is the server provides code actions. The "CodeActionOptions" return type is only
	// valid if the client signals code action literal support via the property
	// "textDocument.codeAction.codeActionLiteralSupport".
	CodeActionProvider interface{} `json:"codeActionProvider,omitempty"` // TODO(zchee): bool | *CodeActionOptions

	// CodeLensProvider is the server provides code lens.
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider,omitempty"`

	// DocumentFormattingProvider is the server provides document formatting.
	DocumentFormattingProvider interface{} `json:"documentFormattingProvider,omitempty"` // TODO(zchee): bool | *DocumentFormattingOptions

	// DocumentRangeFormattingProvider is the server provides document range formatting.
	DocumentRangeFormattingProvider interface{} `json:"documentRangeFormattingProvider,omitempty"` // TODO(zchee): bool | *DocumentRangeFormattingOptions

	// DocumentOnTypeFormattingProvider is the server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`

	// RenameProvider is the server provides rename support. RenameOptions may only be
	// specified if the client states that it supports
	// "prepareSupport" in its initial "initialize" request.
	RenameProvider interface{} `json:"renameProvider,omitempty"` // TODO(zchee): bool | *RenameOptions

	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider,omitempty"`

	// ColorProvider is the server provides color provider support.
	//
	// @since 3.6.0.
	ColorProvider interface{} `json:"colorProvider,omitempty"` // TODO(zchee): bool | *DocumentColorOptions | *DocumentColorRegistrationOptions

	// FoldingRangeProvider is the server provides folding provider support.
	//
	// @since 3.10.0.
	FoldingRangeProvider interface{} `json:"foldingRangeProvider,omitempty"` // TODO(zchee): bool | *FoldingRangeOptions | *FoldingRangeRegistrationOptions

	// SelectionRangeProvider is the server provides selection range support.
	//
	// @since 3.15.0.
	SelectionRangeProvider interface{} `json:"selectionRangeProvider,omitempty"` // TODO(zchee): bool | *EnableSelectionRange | *SelectionRangeOptions | *SelectionRangeRegistrationOptions

	// ExecuteCommandProvider is the server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider,omitempty"`

	// Workspace specific server capabilities
	Workspace *ServerCapabilitiesWorkspace `json:"workspace,omitempty"`

	// LinkedEditingRangeProvider is the server provides linked editing range support.
	//
	// @since 3.16.0.
	LinkedEditingRangeProvider interface{} `json:"linkedEditingRangeProvider,omitempty"` // TODO(zchee): bool | *LinkedEditingRangeOptions | *LinkedEditingRangeRegistrationOptions

	// CallHierarchyProvider is the server provides call hierarchy support.
	//
	// @since 3.16.0.
	CallHierarchyProvider interface{} `json:"callHierarchyProvider,omitempty"` // TODO(zchee): bool | *CallHierarchyOptions | *CallHierarchyRegistrationOptions

	// SemanticTokensProvider is the server provides semantic tokens support.
	//
	// @since 3.16.0.
	SemanticTokensProvider interface{} `json:"semanticTokensProvider,omitempty"` // TODO(zchee): *SemanticTokensOptions | *SemanticTokensRegistrationOptions

	// MonikerProvider whether server provides moniker support.
	//
	// @since 3.16.0.
	MonikerProvider interface{} `json:"monikerProvider,omitempty"` // TODO(zchee): bool | *MonikerOptions | *MonikerRegistrationOptions

	// Experimental server capabilities.
	Experimental interface{} `json:"experimental,omitempty"`
}
