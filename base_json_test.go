// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestCancelParams(t *testing.T) {
	testCancelParams(t, json.Marshal, json.Unmarshal)
}

func TestProgressParams(t *testing.T) {
	testProgressParams(t, json.Marshal, json.Unmarshal)
}
