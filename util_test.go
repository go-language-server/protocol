// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"go.lsp.dev/uri"
)

func TestToURI(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want uri.URI
	}{
		{
			name: "Valid",
			s:    "/path/to/test.go",
			want: uri.URI("file:///path/to/test.go"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, ToURI(tt.s)); diff != "" {
				t.Errorf("(+want, -got)\n%s", diff)
			}
		})
	}
}

func TestNewVersion(t *testing.T) {
	tests := []struct {
		name string
		i    uint64
	}{
		{
			name: "Valid",
			i:    5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := NewVersion(tt.i)
			if got := NewVersion(tt.i); *got != *want {
				t.Errorf("NewVersion(%v) = %v, want %v", tt.i, *got, *want)
			}
		})
	}
}
