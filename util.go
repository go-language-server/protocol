// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

// NewVersion returns the uint64 pointer converted i.
func NewVersion(i int32) *int32 {
	return &i
}
