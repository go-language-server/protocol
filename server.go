// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
)

// Server represents a Language Server Protocol server.
type Server interface {
	Initialize(ctx context.Context, params *InitializeParams) (result *InitializeResult, err error)
	Initialized(ctx context.Context, params *InitializedParams) (err error)
	Shutdown(ctx context.Context) (err error)
	Exit(ctx context.Context) (err error)
	CodeAction(ctx context.Context, params *CodeActionParams) (result []CodeAction, err error)
	CodeLens(ctx context.Context, params *CodeLensParams) (result []CodeLens, err error)
	CodeLensResolve(ctx context.Context, params *CodeLens) (result *CodeLens, err error)
	ColorPresentation(ctx context.Context, params *ColorPresentationParams) (result []ColorPresentation, err error)
	Completion(ctx context.Context, params *CompletionParams) (result *CompletionList, err error)
	CompletionResolve(ctx context.Context, params *CompletionItem) (result *CompletionItem, err error)
	Declaration(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error)
	Definition(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error)
	DidChange(ctx context.Context, params *DidChangeTextDocumentParams) (err error)
	DidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) (err error)
	DidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) (err error)
	DidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) (err error)
	DidClose(ctx context.Context, params *DidCloseTextDocumentParams) (err error)
	DidOpen(ctx context.Context, params *DidOpenTextDocumentParams) (err error)
	DidSave(ctx context.Context, params *DidSaveTextDocumentParams) (err error)
	DocumentColor(ctx context.Context, params *DocumentColorParams) (result []ColorInformation, err error)
	DocumentHighlight(ctx context.Context, params *TextDocumentPositionParams) (result []DocumentHighlight, err error)
	DocumentLink(ctx context.Context, params *DocumentLinkParams) (result []DocumentLink, err error)
	DocumentLinkResolve(ctx context.Context, params *DocumentLink) (result *DocumentLink, err error)
	DocumentSymbol(ctx context.Context, params *DocumentSymbolParams) (result []DocumentSymbol, err error)
	ExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (result interface{}, err error)
	FoldingRanges(ctx context.Context, params *FoldingRangeParams) (result []FoldingRange, err error)
	Formatting(ctx context.Context, params *DocumentFormattingParams) (result []TextEdit, err error)
	Hover(ctx context.Context, params *TextDocumentPositionParams) (result *Hover, err error)
	Implementation(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error)
	OnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) (result []TextEdit, err error)
	PrepareRename(ctx context.Context, params *TextDocumentPositionParams) (result *Range, err error)
	RangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) (result []TextEdit, err error)
	References(ctx context.Context, params *ReferenceParams) (result []Location, err error)
	Rename(ctx context.Context, params *RenameParams) (result *WorkspaceEdit, err error)
	SignatureHelp(ctx context.Context, params *TextDocumentPositionParams) (result *SignatureHelp, err error)
	Symbols(ctx context.Context, params *WorkspaceSymbolParams) (result []SymbolInformation, err error)
	TypeDefinition(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error)
	WillSave(ctx context.Context, params *WillSaveTextDocumentParams) (err error)
	WillSaveWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) (result []TextEdit, err error)
}

