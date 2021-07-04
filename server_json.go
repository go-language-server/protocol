// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
)

// serverDispatch dispatches server.
//nolint:gocognit,funlen,gocyclo,cyclop
func serverDispatch(ctx context.Context, server Server, req *jsonrpc2.Request) (resp interface{}, err error) {
	if ctx.Err() != nil {
		return nil, ErrRequestCancelled
	}

	dec := json.NewDecoder(bytes.NewReader(req.Params))
	logger := LoggerFromContext(ctx)

	switch req.Method {
	case MethodInitialize: // request
		defer logger.Debug(MethodInitialize, zap.Error(err))

		var params InitializeParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodInitialize, err)
		}

		resp, err = server.Initialize(ctx, &params)

		return resp, err

	case MethodShutdown: // request
		defer logger.Debug(MethodShutdown, zap.Error(err))

		if len(req.Params) > 0 {
			return nil, fmt.Errorf("%s: expected no params: %w", MethodShutdown, jsonrpc2.ErrInvalidParams)
		}

		err = server.Shutdown(ctx)

		return nil, err

	case MethodTextDocumentCodeAction: // request
		defer logger.Debug(MethodTextDocumentCodeAction, zap.Error(err))

		var params CodeActionParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentCodeAction, err)
		}

		resp, err = server.CodeAction(ctx, &params)

		return resp, err

	case MethodTextDocumentCodeLens: // request
		defer logger.Debug(MethodTextDocumentCodeLens, zap.Error(err))

		var params CodeLensParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentCodeLens, err)
		}

		resp, err = server.CodeLens(ctx, &params)

		return resp, err

	case MethodCodeLensResolve: // request
		defer logger.Debug(MethodCodeLensResolve, zap.Error(err))

		var params CodeLens
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodCodeLensResolve, err)
		}

		resp, err = server.CodeLensResolve(ctx, &params)

		return resp, err

	case MethodTextDocumentColorPresentation: // request
		defer logger.Debug(MethodTextDocumentColorPresentation, zap.Error(err))

		var params ColorPresentationParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentColorPresentation, err)
		}

		resp, err = server.ColorPresentation(ctx, &params)

		return resp, err

	case MethodTextDocumentCompletion: // request
		defer logger.Debug(MethodTextDocumentCompletion, zap.Error(err))

		var params CompletionParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentCompletion, err)
		}

		resp, err = server.Completion(ctx, &params)

		return resp, err

	case MethodCompletionItemResolve: // request
		defer logger.Debug(MethodCompletionItemResolve, zap.Error(err))

		var params CompletionItem
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodCompletionItemResolve, err)
		}

		resp, err = server.CompletionResolve(ctx, &params)

		return resp, err

	case MethodTextDocumentDeclaration: // request
		defer logger.Debug(MethodTextDocumentDeclaration, zap.Error(err))

		var params DeclarationParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentDeclaration, err)
		}

		resp, err = server.Declaration(ctx, &params)

		return resp, err

	case MethodTextDocumentDefinition: // request
		defer logger.Debug(MethodTextDocumentDefinition, zap.Error(err))

		var params DefinitionParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentDefinition, err)
		}

		resp, err = server.Definition(ctx, &params)

		return resp, err

	case MethodTextDocumentDocumentColor: // request
		defer logger.Debug(MethodTextDocumentDocumentColor, zap.Error(err))

		var params DocumentColorParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentDocumentColor, err)
		}

		resp, err = server.DocumentColor(ctx, &params)

		return resp, err

	case MethodTextDocumentDocumentHighlight: // request
		defer logger.Debug(MethodTextDocumentDocumentHighlight, zap.Error(err))

		var params DocumentHighlightParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentDocumentHighlight, err)
		}

		resp, err = server.DocumentHighlight(ctx, &params)

		return resp, err

	case MethodTextDocumentDocumentLink: // request
		defer logger.Debug(MethodTextDocumentDocumentLink, zap.Error(err))

		var params DocumentLinkParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentDocumentLink, err)
		}

		resp, err = server.DocumentLink(ctx, &params)

		return resp, err

	case MethodDocumentLinkResolve: // request
		defer logger.Debug(MethodDocumentLinkResolve, zap.Error(err))

		var params DocumentLink
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodDocumentLinkResolve, err)
		}

		resp, err = server.DocumentLinkResolve(ctx, &params)

		return resp, err

	case MethodTextDocumentDocumentSymbol: // request
		defer logger.Debug(MethodTextDocumentDocumentSymbol, zap.Error(err))

		var params DocumentSymbolParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentDocumentSymbol, err)
		}

		resp, err = server.DocumentSymbol(ctx, &params)

		return resp, err

	case MethodWorkspaceExecuteCommand: // request
		defer logger.Debug(MethodWorkspaceExecuteCommand, zap.Error(err))

		var params ExecuteCommandParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodWorkspaceExecuteCommand, err)
		}

		resp, err = server.ExecuteCommand(ctx, &params)

		return resp, err

	case MethodTextDocumentFoldingRange: // request
		defer logger.Debug(MethodTextDocumentFoldingRange, zap.Error(err))

		var params FoldingRangeParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentFoldingRange, err)
		}

		resp, err = server.FoldingRanges(ctx, &params)

		return resp, err

	case MethodTextDocumentFormatting: // request
		defer logger.Debug(MethodTextDocumentFormatting, zap.Error(err))

		var params DocumentFormattingParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentFormatting, err)
		}

		resp, err = server.Formatting(ctx, &params)

		return resp, err

	case MethodTextDocumentHover: // request
		defer logger.Debug(MethodTextDocumentHover, zap.Error(err))

		var params HoverParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentHover, err)
		}

		resp, err = server.Hover(ctx, &params)

		return resp, err

	case MethodTextDocumentImplementation: // request
		defer logger.Debug(MethodTextDocumentImplementation, zap.Error(err))

		var params ImplementationParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentImplementation, err)
		}

		resp, err = server.Implementation(ctx, &params)

		return resp, err

	case MethodTextDocumentOnTypeFormatting: // request
		defer logger.Debug(MethodTextDocumentOnTypeFormatting, zap.Error(err))

		var params DocumentOnTypeFormattingParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentOnTypeFormatting, err)
		}

		resp, err = server.OnTypeFormatting(ctx, &params)

		return resp, err

	case MethodTextDocumentPrepareRename: // request
		defer logger.Debug(MethodTextDocumentPrepareRename, zap.Error(err))

		var params PrepareRenameParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentPrepareRename, err)
		}

		resp, err = server.PrepareRename(ctx, &params)

		return resp, err

	case MethodTextDocumentRangeFormatting: // request
		defer logger.Debug(MethodTextDocumentRangeFormatting, zap.Error(err))

		var params DocumentRangeFormattingParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentRangeFormatting, err)
		}

		resp, err = server.RangeFormatting(ctx, &params)

		return resp, err

	case MethodTextDocumentReferences: // request
		defer logger.Debug(MethodTextDocumentReferences, zap.Error(err))

		var params ReferenceParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentReferences, err)
		}

		resp, err = server.References(ctx, &params)

		return resp, err

	case MethodTextDocumentRename: // request
		defer logger.Debug(MethodTextDocumentRename, zap.Error(err))

		var params RenameParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentRename, err)
		}

		resp, err = server.Rename(ctx, &params)

		return resp, err

	case MethodTextDocumentSignatureHelp: // request
		defer logger.Debug(MethodTextDocumentSignatureHelp, zap.Error(err))

		var params SignatureHelpParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentSignatureHelp, err)
		}

		resp, err = server.SignatureHelp(ctx, &params)

		return resp, err

	case MethodWorkspaceSymbol: // request
		defer logger.Debug(MethodWorkspaceSymbol, zap.Error(err))

		var params WorkspaceSymbolParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodWorkspaceSymbol, err)
		}

		resp, err = server.Symbols(ctx, &params)

		return resp, err

	case MethodTextDocumentTypeDefinition: // request
		defer logger.Debug(MethodTextDocumentTypeDefinition, zap.Error(err))

		var params TypeDefinitionParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentTypeDefinition, err)
		}

		resp, err = server.TypeDefinition(ctx, &params)

		return resp, err

	case MethodTextDocumentWillSaveWaitUntil: // request
		defer logger.Debug(MethodTextDocumentWillSaveWaitUntil, zap.Error(err))

		var params WillSaveTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentWillSaveWaitUntil, err)
		}

		resp, err = server.WillSaveWaitUntil(ctx, &params)

		return resp, err

	case MethodShowDocument: // request
		defer logger.Debug(MethodShowDocument, zap.Error(err))

		var params ShowDocumentParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodShowDocument, err)
		}

		resp, err := server.ShowDocument(ctx, &params)

		return resp, err

	case MethodWillCreateFiles: // request
		defer logger.Debug(MethodWillCreateFiles, zap.Error(err))

		var params CreateFilesParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodWillCreateFiles, err)
		}

		resp, err = server.WillCreateFiles(ctx, &params)

		return resp, err

	case MethodCodeLensRefresh: // request
		defer logger.Debug(MethodCodeLensRefresh, zap.Error(err))

		if len(req.Params) > 0 {
			return nil, fmt.Errorf("%s: expected no params: %w", MethodCodeLensRefresh, jsonrpc2.ErrInvalidParams)
		}

		err = server.CodeLensRefresh(ctx)

		return nil, err

	case MethodTextDocumentPrepareCallHierarchy: // request
		defer logger.Debug(MethodTextDocumentPrepareCallHierarchy, zap.Error(err))

		var params CallHierarchyPrepareParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentPrepareCallHierarchy, err)
		}

		resp, err = server.PrepareCallHierarchy(ctx, &params)

		return resp, err

	case MethodCallHierarchyIncomingCalls: // request
		defer logger.Debug(MethodCallHierarchyIncomingCalls, zap.Error(err))

		var params CallHierarchyIncomingCallsParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodCallHierarchyIncomingCalls, err)
		}

		resp, err = server.IncomingCalls(ctx, &params)

		return resp, err

	case MethodCallHierarchyOutgoingCalls: // request
		defer logger.Debug(MethodCallHierarchyOutgoingCalls, zap.Error(err))

		var params CallHierarchyOutgoingCallsParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodCallHierarchyOutgoingCalls, err)
		}

		resp, err = server.OutgoingCalls(ctx, &params)

		return resp, err

	case MethodSemanticTokensFull: // request
		defer logger.Debug(MethodSemanticTokensFull, zap.Error(err))

		var params SemanticTokensParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodSemanticTokensFull, err)
		}

		resp, err = server.SemanticTokensFull(ctx, &params)

		return resp, err

	case MethodSemanticTokensFullDelta: // request
		defer logger.Debug(MethodSemanticTokensFullDelta, zap.Error(err))

		var params SemanticTokensDeltaParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodSemanticTokensFullDelta, err)
		}

		resp, err = server.SemanticTokensFullDelta(ctx, &params)

		return resp, err

	case MethodSemanticTokensRange: // request
		defer logger.Debug(MethodSemanticTokensRange, zap.Error(err))

		var params SemanticTokensRangeParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodSemanticTokensRange, err)
		}

		resp, err = server.SemanticTokensRange(ctx, &params)

		return resp, err

	case MethodSemanticTokensRefresh: // request
		defer logger.Debug(MethodSemanticTokensRefresh, zap.Error(err))

		if len(req.Params) > 0 {
			return nil, fmt.Errorf("%s: expected no params: %w", MethodSemanticTokensRefresh, jsonrpc2.ErrInvalidParams)
		}

		err = server.SemanticTokensRefresh(ctx)

		return nil, err

	case MethodLinkedEditingRange: // request
		defer logger.Debug(MethodLinkedEditingRange, zap.Error(err))

		var params LinkedEditingRangeParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodLinkedEditingRange, err)
		}

		resp, err = server.LinkedEditingRange(ctx, &params)

		return resp, err

	case MethodMoniker: // request
		defer logger.Debug(MethodMoniker, zap.Error(err))

		var params MonikerParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodMoniker, err)
		}

		resp, err = server.Moniker(ctx, &params)

		return resp, err

	default:
		return nil, errors.New("unknown method")
	}
}
