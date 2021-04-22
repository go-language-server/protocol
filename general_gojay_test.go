// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestWorkspaceFolders(t *testing.T) {
	testWorkspaceFolders(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestClientInfo(t *testing.T) {
	testClientInfo(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializeParams(t *testing.T) {
	testInitializeParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestLogTraceParams(t *testing.T) {
	testLogTraceParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSetTraceParams(t *testing.T) {
	testSetTraceParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCreateFilesParams(t *testing.T) {
	testCreateFilesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameFilesParams(t *testing.T) {
	testRenameFilesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeleteFilesParams(t *testing.T) {
	testDeleteFilesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferencesParams(t *testing.T) {
	testReferencesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlightOptions(t *testing.T) {
	testDocumentHighlightOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlightParams(t *testing.T) {
	testDocumentHighlightParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSymbolOptions(t *testing.T) {
	testDocumentSymbolOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceSymbolOptions(t *testing.T) {
	testWorkspaceSymbolOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentFormattingOptions(t *testing.T) {
	testDocumentFormattingOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentRangeFormattingOptions(t *testing.T) {
	testDocumentRangeFormattingOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeclarationOptions(t *testing.T) {
	testDeclarationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeclarationRegistrationOptions(t *testing.T) {
	testDeclarationRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeclarationParams(t *testing.T) {
	testDeclarationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDefinitionOptions(t *testing.T) {
	testDefinitionOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDefinitionParams(t *testing.T) {
	testDefinitionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTypeDefinitionOptions(t *testing.T) {
	testTypeDefinitionOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTypeDefinitionRegistrationOptions(t *testing.T) {
	testTypeDefinitionRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTypeDefinitionParams(t *testing.T) {
	testTypeDefinitionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestImplementationOptions(t *testing.T) {
	testImplementationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestImplementationRegistrationOptions(t *testing.T) {
	testImplementationRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestImplementationParams(t *testing.T) {
	testImplementationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentColorOptions(t *testing.T) {
	testDocumentColorOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentColorRegistrationOptions(t *testing.T) {
	testDocumentColorRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeOptions(t *testing.T) {
	testFoldingRangeOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeRegistrationOptions(t *testing.T) {
	testFoldingRangeRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializeResult(t *testing.T) {
	testInitializeResult(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializeError(t *testing.T) {
	testInitializeError(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestShowDocumentParams(t *testing.T) {
	testShowDocumentParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestShowDocumentResult(t *testing.T) {
	testShowDocumentResult(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferencesOptions(t *testing.T) {
	testReferencesOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeActionOptions(t *testing.T) {
	testCodeActionOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameOptions(t *testing.T) {
	testRenameOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSaveOptions(t *testing.T) {
	testSaveOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentSyncOptions(t *testing.T) {
	testTextDocumentSyncOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestHoverOptions(t *testing.T) {
	testHoverOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestStaticRegistrationOptions(t *testing.T) {
	testStaticRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentLinkRegistrationOptions(t *testing.T) {
	testDocumentLinkRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializedParams(t *testing.T) {
	testInitializedParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
