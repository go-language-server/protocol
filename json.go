// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"sync"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// wireOptions relaxes jsontext validation that the LSP wire contract does not
// require. Duplicate member names decode as last-wins instead of erroring
// (peers are trusted; the union dispatch scanner is presence-based and agrees),
// and strings carrying invalid UTF-8 are mangled to U+FFFD instead of
// rejecting the whole message. Both checks otherwise run per token on every
// hot encode/decode path (the duplicate tracking alone maintains a per-object
// name namespace).
var wireOptions = json.JoinOptions(
	jsontext.AllowDuplicateNames(true),
	jsontext.AllowInvalidUTF8(true),
)

// unmarshalOptions returns the canonical decode option set: relaxed wire
// validation plus the generated union dispatchers. It is built lazily because
// unionUnmarshalers is assigned in a generated init function (a package-level
// var initializer referencing it would either capture nil — var initializers
// run before init — or, if the generator emitted a var initializer instead,
// form the initialization cycle unionUnmarshalers → unmarshalX → decodeWith →
// unmarshalOptions → unionUnmarshalers).
var unmarshalOptions = sync.OnceValue(func() json.Options {
	return json.JoinOptions(wireOptions, json.WithUnmarshalers(unionUnmarshalers))
})

// Marshal encodes v as LSP-conformant JSON.
//
// Sealed-interface union values marshal as their dynamic concrete arm, so no
// custom marshalers are required.
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v, wireOptions)
}

// Unmarshal decodes LSP JSON into v, dispatching every generated union type to
// its discriminating decoder. Duplicate object keys decode as last-wins and
// invalid UTF-8 is replaced with U+FFFD (see wireOptions).
func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v, unmarshalOptions())
}

// decodeWith decodes raw into v with the union unmarshalers applied, so that
// generated union decoders propagate dispatch to nested union fields. The
// generated decoders call this instead of json.Unmarshal directly.
func decodeWith(raw []byte, v any) error {
	return json.Unmarshal(raw, v, unmarshalOptions())
}
