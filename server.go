// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"bytes"
	"context"

	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"

	"github.com/go-language-server/protocol/internal/gojaypool"
)

// ServerInterface represents a Language Server Protocol server.
type ServerInterface interface {
	Run(ctx context.Context) (err error)
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
	Declaration(ctx context.Context, params *TextDocumentPositionParams) (result []LocationLink, err error)
	Definition(ctx context.Context, params *TextDocumentPositionParams) (result []LocationLink, err error)
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
	Implementation(ctx context.Context, params *TextDocumentPositionParams) (result []LocationLink, err error)
	OnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) (result []TextEdit, err error)
	PrepareRename(ctx context.Context, params *TextDocumentPositionParams) (result *Range, err error)
	RangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) (result []TextEdit, err error)
	References(ctx context.Context, params *ReferenceParams) (result []Location, err error)
	Rename(ctx context.Context, params *RenameParams) (result []WorkspaceEdit, err error)
	SignatureHelp(ctx context.Context, params *TextDocumentPositionParams) (result *SignatureHelp, err error)
	Symbols(ctx context.Context, params *WorkspaceSymbolParams) (result []SymbolInformation, err error)
	TypeDefinition(ctx context.Context, params *TextDocumentPositionParams) (result []LocationLink, err error)
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

// Server implements a Language Server Protocol server.
type Server struct {
	*jsonrpc2.Conn
	logger *zap.Logger
}

var _ ServerInterface = (*Server)(nil)

// Run runs the Language Server Protocol server.
func (s *Server) Run(ctx context.Context) (err error) {
	err = s.Conn.Run(ctx)
	return
}

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
func (s *Server) Initialize(ctx context.Context, params *InitializeParams) (result *InitializeResult, err error) {
	s.logger.Debug("call " + MethodInitialize)
	defer s.logger.Debug("end " + MethodInitialize)

	result = new(InitializeResult)
	err = s.Conn.Call(ctx, MethodInitialize, params, result)

	return result, err
}

// Initialized sends the notification from the client to the server after the client received the result of the
// initialize request but before the client is sending any other request or notification to the server.
//
// The server can use the initialized notification for example to dynamically register capabilities.
// The initialized notification may only be sent once.
func (s *Server) Initialized(ctx context.Context, params *InitializedParams) (err error) {
	s.logger.Debug("notify " + MethodInitialized)
	defer s.logger.Debug("end " + MethodInitialized)

	err = s.Conn.Notify(ctx, MethodInitialized, params)
	return err
}

// Shutdown sents the request from the client to the server.
//
// It asks the server to shut down, but to not exit (otherwise the response might not be delivered correctly to the client).
// There is a separate exit notification that asks the server to exit.
//
// Clients must not sent any notifications other than `exit` or requests to a server to which they have sent a shutdown requests.
// If a server receives requests after a shutdown request those requests should be errored with `InvalidRequest`.
func (s *Server) Shutdown(ctx context.Context) (err error) {
	s.logger.Debug("call " + MethodShutdown)
	defer s.logger.Debug("end " + MethodShutdown)

	err = s.Conn.Call(ctx, MethodShutdown, nil, nil)
	return err
}

// Exit a notification to ask the server to exit its process.
//
// The server should exit with success code 0 if the shutdown request has been received before; otherwise with error code 1.
func (s *Server) Exit(ctx context.Context) (err error) {
	s.logger.Debug("notify " + MethodExit)
	defer s.logger.Debug("end " + MethodExit)

	err = s.Conn.Notify(ctx, MethodExit, nil)
	return err
}

// CodeAction sends the request is from the client to the server to compute commands for a given text document and range.
//
// These commands are typically code fixes to either fix problems or to beautify/refactor code. The result of a `textDocument/codeAction`
// request is an array of `Command` literals which are typically presented in the user interface.
//
// To ensure that a server is useful in many clients the commands specified in a code actions should be handled by the
// server and not by the client (see `workspace/executeCommand` and `ServerCapabilities.executeCommandProvider`).
// If the client supports providing edits with a code action then the mode should be used.
func (s *Server) CodeAction(ctx context.Context, params *CodeActionParams) (result []CodeAction, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentCodeAction, params, &result)

	return result, err
}

