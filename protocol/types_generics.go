// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"fmt"

	"go.lsp.dev/uri"
)

// CancelParamsID the request id to cancel.
type CancelParamsID struct {
	Value any `json:"value"`
}

func NewCancelParamsID[T int32 | string](x T) CancelParamsID {
	return CancelParamsID{
		Value: x,
	}
}

func (t CancelParamsID) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case int32:
		return marshal(x)
	case string:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CancelParamsID) UnmarshalJSON(x []byte) error {
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
	return &UnmarshalError{"unmarshal failed to match one of [int32 string]"}
}

// ClientSemanticTokensRequestOptionsFull the client will send the `textDocument/semanticTokens/full` request if the server provides a corresponding handler.
//
// @since 3.18.0 proposed
type ClientSemanticTokensRequestOptionsFull struct {
	Value any `json:"value"`
}

func NewClientSemanticTokensRequestOptionsFull[T bool | ClientSemanticTokensRequestFullDelta](x T) ClientSemanticTokensRequestOptionsFull {
	return ClientSemanticTokensRequestOptionsFull{
		Value: x,
	}
}

func (t ClientSemanticTokensRequestOptionsFull) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case ClientSemanticTokensRequestFullDelta:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ClientSemanticTokensRequestOptionsFull) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 ClientSemanticTokensRequestFullDelta
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool ClientSemanticTokensRequestFullDelta]"}
}

// ClientSemanticTokensRequestOptionsRange the client will send the `textDocument/semanticTokens/range` request if the server provides a corresponding handler.
//
// @since 3.18.0 proposed
type ClientSemanticTokensRequestOptionsRange struct {
	Value any `json:"value"`
}

func NewClientSemanticTokensRequestOptionsRange[T bool | Range](x T) ClientSemanticTokensRequestOptionsRange {
	return ClientSemanticTokensRequestOptionsRange{
		Value: x,
	}
}

func (t ClientSemanticTokensRequestOptionsRange) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case Range:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ClientSemanticTokensRequestOptionsRange) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 Range
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool Range]"}
}

// CompletionItemDefaultsEditRange a default edit range.
//
// @since 3.17.0
type CompletionItemDefaultsEditRange struct {
	Value any `json:"value"`
}

func NewCompletionItemDefaultsEditRange[T Range | EditRangeWithInsertReplace](x T) CompletionItemDefaultsEditRange {
	return CompletionItemDefaultsEditRange{
		Value: x,
	}
}

func (t CompletionItemDefaultsEditRange) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Range:
		return marshal(x)
	case EditRangeWithInsertReplace:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CompletionItemDefaultsEditRange) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Range
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 EditRangeWithInsertReplace
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Range EditRangeWithInsertReplace]"}
}

// CompletionItemDocumentation a human-readable string that represents a doc-comment.
type CompletionItemDocumentation struct {
	Value any `json:"value"`
}

func NewCompletionItemDocumentation[T string | MarkupContent](x T) CompletionItemDocumentation {
	return CompletionItemDocumentation{
		Value: x,
	}
}

func (t CompletionItemDocumentation) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case MarkupContent:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CompletionItemDocumentation) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// CompletionItemTextEdit an TextEdit edit which is applied to a document when selecting this completion. When an edit is provided the value of CompletionItem.insertText insertText is ignored. Most editors support two different operations when accepting a completion item. One is to insert a completion text and the other is to replace an existing text with a completion text. Since this can usually not be predetermined by a server it can report both ranges. Clients need to signal support for `InsertReplaceEdits` via the `textDocument.completion.insertReplaceSupport` client capability property. *Note 1:* The text edit's range as well as both ranges from an insert replace edit must be a [single line] and they must contain the position at which completion has been requested. *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range must be a prefix of the edit's replace range, that means it must be contained and starting at the same position. 3.16.0 additional type `InsertReplaceEdit`.
type CompletionItemTextEdit struct {
	Value any `json:"value"`
}

func NewCompletionItemTextEdit[T TextEdit | InsertReplaceEdit](x T) CompletionItemTextEdit {
	return CompletionItemTextEdit{
		Value: x,
	}
}

func (t CompletionItemTextEdit) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case TextEdit:
		return marshal(x)
	case InsertReplaceEdit:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *CompletionItemTextEdit) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 TextEdit
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 InsertReplaceEdit
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [TextEdit InsertReplaceEdit]"}
}

// DiagnosticCode the diagnostic's code, which usually appear in the user interface.
type DiagnosticCode struct {
	Value any `json:"value"`
}

func NewDiagnosticCode[T int32 | string](x T) DiagnosticCode {
	return DiagnosticCode{
		Value: x,
	}
}

func (t DiagnosticCode) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case int32:
		return marshal(x)
	case string:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DiagnosticCode) UnmarshalJSON(x []byte) error {
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
	return &UnmarshalError{"unmarshal failed to match one of [int32 string]"}
}

type DidChangeConfigurationRegistrationOptionsSection struct {
	Value any `json:"value"`
}

func NewDidChangeConfigurationRegistrationOptionsSection[T string | []string](x T) DidChangeConfigurationRegistrationOptionsSection {
	return DidChangeConfigurationRegistrationOptionsSection{
		Value: x,
	}
}

func (t DidChangeConfigurationRegistrationOptionsSection) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case []string:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DidChangeConfigurationRegistrationOptionsSection) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []string
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string []string]"}
}

// @since 3.17.0
type DocumentDiagnosticReportPartialResultRelatedDocuments struct {
	Value any `json:"value"`
}

