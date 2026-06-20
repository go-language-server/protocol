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
	tests := map[string]struct {
		v    any
		want bool
	}{
		"nil interface": {
			v:    nil,
			want: true,
		},
		"zero generated struct": {
			v:    CodeActionDisabled{},
			want: true,
		},
		"non-zero generated struct": {
			v:    CodeActionDisabled{Reason: "server policy"},
			want: false,
		},
		"zero enum without present-zero semantics": {
			v:    InsertTextFormat(0),
			want: true,
		},
		"non-zero enum": {
			v:    InsertTextFormatPlainText,
			want: false,
		},
		"nil raw json value": {
			v:    LSPAny(nil),
			want: true,
		},
		"non-empty raw json value": {
			v:    LSPAny(`{"x":1}`),
			want: false,
		},
		"custom Zeroer type returning true": {
			v:    customZero{val: 42},
			want: true,
		},
		"custom Zeroer type returning false": {
			v:    customZero{val: 12},
			want: false,
		},
		"Optional with value but not set": {
			v:    Optional[string]{value: "hello", set: false},
			want: true,
		},
		"Optional set with zero value": {
			v:    Optional[string]{value: "", set: true},
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isZeroOmitValue(tt.v); got != tt.want {
				t.Fatalf("isZeroOmitValue(%T) = %v, want %v", tt.v, got, tt.want)
			}
		})
	}
}
