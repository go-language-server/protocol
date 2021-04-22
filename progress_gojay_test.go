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
	testWorkDoneProgressBegin(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressReport(t *testing.T) {
	testWorkDoneProgressReport(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressEnd(t *testing.T) {
	testWorkDoneProgressEnd(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkDoneProgressParams(t *testing.T) {
	testWorkDoneProgressParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestPartialResultParams(t *testing.T) {
	testPartialResultParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
