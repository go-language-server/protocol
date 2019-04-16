// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strconv"
)

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

// String implements fmt.Stringer.
func (m MessageType) String() string {
	switch m {
	case Error:
		return "error"
	case Warning:
		return "warning"
	case Info:
		return "info"
	case Log:
		return "log"
	default:
		return strconv.FormatFloat(float64(m), 'f', -1, 64)
	}
}

// Enabled reports whether the level is enabled.
func (m MessageType) Enabled(level MessageType) bool {
	return m > level
}

// messageTypeMap map of MessageTypes.
var messageTypeMap = map[string]MessageType{
	"error":   Error,
	"warning": Warning,
	"info":    Info,
	"log":     Log,
}

// ToMessageType converts level to the MessageType.
func ToMessageType(level string) MessageType {
	mt, ok := messageTypeMap[level]
	if !ok {
		return MessageType(0) // unknown
	}

	return mt
}

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
