// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"testing"

	gocmp "github.com/google/go-cmp/cmp"
	"go.lsp.dev/uri"
)

func TestGeneratedURIFieldsParseWithExternalURI(t *testing.T) {
	t.Parallel()

	data := []byte(`{"uri":"file:///tmp/my file.go","languageId":"go","version":1,"text":"package main\n"}`)
	var got TextDocumentItem
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal(TextDocumentItem): %v", err)
	}

	want := uri.MustParse("file:///tmp/my file.go")
	if diff := gocmp.Diff(want, got.URI); diff != "" {
		t.Fatalf("TextDocumentItem.URI mismatch (-want +got):\n%s", diff)
	}
	if got.URI.String() != "file:///tmp/my%20file.go" {
		t.Fatalf("TextDocumentItem.URI.String() = %q, want canonical escaped URI", got.URI.String())
	}
}

func TestGeneratedURIMapKeysParseWithExternalURI(t *testing.T) {
	t.Parallel()

	data := []byte(`{"changes":{"file:///tmp/my file.go":[{"range":{"start":{"line":0,"character":1},"end":{"line":0,"character":2}},"newText":"x"}]}}`)
	var got WorkspaceEdit
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal(WorkspaceEdit): %v", err)
	}

	wantKey := uri.MustParse("file:///tmp/my file.go")
	edits, ok := got.Changes[wantKey]
	if !ok {
		t.Fatalf("WorkspaceEdit.Changes missing parsed key %q: %#v", wantKey, got.Changes)
	}
	if diff := gocmp.Diff([]TextEdit{{
		Range: Range{
			Start: Position{Line: 0, Character: 1},
			End:   Position{Line: 0, Character: 2},
		},
		NewText: "x",
	}}, edits); diff != "" {
		t.Fatalf("WorkspaceEdit.Changes[%q] mismatch (-want +got):\n%s", wantKey, diff)
	}
}

func TestRelativePatternBaseURIParsesWithExternalURI(t *testing.T) {
	t.Parallel()

	data := []byte(`{"baseUri":"file:///tmp/my file.go","pattern":"**/*.go"}`)
	var got RelativePattern
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal(RelativePattern): %v", err)
	}

	wantURI := URI(uri.MustParse("file:///tmp/my file.go"))
	if diff := gocmp.Diff(wantURI, got.BaseURI); diff != "" {
		t.Fatalf("RelativePattern.BaseURI mismatch (-want +got):\n%s", diff)
	}
}
