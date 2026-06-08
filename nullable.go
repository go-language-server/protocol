// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "github.com/go-json-experiment/json"

// Nullable wraps a value that is BOTH optional and JSON-nullable, distinguishing
// the three wire states the LSP specification assigns distinct meaning to:
//
//   - absent  — the zero Nullable; omitted on marshal via the ",omitzero" tag;
//   - null    — an explicit JSON null;
//   - value   — a present value.
//
// It is generated only for properties that are simultaneously optional and have
// a null arm (e.g. WorkspaceFoldersInitializeParams.workspaceFolders, where
// absent means "no workspace-folder support" and null means "supported, none
// open"). A plain pointer cannot represent all three states.
type Nullable[T any] struct {
	set   bool
	null  bool
	value T
}

// NewNull returns a Nullable holding an explicit JSON null.
func NewNull[T any]() Nullable[T] { return Nullable[T]{set: true, null: true} }

// NewNullable returns a Nullable holding v.
func NewNullable[T any](v T) Nullable[T] { return Nullable[T]{set: true, value: v} }

// IsZero reports whether the value is absent. It drives the ",omitzero" tag so
// an unset Nullable is omitted entirely.
func (n Nullable[T]) IsZero() bool { return !n.set }

// IsNull reports whether the value is present as an explicit JSON null.
func (n Nullable[T]) IsNull() bool { return n.set && n.null }

// Get returns the wrapped value and whether a non-null value is present.
func (n Nullable[T]) Get() (T, bool) { return n.value, n.set && !n.null }

// MarshalJSON implements json.Marshaler.
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.null {
		return []byte("null"), nil
	}
	return json.Marshal(n.value)
}

// UnmarshalJSON implements json.Unmarshaler. It decodes via decodeWith so that a
// nested union value dispatches to its discriminating decoder.
func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.set = true
	if string(b) == "null" {
		n.null = true
		var zero T
		n.value = zero
		return nil
	}
	n.null = false
	return decodeWith(b, &n.value)
}