const (
	// MethodInitialize method name of "initialize".
	MethodInitialize = "initialize"

	// MethodInitialized method name of "initialized".
	MethodInitialized = "initialized"

	// MethodShutdown method name of "shutdown".
	MethodShutdown = "shutdown"

	// MethodExit method name of "exit".
	MethodExit = "exit"

	// MethodCancelRequest method name of "$/cancelRequest".
	MethodCancelRequest = "$/cancelRequest"

	// MethodTextDocumentCodeAction method name of "textDocument/codeAction".
	MethodTextDocumentCodeAction = "textDocument/codeAction"

	// MethodTextDocumentCodeLens method name of "textDocument/codeLens".
	MethodTextDocumentCodeLens = "textDocument/codeLens"

	// MethodCodeLensResolve method name of "codeLens/resolve".
	MethodCodeLensResolve = "codeLens/resolve"

	// MethodTextDocumentColorPresentation method name of "textDocument/colorPresentation".
	MethodTextDocumentColorPresentation = "textDocument/colorPresentation"

	// MethodTextDocumentCompletion method name of "textDocument/completion".
	MethodTextDocumentCompletion = "textDocument/completion"

	// MethodCompletionItemResolve method name of "completionItem/resolve".
	MethodCompletionItemResolve = "completionItem/resolve"

	// MethodTextDocumentDeclaration method name of "textDocument/declaration".
	MethodTextDocumentDeclaration = "textDocument/declaration"

	// MethodTextDocumentDefinition method name of "textDocument/definition".
	MethodTextDocumentDefinition = "textDocument/definition"

	// MethodTextDocumentDidChange method name of "textDocument/didChange".
	MethodTextDocumentDidChange = "textDocument/didChange"

	// MethodWorkspaceDidChangeConfiguration method name of "workspace/didChangeConfiguration".
	MethodWorkspaceDidChangeConfiguration = "workspace/didChangeConfiguration"

	// MethodWorkspaceDidChangeWatchedFiles method name of "workspace/didChangeWatchedFiles".
	MethodWorkspaceDidChangeWatchedFiles = "workspace/didChangeWatchedFiles"

	// MethodWorkspaceDidChangeWorkspaceFolders method name of "workspace/didChangeWorkspaceFolders".
	MethodWorkspaceDidChangeWorkspaceFolders = "workspace/didChangeWorkspaceFolders"

	// MethodTextDocumentDidClose method name of "textDocument/didClose".
	MethodTextDocumentDidClose = "textDocument/didClose"

	// MethodTextDocumentDidOpen method name of "textDocument/didOpen".
	MethodTextDocumentDidOpen = "textDocument/didOpen"

	// MethodTextDocumentDidSave method name of "textDocument/didSave".
	MethodTextDocumentDidSave = "textDocument/didSave"

	// MethodTextDocumentDocumentColor method name of"textDocument/documentColor".
	MethodTextDocumentDocumentColor = "textDocument/documentColor"

	// MethodTextDocumentDocumentHighlight method name of "textDocument/documentHighlight".
	MethodTextDocumentDocumentHighlight = "textDocument/documentHighlight"

	// MethodTextDocumentDocumentLink method name of "textDocument/documentLink".
	MethodTextDocumentDocumentLink = "textDocument/documentLink"

	// MethodDocumentLinkResolve method name of "documentLink/resolve".
	MethodDocumentLinkResolve = "documentLink/resolve"

	// MethodTextDocumentDocumentSymbol method name of "textDocument/documentSymbol".
	MethodTextDocumentDocumentSymbol = "textDocument/documentSymbol"

	// MethodWorkspaceExecuteCommand method name of "workspace/executeCommand".
	MethodWorkspaceExecuteCommand = "workspace/executeCommand"

	// MethodTextDocumentFoldingRange method name of "textDocument/foldingRange".
	MethodTextDocumentFoldingRange = "textDocument/foldingRange"

	// MethodTextDocumentFormatting method name of "textDocument/formatting".
	MethodTextDocumentFormatting = "textDocument/formatting"

	// MethodTextDocumentHover method name of "textDocument/hover".
	MethodTextDocumentHover = "textDocument/hover"

	// MethodTextDocumentImplementation method name of "textDocument/implementation".
	MethodTextDocumentImplementation = "textDocument/implementation"

	// MethodTextDocumentOnTypeFormatting method name of "textDocument/onTypeFormatting".
	MethodTextDocumentOnTypeFormatting = "textDocument/onTypeFormatting"

	// MethodTextDocumentPrepareRename method name of "textDocument/prepareRename".
	MethodTextDocumentPrepareRename = "textDocument/prepareRename"

	// MethodTextDocumentRangeFormatting method name of "textDocument/rangeFormatting".
	MethodTextDocumentRangeFormatting = "textDocument/rangeFormatting"

	// MethodTextDocumentReferences method name of "textDocument/references".
	MethodTextDocumentReferences = "textDocument/references"

	// MethodTextDocumentRename method name of "textDocument/rename".
	MethodTextDocumentRename = "textDocument/rename"

	// MethodTextDocumentSignatureHelp method name of "textDocument/signatureHelp".
	MethodTextDocumentSignatureHelp = "textDocument/signatureHelp"

	// MethodWorkspaceSymbol method name of "workspace/symbol".
	MethodWorkspaceSymbol = "workspace/symbol"

	// MethodTextDocumentTypeDefinition method name of "textDocument/typeDefinition".
	MethodTextDocumentTypeDefinition = "textDocument/typeDefinition"

	// MethodTextDocumentWillSave method name of "textDocument/willSave".
	MethodTextDocumentWillSave = "textDocument/willSave"

	// MethodTextDocumentWillSaveWaitUntil method name of "textDocument/willSaveWaitUntil".
	MethodTextDocumentWillSaveWaitUntil = "textDocument/willSaveWaitUntil"
)

