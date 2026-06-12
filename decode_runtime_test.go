// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"math"
	"strings"
	"testing"

	gocmp "github.com/google/go-cmp/cmp"
)

func TestDVString(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input   string
		want    string
		wantErr bool
	}{
		"success: plain":                       {input: `"hello"`, want: "hello"},
		"success: empty":                       {input: `""`, want: ""},
		"success: multibyte UTF-8":             {input: `"こんにちは"`, want: "こんにちは"},
		"success: DEL byte stays fast path":    {input: "\"a\x7fb\"", want: "a\x7fb"},
		"success: simple escapes":              {input: `"a\"b\\c\/d\be\ff\ng\rh\ti"`, want: "a\"b\\c/d\be\ff\ng\rh\ti"},
		"success: unicode escape":              {input: `"é"`, want: "é"},
		"success: surrogate pair":              {input: `"😀"`, want: "😀"},
		"success: invalid UTF-8 mangles":       {input: "\"a\xffb\"", want: "a�b"},
		"success: each invalid byte mangles":   {input: "\"\xff\xfe\"", want: "��"},
		"success: truncated rune then ascii":   {input: "\"\xc3a\"", want: "�a"},
		"success: lone high surrogate mangles": {input: `"\ud83d!"`, want: "�!"},
		"success: lone low surrogate mangles":  {input: `"\ude00"`, want: "�"},
		"success: double high surrogate":       {input: `"\ud800\ud800"`, want: "��"},
		"error: surrogate with bad hex pair":   {input: `"\ud800\uzzzz"`, wantErr: true},
		"error: invalid escape":                {input: `"\q"`, wantErr: true},
		"error: bad hex":                       {input: `"\u00zz"`, wantErr: true},
		"error: bare control char":             {input: "\"a\nb\"", wantErr: true},
		"error: unterminated":                  {input: `"abc`, wantErr: true},
		"error: not a string":                  {input: `123`, wantErr: true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, next, err := dvString([]byte(tt.input), 0)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("dvString(%q) = %q, want error", tt.input, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("dvString(%q): %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("dvString(%q) = %q, want %q", tt.input, got, tt.want)
			}
			if next != len(tt.input) {
				t.Errorf("next = %d, want %d", next, len(tt.input))
			}
		})
	}
}

func TestDVNumbers(t *testing.T) {
	t.Parallel()

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			input   string
			want    uint32
			wantErr bool
		}{
			"success: zero":        {input: "0", want: 0},
			"success: max":         {input: "4294967295", want: math.MaxUint32},
			"success: delimited":   {input: "42,", want: 42},
			"error: overflow":      {input: "4294967296", wantErr: true},
			"error: negative":      {input: "-1", wantErr: true},
			"error: exponent":      {input: "1e2", wantErr: true},
			"error: fraction":      {input: "1.5", wantErr: true},
			"error: leading zero":  {input: "01", wantErr: true},
			"error: trailing junk": {input: "12x", wantErr: true},
			"error: empty":         {input: "", wantErr: true},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got, _, err := dvUint32([]byte(tt.input), 0)
				if tt.wantErr {
					if err == nil {
						t.Fatalf("dvUint32(%q) = %d, want error", tt.input, got)
					}
					return
				}
				if err != nil {
					t.Fatalf("dvUint32(%q): %v", tt.input, err)
				}
				if got != tt.want {
					t.Errorf("dvUint32(%q) = %d, want %d", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			input   string
			want    int32
			wantErr bool
		}{
			"success: min":            {input: "-2147483648", want: math.MinInt32},
			"success: max":            {input: "2147483647", want: math.MaxInt32},
			"success: negative zero":  {input: "-0", want: 0},
			"error: overflow":         {input: "2147483648", wantErr: true},
			"error: under min":        {input: "-2147483649", wantErr: true},
			"error: exponent":         {input: "2e1", wantErr: true},
			"error: huge digit count": {input: "999999999999999999999", wantErr: true},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got, _, err := dvInt32([]byte(tt.input), 0)
				if tt.wantErr {
					if err == nil {
						t.Fatalf("dvInt32(%q) = %d, want error", tt.input, got)
					}
					return
				}
				if err != nil {
					t.Fatalf("dvInt32(%q): %v", tt.input, err)
				}
				if got != tt.want {
					t.Errorf("dvInt32(%q) = %d, want %d", tt.input, got, tt.want)
				}
			})
		}
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			input   string
			want    float64
			wantErr bool
		}{
			"success: integer":   {input: "5", want: 5},
			"success: fraction":  {input: "0.25", want: 0.25},
			"success: exponent":  {input: "1e3", want: 1000},
			"success: negative":  {input: "-2.5e-1", want: -0.25},
			"error: bare minus":  {input: "-", wantErr: true},
			"error: dot no frac": {input: "1.", wantErr: true},
			"error: exp no num":  {input: "1e", wantErr: true},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got, _, err := dvFloat64([]byte(tt.input), 0)
				if tt.wantErr {
					if err == nil {
						t.Fatalf("dvFloat64(%q) = %v, want error", tt.input, got)
					}
					return
				}
				if err != nil {
					t.Fatalf("dvFloat64(%q): %v", tt.input, err)
				}
				if got != tt.want {
					t.Errorf("dvFloat64(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})
}

func TestDVSlices(t *testing.T) {
	t.Parallel()

	t.Run("uint32 slice decodes and reuses backing storage", func(t *testing.T) {
		t.Parallel()

		dst := make([]uint32, 2, 8)
		got, next, err := dvUint32Slice([]byte(` [1, 2,3]`), 1, dst)
		if err != nil {
			t.Fatalf("dvUint32Slice: %v", err)
		}
		if diff := gocmp.Diff([]uint32{1, 2, 3}, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
		if next != len(` [1, 2,3]`) {
			t.Errorf("next = %d", next)
		}
		if &got[0] != &dst[:1][0] {
			t.Error("backing storage not reused")
		}
	})

	t.Run("empty array yields non-nil empty slice", func(t *testing.T) {
		t.Parallel()

		got, _, err := dvUint32Slice([]byte(`[]`), 0, []uint32(nil))
		if err != nil {
			t.Fatalf("dvUint32Slice: %v", err)
		}
		if got == nil || len(got) != 0 {
			t.Errorf("got %#v, want non-nil empty", got)
		}
	})

	t.Run("null yields nil slice", func(t *testing.T) {
		t.Parallel()

		got, _, err := dvUint32Slice([]byte(`null`), 0, []uint32{9})
		if err != nil {
			t.Fatalf("dvUint32Slice: %v", err)
		}
		if got != nil {
			t.Errorf("got %#v, want nil", got)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()

		got, _, err := dvStringSlice([]byte(`["a","é"]`), 0, []string(nil))
		if err != nil {
			t.Fatalf("dvStringSlice: %v", err)
		}
		if diff := gocmp.Diff([]string{"a", "é"}, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("uint32 slice rejects invalid number forms", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			input string
		}{
			"error: trailing comma": {input: `[1,]`},
			"error: empty element":  {input: `[1,,2]`},
			"error: exponent":       {input: `[1e2]`},
			"error: fraction":       {input: `[1.5]`},
			"error: leading zero":   {input: `[01]`},
			"error: negative":       {input: `[-1]`},
			"error: plus sign":      {input: `[+1]`},
			"error: overflow":       {input: `[4294967296]`},
			"error: unterminated":   {input: `[1,2`},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if _, _, err := dvUint32Slice([]byte(tt.input), 0, []uint32(nil)); err == nil {
					t.Fatalf("dvUint32Slice(%s) succeeded; want error", tt.input)
				}
			})
		}
	})

	t.Run("slice capacity hint is bounded and positive", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			input   string
			perElem int
			want    int
		}{
			"success: short run":      {input: `1,2,3]`, perElem: 3, want: 3},
			"success: at end":         {input: ``, perElem: 3, want: 1},
			"success: bounded at cap": {input: strings.Repeat("1,", 4096), perElem: 1, want: 4096},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if got := dvSliceCapHint([]byte(tt.input), 0, tt.perElem); got != tt.want {
					t.Errorf("dvSliceCapHint(%q, %d) = %d, want %d", tt.input, tt.perElem, got, tt.want)
				}
			})
		}
	})
}

