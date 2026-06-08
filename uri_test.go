// Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"path/filepath"
	"testing"
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
		want URI
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
		want    URI
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

// TestDocumentURIAlias verifies DocumentURI is interchangeable with URI so that a
// single set of constructors and methods serves both.
func TestDocumentURIAlias(t *testing.T) {
	t.Parallel()

	want := filepath.FromSlash("/tmp/x.go")
	// File returns a URI; the struct field is typed DocumentURI. The assignment
	// compiling at all is the alias proof; Filename then works through the field.
	item := TextDocumentItem{URI: File("/tmp/x.go")}
	if got := item.URI.Filename(); got != want {
		t.Errorf("DocumentURI.Filename() = %q, want %q", got, want)
	}
}
