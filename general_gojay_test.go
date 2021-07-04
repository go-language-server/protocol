// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestWorkspaceFolders(t *testing.T) {
	t.Parallel()

	testWorkspaceFolders(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestClientInfo(t *testing.T) {
	t.Parallel()

	testClientInfo(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializeParams(t *testing.T) {
	t.Parallel()

	testInitializeParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestLogTraceParams(t *testing.T) {
	t.Parallel()

	testLogTraceParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSetTraceParams(t *testing.T) {
	t.Parallel()

	testSetTraceParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCreateFilesParams(t *testing.T) {
	t.Parallel()

	testCreateFilesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameFilesParams(t *testing.T) {
	t.Parallel()

	testRenameFilesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeleteFilesParams(t *testing.T) {
	t.Parallel()

	testDeleteFilesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferencesParams(t *testing.T) {
	t.Parallel()

	testReferencesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlightOptions(t *testing.T) {
	t.Parallel()

	testDocumentHighlightOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlightParams(t *testing.T) {
	t.Parallel()

	testDocumentHighlightParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSymbolOptions(t *testing.T) {
	t.Parallel()

	testDocumentSymbolOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceSymbolOptions(t *testing.T) {
	t.Parallel()

	testWorkspaceSymbolOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentFormattingOptions(t *testing.T) {
	t.Parallel()

	testDocumentFormattingOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentRangeFormattingOptions(t *testing.T) {
	t.Parallel()

	testDocumentRangeFormattingOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeclarationOptions(t *testing.T) {
	t.Parallel()

	testDeclarationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeclarationRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDeclarationRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeclarationParams(t *testing.T) {
	t.Parallel()

	testDeclarationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDefinitionOptions(t *testing.T) {
	t.Parallel()

	testDefinitionOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDefinitionParams(t *testing.T) {
	t.Parallel()

	testDefinitionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTypeDefinitionOptions(t *testing.T) {
	t.Parallel()

	testTypeDefinitionOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTypeDefinitionRegistrationOptions(t *testing.T) {
	t.Parallel()

	testTypeDefinitionRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTypeDefinitionParams(t *testing.T) {
	t.Parallel()

	testTypeDefinitionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestImplementationOptions(t *testing.T) {
	t.Parallel()

	testImplementationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestImplementationRegistrationOptions(t *testing.T) {
	t.Parallel()

	testImplementationRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestImplementationParams(t *testing.T) {
	t.Parallel()

	testImplementationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentColorOptions(t *testing.T) {
	t.Parallel()

	testDocumentColorOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentColorRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDocumentColorRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeOptions(t *testing.T) {
	t.Parallel()

	testFoldingRangeOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeRegistrationOptions(t *testing.T) {
	t.Parallel()

	testFoldingRangeRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializeResult(t *testing.T) {
	t.Parallel()

	testInitializeResult(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializeError(t *testing.T) {
	t.Parallel()

	testInitializeError(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestShowDocumentParams(t *testing.T) {
	t.Parallel()

	testShowDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestShowDocumentResult(t *testing.T) {
	t.Parallel()

	testShowDocumentResult(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferencesOptions(t *testing.T) {
	t.Parallel()

	testReferencesOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeActionOptions(t *testing.T) {
	t.Parallel()

	testCodeActionOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameOptions(t *testing.T) {
	t.Parallel()

	testRenameOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSaveOptions(t *testing.T) {
	t.Parallel()

	testSaveOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentSyncOptions(t *testing.T) {
	t.Parallel()

	testTextDocumentSyncOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestHoverOptions(t *testing.T) {
	t.Parallel()

	testHoverOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestStaticRegistrationOptions(t *testing.T) {
	t.Parallel()

	testStaticRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentLinkRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDocumentLinkRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializedParams(t *testing.T) {
	t.Parallel()

	testInitializedParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
