// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestPublishDiagnosticsParams(t *testing.T) {
	testPublishDiagnosticsParams(t, json.Marshal, json.Unmarshal)
}
