// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "fmt"

type RegularExpressionEngineKind string

// Pattern the glob pattern to watch relative to the base path. Glob patterns can have the following syntax: - `*` to match one or more characters in a path segment - `?` to match on one character in a path segment - `**` to match any number of path segments, including none - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files) - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …) - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
//
// @since 3.17.0
type Pattern string

// NotebookDocumentFilter a notebook document filter denotes a notebook document by different properties. The properties will be match against the notebook's URI (same as with documents)
//
// @since 3.17.0
type NotebookDocumentFilter struct {
	value any
}

func NewNotebookDocumentFilter[T NotebookDocumentFilterNotebookType | NotebookDocumentFilterScheme | NotebookDocumentFilterPattern](val T) NotebookDocumentFilter {
	return NotebookDocumentFilter{
		value: val,
	}
}

func (t NotebookDocumentFilter) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case NotebookDocumentFilterNotebookType:
		return marshal(val)
	case NotebookDocumentFilterScheme:
		return marshal(val)
	case NotebookDocumentFilterPattern:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *NotebookDocumentFilter) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 NotebookDocumentFilterNotebookType
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 NotebookDocumentFilterScheme
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 NotebookDocumentFilterPattern
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// TextDocumentFilter a document filter denotes a document by different properties like the TextDocument.languageId language, the Uri.scheme scheme of its resource, or a glob-pattern that is applied to the TextDocument.fileName path. Glob patterns can have the following syntax: - `*` to match one or more characters in a path segment - `?` to match on one character in a path segment - `**` to match any number of path segments, including none - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files) - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …) - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`) // // Example: A language filter that applies to typescript files on disk: `{ language: 'typescript', scheme: 'file' }` // // Example: A language filter that applies to all package.json paths: `{ language: 'json', pattern: '**package.json' }`
//
// @since 3.17.0
type TextDocumentFilter struct {
	value any
}

func NewTextDocumentFilter[T TextDocumentFilterLanguage | TextDocumentFilterScheme | TextDocumentFilterPattern](val T) TextDocumentFilter {
	return TextDocumentFilter{
		value: val,
	}
}

func (t TextDocumentFilter) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case TextDocumentFilterLanguage:
		return marshal(val)
	case TextDocumentFilterScheme:
		return marshal(val)
	case TextDocumentFilterPattern:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *TextDocumentFilter) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 TextDocumentFilterLanguage
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 TextDocumentFilterScheme
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 TextDocumentFilterPattern
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// GlobPattern the glob pattern. Either a string pattern or a relative pattern.
//
// @since 3.17.0
type GlobPattern struct {
	value any
}

func NewGlobPattern[T Pattern | RelativePattern](val T) GlobPattern {
	return GlobPattern{
		value: val,
	}
}

func (t GlobPattern) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Pattern:
		return marshal(val)
	case RelativePattern:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *GlobPattern) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Pattern
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 RelativePattern
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DocumentFilter a document filter describes a top level text document or a notebook cell document. 3.17.0 - proposed
// support for NotebookCellTextDocumentFilter.
//
// @since 3.17.0 - proposed support for NotebookCellTextDocumentFilter.
type DocumentFilter struct {
	value any
}

func NewDocumentFilter[T TextDocumentFilter | NotebookCellTextDocumentFilter](val T) DocumentFilter {
	return DocumentFilter{
		value: val,
	}
}

func (t DocumentFilter) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case TextDocumentFilter:
		return marshal(val)
	case NotebookCellTextDocumentFilter:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *DocumentFilter) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 TextDocumentFilter
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 NotebookCellTextDocumentFilter
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// MarkedString markedString can be used to render human readable text. It is either a markdown string or a code-block that provides a language and a code snippet. The language identifier is semantically equal to the
// optional language identifier in fenced code blocks in GitHub issues. See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting The pair of a language and a value is an equivalent to markdown: ```${language} ${value} ``` Note that markdown strings will be sanitized - that means html will be escaped. // // Deprecated: use MarkupContent instead.
type MarkedString struct {
	value any
}

func NewMarkedString[T string | MarkedStringWithLanguage](val T) MarkedString {
	return MarkedString{
		value: val,
	}
}

func (t MarkedString) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case MarkedStringWithLanguage:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *MarkedString) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 MarkedStringWithLanguage
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// TextDocumentContentChangeEvent an event describing a change to a text document. If only a text is provided it is considered to be the full content of the document.
type TextDocumentContentChangeEvent struct {
	value any
}

func NewTextDocumentContentChangeEvent[T TextDocumentContentChangePartial | TextDocumentContentChangeWholeDocument](val T) TextDocumentContentChangeEvent {
	return TextDocumentContentChangeEvent{
		value: val,
	}
}

func (t TextDocumentContentChangeEvent) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case TextDocumentContentChangePartial:
		return marshal(val)
	case TextDocumentContentChangeWholeDocument:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *TextDocumentContentChangeEvent) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 TextDocumentContentChangePartial
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 TextDocumentContentChangeWholeDocument
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// WorkspaceDocumentDiagnosticReport a workspace diagnostic document report.
//
// @since 3.17.0
type WorkspaceDocumentDiagnosticReport struct {
	value any
}

