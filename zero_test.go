// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "testing"

func TestIsZeroOmitValue(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want bool
	}{
		{
			name: "nil interface",
			v:    nil,
			want: true,
		},
		{
			name: "zero generated struct",
			v:    CodeActionDisabled{},
			want: true,
		},
		{
			name: "non-zero generated struct",
			v:    CodeActionDisabled{Reason: "server policy"},
			want: false,
		},
		{
			name: "zero enum without present-zero semantics",
			v:    InsertTextFormat(0),
			want: true,
		},
		{
			name: "non-zero enum",
			v:    InsertTextFormatPlainText,
			want: false,
		},
		{
			name: "nil raw json value",
			v:    LSPAny(nil),
			want: true,
		},
		{
			name: "non-empty raw json value",
			v:    LSPAny(`{"x":1}`),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroOmitValue(tt.v); got != tt.want {
				t.Fatalf("isZeroOmitValue(%T) = %v, want %v", tt.v, got, tt.want)
			}
		})
	}
}
