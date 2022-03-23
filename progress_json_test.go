// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"testing"

	"github.com/segmentio/encoding/json"
)

func TestWorkDoneProgressBegin(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressBegin(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressReport(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressReport(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressEnd(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressEnd(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressParams(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressParams(t, json.Marshal, json.Unmarshal)
}

func TestPartialResultParams(t *testing.T) {
	t.Parallel()

	testPartialResultParams(t, json.Marshal, json.Unmarshal)
}
