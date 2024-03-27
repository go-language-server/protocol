// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Protocol represents the LSP protocol
type Protocol struct {
	// The enumerations
	Enumerations []*Enumeration
	// Additional meta data.
	MetaData *MetaData
	// The notifications
	Notifications []*Notification
	// The requests
	Requests []*Request
	// The structures
	Structures []*Structure
	// The type aliases
	TypeAliases []*TypeAlias

	// The client -> server requests
	ClientToServerRequests []*Request
	// The server -> client requests
	ServerToClientRequests []*Request
}

type MetaData struct {
	// The protocol version
	Version string
}
