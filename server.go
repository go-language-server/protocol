// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"github.com/francoispqt/gojay"
	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"
)

// Server represents a implementation of language-server-protocol server.
type Server interface {
	Initialize(context.Context, *InitializeParams) (*InitializeResult, error)
	Initialized(context.Context, *InitializedParams) error
	Shutdown(context.Context) error
	Exit(context.Context) error
	DidChangeWorkspaceFolders(context.Context, *DidChangeWorkspaceFoldersParams) error
	DidChangeConfiguration(context.Context, *DidChangeConfigurationParams) error
	DidChangeWatchedFiles(context.Context, *DidChangeWatchedFilesParams) error
	Symbols(context.Context, *WorkspaceSymbolParams) ([]SymbolInformation, error)
	ExecuteCommand(context.Context, *ExecuteCommandParams) (interface{}, error)
	DidOpen(context.Context, *DidOpenTextDocumentParams) error
	DidChange(context.Context, *DidChangeTextDocumentParams) error
	WillSave(context.Context, *WillSaveTextDocumentParams) error
	WillSaveWaitUntil(context.Context, *WillSaveTextDocumentParams) ([]TextEdit, error)
	DidSave(context.Context, *DidSaveTextDocumentParams) error
	DidClose(context.Context, *DidCloseTextDocumentParams) error
	Completion(context.Context, *CompletionParams) (*CompletionList, error)
	CompletionResolve(context.Context, *CompletionItem) (*CompletionItem, error)
	Hover(context.Context, *TextDocumentPositionParams) (*Hover, error)
	SignatureHelp(context.Context, *TextDocumentPositionParams) (*SignatureHelp, error)
	Definition(context.Context, *TextDocumentPositionParams) ([]Location, error)
	TypeDefinition(context.Context, *TextDocumentPositionParams) ([]Location, error)
	Implementation(context.Context, *TextDocumentPositionParams) ([]Location, error)
	References(context.Context, *ReferenceParams) ([]Location, error)
	DocumentHighlight(context.Context, *TextDocumentPositionParams) ([]DocumentHighlight, error)
	DocumentSymbol(context.Context, *DocumentSymbolParams) ([]DocumentSymbol, error)
	CodeAction(context.Context, *CodeActionParams) ([]CodeAction, error)
	CodeLens(context.Context, *CodeLensParams) ([]CodeLens, error)
	CodeLensResolve(context.Context, *CodeLens) (*CodeLens, error)
	DocumentLink(context.Context, *DocumentLinkParams) ([]DocumentLink, error)
	DocumentLinkResolve(context.Context, *DocumentLink) (*DocumentLink, error)
	DocumentColor(context.Context, *DocumentColorParams) ([]ColorInformation, error)
	ColorPresentation(context.Context, *ColorPresentationParams) ([]ColorPresentation, error)
	Formatting(context.Context, *DocumentFormattingParams) ([]TextEdit, error)
	RangeFormatting(context.Context, *DocumentRangeFormattingParams) ([]TextEdit, error)
	OnTypeFormatting(context.Context, *DocumentOnTypeFormattingParams) ([]TextEdit, error)
	Rename(context.Context, *RenameParams) ([]WorkspaceEdit, error)
	FoldingRanges(context.Context, *FoldingRangeParams) ([]FoldingRange, error)
}

