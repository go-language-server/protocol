// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"
)

// clientHandler represents a client handler.
type clientHandler struct {
	client ClientInterface
}

// compile time check whether the clientHandler implements jsonrpc2.Handler interface.
var _ jsonrpc2.Handler = &clientHandler{}

// Cancel implements Handler interface.
func (clientHandler) Cancel(ctx context.Context, conn *jsonrpc2.Conn, id jsonrpc2.ID, canceled bool) bool {
	return false
}

// Request implements Handler interface.
func (clientHandler) Request(ctx context.Context, conn *jsonrpc2.Conn, direction jsonrpc2.Direction, r *jsonrpc2.WireRequest) context.Context {
	return ctx
}

// Response implements Handler interface.
func (clientHandler) Response(ctx context.Context, conn *jsonrpc2.Conn, direction jsonrpc2.Direction, r *jsonrpc2.WireResponse) context.Context {
	return ctx
}

// Done implements Handler interface.
func (clientHandler) Done(ctx context.Context, err error) {}

// Read implements Handler interface.
func (clientHandler) Read(ctx context.Context, bytes int64) context.Context { return ctx }

// Write implements Handler interface.
func (clientHandler) Write(ctx context.Context, bytes int64) context.Context { return ctx }

// Error implements Handler interface.
func (clientHandler) Error(ctx context.Context, err error) {}

// ClientInterface represents a Language Server Protocol client.
type ClientInterface interface {
	Run(ctx context.Context) (err error)
	LogMessage(ctx context.Context, params *LogMessageParams) (err error)
	PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) (err error)
	ShowMessage(ctx context.Context, params *ShowMessageParams) (err error)
	ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (result *MessageActionItem, err error)
	Telemetry(ctx context.Context, params interface{}) (err error)
	RegisterCapability(ctx context.Context, params *RegistrationParams) (err error)
	UnregisterCapability(ctx context.Context, params *UnregistrationParams) (err error)
	WorkspaceApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (result bool, err error)
	WorkspaceConfiguration(ctx context.Context, params *ConfigurationParams) (result []interface{}, err error)
	WorkspaceFolders(ctx context.Context) (result []WorkspaceFolder, err error)
}

const (
	// MethodWindowShowMessage method name of "window/showMessage".
	MethodWindowShowMessage = "window/showMessage"

	// MethodWindowShowMessageRequest method name of "window/showMessageRequest.
	MethodWindowShowMessageRequest = "window/showMessageRequest"

	// MethodWindowLogMessage method name of "window/logMessage.
	MethodWindowLogMessage = "window/logMessage"

	// MethodTelemetryEvent method name of "telemetry/event.
	MethodTelemetryEvent = "telemetry/event"

	// MethodClientRegisterCapability method name of "client/registerCapability.
	MethodClientRegisterCapability = "client/registerCapability"

	// MethodClientUnregisterCapability method name of "client/unregisterCapability.
	MethodClientUnregisterCapability = "client/unregisterCapability"

	// MethodTextDocumentPublishDiagnostics method name of "textDocument/publishDiagnostics.
	MethodTextDocumentPublishDiagnostics = "textDocument/publishDiagnostics"

	// MethodWorkspaceApplyEdit method name of "workspace/applyEdit.
	MethodWorkspaceApplyEdit = "workspace/applyEdit"

	// MethodWorkspaceConfiguration method name of "workspace/configuration.
	MethodWorkspaceConfiguration = "workspace/configuration"

	// MethodWorkspaceWorkspaceFolders method name of "workspace/workspaceFolders".
	MethodWorkspaceWorkspaceFolders = "workspace/workspaceFolders"
)

// client implements a Language Server Protocol client.
type client struct {
	*jsonrpc2.Conn
	logger *zap.Logger
}

// compiler time check whether the Client implements ClientInterface interface.
var _ ClientInterface = (*client)(nil)

// Run runs the Language Server Protocol client.
func (c *client) Run(ctx context.Context) (err error) {
	err = c.Conn.Run(ctx)
	return
}

