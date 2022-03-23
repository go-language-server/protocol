// SPDX-FileCopyrightText: 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"testing"

	"github.com/segmentio/encoding/json"
)

func TestRegistration(t *testing.T) {
	t.Parallel()

	testRegistration(t, json.Marshal, json.Unmarshal)
}

func TestRegistrationParams(t *testing.T) {
	t.Parallel()

	testRegistrationParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentRegistrationOptions(t *testing.T) {
	t.Parallel()

	testTextDocumentRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestUnregistration(t *testing.T) {
	t.Parallel()

	testUnregistration(t, json.Marshal, json.Unmarshal)
}

func TestUnregistrationParams(t *testing.T) {
	t.Parallel()

	testUnregistrationParams(t, json.Marshal, json.Unmarshal)
}
