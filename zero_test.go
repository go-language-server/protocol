// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "testing"

type customZero struct {
	val int
}

func (c customZero) IsZero() bool {
	return c.val == 42
}

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
		{
			name: "custom Zeroer type returning true",
			v:    customZero{val: 42},
			want: true,
		},
		{
			name: "custom Zeroer type returning false",
			v:    customZero{val: 12},
			want: false,
		},
		{
			name: "Optional with value but not set",
			v:    Optional[string]{value: "hello", set: false},
			want: true,
		},
		{
			name: "Optional set with zero value",
			v:    Optional[string]{value: "", set: true},
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
