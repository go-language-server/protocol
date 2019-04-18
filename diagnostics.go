// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

// PublishDiagnosticsParams represents a params of PublishDiagnostics Notification.
type PublishDiagnosticsParams struct {
	// URI is the URI for which diagnostic information is reported.
	URI DocumentURI `json:"uri"`

	// Diagnostics an array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
}