func NewDocumentDiagnosticReportPartialResultRelatedDocuments[T FullDocumentDiagnosticReport | UnchangedDocumentDiagnosticReport](x T) DocumentDiagnosticReportPartialResultRelatedDocuments {
	return DocumentDiagnosticReportPartialResultRelatedDocuments{
		Value: x,
	}
}

func (t DocumentDiagnosticReportPartialResultRelatedDocuments) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case FullDocumentDiagnosticReport:
		return marshal(x)
	case UnchangedDocumentDiagnosticReport:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *DocumentDiagnosticReportPartialResultRelatedDocuments) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 FullDocumentDiagnosticReport
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 UnchangedDocumentDiagnosticReport
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [FullDocumentDiagnosticReport UnchangedDocumentDiagnosticReport]"}
}

// HoverContents the hover's content.
type HoverContents struct {
	Value any `json:"value"`
}

func NewHoverContents[T MarkupContent | MarkedString | []MarkedString](x T) HoverContents {
	return HoverContents{
		Value: x,
	}
}

func (t HoverContents) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case MarkupContent:
		return marshal(x)
	case MarkedString:
		return marshal(x)
	case []MarkedString:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *HoverContents) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 MarkupContent
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 MarkedString
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 []MarkedString
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [MarkupContent MarkedString []MarkedString]"}
}

// InlayHintLabel the label of this hint. A human readable string or an array of InlayHintLabelPart label parts. *Note* that neither the string nor the label part can be empty.
//
// @since 3.17.0
type InlayHintLabel struct {
	Value any `json:"value"`
}

func NewInlayHintLabel[T string | []InlayHintLabelPart](x T) InlayHintLabel {
	return InlayHintLabel{
		Value: x,
	}
}

func (t InlayHintLabel) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case []InlayHintLabelPart:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlayHintLabel) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []InlayHintLabelPart
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string []InlayHintLabelPart]"}
}

// InlayHintLabelPartTooltip the tooltip text when you hover over this label part. Depending on the client capability `inlayHint.resolveSupport` clients might resolve this property late using the resolve request.
//
// @since 3.17.0
type InlayHintLabelPartTooltip struct {
	Value any `json:"value"`
}

func NewInlayHintLabelPartTooltip[T string | MarkupContent](x T) InlayHintLabelPartTooltip {
	return InlayHintLabelPartTooltip{
		Value: x,
	}
}

func (t InlayHintLabelPartTooltip) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case MarkupContent:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlayHintLabelPartTooltip) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// InlayHintTooltip the tooltip text when you hover over this item.
//
// @since 3.17.0
type InlayHintTooltip struct {
	Value any `json:"value"`
}

func NewInlayHintTooltip[T string | MarkupContent](x T) InlayHintTooltip {
	return InlayHintTooltip{
		Value: x,
	}
}

func (t InlayHintTooltip) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case MarkupContent:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlayHintTooltip) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// InlineCompletionItemInsertText the text to replace the range with. Must be set.
//
// @since 3.18.0 proposed
type InlineCompletionItemInsertText struct {
	Value any `json:"value"`
}

func NewInlineCompletionItemInsertText[T string | StringValue](x T) InlineCompletionItemInsertText {
	return InlineCompletionItemInsertText{
		Value: x,
	}
}

func (t InlineCompletionItemInsertText) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case StringValue:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *InlineCompletionItemInsertText) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 StringValue
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string StringValue]"}
}

// NotebookCellTextDocumentFilterNotebook a filter that matches against the notebook containing the notebook cell. If a string value is provided it matches against the notebook type. '*' matches every notebook.
//
// @since 3.17.0
type NotebookCellTextDocumentFilterNotebook struct {
	Value any `json:"value"`
}

func NewNotebookCellTextDocumentFilterNotebook[T string | NotebookDocumentFilter](x T) NotebookCellTextDocumentFilterNotebook {
	return NotebookCellTextDocumentFilterNotebook{
		Value: x,
	}
}

func (t NotebookCellTextDocumentFilterNotebook) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case NotebookDocumentFilter:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *NotebookCellTextDocumentFilterNotebook) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 NotebookDocumentFilter
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string NotebookDocumentFilter]"}
}

// NotebookDocumentFilterWithCellsNotebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterWithCellsNotebook struct {
	Value any `json:"value"`
}

func NewNotebookDocumentFilterWithCellsNotebook[T string | NotebookDocumentFilter](x T) NotebookDocumentFilterWithCellsNotebook {
	return NotebookDocumentFilterWithCellsNotebook{
		Value: x,
	}
}

func (t NotebookDocumentFilterWithCellsNotebook) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case NotebookDocumentFilter:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *NotebookDocumentFilterWithCellsNotebook) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 NotebookDocumentFilter
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string NotebookDocumentFilter]"}
}

// NotebookDocumentFilterWithNotebookNotebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterWithNotebookNotebook struct {
	Value any `json:"value"`
}

func NewNotebookDocumentFilterWithNotebookNotebook[T string | NotebookDocumentFilter](x T) NotebookDocumentFilterWithNotebookNotebook {
	return NotebookDocumentFilterWithNotebookNotebook{
		Value: x,
	}
}

func (t NotebookDocumentFilterWithNotebookNotebook) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case NotebookDocumentFilter:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *NotebookDocumentFilterWithNotebookNotebook) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 NotebookDocumentFilter
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string NotebookDocumentFilter]"}
}

// NotebookDocumentSyncOptionsNotebookSelector the notebooks to be synced.
//
// @since 3.17.0
type NotebookDocumentSyncOptionsNotebookSelector struct {
	Value any `json:"value"`
}

func NewNotebookDocumentSyncOptionsNotebookSelector[T NotebookDocumentFilterWithNotebook | NotebookDocumentFilterWithCells](x T) NotebookDocumentSyncOptionsNotebookSelector {
	return NotebookDocumentSyncOptionsNotebookSelector{
		Value: x,
	}
}

