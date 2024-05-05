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
	Value any `json:"value"`
}

func NewNotebookDocumentFilter[T NotebookDocumentFilterNotebookType | NotebookDocumentFilterScheme | NotebookDocumentFilterPattern](x T) NotebookDocumentFilter {
	return NotebookDocumentFilter{
		Value: x,
	}
}

func (t NotebookDocumentFilter) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case NotebookDocumentFilterNotebookType:
		return marshal(x)
	case NotebookDocumentFilterScheme:
		return marshal(x)
	case NotebookDocumentFilterPattern:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *NotebookDocumentFilter) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 NotebookDocumentFilterNotebookType
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 NotebookDocumentFilterScheme
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 NotebookDocumentFilterPattern
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// TextDocumentFilter a document filter denotes a document by different properties like the TextDocument.languageId language, the Uri.scheme scheme of its resource, or a glob-pattern that is applied to the TextDocument.fileName path. Glob patterns can have the following syntax: - `*` to match one or more characters in a path segment - `?` to match on one character in a path segment - `**` to match any number of path segments, including none - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files) - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …) - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`) // // Example: A language filter that applies to typescript files on disk: `{ language: 'typescript', scheme: 'file' }` // // Example: A language filter that applies to all package.json paths: `{ language: 'json', pattern: '**package.json' }`
//
// @since 3.17.0
type TextDocumentFilter struct {
	Value any `json:"value"`
}

func NewTextDocumentFilter[T TextDocumentFilterLanguage | TextDocumentFilterScheme | TextDocumentFilterPattern](x T) TextDocumentFilter {
	return TextDocumentFilter{
		Value: x,
	}
}

func (t TextDocumentFilter) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case TextDocumentFilterLanguage:
		return marshal(x)
	case TextDocumentFilterScheme:
		return marshal(x)
	case TextDocumentFilterPattern:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *TextDocumentFilter) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 TextDocumentFilterLanguage
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 TextDocumentFilterScheme
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 TextDocumentFilterPattern
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// GlobPattern the glob pattern. Either a string pattern or a relative pattern.
//
// @since 3.17.0
type GlobPattern struct {
	Value any `json:"value"`
}

func NewGlobPattern[T Pattern | RelativePattern](x T) GlobPattern {
	return GlobPattern{
		Value: x,
	}
}

func (t GlobPattern) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Pattern:
		return marshal(x)
	case RelativePattern:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *GlobPattern) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Pattern
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 RelativePattern
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DocumentFilter a document filter describes a top level text document or a notebook cell document. 3.17.0 - proposed
// support for NotebookCellTextDocumentFilter.
//
// @since 3.17.0 - proposed support for NotebookCellTextDocumentFilter.
type DocumentFilter struct {
	Value any `json:"value"`
}

func NewDocumentFilter[T TextDocumentFilter | NotebookCellTextDocumentFilter](x T) DocumentFilter {
	return DocumentFilter{
		Value: x,
	}
}

func (t DocumentFilter) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case TextDocumentFilter:
		return marshal(x)
	case NotebookCellTextDocumentFilter:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *DocumentFilter) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 TextDocumentFilter
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 NotebookCellTextDocumentFilter
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// MarkedString markedString can be used to render human readable text. It is either a markdown string or a code-block that provides a language and a code snippet. The language identifier is semantically equal to the
// optional language identifier in fenced code blocks in GitHub issues. See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting The pair of a language and a value is an equivalent to markdown: ```${language} ${value} ``` Note that markdown strings will be sanitized - that means html will be escaped. // // Deprecated: use MarkupContent instead.
type MarkedString struct {
	Value any `json:"value"`
}

func NewMarkedString[T string | MarkedStringWithLanguage](x T) MarkedString {
	return MarkedString{
		Value: x,
	}
}

func (t MarkedString) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case MarkedStringWithLanguage:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *MarkedString) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 MarkedStringWithLanguage
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// TextDocumentContentChangeEvent an event describing a change to a text document. If only a text is provided it is considered to be the full content of the document.
type TextDocumentContentChangeEvent struct {
	Value any `json:"value"`
}

func NewTextDocumentContentChangeEvent[T TextDocumentContentChangePartial | TextDocumentContentChangeWholeDocument](x T) TextDocumentContentChangeEvent {
	return TextDocumentContentChangeEvent{
		Value: x,
	}
}

func (t TextDocumentContentChangeEvent) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case TextDocumentContentChangePartial:
		return marshal(x)
	case TextDocumentContentChangeWholeDocument:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *TextDocumentContentChangeEvent) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 TextDocumentContentChangePartial
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 TextDocumentContentChangeWholeDocument
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// WorkspaceDocumentDiagnosticReport a workspace diagnostic document report.
//
// @since 3.17.0
type WorkspaceDocumentDiagnosticReport struct {
	Value any `json:"value"`
}

