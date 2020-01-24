// Copyright 2020 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestDidOpenTextDocumentParams(t *testing.T) {
	testDidOpenTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeTextDocumentParams(t *testing.T) {
	testDidChangeTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocument(t *testing.T) {
	testTextDocument(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentChangeEvent(t *testing.T) {
	testTextDocumentChangeEvent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentWillSaveEvent(t *testing.T) {
	testTextDocumentWillSaveEvent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentContentChangeEvent(t *testing.T) {
	testTextDocumentContentChangeEvent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentChangeRegistrationOptions(t *testing.T) {
	testTextDocumentChangeRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWillSaveTextDocumentParams(t *testing.T) {
	testWillSaveTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidSaveTextDocumentParams(t *testing.T) {
	testDidSaveTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentSaveRegistrationOptions(t *testing.T) {
	testTextDocumentSaveRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidCloseTextDocumentParams(t *testing.T) {
	testDidCloseTextDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
