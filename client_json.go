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
func clientDispatch(ctx context.Context, client Client, reply jsonrpc2.Replier, req jsonrpc2.Request) (bool, error) {
	if ctx.Err() != nil {
		return true, reply(ctx, nil, ErrRequestCancelled)
	}

	dec := json.NewDecoder(bytes.NewReader(req.Params()))
	logger := LoggerFromContext(ctx)

	switch req.Method() {
	case MethodClientRegisterCapability: // request
		var params RegistrationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.RegisterCapability(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodClientRegisterCapability, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodClientUnregisterCapability: // request
		var params UnregistrationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.UnregisterCapability(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodClientUnregisterCapability, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodTelemetryEvent: // notification
		var params interface{}
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.Telemetry(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodTelemetryEvent, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodTextDocumentPublishDiagnostics: // notification
		var params PublishDiagnosticsParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.PublishDiagnostics(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodTextDocumentPublishDiagnostics, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodWindowLogMessage: // notification
		var params LogMessageParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.LogMessage(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodWindowLogMessage, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodWindowShowMessage: // notification
		var params ShowMessageParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		err := client.ShowMessage(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, nil, err)
			if replyErr != nil {
				logger.Error(MethodWindowShowMessage, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, nil, err)

	case MethodWindowShowMessageRequest: // request
		var params ShowMessageRequestParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := client.ShowMessageRequest(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodWindowShowMessageRequest, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodWorkspaceApplyEdit: // request
		var params ApplyWorkspaceEditParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := client.WorkspaceApplyEdit(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodWorkspaceApplyEdit, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodWorkspaceConfiguration: // request
		var params ConfigurationParams
		if err := dec.Decode(&params); err != nil {
			return true, replyParseError(ctx, reply, err)
		}

		resp, err := client.WorkspaceConfiguration(ctx, &params)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodWorkspaceConfiguration, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	case MethodWorkspaceWorkspaceFolders: // request
		if len(req.Params()) > 0 {
			return true, reply(ctx, nil, fmt.Errorf("%w: expected no params", jsonrpc2.ErrInvalidParams))
		}

		resp, err := client.WorkspaceFolders(ctx)
		reply = func(ctx context.Context, result interface{}, err error) error {
			replyErr := reply(ctx, result, err)
			if replyErr != nil {
				logger.Error(MethodWorkspaceWorkspaceFolders, zap.Error(replyErr))
			}
			return replyErr
		}
		return true, reply(ctx, resp, err)

	default:
		return false, nil
	}
}
