// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// TypeDecl is an interface used to classify protocol types represent a LSP type declaration.
type TypeDecl interface {
	isTypeDecl()
}
