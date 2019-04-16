// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

const emptyData = `{}`

type marshalFunc func(v interface{}) ([]byte, error)

type unmarshalFunc func(data []byte, v interface{}) error
