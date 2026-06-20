// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"math"
	"testing"

	"github.com/go-json-experiment/json"
	gocmp "github.com/google/go-cmp/cmp"
)

type semanticTokensPartialResultMarshalShadow struct {
	Data []uint32 `json:"data"`
}

func TestSemanticTokensMarshalFastPathMatchesReflection(t *testing.T) {
	t.Parallel()

	escaped := "r\\\"\n"
	invalidUTF8 := string([]byte{'r', 0xff, 'x'})
	tests := map[string]struct {
		value SemanticTokens
		want  semanticTokensShadow
	}{
		"success: nil data encodes as empty array": {
			value: SemanticTokens{},
			want:  semanticTokensShadow{},
		},
		"success: empty data encodes as empty array": {
			value: SemanticTokens{Data: []uint32{}},
			want:  semanticTokensShadow{Data: []uint32{}},
		},
		"success: max uint32 value": {
			value: SemanticTokens{Data: []uint32{0, 1, math.MaxUint32}},
			want:  semanticTokensShadow{Data: []uint32{0, 1, math.MaxUint32}},
		},
		"success: escaped result id": {
			value: SemanticTokens{ResultID: &escaped, Data: []uint32{1, 2, 3}},
			want:  semanticTokensShadow{ResultID: &escaped, Data: []uint32{1, 2, 3}},
		},
		"success: invalid utf8 result id follows wire options": {
			value: SemanticTokens{ResultID: &invalidUTF8, Data: []uint32{4, 5}},
			want:  semanticTokensShadow{ResultID: &invalidUTF8, Data: []uint32{4, 5}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := Marshal(tt.value)
			if err != nil {
				t.Fatalf("Marshal(SemanticTokens): %v", err)
			}
			want, err := json.Marshal(tt.want, wireOptions)
			if err != nil {
				t.Fatalf("marshal shadow: %v", err)
			}
			if !bytes.Equal(got, want) {
				t.Errorf("Marshal(SemanticTokens) mismatch\ngot:  %s\nwant: %s", got, want)
			}

			gotPtr, err := Marshal(&tt.value)
			if err != nil {
				t.Fatalf("Marshal(*SemanticTokens): %v", err)
			}
			if !bytes.Equal(gotPtr, want) {
				t.Errorf("Marshal(*SemanticTokens) mismatch\ngot:  %s\nwant: %s", gotPtr, want)
			}
		})
	}
}

func TestSemanticTokensPartialResultMarshalFastPathMatchesReflection(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		value SemanticTokensPartialResult
		want  semanticTokensPartialResultMarshalShadow
	}{
		"success: nil data encodes as empty array": {
			value: SemanticTokensPartialResult{},
			want:  semanticTokensPartialResultMarshalShadow{},
		},
		"success: empty data encodes as empty array": {
			value: SemanticTokensPartialResult{Data: []uint32{}},
			want:  semanticTokensPartialResultMarshalShadow{Data: []uint32{}},
		},
		"success: values": {
			value: SemanticTokensPartialResult{Data: []uint32{7, 8, math.MaxUint32}},
			want:  semanticTokensPartialResultMarshalShadow{Data: []uint32{7, 8, math.MaxUint32}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := Marshal(tt.value)
			if err != nil {
				t.Fatalf("Marshal(SemanticTokensPartialResult): %v", err)
			}
			want, err := json.Marshal(tt.want, wireOptions)
			if err != nil {
				t.Fatalf("marshal shadow: %v", err)
			}
			if !bytes.Equal(got, want) {
				t.Errorf("Marshal(SemanticTokensPartialResult) mismatch\ngot:  %s\nwant: %s", got, want)
			}

			gotPtr, err := Marshal(&tt.value)
			if err != nil {
				t.Fatalf("Marshal(*SemanticTokensPartialResult): %v", err)
			}
			if !bytes.Equal(gotPtr, want) {
				t.Errorf("Marshal(*SemanticTokensPartialResult) mismatch\ngot:  %s\nwant: %s", gotPtr, want)
			}
		})
	}
}

func TestSemanticTokensMarshalFastPathNilPointers(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		value any
	}{
		"success: nil semantic tokens pointer": {
			value: (*SemanticTokens)(nil),
		},
		"success: nil semantic tokens partial result pointer": {
			value: (*SemanticTokensPartialResult)(nil),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := Marshal(tt.value)
			if err != nil {
				t.Fatalf("Marshal(%T): %v", tt.value, err)
			}
			if !bytes.Equal(got, []byte("null")) {
				t.Errorf("Marshal(%T) = %s, want null", tt.value, got)
			}
		})
	}
}

func TestAppendSemanticTokensJSONHelpers(t *testing.T) {
	t.Parallel()

	t.Run("appendUint32JSONArray encodes nil and max values", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			data []uint32
			want []byte
		}{
			"success: nil": {
				want: []byte(`prefix:[]`),
			},
			"success: values": {
				data: []uint32{0, 9, math.MaxUint32},
				want: []byte(`prefix:[0,9,4294967295]`),
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got := appendUint32JSONArray([]byte("prefix:"), tt.data)
				if diff := gocmp.Diff(tt.want, got); diff != "" {
					t.Errorf("appendUint32JSONArray mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})
}

func TestSemanticTokensLengthHints(t *testing.T) {
	t.Parallel()

	t.Run("uint32 decimal length", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			value uint32
			want  int
		}{
			"success: one digit":  {value: 0, want: 1},
			"success: two digits": {value: 10, want: 2},
			"success: max":        {value: math.MaxUint32, want: 10},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if got := uint32DecimalLen(tt.value); got != tt.want {
					t.Errorf("uint32DecimalLen(%d) = %d, want %d", tt.value, got, tt.want)
				}
			})
		}
	})

	t.Run("array length hint matches appended length", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			data []uint32
		}{
			"success: nil":    {},
			"success: values": {data: []uint32{0, 11, math.MaxUint32}},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got := appendUint32JSONArray(nil, tt.data)
				if hint := uint32JSONArrayLen(tt.data); hint != len(got) {
					t.Errorf("uint32JSONArrayLen(%v) = %d, want %d", tt.data, hint, len(got))
				}
			})
		}
	})
}
