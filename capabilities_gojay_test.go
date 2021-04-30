// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestWorkspaceClientCapabilities(t *testing.T) {
	testWorkspaceClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentSyncClientCapabilities(t *testing.T) {
	testTextDocumentSyncClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCompletionTextDocumentClientCapabilities(t *testing.T) {
	testCompletionTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestHoverTextDocumentClientCapabilities(t *testing.T) {
	testHoverTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSignatureHelpTextDocumentClientCapabilities(t *testing.T) {
	testSignatureHelpTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferencesTextDocumentClientCapabilities(t *testing.T) {
	testReferencesTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlightClientCapabilities(t *testing.T) {
	testDocumentHighlightClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSymbolClientCapabilities(t *testing.T) {
	testDocumentSymbolClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesFormatting(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesRangeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesRangeFormatting(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilitiesOnTypeFormatting(t *testing.T) {
	testTextDocumentClientCapabilitiesOnTypeFormatting(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDeclarationTextDocumentClientCapabilities(t *testing.T) {
	testDeclarationTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDefinitionTextDocumentClientCapabilities(t *testing.T) {
	testDefinitionTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTypeDefinitionTextDocumentClientCapabilities(t *testing.T) {
	testTypeDefinitionTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestImplementationTextDocumentClientCapabilities(t *testing.T) {
	testImplementationTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeActionClientCapabilities(t *testing.T) {
	testCodeActionClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeLensClientCapabilities(t *testing.T) {
	testCodeLensClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentLinkClientCapabilities(t *testing.T) {
	testDocumentLinkClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentColorClientCapabilities(t *testing.T) {
	testDocumentColorClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameClientCapabilities(t *testing.T) {
	testRenameClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestPublishDiagnosticsClientCapabilities(t *testing.T) {
	testPublishDiagnosticsClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeClientCapabilities(t *testing.T) {
	testFoldingRangeClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentClientCapabilities(t *testing.T) {
	testTextDocumentClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestClientCapabilities(t *testing.T) {
	testClientCapabilities(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
