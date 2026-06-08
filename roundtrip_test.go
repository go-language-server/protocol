// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"slices"
	"testing"

	"github.com/go-json-experiment/json/jsontext"
)

// canon canonicalizes JSON (sorted object keys, minimal whitespace) so that
// round-trip comparisons are independent of field order.
func canon(t *testing.T, b []byte) string {
	t.Helper()
	v := jsontext.Value(slices.Clone(b))
	if err := v.Canonicalize(); err != nil {
		t.Fatalf("canonicalize %s: %v", b, err)
	}
	return string(v)
}

// TestUnionDocumentChanges verifies that every arm of the documentChanges union
// decodes to the correct concrete variant (via type switch) and round-trips.
func TestUnionDocumentChanges(t *testing.T) {
	tests := map[string]struct {
		json   string
		assert func(t *testing.T, v DocumentChange)
	}{
		"success: create (kind discriminator)": {
			json: `{"kind":"create","uri":"file:///a.go"}`,
			assert: func(t *testing.T, v DocumentChange) {
				t.Helper()
				if _, ok := v.(*CreateFile); !ok {
					t.Fatalf("got %T, want *CreateFile", v)
				}
			},
		},
		"success: rename (kind discriminator)": {
			json: `{"kind":"rename","oldUri":"file:///a.go","newUri":"file:///b.go"}`,
			assert: func(t *testing.T, v DocumentChange) {
				t.Helper()
				if _, ok := v.(*RenameFile); !ok {
					t.Fatalf("got %T, want *RenameFile", v)
				}
			},
		},
		"success: delete (kind discriminator)": {
			json: `{"kind":"delete","uri":"file:///a.go"}`,
			assert: func(t *testing.T, v DocumentChange) {
				t.Helper()
				if _, ok := v.(*DeleteFile); !ok {
					t.Fatalf("got %T, want *DeleteFile", v)
				}
			},
		},
		"success: textDocumentEdit (structural, no kind)": {
			json: `{"textDocument":{"uri":"file:///a.go","version":1},` +
				`"edits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}},"newText":"x"}]}`,
			assert: func(t *testing.T, v DocumentChange) {
				t.Helper()
				if _, ok := v.(*TextDocumentEdit); !ok {
					t.Fatalf("got %T, want *TextDocumentEdit", v)
				}
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var v DocumentChange
			if err := Unmarshal([]byte(tt.json), &v); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			tt.assert(t, v)

			out, err := Marshal(v)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if got, want := canon(t, out), canon(t, []byte(tt.json)); got != want {
				t.Errorf("round-trip mismatch:\n got=%s\nwant=%s", got, want)
			}
		})
	}
}

// TestLSPAnyRoundTrip verifies that the raw LSPAny type round-trips every JSON
// shape, including deeply nested objects and arrays, byte-for-byte.
func TestLSPAnyRoundTrip(t *testing.T) {
	tests := map[string]string{
		"success: object":  `{"a":1,"b":"two"}`,
		"success: array":   `[1,2,3]`,
		"success: string":  `"hello"`,
		"success: integer": `42`,
		"success: decimal": `3.14`,
		"success: boolean": `true`,
		"success: null":    `null`,
		"success: nested":  `{"outer":[{"inner":[1,{"deep":true}]},null,"x"]}`,
	}
	for name, in := range tests {
		t.Run(name, func(t *testing.T) {
			var v LSPAny
			if err := Unmarshal([]byte(in), &v); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			out, err := Marshal(v)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if got, want := canon(t, out), canon(t, []byte(in)); got != want {
				t.Errorf("round-trip mismatch: got=%s want=%s", got, want)
			}
		})
	}
}

// TestTupleAndScalarUnion verifies the string|tuple union (ParameterInformation
// label) discriminates by JSON token and round-trips.
func TestTupleAndScalarUnion(t *testing.T) {
	t.Run("success: string arm", func(t *testing.T) {
		var v ParameterInformationLabel
		if err := Unmarshal([]byte(`"param"`), &v); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if got, ok := v.(String); !ok || got != "param" {
			t.Fatalf("got %T %v, want String \"param\"", v, v)
		}
	})
	t.Run("success: tuple arm", func(t *testing.T) {
		var v ParameterInformationLabel
		if err := Unmarshal([]byte(`[3,7]`), &v); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		tup, ok := v.(ParameterInformationLabelTuple)
		if !ok || tup != [2]uint32{3, 7} {
			t.Fatalf("got %T %v, want ParameterInformationLabelTuple{3,7}", v, v)
		}
		out, err := Marshal(v)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		if got, want := canon(t, out), `[3,7]`; got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

// TestEnumRoundTrip verifies a uint32-backed enumeration round-trips.
func TestEnumRoundTrip(t *testing.T) {
	if SymbolKindClass != 5 {
		t.Fatalf("SymbolKindClass = %d, want 5", SymbolKindClass)
	}
	var k SymbolKind
	if err := Unmarshal([]byte(`5`), &k); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if k != SymbolKindClass {
		t.Errorf("got %d, want SymbolKindClass(5)", k)
	}
}

// TestMethodsRegistry verifies the generated method registry is complete.
func TestMethodsRegistry(t *testing.T) {
	if got, want := len(Methods), 95; got != want {
		t.Errorf("len(Methods) = %d, want %d (69 requests + 26 notifications)", got, want)
	}
	if MethodTextDocumentImplementation != "textDocument/implementation" {
		t.Errorf("MethodTextDocumentImplementation = %q", MethodTextDocumentImplementation)
	}
}
