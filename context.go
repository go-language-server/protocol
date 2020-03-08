// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"go.uber.org/zap"
)

type contextKey int

const (
	ctxLogger contextKey = 1 + iota
	ctxClient
)

// WithLogger returns the context with zap.Logger value.
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger, logger)
}

// LoggerFromContext extracts zap.Logger from context.
func LoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(ctxLogger).(*zap.Logger)
	if !ok {
		return zap.NewNop()
	}

	return logger
}

// WithClient returns the context with Client value.
func WithClient(ctx context.Context, client ClientInterface) context.Context {
	return context.WithValue(ctx, ctxClient, client)
}
