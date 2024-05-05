// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// MessageDirection indicates in which direction a message is sent in the protocol.
type MessageDirection string

const (
	MessageDirectionClientToServer MessageDirection = "clientToServer"
	MessageDirectionServerToClient MessageDirection = "serverToClient"
	MessageDirectionBidirectional  MessageDirection = "both"
)
