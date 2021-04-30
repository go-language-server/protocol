// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestWorkspaceClientCapabilities(t *testing.T) {
	testWorkspaceClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentSyncClientCapabilities(t *testing.T) {
	testTextDocumentSyncClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestCompletionTextDocumentClientCapabilities(t *testing.T) {
	testCompletionTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestHoverTextDocumentClientCapabilities(t *testing.T) {
	testHoverTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestSignatureHelpTextDocumentClientCapabilities(t *testing.T) {
	testSignatureHelpTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestReferencesTextDocumentClientCapabilities(t *testing.T) {
	testReferencesTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlightClientCapabilities(t *testing.T) {
	testDocumentHighlightClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDocumentSymbolClientCapabilities(t *testing.T) {
	testDocumentSymbolClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesFormatting(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesRangeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesRangeFormatting(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesOnTypeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesOnTypeFormatting(t, json.Marshal, json.Unmarshal)
}

func TestDeclarationTextDocumentClientCapabilities(t *testing.T) {
	testDeclarationTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDefinitionTextDocumentClientCapabilities(t *testing.T) {
	testDefinitionTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestTypeDefinitionTextDocumentClientCapabilities(t *testing.T) {
	testTypeDefinitionTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestImplementationTextDocumentClientCapabilities(t *testing.T) {
	testImplementationTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestCodeActionClientCapabilities(t *testing.T) {
	testCodeActionClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestCodeLensClientCapabilities(t *testing.T) {
	testCodeLensClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDocumentLinkClientCapabilities(t *testing.T) {
	testDocumentLinkClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDocumentColorClientCapabilities(t *testing.T) {
	testDocumentColorClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestRenameClientCapabilities(t *testing.T) {
	testRenameClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestPublishDiagnosticsClientCapabilities(t *testing.T) {
	testPublishDiagnosticsClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeClientCapabilities(t *testing.T) {
	testFoldingRangeClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilities(t *testing.T) {
	testTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestClientCapabilities(t *testing.T) {
	testClientCapabilities(t, json.Marshal, json.Unmarshal)
}