// server implements a Language Server Protocol server.
type server struct {
	jsonrpc2.Conn
	logger *zap.Logger
}

var _ Server = (*server)(nil)

// Run runs the Language Server Protocol server.
// func (s *server) Run(ctx context.Context) (err error) {
// 	err = s.Conn.Run(ctx)
// 	return
// }

// Initialize sents the request as the first request from the client to the server.
//
// If the server receives a request or notification before the initialize request it should act as follows:
//
// - For a request the response should be an error with code: -32002. The message can be picked by the server.
// - Notifications should be dropped, except for the exit notification. This will allow the exit of a server without an initialize request.
//
// Until the server has responded to the initialize request with an InitializeResult, the client
// must not send any additional requests or notifications to the server.
// In addition the server is not allowed to send any requests or notifications to the client until
// it has responded with an InitializeResult, with the exception that during the initialize request
// the server is allowed to send the notifications window/showMessage, window/logMessage and telemetry/event
// as well as the window/showMessageRequest request to the client.
func (s *server) Initialize(ctx context.Context, params *InitializeParams) (result *InitializeResult, err error) {
	s.logger.Debug("call " + MethodInitialize)
	defer s.logger.Debug("end " + MethodInitialize)

	result = new(InitializeResult)
	_, err = s.Conn.Call(ctx, MethodInitialize, params, result)

	return result, err
}

// Initialized sends the notification from the client to the server after the client received the result of the
// initialize request but before the client is sending any other request or notification to the server.
//
// The server can use the initialized notification for example to dynamically register capabilities.
// The initialized notification may only be sent once.
func (s *server) Initialized(ctx context.Context, params *InitializedParams) (err error) {
	s.logger.Debug("notify " + MethodInitialized)
	defer s.logger.Debug("end " + MethodInitialized)

	return s.Conn.Notify(ctx, MethodInitialized, params)
}

// Shutdown sents the request from the client to the server.
//
// It asks the server to shut down, but to not exit (otherwise the response might not be delivered correctly to the client).
// There is a separate exit notification that asks the server to exit.
//
// Clients must not sent any notifications other than `exit` or requests to a server to which they have sent a shutdown requests.
// If a server receives requests after a shutdown request those requests should be errored with `InvalidRequest`.
func (s *server) Shutdown(ctx context.Context) (err error) {
	s.logger.Debug("call " + MethodShutdown)
	defer s.logger.Debug("end " + MethodShutdown)

	_, err = s.Conn.Call(ctx, MethodShutdown, nil, nil)
	return err
}

// Exit a notification to ask the server to exit its process.
//
// The server should exit with success code 0 if the shutdown request has been received before; otherwise with error code 1.
func (s *server) Exit(ctx context.Context) (err error) {
	s.logger.Debug("notify " + MethodExit)
	defer s.logger.Debug("end " + MethodExit)

	return s.Conn.Notify(ctx, MethodExit, nil)
}

// CodeAction sends the request is from the client to the server to compute commands for a given text document and range.
//
// These commands are typically code fixes to either fix problems or to beautify/refactor code. The result of a `textDocument/codeAction`
// request is an array of `Command` literals which are typically presented in the user interface.
//
// To ensure that a server is useful in many clients the commands specified in a code actions should be handled by the
// server and not by the client (see `workspace/executeCommand` and `ServerCapabilities.executeCommandProvider`).
// If the client supports providing edits with a code action then the mode should be used.
func (s *server) CodeAction(ctx context.Context, params *CodeActionParams) (result []CodeAction, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentCodeAction, params, &result)

	return result, err
}

// CodeLens sends the request from the client to the server to compute code lenses for a given text document.
func (s *server) CodeLens(ctx context.Context, params *CodeLensParams) (result []CodeLens, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentCodeLens, params, &result)

	return result, err
}

