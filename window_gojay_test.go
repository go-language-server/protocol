// SPDX-FileCopyrightText: Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
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