func (t NotebookDocumentSyncOptionsNotebookSelector) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case NotebookDocumentFilterWithNotebook:
		return marshal(x)
	case NotebookDocumentFilterWithCells:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *NotebookDocumentSyncOptionsNotebookSelector) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 NotebookDocumentFilterWithNotebook
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 NotebookDocumentFilterWithCells
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [NotebookDocumentFilterWithNotebook NotebookDocumentFilterWithCells]"}
}

// ParameterInformationDocumentation the human-readable doc-comment of this parameter. Will be shown in the UI but can be omitted.
type ParameterInformationDocumentation struct {
	Value any `json:"value"`
}

func NewParameterInformationDocumentation[T string | MarkupContent](x T) ParameterInformationDocumentation {
	return ParameterInformationDocumentation{
		Value: x,
	}
}

func (t ParameterInformationDocumentation) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case MarkupContent:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ParameterInformationDocumentation) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// ParameterInformationLabel the label of this parameter information. Either a string or an inclusive start and exclusive end offsets within its containing signature label. (see SignatureInformation.label). The offsets are based on a UTF-16 string representation as `Position` and `Range` does. To avoid ambiguities a server should use the [start, end] offset value instead of using a substring. Whether a client support this is controlled via `labelOffsetSupport` client capability. *Note*: a label of type string should be a substring of its containing signature label. Its intended use case is to highlight the parameter label
// part in the `SignatureInformation.label`.
type ParameterInformationLabel struct {
	Value any `json:"value"`
}

func NewParameterInformationLabel[T string | uint32](x T) ParameterInformationLabel {
	return ParameterInformationLabel{
		Value: x,
	}
}

func (t ParameterInformationLabel) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case uint32:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ParameterInformationLabel) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 uint32
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string uint32]"}
}

// RelatedFullDocumentDiagnosticReportRelatedDocuments diagnostics of related documents. This information is useful in programming languages where code in a file A can generate diagnostics in a file B which A depends on. An example of such a language is C/C++ where marco definitions in a file a.cpp and result in errors in a header file b.hpp.
//
// @since 3.17.0
type RelatedFullDocumentDiagnosticReportRelatedDocuments struct {
	Value any `json:"value"`
}

func NewRelatedFullDocumentDiagnosticReportRelatedDocuments[T FullDocumentDiagnosticReport | UnchangedDocumentDiagnosticReport](x T) RelatedFullDocumentDiagnosticReportRelatedDocuments {
	return RelatedFullDocumentDiagnosticReportRelatedDocuments{
		Value: x,
	}
}

func (t RelatedFullDocumentDiagnosticReportRelatedDocuments) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case FullDocumentDiagnosticReport:
		return marshal(x)
	case UnchangedDocumentDiagnosticReport:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *RelatedFullDocumentDiagnosticReportRelatedDocuments) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 FullDocumentDiagnosticReport
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 UnchangedDocumentDiagnosticReport
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [FullDocumentDiagnosticReport UnchangedDocumentDiagnosticReport]"}
}

// RelatedUnchangedDocumentDiagnosticReportRelatedDocuments diagnostics of related documents. This information is useful in programming languages where code in a file A can generate diagnostics in a file B which A depends on. An example of such a language is C/C++ where marco definitions in a file a.cpp and result in errors in a header file b.hpp.
//
// @since 3.17.0
type RelatedUnchangedDocumentDiagnosticReportRelatedDocuments struct {
	Value any `json:"value"`
}

func NewRelatedUnchangedDocumentDiagnosticReportRelatedDocuments[T FullDocumentDiagnosticReport | UnchangedDocumentDiagnosticReport](x T) RelatedUnchangedDocumentDiagnosticReportRelatedDocuments {
	return RelatedUnchangedDocumentDiagnosticReportRelatedDocuments{
		Value: x,
	}
}

func (t RelatedUnchangedDocumentDiagnosticReportRelatedDocuments) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case FullDocumentDiagnosticReport:
		return marshal(x)
	case UnchangedDocumentDiagnosticReport:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *RelatedUnchangedDocumentDiagnosticReportRelatedDocuments) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 FullDocumentDiagnosticReport
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 UnchangedDocumentDiagnosticReport
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [FullDocumentDiagnosticReport UnchangedDocumentDiagnosticReport]"}
}

// RelativePatternBaseURI a workspace folder or a base URI to which this pattern will be matched against relatively.
//
// @since 3.17.0
type RelativePatternBaseURI struct {
	Value any `json:"value"`
}

func NewRelativePatternBaseURI[T WorkspaceFolder | uri.URI](x T) RelativePatternBaseURI {
	return RelativePatternBaseURI{
		Value: x,
	}
}

func (t RelativePatternBaseURI) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case WorkspaceFolder:
		return marshal(x)
	case uri.URI:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *RelativePatternBaseURI) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 WorkspaceFolder
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 uri.URI
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [WorkspaceFolder uri.URI]"}
}

// SemanticTokensOptionsFull server supports providing semantic tokens for a full document.
//
// @since 3.16.0
type SemanticTokensOptionsFull struct {
	Value any `json:"value"`
}

func NewSemanticTokensOptionsFull[T bool | SemanticTokensFullDelta](x T) SemanticTokensOptionsFull {
	return SemanticTokensOptionsFull{
		Value: x,
	}
}

func (t SemanticTokensOptionsFull) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case SemanticTokensFullDelta:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *SemanticTokensOptionsFull) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 SemanticTokensFullDelta
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool SemanticTokensFullDelta]"}
}

