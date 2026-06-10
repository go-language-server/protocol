// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"strings"
	"testing"
)

func TestUnmarshalTypedNilDestinationsReturnError(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		data string
		dst  any
	}{
		"byte walker empty object": {
			data: `{}`,
			dst:  (*CompletionItem)(nil),
		},
		"byte walker null": {
			data: `null`,
			dst:  (*CompletionItem)(nil),
		},
		"byte walker known field": {
			data: `{"label":"x"}`,
			dst:  (*CompletionItem)(nil),
		},
		"union root": {
			data: `"token"`,
			dst:  (*ProgressToken)(nil),
		},
		"nil interface": {
			data: `{}`,
			dst:  nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("Unmarshal panicked for typed nil destination: %v", r)
				}
			}()
			if err := Unmarshal([]byte(tt.data), tt.dst); err == nil {
				t.Fatalf("Unmarshal(%s, %T) succeeded, want invalid destination error", tt.data, tt.dst)
			}
		})
	}
}

// TestOptionalStructZeroValueOmitted documents the one boundary of the Safe-only
// pointer policy: a converted value struct that is present in the input but equal
// to its zero value is omitted on re-marshal (treated as absent). This is benign
// for real data — a non-degenerate Command always carries a non-empty field — but
// a schema-valid all-empty struct does not round-trip byte-for-byte.
func TestOptionalStructZeroValueOmitted(t *testing.T) {
	var v CompletionItem
	if err := Unmarshal([]byte(`{"label":"x","command":{"title":"","command":""}}`), &v); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	out, err := Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if strings.Contains(string(out), "command") {
		t.Fatalf("expected all-zero Command to be omitted, got %s", out)
	}
}

// TestOptionalStructOmitzero pins the optional-struct pointer policy:
//
//   - a "Safe-only" converted value struct (CompletionItem.command, whose zero
//     value is invalid because Command requires non-empty strings) round-trips
//     when present and is omitted via omitzero when absent; and
//   - a zero-meaningful struct kept as a pointer (LocationLink.originSelectionRange,
//     a Range that is valid at the origin) round-trips its present-but-all-zero
//     value instead of being dropped.
func TestOptionalStructOmitzero(t *testing.T) {
	tests := map[string]struct {
		json string
		into func() any
	}{
		"success: converted value struct present (CompletionItem.command)": {
			json: `{"label":"x","command":{"title":"t","command":"c"}}`,
			into: func() any { return new(CompletionItem) },
		},
		"success: converted value struct absent omits the field": {
			json: `{"label":"x"}`,
			into: func() any { return new(CompletionItem) },
		},
		"success: kept-pointer present-zero Range round-trips (LocationLink.originSelectionRange)": {
			json: `{"targetUri":"file:///x",` +
				`"targetRange":{"start":{"line":1,"character":0},"end":{"line":1,"character":2}},` +
				`"targetSelectionRange":{"start":{"line":1,"character":0},"end":{"line":1,"character":1}},` +
				`"originSelectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
			into: func() any { return new(LocationLink) },
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			v := tt.into()
			if err := Unmarshal([]byte(tt.json), v); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
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
