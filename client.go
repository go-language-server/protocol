// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/pkg/xcontext"
)

// NewClient returns the context in which Client is embedded, jsonrpc2.Conn, and the Server.
func NewClientV1(ctx context.Context, clientFunc ClientFunc, listen jsonrpc2.Listener, logger *zap.Logger) (context.Context, jsonrpc2.Conn, Server) {
	client := &client{
		newClient: clientFunc,
	}

	ctx = WithClient(ctx, client)

	listen, err := jsonrpc2.NetPipe(ctx)
	conn.Go(ctx,
		Handlers(
			ClientHandler(client, jsonrpc2.MethodNotFoundHandler),
		),
	)
	server := ServerDispatcher(conn, logger.Named("server"))

	return ctx, conn, server
}

// clientConn implements a Language Server Protocol client.
type clientConn struct {
	conn *jsonrpc2.Connection
}

func (c *clientConn) Close() error {
	return c.conn.Close()
}

func (c *clientConn) Notify(ctx context.Context, method string, params interface{}) error {
	return c.conn.Notify(ctx, method, params)
}

func (c *clientConn) Call(ctx context.Context, method string, params interface{}, result interface{}) (err error) {
	call := c.conn.Call(ctx, method, params)
	err = call.Await(ctx, result)

	if ctx.Err() != nil {
		detachCtx := xcontext.Detach(ctx)
		c.conn.Notify(detachCtx, MethodCancelRequest, &CancelParams{ID: call.ID().Raw()})
	}

	return err
}

// ClientDispatcher returns a Client that dispatches LSP requests across the
// given jsonrpc2 connection.
func ClientDispatcher(conn *jsonrpc2.Connection) Client {
	return &clientConn{
		conn: conn,
	}
}

// ClientOption represents a configures a client.
type ClientOption interface {
	apply(*client)
}

// clientOptionFunc wraps a func so it satisfies the ClientOption interface.
type clientOptionFunc func(*client)

func (f clientOptionFunc) apply(c *client) {
	f(c)
}

// WithInterceptor sets interceptor to Client.
func WithInterceptor(interceptor jsonrpc2.Interceptor) ClientOption {
	return clientOptionFunc(func(c *client) {
		c.interceptors = append(c.interceptors, interceptor)
	})
}

// WithClientLogger sets logger to Client.
func WithClientLogger(logger *zap.Logger) ClientOption {
	return clientOptionFunc(func(c *client) {
		c.logger = logger
	})
}

// ClientFunc is used to construct a Language Server Protocol client for a given server.
type ClientFunc func(ctx context.Context, server Server) Client

// Client binds a Language Server Protocol client to an incoming connection.
type client struct {
	newClient    ClientFunc
	interceptors []jsonrpc2.Interceptor

	logger *zap.Logger
}

// NewClient returns the new Client.
func NewClient(clientFunc ClientFunc, opts ...ClientOption) jsonrpc2.Binder {
	c := &client{
		newClient: clientFunc,
		logger:    zap.NewNop(),
	}

	for _, opt := range opts {
		opt.apply(c)
	}

	for _, interceptor := range c.interceptors {
		c = interceptor(c)
	}

	return c
}

// Bind implements jsonrpc2.Binder.Bind.
func (c *client) Bind(ctx context.Context, conn *jsonrpc2.Connection) (jsonrpc2.Conn, error) {
	server := ServerDispatcher(conn, WithServerLogger(c.logger))
	client := c.newClient(ctx, server)
	ctx = WithClient(ctx, client)

	conn := jsonrpc2.Conn{
		Framer:  jsonrpc2.HeaderFramer(),
		Handler: ClientHandler(client),
	}

	return conn, nil
}

func ClientHandler(client Client) jsonrpc2.Handler {
	return jsonrpc2.HandlerFunc(func(ctx context.Context, req *jsonrpc2.Request) (interface{}, error) {
		if ctx.Err() != nil {
			return nil, RequestCancelledError
		}

		result, err := clientDispatch(ctx, client, req)
		if err != nil {
			return nil, err
		}

		return result, resErr
	})
}

// Client represents a Language Server Protocol client.
type Client interface {
	Progress(ctx context.Context, params *ProgressParams) (err error)
	WorkDoneProgressCreate(ctx context.Context, params *WorkDoneProgressCreateParams) (err error)
	LogMessage(ctx context.Context, params *LogMessageParams) (err error)
	PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) (err error)
	ShowMessage(ctx context.Context, params *ShowMessageParams) (err error)
	ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (result *MessageActionItem, err error)
	Telemetry(ctx context.Context, params interface{}) (err error)
	RegisterCapability(ctx context.Context, params *RegistrationParams) (err error)
	UnregisterCapability(ctx context.Context, params *UnregistrationParams) (err error)
	ApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (result bool, err error)
	Configuration(ctx context.Context, params *ConfigurationParams) (result []interface{}, err error)
	WorkspaceFolders(ctx context.Context) (result []WorkspaceFolder, err error)
}