// CodeLensResolve sends the request from the client to the server to resolve the command for a given code lens item.
func (s *server) CodeLensResolve(ctx context.Context, params *CodeLens) (result *CodeLens, err error) {
	result = new(CodeLens)
	_, err = s.Conn.Call(ctx, MethodCodeLensResolve, params, result)

	return result, err
}

// ColorPresentation sends the request from the client to the server to obtain a list of presentations for a color value at a given location.
//
// Clients can use the result to
//
// - modify a color reference.
// - show in a color picker and let users pick one of the presentations.
func (s *server) ColorPresentation(ctx context.Context, params *ColorPresentationParams) (result []ColorPresentation, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentColorPresentation, params, &result)

	return result, err
}

// Completion sends the request from the client to the server to compute completion items at a given cursor position.
//
// Completion items are presented in the IntelliSense user interface.
// If computing full completion items is expensive, servers can additionally provide a handler for the completion item resolve request (‘completionItem/resolve’).
//
// This request is sent when a completion item is selected in the user interface.
// A typical use case is for example: the ‘textDocument/completion’ request doesn’t fill in the documentation property
// for returned completion items since it is expensive to compute. When the item is selected in the user interface then
// a ‘completionItem/resolve’ request is sent with the selected completion item as a parameter.
//
// The returned completion item should have the documentation property filled in. The request can delay the computation of
// the `detail` and `documentation` properties. However, properties that are needed for the initial sorting and filtering,
// like `sortText`, `filterText`, `insertText`, and `textEdit` must be provided in the `textDocument/completion` response and must not be changed during resolve.
func (s *server) Completion(ctx context.Context, params *CompletionParams) (result *CompletionList, err error) {
	s.logger.Debug("call " + MethodTextDocumentCompletion)
	defer s.logger.Debug("end " + MethodTextDocumentCompletion)

	result = new(CompletionList)
	_, err = s.Conn.Call(ctx, MethodTextDocumentCompletion, params, result)

	return result, err
}

// CompletionResolve sends the request from the client to the server to resolve additional information for a given completion item.
func (s *server) CompletionResolve(ctx context.Context, params *CompletionItem) (result *CompletionItem, err error) {
	result = new(CompletionItem)
	_, err = s.Conn.Call(ctx, MethodCompletionItemResolve, params, result)

	return result, err
}

// Declaration sends the request from the client to the server to resolve the declaration location of a symbol at a given text document position.
//
// The result type LocationLink[] got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.textDocument.declaration.linkSupport`.
//
// Since version 3.14.0.
func (s *server) Declaration(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentDeclaration, params, &result)

	return result, err
}

// Definition sends the request from the client to the server to resolve the definition location of a symbol at a given text document position.
//
// The result type `[]LocationLink` got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.textDocument.definition.linkSupport`.
//
// Since version 3.14.0.
func (s *server) Definition(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error) {
	s.logger.Debug("call " + MethodTextDocumentDefinition)
	defer s.logger.Debug("end " + MethodTextDocumentDefinition)

	_, err = s.Conn.Call(ctx, MethodTextDocumentDefinition, params, &result)

	return result, err
}

// DidChange sends the notification from the client to the server to signal changes to a text document.
//
// In 2.0 the shape of the params has changed to include proper version numbers and language ids.
func (s *server) DidChange(ctx context.Context, params *DidChangeTextDocumentParams) (err error) {
	s.logger.Debug("notify " + MethodTextDocumentDidChange)
	defer s.logger.Debug("end " + MethodTextDocumentDidChange)

	return s.Conn.Notify(ctx, MethodTextDocumentDidChange, params)
}

// DidChangeConfiguration sends the notification from the client to the server to signal the change of configuration settings.
func (s *server) DidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) (err error) {
	return s.Conn.Notify(ctx, MethodWorkspaceDidChangeConfiguration, params)
}

// DidChangeWatchedFiles sends the notification from the client to the server when the client detects changes to files watched by the language client.
//
// It is recommended that servers register for these file events using the registration mechanism.
// In former implementations clients pushed file events without the server actively asking for it.
func (s *server) DidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) (err error) {
	return s.Conn.Notify(ctx, MethodWorkspaceDidChangeWatchedFiles, params)
}