func TestDVObjectGrammar(t *testing.T) {
	t.Parallel()

	t.Run("member key with escaped spelling", func(t *testing.T) {
		t.Parallel()

		raw := []byte(`{"label" : 1}`)
		key, next, err := dvMemberKey(raw, 1)
		if err != nil {
			t.Fatalf("dvMemberKey: %v", err)
		}
		if !keyEquals(key, "label") {
			t.Errorf("keyEquals(%q, label) = false", key)
		}
		if raw[next] != '1' {
			t.Errorf("next points at %q, want value", raw[next])
		}
	})

	t.Run("error: missing colon", func(t *testing.T) {
		t.Parallel()

		if _, _, err := dvMemberKey([]byte(`{"k" 1}`), 1); err == nil {
			t.Error("missing colon accepted")
		}
	})

	t.Run("object separator handling", func(t *testing.T) {
		t.Parallel()

		next, done, err := dvObjectNext([]byte(`{"a":1 , "b":2}`), 6)
		if err != nil || done {
			t.Fatalf("dvObjectNext = done %v err %v", done, err)
		}
		if next != 9 {
			t.Errorf("next = %d, want 9", next)
		}

		next, done, err = dvObjectNext([]byte(`{"a":1}`), 6)
		if err != nil || !done || next != 7 {
			t.Errorf("close: next=%d done=%v err=%v", next, done, err)
		}

		if _, _, err := dvObjectNext([]byte(`{"a":1;}`), 6); err == nil {
			t.Error("bad separator accepted")
		}
	})

	t.Run("dvEnd rejects trailing content", func(t *testing.T) {
		t.Parallel()

		if err := dvEnd([]byte(`{} x`), 2); err == nil {
			t.Error("trailing content accepted")
		}
		if err := dvEnd([]byte("{} \n\t"), 2); err != nil {
			t.Errorf("trailing whitespace rejected: %v", err)
		}
	})
}

func TestDVValueStrictValidation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input   string
		want    string
		wantErr bool
	}{
		"success: nested object": {input: ` {"a":[1,{"b":"c"}]} ,`, want: `{"a":[1,{"b":"c"}]}`},
		"success: duplicate object names": {
			input: `{"a":1,"a":2}`,
			want:  `{"a":1,"a":2}`,
		},
		"success: invalid UTF-8 string allowed": {
			input: "\"\xff\"",
			want:  "\"\xff\"",
		},
		"error: invalid scalar":         {input: `<bad>`, wantErr: true},
		"error: trailing comma array":   {input: `[1,]`, wantErr: true},
		"error: trailing comma object":  {input: `{"a":1,}`, wantErr: true},
		"error: invalid object key":     {input: `{"\q":1}`, wantErr: true},
		"error: invalid nested value":   {input: `{"a":[true, <bad>]}`, wantErr: true},
		"error: unterminated structure": {input: `{"a":[1]`, wantErr: true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, next, err := dvValue([]byte(tt.input), 0)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("dvValue(%q) = %q, next %d, want error", tt.input, got, next)
				}
				return
			}
			if err != nil {
				t.Fatalf("dvValue(%q): %v", tt.input, err)
			}
			if string(got) != tt.want {
				t.Errorf("dvValue(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
