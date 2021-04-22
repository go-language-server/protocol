// SPDX-FileCopyrightText: Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestRegistration(t *testing.T) { testRegistration(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestRegistrationParams(t *testing.T) {
	testRegistrationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestTextDocumentRegistrationOptions(t *testing.T) {
	testTextDocumentRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestUnregistration(t *testing.T) { testUnregistration(t, gojay.Marshal, gojay.Unsafe.Unmarshal) }

func TestUnregistrationParams(t *testing.T) {
	testUnregistrationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
