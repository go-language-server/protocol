// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestWorkDoneProgressBegin(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressBegin(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressReport(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressReport(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressEnd(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressEnd(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressParams(t *testing.T) {
	t.Parallel()

	testWorkDoneProgressParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestPartialResultParams(t *testing.T) {
	t.Parallel()

	testPartialResultParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
