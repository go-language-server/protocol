// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "slices"

func appendCompletionResultJSON(dst []byte, x *CompletionResult) ([]byte, error) {
	if x == nil || *x == nil {
		return append(dst, "null"...), nil
	}
	switch v := (*x).(type) {
	case CompletionItemSlice:
		return v.appendLSPJSON(dst)
	case *CompletionList:
		return v.appendLSPJSON(dst)
	default:
		return appendJSONMarshal(dst, *x)
	}
}

func (x CompletionItemSlice) appendLSPJSON(dst []byte) ([]byte, error) {
	dst = slices.Grow(dst, 2+len(x)*160)
	dst = append(dst, '[')
	for i := range x {
		if i > 0 {
			dst = append(dst, ',')
		}
		var err error
		dst, err = (&x[i]).appendLSPJSON(dst)
		if err != nil {
			return nil, err
		}
	}
	return append(dst, ']'), nil
}

func (x *CompletionList) appendLSPJSON(dst []byte) ([]byte, error) {
	if x == nil {
		return append(dst, "null"...), nil
	}
	dst = slices.Grow(dst, 64+len(x.Items)*160)
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "isIncomplete")
	dst = appendBoolJSON(dst, x.IsIncomplete)
	if x.ItemDefaults != nil {
		dst = appendObjectName(dst, &first, "itemDefaults")
		var err error
		dst, err = appendJSONMarshal(dst, x.ItemDefaults)
		if err != nil {
			return nil, err
		}
	}
	if x.ApplyKind != nil {
		dst = appendObjectName(dst, &first, "applyKind")
		var err error
		dst, err = appendJSONMarshal(dst, x.ApplyKind)
		if err != nil {
			return nil, err
		}
	}
	dst = appendObjectName(dst, &first, "items")
	var err error
	dst, err = CompletionItemSlice(x.Items).appendLSPJSON(dst)
	if err != nil {
		return nil, err
	}
	return append(dst, '}'), nil
}

func (x *CompletionItem) appendLSPJSON(dst []byte) ([]byte, error) {
	if x == nil {
		return append(dst, "null"...), nil
	}
	dst = slices.Grow(dst, 192)
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "label")
	dst = appendJSONString(dst, x.Label)
	if x.LabelDetails != nil {
		dst = appendObjectName(dst, &first, "labelDetails")
		var err error
		dst, err = appendJSONMarshal(dst, x.LabelDetails)
		if err != nil {
			return nil, err
		}
	}
	if x.Kind != 0 {
		dst = appendObjectName(dst, &first, "kind")
		dst = appendUint32JSON(dst, uint32(x.Kind))
	}
	if len(x.Tags) > 0 {
		dst = appendObjectName(dst, &first, "tags")
		var err error
		dst, err = appendJSONMarshal(dst, x.Tags)
		if err != nil {
			return nil, err
		}
	}
	if v, ok := x.Detail.Get(); ok {
		dst = appendObjectName(dst, &first, "detail")
		dst = appendJSONString(dst, v)
	}
	if x.Documentation != nil {
		dst = appendObjectName(dst, &first, "documentation")
		if v, ok := x.Documentation.(String); ok {
			dst = appendJSONString(dst, string(v))
		} else {
			var err error
			dst, err = appendJSONMarshal(dst, x.Documentation)
			if err != nil {
				return nil, err
			}
		}
	}
	if v, ok := x.Deprecated.Get(); ok {
		dst = appendObjectName(dst, &first, "deprecated")
		dst = appendBoolJSON(dst, v)
	}
	if v, ok := x.Preselect.Get(); ok {
		dst = appendObjectName(dst, &first, "preselect")
		dst = appendBoolJSON(dst, v)
	}
	if v, ok := x.SortText.Get(); ok {
		dst = appendObjectName(dst, &first, "sortText")
		dst = appendJSONString(dst, v)
	}
	if v, ok := x.FilterText.Get(); ok {
		dst = appendObjectName(dst, &first, "filterText")
		dst = appendJSONString(dst, v)
	}
	if v, ok := x.InsertText.Get(); ok {
		dst = appendObjectName(dst, &first, "insertText")
		dst = appendJSONString(dst, v)
	}
	if x.InsertTextFormat != 0 {
		dst = appendObjectName(dst, &first, "insertTextFormat")
		dst = appendUint32JSON(dst, uint32(x.InsertTextFormat))
	}
	if x.InsertTextMode != 0 {
		dst = appendObjectName(dst, &first, "insertTextMode")
		dst = appendUint32JSON(dst, uint32(x.InsertTextMode))
	}
	if x.TextEdit != nil {
		dst = appendObjectName(dst, &first, "textEdit")
		var err error
		dst, err = appendJSONMarshal(dst, x.TextEdit)
		if err != nil {
			return nil, err
		}
	}
	if v, ok := x.TextEditText.Get(); ok {
		dst = appendObjectName(dst, &first, "textEditText")
		dst = appendJSONString(dst, v)
	}
	if len(x.AdditionalTextEdits) > 0 {
		dst = appendObjectName(dst, &first, "additionalTextEdits")
		var err error
		dst, err = appendJSONMarshal(dst, x.AdditionalTextEdits)
		if err != nil {
			return nil, err
		}
	}
	if len(x.CommitCharacters) > 0 {
		dst = appendObjectName(dst, &first, "commitCharacters")
		var err error
		dst, err = appendJSONMarshal(dst, x.CommitCharacters)
		if err != nil {
			return nil, err
		}
	}
	if !isZeroCommand(x.Command) {
		dst = appendObjectName(dst, &first, "command")
		var err error
		dst, err = appendJSONMarshal(dst, x.Command)
		if err != nil {
			return nil, err
		}
	}
	if len(x.Data) > 0 {
		dst = appendObjectName(dst, &first, "data")
		var err error
		dst, err = appendRawJSONValue(dst, x.Data)
		if err != nil {
			return nil, err
		}
	}
	return append(dst, '}'), nil
}

