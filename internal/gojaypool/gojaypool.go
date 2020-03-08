// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package gojaypool

import (
	"io"
	_ "unsafe" // required for go:linkname

	"github.com/francoispqt/gojay"
)

//go:linkname borrowDecoder github.com/francoispqt/gojay.borrowDecoder
func borrowDecoder(r io.Reader, bufSize int) *gojay.Decoder

// BorrowSizedDecoder borrows a Decoder from the pool.
// It takes an io.Reader implementation as data input.
//
// In order to benefit from the pool, a borrowed decoder must be released after usage.
func BorrowSizedDecoder(r io.Reader, bufSize int) *gojay.Decoder {
	return borrowDecoder(r, bufSize)
}
