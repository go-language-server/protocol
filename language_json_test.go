// SPDX-FileCopyrightText: 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"testing"

	"github.com/segmentio/encoding/json"
)

func TestCompletionParams(t *testing.T) {
	t.Parallel()

	testCompletionParams(t, json.Marshal, json.Unmarshal)
}

func TestCompletionContext(t *testing.T) {
	t.Parallel()

	testCompletionContext(t, json.Marshal, json.Unmarshal)
}

func TestCompletionList(t *testing.T) {
	t.Parallel()

	testCompletionList(t, json.Marshal, json.Unmarshal)
}

func TestInsertReplaceEdit(t *testing.T) {
	t.Parallel()

	testInsertReplaceEdit(t, json.Marshal, json.Unmarshal)
}

func TestCompletionItem(t *testing.T) {
	t.Parallel()

	testCompletionItem(t, json.Marshal, json.Unmarshal)
}

func TestCompletionRegistrationOptions(t *testing.T) {
	t.Parallel()

	testCompletionRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestHoverParams(t *testing.T) {
	t.Parallel()

	testHoverParams(t, json.Marshal, json.Unmarshal)
}

func TestHover(t *testing.T) {
	t.Parallel()

	testHover(t, json.Marshal, json.Unmarshal)
}

func TestSignatureHelpParams(t *testing.T) {
	t.Parallel()

	testSignatureHelpParams(t, json.Marshal, json.Unmarshal)
}

func TestSignatureHelp(t *testing.T) {
	t.Parallel()

	testSignatureHelp(t, json.Marshal, json.Unmarshal)
}

func TestSignatureInformation(t *testing.T) {
	t.Parallel()

	testSignatureInformation(t, json.Marshal, json.Unmarshal)
}

func TestParameterInformation(t *testing.T) {
	t.Parallel()

	testParameterInformation(t, json.Marshal, json.Unmarshal)
}

func TestSignatureHelpRegistrationOptions(t *testing.T) {
	t.Parallel()

	testSignatureHelpRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestReferenceParams(t *testing.T) {
	t.Parallel()

	testReferenceParams(t, json.Marshal, json.Unmarshal)
}

func TestReferenceContext(t *testing.T) {
	t.Parallel()

	testReferenceContext(t, json.Marshal, json.Unmarshal)
}

func TestDocumentHighlight(t *testing.T) {
	t.Parallel()

	testDocumentHighlight(t, json.Marshal, json.Unmarshal)
}

func TestDocumentSymbolParams(t *testing.T) {
	t.Parallel()

	testDocumentSymbolParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentSymbol(t *testing.T) {
	t.Parallel()

	testDocumentSymbol(t, json.Marshal, json.Unmarshal)
}

func TestCodeActionParams(t *testing.T) {
	t.Parallel()

	testCodeActionParams(t, json.Marshal, json.Unmarshal)
}

func TestCodeActionContext(t *testing.T) {
	t.Parallel()

	testCodeActionContext(t, json.Marshal, json.Unmarshal)
}

func TestCodeAction(t *testing.T) {
	t.Parallel()

	testCodeAction(t, json.Marshal, json.Unmarshal)
}

func TestCodeActionRegistrationOptions(t *testing.T) {
	t.Parallel()

	testCodeActionRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestCodeLensParams(t *testing.T) {
	t.Parallel()

	testCodeLensParams(t, json.Marshal, json.Unmarshal)
}

func TestCodeLens(t *testing.T) {
	t.Parallel()

	testCodeLens(t, json.Marshal, json.Unmarshal)
}

func TestCodeLensRegistrationOptions(t *testing.T) {
	t.Parallel()

	testCodeLensRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentLinkParams(t *testing.T) {
	t.Parallel()

	testDocumentLinkParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentLink(t *testing.T) {
	t.Parallel()

	testDocumentLink(t, json.Marshal, json.Unmarshal)
}

func TestDocumentColorParams(t *testing.T) {
	t.Parallel()

	testDocumentColorParams(t, json.Marshal, json.Unmarshal)
}

func TestColorInformation(t *testing.T) {
	t.Parallel()

	testColorInformation(t, json.Marshal, json.Unmarshal)
}

func TestColor(t *testing.T) {
	t.Parallel()

	testColor(t, json.Marshal, json.Unmarshal)
}

func TestColorPresentationParams(t *testing.T) {
	t.Parallel()

	testColorPresentationParams(t, json.Marshal, json.Unmarshal)
}

func TestColorPresentation(t *testing.T) {
	t.Parallel()

	testColorPresentation(t, json.Marshal, json.Unmarshal)
}

func TestDocumentFormattingParams(t *testing.T) {
	t.Parallel()

	testDocumentFormattingParams(t, json.Marshal, json.Unmarshal)
}

func TestSymbolInformation(t *testing.T) {
	t.Parallel()

	testSymbolInformation(t, json.Marshal, json.Unmarshal)
}

func TestFormattingOptions(t *testing.T) {
	t.Parallel()

	testFormattingOptions(t, json.Marshal, json.Unmarshal)
}

func TestDocumentRangeFormattingParams(t *testing.T) {
	t.Parallel()

	testDocumentRangeFormattingParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentOnTypeFormattingParams(t *testing.T) {
	t.Parallel()

	testDocumentOnTypeFormattingParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentOnTypeFormattingRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDocumentOnTypeFormattingRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestRenameParams(t *testing.T) {
	t.Parallel()

	testRenameParams(t, json.Marshal, json.Unmarshal)
}

func TestRenameRegistrationOptions(t *testing.T) {
	t.Parallel()

	testRenameRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestPrepareRenameParams(t *testing.T) {
	t.Parallel()

	testPrepareRenameParams(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRangeParams(t *testing.T) {
	t.Parallel()

	testFoldingRangeParams(t, json.Marshal, json.Unmarshal)
}

func TestFoldingRange(t *testing.T) {
	t.Parallel()

	testFoldingRange(t, json.Marshal, json.Unmarshal)
}