// list of client methods.
const (
	// MethodProgress method name of "$/progress".
	MethodProgress = "$/progress"

	// MethodWorkDoneProgressCreate method name of "window/workDoneProgress/create".
	MethodWorkDoneProgressCreate = "window/workDoneProgress/create"

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

// Progress is the base protocol offers also support to report progress in a generic fashion.
//
// This mechanism can be used to report any kind of progress including work done progress (usually used to report progress in the user interface using a progress bar) and
// partial result progress to support streaming of results.
//
// @since 3.16.0.
func (c *client) Progress(ctx context.Context, params *ProgressParams) (err error) {
	c.logger.Debug("call " + MethodProgress)
	defer c.logger.Debug("end "+MethodProgress, zap.Error(err))

	return c.Conn.Notify(ctx, MethodProgress, params)
}

// WorkDoneProgressCreate sends the request is sent from the server to the client to ask the client to create a work done progress.
//
// @since 3.16.0.
func (c *client) WorkDoneProgressCreate(ctx context.Context, params *WorkDoneProgressCreateParams) (err error) {
	c.logger.Debug("call " + MethodWorkDoneProgressCreate)
	defer c.logger.Debug("end "+MethodWorkDoneProgressCreate, zap.Error(err))

	return Call(ctx, c.Conn, MethodWorkDoneProgressCreate, params, nil)
}

// LogMessage sends the notification from the server to the client to ask the client to log a particular message.
func (c *client) LogMessage(ctx context.Context, params *LogMessageParams) (err error) {
	c.logger.Debug("call " + MethodWindowLogMessage)
	defer c.logger.Debug("end "+MethodWindowLogMessage, zap.Error(err))

	return c.Conn.Notify(ctx, MethodWindowLogMessage, params)
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
	c.logger.Debug("call " + MethodTextDocumentPublishDiagnostics)
	defer c.logger.Debug("end "+MethodTextDocumentPublishDiagnostics, zap.Error(err))

	return c.Conn.Notify(ctx, MethodTextDocumentPublishDiagnostics, params)
}

// ShowMessage sends the notification from a server to a client to ask the
// client to display a particular message in the user interface.
func (c *client) ShowMessage(ctx context.Context, params *ShowMessageParams) (err error) {
	return c.Conn.Notify(ctx, MethodWindowShowMessage, params)
}

// ShowMessageRequest sends the request from a server to a client to ask the client to display a particular message in the user interface.
//
// In addition to the show message notification the request allows to pass actions and to wait for an answer from the client.
func (c *client) ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (_ *MessageActionItem, err error) {
	c.logger.Debug("call " + MethodWindowShowMessageRequest)
	defer c.logger.Debug("end "+MethodWindowShowMessageRequest, zap.Error(err))

	var result *MessageActionItem
	if err := Call(ctx, c.Conn, MethodWindowShowMessageRequest, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Telemetry sends the notification from the server to the client to ask the client to log a telemetry event.
func (c *client) Telemetry(ctx context.Context, params interface{}) (err error) {
	c.logger.Debug("call " + MethodTelemetryEvent)
	defer c.logger.Debug("end "+MethodTelemetryEvent, zap.Error(err))

	return c.Conn.Notify(ctx, MethodTelemetryEvent, params)
}

// RegisterCapability sends the request from the server to the client to register for a new capability on the client side.
//
// Not all clients need to support dynamic capability registration.
//
// A client opts in via the dynamicRegistration property on the specific client capabilities.
// A client can even provide dynamic registration for capability A but not for capability B (see TextDocumentClientCapabilities as an example).
func (c *client) RegisterCapability(ctx context.Context, params *RegistrationParams) (err error) {
	c.logger.Debug("call " + MethodClientRegisterCapability)
	defer c.logger.Debug("end "+MethodClientRegisterCapability, zap.Error(err))

	return Call(ctx, c.Conn, MethodClientRegisterCapability, params, nil)
}

// UnregisterCapability sends the request from the server to the client to unregister a previously registered capability.
func (c *client) UnregisterCapability(ctx context.Context, params *UnregistrationParams) (err error) {
	c.logger.Debug("call " + MethodClientUnregisterCapability)
	defer c.logger.Debug("end "+MethodClientUnregisterCapability, zap.Error(err))

	return Call(ctx, c.Conn, MethodClientUnregisterCapability, params, nil)
}

// ApplyEdit sends the request from the server to the client to modify resource on the client side.
func (c *client) ApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (result bool, err error) {
	c.logger.Debug("call " + MethodWorkspaceApplyEdit)
	defer c.logger.Debug("end "+MethodWorkspaceApplyEdit, zap.Error(err))

	if err := Call(ctx, c.Conn, MethodWorkspaceApplyEdit, params, &result); err != nil {
		return false, err
	}

	return result, nil
}

// Configuration sends the request from the server to the client to fetch configuration settings from the client.
//
// The request can fetch several configuration settings in one roundtrip.
// The order of the returned configuration settings correspond to the order of the
// passed ConfigurationItems (e.g. the first item in the response is the result for the first configuration item in the params).
func (c *client) Configuration(ctx context.Context, params *ConfigurationParams) (_ []interface{}, err error) {
	c.logger.Debug("call " + MethodWorkspaceConfiguration)
	defer c.logger.Debug("end "+MethodWorkspaceConfiguration, zap.Error(err))

	var result []interface{}
	if err := Call(ctx, c.Conn, MethodWorkspaceConfiguration, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// WorkspaceFolders sends the request from the server to the client to fetch the current open list of workspace folders.
//
// Returns null in the response if only a single file is open in the tool. Returns an empty array if a workspace is open but no folders are configured.
//
// @since 3.6.0.
func (c *client) WorkspaceFolders(ctx context.Context) (result []WorkspaceFolder, err error) {
	c.logger.Debug("call " + MethodWorkspaceWorkspaceFolders)
	defer c.logger.Debug("end "+MethodWorkspaceWorkspaceFolders, zap.Error(err))

	if err := Call(ctx, c.Conn, MethodWorkspaceWorkspaceFolders, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}
