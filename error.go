// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/go-language-server/jsonrpc2"
)

// ErrorInvalidParams reports InvalidParams error.
func ErrorInvalidParams(format string, args ...interface{}) error {
	return jsonrpc2.Errorf(jsonrpc2.CodeInvalidParams, format, args...)
}
