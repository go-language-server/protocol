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

func TestUint64Ptr(t *testing.T) {
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
			want := Uint64Ptr(tt.i)
			if got := Uint64Ptr(tt.i); *got != *want {
				t.Errorf("Uint64Ptr(%v) = %v, want %v", tt.i, *got, *want)
			}
		})
	}
}