// CodeLens sends the request from the client to the server to compute code lenses for a given text document.
func (s *Server) CodeLens(ctx context.Context, params *CodeLensParams) (result []CodeLens, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentCodeLens, params, &result)

	return result, err
}

// CodeLensResolve sends the request from the client to the server to resolve the command for a given code lens item.
func (s *Server) CodeLensResolve(ctx context.Context, params *CodeLens) (result *CodeLens, err error) {
	result = new(CodeLens)
	err = s.Conn.Call(ctx, MethodCodeLensResolve, params, result)

	return result, err
}

// ColorPresentation sends the request from the client to the server to obtain a list of presentations for a color value at a given location.
//
// Clients can use the result to
//
// - modify a color reference.
// - show in a color picker and let users pick one of the presentations
func (s *Server) ColorPresentation(ctx context.Context, params *ColorPresentationParams) (result []ColorPresentation, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentColorPresentation, params, &result)

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
func (s *Server) Completion(ctx context.Context, params *CompletionParams) (result *CompletionList, err error) {
	s.logger.Debug("call " + MethodTextDocumentCompletion)
	defer s.logger.Debug("end " + MethodTextDocumentCompletion)

	result = new(CompletionList)
	err = s.Conn.Call(ctx, MethodTextDocumentCompletion, params, result)

	return result, err
}

// CompletionResolve sends the request from the client to the server to resolve additional information for a given completion item.
func (s *Server) CompletionResolve(ctx context.Context, params *CompletionItem) (result *CompletionItem, err error) {
	result = new(CompletionItem)
	err = s.Conn.Call(ctx, MethodCompletionItemResolve, params, result)

	return result, err
}

// Declaration sends the request from the client to the server to resolve the declaration location of a symbol at a given text document position.
//
// The result type LocationLink[] got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.textDocument.declaration.linkSupport`.
//
// Since version 3.14.0.
func (s *Server) Declaration(ctx context.Context, params *TextDocumentPositionParams) (result []LocationLink, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentDeclaration, params, &result)

	return result, err
}

// Definition sends the request from the client to the server to resolve the definition location of a symbol at a given text document position.
//
// The result type `[]LocationLink` got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.textDocument.definition.linkSupport`.
//
// Since version 3.14.0.
func (s *Server) Definition(ctx context.Context, params *TextDocumentPositionParams) (result []LocationLink, err error) {
	s.logger.Debug("call " + MethodTextDocumentDefinition)
	defer s.logger.Debug("end " + MethodTextDocumentDefinition)

	err = s.Conn.Call(ctx, MethodTextDocumentDefinition, params, &result)

	return result, err
}

// DidChange sends the notification from the client to the server to signal changes to a text document.
//
// In 2.0 the shape of the params has changed to include proper version numbers and language ids.
func (s *Server) DidChange(ctx context.Context, params *DidChangeTextDocumentParams) (err error) {
	s.logger.Debug("notify " + MethodTextDocumentDidChange)
	defer s.logger.Debug("end " + MethodTextDocumentDidChange)

	err = s.Conn.Notify(ctx, MethodTextDocumentDidChange, params)
	return
}

// DidChangeConfiguration sends the notification from the client to the server to signal the change of configuration settings.
func (s *Server) DidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) (err error) {
	err = s.Conn.Notify(ctx, MethodWorkspaceDidChangeConfiguration, params)
	return
}

// DidChangeWatchedFiles sends the notification from the client to the server when the client detects changes to files watched by the language client.
//
// It is recommended that servers register for these file events using the registration mechanism.
// In former implementations clients pushed file events without the server actively asking for it.
func (s *Server) DidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) (err error) {
	err = s.Conn.Notify(ctx, MethodWorkspaceDidChangeWatchedFiles, params)
	return
}

