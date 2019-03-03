// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

// DidOpenTextDocumentParams params of DidOpenTextDocument Notification.
type DidOpenTextDocumentParams struct {

	// TextDocument is the document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}

// DidChangeTextDocumentParams params of DidChangeTextDocument Notification.
type DidChangeTextDocumentParams struct {

	// TextDocument is the document that did change. The version number points
	// to the version after all provided content changes have
	// been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// ContentChanges is the actual content changes. The content changes describe single state changes
	// to the document. So if there are two content changes c1 and c2 for a document
	// in state S then c1 move the document to S' and c2 to S''.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// TextDocumentContentChangeEvent an event describing a change to a text document. If range and rangeLength are omitted
// the new text is considered to be the full content of the document.
type TextDocumentContentChangeEvent struct {

	// Range is the range of the document that changed.
	Range *Range `json:"range,omitempty"`

	// RangeLength is the length of the range that got replaced.
	RangeLength float64 `json:"rangeLength,omitempty"`

	// Text is the new text of the range/document.
	Text string `json:"text"`
}

// TextDocumentChangeRegistrationOptions describe options to be used when registering for text document change events.
type TextDocumentChangeRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// SyncKind how documents are synced to the server. See TextDocumentSyncKind.Full
	// and TextDocumentSyncKind.Incremental.
	SyncKind float64 `json:"syncKind"`
}

// WillSaveTextDocumentParams is the parameters send in a will save text document notification.
type WillSaveTextDocumentParams struct {

	// TextDocument is the document that will be saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Reason is the 'TextDocumentSaveReason'.
	Reason TextDocumentSaveReason `json:"reason,omitempty"`
}

// TextDocumentSaveReason represents reasons why a text document is saved.
type TextDocumentSaveReason float64

const (
	// Manual is the manually triggered, e.g. by the user pressing save, by starting debugging,
	// or by an API call.
	Manual TextDocumentSaveReason = 1

	// AfterDelay is the automatic after a delay.
	AfterDelay TextDocumentSaveReason = 2

	// FocusOut when the editor lost focus.
	FocusOut TextDocumentSaveReason = 3
)

// DidSaveTextDocumentParams params of DidSaveTextDocument Notification.
type DidSaveTextDocumentParams struct {

	// Text optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text string `json:"text,omitempty"`

	// TextDocument is the document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// TextDocumentSaveRegistrationOptions TextDocumentSave Registration options.
type TextDocumentSaveRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// IncludeText is the client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
}

// DidCloseTextDocumentParams params of DidCloseTextDocument Notification.
type DidCloseTextDocumentParams struct {

	// TextDocument the document that was closed.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
