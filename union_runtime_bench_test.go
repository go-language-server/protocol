// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"testing"

	"github.com/go-json-experiment/json/jsontext"
)

func BenchmarkUnionScannerObjectHasAndKnown(b *testing.B) {
	cases := []struct {
		name     string
		raw      jsontext.Value
		required []string
		known    []string
	}{
		{
			name:     "text_edit",
			raw:      jsontext.Value(`{"textDocument":{"uri":"file:///a.go","version":1},"edits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}},"newText":"x"}]}`),
			required: []string{"textDocument", "edits"},
			known:    []string{"textDocument", "edits"},
		},
		{
			name:     "escaped_keys",
			raw:      jsontext.Value(`{"\u0075ri":"file:///a","\u0072ange":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}}}`),
			required: []string{"uri", "range"},
			known:    []string{"uri", "range"},
		},
		{
			name:     "foreign_key",
			raw:      jsontext.Value(`{"uri":"file:///a","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}},"future":true}`),
			required: []string{"uri", "range"},
			known:    []string{"uri", "range"},
		},
	}
	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(tc.raw)))
			for b.Loop() {
				_ = objectHasAndKnownGuard(tc.raw, tc.required, tc.known)
			}
		})
	}
}

func BenchmarkUnionScannerArrayFirstHasAndKnown(b *testing.B) {
	raw := jsontext.Value(`[{"name":"a","kind":1,"location":{"uri":"file:///x"},"data":{"k":1}}]`)
	required := []string{"name", "kind", "location"}
	known := []string{"name", "kind", "location", "data"}

	b.ReportAllocs()
	b.SetBytes(int64(len(raw)))
	for b.Loop() {
		_ = arrayFirstHasAndKnown(raw, required, known)
	}
}

func BenchmarkUnionScannerObjectKind(b *testing.B) {
	cases := []struct {
		name string
		raw  jsontext.Value
	}{
		{"plain", jsontext.Value(`{"kind":"create","uri":"file:///a.go"}`)},
		{"escaped_value", jsontext.Value(`{"kind":"cr\u0065ate","uri":"file:///a.go"}`)},
		{"missing", jsontext.Value(`{"uri":"file:///a.go"}`)},
	}
	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(tc.raw)))
			for b.Loop() {
				_, _ = objectKind(tc.raw)
			}
		})
	}
}

func BenchmarkUnionScannerMalformedNoPanic(b *testing.B) {
	inputs := []struct {
		name string
		raw  jsontext.Value
	}{
		{"unterminated_object", jsontext.Value(`{"a":`)},
		{"unterminated_key", jsontext.Value(`{"`)},
		{"trailing_escape", jsontext.Value(`{"key\`)},
		{"space_only", jsontext.Value(`   `)},
	}
	required := []string{"uri", "range"}
	known := []string{"uri", "range"}

	for _, tc := range inputs {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(tc.raw)))
			for b.Loop() {
				_ = objectHasAndKnownGuard(tc.raw, required, known)
				_ = arrayFirstHasAndKnown(tc.raw, required, known)
				_, _ = objectKind(tc.raw)
			}
		})
	}
}
