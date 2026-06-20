// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"errors"
	"net"
	"sync"
	"testing"
	"time"

	"go.lsp.dev/jsonrpc2"
)

// TestServerNotificationMethods drives every client->server notification sending
// method through the *server dispatcher and asserts it lands on the matching
// method constant via Conn.Notify (not Call).
func TestServerNotificationMethods(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		wantMethod string
		invoke     func(ctx context.Context, srv Server) error
	}{
		"initialized": {
			wantMethod: MethodInitialized,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.Initialized(ctx, &InitializedParams{})
			},
		},
		"exit": {
			wantMethod: MethodExit,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.Exit(ctx)
			},
		},
		"setTrace": {
			wantMethod: MethodSetTrace,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.SetTrace(ctx, &SetTraceParams{})
			},
		},
		"progress": {
			wantMethod: MethodProgress,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.Progress(ctx, &ProgressParams{})
			},
		},
		"workDoneProgressCancel": {
			wantMethod: MethodWindowWorkDoneProgressCancel,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.WorkDoneProgressCancel(ctx, &WorkDoneProgressCancelParams{})
			},
		},
		"didOpen": {
			wantMethod: MethodTextDocumentDidOpen,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidOpen(ctx, &DidOpenTextDocumentParams{})
			},
		},
		"didChange": {
			wantMethod: MethodTextDocumentDidChange,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidChange(ctx, &DidChangeTextDocumentParams{})
			},
		},
		"willSave": {
			wantMethod: MethodTextDocumentWillSave,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.WillSave(ctx, &WillSaveTextDocumentParams{})
			},
		},
		"didSave": {
			wantMethod: MethodTextDocumentDidSave,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidSave(ctx, &DidSaveTextDocumentParams{})
			},
		},
		"didClose": {
			wantMethod: MethodTextDocumentDidClose,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidClose(ctx, &DidCloseTextDocumentParams{})
			},
		},
		"didOpenNotebookDocument": {
			wantMethod: MethodNotebookDocumentDidOpen,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidOpenNotebookDocument(ctx, &DidOpenNotebookDocumentParams{})
			},
		},
		"didChangeNotebookDocument": {
			wantMethod: MethodNotebookDocumentDidChange,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidChangeNotebookDocument(ctx, &DidChangeNotebookDocumentParams{})
			},
		},
		"didSaveNotebookDocument": {
			wantMethod: MethodNotebookDocumentDidSave,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidSaveNotebookDocument(ctx, &DidSaveNotebookDocumentParams{})
			},
		},
		"didCloseNotebookDocument": {
			wantMethod: MethodNotebookDocumentDidClose,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidCloseNotebookDocument(ctx, &DidCloseNotebookDocumentParams{})
			},
		},
		"didChangeConfiguration": {
			wantMethod: MethodWorkspaceDidChangeConfiguration,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidChangeConfiguration(ctx, &DidChangeConfigurationParams{})
			},
		},
		"didChangeWorkspaceFolders": {
			wantMethod: MethodWorkspaceDidChangeWorkspaceFolders,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidChangeWorkspaceFolders(ctx, &DidChangeWorkspaceFoldersParams{})
			},
		},
		"didCreateFiles": {
			wantMethod: MethodWorkspaceDidCreateFiles,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidCreateFiles(ctx, &CreateFilesParams{})
			},
		},
		"didRenameFiles": {
			wantMethod: MethodWorkspaceDidRenameFiles,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidRenameFiles(ctx, &RenameFilesParams{})
			},
		},
		"didDeleteFiles": {
			wantMethod: MethodWorkspaceDidDeleteFiles,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidDeleteFiles(ctx, &DeleteFilesParams{})
			},
		},
		"didChangeWatchedFiles": {
			wantMethod: MethodWorkspaceDidChangeWatchedFiles,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.DidChangeWatchedFiles(ctx, &DidChangeWatchedFilesParams{})
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			conn := &fakeConn{}
			srv := ServerDispatcher(conn)
			if err := tt.invoke(t.Context(), srv); err != nil {
				t.Fatalf("invoke: %v", err)
			}
			method, _, calls, notifies := conn.snapshot()
			if method != tt.wantMethod {
				t.Errorf("method = %q, want %q", method, tt.wantMethod)
			}
			if calls != 0 {
				t.Errorf("Call count = %d, want 0 (notification must not use Call)", calls)
			}
			if notifies != 1 {
				t.Errorf("Notify count = %d, want 1", notifies)
			}
		})
	}
}

