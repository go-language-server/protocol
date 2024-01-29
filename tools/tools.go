// SPDX-FileCopyrightText: 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build tools
// +build tools

package tools

// tools we use during development.
import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/incu6us/goimports-reviser/v3"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
)
