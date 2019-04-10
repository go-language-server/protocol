// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"github.com/francoispqt/gojay"
	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"
)

// Client represents a implementation of language-server-protocol client.
type Client interface {
	Run(ctx context.Context) (err error)
	RegisterCapability(ctx context.Context, params *RegistrationParams) (err error)
	UnregisterCapability(ctx context.Context, params *UnregistrationParams) (err error)
	Telemetry(ctx context.Context, params interface{}) (err error)
	PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) (err error)
	LogMessage(ctx context.Context, params *LogMessageParams) (err error)
	ShowMessage(ctx context.Context, params *ShowMessageParams) (err error)
	ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (result *MessageActionItem, err error)
	WorkspaceApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (result bool, err error)
	WorkspaceConfiguration(ctx context.Context, params *ConfigurationParams) (result []interface{}, err error)
	WorkspaceFolders(ctx context.Context) (result []WorkspaceFolder, err error)
}

const (
	clientRegisterCapability       = "client/registerCapability"
	clientUnregisterCapability     = "client/unregisterCapability"
	telemetryEvent                 = "telemetry/event"
	textDocumentPublishDiagnostics = "textDocument/publishDiagnostics"
	windowLogMessage               = "window/logMessage"
	windowShowMessage              = "window/showMessage"
	windowShowMessageRequest       = "window/showMessageRequest"
	workspaceApplyEdit             = "workspace/applyEdit"
	workspaceConfiguration         = "workspace/configuration"
	workspaceWorkspaceFolders      = "workspace/workspaceFolders"
)

type client struct {
	*jsonrpc2.Conn
}

var _ Client = (*client)(nil)

func (c *client) Run(ctx context.Context) (err error) {
	err = c.Conn.Run(ctx)
	return
}

func (c *client) RegisterCapability(ctx context.Context, params *RegistrationParams) (err error) {
	err = c.Conn.Notify(ctx, clientRegisterCapability, params)
	return
}

func (c *client) UnregisterCapability(ctx context.Context, params *UnregistrationParams) (err error) {
	err = c.Conn.Notify(ctx, clientUnregisterCapability, params)
	return
}

func (c *client) Telemetry(ctx context.Context, params interface{}) (err error) {
	err = c.Conn.Notify(ctx, telemetryEvent, params)
	return
}

func (c *client) PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) (err error) {
	err = c.Conn.Notify(ctx, textDocumentPublishDiagnostics, params)
	return
}

func (c *client) LogMessage(ctx context.Context, params *LogMessageParams) (err error) {
	err = c.Conn.Notify(ctx, windowLogMessage, params)
	return
}

func (c *client) ShowMessage(ctx context.Context, params *ShowMessageParams) (err error) {
	err = c.Conn.Notify(ctx, windowShowMessage, params)
	return
}

func (c *client) ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (result *MessageActionItem, err error) {
	result = new(MessageActionItem)
	err = c.Conn.Call(ctx, windowShowMessageRequest, params, result)

	return result, err
}

func (c *client) WorkspaceApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (_ bool, err error) {
	var result bool
	err = c.Conn.Call(ctx, workspaceApplyEdit, params, &result)

	return result, err
}

func (c *client) WorkspaceConfiguration(ctx context.Context, params *ConfigurationParams) (_ []interface{}, err error) {
	var result []interface{}
	err = c.Conn.Call(ctx, workspaceConfiguration, params, &result)

	return result, err
}

func (c *client) WorkspaceFolders(ctx context.Context) (result []WorkspaceFolder, err error) {
	result = []WorkspaceFolder{}
	err = c.Conn.Call(ctx, workspaceWorkspaceFolders, nil, &result)

	return result, err
}

// ClientHandler returns the client handler.
func ClientHandler(client Client, logger *zap.Logger) jsonrpc2.Handler {
	return func(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
		dec := gojay.BorrowDecoder(r.Params)
		defer dec.Release()

		switch r.Method {
		case cancelRequest:
			var params CancelParams
			if err := dec.Decode(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			conn.Cancel(params.ID)

		case clientRegisterCapability:
			var params RegistrationParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := client.RegisterCapability(ctx, &params); err != nil {
				logger.Error(clientRegisterCapability, zap.Error(err))
			}

		case clientUnregisterCapability:
			var params UnregistrationParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := client.UnregisterCapability(ctx, &params); err != nil {
				logger.Error(clientUnregisterCapability, zap.Error(err))
			}

		case telemetryEvent:
			var params interface{}
			if err := dec.DecodeInterface(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := client.Telemetry(ctx, &params); err != nil {
				logger.Error(telemetryEvent, zap.Error(err))
			}

		case textDocumentPublishDiagnostics:
			var params PublishDiagnosticsParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := client.PublishDiagnostics(ctx, &params); err != nil {
				logger.Error(textDocumentPublishDiagnostics, zap.Error(err))
			}

		case windowLogMessage:
			var params LogMessageParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := client.LogMessage(ctx, &params); err != nil {
				logger.Error(windowLogMessage, zap.Error(err))
			}

		case windowShowMessage:
			var params ShowMessageParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			if err := client.ShowMessage(ctx, &params); err != nil {
				logger.Error(windowShowMessage, zap.Error(err))
			}

		case windowShowMessageRequest:
			var params ShowMessageRequestParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := client.ShowMessageRequest(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(windowShowMessageRequest, zap.Error(err))
			}

		case workspaceApplyEdit:
			var params ApplyWorkspaceEditParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := client.WorkspaceApplyEdit(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(workspaceApplyEdit, zap.Error(err))
			}

		case workspaceConfiguration:
			var params ConfigurationParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			resp, err := client.WorkspaceConfiguration(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(workspaceConfiguration, zap.Error(err))
			}

		case workspaceWorkspaceFolders:
			if r.Params != nil {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeInvalidParams, "Expected no params"))
				return
			}
			resp, err := client.WorkspaceFolders(ctx)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(workspaceWorkspaceFolders, zap.Error(err))
			}

		default:
			if r.IsNotify() {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeMethodNotFound, "method %q not found", r.Method))
			}
		}
	}
}
