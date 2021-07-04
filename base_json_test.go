// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestCancelParams(t *testing.T) {
	t.Parallel()

	testCancelParams(t, json.Marshal, json.Unmarshal)
}

func TestProgressParams(t *testing.T) {
	t.Parallel()

	testProgressParams(t, json.Marshal, json.Unmarshal)
}
