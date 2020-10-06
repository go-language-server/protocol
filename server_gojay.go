// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"bytes"
	"context"
	"fmt"

	"github.com/francoispqt/gojay"
	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/pkg/xcontext"

	"go.lsp.dev/protocol/internal/gojaypool"
)

func ServerHandler(server ServerInterface, handler jsonrpc2.Handler) jsonrpc2.Handler {
	h := func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Requester) error {
		if ctx.Err() != nil {
			ctx := xcontext.Detach(ctx)
			return reply(ctx, nil, RequestCancelledError)
		}
		handled, err := serverDispatch(ctx, server, reply, req)
		if handled || err != nil {
			return err
		}

		// TODO: This code is wrong, it ignores handler and assumes non standard
		// request handles everything
		// non standard request should just be a layered handler.
		var params interface{}
		if err := gojay.Unmarshal(req.Params(), &params); err != nil {
			return sendParseError(ctx, reply, err)
		}

		return reply(ctx, nil, err)
	}

	return h
}

//nolint:funlen,gocognit
func serverDispatch(ctx context.Context, server ServerInterface, reply jsonrpc2.Replier, r jsonrpc2.Requester) (bool, error) {
	if ctx.Err() != nil {
		return true, reply(ctx, nil, RequestCancelledError)
	}

	dec := gojaypool.BorrowSizedDecoder(bytes.NewReader(r.Params()), len(r.Params()))
	defer dec.Release()
	logger := LoggerFromContext(ctx)

	switch r.Method() {
	case MethodInitialize: // request
		var params InitializeParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.Initialize(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodInitialize, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodInitialized: // notification
		var params InitializedParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		err := server.Initialized(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodInitialized, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodShutdown: // request
		if len(r.Params()) > 0 {
			return true, reply(ctx, nil, fmt.Errorf("%w: expected no params", jsonrpc2.ErrInvalidParams))
		}

		err := server.Shutdown(ctx)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodShutdown, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodExit: // notification
		if len(r.Params()) > 0 {
			return true, reply(ctx, nil, fmt.Errorf("%w: expected no params", jsonrpc2.ErrInvalidParams))
		}

		err := server.Exit(ctx)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodExit, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodTextDocumentCodeAction: // request
		var params CodeActionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.CodeAction(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentCodeAction, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentCodeLens: // request
		var params CodeLensParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.CodeLens(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentCodeLens, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodCodeLensResolve: // request
		var params CodeLens
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.CodeLensResolve(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodCodeLensResolve, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentColorPresentation: // request
		var params ColorPresentationParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.ColorPresentation(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentColorPresentation, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentCompletion: // request
		var params CompletionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.Completion(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentCompletion, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodCompletionItemResolve: // request
		var params CompletionItem
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.CompletionResolve(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodCompletionItemResolve, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentDeclaration: // request
		var params TextDocumentPositionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.Declaration(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDeclaration, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentDefinition: // request
		var params TextDocumentPositionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.Definition(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDefinition, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentDidChange: // notification
		var params DidChangeTextDocumentParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		err := server.DidChange(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDidChange, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodWorkspaceDidChangeConfiguration: // notification
		var params DidChangeConfigurationParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		err := server.DidChangeConfiguration(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodWorkspaceDidChangeConfiguration, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodWorkspaceDidChangeWatchedFiles: // notification
		var params DidChangeWatchedFilesParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		err := server.DidChangeWatchedFiles(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodWorkspaceDidChangeWatchedFiles, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodWorkspaceDidChangeWorkspaceFolders: // notification
		var params DidChangeWorkspaceFoldersParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		err := server.DidChangeWorkspaceFolders(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodWorkspaceDidChangeWorkspaceFolders, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodTextDocumentDidClose: // notification
		var params DidCloseTextDocumentParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		err := server.DidClose(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDidClose, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodTextDocumentDidOpen: // notification
		var params DidOpenTextDocumentParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		err := server.DidOpen(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDidOpen, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodTextDocumentDidSave: // notification
		var params DidSaveTextDocumentParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		err := server.DidSave(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDidSave, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodTextDocumentDocumentColor: // request
		var params DocumentColorParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.DocumentColor(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDocumentColor, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentDocumentHighlight: // request
		var params TextDocumentPositionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.DocumentHighlight(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDocumentHighlight, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentDocumentLink: // request
		var params DocumentLinkParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.DocumentLink(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDocumentLink, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodDocumentLinkResolve: // request
		var params DocumentLink
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.DocumentLinkResolve(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodDocumentLinkResolve, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentDocumentSymbol: // request
		var params DocumentSymbolParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.DocumentSymbol(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentDocumentSymbol, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodWorkspaceExecuteCommand: // request
		var params ExecuteCommandParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.ExecuteCommand(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodWorkspaceExecuteCommand, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentFoldingRange: // request
		var params FoldingRangeParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.FoldingRanges(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentFoldingRange, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentFormatting: // request
		var params DocumentFormattingParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.Formatting(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentFormatting, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentHover: // request
		var params TextDocumentPositionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.Hover(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentHover, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentImplementation: // request
		var params TextDocumentPositionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.Implementation(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentImplementation, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentOnTypeFormatting: // request
		var params DocumentOnTypeFormattingParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.OnTypeFormatting(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentOnTypeFormatting, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentPrepareRename: // request
		var params TextDocumentPositionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.PrepareRename(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentPrepareRename, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentRangeFormatting: // request
		var params DocumentRangeFormattingParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.RangeFormatting(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentRangeFormatting, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentReferences: // request
		var params ReferenceParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.References(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentReferences, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentRename: // request
		var params RenameParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Rename(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentRename, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentSignatureHelp: // request
		var params TextDocumentPositionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.SignatureHelp(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentSignatureHelp, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodWorkspaceSymbol: // request
		var params WorkspaceSymbolParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.Symbols(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodWorkspaceSymbol, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentTypeDefinition: // request
		var params TextDocumentPositionParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.TypeDefinition(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentTypeDefinition, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodTextDocumentWillSave: // notification
		var params WillSaveTextDocumentParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		err := server.WillSave(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentWillSave, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodTextDocumentWillSaveWaitUntil: // request
		var params WillSaveTextDocumentParams
		if err := dec.DecodeObject(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}

		resp, err := server.WillSaveWaitUntil(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentWillSaveWaitUntil, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	default:
		var params interface{}
		err := dec.Decode(&params)
		return true, sendParseError(ctx, nil, err)
	}
}