// TestServerRequestMethods drives every client->server request sending method
// through the *server dispatcher and asserts it lands on the matching method
// constant via Conn.Call.
func TestServerRequestMethods(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		wantMethod string
		invoke     func(ctx context.Context, srv Server) error
	}{
		"initialize": {
			wantMethod: MethodInitialize,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Initialize(ctx, &InitializeParams{})
				return err
			},
		},
		"shutdown": {
			wantMethod: MethodShutdown,
			invoke: func(ctx context.Context, srv Server) error {
				return srv.Shutdown(ctx)
			},
		},
		"willSaveWaitUntil": {
			wantMethod: MethodTextDocumentWillSaveWaitUntil,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.WillSaveWaitUntil(ctx, &WillSaveTextDocumentParams{})
				return err
			},
		},
		"declaration": {
			wantMethod: MethodTextDocumentDeclaration,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Declaration(ctx, &DeclarationParams{})
				return err
			},
		},
		"definition": {
			wantMethod: MethodTextDocumentDefinition,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Definition(ctx, &DefinitionParams{})
				return err
			},
		},
		"typeDefinition": {
			wantMethod: MethodTextDocumentTypeDefinition,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.TypeDefinition(ctx, &TypeDefinitionParams{})
				return err
			},
		},
		"implementation": {
			wantMethod: MethodTextDocumentImplementation,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Implementation(ctx, &ImplementationParams{})
				return err
			},
		},
		"references": {
			wantMethod: MethodTextDocumentReferences,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.References(ctx, &ReferenceParams{})
				return err
			},
		},
		"prepareCallHierarchy": {
			wantMethod: MethodTextDocumentPrepareCallHierarchy,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.PrepareCallHierarchy(ctx, &CallHierarchyPrepareParams{})
				return err
			},
		},
		"incomingCalls": {
			wantMethod: MethodCallHierarchyIncomingCalls,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.IncomingCalls(ctx, &CallHierarchyIncomingCallsParams{})
				return err
			},
		},
		"outgoingCalls": {
			wantMethod: MethodCallHierarchyOutgoingCalls,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.OutgoingCalls(ctx, &CallHierarchyOutgoingCallsParams{})
				return err
			},
		},
		"prepareTypeHierarchy": {
			wantMethod: MethodTextDocumentPrepareTypeHierarchy,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.PrepareTypeHierarchy(ctx, &TypeHierarchyPrepareParams{})
				return err
			},
		},
		"supertypes": {
			wantMethod: MethodTypeHierarchySupertypes,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Supertypes(ctx, &TypeHierarchySupertypesParams{})
				return err
			},
		},
		"subtypes": {
			wantMethod: MethodTypeHierarchySubtypes,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Subtypes(ctx, &TypeHierarchySubtypesParams{})
				return err
			},
		},
		"documentHighlight": {
			wantMethod: MethodTextDocumentDocumentHighlight,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.DocumentHighlight(ctx, &DocumentHighlightParams{})
				return err
			},
		},
		"documentLink": {
			wantMethod: MethodTextDocumentDocumentLink,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.DocumentLink(ctx, &DocumentLinkParams{})
				return err
			},
		},
		"documentLinkResolve": {
			wantMethod: MethodDocumentLinkResolve,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.DocumentLinkResolve(ctx, &DocumentLink{})
				return err
			},
		},
		"hover": {
			wantMethod: MethodTextDocumentHover,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Hover(ctx, &HoverParams{})
				return err
			},
		},
		"codeLens": {
			wantMethod: MethodTextDocumentCodeLens,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.CodeLens(ctx, &CodeLensParams{})
				return err
			},
		},
		"codeLensResolve": {
			wantMethod: MethodCodeLensResolve,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.CodeLensResolve(ctx, &CodeLens{})
				return err
			},
		},
		"foldingRanges": {
			wantMethod: MethodTextDocumentFoldingRange,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.FoldingRanges(ctx, &FoldingRangeParams{})
				return err
			},
		},
		"selectionRange": {
			wantMethod: MethodTextDocumentSelectionRange,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.SelectionRange(ctx, &SelectionRangeParams{})
				return err
			},
		},
		"documentSymbol": {
			wantMethod: MethodTextDocumentDocumentSymbol,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.DocumentSymbol(ctx, &DocumentSymbolParams{})
				return err
			},
		},
		"semanticTokensFull": {
			wantMethod: MethodTextDocumentSemanticTokensFull,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.SemanticTokensFull(ctx, &SemanticTokensParams{})
				return err
			},
		},
		"semanticTokensFullDelta": {
			wantMethod: MethodTextDocumentSemanticTokensFullDelta,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.SemanticTokensFullDelta(ctx, &SemanticTokensDeltaParams{})
				return err
			},
		},
		"semanticTokensRange": {
			wantMethod: MethodTextDocumentSemanticTokensRange,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.SemanticTokensRange(ctx, &SemanticTokensRangeParams{})
				return err
			},
		},
		"inlineValue": {
			wantMethod: MethodTextDocumentInlineValue,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.InlineValue(ctx, &InlineValueParams{})
				return err
			},
		},
		"inlayHint": {
			wantMethod: MethodTextDocumentInlayHint,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.InlayHint(ctx, &InlayHintParams{})
				return err
			},
		},
		"inlayHintResolve": {
			wantMethod: MethodInlayHintResolve,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.InlayHintResolve(ctx, &InlayHint{})
				return err
			},
		},
		"moniker": {
			wantMethod: MethodTextDocumentMoniker,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Moniker(ctx, &MonikerParams{})
				return err
			},
		},
		"completion": {
			wantMethod: MethodTextDocumentCompletion,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Completion(ctx, &CompletionParams{})
				return err
			},
		},
		"completionResolve": {
			wantMethod: MethodCompletionItemResolve,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.CompletionResolve(ctx, &CompletionItem{})
				return err
			},
		},
		"diagnostic": {
			wantMethod: MethodTextDocumentDiagnostic,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Diagnostic(ctx, &DocumentDiagnosticParams{})
				return err
			},
		},
		"diagnosticWorkspace": {
			wantMethod: MethodWorkspaceDiagnostic,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.DiagnosticWorkspace(ctx, &WorkspaceDiagnosticParams{})
				return err
			},
		},
		"signatureHelp": {
			wantMethod: MethodTextDocumentSignatureHelp,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.SignatureHelp(ctx, &SignatureHelpParams{})
				return err
			},
		},
		"codeAction": {
			wantMethod: MethodTextDocumentCodeAction,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.CodeAction(ctx, &CodeActionParams{})
				return err
			},
		},
		"codeActionResolve": {
			wantMethod: MethodCodeActionResolve,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.CodeActionResolve(ctx, &CodeAction{})
				return err
			},
		},
		"documentColor": {
			wantMethod: MethodTextDocumentDocumentColor,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.DocumentColor(ctx, &DocumentColorParams{})
				return err
			},
		},
		"colorPresentation": {
			wantMethod: MethodTextDocumentColorPresentation,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.ColorPresentation(ctx, &ColorPresentationParams{})
				return err
			},
		},
		"formatting": {
			wantMethod: MethodTextDocumentFormatting,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Formatting(ctx, &DocumentFormattingParams{})
				return err
			},
		},
		"rangeFormatting": {
			wantMethod: MethodTextDocumentRangeFormatting,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.RangeFormatting(ctx, &DocumentRangeFormattingParams{})
				return err
			},
		},
		"rangesFormatting": {
			wantMethod: MethodTextDocumentRangesFormatting,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.RangesFormatting(ctx, &DocumentRangesFormattingParams{})
				return err
			},
		},
		"onTypeFormatting": {
			wantMethod: MethodTextDocumentOnTypeFormatting,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.OnTypeFormatting(ctx, &DocumentOnTypeFormattingParams{})
				return err
			},
		},
		"rename": {
			wantMethod: MethodTextDocumentRename,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Rename(ctx, &RenameParams{})
				return err
			},
		},
		"prepareRename": {
			wantMethod: MethodTextDocumentPrepareRename,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.PrepareRename(ctx, &PrepareRenameParams{})
				return err
			},
		},
		"linkedEditingRange": {
			wantMethod: MethodTextDocumentLinkedEditingRange,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.LinkedEditingRange(ctx, &LinkedEditingRangeParams{})
				return err
			},
		},
		"inlineCompletion": {
			wantMethod: MethodTextDocumentInlineCompletion,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.InlineCompletion(ctx, &InlineCompletionParams{})
				return err
			},
		},
		"symbols": {
			wantMethod: MethodWorkspaceSymbol,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.Symbols(ctx, &WorkspaceSymbolParams{})
				return err
			},
		},
		"workspaceSymbolResolve": {
			wantMethod: MethodWorkspaceSymbolResolve,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.WorkspaceSymbolResolve(ctx, &WorkspaceSymbol{})
				return err
			},
		},
		"willCreateFiles": {
			wantMethod: MethodWorkspaceWillCreateFiles,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.WillCreateFiles(ctx, &CreateFilesParams{})
				return err
			},
		},
		"willRenameFiles": {
			wantMethod: MethodWorkspaceWillRenameFiles,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.WillRenameFiles(ctx, &RenameFilesParams{})
				return err
			},
		},
		"willDeleteFiles": {
			wantMethod: MethodWorkspaceWillDeleteFiles,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.WillDeleteFiles(ctx, &DeleteFilesParams{})
				return err
			},
		},
		"executeCommand": {
			wantMethod: MethodWorkspaceExecuteCommand,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.ExecuteCommand(ctx, &ExecuteCommandParams{})
				return err
			},
		},
		"textDocumentContent": {
			wantMethod: MethodWorkspaceTextDocumentContent,
			invoke: func(ctx context.Context, srv Server) error {
				_, err := srv.TextDocumentContent(ctx, &TextDocumentContentParams{})
				return err
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			conn := &fakeConn{}
			srv := ServerDispatcher(conn)
			if err := tt.invoke(t.Context(), srv); err != nil {
				t.Fatalf("invoke: %v", err)
			}
			method, _, calls, notifies := conn.snapshot()
			if method != tt.wantMethod {
				t.Errorf("method = %q, want %q", method, tt.wantMethod)
			}
			if calls != 1 {
				t.Errorf("Call count = %d, want 1", calls)
			}
			if notifies != 0 {
				t.Errorf("Notify count = %d, want 0", notifies)
			}
		})
	}
}

