// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"fmt"

	"go.lsp.dev/uri"
)

// CancelParamsID the request id to cancel.
type CancelParamsID struct {
	value any
}

func NewCancelParamsID[T int32 | string](val T) CancelParamsID {
	return CancelParamsID{
		value: val,
	}
}

func (t CancelParamsID) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case int32:
		return marshal(val)
	case string:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CancelParamsID) UnmarshalJSON(val []byte) error {
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
	return &UnmarshalError{"unmarshal failed to match one of [int32 string]"}
}

// ClientSemanticTokensRequestOptionsFull the client will send the `textDocument/semanticTokens/full` request if the server provides a corresponding handler.
type ClientSemanticTokensRequestOptionsFull struct {
	value any
}

func NewClientSemanticTokensRequestOptionsFull[T bool | ClientSemanticTokensRequestFullDelta](val T) *ClientSemanticTokensRequestOptionsFull {
	return &ClientSemanticTokensRequestOptionsFull{
		value: val,
	}
}

func (t ClientSemanticTokensRequestOptionsFull) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case ClientSemanticTokensRequestFullDelta:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ClientSemanticTokensRequestOptionsFull) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 ClientSemanticTokensRequestFullDelta
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool ClientSemanticTokensRequestFullDelta]"}
}

// ClientSemanticTokensRequestOptionsRange the client will send the `textDocument/semanticTokens/range` request if the server provides a corresponding handler.
type ClientSemanticTokensRequestOptionsRange struct {
	value any
}

func NewClientSemanticTokensRequestOptionsRange[T bool | Range](val T) *ClientSemanticTokensRequestOptionsRange {
	return &ClientSemanticTokensRequestOptionsRange{
		value: val,
	}
}

func (t ClientSemanticTokensRequestOptionsRange) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case Range:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ClientSemanticTokensRequestOptionsRange) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 Range
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool Range]"}
}

// CodeActionRequestResult a request to provide commands for the given text document and range.
type CodeActionRequestResult struct {
	value any
}

func NewCodeActionRequestResult[T Command | CodeAction](val T) *CodeActionRequestResult {
	return &CodeActionRequestResult{
		value: val,
	}
}

func (t CodeActionRequestResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Command:
		return marshal(val)
	case CodeAction:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CodeActionRequestResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Command
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 CodeAction
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Command CodeAction]"}
}

// CodeActionRequestResult a request to provide commands for the given text document and range.
type CodeActionRequestResult struct {
	value any
}

func NewCodeActionRequestResult[T Command | CodeAction](val T) *CodeActionRequestResult {
	return &CodeActionRequestResult{
		value: val,
	}
}

func (t CodeActionRequestResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Command:
		return marshal(val)
	case CodeAction:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CodeActionRequestResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Command
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 CodeAction
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Command CodeAction]"}
}

// CompletionItemDefaultsEditRange a default edit range.
//
// @since 3.17.0
type CompletionItemDefaultsEditRange struct {
	value any
}

func NewCompletionItemDefaultsEditRange[T Range | EditRangeWithInsertReplace](val T) *CompletionItemDefaultsEditRange {
	return &CompletionItemDefaultsEditRange{
		value: val,
	}
}

func (t CompletionItemDefaultsEditRange) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Range:
		return marshal(val)
	case EditRangeWithInsertReplace:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CompletionItemDefaultsEditRange) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Range
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 EditRangeWithInsertReplace
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Range EditRangeWithInsertReplace]"}
}

// CompletionItemDocumentation a human-readable string that represents a doc-comment.
type CompletionItemDocumentation struct {
	value any
}

func NewCompletionItemDocumentation[T string | MarkupContent](val T) *CompletionItemDocumentation {
	return &CompletionItemDocumentation{
		value: val,
	}
}

func (t CompletionItemDocumentation) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case MarkupContent:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CompletionItemDocumentation) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// CompletionItemTextEdit an TextEdit edit which is applied to a document when selecting this completion. When an edit is provided the value of CompletionItem.insertText insertText is ignored. Most editors support two different operations when accepting a completion item. One is to insert a completion text and the other is to replace an existing text with a completion text. Since this can usually not be predetermined by a server it can report both ranges. Clients need to signal support for `InsertReplaceEdits` via the `textDocument.completion.insertReplaceSupport` client capability property. *Note 1:* The text edit's range as well as both ranges from an insert replace edit must be a [single line] and they must contain the position at which completion has been requested. *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range must be a prefix of the edit's replace range, that means it must be contained and starting at the same position. 3.16.0 additional type `InsertReplaceEdit`.
//
// @since 3.16.0 additional type `InsertReplaceEdit`
type CompletionItemTextEdit struct {
	value any
}

func NewCompletionItemTextEdit[T TextEdit | InsertReplaceEdit](val T) *CompletionItemTextEdit {
	return &CompletionItemTextEdit{
		value: val,
	}
}

func (t CompletionItemTextEdit) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case TextEdit:
		return marshal(val)
	case InsertReplaceEdit:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CompletionItemTextEdit) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 TextEdit
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 InsertReplaceEdit
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [TextEdit InsertReplaceEdit]"}
}

// CompletionResult request to request completion at a given text document position. The request's parameter is of type TextDocumentPosition the response is of type CompletionItem CompletionItem[] or CompletionList or a Thenable that resolves to such. The request can delay the computation of the CompletionItem.detail `detail` and CompletionItem.documentation `documentation` properties to the `completionItem/resolve` request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
type CompletionResult struct {
	value any
}

func NewCompletionResult[T []CompletionItem | CompletionList](val T) *CompletionResult {
	return &CompletionResult{
		value: val,
	}
}

func (t CompletionResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case []CompletionItem:
		return marshal(val)
	case CompletionList:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CompletionResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 []CompletionItem
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 CompletionList
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [[]CompletionItem CompletionList]"}
}

// CompletionResult request to request completion at a given text document position. The request's parameter is of type TextDocumentPosition the response is of type CompletionItem CompletionItem[] or CompletionList or a Thenable that resolves to such. The request can delay the computation of the CompletionItem.detail `detail` and CompletionItem.documentation `documentation` properties to the `completionItem/resolve` request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
type CompletionResult struct {
	value any
}

func NewCompletionResult[T []CompletionItem | CompletionList](val T) *CompletionResult {
	return &CompletionResult{
		value: val,
	}
}

func (t CompletionResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case []CompletionItem:
		return marshal(val)
	case CompletionList:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CompletionResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 []CompletionItem
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 CompletionList
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [[]CompletionItem CompletionList]"}
}

// DeclarationResult a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Declaration or a
// typed array of DeclarationLink or a Thenable that resolves to such.
type DeclarationResult struct {
	value any
}

