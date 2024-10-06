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
	MethodSetTrace                            ServerMethod = "$/setTrace"                             // server notification
	MethodExit                                ServerMethod = "exit"                                   // server notification
	MethodInitialized                         ServerMethod = "initialized"                            // server notification
	MethodNotebookDocumentDidChange           ServerMethod = "notebookDocument/didChange"             // server notification
	MethodNotebookDocumentDidClose            ServerMethod = "notebookDocument/didClose"              // server notification
	MethodNotebookDocumentDidOpen             ServerMethod = "notebookDocument/didOpen"               // server notification
	MethodNotebookDocumentDidSave             ServerMethod = "notebookDocument/didSave"               // server notification
	MethodTextDocumentDidChange               ServerMethod = "textDocument/didChange"                 // server notification
	MethodTextDocumentDidClose                ServerMethod = "textDocument/didClose"                  // server notification
	MethodTextDocumentDidOpen                 ServerMethod = "textDocument/didOpen"                   // server notification
	MethodTextDocumentDidSave                 ServerMethod = "textDocument/didSave"                   // server notification
	MethodTextDocumentWillSave                ServerMethod = "textDocument/willSave"                  // server notification
	MethodWindowWorkDoneProgressCancel        ServerMethod = "window/workDoneProgress/cancel"         // server notification
	MethodWorkspaceDidChangeConfiguration     ServerMethod = "workspace/didChangeConfiguration"       // server notification
	MethodWorkspaceDidChangeWatchedFiles      ServerMethod = "workspace/didChangeWatchedFiles"        // server notification
	MethodWorkspaceDidChangeWorkspaceFolders  ServerMethod = "workspace/didChangeWorkspaceFolders"    // server notification
	MethodWorkspaceDidCreateFiles             ServerMethod = "workspace/didCreateFiles"               // server notification
	MethodWorkspaceDidDeleteFiles             ServerMethod = "workspace/didDeleteFiles"               // server notification
	MethodWorkspaceDidRenameFiles             ServerMethod = "workspace/didRenameFiles"               // server notification
	MethodCallHierarchyIncomingCalls          ServerMethod = "callHierarchy/incomingCalls"            // server request
	MethodCallHierarchyOutgoingCalls          ServerMethod = "callHierarchy/outgoingCalls"            // server request
	MethodCodeActionResolve                   ServerMethod = "codeAction/resolve"                     // server request
	MethodCodeLensResolve                     ServerMethod = "codeLens/resolve"                       // server request
	MethodCompletionItemResolve               ServerMethod = "completionItem/resolve"                 // server request
	MethodDocumentLinkResolve                 ServerMethod = "documentLink/resolve"                   // server request
	MethodInitialize                          ServerMethod = "initialize"                             // server request
	MethodInlayHintResolve                    ServerMethod = "inlayHint/resolve"                      // server request
	MethodShutdown                            ServerMethod = "shutdown"                               // server request
	MethodTextDocumentCodeAction              ServerMethod = "textDocument/codeAction"                // server request
	MethodTextDocumentCodeLens                ServerMethod = "textDocument/codeLens"                  // server request
	MethodTextDocumentColorPresentation       ServerMethod = "textDocument/colorPresentation"         // server request
	MethodTextDocumentCompletion              ServerMethod = "textDocument/completion"                // server request
	MethodTextDocumentDeclaration             ServerMethod = "textDocument/declaration"               // server request
	MethodTextDocumentDefinition              ServerMethod = "textDocument/definition"                // server request
	MethodTextDocumentDiagnostic              ServerMethod = "textDocument/diagnostic"                // server request
	MethodTextDocumentDocumentColor           ServerMethod = "textDocument/documentColor"             // server request
	MethodTextDocumentDocumentHighlight       ServerMethod = "textDocument/documentHighlight"         // server request
	MethodTextDocumentDocumentLink            ServerMethod = "textDocument/documentLink"              // server request
	MethodTextDocumentDocumentSymbol          ServerMethod = "textDocument/documentSymbol"            // server request
	MethodTextDocumentFoldingRange            ServerMethod = "textDocument/foldingRange"              // server request
	MethodTextDocumentFormatting              ServerMethod = "textDocument/formatting"                // server request
	MethodTextDocumentHover                   ServerMethod = "textDocument/hover"                     // server request
	MethodTextDocumentImplementation          ServerMethod = "textDocument/implementation"            // server request
	MethodTextDocumentInlayHint               ServerMethod = "textDocument/inlayHint"                 // server request
	MethodTextDocumentInlineCompletion        ServerMethod = "textDocument/inlineCompletion"          // server request
	MethodTextDocumentInlineValue             ServerMethod = "textDocument/inlineValue"               // server request
	MethodTextDocumentLinkedEditingRange      ServerMethod = "textDocument/linkedEditingRange"        // server request
	MethodTextDocumentMoniker                 ServerMethod = "textDocument/moniker"                   // server request
	MethodTextDocumentOnTypeFormatting        ServerMethod = "textDocument/onTypeFormatting"          // server request
	MethodTextDocumentPrepareCallHierarchy    ServerMethod = "textDocument/prepareCallHierarchy"      // server request
	MethodTextDocumentPrepareRename           ServerMethod = "textDocument/prepareRename"             // server request
	MethodTextDocumentPrepareTypeHierarchy    ServerMethod = "textDocument/prepareTypeHierarchy"      // server request
	MethodTextDocumentRangeFormatting         ServerMethod = "textDocument/rangeFormatting"           // server request
	MethodTextDocumentRangesFormatting        ServerMethod = "textDocument/rangesFormatting"          // server request
	MethodTextDocumentReferences              ServerMethod = "textDocument/references"                // server request
	MethodTextDocumentRename                  ServerMethod = "textDocument/rename"                    // server request
	MethodTextDocumentSelectionRange          ServerMethod = "textDocument/selectionRange"            // server request
	MethodTextDocumentSemanticTokensFull      ServerMethod = "textDocument/semanticTokens/full"       // server request
	MethodTextDocumentSemanticTokensFullDelta ServerMethod = "textDocument/semanticTokens/full/delta" // server request
	MethodTextDocumentSemanticTokensRange     ServerMethod = "textDocument/semanticTokens/range"      // server request
	MethodTextDocumentSignatureHelp           ServerMethod = "textDocument/signatureHelp"             // server request
	MethodTextDocumentTypeDefinition          ServerMethod = "textDocument/typeDefinition"            // server request
	MethodTextDocumentWillSaveWaitUntil       ServerMethod = "textDocument/willSaveWaitUntil"         // server request
	MethodTypeHierarchySubtypes               ServerMethod = "typeHierarchy/subtypes"                 // server request
	MethodTypeHierarchySupertypes             ServerMethod = "typeHierarchy/supertypes"               // server request
	MethodWorkspaceDiagnostic                 ServerMethod = "workspace/diagnostic"                   // server request
	MethodWorkspaceExecuteCommand             ServerMethod = "workspace/executeCommand"               // server request
	MethodWorkspaceSymbol                     ServerMethod = "workspace/symbol"                       // server request
	MethodWorkspaceTextDocumentContent        ServerMethod = "workspace/textDocumentContent"          // server request
	MethodWorkspaceWillCreateFiles            ServerMethod = "workspace/willCreateFiles"              // server request
	MethodWorkspaceWillDeleteFiles            ServerMethod = "workspace/willDeleteFiles"              // server request
	MethodWorkspaceWillRenameFiles            ServerMethod = "workspace/willRenameFiles"              // server request
	MethodWorkspaceSymbolResolve              ServerMethod = "workspaceSymbol/resolve"                // server request
)

