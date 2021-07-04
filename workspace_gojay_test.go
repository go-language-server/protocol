// SPDX-FileCopyrightText: 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build gojay
// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestWorkspaceFolder(t *testing.T) {
	t.Parallel()

	testWorkspaceFolder(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeWorkspaceFoldersParams(t *testing.T) {
	t.Parallel()

	testDidChangeWorkspaceFoldersParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceFoldersChangeEvent(t *testing.T) {
	t.Parallel()

	testWorkspaceFoldersChangeEvent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeConfigurationParams(t *testing.T) {
	t.Parallel()

	testDidChangeConfigurationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestConfigurationParams(t *testing.T) {
	t.Parallel()

	testConfigurationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestConfigurationItem(t *testing.T) {
	t.Parallel()

	testConfigurationItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeWatchedFilesParams(t *testing.T) {
	t.Parallel()

	testDidChangeWatchedFilesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFileEvent(t *testing.T) {
	t.Parallel()

	testFileEvent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeWatchedFilesRegistrationOptions(t *testing.T) {
	t.Parallel()

	testDidChangeWatchedFilesRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceSymbolParams(t *testing.T) {
	t.Parallel()

	testWorkspaceSymbolParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestExecuteCommandParams(t *testing.T) {
	t.Parallel()

	testExecuteCommandParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestExecuteCommandRegistrationOptions(t *testing.T) {
	t.Parallel()

	testExecuteCommandRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestApplyWorkspaceEditParams(t *testing.T) {
	t.Parallel()

	testApplyWorkspaceEditParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestApplyWorkspaceEditResponse(t *testing.T) {
	t.Parallel()

	testApplyWorkspaceEditResponse(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
