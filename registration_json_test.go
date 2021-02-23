// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestRegistration(t *testing.T) { testRegistration(t, json.Marshal, json.Unmarshal) }

func TestRegistrationParams(t *testing.T) {
	testRegistrationParams(t, json.Marshal, json.Unmarshal)
}

func TestTextDocumentRegistrationOptions(t *testing.T) {
	testTextDocumentRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestUnregistration(t *testing.T) { testUnregistration(t, json.Marshal, json.Unmarshal) }

func TestUnregistrationParams(t *testing.T) {
	testUnregistrationParams(t, json.Marshal, json.Unmarshal)
}