// SemanticTokensOptionsRange server supports providing semantic tokens for a specific range of a document.
//
// @since 3.16.0
type SemanticTokensOptionsRange struct {
	Value any `json:"value"`
}

func NewSemanticTokensOptionsRange[T bool | Range](x T) SemanticTokensOptionsRange {
	return SemanticTokensOptionsRange{
		Value: x,
	}
}

func (t SemanticTokensOptionsRange) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case Range:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *SemanticTokensOptionsRange) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 Range
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool Range]"}
}

// ServerCapabilitiesCallHierarchyProvider the server provides call hierarchy support.
type ServerCapabilitiesCallHierarchyProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesCallHierarchyProvider[T bool | CallHierarchyOptions | CallHierarchyRegistrationOptions](x T) ServerCapabilitiesCallHierarchyProvider {
	return ServerCapabilitiesCallHierarchyProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesCallHierarchyProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case CallHierarchyOptions:
		return marshal(x)
	case CallHierarchyRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesCallHierarchyProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 CallHierarchyOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 CallHierarchyRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool CallHierarchyOptions CallHierarchyRegistrationOptions]"}
}

// ServerCapabilitiesCodeActionProvider the server provides code actions. CodeActionOptions may only be specified if the client states that it supports `codeActionLiteralSupport` in its initial `initialize` request.
type ServerCapabilitiesCodeActionProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesCodeActionProvider[T bool | CodeActionOptions](x T) ServerCapabilitiesCodeActionProvider {
	return ServerCapabilitiesCodeActionProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesCodeActionProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case CodeActionOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesCodeActionProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 CodeActionOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool CodeActionOptions]"}
}

// ServerCapabilitiesColorProvider the server provides color provider support.
type ServerCapabilitiesColorProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesColorProvider[T bool | DocumentColorOptions | DocumentColorRegistrationOptions](x T) ServerCapabilitiesColorProvider {
	return ServerCapabilitiesColorProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesColorProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case DocumentColorOptions:
		return marshal(x)
	case DocumentColorRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesColorProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 DocumentColorOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 DocumentColorRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentColorOptions DocumentColorRegistrationOptions]"}
}

// ServerCapabilitiesDeclarationProvider the server provides Goto Declaration support.
type ServerCapabilitiesDeclarationProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesDeclarationProvider[T bool | DeclarationOptions | DeclarationRegistrationOptions](x T) ServerCapabilitiesDeclarationProvider {
	return ServerCapabilitiesDeclarationProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesDeclarationProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case DeclarationOptions:
		return marshal(x)
	case DeclarationRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDeclarationProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 DeclarationOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 DeclarationRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DeclarationOptions DeclarationRegistrationOptions]"}
}

// ServerCapabilitiesDefinitionProvider the server provides goto definition support.
type ServerCapabilitiesDefinitionProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesDefinitionProvider[T bool | DefinitionOptions](x T) ServerCapabilitiesDefinitionProvider {
	return ServerCapabilitiesDefinitionProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesDefinitionProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case DefinitionOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDefinitionProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 DefinitionOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DefinitionOptions]"}
}

// ServerCapabilitiesDiagnosticProvider the server has support for pull model diagnostics.
type ServerCapabilitiesDiagnosticProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesDiagnosticProvider[T DiagnosticOptions | DiagnosticRegistrationOptions](x T) ServerCapabilitiesDiagnosticProvider {
	return ServerCapabilitiesDiagnosticProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesDiagnosticProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case DiagnosticOptions:
		return marshal(x)
	case DiagnosticRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDiagnosticProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 DiagnosticOptions
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 DiagnosticRegistrationOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [DiagnosticOptions DiagnosticRegistrationOptions]"}
}

// ServerCapabilitiesDocumentFormattingProvider the server provides document formatting.
type ServerCapabilitiesDocumentFormattingProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesDocumentFormattingProvider[T bool | DocumentFormattingOptions](x T) ServerCapabilitiesDocumentFormattingProvider {
	return ServerCapabilitiesDocumentFormattingProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesDocumentFormattingProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case DocumentFormattingOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDocumentFormattingProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 DocumentFormattingOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentFormattingOptions]"}
}

// ServerCapabilitiesDocumentHighlightProvider the server provides document highlight support.
type ServerCapabilitiesDocumentHighlightProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesDocumentHighlightProvider[T bool | DocumentHighlightOptions](x T) ServerCapabilitiesDocumentHighlightProvider {
	return ServerCapabilitiesDocumentHighlightProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesDocumentHighlightProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case DocumentHighlightOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDocumentHighlightProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 DocumentHighlightOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentHighlightOptions]"}
}

// ServerCapabilitiesDocumentRangeFormattingProvider the server provides document range formatting.
type ServerCapabilitiesDocumentRangeFormattingProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesDocumentRangeFormattingProvider[T bool | DocumentRangeFormattingOptions](x T) ServerCapabilitiesDocumentRangeFormattingProvider {
	return ServerCapabilitiesDocumentRangeFormattingProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesDocumentRangeFormattingProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case DocumentRangeFormattingOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDocumentRangeFormattingProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 DocumentRangeFormattingOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentRangeFormattingOptions]"}
}

// ServerCapabilitiesDocumentSymbolProvider the server provides document symbol support.
type ServerCapabilitiesDocumentSymbolProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesDocumentSymbolProvider[T bool | DocumentSymbolOptions](x T) ServerCapabilitiesDocumentSymbolProvider {
	return ServerCapabilitiesDocumentSymbolProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesDocumentSymbolProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case DocumentSymbolOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesDocumentSymbolProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 DocumentSymbolOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool DocumentSymbolOptions]"}
}

