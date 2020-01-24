// Copyright 2020 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

func TestTextDocument(t *testing.T) {
	testTextDocument(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentChangeEvent(t *testing.T) {
	testTextDocumentChangeEvent(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentWillSaveEvent(t *testing.T) {
	testTextDocumentWillSaveEvent(t, json.Marshal, json.Unmarshal)
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