// TestServerRequestResultDecode covers representative request methods across the
// distinct result shapes (pointer, slice, sealed-interface union, opaque LSPAny,
// and the no-param Shutdown), asserting the canned result round-trips through the
// production codec into the caller's typed result.
func TestServerRequestResultDecode(t *testing.T) {
	t.Parallel()

	t.Run("pointer result: Hover", func(t *testing.T) {
		t.Parallel()
		conn := &fakeConn{}
		conn.setResult(&Hover{Contents: String("hi")}, nil)
		srv := ServerDispatcher(conn)
		got, err := srv.Hover(t.Context(), &HoverParams{})
		if err != nil {
			t.Fatalf("Hover: %v", err)
		}
		if got == nil {
			t.Fatal("Hover result is nil")
		}
		if c, ok := got.Contents.(String); !ok || c != "hi" {
			t.Errorf("Hover contents = %#v, want String(\"hi\")", got.Contents)
		}
	})

	t.Run("slice result: References", func(t *testing.T) {
		t.Parallel()
		conn := &fakeConn{}
		conn.setResult([]Location{{URI: "file:///a.go"}}, nil)
		srv := ServerDispatcher(conn)
		got, err := srv.References(t.Context(), &ReferenceParams{})
		if err != nil {
			t.Fatalf("References: %v", err)
		}
		if len(got) != 1 || got[0].URI != "file:///a.go" {
			t.Errorf("References = %+v, want one location at file:///a.go", got)
		}
	})

	t.Run("union result: Definition", func(t *testing.T) {
		t.Parallel()
		conn := &fakeConn{}
		conn.setResult(&Location{URI: "file:///a.go"}, nil)
		srv := ServerDispatcher(conn)
		got, err := srv.Definition(t.Context(), &DefinitionParams{})
		if err != nil {
			t.Fatalf("Definition: %v", err)
		}
		if got == nil {
			t.Fatal("Definition result is nil")
		}
	})

	t.Run("opaque result: ExecuteCommand", func(t *testing.T) {
		t.Parallel()
		conn := &fakeConn{}
		conn.setResult(LSPAny(`{"ok":true}`), nil)
		srv := ServerDispatcher(conn)
		got, err := srv.ExecuteCommand(t.Context(), &ExecuteCommandParams{Command: "do.it"})
		if err != nil {
			t.Fatalf("ExecuteCommand: %v", err)
		}
		if len(got) == 0 {
			t.Error("ExecuteCommand result is empty, want raw JSON")
		}
	})

	t.Run("no-param request: Shutdown", func(t *testing.T) {
		t.Parallel()
		conn := &fakeConn{}
		srv := ServerDispatcher(conn)
		if err := srv.Shutdown(t.Context()); err != nil {
			t.Fatalf("Shutdown: %v", err)
		}
		method, params, calls, _ := conn.snapshot()
		if method != MethodShutdown {
			t.Errorf("method = %q, want %q", method, MethodShutdown)
		}
		if params != nil {
			t.Errorf("params = %v, want nil", params)
		}
		if calls != 1 {
			t.Errorf("Call count = %d, want 1", calls)
		}
	})
}

