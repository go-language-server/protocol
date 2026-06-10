// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"testing"

	"github.com/go-json-experiment/json"
	gocmp "github.com/google/go-cmp/cmp"
)

// TestWireDuplicateKeysLastWins pins the relaxed duplicate-name semantics
// enabled by wireOptions: a duplicated object member decodes as last-wins
// instead of erroring. The union dispatch scanner is presence-based (it reads
// the first occurrence of a discriminator), so for payloads whose duplicates
// agree the two layers cannot diverge; conflicting duplicate discriminators
// are out of contract.
func TestWireDuplicateKeysLastWins(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  any
		into  func() any
	}{
		"success: struct field last wins": {
			input: `{"line":1,"character":2,"line":9}`,
			into:  func() any { return &Position{} },
			want:  &Position{Line: 9, Character: 2},
		},
		"success: nested struct duplicate last wins": {
			input: `{"start":{"line":1,"character":1},"end":{"line":2,"character":2},"start":{"line":7,"character":8}}`,
			into:  func() any { return &Range{} },
			want: &Range{
				Start: Position{Line: 7, Character: 8},
				End:   Position{Line: 2, Character: 2},
			},
		},
		"success: union dispatch tolerates agreeing duplicate": {
			input: `{"kind":"markdown","value":"v","kind":"markdown"}`,
			into:  func() any { var v InlayHintTooltip; return &v },
			want: func() any {
				var v InlayHintTooltip = &MarkupContent{Kind: MarkupKindMarkdown, Value: "v"}
				return &v
			}(),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.into()
			if err := Unmarshal([]byte(tt.input), got); err != nil {
				t.Fatalf("Unmarshal(%s): %v", tt.input, err)
			}
			if diff := gocmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Unmarshal(%s) mismatch (-want +got):\n%s", tt.input, diff)
			}
		})
	}
}

// TestWireInvalidUTF8Mangled pins the relaxed UTF-8 semantics enabled by
// wireOptions: strings carrying invalid UTF-8 decode with the invalid bytes
// replaced by U+FFFD instead of rejecting the message, and marshaling a Go
// string holding invalid UTF-8 succeeds with the same replacement.
func TestWireInvalidUTF8Mangled(t *testing.T) {
	t.Parallel()

	type host struct {
		Message string `json:"message"`
	}

	t.Run("success: decode replaces invalid byte with U+FFFD", func(t *testing.T) {
		t.Parallel()

		var got host
		if err := Unmarshal([]byte("{\"message\":\"a\xffb\"}"), &got); err != nil {
			t.Fatalf("Unmarshal invalid UTF-8: %v", err)
		}
		if want := "a�b"; got.Message != want {
			t.Errorf("Message = %q, want %q", got.Message, want)
		}
	})

	t.Run("success: encode accepts invalid UTF-8 and round-trips mangled", func(t *testing.T) {
		t.Parallel()

		raw, err := Marshal(host{Message: "a\xffb"})
		if err != nil {
			t.Fatalf("Marshal invalid UTF-8: %v", err)
		}
		var got host
		if err := Unmarshal(raw, &got); err != nil {
			t.Fatalf("Unmarshal(%s): %v", raw, err)
		}
		if want := "a�b"; got.Message != want {
			t.Errorf("round-tripped Message = %q, want %q", got.Message, want)
		}
	})
}

// TestOptionalNullableStreamingDispatch proves the MarshalJSONTo /
// UnmarshalJSONFrom conversions of Optional and Nullable keep union dispatch
// working even when the caller uses plain encoding (no protocol options on the
// decoder), because UnmarshalJSONFrom re-applies the union unmarshalers.
func TestOptionalNullableStreamingDispatch(t *testing.T) {
	t.Parallel()

	type optionalHost struct {
		Doc Optional[InlayHintTooltip] `json:"doc,omitzero"`
	}
	type nullableHost struct {
		Doc Nullable[InlayHintTooltip] `json:"doc,omitzero"`
	}

	t.Run("success: union inside Optional dispatches without protocol options", func(t *testing.T) {
		t.Parallel()

		var got optionalHost
		if err := json.Unmarshal([]byte(`{"doc":{"kind":"markdown","value":"x"}}`), &got); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		v, ok := got.Doc.Get()
		if !ok {
			t.Fatal("Optional reported absent for present union value")
		}
		want := &MarkupContent{Kind: MarkupKindMarkdown, Value: "x"}
		if diff := gocmp.Diff(InlayHintTooltip(want), v); diff != "" {
			t.Errorf("Optional union mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("success: null clears Optional", func(t *testing.T) {
		t.Parallel()

		got := optionalHost{Doc: NewOptional[InlayHintTooltip](String("old"))}
		if err := json.Unmarshal([]byte(`{"doc":null}`), &got); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if !got.Doc.IsZero() {
			t.Error("Optional not cleared by explicit null")
		}
	})

	t.Run("success: union inside Nullable dispatches and null is distinct", func(t *testing.T) {
		t.Parallel()

		var got nullableHost
		if err := json.Unmarshal([]byte(`{"doc":{"kind":"plaintext","value":"y"}}`), &got); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		v, ok := got.Doc.Get()
		if !ok {
			t.Fatal("Nullable reported no value for present union value")
		}
		want := &MarkupContent{Kind: MarkupKindPlainText, Value: "y"}
		if diff := gocmp.Diff(InlayHintTooltip(want), v); diff != "" {
			t.Errorf("Nullable union mismatch (-want +got):\n%s", diff)
		}

		var asNull nullableHost
		if err := json.Unmarshal([]byte(`{"doc":null}`), &asNull); err != nil {
			t.Fatalf("json.Unmarshal null: %v", err)
		}
		if !asNull.Doc.IsNull() {
			t.Error("Nullable did not record explicit null")
		}
	})
}
