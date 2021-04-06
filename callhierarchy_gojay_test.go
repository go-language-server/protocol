// Copyright 2021 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestCallHierarchy(t *testing.T) {
	testCallHierarchy(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyOptions(t *testing.T) {
	testCallHierarchyOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyRegistrationOptions(t *testing.T) {
	testCallHierarchyRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyPrepareParams(t *testing.T) {
	testCallHierarchyPrepareParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyItem(t *testing.T) {
	testCallHierarchyItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyIncomingCallsParams(t *testing.T) {
	testCallHierarchyIncomingCallsParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyIncomingCall(t *testing.T) {
	testCallHierarchyIncomingCall(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyOutgoingCallsParams(t *testing.T) {
	testCallHierarchyOutgoingCallsParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyOutgoingCall(t *testing.T) {
	testCallHierarchyOutgoingCall(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
