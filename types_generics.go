// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"go.lsp.dev/uri"
)

// CancelParamsID the request id to cancel.
type CancelParamsID[T int32, U string] = OneOf[T, U]

// ClientSemanticTokensRequestOptionsFull the client will send the `textDocument/semanticTokens/full` request if the server provides a corresponding handler.
type ClientSemanticTokensRequestOptionsFull[T bool, U ClientSemanticTokensRequestFullDelta] = OneOf[T, U]

// ClientSemanticTokensRequestOptionsRange the client will send the `textDocument/semanticTokens/range` request if the server provides a corresponding handler.
type ClientSemanticTokensRequestOptionsRange[T bool, U Range] = OneOf[T, U]

// CodeActionRequestResult a request to provide commands for the given text document and range.
type CodeActionRequestResult[T Command, U CodeAction] = OneOf[T, U]

// CompletionItemDefaultsEditRange a default edit range.
//
// @since 3.17.0
type CompletionItemDefaultsEditRange[T Range, U EditRangeWithInsertReplace] = OneOf[T, U]

// CompletionItemDocumentation a human-readable string that represents a doc-comment.
type CompletionItemDocumentation[T string, U MarkupContent] = OneOf[T, U]

// CompletionItemTextEdit an TextEdit edit which is applied to a document when selecting this completion. When an edit is provided the value of CompletionItem.insertText insertText is ignored. Most editors support two different operations when accepting a completion item. One is to insert a completion text and the other is to replace an existing text with a completion text. Since this can usually not be predetermined by a server it can report both ranges. Clients need to signal support for `InsertReplaceEdits` via the `textDocument.completion.insertReplaceSupport` client capability property. *Note 1:* The text edit's range as well as both ranges from an insert replace edit must be a [single line] and they must contain the position at which completion has been requested. *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range must be a prefix of the edit's replace range, that means it must be contained and starting at the same position. 3.16.0 additional type `InsertReplaceEdit`.
//
// @since 3.16.0 additional type `InsertReplaceEdit`
type CompletionItemTextEdit[T TextEdit, U InsertReplaceEdit] = OneOf[T, U]

// CompletionResult request to request completion at a given text document position. The request's parameter is of type TextDocumentPosition the response is of type CompletionItem CompletionItem[] or CompletionList or a Thenable that resolves to such. The request can delay the computation of the CompletionItem.detail `detail` and CompletionItem.documentation `documentation` properties to the `completionItem/resolve` request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
type CompletionResult[T []CompletionItem, U CompletionList] = OneOf[T, U]

// DeclarationResult a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Declaration or a
// typed array of DeclarationLink or a Thenable that resolves to such.
type DeclarationResult[T Declaration[Location, []Location], U []DeclarationLink] = OneOf[T, U]

// DefinitionResult a request to resolve the definition location of a symbol at a given text document position. The request's parameter is of type TextDocumentPosition the response is of either type Definition or a typed
// array of DefinitionLink or a Thenable that resolves to such.
type DefinitionResult[T Declaration[Location, []Location], U []DefinitionLink] = OneOf[T, U]

// DiagnosticCode the diagnostic's code, which usually appear in the user interface.
type DiagnosticCode[T int32, U string] = OneOf[T, U]

type DidChangeConfigurationRegistrationOptionsSection[T string, U []string] = OneOf[T, U]

type DocumentDiagnosticReportPartialResultRelatedDocuments[T FullDocumentDiagnosticReport, U UnchangedDocumentDiagnosticReport] = OneOf[T, U]

// DocumentSymbolResult a request to list all symbols found in a given text document. The request's parameter is of type TextDocumentIdentifier the response is of type SymbolInformation SymbolInformation[] or a Thenable that
// resolves to such.
type DocumentSymbolResult[T []SymbolInformation, U []DocumentSymbol] = OneOf[T, U]

// HoverContents the hover's content.
type HoverContents[T MarkupContent, U MarkedString[string, MarkedStringWithLanguage], V []MarkedString[string, MarkedStringWithLanguage]] = OneOf3[T, U, V]

// ImplementationResult a request to resolve the implementation locations of a symbol at a given text document position. The
// request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
type ImplementationResult[T Definition[Location, []Location], U []DefinitionLink] = OneOf[T, U]

// InlayHintLabel the label of this hint. A human readable string or an array of InlayHintLabelPart label parts. *Note* that neither the string nor the label part can be empty.
type InlayHintLabel[T string, U []InlayHintLabelPart] = OneOf[T, U]