// TestServerRequestPropagatesError asserts a Call error is returned verbatim with
// a nil typed result for a representative request method.
func TestServerRequestPropagatesError(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("boom")
	conn := &fakeConn{}
	conn.setResult(nil, wantErr)
	srv := ServerDispatcher(conn)

	got, err := srv.Hover(t.Context(), &HoverParams{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("err = %v, want %v", err, wantErr)
	}
	if got != nil {
		t.Errorf("result = %+v, want nil on error", got)
	}
}

// TestServerRequestPassThrough covers Server.Request, the escape hatch for
// non-standard methods: it issues a Call on the supplied method and returns the
// decoded result.
func TestServerRequestPassThrough(t *testing.T) {
	t.Parallel()

	conn := &fakeConn{}
	conn.setResult(LSPAny(`{"custom":1}`), nil)
	srv := ServerDispatcher(conn)

	got, err := srv.Request(t.Context(), "$/customServerMethod", LSPAny(`{"x":1}`))
	if err != nil {
		t.Fatalf("Request: %v", err)
	}
	method, _, calls, _ := conn.snapshot()
	if method != "$/customServerMethod" {
		t.Errorf("method = %q, want $/customServerMethod", method)
	}
	if calls != 1 {
		t.Errorf("Call count = %d, want 1", calls)
	}
	if got == nil {
		t.Error("Request result is nil, want decoded value")
	}
}

// TestServerRequestPassThroughError covers Server.Request's error branch: a Call
// failure is returned verbatim with a nil result.
func TestServerRequestPassThroughError(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("boom")
	conn := &fakeConn{}
	conn.setResult(nil, wantErr)
	srv := ServerDispatcher(conn)

	got, err := srv.Request(t.Context(), "$/customServerMethod", LSPAny(`{"x":1}`))
	if !errors.Is(err, wantErr) {
		t.Fatalf("err = %v, want %v", err, wantErr)
	}
	if got != nil {
		t.Errorf("result = %+v, want nil on error", got)
	}
}

// dispatchRecordingServer records which Server method the production dispatch
// path invoked. Each method signals a buffered channel so the routing test can
// confirm the request reached the intended handler. It embeds UnimplementedServer
// and overrides every method to return a minimal valid (zero) result.
type dispatchRecordingServer struct {
	UnimplementedServer

	mu     sync.Mutex
	called string
	done   chan string
}

func newDispatchRecordingServer() *dispatchRecordingServer {
	return &dispatchRecordingServer{done: make(chan string, 1)}
}

func (s *dispatchRecordingServer) record(method string) {
	s.mu.Lock()
	s.called = method
	s.mu.Unlock()
	select {
	case s.done <- method:
	default:
	}
}

func (s *dispatchRecordingServer) Initialize(context.Context, *InitializeParams) (*InitializeResult, error) {
	s.record("Initialize")
	return nil, nil
}

func (s *dispatchRecordingServer) Initialized(context.Context, *InitializedParams) error {
	s.record("Initialized")
	return nil
}

func (s *dispatchRecordingServer) Shutdown(context.Context) error {
	s.record("Shutdown")
	return nil
}

func (s *dispatchRecordingServer) Exit(context.Context) error {
	s.record("Exit")
	return nil
}

func (s *dispatchRecordingServer) SetTrace(context.Context, *SetTraceParams) error {
	s.record("SetTrace")
	return nil
}

func (s *dispatchRecordingServer) Progress(context.Context, *ProgressParams) error {
	s.record("Progress")
	return nil
}

func (s *dispatchRecordingServer) WorkDoneProgressCancel(context.Context, *WorkDoneProgressCancelParams) error {
	s.record("WorkDoneProgressCancel")
	return nil
}

func (s *dispatchRecordingServer) DidOpen(context.Context, *DidOpenTextDocumentParams) error {
	s.record("DidOpen")
	return nil
}

func (s *dispatchRecordingServer) DidChange(context.Context, *DidChangeTextDocumentParams) error {
	s.record("DidChange")
	return nil
}

func (s *dispatchRecordingServer) WillSave(context.Context, *WillSaveTextDocumentParams) error {
	s.record("WillSave")
	return nil
}

func (s *dispatchRecordingServer) WillSaveWaitUntil(context.Context, *WillSaveTextDocumentParams) ([]TextEdit, error) {
	s.record("WillSaveWaitUntil")
	return nil, nil
}

func (s *dispatchRecordingServer) DidSave(context.Context, *DidSaveTextDocumentParams) error {
	s.record("DidSave")
	return nil
}

func (s *dispatchRecordingServer) DidClose(context.Context, *DidCloseTextDocumentParams) error {
	s.record("DidClose")
	return nil
}

func (s *dispatchRecordingServer) DidOpenNotebookDocument(context.Context, *DidOpenNotebookDocumentParams) error {
	s.record("DidOpenNotebookDocument")
	return nil
}

func (s *dispatchRecordingServer) DidChangeNotebookDocument(context.Context, *DidChangeNotebookDocumentParams) error {
	s.record("DidChangeNotebookDocument")
	return nil
}

func (s *dispatchRecordingServer) DidSaveNotebookDocument(context.Context, *DidSaveNotebookDocumentParams) error {
	s.record("DidSaveNotebookDocument")
	return nil
}

func (s *dispatchRecordingServer) DidCloseNotebookDocument(context.Context, *DidCloseNotebookDocumentParams) error {
	s.record("DidCloseNotebookDocument")
	return nil
}

func (s *dispatchRecordingServer) Declaration(context.Context, *DeclarationParams) (DeclarationResult, error) {
	s.record("Declaration")
	return nil, nil
}

func (s *dispatchRecordingServer) Definition(context.Context, *DefinitionParams) (DefinitionResult, error) {
	s.record("Definition")
	return nil, nil
}

func (s *dispatchRecordingServer) TypeDefinition(context.Context, *TypeDefinitionParams) (DefinitionResult, error) {
	s.record("TypeDefinition")
	return nil, nil
}

func (s *dispatchRecordingServer) Implementation(context.Context, *ImplementationParams) (DefinitionResult, error) {
	s.record("Implementation")
	return nil, nil
}

func (s *dispatchRecordingServer) References(context.Context, *ReferenceParams) ([]Location, error) {
	s.record("References")
	return nil, nil
}

func (s *dispatchRecordingServer) PrepareCallHierarchy(context.Context, *CallHierarchyPrepareParams) ([]CallHierarchyItem, error) {
	s.record("PrepareCallHierarchy")
	return nil, nil
}

func (s *dispatchRecordingServer) IncomingCalls(context.Context, *CallHierarchyIncomingCallsParams) ([]CallHierarchyIncomingCall, error) {
	s.record("IncomingCalls")
	return nil, nil
}

func (s *dispatchRecordingServer) OutgoingCalls(context.Context, *CallHierarchyOutgoingCallsParams) ([]CallHierarchyOutgoingCall, error) {
	s.record("OutgoingCalls")
	return nil, nil
}

func (s *dispatchRecordingServer) PrepareTypeHierarchy(context.Context, *TypeHierarchyPrepareParams) ([]TypeHierarchyItem, error) {
	s.record("PrepareTypeHierarchy")
	return nil, nil
}

func (s *dispatchRecordingServer) Supertypes(context.Context, *TypeHierarchySupertypesParams) ([]TypeHierarchyItem, error) {
	s.record("Supertypes")
	return nil, nil
}

func (s *dispatchRecordingServer) Subtypes(context.Context, *TypeHierarchySubtypesParams) ([]TypeHierarchyItem, error) {
	s.record("Subtypes")
	return nil, nil
}

func (s *dispatchRecordingServer) DocumentHighlight(context.Context, *DocumentHighlightParams) ([]DocumentHighlight, error) {
	s.record("DocumentHighlight")
	return nil, nil
}

func (s *dispatchRecordingServer) DocumentLink(context.Context, *DocumentLinkParams) ([]DocumentLink, error) {
	s.record("DocumentLink")
	return nil, nil
}

func (s *dispatchRecordingServer) DocumentLinkResolve(context.Context, *DocumentLink) (*DocumentLink, error) {
	s.record("DocumentLinkResolve")
	return nil, nil
}

func (s *dispatchRecordingServer) Hover(context.Context, *HoverParams) (*Hover, error) {
	s.record("Hover")
	return nil, nil
}

func (s *dispatchRecordingServer) CodeLens(context.Context, *CodeLensParams) ([]CodeLens, error) {
	s.record("CodeLens")
	return nil, nil
}

func (s *dispatchRecordingServer) CodeLensResolve(context.Context, *CodeLens) (*CodeLens, error) {
	s.record("CodeLensResolve")
	return nil, nil
}

func (s *dispatchRecordingServer) FoldingRanges(context.Context, *FoldingRangeParams) ([]FoldingRange, error) {
	s.record("FoldingRanges")
	return nil, nil
}

func (s *dispatchRecordingServer) SelectionRange(context.Context, *SelectionRangeParams) ([]SelectionRange, error) {
	s.record("SelectionRange")
	return nil, nil
}

func (s *dispatchRecordingServer) DocumentSymbol(context.Context, *DocumentSymbolParams) (DocumentSymbolResult, error) {
	s.record("DocumentSymbol")
	return nil, nil
}

func (s *dispatchRecordingServer) SemanticTokensFull(context.Context, *SemanticTokensParams) (*SemanticTokens, error) {
	s.record("SemanticTokensFull")
	return nil, nil
}

func (s *dispatchRecordingServer) SemanticTokensFullDelta(context.Context, *SemanticTokensDeltaParams) (SemanticTokensDeltaResult, error) {
	s.record("SemanticTokensFullDelta")
	return nil, nil
}

func (s *dispatchRecordingServer) SemanticTokensRange(context.Context, *SemanticTokensRangeParams) (*SemanticTokens, error) {
	s.record("SemanticTokensRange")
	return nil, nil
}

func (s *dispatchRecordingServer) InlineValue(context.Context, *InlineValueParams) ([]InlineValue, error) {
	s.record("InlineValue")
	return nil, nil
}

func (s *dispatchRecordingServer) InlayHint(context.Context, *InlayHintParams) ([]InlayHint, error) {
	s.record("InlayHint")
	return nil, nil
}

func (s *dispatchRecordingServer) InlayHintResolve(context.Context, *InlayHint) (*InlayHint, error) {
	s.record("InlayHintResolve")
	return nil, nil
}

func (s *dispatchRecordingServer) Moniker(context.Context, *MonikerParams) ([]Moniker, error) {
	s.record("Moniker")
	return nil, nil
}

func (s *dispatchRecordingServer) Completion(context.Context, *CompletionParams) (CompletionResult, error) {
	s.record("Completion")
	return nil, nil
}

func (s *dispatchRecordingServer) CompletionResolve(context.Context, *CompletionItem) (*CompletionItem, error) {
	s.record("CompletionResolve")
	return nil, nil
}

func (s *dispatchRecordingServer) Diagnostic(context.Context, *DocumentDiagnosticParams) (DocumentDiagnosticReport, error) {
	s.record("Diagnostic")
	return nil, nil
}

func (s *dispatchRecordingServer) DiagnosticWorkspace(context.Context, *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error) {
	s.record("DiagnosticWorkspace")
	return nil, nil
}

func (s *dispatchRecordingServer) SignatureHelp(context.Context, *SignatureHelpParams) (*SignatureHelp, error) {
	s.record("SignatureHelp")
	return nil, nil
}

func (s *dispatchRecordingServer) CodeAction(context.Context, *CodeActionParams) ([]CommandOrCodeAction, error) {
	s.record("CodeAction")
	return nil, nil
}

func (s *dispatchRecordingServer) CodeActionResolve(context.Context, *CodeAction) (*CodeAction, error) {
	s.record("CodeActionResolve")
	return nil, nil
}

func (s *dispatchRecordingServer) DocumentColor(context.Context, *DocumentColorParams) ([]ColorInformation, error) {
	s.record("DocumentColor")
	return nil, nil
}

func (s *dispatchRecordingServer) ColorPresentation(context.Context, *ColorPresentationParams) ([]ColorPresentation, error) {
	s.record("ColorPresentation")
	return nil, nil
}

func (s *dispatchRecordingServer) Formatting(context.Context, *DocumentFormattingParams) ([]TextEdit, error) {
	s.record("Formatting")
	return nil, nil
}

func (s *dispatchRecordingServer) RangeFormatting(context.Context, *DocumentRangeFormattingParams) ([]TextEdit, error) {
	s.record("RangeFormatting")
	return nil, nil
}

func (s *dispatchRecordingServer) RangesFormatting(context.Context, *DocumentRangesFormattingParams) ([]TextEdit, error) {
	s.record("RangesFormatting")
	return nil, nil
}

func (s *dispatchRecordingServer) OnTypeFormatting(context.Context, *DocumentOnTypeFormattingParams) ([]TextEdit, error) {
	s.record("OnTypeFormatting")
	return nil, nil
}

func (s *dispatchRecordingServer) Rename(context.Context, *RenameParams) (*WorkspaceEdit, error) {
	s.record("Rename")
	return nil, nil
}

func (s *dispatchRecordingServer) PrepareRename(context.Context, *PrepareRenameParams) (PrepareRenameResult, error) {
	s.record("PrepareRename")
	return nil, nil
}

func (s *dispatchRecordingServer) LinkedEditingRange(context.Context, *LinkedEditingRangeParams) (*LinkedEditingRanges, error) {
	s.record("LinkedEditingRange")
	return nil, nil
}

func (s *dispatchRecordingServer) InlineCompletion(context.Context, *InlineCompletionParams) (InlineCompletionResult, error) {
	s.record("InlineCompletion")
	return nil, nil
}

func (s *dispatchRecordingServer) Symbols(context.Context, *WorkspaceSymbolParams) (WorkspaceSymbolResult, error) {
	s.record("Symbols")
	return nil, nil
}

func (s *dispatchRecordingServer) WorkspaceSymbolResolve(context.Context, *WorkspaceSymbol) (*WorkspaceSymbol, error) {
	s.record("WorkspaceSymbolResolve")
	return nil, nil
}

func (s *dispatchRecordingServer) DidChangeConfiguration(context.Context, *DidChangeConfigurationParams) error {
	s.record("DidChangeConfiguration")
	return nil
}

func (s *dispatchRecordingServer) DidChangeWorkspaceFolders(context.Context, *DidChangeWorkspaceFoldersParams) error {
	s.record("DidChangeWorkspaceFolders")
	return nil
}

func (s *dispatchRecordingServer) WillCreateFiles(context.Context, *CreateFilesParams) (*WorkspaceEdit, error) {
	s.record("WillCreateFiles")
	return nil, nil
}

func (s *dispatchRecordingServer) WillRenameFiles(context.Context, *RenameFilesParams) (*WorkspaceEdit, error) {
	s.record("WillRenameFiles")
	return nil, nil
}

func (s *dispatchRecordingServer) WillDeleteFiles(context.Context, *DeleteFilesParams) (*WorkspaceEdit, error) {
	s.record("WillDeleteFiles")
	return nil, nil
}

func (s *dispatchRecordingServer) DidCreateFiles(context.Context, *CreateFilesParams) error {
	s.record("DidCreateFiles")
	return nil
}

func (s *dispatchRecordingServer) DidRenameFiles(context.Context, *RenameFilesParams) error {
	s.record("DidRenameFiles")
	return nil
}

func (s *dispatchRecordingServer) DidDeleteFiles(context.Context, *DeleteFilesParams) error {
	s.record("DidDeleteFiles")
	return nil
}

func (s *dispatchRecordingServer) DidChangeWatchedFiles(context.Context, *DidChangeWatchedFilesParams) error {
	s.record("DidChangeWatchedFiles")
	return nil
}

func (s *dispatchRecordingServer) ExecuteCommand(context.Context, *ExecuteCommandParams) (LSPAny, error) {
	s.record("ExecuteCommand")
	return nil, nil
}

func (s *dispatchRecordingServer) TextDocumentContent(context.Context, *TextDocumentContentParams) (*TextDocumentContentResult, error) {
	s.record("TextDocumentContent")
	return nil, nil
}

// newServerHandlerConnPair wires a caller conn to a served conn running the
// production ServerHandler for srv, mirroring newClientHandlerConnPair.
func newServerHandlerConnPair(ctx context.Context, t *testing.T, srv Server, fallback jsonrpc2.Handler) (caller, served jsonrpc2.Conn) {
	t.Helper()

	callerEnd, servedEnd := net.Pipe()
	caller = jsonrpc2.NewConn(jsonrpc2.NewStream(callerEnd), jsonrpc2.WithCodec(lspCodec{}))
	served = jsonrpc2.NewConn(jsonrpc2.NewStream(servedEnd), jsonrpc2.WithCodec(lspCodec{}))
	caller.Go(ctx, jsonrpc2.MethodNotFoundHandler)
	served.Go(ctx, ServerHandler(srv, fallback))

	return caller, served
}

// TestServerDispatchRouting fires one wire message per standard server method
// through the production ServerHandler over an in-memory pipe and asserts
// serverDispatch routed it to the matching Server method. This exercises every
// case arm of serverDispatch the way the transport actually builds requests.
func TestServerDispatchRouting(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		method     string
		params     any
		isCall     bool
		wantMethod string
	}{
		"initialize":                {MethodInitialize, &InitializeParams{}, true, "Initialize"},
		"initialized":               {MethodInitialized, &InitializedParams{}, false, "Initialized"},
		"shutdown":                  {MethodShutdown, nil, true, "Shutdown"},
		"exit":                      {MethodExit, nil, false, "Exit"},
		"setTrace":                  {MethodSetTrace, &SetTraceParams{}, false, "SetTrace"},
		"progress":                  {MethodProgress, &ProgressParams{}, false, "Progress"},
		"workDoneProgressCancel":    {MethodWindowWorkDoneProgressCancel, &WorkDoneProgressCancelParams{}, false, "WorkDoneProgressCancel"},
		"didOpen":                   {MethodTextDocumentDidOpen, &DidOpenTextDocumentParams{}, false, "DidOpen"},
		"didChange":                 {MethodTextDocumentDidChange, &DidChangeTextDocumentParams{}, false, "DidChange"},
		"willSave":                  {MethodTextDocumentWillSave, &WillSaveTextDocumentParams{}, false, "WillSave"},
		"willSaveWaitUntil":         {MethodTextDocumentWillSaveWaitUntil, &WillSaveTextDocumentParams{}, true, "WillSaveWaitUntil"},
		"didSave":                   {MethodTextDocumentDidSave, &DidSaveTextDocumentParams{}, false, "DidSave"},
		"didClose":                  {MethodTextDocumentDidClose, &DidCloseTextDocumentParams{}, false, "DidClose"},
		"didOpenNotebookDocument":   {MethodNotebookDocumentDidOpen, &DidOpenNotebookDocumentParams{}, false, "DidOpenNotebookDocument"},
		"didChangeNotebookDocument": {MethodNotebookDocumentDidChange, &DidChangeNotebookDocumentParams{}, false, "DidChangeNotebookDocument"},
		"didSaveNotebookDocument":   {MethodNotebookDocumentDidSave, &DidSaveNotebookDocumentParams{}, false, "DidSaveNotebookDocument"},
		"didCloseNotebookDocument":  {MethodNotebookDocumentDidClose, &DidCloseNotebookDocumentParams{}, false, "DidCloseNotebookDocument"},
		"declaration":               {MethodTextDocumentDeclaration, &DeclarationParams{}, true, "Declaration"},
		"definition":                {MethodTextDocumentDefinition, &DefinitionParams{}, true, "Definition"},
		"typeDefinition":            {MethodTextDocumentTypeDefinition, &TypeDefinitionParams{}, true, "TypeDefinition"},
		"implementation":            {MethodTextDocumentImplementation, &ImplementationParams{}, true, "Implementation"},
		"references":                {MethodTextDocumentReferences, &ReferenceParams{}, true, "References"},
		"prepareCallHierarchy":      {MethodTextDocumentPrepareCallHierarchy, &CallHierarchyPrepareParams{}, true, "PrepareCallHierarchy"},
		"incomingCalls":             {MethodCallHierarchyIncomingCalls, &CallHierarchyIncomingCallsParams{}, true, "IncomingCalls"},
		"outgoingCalls":             {MethodCallHierarchyOutgoingCalls, &CallHierarchyOutgoingCallsParams{}, true, "OutgoingCalls"},
		"prepareTypeHierarchy":      {MethodTextDocumentPrepareTypeHierarchy, &TypeHierarchyPrepareParams{}, true, "PrepareTypeHierarchy"},
		"supertypes":                {MethodTypeHierarchySupertypes, &TypeHierarchySupertypesParams{}, true, "Supertypes"},
		"subtypes":                  {MethodTypeHierarchySubtypes, &TypeHierarchySubtypesParams{}, true, "Subtypes"},
		"documentHighlight":         {MethodTextDocumentDocumentHighlight, &DocumentHighlightParams{}, true, "DocumentHighlight"},
		"documentLink":              {MethodTextDocumentDocumentLink, &DocumentLinkParams{}, true, "DocumentLink"},
		"documentLinkResolve":       {MethodDocumentLinkResolve, &DocumentLink{}, true, "DocumentLinkResolve"},
		"hover":                     {MethodTextDocumentHover, &HoverParams{}, true, "Hover"},
		"codeLens":                  {MethodTextDocumentCodeLens, &CodeLensParams{}, true, "CodeLens"},
		"codeLensResolve":           {MethodCodeLensResolve, &CodeLens{}, true, "CodeLensResolve"},
		"foldingRanges":             {MethodTextDocumentFoldingRange, &FoldingRangeParams{}, true, "FoldingRanges"},
		"selectionRange":            {MethodTextDocumentSelectionRange, &SelectionRangeParams{}, true, "SelectionRange"},
		"documentSymbol":            {MethodTextDocumentDocumentSymbol, &DocumentSymbolParams{}, true, "DocumentSymbol"},
		"semanticTokensFull":        {MethodTextDocumentSemanticTokensFull, &SemanticTokensParams{}, true, "SemanticTokensFull"},
		"semanticTokensFullDelta":   {MethodTextDocumentSemanticTokensFullDelta, &SemanticTokensDeltaParams{}, true, "SemanticTokensFullDelta"},
		"semanticTokensRange":       {MethodTextDocumentSemanticTokensRange, &SemanticTokensRangeParams{}, true, "SemanticTokensRange"},
		"inlineValue":               {MethodTextDocumentInlineValue, &InlineValueParams{}, true, "InlineValue"},
		"inlayHint":                 {MethodTextDocumentInlayHint, &InlayHintParams{}, true, "InlayHint"},
		"inlayHintResolve":          {MethodInlayHintResolve, &InlayHint{}, true, "InlayHintResolve"},
		"moniker":                   {MethodTextDocumentMoniker, &MonikerParams{}, true, "Moniker"},
		"completion":                {MethodTextDocumentCompletion, &CompletionParams{}, true, "Completion"},
		"completionResolve":         {MethodCompletionItemResolve, &CompletionItem{}, true, "CompletionResolve"},
		"diagnostic":                {MethodTextDocumentDiagnostic, &DocumentDiagnosticParams{}, true, "Diagnostic"},
		"diagnosticWorkspace":       {MethodWorkspaceDiagnostic, &WorkspaceDiagnosticParams{}, true, "DiagnosticWorkspace"},
		"signatureHelp":             {MethodTextDocumentSignatureHelp, &SignatureHelpParams{}, true, "SignatureHelp"},
		"codeAction":                {MethodTextDocumentCodeAction, &CodeActionParams{}, true, "CodeAction"},
		"codeActionResolve":         {MethodCodeActionResolve, &CodeAction{}, true, "CodeActionResolve"},
		"documentColor":             {MethodTextDocumentDocumentColor, &DocumentColorParams{}, true, "DocumentColor"},
		"colorPresentation":         {MethodTextDocumentColorPresentation, &ColorPresentationParams{}, true, "ColorPresentation"},
		"formatting":                {MethodTextDocumentFormatting, &DocumentFormattingParams{}, true, "Formatting"},
		"rangeFormatting":           {MethodTextDocumentRangeFormatting, &DocumentRangeFormattingParams{}, true, "RangeFormatting"},
		"rangesFormatting":          {MethodTextDocumentRangesFormatting, &DocumentRangesFormattingParams{}, true, "RangesFormatting"},
		"onTypeFormatting":          {MethodTextDocumentOnTypeFormatting, &DocumentOnTypeFormattingParams{}, true, "OnTypeFormatting"},
		"rename":                    {MethodTextDocumentRename, &RenameParams{}, true, "Rename"},
		"prepareRename":             {MethodTextDocumentPrepareRename, &PrepareRenameParams{}, true, "PrepareRename"},
		"linkedEditingRange":        {MethodTextDocumentLinkedEditingRange, &LinkedEditingRangeParams{}, true, "LinkedEditingRange"},
		"inlineCompletion":          {MethodTextDocumentInlineCompletion, &InlineCompletionParams{}, true, "InlineCompletion"},
		"symbols":                   {MethodWorkspaceSymbol, &WorkspaceSymbolParams{}, true, "Symbols"},
		"workspaceSymbolResolve":    {MethodWorkspaceSymbolResolve, &WorkspaceSymbol{}, true, "WorkspaceSymbolResolve"},
		"didChangeConfiguration":    {MethodWorkspaceDidChangeConfiguration, &DidChangeConfigurationParams{}, false, "DidChangeConfiguration"},
		"didChangeWorkspaceFolders": {MethodWorkspaceDidChangeWorkspaceFolders, &DidChangeWorkspaceFoldersParams{}, false, "DidChangeWorkspaceFolders"},
		"willCreateFiles":           {MethodWorkspaceWillCreateFiles, &CreateFilesParams{}, true, "WillCreateFiles"},
		"willRenameFiles":           {MethodWorkspaceWillRenameFiles, &RenameFilesParams{}, true, "WillRenameFiles"},
		"willDeleteFiles":           {MethodWorkspaceWillDeleteFiles, &DeleteFilesParams{}, true, "WillDeleteFiles"},
		"didCreateFiles":            {MethodWorkspaceDidCreateFiles, &CreateFilesParams{}, false, "DidCreateFiles"},
		"didRenameFiles":            {MethodWorkspaceDidRenameFiles, &RenameFilesParams{}, false, "DidRenameFiles"},
		"didDeleteFiles":            {MethodWorkspaceDidDeleteFiles, &DeleteFilesParams{}, false, "DidDeleteFiles"},
		"didChangeWatchedFiles":     {MethodWorkspaceDidChangeWatchedFiles, &DidChangeWatchedFilesParams{}, false, "DidChangeWatchedFiles"},
		"executeCommand":            {MethodWorkspaceExecuteCommand, &ExecuteCommandParams{}, true, "ExecuteCommand"},
		"textDocumentContent":       {MethodWorkspaceTextDocumentContent, &TextDocumentContentParams{}, true, "TextDocumentContent"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
			defer cancel()

			srv := newDispatchRecordingServer()
			caller, served := newServerHandlerConnPair(ctx, t, srv, jsonrpc2.MethodNotFoundHandler)
			defer closeJSONRPCConns(t, caller, served)

			if tt.isCall {
				if _, err := caller.Call(ctx, tt.method, tt.params, nil); err != nil {
					t.Fatalf("Call(%s): %v", tt.method, err)
				}
			} else {
				if err := caller.Notify(ctx, tt.method, tt.params); err != nil {
					t.Fatalf("Notify(%s): %v", tt.method, err)
				}
			}

			select {
			case got := <-srv.done:
				if got != tt.wantMethod {
					t.Errorf("dispatched to %q, want %q", got, tt.wantMethod)
				}
			case <-ctx.Done():
				t.Fatalf("server method %q was not dispatched: %v", tt.wantMethod, ctx.Err())
			}
		})
	}
}

