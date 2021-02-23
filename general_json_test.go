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

func TestWorkspaceClientCapabilities(t *testing.T) {
	testWorkspaceClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesSynchronization(t *testing.T) {
	testTextDocumentClientCapabilitiesSynchronization(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesCompletion(t *testing.T) {
	testTextDocumentClientCapabilitiesCompletion(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesHover(t *testing.T) {
	testTextDocumentClientCapabilitiesHover(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesSignatureHelp(t *testing.T) {
	testTextDocumentClientCapabilitiesSignatureHelp(t, json.Marshal, json.Unmarshal)
}

func TestReferencesParams(t *testing.T) {
	testReferencesParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesReferences(t *testing.T) {
	testTextDocumentClientCapabilitiesReferences(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlightOptions(t *testing.T) {
	testDocumentHighlightOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlightParams(t *testing.T) {
	testDocumentHighlightParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDocumentHighlight(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentHighlight(t, json.Marshal, json.Unmarshal)
}

func TestDocumentSymbolOptions(t *testing.T) {
	testDocumentSymbolOptions(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDocumentSymbol(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentSymbol(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceSymbolOptions(t *testing.T) {
	testWorkspaceSymbolOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentFormattingOptions(t *testing.T) {
	testDocumentFormattingOptions(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesFormatting(t, json.Marshal, json.Unmarshal)
}

func TestDocumentRangeFormattingOptions(t *testing.T) {
	testDocumentRangeFormattingOptions(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesRangeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesRangeFormatting(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesOnTypeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesOnTypeFormatting(t, json.Marshal, json.Unmarshal)
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

func TestTextDocumentClientCapabilitiesDeclaration(t *testing.T) {
	testTextDocumentClientCapabilitiesDeclaration(t, json.Marshal, json.Unmarshal)
}

func TestDefinitionOptions(t *testing.T) {
	testDefinitionOptions(t, json.Marshal, json.Unmarshal)
}

func TestDefinitionParams(t *testing.T) {
	testDefinitionParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDefinition(t *testing.T) {
	testTextDocumentClientCapabilitiesDefinition(t, json.Marshal, json.Unmarshal)
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

func TestTextDocumentClientCapabilitiesTypeDefinition(t *testing.T) {
	testTextDocumentClientCapabilitiesTypeDefinition(t, json.Marshal, json.Unmarshal)
}

func TestImplementationParams(t *testing.T) {
	testImplementationParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesImplementation(t *testing.T) {
	testTextDocumentClientCapabilitiesImplementation(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesCodeAction(t *testing.T) {
	testTextDocumentClientCapabilitiesCodeAction(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesCodeLens(t *testing.T) {
	testTextDocumentClientCapabilitiesCodeLens(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDocumentLink(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentLink(t, json.Marshal, json.Unmarshal)
}

func TestDocumentColorOptions(t *testing.T) {
	testDocumentColorOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentColorRegistrationOptions(t *testing.T) {
	testDocumentColorRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesColorProvider(t *testing.T) {
	testTextDocumentClientCapabilitiesColorProvider(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesRename(t *testing.T) {
	testTextDocumentClientCapabilitiesRename(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesPublishDiagnostics(t *testing.T) {
	testTextDocumentClientCapabilitiesPublishDiagnostics(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeOptions(t *testing.T) {
	testFoldingRangeOptions(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeRegistrationOptions(t *testing.T) {
	testFoldingRangeRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesFoldingRange(t *testing.T) {
	testTextDocumentClientCapabilitiesFoldingRange(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilities(t *testing.T) {
	testTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestClientCapabilities(t *testing.T) { testClientCapabilities(t, json.Marshal, json.Unmarshal) }

func TestInitializeResult(t *testing.T) { testInitializeResult(t, json.Marshal, json.Unmarshal) }

func TestDocumentLinkRegistrationOptions(t *testing.T) {
	testDocumentLinkRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestInitializedParams(t *testing.T) { testInitializedParams(t, json.Marshal, json.Unmarshal) }
