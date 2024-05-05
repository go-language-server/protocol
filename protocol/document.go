// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"go.lsp.dev/uri"
)

// TextDocumentSaveReason represents reasons why a text document is saved.
type TextDocumentSaveReason uint32

const (
	// ManualTextDocumentSaveReason manually triggered, e.g. by the user pressing save, by starting debugging, or by an API call.
	ManualTextDocumentSaveReason TextDocumentSaveReason = 1

	// AfterDelayTextDocumentSaveReason automatic after a delay.
	AfterDelayTextDocumentSaveReason TextDocumentSaveReason = 2

	// FocusOutTextDocumentSaveReason when the editor lost focus.
	FocusOutTextDocumentSaveReason TextDocumentSaveReason = 3
)

// NotebookCellKind a notebook cell kind.
//
// @since 3.17.0
type NotebookCellKind uint32

const (
	// MarkupNotebookCellKind a markup-cell is formatted source that is used for display.
	MarkupNotebookCellKind NotebookCellKind = 1

	// CodeNotebookCellKind a code-cell is source code.
	CodeNotebookCellKind NotebookCellKind = 2
)

type ExecutionSummary struct {
	// ExecutionOrder a strict monotonically increasing value indicating the execution order of a cell inside a notebook.
	ExecutionOrder uint32 `json:"executionOrder"`

	// Success whether the execution was successful or not if known by the client.
	Success bool `json:"success,omitempty"`
}

// NotebookCell a notebook cell. A cell's document URI must be unique across ALL notebook cells and can therefore be
// used to uniquely identify a notebook cell or the cell's text document.
//
// @since 3.17.0
type NotebookCell struct {
	// Kind the cell's kind.
	//
	// @since 3.17.0
	Kind NotebookCellKind `json:"kind"`

	// Document the URI of the cell's text document content.
	//
	// @since 3.17.0
	Document DocumentURI `json:"document"`

	// Metadata additional metadata stored with the cell. Note: should always be an object literal (e.g. LSPObject).
	//
	// @since 3.17.0
	Metadata map[string]any `json:"metadata,omitempty"`

	// ExecutionSummary additional execution summary information if supported by the client.
	//
	// @since 3.17.0
	ExecutionSummary *ExecutionSummary `json:"executionSummary,omitempty"`
}

// NotebookDocument a notebook document.
//
// @since 3.17.0
type NotebookDocument struct {
	// URI the notebook document's uri.
	//
	// @since 3.17.0
	URI uri.URI `json:"uri"`

	// NotebookType the type of the notebook.
	//
	// @since 3.17.0
	NotebookType string `json:"notebookType"`

	// Version the version number of this document (it will increase after each change, including undo/redo).
	//
	// @since 3.17.0
	Version int32 `json:"version"`

	// Metadata additional metadata stored with the notebook document. Note: should always be an object literal (e.g. LSPObject).
	//
	// @since 3.17.0
	Metadata map[string]any `json:"metadata,omitempty"`

	// Cells the cells of a notebook.
	//
	// @since 3.17.0
	Cells []NotebookCell `json:"cells"`
}

// DidOpenNotebookDocumentParams the params sent in an open notebook document notification.
//
// @since 3.17.0
type DidOpenNotebookDocumentParams struct {
	// NotebookDocument the notebook document that got opened.
	//
	// @since 3.17.0
	NotebookDocument NotebookDocument `json:"notebookDocument"`

	// CellTextDocuments the text documents that represent the content of a notebook cell.
	//
	// @since 3.17.0
	CellTextDocuments []TextDocumentItem `json:"cellTextDocuments"`
}

// NotebookCellLanguage.
//
// @since 3.18.0 proposed
type NotebookCellLanguage struct {
	// @since 3.18.0 proposed
	Language string `json:"language"`
}

// NotebookDocumentFilterWithCells.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterWithCells struct {
	// Notebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
	//
	// @since 3.18.0 proposed
	Notebook NotebookDocumentFilterWithCellsNotebook `json:"notebook,omitempty"`

	// Cells the cells of the matching notebook to be synced.
	//
	// @since 3.18.0 proposed
	Cells []NotebookCellLanguage `json:"cells"`
}

// NotebookDocumentFilterWithNotebook.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterWithNotebook struct {
	// Notebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
	//
	// @since 3.18.0 proposed
	Notebook NotebookDocumentFilterWithNotebookNotebook `json:"notebook"`

	// Cells the cells of the matching notebook to be synced.
	//
	// @since 3.18.0 proposed
	Cells []NotebookCellLanguage `json:"cells,omitempty"`
}

// NotebookDocumentSyncOptions options specific to a notebook plus its cells to be synced to the server. If a selector provides a notebook document filter but no cell selector all cells of a matching notebook document will be synced. If a selector provides no notebook document filter but only a cell selector all notebook document
// that contain at least one matching cell will be synced.
//
// @since 3.17.0
type NotebookDocumentSyncOptions struct {
	// NotebookSelector the notebooks to be synced.
	//
	// @since 3.17.0
	NotebookSelector NotebookDocumentSyncOptionsNotebookSelector `json:"notebookSelector"`

	// Save whether save notification should be forwarded to the server. Will only be honored if mode === `notebook`.
	//
	// @since 3.17.0
	Save bool `json:"save,omitempty"`
}

