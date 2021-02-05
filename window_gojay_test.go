// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestShowMessageParams(t *testing.T) {
	testShowMessageParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestShowMessageRequestParams(t *testing.T) {
	testShowMessageRequestParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestMessageActionItem(t *testing.T) {
	testMessageActionItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestLogMessageParams(t *testing.T) {
	testLogMessageParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressCreateParams(t *testing.T) {
	testWorkDoneProgressCreateParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressCancelParams(t *testing.T) {
	testWorkDoneProgressCancelParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
