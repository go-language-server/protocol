// SPDX-FileCopyrightText: 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestRegistration(t *testing.T) {
	t.Parallel()

	testRegistration(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestRegistrationParams(t *testing.T) {
	t.Parallel()

	testRegistrationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentRegistrationOptions(t *testing.T) {
	t.Parallel()

	testTextDocumentRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestUnregistration(t *testing.T) {
	t.Parallel()

	testUnregistration(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestUnregistrationParams(t *testing.T) {
	t.Parallel()

	testUnregistrationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
