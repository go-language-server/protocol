// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestShowMessageParams(t *testing.T) { testShowMessageParams(t, json.Marshal, json.Unmarshal) }

func TestShowMessageRequestParams(t *testing.T) {
	testShowMessageRequestParams(t, json.Marshal, json.Unmarshal)
}

func TestMessageActionItem(t *testing.T) {
	testMessageActionItem(t, json.Marshal, json.Unmarshal)
}

func TestLogMessageParams(t *testing.T) {
	testLogMessageParams(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressCreateParams(t *testing.T) {
	testWorkDoneProgressCreateParams(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressCancelParams(t *testing.T) {
	testWorkDoneProgressCancelParams(t, json.Marshal, json.Unmarshal)
}
