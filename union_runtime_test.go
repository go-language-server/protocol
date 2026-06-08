// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// loadCorpusInputs returns the raw bytes of every testdata/corpus payload, so
// the differential oracle and fuzz seeds exercise real LSP shapes.
func loadCorpusInputs(tb testing.TB) [][]byte {
	tb.Helper()
	matches, err := filepath.Glob(filepath.Join("testdata", "corpus", "*.json"))
	if err != nil {
		tb.Fatalf("glob corpus: %v", err)
	}
	out := make([][]byte, 0, len(matches))
	for _, m := range matches {
		b, err := os.ReadFile(m)
		if err != nil {
			tb.Fatalf("read %s: %v", m, err)
		}
		out = append(out, b)
	}
	return out
}

// The functions below are the pre-optimization map-based implementations of the
// union-dispatch predicates, retained here as a differential oracle. The
// production scanner in union_runtime.go must agree with them on every input;
// TestScannerMatchesMapOracle enforces that.

func oracleObjectFields(raw jsontext.Value) map[string]jsontext.Value {
	var m map[string]jsontext.Value
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil
	}
	return m
}

func oracleObjectKind(raw jsontext.Value) (string, bool) {
	v, ok := oracleObjectFields(raw)["kind"]
	if !ok {
		return "", false
	}
	var s string
	if err := json.Unmarshal(v, &s); err != nil {
		return "", false
	}
	return s, true
}

func oracleObjectHasKeys(raw jsontext.Value, keys ...string) bool {
	fields := oracleObjectFields(raw)
	if fields == nil {
		return false
	}
	for _, k := range keys {
		if _, ok := fields[k]; !ok {
			return false
		}
	}
	return true
}

func oracleObjectKeysKnown(raw jsontext.Value, known ...string) bool {
	fields := oracleObjectFields(raw)
	if fields == nil {
		return false
	}
	set := make(map[string]struct{}, len(known))
	for _, k := range known {
		set[k] = struct{}{}
	}
	for k := range fields {
		if _, ok := set[k]; !ok {
			return false
		}
	}
	return true
}

func oracleArrayFirst(raw jsontext.Value) (jsontext.Value, bool) {
	var elems []jsontext.Value
	if err := json.Unmarshal(raw, &elems); err != nil || len(elems) == 0 {
		return nil, false
	}
	return elems[0], true
}

func oracleArrayFirstHasKeys(raw jsontext.Value, keys ...string) bool {
	first, ok := oracleArrayFirst(raw)
	if !ok {
		return false
	}
	return oracleObjectHasKeys(first, keys...)
}

func oracleArrayFirstKeysKnown(raw jsontext.Value, known ...string) bool {
	first, ok := oracleArrayFirst(raw)
	if !ok {
		return false
	}
	return oracleObjectKeysKnown(first, known...)
}

// wellFormedInputs is the set of syntactically valid JSON values the predicates
// can actually receive in production: dec.ReadValue() rejects malformed input
// before any dispatch predicate runs, so the scanner is only contracted to
// agree with the map oracle on valid JSON. These cover corpus payloads, the
// dispatch collision shapes, and valid adversarial edge cases (escaped keys,
// structural characters inside strings, deep nesting, non-objects).
func wellFormedInputs(tb testing.TB) [][]byte {
	tb.Helper()
	corpus := loadCorpusInputs(tb)
	in := make([][]byte, 0, len(staticWellFormed)+len(corpus))
	in = append(in, staticWellFormed...)
	in = append(in, corpus...)
	return in
}

// staticWellFormed holds the hand-written valid JSON shapes: dispatch collision
// cases and valid adversarial edge cases (escaped keys, structural characters
// inside strings, deep nesting, non-objects).
var staticWellFormed = [][]byte{
	[]byte(`{"uri":"file:///a","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}}}`),
	[]byte(`{"kind":"create","uri":"file:///a.go"}`),
	[]byte(`{"title":"x","command":"do.it"}`),
	[]byte(`{"title":"Fix","command":{"title":"y","command":"do.it"}}`),
	[]byte(`{"language":"go"}`),
	[]byte(`{"scheme":"file"}`),
	[]byte(`{"pattern":"**/*.go"}`),
	[]byte(`{"notebook":"jupyter-notebook","language":"python"}`),
	[]byte(`{"notebookSelector":[],"id":"reg-1"}`),
	// brace, comma and colon inside a string value must not confuse skipping.
	[]byte(`{"a":"}{,:","b":"[]","c":1}`),
	// escaped quote inside a string value.
	[]byte(`{"a":"he said \"hi\"","b":2}`),
	// escaped backslash immediately before the closing quote.
	[]byte(`{"a":"trailing\\","b":2}`),
	// JSON-escaped object KEYS (valid JSON): the scanner compares raw key bytes
	// against literals, so it must decode these to match. Each \uXXXX is a real
	// escape in the JSON (a backtick raw string keeps the backslash literal).
	[]byte(`{"\u0075ri":"file:///a","\u0072ange":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}}}`),
	[]byte(`{"\u006bind":"create","uri":"file:///a.go"}`),
	[]byte(`{"comma\u006ed":"do.it","title":"x"}`),
	// escaped discriminator VALUE: "cr\u0065ate" decodes to "create".
	[]byte(`{"kind":"cr\u0065ate","uri":"file:///a.go"}`),
	// nested objects/arrays.
	[]byte(`{"outer":{"inner":{"k":[1,2,{"x":3}]}},"flag":true}`),
	// whitespace everywhere.
	[]byte("  {  \"a\" : 1 , \"b\" : { } }  "),
	// empty object.
	[]byte(`{}`),
	// non-objects.
	[]byte(`[1,2,3]`),
	[]byte(`"string"`),
	[]byte(`42`),
	[]byte(`null`),
	[]byte(`true`),
	// arrays of objects (for arrayFirst*).
	[]byte(`[{"name":"a","kind":1,"location":{"uri":"file:///x"},"data":{"k":1}}]`),
	[]byte(`[{"name":"a","kind":1,"location":{"uri":"file:///x","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}}}}]`),
	[]byte(`[]`),
}

// malformedInputs are syntactically invalid values. They never reach the
// predicates in production (ReadValue rejects them first), so the scanner is not
// contracted to match the oracle on them; the robustness test only requires the
// scanner not to panic.
func malformedInputs() [][]byte {
	return [][]byte{
		[]byte(`{"a":`),
		[]byte(`{"a"`),
		[]byte(`{`),
		[]byte(`{"`), // fuzz regression: unterminated key right after the brace
		[]byte(`{"a":1,`),
		[]byte(`{"a":"unterminated`),
		[]byte(`{"key\`), // unterminated with a trailing escape
		[]byte(``),
		[]byte(`   `),
	}
}

// keySets are representative required/known key tuples drawn from the generated
// dispatch code, used to exercise the predicates across realistic arguments.
var keySets = [][]string{
	{},
	{"uri"},
	{"uri", "range"},
	{"kind"},
	{"command"},
	{"title", "command"},
	{"language", "scheme", "pattern"},
	{"name", "kind", "location", "data"},
	{"notebook", "language"},
	{"start", "end"},
}

// TestScannerMatchesMapOracle is the differential guardrail: for every
// well-formed input and every key set, the production scanner must return
// exactly what the retained map-based oracle returns.
func TestScannerMatchesMapOracle(t *testing.T) {
	for _, raw := range wellFormedInputs(t) {
		raw := jsontext.Value(raw)

		if g, w := mustKind(objectKind(raw)), mustKind(oracleObjectKind(raw)); g != w {
			t.Errorf("objectKind(%s) = %q, oracle = %q", raw, g, w)
		}

		for _, ks := range keySets {
			if g, w := objectHasKeys(raw, ks...), oracleObjectHasKeys(raw, ks...); g != w {
				t.Errorf("objectHasKeys(%s, %v) = %v, oracle = %v", raw, ks, g, w)
			}
			if g, w := objectKeysKnown(raw, ks...), oracleObjectKeysKnown(raw, ks...); g != w {
				t.Errorf("objectKeysKnown(%s, %v) = %v, oracle = %v", raw, ks, g, w)
			}
			if g, w := arrayFirstHasKeys(raw, ks...), oracleArrayFirstHasKeys(raw, ks...); g != w {
				t.Errorf("arrayFirstHasKeys(%s, %v) = %v, oracle = %v", raw, ks, g, w)
			}
			if g, w := arrayFirstKeysKnown(raw, ks...), oracleArrayFirstKeysKnown(raw, ks...); g != w {
				t.Errorf("arrayFirstKeysKnown(%s, %v) = %v, oracle = %v", raw, ks, g, w)
			}

			// The fused guards the generated dispatch actually calls must equal
			// the separate-predicate pair they replaced.
			fusedObj := objectHasAndKnownGuard(raw, ks, ks)
			wantObj := oracleObjectHasKeys(raw, ks...) && oracleObjectKeysKnown(raw, ks...)
			if fusedObj != wantObj {
				t.Errorf("objectHasAndKnownGuard(%s, %v) = %v, oracle = %v", raw, ks, fusedObj, wantObj)
			}
			fusedArr := arrayFirstHasAndKnown(raw, ks, ks)
			wantArr := oracleArrayFirstHasKeys(raw, ks...) && oracleArrayFirstKeysKnown(raw, ks...)
			if fusedArr != wantArr {
				t.Errorf("arrayFirstHasAndKnown(%s, %v) = %v, oracle = %v", raw, ks, fusedArr, wantArr)
			}
		}
	}
}

// TestScannerMalformedNoPanic asserts the scanner is robust on syntactically
// invalid input that cannot reach it in production: it must terminate without
// panicking. No oracle agreement is required here (the inputs are not valid
// JSON), only safety. A panic would crash the test; the explicit recover turns
// it into a readable failure naming the offending input.
func TestScannerMalformedNoPanic(t *testing.T) {
	for _, raw := range malformedInputs() {
		raw := jsontext.Value(raw)
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("panic on malformed input %q: %v", raw, r)
				}
			}()
			_, _ = objectKind(raw)
			for _, ks := range keySets {
				_ = objectHasKeys(raw, ks...)
				_ = objectKeysKnown(raw, ks...)
				_ = arrayFirstHasKeys(raw, ks...)
				_ = arrayFirstKeysKnown(raw, ks...)
			}
		}()
	}
}

// mustKind collapses (string, bool) into a comparable string so a present/absent
// difference is caught alongside a value difference.
func mustKind(s string, ok bool) string {
	if !ok {
		return "\x00absent"
	}
	return s
}
