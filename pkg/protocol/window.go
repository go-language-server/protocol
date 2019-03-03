// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

// ShowMessageParams params of ShowMessage Notification.
type ShowMessageParams struct {

	// Message is the actual message.
	Message string `json:"message"`

	// Type is the message type.
	Type MessageType `json:"type"`
}

// MessageType type of ShowMessageParams type.
type MessageType float64

const (
	// Error an error message.
	Error MessageType = 1
	// Warning a warning message.
	Warning MessageType = 2
	// Info an information message.
	Info MessageType = 3
	// Log a log message.
	Log MessageType = 4
)

// ShowMessageRequestParams params of ShowMessage Request.
type ShowMessageRequestParams struct {

	// Actions is the message action items to present.
	Actions []MessageActionItem `json:"actions"`

	// Message is the actual message
	Message string `json:"message"`

	// Type is the message type. See {@link MessageType}
	Type MessageType `json:"type"`
}

// MessageActionItem item of ShowMessageRequestParams action.
type MessageActionItem struct {

	// Title a short title like 'Retry', 'Open Log' etc.
	Title string `json:"title"`
}

// LogMessageParams params of LogMessage Notification.
type LogMessageParams struct {

	// Message is the actual message
	Message string `json:"message"`

	// Type is the message type. See {@link MessageType}
	Type MessageType `json:"type"`
}
