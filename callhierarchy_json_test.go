// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"testing"

	"github.com/segmentio/encoding/json"
)

func TestCallHierarchy(t *testing.T) {
	t.Parallel()

	testCallHierarchy(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyOptions(t *testing.T) {
	t.Parallel()

	testCallHierarchyOptions(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyRegistrationOptions(t *testing.T) {
	t.Parallel()

	testCallHierarchyRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyPrepareParams(t *testing.T) {
	t.Parallel()

	testCallHierarchyPrepareParams(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyItem(t *testing.T) {
	t.Parallel()

	testCallHierarchyItem(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyIncomingCallsParams(t *testing.T) {
	t.Parallel()

	testCallHierarchyIncomingCallsParams(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyIncomingCall(t *testing.T) {
	t.Parallel()

	testCallHierarchyIncomingCall(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyOutgoingCallsParams(t *testing.T) {
	t.Parallel()

	testCallHierarchyOutgoingCallsParams(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyOutgoingCall(t *testing.T) {
	t.Parallel()

	testCallHierarchyOutgoingCall(t, json.Marshal, json.Unmarshal)
}
