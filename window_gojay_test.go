// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestShowMessageParams(t *testing.T) {
	t.Parallel()

	testShowMessageParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestShowMessageRequestParams(t *testing.T) {
	t.Parallel()

	testShowMessageRequestParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestMessageActionItem(t *testing.T) {
	t.Parallel()

	testMessageActionItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestLogMessageParams(t *testing.T) {
	t.Parallel()

	testLogMessageParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressCreateParams(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressCreateParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressCancelParams(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressCancelParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
