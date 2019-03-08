// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strconv"
)

var (
	nullUint64, _ = strconv.ParseUint("null", 10, 64)
)

// Uint64 returns the i pointers.
func Uint64(i uint64) *uint64 {
	return &i
}