// ServerCapabilitiesFoldingRangeProvider the server provides folding provider support.
type ServerCapabilitiesFoldingRangeProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesFoldingRangeProvider[T bool | FoldingRangeOptions | FoldingRangeRegistrationOptions](x T) ServerCapabilitiesFoldingRangeProvider {
	return ServerCapabilitiesFoldingRangeProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesFoldingRangeProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case FoldingRangeOptions:
		return marshal(x)
	case FoldingRangeRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesFoldingRangeProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 FoldingRangeOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 FoldingRangeRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool FoldingRangeOptions FoldingRangeRegistrationOptions]"}
}

// ServerCapabilitiesHoverProvider the server provides hover support.
type ServerCapabilitiesHoverProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesHoverProvider[T bool | HoverOptions](x T) ServerCapabilitiesHoverProvider {
	return ServerCapabilitiesHoverProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesHoverProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case HoverOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesHoverProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 HoverOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool HoverOptions]"}
}

// ServerCapabilitiesImplementationProvider the server provides Goto Implementation support.
type ServerCapabilitiesImplementationProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesImplementationProvider[T bool | ImplementationOptions | ImplementationRegistrationOptions](x T) ServerCapabilitiesImplementationProvider {
	return ServerCapabilitiesImplementationProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesImplementationProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case ImplementationOptions:
		return marshal(x)
	case ImplementationRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesImplementationProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 ImplementationOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 ImplementationRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool ImplementationOptions ImplementationRegistrationOptions]"}
}

// ServerCapabilitiesInlayHintProvider the server provides inlay hints.
type ServerCapabilitiesInlayHintProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesInlayHintProvider[T bool | InlayHintOptions | InlayHintRegistrationOptions](x T) ServerCapabilitiesInlayHintProvider {
	return ServerCapabilitiesInlayHintProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesInlayHintProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case InlayHintOptions:
		return marshal(x)
	case InlayHintRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesInlayHintProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 InlayHintOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 InlayHintRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool InlayHintOptions InlayHintRegistrationOptions]"}
}

// ServerCapabilitiesInlineCompletionProvider inline completion options used during static registration.  3.18.0 @proposed.
type ServerCapabilitiesInlineCompletionProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesInlineCompletionProvider[T bool | InlineCompletionOptions](x T) ServerCapabilitiesInlineCompletionProvider {
	return ServerCapabilitiesInlineCompletionProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesInlineCompletionProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case InlineCompletionOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesInlineCompletionProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 InlineCompletionOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool InlineCompletionOptions]"}
}

// ServerCapabilitiesInlineValueProvider the server provides inline values.
type ServerCapabilitiesInlineValueProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesInlineValueProvider[T bool | InlineValueOptions | InlineValueRegistrationOptions](x T) ServerCapabilitiesInlineValueProvider {
	return ServerCapabilitiesInlineValueProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesInlineValueProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case InlineValueOptions:
		return marshal(x)
	case InlineValueRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesInlineValueProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 InlineValueOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 InlineValueRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool InlineValueOptions InlineValueRegistrationOptions]"}
}

// ServerCapabilitiesLinkedEditingRangeProvider the server provides linked editing range support.
type ServerCapabilitiesLinkedEditingRangeProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesLinkedEditingRangeProvider[T bool | LinkedEditingRangeOptions | LinkedEditingRangeRegistrationOptions](x T) ServerCapabilitiesLinkedEditingRangeProvider {
	return ServerCapabilitiesLinkedEditingRangeProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesLinkedEditingRangeProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case LinkedEditingRangeOptions:
		return marshal(x)
	case LinkedEditingRangeRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesLinkedEditingRangeProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 LinkedEditingRangeOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 LinkedEditingRangeRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool LinkedEditingRangeOptions LinkedEditingRangeRegistrationOptions]"}
}

// ServerCapabilitiesMonikerProvider the server provides moniker support.
type ServerCapabilitiesMonikerProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesMonikerProvider[T bool | MonikerOptions | MonikerRegistrationOptions](x T) ServerCapabilitiesMonikerProvider {
	return ServerCapabilitiesMonikerProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesMonikerProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case MonikerOptions:
		return marshal(x)
	case MonikerRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesMonikerProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 MonikerOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 MonikerRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool MonikerOptions MonikerRegistrationOptions]"}
}

// ServerCapabilitiesNotebookDocumentSync defines how notebook documents are synced.
type ServerCapabilitiesNotebookDocumentSync struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesNotebookDocumentSync[T NotebookDocumentSyncOptions | NotebookDocumentSyncRegistrationOptions](x T) ServerCapabilitiesNotebookDocumentSync {
	return ServerCapabilitiesNotebookDocumentSync{
		Value: x,
	}
}

func (t ServerCapabilitiesNotebookDocumentSync) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case NotebookDocumentSyncOptions:
		return marshal(x)
	case NotebookDocumentSyncRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesNotebookDocumentSync) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 NotebookDocumentSyncOptions
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 NotebookDocumentSyncRegistrationOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [NotebookDocumentSyncOptions NotebookDocumentSyncRegistrationOptions]"}
}

// ServerCapabilitiesReferencesProvider the server provides find references support.
type ServerCapabilitiesReferencesProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesReferencesProvider[T bool | ReferenceOptions](x T) ServerCapabilitiesReferencesProvider {
	return ServerCapabilitiesReferencesProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesReferencesProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case ReferenceOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesReferencesProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 ReferenceOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool ReferenceOptions]"}
}

