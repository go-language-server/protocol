// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import "sync"

// Pooler represents a poolable interface.
type Pooler interface {
	Reset()
}

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

	// general

	// InitializeParamsPool represents a pool the InitializeParams.
	InitializeParamsPool *sync.Pool

	// WorkspaceClientCapabilitiesWorkspaceEditPool represents a pool the WorkspaceClientCapabilitiesWorkspaceEdit.
	WorkspaceClientCapabilitiesWorkspaceEditPool *sync.Pool

	// WorkspaceClientCapabilitiesDidChangeConfigurationPool represents a pool the WorkspaceClientCapabilitiesDidChangeConfiguration.
	WorkspaceClientCapabilitiesDidChangeConfigurationPool *sync.Pool

	// WorkspaceClientCapabilitiesDidChangeWatchedFilesPool represents a pool the WorkspaceClientCapabilitiesDidChangeWatchedFiles.
	WorkspaceClientCapabilitiesDidChangeWatchedFilesPool *sync.Pool

	// WorkspaceClientCapabilitiesSymbolPool represents a pool the WorkspaceClientCapabilitiesSymbol.
	WorkspaceClientCapabilitiesSymbolPool *sync.Pool

	// WorkspaceClientCapabilitiesSymbolKindPool represents a pool the WorkspaceClientCapabilitiesSymbolKind.
	WorkspaceClientCapabilitiesSymbolKindPool *sync.Pool

	// WorkspaceClientCapabilitiesExecuteCommandPool represents a pool the WorkspaceClientCapabilitiesExecuteCommand.
	WorkspaceClientCapabilitiesExecuteCommandPool *sync.Pool

	// WorkspaceClientCapabilitiesPool represents a pool the WorkspaceClientCapabilities.
	WorkspaceClientCapabilitiesPool *sync.Pool

	// TextDocumentClientCapabilitiesSynchronizationPool represents a pool the TextDocumentClientCapabilitiesSynchronization.
	TextDocumentClientCapabilitiesSynchronizationPool *sync.Pool

	// TextDocumentClientCapabilitiesCompletionPool represents a pool the TextDocumentClientCapabilitiesCompletion.
	TextDocumentClientCapabilitiesCompletionPool *sync.Pool

	// TextDocumentClientCapabilitiesCompletionItemPool represents a pool the TextDocumentClientCapabilitiesCompletionItem.
	TextDocumentClientCapabilitiesCompletionItemPool *sync.Pool

	// TextDocumentClientCapabilitiesHoverPool represents a pool the TextDocumentClientCapabilitiesHover.
	TextDocumentClientCapabilitiesHoverPool *sync.Pool

	// TextDocumentClientCapabilitiesSignatureHelpPool represents a pool the TextDocumentClientCapabilitiesSignatureHelp.
	TextDocumentClientCapabilitiesSignatureHelpPool *sync.Pool

	// TextDocumentClientCapabilitiesSignatureInformationPool represents a pool the TextDocumentClientCapabilitiesSignatureInformation.
	TextDocumentClientCapabilitiesSignatureInformationPool *sync.Pool

	// TextDocumentClientCapabilitiesReferencesPool represents a pool the TextDocumentClientCapabilitiesReferences.
	TextDocumentClientCapabilitiesReferencesPool *sync.Pool

	// TextDocumentClientCapabilitiesDocumentHighlightPool represents a pool the TextDocumentClientCapabilitiesDocumentHighlight.
	TextDocumentClientCapabilitiesDocumentHighlightPool *sync.Pool

	// TextDocumentClientCapabilitiesDocumentSymbolPool represents a pool the TextDocumentClientCapabilitiesDocumentSymbol.
	TextDocumentClientCapabilitiesDocumentSymbolPool *sync.Pool

	// TextDocumentClientCapabilitiesFormattingPool represents a pool the TextDocumentClientCapabilitiesFormatting.
	TextDocumentClientCapabilitiesFormattingPool *sync.Pool

	// TextDocumentClientCapabilitiesRangeFormattingPool represents a pool the TextDocumentClientCapabilitiesRangeFormatting.
	TextDocumentClientCapabilitiesRangeFormattingPool *sync.Pool

	// TextDocumentClientCapabilitiesOnTypeFormattingPool represents a pool the TextDocumentClientCapabilitiesOnTypeFormatting.
	TextDocumentClientCapabilitiesOnTypeFormattingPool *sync.Pool

	// TextDocumentClientCapabilitiesDeclarationPool represents a pool the TextDocumentClientCapabilitiesDeclaration.
	TextDocumentClientCapabilitiesDeclarationPool *sync.Pool

	// TextDocumentClientCapabilitiesDefinitionPool represents a pool the TextDocumentClientCapabilitiesDefinition.
	TextDocumentClientCapabilitiesDefinitionPool *sync.Pool

	// TextDocumentClientCapabilitiesTypeDefinitionPool represents a pool the TextDocumentClientCapabilitiesTypeDefinition.
	TextDocumentClientCapabilitiesTypeDefinitionPool *sync.Pool

	// TextDocumentClientCapabilitiesImplementationPool represents a pool the TextDocumentClientCapabilitiesImplementation.
	TextDocumentClientCapabilitiesImplementationPool *sync.Pool

	// TextDocumentClientCapabilitiesCodeActionPool represents a pool the TextDocumentClientCapabilitiesCodeAction.
	TextDocumentClientCapabilitiesCodeActionPool *sync.Pool

	// TextDocumentClientCapabilitiesCodeActionLiteralSupportPool represents a pool the TextDocumentClientCapabilitiesCodeActionLiteralSupport.
	TextDocumentClientCapabilitiesCodeActionLiteralSupportPool *sync.Pool

	// TextDocumentClientCapabilitiesCodeActionKindPool represents a pool the TextDocumentClientCapabilitiesCodeActionKind.
	TextDocumentClientCapabilitiesCodeActionKindPool *sync.Pool

	// TextDocumentClientCapabilitiesCodeLensPool represents a pool the TextDocumentClientCapabilitiesCodeLens.
	TextDocumentClientCapabilitiesCodeLensPool *sync.Pool

	// TextDocumentClientCapabilitiesDocumentLinkPool represents a pool the TextDocumentClientCapabilitiesDocumentLink.
	TextDocumentClientCapabilitiesDocumentLinkPool *sync.Pool

	// TextDocumentClientCapabilitiesColorProviderPool represents a pool the TextDocumentClientCapabilitiesColorProvider.
	TextDocumentClientCapabilitiesColorProviderPool *sync.Pool

	// TextDocumentClientCapabilitiesRenamePool represents a pool the TextDocumentClientCapabilitiesRename.
	TextDocumentClientCapabilitiesRenamePool *sync.Pool

	// TextDocumentClientCapabilitiesPublishDiagnosticsPool represents a pool the TextDocumentClientCapabilitiesPublishDiagnostics.
	TextDocumentClientCapabilitiesPublishDiagnosticsPool *sync.Pool

	// TextDocumentClientCapabilitiesFoldingRangePool represents a pool the TextDocumentClientCapabilitiesFoldingRange.
	TextDocumentClientCapabilitiesFoldingRangePool *sync.Pool

	// TextDocumentClientCapabilitiesSelectionRangePool represents a pool the TextDocumentClientCapabilitiesSelectionRange.
	TextDocumentClientCapabilitiesSelectionRangePool *sync.Pool

	// TextDocumentClientCapabilitiesPool represents a pool the TextDocumentClientCapabilities.
	TextDocumentClientCapabilitiesPool *sync.Pool

	// TextDocumentClientCapabilitiesParameterInformationPool represents a pool the TextDocumentClientCapabilitiesParameterInformation.
	TextDocumentClientCapabilitiesParameterInformationPool *sync.Pool

	// ClientCapabilitiesPool represents a pool the ClientCapabilities.
	ClientCapabilitiesPool *sync.Pool

	// InitializeResultPool represents a pool the InitializeResult.
	InitializeResultPool *sync.Pool

	// InitializeErrorPool represents a pool the InitializeError.
	InitializeErrorPool *sync.Pool

	// CompletionOptionsPool represents a pool the CompletionOptions.
	CompletionOptionsPool *sync.Pool

	// SignatureHelpOptionsPool represents a pool the SignatureHelpOptions.
	SignatureHelpOptionsPool *sync.Pool

	// CodeActionOptionsPool represents a pool the CodeActionOptions.
	CodeActionOptionsPool *sync.Pool

	// CodeLensOptionsPool represents a pool the CodeLensOptions.
	CodeLensOptionsPool *sync.Pool

	// DocumentOnTypeFormattingOptionsPool represents a pool the DocumentOnTypeFormattingOptions.
	DocumentOnTypeFormattingOptionsPool *sync.Pool

	// RenameOptionsPool represents a pool the RenameOptions.
	RenameOptionsPool *sync.Pool

	// DocumentLinkOptionsPool represents a pool the DocumentLinkOptions.
	DocumentLinkOptionsPool *sync.Pool

	// ExecuteCommandOptionsPool represents a pool the ExecuteCommandOptions.
	ExecuteCommandOptionsPool *sync.Pool

	// SaveOptionsPool represents a pool the SaveOptions.
	SaveOptionsPool *sync.Pool

	// TextDocumentSyncOptionsPool represents a pool the TextDocumentSyncOptions.
	TextDocumentSyncOptionsPool *sync.Pool

	// StaticRegistrationOptionsPool represents a pool the StaticRegistrationOptions.
	StaticRegistrationOptionsPool *sync.Pool

	// ServerCapabilitiesWorkspacePool represents a pool the ServerCapabilitiesWorkspace.
	ServerCapabilitiesWorkspacePool *sync.Pool

	// ServerCapabilitiesWorkspaceFoldersPool represents a pool the ServerCapabilitiesWorkspaceFolders.
	ServerCapabilitiesWorkspaceFoldersPool *sync.Pool

	// ServerCapabilitiesPool represents a pool the ServerCapabilities.
	ServerCapabilitiesPool *sync.Pool

	// DocumentLinkRegistrationOptionsPool represents a pool the DocumentLinkRegistrationOptions.
	DocumentLinkRegistrationOptionsPool *sync.Pool
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

	// general
	InitializeParamsPool = &sync.Pool{
		New: func() interface{} {
			return &InitializeParams{}
		},
	}
	WorkspaceClientCapabilitiesWorkspaceEditPool = &sync.Pool{
		New: func() interface{} {
			return &WorkspaceClientCapabilitiesWorkspaceEdit{}
		},
	}
	WorkspaceClientCapabilitiesDidChangeConfigurationPool = &sync.Pool{
		New: func() interface{} {
			return &WorkspaceClientCapabilitiesDidChangeConfiguration{}
		},
	}
	WorkspaceClientCapabilitiesDidChangeWatchedFilesPool = &sync.Pool{
		New: func() interface{} {
			return &WorkspaceClientCapabilitiesDidChangeWatchedFiles{}
		},
	}
	WorkspaceClientCapabilitiesSymbolKindPool = &sync.Pool{
		New: func() interface{} {
			return &WorkspaceClientCapabilitiesSymbolKind{}
		},
	}
	WorkspaceClientCapabilitiesExecuteCommandPool = &sync.Pool{
		New: func() interface{} {
			return &WorkspaceClientCapabilitiesExecuteCommand{}
		},
	}
	WorkspaceClientCapabilitiesPool = &sync.Pool{
		New: func() interface{} {
			return &WorkspaceClientCapabilities{}
		},
	}
	TextDocumentClientCapabilitiesSynchronizationPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesSynchronization{}
		},
	}
	TextDocumentClientCapabilitiesCompletionPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesCompletion{}
		},
	}
	TextDocumentClientCapabilitiesCompletionItemPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesCompletionItem{}
		},
	}
	TextDocumentClientCapabilitiesHoverPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesHover{}
		},
	}
	TextDocumentClientCapabilitiesSignatureHelpPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesSignatureHelp{}
		},
	}
	TextDocumentClientCapabilitiesSignatureInformationPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesSignatureInformation{}
		},
	}
	TextDocumentClientCapabilitiesReferencesPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesReferences{}
		},
	}
	TextDocumentClientCapabilitiesDocumentHighlightPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesDocumentHighlight{}
		},
	}
	TextDocumentClientCapabilitiesDocumentSymbolPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesDocumentSymbol{}
		},
	}
	TextDocumentClientCapabilitiesFormattingPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesFormatting{}
		},
	}
	TextDocumentClientCapabilitiesRangeFormattingPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesRangeFormatting{}
		},
	}
	TextDocumentClientCapabilitiesOnTypeFormattingPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesOnTypeFormatting{}
		},
	}
	TextDocumentClientCapabilitiesDeclarationPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesDeclaration{}
		},
	}
	TextDocumentClientCapabilitiesDefinitionPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesDefinition{}
		},
	}
	TextDocumentClientCapabilitiesTypeDefinitionPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesTypeDefinition{}
		},
	}
	TextDocumentClientCapabilitiesImplementationPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesImplementation{}
		},
	}
	TextDocumentClientCapabilitiesCodeActionPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesCodeAction{}
		},
	}
	TextDocumentClientCapabilitiesCodeActionLiteralSupportPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesCodeActionLiteralSupport{}
		},
	}
	TextDocumentClientCapabilitiesCodeActionKindPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesCodeActionKind{}
		},
	}
	TextDocumentClientCapabilitiesCodeLensPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesCodeLens{}
		},
	}
	TextDocumentClientCapabilitiesDocumentLinkPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesDocumentLink{}
		},
	}
	TextDocumentClientCapabilitiesColorProviderPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesColorProvider{}
		},
	}
	TextDocumentClientCapabilitiesRenamePool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesRename{}
		},
	}
	TextDocumentClientCapabilitiesPublishDiagnosticsPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesPublishDiagnostics{}
		},
	}
	TextDocumentClientCapabilitiesFoldingRangePool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesFoldingRange{}
		},
	}
	TextDocumentClientCapabilitiesSelectionRangePool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesSelectionRange{}
		},
	}
	TextDocumentClientCapabilitiesPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilities{}
		},
	}
	TextDocumentClientCapabilitiesParameterInformationPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentClientCapabilitiesParameterInformation{}
		},
	}
	ClientCapabilitiesPool = &sync.Pool{
		New: func() interface{} {
			return &ClientCapabilities{}
		},
	}
	InitializeResultPool = &sync.Pool{
		New: func() interface{} {
			return &InitializeResult{}
		},
	}
	InitializeErrorPool = &sync.Pool{
		New: func() interface{} {
			return &InitializeError{}
		},
	}
	CompletionOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &CompletionOptions{}
		},
	}
	SignatureHelpOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &SignatureHelpOptions{}
		},
	}
	CodeActionOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &CodeActionOptions{}
		},
	}
	CodeLensOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &CodeLensOptions{}
		},
	}
	DocumentOnTypeFormattingOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &DocumentOnTypeFormattingOptions{}
		},
	}
	RenameOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &RenameOptions{}
		},
	}
	DocumentLinkOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &DocumentLinkOptions{}
		},
	}
	ExecuteCommandOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &ExecuteCommandOptions{}
		},
	}
	SaveOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &SaveOptions{}
		},
	}
	TextDocumentSyncOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &TextDocumentSyncOptions{}
		},
	}
	StaticRegistrationOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &StaticRegistrationOptions{}
		},
	}
	ServerCapabilitiesWorkspacePool = &sync.Pool{
		New: func() interface{} {
			return &ServerCapabilitiesWorkspace{}
		},
	}
	ServerCapabilitiesWorkspaceFoldersPool = &sync.Pool{
		New: func() interface{} {
			return &ServerCapabilitiesWorkspaceFolders{}
		},
	}
	ServerCapabilitiesPool = &sync.Pool{
		New: func() interface{} {
			return &ServerCapabilities{}
		},
	}
	DocumentLinkRegistrationOptionsPool = &sync.Pool{
		New: func() interface{} {
			return &DocumentLinkRegistrationOptions{}
		},
	}
}
