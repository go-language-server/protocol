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

func TestWorkspaceFolder(t *testing.T) {
	testWorkspaceFolder(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeWorkspaceFoldersParams(t *testing.T) {
	testDidChangeWorkspaceFoldersParams(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceFoldersChangeEvent(t *testing.T) {
	testWorkspaceFoldersChangeEvent(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeConfigurationParams(t *testing.T) {
	testDidChangeConfigurationParams(t, json.Marshal, json.Unmarshal)
}

func TestConfigurationParams(t *testing.T) {
	testConfigurationParams(t, json.Marshal, json.Unmarshal)
}

func TestConfigurationItem(t *testing.T) {
	testConfigurationItem(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeWatchedFilesParams(t *testing.T) {
	testDidChangeWatchedFilesParams(t, json.Marshal, json.Unmarshal)
}

func TestFileEvent(t *testing.T) {
	testFileEvent(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeWatchedFilesRegistrationOptions(t *testing.T) {
	testDidChangeWatchedFilesRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceSymbolParams(t *testing.T) {
	testWorkspaceSymbolParams(t, json.Marshal, json.Unmarshal)
}

func TestExecuteCommandParams(t *testing.T) {
	testExecuteCommandParams(t, json.Marshal, json.Unmarshal)
}

func TestExecuteCommandRegistrationOptions(t *testing.T) {
	testExecuteCommandRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestApplyWorkspaceEditParams(t *testing.T) {
	testApplyWorkspaceEditParams(t, json.Marshal, json.Unmarshal)
}

func TestApplyWorkspaceEditResponse(t *testing.T) {
	testApplyWorkspaceEditResponse(t, json.Marshal, json.Unmarshal)
}
