// SPDX-FileCopyrightText: Copyright 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestDidOpenTextDocumentParams(t *testing.T) {
	t.Parallel()

	testDidOpenTextDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeTextDocumentParams(t *testing.T) {
	t.Parallel()

	testDidChangeTextDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentContentChangeEvent(t *testing.T) {
	t.Parallel()

	testTextDocumentContentChangeEvent(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentChangeRegistrationOptions(t *testing.T) {
	t.Parallel()

	testTextDocumentChangeRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestWillSaveTextDocumentParams(t *testing.T) {
	t.Parallel()

	testWillSaveTextDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestDidSaveTextDocumentParams(t *testing.T) {
	t.Parallel()

	testDidSaveTextDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentSaveRegistrationOptions(t *testing.T) {
	t.Parallel()

	testTextDocumentSaveRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDidCloseTextDocumentParams(t *testing.T) {
	t.Parallel()

	testDidCloseTextDocumentParams(t, json.Marshal, json.Unmarshal)
}