// DidChangeWorkspaceFolders sents the notification from the client to the server to inform the server about workspace folder configuration changes.
//
// The notification is sent by default if both ServerCapabilities/workspace/workspaceFolders and ClientCapabilities/workspace/workspaceFolders are true;
// or if the server has registered itself to receive this notification.
// To register for the workspace/didChangeWorkspaceFolders send a client/registerCapability request from the server to the client.
//
// The registration parameter must have a registrations item of the following form, where id is a unique id used to unregister the capability (the example uses a UUID).
func (s *server) DidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) (err error) {
	return s.Conn.Notify(ctx, MethodWorkspaceDidChangeWorkspaceFolders, params)
}

// DidClose sends the notification from the client to the server when the document got closed in the client.
//
// The document’s truth now exists where the document’s Uri points to (e.g. if the document’s Uri is a file Uri the truth now exists on disk).
// As with the open notification the close notification is about managing the document’s content.
// Receiving a close notification doesn’t mean that the document was open in an editor before.
//
// A close notification requires a previous open notification to be sent.
// Note that a server’s ability to fulfill requests is independent of whether a text document is open or closed.
func (s *server) DidClose(ctx context.Context, params *DidCloseTextDocumentParams) (err error) {
	return s.Conn.Notify(ctx, MethodTextDocumentDidClose, params)
}

// DidOpen sends the open notification from the client to the server to signal newly opened text documents.
//
// The document’s truth is now managed by the client and the server must not try to read the document’s truth using the document’s Uri.
// Open in this sense means it is managed by the client. It doesn’t necessarily mean that its content is presented in an editor.
//
// An open notification must not be sent more than once without a corresponding close notification send before.
// This means open and close notification must be balanced and the max open count for a particular textDocument is one.
// Note that a server’s ability to fulfill requests is independent of whether a text document is open or closed.
func (s *server) DidOpen(ctx context.Context, params *DidOpenTextDocumentParams) (err error) {
	s.logger.Debug("call " + MethodTextDocumentDidOpen)
	defer s.logger.Debug("end " + MethodTextDocumentDidOpen)

	return s.Conn.Notify(ctx, MethodTextDocumentDidOpen, params)
}

// DidSave sends the notification from the client to the server when the document was saved in the client.
func (s *server) DidSave(ctx context.Context, params *DidSaveTextDocumentParams) (err error) {
	return s.Conn.Notify(ctx, MethodTextDocumentDidSave, params)
}

// DocumentColor sends the request from the client to the server to list all color references found in a given text document.
//
// Along with the range, a color value in RGB is returned.
//
// Clients can use the result to decorate color references in an editor.
// For example:
//
// - Color boxes showing the actual color next to the reference
// - Show a color picker when a color reference is edited.
func (s *server) DocumentColor(ctx context.Context, params *DocumentColorParams) (result []ColorInformation, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentDocumentColor, params, &result)

	return result, err
}

// DocumentHighlight sends the request is from the client to the server to resolve a document highlights for a given text document position.
//
// For programming languages this usually highlights all references to the symbol scoped to this file.
// However we kept ‘textDocument/documentHighlight’ and ‘textDocument/references’ separate requests since the first one is allowed to be more fuzzy.
//
// Symbol matches usually have a `DocumentHighlightKind` of `Read` or `Write` whereas fuzzy or textual matches use `Text` as the kind.
func (s *server) DocumentHighlight(ctx context.Context, params *TextDocumentPositionParams) (result []DocumentHighlight, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentDocumentHighlight, params, &result)

	return result, err
}

// DocumentLink sends the request from the client to the server to request the location of links in a document.
func (s *server) DocumentLink(ctx context.Context, params *DocumentLinkParams) (result []DocumentLink, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentDocumentLink, params, &result)

	return result, err
}

// DocumentLinkResolve sends the request from the client to the server to resolve the target of a given document link.
func (s *server) DocumentLinkResolve(ctx context.Context, params *DocumentLink) (result *DocumentLink, err error) {
	result = new(DocumentLink)
	_, err = s.Conn.Call(ctx, MethodDocumentLinkResolve, params, result)

	return result, err
}

// DocumentSymbol sends the request from the client to the server to return a flat list of all symbols found in a given text document.
//
// Neither the symbol’s location range nor the symbol’s container name should be used to infer a hierarchy.
func (s *server) DocumentSymbol(ctx context.Context, params *DocumentSymbolParams) (result []DocumentSymbol, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentDocumentSymbol, params, &result)

	return result, err
}

