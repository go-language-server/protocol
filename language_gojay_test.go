// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestCompletionParams(t *testing.T) {
	testCompletionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCompletionContext(t *testing.T) {
	testCompletionContext(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCompletionList(t *testing.T) {
	testCompletionList(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCompletionItem(t *testing.T) { testCompletionItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestCompletionRegistrationOptions(t *testing.T) {
	testCompletionRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestHover(t *testing.T) { testHover(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestSignatureHelp(t *testing.T) { testSignatureHelp(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestSignatureInformation(t *testing.T) {
	testSignatureInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestParameterInformation(t *testing.T) {
	testParameterInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSignatureHelpRegistrationOptions(t *testing.T) {
	testSignatureHelpRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestReferenceParams(t *testing.T) { testReferenceParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestReferenceContext(t *testing.T) {
	testReferenceContext(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentHighlight(t *testing.T) {
	testDocumentHighlight(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSymbolParams(t *testing.T) {
	testDocumentSymbolParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentSymbol(t *testing.T) { testDocumentSymbol(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestCodeActionParams(t *testing.T) {
	testCodeActionParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeActionContext(t *testing.T) {
	testCodeActionContext(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeAction(t *testing.T) { testCodeAction(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestCodeActionRegistrationOptions(t *testing.T) {
	testCodeActionRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCodeLensParams(t *testing.T) { testCodeLensParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestCodeLens(t *testing.T) { testCodeLens(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestCodeLensRegistrationOptions(t *testing.T) {
	testCodeLensRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentLinkParams(t *testing.T) {
	testDocumentLinkParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentLink(t *testing.T) { testDocumentLink(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestDocumentColorParams(t *testing.T) {
	testDocumentColorParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestColorInformation(t *testing.T) {
	testColorInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestColor(t *testing.T) { testColor(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestColorPresentationParams(t *testing.T) {
	testColorPresentationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestColorPresentation(t *testing.T) {
	testColorPresentation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentFormattingParams(t *testing.T) {
	testDocumentFormattingParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestSymbolInformation(t *testing.T) {
	testSymbolInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFormattingOptions(t *testing.T) {
	testFormattingOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentRangeFormattingParams(t *testing.T) {
	testDocumentRangeFormattingParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentOnTypeFormattingParams(t *testing.T) {
	testDocumentOnTypeFormattingParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDocumentOnTypeFormattingRegistrationOptions(t *testing.T) {
	testDocumentOnTypeFormattingRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameParams(t *testing.T) {
	testRenameParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRenameRegistrationOptions(t *testing.T) {
	testRenameRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRangeParams(t *testing.T) {
	testFoldingRangeParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFoldingRange(t *testing.T) {
	testFoldingRange(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