// DidChangeWorkspaceFolders sents the notification from the client to the server to inform the server about workspace folder configuration changes.
//
// The notification is sent by default if both ServerCapabilities/workspace/workspaceFolders and ClientCapabilities/workspace/workspaceFolders are true;
// or if the server has registered itself to receive this notification.
// To register for the workspace/didChangeWorkspaceFolders send a client/registerCapability request from the server to the client.
//
// The registration parameter must have a registrations item of the following form, where id is a unique id used to unregister the capability (the example uses a UUID)
func (s *Server) DidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) (err error) {
	err = s.Conn.Notify(ctx, MethodWorkspaceDidChangeWorkspaceFolders, params)
	return
}

// DidClose sends the notification from the client to the server when the document got closed in the client.
//
// The document’s truth now exists where the document’s Uri points to (e.g. if the document’s Uri is a file Uri the truth now exists on disk).
// As with the open notification the close notification is about managing the document’s content.
// Receiving a close notification doesn’t mean that the document was open in an editor before.
//
// A close notification requires a previous open notification to be sent.
// Note that a server’s ability to fulfill requests is independent of whether a text document is open or closed.
func (s *Server) DidClose(ctx context.Context, params *DidCloseTextDocumentParams) (err error) {
	err = s.Conn.Notify(ctx, MethodTextDocumentDidClose, params)
	return
}

// DidOpen sends the open notification from the client to the server to signal newly opened text documents.
//
// The document’s truth is now managed by the client and the server must not try to read the document’s truth using the document’s Uri.
// Open in this sense means it is managed by the client. It doesn’t necessarily mean that its content is presented in an editor.
//
// An open notification must not be sent more than once without a corresponding close notification send before.
// This means open and close notification must be balanced and the max open count for a particular textDocument is one.
// Note that a server’s ability to fulfill requests is independent of whether a text document is open or closed.
func (s *Server) DidOpen(ctx context.Context, params *DidOpenTextDocumentParams) (err error) {
	s.logger.Debug("call " + MethodTextDocumentDidOpen)
	defer s.logger.Debug("end " + MethodTextDocumentDidOpen)

	err = s.Conn.Notify(ctx, MethodTextDocumentDidOpen, params)
	return
}

// DidSave sends the notification from the client to the server when the document was saved in the client.
func (s *Server) DidSave(ctx context.Context, params *DidSaveTextDocumentParams) (err error) {
	err = s.Conn.Notify(ctx, MethodTextDocumentDidSave, params)
	return
}

// DocumentColor sends the request from the client to the server to list all color references found in a given text document.
//
// Along with the range, a color value in RGB is returned.
//
// Clients can use the result to decorate color references in an editor.
// For example:
//
// - Color boxes showing the actual color next to the reference
// - Show a color picker when a color reference is edited
func (s *Server) DocumentColor(ctx context.Context, params *DocumentColorParams) (result []ColorInformation, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentDocumentColor, params, &result)

	return result, err
}

// DocumentHighlight sends the request is from the client to the server to resolve a document highlights for a given text document position.
//
// For programming languages this usually highlights all references to the symbol scoped to this file.
// However we kept ‘textDocument/documentHighlight’ and ‘textDocument/references’ separate requests since the first one is allowed to be more fuzzy.
//
// Symbol matches usually have a `DocumentHighlightKind` of `Read` or `Write` whereas fuzzy or textual matches use `Text` as the kind.
func (s *Server) DocumentHighlight(ctx context.Context, params *TextDocumentPositionParams) (result []DocumentHighlight, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentDocumentHighlight, params, &result)

	return result, err
}

// DocumentLink sends the request from the client to the server to request the location of links in a document.
func (s *Server) DocumentLink(ctx context.Context, params *DocumentLinkParams) (result []DocumentLink, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentDocumentLink, params, &result)

	return result, err
}

// DocumentLinkResolve sends the request from the client to the server to resolve the target of a given document link.
func (s *Server) DocumentLinkResolve(ctx context.Context, params *DocumentLink) (result *DocumentLink, err error) {
	result = new(DocumentLink)
	err = s.Conn.Call(ctx, MethodDocumentLinkResolve, params, result)

	return result, err
}

