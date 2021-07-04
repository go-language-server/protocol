// SPDX-FileCopyrightText: 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gojay
// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestWorkspaceFolder(t *testing.T) {
	t.Parallel()

	testWorkspaceFolder(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeWorkspaceFoldersParams(t *testing.T) {
	t.Parallel()

	testDidChangeWorkspaceFoldersParams(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceFoldersChangeEvent(t *testing.T) {
	t.Parallel()

	testWorkspaceFoldersChangeEvent(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeConfigurationParams(t *testing.T) {
	t.Parallel()

	testDidChangeConfigurationParams(t, json.Marshal, json.Unmarshal)
}

func TestConfigurationParams(t *testing.T) {
	t.Parallel()

	testConfigurationParams(t, json.Marshal, json.Unmarshal)
}

func TestConfigurationItem(t *testing.T) {
	t.Parallel()

	testConfigurationItem(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeWatchedFilesParams(t *testing.T) {
	t.Parallel()

	testDidChangeWatchedFilesParams(t, json.Marshal, json.Unmarshal)
}

func TestFileEvent(t *testing.T) {
	t.Parallel()

	testFileEvent(t, json.Marshal, json.Unmarshal)
}

func TestDidChangeWatchedFilesRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDidChangeWatchedFilesRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestWorkspaceSymbolParams(t *testing.T) {
	t.Parallel()

	testWorkspaceSymbolParams(t, json.Marshal, json.Unmarshal)
}

func TestExecuteCommandParams(t *testing.T) {
	t.Parallel()

	testExecuteCommandParams(t, json.Marshal, json.Unmarshal)
}

func TestExecuteCommandRegistrationOptions(t *testing.T) {
	t.Parallel()

	testExecuteCommandRegistrationOptions(t, json.Marshal, json.Unmarshal)
}

func TestApplyWorkspaceEditParams(t *testing.T) {
	t.Parallel()

	testApplyWorkspaceEditParams(t, json.Marshal, json.Unmarshal)
}

func TestApplyWorkspaceEditResponse(t *testing.T) {
	t.Parallel()

	testApplyWorkspaceEditResponse(t, json.Marshal, json.Unmarshal)
}
