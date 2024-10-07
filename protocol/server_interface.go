// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"

	"go.lsp.dev/jsonrpc2"
)

const (
	MethodServerCancelRequest                 ServerMethod = "$/cancelRequest"                        // bidirect server notification
	MethodServerProgress                      ServerMethod = "$/progress"                             // bidirect server notification
	MethodWorkspaceDidChangeConfiguration     ServerMethod = "workspace/didChangeConfiguration"       // server notification
	MethodNotebookDocumentDidChange           ServerMethod = "notebookDocument/didChange"             // server notification
	MethodTextDocumentDidChange               ServerMethod = "textDocument/didChange"                 // server notification
	MethodWorkspaceDidChangeWatchedFiles      ServerMethod = "workspace/didChangeWatchedFiles"        // server notification
	MethodWorkspaceDidChangeWorkspaceFolders  ServerMethod = "workspace/didChangeWorkspaceFolders"    // server notification
	MethodNotebookDocumentDidClose            ServerMethod = "notebookDocument/didClose"              // server notification
	MethodTextDocumentDidClose                ServerMethod = "textDocument/didClose"                  // server notification
	MethodWorkspaceDidCreateFiles             ServerMethod = "workspace/didCreateFiles"               // server notification
	MethodWorkspaceDidDeleteFiles             ServerMethod = "workspace/didDeleteFiles"               // server notification
	MethodNotebookDocumentDidOpen             ServerMethod = "notebookDocument/didOpen"               // server notification
	MethodTextDocumentDidOpen                 ServerMethod = "textDocument/didOpen"                   // server notification
	MethodWorkspaceDidRenameFiles             ServerMethod = "workspace/didRenameFiles"               // server notification
	MethodNotebookDocumentDidSave             ServerMethod = "notebookDocument/didSave"               // server notification
	MethodTextDocumentDidSave                 ServerMethod = "textDocument/didSave"                   // server notification
	MethodExit                                ServerMethod = "exit"                                   // server notification
	MethodInitialized                         ServerMethod = "initialized"                            // server notification
	MethodSetTrace                            ServerMethod = "$/setTrace"                             // server notification
	MethodTextDocumentWillSave                ServerMethod = "textDocument/willSave"                  // server notification
	MethodWindowWorkDoneProgressCancel        ServerMethod = "window/workDoneProgress/cancel"         // server notification
	MethodCallHierarchyIncomingCalls          ServerMethod = "callHierarchy/incomingCalls"            // server request
	MethodCallHierarchyOutgoingCalls          ServerMethod = "callHierarchy/outgoingCalls"            // server request
	MethodTextDocumentPrepareCallHierarchy    ServerMethod = "textDocument/prepareCallHierarchy"      // server request
	MethodTextDocumentCodeAction              ServerMethod = "textDocument/codeAction"                // server request
	MethodCodeActionResolve                   ServerMethod = "codeAction/resolve"                     // server request
	MethodTextDocumentCodeLens                ServerMethod = "textDocument/codeLens"                  // server request
	MethodCodeLensResolve                     ServerMethod = "codeLens/resolve"                       // server request
	MethodTextDocumentColorPresentation       ServerMethod = "textDocument/colorPresentation"         // server request
	MethodTextDocumentCompletion              ServerMethod = "textDocument/completion"                // server request
	MethodCompletionItemResolve               ServerMethod = "completionItem/resolve"                 // server request
	MethodTextDocumentDeclaration             ServerMethod = "textDocument/declaration"               // server request
	MethodTextDocumentDefinition              ServerMethod = "textDocument/definition"                // server request
	MethodTextDocumentDocumentColor           ServerMethod = "textDocument/documentColor"             // server request
	MethodTextDocumentDiagnostic              ServerMethod = "textDocument/diagnostic"                // server request
	MethodTextDocumentFormatting              ServerMethod = "textDocument/formatting"                // server request
	MethodTextDocumentDocumentHighlight       ServerMethod = "textDocument/documentHighlight"         // server request
	MethodTextDocumentDocumentLink            ServerMethod = "textDocument/documentLink"              // server request
	MethodDocumentLinkResolve                 ServerMethod = "documentLink/resolve"                   // server request
	MethodTextDocumentOnTypeFormatting        ServerMethod = "textDocument/onTypeFormatting"          // server request
	MethodTextDocumentRangeFormatting         ServerMethod = "textDocument/rangeFormatting"           // server request
	MethodTextDocumentRangesFormatting        ServerMethod = "textDocument/rangesFormatting"          // server request
	MethodTextDocumentDocumentSymbol          ServerMethod = "textDocument/documentSymbol"            // server request
	MethodWorkspaceExecuteCommand             ServerMethod = "workspace/executeCommand"               // server request
	MethodTextDocumentFoldingRange            ServerMethod = "textDocument/foldingRange"              // server request
	MethodTextDocumentHover                   ServerMethod = "textDocument/hover"                     // server request
	MethodTextDocumentImplementation          ServerMethod = "textDocument/implementation"            // server request
	MethodInitialize                          ServerMethod = "initialize"                             // server request
	MethodTextDocumentInlayHint               ServerMethod = "textDocument/inlayHint"                 // server request
	MethodInlayHintResolve                    ServerMethod = "inlayHint/resolve"                      // server request
	MethodTextDocumentInlineCompletion        ServerMethod = "textDocument/inlineCompletion"          // server request
	MethodTextDocumentInlineValue             ServerMethod = "textDocument/inlineValue"               // server request
	MethodTextDocumentLinkedEditingRange      ServerMethod = "textDocument/linkedEditingRange"        // server request
	MethodTextDocumentMoniker                 ServerMethod = "textDocument/moniker"                   // server request
	MethodTextDocumentPrepareRename           ServerMethod = "textDocument/prepareRename"             // server request
	MethodTextDocumentReferences              ServerMethod = "textDocument/references"                // server request
	MethodTextDocumentRename                  ServerMethod = "textDocument/rename"                    // server request
	MethodTextDocumentSelectionRange          ServerMethod = "textDocument/selectionRange"            // server request
	MethodTextDocumentSemanticTokensFullDelta ServerMethod = "textDocument/semanticTokens/full/delta" // server request
	MethodTextDocumentSemanticTokensRange     ServerMethod = "textDocument/semanticTokens/range"      // server request
	MethodTextDocumentSemanticTokensFull      ServerMethod = "textDocument/semanticTokens/full"       // server request
	MethodShutdown                            ServerMethod = "shutdown"                               // server request
	MethodTextDocumentSignatureHelp           ServerMethod = "textDocument/signatureHelp"             // server request
	MethodWorkspaceTextDocumentContent        ServerMethod = "workspace/textDocumentContent"          // server request
	MethodTextDocumentTypeDefinition          ServerMethod = "textDocument/typeDefinition"            // server request
	MethodTextDocumentPrepareTypeHierarchy    ServerMethod = "textDocument/prepareTypeHierarchy"      // server request
	MethodTypeHierarchySubtypes               ServerMethod = "typeHierarchy/subtypes"                 // server request
	MethodTypeHierarchySupertypes             ServerMethod = "typeHierarchy/supertypes"               // server request
	MethodWorkspaceWillCreateFiles            ServerMethod = "workspace/willCreateFiles"              // server request
	MethodWorkspaceWillDeleteFiles            ServerMethod = "workspace/willDeleteFiles"              // server request
	MethodWorkspaceWillRenameFiles            ServerMethod = "workspace/willRenameFiles"              // server request
	MethodTextDocumentWillSaveWaitUntil       ServerMethod = "textDocument/willSaveWaitUntil"         // server request
	MethodWorkspaceDiagnostic                 ServerMethod = "workspace/diagnostic"                   // server request
	MethodWorkspaceSymbol                     ServerMethod = "workspace/symbol"                       // server request
	MethodWorkspaceSymbolResolve              ServerMethod = "workspaceSymbol/resolve"                // server request
)

