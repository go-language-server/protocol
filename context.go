// SPDX-FileCopyrightText: 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"context"
	"io"
	"log/slog"
)

type (
	ctxLogger struct{}
	ctxClient struct{}
)

// WithLogger returns the context with slog.Logger value.
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, logger)
}

// LoggerFromContext extracts slog.Logger from context.
func LoggerFromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(ctxLogger{}).(*slog.Logger)
	if !ok {
		return slog.New(slog.NewTextHandler(io.Discard, nil))
	}

	return logger
}

// WithClient returns the context with Client value.
func WithClient(ctx context.Context, client Client) context.Context {
	return context.WithValue(ctx, ctxClient{}, client)
}
