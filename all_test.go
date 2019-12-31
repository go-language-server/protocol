// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import "go/build"

// marshalFunc helper function type of testing.
type marshalFunc func(v interface{}) ([]byte, error)

// unmarshalFunc helper function type of testing.
type unmarshalFunc func(data []byte, v interface{}) error

// isGoJay reports whether the testing with "gojay" build tag.
func isGoJay() bool {
	for _, tag := range build.Default.BuildTags {
		if tag == "gojay" {
			return true
		}
	}
	return false
}
