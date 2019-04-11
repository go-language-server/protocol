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

// ServerInterface represents a implementation of language-server-protocol server.
type ServerInterface interface {
	Run(ctx context.Context) error
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
	RangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) (result []TextEdit, err error)
	References(ctx context.Context, params *ReferenceParams) (result []Location, err error)
	Rename(ctx context.Context, params *RenameParams) (result []WorkspaceEdit, err error)
	SignatureHelp(ctx context.Context, params *TextDocumentPositionParams) (result *SignatureHelp, err error)
	Symbols(ctx context.Context, params *WorkspaceSymbolParams) (result []SymbolInformation, err error)
	TypeDefinition(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error)
	WillSave(ctx context.Context, params *WillSaveTextDocumentParams) (err error)
	WillSaveWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) (result []TextEdit, err error)
}

const (
	initialize                         = "initialize"
	initialized                        = "initialized"
	shutdown                           = "shutdown"
	exit                               = "exit"
	cancelRequest                      = "$/cancelRequest"
	textDocumentCodeAction             = "textDocument/codeAction"
	textDocumentCodeLens               = "textDocument/codeLens"
	codeLensResolve                    = "codeLens/resolve"
	textDocumentColorPresentation      = "textDocument/colorPresentation"
	textDocumentCompletion             = "textDocument/completion"
	completionItemResolve              = "completionItem/resolve"
	textDocumentDefinition             = "textDocument/definition"
	textDocumentDidChange              = "textDocument/didChange"
	workspaceDidChangeConfiguration    = "workspace/didChangeConfiguration"
	workspaceDidChangeWatchedFiles     = "workspace/didChangeWatchedFiles"
	workspaceDidChangeWorkspaceFolders = "workspace/didChangeWorkspaceFolders"
	textDocumentDidClose               = "textDocument/didClose"
	textDocumentDidOpen                = "textDocument/didOpen"
	textDocumentDidSave                = "textDocument/didSave"
	textDocumentDocumentColor          = "textDocument/documentColor"
	textDocumentDocumentHighlight      = "textDocument/documentHighlight"
	textDocumentDocumentLink           = "textDocument/documentLink"
	documentLinkResolve                = "documentLink/resolve"
	textDocumentDocumentSymbol         = "textDocument/documentSymbol"
	workspaceExecuteCommand            = "workspace/executeCommand"
	textDocumentFoldingRange           = "textDocument/foldingRange"
	textDocumentFormatting             = "textDocument/formatting"
	textDocumentHover                  = "textDocument/hover"
	textDocumentImplementation         = "textDocument/implementation"
	textDocumentOnTypeFormatting       = "textDocument/onTypeFormatting"
	textDocumentRangeFormatting        = "textDocument/rangeFormatting"
	textDocumentReferences             = "textDocument/references"
	textDocumentRename                 = "textDocument/rename"
	textDocumentSignatureHelp          = "textDocument/signatureHelp"
	workspaceSymbol                    = "workspace/symbol"
	textDocumentTypeDefinition         = "textDocument/typeDefinition"
	textDocumentWillSave               = "textDocument/willSave"
	textDocumentWillSaveWaitUntil      = "textDocument/willSaveWaitUntil"
)

type Server struct {
	*jsonrpc2.Conn
}

var _ ServerInterface = (*Server)(nil)

func (s *Server) Run(ctx context.Context) error {
	return s.Conn.Run(ctx)
}