type Server interface {
	Cancel(ctx context.Context, params *CancelParams) error

	Progress(ctx context.Context, params *ProgressParams) error

	// DidChangeConfiguration the configuration change notification is sent from the client to the server when the client's configuration has changed. The notification contains the changed configuration as defined by the language client.
	DidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) error

	DidChangeNotebookDocument(ctx context.Context, params *DidChangeNotebookDocumentParams) error

	// DidChangeTextDocument the document change notification is sent from the client to the server to signal changes to a text document.
	DidChangeTextDocument(ctx context.Context, params *DidChangeTextDocumentParams) error

	// DidChangeWatchedFiles the watched files notification is sent from the client to the server when the client detects changes
	// to file watched by the language client.
	DidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) error

	// DidChangeWorkspaceFolders the `workspace/didChangeWorkspaceFolders` notification is sent from the client to the server when the workspace folder configuration changes.
	DidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) error

	// DidCloseNotebookDocument a notification sent when a notebook closes.
	//
	// @since 3.17.0
	DidCloseNotebookDocument(ctx context.Context, params *DidCloseNotebookDocumentParams) error

	// DidCloseTextDocument the document close notification is sent from the client to the server when the document got closed in the client. The document's truth now exists where the document's uri points to (e.g. if the document's uri is a file uri the truth now exists on disk). As with the open notification the close notification is about managing the document's content. Receiving a close notification doesn't mean that the document was open in an editor before. A close notification requires a previous open notification to be sent.
	DidCloseTextDocument(ctx context.Context, params *DidCloseTextDocumentParams) error

	// DidCreateFiles the did create files notification is sent from the client to the server when files were created from
	// within the client.
	//
	// @since 3.16.0
	DidCreateFiles(ctx context.Context, params *CreateFilesParams) error

	// DidDeleteFiles the will delete files request is sent from the client to the server before files are actually deleted as long as the deletion is triggered from within the client.
	//
	// @since 3.16.0
	DidDeleteFiles(ctx context.Context, params *DeleteFilesParams) error

	// DidOpenNotebookDocument a notification sent when a notebook opens.
	//
	// @since 3.17.0
	DidOpenNotebookDocument(ctx context.Context, params *DidOpenNotebookDocumentParams) error

	// DidOpenTextDocument the document open notification is sent from the client to the server to signal newly opened text documents. The document's truth is now managed by the client and the server must not try to read the document's truth using the document's uri. Open in this sense means it is managed by the client. It doesn't necessarily mean that its content is presented in an editor. An open notification must not be sent more than once without a corresponding close notification send before. This means open and close notification must be balanced and the max open count is one.
	DidOpenTextDocument(ctx context.Context, params *DidOpenTextDocumentParams) error

	// DidRenameFiles the did rename files notification is sent from the client to the server when files were renamed from
	// within the client.
	//
	// @since 3.16.0
	DidRenameFiles(ctx context.Context, params *RenameFilesParams) error

	// DidSaveNotebookDocument a notification sent when a notebook document is saved.
	//
	// @since 3.17.0
	DidSaveNotebookDocument(ctx context.Context, params *DidSaveNotebookDocumentParams) error

	// DidSaveTextDocument the document save notification is sent from the client to the server when the document got saved in the client.
	DidSaveTextDocument(ctx context.Context, params *DidSaveTextDocumentParams) error

	// Exit the exit event is sent from the client to the server to ask the server to exit its process.
	Exit(ctx context.Context) error

	// Initialized the initialized notification is sent from the client to the server after the client is fully initialized and the server is allowed to send requests from the server to the client.
	Initialized(ctx context.Context, params *InitializedParams) error

	SetTrace(ctx context.Context, params *SetTraceParams) error

	// WillSaveTextDocument a document will save notification is sent from the client to the server before the document is actually saved.
	WillSaveTextDocument(ctx context.Context, params *WillSaveTextDocumentParams) error

	// WorkDoneProgressCancel the `window/workDoneProgress/cancel` notification is sent from the client to the server to cancel a progress initiated on the server side.
	WorkDoneProgressCancel(ctx context.Context, params *WorkDoneProgressCancelParams) error
	// CallHierarchyIncomingCalls a request to resolve the incoming calls for a given `CallHierarchyItem`.
	//
	// @since 3.16.0
	CallHierarchyIncomingCalls(ctx context.Context, params *CallHierarchyIncomingCallsParams) ([]*CallHierarchyIncomingCall, error)

	// CallHierarchyOutgoingCalls a request to resolve the outgoing calls for a given `CallHierarchyItem`.
	//
	// @since 3.16.0
	CallHierarchyOutgoingCalls(ctx context.Context, params *CallHierarchyOutgoingCallsParams) ([]*CallHierarchyOutgoingCall, error)

	// CallHierarchyPrepare a request to result a `CallHierarchyItem` in a document at a given position. Can be used as an input
	// to an incoming or outgoing call hierarchy.
	//
	// @since 3.16.0
	CallHierarchyPrepare(ctx context.Context, params *CallHierarchyPrepareParams) ([]*CallHierarchyItem, error)

	// CodeAction a request to provide commands for the given text document and range.
	CodeAction(ctx context.Context, params *CodeActionParams) (*CodeActionRequestResult, error)

	// CodeActionResolve request to resolve additional information for a given code action.The request's parameter is of type
	// CodeAction the response is of type CodeAction or a Thenable that resolves to such.
	CodeActionResolve(ctx context.Context, params *CodeAction) (*CodeAction, error)

	// CodeLens a request to provide code lens for the given text document.
	CodeLens(ctx context.Context, params *CodeLensParams) ([]*CodeLens, error)

	// CodeLensResolve a request to resolve a command for a given code lens.
	CodeLensResolve(ctx context.Context, params *CodeLens) (*CodeLens, error)

	// ColorPresentation a request to list all presentation for a color. The request's parameter is of type ColorPresentationParams the response is of type ColorInformation ColorInformation[] or a Thenable that resolves to such.
	ColorPresentation(ctx context.Context, params *ColorPresentationParams) ([]*ColorPresentation, error)

	// Completion request to request completion at a given text document position. The request's parameter is of type TextDocumentPosition the response is of type CompletionItem CompletionItem[] or CompletionList or a Thenable that resolves to such. The request can delay the computation of the CompletionItem.detail `detail` and CompletionItem.documentation `documentation` properties to the `completionItem/resolve` request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
	// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
	Completion(ctx context.Context, params *CompletionParams) (*CompletionResult, error)

	// CompletionResolve request to resolve additional information for a given completion item.The request's parameter is of type CompletionItem the response is of type CompletionItem or a Thenable that resolves to such.
	CompletionResolve(ctx context.Context, params *CompletionItem) (*CompletionItem, error)

	// Declaration a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Declaration or a
	// typed array of DeclarationLink or a Thenable that resolves to such.
	Declaration(ctx context.Context, params *DeclarationParams) (*DeclarationResult, error)

	// Definition a request to resolve the definition location of a symbol at a given text document position. The request's parameter is of type TextDocumentPosition the response is of either type Definition or a typed
	// array of DefinitionLink or a Thenable that resolves to such.
	Definition(ctx context.Context, params *DefinitionParams) (*DefinitionResult, error)

	// DocumentColor a request to list all color symbols found in a given text document. The request's parameter is of type DocumentColorParams the response is of type ColorInformation ColorInformation[] or a Thenable that resolves to such.
	DocumentColor(ctx context.Context, params *DocumentColorParams) ([]*ColorInformation, error)

	// DocumentDiagnostic the document diagnostic request definition.
	//
	// @since 3.17.0
	DocumentDiagnostic(ctx context.Context, params *DocumentDiagnosticParams) (*DocumentDiagnosticReport, error)

	// DocumentFormatting a request to format a whole document.
	DocumentFormatting(ctx context.Context, params *DocumentFormattingParams) ([]*TextEdit, error)

	// DocumentHighlight request to resolve a DocumentHighlight for a given text document position. The request's parameter is of type TextDocumentPosition the request response is an array of type DocumentHighlight or a Thenable that resolves to such.
	DocumentHighlight(ctx context.Context, params *DocumentHighlightParams) ([]*DocumentHighlight, error)

	// DocumentLink a request to provide document links.
	DocumentLink(ctx context.Context, params *DocumentLinkParams) ([]*DocumentLink, error)

	// DocumentLinkResolve request to resolve additional information for a given document link. The request's parameter is of type DocumentLink the response is of type DocumentLink or a Thenable that resolves to such.
	DocumentLinkResolve(ctx context.Context, params *DocumentLink) (*DocumentLink, error)

	// DocumentOnTypeFormatting a request to format a document on type.
	DocumentOnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) ([]*TextEdit, error)

	// DocumentRangeFormatting a request to format a range in a document.
	DocumentRangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) ([]*TextEdit, error)

	// DocumentRangesFormatting a request to format ranges in a document.  3.18.0 @proposed.
	//
	// @since 3.18.0 proposed
	DocumentRangesFormatting(ctx context.Context, params *DocumentRangesFormattingParams) ([]*TextEdit, error)

	// DocumentSymbol a request to list all symbols found in a given text document. The request's parameter is of type TextDocumentIdentifier the response is of type SymbolInformation SymbolInformation[] or a Thenable that
	// resolves to such.
	DocumentSymbol(ctx context.Context, params *DocumentSymbolParams) (*DocumentSymbolResult, error)

	// ExecuteCommand a request send from the client to the server to execute a command. The request might return a workspace edit which the client will apply to the workspace.
	ExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (any, error)

	// FoldingRange a request to provide folding ranges in a document. The request's parameter is of type FoldingRangeParams, the response is of type FoldingRangeList or a Thenable that resolves to such.
	FoldingRange(ctx context.Context, params *FoldingRangeParams) ([]*FoldingRange, error)

	// Hover request to request hover information at a given text document position. The request's parameter is of type TextDocumentPosition the response is of type Hover or a Thenable that resolves to such.
	Hover(ctx context.Context, params *HoverParams) (*Hover, error)

	// Implementation a request to resolve the implementation locations of a symbol at a given text document position. The
	// request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
	Implementation(ctx context.Context, params *ImplementationParams) (*ImplementationResult, error)

	// Initialize the initialize request is sent from the client to the server. It is sent once as the request after starting up the server. The requests parameter is of type InitializeParams the response if of type InitializeResult of a Thenable that resolves to such.
	Initialize(ctx context.Context, params *InitializeParams) (*InitializeResult, error)

	// InlayHint a request to provide inlay hints in a document. The request's parameter is of type InlayHintsParams,
	// the response is of type InlayHint InlayHint[] or a Thenable that resolves to such.
	//
	// @since 3.17.0
	InlayHint(ctx context.Context, params *InlayHintParams) ([]*InlayHint, error)

	// InlayHintResolve a request to resolve additional properties for an inlay hint. The request's parameter is of type InlayHint, the response is of type InlayHint or a Thenable that resolves to such.
	//
	// @since 3.17.0
	InlayHintResolve(ctx context.Context, params *InlayHint) (*InlayHint, error)

	// InlineCompletion a request to provide inline completions in a document. The request's parameter is of type InlineCompletionParams, the response is of type InlineCompletion InlineCompletion[] or a Thenable that resolves to such. 3.18.0 @proposed.
	//
	// @since 3.18.0 proposed
	InlineCompletion(ctx context.Context, params *InlineCompletionParams) (*InlineCompletionResult, error)

	// InlineValue a request to provide inline values in a document. The request's parameter is of type InlineValueParams, the response is of type InlineValue InlineValue[] or a Thenable that resolves to such.
	//
	// @since 3.17.0
	InlineValue(ctx context.Context, params *InlineValueParams) ([]*InlineValue, error)

	// LinkedEditingRange a request to provide ranges that can be edited together.
	//
	// @since 3.16.0
	LinkedEditingRange(ctx context.Context, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error)

	// Moniker a request to get the moniker of a symbol at a given text document position. The request parameter is
	// of type TextDocumentPositionParams. The response is of type Moniker Moniker[] or `null`.
	Moniker(ctx context.Context, params *MonikerParams) ([]*Moniker, error)

	// PrepareRename a request to test and perform the setup necessary for a rename. 3.16 - support for default behavior.
	//
	// @since 3.16 - support for default behavior
	PrepareRename(ctx context.Context, params *PrepareRenameParams) (*PrepareRenameResult, error)

	// References a request to resolve project-wide references for the symbol denoted by the given text document position. The request's parameter is of type ReferenceParams the response is of type Location Location[] or a Thenable that resolves to such.
	References(ctx context.Context, params *ReferenceParams) ([]*Location, error)

	// Rename a request to rename a symbol.
	Rename(ctx context.Context, params *RenameParams) (*WorkspaceEdit, error)

	// SelectionRange a request to provide selection ranges in a document. The request's parameter is of type SelectionRangeParams, the response is of type SelectionRange SelectionRange[] or a Thenable that resolves to such.
	SelectionRange(ctx context.Context, params *SelectionRangeParams) ([]*SelectionRange, error)

	// SemanticTokensDelta.
	//
	// @since 3.16.0
	SemanticTokensDelta(ctx context.Context, params *SemanticTokensDeltaParams) (*SemanticTokensDeltaResult, error)

	// SemanticTokensRange.
	//
	// @since 3.16.0
	SemanticTokensRange(ctx context.Context, params *SemanticTokensRangeParams) (*SemanticTokens, error)

	// SemanticTokens.
	//
	// @since 3.16.0
	SemanticTokens(ctx context.Context, params *SemanticTokensParams) (*SemanticTokens, error)

	// Shutdown a shutdown request is sent from the client to the server. It is sent once when the client decides to
	// shutdown the server. The only notification that is sent after a shutdown request is the exit event.
	Shutdown(ctx context.Context) error

	SignatureHelp(ctx context.Context, params *SignatureHelpParams) (*SignatureHelp, error)

	// TextDocumentContent the `workspace/textDocumentContent` request is sent from the client to the server to request the content of a text document. 3.18.0 @proposed.
	//
	// @since 3.18.0 proposed
	TextDocumentContent(ctx context.Context, params *TextDocumentContentParams) (*TextDocumentContentResult, error)

	// TypeDefinition a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
	TypeDefinition(ctx context.Context, params *TypeDefinitionParams) (*TypeDefinitionResult, error)

	// TypeHierarchyPrepare a request to result a `TypeHierarchyItem` in a document at a given position. Can be used as an input
	// to a subtypes or supertypes type hierarchy.
	//
	// @since 3.17.0
	TypeHierarchyPrepare(ctx context.Context, params *TypeHierarchyPrepareParams) ([]*TypeHierarchyItem, error)

	// TypeHierarchySubtypes a request to resolve the subtypes for a given `TypeHierarchyItem`.
	//
	// @since 3.17.0
	TypeHierarchySubtypes(ctx context.Context, params *TypeHierarchySubtypesParams) ([]*TypeHierarchyItem, error)

	// TypeHierarchySupertypes a request to resolve the supertypes for a given `TypeHierarchyItem`.
	//
	// @since 3.17.0
	TypeHierarchySupertypes(ctx context.Context, params *TypeHierarchySupertypesParams) ([]*TypeHierarchyItem, error)

	// WillCreateFiles the will create files request is sent from the client to the server before files are actually created as long as the creation is triggered from within the client. The request can return a `WorkspaceEdit` which will be applied to workspace before the files are created. Hence the `WorkspaceEdit` can not manipulate the content of the file to be created.
	//
	// @since 3.16.0
	WillCreateFiles(ctx context.Context, params *CreateFilesParams) (*WorkspaceEdit, error)

	// WillDeleteFiles the did delete files notification is sent from the client to the server when files were deleted from
	// within the client.
	//
	// @since 3.16.0
	WillDeleteFiles(ctx context.Context, params *DeleteFilesParams) (*WorkspaceEdit, error)

	// WillRenameFiles the will rename files request is sent from the client to the server before files are actually renamed as long as the rename is triggered from within the client.
	//
	// @since 3.16.0
	WillRenameFiles(ctx context.Context, params *RenameFilesParams) (*WorkspaceEdit, error)

	// WillSaveTextDocumentWaitUntil a document will save request is sent from the client to the server before the document is actually saved. The request can return an array of TextEdits which will be applied to the text document before
	// it is saved. Please note that clients might drop results if computing the text edits took too long or if a server constantly fails on this request. This is done to keep the save fast and reliable.
	WillSaveTextDocumentWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) ([]*TextEdit, error)

	// WorkspaceDiagnostic the workspace diagnostic request definition.
	//
	// @since 3.17.0
	WorkspaceDiagnostic(ctx context.Context, params *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error)

	// WorkspaceSymbol a request to list project-wide symbols matching the query string given by the WorkspaceSymbolParams.
	// The response is of type SymbolInformation SymbolInformation[] or a Thenable that resolves to such. 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
	//
	// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
	WorkspaceSymbol(ctx context.Context, params *WorkspaceSymbolParams) (*WorkspaceSymbolResult, error)

	// WorkspaceSymbolResolve a request to resolve the range inside the workspace symbol's location.
	//
	// @since 3.17.0
	WorkspaceSymbolResolve(ctx context.Context, params *WorkspaceSymbol) (*WorkspaceSymbol, error)

	Request(ctx context.Context, method string, params any) (any, error)
}