const (
	initialize                         = "initialize"
	initialized                        = "initialized"
	shutdown                           = "shutdown"
	exit                               = "exit"
	cancelRequest                      = "$/cancelRequest"
	workspaceDidChangeWorkspaceFolders = "workspace/didChangeWorkspaceFolders"
	workspaceDidChangeConfiguration    = "workspace/didChangeConfiguration"
	workspaceDidChangeWatchedFiles     = "workspace/didChangeWatchedFiles"
	workspaceSymbol                    = "workspace/symbol"
	workspaceExecuteCommand            = "workspace/executeCommand"
	textDocumentDidOpen                = "textDocument/didOpen"
	textDocumentDidChange              = "textDocument/didChange"
	textDocumentWillSave               = "textDocument/willSave"
	textDocumentWillSaveWaitUntil      = "textDocument/willSaveWaitUntil"
	textDocumentDidSave                = "textDocument/didSave"
	textDocumentDidClose               = "textDocument/didClose"
	textDocumentCompletion             = "textDocument/completion"
	completionItemResolve              = "completionItem/resolve"
	textDocumentHover                  = "textDocument/hover"
	textDocumentSignatureHelp          = "textDocument/signatureHelp"
	textDocumentDefinition             = "textDocument/definition"
	textDocumentTypeDefinition         = "textDocument/typeDefinition"
	textDocumentImplementation         = "textDocument/implementation"
	textDocumentReferences             = "textDocument/references"
	textDocumentDocumentHighlight      = "textDocument/documentHighlight"
	textDocumentDocumentSymbol         = "textDocument/documentSymbol"
	textDocumentCodeAction             = "textDocument/codeAction"
	textDocumentCodeLens               = "textDocument/codeLens"
	codeLensResolve                    = "codeLens/resolve"
	textDocumentDocumentLink           = "textDocument/documentLink"
	documentLinkResolve                = "documentLink/resolve"
	textDocumentDocumentColor          = "textDocument/documentColor"
	textDocumentColorPresentation      = "textDocument/colorPresentation"
	textDocumentFormatting             = "textDocument/formatting"
	textDocumentRangeFormatting        = "textDocument/rangeFormatting"
	textDocumentOnTypeFormatting       = "textDocument/onTypeFormatting"
	textDocumentRename                 = "textDocument/rename"
	textDocumentFoldingRange           = "textDocument/foldingRange"
)

func sendParseError(ctx context.Context, logger *zap.Logger, conn *jsonrpc2.Conn, req *jsonrpc2.Request, err error) {
	if _, ok := err.(*jsonrpc2.Error); !ok {
		err = jsonrpc2.Errorf(jsonrpc2.CodeParseError, "%v", err)
	}
	if err := conn.Reply(ctx, req, nil, err); err != nil {
		logger.Error("sendParseError", zap.Error(err))
	}
}

