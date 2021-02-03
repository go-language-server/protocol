// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"bytes"
	"context"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"

	"go.lsp.dev/protocol/internal/gojaypool"
)

// Deliver implements jsonrpc2.Handler.
//nolint:funlen,gocognit
func (h clientHandler) Deliver(ctx context.Context, r *jsonrpc2.Request, delivered bool) bool {
	if delivered {
		return false
	}

	if ctx.Err() != nil {
		r.Reply(ctx, nil, jsonrpc2.NewError(RequestCancelledError, ctx.Err().Error()))
		return true
	}

	dec := gojaypool.BorrowSizedDecoder(bytes.NewReader(*r.Params), len(*r.Params))
	defer dec.Release()
	logger := LoggerFromContext(ctx)

	switch r.Method {
	case MethodClientRegisterCapability: // request
		var params RegistrationParams
		if err := dec.DecodeObject(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		err := h.client.RegisterCapability(ctx, &params)
		if err := r.Reply(ctx, nil, err); err != nil {
			logger.Error(MethodClientRegisterCapability, zap.Error(err))
		}
		return true

	case MethodClientUnregisterCapability: // request
		var params UnregistrationParams
		if err := dec.DecodeObject(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		err := h.client.UnregisterCapability(ctx, &params)
		if err := r.Reply(ctx, nil, err); err != nil {
			logger.Error(MethodClientUnregisterCapability, zap.Error(err))
		}
		return true

	case MethodTelemetryEvent: // notification
		var params interface{}
		if err := dec.Decode(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.client.Telemetry(ctx, &params); err != nil {
			logger.Error(MethodTelemetryEvent, zap.Error(err))
		}
		return true

	case MethodTextDocumentPublishDiagnostics: // notification
		var params PublishDiagnosticsParams
		if err := dec.DecodeObject(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.client.PublishDiagnostics(ctx, &params); err != nil {
			logger.Error(MethodTextDocumentPublishDiagnostics, zap.Error(err))
		}
		return true

	case MethodWindowLogMessage: // notification
		var params LogMessageParams
		if err := dec.DecodeObject(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.client.LogMessage(ctx, &params); err != nil {
			logger.Error(MethodWindowLogMessage, zap.Error(err))
		}
		return true

	case MethodWindowShowMessage: // notification
		var params ShowMessageParams
		if err := dec.DecodeObject(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		if err := h.client.ShowMessage(ctx, &params); err != nil {
			logger.Error(MethodWindowShowMessage, zap.Error(err))
		}
		return true

	case MethodWindowShowMessageRequest: // request
		var params ShowMessageRequestParams
		if err := dec.DecodeObject(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.client.ShowMessageRequest(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodWindowShowMessageRequest, zap.Error(err))
		}
		return true

	case MethodWorkspaceApplyEdit: // request
		var params ApplyWorkspaceEditParams
		if err := dec.DecodeObject(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.client.WorkspaceApplyEdit(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodWorkspaceApplyEdit, zap.Error(err))
		}
		return true

	case MethodWorkspaceConfiguration: // request
		var params ConfigurationParams
		if err := dec.DecodeObject(&params); err != nil {
			ReplyError(ctx, err, r)
			return true
		}
		resp, err := h.client.WorkspaceConfiguration(ctx, &params)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodWorkspaceConfiguration, zap.Error(err))
		}
		return true

	case MethodWorkspaceWorkspaceFolders: // request
		if r.Params != nil {
			if err := r.Reply(ctx, nil, jsonrpc2.Errorf(jsonrpc2.InvalidParams, "Expected no params")); err != nil {
				logger.Error(MethodWorkspaceWorkspaceFolders, zap.Error(err))
			}
			return true
		}
		resp, err := h.client.WorkspaceFolders(ctx)
		if err := r.Reply(ctx, resp, err); err != nil {
			logger.Error(MethodWorkspaceWorkspaceFolders, zap.Error(err))
		}
		return true

	default:
		return false
	}
}
