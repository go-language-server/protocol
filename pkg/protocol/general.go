// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/go-language-server/jsonrpc2"
)

// CancelParams params of cancelRequest.
type CancelParams struct {

	// ID is the request id to cancel.
	ID jsonrpc2.ID `json:"id"`
}

// InitializeParams params of Initialize Request.
type InitializeParams struct {

	// ProcessID is the process Id of the parent process that started
	// the server. Is null if the process has not been started by another process.
	// If the parent process is not alive then the server should exit (see exit notification) its process.
	ProcessID *float64 `json:"processId"`

	// RootPath is the rootPath of the workspace. Is null
	// if no folder is open.
	RootPath *string `json:"rootPath,omitempty"`

	// RootURI is the rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins.
	RootURI *DocumentURI `json:"rootUri"`

	// InitializationOptions user provided initialization options.
	InitializationOptions interface{} `json:"initializationOptions,omitempty"`

	// Capabilities is the capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`

	// Trace is the initial trace setting. If omitted trace is disabled ('off').
	Trace string `json:"trace,omitempty"`

	// WorkspaceFolders is the workspace folders configured in the client when the server starts.
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	//
	// Since 3.6.0
	WorkspaceFolders []*WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

// WorkspaceClientCapabilities Workspace specific client capabilities.
type WorkspaceClientCapabilities struct {

	// The client supports applying batch edits to the workspace by supporting
	// the request 'workspace/applyEdit'
	ApplyEdit bool `json:"applyEdit,omitempty"`

	// WorkspaceEdit capabilities specific to `WorkspaceEdit`s
	WorkspaceEdit *struct {
		// DocumentChanges is the client supports versioned document changes in `WorkspaceEdit`s
		DocumentChanges bool `json:"documentChanges,omitempty"`

		// FailureHandling is the failure handling strategy of a client if applying the workspace edit
		// fails.
		FailureHandling string `json:"failureHandling,omitempty"`

		// ResourceOperations is the resource operations the client supports. Clients should at least
		// support 'create', 'rename' and 'delete' files and folders.
		ResourceOperations []string `json:"resourceOperations,omitempty"`
	} `json:"workspaceEdit,omitempty"`

	// DidChangeConfiguration capabilities specific to the `workspace/didChangeConfiguration` notification.
	DidChangeConfiguration *struct {
		// Did change configuration notification supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"didChangeConfiguration,omitempty"`

	// DidChangeWatchedFiles capabilities specific to the `workspace/didChangeWatchedFiles` notification.
	DidChangeWatchedFiles *struct {
		// Did change watched files notification supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"didChangeWatchedFiles,omitempty"`

	// Symbol capabilities specific to the `workspace/symbol` request.
	Symbol *struct {
		// Symbol request supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		// Specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
		SymbolKind *struct {
			/**
			 * ValueSet is the symbol kind values the client supports. When this
			 * property exists the client also guarantees that it will
			 * handle values outside its set gracefully and falls back
			 * to a default value when unknown.
			 *
			 * If this property is not present the client only supports
			 * the symbol kinds from `File` to `Array` as defined in
			 * the initial version of the protocol.
			 */
			ValueSet []*SymbolKind `json:"valueSet,omitempty"`
		} `json:"symbolKind,omitempty"`
	} `json:"symbol,omitempty"`

	// ExecuteCommand capabilities specific to the `workspace/executeCommand` request.
	ExecuteCommand *struct {
		// DynamicRegistration Execute command supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"executeCommand,omitempty"`

	// WorkspaceFolders is the client has support for workspace folders.
	//
	// Since 3.6.0
	WorkspaceFolders bool `json:"workspaceFolders,omitempty"`

	// Configuration is the client supports `workspace/configuration` requests.
	//
	// Since 3.6.0
	Configuration bool `json:"configuration,omitempty"`
}

// TextDocumentClientCapabilities Text document specific client capabilities.
type TextDocumentClientCapabilities struct {
	Synchronization *struct {

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
	} `json:"synchronization,omitempty"`

	// Completion Capabilities specific to the `textDocument/completion`
	Completion *struct {

		// Whether completion supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		// The client supports the following `CompletionItem` specific
		// capabilities.
		CompletionItem *struct {
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
		} `json:"completionItem,omitempty"`

		CompletionItemKind *CompletionItemKind `json:"completionItemKind,omitempty"`

		// ContextSupport is the client supports to send additional context information for a
		// `textDocument/completion` request.
		ContextSupport bool `json:"contextSupport,omitempty"`
	} `json:"completion,omitempty"`

	// Hover capabilities specific to the `textDocument/hover`
	Hover *struct {
		/**
		 * Whether hover supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		/**
		 * Client supports the follow content formats for the content
		 * property. The order describes the preferred format of the client.
		 */
		ContentFormat []MarkupKind `json:"contentFormat,omitempty"`
	} `json:"hover,omitempty"`

	// SignatureHelp capabilities specific to the `textDocument/signatureHelp`
	SignatureHelp *struct {
		/**
		 * DynamicRegistration whether signature help supports dynamic registration.
		 */
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		/**
		 * The client supports the following `SignatureInformation`
		 * specific properties.
		 */
		SignatureInformation struct {
			/**
			 * Client supports the follow content formats for the documentation
			 * property. The order describes the preferred format of the client.
			 */
			DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`
		} `json:"signatureInformation,omitempty"`
	} `json:"signatureHelp,omitempty"`

	// References capabilities specific to the `textDocument/references`
	References *struct {

		// DynamicRegistration whether references supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"references,omitempty"`

	// DocumentHighlight capabilities specific to the `textDocument/documentHighlight`
	DocumentHighlight *struct {
		// DynamicRegistration Whether document highlight supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"documentHighlight,omitempty"`

	// DocumentSymbol capabilities specific to the `textDocument/documentSymbol`
	DocumentSymbol *struct {
		// DynamicRegistration whether document symbol supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		// SymbolKind specific capabilities for the `SymbolKind`.
		SymbolKind *struct {
			// ValueSet is the symbol kind values the client supports. When this
			// property exists the client also guarantees that it will
			// handle values outside its set gracefully and falls back
			// to a default value when unknown.
			//
			// If this property is not present the client only supports
			// the symbol kinds from `File` to `Array` as defined in
			// the initial version of the protocol.
			ValueSet []SymbolKind `json:"valueSet,omitempty"`
		} `json:"symbolKind,omitempty"`

		// HierarchicalDocumentSymbolSupport is the client support hierarchical document symbols.
		HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`
	} `json:"documentSymbol,omitempty"`

	// Formatting capabilities specific to the `textDocument/formatting`
	Formatting *struct {

		// DynamicRegistration whether formatting supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"formatting,omitempty"`

	// RangeFormatting capabilities specific to the `textDocument/rangeFormatting`
	RangeFormatting *struct {

		// DynamicRegistration whether range formatting supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"rangeFormatting,omitempty"`

	// OnTypeFormatting Capabilities specific to the `textDocument/onTypeFormatting`
	OnTypeFormatting *struct {

		// DynamicRegistration whether on type formatting supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"onTypeFormatting,omitempty"`

	// Declaration capabilities specific to the `textDocument/declaration`
	Declaration *struct {

		// DynamicRegistration whether declaration supports dynamic registration. If this is set to `true`
		// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
		// return value for the corresponding server capability as well.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		// LinkSupport is the client supports additional metadata in the form of declaration links.
		//
		// Since 3.14.0
		LinkSupport bool `json:"linkSupport,omitempty"`
	} `json:"declaration,omitempty"`

	// Definition capabilities specific to the `textDocument/definition`.
	//
	// Since 3.14.0
	Definition *struct {

		// DynamicRegistration whether definition supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		// LinkSupport is the client supports additional metadata in the form of definition links.
		LinkSupport bool `json:"linkSupport,omitempty"`
	} `json:"definition,omitempty"`

	// TypeDefinition capabilities specific to the `textDocument/typeDefinition`
	//
	// Since 3.6.0
	TypeDefinition *struct {

		// DynamicRegistration whether typeDefinition supports dynamic registration. If this is set to `true`
		// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
		// return value for the corresponding server capability as well.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		// LinkSupport is the client supports additional metadata in the form of definition links.
		//
		// Since 3.14.0
		LinkSupport bool `json:"linkSupport,omitempty"`
	} `json:"typeDefinition,omitempty"`

	// Implementation capabilities specific to the `textDocument/implementation`.
	//
	// Since 3.6.0
	Implementation *struct {

		// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true`
		// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
		// return value for the corresponding server capability as well.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		// LinkSupport is the client supports additional metadata in the form of definition links.
		//
		// Since 3.14.0
		LinkSupport bool `json:"linkSupport,omitempty"`
	} `json:"implementation,omitempty"`

	// CodeAction capabilities specific to the `textDocument/codeAction`
	CodeAction *struct {
		// DynamicRegistration whether code action supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
		// CodeActionLiteralSupport is the client support code action literals as a valid
		// response of the `textDocument/codeAction` request.
		//
		// Since 3.8.0
		CodeActionLiteralSupport *struct {
			// CodeActionKind is the code action kind is support with the following value
			// set.
			CodeActionKind struct {

				/**
				 * ValueSet is the code action kind values the client supports. When this
				 * property exists the client also guarantees that it will
				 * handle values outside its set gracefully and falls back
				 * to a default value when unknown.
				 */
				ValueSet []CodeActionKind `json:"valueSet"`
			} `json:"codeActionKind"`
		} `json:"codeActionLiteralSupport,omitempty"`
	} `json:"codeAction,omitempty"`

	// CodeLens capabilities specific to the `textDocument/codeLens`
	CodeLens *struct {
		// DynamicRegistration Whether code lens supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"codeLens,omitempty"`

	// DocumentLink capabilities specific to the `textDocument/documentLink`
	DocumentLink *struct {
		// DynamicRegistration whether document link supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"documentLink,omitempty"`

	// ColorProvider capabilities specific to the `textDocument/documentColor` and the
	// `textDocument/colorPresentation` request.
	//
	// Since 3.6.0
	ColorProvider *struct {
		// DynamicRegistration whether colorProvider supports dynamic registration. If this is set to `true`
		// the client supports the new `(ColorProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)`
		// return value for the corresponding server capability as well.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	} `json:"colorProvider,omitempty"`

	// Rename capabilities specific to the `textDocument/rename`
	Rename *struct {
		// DynamicRegistration whether rename supports dynamic registration.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		// PrepareSupport is the client supports testing for validity of rename operations
		// before execution.
		PrepareSupport bool `json:"prepareSupport,omitempty"`
	} `json:"rename,omitempty"`

	// PublishDiagnostics capabilities specific to `textDocument/publishDiagnostics`.
	PublishDiagnostics *struct {

		// RelatedInformation whether the clients accepts diagnostics with related information.
		RelatedInformation bool `json:"relatedInformation,omitempty"`
	} `json:"publishDiagnostics,omitempty"`

	// FoldingRange capabilities specific to `textDocument/foldingRange` requests.
	//
	// Since 3.10.0
	FoldingRange *struct {
		// DynamicRegistration whether implementation supports dynamic registration for folding range providers. If this is set to `true`
		// the client supports the new `(FoldingRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions)`
		// return value for the corresponding server capability as well.
		DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

		// RangeLimit is the maximum number of folding ranges that the client prefers to receive per document. The value serves as a
		// hint, servers are free to follow the limit.
		RangeLimit float64 `json:"rangeLimit,omitempty"`
		// LineFoldingOnly if set, the client signals that it only supports folding complete lines. If set, client will
		// ignore specified `startCharacter` and `endCharacter` properties in a FoldingRange.
		LineFoldingOnly bool `json:"lineFoldingOnly,omitempty"`
	} `json:"foldingRange,omitempty"`
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

	// Experimental client capabilities.
	Experimental interface{} `json:"experimental,omitempty"`
}

// InitializeResult result of ClientCapabilities.
type InitializeResult struct {

	// Capabilities is the capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`
}

// InitializeError known error codes for an `InitializeError`.
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
	// None documents should not be synced at all.
	None TextDocumentSyncKind = 0

	// Full documents are synced by always sending the full content
	// of the document.
	Full TextDocumentSyncKind = 1

	// Incremental documents are synced by sending the full content on open.
	// After that only incremental updates to the document are
	// send.
	Incremental TextDocumentSyncKind = 2
)

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
}

// CodeActionOptions CodeAction options.
type CodeActionOptions struct {

	// CodeActionKinds that this server may return.
	//
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server
	// may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`
}

// CodeLensOptions CodeLens options.
type CodeLensOptions struct {

	// Code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DocumentOnTypeFormattingOptions format document on type options.
type DocumentOnTypeFormattingOptions struct {

	// FirstTriggerCharacter a character on which formatting should be triggered, like `}`.
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

// ColorProviderOptions ColorProvider options.
type ColorProviderOptions struct{}

// FoldingRangeProviderOptions FoldingRangeProvider options.
type FoldingRangeProviderOptions struct{}

// TextDocumentSyncOptions TextDocumentSync options.
type TextDocumentSyncOptions struct {

	// OpenClose open and close notifications are sent to the server.
	OpenClose bool `json:"openClose,omitempty"`

	// Change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full
	// and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
	Change float64 `json:"change,omitempty"`

	// WillSave notifications are sent to the server.
	WillSave bool `json:"willSave,omitempty"`

	// WillSaveWaitUntil will save wait until requests are sent to the server.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

	// Save notifications are sent to the server.
	Save *SaveOptions `json:"save,omitempty"`
}

// StaticRegistrationOptions StaticRegistration options to be returned in the initialize request.
type StaticRegistrationOptions struct {

	// ID is the id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	ID string `json:"id,omitempty"`
}

// ServerCapabilities server capabilities.
type ServerCapabilities struct {

	// TextDocumentSync defines how text documents are synced. Is either a detailed structure defining each notification or
	// for backwards compatibility the TextDocumentSyncKind number. If omitted it defaults to `TextDocumentSyncKind.None`.
	TextDocumentSync interface{} `json:"textDocumentSync,omitempty"`

	// HoverProvider is the server provides hover support.
	HoverProvider bool `json:"hoverProvider,omitempty"`

	// CompletionProvider is the server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	// SignatureHelpProvider is the server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	// DefinitionProvider is the server provides goto definition support.
	DefinitionProvider bool `json:"definitionProvider,omitempty"`

	// TypeDefinitionProvider is the server provides Goto Type Definition support.
	//
	// Since 3.6.0
	TypeDefinitionProvider interface{} `json:"typeDefinitionProvider,omitempty"`

	// ImplementationProvider is the server provides Goto Implementation support.
	//
	// Since 3.6.0
	ImplementationProvider interface{} `json:"implementationProvider,omitempty"`

	// ReferencesProvider is the server provides find references support.
	ReferencesProvider bool `json:"referencesProvider,omitempty"`

	// DocumentHighlightProvider is the server provides document highlight support.
	DocumentHighlightProvider bool `json:"documentHighlightProvider,omitempty"`

	// DocumentSymbolProvider is the server provides document symbol support.
	DocumentSymbolProvider bool `json:"documentSymbolProvider,omitempty"`

	// WorkspaceSymbolProvider is the server provides workspace symbol support.
	WorkspaceSymbolProvider bool `json:"workspaceSymbolProvider,omitempty"`

	// CodeActionProvider is the server provides code actions. The `CodeActionOptions` return type is only
	// valid if the client signals code action literal support via the property
	// `textDocument.codeAction.codeActionLiteralSupport`.
	CodeActionProvider bool `json:"codeActionProvider,omitempty"`

	// CodeLensProvider is the server provides code lens.
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider,omitempty"`

	// DocumentFormattingProvider is the server provides document formatting.
	DocumentFormattingProvider bool `json:"documentFormattingProvider,omitempty"`

	// DocumentRangeFormattingProvider is the server provides document range formatting.
	DocumentRangeFormattingProvider bool `json:"documentRangeFormattingProvider,omitempty"`

	// DocumentOnTypeFormattingProvider is the server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`

	// RenameProvider is the server provides rename support. RenameOptions may only be
	// specified if the client states that it supports
	// `prepareSupport` in its initial `initialize` request.
	RenameProvider bool `json:"renameProvider,omitempty"`

	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider,omitempty"`

	// ColorProvider is the server provides color provider support.
	//
	// Since 3.6.0
	ColorProvider interface{} `json:"colorProvider,omitempty"`

	// FoldingRangeProvider is the server provides folding provider support.
	//
	// Since 3.10.0
	FoldingRangeProvider interface{} `json:"foldingRangeProvider,omitempty"`

	// ExecuteCommandProvider is the server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider,omitempty"`

	// Workspace specific server capabilities
	Workspace *struct {
		WorkspaceFolders struct {
			/**
			 * The server has support for workspace folders
			 */
			Supported bool `json:"supported,omitempty"`
			/**
			 * Whether the server wants to receive workspace folder
			 * change notifications.
			 *
			 * If a strings is provided the string is treated as a ID
			 * under which the notification is registered on the client
			 * side. The ID can be used to unregister for these events
			 * using the `client/unregisterCapability` request.
			 */
			ChangeNotifications interface{} `json:"changeNotifications,omitempty"` // string | boolean
		} `json:"workspaceFolders,omitempty"`
	} `json:"workspace,omitempty"`

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

// ResourceOperationKind is the file event type.
type ResourceOperationKind string

const (
	// CreateResourceOperation supports creating new files and folders.
	CreateResourceOperation ResourceOperationKind = "create"

	// RenameResourceOperation supports renaming existing files and folders.
	RenameResourceOperation ResourceOperationKind = "rename"

	// DeleteResourceOperation supports deleting existing files and folders.
	DeleteResourceOperation ResourceOperationKind = "delete"
)

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