func (x *PublishDiagnosticsParams) appendLSPJSON(dst []byte) ([]byte, error) {
	if x == nil {
		return append(dst, "null"...), nil
	}
	dst = slices.Grow(dst, 64+len(x.Diagnostics)*200)
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "uri")
	dst = appendJSONString(dst, string(x.URI))
	if v, ok := x.Version.Get(); ok {
		dst = appendObjectName(dst, &first, "version")
		dst = appendInt32JSON(dst, v)
	}
	dst = appendObjectName(dst, &first, "diagnostics")
	var err error
	dst, err = appendDiagnosticSliceJSON(dst, x.Diagnostics)
	if err != nil {
		return nil, err
	}
	return append(dst, '}'), nil
}

func appendDiagnosticSliceJSON(dst []byte, x []Diagnostic) ([]byte, error) {
	dst = slices.Grow(dst, 2+len(x)*200)
	dst = append(dst, '[')
	for i := range x {
		if i > 0 {
			dst = append(dst, ',')
		}
		var err error
		dst, err = (&x[i]).appendLSPJSON(dst)
		if err != nil {
			return nil, err
		}
	}
	return append(dst, ']'), nil
}

func (x *Diagnostic) appendLSPJSON(dst []byte) ([]byte, error) {
	if x == nil {
		return append(dst, "null"...), nil
	}
	dst = slices.Grow(dst, 200)
	dst = slices.Grow(dst, 128)
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "range")
	dst = appendRangeJSON(dst, x.Range)
	if x.Severity != 0 {
		dst = appendObjectName(dst, &first, "severity")
		dst = appendUint32JSON(dst, uint32(x.Severity))
	}
	if x.Code != nil {
		dst = appendObjectName(dst, &first, "code")
		var err error
		dst, err = appendProgressTokenJSON(dst, x.Code)
		if err != nil {
			return nil, err
		}
	}
	if x.CodeDescription != (CodeDescription{}) {
		dst = appendObjectName(dst, &first, "codeDescription")
		var err error
		dst, err = appendJSONMarshal(dst, x.CodeDescription)
		if err != nil {
			return nil, err
		}
	}
	if v, ok := x.Source.Get(); ok {
		dst = appendObjectName(dst, &first, "source")
		dst = appendJSONString(dst, v)
	}
	dst = appendObjectName(dst, &first, "message")
	if v, ok := x.Message.(String); ok {
		dst = appendJSONString(dst, string(v))
	} else if v, ok := x.Message.(*String); ok {
		if v == nil {
			dst = append(dst, "null"...)
		} else {
			dst = appendJSONString(dst, string(*v))
		}
	} else {
		var err error
		dst, err = appendJSONMarshal(dst, x.Message)
		if err != nil {
			return nil, err
		}
	}
	if !x.Tags.IsZero() {
		dst = appendObjectName(dst, &first, "tags")
		dst = appendDiagnosticTagsJSON(dst, x.Tags)
	}
	if len(x.RelatedInformation) > 0 {
		dst = appendObjectName(dst, &first, "relatedInformation")
		var err error
		dst, err = appendJSONMarshal(dst, x.RelatedInformation)
		if err != nil {
			return nil, err
		}
	}
	if len(x.Data) > 0 {
		dst = appendObjectName(dst, &first, "data")
		var err error
		dst, err = appendRawJSONValue(dst, x.Data)
		if err != nil {
			return nil, err
		}
	}
	return append(dst, '}'), nil
}