// InlayHintLabelPartTooltip the tooltip text when you hover over this label part. Depending on the client capability `inlayHint.resolveSupport` clients might resolve this property late using the resolve request.
type InlayHintLabelPartTooltip[T string, U MarkupContent] = OneOf[T, U]

// InlayHintTooltip the tooltip text when you hover over this item.
type InlayHintTooltip[T string, U MarkupContent] = OneOf[T, U]

// InlineCompletionItemInsertText the text to replace the range with. Must be set.
type InlineCompletionItemInsertText[T string, U StringValue] = OneOf[T, U]

// InlineCompletionResult a request to provide inline completions in a document. The request's parameter is of type InlineCompletionParams, the response is of type InlineCompletion InlineCompletion[] or a Thenable that resolves to such. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionResult[T InlineCompletionList, U []InlineCompletionItem] = OneOf[T, U]

// NotebookCellTextDocumentFilterNotebook a filter that matches against the notebook containing the notebook cell. If a string value is provided it matches against the notebook type. '*' matches every notebook.
type NotebookCellTextDocumentFilterNotebook[T string, U NotebookDocumentFilter[NotebookDocumentFilterNotebookType, NotebookDocumentFilterScheme, NotebookDocumentFilterPattern]] = OneOf[T, U]

// NotebookDocumentFilterWithCellsNotebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
type NotebookDocumentFilterWithCellsNotebook[T string, U NotebookDocumentFilter[NotebookDocumentFilterNotebookType, NotebookDocumentFilterScheme, NotebookDocumentFilterPattern]] = OneOf[T, U]

// NotebookDocumentFilterWithNotebookNotebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
type NotebookDocumentFilterWithNotebookNotebook[T string, U NotebookDocumentFilter[NotebookDocumentFilterNotebookType, NotebookDocumentFilterScheme, NotebookDocumentFilterPattern]] = OneOf[T, U]

// NotebookDocumentSyncOptionsNotebookSelector the notebooks to be synced.
type NotebookDocumentSyncOptionsNotebookSelector[T NotebookDocumentFilterWithNotebook, U NotebookDocumentFilterWithCells] = OneOf[T, U]

// ParameterInformationDocumentation the human-readable doc-comment of this parameter. Will be shown in the UI but can be omitted.
type ParameterInformationDocumentation[T string, U MarkupContent] = OneOf[T, U]

// ParameterInformationLabel the label of this parameter information. Either a string or an inclusive start and exclusive end offsets within its containing signature label. (see SignatureInformation.label). The offsets are based on a UTF-16 string representation as `Position` and `Range` does. To avoid ambiguities a server should use the [start, end] offset value instead of using a substring. Whether a client support this is controlled via `labelOffsetSupport` client capability. *Note*: a label of type string should be a substring of its containing signature label. Its intended use case is to highlight the parameter label
// part in the `SignatureInformation.label`.
type ParameterInformationLabel[T string, U uint32] = OneOf[T, U]

// RelatedFullDocumentDiagnosticReportRelatedDocuments diagnostics of related documents. This information is useful in programming languages where code in a file A can generate diagnostics in a file B which A depends on. An example of such a language is C/C++ where marco definitions in a file a.cpp and result in errors in a header file b.hpp.
//
// @since 3.17.0
type RelatedFullDocumentDiagnosticReportRelatedDocuments[T FullDocumentDiagnosticReport, U UnchangedDocumentDiagnosticReport] = OneOf[T, U]

// RelatedUnchangedDocumentDiagnosticReportRelatedDocuments diagnostics of related documents. This information is useful in programming languages where code in a file A can generate diagnostics in a file B which A depends on. An example of such a language is C/C++ where marco definitions in a file a.cpp and result in errors in a header file b.hpp.
//
// @since 3.17.0
type RelatedUnchangedDocumentDiagnosticReportRelatedDocuments[T FullDocumentDiagnosticReport, U UnchangedDocumentDiagnosticReport] = OneOf[T, U]

// RelativePatternBaseURI a workspace folder or a base URI to which this pattern will be matched against relatively.
type RelativePatternBaseURI[T WorkspaceFolder, U uri.URI] = OneOf[T, U]

// SemanticTokensDeltaResult.
//
// @since 3.16.0
type SemanticTokensDeltaResult[T SemanticTokens, U SemanticTokensDelta] = OneOf[T, U]

// SemanticTokensOptionsFull server supports providing semantic tokens for a full document.
type SemanticTokensOptionsFull[T bool, U SemanticTokensFullDelta] = OneOf[T, U]