// ServerCapabilitiesRenameProvider the server provides rename support. RenameOptions may only be specified if the client states that it
// supports `prepareSupport` in its initial `initialize` request.
type ServerCapabilitiesRenameProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesRenameProvider[T bool | RenameOptions](x T) ServerCapabilitiesRenameProvider {
	return ServerCapabilitiesRenameProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesRenameProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case RenameOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesRenameProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 RenameOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool RenameOptions]"}
}

// ServerCapabilitiesSelectionRangeProvider the server provides selection range support.
type ServerCapabilitiesSelectionRangeProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesSelectionRangeProvider[T bool | SelectionRangeOptions | SelectionRangeRegistrationOptions](x T) ServerCapabilitiesSelectionRangeProvider {
	return ServerCapabilitiesSelectionRangeProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesSelectionRangeProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case SelectionRangeOptions:
		return marshal(x)
	case SelectionRangeRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesSelectionRangeProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 SelectionRangeOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 SelectionRangeRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool SelectionRangeOptions SelectionRangeRegistrationOptions]"}
}

// ServerCapabilitiesSemanticTokensProvider the server provides semantic tokens support.
type ServerCapabilitiesSemanticTokensProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesSemanticTokensProvider[T SemanticTokensOptions | SemanticTokensRegistrationOptions](x T) ServerCapabilitiesSemanticTokensProvider {
	return ServerCapabilitiesSemanticTokensProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesSemanticTokensProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case SemanticTokensOptions:
		return marshal(x)
	case SemanticTokensRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesSemanticTokensProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 SemanticTokensOptions
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 SemanticTokensRegistrationOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [SemanticTokensOptions SemanticTokensRegistrationOptions]"}
}

// ServerCapabilitiesTextDocumentSync defines how text documents are synced. Is either a detailed structure defining each notification or for backwards compatibility the TextDocumentSyncKind number.
type ServerCapabilitiesTextDocumentSync struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesTextDocumentSync[T TextDocumentSyncOptions | TextDocumentSyncKind](x T) ServerCapabilitiesTextDocumentSync {
	return ServerCapabilitiesTextDocumentSync{
		Value: x,
	}
}

func (t ServerCapabilitiesTextDocumentSync) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case TextDocumentSyncOptions:
		return marshal(x)
	case TextDocumentSyncKind:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesTextDocumentSync) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 TextDocumentSyncOptions
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 TextDocumentSyncKind
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [TextDocumentSyncOptions TextDocumentSyncKind]"}
}

// ServerCapabilitiesTypeDefinitionProvider the server provides Goto Type Definition support.
type ServerCapabilitiesTypeDefinitionProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesTypeDefinitionProvider[T bool | TypeDefinitionOptions | TypeDefinitionRegistrationOptions](x T) ServerCapabilitiesTypeDefinitionProvider {
	return ServerCapabilitiesTypeDefinitionProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesTypeDefinitionProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case TypeDefinitionOptions:
		return marshal(x)
	case TypeDefinitionRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesTypeDefinitionProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 TypeDefinitionOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 TypeDefinitionRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool TypeDefinitionOptions TypeDefinitionRegistrationOptions]"}
}

// ServerCapabilitiesTypeHierarchyProvider the server provides type hierarchy support.
type ServerCapabilitiesTypeHierarchyProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesTypeHierarchyProvider[T bool | TypeHierarchyOptions | TypeHierarchyRegistrationOptions](x T) ServerCapabilitiesTypeHierarchyProvider {
	return ServerCapabilitiesTypeHierarchyProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesTypeHierarchyProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case TypeHierarchyOptions:
		return marshal(x)
	case TypeHierarchyRegistrationOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesTypeHierarchyProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 TypeHierarchyOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 TypeHierarchyRegistrationOptions
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool TypeHierarchyOptions TypeHierarchyRegistrationOptions]"}
}

// ServerCapabilitiesWorkspaceSymbolProvider the server provides workspace symbol support.
type ServerCapabilitiesWorkspaceSymbolProvider struct {
	Value any `json:"value"`
}

func NewServerCapabilitiesWorkspaceSymbolProvider[T bool | WorkspaceSymbolOptions](x T) ServerCapabilitiesWorkspaceSymbolProvider {
	return ServerCapabilitiesWorkspaceSymbolProvider{
		Value: x,
	}
}

func (t ServerCapabilitiesWorkspaceSymbolProvider) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case WorkspaceSymbolOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *ServerCapabilitiesWorkspaceSymbolProvider) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 WorkspaceSymbolOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool WorkspaceSymbolOptions]"}
}

// SignatureInformationDocumentation the human-readable doc-comment of this signature. Will be shown in the UI but can be omitted.
type SignatureInformationDocumentation struct {
	Value any `json:"value"`
}

func NewSignatureInformationDocumentation[T string | MarkupContent](x T) SignatureInformationDocumentation {
	return SignatureInformationDocumentation{
		Value: x,
	}
}

func (t SignatureInformationDocumentation) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case MarkupContent:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *SignatureInformationDocumentation) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 MarkupContent
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string MarkupContent]"}
}

// TextDocumentCodeActionResult a request to provide commands for the given text document and range.
type TextDocumentCodeActionResult struct {
	Value any `json:"value"`
}

func NewTextDocumentCodeActionResult[T Command | CodeAction](x T) TextDocumentCodeActionResult {
	return TextDocumentCodeActionResult{
		Value: x,
	}
}

func (t TextDocumentCodeActionResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Command:
		return marshal(x)
	case CodeAction:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentCodeActionResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Command
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 CodeAction
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Command CodeAction]"}
}