// TestServerHandlerNonStandardMethodRoutesToRequest covers the ServerHandler
// fall-through: a method not in the standard set is decoded as opaque LSPAny and
// routed to Server.Request.
func TestServerHandlerNonStandardMethodRoutesToRequest(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	srv := newRequestRoutingServer()
	caller, served := newServerHandlerConnPair(ctx, t, srv, jsonrpc2.MethodNotFoundHandler)
	defer closeJSONRPCConns(t, caller, served)

	const method = "$/customServerRequest"
	var got jsonrpc2.RawMessage
	if _, err := caller.Call(ctx, method, jsonrpc2.RawMessage(`{"x":1}`), &got); err != nil {
		t.Fatalf("Call(%s): %v", method, err)
	}
	select {
	case gotMethod := <-srv.done:
		if gotMethod != method {
			t.Errorf("Request method = %q, want %q", gotMethod, method)
		}
	case <-ctx.Done():
		t.Fatalf("Server.Request was not called: %v", ctx.Err())
	}
}

// requestRoutingServer records the method passed to Server.Request, the
// non-standard-method escape hatch in ServerHandler.
type requestRoutingServer struct {
	UnimplementedServer

	done chan string
}

func newRequestRoutingServer() *requestRoutingServer {
	return &requestRoutingServer{done: make(chan string, 1)}
}