func (s *Server) Initialize(ctx context.Context, params *InitializeParams) (result *InitializeResult, err error) {
	result = new(InitializeResult)
	if err := s.Conn.Call(ctx, initialize, params, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Initialized(ctx context.Context, params *InitializedParams) (err error) {
	err = s.Conn.Notify(ctx, initialized, params)
	return
}

func (s *Server) Shutdown(ctx context.Context) (err error) {
	err = s.Conn.Call(ctx, shutdown, nil, nil)
	return
}

func (s *Server) Exit(ctx context.Context) (err error) {
	err = s.Conn.Notify(ctx, exit, nil)
	return
}

func (s *Server) CodeAction(ctx context.Context, params *CodeActionParams) (result []CodeAction, err error) {
	if err = s.Conn.Call(ctx, textDocumentCodeAction, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) CodeLens(ctx context.Context, params *CodeLensParams) (result []CodeLens, err error) {
	if err = s.Conn.Call(ctx, textDocumentCodeLens, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) CodeLensResolve(ctx context.Context, params *CodeLens) (result *CodeLens, err error) {
	result = new(CodeLens)
	if err = s.Conn.Call(ctx, codeLensResolve, params, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) ColorPresentation(ctx context.Context, params *ColorPresentationParams) (result []ColorPresentation, err error) {
	if err = s.Conn.Call(ctx, textDocumentColorPresentation, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Completion(ctx context.Context, params *CompletionParams) (result *CompletionList, err error) {
	result = new(CompletionList)
	if err = s.Conn.Call(ctx, textDocumentCompletion, params, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) CompletionResolve(ctx context.Context, params *CompletionItem) (result *CompletionItem, err error) {
	result = new(CompletionItem)
	if err = s.Conn.Call(ctx, completionItemResolve, params, result); err != nil {
		return nil, err

	}

	return result, nil
}

func (s *Server) Definition(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error) {
	if err = s.Conn.Call(ctx, textDocumentDefinition, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) DidChange(ctx context.Context, params *DidChangeTextDocumentParams) (err error) {
	err = s.Conn.Notify(ctx, textDocumentDidChange, params)
	return
}

func (s *Server) DidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) (err error) {
	err = s.Conn.Notify(ctx, workspaceDidChangeConfiguration, params)
	return
}

func (s *Server) DidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) (err error) {
	err = s.Conn.Notify(ctx, workspaceDidChangeWatchedFiles, params)
	return
}

func (s *Server) DidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) (err error) {
	err = s.Conn.Notify(ctx, workspaceDidChangeWorkspaceFolders, params)
	return
}

func (s *Server) DidClose(ctx context.Context, params *DidCloseTextDocumentParams) (err error) {
	err = s.Conn.Notify(ctx, textDocumentDidClose, params)
	return
}

func (s *Server) DidOpen(ctx context.Context, params *DidOpenTextDocumentParams) (err error) {
	err = s.Conn.Notify(ctx, textDocumentDidOpen, params)
	return
}

func (s *Server) DidSave(ctx context.Context, params *DidSaveTextDocumentParams) (err error) {
	err = s.Conn.Notify(ctx, textDocumentDidSave, params)
	return
}

func (s *Server) DocumentColor(ctx context.Context, params *DocumentColorParams) (result []ColorInformation, err error) {
	if err = s.Conn.Call(ctx, textDocumentDocumentColor, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) DocumentHighlight(ctx context.Context, params *TextDocumentPositionParams) (result []DocumentHighlight, err error) {
	if err = s.Conn.Call(ctx, textDocumentDocumentHighlight, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) DocumentLink(ctx context.Context, params *DocumentLinkParams) (result []DocumentLink, err error) {
	if err = s.Conn.Call(ctx, textDocumentDocumentLink, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) DocumentLinkResolve(ctx context.Context, params *DocumentLink) (result *DocumentLink, err error) {
	result = new(DocumentLink)
	if err = s.Conn.Call(ctx, documentLinkResolve, params, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) DocumentSymbol(ctx context.Context, params *DocumentSymbolParams) (result []DocumentSymbol, err error) {
	if err = s.Conn.Call(ctx, textDocumentDocumentSymbol, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) ExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (result interface{}, err error) {
	if err = s.Conn.Call(ctx, workspaceExecuteCommand, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) FoldingRanges(ctx context.Context, params *FoldingRangeParams) (result []FoldingRange, err error) {
	if err = s.Conn.Call(ctx, textDocumentFoldingRange, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Formatting(ctx context.Context, params *DocumentFormattingParams) (result []TextEdit, err error) {
	if err = s.Conn.Call(ctx, textDocumentFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Hover(ctx context.Context, params *TextDocumentPositionParams) (result *Hover, err error) {
	result = new(Hover)
	if err = s.Conn.Call(ctx, textDocumentHover, params, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Implementation(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error) {
	if err = s.Conn.Call(ctx, textDocumentImplementation, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) OnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) (result []TextEdit, err error) {
	if err = s.Conn.Call(ctx, textDocumentOnTypeFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) RangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) (result []TextEdit, err error) {
	if err = s.Conn.Call(ctx, textDocumentRangeFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) References(ctx context.Context, params *ReferenceParams) (result []Location, err error) {
	if err = s.Conn.Call(ctx, textDocumentReferences, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Rename(ctx context.Context, params *RenameParams) (result []WorkspaceEdit, err error) {
	if err = s.Conn.Call(ctx, textDocumentRename, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) SignatureHelp(ctx context.Context, params *TextDocumentPositionParams) (result *SignatureHelp, err error) {
	result = new(SignatureHelp)
	if err := s.Conn.Call(ctx, textDocumentSignatureHelp, params, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Symbols(ctx context.Context, params *WorkspaceSymbolParams) (result []SymbolInformation, err error) {
	if err = s.Conn.Call(ctx, workspaceSymbol, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) TypeDefinition(ctx context.Context, params *TextDocumentPositionParams) (result []Location, err error) {
	if err = s.Conn.Call(ctx, textDocumentTypeDefinition, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) WillSave(ctx context.Context, params *WillSaveTextDocumentParams) (err error) {
	err = s.Conn.Notify(ctx, textDocumentWillSave, params)
	return
}

func (s *Server) WillSaveWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) (result []TextEdit, err error) {
	if err = s.Conn.Call(ctx, textDocumentWillSaveWaitUntil, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ServerHandler returns the client handler.
func ServerHandler(server ServerInterface, logger *zap.Logger) jsonrpc2.Handler {
	return func(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
		dec := gojay.Unsafe

		switch r.Method {
		case initialize:
			var params InitializeParams
			if err := dec.Unmarshal(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Initialize(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(initialize, zap.Error(err))
			}

		case initialized:
			var params InitializedParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
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
			if err := dec.Unmarshal(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			conn.Cancel(params.ID)

		case textDocumentCodeAction:
			var params CodeActionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.CodeAction(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentCodeAction, zap.Error(err))
			}

		case textDocumentCodeLens:
			var params CodeLensParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.CodeLens(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentCodeLens, zap.Error(err))
			}

		case codeLensResolve:
			var params CodeLens
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.CodeLensResolve(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(codeLensResolve, zap.Error(err))
			}

		case textDocumentColorPresentation:
			var params ColorPresentationParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.ColorPresentation(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentColorPresentation, zap.Error(err))
			}

		case textDocumentCompletion:
			var params CompletionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Completion(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentCompletion, zap.Error(err))
			}

		case completionItemResolve:
			var params CompletionItem
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.CompletionResolve(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(completionItemResolve, zap.Error(err))
			}

		case textDocumentDefinition:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Definition(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDefinition, zap.Error(err))
			}

		case textDocumentDidChange:
			var params DidChangeTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidChange(ctx, &params); err != nil {
				logger.Error(textDocumentDidChange, zap.Error(err))
			}

		case workspaceDidChangeConfiguration:
			var params DidChangeConfigurationParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidChangeConfiguration(ctx, &params); err != nil {
				logger.Error(workspaceDidChangeConfiguration, zap.Error(err))
			}

		case workspaceDidChangeWatchedFiles:
			var params DidChangeWatchedFilesParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidChangeWatchedFiles(ctx, &params); err != nil {
				logger.Error(workspaceDidChangeWatchedFiles, zap.Error(err))
			}

		case workspaceDidChangeWorkspaceFolders:
			var params DidChangeWorkspaceFoldersParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidChangeWorkspaceFolders(ctx, &params); err != nil {
				logger.Error(workspaceDidChangeWorkspaceFolders, zap.Error(err))
			}

		case textDocumentDidClose:
			var params DidCloseTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidClose(ctx, &params); err != nil {
				logger.Error(textDocumentDidClose, zap.Error(err))
			}

		case textDocumentDidOpen:
			var params DidOpenTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidOpen(ctx, &params); err != nil {
				logger.Error(textDocumentDidOpen, zap.Error(err))
			}

		case textDocumentDidSave:
			var params DidSaveTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.DidSave(ctx, &params); err != nil {
				logger.Error(textDocumentDidSave, zap.Error(err))
			}

		case textDocumentDocumentColor:
			var params DocumentColorParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentColor(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDocumentColor, zap.Error(err))
			}

		case textDocumentDocumentHighlight:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentHighlight(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDocumentHighlight, zap.Error(err))
			}

		case textDocumentDocumentLink:
			var params DocumentLinkParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentLink(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDocumentLink, zap.Error(err))
			}

		case documentLinkResolve:
			var params DocumentLink
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentLinkResolve(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(documentLinkResolve, zap.Error(err))
			}

		case textDocumentDocumentSymbol:
			var params DocumentSymbolParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.DocumentSymbol(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentDocumentSymbol, zap.Error(err))
			}

		case workspaceExecuteCommand:
			var params ExecuteCommandParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.ExecuteCommand(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(workspaceExecuteCommand, zap.Error(err))
			}

		case textDocumentFoldingRange:
			var params FoldingRangeParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.FoldingRanges(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentFoldingRange, zap.Error(err))
			}

		case textDocumentFormatting:
			var params DocumentFormattingParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Formatting(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentFormatting, zap.Error(err))
			}

		case textDocumentHover:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Hover(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentHover, zap.Error(err))
			}

		case textDocumentImplementation:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Implementation(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentImplementation, zap.Error(err))
			}

		case textDocumentOnTypeFormatting:
			var params DocumentOnTypeFormattingParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.OnTypeFormatting(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentOnTypeFormatting, zap.Error(err))
			}

		case textDocumentRangeFormatting:
			var params DocumentRangeFormattingParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.RangeFormatting(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentRangeFormatting, zap.Error(err))
			}

		case textDocumentReferences:
			var params ReferenceParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.References(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentReferences, zap.Error(err))
			}

		case textDocumentRename:
			var params RenameParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Rename(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentRename, zap.Error(err))
			}

		case textDocumentSignatureHelp:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.SignatureHelp(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentSignatureHelp, zap.Error(err))
			}

		case workspaceSymbol:
			var params WorkspaceSymbolParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.Symbols(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(workspaceSymbol, zap.Error(err))
			}

		case textDocumentTypeDefinition:
			var params TextDocumentPositionParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.TypeDefinition(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentTypeDefinition, zap.Error(err))
			}

		case textDocumentWillSave:
			var params WillSaveTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := server.WillSave(ctx, &params); err != nil {
				logger.Error(textDocumentWillSave, zap.Error(err))
			}

		case textDocumentWillSaveWaitUntil:
			var params WillSaveTextDocumentParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := server.WillSaveWaitUntil(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(textDocumentWillSaveWaitUntil, zap.Error(err))
			}

		default:
			if r.IsNotify() {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeMethodNotFound, "method %q not found", r.Method))
			}
		}
	}
}
