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

// ClientDispatcher returns a Client that dispatches LSP requests across the
// given jsonrpc2 connection.
func ClientDispatcher(conn jsonrpc2.Conn, logger *zap.Logger) Client {
	return &client{
		Conn:   conn,
		logger: logger,
	}
}

// ClientHandler handler of LSP client.
func ClientHandler(client Client, handler jsonrpc2.Handler) jsonrpc2.Handler {
	h := func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		if ctx.Err() != nil {
			xctx := context.WithoutCancel(ctx)

			return reply(xctx, nil, ErrRequestCancelled)
		}

		handled, err := clientDispatch(ctx, client, reply, req)
		if handled || err != nil {
			return err
		}

		return handler(ctx, reply, req)
	}

	return h
}

// clientDispatch implements jsonrpc2.Handler.
//
//nolint:funlen,cyclop
func clientDispatch(ctx context.Context, client Client, reply jsonrpc2.Replier, req jsonrpc2.Request) (handled bool, err error) {
	if ctx.Err() != nil {
		return true, reply(ctx, nil, ErrRequestCancelled)
	}

	dec := newDecoder(bytes.NewReader(req.Params()))
	logger := LoggerFromContext(ctx)

	switch req.Method() {
	case MethodClientProgress: // notification
		defer logger.Debug(MethodClientProgress, zap.Error(err))

		var params ProgressParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.Progress(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodLogTrace: // notification
		defer logger.Debug(MethodLogTrace, zap.Error(err))

		var params LogTraceParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.LogTrace(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodTelemetryEvent: // notification
		defer logger.Debug(MethodTelemetryEvent, zap.Error(err))

		var params any
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.TelemetryEvent(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodTextDocumentPublishDiagnostics: // notification
		defer logger.Debug(MethodTextDocumentPublishDiagnostics, zap.Error(err))

		var params PublishDiagnosticsParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.TextDocumentPublishDiagnostics(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWindowLogMessage: // notification
		defer logger.Debug(MethodWindowLogMessage, zap.Error(err))

		var params LogMessageParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.WindowLogMessage(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWindowShowMessage: // notification
		defer logger.Debug(MethodWindowShowMessage, zap.Error(err))

		var params ShowMessageParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.WindowShowMessage(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodClientRegisterCapability: // request
		defer logger.Debug(MethodClientRegisterCapability, zap.Error(err))

		var params RegistrationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.ClientRegisterCapability(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodClientUnregisterCapability: // request
		defer logger.Debug(MethodClientUnregisterCapability, zap.Error(err))

		var params UnregistrationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.ClientUnregisterCapability(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWindowShowDocument: // request
		defer logger.Debug(MethodWindowShowDocument, zap.Error(err))

		var params ShowDocumentParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := client.WindowShowDocument(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWindowShowMessageRequest: // request
		defer logger.Debug(MethodWindowShowMessageRequest, zap.Error(err))

		var params ShowMessageRequestParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := client.WindowShowMessageRequest(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWindowWorkDoneProgressCreate: // request
		defer logger.Debug(MethodWindowWorkDoneProgressCreate, zap.Error(err))

		var params WorkDoneProgressCreateParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.WindowWorkDoneProgressCreate(ctx, &params)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceApplyEdit: // request
		defer logger.Debug(MethodWorkspaceApplyEdit, zap.Error(err))

		var params ApplyWorkspaceEditParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := client.WorkspaceApplyEdit(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWorkspaceCodeLensRefresh: // request
		defer logger.Debug(MethodWorkspaceCodeLensRefresh, zap.Error(err))

		err := client.WorkspaceCodeLensRefresh(ctx)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceConfiguration: // request
		defer logger.Debug(MethodWorkspaceConfiguration, zap.Error(err))

		var params ConfigurationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := client.WorkspaceConfiguration(ctx, &params)

		return true, reply(ctx, resp, err)

	case MethodWorkspaceDiagnosticRefresh: // request
		defer logger.Debug(MethodWorkspaceDiagnosticRefresh, zap.Error(err))

		err := client.WorkspaceDiagnosticRefresh(ctx)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceFoldingRangeRefresh: // request
		defer logger.Debug(MethodWorkspaceFoldingRangeRefresh, zap.Error(err))

		err := client.WorkspaceFoldingRangeRefresh(ctx)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceInlayHintRefresh: // request
		defer logger.Debug(MethodWorkspaceInlayHintRefresh, zap.Error(err))

		err := client.WorkspaceInlayHintRefresh(ctx)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceSemanticTokensRefresh: // request
		defer logger.Debug(MethodWorkspaceSemanticTokensRefresh, zap.Error(err))

		err := client.WorkspaceSemanticTokensRefresh(ctx)

		return true, reply(ctx, nil, err)

	case MethodWorkspaceWorkspaceFolders: // request
		defer logger.Debug(MethodWorkspaceWorkspaceFolders, zap.Error(err))

		if len(req.Params()) > 0 {
			return true, reply(ctx, nil, fmt.Errorf("expected no params: %w", jsonrpc2.ErrInvalidParams))
		}

		resp, err := client.WorkspaceWorkspaceFolders(ctx)

		return true, reply(ctx, resp, err)

	default:
		return false, nil
	}
}

// client implements a Language Server Protocol client.
type client struct {
	jsonrpc2.Conn

	logger *zap.Logger
}

// compiler time check whether the Client implements ClientInterface interface.
var _ Client = (*client)(nil)

func (c *client) CancelRequest(ctx context.Context, params *CancelParams) (err error) {
	c.logger.Debug("notify " + MethodClientCancelRequest)
	defer c.logger.Debug("end "+MethodClientCancelRequest, zap.Error(err))

	return c.Conn.Notify(ctx, MethodClientCancelRequest, params)
}

// Progress is the base protocol offers also support to report progress in a generic fashion.
//
// This mechanism can be used to report any kind of progress including work done progress (usually used to report progress in the user interface using a progress bar) and
// partial result progress to support streaming of results.
//
// @since 3.16.0.
func (c *client) Progress(ctx context.Context, params *ProgressParams) (err error) {
	c.logger.Debug("notify " + MethodClientProgress)
	defer c.logger.Debug("end "+MethodClientProgress, zap.Error(err))

	return c.Conn.Notify(ctx, MethodClientProgress, params)
}

func (c *client) LogTrace(ctx context.Context, params *LogTraceParams) (err error) {
	c.logger.Debug("notify " + MethodLogTrace)
	defer c.logger.Debug("end "+MethodLogTrace, zap.Error(err))

	return c.Conn.Notify(ctx, MethodLogTrace, params)
}

// Telemetry sends the notification from the server to the client to ask the client to log a telemetry event.
func (c *client) TelemetryEvent(ctx context.Context, params any) (err error) {
	c.logger.Debug("notify " + MethodTelemetryEvent)
	defer c.logger.Debug("end "+MethodTelemetryEvent, zap.Error(err))

	return c.Conn.Notify(ctx, MethodTelemetryEvent, params)
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
func (c *client) TextDocumentPublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) (err error) {
	c.logger.Debug("notify " + MethodTextDocumentPublishDiagnostics)
	defer c.logger.Debug("end "+MethodTextDocumentPublishDiagnostics, zap.Error(err))

	return c.Conn.Notify(ctx, MethodTextDocumentPublishDiagnostics, params)
}

// LogMessage sends the notification from the server to the client to ask the client to log a particular message.
func (c *client) WindowLogMessage(ctx context.Context, params *LogMessageParams) (err error) {
	c.logger.Debug("notify " + MethodWindowLogMessage)
	defer c.logger.Debug("end "+MethodWindowLogMessage, zap.Error(err))

	return c.Conn.Notify(ctx, MethodWindowLogMessage, params)
}

// ShowMessage sends the notification from a server to a client to ask the
// client to display a particular message in the user interface.
func (c *client) WindowShowMessage(ctx context.Context, params *ShowMessageParams) (err error) {
	c.logger.Debug("notify " + MethodWindowShowMessage)
	defer c.logger.Debug("end "+MethodWindowShowMessage, zap.Error(err))

	return c.Conn.Notify(ctx, MethodWindowShowMessage, params)
}

// RegisterCapability sends the request from the server to the client to register for a new capability on the client side.
//
// Not all clients need to support dynamic capability registration.
//
// A client opts in via the dynamicRegistration property on the specific client capabilities.
// A client can even provide dynamic registration for capability A but not for capability B (see TextDocumentClientCapabilities as an example).
func (c *client) ClientRegisterCapability(ctx context.Context, params *RegistrationParams) (err error) {
	c.logger.Debug("call " + MethodClientRegisterCapability)
	defer c.logger.Debug("end "+MethodClientRegisterCapability, zap.Error(err))

	return Call(ctx, c.Conn, MethodClientRegisterCapability, params, nil)
}

// UnregisterCapability sends the request from the server to the client to unregister a previously registered capability.
func (c *client) ClientUnregisterCapability(ctx context.Context, params *UnregistrationParams) (err error) {
	c.logger.Debug("call " + MethodClientUnregisterCapability)
	defer c.logger.Debug("end "+MethodClientUnregisterCapability, zap.Error(err))

	return Call(ctx, c.Conn, MethodClientUnregisterCapability, params, nil)
}

// ShowMessage sends the notification from a server to a client to ask the
// client to display a particular message in the user interface.
func (c *client) WindowShowDocument(ctx context.Context, params *ShowDocumentParams) (_ *ShowDocumentResult, err error) {
	c.logger.Debug("call " + MethodWindowShowDocument)
	defer c.logger.Debug("end "+MethodWindowShowDocument, zap.Error(err))

	var result *ShowDocumentResult
	if err := Call(ctx, c.Conn, MethodWindowShowDocument, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ShowMessageRequest sends the request from a server to a client to ask the client to display a particular message in the user interface.
//
// In addition to the show message notification the request allows to pass actions and to wait for an answer from the client.
func (c *client) WindowShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (_ *MessageActionItem, err error) {
	c.logger.Debug("call " + MethodWindowShowMessageRequest)
	defer c.logger.Debug("end "+MethodWindowShowMessageRequest, zap.Error(err))

	var result *MessageActionItem
	if err := Call(ctx, c.Conn, MethodWindowShowMessageRequest, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// WorkDoneProgressCreate sends the request is sent from the server to the client to ask the client to create a work done progress.
//
// @since 3.16.0.
func (c *client) WindowWorkDoneProgressCreate(ctx context.Context, params *WorkDoneProgressCreateParams) (err error) {
	c.logger.Debug("call " + MethodWindowWorkDoneProgressCreate)
	defer c.logger.Debug("end "+MethodWindowWorkDoneProgressCreate, zap.Error(err))

	return Call(ctx, c.Conn, MethodWindowWorkDoneProgressCreate, params, nil)
}

// ApplyEdit sends the request from the server to the client to modify resource on the client side.
func (c *client) WorkspaceApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (result *ApplyWorkspaceEditResult, err error) {
	c.logger.Debug("call " + MethodWorkspaceApplyEdit)
	defer c.logger.Debug("end "+MethodWorkspaceApplyEdit, zap.Error(err))

	if err := Call(ctx, c.Conn, MethodWorkspaceApplyEdit, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) WorkspaceCodeLensRefresh(ctx context.Context) (err error) {
	return c.refresh(ctx, MethodWorkspaceCodeLensRefresh)
}

// Configuration sends the request from the server to the client to fetch configuration settings from the client.
//
// The request can fetch several configuration settings in one roundtrip.
// The order of the returned configuration settings correspond to the order of the
// passed ConfigurationItems (e.g. the first item in the response is the result for the first configuration item in the params).
func (c *client) WorkspaceConfiguration(ctx context.Context, params *ConfigurationParams) (_ []any, err error) {
	c.logger.Debug("call " + MethodWorkspaceConfiguration)
	defer c.logger.Debug("end "+MethodWorkspaceConfiguration, zap.Error(err))

	var result []any
	if err := Call(ctx, c.Conn, MethodWorkspaceConfiguration, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) refresh(ctx context.Context, method string) (err error) {
	c.logger.Debug("call " + method)
	defer c.logger.Debug("end "+method, zap.Error(err))

	return c.Conn.Notify(ctx, method, nil)
}

func (c *client) WorkspaceDiagnosticRefresh(ctx context.Context) (err error) {
	return c.refresh(ctx, MethodWorkspaceDiagnosticRefresh)
}

func (c *client) WorkspaceFoldingRangeRefresh(ctx context.Context) (err error) {
	return c.refresh(ctx, MethodWorkspaceFoldingRangeRefresh)
}

func (c *client) WorkspaceInlayHintRefresh(ctx context.Context) (err error) {
	return c.refresh(ctx, MethodWorkspaceInlayHintRefresh)
}

func (c *client) WorkspaceInlineValueRefresh(ctx context.Context) (err error) {
	return c.refresh(ctx, MethodWorkspaceInlineValueRefresh)
}

func (c *client) WorkspaceSemanticTokensRefresh(ctx context.Context) (err error) {
	return c.refresh(ctx, MethodWorkspaceSemanticTokensRefresh)
}

// WorkspaceTextDocumentContentRefresh the `workspace/textDocumentContent` request is sent from the server to the client to refresh the content of a specific text document. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
func (c *client) WorkspaceTextDocumentContentRefresh(ctx context.Context, params *TextDocumentContentRefreshParams) (err error) {
	c.logger.Debug("call " + MethodWorkspaceTextDocumentContentRefresh)
	defer c.logger.Debug("end "+MethodWorkspaceTextDocumentContentRefresh, zap.Error(err))

	if err := Call(ctx, c.Conn, MethodWorkspaceTextDocumentContentRefresh, params, nil); err != nil {
		return err
	}

	return nil
}

// WorkspaceFolders sends the request from the server to the client to fetch the current open list of workspace folders.
//
// Returns null in the response if only a single file is open in the tool. Returns an empty array if a workspace is open but no folders are configured.
//
// @since 3.6.0.
func (c *client) WorkspaceWorkspaceFolders(ctx context.Context) (result []*WorkspaceFolder, err error) {
	c.logger.Debug("call " + MethodWorkspaceWorkspaceFolders)
	defer c.logger.Debug("end "+MethodWorkspaceWorkspaceFolders, zap.Error(err))

	if err := Call(ctx, c.Conn, MethodWorkspaceWorkspaceFolders, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}
