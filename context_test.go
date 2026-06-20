// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"context"
	"log/slog"
	"testing"
)

// TestWithLoggerRoundTrip asserts a logger stored by WithLogger is returned by
// LoggerFromContext, and that the returned logger actually routes records to the
// stored handler.
func TestWithLoggerRoundTrip(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	want := slog.New(slog.NewTextHandler(&buf, nil))
	ctx := WithLogger(t.Context(), want)

	got := LoggerFromContext(ctx)
	if got != want {
		t.Fatalf("LoggerFromContext returned %p, want stored logger %p", got, want)
	}

	got.Info("hello", "k", "v")
	if !bytes.Contains(buf.Bytes(), []byte("hello")) {
		t.Errorf("record not routed to stored handler; buffer = %q", buf.String())
	}
}

// TestLoggerFromContextAbsentReturnsNop covers the fallback branch: a context
// carrying no logger yields the process-wide nop logger (never nil) that
// discards every record.
func TestLoggerFromContextAbsentReturnsNop(t *testing.T) {
	t.Parallel()

	got := LoggerFromContext(t.Context())
	if got == nil {
		t.Fatal("LoggerFromContext returned nil, want nop logger")
	}
	if got != nopLogger {
		t.Errorf("LoggerFromContext returned %p, want nopLogger %p", got, nopLogger)
	}

	// The nop logger must be safe to use unconditionally.
	got.Info("discarded", "k", "v")
}

// TestLoggerFromContextWrongTypeReturnsNop covers the type-assertion failure
// branch: a non-logger value stored under the logger key must not be returned;
// the nop logger is returned instead.
func TestLoggerFromContextWrongTypeReturnsNop(t *testing.T) {
	t.Parallel()

	ctx := context.WithValue(t.Context(), ctxLogger{}, "not a logger")
	if got := LoggerFromContext(ctx); got != nopLogger {
		t.Errorf("LoggerFromContext returned %p for wrong-typed value, want nopLogger %p", got, nopLogger)
	}
}

// TestWithClientRoundTrip asserts a client stored by WithClient is returned by
// ClientFromContext with ok=true.
func TestWithClientRoundTrip(t *testing.T) {
	t.Parallel()

	want := UnimplementedClient{}
	ctx := WithClient(t.Context(), want)

	got, ok := ClientFromContext(ctx)
	if !ok {
		t.Fatal("ClientFromContext ok = false, want true")
	}
	if got != Client(want) {
		t.Errorf("ClientFromContext returned %#v, want %#v", got, want)
	}
}

// TestClientFromContextAbsent covers the absent branch: a context carrying no
// client yields a nil client and ok=false.
func TestClientFromContextAbsent(t *testing.T) {
	t.Parallel()

	got, ok := ClientFromContext(t.Context())
	if ok {
		t.Errorf("ClientFromContext ok = true, want false")
	}
	if got != nil {
		t.Errorf("ClientFromContext returned %#v, want nil", got)
	}
}
