// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"context"
	"fmt"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
)

// ServerDispatcher returns a Server that dispatches LSP requests across the
// given jsonrpc2 connection.
func ServerDispatcher(conn jsonrpc2.Conn, logger *zap.Logger) Server {
	return &server{
		Conn:   conn,
		logger: logger,
	}
}

// ServerHandler jsonrpc2.Handler of Language Server Prococol Server.
//
//nolint:unparam,revive
func ServerHandler(server Server, handler jsonrpc2.Handler) jsonrpc2.Handler {
	h := func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		if ctx.Err() != nil {
			xctx := context.WithoutCancel(ctx)

			return reply(xctx, nil, ErrRequestCancelled)
		}
		handled, err := serverDispatch(ctx, server, reply, req)
		if handled || err != nil {
			return err
		}

		// TODO: This code is wrong, it ignores handler and assumes non standard
		// request handles everything
		// non standard request should just be a layered handler.
		var params any
		if err := unmarshal(req.Params(), &params); err != nil {
			return replyParseError(ctx, reply, err)
		}

		resp, err := server.Request(ctx, req.Method(), params)

		return reply(ctx, resp, err)
	}

	return h
}

// serverDispatch implements jsonrpc2.Handler.
//
//nolint:gocognit,funlen,gocyclo,cyclop,maintidx
func serverDispatch(ctx context.Context, server Server, reply jsonrpc2.Replier, req jsonrpc2.Request) (handled bool, err error) {
	if ctx.Err() != nil {
		return true, reply(ctx, nil, ErrRequestCancelled)
	}

	dec := newDecoder(bytes.NewReader(req.Params()))
	logger := LoggerFromContext(ctx)

	switch req.Method() {
	case MethodServerProgress: // notification
		defer logger.Debug(MethodServerProgress, zap.Error(err))

		var params ProgressParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.Progress(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodSetTrace: // notification
		defer logger.Debug(MethodSetTrace, zap.Error(err))

		var params SetTraceParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.SetTrace(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodExit: // notification
		defer logger.Debug(MethodExit, zap.Error(err))

		if len(req.Params()) > 0 {
			return true, reply(ctx, nil, fmt.Errorf("expected no params: %w", jsonrpc2.ErrInvalidParams))
		}

		err := server.Exit(ctx)

		return true, reply(ctx, nil, err)

	case MethodInitialized: // notification
		defer logger.Debug(MethodInitialized, zap.Error(err))

		var params InitializedParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.Initialized(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodNotebookDocumentDidChange: // notification
		defer logger.Debug(MethodNotebookDocumentDidChange, zap.Error(err))

		var params DidChangeNotebookDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.NotebookDocumentDidChange(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodNotebookDocumentDidClose: // notification
		defer logger.Debug(MethodNotebookDocumentDidClose, zap.Error(err))

		var params DidCloseNotebookDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.NotebookDocumentDidClose(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodNotebookDocumentDidOpen: // notification
		defer logger.Debug(MethodNotebookDocumentDidOpen, zap.Error(err))

		var params DidOpenNotebookDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.NotebookDocumentDidOpen(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodNotebookDocumentDidSave: // notification
		defer logger.Debug(MethodNotebookDocumentDidSave, zap.Error(err))

		var params DidSaveNotebookDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.NotebookDocumentDidSave(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodTextDocumentDidChange: // notification
		defer logger.Debug(MethodTextDocumentDidChange, zap.Error(err))

		var params DidChangeTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.TextDocumentDidChange(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodTextDocumentDidClose: // notification
		defer logger.Debug(MethodTextDocumentDidClose, zap.Error(err))

		var params DidCloseTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.TextDocumentDidClose(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodTextDocumentDidOpen: // notification
		defer logger.Debug(MethodTextDocumentDidOpen, zap.Error(err))

		var params DidOpenTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.TextDocumentDidOpen(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodTextDocumentDidSave: // notification
		defer logger.Debug(MethodTextDocumentDidSave, zap.Error(err))

		var params DidSaveTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.TextDocumentDidSave(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodTextDocumentWillSave: // notification
		defer logger.Debug(MethodTextDocumentWillSave, zap.Error(err))

		var params WillSaveTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.TextDocumentWillSave(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWindowWorkDoneProgressCancel: // notification
		defer logger.Debug(MethodWindowWorkDoneProgressCancel, zap.Error(err))

		var params WorkDoneProgressCancelParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.WindowWorkDoneProgressCancel(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceDidChangeConfiguration: // notification
		defer logger.Debug(MethodWorkspaceDidChangeConfiguration, zap.Error(err))

		var params DidChangeConfigurationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.WorkspaceDidChangeConfiguration(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceDidChangeWatchedFiles: // notification
		defer logger.Debug(MethodWorkspaceDidChangeWatchedFiles, zap.Error(err))

		var params DidChangeWatchedFilesParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.WorkspaceDidChangeWatchedFiles(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceDidChangeWorkspaceFolders: // notification
		defer logger.Debug(MethodWorkspaceDidChangeWorkspaceFolders, zap.Error(err))

		var params DidChangeWorkspaceFoldersParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.WorkspaceDidChangeWorkspaceFolders(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceDidCreateFiles: // notification
		defer logger.Debug(MethodWorkspaceDidCreateFiles, zap.Error(err))

		var params CreateFilesParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.WorkspaceDidCreateFiles(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceDidDeleteFiles: // notification
		defer logger.Debug(MethodWorkspaceDidDeleteFiles, zap.Error(err))

		var params DeleteFilesParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.WorkspaceDidDeleteFiles(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceDidRenameFiles: // notification
		defer logger.Debug(MethodWorkspaceDidRenameFiles, zap.Error(err))

		var params RenameFilesParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := server.WorkspaceDidRenameFiles(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodCallHierarchyIncomingCalls: // request
		defer logger.Debug(MethodCallHierarchyIncomingCalls, zap.Error(err))

		var params CallHierarchyIncomingCallsParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.CallHierarchyIncomingCalls(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodCallHierarchyOutgoingCalls: // request
		defer logger.Debug(MethodCallHierarchyOutgoingCalls, zap.Error(err))

		var params CallHierarchyOutgoingCallsParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.CallHierarchyOutgoingCalls(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodCodeActionResolve: // request
		defer logger.Debug(MethodCodeActionResolve, zap.Error(err))

		var params CodeAction
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.CodeActionResolve(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodCodeLensResolve: // request
		defer logger.Debug(MethodCodeLensResolve, zap.Error(err))

		var params CodeLens
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.CodeLensResolve(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodCompletionItemResolve: // request
		defer logger.Debug(MethodCompletionItemResolve, zap.Error(err))

		var params CompletionItem
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.CompletionItemResolve(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodDocumentLinkResolve: // request
		defer logger.Debug(MethodDocumentLinkResolve, zap.Error(err))

		var params DocumentLink
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.DocumentLinkResolve(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodInitialize: // request
		defer logger.Debug(MethodInitialize, zap.Error(err))

		var params InitializeParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.Initialize(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodInlayHintResolve: // request
		defer logger.Debug(MethodInlayHintResolve, zap.Error(err))

		var params InlayHint
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.InlayHintResolve(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodShutdown: // request
		defer logger.Debug(MethodShutdown, zap.Error(err))

		if len(req.Params()) > 0 {
			return true, reply(ctx, nil, fmt.Errorf("expected no params: %w", jsonrpc2.ErrInvalidParams))
		}

		err := server.Shutdown(ctx)

		return true, reply(ctx, nil, err)

	case MethodTextDocumentCodeAction: // request
		defer logger.Debug(MethodTextDocumentCodeAction, zap.Error(err))

		var params CodeActionParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentCodeAction(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentCodeLens: // request
		defer logger.Debug(MethodTextDocumentCodeLens, zap.Error(err))

		var params CodeLensParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentCodeLens(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentColorPresentation: // request
		defer logger.Debug(MethodTextDocumentColorPresentation, zap.Error(err))

		var params ColorPresentationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentColorPresentation(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentCompletion: // request
		defer logger.Debug(MethodTextDocumentCompletion, zap.Error(err))

		var params CompletionParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentCompletion(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentDeclaration: // request
		defer logger.Debug(MethodTextDocumentDeclaration, zap.Error(err))

		var params DeclarationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentDeclaration(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentDefinition: // request
		defer logger.Debug(MethodTextDocumentDefinition, zap.Error(err))

		var params DefinitionParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentDefinition(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentDiagnostic: // request
		defer logger.Debug(MethodTextDocumentDiagnostic, zap.Error(err))

		var params DocumentDiagnosticParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentDiagnostic(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentDocumentColor: // request
		defer logger.Debug(MethodTextDocumentDocumentColor, zap.Error(err))

		var params DocumentColorParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentDocumentColor(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentDocumentHighlight: // request
		defer logger.Debug(MethodTextDocumentDocumentHighlight, zap.Error(err))

		var params DocumentHighlightParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentDocumentHighlight(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentDocumentLink: // request
		defer logger.Debug(MethodTextDocumentDocumentLink, zap.Error(err))

		var params DocumentLinkParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentDocumentLink(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentDocumentSymbol: // request
		defer logger.Debug(MethodTextDocumentDocumentSymbol, zap.Error(err))

		var params DocumentSymbolParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentDocumentSymbol(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentFoldingRange: // request
		defer logger.Debug(MethodTextDocumentFoldingRange, zap.Error(err))

		var params FoldingRangeParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentFoldingRange(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentFormatting: // request
		defer logger.Debug(MethodTextDocumentFormatting, zap.Error(err))

		var params DocumentFormattingParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentFormatting(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentHover: // request
		defer logger.Debug(MethodTextDocumentHover, zap.Error(err))

		var params HoverParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentHover(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentImplementation: // request
		defer logger.Debug(MethodTextDocumentImplementation, zap.Error(err))

		var params ImplementationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentImplementation(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentInlayHint: // request
		defer logger.Debug(MethodTextDocumentInlayHint, zap.Error(err))

		var params InlayHintParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentInlayHint(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentInlineCompletion: // request
		defer logger.Debug(MethodTextDocumentInlineCompletion, zap.Error(err))

		var params InlineCompletionParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentInlineCompletion(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentInlineValue: // request
		defer logger.Debug(MethodTextDocumentInlineValue, zap.Error(err))

		var params InlineValueParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentInlineValue(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentLinkedEditingRange: // request
		defer logger.Debug(MethodTextDocumentLinkedEditingRange, zap.Error(err))

		var params LinkedEditingRangeParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentLinkedEditingRange(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentMoniker: // request
		defer logger.Debug(MethodTextDocumentMoniker, zap.Error(err))

		var params MonikerParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentMoniker(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentOnTypeFormatting: // request
		defer logger.Debug(MethodTextDocumentOnTypeFormatting, zap.Error(err))

		var params DocumentOnTypeFormattingParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentOnTypeFormatting(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentPrepareCallHierarchy: // request
		defer logger.Debug(MethodTextDocumentPrepareCallHierarchy, zap.Error(err))

		var params CallHierarchyPrepareParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentPrepareCallHierarchy(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentPrepareRename: // request
		defer logger.Debug(MethodTextDocumentPrepareRename, zap.Error(err))

		var params PrepareRenameParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentPrepareRename(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentPrepareTypeHierarchy: // request
		defer logger.Debug(MethodTextDocumentPrepareTypeHierarchy, zap.Error(err))

		var params TypeHierarchyPrepareParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentPrepareTypeHierarchy(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentRangeFormatting: // request
		defer logger.Debug(MethodTextDocumentRangeFormatting, zap.Error(err))

		var params DocumentRangeFormattingParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentRangeFormatting(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentRangesFormatting: // request
		defer logger.Debug(MethodTextDocumentRangesFormatting, zap.Error(err))

		var params DocumentRangesFormattingParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentRangesFormatting(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentReferences: // request
		defer logger.Debug(MethodTextDocumentReferences, zap.Error(err))

		var params ReferenceParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentReferences(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentRename: // request
		defer logger.Debug(MethodTextDocumentRename, zap.Error(err))

		var params RenameParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentRename(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentSelectionRange: // request
		defer logger.Debug(MethodTextDocumentSelectionRange, zap.Error(err))

		var params SelectionRangeParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentSelectionRange(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentSemanticTokensFull: // request
		defer logger.Debug(MethodTextDocumentSemanticTokensFull, zap.Error(err))

		var params SemanticTokensParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentSemanticTokensFull(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentSemanticTokensFullDelta: // request
		defer logger.Debug(MethodTextDocumentSemanticTokensFullDelta, zap.Error(err))

		var params SemanticTokensDeltaParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentSemanticTokensFullDelta(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentSemanticTokensRange: // request
		defer logger.Debug(MethodTextDocumentSemanticTokensRange, zap.Error(err))

		var params SemanticTokensRangeParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentSemanticTokensRange(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentSignatureHelp: // request
		defer logger.Debug(MethodTextDocumentSignatureHelp, zap.Error(err))

		var params SignatureHelpParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentSignatureHelp(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentTypeDefinition: // request
		defer logger.Debug(MethodTextDocumentTypeDefinition, zap.Error(err))

		var params TypeDefinitionParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentTypeDefinition(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTextDocumentWillSaveWaitUntil: // request
		defer logger.Debug(MethodTextDocumentWillSaveWaitUntil, zap.Error(err))

		var params WillSaveTextDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TextDocumentWillSaveWaitUntil(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTypeHierarchySubtypes: // request
		defer logger.Debug(MethodTypeHierarchySubtypes, zap.Error(err))

		var params TypeHierarchySubtypesParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TypeHierarchySubtypes(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodTypeHierarchySupertypes: // request
		defer logger.Debug(MethodTypeHierarchySupertypes, zap.Error(err))

		var params TypeHierarchySupertypesParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.TypeHierarchySupertypes(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWorkspaceDiagnostic: // request
		defer logger.Debug(MethodWorkspaceDiagnostic, zap.Error(err))

		var params WorkspaceDiagnosticParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.WorkspaceDiagnostic(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWorkspaceExecuteCommand: // request
		defer logger.Debug(MethodWorkspaceExecuteCommand, zap.Error(err))

		var params ExecuteCommandParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.WorkspaceExecuteCommand(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWorkspaceSymbol: // request
		defer logger.Debug(MethodWorkspaceSymbol, zap.Error(err))

		var params WorkspaceSymbolParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.WorkspaceSymbol(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWorkspaceWillCreateFiles: // request
		defer logger.Debug(MethodWorkspaceWillCreateFiles, zap.Error(err))

		var params CreateFilesParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.WorkspaceWillCreateFiles(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWorkspaceWillDeleteFiles: // request
		defer logger.Debug(MethodWorkspaceWillDeleteFiles, zap.Error(err))

		var params DeleteFilesParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.WorkspaceWillDeleteFiles(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWorkspaceWillRenameFiles: // request
		defer logger.Debug(MethodWorkspaceWillRenameFiles, zap.Error(err))

		var params RenameFilesParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.WorkspaceWillRenameFiles(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWorkspaceSymbolResolve: // request
		defer logger.Debug(MethodWorkspaceSymbolResolve, zap.Error(err))

		var params WorkspaceSymbol
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := server.WorkspaceSymbolResolve(ctx, &params)

		return true, reply(ctx, resp, err)

	default:
		return false, nil
	}
}

// server implements a Language Server Protocol server.
type server struct {
	jsonrpc2.Conn

	logger *zap.Logger
}

var _ Server = (*server)(nil)

func (s *server) CancelRequest(ctx context.Context, params *CancelParams) (err error) {
	s.logger.Debug("notify " + MethodClientCancelRequest)
	defer s.logger.Debug("end "+MethodClientCancelRequest, zap.Error(err))

	return s.Conn.Notify(ctx, MethodClientCancelRequest, params)
}

// Progress is the base protocol offers also support to report progress in a generic fashion.
//
// This mechanism can be used to report any kind of progress including work done progress (usually used to report progress in the user interface using a progress bar) and
// partial result progress to support streaming of results.
//
// @since 3.16.0.
func (s *server) Progress(ctx context.Context, params *ProgressParams) (err error) {
	s.logger.Debug("notify " + MethodClientProgress)
	defer s.logger.Debug("end "+MethodClientProgress, zap.Error(err))

	return s.Conn.Notify(ctx, MethodClientProgress, params)
}

// SetTrace a notification that should be used by the client to modify the trace setting of the server.
//
// @since 3.16.0.
func (s *server) SetTrace(ctx context.Context, params *SetTraceParams) (err error) {
	s.logger.Debug("notify " + MethodSetTrace)
	defer s.logger.Debug("end "+MethodSetTrace, zap.Error(err))

	return s.Conn.Notify(ctx, MethodSetTrace, params)
}

// Exit a notification to ask the server to exit its process.
//
// The server should exit with success code 0 if the shutdown request has been received before; otherwise with error code 1.
func (s *server) Exit(ctx context.Context) (err error) {
	s.logger.Debug("notify " + MethodExit)
	defer s.logger.Debug("end "+MethodExit, zap.Error(err))

	return s.Conn.Notify(ctx, MethodExit, nil)
}

// Initialized sends the notification from the client to the server after the client received the result of the
// initialize request but before the client is sending any other request or notification to the server.
//
// The server can use the initialized notification for example to dynamically register capabilities.
// The initialized notification may only be sent once.
func (s *server) Initialized(ctx context.Context, params *InitializedParams) (err error) {
	s.logger.Debug("notify " + MethodInitialized)
	defer s.logger.Debug("end "+MethodInitialized, zap.Error(err))

	return s.Conn.Notify(ctx, MethodInitialized, params)
}

func (s *server) NotebookDocumentDidChange(ctx context.Context, params *DidChangeNotebookDocumentParams) (err error) {
	s.logger.Debug("call " + MethodNotebookDocumentDidChange)
	defer s.logger.Debug("end "+MethodNotebookDocumentDidChange, zap.Error(err))

	return s.Conn.Notify(ctx, MethodNotebookDocumentDidChange, params)
}

// NotebookDocumentDidClose a notification sent when a notebook closes.
//
// @since 3.17.0
func (s *server) NotebookDocumentDidClose(ctx context.Context, params *DidCloseNotebookDocumentParams) (err error) {
	s.logger.Debug("call " + MethodNotebookDocumentDidClose)
	defer s.logger.Debug("end "+MethodNotebookDocumentDidClose, zap.Error(err))

	return s.Conn.Notify(ctx, MethodNotebookDocumentDidClose, params)
}

// NotebookDocumentDidOpen a notification sent when a notebook opens.
//
// @since 3.17.0
func (s *server) NotebookDocumentDidOpen(ctx context.Context, params *DidOpenNotebookDocumentParams) (err error) {
	s.logger.Debug("call " + MethodNotebookDocumentDidOpen)
	defer s.logger.Debug("end "+MethodNotebookDocumentDidOpen, zap.Error(err))

	return s.Conn.Notify(ctx, MethodNotebookDocumentDidOpen, params)
}

// NotebookDocumentDidSave a notification sent when a notebook document is saved.
//
// @since 3.17.0
func (s *server) NotebookDocumentDidSave(ctx context.Context, params *DidSaveNotebookDocumentParams) (err error) {
	s.logger.Debug("call " + MethodNotebookDocumentDidSave)
	defer s.logger.Debug("end "+MethodNotebookDocumentDidSave, zap.Error(err))

	return s.Conn.Notify(ctx, MethodNotebookDocumentDidSave, params)
}

// DidChange sends the notification from the client to the server to signal changes to a text document.
//
// In 2.0 the shape of the params has changed to include proper version numbers and language ids.
func (s *server) TextDocumentDidChange(ctx context.Context, params *DidChangeTextDocumentParams) (err error) {
	s.logger.Debug("notify " + MethodTextDocumentDidChange)
	defer s.logger.Debug("end "+MethodTextDocumentDidChange, zap.Error(err))

	return s.Conn.Notify(ctx, MethodTextDocumentDidChange, params)
}

// DidClose sends the notification from the client to the server when the document got closed in the client.
//
// The document’s truth now exists where the document’s Uri points to (e.g. if the document’s Uri is a file Uri the truth now exists on disk).
// As with the open notification the close notification is about managing the document’s content.
// Receiving a close notification doesn’t mean that the document was open in an editor before.
//
// A close notification requires a previous open notification to be sent.
// Note that a server’s ability to fulfill requests is independent of whether a text document is open or closed.
func (s *server) TextDocumentDidClose(ctx context.Context, params *DidCloseTextDocumentParams) (err error) {
	s.logger.Debug("call " + MethodTextDocumentDidClose)
	defer s.logger.Debug("end "+MethodTextDocumentDidClose, zap.Error(err))

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
func (s *server) TextDocumentDidOpen(ctx context.Context, params *DidOpenTextDocumentParams) (err error) {
	s.logger.Debug("call " + MethodTextDocumentDidOpen)
	defer s.logger.Debug("end "+MethodTextDocumentDidOpen, zap.Error(err))

	return s.Conn.Notify(ctx, MethodTextDocumentDidOpen, params)
}

// DidSave sends the notification from the client to the server when the document was saved in the client.
func (s *server) TextDocumentDidSave(ctx context.Context, params *DidSaveTextDocumentParams) (err error) {
	s.logger.Debug("call " + MethodTextDocumentDidSave)
	defer s.logger.Debug("end "+MethodTextDocumentDidSave, zap.Error(err))

	return s.Conn.Notify(ctx, MethodTextDocumentDidSave, params)
}

// WillSave sends the notification from the client to the server before the document is actually saved.
func (s *server) TextDocumentWillSave(ctx context.Context, params *WillSaveTextDocumentParams) (err error) {
	s.logger.Debug("call " + MethodTextDocumentWillSave)
	defer s.logger.Debug("end "+MethodTextDocumentWillSave, zap.Error(err))

	return s.Conn.Notify(ctx, MethodTextDocumentWillSave, params)
}

// WindowWorkDoneProgressCancel is the sends notification from the client to the server to cancel a progress initiated on the
// server side using the "window/workDoneProgress/create".
func (s *server) WindowWorkDoneProgressCancel(ctx context.Context, params *WorkDoneProgressCancelParams) (err error) {
	s.logger.Debug("call " + MethodWindowWorkDoneProgressCancel)
	defer s.logger.Debug("end "+MethodWindowWorkDoneProgressCancel, zap.Error(err))

	return s.Conn.Notify(ctx, MethodWindowWorkDoneProgressCancel, params)
}

// DidChangeConfiguration sends the notification from the client to the server to signal the change of configuration settings.
func (s *server) WorkspaceDidChangeConfiguration(ctx context.Context, params *DidChangeConfigurationParams) (err error) {
	s.logger.Debug("call " + MethodWorkspaceDidChangeConfiguration)
	defer s.logger.Debug("end "+MethodWorkspaceDidChangeConfiguration, zap.Error(err))

	return s.Conn.Notify(ctx, MethodWorkspaceDidChangeConfiguration, params)
}

// DidChangeWatchedFiles sends the notification from the client to the server when the client detects changes to files watched by the language client.
//
// It is recommended that servers register for these file events using the registration mechanism.
// In former implementations clients pushed file events without the server actively asking for it.
func (s *server) WorkspaceDidChangeWatchedFiles(ctx context.Context, params *DidChangeWatchedFilesParams) (err error) {
	s.logger.Debug("call " + MethodWorkspaceDidChangeWatchedFiles)
	defer s.logger.Debug("end "+MethodWorkspaceDidChangeWatchedFiles, zap.Error(err))

	return s.Conn.Notify(ctx, MethodWorkspaceDidChangeWatchedFiles, params)
}

// DidChangeWorkspaceFolders sents the notification from the client to the server to inform the server about workspace folder configuration changes.
//
// The notification is sent by default if both ServerCapabilities/workspace/workspaceFolders and ClientCapabilities/workspace/workspaceFolders are true;
// or if the server has registered itself to receive this notification.
// To register for the workspace/didChangeWorkspaceFolders send a client/registerCapability request from the server to the client.
//
// The registration parameter must have a registrations item of the following form, where id is a unique id used to unregister the capability (the example uses a UUID).
func (s *server) WorkspaceDidChangeWorkspaceFolders(ctx context.Context, params *DidChangeWorkspaceFoldersParams) (err error) {
	s.logger.Debug("call " + MethodWorkspaceDidChangeWorkspaceFolders)
	defer s.logger.Debug("end "+MethodWorkspaceDidChangeWorkspaceFolders, zap.Error(err))

	return s.Conn.Notify(ctx, MethodWorkspaceDidChangeWorkspaceFolders, params)
}

// DidCreateFiles sends the did create files notification is sent from the client to the server when files were created from within the client.
//
// @since 3.16.0.
func (s *server) WorkspaceDidCreateFiles(ctx context.Context, params *CreateFilesParams) (err error) {
	s.logger.Debug("call " + MethodWorkspaceDidCreateFiles)
	defer s.logger.Debug("end "+MethodWorkspaceDidCreateFiles, zap.Error(err))

	return s.Conn.Notify(ctx, MethodWorkspaceDidCreateFiles, params)
}

// DidDeleteFiles sends the did delete files notification is sent from the client to the server when files were deleted from within the client.
//
// @since 3.16.0.
func (s *server) WorkspaceDidDeleteFiles(ctx context.Context, params *DeleteFilesParams) (err error) {
	s.logger.Debug("call " + MethodWorkspaceDidDeleteFiles)
	defer s.logger.Debug("end "+MethodWorkspaceDidDeleteFiles, zap.Error(err))

	return s.Conn.Notify(ctx, MethodWorkspaceDidDeleteFiles, params)
}

// DidRenameFiles sends the did rename files notification is sent from the client to the server when files were renamed from within the client.
//
// @since 3.16.0.
func (s *server) WorkspaceDidRenameFiles(ctx context.Context, params *RenameFilesParams) (err error) {
	s.logger.Debug("call " + MethodWorkspaceDidRenameFiles)
	defer s.logger.Debug("end "+MethodWorkspaceDidRenameFiles, zap.Error(err))

	return s.Conn.Notify(ctx, MethodWorkspaceDidRenameFiles, params)
}

// IncomingCalls is the request is sent from the client to the server to resolve incoming calls for a given call hierarchy item.
//
// The request doesn’t define its own client and server capabilities. It is only issued if a server registers for the "textDocument/prepareCallHierarchy" request.
//
// @since 3.16.0.
func (s *server) CallHierarchyIncomingCalls(ctx context.Context, params *CallHierarchyIncomingCallsParams) (_ []*CallHierarchyIncomingCall, err error) {
	s.logger.Debug("call " + MethodCallHierarchyIncomingCalls)
	defer s.logger.Debug("end "+MethodCallHierarchyIncomingCalls, zap.Error(err))

	var result []*CallHierarchyIncomingCall
	if err := Call(ctx, s.Conn, MethodCallHierarchyIncomingCalls, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// OutgoingCalls is the request is sent from the client to the server to resolve outgoing calls for a given call hierarchy item.
//
// The request doesn’t define its own client and server capabilities. It is only issued if a server registers for the "textDocument/prepareCallHierarchy" request.
//
// @since 3.16.0.
func (s *server) CallHierarchyOutgoingCalls(ctx context.Context, params *CallHierarchyOutgoingCallsParams) (_ []*CallHierarchyOutgoingCall, err error) {
	s.logger.Debug("call " + MethodCallHierarchyOutgoingCalls)
	defer s.logger.Debug("end "+MethodCallHierarchyOutgoingCalls, zap.Error(err))

	var result []*CallHierarchyOutgoingCall
	if err := Call(ctx, s.Conn, MethodCallHierarchyOutgoingCalls, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) CodeActionResolve(ctx context.Context, params *CodeAction) (_ *CodeAction, err error) {
	s.logger.Debug("call " + MethodCodeActionResolve)
	defer s.logger.Debug("end "+MethodCodeActionResolve, zap.Error(err))

	var result *CodeAction
	if err := Call(ctx, s.Conn, MethodCodeActionResolve, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CodeLensResolve sends the request from the client to the server to resolve the command for a given code lens item.
func (s *server) CodeLensResolve(ctx context.Context, params *CodeLens) (_ *CodeLens, err error) {
	s.logger.Debug("call " + MethodCodeLensResolve)
	defer s.logger.Debug("end "+MethodCodeLensResolve, zap.Error(err))

	var result *CodeLens
	if err := Call(ctx, s.Conn, MethodCodeLensResolve, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CompletionResolve sends the request from the client to the server to resolve additional information for a given completion item.
func (s *server) CompletionItemResolve(ctx context.Context, params *CompletionItem) (_ *CompletionItem, err error) {
	s.logger.Debug("call " + MethodCompletionItemResolve)
	defer s.logger.Debug("end "+MethodCompletionItemResolve, zap.Error(err))

	var result *CompletionItem
	if err := Call(ctx, s.Conn, MethodCompletionItemResolve, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// DocumentLinkResolve sends the request from the client to the server to resolve the target of a given document link.
func (s *server) DocumentLinkResolve(ctx context.Context, params *DocumentLink) (_ *DocumentLink, err error) {
	s.logger.Debug("call " + MethodDocumentLinkResolve)
	defer s.logger.Debug("end "+MethodDocumentLinkResolve, zap.Error(err))

	var result *DocumentLink
	if err := Call(ctx, s.Conn, MethodDocumentLinkResolve, params, &result); err != nil {
		return nil, err
	}

	return result, nil
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
func (s *server) Initialize(ctx context.Context, params *InitializeParams) (_ *InitializeResult, err error) {
	s.logger.Debug("call " + MethodInitialize)
	defer s.logger.Debug("end "+MethodInitialize, zap.Error(err))

	var result *InitializeResult
	if err := Call(ctx, s.Conn, MethodInitialize, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) InlayHintResolve(ctx context.Context, params *InlayHint) (_ *InlayHint, err error) {
	s.logger.Debug("call " + MethodInlayHintResolve)
	defer s.logger.Debug("end "+MethodInlayHintResolve, zap.Error(err))

	var result *InlayHint
	if err := Call(ctx, s.Conn, MethodInlayHintResolve, params, &result); err != nil {
		return nil, err
	}

	return result, nil
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
	defer s.logger.Debug("end "+MethodShutdown, zap.Error(err))

	return Call(ctx, s.Conn, MethodShutdown, nil, nil)
}

// CodeAction sends the request is from the client to the server to compute commands for a given text document and range.
//
// These commands are typically code fixes to either fix problems or to beautify/refactor code. The result of a `textDocument/codeAction`
// request is an array of `Command` literals which are typically presented in the user interface.
//
// To ensure that a server is useful in many clients the commands specified in a code actions should be handled by the
// server and not by the client (see `workspace/executeCommand` and `ServerCapabilities.executeCommandProvider`).
// If the client supports providing edits with a code action then the mode should be used.
func (s *server) TextDocumentCodeAction(ctx context.Context, params *CodeActionParams) (_ *TextDocumentCodeActionResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentCodeAction)
	defer s.logger.Debug("end "+MethodTextDocumentCodeAction, zap.Error(err))

	var result *TextDocumentCodeActionResult
	if err := Call(ctx, s.Conn, MethodTextDocumentCodeAction, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CodeLens sends the request from the client to the server to compute code lenses for a given text document.
func (s *server) TextDocumentCodeLens(ctx context.Context, params *CodeLensParams) (_ []*CodeLens, err error) {
	s.logger.Debug("call " + MethodTextDocumentCodeLens)
	defer s.logger.Debug("end "+MethodTextDocumentCodeLens, zap.Error(err))

	var result []*CodeLens
	if err := Call(ctx, s.Conn, MethodTextDocumentCodeLens, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ColorPresentation sends the request from the client to the server to obtain a list of presentations for a color value at a given location.
//
// # Clients can use the result to
//
// - modify a color reference.
// - show in a color picker and let users pick one of the presentations.
func (s *server) TextDocumentColorPresentation(ctx context.Context, params *ColorPresentationParams) (_ []*ColorPresentation, err error) {
	s.logger.Debug("call " + MethodTextDocumentColorPresentation)
	defer s.logger.Debug("end "+MethodTextDocumentColorPresentation, zap.Error(err))

	var result []*ColorPresentation
	if err := Call(ctx, s.Conn, MethodTextDocumentColorPresentation, params, &result); err != nil {
		return nil, err
	}

	return result, nil
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
func (s *server) TextDocumentCompletion(ctx context.Context, params *CompletionParams) (_ *TextDocumentCompletionResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentCompletion)
	defer s.logger.Debug("end "+MethodTextDocumentCompletion, zap.Error(err))

	var result *TextDocumentCompletionResult
	if err := Call(ctx, s.Conn, MethodTextDocumentCompletion, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Declaration sends the request from the client to the server to resolve the declaration location of a symbol at a given text document position.
//
// The result type LocationLink[] got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.textDocument.declaration.linkSupport`.
//
// @since 3.14.0.
func (s *server) TextDocumentDeclaration(ctx context.Context, params *DeclarationParams) (_ *TextDocumentDeclarationResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentDeclaration)
	defer s.logger.Debug("end "+MethodTextDocumentDeclaration, zap.Error(err))

	var result *TextDocumentDeclarationResult
	if err := Call(ctx, s.Conn, MethodTextDocumentDeclaration, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Definition sends the request from the client to the server to resolve the definition location of a symbol at a given text document position.
//
// The result type `[]LocationLink` got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.textDocument.definition.linkSupport`.
//
// @since 3.14.0.
func (s *server) TextDocumentDefinition(ctx context.Context, params *DefinitionParams) (_ *TextDocumentDefinitionResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentDefinition)
	defer s.logger.Debug("end "+MethodTextDocumentDefinition, zap.Error(err))

	var result *TextDocumentDefinitionResult
	if err := Call(ctx, s.Conn, MethodTextDocumentDefinition, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) TextDocumentDiagnostic(ctx context.Context, params *DocumentDiagnosticParams) (_ *DocumentDiagnosticReport, err error) {
	s.logger.Debug("call " + MethodTextDocumentDiagnostic)
	defer s.logger.Debug("end "+MethodTextDocumentDiagnostic, zap.Error(err))

	var result *DocumentDiagnosticReport
	if err := Call(ctx, s.Conn, MethodTextDocumentDiagnostic, params, &result); err != nil {
		return nil, err
	}

	return result, nil
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
func (s *server) TextDocumentDocumentColor(ctx context.Context, params *DocumentColorParams) (_ []*ColorInformation, err error) {
	s.logger.Debug("call " + MethodTextDocumentDocumentColor)
	defer s.logger.Debug("end "+MethodTextDocumentDocumentColor, zap.Error(err))

	var result []*ColorInformation
	if err := Call(ctx, s.Conn, MethodTextDocumentDocumentColor, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// DocumentHighlight sends the request is from the client to the server to resolve a document highlights for a given text document position.
//
// For programming languages this usually highlights all references to the symbol scoped to this file.
// However we kept ‘textDocument/documentHighlight’ and ‘textDocument/references’ separate requests since the first one is allowed to be more fuzzy.
//
// Symbol matches usually have a `DocumentHighlightKind` of `Read` or `Write` whereas fuzzy or textual matches use `Text` as the kind.
func (s *server) TextDocumentDocumentHighlight(ctx context.Context, params *DocumentHighlightParams) (_ []*DocumentHighlight, err error) {
	s.logger.Debug("call " + MethodTextDocumentDocumentHighlight)
	defer s.logger.Debug("end "+MethodTextDocumentDocumentHighlight, zap.Error(err))

	var result []*DocumentHighlight
	if err := Call(ctx, s.Conn, MethodTextDocumentDocumentHighlight, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// DocumentLink sends the request from the client to the server to request the location of links in a document.
func (s *server) TextDocumentDocumentLink(ctx context.Context, params *DocumentLinkParams) (_ []*DocumentLink, err error) {
	s.logger.Debug("call " + MethodTextDocumentDocumentLink)
	defer s.logger.Debug("end "+MethodTextDocumentDocumentLink, zap.Error(err))

	var result []*DocumentLink
	if err := Call(ctx, s.Conn, MethodTextDocumentDocumentLink, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// DocumentSymbol sends the request from the client to the server to return a flat list of all symbols found in a given text document.
//
// Neither the symbol’s location range nor the symbol’s container name should be used to infer a hierarchy.
func (s *server) TextDocumentDocumentSymbol(ctx context.Context, params *DocumentSymbolParams) (_ *TextDocumentDocumentSymbolResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentDocumentSymbol)
	defer s.logger.Debug("end "+MethodTextDocumentDocumentSymbol, zap.Error(err))

	var result *TextDocumentDocumentSymbolResult
	if err := Call(ctx, s.Conn, MethodTextDocumentDocumentSymbol, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// FoldingRanges sends the request from the client to the server to return all folding ranges found in a given text document.
//
// @since version 3.10.0.
func (s *server) TextDocumentFoldingRange(ctx context.Context, params *FoldingRangeParams) (_ []*FoldingRange, err error) {
	s.logger.Debug("call " + MethodTextDocumentFoldingRange)
	defer s.logger.Debug("end "+MethodTextDocumentFoldingRange, zap.Error(err))

	var result []*FoldingRange
	if err := Call(ctx, s.Conn, MethodTextDocumentFoldingRange, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Formatting sends the request from the client to the server to format a whole document.
func (s *server) TextDocumentFormatting(ctx context.Context, params *DocumentFormattingParams) (_ []*TextEdit, err error) {
	s.logger.Debug("call " + MethodTextDocumentFormatting)
	defer s.logger.Debug("end "+MethodTextDocumentFormatting, zap.Error(err))

	var result []*TextEdit
	if err := Call(ctx, s.Conn, MethodTextDocumentFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Hover sends the request is from the client to the server to request hover information at a given text document position.
func (s *server) TextDocumentHover(ctx context.Context, params *HoverParams) (_ *Hover, err error) {
	s.logger.Debug("call " + MethodTextDocumentHover)
	defer s.logger.Debug("end "+MethodTextDocumentHover, zap.Error(err))

	var result *Hover
	if err := Call(ctx, s.Conn, MethodTextDocumentHover, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Implementation sends the request from the client to the server to resolve the implementation location of a symbol at a given text document position.
//
// The result type `[]LocationLink` got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.implementation.typeDefinition.linkSupport`.
func (s *server) TextDocumentImplementation(ctx context.Context, params *ImplementationParams) (_ *TextDocumentImplementationResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentImplementation)
	defer s.logger.Debug("end "+MethodTextDocumentImplementation, zap.Error(err))

	var result *TextDocumentImplementationResult
	if err := Call(ctx, s.Conn, MethodTextDocumentImplementation, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) TextDocumentInlayHint(ctx context.Context, params *InlayHintParams) (_ []*InlayHint, err error) {
	s.logger.Debug("call " + MethodTextDocumentInlayHint)
	defer s.logger.Debug("end "+MethodTextDocumentInlayHint, zap.Error(err))

	var result []*InlayHint
	if err := Call(ctx, s.Conn, MethodTextDocumentInlayHint, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) TextDocumentInlineCompletion(ctx context.Context, params *InlineCompletionParams) (_ *TextDocumentInlineCompletionResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentInlineCompletion)
	defer s.logger.Debug("end "+MethodTextDocumentInlineCompletion, zap.Error(err))

	var result *TextDocumentInlineCompletionResult
	if err := Call(ctx, s.Conn, MethodTextDocumentInlineCompletion, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) TextDocumentInlineValue(ctx context.Context, params *InlineValueParams) (_ []*InlineValue, err error) {
	s.logger.Debug("call " + MethodTextDocumentInlineValue)
	defer s.logger.Debug("end "+MethodTextDocumentInlineValue, zap.Error(err))

	var result []*InlineValue
	if err := Call(ctx, s.Conn, MethodTextDocumentInlineValue, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// LinkedEditingRange is the linked editing request is sent from the client to the server to return for a given position in a document the range of the symbol at the position and all ranges that have the same content.
//
// Optionally a word pattern can be returned to describe valid contents.
//
// A rename to one of the ranges can be applied to all other ranges if the new content is valid. If no result-specific word pattern is provided, the word pattern from the client’s language configuration is used.
//
// @since 3.16.0.
func (s *server) TextDocumentLinkedEditingRange(ctx context.Context, params *LinkedEditingRangeParams) (_ *LinkedEditingRanges, err error) {
	s.logger.Debug("call " + MethodTextDocumentLinkedEditingRange)
	defer s.logger.Debug("end "+MethodTextDocumentLinkedEditingRange, zap.Error(err))

	var result *LinkedEditingRanges
	if err := Call(ctx, s.Conn, MethodTextDocumentLinkedEditingRange, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Moniker is the request is sent from the client to the server to get the symbol monikers for a given text document position.
//
// An array of Moniker types is returned as response to indicate possible monikers at the given location.
//
// If no monikers can be calculated, an empty array or null should be returned.
//
// @since 3.16.0.
func (s *server) TextDocumentMoniker(ctx context.Context, params *MonikerParams) (_ []*Moniker, err error) {
	s.logger.Debug("call " + MethodTextDocumentMoniker)
	defer s.logger.Debug("end "+MethodTextDocumentMoniker, zap.Error(err))

	var result []*Moniker
	if err := Call(ctx, s.Conn, MethodTextDocumentMoniker, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// OnTypeFormatting sends the request from the client to the server to format parts of the document during typing.
func (s *server) TextDocumentOnTypeFormatting(ctx context.Context, params *DocumentOnTypeFormattingParams) (_ []*TextEdit, err error) {
	s.logger.Debug("call " + MethodTextDocumentOnTypeFormatting)
	defer s.logger.Debug("end "+MethodTextDocumentOnTypeFormatting, zap.Error(err))

	var result []*TextEdit
	if err := Call(ctx, s.Conn, MethodTextDocumentOnTypeFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// PrepareCallHierarchy sent from the client to the server to return a call hierarchy for the language element of given text document positions.
//
// The call hierarchy requests are executed in two steps:
//  1. first a call hierarchy item is resolved for the given text document position
//  2. for a call hierarchy item the incoming or outgoing call hierarchy items are resolved.
//
// @since 3.16.0.
func (s *server) TextDocumentPrepareCallHierarchy(ctx context.Context, params *CallHierarchyPrepareParams) (_ []*CallHierarchyItem, err error) {
	s.logger.Debug("call " + MethodTextDocumentPrepareCallHierarchy)
	defer s.logger.Debug("end "+MethodTextDocumentPrepareCallHierarchy, zap.Error(err))

	var result []*CallHierarchyItem
	if err := Call(ctx, s.Conn, MethodTextDocumentPrepareCallHierarchy, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// PrepareRename sends the request from the client to the server to setup and test the validity of a rename operation at a given location.
//
// @since version 3.12.0.
func (s *server) TextDocumentPrepareRename(ctx context.Context, params *PrepareRenameParams) (_ *PrepareRenameResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentPrepareRename)
	defer s.logger.Debug("end "+MethodTextDocumentPrepareRename, zap.Error(err))

	var result *PrepareRenameResult
	if err := Call(ctx, s.Conn, MethodTextDocumentPrepareRename, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) TextDocumentPrepareTypeHierarchy(ctx context.Context, params *TypeHierarchyPrepareParams) (_ []*TypeHierarchyItem, err error) {
	s.logger.Debug("call " + MethodTextDocumentPrepareTypeHierarchy)
	defer s.logger.Debug("end "+MethodTextDocumentPrepareTypeHierarchy, zap.Error(err))

	var result []*TypeHierarchyItem
	if err := Call(ctx, s.Conn, MethodTextDocumentPrepareTypeHierarchy, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// RangeFormatting sends the request from the client to the server to format a given range in a document.
func (s *server) TextDocumentRangeFormatting(ctx context.Context, params *DocumentRangeFormattingParams) (_ []*TextEdit, err error) {
	s.logger.Debug("call " + MethodTextDocumentRangeFormatting)
	defer s.logger.Debug("end "+MethodTextDocumentRangeFormatting, zap.Error(err))

	var result []*TextEdit
	if err := Call(ctx, s.Conn, MethodTextDocumentRangeFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) TextDocumentRangesFormatting(ctx context.Context, params *DocumentRangesFormattingParams) (_ []*TextEdit, err error) {
	s.logger.Debug("call " + MethodTextDocumentRangesFormatting)
	defer s.logger.Debug("end "+MethodTextDocumentRangesFormatting, zap.Error(err))

	var result []*TextEdit
	if err := Call(ctx, s.Conn, MethodTextDocumentRangesFormatting, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// References sends the request from the client to the server to resolve project-wide references for the symbol denoted by the given text document position.
func (s *server) TextDocumentReferences(ctx context.Context, params *ReferenceParams) (_ []*Location, err error) {
	s.logger.Debug("call " + MethodTextDocumentReferences)
	defer s.logger.Debug("end "+MethodTextDocumentReferences, zap.Error(err))

	var result []*Location
	if err := Call(ctx, s.Conn, MethodTextDocumentReferences, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Rename sends the request from the client to the server to perform a workspace-wide rename of a symbol.
func (s *server) TextDocumentRename(ctx context.Context, params *RenameParams) (_ *WorkspaceEdit, err error) {
	s.logger.Debug("call " + MethodTextDocumentRename)
	defer s.logger.Debug("end "+MethodTextDocumentRename, zap.Error(err))

	var result *WorkspaceEdit
	if err := Call(ctx, s.Conn, MethodTextDocumentRename, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) TextDocumentSelectionRange(ctx context.Context, params *SelectionRangeParams) (_ []*SelectionRange, err error) {
	s.logger.Debug("call " + MethodTextDocumentSelectionRange)
	defer s.logger.Debug("end "+MethodTextDocumentSelectionRange, zap.Error(err))

	var result []*SelectionRange
	if err := Call(ctx, s.Conn, MethodTextDocumentSelectionRange, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SemanticTokensFull is the request is sent from the client to the server to resolve semantic tokens for a given file.
//
// Semantic tokens are used to add additional color information to a file that depends on language specific symbol information.
//
// A semantic token request usually produces a large result. The protocol therefore supports encoding tokens with numbers.
//
// @since 3.16.0.
func (s *server) TextDocumentSemanticTokensFull(ctx context.Context, params *SemanticTokensParams) (_ *SemanticTokens, err error) {
	s.logger.Debug("call " + MethodTextDocumentSemanticTokensFull)
	defer s.logger.Debug("end "+MethodTextDocumentSemanticTokensFull, zap.Error(err))

	var result *SemanticTokens
	if err := Call(ctx, s.Conn, MethodTextDocumentSemanticTokensFull, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SemanticTokensFullDelta is the request is sent from the client to the server to resolve semantic token delta for a given file.
//
// Semantic tokens are used to add additional color information to a file that depends on language specific symbol information.
//
// A semantic token request usually produces a large result. The protocol therefore supports encoding tokens with numbers.
//
// @since 3.16.0.
func (s *server) TextDocumentSemanticTokensFullDelta(ctx context.Context, params *SemanticTokensDeltaParams) (_ *TextDocumentSemanticTokensFullDeltaResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentSemanticTokensFullDelta)
	defer s.logger.Debug("end "+MethodTextDocumentSemanticTokensFullDelta, zap.Error(err))

	var result *TextDocumentSemanticTokensFullDeltaResult
	if err := Call(ctx, s.Conn, MethodTextDocumentSemanticTokensFullDelta, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SemanticTokensRange is the request is sent from the client to the server to resolve semantic token delta for a given file.
//
// When a user opens a file it can be beneficial to only compute the semantic tokens for the visible range (faster rendering of the tokens in the user interface).
// If a server can compute these tokens faster than for the whole file it can provide a handler for the "textDocument/semanticTokens/range" request to handle this case special.
//
// Please note that if a client also announces that it will send the "textDocument/semanticTokens/range" server should implement this request as well to allow for flicker free scrolling and semantic coloring of a minimap.
//
// @since 3.16.0.
func (s *server) TextDocumentSemanticTokensRange(ctx context.Context, params *SemanticTokensRangeParams) (_ *SemanticTokens, err error) {
	s.logger.Debug("call " + MethodTextDocumentSemanticTokensRange)
	defer s.logger.Debug("end "+MethodTextDocumentSemanticTokensRange, zap.Error(err))

	var result *SemanticTokens
	if err := Call(ctx, s.Conn, MethodTextDocumentSemanticTokensRange, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SignatureHelp sends the request from the client to the server to request signature information at a given cursor position.
func (s *server) TextDocumentSignatureHelp(ctx context.Context, params *SignatureHelpParams) (_ *SignatureHelp, err error) {
	s.logger.Debug("call " + MethodTextDocumentSignatureHelp)
	defer s.logger.Debug("end "+MethodTextDocumentSignatureHelp, zap.Error(err))

	var result *SignatureHelp
	if err := Call(ctx, s.Conn, MethodTextDocumentSignatureHelp, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TypeDefinition sends the request from the client to the server to resolve the type definition location of a symbol at a given text document position.
//
// The result type `[]LocationLink` got introduce with version 3.14.0 and depends in the corresponding client capability `clientCapabilities.textDocument.typeDefinition.linkSupport`.
//
// @since version 3.6.0.
func (s *server) TextDocumentTypeDefinition(ctx context.Context, params *TypeDefinitionParams) (_ *TextDocumentTypeDefinitionResult, err error) {
	s.logger.Debug("call " + MethodTextDocumentTypeDefinition)
	defer s.logger.Debug("end "+MethodTextDocumentTypeDefinition, zap.Error(err))

	var result *TextDocumentTypeDefinitionResult
	if err := Call(ctx, s.Conn, MethodTextDocumentTypeDefinition, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// WillSaveWaitUntil sends the request from the client to the server before the document is actually saved.
//
// The request can return an array of TextEdits which will be applied to the text document before it is saved.
// Please note that clients might drop results if computing the text edits took too long or if a server constantly fails on this request.
// This is done to keep the save fast and reliable.
func (s *server) TextDocumentWillSaveWaitUntil(ctx context.Context, params *WillSaveTextDocumentParams) (_ []*TextEdit, err error) {
	s.logger.Debug("call " + MethodTextDocumentWillSaveWaitUntil)
	defer s.logger.Debug("end "+MethodTextDocumentWillSaveWaitUntil, zap.Error(err))

	var result []*TextEdit
	if err := Call(ctx, s.Conn, MethodTextDocumentWillSaveWaitUntil, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) TypeHierarchySubtypes(ctx context.Context, params *TypeHierarchySubtypesParams) (_ []*TypeHierarchyItem, err error) {
	s.logger.Debug("call " + MethodTypeHierarchySubtypes)
	defer s.logger.Debug("end "+MethodTypeHierarchySubtypes, zap.Error(err))

	var result []*TypeHierarchyItem
	if err := Call(ctx, s.Conn, MethodTypeHierarchySubtypes, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) TypeHierarchySupertypes(ctx context.Context, params *TypeHierarchySupertypesParams) (_ []*TypeHierarchyItem, err error) {
	s.logger.Debug("call " + MethodTypeHierarchySupertypes)
	defer s.logger.Debug("end "+MethodTypeHierarchySupertypes, zap.Error(err))

	var result []*TypeHierarchyItem
	if err := Call(ctx, s.Conn, MethodTypeHierarchySupertypes, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) WorkspaceDiagnostic(ctx context.Context, params *WorkspaceDiagnosticParams) (_ *WorkspaceDiagnosticReport, err error) {
	s.logger.Debug("call " + MethodWorkspaceDiagnostic)
	defer s.logger.Debug("end "+MethodWorkspaceDiagnostic, zap.Error(err))

	var result *WorkspaceDiagnosticReport
	if err := Call(ctx, s.Conn, MethodWorkspaceDiagnostic, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ExecuteCommand sends the request from the client to the server to trigger command execution on the server.
//
// In most cases the server creates a `WorkspaceEdit` structure and applies the changes to the workspace using the
// request `workspace/applyEdit` which is sent from the server to the client.
func (s *server) WorkspaceExecuteCommand(ctx context.Context, params *ExecuteCommandParams) (result any, err error) {
	s.logger.Debug("call " + MethodWorkspaceExecuteCommand)
	defer s.logger.Debug("end "+MethodWorkspaceExecuteCommand, zap.Error(err))

	if err := Call(ctx, s.Conn, MethodWorkspaceExecuteCommand, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Symbols sends the request from the client to the server to list project-wide symbols matching the query string.
func (s *server) WorkspaceSymbol(ctx context.Context, params *WorkspaceSymbolParams) (_ *WorkspaceSymbolResult, err error) {
	s.logger.Debug("call " + MethodWorkspaceSymbol)
	defer s.logger.Debug("end "+MethodWorkspaceSymbol, zap.Error(err))

	var result *WorkspaceSymbolResult
	if err := Call(ctx, s.Conn, MethodWorkspaceSymbol, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// WillCreateFiles sends the will create files request is sent from the client to the server before files are actually created as long as the creation is triggered from within the client.
//
// The request can return a WorkspaceEdit which will be applied to workspace before the files are created.
//
// Please note that clients might drop results if computing the edit took too long or if a server constantly fails on this request. This is done to keep creates fast and reliable.
//
// @since 3.16.0.
func (s *server) WorkspaceWillCreateFiles(ctx context.Context, params *CreateFilesParams) (_ *WorkspaceEdit, err error) {
	s.logger.Debug("call " + MethodWorkspaceWillCreateFiles)
	defer s.logger.Debug("end "+MethodWorkspaceWillCreateFiles, zap.Error(err))

	var result *WorkspaceEdit
	if err := Call(ctx, s.Conn, MethodWorkspaceWillCreateFiles, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// WillDeleteFiles sends the will delete files request is sent from the client to the server before files are actually deleted as long as the deletion is triggered from within the client.
//
// The request can return a WorkspaceEdit which will be applied to workspace before the files are deleted.
//
// Please note that clients might drop results if computing the edit took too long or if a server constantly fails on this request. This is done to keep deletes fast and reliable.
//
// @since 3.16.0.
func (s *server) WorkspaceWillDeleteFiles(ctx context.Context, params *DeleteFilesParams) (_ *WorkspaceEdit, err error) {
	s.logger.Debug("call " + MethodWorkspaceWillDeleteFiles)
	defer s.logger.Debug("end "+MethodWorkspaceWillDeleteFiles, zap.Error(err))

	var result *WorkspaceEdit
	if err := Call(ctx, s.Conn, MethodWorkspaceWillDeleteFiles, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// WillRenameFiles sends the will rename files request is sent from the client to the server before files are actually renamed as long as the rename is triggered from within the client.
//
// The request can return a WorkspaceEdit which will be applied to workspace before the files are renamed.
//
// Please note that clients might drop results if computing the edit took too long or if a server constantly fails on this request. This is done to keep renames fast and reliable.
//
// @since 3.16.0.
func (s *server) WorkspaceWillRenameFiles(ctx context.Context, params *RenameFilesParams) (_ *WorkspaceEdit, err error) {
	s.logger.Debug("call " + MethodWorkspaceWillRenameFiles)
	defer s.logger.Debug("end "+MethodWorkspaceWillRenameFiles, zap.Error(err))

	var result *WorkspaceEdit
	if err := Call(ctx, s.Conn, MethodWorkspaceWillRenameFiles, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *server) WorkspaceSymbolResolve(ctx context.Context, params *WorkspaceSymbol) (_ *WorkspaceSymbol, err error) {
	s.logger.Debug("call " + MethodWorkspaceSymbolResolve)
	defer s.logger.Debug("end "+MethodWorkspaceSymbolResolve, zap.Error(err))

	var result *WorkspaceSymbol
	if err := Call(ctx, s.Conn, MethodWorkspaceSymbolResolve, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Request sends a request from the client to the server that non-compliant with the Language Server Protocol specifications.
func (s *server) Request(ctx context.Context, method string, params any) (any, error) {
	s.logger.Debug("call " + method)
	defer s.logger.Debug("end " + method)

	var result any
	if err := Call(ctx, s.Conn, method, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}
