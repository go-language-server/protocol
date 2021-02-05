// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

func TestWorkspaceClientCapabilities(t *testing.T) {
	testWorkspaceClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesSynchronization(t *testing.T) {
	testTextDocumentClientCapabilitiesSynchronization(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesCompletion(t *testing.T) {
	testTextDocumentClientCapabilitiesCompletion(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesHover(t *testing.T) {
	testTextDocumentClientCapabilitiesHover(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesSignatureHelp(t *testing.T) {
	testTextDocumentClientCapabilitiesSignatureHelp(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferencesParams(t *testing.T) {
	testReferencesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesReferences(t *testing.T) {
	testTextDocumentClientCapabilitiesReferences(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlightOptions(t *testing.T) {
	testDocumentHighlightOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlightParams(t *testing.T) {
	testDocumentHighlightParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDocumentHighlight(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentHighlight(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSymbolOptions(t *testing.T) {
	testDocumentSymbolOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDocumentSymbol(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentSymbol(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceSymbolOptions(t *testing.T) {
	testWorkspaceSymbolOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentFormattingOptions(t *testing.T) {
	testDocumentFormattingOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesFormatting(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentRangeFormattingOptions(t *testing.T) {
	testDocumentRangeFormattingOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesRangeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesRangeFormatting(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesOnTypeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesOnTypeFormatting(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
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

func TestTextDocumentClientCapabilitiesDeclaration(t *testing.T) {
	testTextDocumentClientCapabilitiesDeclaration(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDefinitionOptions(t *testing.T) {
	testDefinitionOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDefinitionParams(t *testing.T) {
	testDefinitionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDefinition(t *testing.T) {
	testTextDocumentClientCapabilitiesDefinition(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
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

func TestTextDocumentClientCapabilitiesTypeDefinition(t *testing.T) {
	testTextDocumentClientCapabilitiesTypeDefinition(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestImplementationParams(t *testing.T) {
	testImplementationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesImplementation(t *testing.T) {
	testTextDocumentClientCapabilitiesImplementation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesCodeAction(t *testing.T) {
	testTextDocumentClientCapabilitiesCodeAction(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesCodeLens(t *testing.T) {
	testTextDocumentClientCapabilitiesCodeLens(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDocumentLink(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentLink(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentColorOptions(t *testing.T) {
	testDocumentColorOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentColorRegistrationOptions(t *testing.T) {
	testDocumentColorRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesColorProvider(t *testing.T) {
	testTextDocumentClientCapabilitiesColorProvider(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesRename(t *testing.T) {
	testTextDocumentClientCapabilitiesRename(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesPublishDiagnostics(t *testing.T) {
	testTextDocumentClientCapabilitiesPublishDiagnostics(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeOptions(t *testing.T) {
	testFoldingRangeOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeRegistrationOptions(t *testing.T) {
	testFoldingRangeRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesFoldingRange(t *testing.T) {
	testTextDocumentClientCapabilitiesFoldingRange(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilities(t *testing.T) {
	testTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestClientCapabilities(t *testing.T) {
	testClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializeResult(t *testing.T) {
	testInitializeResult(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentLinkRegistrationOptions(t *testing.T) {
	testDocumentLinkRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInitializedParams(t *testing.T) {
	testInitializedParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
