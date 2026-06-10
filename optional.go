// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "github.com/go-json-experiment/json"

// Optional wraps a non-nullable optional LSP property, preserving whether a
// zero value such as "" or false was present on the wire. A JSON null clears the
// value to match the legacy pointer representation, where null and absent both
// decoded to nil and were omitted on marshal.
type Optional[T any] struct {
	set   bool
	value T
}

// NewOptional returns an Optional holding v.
func NewOptional[T any](v T) Optional[T] { return Optional[T]{set: true, value: v} }

// IsZero reports whether the value is absent. It drives the ",omitzero" tag so
// an unset Optional is omitted entirely.
func (o Optional[T]) IsZero() bool { return !o.set }

// Get returns the wrapped value and whether it is present.
func (o Optional[T]) Get() (T, bool) { return o.value, o.set }

// Set marks the Optional present with v.
func (o *Optional[T]) Set(v T) {
	o.set = true
	o.value = v
}

// Clear marks the Optional absent and clears the stored value.
func (o *Optional[T]) Clear() {
	var zero T
	o.set = false
	o.value = zero
}

// MarshalJSON implements json.Marshaler.
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.set {
		return []byte("null"), nil
	}
	return json.Marshal(o.value)
}

// UnmarshalJSON implements json.Unmarshaler. It decodes via decodeWith so that
// nested union values continue to use their generated discriminating decoders.
func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		o.Clear()
		return nil
	}
	o.set = true
	return decodeWith(b, &o.value)
}