// UnimplementedServer should be embedded to have forward compatible implementations.
type UnimplementedServer struct{}

func (UnimplementedServer) Cancel(ctx context.Context, params *CancelParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) Progress(ctx context.Context, params *ProgressParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidChangeNotebookDocument(ctx context.Context, params *DidChangeNotebookDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidChangeTextDocument(ctx context.Context, params *DidChangeTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidCloseNotebookDocument(ctx context.Context, params *DidCloseNotebookDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidCloseTextDocument(ctx context.Context, params *DidCloseTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidCreateFiles(ctx context.Context, params *CreateFilesParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidDeleteFiles(ctx context.Context, params *DeleteFilesParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidOpenNotebookDocument(ctx context.Context, params *DidOpenNotebookDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidOpenTextDocument(ctx context.Context, params *DidOpenTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidRenameFiles(ctx context.Context, params *RenameFilesParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidSaveNotebookDocument(ctx context.Context, params *DidSaveNotebookDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) DidSaveTextDocument(ctx context.Context, params *DidSaveTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) Exit(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) Initialized(ctx context.Context, params *InitializedParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) SetTrace(ctx context.Context, params *SetTraceParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) WillSaveTextDocument(ctx context.Context, params *WillSaveTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkDoneProgressCancel(ctx context.Context, params *WorkDoneProgressCancelParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) CallHierarchyIncomingCalls(ctx context.Context, params *CallHierarchyIncomingCallsParams) ([]*CallHierarchyIncomingCall, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CallHierarchyOutgoingCalls(ctx context.Context, params *CallHierarchyOutgoingCallsParams) ([]*CallHierarchyOutgoingCall, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CallHierarchyPrepare(ctx context.Context, params *CallHierarchyPrepareParams) ([]*CallHierarchyItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CodeAction(ctx context.Context, params *CodeActionParams) (*CodeActionRequestResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CodeActionResolve(ctx context.Context, params *CodeAction) (*CodeAction, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CodeLens(ctx context.Context, params *CodeLensParams) ([]*CodeLens, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CodeLensResolve(ctx context.Context, params *CodeLens) (*CodeLens, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) ColorPresentation(ctx context.Context, params *ColorPresentationParams) ([]*ColorPresentation, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Completion(ctx context.Context, params *CompletionParams) (*CompletionResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CompletionResolve(ctx context.Context, params *CompletionItem) (*CompletionItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Declaration(ctx context.Context, params *DeclarationParams) (*DeclarationResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Definition(ctx context.Context, params *DefinitionParams) (*DefinitionResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentColor(ctx context.Context, params *DocumentColorParams) ([]*ColorInformation, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentDiagnostic(ctx context.Context, params *DocumentDiagnosticParams) (*DocumentDiagnosticReport, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentFormatting(ctx context.Context, params *DocumentFormattingParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentHighlight(ctx context.Context, params *DocumentHighlightParams) ([]*DocumentHighlight, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentLink(ctx context.Context, params *DocumentLinkParams) ([]*DocumentLink, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentLinkResolve(ctx context.Context, params *DocumentLink) (*DocumentLink, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentOnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentRangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentRangesFormatting(ctx context.Context, params *DocumentRangesFormattingParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentSymbol(ctx context.Context, params *DocumentSymbolParams) (*DocumentSymbolResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) ExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (any, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) FoldingRange(ctx context.Context, params *FoldingRangeParams) ([]*FoldingRange, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Hover(ctx context.Context, params *HoverParams) (*Hover, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Implementation(ctx context.Context, params *ImplementationParams) (*ImplementationResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Initialize(ctx context.Context, params *InitializeParams) (*InitializeResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) InlayHint(ctx context.Context, params *InlayHintParams) ([]*InlayHint, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) InlayHintResolve(ctx context.Context, params *InlayHint) (*InlayHint, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) InlineCompletion(ctx context.Context, params *InlineCompletionParams) (*InlineCompletionResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) InlineValue(ctx context.Context, params *InlineValueParams) ([]*InlineValue, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) LinkedEditingRange(ctx context.Context, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Moniker(ctx context.Context, params *MonikerParams) ([]*Moniker, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) PrepareRename(ctx context.Context, params *PrepareRenameParams) (*PrepareRenameResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) References(ctx context.Context, params *ReferenceParams) ([]*Location, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Rename(ctx context.Context, params *RenameParams) (*WorkspaceEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) SelectionRange(ctx context.Context, params *SelectionRangeParams) ([]*SelectionRange, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) SemanticTokensDelta(ctx context.Context, params *SemanticTokensDeltaParams) (*SemanticTokensDeltaResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) SemanticTokensRange(ctx context.Context, params *SemanticTokensRangeParams) (*SemanticTokens, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) SemanticTokens(ctx context.Context, params *SemanticTokensParams) (*SemanticTokens, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Shutdown(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) SignatureHelp(ctx context.Context, params *SignatureHelpParams) (*SignatureHelp, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentContent(ctx context.Context, params *TextDocumentContentParams) (*TextDocumentContentResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TypeDefinition(ctx context.Context, params *TypeDefinitionParams) (*TypeDefinitionResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TypeHierarchyPrepare(ctx context.Context, params *TypeHierarchyPrepareParams) ([]*TypeHierarchyItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TypeHierarchySubtypes(ctx context.Context, params *TypeHierarchySubtypesParams) ([]*TypeHierarchyItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TypeHierarchySupertypes(ctx context.Context, params *TypeHierarchySupertypesParams) ([]*TypeHierarchyItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WillCreateFiles(ctx context.Context, params *CreateFilesParams) (*WorkspaceEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WillDeleteFiles(ctx context.Context, params *DeleteFilesParams) (*WorkspaceEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WillRenameFiles(ctx context.Context, params *RenameFilesParams) (*WorkspaceEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WillSaveTextDocumentWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceDiagnostic(ctx context.Context, params *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceSymbol(ctx context.Context, params *WorkspaceSymbolParams) (*WorkspaceSymbolResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceSymbolResolve(ctx context.Context, params *WorkspaceSymbol) (*WorkspaceSymbol, error) {
	return nil, jsonrpc2.ErrInternal
}
