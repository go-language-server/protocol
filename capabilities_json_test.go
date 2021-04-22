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

func TestTextDocumentClientCapabilitiesReferences(t *testing.T) {
	testTextDocumentClientCapabilitiesReferences(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDocumentHighlight(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentHighlight(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDocumentSymbol(t *testing.T) {
	testTextDocumentClientCapabilitiesDocumentSymbol(t, json.Marshal, json.Unmarshal)
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

func TestTextDocumentClientCapabilitiesDeclaration(t *testing.T) {
	testTextDocumentClientCapabilitiesDeclaration(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesDefinition(t *testing.T) {
	testTextDocumentClientCapabilitiesDefinition(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesTypeDefinition(t *testing.T) {
	testTextDocumentClientCapabilitiesTypeDefinition(t, json.Marshal, json.Unmarshal)
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

func TestTextDocumentClientCapabilitiesColorProvider(t *testing.T) {
	testTextDocumentClientCapabilitiesColorProvider(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesRename(t *testing.T) {
	testTextDocumentClientCapabilitiesRename(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesPublishDiagnostics(t *testing.T) {
	testTextDocumentClientCapabilitiesPublishDiagnostics(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilitiesFoldingRange(t *testing.T) {
	testTextDocumentClientCapabilitiesFoldingRange(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentClientCapabilities(t *testing.T) {
	testTextDocumentClientCapabilities(t, json.Marshal, json.Unmarshal)
}

func TestClientCapabilities(t *testing.T) {
	testClientCapabilities(t, json.Marshal, json.Unmarshal)
}
