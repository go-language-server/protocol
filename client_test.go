// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	gocmp "github.com/google/go-cmp/cmp"

	"go.lsp.dev/jsonrpc2"
)

// TestClientNotificationMethods drives every server->client notification sending
// method through the *client dispatcher and asserts it lands on the matching
// method constant via Conn.Notify (not Call).
func TestClientNotificationMethods(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		wantMethod string
		invoke     func(ctx context.Context, c Client) error
	}{
		"progress": {
			wantMethod: MethodProgress,
			invoke: func(ctx context.Context, c Client) error {
				return c.Progress(ctx, &ProgressParams{Token: String("t")})
			},
		},
		"logTrace": {
			wantMethod: MethodLogTrace,
			invoke: func(ctx context.Context, c Client) error {
				return c.LogTrace(ctx, &LogTraceParams{Message: "m"})
			},
		},
		"showMessage": {
			wantMethod: MethodWindowShowMessage,
			invoke: func(ctx context.Context, c Client) error {
				return c.ShowMessage(ctx, &ShowMessageParams{Type: MessageTypeInfo, Message: "hi"})
			},
		},
		"logMessage": {
			wantMethod: MethodWindowLogMessage,
			invoke: func(ctx context.Context, c Client) error {
				return c.LogMessage(ctx, &LogMessageParams{Type: MessageTypeLog, Message: "log"})
			},
		},
		"telemetry": {
			wantMethod: MethodTelemetryEvent,
			invoke: func(ctx context.Context, c Client) error {
				return c.Telemetry(ctx, LSPAny(`{"k":"v"}`))
			},
		},
		"publishDiagnostics": {
			wantMethod: MethodTextDocumentPublishDiagnostics,
			invoke: func(ctx context.Context, c Client) error {
				return c.PublishDiagnostics(ctx, &PublishDiagnosticsParams{URI: "file:///a.go"})
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			conn := &fakeConn{}
			c := ClientDispatcher(conn)
			if err := tt.invoke(t.Context(), c); err != nil {
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

// TestClientRequestMethodsNoResult drives every server->client request sending
// method that discards its result (passes nil to Call) and asserts it lands on
// the matching method constant via Conn.Call.
func TestClientRequestMethodsNoResult(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		wantMethod string
		invoke     func(ctx context.Context, c Client) error
	}{
		"registerCapability": {
			wantMethod: MethodClientRegisterCapability,
			invoke: func(ctx context.Context, c Client) error {
				return c.RegisterCapability(ctx, &RegistrationParams{})
			},
		},
		"unregisterCapability": {
			wantMethod: MethodClientUnregisterCapability,
			invoke: func(ctx context.Context, c Client) error {
				return c.UnregisterCapability(ctx, &UnregistrationParams{})
			},
		},
		"workDoneProgressCreate": {
			wantMethod: MethodWindowWorkDoneProgressCreate,
			invoke: func(ctx context.Context, c Client) error {
				return c.WorkDoneProgressCreate(ctx, &WorkDoneProgressCreateParams{Token: String("t")})
			},
		},
		"codeLensRefresh": {
			wantMethod: MethodWorkspaceCodeLensRefresh,
			invoke:     func(ctx context.Context, c Client) error { return c.CodeLensRefresh(ctx) },
		},
		"foldingRangeRefresh": {
			wantMethod: MethodWorkspaceFoldingRangeRefresh,
			invoke:     func(ctx context.Context, c Client) error { return c.FoldingRangeRefresh(ctx) },
		},
		"semanticTokensRefresh": {
			wantMethod: MethodWorkspaceSemanticTokensRefresh,
			invoke:     func(ctx context.Context, c Client) error { return c.SemanticTokensRefresh(ctx) },
		},
		"inlineValueRefresh": {
			wantMethod: MethodWorkspaceInlineValueRefresh,
			invoke:     func(ctx context.Context, c Client) error { return c.InlineValueRefresh(ctx) },
		},
		"inlayHintRefresh": {
			wantMethod: MethodWorkspaceInlayHintRefresh,
			invoke:     func(ctx context.Context, c Client) error { return c.InlayHintRefresh(ctx) },
		},
		"diagnosticRefresh": {
			wantMethod: MethodWorkspaceDiagnosticRefresh,
			invoke:     func(ctx context.Context, c Client) error { return c.DiagnosticRefresh(ctx) },
		},
		"textDocumentContentRefresh": {
			wantMethod: MethodWorkspaceTextDocumentContentRefresh,
			invoke: func(ctx context.Context, c Client) error {
				return c.TextDocumentContentRefresh(ctx, &TextDocumentContentRefreshParams{URI: "file:///a.go"})
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			conn := &fakeConn{}
			c := ClientDispatcher(conn)
			if err := tt.invoke(t.Context(), c); err != nil {
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

// TestClientShowMessageRequestResult covers the request path that decodes a
// pointer result, asserting both the method constant and the decoded value.
func TestClientShowMessageRequestResult(t *testing.T) {
	t.Parallel()

	conn := &fakeConn{}
	conn.setResult(&MessageActionItem{Title: "Apply"}, nil)
	c := ClientDispatcher(conn)

	got, err := c.ShowMessageRequest(t.Context(), &ShowMessageRequestParams{Message: "Proceed?"})
	if err != nil {
		t.Fatalf("ShowMessageRequest: %v", err)
	}
	method, _, calls, _ := conn.snapshot()
	if method != MethodWindowShowMessageRequest {
		t.Errorf("method = %q, want %q", method, MethodWindowShowMessageRequest)
	}
	if calls != 1 {
		t.Errorf("Call count = %d, want 1", calls)
	}
	if diff := gocmp.Diff(&MessageActionItem{Title: "Apply"}, got); diff != "" {
		t.Errorf("result mismatch (-want +got):\n%s", diff)
	}
}

// TestClientShowDocumentResult covers ShowDocument's pointer-result decode path.
func TestClientShowDocumentResult(t *testing.T) {
	t.Parallel()

	conn := &fakeConn{}
	conn.setResult(&ShowDocumentResult{Success: true}, nil)
	c := ClientDispatcher(conn)

	got, err := c.ShowDocument(t.Context(), &ShowDocumentParams{URI: "file:///a.go"})
	if err != nil {
		t.Fatalf("ShowDocument: %v", err)
	}
	method, _, _, _ := conn.snapshot()
	if method != MethodWindowShowDocument {
		t.Errorf("method = %q, want %q", method, MethodWindowShowDocument)
	}
	if got == nil || !got.Success {
		t.Errorf("result = %+v, want Success=true", got)
	}
}

// TestClientApplyEditResult covers ApplyEdit's pointer-result decode path.
func TestClientApplyEditResult(t *testing.T) {
	t.Parallel()

	conn := &fakeConn{}
	conn.setResult(&ApplyWorkspaceEditResult{Applied: true}, nil)
	c := ClientDispatcher(conn)

	got, err := c.ApplyEdit(t.Context(), &ApplyWorkspaceEditParams{Edit: WorkspaceEdit{}})
	if err != nil {
		t.Fatalf("ApplyEdit: %v", err)
	}
	method, _, _, _ := conn.snapshot()
	if method != MethodWorkspaceApplyEdit {
		t.Errorf("method = %q, want %q", method, MethodWorkspaceApplyEdit)
	}
	if got == nil || !got.Applied {
		t.Errorf("result = %+v, want Applied=true", got)
	}
}

// TestClientConfigurationResult covers Configuration's slice-result decode path.
func TestClientConfigurationResult(t *testing.T) {
	t.Parallel()

	conn := &fakeConn{}
	conn.setResult([]LSPAny{LSPAny(`"v1"`), LSPAny(`"v2"`)}, nil)
	c := ClientDispatcher(conn)

	got, err := c.Configuration(t.Context(), &ConfigurationParams{Items: []ConfigurationItem{{}}})
	if err != nil {
		t.Fatalf("Configuration: %v", err)
	}
	method, _, _, _ := conn.snapshot()
	if method != MethodWorkspaceConfiguration {
		t.Errorf("method = %q, want %q", method, MethodWorkspaceConfiguration)
	}
	if len(got) != 2 {
		t.Errorf("len(result) = %d, want 2", len(got))
	}
}

// TestClientWorkspaceFoldersResult covers WorkspaceFolders, the lone client
// request with no params and a slice result.
func TestClientWorkspaceFoldersResult(t *testing.T) {
	t.Parallel()

	conn := &fakeConn{}
	conn.setResult([]WorkspaceFolder{{URI: "file:///root", Name: "root"}}, nil)
	c := ClientDispatcher(conn)

	got, err := c.WorkspaceFolders(t.Context())
	if err != nil {
		t.Fatalf("WorkspaceFolders: %v", err)
	}
	method, params, _, _ := conn.snapshot()
	if method != MethodWorkspaceWorkspaceFolders {
		t.Errorf("method = %q, want %q", method, MethodWorkspaceWorkspaceFolders)
	}
	if params != nil {
		t.Errorf("params = %v, want nil", params)
	}
	if len(got) != 1 || got[0].Name != "root" {
		t.Errorf("result = %+v, want one folder named root", got)
	}
}

// TestClientRequestMethodsPropagateError asserts that a Call error is returned
// verbatim and the typed result is the zero value (nil) for the result-bearing
// client request methods.
func TestClientRequestMethodsPropagateError(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("boom")

	t.Run("showMessageRequest", func(t *testing.T) {
		t.Parallel()

		conn := &fakeConn{}
		conn.setResult(nil, wantErr)
		c := ClientDispatcher(conn)
		got, err := c.ShowMessageRequest(t.Context(), &ShowMessageRequestParams{})
		if !errors.Is(err, wantErr) {
			t.Fatalf("err = %v, want %v", err, wantErr)
		}
		if got != nil {
			t.Errorf("result = %+v, want nil on error", got)
		}
	})

	t.Run("configuration", func(t *testing.T) {
		t.Parallel()

		conn := &fakeConn{}
		conn.setResult(nil, wantErr)
		c := ClientDispatcher(conn)
		got, err := c.Configuration(t.Context(), &ConfigurationParams{})
		if !errors.Is(err, wantErr) {
			t.Fatalf("err = %v, want %v", err, wantErr)
		}
		if got != nil {
			t.Errorf("result = %+v, want nil on error", got)
		}
	})

	t.Run("registerCapability", func(t *testing.T) {
		t.Parallel()

		conn := &fakeConn{}
		conn.setResult(nil, wantErr)
		c := ClientDispatcher(conn)
		if err := c.RegisterCapability(t.Context(), &RegistrationParams{}); !errors.Is(err, wantErr) {
			t.Fatalf("err = %v, want %v", err, wantErr)
		}
	})
}

// dispatchRecordingClient records which Client method the production dispatch
// path invoked. Each method signals a buffered channel so a routing test can
// confirm the request reached the intended handler. It embeds UnimplementedClient
// so only the methods under test are overridden, and returns minimal valid
// results for the request methods.
type dispatchRecordingClient struct {
	UnimplementedClient

	mu     sync.Mutex
	called string
	done   chan string
}

func newDispatchRecordingClient() *dispatchRecordingClient {
	return &dispatchRecordingClient{done: make(chan string, 1)}
}

func (c *dispatchRecordingClient) record(method string) {
	c.mu.Lock()
	c.called = method
	c.mu.Unlock()
	select {
	case c.done <- method:
	default:
	}
}

func (c *dispatchRecordingClient) Progress(context.Context, *ProgressParams) error {
	c.record("Progress")
	return nil
}

func (c *dispatchRecordingClient) LogTrace(context.Context, *LogTraceParams) error {
	c.record("LogTrace")
	return nil
}

func (c *dispatchRecordingClient) RegisterCapability(context.Context, *RegistrationParams) error {
	c.record("RegisterCapability")
	return nil
}

func (c *dispatchRecordingClient) UnregisterCapability(context.Context, *UnregistrationParams) error {
	c.record("UnregisterCapability")
	return nil
}

func (c *dispatchRecordingClient) ShowMessage(context.Context, *ShowMessageParams) error {
	c.record("ShowMessage")
	return nil
}

func (c *dispatchRecordingClient) LogMessage(context.Context, *LogMessageParams) error {
	c.record("LogMessage")
	return nil
}

func (c *dispatchRecordingClient) ShowMessageRequest(context.Context, *ShowMessageRequestParams) (*MessageActionItem, error) {
	c.record("ShowMessageRequest")
	return &MessageActionItem{Title: "ok"}, nil
}

func (c *dispatchRecordingClient) ShowDocument(context.Context, *ShowDocumentParams) (*ShowDocumentResult, error) {
	c.record("ShowDocument")
	return &ShowDocumentResult{Success: true}, nil
}

func (c *dispatchRecordingClient) WorkDoneProgressCreate(context.Context, *WorkDoneProgressCreateParams) error {
	c.record("WorkDoneProgressCreate")
	return nil
}

func (c *dispatchRecordingClient) Telemetry(context.Context, LSPAny) error {
	c.record("Telemetry")
	return nil
}

func (c *dispatchRecordingClient) PublishDiagnostics(context.Context, *PublishDiagnosticsParams) error {
	c.record("PublishDiagnostics")
	return nil
}

func (c *dispatchRecordingClient) Configuration(context.Context, *ConfigurationParams) ([]LSPAny, error) {
	c.record("Configuration")
	return []LSPAny{}, nil
}

func (c *dispatchRecordingClient) WorkspaceFolders(context.Context) ([]WorkspaceFolder, error) {
	c.record("WorkspaceFolders")
	return []WorkspaceFolder{}, nil
}

func (c *dispatchRecordingClient) ApplyEdit(context.Context, *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error) {
	c.record("ApplyEdit")
	return &ApplyWorkspaceEditResult{Applied: true}, nil
}

func (c *dispatchRecordingClient) CodeLensRefresh(context.Context) error {
	c.record("CodeLensRefresh")
	return nil
}

func (c *dispatchRecordingClient) FoldingRangeRefresh(context.Context) error {
	c.record("FoldingRangeRefresh")
	return nil
}

func (c *dispatchRecordingClient) SemanticTokensRefresh(context.Context) error {
	c.record("SemanticTokensRefresh")
	return nil
}

func (c *dispatchRecordingClient) InlineValueRefresh(context.Context) error {
	c.record("InlineValueRefresh")
	return nil
}

func (c *dispatchRecordingClient) InlayHintRefresh(context.Context) error {
	c.record("InlayHintRefresh")
	return nil
}

func (c *dispatchRecordingClient) DiagnosticRefresh(context.Context) error {
	c.record("DiagnosticRefresh")
	return nil
}

func (c *dispatchRecordingClient) TextDocumentContentRefresh(context.Context, *TextDocumentContentRefreshParams) error {
	c.record("TextDocumentContentRefresh")
	return nil
}

// TestClientDispatchRouting fires one wire message per standard client method
// through the production ClientHandler over an in-memory pipe and asserts
// clientDispatch routed it to the matching Client method. This exercises every
// case arm of clientDispatch the way the transport actually builds requests.
func TestClientDispatchRouting(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		method     string
		params     any
		isCall     bool
		wantMethod string
	}{
		"progress":                   {MethodProgress, &ProgressParams{Token: String("t")}, false, "Progress"},
		"logTrace":                   {MethodLogTrace, &LogTraceParams{Message: "m"}, false, "LogTrace"},
		"registerCapability":         {MethodClientRegisterCapability, &RegistrationParams{}, true, "RegisterCapability"},
		"unregisterCapability":       {MethodClientUnregisterCapability, &UnregistrationParams{}, true, "UnregisterCapability"},
		"showMessage":                {MethodWindowShowMessage, &ShowMessageParams{Type: MessageTypeInfo, Message: "m"}, false, "ShowMessage"},
		"showMessageRequest":         {MethodWindowShowMessageRequest, &ShowMessageRequestParams{Message: "m"}, true, "ShowMessageRequest"},
		"logMessage":                 {MethodWindowLogMessage, &LogMessageParams{Type: MessageTypeLog, Message: "m"}, false, "LogMessage"},
		"showDocument":               {MethodWindowShowDocument, &ShowDocumentParams{URI: "file:///a"}, true, "ShowDocument"},
		"workDoneProgressCreate":     {MethodWindowWorkDoneProgressCreate, &WorkDoneProgressCreateParams{Token: String("t")}, true, "WorkDoneProgressCreate"},
		"telemetry":                  {MethodTelemetryEvent, LSPAny(`{"k":"v"}`), false, "Telemetry"},
		"publishDiagnostics":         {MethodTextDocumentPublishDiagnostics, &PublishDiagnosticsParams{URI: "file:///a"}, false, "PublishDiagnostics"},
		"configuration":              {MethodWorkspaceConfiguration, &ConfigurationParams{}, true, "Configuration"},
		"workspaceFolders":           {MethodWorkspaceWorkspaceFolders, nil, true, "WorkspaceFolders"},
		"applyEdit":                  {MethodWorkspaceApplyEdit, &ApplyWorkspaceEditParams{Edit: WorkspaceEdit{}}, true, "ApplyEdit"},
		"codeLensRefresh":            {MethodWorkspaceCodeLensRefresh, nil, true, "CodeLensRefresh"},
		"foldingRangeRefresh":        {MethodWorkspaceFoldingRangeRefresh, nil, true, "FoldingRangeRefresh"},
		"semanticTokensRefresh":      {MethodWorkspaceSemanticTokensRefresh, nil, true, "SemanticTokensRefresh"},
		"inlineValueRefresh":         {MethodWorkspaceInlineValueRefresh, nil, true, "InlineValueRefresh"},
		"inlayHintRefresh":           {MethodWorkspaceInlayHintRefresh, nil, true, "InlayHintRefresh"},
		"diagnosticRefresh":          {MethodWorkspaceDiagnosticRefresh, nil, true, "DiagnosticRefresh"},
		"textDocumentContentRefresh": {MethodWorkspaceTextDocumentContentRefresh, &TextDocumentContentRefreshParams{URI: "file:///a"}, true, "TextDocumentContentRefresh"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
			defer cancel()

			client := newDispatchRecordingClient()
			caller, served := newClientHandlerConnPair(ctx, t, client, jsonrpc2.MethodNotFoundHandler)
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
			case got := <-client.done:
				if got != tt.wantMethod {
					t.Errorf("dispatched to %q, want %q", got, tt.wantMethod)
				}
			case <-ctx.Done():
				t.Fatalf("client method %q was not dispatched: %v", tt.wantMethod, ctx.Err())
			}
		})
	}
}

// TestClientResultMethodsPropagateError drives every result-bearing client
// request method against a conn that returns an error, asserting the error
// propagates verbatim. This covers each method's error-return branch.
func TestClientResultMethodsPropagateError(t *testing.T) {
	t.Parallel()

	invokers := map[string]func(ctx context.Context, c Client) error{
		"showMessageRequest": func(ctx context.Context, c Client) error {
			_, err := c.ShowMessageRequest(ctx, &ShowMessageRequestParams{})
			return err
		},
		"showDocument": func(ctx context.Context, c Client) error {
			_, err := c.ShowDocument(ctx, &ShowDocumentParams{})
			return err
		},
		"configuration": func(ctx context.Context, c Client) error {
			_, err := c.Configuration(ctx, &ConfigurationParams{})
			return err
		},
		"workspaceFolders": func(ctx context.Context, c Client) error {
			_, err := c.WorkspaceFolders(ctx)
			return err
		},
		"applyEdit": func(ctx context.Context, c Client) error {
			_, err := c.ApplyEdit(ctx, &ApplyWorkspaceEditParams{})
			return err
		},
	}
	wantErr := errors.New("boom")
	for name, invoke := range invokers {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			conn := &fakeConn{}
			conn.setResult(nil, wantErr)
			c := ClientDispatcher(conn)
			if err := invoke(t.Context(), c); !errors.Is(err, wantErr) {
				t.Fatalf("err = %v, want %v", err, wantErr)
			}
		})
	}
}

// TestClientDispatchMalformedParams asserts clientDispatch reports a parse error
// (wrapping jsonrpc2.ErrParse) for a request whose params fail to decode, without
// invoking the client method. It is driven through ClientHandler so the request
// is built by the production scanner.
func TestClientDispatchMalformedParams(t *testing.T) {
	t.Parallel()

	// Every standard client request method whose params are a JSON object: a JSON
	// array is accepted by the framer but rejected by the per-case struct decode,
	// exercising that case's replyParseError branch.
	methods := map[string]string{
		"registerCapability":         MethodClientRegisterCapability,
		"unregisterCapability":       MethodClientUnregisterCapability,
		"showMessageRequest":         MethodWindowShowMessageRequest,
		"showDocument":               MethodWindowShowDocument,
		"workDoneProgressCreate":     MethodWindowWorkDoneProgressCreate,
		"configuration":              MethodWorkspaceConfiguration,
		"applyEdit":                  MethodWorkspaceApplyEdit,
		"textDocumentContentRefresh": MethodWorkspaceTextDocumentContentRefresh,
	}
	for name, method := range methods {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
			defer cancel()

			client := newDispatchRecordingClient()
			caller, served := newClientHandlerConnPair(ctx, t, client, jsonrpc2.MethodNotFoundHandler)
			defer closeJSONRPCConns(t, caller, served)

			_, err := caller.Call(ctx, method, jsonrpc2.RawMessage(`[1,2,3]`), nil)
			if !errors.Is(err, jsonrpc2.ErrParse) {
				t.Fatalf("Call(%s) malformed params error = %v, want wrapping %v", method, err, jsonrpc2.ErrParse)
			}
			select {
			case got := <-client.done:
				t.Fatalf("client method %q was invoked despite malformed params", got)
			case <-time.After(100 * time.Millisecond):
			}
		})
	}
}
