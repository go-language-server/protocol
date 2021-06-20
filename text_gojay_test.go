// SPDX-FileCopyrightText: Copyright 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestDidOpenTextDocumentParams(t *testing.T) {
	t.Parallel()

	testDidOpenTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeTextDocumentParams(t *testing.T) {
	t.Parallel()

	testDidChangeTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentContentChangeEvent(t *testing.T) {
	t.Parallel()

	testTextDocumentContentChangeEvent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentChangeRegistrationOptions(t *testing.T) {
	t.Parallel()

	testTextDocumentChangeRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWillSaveTextDocumentParams(t *testing.T) {
	t.Parallel()

	testWillSaveTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidSaveTextDocumentParams(t *testing.T) {
	t.Parallel()

	testDidSaveTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentSaveRegistrationOptions(t *testing.T) {
	t.Parallel()

	testTextDocumentSaveRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidCloseTextDocumentParams(t *testing.T) {
	t.Parallel()

	testDidCloseTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