func NewDeclarationResult[T Declaration | []DeclarationLink](val T) *DeclarationResult {
	return &DeclarationResult{
		value: val,
	}
}

func (t DeclarationResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Declaration:
		return marshal(val)
	case []DeclarationLink:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DeclarationResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Declaration
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DeclarationLink
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Declaration []DeclarationLink]"}
}

// DeclarationResult a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Declaration or a
// typed array of DeclarationLink or a Thenable that resolves to such.
type DeclarationResult struct {
	value any
}

func NewDeclarationResult[T Declaration | []DeclarationLink](val T) *DeclarationResult {
	return &DeclarationResult{
		value: val,
	}
}

func (t DeclarationResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Declaration:
		return marshal(val)
	case []DeclarationLink:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DeclarationResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Declaration
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DeclarationLink
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Declaration []DeclarationLink]"}
}

// DefinitionResult a request to resolve the definition location of a symbol at a given text document position. The request's parameter is of type TextDocumentPosition the response is of either type Definition or a typed
// array of DefinitionLink or a Thenable that resolves to such.
type DefinitionResult struct {
	value any
}

func NewDefinitionResult[T Definition | []DefinitionLink](val T) *DefinitionResult {
	return &DefinitionResult{
		value: val,
	}
}

func (t DefinitionResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Definition:
		return marshal(val)
	case []DefinitionLink:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DefinitionResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Definition
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DefinitionLink
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Definition []DefinitionLink]"}
}

// DefinitionResult a request to resolve the definition location of a symbol at a given text document position. The request's parameter is of type TextDocumentPosition the response is of either type Definition or a typed
// array of DefinitionLink or a Thenable that resolves to such.
type DefinitionResult struct {
	value any
}

func NewDefinitionResult[T Definition | []DefinitionLink](val T) *DefinitionResult {
	return &DefinitionResult{
		value: val,
	}
}

func (t DefinitionResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Definition:
		return marshal(val)
	case []DefinitionLink:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DefinitionResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Definition
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DefinitionLink
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Definition []DefinitionLink]"}
}

// DiagnosticCode the diagnostic's code, which usually appear in the user interface.
type DiagnosticCode struct {
	value any
}

func NewDiagnosticCode[T int32 | string](val T) DiagnosticCode {
	return DiagnosticCode{
		value: val,
	}
}

func (t DiagnosticCode) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case int32:
		return marshal(val)
	case string:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DiagnosticCode) UnmarshalJSON(val []byte) error {
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
	return &UnmarshalError{"unmarshal failed to match one of [int32 string]"}
}

type DidChangeConfigurationRegistrationOptionsSection struct {
	value any
}

func NewDidChangeConfigurationRegistrationOptionsSection[T string | []string](val T) DidChangeConfigurationRegistrationOptionsSection {
	return DidChangeConfigurationRegistrationOptionsSection{
		value: val,
	}
}

func (t DidChangeConfigurationRegistrationOptionsSection) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case []string:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DidChangeConfigurationRegistrationOptionsSection) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []string
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string []string]"}
}

// @since 3.17.0
type DocumentDiagnosticReportPartialResultRelatedDocuments struct {
	value any
}

func NewDocumentDiagnosticReportPartialResultRelatedDocuments[T FullDocumentDiagnosticReport | UnchangedDocumentDiagnosticReport](val T) *DocumentDiagnosticReportPartialResultRelatedDocuments {
	return &DocumentDiagnosticReportPartialResultRelatedDocuments{
		value: val,
	}
}

func (t DocumentDiagnosticReportPartialResultRelatedDocuments) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case FullDocumentDiagnosticReport:
		return marshal(val)
	case UnchangedDocumentDiagnosticReport:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DocumentDiagnosticReportPartialResultRelatedDocuments) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 FullDocumentDiagnosticReport
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 UnchangedDocumentDiagnosticReport
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [FullDocumentDiagnosticReport UnchangedDocumentDiagnosticReport]"}
}

// DocumentSymbolResult a request to list all symbols found in a given text document. The request's parameter is of type TextDocumentIdentifier the response is of type SymbolInformation SymbolInformation[] or a Thenable that
// resolves to such.
type DocumentSymbolResult struct {
	value any
}

func NewDocumentSymbolResult[T []SymbolInformation | []DocumentSymbol](val T) DocumentSymbolResult {
	return DocumentSymbolResult{
		value: val,
	}
}

func (t DocumentSymbolResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case []SymbolInformation:
		return marshal(val)
	case []DocumentSymbol:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DocumentSymbolResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 []SymbolInformation
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DocumentSymbol
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [[]SymbolInformation []DocumentSymbol]"}
}

// DocumentSymbolResult a request to list all symbols found in a given text document. The request's parameter is of type TextDocumentIdentifier the response is of type SymbolInformation SymbolInformation[] or a Thenable that
// resolves to such.
type DocumentSymbolResult struct {
	value any
}

func NewDocumentSymbolResult[T []SymbolInformation | []DocumentSymbol](val T) DocumentSymbolResult {
	return DocumentSymbolResult{
		value: val,
	}
}

func (t DocumentSymbolResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case []SymbolInformation:
		return marshal(val)
	case []DocumentSymbol:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DocumentSymbolResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 []SymbolInformation
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DocumentSymbol
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [[]SymbolInformation []DocumentSymbol]"}
}

// HoverContents the hover's content.
type HoverContents struct {
	value any
}

func NewHoverContents[T MarkupContent | MarkedString | []MarkedString](val T) *HoverContents {
	return &HoverContents{
		value: val,
	}
}

func (t HoverContents) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case MarkupContent:
		return marshal(val)
	case MarkedString:
		return marshal(val)
	case []MarkedString:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *HoverContents) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 MarkupContent
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 MarkedString
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 []MarkedString
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [MarkupContent MarkedString []MarkedString]"}
}

// ImplementationResult a request to resolve the implementation locations of a symbol at a given text document position. The
// request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
type ImplementationResult struct {
	value any
}

func NewImplementationResult[T Definition | []DefinitionLink](val T) *ImplementationResult {
	return &ImplementationResult{
		value: val,
	}
}

func (t ImplementationResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Definition:
		return marshal(val)
	case []DefinitionLink:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ImplementationResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Definition
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DefinitionLink
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Definition []DefinitionLink]"}
}

// ImplementationResult a request to resolve the implementation locations of a symbol at a given text document position. The
// request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
type ImplementationResult struct {
	value any
}

func NewImplementationResult[T Definition | []DefinitionLink](val T) *ImplementationResult {
	return &ImplementationResult{
		value: val,
	}
}

