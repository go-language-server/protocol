// SPDX-FileCopyrightText: 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestCompletionParams(t *testing.T) {
	t.Parallel()

	testCompletionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCompletionContext(t *testing.T) {
	t.Parallel()

	testCompletionContext(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCompletionList(t *testing.T) {
	t.Parallel()

	testCompletionList(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestInsertReplaceEdit(t *testing.T) {
	t.Parallel()

	testInsertReplaceEdit(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCompletionItem(t *testing.T) {
	t.Parallel()

	testCompletionItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCompletionRegistrationOptions(t *testing.T) {
	t.Parallel()

	testCompletionRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestHoverParams(t *testing.T) {
	t.Parallel()

	testHoverParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestHover(t *testing.T) {
	t.Parallel()

	testHover(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSignatureHelpParams(t *testing.T) {
	t.Parallel()

	testSignatureHelpParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSignatureHelp(t *testing.T) {
	t.Parallel()

	testSignatureHelp(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSignatureInformation(t *testing.T) {
	t.Parallel()

	testSignatureInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestParameterInformation(t *testing.T) {
	t.Parallel()

	testParameterInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSignatureHelpRegistrationOptions(t *testing.T) {
	t.Parallel()

	testSignatureHelpRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferenceParams(t *testing.T) {
	t.Parallel()

	testReferenceParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferenceContext(t *testing.T) {
	t.Parallel()

	testReferenceContext(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlight(t *testing.T) {
	t.Parallel()

	testDocumentHighlight(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSymbolParams(t *testing.T) {
	t.Parallel()

	testDocumentSymbolParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSymbol(t *testing.T) {
	t.Parallel()

	testDocumentSymbol(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeActionParams(t *testing.T) {
	t.Parallel()

	testCodeActionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeActionContext(t *testing.T) {
	t.Parallel()

	testCodeActionContext(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeAction(t *testing.T) {
	t.Parallel()

	testCodeAction(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeActionRegistrationOptions(t *testing.T) {
	t.Parallel()

	testCodeActionRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeLensParams(t *testing.T) {
	t.Parallel()

	testCodeLensParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeLens(t *testing.T) {
	t.Parallel()

	testCodeLens(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeLensRegistrationOptions(t *testing.T) {
	t.Parallel()

	testCodeLensRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentLinkParams(t *testing.T) {
	t.Parallel()

	testDocumentLinkParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentLink(t *testing.T) {
	t.Parallel()

	testDocumentLink(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentColorParams(t *testing.T) {
	t.Parallel()

	testDocumentColorParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestColorInformation(t *testing.T) {
	t.Parallel()

	testColorInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestColor(t *testing.T) {
	t.Parallel()

	testColor(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestColorPresentationParams(t *testing.T) {
	t.Parallel()

	testColorPresentationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestColorPresentation(t *testing.T) {
	t.Parallel()

	testColorPresentation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentFormattingParams(t *testing.T) {
	t.Parallel()

	testDocumentFormattingParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSymbolInformation(t *testing.T) {
	t.Parallel()

	testSymbolInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFormattingOptions(t *testing.T) {
	t.Parallel()

	testFormattingOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentRangeFormattingParams(t *testing.T) {
	t.Parallel()

	testDocumentRangeFormattingParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentOnTypeFormattingParams(t *testing.T) {
	t.Parallel()

	testDocumentOnTypeFormattingParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentOnTypeFormattingRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDocumentOnTypeFormattingRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameParams(t *testing.T) {
	t.Parallel()

	testRenameParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameRegistrationOptions(t *testing.T) {
	t.Parallel()

	testRenameRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestPrepareRenameParams(t *testing.T) {
	t.Parallel()

	testPrepareRenameParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeParams(t *testing.T) {
	t.Parallel()

	testFoldingRangeParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRange(t *testing.T) {
	t.Parallel()

	testFoldingRange(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
