// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"

	"go.lsp.dev/jsonrpc2"
)

const (
	MethodClientCancelRequest                 ClientMethod = "$/cancelRequest"                       // bidirect client notification
	MethodClientProgress                      ClientMethod = "$/progress"                            // bidirect client notification
	MethodLogTrace                            ClientMethod = "$/logTrace"                            // client notification
	MethodTelemetryEvent                      ClientMethod = "telemetry/event"                       // client notification
	MethodTextDocumentPublishDiagnostics      ClientMethod = "textDocument/publishDiagnostics"       // client notification
	MethodWindowLogMessage                    ClientMethod = "window/logMessage"                     // client notification
	MethodWindowShowMessage                   ClientMethod = "window/showMessage"                    // client notification
	MethodClientRegisterCapability            ClientMethod = "client/registerCapability"             // client request
	MethodClientUnregisterCapability          ClientMethod = "client/unregisterCapability"           // client request
	MethodWindowShowDocument                  ClientMethod = "window/showDocument"                   // client request
	MethodWindowShowMessageRequest            ClientMethod = "window/showMessageRequest"             // client request
	MethodWindowWorkDoneProgressCreate        ClientMethod = "window/workDoneProgress/create"        // client request
	MethodWorkspaceApplyEdit                  ClientMethod = "workspace/applyEdit"                   // client request
	MethodWorkspaceCodeLensRefresh            ClientMethod = "workspace/codeLens/refresh"            // client request
	MethodWorkspaceConfiguration              ClientMethod = "workspace/configuration"               // client request
	MethodWorkspaceDiagnosticRefresh          ClientMethod = "workspace/diagnostic/refresh"          // client request
	MethodWorkspaceFoldingRangeRefresh        ClientMethod = "workspace/foldingRange/refresh"        // client request
	MethodWorkspaceInlayHintRefresh           ClientMethod = "workspace/inlayHint/refresh"           // client request
	MethodWorkspaceInlineValueRefresh         ClientMethod = "workspace/inlineValue/refresh"         // client request
	MethodWorkspaceSemanticTokensRefresh      ClientMethod = "workspace/semanticTokens/refresh"      // client request
	MethodWorkspaceTextDocumentContentRefresh ClientMethod = "workspace/textDocumentContent/refresh" // client request
	MethodWorkspaceWorkspaceFolders           ClientMethod = "workspace/workspaceFolders"            // client request
)

type Client interface {
	Cancel(ctx context.Context, params *CancelParams) error

	Progress(ctx context.Context, params *ProgressParams) error

	LogTrace(ctx context.Context, params *LogTraceParams) error

	// TelemetryEvent the telemetry event notification is sent from the server to the client to ask the client to log telemetry data.
	TelemetryEvent(ctx context.Context, params any) error

	// PublishDiagnostics diagnostics notification are sent from the server to the client to signal results of validation runs.
	PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) error

	// LogMessage the log message notification is sent from the server to the client to ask the client to log a particular message.
	LogMessage(ctx context.Context, params *LogMessageParams) error

	// ShowMessage the show message notification is sent from a server to a client to ask the client to display a particular message in the user interface.
	ShowMessage(ctx context.Context, params *ShowMessageParams) error
	// Registration the `client/registerCapability` request is sent from the server to the client to register a new capability handler on the client side.
	Registration(ctx context.Context, params *RegistrationParams) error

	// Unregistration the `client/unregisterCapability` request is sent from the server to the client to unregister a previously registered capability handler on the client side.
	Unregistration(ctx context.Context, params *UnregistrationParams) error

	// ShowDocument a request to show a document. This request might open an external program depending on the value of the URI to open. For example a request to open `https://code.visualstudio.com/` will very likely open the URI in a WEB browser.
	//
	// @since 3.16.0
	ShowDocument(ctx context.Context, params *ShowDocumentParams) (*ShowDocumentResult, error)

	// ShowMessageRequest the show message request is sent from the server to the client to show a message and a set of options actions to the user.
	ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (*MessageActionItem, error)

	// WorkDoneProgressCreate the `window/workDoneProgress/create` request is sent from the server to the client to initiate progress reporting from the server.
	WorkDoneProgressCreate(ctx context.Context, params *WorkDoneProgressCreateParams) error

	// ApplyWorkspaceEdit a request sent from the server to the client to modified certain resources.
	ApplyWorkspaceEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error)

	// CodeLensRefresh a request to refresh all code actions
	//
	// @since 3.16.0
	CodeLensRefresh(ctx context.Context) error

	// Configuration the 'workspace/configuration' request is sent from the server to the client to fetch a certain configuration setting. This pull model replaces the old push model were the client signaled configuration
	// change via an event. If the server still needs to react to configuration changes (since the server caches the result of `workspace/configuration` requests) the server should register for an empty configuration change event and empty the cache if such an event is received.
	Configuration(ctx context.Context, params *ConfigurationParams) ([]any, error)

	// DiagnosticRefresh the diagnostic refresh request definition.
	//
	// @since 3.17.0
	DiagnosticRefresh(ctx context.Context) error

	// FoldingRangeRefresh.
	//
	// @since 3.18.0 proposed
	FoldingRangeRefresh(ctx context.Context) error

	// InlayHintRefresh.
	//
	// @since 3.17.0
	InlayHintRefresh(ctx context.Context) error

	// InlineValueRefresh.
	//
	// @since 3.17.0
	InlineValueRefresh(ctx context.Context) error

	// SemanticTokensRefresh.
	//
	// @since 3.16.0
	SemanticTokensRefresh(ctx context.Context) error

	// TextDocumentContentRefresh the `workspace/textDocumentContent` request is sent from the server to the client to refresh the content of a specific text document. 3.18.0 @proposed.
	//
	// @since 3.18.0 proposed
	TextDocumentContentRefresh(ctx context.Context, params *TextDocumentContentRefreshParams) error

	// WorkspaceFolders the `workspace/workspaceFolders` is sent from the server to the client to fetch the open workspace folders.
	WorkspaceFolders(ctx context.Context) ([]*WorkspaceFolder, error)
}

// UnimplementedClient should be embedded to have forward compatible implementations.
type UnimplementedClient struct{}

var _ Client = UnimplementedClient{}

func (UnimplementedClient) Cancel(ctx context.Context, params *CancelParams) error {
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

func (UnimplementedClient) PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) LogMessage(ctx context.Context, params *LogMessageParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) ShowMessage(ctx context.Context, params *ShowMessageParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) Registration(ctx context.Context, params *RegistrationParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) Unregistration(ctx context.Context, params *UnregistrationParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) ShowDocument(ctx context.Context, params *ShowDocumentParams) (*ShowDocumentResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedClient) ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (*MessageActionItem, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkDoneProgressCreate(ctx context.Context, params *WorkDoneProgressCreateParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) ApplyWorkspaceEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedClient) CodeLensRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) Configuration(ctx context.Context, params *ConfigurationParams) ([]any, error) {
	return nil, jsonrpc2.ErrInternal
}

func (UnimplementedClient) DiagnosticRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) FoldingRangeRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) InlayHintRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) InlineValueRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) SemanticTokensRefresh(ctx context.Context) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) TextDocumentContentRefresh(ctx context.Context, params *TextDocumentContentRefreshParams) error {
	return jsonrpc2.ErrInternal
}

func (UnimplementedClient) WorkspaceFolders(ctx context.Context) ([]*WorkspaceFolder, error) {
	return nil, jsonrpc2.ErrInternal
}