// DocumentSymbol sends the request from the client to the server to return a flat list of all symbols found in a given text document.
//
// Neither the symbol’s location range nor the symbol’s container name should be used to infer a hierarchy.
func (s *Server) DocumentSymbol(ctx context.Context, params *DocumentSymbolParams) (result []DocumentSymbol, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentDocumentSymbol, params, &result)

	return result, err
}

// ExecuteCommand sends the request from the client to the server to trigger command execution on the server.
//
// In most cases the server creates a `WorkspaceEdit` structure and applies the changes to the workspace using the
// request `workspace/applyEdit` which is sent from the server to the client.
func (s *Server) ExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (result interface{}, err error) {
	err = s.Conn.Call(ctx, MethodWorkspaceExecuteCommand, params, &result)

	return result, err
}

// FoldingRanges sends the request from the client to the server to return all folding ranges found in a given text document.
//
// Since version 3.10.0.
func (s *Server) FoldingRanges(ctx context.Context, params *FoldingRangeParams) (result []FoldingRange, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentFoldingRange, params, &result)

	return result, err
}

// Formatting sends the request from the client to the server to format a whole document.
func (s *Server) Formatting(ctx context.Context, params *DocumentFormattingParams) (result []TextEdit, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentFormatting, params, &result)

	return result, err
}

// Hover sends the request is from the client to the server to request hover information at a given text document position.
func (s *Server) Hover(ctx context.Context, params *TextDocumentPositionParams) (result *Hover, err error) {
	result = new(Hover)
	err = s.Conn.Call(ctx, MethodTextDocumentHover, params, result)

	return result, err
}

// Implementation sends the request from the client to the server to resolve the implementation location of a symbol at a given text document position.
//
// The result type `[]LocationLink` got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.implementation.typeDefinition.linkSupport`.
func (s *Server) Implementation(ctx context.Context, params *TextDocumentPositionParams) (result []LocationLink, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentImplementation, params, &result)

	return result, err
}

// OnTypeFormatting sends the request from the client to the server to format parts of the document during typing.
func (s *Server) OnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) (result []TextEdit, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentOnTypeFormatting, params, &result)

	return result, err
}

// PrepareRename sends the request from the client to the server to setup and test the validity of a rename operation at a given location.
//
// Since version 3.12.0.
func (s *Server) PrepareRename(ctx context.Context, params *TextDocumentPositionParams) (result *Range, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentPrepareRename, params, &result)

	return result, err
}

// RangeFormatting sends the request from the client to the server to format a given range in a document.
func (s *Server) RangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) (result []TextEdit, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentRangeFormatting, params, &result)

	return result, err
}

// References sends the request from the client to the server to resolve project-wide references for the symbol denoted by the given text document position.
func (s *Server) References(ctx context.Context, params *ReferenceParams) (result []Location, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentReferences, params, &result)

	return result, err
}

// Rename sends the request from the client to the server to perform a workspace-wide rename of a symbol.
func (s *Server) Rename(ctx context.Context, params *RenameParams) (result []WorkspaceEdit, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentRename, params, &result)

	return result, err
}

// SignatureHelp sends the request from the client to the server to request signature information at a given cursor position.
func (s *Server) SignatureHelp(ctx context.Context, params *TextDocumentPositionParams) (result *SignatureHelp, err error) {
	result = new(SignatureHelp)
	err = s.Conn.Call(ctx, MethodTextDocumentSignatureHelp, params, result)

	return result, err
}

// Symbols sends the request from the client to the server to list project-wide symbols matching the query string.
func (s *Server) Symbols(ctx context.Context, params *WorkspaceSymbolParams) (result []SymbolInformation, err error) {
	err = s.Conn.Call(ctx, MethodWorkspaceSymbol, params, &result)

	return result, err
}

// TypeDefinition sends the request from the client to the server to resolve the type definition location of a symbol at a given text document position.
//
// The result type `[]LocationLink` got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.textDocument.typeDefinition.linkSupport`.
//
// Since version 3.6.0.
func (s *Server) TypeDefinition(ctx context.Context, params *TextDocumentPositionParams) (result []LocationLink, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentTypeDefinition, params, &result)

	return result, err
}

