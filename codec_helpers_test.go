// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"strings"
	"testing"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

func TestDecodePositionFrom(t *testing.T) {
	pos := Position{Line: 9, Character: 8}
	if err := decodePositionFrom(jsontext.NewDecoder(strings.NewReader(`{"line":1,"character":2,"extra":true}`)), &pos); err != nil {
		t.Fatalf("decode position: %v", err)
	}
	if pos != (Position{Line: 1, Character: 2}) {
		t.Fatalf("decoded position = %#v, want line 1 character 2", pos)
	}

	if err := decodePositionFrom(jsontext.NewDecoder(strings.NewReader(`{"line":3}`)), &pos); err != nil {
		t.Fatalf("merge position: %v", err)
	}
	if pos != (Position{Line: 3, Character: 2}) {
		t.Fatalf("merged position = %#v, want line 3 character 2", pos)
	}

	if err := decodePositionFrom(jsontext.NewDecoder(strings.NewReader(`null`)), &pos); err != nil {
		t.Fatalf("decode null position: %v", err)
	}
	if pos != (Position{}) {
		t.Fatalf("null position = %#v, want zero", pos)
	}

	if err := decodePositionFrom(jsontext.NewDecoder(strings.NewReader(`{"line":"bad"}`)), &pos); err == nil {
		t.Fatal("decode string line position succeeded")
	}
}

func TestDecodeRangeFrom(t *testing.T) {
	rng := Range{
		Start: Position{Line: 7, Character: 8},
		End:   Position{Line: 9, Character: 10},
	}
	data := `{"start":{"line":1,"character":2},"end":{"line":3,"character":4},"extra":{"ignored":true}}`
	if err := decodeRangeFrom(jsontext.NewDecoder(strings.NewReader(data)), &rng); err != nil {
		t.Fatalf("decode range: %v", err)
	}
	want := Range{Start: Position{Line: 1, Character: 2}, End: Position{Line: 3, Character: 4}}
	if rng != want {
		t.Fatalf("decoded range = %#v, want %#v", rng, want)
	}

	if err := decodeRangeFrom(jsontext.NewDecoder(strings.NewReader(`{"start":{"line":5}}`)), &rng); err != nil {
		t.Fatalf("merge range: %v", err)
	}
	want = Range{Start: Position{Line: 5, Character: 2}, End: Position{Line: 3, Character: 4}}
	if rng != want {
		t.Fatalf("merged range = %#v, want %#v", rng, want)
	}

	if err := decodeRangeFrom(jsontext.NewDecoder(strings.NewReader(`null`)), &rng); err != nil {
		t.Fatalf("decode null range: %v", err)
	}
	if rng != (Range{}) {
		t.Fatalf("null range = %#v, want zero", rng)
	}
}

func TestDecodeProgressTokenFrom(t *testing.T) {
	var got ProgressToken
	if err := decodeProgressTokenFrom(jsontext.NewDecoder(strings.NewReader(`"token"`)), &got); err != nil {
		t.Fatalf("decode string token: %v", err)
	}
	if v, ok := got.(String); !ok || string(v) != "token" {
		t.Fatalf("string token = %#v, want String token", got)
	}

	if err := decodeProgressTokenFrom(jsontext.NewDecoder(strings.NewReader(`42`)), &got); err != nil {
		t.Fatalf("decode integer token: %v", err)
	}
	if v, ok := got.(Integer); !ok || v != 42 {
		t.Fatalf("integer token = %#v, want Integer 42", got)
	}

	if err := decodeProgressTokenFrom(jsontext.NewDecoder(strings.NewReader(`null`)), &got); err != nil {
		t.Fatalf("decode null token: %v", err)
	}
	if got != nil {
		t.Fatalf("null token = %#v, want nil", got)
	}

	if err := decodeProgressTokenFrom(jsontext.NewDecoder(strings.NewReader(`false`)), &got); err == nil {
		t.Fatal("decode bool progress token succeeded")
	}
}

func TestDecodeDiagnosticTagsFrom(t *testing.T) {
	got := NewDiagnosticTags(DiagnosticTagDeprecated)
	if err := decodeDiagnosticTagsFrom(jsontext.NewDecoder(strings.NewReader(`[1,2]`)), &got); err != nil {
		t.Fatalf("decode diagnostic tags: %v", err)
	}
	want := []DiagnosticTag{DiagnosticTagUnnecessary, DiagnosticTagDeprecated}
	if gotTags := got.Slice(); len(gotTags) != len(want) || gotTags[0] != want[0] || gotTags[1] != want[1] {
		t.Fatalf("diagnostic tags = %#v, want %#v", gotTags, want)
	}

	if err := decodeDiagnosticTagsFrom(jsontext.NewDecoder(strings.NewReader(`null`)), &got); err != nil {
		t.Fatalf("decode null diagnostic tags: %v", err)
	}
	if !got.IsZero() || got.Len() != 0 {
		t.Fatalf("null diagnostic tags = %#v, want nil", got)
	}

	if err := decodeDiagnosticTagsFrom(jsontext.NewDecoder(strings.NewReader(`["bad"]`)), &got); err == nil {
		t.Fatal("decode string diagnostic tag succeeded")
	}
}