func (t ImplementationResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Definition:
		return marshal(val)
	case []DefinitionLink:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ImplementationResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Definition
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DefinitionLink
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Definition []DefinitionLink]"}
}

// InlayHintLabel the label of this hint. A human readable string or an array of InlayHintLabelPart label parts. *Note* that neither the string nor the label part can be empty.
type InlayHintLabel struct {
	value any
}

func NewInlayHintLabel[T string | []InlayHintLabelPart](val T) InlayHintLabel {
	return InlayHintLabel{
		value: val,
	}
}

func (t InlayHintLabel) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case []InlayHintLabelPart:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlayHintLabel) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []InlayHintLabelPart
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string []InlayHintLabelPart]"}
}

// InlayHintLabelPartTooltip the tooltip text when you hover over this label part. Depending on the client capability `inlayHint.resolveSupport` clients might resolve this property late using the resolve request.
type InlayHintLabelPartTooltip struct {
	value any
}

func NewInlayHintLabelPartTooltip[T string | MarkupContent](val T) *InlayHintLabelPartTooltip {
	return &InlayHintLabelPartTooltip{
		value: val,
	}
}

func (t InlayHintLabelPartTooltip) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case MarkupContent:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlayHintLabelPartTooltip) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// InlayHintTooltip the tooltip text when you hover over this item.
type InlayHintTooltip struct {
	value any
}

func NewInlayHintTooltip[T string | MarkupContent](val T) *InlayHintTooltip {
	return &InlayHintTooltip{
		value: val,
	}
}

func (t InlayHintTooltip) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case MarkupContent:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlayHintTooltip) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// InlineCompletionItemInsertText the text to replace the range with. Must be set.
type InlineCompletionItemInsertText struct {
	value any
}

func NewInlineCompletionItemInsertText[T string | StringValue](val T) *InlineCompletionItemInsertText {
	return &InlineCompletionItemInsertText{
		value: val,
	}
}

func (t InlineCompletionItemInsertText) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case StringValue:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlineCompletionItemInsertText) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 StringValue
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string StringValue]"}
}

// InlineCompletionResult a request to provide inline completions in a document. The request's parameter is of type InlineCompletionParams, the response is of type InlineCompletion InlineCompletion[] or a Thenable that resolves to such. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionResult struct {
	value any
}

func NewInlineCompletionResult[T InlineCompletionList | []InlineCompletionItem](val T) *InlineCompletionResult {
	return &InlineCompletionResult{
		value: val,
	}
}

func (t InlineCompletionResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case InlineCompletionList:
		return marshal(val)
	case []InlineCompletionItem:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlineCompletionResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 InlineCompletionList
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []InlineCompletionItem
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [InlineCompletionList []InlineCompletionItem]"}
}

// InlineCompletionResult a request to provide inline completions in a document. The request's parameter is of type InlineCompletionParams, the response is of type InlineCompletion InlineCompletion[] or a Thenable that resolves to such. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionResult struct {
	value any
}

func NewInlineCompletionResult[T InlineCompletionList | []InlineCompletionItem](val T) *InlineCompletionResult {
	return &InlineCompletionResult{
		value: val,
	}
}

func (t InlineCompletionResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case InlineCompletionList:
		return marshal(val)
	case []InlineCompletionItem:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlineCompletionResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 InlineCompletionList
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []InlineCompletionItem
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [InlineCompletionList []InlineCompletionItem]"}
}

// NotebookCellTextDocumentFilterNotebook a filter that matches against the notebook containing the notebook cell. If a string value is provided it matches against the notebook type. '*' matches every notebook.
type NotebookCellTextDocumentFilterNotebook struct {
	value any
}

func NewNotebookCellTextDocumentFilterNotebook[T string | NotebookDocumentFilter](val T) *NotebookCellTextDocumentFilterNotebook {
	return &NotebookCellTextDocumentFilterNotebook{
		value: val,
	}
}

func (t NotebookCellTextDocumentFilterNotebook) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case NotebookDocumentFilter:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *NotebookCellTextDocumentFilterNotebook) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 NotebookDocumentFilter
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string NotebookDocumentFilter]"}
}

// NotebookDocumentFilterWithCellsNotebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
type NotebookDocumentFilterWithCellsNotebook struct {
	value any
}

func NewNotebookDocumentFilterWithCellsNotebook[T string | NotebookDocumentFilter](val T) *NotebookDocumentFilterWithCellsNotebook {
	return &NotebookDocumentFilterWithCellsNotebook{
		value: val,
	}
}

func (t NotebookDocumentFilterWithCellsNotebook) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case NotebookDocumentFilter:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *NotebookDocumentFilterWithCellsNotebook) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 NotebookDocumentFilter
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string NotebookDocumentFilter]"}
}

// NotebookDocumentFilterWithNotebookNotebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
type NotebookDocumentFilterWithNotebookNotebook struct {
	value any
}

func NewNotebookDocumentFilterWithNotebookNotebook[T string | NotebookDocumentFilter](val T) *NotebookDocumentFilterWithNotebookNotebook {
	return &NotebookDocumentFilterWithNotebookNotebook{
		value: val,
	}
}

func (t NotebookDocumentFilterWithNotebookNotebook) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case NotebookDocumentFilter:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *NotebookDocumentFilterWithNotebookNotebook) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 NotebookDocumentFilter
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string NotebookDocumentFilter]"}
}

// NotebookDocumentSyncOptionsNotebookSelector the notebooks to be synced.
//
// @since 3.17.0
type NotebookDocumentSyncOptionsNotebookSelector struct {
	value any
}

func NewNotebookDocumentSyncOptionsNotebookSelector[T NotebookDocumentFilterWithNotebook | NotebookDocumentFilterWithCells](val T) *NotebookDocumentSyncOptionsNotebookSelector {
	return &NotebookDocumentSyncOptionsNotebookSelector{
		value: val,
	}
}

func (t NotebookDocumentSyncOptionsNotebookSelector) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case NotebookDocumentFilterWithNotebook:
		return marshal(val)
	case NotebookDocumentFilterWithCells:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *NotebookDocumentSyncOptionsNotebookSelector) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 NotebookDocumentFilterWithNotebook
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 NotebookDocumentFilterWithCells
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [NotebookDocumentFilterWithNotebook NotebookDocumentFilterWithCells]"}
}

// ParameterInformationDocumentation the human-readable doc-comment of this parameter. Will be shown in the UI but can be omitted.
type ParameterInformationDocumentation struct {
	value any
}