func (s *requestRoutingServer) Request(_ context.Context, method string, _ any) (any, error) {
	select {
	case s.done <- method:
	default:
	}

	return LSPAny(`{"ok":true}`), nil
}

// TestServerDispatchCanceledContext asserts ServerHandler short-circuits to
// ErrRequestCancelled when the context is already done, without invoking the
// server.
func TestServerDispatchCanceledContext(t *testing.T) {
	t.Parallel()

	srv := newDispatchRecordingServer()
	handler := ServerHandler(srv, jsonrpc2.MethodNotFoundHandler)
	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	_, err := handler(ctx, &jsonrpc2.Request{})
	if !errors.Is(err, ErrRequestCancelled) {
		t.Fatalf("err = %v, want %v", err, ErrRequestCancelled)
	}
	if srv.called != "" {
		t.Errorf("server was invoked (%q) for canceled context", srv.called)
	}
}

// TestServerRequestMethodsPropagateError drives every result-bearing server
// request method against a conn that returns an error, asserting the error
// propagates verbatim. This covers each method's error-return branch.
func TestServerRequestMethodsPropagateError(t *testing.T) {
	t.Parallel()

	invokers := map[string]func(ctx context.Context, srv Server) error{
		"initialize": func(ctx context.Context, srv Server) error {
			_, err := srv.Initialize(ctx, &InitializeParams{})
			return err
		},
		"willSaveWaitUntil": func(ctx context.Context, srv Server) error {
			_, err := srv.WillSaveWaitUntil(ctx, &WillSaveTextDocumentParams{})
			return err
		},
		"declaration": func(ctx context.Context, srv Server) error {
			_, err := srv.Declaration(ctx, &DeclarationParams{})
			return err
		},
		"definition": func(ctx context.Context, srv Server) error {
			_, err := srv.Definition(ctx, &DefinitionParams{})
			return err
		},
		"typeDefinition": func(ctx context.Context, srv Server) error {
			_, err := srv.TypeDefinition(ctx, &TypeDefinitionParams{})
			return err
		},
		"implementation": func(ctx context.Context, srv Server) error {
			_, err := srv.Implementation(ctx, &ImplementationParams{})
			return err
		},
		"references": func(ctx context.Context, srv Server) error {
			_, err := srv.References(ctx, &ReferenceParams{})
			return err
		},
		"prepareCallHierarchy": func(ctx context.Context, srv Server) error {
			_, err := srv.PrepareCallHierarchy(ctx, &CallHierarchyPrepareParams{})
			return err
		},
		"incomingCalls": func(ctx context.Context, srv Server) error {
			_, err := srv.IncomingCalls(ctx, &CallHierarchyIncomingCallsParams{})
			return err
		},
		"outgoingCalls": func(ctx context.Context, srv Server) error {
			_, err := srv.OutgoingCalls(ctx, &CallHierarchyOutgoingCallsParams{})
			return err
		},
		"prepareTypeHierarchy": func(ctx context.Context, srv Server) error {
			_, err := srv.PrepareTypeHierarchy(ctx, &TypeHierarchyPrepareParams{})
			return err
		},
		"supertypes": func(ctx context.Context, srv Server) error {
			_, err := srv.Supertypes(ctx, &TypeHierarchySupertypesParams{})
			return err
		},
		"subtypes": func(ctx context.Context, srv Server) error {
			_, err := srv.Subtypes(ctx, &TypeHierarchySubtypesParams{})
			return err
		},
		"documentHighlight": func(ctx context.Context, srv Server) error {
			_, err := srv.DocumentHighlight(ctx, &DocumentHighlightParams{})
			return err
		},
		"documentLink": func(ctx context.Context, srv Server) error {
			_, err := srv.DocumentLink(ctx, &DocumentLinkParams{})
			return err
		},
		"documentLinkResolve": func(ctx context.Context, srv Server) error {
			_, err := srv.DocumentLinkResolve(ctx, &DocumentLink{})
			return err
		},
		"hover": func(ctx context.Context, srv Server) error {
			_, err := srv.Hover(ctx, &HoverParams{})
			return err
		},
		"codeLens": func(ctx context.Context, srv Server) error {
			_, err := srv.CodeLens(ctx, &CodeLensParams{})
			return err
		},
		"codeLensResolve": func(ctx context.Context, srv Server) error {
			_, err := srv.CodeLensResolve(ctx, &CodeLens{})
			return err
		},
		"foldingRanges": func(ctx context.Context, srv Server) error {
			_, err := srv.FoldingRanges(ctx, &FoldingRangeParams{})
			return err
		},
		"selectionRange": func(ctx context.Context, srv Server) error {
			_, err := srv.SelectionRange(ctx, &SelectionRangeParams{})
			return err
		},
		"documentSymbol": func(ctx context.Context, srv Server) error {
			_, err := srv.DocumentSymbol(ctx, &DocumentSymbolParams{})
			return err
		},
		"semanticTokensFull": func(ctx context.Context, srv Server) error {
			_, err := srv.SemanticTokensFull(ctx, &SemanticTokensParams{})
			return err
		},
		"semanticTokensFullDelta": func(ctx context.Context, srv Server) error {
			_, err := srv.SemanticTokensFullDelta(ctx, &SemanticTokensDeltaParams{})
			return err
		},
		"semanticTokensRange": func(ctx context.Context, srv Server) error {
			_, err := srv.SemanticTokensRange(ctx, &SemanticTokensRangeParams{})
			return err
		},
		"inlineValue": func(ctx context.Context, srv Server) error {
			_, err := srv.InlineValue(ctx, &InlineValueParams{})
			return err
		},
		"inlayHint": func(ctx context.Context, srv Server) error {
			_, err := srv.InlayHint(ctx, &InlayHintParams{})
			return err
		},
		"inlayHintResolve": func(ctx context.Context, srv Server) error {
			_, err := srv.InlayHintResolve(ctx, &InlayHint{})
			return err
		},
		"moniker": func(ctx context.Context, srv Server) error {
			_, err := srv.Moniker(ctx, &MonikerParams{})
			return err
		},
		"completion": func(ctx context.Context, srv Server) error {
			_, err := srv.Completion(ctx, &CompletionParams{})
			return err
		},
		"completionResolve": func(ctx context.Context, srv Server) error {
			_, err := srv.CompletionResolve(ctx, &CompletionItem{})
			return err
		},
		"diagnostic": func(ctx context.Context, srv Server) error {
			_, err := srv.Diagnostic(ctx, &DocumentDiagnosticParams{})
			return err
		},
		"diagnosticWorkspace": func(ctx context.Context, srv Server) error {
			_, err := srv.DiagnosticWorkspace(ctx, &WorkspaceDiagnosticParams{})
			return err
		},
		"signatureHelp": func(ctx context.Context, srv Server) error {
			_, err := srv.SignatureHelp(ctx, &SignatureHelpParams{})
			return err
		},
		"codeAction": func(ctx context.Context, srv Server) error {
			_, err := srv.CodeAction(ctx, &CodeActionParams{})
			return err
		},
		"codeActionResolve": func(ctx context.Context, srv Server) error {
			_, err := srv.CodeActionResolve(ctx, &CodeAction{})
			return err
		},
		"documentColor": func(ctx context.Context, srv Server) error {
			_, err := srv.DocumentColor(ctx, &DocumentColorParams{})
			return err
		},
		"colorPresentation": func(ctx context.Context, srv Server) error {
			_, err := srv.ColorPresentation(ctx, &ColorPresentationParams{})
			return err
		},
		"formatting": func(ctx context.Context, srv Server) error {
			_, err := srv.Formatting(ctx, &DocumentFormattingParams{})
			return err
		},
		"rangeFormatting": func(ctx context.Context, srv Server) error {
			_, err := srv.RangeFormatting(ctx, &DocumentRangeFormattingParams{})
			return err
		},
		"rangesFormatting": func(ctx context.Context, srv Server) error {
			_, err := srv.RangesFormatting(ctx, &DocumentRangesFormattingParams{})
			return err
		},
		"onTypeFormatting": func(ctx context.Context, srv Server) error {
			_, err := srv.OnTypeFormatting(ctx, &DocumentOnTypeFormattingParams{})
			return err
		},
		"rename": func(ctx context.Context, srv Server) error {
			_, err := srv.Rename(ctx, &RenameParams{})
			return err
		},
		"prepareRename": func(ctx context.Context, srv Server) error {
			_, err := srv.PrepareRename(ctx, &PrepareRenameParams{})
			return err
		},
		"linkedEditingRange": func(ctx context.Context, srv Server) error {
			_, err := srv.LinkedEditingRange(ctx, &LinkedEditingRangeParams{})
			return err
		},
		"inlineCompletion": func(ctx context.Context, srv Server) error {
			_, err := srv.InlineCompletion(ctx, &InlineCompletionParams{})
			return err
		},
		"symbols": func(ctx context.Context, srv Server) error {
			_, err := srv.Symbols(ctx, &WorkspaceSymbolParams{})
			return err
		},
		"workspaceSymbolResolve": func(ctx context.Context, srv Server) error {
			_, err := srv.WorkspaceSymbolResolve(ctx, &WorkspaceSymbol{})
			return err
		},
		"willCreateFiles": func(ctx context.Context, srv Server) error {
			_, err := srv.WillCreateFiles(ctx, &CreateFilesParams{})
			return err
		},
		"willRenameFiles": func(ctx context.Context, srv Server) error {
			_, err := srv.WillRenameFiles(ctx, &RenameFilesParams{})
			return err
		},
		"willDeleteFiles": func(ctx context.Context, srv Server) error {
			_, err := srv.WillDeleteFiles(ctx, &DeleteFilesParams{})
			return err
		},
		"executeCommand": func(ctx context.Context, srv Server) error {
			_, err := srv.ExecuteCommand(ctx, &ExecuteCommandParams{})
			return err
		},
		"textDocumentContent": func(ctx context.Context, srv Server) error {
			_, err := srv.TextDocumentContent(ctx, &TextDocumentContentParams{})
			return err
		},
	}
	wantErr := errors.New("boom")
	for name, invoke := range invokers {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			conn := &fakeConn{}
			conn.setResult(nil, wantErr)
			srv := ServerDispatcher(conn)
			if err := invoke(t.Context(), srv); !errors.Is(err, wantErr) {
				t.Fatalf("err = %v, want %v", err, wantErr)
			}
		})
	}
}