func appendWorkspaceSymbolResultJSON(dst []byte, x *WorkspaceSymbolResult) ([]byte, error) {
	if x == nil || *x == nil {
		return append(dst, "null"...), nil
	}
	switch v := (*x).(type) {
	case WorkspaceSymbolSlice:
		return v.appendLSPJSON(dst)
	case SymbolInformationSlice:
		return v.appendLSPJSON(dst)
	default:
		return appendJSONMarshal(dst, *x)
	}
}

func (x WorkspaceSymbolSlice) appendLSPJSON(dst []byte) ([]byte, error) {
	dst = slices.Grow(dst, 2+len(x)*128)
	dst = append(dst, '[')
	for i := range x {
		if i > 0 {
			dst = append(dst, ',')
		}
		var err error
		dst, err = (&x[i]).appendLSPJSON(dst)
		if err != nil {
			return nil, err
		}
	}
	return append(dst, ']'), nil
}

func (x SymbolInformationSlice) appendLSPJSON(dst []byte) ([]byte, error) {
	dst = slices.Grow(dst, 2+len(x)*176)
	dst = append(dst, '[')
	for i := range x {
		if i > 0 {
			dst = append(dst, ',')
		}
		var err error
		dst, err = (&x[i]).appendLSPJSON(dst)
		if err != nil {
			return nil, err
		}
	}
	return append(dst, ']'), nil
}

func (x *WorkspaceSymbol) appendLSPJSON(dst []byte) ([]byte, error) {
	if x == nil {
		return append(dst, "null"...), nil
	}
	dst = slices.Grow(dst, 128)
	dst = append(dst, '{')
	first := true
	dst = appendBaseSymbolInformationFields(dst, &first, x.BaseSymbolInformation)
	dst = appendObjectName(dst, &first, "location")
	var err error
	dst, err = appendWorkspaceSymbolLocationJSON(dst, x.Location)
	if err != nil {
		return nil, err
	}
	if len(x.Data) > 0 {
		dst = appendObjectName(dst, &first, "data")
		var err error
		dst, err = appendRawJSONValue(dst, x.Data)
		if err != nil {
			return nil, err
		}
	}
	return append(dst, '}'), nil
}

func (x *SymbolInformation) appendLSPJSON(dst []byte) ([]byte, error) {
	if x == nil {
		return append(dst, "null"...), nil
	}
	dst = append(dst, '{')
	first := true
	dst = appendBaseSymbolInformationFields(dst, &first, x.BaseSymbolInformation)
	if x.Deprecated != nil {
		dst = appendObjectName(dst, &first, "deprecated")
		dst = appendBoolJSON(dst, *x.Deprecated)
	}
	dst = appendObjectName(dst, &first, "location")
	dst = appendLocationJSON(dst, x.Location)
	return append(dst, '}'), nil
}

func appendBaseSymbolInformationFields(dst []byte, first *bool, x BaseSymbolInformation) []byte {
	dst = appendObjectName(dst, first, "name")
	dst = appendJSONString(dst, x.Name)
	dst = appendObjectName(dst, first, "kind")
	dst = appendUint32JSON(dst, uint32(x.Kind))
	if len(x.Tags) > 0 {
		dst = appendObjectName(dst, first, "tags")
		dst = appendDiagnosticTagLikeSlice(dst, x.Tags)
	}
	if x.ContainerName != nil {
		dst = appendObjectName(dst, first, "containerName")
		dst = appendJSONString(dst, *x.ContainerName)
	}
	return dst
}

func appendDiagnosticTagLikeSlice[T ~uint32](dst []byte, x []T) []byte {
	dst = append(dst, '[')
	for i, v := range x {
		if i > 0 {
			dst = append(dst, ',')
		}
		dst = appendUint32JSON(dst, uint32(v))
	}
	return append(dst, ']')
}

func appendDiagnosticTagsJSON(dst []byte, x DiagnosticTags) []byte {
	dst = append(dst, '[')
	if x.n > 0 {
		dst = appendUint32JSON(dst, uint32(x.first))
		for _, v := range x.rest[:max(x.n-1, 0)] {
			dst = append(dst, ',')
			dst = appendUint32JSON(dst, uint32(v))
		}
	}
	return append(dst, ']')
}

