// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"go.lsp.dev/uri"
)

// MessageType the message type.
type MessageType uint32

const (
	// ErrorMessageType an error message.
	ErrorMessageType MessageType = 1

	// WarningMessageType a warning message.
	WarningMessageType MessageType = 2

	// InfoMessageType an information message.
	InfoMessageType MessageType = 3

	// LogMessageType a log message.
	LogMessageType MessageType = 4

	// DebugMessageType a debug message.  3.18.0 @proposed.
	//
	// @since 3.18.0 proposed
	DebugMessageType MessageType = 5
)

type WorkDoneProgressCreateParams struct {
	// Token the token to be used to report progress.
	Token ProgressToken `json:"token"`
}

type WorkDoneProgressCancelParams struct {
	// Token the token to be used to report progress.
	Token ProgressToken `json:"token"`
}

// ShowDocumentParams params to show a resource in the UI.
//
// @since 3.16.0
type ShowDocumentParams struct {
	// URI the uri to show.
	//
	// @since 3.16.0
	URI uri.URI `json:"uri"`

	// External indicates to show the resource in an external program. To show, for example, `https://code.visualstudio.com/` in the default WEB browser set `external` to `true`.
	//
	// @since 3.16.0
	External bool `json:"external,omitempty"`

	// TakeFocus an optional property to indicate whether the editor showing the document should take focus or not. Clients might ignore this property if an external program is started.
	//
	// @since 3.16.0
	TakeFocus bool `json:"takeFocus,omitempty"`

	// Selection an optional selection range if the document is a text document. Clients might ignore the property if
	// an external program is started or the file is not a text file.
	//
	// @since 3.16.0
	Selection *Range `json:"selection,omitempty"`
}

// ShowDocumentResult the result of a showDocument request.
//
// @since 3.16.0
type ShowDocumentResult struct {
	// Success a boolean indicating if the show was successful.
	//
	// @since 3.16.0
	Success bool `json:"success"`
}

// ShowMessageParams the parameters of a notification message.
type ShowMessageParams struct {
	// Type the message type. See MessageType.
	Type MessageType `json:"type"`

	// Message the actual message.
	Message string `json:"message"`
}

type MessageActionItem struct {
	// Title a short title like 'Retry', 'Open Log' etc.
	Title string `json:"title"`
}

type ShowMessageRequestParams struct {
	// Type the message type. See MessageType.
	Type MessageType `json:"type"`

	// Message the actual message.
	Message string `json:"message"`

	// Actions the message action items to present.
	Actions []MessageActionItem `json:"actions,omitempty"`
}

// LogMessageParams the log message parameters.
type LogMessageParams struct {
	// Type the message type. See MessageType.
	Type MessageType `json:"type"`

	// Message the actual message.
	Message string `json:"message"`
}
