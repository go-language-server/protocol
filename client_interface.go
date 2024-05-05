// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"

	"go.lsp.dev/jsonrpc2"
)

const (
	MethodClientCancelRequest            ClientMethod = "$/cancelRequest"                  // bidirect client notification
	MethodClientProgress                 ClientMethod = "$/progress"                       // bidirect client notification
	MethodLogTrace                       ClientMethod = "$/logTrace"                       // client notification
	MethodTelemetryEvent                 ClientMethod = "telemetry/event"                  // client notification
	MethodTextDocumentPublishDiagnostics ClientMethod = "textDocument/publishDiagnostics"  // client notification
	MethodWindowLogMessage               ClientMethod = "window/logMessage"                // client notification
	MethodWindowShowMessage              ClientMethod = "window/showMessage"               // client notification
	MethodClientRegisterCapability       ClientMethod = "client/registerCapability"        // client request
	MethodClientUnregisterCapability     ClientMethod = "client/unregisterCapability"      // client request
	MethodWindowShowDocument             ClientMethod = "window/showDocument"              // client request
	MethodWindowShowMessageRequest       ClientMethod = "window/showMessageRequest"        // client request
	MethodWindowWorkDoneProgressCreate   ClientMethod = "window/workDoneProgress/create"   // client request
	MethodWorkspaceApplyEdit             ClientMethod = "workspace/applyEdit"              // client request
	MethodWorkspaceCodeLensRefresh       ClientMethod = "workspace/codeLens/refresh"       // client request
	MethodWorkspaceConfiguration         ClientMethod = "workspace/configuration"          // client request
	MethodWorkspaceDiagnosticRefresh     ClientMethod = "workspace/diagnostic/refresh"     // client request
	MethodWorkspaceFoldingRangeRefresh   ClientMethod = "workspace/foldingRange/refresh"   // client request
	MethodWorkspaceInlayHintRefresh      ClientMethod = "workspace/inlayHint/refresh"      // client request
	MethodWorkspaceInlineValueRefresh    ClientMethod = "workspace/inlineValue/refresh"    // client request
	MethodWorkspaceSemanticTokensRefresh ClientMethod = "workspace/semanticTokens/refresh" // client request
	MethodWorkspaceWorkspaceFolders      ClientMethod = "workspace/workspaceFolders"       // client request
)

type Client interface {
	CancelRequest(ctx context.Context, params *CancelParams) error

	Progress(ctx context.Context, params *ProgressParams) error

	LogTrace(ctx context.Context, params *LogTraceParams) error

	// TelemetryEvent the telemetry event notification is sent from the server to the client to ask the client to log telemetry data.
	TelemetryEvent(ctx context.Context, params any) error

	// TextDocumentPublishDiagnostics diagnostics notification are sent from the server to the client to signal results of validation runs.
	TextDocumentPublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) error

	// WindowLogMessage the log message notification is sent from the server to the client to ask the client to log a particular message.
	WindowLogMessage(ctx context.Context, params *LogMessageParams) error

	// WindowShowMessage the show message notification is sent from a server to a client to ask the client to display a particular message in the user interface.
	WindowShowMessage(ctx context.Context, params *ShowMessageParams) error
	// ClientRegisterCapability the `client/registerCapability` request is sent from the server to the client to register a new capability handler on the client side.
	ClientRegisterCapability(ctx context.Context, params *RegistrationParams) error

	// ClientUnregisterCapability the `client/unregisterCapability` request is sent from the server to the client to unregister a previously registered capability handler on the client side.
	ClientUnregisterCapability(ctx context.Context, params *UnregistrationParams) error

	// WindowShowDocument a request to show a document. This request might open an external program depending on the value of the URI to open. For example a request to open `https://code.visualstudio.com/` will very likely open the URI in a WEB browser.
	//
	// @since 3.16.0
	WindowShowDocument(ctx context.Context, params *ShowDocumentParams) (*ShowDocumentResult, error)

	// WindowShowMessageRequest the show message request is sent from the server to the client to show a message and a set of options actions to the user.
	WindowShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (*MessageActionItem, error)

	// WindowWorkDoneProgressCreate the `window/workDoneProgress/create` request is sent from the server to the client to initiate progress reporting from the server.
	WindowWorkDoneProgressCreate(ctx context.Context, params *WorkDoneProgressCreateParams) error

	// WorkspaceApplyEdit a request sent from the server to the client to modified certain resources.
	WorkspaceApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error)

	// WorkspaceCodeLensRefresh a request to refresh all code actions
	//
	// @since 3.16.0
	WorkspaceCodeLensRefresh(ctx context.Context) error

	// WorkspaceConfiguration the 'workspace/configuration' request is sent from the server to the client to fetch a certain configuration setting. This pull model replaces the old push model were the client signaled configuration
	// change via an event. If the server still needs to react to configuration changes (since the server caches the result of `workspace/configuration` requests) the server should register for an empty configuration change event and empty the cache if such an event is received.
	WorkspaceConfiguration(ctx context.Context, params *ConfigurationParams) ([]any, error)

	// WorkspaceDiagnosticRefresh the diagnostic refresh request definition.
	//
	// @since 3.17.0
	WorkspaceDiagnosticRefresh(ctx context.Context) error

	// WorkspaceFoldingRangeRefresh.
	//
	// @since 3.18.0 proposed
	WorkspaceFoldingRangeRefresh(ctx context.Context) error

	// WorkspaceInlayHintRefresh.
	//
	// @since 3.17.0
	WorkspaceInlayHintRefresh(ctx context.Context) error

	// WorkspaceInlineValueRefresh.
	//
	// @since 3.17.0
	WorkspaceInlineValueRefresh(ctx context.Context) error

	// WorkspaceSemanticTokensRefresh.
	//
	// @since 3.16.0
	WorkspaceSemanticTokensRefresh(ctx context.Context) error

	// WorkspaceWorkspaceFolders the `workspace/workspaceFolders` is sent from the server to the client to fetch the open workspace folders.
	WorkspaceWorkspaceFolders(ctx context.Context) ([]*WorkspaceFolder, error)
}

// UnimplementedClient should be embedded to have forward compatible implementations.
type UnimplementedClient struct{}

func (UnimplementedClient) CancelRequest(ctx context.Context, params *CancelParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) Progress(ctx context.Context, params *ProgressParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) LogTrace(ctx context.Context, params *LogTraceParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) TelemetryEvent(ctx context.Context, params any) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) TextDocumentPublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WindowLogMessage(ctx context.Context, params *LogMessageParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WindowShowMessage(ctx context.Context, params *ShowMessageParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) ClientRegisterCapability(ctx context.Context, params *RegistrationParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) ClientUnregisterCapability(ctx context.Context, params *UnregistrationParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WindowShowDocument(ctx context.Context, params *ShowDocumentParams) (*ShowDocumentResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedClient) WindowShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (*MessageActionItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedClient) WindowWorkDoneProgressCreate(ctx context.Context, params *WorkDoneProgressCreateParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceCodeLensRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceConfiguration(ctx context.Context, params *ConfigurationParams) ([]any, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceDiagnosticRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceFoldingRangeRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceInlayHintRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceInlineValueRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceSemanticTokensRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceWorkspaceFolders(ctx context.Context) ([]*WorkspaceFolder, error) {
	return nil, jsonrpc2.ErrInternal
}
