// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// objectFields decodes a raw JSON object into its top-level members. It returns
// nil if raw is not an object.
func objectFields(raw jsontext.Value) map[string]jsontext.Value {
	var m map[string]jsontext.Value
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil
	}
	return m
}

// objectKind returns the value of a JSON object's "kind" string member, used as
// a discriminator when disambiguating union arms.
func objectKind(raw jsontext.Value) (string, bool) {
	v, ok := objectFields(raw)["kind"]
	if !ok {
		return "", false
	}
	var s string
	if err := json.Unmarshal(v, &s); err != nil {
		return "", false
	}
	return s, true
}

// objectHasKeys reports whether the JSON object contains every given top-level
// key, used to require an arm's mandatory fields before selecting it.
func objectHasKeys(raw jsontext.Value, keys ...string) bool {
	fields := objectFields(raw)
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

// arrayFirstHasKeys reports whether raw is a non-empty JSON array whose first
// element is an object containing every given key.
func arrayFirstHasKeys(raw jsontext.Value, keys ...string) bool {
	var elems []jsontext.Value
	if err := json.Unmarshal(raw, &elems); err != nil || len(elems) == 0 {
		return false
	}
	return objectHasKeys(elems[0], keys...)
}

// objectKeysKnown reports whether every top-level key of the JSON object raw is
// present in the given set of known field names. It is used to ensure a union
// arm is not selected when the payload carries a field that arm cannot hold.
func objectKeysKnown(raw jsontext.Value, known ...string) bool {
	fields := objectFields(raw)
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

// arrayFirstKeysKnown reports whether raw is a non-empty JSON array whose first
// element is an object all of whose keys are in the given known set.
func arrayFirstKeysKnown(raw jsontext.Value, known ...string) bool {
	var elems []jsontext.Value
	if err := json.Unmarshal(raw, &elems); err != nil || len(elems) == 0 {
		return false
	}
	return objectKeysKnown(elems[0], known...)
}
