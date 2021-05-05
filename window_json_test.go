// SPDX-FileCopyrightText: Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestShowMessageParams(t *testing.T) {
	t.Parallel()

	testShowMessageParams(t, json.Marshal, json.Unmarshal)
}

func TestShowMessageRequestParams(t *testing.T) {
	t.Parallel()

	testShowMessageRequestParams(t, json.Marshal, json.Unmarshal)
}

func TestMessageActionItem(t *testing.T) {
	t.Parallel()

	testMessageActionItem(t, json.Marshal, json.Unmarshal)
}

func TestLogMessageParams(t *testing.T) {
	t.Parallel()

	testLogMessageParams(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressCreateParams(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressCreateParams(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressCancelParams(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressCancelParams(t, json.Marshal, json.Unmarshal)
}