// ExecuteCommand sends the request from the client to the server to trigger command execution on the server.
//
// In most cases the server creates a `WorkspaceEdit` structure and applies the changes to the workspace using the
// request `workspace/applyEdit` which is sent from the server to the client.
func (s *server) ExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (result interface{}, err error) {
	_, err = s.Conn.Call(ctx, MethodWorkspaceExecuteCommand, params, &result)

	return result, err
}

// FoldingRanges sends the request from the client to the server to return all folding ranges found in a given text document.
//
// Since version 3.10.0.
func (s *server) FoldingRanges(ctx context.Context, params *FoldingRangeParams) (result []FoldingRange, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentFoldingRange, params, &result)

	return result, err
}

// Formatting sends the request from the client to the server to format a whole document.
func (s *server) Formatting(ctx context.Context, params *DocumentFormattingParams) (result []TextEdit, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentFormatting, params, &result)

	return result, err
}

// Hover sends the request is from the client to the server to request hover information at a given text document position.
func (s *server) Hover(ctx context.Context, params *TextDocumentPositionParams) (result *Hover, err error) {
	result = new(Hover)
	_, err = s.Conn.Call(ctx, MethodTextDocumentHover, params, result)

	return result, err
}

// Implementation sends the request from the client to the server to resolve the implementation location of a symbol at a given text document position.
//
// The result type `[]LocationLink` got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.implementation.typeDefinition.linkSupport`.
func (s *server) Implementation(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentImplementation, params, &result)

	return result, err
}

// OnTypeFormatting sends the request from the client to the server to format parts of the document during typing.
func (s *server) OnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) (result []TextEdit, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentOnTypeFormatting, params, &result)

	return result, err
}

// PrepareRename sends the request from the client to the server to setup and test the validity of a rename operation at a given location.
//
// Since version 3.12.0.
func (s *server) PrepareRename(ctx context.Context, params *TextDocumentPositionParams) (result *Range, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentPrepareRename, params, &result)

	return result, err
}

// RangeFormatting sends the request from the client to the server to format a given range in a document.
func (s *server) RangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) (result []TextEdit, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentRangeFormatting, params, &result)

	return result, err
}

// References sends the request from the client to the server to resolve project-wide references for the symbol denoted by the given text document position.
func (s *server) References(ctx context.Context, params *ReferenceParams) (result []Location, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentReferences, params, &result)

	return result, err
}

// Rename sends the request from the client to the server to perform a workspace-wide rename of a symbol.
func (s *server) Rename(ctx context.Context, params *RenameParams) (result *WorkspaceEdit, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentRename, params, &result)

	return result, err
}

// SignatureHelp sends the request from the client to the server to request signature information at a given cursor position.
func (s *server) SignatureHelp(ctx context.Context, params *TextDocumentPositionParams) (result *SignatureHelp, err error) {
	result = new(SignatureHelp)
	_, err = s.Conn.Call(ctx, MethodTextDocumentSignatureHelp, params, result)

	return result, err
}

// Symbols sends the request from the client to the server to list project-wide symbols matching the query string.
func (s *server) Symbols(ctx context.Context, params *WorkspaceSymbolParams) (result []SymbolInformation, err error) {
	_, err = s.Conn.Call(ctx, MethodWorkspaceSymbol, params, &result)

	return result, err
}

// TypeDefinition sends the request from the client to the server to resolve the type definition location of a symbol at a given text document position.
//
// The result type `[]LocationLink` got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.textDocument.typeDefinition.linkSupport`.
//
// Since version 3.6.0.
func (s *server) TypeDefinition(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentTypeDefinition, params, &result)

	return result, err
}

// WillSave sends the notification from the client to the server before the document is actually saved.
func (s *server) WillSave(ctx context.Context, params *WillSaveTextDocumentParams) (err error) {
	return s.Conn.Notify(ctx, MethodTextDocumentWillSave, params)
}

// WillSaveWaitUntil sends the request from the client to the server before the document is actually saved.
//
// The request can return an array of TextEdits which will be applied to the text document before it is saved.
// Please note that clients might drop results if computing the text edits took too long or if a server constantly fails on this request.
// This is done to keep the save fast and reliable.
func (s *server) WillSaveWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) (result []TextEdit, err error) {
	_, err = s.Conn.Call(ctx, MethodTextDocumentWillSaveWaitUntil, params, &result)

	return result, err
}
