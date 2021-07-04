// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestDiagnostic(t *testing.T) {
	t.Parallel()

	testDiagnostic(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDiagnosticRelatedInformation(t *testing.T) {
	t.Parallel()

	testDiagnosticRelatedInformation(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestPublishDiagnosticsParams(t *testing.T) {
	t.Parallel()

	testPublishDiagnosticsParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
