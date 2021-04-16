// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

// CancelParams params of cancelRequest.
type CancelParams struct {
	// ID is the request id to cancel.
	ID interface{} `json:"id"` // int32 | string
}
