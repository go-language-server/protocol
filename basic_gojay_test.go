// SPDX-FileCopyrightText: Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestPosition(t *testing.T) { testPosition(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestRange(t *testing.T) { testRange(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestLocation(t *testing.T) { testLocation(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestLocationLink(t *testing.T) { testLocationLink(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestCodeDescription(t *testing.T) { testCodeDescription(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestCommand(t *testing.T) { testCommand(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestChangeAnnotation(t *testing.T) {
	testChangeAnnotation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestAnnotatedTextEdit(t *testing.T) {
	testAnnotatedTextEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextEdit(t *testing.T) { testTextEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestTextDocumentEdit(t *testing.T) {
	testTextDocumentEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCreateFileOptions(t *testing.T) {
	testCreateFileOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCreateFile(t *testing.T) { testCreateFile(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestRenameFileOptions(t *testing.T) {
	testRenameFileOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameFile(t *testing.T) { testRenameFile(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestDeleteFileOptions(t *testing.T) {
	testDeleteFileOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeleteFile(t *testing.T) { testDeleteFile(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestWorkspaceEdit(t *testing.T) { testWorkspaceEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestTextDocumentIdentifier(t *testing.T) {
	testTextDocumentIdentifier(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentItem(t *testing.T) {
	testTextDocumentItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestVersionedTextDocumentIdentifier(t *testing.T) {
	testVersionedTextDocumentIdentifier(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestOptionalVersionedTextDocumentIdentifier(t *testing.T) {
	testOptionalVersionedTextDocumentIdentifier(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentPositionParams(t *testing.T) {
	testTextDocumentPositionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentFilter(t *testing.T) { testDocumentFilter(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestDocumentSelector(t *testing.T) {
	testDocumentSelector(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestMarkupContent(t *testing.T) { testMarkupContent(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }
