// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"testing"

	"github.com/go-json-experiment/json/jsontext"
)

// FuzzUnionDispatchOracle fuzzes the union-dispatch predicates against the
// retained map-based oracle. For any input that is valid JSON, the scanner must
// agree with the oracle; for invalid JSON it must not panic. This hardens the
// zero-alloc scanner rewrite against inputs the static corpus does not cover.
func FuzzUnionDispatchOracle(f *testing.F) {
	for _, seed := range wellFormedInputs(f) {
		f.Add(seed)
	}
	for _, seed := range malformedInputs() {
		f.Add(seed)
	}

	keys := []string{"kind", "uri", "range", "command", "title", "language"}

	f.Fuzz(func(t *testing.T, data []byte) {
		raw := jsontext.Value(data)
		valid := raw.IsValid()

		// objectKind must agree with the oracle on valid JSON.
		gk, gok := objectKind(raw)
		ok2, ook := oracleObjectKind(raw)
		if valid && (gok != ook || (gok && gk != ok2)) {
			t.Fatalf("objectKind(%s): scanner=(%q,%v) oracle=(%q,%v)", raw, gk, gok, ok2, ook)
		}

		// Try several key subsets.
		for n := 0; n <= len(keys); n++ {
			ks := keys[:n]
			if g, w := objectHasKeys(raw, ks...), oracleObjectHasKeys(raw, ks...); valid && g != w {
				t.Fatalf("objectHasKeys(%s,%v): scanner=%v oracle=%v", raw, ks, g, w)
			}
			if g, w := objectKeysKnown(raw, ks...), oracleObjectKeysKnown(raw, ks...); valid && g != w {
				t.Fatalf("objectKeysKnown(%s,%v): scanner=%v oracle=%v", raw, ks, g, w)
			}
			if g, w := arrayFirstHasKeys(raw, ks...), oracleArrayFirstHasKeys(raw, ks...); valid && g != w {
				t.Fatalf("arrayFirstHasKeys(%s,%v): scanner=%v oracle=%v", raw, ks, g, w)
			}
			if g, w := arrayFirstKeysKnown(raw, ks...), oracleArrayFirstKeysKnown(raw, ks...); valid && g != w {
				t.Fatalf("arrayFirstKeysKnown(%s,%v): scanner=%v oracle=%v", raw, ks, g, w)
			}

			// The fused guards are what the generated dispatch actually calls
			// (objectHasAndKnownGuard at 76 sites, arrayFirstHasAndKnown). Prove
			// them equivalent to the separate objectHasKeys && objectKeysKnown
			// pair the generator previously emitted, directly under the fuzzer.
			fusedObj := objectHasAndKnownGuard(raw, ks, ks)
			wantObj := oracleObjectHasKeys(raw, ks...) && oracleObjectKeysKnown(raw, ks...)
			if valid && fusedObj != wantObj {
				t.Fatalf("objectHasAndKnownGuard(%s,%v): scanner=%v oracle=%v", raw, ks, fusedObj, wantObj)
			}
			fusedArr := arrayFirstHasAndKnown(raw, ks, ks)
			wantArr := oracleArrayFirstHasKeys(raw, ks...) && oracleArrayFirstKeysKnown(raw, ks...)
			if valid && fusedArr != wantArr {
				t.Fatalf("arrayFirstHasAndKnown(%s,%v): scanner=%v oracle=%v", raw, ks, fusedArr, wantArr)
			}
		}
	})
}

// FuzzRoundTrip fuzzes decode->encode->decode stability for the union-bearing
// result types. A payload that decodes successfully must re-encode and decode
// again to an equal canonical form, proving the dispatch rewrite does not lose
// or corrupt data.
func FuzzRoundTrip(f *testing.F) {
	for _, seed := range wellFormedInputs(f) {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		var first CommandOrCodeAction
		if err := Unmarshal(data, &first); err != nil {
			return // not a valid CommandOrCodeAction; nothing to assert
		}
		out, err := Marshal(first)
		if err != nil {
			t.Fatalf("marshal after decode: %v", err)
		}
		var second CommandOrCodeAction
		if err := Unmarshal(out, &second); err != nil {
			t.Fatalf("re-decode of own output failed: %v\noutput=%s", err, out)
		}
		out2, err := Marshal(second)
		if err != nil {
			t.Fatalf("marshal second: %v", err)
		}
		if !jsontext.Value(out).IsValid() {
			t.Fatalf("re-encoded output is not valid JSON: %s", out)
		}
		if !bytes.Equal(canonicalize(t, out), canonicalize(t, out2)) {
			t.Fatalf("round-trip not stable:\n first=%s\nsecond=%s", out, out2)
		}
	})
}

// canonicalize returns the canonical JSON form of b for order-independent
// comparison.
func canonicalize(t *testing.T, b []byte) []byte {
	t.Helper()
	v := jsontext.Value(append([]byte(nil), b...))
	if err := v.Canonicalize(); err != nil {
		t.Fatalf("canonicalize %s: %v", b, err)
	}
	return v
}