func NewParameterInformationDocumentation[T string | MarkupContent](val T) *ParameterInformationDocumentation {
	return &ParameterInformationDocumentation{
		value: val,
	}
}

func (t ParameterInformationDocumentation) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case MarkupContent:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ParameterInformationDocumentation) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// ParameterInformationLabel the label of this parameter information. Either a string or an inclusive start and exclusive end offsets within its containing signature label. (see SignatureInformation.label). The offsets are based on a UTF-16 string representation as `Position` and `Range` does. To avoid ambiguities a server should use the [start, end] offset value instead of using a substring. Whether a client support this is controlled via `labelOffsetSupport` client capability. *Note*: a label of type string should be a substring of its containing signature label. Its intended use case is to highlight the parameter label
// part in the `SignatureInformation.label`.
type ParameterInformationLabel struct {
	value any
}

func NewParameterInformationLabel[T string | uint32](val T) *ParameterInformationLabel {
	return &ParameterInformationLabel{
		value: val,
	}
}

func (t ParameterInformationLabel) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case uint32:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ParameterInformationLabel) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 uint32
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string uint32]"}
}

// RelatedFullDocumentDiagnosticReportRelatedDocuments diagnostics of related documents. This information is useful in programming languages where code in a file A can generate diagnostics in a file B which A depends on. An example of such a language is C/C++ where marco definitions in a file a.cpp and result in errors in a header file b.hpp.
//
// @since 3.17.0
type RelatedFullDocumentDiagnosticReportRelatedDocuments struct {
	value any
}

func NewRelatedFullDocumentDiagnosticReportRelatedDocuments[T FullDocumentDiagnosticReport | UnchangedDocumentDiagnosticReport](val T) *RelatedFullDocumentDiagnosticReportRelatedDocuments {
	return &RelatedFullDocumentDiagnosticReportRelatedDocuments{
		value: val,
	}
}

func (t RelatedFullDocumentDiagnosticReportRelatedDocuments) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case FullDocumentDiagnosticReport:
		return marshal(val)
	case UnchangedDocumentDiagnosticReport:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *RelatedFullDocumentDiagnosticReportRelatedDocuments) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 FullDocumentDiagnosticReport
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 UnchangedDocumentDiagnosticReport
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [FullDocumentDiagnosticReport UnchangedDocumentDiagnosticReport]"}
}

// RelatedUnchangedDocumentDiagnosticReportRelatedDocuments diagnostics of related documents. This information is useful in programming languages where code in a file A can generate diagnostics in a file B which A depends on. An example of such a language is C/C++ where marco definitions in a file a.cpp and result in errors in a header file b.hpp.
//
// @since 3.17.0
type RelatedUnchangedDocumentDiagnosticReportRelatedDocuments struct {
	value any
}

func NewRelatedUnchangedDocumentDiagnosticReportRelatedDocuments[T FullDocumentDiagnosticReport | UnchangedDocumentDiagnosticReport](val T) *RelatedUnchangedDocumentDiagnosticReportRelatedDocuments {
	return &RelatedUnchangedDocumentDiagnosticReportRelatedDocuments{
		value: val,
	}
}

func (t RelatedUnchangedDocumentDiagnosticReportRelatedDocuments) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case FullDocumentDiagnosticReport:
		return marshal(val)
	case UnchangedDocumentDiagnosticReport:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *RelatedUnchangedDocumentDiagnosticReportRelatedDocuments) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 FullDocumentDiagnosticReport
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 UnchangedDocumentDiagnosticReport
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [FullDocumentDiagnosticReport UnchangedDocumentDiagnosticReport]"}
}

// RelativePatternBaseURI a workspace folder or a base URI to which this pattern will be matched against relatively.
type RelativePatternBaseURI struct {
	value any
}

func NewRelativePatternBaseURI[T WorkspaceFolder | uri.URI](val T) *RelativePatternBaseURI {
	return &RelativePatternBaseURI{
		value: val,
	}
}

func (t RelativePatternBaseURI) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case WorkspaceFolder:
		return marshal(val)
	case uri.URI:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *RelativePatternBaseURI) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 WorkspaceFolder
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 uri.URI
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [WorkspaceFolder uri.URI]"}
}

// SemanticTokensDeltaResult.
//
// @since 3.16.0
type SemanticTokensDeltaResult struct {
	value any
}

func NewSemanticTokensDeltaResult[T SemanticTokens | SemanticTokensDelta](val T) *SemanticTokensDeltaResult {
	return &SemanticTokensDeltaResult{
		value: val,
	}
}

func (t SemanticTokensDeltaResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case SemanticTokens:
		return marshal(val)
	case SemanticTokensDelta:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *SemanticTokensDeltaResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 SemanticTokens
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 SemanticTokensDelta
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [SemanticTokens SemanticTokensDelta]"}
}

// SemanticTokensDeltaResult.
//
// @since 3.16.0
type SemanticTokensDeltaResult struct {
	value any
}

func NewSemanticTokensDeltaResult[T SemanticTokens | SemanticTokensDelta](val T) *SemanticTokensDeltaResult {
	return &SemanticTokensDeltaResult{
		value: val,
	}
}

func (t SemanticTokensDeltaResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case SemanticTokens:
		return marshal(val)
	case SemanticTokensDelta:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *SemanticTokensDeltaResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 SemanticTokens
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 SemanticTokensDelta
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [SemanticTokens SemanticTokensDelta]"}
}

// SemanticTokensOptionsFull server supports providing semantic tokens for a full document.
type SemanticTokensOptionsFull struct {
	value any
}

func NewSemanticTokensOptionsFull[T bool | SemanticTokensFullDelta](val T) *SemanticTokensOptionsFull {
	return &SemanticTokensOptionsFull{
		value: val,
	}
}

func (t SemanticTokensOptionsFull) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case SemanticTokensFullDelta:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *SemanticTokensOptionsFull) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 SemanticTokensFullDelta
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool SemanticTokensFullDelta]"}
}

// SemanticTokensOptionsRange server supports providing semantic tokens for a specific range of a document.
type SemanticTokensOptionsRange struct {
	value any
}

func NewSemanticTokensOptionsRange[T bool | Range](val T) *SemanticTokensOptionsRange {
	return &SemanticTokensOptionsRange{
		value: val,
	}
}

func (t SemanticTokensOptionsRange) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case Range:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *SemanticTokensOptionsRange) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 Range
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool Range]"}
}

// ServerCapabilitiesCallHierarchyProvider the server provides call hierarchy support.
//
// @since 3.16.0
type ServerCapabilitiesCallHierarchyProvider struct {
	value any
}