func NewWorkspaceDocumentDiagnosticReport[T WorkspaceFullDocumentDiagnosticReport | WorkspaceUnchangedDocumentDiagnosticReport](x T) WorkspaceDocumentDiagnosticReport {
	return WorkspaceDocumentDiagnosticReport{
		Value: x,
	}
}

func (t WorkspaceDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case WorkspaceFullDocumentDiagnosticReport:
		return marshal(x)
	case WorkspaceUnchangedDocumentDiagnosticReport:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *WorkspaceDocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 WorkspaceFullDocumentDiagnosticReport
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 WorkspaceUnchangedDocumentDiagnosticReport
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// ChangeAnnotationIdentifier an identifier to refer to a change annotation stored with a workspace edit.
type ChangeAnnotationIdentifier string

type ProgressToken struct {
	Value any `json:"value"`
}

func NewProgressToken[T int32 | string](x T) ProgressToken {
	return ProgressToken{
		Value: x,
	}
}

func (t ProgressToken) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case int32:
		return marshal(x)
	case string:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *ProgressToken) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 int32
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 string
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DocumentSelector a document selector is the combination of one or many document filters. // // Example: `let sel:DocumentSelector = [{ language: 'typescript' }, { language: 'json', pattern: '**∕tsconfig.json' }]`; The use of a string as a document filter is deprecated
//
// @since 3.16.0.
type DocumentSelector []DocumentFilter

type PrepareRenameResult struct {
	Value any `json:"value"`
}

func NewPrepareRenameResult[T Range | PrepareRenamePlaceholder | PrepareRenameDefaultBehavior](x T) PrepareRenameResult {
	return PrepareRenameResult{
		Value: x,
	}
}

func (t PrepareRenameResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Range:
		return marshal(x)
	case PrepareRenamePlaceholder:
		return marshal(x)
	case PrepareRenameDefaultBehavior:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *PrepareRenameResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Range
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 PrepareRenamePlaceholder
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 PrepareRenameDefaultBehavior
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DocumentDiagnosticReport the result of a document diagnostic pull request. A report can either be a full report containing all diagnostics for the requested document or an unchanged report indicating that nothing has changed in terms of diagnostics in comparison to the last pull request.
//
// @since 3.17.0
type DocumentDiagnosticReport struct {
	Value any `json:"value"`
}

func NewDocumentDiagnosticReport[T RelatedFullDocumentDiagnosticReport | RelatedUnchangedDocumentDiagnosticReport](x T) DocumentDiagnosticReport {
	return DocumentDiagnosticReport{
		Value: x,
	}
}

func (t DocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case RelatedFullDocumentDiagnosticReport:
		return marshal(x)
	case RelatedUnchangedDocumentDiagnosticReport:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *DocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 RelatedFullDocumentDiagnosticReport
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 RelatedUnchangedDocumentDiagnosticReport
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// InlineValue inline value information can be provided by different means: - directly as a text value (class InlineValueText). - as a name to use for a variable lookup (class InlineValueVariableLookup) - as an evaluatable expression (class InlineValueEvaluatableExpression) The InlineValue types combines all inline value types into one type.
//
// @since 3.17.0
type InlineValue struct {
	Value any `json:"value"`
}

func NewInlineValue[T InlineValueText | InlineValueVariableLookup | InlineValueEvaluatableExpression](x T) InlineValue {
	return InlineValue{
		Value: x,
	}
}

func (t InlineValue) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case InlineValueText:
		return marshal(x)
	case InlineValueVariableLookup:
		return marshal(x)
	case InlineValueEvaluatableExpression:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *InlineValue) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 InlineValueText
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 InlineValueVariableLookup
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 InlineValueEvaluatableExpression
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DeclarationLink information about where a symbol is declared. Provides additional metadata over normal Location location declarations, including the range of the declaring symbol. Servers should prefer returning `DeclarationLink` over `Declaration` if supported by the client.
type DeclarationLink LocationLink

// Declaration the declaration of a symbol representation as one or many Location locations.
type Declaration struct {
	Value any `json:"value"`
}

func NewDeclaration[T Location | []Location](x T) Declaration {
	return Declaration{
		Value: x,
	}
}

func (t Declaration) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Location:
		return marshal(x)
	case []Location:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *Declaration) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Location
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []Location
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}

// DefinitionLink information about where a symbol is defined. Provides additional metadata over normal Location location definitions, including the range of the defining symbol.
type DefinitionLink LocationLink

// Definition the definition of a symbol represented as one or many Location locations. For most programming languages there is only one location at which a symbol is defined. Servers should prefer returning `DefinitionLink` over `Definition` if supported by the client.
type Definition struct {
	Value any `json:"value"`
}

func NewDefinition[T Location | []Location](x T) Definition {
	return Definition{
		Value: x,
	}
}

func (t Definition) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Location:
		return marshal(x)
	case []Location:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unkonwn type: %T", t)
}

func (t *Definition) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Location
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []Location
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}
}
