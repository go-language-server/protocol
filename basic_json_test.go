// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestPosition(t *testing.T) { testPosition(t, json.Marshal, json.Unmarshal) }

func TestRange(t *testing.T) { testRange(t, json.Marshal, json.Unmarshal) }

func TestLocation(t *testing.T) { testLocation(t, json.Marshal, json.Unmarshal) }

func TestLocationLink(t *testing.T) { testLocationLink(t, json.Marshal, json.Unmarshal) }

func TestDiagnostic(t *testing.T) { testDiagnostic(t, json.Marshal, json.Unmarshal) }

func TestDiagnosticRelatedInformation(t *testing.T) {
	testDiagnosticRelatedInformation(t, json.Marshal, json.Unmarshal)
}

func TestCommand(t *testing.T) { testCommand(t, json.Marshal, json.Unmarshal) }

func TestTextEdit(t *testing.T) { testTextEdit(t, json.Marshal, json.Unmarshal) }

func TestTextDocumentEdit(t *testing.T) { testTextDocumentEdit(t, json.Marshal, json.Unmarshal) }

func TestCreateFileOptions(t *testing.T) { testCreateFileOptions(t, json.Marshal, json.Unmarshal) }

func TestCreateFile(t *testing.T) { testCreateFile(t, json.Marshal, json.Unmarshal) }

func TestRenameFileOptions(t *testing.T) { testRenameFileOptions(t, json.Marshal, json.Unmarshal) }

func TestRenameFile(t *testing.T) { testRenameFile(t, json.Marshal, json.Unmarshal) }

func TestDeleteFileOptions(t *testing.T) { testDeleteFileOptions(t, json.Marshal, json.Unmarshal) }

func TestDeleteFile(t *testing.T) { testDeleteFile(t, json.Marshal, json.Unmarshal) }

func TestWorkspaceEdit(t *testing.T) { testWorkspaceEdit(t, json.Marshal, json.Unmarshal) }

func TestTextDocumentIdentifier(t *testing.T) {
	testTextDocumentIdentifier(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentItem(t *testing.T) { testTextDocumentItem(t, json.Marshal, json.Unmarshal) }

func TestVersionedTextDocumentIdentifier(t *testing.T) {
	testVersionedTextDocumentIdentifier(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentPositionParams(t *testing.T) {
	testTextDocumentPositionParams(t, json.Marshal, json.Unmarshal)
}

func TestDocumentFilter(t *testing.T) { testDocumentFilter(t, json.Marshal, json.Unmarshal) }

func TestDocumentSelector(t *testing.T) { testDocumentSelector(t, json.Marshal, json.Unmarshal) }

func TestMarkupContent(t *testing.T) { testMarkupContent(t, json.Marshal, json.Unmarshal) }
