// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

// PublishDiagnosticsParams represents a params of PublishDiagnostics Notification.
type PublishDiagnosticsParams struct {
	// URI is the URI for which diagnostic information is reported.
	URI DocumentURI `json:"uri"`

	// Version optional the version number of the document the diagnostics are published for.
	//
	// @since 3.15
	Version uint32 `json:"version,omitempty"`

	// Diagnostics an array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
}