func TestDiagnosticTagsAccessorsAndJSON(t *testing.T) {
	var tags DiagnosticTags
	if !tags.IsZero() || tags.Len() != 0 || tags.Slice() != nil {
		t.Fatalf("zero diagnostic tags = %#v, want empty", tags)
	}

	tags.Set([]DiagnosticTag{DiagnosticTagUnnecessary})
	if tags.IsZero() || tags.Len() != 1 {
		t.Fatalf("single diagnostic tag state = %#v, want one present", tags)
	}
	if got := tags.Slice(); len(got) != 1 || got[0] != DiagnosticTagUnnecessary {
		t.Fatalf("single diagnostic tag slice = %#v, want [1]", got)
	}

	tags.Set([]DiagnosticTag{DiagnosticTagUnnecessary, DiagnosticTagDeprecated})
	if tags.Len() != 2 {
		t.Fatalf("two diagnostic tag length = %d, want 2", tags.Len())
	}
	gotJSON, err := json.Marshal(tags)
	if err != nil {
		t.Fatalf("marshal diagnostic tags: %v", err)
	}
	if string(gotJSON) != `[1,2]` {
		t.Fatalf("marshal diagnostic tags = %s, want [1,2]", gotJSON)
	}

	tags.Clear()
	if !tags.IsZero() || tags.Len() != 0 {
		t.Fatalf("cleared diagnostic tags = %#v, want empty", tags)
	}
	if err := json.Unmarshal([]byte(`[2]`), &tags); err != nil {
		t.Fatalf("unmarshal diagnostic tags: %v", err)
	}
	if got := tags.Slice(); len(got) != 1 || got[0] != DiagnosticTagDeprecated {
		t.Fatalf("unmarshal diagnostic tags = %#v, want [2]", got)
	}
}

func TestEncodeProgressTokenTo(t *testing.T) {
	tests := map[string]struct {
		token ProgressToken
		want  string
	}{
		"string":  {token: String("token"), want: `"token"`},
		"integer": {token: Integer(42), want: `42`},
		"nil":     {token: nil, want: `null`},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var b strings.Builder
			enc := jsontext.NewEncoder(&b)
			if err := encodeProgressTokenTo(enc, tc.token); err != nil {
				t.Fatalf("encode progress token: %v", err)
			}
			if got := strings.TrimSpace(b.String()); got != tc.want {
				t.Fatalf("encoded progress token = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestEncodeDiagnosticTagsTo(t *testing.T) {
	var b strings.Builder
	enc := jsontext.NewEncoder(&b)
	tags := NewDiagnosticTags(DiagnosticTagUnnecessary, DiagnosticTagDeprecated)
	if err := encodeDiagnosticTagsTo(enc, tags); err != nil {
		t.Fatalf("encode diagnostic tags: %v", err)
	}
	if got, want := strings.TrimSpace(b.String()), `[1,2]`; got != want {
		t.Fatalf("encoded diagnostic tags = %s, want %s", got, want)
	}
}

func TestEncodePositionTo(t *testing.T) {
	var b strings.Builder
	enc := jsontext.NewEncoder(&b)
	if err := encodePositionTo(enc, Position{Line: 1, Character: 2}); err != nil {
		t.Fatalf("encode position: %v", err)
	}
	if got, want := strings.TrimSpace(b.String()), `{"line":1,"character":2}`; got != want {
		t.Fatalf("encoded position = %s, want %s", got, want)
	}
}

func TestEncodeRangeTo(t *testing.T) {
	var b strings.Builder
	enc := jsontext.NewEncoder(&b)
	rng := Range{Start: Position{Line: 1, Character: 2}, End: Position{Line: 3, Character: 4}}
	if err := encodeRangeTo(enc, rng); err != nil {
		t.Fatalf("encode range: %v", err)
	}
	want := `{"start":{"line":1,"character":2},"end":{"line":3,"character":4}}`
	if got := strings.TrimSpace(b.String()); got != want {
		t.Fatalf("encoded range = %s, want %s", got, want)
	}
}
