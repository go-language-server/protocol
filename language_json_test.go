// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestCompletionParams(t *testing.T) {
	testCompletionParams(t, json.Marshal, json.Unmarshal)
}

func TestCompletionContext(t *testing.T) {
	testCompletionContext(t, json.Marshal, json.Unmarshal)
}

func TestCompletionList(t *testing.T) {
	testCompletionList(t, json.Marshal, json.Unmarshal)
}

func TestCompletionItem(t *testing.T) { testCompletionItem(t, json.Marshal, json.Unmarshal) }

func TestCompletionRegistrationOptions(t *testing.T) {
	testCompletionRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestHoverParams(t *testing.T) {
	testHoverParams(t, json.Marshal, json.Unmarshal)
}

func TestHover(t *testing.T) { testHover(t, json.Marshal, json.Unmarshal) }

func TestSignatureHelpParams(t *testing.T) {
	testSignatureHelpParams(t, json.Marshal, json.Unmarshal)
}

func TestSignatureHelp(t *testing.T) { testSignatureHelp(t, json.Marshal, json.Unmarshal) }

func TestSignatureInformation(t *testing.T) {
	testSignatureInformation(t, json.Marshal, json.Unmarshal)
}

func TestParameterInformation(t *testing.T) {
	testParameterInformation(t, json.Marshal, json.Unmarshal)
}

func TestSignatureHelpRegistrationOptions(t *testing.T) {
	testSignatureHelpRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestReferenceParams(t *testing.T) { testReferenceParams(t, json.Marshal, json.Unmarshal) }

func TestReferenceContext(t *testing.T) {
	testReferenceContext(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlight(t *testing.T) { testDocumentHighlight(t, json.Marshal, json.Unmarshal) }

func TestDocumentSymbolParams(t *testing.T) {
	testDocumentSymbolParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentSymbol(t *testing.T) { testDocumentSymbol(t, json.Marshal, json.Unmarshal) }

func TestCodeActionParams(t *testing.T) { testCodeActionParams(t, json.Marshal, json.Unmarshal) }

func TestCodeActionContext(t *testing.T) { testCodeActionContext(t, json.Marshal, json.Unmarshal) }

func TestCodeAction(t *testing.T) { testCodeAction(t, json.Marshal, json.Unmarshal) }

func TestCodeActionRegistrationOptions(t *testing.T) {
	testCodeActionRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestCodeLensParams(t *testing.T) { testCodeLensParams(t, json.Marshal, json.Unmarshal) }

func TestCodeLens(t *testing.T) { testCodeLens(t, json.Marshal, json.Unmarshal) }

func TestCodeLensRegistrationOptions(t *testing.T) {
	testCodeLensRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentLinkParams(t *testing.T) {
	testDocumentLinkParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentLink(t *testing.T) { testDocumentLink(t, json.Marshal, json.Unmarshal) }

func TestDocumentColorParams(t *testing.T) {
	testDocumentColorParams(t, json.Marshal, json.Unmarshal)
}

func TestColorInformation(t *testing.T) {
	testColorInformation(t, json.Marshal, json.Unmarshal)
}

func TestColor(t *testing.T) { testColor(t, json.Marshal, json.Unmarshal) }

func TestColorPresentationParams(t *testing.T) {
	testColorPresentationParams(t, json.Marshal, json.Unmarshal)
}

func TestColorPresentation(t *testing.T) {
	testColorPresentation(t, json.Marshal, json.Unmarshal)
}

func TestDocumentFormattingParams(t *testing.T) {
	testDocumentFormattingParams(t, json.Marshal, json.Unmarshal)
}

func TestSymbolInformation(t *testing.T) {
	testSymbolInformation(t, json.Marshal, json.Unmarshal)
}

func TestFormattingOptions(t *testing.T) {
	testFormattingOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentRangeFormattingParams(t *testing.T) {
	testDocumentRangeFormattingParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentOnTypeFormattingParams(t *testing.T) {
	testDocumentOnTypeFormattingParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentOnTypeFormattingRegistrationOptions(t *testing.T) {
	testDocumentOnTypeFormattingRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestRenameParams(t *testing.T) {
	testRenameParams(t, json.Marshal, json.Unmarshal)
}

func TestRenameRegistrationOptions(t *testing.T) {
	testRenameRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestPrepareRenameParams(t *testing.T) {
	testPrepareRenameParams(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeParams(t *testing.T) {
	testFoldingRangeParams(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRange(t *testing.T) {
	testFoldingRange(t, json.Marshal, json.Unmarshal)
}
