// Copyright 2021 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"encoding/json"
	"fmt"
)

// ProgressToken is the progress token provided by the client or server.
//
// @since 3.15.0.
type ProgressToken struct {
	name   string
	number int64
}

// compile time check whether the ProgressToken implements a fmt.Formatter, fmt.Stringer, json.Marshaler and json.Unmarshaler interfaces.
var (
	_ fmt.Formatter    = (*ProgressToken)(nil)
	_ fmt.Stringer     = (*ProgressToken)(nil)
	_ json.Marshaler   = (*ProgressToken)(nil)
	_ json.Unmarshaler = (*ProgressToken)(nil)
)

// NewProgressToken returns a new ProgressToken.
func NewProgressToken(s string) *ProgressToken {
	return &ProgressToken{name: s}
}

// NewNumberProgressToken returns a new number ProgressToken.
func NewNumberProgressToken(n int64) *ProgressToken {
	return &ProgressToken{number: n}
}

// Format writes the ProgressToken to the formatter.
//
// If the rune is q the representation is non ambiguous,
// string forms are quoted.
func (v ProgressToken) Format(f fmt.State, r rune) {
	const numF = `%d`
	strF := `%s`
	if r == 'q' {
		strF = `%q`
	}

	switch {
	case v.name != "":
		fmt.Fprintf(f, strF, v.name)
	default:
		fmt.Fprintf(f, numF, v.number)
	}
}

// String returns a string representation of the type.
func (v ProgressToken) String() string {
	return fmt.Sprint(v)
}

// MarshalJSON implements json.Marshaler.
func (v *ProgressToken) MarshalJSON() ([]byte, error) {
	if v.name != "" {
		return json.Marshal(v.name)
	}
	return json.Marshal(v.number)
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *ProgressToken) UnmarshalJSON(data []byte) error {
	*v = ProgressToken{}
	if err := json.Unmarshal(data, &v.number); err == nil {
		return nil
	}
	return json.Unmarshal(data, &v.name)
}

// ProgressParams params of Progress netification.
//
// @since 3.15.0.
type ProgressParams struct {
	// Token is the progress token provided by the client or server.
	Token ProgressToken `json:"token"`

	// Value is the progress data.
	Value interface{} `json:"value"`
}

// WorkDoneProgressKind kind of WorkDoneProgress.
//
// @since 3.15.0.
type WorkDoneProgressKind string

// list of WorkDoneProgressKind.
const (
	WorkDoneProgressKindBegin  WorkDoneProgressKind = "begin"
	WorkDoneProgressKindReport WorkDoneProgressKind = "report"
	WorkDoneProgressKindEnd    WorkDoneProgressKind = "end"
)

// WorkDoneProgressBegin is the to start progress reporting a "$/progress" notification.
//
// @since 3.15.0.
type WorkDoneProgressBegin struct {
	// Kind is the kind of WorkDoneProgressBegin.
	//
	// It must be WorkDoneProgressKindBegin.
	Kind WorkDoneProgressKind `json:"kind"`

	// Title mandatory title of the progress operation. Used to briefly inform about
	// the kind of operation being performed.
	//
	// Examples: "Indexing" or "Linking dependencies".
	Title string `json:"title"`

	// Cancellable controls if a cancel button should show to allow the user to cancel the
	// long running operation. Clients that don't support cancellation are allowed
	// to ignore the setting.
	Cancellable bool `json:"cancellable,omitempty"`

	// Message is optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message string `json:"message,omitempty"`

	// Percentage is optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule.
	Percentage float64 `json:"percentage,omitempty"`
}

// WorkDoneProgressReport is the reporting progress is done.
//
// @since 3.15.0.
type WorkDoneProgressReport struct {
	// Kind is the kind of WorkDoneProgressReport.
	//
	// It must be WorkDoneProgressKindReport.
	Kind WorkDoneProgressKind `json:"kind"`

	// Cancellable controls enablement state of a cancel button.
	//
	// Clients that don't support cancellation or don't support controlling the button's
	// enablement state are allowed to ignore the property.
	Cancellable bool `json:"cancellable,omitempty"`

	// Message is optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message string `json:"message,omitempty"`

	// Percentage is optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule.
	Percentage float64 `json:"percentage,omitempty"`
}

// WorkDoneProgressEnd is the signaling the end of a progress reporting is done.
//
// @since 3.15.0.
type WorkDoneProgressEnd struct {
	// Kind is the kind of WorkDoneProgressEnd.
	//
	// It must be WorkDoneProgressKindEnd.
	Kind WorkDoneProgressKind `json:"kind"`

	// Message is optional, a final message indicating to for example indicate the outcome
	// of the operation.
	Message string `json:"message,omitempty"`
}

// WorkDoneProgressParams is a parameter property of report work done progress.
//
// @since 3.15.0.
type WorkDoneProgressParams struct {
	// WorkDoneToken an optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}

// PartialResultParams is the parameter literal used to pass a partial result token.
//
// @since 3.15.0.
type PartialResultParams struct {
	// PartialResultToken an optional token that a server can use to report partial results
	// (for example, streaming) to the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
}
