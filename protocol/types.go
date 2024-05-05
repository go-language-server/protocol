// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

type DocumentURI string

type ClientMethod = string

type ServerMethod = string

// UnmarshalError indicates that a JSON value did not conform to
// one of the expected cases of an LSP union type.
type UnmarshalError struct {
	msg string
}

func (e UnmarshalError) Error() string {
	return e.msg
}
