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
	testDidOpenTextDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeTextDocumentParams(t *testing.T) {
	testDidChangeTextDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentContentChangeEvent(t *testing.T) {
	testTextDocumentContentChangeEvent(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentChangeRegistrationOptions(t *testing.T) {
	testTextDocumentChangeRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestWillSaveTextDocumentParams(t *testing.T) {
	testWillSaveTextDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestDidSaveTextDocumentParams(t *testing.T) {
	testDidSaveTextDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentSaveRegistrationOptions(t *testing.T) {
	testTextDocumentSaveRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDidCloseTextDocumentParams(t *testing.T) {
	testDidCloseTextDocumentParams(t, json.Marshal, json.Unmarshal)
}