// LogMessage sends the notification from the server to the client to ask the client to log a particular message.
func (c *client) LogMessage(ctx context.Context, params *LogMessageParams) (err error) {
	err = c.Conn.Notify(ctx, MethodWindowLogMessage, params)
	return
}

// PublishDiagnostics sends the notification from the server to the client to signal results of validation runs.
//
// Diagnostics are “owned” by the server so it is the server’s responsibility to clear them if necessary. The following rule is used for VS Code servers that generate diagnostics:
//
// - if a language is single file only (for example HTML) then diagnostics are cleared by the server when the file is closed.
// - if a language has a project system (for example C#) diagnostics are not cleared when a file closes. When a project is opened all diagnostics for all files are recomputed (or read from a cache).
//
// When a file changes it is the server’s responsibility to re-compute diagnostics and push them to the client.
// If the computed set is empty it has to push the empty array to clear former diagnostics.
// Newly pushed diagnostics always replace previously pushed diagnostics. There is no merging that happens on the client side.
func (c *client) PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) (err error) {
	err = c.Conn.Notify(ctx, MethodTextDocumentPublishDiagnostics, params)
	return
}

// ShowMessage sends the notification from a server to a client to ask the
// client to display a particular message in the user interface.
func (c *client) ShowMessage(ctx context.Context, params *ShowMessageParams) (err error) {
	err = c.Conn.Notify(ctx, MethodWindowShowMessage, params)
	return
}

// ShowMessageRequest sends the request from a server to a client to ask the client to display a particular message in the user interface.
//
// In addition to the show message notification the request allows to pass actions and to wait for an answer from the client.
func (c *client) ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (result *MessageActionItem, err error) {
	result = new(MessageActionItem)
	err = c.Conn.Call(ctx, MethodWindowShowMessageRequest, params, result)

	return result, err
}

// Telemetry sends the notification from the server to the client to ask the client to log a telemetry event.
func (c *client) Telemetry(ctx context.Context, params interface{}) (err error) {
	err = c.Conn.Notify(ctx, MethodTelemetryEvent, params)
	return
}

// RegisterCapability sends the request from the server to the client to register for a new capability on the client side.
//
// Not all clients need to support dynamic capability registration.
//
// A client opts in via the dynamicRegistration property on the specific client capabilities.
// A client can even provide dynamic registration for capability A but not for capability B (see TextDocumentClientCapabilities as an example).
func (c *client) RegisterCapability(ctx context.Context, params *RegistrationParams) (err error) {
	err = c.Conn.Call(ctx, MethodClientRegisterCapability, params, nil)
	return
}

// UnregisterCapability sends the request from the server to the client to unregister a previously registered capability.
func (c *client) UnregisterCapability(ctx context.Context, params *UnregistrationParams) (err error) {
	err = c.Conn.Call(ctx, MethodClientUnregisterCapability, params, nil)
	return
}

// WorkspaceApplyEdit sends the request from the server to the client to modify resource on the client side.
func (c *client) WorkspaceApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (result bool, err error) {
	err = c.Conn.Call(ctx, MethodWorkspaceApplyEdit, params, &result)

	return result, err
}

// WorkspaceConfiguration sends the request from the server to the client to fetch configuration settings from the client.
//
// The request can fetch several configuration settings in one roundtrip.
// The order of the returned configuration settings correspond to the order of the
// passed ConfigurationItems (e.g. the first item in the response is the result for the first configuration item in the params).
func (c *client) WorkspaceConfiguration(ctx context.Context, params *ConfigurationParams) ([]interface{}, error) {
	var result []interface{}
	err := c.Conn.Call(ctx, MethodWorkspaceConfiguration, params, &result)

	return result, err
}

// WorkspaceFolders sends the request from the server to the client to fetch the current open list of workspace folders.
//
// Returns null in the response if only a single file is open in the tool. Returns an empty array if a workspace is open but no folders are configured.
//
// Since version 3.6.0.
func (c *client) WorkspaceFolders(ctx context.Context) (result []WorkspaceFolder, err error) {
	err = c.Conn.Call(ctx, MethodWorkspaceWorkspaceFolders, nil, &result)

	return result, err
}