// TextDocumentCompletionResult request to request completion at a given text document position. The request's parameter is of type TextDocumentPosition the response is of type CompletionItem CompletionItem[] or CompletionList or a Thenable that resolves to such. The request can delay the computation of the CompletionItem.detail `detail` and CompletionItem.documentation `documentation` properties to the `completionItem/resolve` request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
type TextDocumentCompletionResult struct {
	Value any `json:"value"`
}

func NewTextDocumentCompletionResult[T []CompletionItem | CompletionList](x T) TextDocumentCompletionResult {
	return TextDocumentCompletionResult{
		Value: x,
	}
}

func (t TextDocumentCompletionResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case []CompletionItem:
		return marshal(x)
	case CompletionList:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentCompletionResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 []CompletionItem
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 CompletionList
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [[]CompletionItem CompletionList]"}
}

// TextDocumentDeclarationResult a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Declaration or a
// typed array of DeclarationLink or a Thenable that resolves to such.
type TextDocumentDeclarationResult struct {
	Value any `json:"value"`
}

func NewTextDocumentDeclarationResult[T Declaration | []DeclarationLink](x T) TextDocumentDeclarationResult {
	return TextDocumentDeclarationResult{
		Value: x,
	}
}

func (t TextDocumentDeclarationResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Declaration:
		return marshal(x)
	case []DeclarationLink:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentDeclarationResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Declaration
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []DeclarationLink
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Declaration []DeclarationLink]"}
}

// TextDocumentDefinitionResult a request to resolve the definition location of a symbol at a given text document position. The request's parameter is of type TextDocumentPosition the response is of either type Definition or a typed
// array of DefinitionLink or a Thenable that resolves to such.
type TextDocumentDefinitionResult struct {
	Value any `json:"value"`
}

func NewTextDocumentDefinitionResult[T Definition | []DefinitionLink](x T) TextDocumentDefinitionResult {
	return TextDocumentDefinitionResult{
		Value: x,
	}
}

func (t TextDocumentDefinitionResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Definition:
		return marshal(x)
	case []DefinitionLink:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentDefinitionResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Definition
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []DefinitionLink
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Definition []DefinitionLink]"}
}

// TextDocumentDocumentSymbolResult a request to list all symbols found in a given text document. The request's parameter is of type TextDocumentIdentifier the response is of type SymbolInformation SymbolInformation[] or a Thenable that
// resolves to such.
type TextDocumentDocumentSymbolResult struct {
	Value any `json:"value"`
}

func NewTextDocumentDocumentSymbolResult[T []SymbolInformation | []DocumentSymbol](x T) TextDocumentDocumentSymbolResult {
	return TextDocumentDocumentSymbolResult{
		Value: x,
	}
}

func (t TextDocumentDocumentSymbolResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case []SymbolInformation:
		return marshal(x)
	case []DocumentSymbol:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentDocumentSymbolResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 []SymbolInformation
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []DocumentSymbol
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [[]SymbolInformation []DocumentSymbol]"}
}

// TextDocumentEditEdits the edits to be applied. 3.16.0 - support for AnnotatedTextEdit. This is guarded using a client capability.
type TextDocumentEditEdits struct {
	Value any `json:"value"`
}

func NewTextDocumentEditEdits[T TextEdit | AnnotatedTextEdit](x T) TextDocumentEditEdits {
	return TextDocumentEditEdits{
		Value: x,
	}
}

func (t TextDocumentEditEdits) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case TextEdit:
		return marshal(x)
	case AnnotatedTextEdit:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentEditEdits) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 TextEdit
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 AnnotatedTextEdit
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [TextEdit AnnotatedTextEdit]"}
}

// TextDocumentImplementationResult a request to resolve the implementation locations of a symbol at a given text document position. The
// request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
type TextDocumentImplementationResult struct {
	Value any `json:"value"`
}

func NewTextDocumentImplementationResult[T Definition | []DefinitionLink](x T) TextDocumentImplementationResult {
	return TextDocumentImplementationResult{
		Value: x,
	}
}

func (t TextDocumentImplementationResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Definition:
		return marshal(x)
	case []DefinitionLink:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentImplementationResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Definition
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []DefinitionLink
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Definition []DefinitionLink]"}
}

// TextDocumentInlineCompletionResult a request to provide inline completions in a document. The request's parameter is of type InlineCompletionParams, the response is of type InlineCompletion InlineCompletion[] or a Thenable that resolves to such. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
type TextDocumentInlineCompletionResult struct {
	Value any `json:"value"`
}

func NewTextDocumentInlineCompletionResult[T InlineCompletionList | []InlineCompletionItem](x T) TextDocumentInlineCompletionResult {
	return TextDocumentInlineCompletionResult{
		Value: x,
	}
}

func (t TextDocumentInlineCompletionResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case InlineCompletionList:
		return marshal(x)
	case []InlineCompletionItem:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentInlineCompletionResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 InlineCompletionList
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []InlineCompletionItem
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [InlineCompletionList []InlineCompletionItem]"}
}

// TextDocumentSemanticTokensFullDeltaResult.
//
// @since 3.16.0
type TextDocumentSemanticTokensFullDeltaResult struct {
	Value any `json:"value"`
}

func NewTextDocumentSemanticTokensFullDeltaResult[T SemanticTokens | SemanticTokensDelta](x T) TextDocumentSemanticTokensFullDeltaResult {
	return TextDocumentSemanticTokensFullDeltaResult{
		Value: x,
	}
}

func (t TextDocumentSemanticTokensFullDeltaResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case SemanticTokens:
		return marshal(x)
	case SemanticTokensDelta:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentSemanticTokensFullDeltaResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 SemanticTokens
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 SemanticTokensDelta
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [SemanticTokens SemanticTokensDelta]"}
}