// WillSave sends the notification from the client to the server before the document is actually saved.
func (s *Server) WillSave(ctx context.Context, params *WillSaveTextDocumentParams) (err error) {
	err = s.Conn.Notify(ctx, MethodTextDocumentWillSave, params)
	return
}

// WillSaveWaitUntil sends the request from the client to the server before the document is actually saved.
//
// The request can return an array of TextEdits which will be applied to the text document before it is saved.
// Please note that clients might drop results if computing the text edits took too long or if a server constantly fails on this request.
// This is done to keep the save fast and reliable.
func (s *Server) WillSaveWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) (result []TextEdit, err error) {
	err = s.Conn.Call(ctx, MethodTextDocumentWillSaveWaitUntil, params, &result)

	return result, err
}

// ServerHandler returns the client handler.
func ServerHandler(ctx context.Context, server ServerInterface, logger *zap.Logger) jsonrpc2.Handler {
	return func(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
		dec := gojaypool.BorrowSizedDecoder(bytes.NewReader(*r.Params), len(*r.Params))
		defer dec.Release()

		switch r.Method {
		case MethodInitialize:
			var params InitializeParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Initialize(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodInitialize, zap.Error(err))
			}

		case MethodInitialized:
			var params InitializedParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.Initialized(ctx, &params); err != nil {
				logger.Error(MethodInitialized, zap.Error(err))
			}

		case MethodShutdown:
			if r.Params != nil {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeInvalidParams, "Expected no params"))
				return
			}
			if err := server.Shutdown(ctx); err != nil {
				logger.Error(MethodShutdown, zap.Error(err))
			}

		case MethodExit:
			if r.Params != nil {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeInvalidParams, "Expected no params"))
				return
			}
			if err := server.Exit(ctx); err != nil {
				logger.Error(MethodExit, zap.Error(err))
			}

		case MethodCancelRequest:
			var params CancelParams
			if err := dec.Decode(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			conn.Cancel(params.ID)

		case MethodTextDocumentCodeAction:
			var params CodeActionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.CodeAction(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentCodeAction, zap.Error(err))
			}

		case MethodTextDocumentCodeLens:
			var params CodeLensParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.CodeLens(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentCodeLens, zap.Error(err))
			}

		case MethodCodeLensResolve:
			var params CodeLens
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.CodeLensResolve(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodCodeLensResolve, zap.Error(err))
			}

		case MethodTextDocumentColorPresentation:
			var params ColorPresentationParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.ColorPresentation(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentColorPresentation, zap.Error(err))
			}

		case MethodTextDocumentCompletion:
			var params CompletionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Completion(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentCompletion, zap.Error(err))
			}

		case MethodCompletionItemResolve:
			var params CompletionItem
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.CompletionResolve(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodCompletionItemResolve, zap.Error(err))
			}

		case MethodTextDocumentDeclaration:
			var params TextDocumentPositionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Declaration(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentDeclaration, zap.Error(err))
			}

		case MethodTextDocumentDefinition:
			var params TextDocumentPositionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Definition(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentDefinition, zap.Error(err))
			}

		case MethodTextDocumentDidChange:
			var params DidChangeTextDocumentParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidChange(ctx, &params); err != nil {
				logger.Error(MethodTextDocumentDidChange, zap.Error(err))
			}

		case MethodWorkspaceDidChangeConfiguration:
			var params DidChangeConfigurationParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidChangeConfiguration(ctx, &params); err != nil {
				logger.Error(MethodWorkspaceDidChangeConfiguration, zap.Error(err))
			}

		case MethodWorkspaceDidChangeWatchedFiles:
			var params DidChangeWatchedFilesParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidChangeWatchedFiles(ctx, &params); err != nil {
				logger.Error(MethodWorkspaceDidChangeWatchedFiles, zap.Error(err))
			}

		case MethodWorkspaceDidChangeWorkspaceFolders:
			var params DidChangeWorkspaceFoldersParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidChangeWorkspaceFolders(ctx, &params); err != nil {
				logger.Error(MethodWorkspaceDidChangeWorkspaceFolders, zap.Error(err))
			}

		case MethodTextDocumentDidClose:
			var params DidCloseTextDocumentParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidClose(ctx, &params); err != nil {
				logger.Error(MethodTextDocumentDidClose, zap.Error(err))
			}

		case MethodTextDocumentDidOpen:
			var params DidOpenTextDocumentParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidOpen(ctx, &params); err != nil {
				logger.Error(MethodTextDocumentDidOpen, zap.Error(err))
			}

		case MethodTextDocumentDidSave:
			var params DidSaveTextDocumentParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidSave(ctx, &params); err != nil {
				logger.Error(MethodTextDocumentDidSave, zap.Error(err))
			}

		case MethodTextDocumentDocumentColor:
			var params DocumentColorParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentColor(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentDocumentColor, zap.Error(err))
			}

		case MethodTextDocumentDocumentHighlight:
			var params TextDocumentPositionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentHighlight(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentDocumentHighlight, zap.Error(err))
			}

		case MethodTextDocumentDocumentLink:
			var params DocumentLinkParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentLink(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentDocumentLink, zap.Error(err))
			}

		case MethodDocumentLinkResolve:
			var params DocumentLink
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentLinkResolve(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodDocumentLinkResolve, zap.Error(err))
			}

		case MethodTextDocumentDocumentSymbol:
			var params DocumentSymbolParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentSymbol(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentDocumentSymbol, zap.Error(err))
			}

		case MethodWorkspaceExecuteCommand:
			var params ExecuteCommandParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.ExecuteCommand(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodWorkspaceExecuteCommand, zap.Error(err))
			}

		case MethodTextDocumentFoldingRange:
			var params FoldingRangeParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.FoldingRanges(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentFoldingRange, zap.Error(err))
			}

		case MethodTextDocumentFormatting:
			var params DocumentFormattingParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Formatting(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentFormatting, zap.Error(err))
			}

		case MethodTextDocumentHover:
			var params TextDocumentPositionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Hover(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentHover, zap.Error(err))
			}

		case MethodTextDocumentImplementation:
			var params TextDocumentPositionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Implementation(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentImplementation, zap.Error(err))
			}

		case MethodTextDocumentOnTypeFormatting:
			var params DocumentOnTypeFormattingParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.OnTypeFormatting(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentOnTypeFormatting, zap.Error(err))
			}

		case MethodTextDocumentPrepareRename:
			var params TextDocumentPositionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.PrepareRename(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentPrepareRename, zap.Error(err))
			}

		case MethodTextDocumentRangeFormatting:
			var params DocumentRangeFormattingParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.RangeFormatting(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentRangeFormatting, zap.Error(err))
			}

		case MethodTextDocumentReferences:
			var params ReferenceParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.References(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentReferences, zap.Error(err))
			}

		case MethodTextDocumentRename:
			var params RenameParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Rename(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentRename, zap.Error(err))
			}

		case MethodTextDocumentSignatureHelp:
			var params TextDocumentPositionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.SignatureHelp(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentSignatureHelp, zap.Error(err))
			}

		case MethodWorkspaceSymbol:
			var params WorkspaceSymbolParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Symbols(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodWorkspaceSymbol, zap.Error(err))
			}

		case MethodTextDocumentTypeDefinition:
			var params TextDocumentPositionParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.TypeDefinition(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentTypeDefinition, zap.Error(err))
			}

		case MethodTextDocumentWillSave:
			var params WillSaveTextDocumentParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.WillSave(ctx, &params); err != nil {
				logger.Error(MethodTextDocumentWillSave, zap.Error(err))
			}

		case MethodTextDocumentWillSaveWaitUntil:
			var params WillSaveTextDocumentParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.WillSaveWaitUntil(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodTextDocumentWillSaveWaitUntil, zap.Error(err))
			}

		default:
			if r.IsNotify() {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeMethodNotFound, "method %q not found", r.Method))
			}
		}
	}
}
