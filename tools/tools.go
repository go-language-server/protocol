// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/zchee/goimportz/cmd/goimportz"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
)
