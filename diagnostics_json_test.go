// SPDX-FileCopyrightText: Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestDiagnostic(t *testing.T) {
	testDiagnostic(t, json.Marshal, json.Unmarshal)
}

func TestDiagnosticRelatedInformation(t *testing.T) {
	testDiagnosticRelatedInformation(t, json.Marshal, json.Unmarshal)
}

func TestPublishDiagnosticsParams(t *testing.T) {
	testPublishDiagnosticsParams(t, json.Marshal, json.Unmarshal)
}