type Server interface {
	CancelRequest(ctx context.Context, params *CancelParams) error

	Progress(ctx context.Context, params *ProgressParams) error

	SetTrace(ctx context.Context, params *SetTraceParams) error

	// Exit the exit event is sent from the client to the server to ask the server to exit its process.
	Exit(ctx context.Context) error

	// Initialized the initialized notification is sent from the client to the server after the client is fully initialized and the server is allowed to send requests from the server to the client.
	Initialized(ctx context.Context, params *InitializedParams) error

	NotebookDocumentDidChange(ctx context.Context, params *DidChangeNotebookDocumentParams) error

	// NotebookDocumentDidClose a notification sent when a notebook closes.
	//
	// @since 3.17.0
	NotebookDocumentDidClose(ctx context.Context, params *DidCloseNotebookDocumentParams) error

	// NotebookDocumentDidOpen a notification sent when a notebook opens.
	//
	// @since 3.17.0
	NotebookDocumentDidOpen(ctx context.Context, params *DidOpenNotebookDocumentParams) error

	// NotebookDocumentDidSave a notification sent when a notebook document is saved.
	//
	// @since 3.17.0
	NotebookDocumentDidSave(ctx context.Context, params *DidSaveNotebookDocumentParams) error

	// TextDocumentDidChange the document change notification is sent from the client to the server to signal changes to a text document.
	TextDocumentDidChange(ctx context.Context, params *DidChangeTextDocumentParams) error

	// TextDocumentDidClose the document close notification is sent from the client to the server when the document got closed in the client. The document's truth now exists where the document's uri points to (e.g. if the document's uri is a file uri the truth now exists on disk). As with the open notification the close notification is about managing the document's content. Receiving a close notification doesn't mean that the document was open in an editor before. A close notification requires a previous open notification to be sent.
	TextDocumentDidClose(ctx context.Context, params *DidCloseTextDocumentParams) error

	// TextDocumentDidOpen the document open notification is sent from the client to the server to signal newly opened text documents. The document's truth is now managed by the client and the server must not try to read the document's truth using the document's uri. Open in this sense means it is managed by the client. It doesn't necessarily mean that its content is presented in an editor. An open notification must not be sent more than once without a corresponding close notification send before. This means open and close notification must be balanced and the max open count is one.
	TextDocumentDidOpen(ctx context.Context, params *DidOpenTextDocumentParams) error

	// TextDocumentDidSave the document save notification is sent from the client to the server when the document got saved in the client.
	TextDocumentDidSave(ctx context.Context, params *DidSaveTextDocumentParams) error

	// TextDocumentWillSave a document will save notification is sent from the client to the server before the document is actually saved.
	TextDocumentWillSave(ctx context.Context, params *WillSaveTextDocumentParams) error

	// WindowWorkDoneProgressCancel the `window/workDoneProgress/cancel` notification is sent from the client to the server to cancel a progress initiated on the server side.
	WindowWorkDoneProgressCancel(ctx context.Context, params *WorkDoneProgressCancelParams) error

	// WorkspaceDidChangeConfiguration the configuration change notification is sent from the client to the server when the client's configuration has changed. The notification contains the changed configuration as defined by the language client.
	WorkspaceDidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) error

	// WorkspaceDidChangeWatchedFiles the watched files notification is sent from the client to the server when the client detects changes
	// to file watched by the language client.
	WorkspaceDidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) error

	// WorkspaceDidChangeWorkspaceFolders the `workspace/didChangeWorkspaceFolders` notification is sent from the client to the server when the workspace folder configuration changes.
	WorkspaceDidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) error

	// WorkspaceDidCreateFiles the did create files notification is sent from the client to the server when files were created from
	// within the client.
	//
	// @since 3.16.0
	WorkspaceDidCreateFiles(ctx context.Context, params *CreateFilesParams) error

	// WorkspaceDidDeleteFiles the will delete files request is sent from the client to the server before files are actually deleted as long as the deletion is triggered from within the client.
	//
	// @since 3.16.0
	WorkspaceDidDeleteFiles(ctx context.Context, params *DeleteFilesParams) error

	// WorkspaceDidRenameFiles the did rename files notification is sent from the client to the server when files were renamed from
	// within the client.
	//
	// @since 3.16.0
	WorkspaceDidRenameFiles(ctx context.Context, params *RenameFilesParams) error
	// CallHierarchyIncomingCalls a request to resolve the incoming calls for a given `CallHierarchyItem`.
	//
	// @since 3.16.0
	CallHierarchyIncomingCalls(ctx context.Context, params *CallHierarchyIncomingCallsParams) ([]*CallHierarchyIncomingCall, error)

	// CallHierarchyOutgoingCalls a request to resolve the outgoing calls for a given `CallHierarchyItem`.
	//
	// @since 3.16.0
	CallHierarchyOutgoingCalls(ctx context.Context, params *CallHierarchyOutgoingCallsParams) ([]*CallHierarchyOutgoingCall, error)

	// CodeActionResolve request to resolve additional information for a given code action.The request's parameter is of type
	// CodeAction the response is of type CodeAction or a Thenable that resolves to such.
	CodeActionResolve(ctx context.Context, params *CodeAction) (*CodeAction, error)

	// CodeLensResolve a request to resolve a command for a given code lens.
	CodeLensResolve(ctx context.Context, params *CodeLens) (*CodeLens, error)

	// CompletionItemResolve request to resolve additional information for a given completion item.The request's parameter is of type CompletionItem the response is of type CompletionItem or a Thenable that resolves to such.
	CompletionItemResolve(ctx context.Context, params *CompletionItem) (*CompletionItem, error)

	// DocumentLinkResolve request to resolve additional information for a given document link. The request's parameter is of type DocumentLink the response is of type DocumentLink or a Thenable that resolves to such.
	DocumentLinkResolve(ctx context.Context, params *DocumentLink) (*DocumentLink, error)

	// Initialize the initialize request is sent from the client to the server. It is sent once as the request after starting up the server. The requests parameter is of type InitializeParams the response if of type InitializeResult of a Thenable that resolves to such.
	Initialize(ctx context.Context, params *InitializeParams) (*InitializeResult, error)

	// InlayHintResolve a request to resolve additional properties for an inlay hint. The request's parameter is of type InlayHint, the response is of type InlayHint or a Thenable that resolves to such.
	//
	// @since 3.17.0
	InlayHintResolve(ctx context.Context, params *InlayHint) (*InlayHint, error)

	// Shutdown a shutdown request is sent from the client to the server. It is sent once when the client decides to
	// shutdown the server. The only notification that is sent after a shutdown request is the exit event.
	Shutdown(ctx context.Context) error

	// TextDocumentCodeAction a request to provide commands for the given text document and range.
	TextDocumentCodeAction(ctx context.Context, params *CodeActionParams) (*TextDocumentCodeActionResult, error)

	// TextDocumentCodeLens a request to provide code lens for the given text document.
	TextDocumentCodeLens(ctx context.Context, params *CodeLensParams) ([]*CodeLens, error)

	// TextDocumentColorPresentation a request to list all presentation for a color. The request's parameter is of type ColorPresentationParams the response is of type ColorInformation ColorInformation[] or a Thenable that resolves to such.
	TextDocumentColorPresentation(ctx context.Context, params *ColorPresentationParams) ([]*ColorPresentation, error)

	// TextDocumentCompletion request to request completion at a given text document position. The request's parameter is of type TextDocumentPosition the response is of type CompletionItem CompletionItem[] or CompletionList or a Thenable that resolves to such. The request can delay the computation of the CompletionItem.detail `detail` and CompletionItem.documentation `documentation` properties to the `completionItem/resolve` request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
	// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
	TextDocumentCompletion(ctx context.Context, params *CompletionParams) (*TextDocumentCompletionResult, error)

	// TextDocumentDeclaration a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Declaration or a
	// typed array of DeclarationLink or a Thenable that resolves to such.
	TextDocumentDeclaration(ctx context.Context, params *DeclarationParams) (*TextDocumentDeclarationResult, error)

	// TextDocumentDefinition a request to resolve the definition location of a symbol at a given text document position. The request's parameter is of type TextDocumentPosition the response is of either type Definition or a typed
	// array of DefinitionLink or a Thenable that resolves to such.
	TextDocumentDefinition(ctx context.Context, params *DefinitionParams) (*TextDocumentDefinitionResult, error)

	// TextDocumentDiagnostic the document diagnostic request definition.
	//
	// @since 3.17.0
	TextDocumentDiagnostic(ctx context.Context, params *DocumentDiagnosticParams) (*DocumentDiagnosticReport, error)

	// TextDocumentDocumentColor a request to list all color symbols found in a given text document. The request's parameter is of type DocumentColorParams the response is of type ColorInformation ColorInformation[] or a Thenable that resolves to such.
	TextDocumentDocumentColor(ctx context.Context, params *DocumentColorParams) ([]*ColorInformation, error)

	// TextDocumentDocumentHighlight request to resolve a DocumentHighlight for a given text document position. The request's parameter is of type TextDocumentPosition the request response is an array of type DocumentHighlight or a Thenable that resolves to such.
	TextDocumentDocumentHighlight(ctx context.Context, params *DocumentHighlightParams) ([]*DocumentHighlight, error)

	// TextDocumentDocumentLink a request to provide document links.
	TextDocumentDocumentLink(ctx context.Context, params *DocumentLinkParams) ([]*DocumentLink, error)

	// TextDocumentDocumentSymbol a request to list all symbols found in a given text document. The request's parameter is of type TextDocumentIdentifier the response is of type SymbolInformation SymbolInformation[] or a Thenable that
	// resolves to such.
	TextDocumentDocumentSymbol(ctx context.Context, params *DocumentSymbolParams) (*TextDocumentDocumentSymbolResult, error)

	// TextDocumentFoldingRange a request to provide folding ranges in a document. The request's parameter is of type FoldingRangeParams, the response is of type FoldingRangeList or a Thenable that resolves to such.
	TextDocumentFoldingRange(ctx context.Context, params *FoldingRangeParams) ([]*FoldingRange, error)

	// TextDocumentFormatting a request to format a whole document.
	TextDocumentFormatting(ctx context.Context, params *DocumentFormattingParams) ([]*TextEdit, error)

	// TextDocumentHover request to request hover information at a given text document position. The request's parameter is of type TextDocumentPosition the response is of type Hover or a Thenable that resolves to such.
	TextDocumentHover(ctx context.Context, params *HoverParams) (*Hover, error)

	// TextDocumentImplementation a request to resolve the implementation locations of a symbol at a given text document position. The
	// request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
	TextDocumentImplementation(ctx context.Context, params *ImplementationParams) (*TextDocumentImplementationResult, error)

	// TextDocumentInlayHint a request to provide inlay hints in a document. The request's parameter is of type InlayHintsParams,
	// the response is of type InlayHint InlayHint[] or a Thenable that resolves to such.
	//
	// @since 3.17.0
	TextDocumentInlayHint(ctx context.Context, params *InlayHintParams) ([]*InlayHint, error)

	// TextDocumentInlineCompletion a request to provide inline completions in a document. The request's parameter is of type InlineCompletionParams, the response is of type InlineCompletion InlineCompletion[] or a Thenable that resolves to such. 3.18.0 @proposed.
	//
	// @since 3.18.0 proposed
	TextDocumentInlineCompletion(ctx context.Context, params *InlineCompletionParams) (*TextDocumentInlineCompletionResult, error)

	// TextDocumentInlineValue a request to provide inline values in a document. The request's parameter is of type InlineValueParams, the response is of type InlineValue InlineValue[] or a Thenable that resolves to such.
	//
	// @since 3.17.0
	TextDocumentInlineValue(ctx context.Context, params *InlineValueParams) ([]*InlineValue, error)

	// TextDocumentLinkedEditingRange a request to provide ranges that can be edited together.
	//
	// @since 3.16.0
	TextDocumentLinkedEditingRange(ctx context.Context, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error)

	// TextDocumentMoniker a request to get the moniker of a symbol at a given text document position. The request parameter is
	// of type TextDocumentPositionParams. The response is of type Moniker Moniker[] or `null`.
	TextDocumentMoniker(ctx context.Context, params *MonikerParams) ([]*Moniker, error)

	// TextDocumentOnTypeFormatting a request to format a document on type.
	TextDocumentOnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) ([]*TextEdit, error)

	// TextDocumentPrepareCallHierarchy a request to result a `CallHierarchyItem` in a document at a given position. Can be used as an input
	// to an incoming or outgoing call hierarchy.
	//
	// @since 3.16.0
	TextDocumentPrepareCallHierarchy(ctx context.Context, params *CallHierarchyPrepareParams) ([]*CallHierarchyItem, error)

	// TextDocumentPrepareRename a request to test and perform the setup necessary for a rename. 3.16 - support for default behavior.
	//
	// @since 3.16 - support for default behavior
	TextDocumentPrepareRename(ctx context.Context, params *PrepareRenameParams) (*PrepareRenameResult, error)

	// TextDocumentPrepareTypeHierarchy a request to result a `TypeHierarchyItem` in a document at a given position. Can be used as an input
	// to a subtypes or supertypes type hierarchy.
	//
	// @since 3.17.0
	TextDocumentPrepareTypeHierarchy(ctx context.Context, params *TypeHierarchyPrepareParams) ([]*TypeHierarchyItem, error)

	// TextDocumentRangeFormatting a request to format a range in a document.
	TextDocumentRangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) ([]*TextEdit, error)

	// TextDocumentRangesFormatting a request to format ranges in a document.  3.18.0 @proposed.
	//
	// @since 3.18.0 proposed
	TextDocumentRangesFormatting(ctx context.Context, params *DocumentRangesFormattingParams) ([]*TextEdit, error)

	// TextDocumentReferences a request to resolve project-wide references for the symbol denoted by the given text document position. The request's parameter is of type ReferenceParams the response is of type Location Location[] or a Thenable that resolves to such.
	TextDocumentReferences(ctx context.Context, params *ReferenceParams) ([]*Location, error)

	// TextDocumentRename a request to rename a symbol.
	TextDocumentRename(ctx context.Context, params *RenameParams) (*WorkspaceEdit, error)

	// TextDocumentSelectionRange a request to provide selection ranges in a document. The request's parameter is of type SelectionRangeParams, the response is of type SelectionRange SelectionRange[] or a Thenable that resolves to such.
	TextDocumentSelectionRange(ctx context.Context, params *SelectionRangeParams) ([]*SelectionRange, error)

	// TextDocumentSemanticTokensFull.
	//
	// @since 3.16.0
	TextDocumentSemanticTokensFull(ctx context.Context, params *SemanticTokensParams) (*SemanticTokens, error)

	// TextDocumentSemanticTokensFullDelta.
	//
	// @since 3.16.0
	TextDocumentSemanticTokensFullDelta(ctx context.Context, params *SemanticTokensDeltaParams) (*TextDocumentSemanticTokensFullDeltaResult, error)

	// TextDocumentSemanticTokensRange.
	//
	// @since 3.16.0
	TextDocumentSemanticTokensRange(ctx context.Context, params *SemanticTokensRangeParams) (*SemanticTokens, error)

	TextDocumentSignatureHelp(ctx context.Context, params *SignatureHelpParams) (*SignatureHelp, error)

	// TextDocumentTypeDefinition a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
	TextDocumentTypeDefinition(ctx context.Context, params *TypeDefinitionParams) (*TextDocumentTypeDefinitionResult, error)

	// TextDocumentWillSaveWaitUntil a document will save request is sent from the client to the server before the document is actually saved. The request can return an array of TextEdits which will be applied to the text document before
	// it is saved. Please note that clients might drop results if computing the text edits took too long or if a server constantly fails on this request. This is done to keep the save fast and reliable.
	TextDocumentWillSaveWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) ([]*TextEdit, error)

	// TypeHierarchySubtypes a request to resolve the subtypes for a given `TypeHierarchyItem`.
	//
	// @since 3.17.0
	TypeHierarchySubtypes(ctx context.Context, params *TypeHierarchySubtypesParams) ([]*TypeHierarchyItem, error)

	// TypeHierarchySupertypes a request to resolve the supertypes for a given `TypeHierarchyItem`.
	//
	// @since 3.17.0
	TypeHierarchySupertypes(ctx context.Context, params *TypeHierarchySupertypesParams) ([]*TypeHierarchyItem, error)

	// WorkspaceDiagnostic the workspace diagnostic request definition.
	//
	// @since 3.17.0
	WorkspaceDiagnostic(ctx context.Context, params *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error)

	// WorkspaceExecuteCommand a request send from the client to the server to execute a command. The request might return a workspace edit which the client will apply to the workspace.
	WorkspaceExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (any, error)

	// WorkspaceSymbol a request to list project-wide symbols matching the query string given by the WorkspaceSymbolParams.
	// The response is of type SymbolInformation SymbolInformation[] or a Thenable that resolves to such. 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
	//
	// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
	WorkspaceSymbol(ctx context.Context, params *WorkspaceSymbolParams) (*WorkspaceSymbolResult, error)

	// WorkspaceTextDocumentContent the `workspace/textDocumentContent` request is sent from the client to the server to request the content of a text document. 3.18.0 @proposed.
	//
	// @since 3.18.0 proposed
	WorkspaceTextDocumentContent(ctx context.Context, params *TextDocumentContentParams) (*TextDocumentContentResult, error)

	// WorkspaceWillCreateFiles the will create files request is sent from the client to the server before files are actually created as long as the creation is triggered from within the client. The request can return a `WorkspaceEdit` which will be applied to workspace before the files are created. Hence the `WorkspaceEdit` can not manipulate the content of the file to be created.
	//
	// @since 3.16.0
	WorkspaceWillCreateFiles(ctx context.Context, params *CreateFilesParams) (*WorkspaceEdit, error)

	// WorkspaceWillDeleteFiles the did delete files notification is sent from the client to the server when files were deleted from
	// within the client.
	//
	// @since 3.16.0
	WorkspaceWillDeleteFiles(ctx context.Context, params *DeleteFilesParams) (*WorkspaceEdit, error)

	// WorkspaceWillRenameFiles the will rename files request is sent from the client to the server before files are actually renamed as long as the rename is triggered from within the client.
	//
	// @since 3.16.0
	WorkspaceWillRenameFiles(ctx context.Context, params *RenameFilesParams) (*WorkspaceEdit, error)

	// WorkspaceSymbolResolve a request to resolve the range inside the workspace symbol's location.
	//
	// @since 3.17.0
	WorkspaceSymbolResolve(ctx context.Context, params *WorkspaceSymbol) (*WorkspaceSymbol, error)

	Request(ctx context.Context, method string, params any) (any, error)
}