func NewServerCapabilitiesCallHierarchyProvider[T bool | CallHierarchyOptions | CallHierarchyRegistrationOptions](val T) *ServerCapabilitiesCallHierarchyProvider {
	return &ServerCapabilitiesCallHierarchyProvider{
		value: val,
	}
}

func (t ServerCapabilitiesCallHierarchyProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case CallHierarchyOptions:
		return marshal(val)
	case CallHierarchyRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesCallHierarchyProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 CallHierarchyOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 CallHierarchyRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool CallHierarchyOptions CallHierarchyRegistrationOptions]"}
}

// ServerCapabilitiesCodeActionProvider the server provides code actions. CodeActionOptions may only be specified if the client states that it supports `codeActionLiteralSupport` in its initial `initialize` request.
type ServerCapabilitiesCodeActionProvider struct {
	value any
}

func NewServerCapabilitiesCodeActionProvider[T bool | CodeActionOptions](val T) *ServerCapabilitiesCodeActionProvider {
	return &ServerCapabilitiesCodeActionProvider{
		value: val,
	}
}

func (t ServerCapabilitiesCodeActionProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case CodeActionOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesCodeActionProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 CodeActionOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool CodeActionOptions]"}
}

// ServerCapabilitiesColorProvider the server provides color provider support.
type ServerCapabilitiesColorProvider struct {
	value any
}

func NewServerCapabilitiesColorProvider[T bool | DocumentColorOptions | DocumentColorRegistrationOptions](val T) *ServerCapabilitiesColorProvider {
	return &ServerCapabilitiesColorProvider{
		value: val,
	}
}

func (t ServerCapabilitiesColorProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case DocumentColorOptions:
		return marshal(val)
	case DocumentColorRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesColorProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 DocumentColorOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 DocumentColorRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentColorOptions DocumentColorRegistrationOptions]"}
}

// ServerCapabilitiesDeclarationProvider the server provides Goto Declaration support.
type ServerCapabilitiesDeclarationProvider struct {
	value any
}

func NewServerCapabilitiesDeclarationProvider[T bool | DeclarationOptions | DeclarationRegistrationOptions](val T) *ServerCapabilitiesDeclarationProvider {
	return &ServerCapabilitiesDeclarationProvider{
		value: val,
	}
}

func (t ServerCapabilitiesDeclarationProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case DeclarationOptions:
		return marshal(val)
	case DeclarationRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDeclarationProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 DeclarationOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 DeclarationRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DeclarationOptions DeclarationRegistrationOptions]"}
}

// ServerCapabilitiesDefinitionProvider the server provides goto definition support.
type ServerCapabilitiesDefinitionProvider struct {
	value any
}

func NewServerCapabilitiesDefinitionProvider[T bool | DefinitionOptions](val T) *ServerCapabilitiesDefinitionProvider {
	return &ServerCapabilitiesDefinitionProvider{
		value: val,
	}
}

func (t ServerCapabilitiesDefinitionProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case DefinitionOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDefinitionProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 DefinitionOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DefinitionOptions]"}
}

// ServerCapabilitiesDiagnosticProvider the server has support for pull model diagnostics.
//
// @since 3.17.0
type ServerCapabilitiesDiagnosticProvider struct {
	value any
}

func NewServerCapabilitiesDiagnosticProvider[T DiagnosticOptions | DiagnosticRegistrationOptions](val T) *ServerCapabilitiesDiagnosticProvider {
	return &ServerCapabilitiesDiagnosticProvider{
		value: val,
	}
}

func (t ServerCapabilitiesDiagnosticProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case DiagnosticOptions:
		return marshal(val)
	case DiagnosticRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDiagnosticProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 DiagnosticOptions
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 DiagnosticRegistrationOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [DiagnosticOptions DiagnosticRegistrationOptions]"}
}

// ServerCapabilitiesDocumentFormattingProvider the server provides document formatting.
type ServerCapabilitiesDocumentFormattingProvider struct {
	value any
}

func NewServerCapabilitiesDocumentFormattingProvider[T bool | DocumentFormattingOptions](val T) *ServerCapabilitiesDocumentFormattingProvider {
	return &ServerCapabilitiesDocumentFormattingProvider{
		value: val,
	}
}

func (t ServerCapabilitiesDocumentFormattingProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case DocumentFormattingOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDocumentFormattingProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 DocumentFormattingOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentFormattingOptions]"}
}

// ServerCapabilitiesDocumentHighlightProvider the server provides document highlight support.
type ServerCapabilitiesDocumentHighlightProvider struct {
	value any
}

func NewServerCapabilitiesDocumentHighlightProvider[T bool | DocumentHighlightOptions](val T) *ServerCapabilitiesDocumentHighlightProvider {
	return &ServerCapabilitiesDocumentHighlightProvider{
		value: val,
	}
}

func (t ServerCapabilitiesDocumentHighlightProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case DocumentHighlightOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDocumentHighlightProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 DocumentHighlightOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentHighlightOptions]"}
}

// ServerCapabilitiesDocumentRangeFormattingProvider the server provides document range formatting.
type ServerCapabilitiesDocumentRangeFormattingProvider struct {
	value any
}

func NewServerCapabilitiesDocumentRangeFormattingProvider[T bool | DocumentRangeFormattingOptions](val T) *ServerCapabilitiesDocumentRangeFormattingProvider {
	return &ServerCapabilitiesDocumentRangeFormattingProvider{
		value: val,
	}
}

func (t ServerCapabilitiesDocumentRangeFormattingProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case DocumentRangeFormattingOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDocumentRangeFormattingProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 DocumentRangeFormattingOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentRangeFormattingOptions]"}
}

// ServerCapabilitiesDocumentSymbolProvider the server provides document symbol support.
type ServerCapabilitiesDocumentSymbolProvider struct {
	value any
}

func NewServerCapabilitiesDocumentSymbolProvider[T bool | DocumentSymbolOptions](val T) *ServerCapabilitiesDocumentSymbolProvider {
	return &ServerCapabilitiesDocumentSymbolProvider{
		value: val,
	}
}

func (t ServerCapabilitiesDocumentSymbolProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case DocumentSymbolOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDocumentSymbolProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 DocumentSymbolOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentSymbolOptions]"}
}

// ServerCapabilitiesFoldingRangeProvider the server provides folding provider support.
type ServerCapabilitiesFoldingRangeProvider struct {
	value any
}

func NewServerCapabilitiesFoldingRangeProvider[T bool | FoldingRangeOptions | FoldingRangeRegistrationOptions](val T) *ServerCapabilitiesFoldingRangeProvider {
	return &ServerCapabilitiesFoldingRangeProvider{
		value: val,
	}
}

func (t ServerCapabilitiesFoldingRangeProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case FoldingRangeOptions:
		return marshal(val)
	case FoldingRangeRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesFoldingRangeProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 FoldingRangeOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 FoldingRangeRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool FoldingRangeOptions FoldingRangeRegistrationOptions]"}
}

// ServerCapabilitiesHoverProvider the server provides hover support.
type ServerCapabilitiesHoverProvider struct {
	value any
}

func NewServerCapabilitiesHoverProvider[T bool | HoverOptions](val T) *ServerCapabilitiesHoverProvider {
	return &ServerCapabilitiesHoverProvider{
		value: val,
	}
}

func (t ServerCapabilitiesHoverProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case HoverOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesHoverProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 HoverOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool HoverOptions]"}
}

// ServerCapabilitiesImplementationProvider the server provides Goto Implementation support.
type ServerCapabilitiesImplementationProvider struct {
	value any
}

func NewServerCapabilitiesImplementationProvider[T bool | ImplementationOptions | ImplementationRegistrationOptions](val T) *ServerCapabilitiesImplementationProvider {
	return &ServerCapabilitiesImplementationProvider{
		value: val,
	}
}

func (t ServerCapabilitiesImplementationProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case ImplementationOptions:
		return marshal(val)
	case ImplementationRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesImplementationProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 ImplementationOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 ImplementationRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool ImplementationOptions ImplementationRegistrationOptions]"}
}

// ServerCapabilitiesInlayHintProvider the server provides inlay hints.
//
// @since 3.17.0
type ServerCapabilitiesInlayHintProvider struct {
	value any
}

func NewServerCapabilitiesInlayHintProvider[T bool | InlayHintOptions | InlayHintRegistrationOptions](val T) *ServerCapabilitiesInlayHintProvider {
	return &ServerCapabilitiesInlayHintProvider{
		value: val,
	}
}

func (t ServerCapabilitiesInlayHintProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case InlayHintOptions:
		return marshal(val)
	case InlayHintRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesInlayHintProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 InlayHintOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 InlayHintRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool InlayHintOptions InlayHintRegistrationOptions]"}
}

// ServerCapabilitiesInlineCompletionProvider inline completion options used during static registration.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type ServerCapabilitiesInlineCompletionProvider struct {
	value any
}

func NewServerCapabilitiesInlineCompletionProvider[T bool | InlineCompletionOptions](val T) *ServerCapabilitiesInlineCompletionProvider {
	return &ServerCapabilitiesInlineCompletionProvider{
		value: val,
	}
}

func (t ServerCapabilitiesInlineCompletionProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case InlineCompletionOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesInlineCompletionProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 InlineCompletionOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool InlineCompletionOptions]"}
}

// ServerCapabilitiesInlineValueProvider the server provides inline values.
//
// @since 3.17.0
type ServerCapabilitiesInlineValueProvider struct {
	value any
}

func NewServerCapabilitiesInlineValueProvider[T bool | InlineValueOptions | InlineValueRegistrationOptions](val T) *ServerCapabilitiesInlineValueProvider {
	return &ServerCapabilitiesInlineValueProvider{
		value: val,
	}
}

func (t ServerCapabilitiesInlineValueProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case InlineValueOptions:
		return marshal(val)
	case InlineValueRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesInlineValueProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 InlineValueOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 InlineValueRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool InlineValueOptions InlineValueRegistrationOptions]"}
}

// ServerCapabilitiesLinkedEditingRangeProvider the server provides linked editing range support.
//
// @since 3.16.0
type ServerCapabilitiesLinkedEditingRangeProvider struct {
	value any
}

func NewServerCapabilitiesLinkedEditingRangeProvider[T bool | LinkedEditingRangeOptions | LinkedEditingRangeRegistrationOptions](val T) *ServerCapabilitiesLinkedEditingRangeProvider {
	return &ServerCapabilitiesLinkedEditingRangeProvider{
		value: val,
	}
}

func (t ServerCapabilitiesLinkedEditingRangeProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case LinkedEditingRangeOptions:
		return marshal(val)
	case LinkedEditingRangeRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesLinkedEditingRangeProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 LinkedEditingRangeOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 LinkedEditingRangeRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool LinkedEditingRangeOptions LinkedEditingRangeRegistrationOptions]"}
}

// ServerCapabilitiesMonikerProvider the server provides moniker support.
//
// @since 3.16.0
type ServerCapabilitiesMonikerProvider struct {
	value any
}

func NewServerCapabilitiesMonikerProvider[T bool | MonikerOptions | MonikerRegistrationOptions](val T) *ServerCapabilitiesMonikerProvider {
	return &ServerCapabilitiesMonikerProvider{
		value: val,
	}
}

func (t ServerCapabilitiesMonikerProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case MonikerOptions:
		return marshal(val)
	case MonikerRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesMonikerProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 MonikerOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 MonikerRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool MonikerOptions MonikerRegistrationOptions]"}
}

// ServerCapabilitiesNotebookDocumentSync defines how notebook documents are synced.
//
// @since 3.17.0
type ServerCapabilitiesNotebookDocumentSync struct {
	value any
}

func NewServerCapabilitiesNotebookDocumentSync[T NotebookDocumentSyncOptions | NotebookDocumentSyncRegistrationOptions](val T) *ServerCapabilitiesNotebookDocumentSync {
	return &ServerCapabilitiesNotebookDocumentSync{
		value: val,
	}
}

func (t ServerCapabilitiesNotebookDocumentSync) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case NotebookDocumentSyncOptions:
		return marshal(val)
	case NotebookDocumentSyncRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesNotebookDocumentSync) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 NotebookDocumentSyncOptions
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 NotebookDocumentSyncRegistrationOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [NotebookDocumentSyncOptions NotebookDocumentSyncRegistrationOptions]"}
}

// ServerCapabilitiesReferencesProvider the server provides find references support.
type ServerCapabilitiesReferencesProvider struct {
	value any
}

func NewServerCapabilitiesReferencesProvider[T bool | ReferenceOptions](val T) *ServerCapabilitiesReferencesProvider {
	return &ServerCapabilitiesReferencesProvider{
		value: val,
	}
}

func (t ServerCapabilitiesReferencesProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case ReferenceOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesReferencesProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 ReferenceOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool ReferenceOptions]"}
}

// ServerCapabilitiesRenameProvider the server provides rename support. RenameOptions may only be specified if the client states that it
// supports `prepareSupport` in its initial `initialize` request.
type ServerCapabilitiesRenameProvider struct {
	value any
}