func appendProgressTokenJSON(dst []byte, x ProgressToken) ([]byte, error) {
	switch v := x.(type) {
	case Integer:
		return appendInt32JSON(dst, int32(v)), nil
	case String:
		return appendJSONString(dst, string(v)), nil
	case *String:
		if v == nil {
			return append(dst, "null"...), nil
		}
		return appendJSONString(dst, string(*v)), nil
	default:
		return appendJSONMarshal(dst, x)
	}
}

func appendWorkspaceSymbolLocationJSON(dst []byte, x WorkspaceSymbolLocation) ([]byte, error) {
	switch v := x.(type) {
	case *Location:
		if v == nil {
			return append(dst, "null"...), nil
		}
		return appendLocationJSON(dst, *v), nil
	case *LocationUriOnly:
		if v == nil {
			return append(dst, "null"...), nil
		}
		return appendLocationURIOnlyJSON(dst, *v), nil
	default:
		return appendJSONMarshal(dst, x)
	}
}

func appendLocationURIOnlyJSON(dst []byte, x LocationUriOnly) []byte {
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "uri")
	dst = appendJSONString(dst, string(x.URI))
	return append(dst, '}')
}

func appendLocationJSON(dst []byte, x Location) []byte {
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "uri")
	dst = appendJSONString(dst, string(x.URI))
	dst = appendObjectName(dst, &first, "range")
	dst = appendRangeJSON(dst, x.Range)
	return append(dst, '}')
}

func appendRangeJSON(dst []byte, x Range) []byte {
	dst = append(dst, `{"start":`...)
	dst = appendPositionJSON(dst, x.Start)
	dst = append(dst, `,"end":`...)
	dst = appendPositionJSON(dst, x.End)
	return append(dst, '}')
}

func appendPositionJSON(dst []byte, x Position) []byte {
	dst = append(dst, `{"line":`...)
	dst = appendUint32JSON(dst, x.Line)
	dst = append(dst, `,"character":`...)
	dst = appendUint32JSON(dst, x.Character)
	return append(dst, '}')
}

func (x *DidChangeTextDocumentParams) appendLSPJSON(dst []byte) ([]byte, error) {
	if x == nil {
		return append(dst, "null"...), nil
	}
	dst = slices.Grow(dst, 64+len(x.ContentChanges)*128)
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "textDocument")
	dst = appendVersionedTextDocumentIdentifierJSON(dst, x.TextDocument)
	dst = appendObjectName(dst, &first, "contentChanges")
	var err error
	dst, err = appendTextDocumentContentChangeEventSliceJSON(dst, x.ContentChanges)
	if err != nil {
		return nil, err
	}
	return append(dst, '}'), nil
}

func appendVersionedTextDocumentIdentifierJSON(dst []byte, x VersionedTextDocumentIdentifier) []byte {
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "uri")
	dst = appendJSONString(dst, string(x.URI))
	dst = appendObjectName(dst, &first, "version")
	dst = appendInt32JSON(dst, x.Version)
	return append(dst, '}')
}

func appendTextDocumentContentChangeEventSliceJSON(dst []byte, x []TextDocumentContentChangeEvent) ([]byte, error) {
	dst = slices.Grow(dst, 2+len(x)*128)
	dst = append(dst, '[')
	for i, v := range x {
		if i > 0 {
			dst = append(dst, ',')
		}
		var err error
		dst, err = appendTextDocumentContentChangeEventJSON(dst, v)
		if err != nil {
			return nil, err
		}
	}
	return append(dst, ']'), nil
}

func appendTextDocumentContentChangeEventJSON(dst []byte, x TextDocumentContentChangeEvent) ([]byte, error) {
	switch v := x.(type) {
	case *TextDocumentContentChangePartial:
		return v.appendLSPJSON(dst)
	case *TextDocumentContentChangeWholeDocument:
		return v.appendLSPJSON(dst)
	default:
		return appendJSONMarshal(dst, x)
	}
}

func (x *TextDocumentContentChangePartial) appendLSPJSON(dst []byte) ([]byte, error) {
	if x == nil {
		return append(dst, "null"...), nil
	}
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "range")
	dst = appendRangeJSON(dst, x.Range)
	if x.RangeLength != nil {
		dst = appendObjectName(dst, &first, "rangeLength")
		dst = appendUint32JSON(dst, *x.RangeLength)
	}
	dst = appendObjectName(dst, &first, "text")
	dst = appendJSONString(dst, x.Text)
	return append(dst, '}'), nil
}

func (x *TextDocumentContentChangeWholeDocument) appendLSPJSON(dst []byte) ([]byte, error) {
	if x == nil {
		return append(dst, "null"...), nil
	}
	dst = append(dst, '{')
	first := true
	dst = appendObjectName(dst, &first, "text")
	dst = appendJSONString(dst, x.Text)
	return append(dst, '}'), nil
}