func serverHandler(server Server, logger *zap.Logger) jsonrpc2.Handler {
	return func(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
		dec := gojay.Unsafe

		switch r.Method {
		case initialize:
			var params InitializeParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.Initialize(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(initialize, zap.Error(err))
			}

		case initialized:
			var params InitializedParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			if err := server.Initialized(ctx, &params); err != nil {
				logger.Error(initialized, zap.Error(err))
			}

		case shutdown:
			if r.Params != nil {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeInvalidParams, "Expected no params"))
				return
			}
			if err := server.Shutdown(ctx); err != nil {
				logger.Error(shutdown, zap.Error(err))
			}

		case exit:
			if r.Params != nil {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeInvalidParams, "Expected no params"))
				return
			}
			if err := server.Exit(ctx); err != nil {
				logger.Error(exit, zap.Error(err))
			}

		case cancelRequest:
			var params CancelParams
			if err := dec.Unmarshal(*r.Params, params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			conn.Cancel(params.ID)

		case workspaceDidChangeWorkspaceFolders:
			var params DidChangeWorkspaceFoldersParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			if err := server.DidChangeWorkspaceFolders(ctx, &params); err != nil {
				logger.Error(workspaceDidChangeWorkspaceFolders, zap.Error(err))
			}

		case workspaceDidChangeConfiguration:
			var params DidChangeConfigurationParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			if err := server.DidChangeConfiguration(ctx, &params); err != nil {
				logger.Error(workspaceDidChangeConfiguration, zap.Error(err))
			}

		case workspaceDidChangeWatchedFiles:
			var params DidChangeWatchedFilesParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			if err := server.DidChangeWatchedFiles(ctx, &params); err != nil {
				logger.Error(workspaceDidChangeWatchedFiles, zap.Error(err))
			}

		case workspaceSymbol:
			var params WorkspaceSymbolParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.Symbols(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(workspaceSymbol, zap.Error(err))
			}

		case workspaceExecuteCommand:
			var params ExecuteCommandParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.ExecuteCommand(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(workspaceExecuteCommand, zap.Error(err))
			}

		case textDocumentDidOpen:
			var params DidOpenTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			if err := server.DidOpen(ctx, &params); err != nil {
				logger.Error(textDocumentDidOpen, zap.Error(err))
			}

		case textDocumentDidChange:
			var params DidChangeTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			if err := server.DidChange(ctx, &params); err != nil {
				logger.Error(textDocumentDidChange, zap.Error(err))
			}

		case textDocumentWillSave:
			var params WillSaveTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			if err := server.WillSave(ctx, &params); err != nil {
				logger.Error(textDocumentWillSave, zap.Error(err))
			}

		case textDocumentWillSaveWaitUntil:
			var params WillSaveTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.WillSaveWaitUntil(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentWillSaveWaitUntil, zap.Error(err))
			}

		case textDocumentDidSave:
			var params DidSaveTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			if err := server.DidSave(ctx, &params); err != nil {
				logger.Error(textDocumentDidSave, zap.Error(err))
			}

		case textDocumentDidClose:
			var params DidCloseTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			if err := server.DidClose(ctx, &params); err != nil {
				logger.Error(textDocumentDidClose, zap.Error(err))
			}

		case textDocumentCompletion:
			var params CompletionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.Completion(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentCompletion, zap.Error(err))
			}

		case completionItemResolve:
			var params CompletionItem
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.CompletionResolve(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(completionItemResolve, zap.Error(err))
			}

		case textDocumentHover:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.Hover(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentHover, zap.Error(err))
			}

		case textDocumentSignatureHelp:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.SignatureHelp(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentSignatureHelp, zap.Error(err))
			}

		case textDocumentDefinition:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.Definition(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDefinition, zap.Error(err))
			}

		case textDocumentTypeDefinition:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.TypeDefinition(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentTypeDefinition, zap.Error(err))
			}

		case textDocumentImplementation:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.Implementation(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentImplementation, zap.Error(err))
			}

		case textDocumentReferences:
			var params ReferenceParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.References(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentReferences, zap.Error(err))
			}

		case textDocumentDocumentHighlight:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.DocumentHighlight(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDocumentHighlight, zap.Error(err))
			}

		case textDocumentDocumentSymbol:
			var params DocumentSymbolParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.DocumentSymbol(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDocumentSymbol, zap.Error(err))
			}

		case textDocumentCodeAction:
			var params CodeActionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.CodeAction(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentCodeAction, zap.Error(err))
			}

		case textDocumentCodeLens:
			var params CodeLensParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.CodeLens(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentCodeLens, zap.Error(err))
			}

		case codeLensResolve:
			var params CodeLens
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.CodeLensResolve(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(codeLensResolve, zap.Error(err))
			}

		case textDocumentDocumentLink:
			var params DocumentLinkParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.DocumentLink(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDocumentLink, zap.Error(err))
			}

		case documentLinkResolve:
			var params DocumentLink
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.DocumentLinkResolve(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(documentLinkResolve, zap.Error(err))
			}

		case textDocumentDocumentColor:
			var params DocumentColorParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.DocumentColor(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDocumentColor, zap.Error(err))
			}

		case textDocumentColorPresentation:
			var params ColorPresentationParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.ColorPresentation(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentColorPresentation, zap.Error(err))
			}

		case textDocumentFormatting:
			var params DocumentFormattingParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.Formatting(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentFormatting, zap.Error(err))
			}

		case textDocumentRangeFormatting:
			var params DocumentRangeFormattingParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.RangeFormatting(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentRangeFormatting, zap.Error(err))
			}

		case textDocumentOnTypeFormatting:
			var params DocumentOnTypeFormattingParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.OnTypeFormatting(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentOnTypeFormatting, zap.Error(err))
			}

		case textDocumentRename:
			var params RenameParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.Rename(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentRename, zap.Error(err))
			}

		case textDocumentFoldingRange:
			var params FoldingRangeParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				sendParseError(ctx, logger, conn, r, err)
				return
			}
			resp, err := server.FoldingRanges(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentFoldingRange, zap.Error(err))
			}

		default:
			if r.IsNotify() {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeMethodNotFound, "method %q not found", r.Method))
			}
		}
	}
}

