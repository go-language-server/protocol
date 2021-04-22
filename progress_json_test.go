// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestWorkDoneProgressBegin(t *testing.T) {
	testWorkDoneProgressBegin(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressReport(t *testing.T) {
	testWorkDoneProgressReport(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressEnd(t *testing.T) {
	testWorkDoneProgressEnd(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressParams(t *testing.T) {
	testWorkDoneProgressParams(t, json.Marshal, json.Unmarshal)
}

func TestPartialResultParams(t *testing.T) {
	testPartialResultParams(t, json.Marshal, json.Unmarshal)
}
