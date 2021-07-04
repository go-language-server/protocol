// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestPosition(t *testing.T) {
	t.Parallel()

	testPosition(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRange(t *testing.T) {
	t.Parallel()

	testRange(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestLocation(t *testing.T) {
	t.Parallel()

	testLocation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestLocationLink(t *testing.T) {
	t.Parallel()

	testLocationLink(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeDescription(t *testing.T) {
	t.Parallel()

	testCodeDescription(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCommand(t *testing.T) {
	t.Parallel()

	testCommand(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestChangeAnnotation(t *testing.T) {
	t.Parallel()

	testChangeAnnotation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestAnnotatedTextEdit(t *testing.T) {
	t.Parallel()

	testAnnotatedTextEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextEdit(t *testing.T) {
	t.Parallel()

	testTextEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentEdit(t *testing.T) {
	t.Parallel()

	testTextDocumentEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCreateFileOptions(t *testing.T) {
	t.Parallel()

	testCreateFileOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCreateFile(t *testing.T) {
	t.Parallel()

	testCreateFile(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameFileOptions(t *testing.T) {
	t.Parallel()

	testRenameFileOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameFile(t *testing.T) {
	t.Parallel()

	testRenameFile(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeleteFileOptions(t *testing.T) {
	t.Parallel()

	testDeleteFileOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeleteFile(t *testing.T) {
	t.Parallel()

	testDeleteFile(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceEdit(t *testing.T) {
	t.Parallel()

	testWorkspaceEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentIdentifier(t *testing.T) {
	t.Parallel()

	testTextDocumentIdentifier(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentItem(t *testing.T) {
	t.Parallel()

	testTextDocumentItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestVersionedTextDocumentIdentifier(t *testing.T) {
	t.Parallel()

	testVersionedTextDocumentIdentifier(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestOptionalVersionedTextDocumentIdentifier(t *testing.T) {
	t.Parallel()

	testOptionalVersionedTextDocumentIdentifier(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentPositionParams(t *testing.T) {
	t.Parallel()

	testTextDocumentPositionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentFilter(t *testing.T) {
	t.Parallel()

	testDocumentFilter(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSelector(t *testing.T) {
	t.Parallel()

	testDocumentSelector(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestMarkupContent(t *testing.T) {
	t.Parallel()

	testMarkupContent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
