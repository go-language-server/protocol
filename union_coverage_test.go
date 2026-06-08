// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "testing"

// rng is a minimal valid Range payload reused across union fixtures.
const rng = `{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}`

// roundTripUnion unmarshals in into a fresh value of union type T, asserts the
// resolved dynamic type, and verifies the value re-marshals to canonically the
// same JSON. It is the shared workhorse for the broad union sweep below.
func roundTripUnion[T any](t *testing.T, in, wantType string) {
	t.Helper()
	var v T
	if err := Unmarshal([]byte(in), &v); err != nil {
		t.Fatalf("unmarshal %s: %v", in, err)
	}
	if got := typeName(v); got != wantType {
		t.Fatalf("dispatch %s: got %s, want %s", in, got, wantType)
	}
	out, err := Marshal(v)
	if err != nil {
		t.Fatalf("marshal %s: %v", in, err)
	}
	if got, want := canon(t, out), canon(t, []byte(in)); got != want {
		t.Errorf("round-trip %s:\n got=%s\nwant=%s", in, got, want)
	}
}

// TestUnionCoverageSweep is a permanent, breadth-oriented dispatch+round-trip
// sweep over a structurally diverse set of unions: scalar-token discrimination,
// kind-string discriminators, subset/superset object arms, slice arms, and the
// multi-arm structural families. It guards against silent dispatch regressions
// that a handful of hand-picked cases would miss.
func TestUnionCoverageSweep(t *testing.T) {
	// Scalar-token discrimination: the JSON token alone selects the arm.
	t.Run("ProgressToken", func(t *testing.T) {
		roundTripUnion[ProgressToken](t, `42`, "protocol.Integer")
		roundTripUnion[ProgressToken](t, `"tok-1"`, "protocol.String")
	})

	// String vs object glob pattern (RelativePattern requires baseUri+pattern).
	t.Run("GlobPattern", func(t *testing.T) {
		roundTripUnion[GlobPattern](t, `"**/*.go"`, "protocol.Pattern")
		roundTripUnion[GlobPattern](t, `{"baseUri":"file:///w","pattern":"*.go"}`, "*protocol.RelativePattern")
	})

	// String vs []string.
	t.Run("DidChangeConfigurationRegistrationOptionsSection", func(t *testing.T) {
		roundTripUnion[DidChangeConfigurationRegistrationOptionsSection](t, `"one"`, "protocol.String")
		roundTripUnion[DidChangeConfigurationRegistrationOptionsSection](t, `["a","b"]`, "protocol.StringSlice")
	})

	// kind-string discriminator: only the snippet object carries kind:"snippet".
	t.Run("InlineCompletionItemInsertText", func(t *testing.T) {
		roundTripUnion[InlineCompletionItemInsertText](t, `"plain"`, "protocol.String")
		roundTripUnion[InlineCompletionItemInsertText](t, `{"kind":"snippet","value":"$1"}`, "*protocol.StringValue")
	})

	// Object arm vs string arm.
	t.Run("RelativePatternBaseURI", func(t *testing.T) {
		roundTripUnion[RelativePatternBaseURI](t, `{"uri":"file:///w","name":"w"}`, "*protocol.WorkspaceFolder")
		roundTripUnion[RelativePatternBaseURI](t, `"file:///w"`, "protocol.URI")
	})

	// Subset/superset object arms: LocationUriOnly is a strict subset of
	// Location, so the superset must win when range is present and the subset
	// must not be shadowed when it is absent.
	t.Run("WorkspaceSymbolLocation", func(t *testing.T) {
		roundTripUnion[WorkspaceSymbolLocation](t, `{"uri":"file:///x","range":`+rng+`}`, "*protocol.Location")
		roundTripUnion[WorkspaceSymbolLocation](t, `{"uri":"file:///x"}`, "*protocol.LocationUriOnly")
	})

	// Declaration: single object arm plus its slice arm.
	t.Run("Declaration", func(t *testing.T) {
		roundTripUnion[Declaration](t, `{"uri":"file:///x","range":`+rng+`}`, "*protocol.Location")
		roundTripUnion[Declaration](t, `[{"uri":"file:///x","range":`+rng+`}]`, "protocol.LocationSlice")
	})

	// Differ by a single distinctive key set (data vs edits).
	t.Run("SemanticTokensDeltaResult", func(t *testing.T) {
		roundTripUnion[SemanticTokensDeltaResult](t, `{"data":[0,0,1,0,0]}`, "*protocol.SemanticTokens")
		roundTripUnion[SemanticTokensDeltaResult](t, `{"edits":[]}`, "*protocol.SemanticTokensDelta")
	})

	// InsertReplaceEdit is a superset of TextEdit's shape via insert/replace.
	t.Run("CompletionItemTextEdit", func(t *testing.T) {
		roundTripUnion[CompletionItemTextEdit](t, `{"range":`+rng+`,"newText":"x"}`, "*protocol.TextEdit")
		roundTripUnion[CompletionItemTextEdit](t, `{"newText":"x","insert":`+rng+`,"replace":`+rng+`}`, "*protocol.InsertReplaceEdit")
	})

	// Three structurally overlapping arms keyed by their distinctive field.
	t.Run("InlineValue", func(t *testing.T) {
		roundTripUnion[InlineValue](t, `{"range":`+rng+`,"text":"v"}`, "*protocol.InlineValueText")
		roundTripUnion[InlineValue](t, `{"range":`+rng+`,"variableName":"v","caseSensitiveLookup":true}`, "*protocol.InlineValueVariableLookup")
		roundTripUnion[InlineValue](t, `{"range":`+rng+`,"expression":"a+b"}`, "*protocol.InlineValueEvaluatableExpression")
	})
}
