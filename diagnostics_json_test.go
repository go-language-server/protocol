// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"testing"

	"github.com/segmentio/encoding/json"
)

func TestDiagnostic(t *testing.T) {
	t.Parallel()

	testDiagnostic(t, json.Marshal, json.Unmarshal)
}

func TestDiagnosticRelatedInformation(t *testing.T) {
	t.Parallel()

	testDiagnosticRelatedInformation(t, json.Marshal, json.Unmarshal)
}

func TestPublishDiagnosticsParams(t *testing.T) {
	t.Parallel()

	testPublishDiagnosticsParams(t, json.Marshal, json.Unmarshal)
}