// TextDocumentSyncOptionsSave if present save notifications are sent to the server. If omitted the notification should not be sent.
type TextDocumentSyncOptionsSave struct {
	Value any `json:"value"`
}

func NewTextDocumentSyncOptionsSave[T bool | SaveOptions](x T) TextDocumentSyncOptionsSave {
	return TextDocumentSyncOptionsSave{
		Value: x,
	}
}

func (t TextDocumentSyncOptionsSave) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case bool:
		return marshal(x)
	case SaveOptions:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentSyncOptionsSave) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 bool
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 SaveOptions
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [bool SaveOptions]"}
}

// TextDocumentTypeDefinitionResult a request to resolve the type definition locations of a symbol at a given text document position. The request's parameter is of type TextDocumentPositionParams the response is of type Definition or a Thenable that resolves to such.
type TextDocumentTypeDefinitionResult struct {
	Value any `json:"value"`
}

func NewTextDocumentTypeDefinitionResult[T Definition | []DefinitionLink](x T) TextDocumentTypeDefinitionResult {
	return TextDocumentTypeDefinitionResult{
		Value: x,
	}
}

func (t TextDocumentTypeDefinitionResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Definition:
		return marshal(x)
	case []DefinitionLink:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *TextDocumentTypeDefinitionResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Definition
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []DefinitionLink
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Definition []DefinitionLink]"}
}

// WorkspaceEditDocumentChanges depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes are
// either an array of `TextDocumentEdit`s to express changes to n different text documents where each text document edit addresses a specific version of a text document. Or it can contain above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations. Whether a client supports versioned document edits is expressed via `workspace.workspaceEdit.documentChanges` client capability. If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then only plain `TextEdit`s using the `changes` property are supported.
type WorkspaceEditDocumentChanges struct {
	Value any `json:"value"`
}

func NewWorkspaceEditDocumentChanges[T TextDocumentEdit | CreateFile | RenameFile | DeleteFile](x T) WorkspaceEditDocumentChanges {
	return WorkspaceEditDocumentChanges{
		Value: x,
	}
}

func (t WorkspaceEditDocumentChanges) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case TextDocumentEdit:
		return marshal(x)
	case CreateFile:
		return marshal(x)
	case RenameFile:
		return marshal(x)
	case DeleteFile:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceEditDocumentChanges) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 TextDocumentEdit
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 CreateFile
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 RenameFile
	if err := unmarshal(x, &h2); err == nil {
		t.Value = h2
		return nil
	}
	var h3 DeleteFile
	if err := unmarshal(x, &h3); err == nil {
		t.Value = h3
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [TextDocumentEdit CreateFile RenameFile DeleteFile]"}
}

// WorkspaceFoldersServerCapabilitiesChangeNotifications whether the server wants to receive workspace folder change notifications. If a string is provided the string is treated as an ID under which the notification is registered on the client side. The ID can be used to unregister for these events using the `client/unregisterCapability` request.
type WorkspaceFoldersServerCapabilitiesChangeNotifications struct {
	Value any `json:"value"`
}

func NewWorkspaceFoldersServerCapabilitiesChangeNotifications[T string | bool](x T) WorkspaceFoldersServerCapabilitiesChangeNotifications {
	return WorkspaceFoldersServerCapabilitiesChangeNotifications{
		Value: x,
	}
}

func (t WorkspaceFoldersServerCapabilitiesChangeNotifications) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case string:
		return marshal(x)
	case bool:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceFoldersServerCapabilitiesChangeNotifications) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 string
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 bool
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [string bool]"}
}

// WorkspaceSymbolLocation the location of the symbol. Whether a server is allowed to return a location without a range depends
// on the client capability `workspace.symbol.resolveSupport`. See SymbolInformation#location for
// more details.
//
// @since 3.17.0
type WorkspaceSymbolLocation struct {
	Value any `json:"value"`
}

func NewWorkspaceSymbolLocation[T Location | LocationURIOnly](x T) WorkspaceSymbolLocation {
	return WorkspaceSymbolLocation{
		Value: x,
	}
}

func (t WorkspaceSymbolLocation) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case Location:
		return marshal(x)
	case LocationURIOnly:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceSymbolLocation) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 Location
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 LocationURIOnly
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [Location LocationURIOnly]"}
}

// WorkspaceSymbolResult a request to list project-wide symbols matching the query string given by the WorkspaceSymbolParams.
// The response is of type SymbolInformation SymbolInformation[] or a Thenable that resolves to such. 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
//
// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients need to advertise support for WorkspaceSymbols via the client capability `workspace.symbol.resolveSupport`.
type WorkspaceSymbolResult struct {
	Value any `json:"value"`
}

func NewWorkspaceSymbolResult[T []SymbolInformation | []WorkspaceSymbol](x T) WorkspaceSymbolResult {
	return WorkspaceSymbolResult{
		Value: x,
	}
}

func (t WorkspaceSymbolResult) MarshalJSON() ([]byte, error) {
	switch x := t.Value.(type) {
	case []SymbolInformation:
		return marshal(x)
	case []WorkspaceSymbol:
		return marshal(x)
	case nil:
		return []byte("null"), nil
	}
	return nil, fmt.Errorf("unknown type: %T", t)
}

func (t *WorkspaceSymbolResult) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 []SymbolInformation
	if err := unmarshal(x, &h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 []WorkspaceSymbol
	if err := unmarshal(x, &h1); err == nil {
		t.Value = h1
		return nil
	}
	return &UnmarshalError{"unmarshal failed to match one of [[]SymbolInformation []WorkspaceSymbol]"}
}
