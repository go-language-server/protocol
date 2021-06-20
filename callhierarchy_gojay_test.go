// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestCallHierarchy(t *testing.T) {
	t.Parallel()

	testCallHierarchy(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyOptions(t *testing.T) {
	t.Parallel()

	testCallHierarchyOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyRegistrationOptions(t *testing.T) {
	t.Parallel()

	testCallHierarchyRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyPrepareParams(t *testing.T) {
	t.Parallel()

	testCallHierarchyPrepareParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyItem(t *testing.T) {
	t.Parallel()

	testCallHierarchyItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyIncomingCallsParams(t *testing.T) {
	t.Parallel()

	testCallHierarchyIncomingCallsParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyIncomingCall(t *testing.T) {
	t.Parallel()

	testCallHierarchyIncomingCall(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyOutgoingCallsParams(t *testing.T) {
	t.Parallel()

	testCallHierarchyOutgoingCallsParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestCallHierarchyOutgoingCall(t *testing.T) {
	t.Parallel()

	testCallHierarchyOutgoingCall(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
