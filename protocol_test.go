// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

// marshalFunc helper function type of testing.
type marshalFunc func(v interface{}) ([]byte, error)

// unmarshalFunc helper function type of testing.
type unmarshalFunc func(data []byte, v interface{}) error
