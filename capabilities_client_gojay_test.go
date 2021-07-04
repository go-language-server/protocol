// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestClientCapabilities(t *testing.T) {
	t.Parallel()

	testClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceClientCapabilities(t *testing.T) {
	t.Parallel()

	testWorkspaceClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceClientCapabilitiesWorkspaceEdit(t *testing.T) {
	t.Parallel()

	testWorkspaceClientCapabilitiesWorkspaceEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentSyncClientCapabilities(t *testing.T) {
	t.Parallel()

	testTextDocumentSyncClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCompletionTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testCompletionTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestHoverTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testHoverTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSignatureHelpTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testSignatureHelpTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferencesTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testReferencesTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlightClientCapabilities(t *testing.T) {
	t.Parallel()

	testDocumentHighlightClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSymbolClientCapabilities(t *testing.T) {
	t.Parallel()

	testDocumentSymbolClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeclarationTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testDeclarationTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDefinitionTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testDefinitionTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTypeDefinitionTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testTypeDefinitionTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestImplementationTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testImplementationTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeActionClientCapabilities(t *testing.T) {
	t.Parallel()

	testCodeActionClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeLensClientCapabilities(t *testing.T) {
	t.Parallel()

	testCodeLensClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentLinkClientCapabilities(t *testing.T) {
	t.Parallel()

	testDocumentLinkClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentColorClientCapabilities(t *testing.T) {
	t.Parallel()

	testDocumentColorClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameClientCapabilities(t *testing.T) {
	t.Parallel()

	testRenameClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestPublishDiagnosticsClientCapabilities(t *testing.T) {
	t.Parallel()

	testPublishDiagnosticsClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeClientCapabilities(t *testing.T) {
	t.Parallel()

	testFoldingRangeClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilities(t *testing.T) {
	t.Parallel()

	testTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSemanticTokensClientCapabilities(t *testing.T) {
	t.Parallel()

	testSemanticTokensClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
