// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"go.lsp.dev/jsonrpc2"
)

// clientDispatch dispatches client.
//nolint:funlen,cyclop
func clientDispatch(ctx context.Context, client Client, req *jsonrpc2.Request) (resp interface{}, err error) {
	if ctx.Err() != nil {
		return nil, ErrRequestCancelled
	}

	dec := json.NewDecoder(bytes.NewReader(req.Params))
	logger := LoggerFromContext(ctx)

	switch req.Method {
	case MethodProgress: // notification
		defer logger.Debug(MethodProgress, zap.Error(err))

		var params ProgressParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodInitialize, err)
		}

		err = client.Progress(ctx, &params)

		return nil, err

	case MethodWorkDoneProgressCreate: // request
		defer logger.Debug(MethodWorkDoneProgressCreate, zap.Error(err))

		var params WorkDoneProgressCreateParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodWorkDoneProgressCreate, err)
		}

		err = client.WorkDoneProgressCreate(ctx, &params)

		return nil, err

	case MethodWindowLogMessage: // notification
		defer logger.Debug(MethodWindowLogMessage, zap.Error(err))

		var params LogMessageParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodWindowLogMessage, err)
		}

		err = client.LogMessage(ctx, &params)

		return nil, err

	case MethodTextDocumentPublishDiagnostics: // notification
		defer logger.Debug(MethodTextDocumentPublishDiagnostics, zap.Error(err))

		var params PublishDiagnosticsParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTextDocumentPublishDiagnostics, err)
		}

		err = client.PublishDiagnostics(ctx, &params)

		return nil, err

	case MethodWindowShowMessage: // notification
		defer logger.Debug(MethodWindowShowMessage, zap.Error(err))

		var params ShowMessageParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodWindowShowMessage, err)
		}

		err = client.ShowMessage(ctx, &params)

		return nil, err

	case MethodWindowShowMessageRequest: // request
		defer logger.Debug(MethodWindowShowMessageRequest, zap.Error(err))

		var params ShowMessageRequestParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodWindowShowMessageRequest, err)
		}

		resp, err = client.ShowMessageRequest(ctx, &params)

		return resp, err

	case MethodTelemetryEvent: // notification
		defer logger.Debug(MethodTelemetryEvent, zap.Error(err))

		var params interface{}
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodTelemetryEvent, err)
		}

		err = client.Telemetry(ctx, &params)

		return nil, err

	case MethodClientRegisterCapability: // request
		defer logger.Debug(MethodClientRegisterCapability, zap.Error(err))

		var params RegistrationParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodClientRegisterCapability, err)
		}

		err = client.RegisterCapability(ctx, &params)

		return nil, err

	case MethodClientUnregisterCapability: // request
		defer logger.Debug(MethodClientUnregisterCapability, zap.Error(err))

		var params UnregistrationParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodClientUnregisterCapability, err)
		}

		err = client.UnregisterCapability(ctx, &params)

		return nil, err

	case MethodWorkspaceApplyEdit: // request
		defer logger.Debug(MethodWorkspaceApplyEdit, zap.Error(err))

		var params ApplyWorkspaceEditParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodWorkspaceApplyEdit, err)
		}

		resp, err = client.ApplyEdit(ctx, &params)

		return resp, err

	case MethodWorkspaceConfiguration: // request
		defer logger.Debug(MethodWorkspaceConfiguration, zap.Error(err))

		var params ConfigurationParams
		if err := dec.Decode(&params); err != nil {
			return nil, fmt.Errorf("%w: %s: %s", jsonrpc2.ErrParse, MethodWorkspaceConfiguration, err)
		}

		resp, err = client.Configuration(ctx, &params)

		return resp, err

	case MethodWorkspaceWorkspaceFolders: // request
		defer logger.Debug(MethodWorkspaceWorkspaceFolders, zap.Error(err))

		if len(req.Params) > 0 {
			return nil, fmt.Errorf("%s: expected no params: %w", MethodWorkspaceWorkspaceFolders, jsonrpc2.ErrInvalidParams)
		}

		resp, err = client.WorkspaceFolders(ctx)

		return resp, err

	default:
		return nil, errors.New("unknown method")
	}
}
