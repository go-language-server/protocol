// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import "go.lsp.dev/uri"

// ToURI returns the new DocumentURI from s.
func ToURI(s string) uri.URI {
	return uri.File(s)
}

// NewVersion returns the uint64 pointer converted i.
func NewVersion(i uint64) *uint64 {
	return &i
}