// SemanticTokensOptionsRange server supports providing semantic tokens for a specific range of a document.
type SemanticTokensOptionsRange[T bool, U Range] = OneOf[T, U]

// ServerCapabilitiesCallHierarchyProvider the server provides call hierarchy support.
//
// @since 3.16.0
type ServerCapabilitiesCallHierarchyProvider[T bool, U CallHierarchyOptions, V CallHierarchyRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesCodeActionProvider the server provides code actions. CodeActionOptions may only be specified if the client states that it supports `codeActionLiteralSupport` in its initial `initialize` request.
type ServerCapabilitiesCodeActionProvider[T bool, U CodeActionOptions] = OneOf[T, U]

// ServerCapabilitiesColorProvider the server provides color provider support.
type ServerCapabilitiesColorProvider[T bool, U DocumentColorOptions, V DocumentColorRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesDeclarationProvider the server provides Goto Declaration support.
type ServerCapabilitiesDeclarationProvider[T bool, U DeclarationOptions, V DeclarationRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesDefinitionProvider the server provides goto definition support.
type ServerCapabilitiesDefinitionProvider[T bool, U DefinitionOptions] = OneOf[T, U]

// ServerCapabilitiesDiagnosticProvider the server has support for pull model diagnostics.
//
// @since 3.17.0
type ServerCapabilitiesDiagnosticProvider[T DiagnosticOptions, U DiagnosticRegistrationOptions] = OneOf[T, U]

// ServerCapabilitiesDocumentFormattingProvider the server provides document formatting.
type ServerCapabilitiesDocumentFormattingProvider[T bool, U DocumentFormattingOptions] = OneOf[T, U]

// ServerCapabilitiesDocumentHighlightProvider the server provides document highlight support.
type ServerCapabilitiesDocumentHighlightProvider[T bool, U DocumentHighlightOptions] = OneOf[T, U]

// ServerCapabilitiesDocumentRangeFormattingProvider the server provides document range formatting.
type ServerCapabilitiesDocumentRangeFormattingProvider[T bool, U DocumentRangeFormattingOptions] = OneOf[T, U]

// ServerCapabilitiesDocumentSymbolProvider the server provides document symbol support.
type ServerCapabilitiesDocumentSymbolProvider[T bool, U DocumentSymbolOptions] = OneOf[T, U]

// ServerCapabilitiesFoldingRangeProvider the server provides folding provider support.
type ServerCapabilitiesFoldingRangeProvider[T bool, U FoldingRangeOptions, V FoldingRangeRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesHoverProvider the server provides hover support.
type ServerCapabilitiesHoverProvider[T bool, U HoverOptions] = OneOf[T, U]

// ServerCapabilitiesImplementationProvider the server provides Goto Implementation support.
type ServerCapabilitiesImplementationProvider[T bool, U ImplementationOptions, V ImplementationRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesInlayHintProvider the server provides inlay hints.
//
// @since 3.17.0
type ServerCapabilitiesInlayHintProvider[T bool, U InlayHintOptions, V InlayHintRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesInlineCompletionProvider inline completion options used during static registration.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type ServerCapabilitiesInlineCompletionProvider[T bool, U InlineCompletionOptions] = OneOf[T, U]

// ServerCapabilitiesInlineValueProvider the server provides inline values.
//
// @since 3.17.0
type ServerCapabilitiesInlineValueProvider[T bool, U InlineValueOptions, V InlineValueRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesLinkedEditingRangeProvider the server provides linked editing range support.
//
// @since 3.16.0
type ServerCapabilitiesLinkedEditingRangeProvider[T bool, U LinkedEditingRangeOptions, V LinkedEditingRangeRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesMonikerProvider the server provides moniker support.
//
// @since 3.16.0
type ServerCapabilitiesMonikerProvider[T bool, U MonikerOptions, V MonikerRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesNotebookDocumentSync defines how notebook documents are synced.
//
// @since 3.17.0
type ServerCapabilitiesNotebookDocumentSync[T NotebookDocumentSyncOptions, U NotebookDocumentSyncRegistrationOptions] = OneOf[T, U]

// ServerCapabilitiesReferencesProvider the server provides find references support.
type ServerCapabilitiesReferencesProvider[T bool, U ReferenceOptions] = OneOf[T, U]

// ServerCapabilitiesRenameProvider the server provides rename support. RenameOptions may only be specified if the client states that it
// supports `prepareSupport` in its initial `initialize` request.
type ServerCapabilitiesRenameProvider[T bool, U RenameOptions] = OneOf[T, U]

// ServerCapabilitiesSelectionRangeProvider the server provides selection range support.
type ServerCapabilitiesSelectionRangeProvider[T bool, U SelectionRangeOptions, V SelectionRangeRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesSemanticTokensProvider the server provides semantic tokens support.
//
// @since 3.16.0
type ServerCapabilitiesSemanticTokensProvider[T SemanticTokensOptions, U SemanticTokensRegistrationOptions] = OneOf[T, U]

// ServerCapabilitiesTextDocumentSync defines how text documents are synced. Is either a detailed structure defining each notification or for backwards compatibility the TextDocumentSyncKind number.
type ServerCapabilitiesTextDocumentSync[T TextDocumentSyncOptions, U TextDocumentSyncKind] = OneOf[T, U]

// ServerCapabilitiesTypeDefinitionProvider the server provides Goto Type Definition support.
type ServerCapabilitiesTypeDefinitionProvider[T bool, U TypeDefinitionOptions, V TypeDefinitionRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesTypeHierarchyProvider the server provides type hierarchy support.
//
// @since 3.17.0
type ServerCapabilitiesTypeHierarchyProvider[T bool, U TypeHierarchyOptions, V TypeHierarchyRegistrationOptions] = OneOf3[T, U, V]

// ServerCapabilitiesWorkspaceSymbolProvider the server provides workspace symbol support.
type ServerCapabilitiesWorkspaceSymbolProvider[T bool, U WorkspaceSymbolOptions] = OneOf[T, U]

// SignatureInformationDocumentation the human-readable doc-comment of this signature. Will be shown in the UI but can be omitted.
type SignatureInformationDocumentation[T string, U MarkupContent] = OneOf[T, U]

// TextDocumentEditEdits the edits to be applied. 3.16.0 - support for AnnotatedTextEdit. This is guarded using a client capability. 3.18.0 - support for SnippetTextEdit. This is guarded using a client capability.
//
// @since 3.18.0 - support for SnippetTextEdit. This is guarded using a client capability.
type TextDocumentEditEdits[T TextEdit, U AnnotatedTextEdit, V SnippetTextEdit] = OneOf3[T, U, V]

// TextDocumentSyncOptionsSave if present save notifications are sent to the server. If omitted the notification should not be sent.
type TextDocumentSyncOptionsSave[T bool, U SaveOptions] = OneOf[T, U]

// TypeDefinitionResult a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
type TypeDefinitionResult[T Definition[Location, []Location], U []DefinitionLink] = OneOf[T, U]

// WorkspaceEditDocumentChanges depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes are
// either an array of `TextDocumentEdit`s to express changes to n different text documents where each text document edit addresses a specific version of a text document. Or it can contain above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations. Whether a client supports versioned document edits is expressed via `workspace.workspaceEdit.documentChanges` client capability. If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then only plain `TextEdit`s using the `changes` property are supported.
type WorkspaceEditDocumentChanges[T TextDocumentEdit, U CreateFile, V RenameFile, Y DeleteFile] OneOf4[T, U, V, Y]

// WorkspaceFoldersServerCapabilitiesChangeNotifications whether the server wants to receive workspace folder change notifications. If a string is provided the string is treated as an ID under which the notification is registered on the client side. The ID can be used to unregister for these events using the `client/unregisterCapability` request.
type WorkspaceFoldersServerCapabilitiesChangeNotifications[T string, U bool] = OneOf[T, U]

// WorkspaceOptionsTextDocumentContent the server supports the `workspace/textDocumentContent` request.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type WorkspaceOptionsTextDocumentContent[T TextDocumentContentOptions, U TextDocumentContentRegistrationOptions] = OneOf[T, U]

// WorkspaceSymbolLocation the location of the symbol. Whether a server is allowed to return a location without a range depends
// on the client capability `workspace.symbol.resolveSupport`. See SymbolInformation#location for
// more details.
type WorkspaceSymbolLocation[T Location, U LocationURIOnly] = OneOf[T, U]

// WorkspaceSymbolResult a request to list project-wide symbols matching the query string given by the WorkspaceSymbolParams.
// The response is of type SymbolInformation SymbolInformation[] or a Thenable that resolves to such. 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
//
// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
type WorkspaceSymbolResult[T []SymbolInformation, U []WorkspaceSymbol] = OneOf[T, U]
