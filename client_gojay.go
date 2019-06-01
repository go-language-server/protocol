// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"bytes"
	"context"

	"github.com/go-language-server/jsonrpc2"
	"go.uber.org/zap"

	"github.com/go-language-server/protocol/internal/gojaypool"
)

// ClientHandler returns the client handler.
func ClientHandler(ctx context.Context, client ClientInterface, logger *zap.Logger) jsonrpc2.Handler {
	return func(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
		dec := gojaypool.BorrowSizedDecoder(bytes.NewReader(*r.Params), len(*r.Params))
		defer dec.Release()

		switch r.Method {
		case MethodCancelRequest:
			var params CancelParams
			if err := dec.Decode(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}
			conn.Cancel(params.ID)

		case MethodClientRegisterCapability:
			var params RegistrationParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}

			err := client.RegisterCapability(ctx, &params)
			if err := conn.Reply(ctx, r, nil, err); err != nil {
				logger.Error(MethodClientRegisterCapability, zap.Error(err))
			}

		case MethodClientUnregisterCapability:
			var params UnregistrationParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}

			err := client.UnregisterCapability(ctx, &params)
			if err := conn.Reply(ctx, r, nil, err); err != nil {
				logger.Error(MethodClientUnregisterCapability, zap.Error(err))
			}

		case MethodTelemetryEvent:
			var params interface{}
			if err := dec.Decode(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}

			if err := client.Telemetry(ctx, &params); err != nil {
				logger.Error(MethodTelemetryEvent, zap.Error(err))
			}

		case MethodTextDocumentPublishDiagnostics:
			var params PublishDiagnosticsParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}

			if err := client.PublishDiagnostics(ctx, &params); err != nil {
				logger.Error(MethodTextDocumentPublishDiagnostics, zap.Error(err))
			}

		case MethodWindowLogMessage:
			var params LogMessageParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}

			if err := client.LogMessage(ctx, &params); err != nil {
				logger.Error(MethodWindowLogMessage, zap.Error(err))
			}

		case MethodWindowShowMessage:
			var params ShowMessageParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}

			if err := client.ShowMessage(ctx, &params); err != nil {
				logger.Error(MethodWindowShowMessage, zap.Error(err))
			}

		case MethodWindowShowMessageRequest:
			var params ShowMessageRequestParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}

			resp, err := client.ShowMessageRequest(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodWindowShowMessageRequest, zap.Error(err))
			}

		case MethodWorkspaceApplyEdit:
			var params ApplyWorkspaceEditParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}

			resp, err := client.WorkspaceApplyEdit(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodWorkspaceApplyEdit, zap.Error(err))
			}

		case MethodWorkspaceConfiguration:
			var params ConfigurationParams
			if err := dec.DecodeObject(&params); err != nil {
				ReplyError(ctx, err, conn, r, logger)
				return
			}

			resp, err := client.WorkspaceConfiguration(ctx, &params)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodWorkspaceConfiguration, zap.Error(err))
			}

		case MethodWorkspaceWorkspaceFolders:
			if r.Params != nil {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeInvalidParams, "Expected no params"))
				return
			}

			resp, err := client.WorkspaceFolders(ctx)
			if err := conn.Reply(ctx, r, resp, err); err != nil {
				logger.Error(MethodWorkspaceWorkspaceFolders, zap.Error(err))
			}

		default:
			if r.IsNotify() {
				conn.Reply(ctx, r, nil, jsonrpc2.Errorf(jsonrpc2.CodeMethodNotFound, "method %q not found", r.Method))
			}
		}
	}
}
