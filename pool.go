// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"sync"
)

var (
	// basic

	// PositionPool represents a pool the Position.
	PositionPool *sync.Pool

	// RangePool represents a pool the Range.
	RangePool *sync.Pool

	// LocationPool represents a pool the Location.
	LocationPool *sync.Pool

	// LocationLinkPool represents a pool the LocationLink.
	LocationLinkPool *sync.Pool

	// DiagnosticPool represents a pool the Diagnostic.
	DiagnosticPool *sync.Pool

	// DiagnosticRelatedInformationPool represents a pool the DiagnosticRelatedInformation.
	DiagnosticRelatedInformationPool *sync.Pool

	// CommandPool represents a pool the Command.
	CommandPool *sync.Pool

	// TextEditPool represents a pool the TextEdit.
	TextEditPool *sync.Pool

	// TextDocumentEditPool represents a pool the TextDocumentEdit.
	TextDocumentEditPool *sync.Pool

	// CreateFileOptionsPool represents a pool the CreateFileOptions.
	CreateFileOptionsPool *sync.Pool

	// CreateFilePool represents a pool the CreateFile.
	CreateFilePool *sync.Pool

	// RenameFileOptionsPool represents a pool the RenameFileOptions.
	RenameFileOptionsPool *sync.Pool

	// RenameFilePool represents a pool the RenameFile.
	RenameFilePool *sync.Pool

	// DeleteFileOptionsPool represents a pool the DeleteFileOptions.
	DeleteFileOptionsPool *sync.Pool

	// DeleteFilePool represents a pool the DeleteFile.
	DeleteFilePool *sync.Pool

	// WorkspaceEditPool represents a pool the WorkspaceEdit.
	WorkspaceEditPool *sync.Pool

	// TextDocumentIdentifierPool represents a pool the TextDocumentIdentifier.
	TextDocumentIdentifierPool *sync.Pool

	// TextDocumentItemPool represents a pool the TextDocumentItem.
	TextDocumentItemPool *sync.Pool

	// VersionedTextDocumentIdentifierPool represents a pool the VersionedTextDocumentIdentifier.
	VersionedTextDocumentIdentifierPool *sync.Pool

	// TextDocumentPositionParamsPool represents a pool the TextDocumentPositionParams.
	TextDocumentPositionParamsPool *sync.Pool

	// DocumentFilterPool represents a pool the DocumentFilter.
	DocumentFilterPool *sync.Pool

	// DocumentSelectorPool represents a pool the DocumentSelector.
	DocumentSelectorPool *sync.Pool

	// MarkupContentPool represents a pool the MarkupContent.
	MarkupContentPool *sync.Pool

	// diagnostics

	// PublishDiagnosticsParamsPool represents a pool the PublishDiagnosticsParams.
	PublishDiagnosticsParamsPool *sync.Pool
)

func init() {
	// basic
	PositionPool = &sync.Pool{
		New: func() interface{} {
			return &Position{}
		},
	}
	RangePool = &sync.Pool{
		New: func() interface{} {
			return &Range{}
		},
	}
	LocationPool = &sync.Pool{
		New: func() interface{} {
			return &Location{}
		},
	}
	DiagnosticPool = &sync.Pool{
		New: func() interface{} {
			return &Diagnostic{}
		},
	}
	LocationLinkPool = &sync.Pool{
		New: func() interface{} {
			return &LocationLink{}
		},
	}
	DiagnosticRelatedInformationPool = &sync.Pool{
		New: func() interface{} {
			return &DiagnosticRelatedInformation{}
		},
	}
	CommandPool = &sync.Pool{
		New: func() interface{} {
			return &Command{}
		},
	}
	TextEditPool = &sync.Pool{
		New: func() interface{} {
			return &TextEdit{}
		},
	}
	TextDocumentEditPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentEdit{}
		},
	}
	CreateFileOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &CreateFileOptions{}
		},
	}
	CreateFilePool = &sync.Pool{
		New: func() interface{} {
			return &CreateFile{}
		},
	}
	RenameFileOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &RenameFileOptions{}
		},
	}
	RenameFilePool = &sync.Pool{
		New: func() interface{} {
			return &RenameFile{}
		},
	}
	DeleteFileOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &DeleteFileOptions{}
		},
	}
	DeleteFilePool = &sync.Pool{
		New: func() interface{} {
			return &DeleteFile{}
		},
	}
	WorkspaceEditPool = &sync.Pool{
		New: func() interface{} {
			return &WorkspaceEdit{}
		},
	}
	TextDocumentIdentifierPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentIdentifier{}
		},
	}
	TextDocumentItemPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentItem{}
		},
	}
	VersionedTextDocumentIdentifierPool = &sync.Pool{
		New: func() interface{} {
			return &VersionedTextDocumentIdentifier{}
		},
	}
	TextDocumentPositionParamsPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentPositionParams{}
		},
	}
	DocumentFilterPool = &sync.Pool{
		New: func() interface{} {
			return &DocumentFilter{}
		},
	}
	DocumentSelectorPool = &sync.Pool{
		New: func() interface{} {
			return &DocumentSelector{}
		},
	}
	MarkupContentPool = &sync.Pool{
		New: func() interface{} {
			return &MarkupContent{}
		},
	}

	// diagnostics
	PublishDiagnosticsParamsPool = &sync.Pool{
		New: func() interface{} {
			return &PublishDiagnosticsParams{}
		},
	}
}
