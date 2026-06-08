// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "github.com/go-json-experiment/json"

// Marshal encodes v as LSP-conformant JSON.
//
// Sealed-interface union values marshal as their dynamic concrete arm, so no
// custom marshalers are required.
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal decodes LSP JSON into v, dispatching every generated union type to
// its discriminating decoder.
func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v, json.WithUnmarshalers(unionUnmarshalers))
}

// decodeWith decodes raw into v with the union unmarshalers applied, so that
// generated union decoders propagate dispatch to nested union fields. The
// generated decoders call this instead of json.Unmarshal directly.
func decodeWith(raw []byte, v any) error {
	return json.Unmarshal(raw, v, json.WithUnmarshalers(unionUnmarshalers))
}
