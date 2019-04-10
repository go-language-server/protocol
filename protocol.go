// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"context"

	"github.com/go-language-server/jsonrpc2"
)

const (
	// Version is the version of the language-server-protocol specification being implemented.
	Version = "3.14.0"
)

// DefaultBufferSize default message buffer size.
const DefaultBufferSize = 20

// DefaultCanceller returns the default canceler function.
func DefaultCanceller(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	conn.Notify(context.Background(), cancelRequest, &CancelParams{ID: *req.ID})
}
