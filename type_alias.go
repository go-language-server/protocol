// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

type RegularExpressionEngineKind string

// Pattern the glob pattern to watch relative to the base path. Glob patterns can have the following syntax: - `*` to match one or more characters in a path segment - `?` to match on one character in a path segment - `**` to match any number of path segments, including none - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files) - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …) - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
//
// @since 3.17.0
type Pattern string

type NotebookDocumentFilter[T NotebookDocumentFilterNotebookType, U NotebookDocumentFilterScheme, V NotebookDocumentFilterPattern] = OneOf3[T, U, V]

// TextDocumentFilter a document filter denotes a document by different properties like the TextDocument.languageId language, the Uri.scheme scheme of its resource, or a glob-pattern that is applied to the TextDocument.fileName path. Glob patterns can have the following syntax: - `*` to match one or more characters in a path segment - `?` to match on one character in a path segment - `**` to match any number of path segments, including none - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files) - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …) - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`) // // Example: A language filter that applies to typescript files on disk: `{ language: 'typescript', scheme: 'file' }` // // Example: A language filter that applies to all package.json paths: `{ language: 'json', pattern: '**package.json' }`
//
// @since 3.17.0
type TextDocumentFilter[T TextDocumentFilterLanguage, U TextDocumentFilterScheme, V TextDocumentFilterPattern] = OneOf3[T, U, V]

// GlobPattern the glob pattern. Either a string pattern or a relative pattern.
//
// @since 3.17.0
type GlobPattern[T Pattern, U RelativePattern] = OneOf[T, U]

// DocumentFilter a document filter describes a top level text document or a notebook cell document. 3.17.0 - support for NotebookCellTextDocumentFilter.
//
// @since 3.17.0 - support for NotebookCellTextDocumentFilter.
type DocumentFilter[T TextDocumentFilter[TextDocumentFilterLanguage, TextDocumentFilterScheme, TextDocumentFilterPattern], U NotebookCellTextDocumentFilter] = OneOf[T, U]

// MarkedString markedString can be used to render human readable text. It is either a markdown string or a code-block that provides a language and a code snippet. The language identifier is semantically equal to the
// optional language identifier in fenced code blocks in GitHub issues. See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting The pair of a language and a value is an equivalent to markdown: ```${language} ${value} ``` Note that markdown strings will be sanitized - that means html will be escaped. // // Deprecated: use MarkupContent instead.
type MarkedString[T string, U MarkedStringWithLanguage] = OneOf[T, U]

// TextDocumentContentChangeEvent an event describing a change to a text document. If only a text is provided it is considered to be the full content of the document.
type TextDocumentContentChangeEvent[T TextDocumentContentChangePartial, U TextDocumentContentChangeWholeDocument] = OneOf[T, U]

// WorkspaceDocumentDiagnosticReport a workspace diagnostic document report.
//
// @since 3.17.0
type WorkspaceDocumentDiagnosticReport[T WorkspaceFullDocumentDiagnosticReport, U WorkspaceUnchangedDocumentDiagnosticReport] = OneOf[T, U]

// ChangeAnnotationIdentifier an identifier to refer to a change annotation stored with a workspace edit.
type ChangeAnnotationIdentifier string

type ProgressToken[T int32, U string] = OneOf[T, U]

// DocumentSelector a document selector is the combination of one or many document filters. // // Example: `let sel:DocumentSelector = [{ language: 'typescript' }, { language: 'json', pattern: '**∕tsconfig.json' }]`; The use of a string as a document filter is deprecated
//
// @since 3.16.0.
type DocumentSelector []DocumentFilter[TextDocumentFilter[TextDocumentFilterLanguage, TextDocumentFilterScheme, TextDocumentFilterPattern], NotebookCellTextDocumentFilter]

type PrepareRenameResult[T Range, U PrepareRenamePlaceholder, V PrepareRenameDefaultBehavior] = OneOf3[T, U, V]

// DocumentDiagnosticReport the result of a document diagnostic pull request. A report can either be a full report containing all diagnostics for the requested document or an unchanged report indicating that nothing has changed in terms of diagnostics in comparison to the last pull request.
//
// @since 3.17.0
type DocumentDiagnosticReport[T RelatedFullDocumentDiagnosticReport, U RelatedUnchangedDocumentDiagnosticReport] = OneOf[T, U]

// InlineValue inline value information can be provided by different means: - directly as a text value (class InlineValueText). - as a name to use for a variable lookup (class InlineValueVariableLookup) - as an evaluatable expression (class InlineValueEvaluatableExpression) The InlineValue types combines all inline value types into one type.
//
// @since 3.17.0
type InlineValue[T InlineValueText, U InlineValueVariableLookup, V InlineValueEvaluatableExpression] = OneOf3[T, U, V]

// DeclarationLink information about where a symbol is declared. Provides additional metadata over normal Location location declarations, including the range of the declaring symbol. Servers should prefer returning `DeclarationLink` over `Declaration` if supported by the client.
type DeclarationLink LocationLink

// Declaration the declaration of a symbol representation as one or many Location locations.
type Declaration[T Location, U []Location] = OneOf[T, U]

// DefinitionLink information about where a symbol is defined. Provides additional metadata over normal Location location definitions, including the range of the defining symbol.
type DefinitionLink LocationLink

// Definition the definition of a symbol represented as one or many Location locations. For most programming languages there is only one location at which a symbol is defined. Servers should prefer returning `DefinitionLink` over `Definition` if supported by the client.
type Definition[T Location, U []Location] = OneOf[T, U]
