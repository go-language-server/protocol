// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strconv"

	"go.lsp.dev/uri"
)

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

// TextDocument is a simple text document. Not to be implemented.
type TextDocument struct {
	// URI is the associated URI for this document. Most documents have the __file__-scheme, indicating that they
	// represent files on disk. However, some documents may have other schemes indicating that they are not
	// available on disk.
	//
	// @readonly
	URI uri.URI `json:"uri"`

	// LanguageID is the identifier of the language associated with this document.
	//
	// @readonly
	LanguageID string `json:"languageId"`

	// Version is the version number of this document (it will increase after each
	// change, including undo/redo).
	//
	// @readonly
	Version float64 `json:"version"`

	// LineCount is the number of lines in this document.
	//
	// @readonly
	LineCount float64 `json:"lineCount"`
}

// TextDocumentChangeEvent Event to signal changes to a simple text document.
type TextDocumentChangeEvent struct {
	// Document is the document that has changed.
	Document TextDocument `json:"document"`
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

// String implements fmt.Stringer.
func (t TextDocumentSaveReason) String() string {
	switch t {
	case Manual:
		return "Manual"
	case AfterDelay:
		return "AfterDelay"
	case FocusOut:
		return "FocusOut"
	default:
		return strconv.FormatFloat(float64(t), 'f', -10, 64)
	}
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

// DidSaveTextDocumentParams params of DidSaveTextDocument Notification.
type DidSaveTextDocumentParams struct {
	// Text optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text string `json:"text,omitempty"`

	// TextDocument is the document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// TextDocumentContentChangeEvent an event describing a change to a text document. If range and rangeLength are omitted
// the new text is considered to be the full content of the document.
type TextDocumentContentChangeEvent struct {
	// Range is the range of the document that changed.
	Range *Range `json:"range,omitempty"`

	// RangeLength is the length of the range that got replaced.
	RangeLength float64 `json:"rangeLength,omitempty"`

	// Text is the new text of the document.
	Text string `json:"text"`
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
