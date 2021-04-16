// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strconv"
)

// TraceValue represents a InitializeParams Trace mode.
type TraceValue string

// list of TraceValue.
const (
	// TraceOff disable tracing.
	TraceOff TraceValue = "off"

	// TraceMessage normal tracing mode.
	TraceMessage TraceValue = "message"

	// TraceVerbose verbose tracing mode.
	TraceVerbose TraceValue = "verbose"
)

// ClientInfo information about the client.
//
// @since 3.15.0.
type ClientInfo struct {
	// Name is the name of the client as defined by the client.
	Name string `json:"name"`

	// Version is the client's version as defined by the client.
	Version string `json:"version,omitempty"`
}

// InitializeParams params of Initialize Request.
type InitializeParams struct {
	WorkDoneProgressParams

	// ProcessID is the process Id of the parent process that started
	// the server. Is null if the process has not been started by another process.
	// If the parent process is not alive then the server should exit (see exit notification) its process.
	ProcessID int32 `json:"processId"`

	// ClientInfo is the information about the client.
	//
	// @since 3.15.0
	ClientInfo *ClientInfo `json:"clientInfo,omitempty"`

	// Locale is the locale the client is currently showing the user interface
	// in. This must not necessarily be the locale of the operating
	// system.
	//
	// Uses IETF language tags as the value's syntax
	// (See https://en.wikipedia.org/wiki/IETF_language_tag)
	//
	// @since 3.16.0.
	Locale string `json:"locale,omitempty"`

	// RootPath is the rootPath of the workspace. Is null
	// if no folder is open.
	//
	// Deprecated: Use RootURI instead.
	RootPath string `json:"rootPath,omitempty"`

	// RootURI is the rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and "rootUri" are set
	// "rootUri" wins.
	//
	// Deprecated: Use WorkspaceFolders instead.
	RootURI DocumentURI `json:"rootUri,omitempty"`

	// InitializationOptions user provided initialization options.
	InitializationOptions interface{} `json:"initializationOptions,omitempty"`

	// Capabilities is the capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`

	// Trace is the initial trace setting. If omitted trace is disabled ('off').
	Trace TraceValue `json:"trace,omitempty"`

	// WorkspaceFolders is the workspace folders configured in the client when the server starts.
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	//
	// @since 3.6.0.
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

// LogTraceParams params of LogTrace notification.
//
// @since 3.16.0.
type LogTraceParams struct {
	// Message is the message to be logged.
	Message string `json:"message"`

	// Verbose is the additional information that can be computed if the "trace" configuration
	// is set to "verbose".
	Verbose TraceValue `json:"verbose,omitempty"`
}

// SetTraceParams params of SetTrace notification.
//
// @since 3.16.0.
type SetTraceParams struct {
	// Value is the new value that should be assigned to the trace setting.
	Value TraceValue `json:"value"`
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

// FileOperationRegistrationOptions is the options to register for file operations.
//
// @since 3.16.0.
type FileOperationRegistrationOptions struct {
	// filters is the actual filters.
	Filters []FileOperationFilter `json:"filters"`
}

// FileOperationPatternKind is a pattern kind describing if a glob pattern matches a file a folder or
// both.
//
// @since 3.16.0.
type FileOperationPatternKind string

// list of FileOperationPatternKind.
const (
	// FileOperationPatternKindFile is the pattern matches a file only.
	FileOperationPatternKindFile FileOperationPatternKind = "file"

	// FileOperationPatternKindFolder is the pattern matches a folder only.
	FileOperationPatternKindFolder FileOperationPatternKind = "folder"
)

// FileOperationPatternOptions matching options for the file operation pattern.
//
// @since 3.16.0.
type FileOperationPatternOptions struct {
	// IgnoreCase is The pattern should be matched ignoring casing.
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}

// FileOperationPattern a pattern to describe in which file operation requests or notifications
// the server is interested in.
//
// @since 3.16.0.
type FileOperationPattern struct {
	// The glob pattern to match. Glob patterns can have the following syntax:
	//  - `*` to match one or more characters in a path segment
	//  - `?` to match on one character in a path segment
	//  - `**` to match any number of path segments, including none
	//  - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript
	//    and JavaScript files)
	//  - `[]` to declare a range of characters to match in a path segment
	//    (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	//  - `[!...]` to negate a range of characters to match in a path segment
	//    (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but
	//    not `example.0`)
	Glob string `json:"glob"`

	// Matches whether to match files or folders with this pattern.
	//
	// Matches both if undefined.
	Matches FileOperationPatternKind `json:"matches,omitempty"`

	// Options additional options used during matching.
	Options FileOperationPatternOptions `json:"options,omitempty"`
}

// FileOperationFilter is a filter to describe in which file operation requests or notifications
// the server is interested in.
//
// @since 3.16.0.
type FileOperationFilter struct {
	// Scheme is a URI like "file" or "untitled".
	Scheme string `json:"scheme,omitempty"`

	// Pattern is the actual file operation pattern.
	Pattern FileOperationPattern `json:"pattern"`
}

// CreateFilesParams is the parameters sent in notifications/requests for user-initiated creation
// of files.
//
// @since 3.16.0.
type CreateFilesParams struct {
	// Files an array of all files/folders created in this operation.
	Files []FileCreate `json:"files"`
}

// FileCreate nepresents information on a file/folder create.
//
// @since 3.16.0.
type FileCreate struct {
	// URI is a file:// URI for the location of the file/folder being created.
	URI string `json:"uri"`
}

// RenameFilesParams is the parameters sent in notifications/requests for user-initiated renames
// of files.
//
// @since 3.16.0.
type RenameFilesParams struct {
	// Files an array of all files/folders renamed in this operation. When a folder
	// is renamed, only the folder will be included, and not its children.
	Files []FileRename `json:"files"`
}

// FileRename represents information on a file/folder rename.
//
// @since 3.16.0.
type FileRename struct {
	// OldURI is a file:// URI for the original location of the file/folder being renamed.
	OldURI string `json:"oldUri"`

	// NewURI is a file:// URI for the new location of the file/folder being renamed.
	NewURI string `json:"newUri"`
}

// DeleteFilesParams is the parameters sent in notifications/requests for user-initiated deletes
// of files.
//
// @since 3.16.0.
type DeleteFilesParams struct {
	// Files an array of all files/folders deleted in this operation.
	Files []FileDelete `json:"files"`
}

// FileDelete represents information on a file/folder delete.
//
// @since 3.16.0.
type FileDelete struct {
	// URI is a file:// URI for the location of the file/folder being deleted.
	URI string `json:"uri"`
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

// ReferencesParams params of References Request.
//
// @since 3.15.0.
type ReferencesParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	// Context is the ReferenceParams context.
	Context ReferenceContext `json:"context"`
}

// TextDocumentClientCapabilitiesReferences capabilities specific to the "textDocument/references".
type TextDocumentClientCapabilitiesReferences struct {
	// DynamicRegistration whether references supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentHighlightOptions registration option of DocumentHighlight server capability.
//
// @since 3.15.0.
type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

// DocumentHighlightParams params of DocumentHighlight Request.
//
// @since 3.15.0.
type DocumentHighlightParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// TextDocumentClientCapabilitiesDocumentHighlight capabilities specific to the "textDocument/documentHighlight".
type TextDocumentClientCapabilitiesDocumentHighlight struct {
	// DynamicRegistration Whether document highlight supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentSymbolOptions registration option of DocumentSymbol server capability.
//
// @since 3.15.0.
type DocumentSymbolOptions struct {
	WorkDoneProgressOptions

	// Label a human-readable string that is shown when multiple outlines trees
	// are shown for the same document.
	//
	// @since 3.16.0.
	Label string `json:"label,omitempty"`
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

// WorkspaceSymbolOptions registration option of WorkspaceSymbol server capability.
//
// @since 3.15.0.
type WorkspaceSymbolOptions struct {
	WorkDoneProgressOptions
}

// DocumentFormattingOptions registration option of DocumentFormatting server capability.
//
// @since 3.15.0.
type DocumentFormattingOptions struct {
	WorkDoneProgressOptions
}

// TextDocumentClientCapabilitiesFormatting capabilities specific to the textDocument/formatting.
type TextDocumentClientCapabilitiesFormatting struct {
	// DynamicRegistration whether formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentRangeFormattingOptions registration option of DocumentRangeFormatting server capability.
//
// @since 3.15.0.
type DocumentRangeFormattingOptions struct {
	WorkDoneProgressOptions
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

// DeclarationOptions registration option of Declaration server capability.
//
// @since 3.15.0.
type DeclarationOptions struct {
	WorkDoneProgressOptions
}

// DeclarationRegistrationOptions registration option of Declaration server capability.
//
// @since 3.15.0.
type DeclarationRegistrationOptions struct {
	DeclarationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// DeclarationParams params of Declaration Request.
//
// @since 3.15.0.
type DeclarationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
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

// DefinitionOptions registration option of Definition server capability.
//
// @since 3.15.0.
type DefinitionOptions struct {
	WorkDoneProgressOptions
}

// DefinitionParams params of Definition Request.
//
// @since 3.15.0.
type DefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
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

// TypeDefinitionOptions registration option of TypeDefinition server capability.
//
// @since 3.15.0.
type TypeDefinitionOptions struct {
	WorkDoneProgressOptions
}

// TypeDefinitionRegistrationOptions registration option of TypeDefinition server capability.
//
// @since 3.15.0.
type TypeDefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	TypeDefinitionOptions
	StaticRegistrationOptions
}

// TypeDefinitionParams params of TypeDefinition Request.
//
// @since 3.15.0.
type TypeDefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
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

// ImplementationOptions registration option of Implementation server capability.
//
// @since 3.15.0.
type ImplementationOptions struct {
	WorkDoneProgressOptions
}

// ImplementationRegistrationOptions registration option of Implementation server capability.
//
// @since 3.15.0.
type ImplementationRegistrationOptions struct {
	TextDocumentRegistrationOptions
	ImplementationOptions
	StaticRegistrationOptions
}

// ImplementationParams params of Implementation Request.
//
// @since 3.15.0.
type ImplementationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
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

// DocumentColorOptions registration option of DocumentColor server capability.
//
// @since 3.15.0.
type DocumentColorOptions struct {
	WorkDoneProgressOptions
}

// DocumentColorRegistrationOptions registration option of DocumentColor server capability.
//
// @since 3.15.0.
type DocumentColorRegistrationOptions struct {
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
	DocumentColorOptions
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

// PrepareSupportDefaultBehavior default behavior of PrepareSupport.
//
// @since 3.16.0.
type PrepareSupportDefaultBehavior float64

// list of PrepareSupportDefaultBehavior.
const (
	// PrepareSupportDefaultBehaviorIdentifier is the client's default behavior is to select the identifier
	// according the to language's syntax rule.
	PrepareSupportDefaultBehaviorIdentifier PrepareSupportDefaultBehavior = 1
)

// String returns a string representation of the PrepareSupportDefaultBehavior.
func (k PrepareSupportDefaultBehavior) String() string {
	switch k {
	case PrepareSupportDefaultBehaviorIdentifier:
		return "Identifier"
	default:
		return strconv.FormatFloat(float64(k), 'f', -10, 64)
	}
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

// FoldingRangeOptions registration option of FoldingRange server capability.
//
// @since 3.15.0.
type FoldingRangeOptions struct {
	WorkDoneProgressOptions
}

// FoldingRangeRegistrationOptions registration option of FoldingRange server capability.
//
// @since 3.15.0.
type FoldingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	FoldingRangeOptions
	StaticRegistrationOptions
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

// ShowDocumentParams params to show a document.
//
// @since 3.16.0.
type ShowDocumentParams struct {
	// URI is the document uri to show.
	URI URI `json:"uri"`

	// External indicates to show the resource in an external program.
	// To show for example `https://code.visualstudio.com/`
	// in the default WEB browser set `external` to `true`.
	External bool `json:"external,omitempty"`

	// TakeFocus an optional property to indicate whether the editor
	// showing the document should take focus or not.
	// Clients might ignore this property if an external
	// program is started.
	TakeFocus bool `json:"takeFocus,omitempty"`

	// Selection an optional selection range if the document is a text
	// document. Clients might ignore the property if an
	// external program is started or the file is not a text
	// file.
	Selection *Range `json:"selection,omitempty"`
}

// ShowDocumentResult is the result of an show document request.
//
// @since 3.16.0.
type ShowDocumentResult struct {
	// Success a boolean indicating if the show was successful.
	Success bool `json:"success"`
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

// ServerInfo Information about the server.
//
// @since 3.15.0.
type ServerInfo struct {
	// Name is the name of the server as defined by the server.
	Name string `json:"name"`

	// Version is the server's version as defined by the server.
	Version string `json:"version,omitempty"`
}

// InitializeError known error codes for an "InitializeError".
type InitializeError struct {
	// Retry indicates whether the client execute the following retry logic:
	// (1) show the message provided by the ResponseError to the user
	// (2) user selects retry or cancel
	// (3) if user selected retry the initialize method is sent again.
	Retry bool `json:"retry,omitempty"`
}

// TextDocumentSyncKind defines how the host (editor) should sync document changes to the language server.
type TextDocumentSyncKind float64

const (
	// TextDocumentSyncKindNone documents should not be synced at all.
	TextDocumentSyncKindNone TextDocumentSyncKind = 0

	// TextDocumentSyncKindFull documents are synced by always sending the full content
	// of the document.
	TextDocumentSyncKindFull TextDocumentSyncKind = 1

	// TextDocumentSyncKindIncremental documents are synced by sending the full content on open.
	// After that only incremental updates to the document are
	// send.
	TextDocumentSyncKindIncremental TextDocumentSyncKind = 2
)

// String implements fmt.Stringer.
func (k TextDocumentSyncKind) String() string {
	switch k {
	case TextDocumentSyncKindNone:
		return "None"
	case TextDocumentSyncKindFull:
		return "Full"
	case TextDocumentSyncKindIncremental:
		return "Incremental"
	default:
		return strconv.FormatFloat(float64(k), 'f', -10, 64)
	}
}

// CompletionOptions Completion options.
type CompletionOptions struct {
	// The server provides support to resolve additional
	// information for a completion item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	// The characters that trigger completion automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
}

// SignatureHelpOptions SignatureHelp options.
type SignatureHelpOptions struct {
	// The characters that trigger signature help
	// automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// RetriggerCharacters is the slist of characters that re-trigger signature help.
	//
	// These trigger characters are only active when signature help is already
	// showing.
	// All trigger characters are also counted as re-trigger characters.
	//
	// @since 3.15.0.
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

// ReferencesOptions ReferencesProvider options.
//
// @since 3.15.0.
type ReferencesOptions struct {
	WorkDoneProgressOptions
}

// WorkDoneProgressOptions WorkDoneProgress options.
//
// @since 3.15.0.
type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

// CodeActionOptions CodeAction options.
type CodeActionOptions struct {
	// CodeActionKinds that this server may return.
	//
	// The list of kinds may be generic, such as "CodeActionKind.Refactor", or the server
	// may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`

	// ResolveProvider is the server provides support to resolve additional
	// information for a code action.
	//
	// @since 3.16.0.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// CodeLensOptions CodeLens options.
type CodeLensOptions struct {
	// Code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DocumentOnTypeFormattingOptions format document on type options.
type DocumentOnTypeFormattingOptions struct {
	// FirstTriggerCharacter a character on which formatting should be triggered, like "}".
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

	// MoreTriggerCharacter more trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
}

// RenameOptions rename options.
type RenameOptions struct {
	// PrepareProvider renames should be checked and tested before being executed.
	PrepareProvider bool `json:"prepareProvider,omitempty"`
}

// DocumentLinkOptions document link options.
type DocumentLinkOptions struct {
	// ResolveProvider document links have a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// ExecuteCommandOptions execute command options.
type ExecuteCommandOptions struct {
	// Commands is the commands to be executed on the server
	Commands []string `json:"commands"`
}

// SaveOptions save options.
type SaveOptions struct {
	// IncludeText is the client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
}

// TextDocumentSyncOptions TextDocumentSync options.
type TextDocumentSyncOptions struct {
	// OpenClose open and close notifications are sent to the server.
	OpenClose bool `json:"openClose,omitempty"`

	// Change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full
	// and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
	Change TextDocumentSyncKind `json:"change,omitempty"`

	// WillSave notifications are sent to the server.
	WillSave bool `json:"willSave,omitempty"`

	// WillSaveWaitUntil will save wait until requests are sent to the server.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

	// Save notifications are sent to the server.
	Save *SaveOptions `json:"save,omitempty"`
}

// HoverOptions option of hover provider server capabilities.
type HoverOptions struct {
	WorkDoneProgressOptions
}

// SemanticTokensOptions option of semantic tokens provider server capabilities.
//
// @since 3.16.0.
type SemanticTokensOptions struct {
	WorkDoneProgressOptions
}

// SemanticTokensRegistrationOptions registration option of semantic tokens provider server capabilities.
//
// @since 3.16.0.
type SemanticTokensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SemanticTokensOptions
	StaticRegistrationOptions
}

// LinkedEditingRangeOptions option of linked editing range provider server capabilities.
//
// @since 3.16.0.
type LinkedEditingRangeOptions struct {
	WorkDoneProgressOptions
}

// LinkedEditingRangeRegistrationOptions registration option of linked editing range provider server capabilities.
//
// @since 3.16.0.
type LinkedEditingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	LinkedEditingRangeOptions
	StaticRegistrationOptions
}

// LinkedEditingRangeParams params for the LinkedEditingRange request.
//
// @since 3.16.0.
type LinkedEditingRangeParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// LinkedEditingRanges result of LinkedEditingRange request.
//
// @since 3.16.0.
type LinkedEditingRanges struct {
	// Ranges a list of ranges that can be renamed together.
	//
	// The ranges must have identical length and contain identical text content.
	//
	// The ranges cannot overlap.
	Ranges []Range `json:"ranges"`

	// WordPattern an optional word pattern (regular expression) that describes valid contents for
	// the given ranges.
	//
	// If no pattern is provided, the client configuration's word pattern will be used.
	WordPattern string `json:"wordPattern,omitempty"`
}

// MonikerOptions option of moniker provider server capabilities.
//
// @since 3.16.0.
type MonikerOptions struct {
	WorkDoneProgressOptions
}

// MonikerRegistrationOptions registration option of moniker provider server capabilities.
//
// @since 3.16.0.
type MonikerRegistrationOptions struct {
	TextDocumentRegistrationOptions
	MonikerOptions
}

// MonikerParams params for the Moniker request.
//
// @since 3.16.0.
type MonikerParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// UniquenessLevel is the Moniker uniqueness level to define scope of the moniker.
//
// @since 3.16.0.
type UniquenessLevel string

// list of UniquenessLevel.
const (
	// UniquenessLevelDocument is the moniker is only unique inside a document.
	UniquenessLevelDocument UniquenessLevel = "document"

	// UniquenessLevelProject is the moniker is unique inside a project for which a dump got created.
	UniquenessLevelProject UniquenessLevel = "project"

	// UniquenessLevelGroup is the moniker is unique inside the group to which a project belongs.
	UniquenessLevelGroup UniquenessLevel = "group"

	// UniquenessLevelScheme is the moniker is unique inside the moniker scheme.
	UniquenessLevelScheme UniquenessLevel = "scheme"

	// UniquenessLevelGlobal is the moniker is globally unique.
	UniquenessLevelGlobal UniquenessLevel = "global"
)

// MonikerKind is the moniker kind.
//
// @since 3.16.0.
type MonikerKind string

// list of MonikerKind.
const (
	// MonikerKindImport is the moniker represent a symbol that is imported into a project.
	MonikerKindImport MonikerKind = "import"

	// MonikerKindExport is the moniker represents a symbol that is exported from a project.
	MonikerKindExport MonikerKind = "export"

	// MonikerKindLocal is the moniker represents a symbol that is local to a project (e.g. a local
	// variable of a function, a class not visible outside the project, ...).
	MonikerKindLocal MonikerKind = "local"
)

// Moniker definition to match LSIF 0.5 moniker definition.
//
// @since 3.16.0.
type Moniker struct {
	// Scheme is the scheme of the moniker. For example tsc or .Net.
	Scheme string `json:"scheme"`

	// Identifier is the identifier of the moniker.
	//
	// The value is opaque in LSIF however schema owners are allowed to define the structure if they want.
	Identifier string `json:"identifier"`

	// Unique is the scope in which the moniker is unique.
	Unique UniquenessLevel `json:"unique"`

	// Kind is the moniker kind if known.
	Kind MonikerKind `json:"kind,omitempty"`
}

// StaticRegistrationOptions staticRegistration options to be returned in the initialize request.
type StaticRegistrationOptions struct {
	// ID is the id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	ID string `json:"id,omitempty"`
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

// DocumentLinkRegistrationOptions DocumentLinkRegistration options.
type DocumentLinkRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// ResolveProvider document links have a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// InitializedParams params of Initialized Notification.
type InitializedParams struct{}

// FailureHandlingKind is the kind of failure handling .
type FailureHandlingKind string

const (
	// Abort applying the workspace change is simply aborted if one of the changes provided
	// fails. All operations executed before the failing operation stay executed.
	Abort FailureHandlingKind = "abort"

	// Transactional all operations are executed transactional. That means they either all
	// succeed or no changes at all are applied to the workspace.
	Transactional FailureHandlingKind = "transactional"

	// TextOnlyTransactional if the workspace edit contains only textual file changes they are executed transactional.
	// If resource changes (create, rename or delete file) are part of the change the failure
	// handling strategy is abort.
	TextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"

	// Undo the client tries to undo the operations already executed. But there is no
	// guarantee that this is succeeding.
	Undo FailureHandlingKind = "undo"
)

// WorkspaceFolders represents a slice of WorkspaceFolder.
type WorkspaceFolders []WorkspaceFolder
