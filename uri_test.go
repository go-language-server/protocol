// Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"path/filepath"
	"testing"

	uripkg "go.lsp.dev/uri"
)

func TestFileFilenameRoundTrip(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		path string
	}{
		"success: absolute path":    {path: "/tmp/example.go"},
		"success: nested path":      {path: "/a/b/c/d.go"},
		"success: path with spaces": {path: "/tmp/my file.go"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			u := File(tt.path)
			if got, want := u.Filename(), filepath.FromSlash(tt.path); got != want {
				t.Errorf("File(%q).Filename() = %q, want %q", tt.path, got, want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in   string
		want uripkg.URI
	}{
		"success: already a file URI": {in: "file:///tmp/x.go", want: "file:///tmp/x.go"},
		"success: plain path":         {in: "/tmp/x.go", want: "file:///tmp/x.go"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := New(tt.in); got != tt.want {
				t.Errorf("New(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in      string
		want    uripkg.URI
		wantErr bool
	}{
		"success: file scheme":  {in: "file:///tmp/x.go", want: "file:///tmp/x.go"},
		"success: https scheme": {in: "https://example.com/a", want: "https://example.com/a"},
		"error: unknown scheme": {in: "ftp://example.com", wantErr: true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := Parse(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Parse(%q) error = %v, wantErr %v", tt.in, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Parse(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

// TestGeneratedURIFieldsUseExternalURI verifies ordinary generated URI and
// URI field contexts accept go.lsp.dev/uri.URI values directly.
func TestGeneratedURIFieldsUseExternalURI(t *testing.T) {
	t.Parallel()

	docURI := uripkg.File("/tmp/x.go")
	item := TextDocumentItem{URI: docURI}
	if got := item.URI; got != docURI {
		t.Fatalf("TextDocumentItem.URI = %q, want %q", got, docURI)
	}
	if got, want := item.URI.Filename(), filepath.FromSlash("/tmp/x.go"); got != want {
		t.Fatalf("TextDocumentItem.URI.Filename() = %q, want %q", got, want)
	}
}

func TestGeneratedURIRoundTripsAsJSONString(t *testing.T) {
	t.Parallel()

	docURI := uripkg.URI("file:///tmp/x.go")
	link := DocumentLink{
		Range: Range{
			Start: Position{Line: 1, Character: 2},
			End:   Position{Line: 3, Character: 4},
		},
		Target: &docURI,
	}
	out, err := Marshal(link)
	if err != nil {
		t.Fatalf("Marshal(DocumentLink): %v", err)
	}
	want := []byte(`{"range":{"start":{"line":1,"character":2},"end":{"line":3,"character":4}},"target":"file:///tmp/x.go"}`)
	if got := canon(t, out); got != canon(t, want) {
		t.Fatalf("Marshal(DocumentLink) = %s, want %s", got, canon(t, want))
	}

	var got DocumentLink
	if err := Unmarshal(out, &got); err != nil {
		t.Fatalf("Unmarshal(DocumentLink): %v", err)
	}
	if got.Target == nil || *got.Target != docURI {
		t.Fatalf("DocumentLink.Target = %v, want %q", got.Target, docURI)
	}
}

func TestGeneratedURIMapKeysRoundTrip(t *testing.T) {
	t.Parallel()

	docURI := uripkg.URI("file:///tmp/x.go")
	edit := WorkspaceEdit{Changes: map[uripkg.URI][]TextEdit{
		docURI: {{
			Range: Range{
				Start: Position{Line: 0, Character: 1},
				End:   Position{Line: 0, Character: 2},
			},
			NewText: "x",
		}},
	}}
	out, err := Marshal(edit)
	if err != nil {
		t.Fatalf("Marshal(WorkspaceEdit): %v", err)
	}

	var got WorkspaceEdit
	if err := Unmarshal(out, &got); err != nil {
		t.Fatalf("Unmarshal(WorkspaceEdit): %v", err)
	}
	edits, ok := got.Changes[docURI]
	if !ok {
		t.Fatalf("WorkspaceEdit.Changes missing key %q after round trip: %#v", docURI, got.Changes)
	}
	if len(edits) != 1 || edits[0].NewText != "x" {
		t.Fatalf("WorkspaceEdit.Changes[%q] = %#v, want one edit with NewText x", docURI, edits)
	}
}

func TestLocalURIBridgeForRelativePatternBaseURI(t *testing.T) {
	t.Parallel()

	docURI := uripkg.URI("file:///tmp/x.go")
	var base RelativePatternBaseURI = URI(docURI)
	out, err := Marshal(base)
	if err != nil {
		t.Fatalf("Marshal(RelativePatternBaseURI URI arm): %v", err)
	}
	if got, want := string(out), `"file:///tmp/x.go"`; got != want {
		t.Fatalf("Marshal(RelativePatternBaseURI URI arm) = %s, want %s", got, want)
	}

	var decoded RelativePatternBaseURI
	if err := Unmarshal(out, &decoded); err != nil {
		t.Fatalf("Unmarshal(RelativePatternBaseURI URI arm): %v", err)
	}
	if got, ok := decoded.(URI); !ok || got != URI(docURI) {
		t.Fatalf("decoded RelativePatternBaseURI = %T %[1]v, want protocol.URI(%q)", decoded, docURI)
	}
}

func TestCompatibilityURIBridgeFilename(t *testing.T) {
	t.Parallel()

	docURI := uripkg.File("/tmp/x.go")
	if got, want := URI(docURI).Filename(), filepath.FromSlash("/tmp/x.go"); got != want {
		t.Fatalf("URI(%q).Filename() = %q, want %q", docURI, got, want)
	}
}
