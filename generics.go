// Copyright 2025 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"fmt"
)

// Nil represents a empty sturct that is provides nil type.
type Nil struct{}

var (
	trueBytes  = []byte("true")
	falseBytes = []byte("false")
	nullBytes  = []byte("null")
)

// OneOf represents a JSON-unmarshals as either a T or a U.
type OneOf[T, U any] struct {
	// The underlying value
	value any
}

var (
	_ Marshaler   = (*OneOf[any, any])(nil)
	_ Unmarshaler = (*OneOf[any, any])(nil)
)

// Any returns the underlying value.
func (o *OneOf[T, U]) Any() any {
	if o == nil {
		return nil
	}
	return o.value
}

// MarshalJSON implements [Marshaler].
func (o *OneOf[T, U]) MarshalJSON() ([]byte, error) {
	if o == nil {
		return nullBytes, nil
	}
	return marshal(o.value)
}

// UnmarshalJSON implements [Unmarshaler].
func (o *OneOf[T, U]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		return nil
	}

	// Try to unmarshal as T
	var valueT T
	errT := unmarshal(data, &valueT)
	if errT == nil {
		o.value = valueT
		return nil
	}

	// Try to unmarshal as U
	var valueU U
	errU := unmarshal(data, &valueU)
	if errU == nil {
		o.value = valueU
		return nil
	}

	return fmt.Errorf("cannot unmarshal to either type: %v or %v", errT, errU)
}

// OneOf3 represents a JSON-unmarshals as either a T, U or V.
type OneOf3[T, U, V any] struct {
	// The underlying value
	value any
}

var (
	_ Marshaler   = (*OneOf3[any, any, any])(nil)
	_ Unmarshaler = (*OneOf3[any, any, any])(nil)
)

// Any returns the underlying value.
func (o *OneOf3[T, U, V]) Any() any {
	if o == nil {
		return nil
	}
	return o.value
}

// MarshalJSON implements [Marshaler].
func (o *OneOf3[T, U, V]) MarshalJSON() ([]byte, error) {
	if o == nil {
		return nullBytes, nil
	}
	return marshal(o.value)
}

// UnmarshalJSON implements [Unmarshaler].
func (o *OneOf3[T, U, V]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		return nil
	}

	// Try to unmarshal as T
	var valueT T
	errT := unmarshal(data, &valueT)
	if errT == nil {
		o.value = valueT
		return nil
	}

	// Try to unmarshal as U
	var valueU U
	errU := unmarshal(data, &valueU)
	if errU == nil {
		o.value = valueU
		return nil
	}

	// Try to unmarshal as V
	var valueV V
	errV := unmarshal(data, &valueV)
	if errV == nil {
		o.value = valueV
		return nil
	}

	return fmt.Errorf("cannot unmarshal to either type: %v or %v or %v", errT, errU, errV)
}

// OneOf4 represents a JSON-unmarshals as either a T, U, V or V.
type OneOf4[T, U, V, Y any] struct {
	// The underlying value
	value any
}

var (
	_ Marshaler   = (*OneOf4[any, any, any, any])(nil)
	_ Unmarshaler = (*OneOf4[any, any, any, any])(nil)
)

// Any returns the underlying value.
func (o *OneOf4[T, U, V, Y]) Any() any {
	if o == nil {
		return nil
	}
	return o.value
}

// MarshalJSON implements [Marshaler].
func (o *OneOf4[T, U, V, Y]) MarshalJSON() ([]byte, error) {
	if o == nil {
		return nullBytes, nil
	}
	return marshal(o.value)
}

// UnmarshalJSON implements [Unmarshaler].
func (o *OneOf4[T, U, V, Y]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		return nil
	}

	// Try to unmarshal as T
	var valueT T
	errT := unmarshal(data, &valueT)
	if errT == nil {
		o.value = valueT
		return nil
	}

	// Try to unmarshal as U
	var valueU U
	errU := unmarshal(data, &valueU)
	if errU == nil {
		o.value = valueU
		return nil
	}

	// Try to unmarshal as V
	var valueV V
	errV := unmarshal(data, &valueV)
	if errV == nil {
		o.value = valueV
		return nil
	}

	// Try to unmarshal as V
	var valueY Y
	errY := unmarshal(data, &valueY)
	if errY == nil {
		o.value = valueV
		return nil
	}

	return fmt.Errorf("cannot unmarshal to either type: %v or %v or %v or %v", errT, errU, errV, errY)
}
