// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"bytes"
	"context"

	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"

	"github.com/go-language-server/protocol/internal/gojaypool"
)

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