func NewServerCapabilitiesRenameProvider[T bool | RenameOptions](val T) *ServerCapabilitiesRenameProvider {
	return &ServerCapabilitiesRenameProvider{
		value: val,
	}
}

func (t ServerCapabilitiesRenameProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case RenameOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesRenameProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 RenameOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool RenameOptions]"}
}

// ServerCapabilitiesSelectionRangeProvider the server provides selection range support.
type ServerCapabilitiesSelectionRangeProvider struct {
	value any
}

func NewServerCapabilitiesSelectionRangeProvider[T bool | SelectionRangeOptions | SelectionRangeRegistrationOptions](val T) *ServerCapabilitiesSelectionRangeProvider {
	return &ServerCapabilitiesSelectionRangeProvider{
		value: val,
	}
}

func (t ServerCapabilitiesSelectionRangeProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case SelectionRangeOptions:
		return marshal(val)
	case SelectionRangeRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesSelectionRangeProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 SelectionRangeOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 SelectionRangeRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool SelectionRangeOptions SelectionRangeRegistrationOptions]"}
}

// ServerCapabilitiesSemanticTokensProvider the server provides semantic tokens support.
//
// @since 3.16.0
type ServerCapabilitiesSemanticTokensProvider struct {
	value any
}

func NewServerCapabilitiesSemanticTokensProvider[T SemanticTokensOptions | SemanticTokensRegistrationOptions](val T) *ServerCapabilitiesSemanticTokensProvider {
	return &ServerCapabilitiesSemanticTokensProvider{
		value: val,
	}
}

func (t ServerCapabilitiesSemanticTokensProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case SemanticTokensOptions:
		return marshal(val)
	case SemanticTokensRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesSemanticTokensProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 SemanticTokensOptions
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 SemanticTokensRegistrationOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [SemanticTokensOptions SemanticTokensRegistrationOptions]"}
}

// ServerCapabilitiesTextDocumentSync defines how text documents are synced. Is either a detailed structure defining each notification or for backwards compatibility the TextDocumentSyncKind number.
type ServerCapabilitiesTextDocumentSync struct {
	value any
}

func NewServerCapabilitiesTextDocumentSync[T TextDocumentSyncOptions | TextDocumentSyncKind](val T) *ServerCapabilitiesTextDocumentSync {
	return &ServerCapabilitiesTextDocumentSync{
		value: val,
	}
}

func (t ServerCapabilitiesTextDocumentSync) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case TextDocumentSyncOptions:
		return marshal(val)
	case TextDocumentSyncKind:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesTextDocumentSync) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 TextDocumentSyncOptions
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 TextDocumentSyncKind
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [TextDocumentSyncOptions TextDocumentSyncKind]"}
}

// ServerCapabilitiesTypeDefinitionProvider the server provides Goto Type Definition support.
type ServerCapabilitiesTypeDefinitionProvider struct {
	value any
}

func NewServerCapabilitiesTypeDefinitionProvider[T bool | TypeDefinitionOptions | TypeDefinitionRegistrationOptions](val T) *ServerCapabilitiesTypeDefinitionProvider {
	return &ServerCapabilitiesTypeDefinitionProvider{
		value: val,
	}
}

func (t ServerCapabilitiesTypeDefinitionProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case TypeDefinitionOptions:
		return marshal(val)
	case TypeDefinitionRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesTypeDefinitionProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 TypeDefinitionOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 TypeDefinitionRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool TypeDefinitionOptions TypeDefinitionRegistrationOptions]"}
}

// ServerCapabilitiesTypeHierarchyProvider the server provides type hierarchy support.
//
// @since 3.17.0
type ServerCapabilitiesTypeHierarchyProvider struct {
	value any
}

func NewServerCapabilitiesTypeHierarchyProvider[T bool | TypeHierarchyOptions | TypeHierarchyRegistrationOptions](val T) *ServerCapabilitiesTypeHierarchyProvider {
	return &ServerCapabilitiesTypeHierarchyProvider{
		value: val,
	}
}

func (t ServerCapabilitiesTypeHierarchyProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case TypeHierarchyOptions:
		return marshal(val)
	case TypeHierarchyRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesTypeHierarchyProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 TypeHierarchyOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 TypeHierarchyRegistrationOptions
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool TypeHierarchyOptions TypeHierarchyRegistrationOptions]"}
}

// ServerCapabilitiesWorkspaceSymbolProvider the server provides workspace symbol support.
type ServerCapabilitiesWorkspaceSymbolProvider struct {
	value any
}

func NewServerCapabilitiesWorkspaceSymbolProvider[T bool | WorkspaceSymbolOptions](val T) *ServerCapabilitiesWorkspaceSymbolProvider {
	return &ServerCapabilitiesWorkspaceSymbolProvider{
		value: val,
	}
}

func (t ServerCapabilitiesWorkspaceSymbolProvider) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case WorkspaceSymbolOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesWorkspaceSymbolProvider) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 WorkspaceSymbolOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool WorkspaceSymbolOptions]"}
}

// SignatureInformationDocumentation the human-readable doc-comment of this signature. Will be shown in the UI but can be omitted.
type SignatureInformationDocumentation struct {
	value any
}

func NewSignatureInformationDocumentation[T string | MarkupContent](val T) *SignatureInformationDocumentation {
	return &SignatureInformationDocumentation{
		value: val,
	}
}

func (t SignatureInformationDocumentation) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case MarkupContent:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *SignatureInformationDocumentation) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// TextDocumentEditEdits the edits to be applied. 3.16.0 - support for AnnotatedTextEdit. This is guarded using a client capability. 3.18.0 - support for SnippetTextEdit. This is guarded using a client capability.
type TextDocumentEditEdits struct {
	value any
}

func NewTextDocumentEditEdits[T TextEdit | AnnotatedTextEdit | SnippetTextEdit](val T) *TextDocumentEditEdits {
	return &TextDocumentEditEdits{
		value: val,
	}
}

func (t TextDocumentEditEdits) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case TextEdit:
		return marshal(val)
	case AnnotatedTextEdit:
		return marshal(val)
	case SnippetTextEdit:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentEditEdits) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 TextEdit
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 AnnotatedTextEdit
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 SnippetTextEdit
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [TextEdit AnnotatedTextEdit SnippetTextEdit]"}
}

// TextDocumentSyncOptionsSave if present save notifications are sent to the server. If omitted the notification should not be sent.
type TextDocumentSyncOptionsSave struct {
	value any
}

func NewTextDocumentSyncOptionsSave[T bool | SaveOptions](val T) *TextDocumentSyncOptionsSave {
	return &TextDocumentSyncOptionsSave{
		value: val,
	}
}

