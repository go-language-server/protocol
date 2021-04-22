// SPDX-FileCopyrightText: Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestDiagnostic(t *testing.T) {
	testDiagnostic(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDiagnosticRelatedInformation(t *testing.T) {
	testDiagnosticRelatedInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestPublishDiagnosticsParams(t *testing.T) {
	testPublishDiagnosticsParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
