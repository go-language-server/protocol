// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !gojay

package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
)

// clientDispatch implements jsonrpc2.Handler.
//nolint:gocognit,funlen
func clientDispatch(ctx context.Context, client Client, reply jsonrpc2.Replier, req jsonrpc2.Request) (handled bool, err error) {
	if ctx.Err() != nil {
		return true, reply(ctx, nil, ErrRequestCancelled)
	}

	dec := json.NewDecoder(bytes.NewReader(req.Params()))
	logger := LoggerFromContext(ctx)

	switch req.Method() {
	case MethodWindowLogMessage: // notification
		defer logger.Debug(MethodWindowLogMessage, zap.Error(err))

		var params LogMessageParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}
		err := client.LogMessage(ctx, &params)
		return true, reply(ctx, nil, err)

	case MethodTextDocumentPublishDiagnostics: // notification
		defer logger.Debug(MethodTextDocumentPublishDiagnostics, zap.Error(err))

		var params PublishDiagnosticsParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}
		err := client.PublishDiagnostics(ctx, &params)
		return true, reply(ctx, nil, err)

	case MethodWindowShowMessage: // notification
		defer logger.Debug(MethodWindowShowMessage, zap.Error(err))

		var params ShowMessageParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}
		err := client.ShowMessage(ctx, &params)
		return true, reply(ctx, nil, err)

	case MethodWindowShowMessageRequest: // request
		defer logger.Debug(MethodWindowShowMessageRequest, zap.Error(err))

		var params ShowMessageRequestParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}
		resp, err := client.ShowMessageRequest(ctx, &params)
		return true, reply(ctx, resp, err)

	case MethodTelemetryEvent: // notification
		defer logger.Debug(MethodTelemetryEvent, zap.Error(err))

		var params interface{}
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}
		err := client.Telemetry(ctx, &params)
		return true, reply(ctx, nil, err)

	case MethodClientRegisterCapability: // request
		defer logger.Debug(MethodClientRegisterCapability, zap.Error(err))

		var params RegistrationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}
		err := client.RegisterCapability(ctx, &params)
		return true, reply(ctx, nil, err)

	case MethodClientUnregisterCapability: // request
		defer logger.Debug(MethodClientUnregisterCapability, zap.Error(err))

		var params UnregistrationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}
		err := client.UnregisterCapability(ctx, &params)
		return true, reply(ctx, nil, err)

	case MethodWorkspaceApplyEdit: // request
		defer logger.Debug(MethodWorkspaceApplyEdit, zap.Error(err))

		var params ApplyWorkspaceEditParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}
		resp, err := client.WorkspaceApplyEdit(ctx, &params)
		return true, reply(ctx, resp, err)

	case MethodWorkspaceConfiguration: // request
		defer logger.Debug(MethodWorkspaceConfiguration, zap.Error(err))

		var params ConfigurationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}
		resp, err := client.WorkspaceConfiguration(ctx, &params)
		return true, reply(ctx, resp, err)

	case MethodWorkspaceWorkspaceFolders: // request
		defer logger.Debug(MethodWorkspaceWorkspaceFolders, zap.Error(err))

		if len(req.Params()) > 0 {
			return true, reply(ctx, nil, fmt.Errorf("expected no params: %w", jsonrpc2.ErrInvalidParams))
		}
		resp, err := client.WorkspaceFolders(ctx)
		return true, reply(ctx, resp, err)

	default:
		return false, nil
	}
}
