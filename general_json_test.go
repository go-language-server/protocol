// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestWorkspaceFolders(t *testing.T) { testWorkspaceFolders(t, json.Marshal, json.Unmarshal) }

func TestClientInfo(t *testing.T) { testClientInfo(t, json.Marshal, json.Unmarshal) }

func TestInitializeParams(t *testing.T) { testInitializeParams(t, json.Marshal, json.Unmarshal) }

func TestLogTraceParams(t *testing.T) {
	testLogTraceParams(t, json.Marshal, json.Unmarshal)
}

func TestSetTraceParams(t *testing.T) {
	testSetTraceParams(t, json.Marshal, json.Unmarshal)
}

func TestCreateFilesParams(t *testing.T) {
	testCreateFilesParams(t, json.Marshal, json.Unmarshal)
}

func TestRenameFilesParams(t *testing.T) {
	testRenameFilesParams(t, json.Marshal, json.Unmarshal)
}

func TestDeleteFilesParams(t *testing.T) {
	testDeleteFilesParams(t, json.Marshal, json.Unmarshal)
}

func TestReferencesParams(t *testing.T) {
	testReferencesParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlightOptions(t *testing.T) {
	testDocumentHighlightOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlightParams(t *testing.T) {
	testDocumentHighlightParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentSymbolOptions(t *testing.T) {
	testDocumentSymbolOptions(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceSymbolOptions(t *testing.T) {
	testWorkspaceSymbolOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentFormattingOptions(t *testing.T) {
	testDocumentFormattingOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentRangeFormattingOptions(t *testing.T) {
	testDocumentRangeFormattingOptions(t, json.Marshal, json.Unmarshal)
}

func TestDeclarationOptions(t *testing.T) {
	testDeclarationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDeclarationRegistrationOptions(t *testing.T) {
	testDeclarationRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDeclarationParams(t *testing.T) {
	testDeclarationParams(t, json.Marshal, json.Unmarshal)
}

func TestDefinitionOptions(t *testing.T) {
	testDefinitionOptions(t, json.Marshal, json.Unmarshal)
}

func TestDefinitionParams(t *testing.T) {
	testDefinitionParams(t, json.Marshal, json.Unmarshal)
}

func TestTypeDefinitionOptions(t *testing.T) {
	testTypeDefinitionOptions(t, json.Marshal, json.Unmarshal)
}

func TestTypeDefinitionRegistrationOptions(t *testing.T) {
	testTypeDefinitionRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestTypeDefinitionParams(t *testing.T) {
	testTypeDefinitionParams(t, json.Marshal, json.Unmarshal)
}

func TestImplementationOptions(t *testing.T) {
	testImplementationOptions(t, json.Marshal, json.Unmarshal)
}

func TestImplementationRegistrationOptions(t *testing.T) {
	testImplementationRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestImplementationParams(t *testing.T) {
	testImplementationParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentColorOptions(t *testing.T) {
	testDocumentColorOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentColorRegistrationOptions(t *testing.T) {
	testDocumentColorRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeOptions(t *testing.T) {
	testFoldingRangeOptions(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeRegistrationOptions(t *testing.T) {
	testFoldingRangeRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestInitializeResult(t *testing.T) { testInitializeResult(t, json.Marshal, json.Unmarshal) }

func TestInitializeError(t *testing.T) {
	testInitializeError(t, json.Marshal, json.Unmarshal)
}

func TestShowDocumentParams(t *testing.T) {
	testShowDocumentParams(t, json.Marshal, json.Unmarshal)
}

func TestShowDocumentResult(t *testing.T) {
	testShowDocumentResult(t, json.Marshal, json.Unmarshal)
}

func TestReferencesOptions(t *testing.T) {
	testReferencesOptions(t, json.Marshal, json.Unmarshal)
}

func TestCodeActionOptions(t *testing.T) {
	testCodeActionOptions(t, json.Marshal, json.Unmarshal)
}

func TestRenameOptions(t *testing.T) {
	testRenameOptions(t, json.Marshal, json.Unmarshal)
}

func TestSaveOptions(t *testing.T) {
	testSaveOptions(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentSyncOptions(t *testing.T) {
	testTextDocumentSyncOptions(t, json.Marshal, json.Unmarshal)
}

func TestHoverOptions(t *testing.T) {
	testHoverOptions(t, json.Marshal, json.Unmarshal)
}

func TestStaticRegistrationOptions(t *testing.T) {
	testStaticRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentLinkRegistrationOptions(t *testing.T) {
	testDocumentLinkRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestInitializedParams(t *testing.T) { testInitializedParams(t, json.Marshal, json.Unmarshal) }
