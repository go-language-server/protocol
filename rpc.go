// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"github.com/go-language-server/jsonrpc2"
)

// DefaultBufferSize default message buffer size.
const DefaultBufferSize = 20

// Canceller returns the default canceler function.
func Canceller(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	conn.Notify(context.Background(), cancelRequest, &CancelParams{ID: *req.ID})
}