type serverDispatcher struct {
	*jsonrpc2.Conn
}

func (s *serverDispatcher) Initialize(ctx context.Context, params *InitializeParams) (*InitializeResult, error) {
	var result InitializeResult
	if err := s.Conn.Call(ctx, initialize, params, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *serverDispatcher) Initialized(ctx context.Context, params *InitializedParams) error {
	return s.Conn.Notify(ctx, initialized, params)
}

func (s *serverDispatcher) Shutdown(ctx context.Context) error {
	return s.Conn.Call(ctx, shutdown, nil, nil)
}

func (s *serverDispatcher) Exit(ctx context.Context) error {
	return s.Conn.Notify(ctx, exit, nil)
}

func (s *serverDispatcher) DidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) error {
	return s.Conn.Notify(ctx, workspaceDidChangeWorkspaceFolders, params)
}

func (s *serverDispatcher) DidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) error {
	return s.Conn.Notify(ctx, workspaceDidChangeConfiguration, params)
}

func (s *serverDispatcher) DidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) error {
	return s.Conn.Notify(ctx, workspaceDidChangeWatchedFiles, params)
}

func (s *serverDispatcher) Symbols(ctx context.Context, params *WorkspaceSymbolParams) ([]SymbolInformation, error) {
	var result []SymbolInformation
	if err := s.Conn.Call(ctx, workspaceSymbol, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) ExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (interface{}, error) {
	var result interface{}
	if err := s.Conn.Call(ctx, workspaceExecuteCommand, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) DidOpen(ctx context.Context, params *DidOpenTextDocumentParams) error {
	return s.Conn.Notify(ctx, textDocumentDidOpen, params)
}

func (s *serverDispatcher) DidChange(ctx context.Context, params *DidChangeTextDocumentParams) error {
	return s.Conn.Notify(ctx, textDocumentDidChange, params)
}

func (s *serverDispatcher) WillSave(ctx context.Context, params *WillSaveTextDocumentParams) error {
	return s.Conn.Notify(ctx, textDocumentWillSave, params)
}

func (s *serverDispatcher) WillSaveWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) ([]TextEdit, error) {
	var result []TextEdit
	if err := s.Conn.Call(ctx, textDocumentWillSaveWaitUntil, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) DidSave(ctx context.Context, params *DidSaveTextDocumentParams) error {
	return s.Conn.Notify(ctx, textDocumentDidSave, params)
}

func (s *serverDispatcher) DidClose(ctx context.Context, params *DidCloseTextDocumentParams) error {
	return s.Conn.Notify(ctx, textDocumentDidClose, params)
}

func (s *serverDispatcher) Completion(ctx context.Context, params *CompletionParams) (*CompletionList, error) {
	var result CompletionList
	if err := s.Conn.Call(ctx, textDocumentCompletion, params, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *serverDispatcher) CompletionResolve(ctx context.Context, params *CompletionItem) (*CompletionItem, error) {
	var result CompletionItem
	if err := s.Conn.Call(ctx, completionItemResolve, params, &result); err != nil {
		return nil, err

	}
	return &result, nil
}

func (s *serverDispatcher) Hover(ctx context.Context, params *TextDocumentPositionParams) (*Hover, error) {
	var result Hover
	if err := s.Conn.Call(ctx, textDocumentHover, params, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *serverDispatcher) SignatureHelp(ctx context.Context, params *TextDocumentPositionParams) (*SignatureHelp, error) {
	var result SignatureHelp
	if err := s.Conn.Call(ctx, textDocumentSignatureHelp, params, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *serverDispatcher) Definition(ctx context.Context, params *TextDocumentPositionParams) ([]Location, error) {
	var result []Location
	if err := s.Conn.Call(ctx, textDocumentDefinition, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) TypeDefinition(ctx context.Context, params *TextDocumentPositionParams) ([]Location, error) {
	var result []Location
	if err := s.Conn.Call(ctx, textDocumentTypeDefinition, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) Implementation(ctx context.Context, params *TextDocumentPositionParams) ([]Location, error) {
	var result []Location
	if err := s.Conn.Call(ctx, textDocumentImplementation, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) References(ctx context.Context, params *ReferenceParams) ([]Location, error) {
	var result []Location
	if err := s.Conn.Call(ctx, textDocumentReferences, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) DocumentHighlight(ctx context.Context, params *TextDocumentPositionParams) ([]DocumentHighlight, error) {
	var result []DocumentHighlight
	if err := s.Conn.Call(ctx, textDocumentDocumentHighlight, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) DocumentSymbol(ctx context.Context, params *DocumentSymbolParams) ([]DocumentSymbol, error) {
	var result []DocumentSymbol
	if err := s.Conn.Call(ctx, textDocumentDocumentSymbol, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) CodeAction(ctx context.Context, params *CodeActionParams) ([]CodeAction, error) {
	var result []CodeAction
	if err := s.Conn.Call(ctx, textDocumentCodeAction, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) CodeLens(ctx context.Context, params *CodeLensParams) ([]CodeLens, error) {
	var result []CodeLens
	if err := s.Conn.Call(ctx, textDocumentCodeLens, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) CodeLensResolve(ctx context.Context, params *CodeLens) (*CodeLens, error) {
	var result CodeLens
	if err := s.Conn.Call(ctx, codeLensResolve, params, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *serverDispatcher) DocumentLink(ctx context.Context, params *DocumentLinkParams) ([]DocumentLink, error) {
	var result []DocumentLink
	if err := s.Conn.Call(ctx, textDocumentDocumentLink, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) DocumentLinkResolve(ctx context.Context, params *DocumentLink) (*DocumentLink, error) {
	var result DocumentLink
	if err := s.Conn.Call(ctx, documentLinkResolve, params, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *serverDispatcher) DocumentColor(ctx context.Context, params *DocumentColorParams) ([]ColorInformation, error) {
	var result []ColorInformation
	if err := s.Conn.Call(ctx, textDocumentDocumentColor, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) ColorPresentation(ctx context.Context, params *ColorPresentationParams) ([]ColorPresentation, error) {
	var result []ColorPresentation
	if err := s.Conn.Call(ctx, textDocumentColorPresentation, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) Formatting(ctx context.Context, params *DocumentFormattingParams) ([]TextEdit, error) {
	var result []TextEdit
	if err := s.Conn.Call(ctx, textDocumentFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) RangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) ([]TextEdit, error) {
	var result []TextEdit
	if err := s.Conn.Call(ctx, textDocumentRangeFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) OnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) ([]TextEdit, error) {
	var result []TextEdit
	if err := s.Conn.Call(ctx, textDocumentOnTypeFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) Rename(ctx context.Context, params *RenameParams) ([]WorkspaceEdit, error) {
	var result []WorkspaceEdit
	if err := s.Conn.Call(ctx, textDocumentRename, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *serverDispatcher) FoldingRanges(ctx context.Context, params *FoldingRangeParams) ([]FoldingRange, error) {
	var result []FoldingRange
	if err := s.Conn.Call(ctx, textDocumentFoldingRange, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}