// UnimplementedServer should be embedded to have forward compatible implementations.
type UnimplementedServer struct{}

func (UnimplementedServer) CancelRequest(ctx context.Context, params *CancelParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) Progress(ctx context.Context, params *ProgressParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) SetTrace(ctx context.Context, params *SetTraceParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) Exit(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) Initialized(ctx context.Context, params *InitializedParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) NotebookDocumentDidChange(ctx context.Context, params *DidChangeNotebookDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) NotebookDocumentDidClose(ctx context.Context, params *DidCloseNotebookDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) NotebookDocumentDidOpen(ctx context.Context, params *DidOpenNotebookDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) NotebookDocumentDidSave(ctx context.Context, params *DidSaveNotebookDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDidChange(ctx context.Context, params *DidChangeTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDidClose(ctx context.Context, params *DidCloseTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDidOpen(ctx context.Context, params *DidOpenTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDidSave(ctx context.Context, params *DidSaveTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentWillSave(ctx context.Context, params *WillSaveTextDocumentParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) WindowWorkDoneProgressCancel(ctx context.Context, params *WorkDoneProgressCancelParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceDidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceDidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceDidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceDidCreateFiles(ctx context.Context, params *CreateFilesParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceDidDeleteFiles(ctx context.Context, params *DeleteFilesParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceDidRenameFiles(ctx context.Context, params *RenameFilesParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) CallHierarchyIncomingCalls(ctx context.Context, params *CallHierarchyIncomingCallsParams) ([]*CallHierarchyIncomingCall, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CallHierarchyOutgoingCalls(ctx context.Context, params *CallHierarchyOutgoingCallsParams) ([]*CallHierarchyOutgoingCall, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CodeActionResolve(ctx context.Context, params *CodeAction) (*CodeAction, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CodeLensResolve(ctx context.Context, params *CodeLens) (*CodeLens, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) CompletionItemResolve(ctx context.Context, params *CompletionItem) (*CompletionItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) DocumentLinkResolve(ctx context.Context, params *DocumentLink) (*DocumentLink, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Initialize(ctx context.Context, params *InitializeParams) (*InitializeResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) InlayHintResolve(ctx context.Context, params *InlayHint) (*InlayHint, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) Shutdown(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentCodeAction(ctx context.Context, params *CodeActionParams) (*TextDocumentCodeActionResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentCodeLens(ctx context.Context, params *CodeLensParams) ([]*CodeLens, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentColorPresentation(ctx context.Context, params *ColorPresentationParams) ([]*ColorPresentation, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentCompletion(ctx context.Context, params *CompletionParams) (*TextDocumentCompletionResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDeclaration(ctx context.Context, params *DeclarationParams) (*TextDocumentDeclarationResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDefinition(ctx context.Context, params *DefinitionParams) (*TextDocumentDefinitionResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDiagnostic(ctx context.Context, params *DocumentDiagnosticParams) (*DocumentDiagnosticReport, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDocumentColor(ctx context.Context, params *DocumentColorParams) ([]*ColorInformation, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDocumentHighlight(ctx context.Context, params *DocumentHighlightParams) ([]*DocumentHighlight, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDocumentLink(ctx context.Context, params *DocumentLinkParams) ([]*DocumentLink, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentDocumentSymbol(ctx context.Context, params *DocumentSymbolParams) (*TextDocumentDocumentSymbolResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentFoldingRange(ctx context.Context, params *FoldingRangeParams) ([]*FoldingRange, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentFormatting(ctx context.Context, params *DocumentFormattingParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentHover(ctx context.Context, params *HoverParams) (*Hover, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentImplementation(ctx context.Context, params *ImplementationParams) (*TextDocumentImplementationResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentInlayHint(ctx context.Context, params *InlayHintParams) ([]*InlayHint, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentInlineCompletion(ctx context.Context, params *InlineCompletionParams) (*TextDocumentInlineCompletionResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentInlineValue(ctx context.Context, params *InlineValueParams) ([]*InlineValue, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentLinkedEditingRange(ctx context.Context, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentMoniker(ctx context.Context, params *MonikerParams) ([]*Moniker, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentOnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentPrepareCallHierarchy(ctx context.Context, params *CallHierarchyPrepareParams) ([]*CallHierarchyItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentPrepareRename(ctx context.Context, params *PrepareRenameParams) (*PrepareRenameResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentPrepareTypeHierarchy(ctx context.Context, params *TypeHierarchyPrepareParams) ([]*TypeHierarchyItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentRangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentRangesFormatting(ctx context.Context, params *DocumentRangesFormattingParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentReferences(ctx context.Context, params *ReferenceParams) ([]*Location, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentRename(ctx context.Context, params *RenameParams) (*WorkspaceEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentSelectionRange(ctx context.Context, params *SelectionRangeParams) ([]*SelectionRange, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentSemanticTokensFull(ctx context.Context, params *SemanticTokensParams) (*SemanticTokens, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentSemanticTokensFullDelta(ctx context.Context, params *SemanticTokensDeltaParams) (*TextDocumentSemanticTokensFullDeltaResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentSemanticTokensRange(ctx context.Context, params *SemanticTokensRangeParams) (*SemanticTokens, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentSignatureHelp(ctx context.Context, params *SignatureHelpParams) (*SignatureHelp, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentTypeDefinition(ctx context.Context, params *TypeDefinitionParams) (*TextDocumentTypeDefinitionResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TextDocumentWillSaveWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) ([]*TextEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TypeHierarchySubtypes(ctx context.Context, params *TypeHierarchySubtypesParams) ([]*TypeHierarchyItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) TypeHierarchySupertypes(ctx context.Context, params *TypeHierarchySupertypesParams) ([]*TypeHierarchyItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceDiagnostic(ctx context.Context, params *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (any, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceSymbol(ctx context.Context, params *WorkspaceSymbolParams) (*WorkspaceSymbolResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceTextDocumentContent(ctx context.Context, params *TextDocumentContentParams) (*TextDocumentContentResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceWillCreateFiles(ctx context.Context, params *CreateFilesParams) (*WorkspaceEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceWillDeleteFiles(ctx context.Context, params *DeleteFilesParams) (*WorkspaceEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceWillRenameFiles(ctx context.Context, params *RenameFilesParams) (*WorkspaceEdit, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedServer) WorkspaceSymbolResolve(ctx context.Context, params *WorkspaceSymbol) (*WorkspaceSymbol, error) {
	return nil, jsonrpc2.ErrInternal
}