func (t TextDocumentSyncOptionsSave) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case bool:
		return marshal(val)
	case SaveOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentSyncOptionsSave) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 SaveOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool SaveOptions]"}
}

// TypeDefinitionResult a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
type TypeDefinitionResult struct {
	value any
}

func NewTypeDefinitionResult[T Definition | []DefinitionLink](val T) *TypeDefinitionResult {
	return &TypeDefinitionResult{
		value: val,
	}
}

func (t TypeDefinitionResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Definition:
		return marshal(val)
	case []DefinitionLink:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TypeDefinitionResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Definition
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DefinitionLink
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Definition []DefinitionLink]"}
}

// TypeDefinitionResult a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
type TypeDefinitionResult struct {
	value any
}

func NewTypeDefinitionResult[T Definition | []DefinitionLink](val T) *TypeDefinitionResult {
	return &TypeDefinitionResult{
		value: val,
	}
}

func (t TypeDefinitionResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Definition:
		return marshal(val)
	case []DefinitionLink:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TypeDefinitionResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Definition
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []DefinitionLink
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Definition []DefinitionLink]"}
}

// WorkspaceEditDocumentChanges depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes are
// either an array of `TextDocumentEdit`s to express changes to n different text documents where each text document edit addresses a specific version of a text document. Or it can contain above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations. Whether a client supports versioned document edits is expressed via `workspace.workspaceEdit.documentChanges` client capability. If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then only plain `TextEdit`s using the `changes` property are supported.
type WorkspaceEditDocumentChanges struct {
	value any
}

func NewWorkspaceEditDocumentChanges[T TextDocumentEdit | CreateFile | RenameFile | DeleteFile](val T) *WorkspaceEditDocumentChanges {
	return &WorkspaceEditDocumentChanges{
		value: val,
	}
}

func (t WorkspaceEditDocumentChanges) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case TextDocumentEdit:
		return marshal(val)
	case CreateFile:
		return marshal(val)
	case RenameFile:
		return marshal(val)
	case DeleteFile:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceEditDocumentChanges) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 TextDocumentEdit
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 CreateFile
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	var h2 RenameFile
	if err := unmarshal(val, &h2); err == nil {
		t.value = h2
		return nil
	}
	var h3 DeleteFile
	if err := unmarshal(val, &h3); err == nil {
		t.value = h3
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [TextDocumentEdit CreateFile RenameFile DeleteFile]"}
}

// WorkspaceFoldersServerCapabilitiesChangeNotifications whether the server wants to receive workspace folder change notifications. If a string is provided the string is treated as an ID under which the notification is registered on the client side. The ID can be used to unregister for these events using the `client/unregisterCapability` request.
type WorkspaceFoldersServerCapabilitiesChangeNotifications struct {
	value any
}

func NewWorkspaceFoldersServerCapabilitiesChangeNotifications[T string | bool](val T) WorkspaceFoldersServerCapabilitiesChangeNotifications {
	return WorkspaceFoldersServerCapabilitiesChangeNotifications{
		value: val,
	}
}

func (t WorkspaceFoldersServerCapabilitiesChangeNotifications) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case string:
		return marshal(val)
	case bool:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceFoldersServerCapabilitiesChangeNotifications) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 bool
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string bool]"}
}

// WorkspaceOptionsTextDocumentContent the server supports the `workspace/textDocumentContent` request.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type WorkspaceOptionsTextDocumentContent struct {
	value any
}

func NewWorkspaceOptionsTextDocumentContent[T TextDocumentContentOptions | TextDocumentContentRegistrationOptions](val T) *WorkspaceOptionsTextDocumentContent {
	return &WorkspaceOptionsTextDocumentContent{
		value: val,
	}
}

func (t WorkspaceOptionsTextDocumentContent) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case TextDocumentContentOptions:
		return marshal(val)
	case TextDocumentContentRegistrationOptions:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceOptionsTextDocumentContent) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 TextDocumentContentOptions
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 TextDocumentContentRegistrationOptions
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [TextDocumentContentOptions TextDocumentContentRegistrationOptions]"}
}

// WorkspaceSymbolLocation the location of the symbol. Whether a server is allowed to return a location without a range depends
// on the client capability `workspace.symbol.resolveSupport`. See SymbolInformation#location for
// more details.
type WorkspaceSymbolLocation struct {
	value any
}

func NewWorkspaceSymbolLocation[T Location | LocationURIOnly](val T) *WorkspaceSymbolLocation {
	return &WorkspaceSymbolLocation{
		value: val,
	}
}

func (t WorkspaceSymbolLocation) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case Location:
		return marshal(val)
	case LocationURIOnly:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceSymbolLocation) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 Location
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 LocationURIOnly
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Location LocationURIOnly]"}
}

// WorkspaceSymbolResult a request to list project-wide symbols matching the query string given by the WorkspaceSymbolParams.
// The response is of type SymbolInformation SymbolInformation[] or a Thenable that resolves to such. 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
//
// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
type WorkspaceSymbolResult struct {
	value any
}

func NewWorkspaceSymbolResult[T []SymbolInformation | []WorkspaceSymbol](val T) WorkspaceSymbolResult {
	return WorkspaceSymbolResult{
		value: val,
	}
}

func (t WorkspaceSymbolResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case []SymbolInformation:
		return marshal(val)
	case []WorkspaceSymbol:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceSymbolResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 []SymbolInformation
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []WorkspaceSymbol
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [[]SymbolInformation []WorkspaceSymbol]"}
}

// WorkspaceSymbolResult a request to list project-wide symbols matching the query string given by the WorkspaceSymbolParams.
// The response is of type SymbolInformation SymbolInformation[] or a Thenable that resolves to such. 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
//
// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
type WorkspaceSymbolResult struct {
	value any
}

func NewWorkspaceSymbolResult[T []SymbolInformation | []WorkspaceSymbol](val T) WorkspaceSymbolResult {
	return WorkspaceSymbolResult{
		value: val,
	}
}

func (t WorkspaceSymbolResult) MarshalJSON() ([]byte, error) {
	switch val := t.value.(type) {
	case []SymbolInformation:
		return marshal(val)
	case []WorkspaceSymbol:
		return marshal(val)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceSymbolResult) UnmarshalJSON(val []byte) error {
	if string(val) == "null" {
		t.value = nil
		return nil
	}
	var h0 []SymbolInformation
	if err := unmarshal(val, &h0); err == nil {
		t.value = h0
		return nil
	}
	var h1 []WorkspaceSymbol
	if err := unmarshal(val, &h1); err == nil {
		t.value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [[]SymbolInformation []WorkspaceSymbol]"}
}
