// Copyright 2020 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"testing"

	"github.com/francoispqt/gojay"
)

func TestWorkspaceFolder(t *testing.T) {
	testWorkspaceFolder(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeWorkspaceFoldersParams(t *testing.T) {
	testDidChangeWorkspaceFoldersParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceFoldersChangeEvent(t *testing.T) {
	testWorkspaceFoldersChangeEvent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeConfigurationParams(t *testing.T) {
	testDidChangeConfigurationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestConfigurationParams(t *testing.T) {
	testConfigurationParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestConfigurationItem(t *testing.T) {
	testConfigurationItem(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeWatchedFilesParams(t *testing.T) {
	testDidChangeWatchedFilesParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestFileEvent(t *testing.T) {
	testFileEvent(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestDidChangeWatchedFilesRegistrationOptions(t *testing.T) {
	testDidChangeWatchedFilesRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestWorkspaceSymbolParams(t *testing.T) {
	testWorkspaceSymbolParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestExecuteCommandParams(t *testing.T) {
	testExecuteCommandParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestExecuteCommandRegistrationOptions(t *testing.T) {
	testExecuteCommandRegistrationOptions(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestApplyWorkspaceEditParams(t *testing.T) {
	testApplyWorkspaceEditParams(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}

func TestApplyWorkspaceEditResponse(t *testing.T) {
	testApplyWorkspaceEditResponse(t, gojay.Marshal, gojay.Unsafe.Unmarshal)
}
