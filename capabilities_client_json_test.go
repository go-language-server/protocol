// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"testing"

	"github.com/segmentio/encoding/json"
)

func TestClientCapabilities(t *testing.T) {
	t.Parallel()

	testClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceClientCapabilities(t *testing.T) {
	t.Parallel()

	testWorkspaceClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceClientCapabilitiesWorkspaceEdit(t *testing.T) {
	t.Parallel()

	testWorkspaceClientCapabilitiesWorkspaceEdit(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentSyncClientCapabilities(t *testing.T) {
	t.Parallel()

	testTextDocumentSyncClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestCompletionTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testCompletionTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestHoverTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testHoverTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestSignatureHelpTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testSignatureHelpTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestReferencesTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testReferencesTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlightClientCapabilities(t *testing.T) {
	t.Parallel()

	testDocumentHighlightClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDocumentSymbolClientCapabilities(t *testing.T) {
	t.Parallel()

	testDocumentSymbolClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDeclarationTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testDeclarationTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDefinitionTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testDefinitionTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestTypeDefinitionTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testTypeDefinitionTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestImplementationTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testImplementationTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestCodeActionClientCapabilities(t *testing.T) {
	t.Parallel()

	testCodeActionClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestCodeLensClientCapabilities(t *testing.T) {
	t.Parallel()

	testCodeLensClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDocumentLinkClientCapabilities(t *testing.T) {
	t.Parallel()

	testDocumentLinkClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestDocumentColorClientCapabilities(t *testing.T) {
	t.Parallel()

	testDocumentColorClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestRenameClientCapabilities(t *testing.T) {
	t.Parallel()

	testRenameClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestPublishDiagnosticsClientCapabilities(t *testing.T) {
	t.Parallel()

	testPublishDiagnosticsClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeClientCapabilities(t *testing.T) {
	t.Parallel()

	testFoldingRangeClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestSemanticTokensClientCapabilities(t *testing.T) {
	t.Parallel()

	testSemanticTokensClientCapabilities(t, json.Marshal, json.Unmarshal)
}
