// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !gojay

package protocol

import (
	"bytes"
	"context"
	"encoding/json"

	"go.lsp.dev/jsonrpc2"
	"go.uber.org/zap"
)

// Deliver implements jsonrpc2.Handler.
func (h serverHandler) Deliver(ctx context.Context, r *jsonrpc2.Request, delivered bool) bool {
	if delivered {
		return false
	}

	if ctx.Err() != nil {
		r.Reply(ctx, nil, jsonrpc2.NewError(RequestCancelledError, ctx.Err().Error()))
		return true
	}

	dec := json.NewDecoder(bytes.NewReader(*r.Params))
	logger := LoggerFromContext(ctx)

	switch r.Method {
	case MethodInitialize: // request
		var params InitializeParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.Initialize(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodInitialize, zap.Error(err))
		}
		return true

	case MethodInitialized: // notification
		var params InitializedParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.server.Initialized(ctx, &params); err != nil {
			logger.Error(MethodInitialized, zap.Error(err))
		}
		return true

	case MethodShutdown: // request
		if r.Params != nil {
			if err := r.Reply(ctx, nil, jsonrpc2.Errorf(jsonrpc2.InvalidParams, "Expected no params")); err != nil {
				logger.Error(MethodShutdown, zap.Error(err))
			}
			return true
		}
		if err := h.server.Shutdown(ctx); err != nil {
			logger.Error(MethodShutdown, zap.Error(err))
		}
		return true

	case MethodExit: // notification
		if r.Params != nil {
			if err := r.Reply(ctx, nil, jsonrpc2.Errorf(jsonrpc2.InvalidParams, "Expected no params")); err != nil {
				logger.Error(MethodExit, zap.Error(err))
			}
			return true
		}
		if err := h.server.Exit(ctx); err != nil {
			logger.Error(MethodExit, zap.Error(err))
		}
		return true

	case MethodTextDocumentCodeAction: // request
		var params CodeActionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.CodeAction(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentCodeAction, zap.Error(err))
		}
		return true

	case MethodTextDocumentCodeLens: // request
		var params CodeLensParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.CodeLens(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentCodeLens, zap.Error(err))
		}
		return true

	case MethodCodeLensResolve: // request
		var params CodeLens
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.CodeLensResolve(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodCodeLensResolve, zap.Error(err))
		}
		return true

	case MethodTextDocumentColorPresentation: // request
		var params ColorPresentationParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.ColorPresentation(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentColorPresentation, zap.Error(err))
		}
		return true

	case MethodTextDocumentCompletion: // request
		var params CompletionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.Completion(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentCompletion, zap.Error(err))
		}
		return true

	case MethodCompletionItemResolve: // request
		var params CompletionItem
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.CompletionResolve(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodCompletionItemResolve, zap.Error(err))
		}
		return true

	case MethodTextDocumentDeclaration: // request
		var params TextDocumentPositionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.Declaration(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentDeclaration, zap.Error(err))
		}
		return true

	case MethodTextDocumentDefinition: // request
		var params TextDocumentPositionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.Definition(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentDefinition, zap.Error(err))
		}
		return true

	case MethodTextDocumentDidChange: // notification
		var params DidChangeTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.server.DidChange(ctx, &params); err != nil {
			logger.Error(MethodTextDocumentDidChange, zap.Error(err))
		}
		return true

	case MethodWorkspaceDidChangeConfiguration: // notification
		var params DidChangeConfigurationParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.server.DidChangeConfiguration(ctx, &params); err != nil {
			logger.Error(MethodWorkspaceDidChangeConfiguration, zap.Error(err))
		}
		return true

	case MethodWorkspaceDidChangeWatchedFiles: // notification
		var params DidChangeWatchedFilesParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.server.DidChangeWatchedFiles(ctx, &params); err != nil {
			logger.Error(MethodWorkspaceDidChangeWatchedFiles, zap.Error(err))
		}
		return true

	case MethodWorkspaceDidChangeWorkspaceFolders: // notification
		var params DidChangeWorkspaceFoldersParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.server.DidChangeWorkspaceFolders(ctx, &params); err != nil {
			logger.Error(MethodWorkspaceDidChangeWorkspaceFolders, zap.Error(err))
		}
		return true

	case MethodTextDocumentDidClose: // notification
		var params DidCloseTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.server.DidClose(ctx, &params); err != nil {
			logger.Error(MethodTextDocumentDidClose, zap.Error(err))
		}
		return true

	case MethodTextDocumentDidOpen: // notification
		var params DidOpenTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.server.DidOpen(ctx, &params); err != nil {
			logger.Error(MethodTextDocumentDidOpen, zap.Error(err))
		}
		return true

	case MethodTextDocumentDidSave: // notification
		var params DidSaveTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.server.DidSave(ctx, &params); err != nil {
			logger.Error(MethodTextDocumentDidSave, zap.Error(err))
		}
		return true

	case MethodTextDocumentDocumentColor: // request
		var params DocumentColorParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.DocumentColor(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentDocumentColor, zap.Error(err))
		}
		return true

	case MethodTextDocumentDocumentHighlight: // request
		var params TextDocumentPositionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.DocumentHighlight(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentDocumentHighlight, zap.Error(err))
		}
		return true

	case MethodTextDocumentDocumentLink: // request
		var params DocumentLinkParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.DocumentLink(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentDocumentLink, zap.Error(err))
		}
		return true

	case MethodDocumentLinkResolve: // request
		var params DocumentLink
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.DocumentLinkResolve(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodDocumentLinkResolve, zap.Error(err))
		}
		return true

	case MethodTextDocumentDocumentSymbol: // request
		var params DocumentSymbolParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.DocumentSymbol(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentDocumentSymbol, zap.Error(err))
		}
		return true

	case MethodWorkspaceExecuteCommand: // request
		var params ExecuteCommandParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.ExecuteCommand(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodWorkspaceExecuteCommand, zap.Error(err))
		}
		return true

	case MethodTextDocumentFoldingRange: // request
		var params FoldingRangeParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.FoldingRanges(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentFoldingRange, zap.Error(err))
		}
		return true

	case MethodTextDocumentFormatting: // request
		var params DocumentFormattingParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.Formatting(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentFormatting, zap.Error(err))
		}
		return true

	case MethodTextDocumentHover: // request
		var params TextDocumentPositionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.Hover(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentHover, zap.Error(err))
		}
		return true

	case MethodTextDocumentImplementation: // request
		var params TextDocumentPositionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.Implementation(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentImplementation, zap.Error(err))
		}
		return true

	case MethodTextDocumentOnTypeFormatting: // request
		var params DocumentOnTypeFormattingParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.OnTypeFormatting(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentOnTypeFormatting, zap.Error(err))
		}
		return true

	case MethodTextDocumentPrepareRename: // request
		var params TextDocumentPositionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.PrepareRename(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentPrepareRename, zap.Error(err))
		}
		return true

	case MethodTextDocumentRangeFormatting: // request
		var params DocumentRangeFormattingParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.RangeFormatting(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentRangeFormatting, zap.Error(err))
		}
		return true

	case MethodTextDocumentReferences: // request
		var params ReferenceParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.References(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentReferences, zap.Error(err))
		}
		return true

	case MethodTextDocumentRename: // request
		var params RenameParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.Rename(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentRename, zap.Error(err))
		}
		return true

	case MethodTextDocumentSignatureHelp: // request
		var params TextDocumentPositionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.SignatureHelp(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentSignatureHelp, zap.Error(err))
		}
		return true

	case MethodWorkspaceSymbol: // request
		var params WorkspaceSymbolParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.Symbols(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodWorkspaceSymbol, zap.Error(err))
		}
		return true

	case MethodTextDocumentTypeDefinition: // request
		var params TextDocumentPositionParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.TypeDefinition(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentTypeDefinition, zap.Error(err))
		}
		return true

	case MethodTextDocumentWillSave: // notification
		var params WillSaveTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.server.WillSave(ctx, &params); err != nil {
			logger.Error(MethodTextDocumentWillSave, zap.Error(err))
		}
		return true

	case MethodTextDocumentWillSaveWaitUntil: // request
		var params WillSaveTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.server.WillSaveWaitUntil(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodTextDocumentWillSaveWaitUntil, zap.Error(err))
		}
		return true

	default:
		var params interface{}
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		return true
	}
}