// TestServerDispatchMalformedParams asserts serverDispatch reports a parse error
// (wrapping jsonrpc2.ErrParse) for a request whose params fail to decode, without
// invoking the server method. It is driven through ServerHandler so the request
// is built by the production scanner.
func TestServerDispatchMalformedParams(t *testing.T) {
	t.Parallel()

	// Every standard request method whose params are a JSON object: a JSON array
	// is well-formed JSON that the framer accepts but the per-case struct decode
	// rejects, exercising that case's replyParseError branch.
	methods := map[string]string{
		"initialize":              MethodInitialize,
		"willSaveWaitUntil":       MethodTextDocumentWillSaveWaitUntil,
		"declaration":             MethodTextDocumentDeclaration,
		"definition":              MethodTextDocumentDefinition,
		"typeDefinition":          MethodTextDocumentTypeDefinition,
		"implementation":          MethodTextDocumentImplementation,
		"references":              MethodTextDocumentReferences,
		"prepareCallHierarchy":    MethodTextDocumentPrepareCallHierarchy,
		"incomingCalls":           MethodCallHierarchyIncomingCalls,
		"outgoingCalls":           MethodCallHierarchyOutgoingCalls,
		"prepareTypeHierarchy":    MethodTextDocumentPrepareTypeHierarchy,
		"supertypes":              MethodTypeHierarchySupertypes,
		"subtypes":                MethodTypeHierarchySubtypes,
		"documentHighlight":       MethodTextDocumentDocumentHighlight,
		"documentLink":            MethodTextDocumentDocumentLink,
		"documentLinkResolve":     MethodDocumentLinkResolve,
		"hover":                   MethodTextDocumentHover,
		"codeLens":                MethodTextDocumentCodeLens,
		"codeLensResolve":         MethodCodeLensResolve,
		"foldingRanges":           MethodTextDocumentFoldingRange,
		"selectionRange":          MethodTextDocumentSelectionRange,
		"documentSymbol":          MethodTextDocumentDocumentSymbol,
		"semanticTokensFull":      MethodTextDocumentSemanticTokensFull,
		"semanticTokensFullDelta": MethodTextDocumentSemanticTokensFullDelta,
		"semanticTokensRange":     MethodTextDocumentSemanticTokensRange,
		"inlineValue":             MethodTextDocumentInlineValue,
		"inlayHint":               MethodTextDocumentInlayHint,
		"inlayHintResolve":        MethodInlayHintResolve,
		"moniker":                 MethodTextDocumentMoniker,
		"completion":              MethodTextDocumentCompletion,
		"completionResolve":       MethodCompletionItemResolve,
		"diagnostic":              MethodTextDocumentDiagnostic,
		"diagnosticWorkspace":     MethodWorkspaceDiagnostic,
		"signatureHelp":           MethodTextDocumentSignatureHelp,
		"codeAction":              MethodTextDocumentCodeAction,
		"codeActionResolve":       MethodCodeActionResolve,
		"documentColor":           MethodTextDocumentDocumentColor,
		"colorPresentation":       MethodTextDocumentColorPresentation,
		"formatting":              MethodTextDocumentFormatting,
		"rangeFormatting":         MethodTextDocumentRangeFormatting,
		"rangesFormatting":        MethodTextDocumentRangesFormatting,
		"onTypeFormatting":        MethodTextDocumentOnTypeFormatting,
		"rename":                  MethodTextDocumentRename,
		"prepareRename":           MethodTextDocumentPrepareRename,
		"linkedEditingRange":      MethodTextDocumentLinkedEditingRange,
		"inlineCompletion":        MethodTextDocumentInlineCompletion,
		"symbols":                 MethodWorkspaceSymbol,
		"workspaceSymbolResolve":  MethodWorkspaceSymbolResolve,
		"willCreateFiles":         MethodWorkspaceWillCreateFiles,
		"willRenameFiles":         MethodWorkspaceWillRenameFiles,
		"willDeleteFiles":         MethodWorkspaceWillDeleteFiles,
		"executeCommand":          MethodWorkspaceExecuteCommand,
		"textDocumentContent":     MethodWorkspaceTextDocumentContent,
	}
	for name, method := range methods {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
			defer cancel()

			srv := newDispatchRecordingServer()
			caller, served := newServerHandlerConnPair(ctx, t, srv, jsonrpc2.MethodNotFoundHandler)
			defer closeJSONRPCConns(t, caller, served)

			_, err := caller.Call(ctx, method, jsonrpc2.RawMessage(`[1,2,3]`), nil)
			if !errors.Is(err, jsonrpc2.ErrParse) {
				t.Fatalf("Call(%s) malformed params error = %v, want wrapping %v", method, err, jsonrpc2.ErrParse)
			}
			select {
			case got := <-srv.done:
				t.Fatalf("server method %q was invoked despite malformed params", got)
			case <-time.After(100 * time.Millisecond):
			}
		})
	}
}
