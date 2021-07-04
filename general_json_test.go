// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestWorkspaceFolders(t *testing.T) {
	t.Parallel()

	testWorkspaceFolders(t, json.Marshal, json.Unmarshal)
}

func TestClientInfo(t *testing.T) {
	t.Parallel()

	testClientInfo(t, json.Marshal, json.Unmarshal)
}

func TestInitializeParams(t *testing.T) {
	t.Parallel()

	testInitializeParams(t, json.Marshal, json.Unmarshal)
}

func TestLogTraceParams(t *testing.T) {
	t.Parallel()

	testLogTraceParams(t, json.Marshal, json.Unmarshal)
}

func TestSetTraceParams(t *testing.T) {
	t.Parallel()

	testSetTraceParams(t, json.Marshal, json.Unmarshal)
}

func TestCreateFilesParams(t *testing.T) {
	t.Parallel()

	testCreateFilesParams(t, json.Marshal, json.Unmarshal)
}

func TestRenameFilesParams(t *testing.T) {
	t.Parallel()

	testRenameFilesParams(t, json.Marshal, json.Unmarshal)
}

func TestDeleteFilesParams(t *testing.T) {
	t.Parallel()

	testDeleteFilesParams(t, json.Marshal, json.Unmarshal)
}

func TestReferencesParams(t *testing.T) {
	t.Parallel()

	testReferencesParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlightOptions(t *testing.T) {
	t.Parallel()

	testDocumentHighlightOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlightParams(t *testing.T) {
	t.Parallel()

	testDocumentHighlightParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentSymbolOptions(t *testing.T) {
	t.Parallel()

	testDocumentSymbolOptions(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceSymbolOptions(t *testing.T) {
	t.Parallel()

	testWorkspaceSymbolOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentFormattingOptions(t *testing.T) {
	t.Parallel()

	testDocumentFormattingOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentRangeFormattingOptions(t *testing.T) {
	t.Parallel()

	testDocumentRangeFormattingOptions(t, json.Marshal, json.Unmarshal)
}

func TestDeclarationOptions(t *testing.T) {
	t.Parallel()

	testDeclarationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDeclarationRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDeclarationRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDeclarationParams(t *testing.T) {
	t.Parallel()

	testDeclarationParams(t, json.Marshal, json.Unmarshal)
}

func TestDefinitionOptions(t *testing.T) {
	t.Parallel()

	testDefinitionOptions(t, json.Marshal, json.Unmarshal)
}

func TestDefinitionParams(t *testing.T) {
	t.Parallel()

	testDefinitionParams(t, json.Marshal, json.Unmarshal)
}

func TestTypeDefinitionOptions(t *testing.T) {
	t.Parallel()

	testTypeDefinitionOptions(t, json.Marshal, json.Unmarshal)
}

func TestTypeDefinitionRegistrationOptions(t *testing.T) {
	t.Parallel()

	testTypeDefinitionRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestTypeDefinitionParams(t *testing.T) {
	t.Parallel()

	testTypeDefinitionParams(t, json.Marshal, json.Unmarshal)
}

func TestImplementationOptions(t *testing.T) {
	t.Parallel()

	testImplementationOptions(t, json.Marshal, json.Unmarshal)
}

func TestImplementationRegistrationOptions(t *testing.T) {
	t.Parallel()

	testImplementationRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestImplementationParams(t *testing.T) {
	t.Parallel()

	testImplementationParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentColorOptions(t *testing.T) {
	t.Parallel()

	testDocumentColorOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentColorRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDocumentColorRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeOptions(t *testing.T) {
	t.Parallel()

	testFoldingRangeOptions(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeRegistrationOptions(t *testing.T) {
	t.Parallel()

	testFoldingRangeRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestInitializeResult(t *testing.T) {
	t.Parallel()

	testInitializeResult(t, json.Marshal, json.Unmarshal)
}

func TestInitializeError(t *testing.T) {
	t.Parallel()

	testInitializeError(t, json.Marshal, json.Unmarshal)
}

func TestShowDocumentParams(t *testing.T) {
	t.Parallel()

	testShowDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestShowDocumentResult(t *testing.T) {
	t.Parallel()

	testShowDocumentResult(t, json.Marshal, json.Unmarshal)
}

func TestReferencesOptions(t *testing.T) {
	t.Parallel()

	testReferencesOptions(t, json.Marshal, json.Unmarshal)
}

func TestCodeActionOptions(t *testing.T) {
	t.Parallel()

	testCodeActionOptions(t, json.Marshal, json.Unmarshal)
}

func TestRenameOptions(t *testing.T) {
	t.Parallel()

	testRenameOptions(t, json.Marshal, json.Unmarshal)
}

func TestSaveOptions(t *testing.T) {
	t.Parallel()

	testSaveOptions(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentSyncOptions(t *testing.T) {
	t.Parallel()

	testTextDocumentSyncOptions(t, json.Marshal, json.Unmarshal)
}

func TestHoverOptions(t *testing.T) {
	t.Parallel()

	testHoverOptions(t, json.Marshal, json.Unmarshal)
}

func TestStaticRegistrationOptions(t *testing.T) {
	t.Parallel()

	testStaticRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentLinkRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDocumentLinkRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestInitializedParams(t *testing.T) {
	t.Parallel()

	testInitializedParams(t, json.Marshal, json.Unmarshal)
}
