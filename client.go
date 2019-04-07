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
	ShowMessage(context.Context, *ShowMessageParams) error
	ShowMessageRequest(context.Context, *ShowMessageRequestParams) (*MessageActionItem, error)
	LogMessage(context.Context, *LogMessageParams) error
	Telemetry(context.Context, interface{}) error
	RegisterCapability(context.Context, *RegistrationParams) error
	UnregisterCapability(context.Context, *UnregistrationParams) error
	WorkspaceFolders(context.Context) ([]WorkspaceFolder, error)
	Configuration(context.Context, *ConfigurationParams) ([]interface{}, error)
	ApplyEdit(context.Context, *ApplyWorkspaceEditParams) (bool, error)
	PublishDiagnostics(context.Context, *PublishDiagnosticsParams) error
}

const (
	windowShowMessage              = "window/showMessage"
	windowShowMessageRequest       = "window/showMessageRequest"
	windowLogMessage               = "window/logMessage"
	telemetryEvent                 = "telemetry/event"
	clientRegisterCapability       = "client/registerCapability"
	clientUnregisterCapability     = "client/unregisterCapability"
	workspaceWorkspaceFolders      = "workspace/workspaceFolders"
	workspaceConfiguration         = "workspace/configuration"
	workspaceApplyEdit             = "workspace/applyEdit"
	textDocumentPublishDiagnostics = "textDocument/publishDiagnostics"
)

func clientHandler(client Client, log *zap.Logger) jsonrpc2.Handler {
	return func(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
		dec := gojay.Unsafe

		switch r.Method {
		case cancelRequest:
			var params CancelParams
			if err := dec.Unmarshal(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			conn.Cancel(params.ID)

		case windowShowMessage:
			var params ShowMessageParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			if err := client.ShowMessage(ctx, &params); err != nil {
				log.Error(windowShowMessage, zap.Error(err))
			}

		case windowShowMessageRequest:
			var params ShowMessageRequestParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			resp, err := client.ShowMessageRequest(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				log.Error(windowShowMessageRequest, zap.Error(err))
			}

		case windowLogMessage:
			var params LogMessageParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			if err := client.LogMessage(ctx, &params); err != nil {
				log.Error(windowLogMessage, zap.Error(err))
			}

		case telemetryEvent:
			var params interface{}
			if err := dec.Unmarshal(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			if err := client.Telemetry(ctx, &params); err != nil {
				log.Error(telemetryEvent, zap.Error(err))
			}

		case clientRegisterCapability:
			var params RegistrationParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			if err := client.RegisterCapability(ctx, &params); err != nil {
				log.Error(clientRegisterCapability, zap.Error(err))
			}

		case clientUnregisterCapability:
			var params UnregistrationParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			if err := client.UnregisterCapability(ctx, &params); err != nil {
				log.Error(clientUnregisterCapability, zap.Error(err))
			}

		case workspaceWorkspaceFolders:
			if r.Params != nil {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeInvalidParams, "Expected no params"))
				return
			}
			resp, err := client.WorkspaceFolders(ctx)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				log.Error(workspaceWorkspaceFolders, zap.Error(err))
			}

		case workspaceConfiguration:
			var params ConfigurationParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			resp, err := client.Configuration(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				log.Error(workspaceConfiguration, zap.Error(err))
			}

		case workspaceApplyEdit:
			var params ApplyWorkspaceEditParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			resp, err := client.ApplyEdit(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				log.Error(workspaceApplyEdit, zap.Error(err))
			}

		case textDocumentPublishDiagnostics:
			var params PublishDiagnosticsParams
			if err := dec.UnmarshalJSONObject(*r.Params, &params); err != nil {
				replyError(ctx, log, conn, r, err)
				return
			}
			if err := client.PublishDiagnostics(ctx, &params); err != nil {
				log.Error(textDocumentPublishDiagnostics, zap.Error(err))
			}

		default:
			if r.IsNotify() {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeMethodNotFound, "method %q not found", r.Method))
			}
		}
	}
}

type client struct {
	*jsonrpc2.Conn
}

func (c *client) ShowMessage(ctx context.Context, params *ShowMessageParams) error {
	return c.Conn.Notify(ctx, windowShowMessage, params)
}

func (c *client) ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (*MessageActionItem, error) {
	var result MessageActionItem
	if err := c.Conn.Call(ctx, windowShowMessageRequest, params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) LogMessage(ctx context.Context, params *LogMessageParams) error {
	return c.Conn.Notify(ctx, windowLogMessage, params)
}

func (c *client) Telemetry(ctx context.Context, params interface{}) error {
	return c.Conn.Notify(ctx, telemetryEvent, params)
}

func (c *client) RegisterCapability(ctx context.Context, params *RegistrationParams) error {
	return c.Conn.Notify(ctx, clientRegisterCapability, params)
}

func (c *client) UnregisterCapability(ctx context.Context, params *UnregistrationParams) error {
	return c.Conn.Notify(ctx, clientUnregisterCapability, params)
}

func (c *client) WorkspaceFolders(ctx context.Context) ([]WorkspaceFolder, error) {
	var result []WorkspaceFolder
	if err := c.Conn.Call(ctx, workspaceWorkspaceFolders, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) Configuration(ctx context.Context, params *ConfigurationParams) ([]interface{}, error) {
	var result []interface{}
	if err := c.Conn.Call(ctx, workspaceConfiguration, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) ApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (bool, error) {
	var result bool
	if err := c.Conn.Call(ctx, workspaceApplyEdit, params, &result); err != nil {
		return false, err
	}
	return result, nil
}

func (c *client) PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) error {
	return c.Conn.Notify(ctx, textDocumentPublishDiagnostics, params)
}
