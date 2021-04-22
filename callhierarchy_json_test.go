// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestCallHierarchy(t *testing.T) {
	testCallHierarchy(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyOptions(t *testing.T) {
	testCallHierarchyOptions(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyRegistrationOptions(t *testing.T) {
	testCallHierarchyRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyPrepareParams(t *testing.T) {
	testCallHierarchyPrepareParams(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyItem(t *testing.T) {
	testCallHierarchyItem(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyIncomingCallsParams(t *testing.T) {
	testCallHierarchyIncomingCallsParams(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyIncomingCall(t *testing.T) {
	testCallHierarchyIncomingCall(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyOutgoingCallsParams(t *testing.T) {
	testCallHierarchyOutgoingCallsParams(t, json.Marshal, json.Unmarshal)
}

func TestCallHierarchyOutgoingCall(t *testing.T) {
	testCallHierarchyOutgoingCall(t, json.Marshal, json.Unmarshal)
}