func NewWorkspaceDocumentDiagnosticReport[T WorkspaceFullDocumentDiagnosticReport | WorkspaceUnchangedDocumentDiagnosticReport](val T) WorkspaceDocumentDiagnosticReport {
	return WorkspaceDocumentDiagnosticReport{
		value: val,
	}
}

func (t WorkspaceDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case WorkspaceFullDocumentDiagnosticReport:
		return marshal(val)
	case WorkspaceUnchangedDocumentDiagnosticReport:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *WorkspaceDocumentDiagnosticReport) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 WorkspaceFullDocumentDiagnosticReport
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 WorkspaceUnchangedDocumentDiagnosticReport
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// ChangeAnnotationIdentifier an identifier to refer to a change annotation stored with a workspace edit.
type ChangeAnnotationIdentifier string

type ProgressToken struct {
	value any
}

func NewProgressToken[T int32 | string](val T) ProgressToken {
	return ProgressToken{
		value: val,
	}
}

func (t ProgressToken) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case int32:
		return marshal(val)
	case string:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *ProgressToken) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 int32
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 string
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DocumentSelector a document selector is the combination of one or many document filters. // // Example: `let sel:DocumentSelector = [{ language: 'typescript' }, { language: 'json', pattern: '**∕tsconfig.json' }]`; The use of a string as a document filter is deprecated
//
// @since 3.16.0.
type DocumentSelector []DocumentFilter

type PrepareRenameResult struct {
	value any
}

func NewPrepareRenameResult[T Range | PrepareRenamePlaceholder | PrepareRenameDefaultBehavior](val T) PrepareRenameResult {
	return PrepareRenameResult{
		value: val,
	}
}

func (t PrepareRenameResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Range:
		return marshal(val)
	case PrepareRenamePlaceholder:
		return marshal(val)
	case PrepareRenameDefaultBehavior:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *PrepareRenameResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Range
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 PrepareRenamePlaceholder
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 PrepareRenameDefaultBehavior
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DocumentDiagnosticReport the result of a document diagnostic pull request. A report can either be a full report containing all diagnostics for the requested document or an unchanged report indicating that nothing has changed in terms of diagnostics in comparison to the last pull request.
//
// @since 3.17.0
type DocumentDiagnosticReport struct {
	value any
}

func NewDocumentDiagnosticReport[T RelatedFullDocumentDiagnosticReport | RelatedUnchangedDocumentDiagnosticReport](val T) DocumentDiagnosticReport {
	return DocumentDiagnosticReport{
		value: val,
	}
}

func (t DocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case RelatedFullDocumentDiagnosticReport:
		return marshal(val)
	case RelatedUnchangedDocumentDiagnosticReport:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *DocumentDiagnosticReport) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 RelatedFullDocumentDiagnosticReport
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 RelatedUnchangedDocumentDiagnosticReport
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// InlineValue inline value information can be provided by different means: - directly as a text value (class InlineValueText). - as a name to use for a variable lookup (class InlineValueVariableLookup) - as an evaluatable expression (class InlineValueEvaluatableExpression) The InlineValue types combines all inline value types into one type.
//
// @since 3.17.0
type InlineValue struct {
	value any
}

func NewInlineValue[T InlineValueText | InlineValueVariableLookup | InlineValueEvaluatableExpression](val T) InlineValue {
	return InlineValue{
		value: val,
	}
}

func (t InlineValue) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case InlineValueText:
		return marshal(val)
	case InlineValueVariableLookup:
		return marshal(val)
	case InlineValueEvaluatableExpression:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *InlineValue) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 InlineValueText
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 InlineValueVariableLookup
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 InlineValueEvaluatableExpression
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DeclarationLink information about where a symbol is declared. Provides additional metadata over normal Location location declarations, including the range of the declaring symbol. Servers should prefer returning `DeclarationLink` over `Declaration` if supported by the client.
type DeclarationLink LocationLink

// Declaration the declaration of a symbol representation as one or many Location locations.
type Declaration struct {
	value any
}

func NewDeclaration[T Location | []Location](val T) Declaration {
	return Declaration{
		value: val,
	}
}

func (t Declaration) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Location:
		return marshal(val)
	case []Location:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *Declaration) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Location
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []Location
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DefinitionLink information about where a symbol is defined. Provides additional metadata over normal Location location definitions, including the range of the defining symbol.
type DefinitionLink LocationLink

// Definition the definition of a symbol represented as one or many Location locations. For most programming languages there is only one location at which a symbol is defined. Servers should prefer returning `DefinitionLink` over `Definition` if supported by the client.
type Definition struct {
	value any
}

func NewDefinition[T Location | []Location](val T) Definition {
	return Definition{
		value: val,
	}
}

func (t Definition) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Location:
		return marshal(val)
	case []Location:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *Definition) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Location
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []Location
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}