// NotebookDocumentSyncRegistrationOptions registration options specific to a notebook.
//
// @since 3.17.0
type NotebookDocumentSyncRegistrationOptions struct {
	// extends
	NotebookDocumentSyncOptions
	// mixins
	StaticRegistrationOptions
}

// VersionedNotebookDocumentIdentifier a versioned notebook document identifier.
//
// @since 3.17.0
type VersionedNotebookDocumentIdentifier struct {
	// Version the version number of this notebook document.
	//
	// @since 3.17.0
	Version int32 `json:"version"`

	// URI the notebook document's uri.
	//
	// @since 3.17.0
	URI uri.URI `json:"uri"`
}

// NotebookCellArrayChange a change describing how to move a `NotebookCell` array from state S to S'.
//
// @since 3.17.0
type NotebookCellArrayChange struct {
	// Start the start oftest of the cell that changed.
	//
	// @since 3.17.0
	Start uint32 `json:"start"`

	// DeleteCount the deleted cells.
	//
	// @since 3.17.0
	DeleteCount uint32 `json:"deleteCount"`

	// Cells the new cells, if any.
	//
	// @since 3.17.0
	Cells []NotebookCell `json:"cells,omitempty"`
}

// NotebookDocumentCellChangeStructure structural changes to cells in a notebook document.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type NotebookDocumentCellChangeStructure struct {
	// Array the change to the cell array.
	//
	// @since 3.18.0 proposed
	Array NotebookCellArrayChange `json:"array"`

	// DidOpen additional opened cell text documents.
	//
	// @since 3.18.0 proposed
	DidOpen []TextDocumentItem `json:"didOpen,omitempty"`

	// DidClose additional closed cell text documents.
	//
	// @since 3.18.0 proposed
	DidClose []TextDocumentIdentifier `json:"didClose,omitempty"`
}

// NotebookDocumentCellContentChanges content changes to a cell in a notebook document.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type NotebookDocumentCellContentChanges struct {
	// @since 3.18.0 proposed
	Document VersionedTextDocumentIdentifier `json:"document"`

	// @since 3.18.0 proposed
	Changes []TextDocumentContentChangeEvent `json:"changes"`
}

// NotebookDocumentCellChanges cell changes to a notebook document.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type NotebookDocumentCellChanges struct {
	// Structure changes to the cell structure to add or remove cells.
	//
	// @since 3.18.0 proposed
	Structure *NotebookDocumentCellChangeStructure `json:"structure,omitempty"`

	// Data changes to notebook cells properties like its kind, execution summary or metadata.
	//
	// @since 3.18.0 proposed
	Data []NotebookCell `json:"data,omitempty"`

	// TextContent changes to the text content of notebook cells.
	//
	// @since 3.18.0 proposed
	TextContent []NotebookDocumentCellContentChanges `json:"textContent,omitempty"`
}

// NotebookDocumentChangeEvent a change event for a notebook document.
//
// @since 3.17.0
type NotebookDocumentChangeEvent struct {
	// Metadata the changed meta data if any. Note: should always be an object literal (e.g. LSPObject).
	//
	// @since 3.17.0
	Metadata map[string]any `json:"metadata,omitempty"`

	// Cells changes to cells.
	//
	// @since 3.17.0
	Cells *NotebookDocumentCellChanges `json:"cells,omitempty"`
}

// DidChangeNotebookDocumentParams the params sent in a change notebook document notification.
//
// @since 3.17.0
type DidChangeNotebookDocumentParams struct {
	// NotebookDocument the notebook document that did change. The version number points to the version after all provided changes have been applied. If only the text document content of a cell changes the notebook version doesn't necessarily have to change.
	//
	// @since 3.17.0
	NotebookDocument VersionedNotebookDocumentIdentifier `json:"notebookDocument"`

	// Change the actual changes to the notebook document. The changes describe single state changes to the notebook document. So if there are two changes c1 (at array index 0) and c2 (at array index 1) for a notebook in state S then c1 moves the notebook from S to S' and c2 from S' to S''. So c1 is computed on the state S and c2 is computed on the state S'. To mirror the content of a notebook using change events use the following approach: - start with the same initial content - apply the 'notebookDocument/didChange' notifications in the order you receive them. - apply the `NotebookChangeEvent`s in a single notification in the order you receive them.
	//
	// @since 3.17.0
	Change NotebookDocumentChangeEvent `json:"change"`
}

// NotebookDocumentIdentifier a literal to identify a notebook document in the client.
//
// @since 3.17.0
type NotebookDocumentIdentifier struct {
	// URI the notebook document's uri.
	//
	// @since 3.17.0
	URI uri.URI `json:"uri"`
}

