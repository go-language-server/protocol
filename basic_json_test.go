// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestPosition(t *testing.T) {
	t.Parallel()

	testPosition(t, json.Marshal, json.Unmarshal)
}

func TestRange(t *testing.T) {
	t.Parallel()

	testRange(t, json.Marshal, json.Unmarshal)
}

func TestLocation(t *testing.T) {
	t.Parallel()

	testLocation(t, json.Marshal, json.Unmarshal)
}

func TestLocationLink(t *testing.T) {
	t.Parallel()

	testLocationLink(t, json.Marshal, json.Unmarshal)
}

func TestCodeDescription(t *testing.T) {
	t.Parallel()

	testCodeDescription(t, json.Marshal, json.Unmarshal)
}

func TestCommand(t *testing.T) {
	t.Parallel()

	testCommand(t, json.Marshal, json.Unmarshal)
}

func TestChangeAnnotation(t *testing.T) {
	t.Parallel()

	testChangeAnnotation(t, json.Marshal, json.Unmarshal)
}

func TestAnnotatedTextEdit(t *testing.T) {
	t.Parallel()

	testAnnotatedTextEdit(t, json.Marshal, json.Unmarshal)
}

func TestTextEdit(t *testing.T) {
	t.Parallel()

	testTextEdit(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentEdit(t *testing.T) {
	t.Parallel()

	testTextDocumentEdit(t, json.Marshal, json.Unmarshal)
}

func TestCreateFileOptions(t *testing.T) {
	t.Parallel()

	testCreateFileOptions(t, json.Marshal, json.Unmarshal)
}

func TestCreateFile(t *testing.T) {
	t.Parallel()

	testCreateFile(t, json.Marshal, json.Unmarshal)
}

func TestRenameFileOptions(t *testing.T) {
	t.Parallel()

	testRenameFileOptions(t, json.Marshal, json.Unmarshal)
}

func TestRenameFile(t *testing.T) {
	t.Parallel()

	testRenameFile(t, json.Marshal, json.Unmarshal)
}

func TestDeleteFileOptions(t *testing.T) {
	t.Parallel()

	testDeleteFileOptions(t, json.Marshal, json.Unmarshal)
}

func TestDeleteFile(t *testing.T) {
	t.Parallel()

	testDeleteFile(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceEdit(t *testing.T) {
	t.Parallel()

	testWorkspaceEdit(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentIdentifier(t *testing.T) {
	t.Parallel()

	testTextDocumentIdentifier(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentItem(t *testing.T) {
	t.Parallel()

	testTextDocumentItem(t, json.Marshal, json.Unmarshal)
}

func TestVersionedTextDocumentIdentifier(t *testing.T) {
	t.Parallel()

	testVersionedTextDocumentIdentifier(t, json.Marshal, json.Unmarshal)
}

func TestOptionalVersionedTextDocumentIdentifier(t *testing.T) {
	t.Parallel()

	testOptionalVersionedTextDocumentIdentifier(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentPositionParams(t *testing.T) {
	t.Parallel()

	testTextDocumentPositionParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentFilter(t *testing.T) {
	t.Parallel()

	testDocumentFilter(t, json.Marshal, json.Unmarshal)
}

func TestDocumentSelector(t *testing.T) {
	t.Parallel()

	testDocumentSelector(t, json.Marshal, json.Unmarshal)
}

func TestMarkupContent(t *testing.T) {
	t.Parallel()

	testMarkupContent(t, json.Marshal, json.Unmarshal)
}
