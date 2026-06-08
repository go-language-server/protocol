// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"reflect"
	"testing"
)

// assertWireRoundTrip decodes jsonStr into T, re-marshals, re-decodes, and asserts
// the two decoded values are deeply equal — i.e. the generated 3.18 type both
// accepts the wire fixture and reproduces it.
func assertWireRoundTrip[T any](t *testing.T, jsonStr string) {
	t.Helper()

	var first T
	if err := Unmarshal([]byte(jsonStr), &first); err != nil {
		t.Fatalf("decode %q: %v", jsonStr, err)
	}
	out, err := Marshal(first)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var second T
	if err := Unmarshal(out, &second); err != nil {
		t.Fatalf("re-decode %q: %v", out, err)
	}
	if !reflect.DeepEqual(first, second) {
		t.Errorf("round-trip mismatch\n in:  %s\n out: %s", jsonStr, out)
	}
}

// TestComplianceFixtures replays representative LSP wire-format JSON (message
// shapes stable across 3.15..3.18) through the generated 3.18 types. They are an
// independent wire-format oracle: real LSP message bodies must decode and
// round-trip. A failure here is triaged as a generator bug vs a legitimate 3.18
// shape change, per the migration plan.
func TestComplianceFixtures(t *testing.T) {
	t.Parallel()

	t.Run("Position", func(t *testing.T) {
		t.Parallel()
		assertWireRoundTrip[Position](t, `{"line":5,"character":10}`)
	})
	t.Run("Range", func(t *testing.T) {
		t.Parallel()
		assertWireRoundTrip[Range](t, `{"start":{"line":0,"character":0},"end":{"line":1,"character":2}}`)
	})
	t.Run("Location", func(t *testing.T) {
		t.Parallel()
		assertWireRoundTrip[Location](t, `{"uri":"file:///x.go","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":4}}}`)
	})
	t.Run("TextEdit", func(t *testing.T) {
		t.Parallel()
		assertWireRoundTrip[TextEdit](t, `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":4}},"newText":"hello"}`)
	})
	t.Run("MarkupContent", func(t *testing.T) {
		t.Parallel()
		assertWireRoundTrip[MarkupContent](t, `{"kind":"markdown","value":"# Title"}`)
	})
	t.Run("Diagnostic", func(t *testing.T) {
		t.Parallel()
		assertWireRoundTrip[Diagnostic](t, `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":3}},"severity":1,"message":"boom"}`)
	})
}

// TestForwardCompatIgnoresUnknownFields asserts the decoder ignores object members
// the 3.18 types do not declare. LSP requires this so a peer tolerates a newer
// counterpart; a decoder that rejected unknown members would break interop.
func TestForwardCompatIgnoresUnknownFields(t *testing.T) {
	t.Parallel()

	data := []byte(`{"uri":"file:///x.go","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":4}},"futureField":42}`)
	var loc Location
	if err := Unmarshal(data, &loc); err != nil {
		t.Fatalf("decode with an unknown field must succeed for forward compatibility, got: %v", err)
	}
	if loc.URI != "file:///x.go" {
		t.Errorf("URI = %q, want file:///x.go", loc.URI)
	}
}