// DidSaveNotebookDocumentParams the params sent in a save notebook document notification.
//
// @since 3.17.0
type DidSaveNotebookDocumentParams struct {
	// NotebookDocument the notebook document that got saved.
	//
	// @since 3.17.0
	NotebookDocument NotebookDocumentIdentifier `json:"notebookDocument"`
}

// DidCloseNotebookDocumentParams the params sent in a close notebook document notification.
//
// @since 3.17.0
type DidCloseNotebookDocumentParams struct {
	// NotebookDocument the notebook document that got closed.
	//
	// @since 3.17.0
	NotebookDocument NotebookDocumentIdentifier `json:"notebookDocument"`

	// CellTextDocuments the text documents that represent the content of a notebook cell that got closed.
	//
	// @since 3.17.0
	CellTextDocuments []TextDocumentIdentifier `json:"cellTextDocuments"`
}

// SaveOptions save options.
type SaveOptions struct {
	// IncludeText the client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
}

type TextDocumentSyncOptions struct {
	// OpenClose open and close notifications are sent to the server. If omitted open close notification should not be sent.
	OpenClose bool `json:"openClose,omitempty"`

	// Change change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
	Change TextDocumentSyncKind `json:"change,omitempty"`

	// WillSave if present will save notifications are sent to the server. If omitted the notification should not be
	// sent.
	WillSave bool `json:"willSave,omitempty"`

	// WillSaveWaitUntil if present will save wait until requests are sent to the server. If omitted the request should not be sent.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

	// Save if present save notifications are sent to the server. If omitted the notification should not be sent.
	Save TextDocumentSyncOptionsSave `json:"save,omitempty"`
}

// DidOpenTextDocumentParams the parameters sent in an open text document notification.
type DidOpenTextDocumentParams struct {
	// TextDocument the document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}

// DidChangeTextDocumentParams the change text document notification's parameters.
type DidChangeTextDocumentParams struct {
	// TextDocument the document that did change. The version number points to the version after all provided content changes have been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// ContentChanges the actual content changes. The content changes describe single state changes to the document. So if
	// there are two content changes c1 (at array index 0) and c2 (at array index 1) for a document in state S then c1 moves the document from S to S' and c2 from S' to S''. So c1 is computed on the state S and c2 is computed on the state S'. To mirror the content of a document using change events use the following approach: - start with the same initial content - apply the 'textDocument/didChange'
	// notifications in the order you receive them. - apply the `TextDocumentContentChangeEvent`s in a single notification in the order you receive them.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// TextDocumentChangeRegistrationOptions describe options to be used when registered for text document change events.
type TextDocumentChangeRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions

	// SyncKind how documents are synced to the server.
	SyncKind TextDocumentSyncKind `json:"syncKind"`
}

// DidCloseTextDocumentParams the parameters sent in a close text document notification.
type DidCloseTextDocumentParams struct {
	// TextDocument the document that was closed.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// DidSaveTextDocumentParams the parameters sent in a save text document notification.
type DidSaveTextDocumentParams struct {
	// TextDocument the document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Text optional the content when saved. Depends on the includeText value when the save notification was requested.
	Text string `json:"text,omitempty"`
}

// TextDocumentSaveRegistrationOptions save registration options.
type TextDocumentSaveRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	SaveOptions
}

// WillSaveTextDocumentParams the parameters sent in a will save text document notification.
type WillSaveTextDocumentParams struct {
	// TextDocument the document that will be saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Reason the 'TextDocumentSaveReason'.
	Reason TextDocumentSaveReason `json:"reason"`
}

// NotebookDocumentFilterNotebookType a notebook document filter where `notebookType` is required field.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterNotebookType struct {
	// NotebookType the type of the enclosing notebook.
	//
	// @since 3.18.0 proposed
	NotebookType string `json:"notebookType"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0 proposed
	Scheme string `json:"scheme,omitempty"`

	// Pattern a glob pattern.
	//
	// @since 3.18.0 proposed
	Pattern string `json:"pattern,omitempty"`
}

// NotebookDocumentFilterScheme a notebook document filter where `scheme` is required field.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterScheme struct {
	// NotebookType the type of the enclosing notebook.
	//
	// @since 3.18.0 proposed
	NotebookType string `json:"notebookType,omitempty"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0 proposed
	Scheme string `json:"scheme"`

	// Pattern a glob pattern.
	//
	// @since 3.18.0 proposed
	Pattern string `json:"pattern,omitempty"`
}

// NotebookDocumentFilterPattern a notebook document filter where `pattern` is required field.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterPattern struct {
	// NotebookType the type of the enclosing notebook.
	//
	// @since 3.18.0 proposed
	NotebookType string `json:"notebookType,omitempty"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0 proposed
	Scheme string `json:"scheme,omitempty"`

	// Pattern a glob pattern.
	//
	// @since 3.18.0 proposed
	Pattern string `json:"pattern"`
}
