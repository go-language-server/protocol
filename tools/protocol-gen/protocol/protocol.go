// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package protocol

// Protocol represents the LSP protocol.
type Protocol struct {
	// The enumerations.
	Enumerations []*Enumeration

	// Additional meta data.
	MetaData *MetaData

	// The notifications.
	Notifications []*Notification

	// The client -> server notifications.
	ClientToServerNotifications []*Notification

	// The server -> client notifications.
	ServerToClientNotifications []*Notification

	// The client -> server and server -> client notifications.
	BidirectionalNotifications []*Notification

	// The requests.
	Requests []*Request

	// The client -> server requests.
	ClientToServerRequests []*Request

	// The server -> client requests.
	ServerToClientRequests []*Request

	// The client -> server and server -> client requests.
	BidirectionalRequests []*Request

	// The structures.
	Structures []*Structure

	// The type aliases.
	TypeAliases []*TypeAlias
}

// MetaData is the protocol version.
type MetaData struct {
	Version string
}
