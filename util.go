// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import "go.lsp.dev/uri"

// ToURI returns the new DocumentURI from s.
func ToURI(s string) uri.URI {
	return uri.File(s)
}

// Uint64Ptr converts i to uint64 pointer.
func Uint64Ptr(i uint64) *uint64 {
	return &i
}
