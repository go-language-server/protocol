// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestCancelParams(t *testing.T) {
	t.Parallel()

	testCancelParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestProgressParams(t *testing.T) {
	t.Parallel()

	testProgressParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
