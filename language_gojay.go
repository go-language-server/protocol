// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyContext:
		if v.Context == nil {
			v.Context = &CompletionContext{}
		}
		return dec.Object(v.Context)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ObjectKey(keyPosition, &v.Position)
	enc.ObjectKeyOmitEmpty(keyContext, v.Context)
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionContext) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTriggerCharacter:
		return dec.String(&v.TriggerCharacter)
	case keyTriggerKind:
		return dec.Float64((*float64)(&v.TriggerKind))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionContext) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionContext) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty(keyTriggerCharacter, v.TriggerCharacter)
	enc.Float64Key(keyTriggerKind, float64(v.TriggerKind))
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionContext) IsNil() bool { return v == nil }

type items []CompletionItem

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *items) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := CompletionItem{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *items) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *items) MarshalJSONArray(enc *gojay.Encoder) {
	vv := *v
	for i := range vv {
		enc.ObjectOmitEmpty(&vv[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *items) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionList) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyIsIncomplete:
		return dec.Bool(&v.IsIncomplete)
	case keyItems:
		return dec.Array((*items)(&v.Items))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionList) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionList) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey(keyIsIncomplete, v.IsIncomplete)
	enc.ArrayKey(keyItems, (*items)(&v.Items))
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionList) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyAdditionalTextEdits:
		return dec.Array((*TextEdits)(&v.AdditionalTextEdits))
	case keyCommand:
		return dec.Object(v.Command)
	case keyCommitCharacters:
		return dec.Array((*Strings)(&v.CommitCharacters))
	case keyData:
		return dec.Interface(&v.Data)
	case keyDeprecated:
		return dec.Bool(&v.Deprecated)
	case keyDetail:
		return dec.String(&v.Detail)
	case keyDocumentation:
		return dec.Interface(&v.Documentation)
	case keyFilterText:
		return dec.String(&v.FilterText)
	case keyInsertText:
		return dec.String(&v.InsertText)
	case keyInsertTextFormat:
		return dec.Float64((*float64)(&v.InsertTextFormat))
	case keyKind:
		return dec.Float64(&v.Kind)
	case keyLabel:
		return dec.String(&v.Label)
	case keyPreselect:
		return dec.Bool(&v.Preselect)
	case keySortText:
		return dec.String(&v.SortText)
	case keyTextEdit:
		if v.TextEdit == nil {
			v.TextEdit = &TextEdit{}
		}
		return dec.Object(v.TextEdit)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionItem) NKeys() int { return 15 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKeyOmitEmpty(keyAdditionalTextEdits, (*TextEdits)(&v.AdditionalTextEdits))
	enc.ObjectKeyOmitEmpty(keyCommand, v.Command)
	enc.AddArrayKeyOmitEmpty(keyCommitCharacters, (*Strings)(&v.CommitCharacters))
	enc.AddInterfaceKeyOmitEmpty(keyData, &v.Data)
	enc.BoolKeyOmitEmpty(keyDeprecated, v.Deprecated)
	enc.StringKeyOmitEmpty(keyDetail, v.Detail)
	enc.AddInterfaceKeyOmitEmpty(keyDocumentation, &v.Documentation)
	enc.StringKeyOmitEmpty(keyFilterText, v.FilterText)
	enc.StringKeyOmitEmpty(keyInsertText, v.InsertText)
	enc.Float64KeyOmitEmpty(keyInsertTextFormat, float64(v.InsertTextFormat))
	enc.Float64KeyOmitEmpty(keyKind, v.Kind)
	enc.StringKeyOmitEmpty(keyLabel, v.Label)
	enc.BoolKeyOmitEmpty(keyPreselect, v.Preselect)
	enc.StringKeyOmitEmpty(keySortText, v.SortText)
	enc.ObjectKeyOmitEmpty(keyTextEdit, v.TextEdit)
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		return dec.Array(&v.DocumentSelector)
	case keyTriggerCharacters:
		return dec.Array((*Strings)(&v.TriggerCharacters))
	case keyResolveProvider:
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionRegistrationOptions) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.AddArrayKeyOmitEmpty(keyTriggerCharacters, (*Strings)(&v.TriggerCharacters))
	enc.BoolKeyOmitEmpty(keyResolveProvider, v.ResolveProvider)
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Hover) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyContents:
		return dec.Object(&v.Contents)
	case keyRange:
		return dec.Object(&v.Range)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Hover) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Hover) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyContents, &v.Contents)
	enc.ObjectKey(keyRange, &v.Range)
}

// IsNil returns wether the structure is nil value or not.
func (v *Hover) IsNil() bool { return v == nil }

// IsNil returns wether the structure is nil value or not.
func (v *CompletionRegistrationOptions) IsNil() bool { return v == nil }

type signatures []SignatureInformation

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *signatures) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := SignatureInformation{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *signatures) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *signatures) MarshalJSONArray(enc *gojay.Encoder) {
	vv := *v
	for i := range vv {
		enc.ObjectOmitEmpty(&vv[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *signatures) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SignatureHelp) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keySignatures:
		return dec.Array((*signatures)(&v.Signatures))
	case keyActiveParameter:
		return dec.Float64(&v.ActiveParameter)
	case keyActiveSignature:
		return dec.Float64(&v.ActiveSignature)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SignatureHelp) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SignatureHelp) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keySignatures, (*signatures)(&v.Signatures))
	enc.Float64KeyOmitEmpty(keyActiveParameter, v.ActiveParameter)
	enc.Float64KeyOmitEmpty(keyActiveSignature, v.ActiveSignature)
}

// IsNil returns wether the structure is nil value or not.
func (v *SignatureHelp) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SignatureInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentationFormat:
		return dec.Array((*MarkupKinds)(&v.DocumentationFormat))
	case keyParameterInformation:
		if v.ParameterInformation == nil {
			v.ParameterInformation = &ParameterInformation{}
		}
		return dec.Object(v.ParameterInformation)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SignatureInformation) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SignatureInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyDocumentationFormat, (*MarkupKinds)(&v.DocumentationFormat))
	enc.ObjectKeyOmitEmpty(keyParameterInformation, v.ParameterInformation)
}

// IsNil returns wether the structure is nil value or not.
func (v *SignatureInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ParameterInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyLabel:
		return dec.String(&v.Label)
	case keyDocumentation:
		return dec.Interface(&v.Documentation)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ParameterInformation) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ParameterInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyLabel, v.Label)
	enc.AddInterfaceKeyOmitEmpty(keyDocumentation, &v.Documentation)
}

// IsNil returns wether the structure is nil value or not.
func (v *ParameterInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SignatureHelpRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		return dec.Array(&v.DocumentSelector)
	case keyTriggerCharacters:
		return dec.Array((*Strings)(&v.TriggerCharacters))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SignatureHelpRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SignatureHelpRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.AddArrayKeyOmitEmpty(keyTriggerCharacters, (*Strings)(&v.TriggerCharacters))
}

// IsNil returns wether the structure is nil value or not.
func (v *SignatureHelpRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ReferenceContext) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyIncludeDeclaration {
		return dec.Bool(&v.IncludeDeclaration)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ReferenceContext) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ReferenceContext) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey(keyIncludeDeclaration, v.IncludeDeclaration)
}

// IsNil returns wether the structure is nil value or not.
func (v *ReferenceContext) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ReferenceParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyContext:
		return dec.Object(&v.Context)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ReferenceParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ReferenceParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ObjectKey(keyPosition, &v.Position)
	enc.ObjectKey(keyContext, &v.Context)
}

// IsNil returns wether the structure is nil value or not.
func (v *ReferenceParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentHighlight) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRange:
		return dec.Object(&v.Range)
	case keyKind:
		return dec.Int((*int)(&v.Kind))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentHighlight) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentHighlight) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyRange, &v.Range)
	enc.IntKeyOmitEmpty(keyKind, int(v.Kind))
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentHighlight) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentSymbolParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTextDocument {
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentSymbolParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentSymbolParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentSymbolParams) IsNil() bool { return v == nil }

type documentSymbols []DocumentSymbol

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *documentSymbols) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := DocumentSymbol{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *documentSymbols) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *documentSymbols) MarshalJSONArray(enc *gojay.Encoder) {
	vv := *v
	for i := range vv {
		enc.ObjectOmitEmpty(&vv[i])
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *documentSymbols) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentSymbol) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyName:
		return dec.String(&v.Name)
	case keyDetail:
		return dec.String(&v.Detail)
	case keyKind:
		return dec.Float64((*float64)(&v.Kind))
	case keyDeprecated:
		return dec.Bool(&v.Deprecated)
	case keyRange:
		return dec.Object(&v.Range)
	case keySelectionRange:
		return dec.Object(&v.SelectionRange)
	case keyChildren:
		if v.Children == nil {
			v.Children = []DocumentSymbol{}
		}
		return dec.Array((*documentSymbols)(&v.Children))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentSymbol) NKeys() int { return 7 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyName, v.Name)
	enc.StringKeyOmitEmpty(keyDetail, v.Detail)
	enc.Float64Key(keyKind, float64(v.Kind))
	enc.BoolKeyOmitEmpty(keyDeprecated, v.Deprecated)
	enc.ObjectKey(keyRange, &v.Range)
	enc.ObjectKey(keySelectionRange, &v.SelectionRange)
	enc.ArrayKeyOmitEmpty(keyChildren, (*documentSymbols)(&v.Children))
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentSymbol) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentFormattingParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyOptions:
		return dec.Object(&v.Options)
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentFormattingParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentFormattingParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyOptions, &v.Options)
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentFormattingParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SymbolInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyName:
		return dec.String(&v.Name)
	case keyKind:
		return dec.Float64(&v.Kind)
	case keyDeprecated:
		return dec.Bool(&v.Deprecated)
	case keyLocation:
		return dec.Object(&v.Location)
	case keyContainerName:
		return dec.String(&v.ContainerName)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SymbolInformation) NKeys() int { return 5 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SymbolInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyName, v.Name)
	enc.Float64Key(keyKind, v.Kind)
	enc.BoolKeyOmitEmpty(keyDeprecated, v.Deprecated)
	enc.ObjectKey(keyLocation, &v.Location)
	enc.StringKeyOmitEmpty(keyContainerName, v.ContainerName)
}

// IsNil returns wether the structure is nil value or not.
func (v *SymbolInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeActionParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyContext:
		return dec.Object(&v.Context)
	case keyRange:
		return dec.Object(&v.Range)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeActionParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ObjectKey(keyContext, &v.Context)
	enc.ObjectKey(keyRange, &v.Range)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeActionParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeActionContext) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDiagnostics:
		return dec.Array((*Diagnostics)(&v.Diagnostics))
	case keyOnly:
		return dec.Array((*CodeActionKinds)(&v.Only))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionContext) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeActionContext) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDiagnostics, Diagnostics(v.Diagnostics))
	enc.ArrayKey(keyOnly, CodeActionKinds(v.Only))
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeActionContext) IsNil() bool { return v == nil }

// IsNil returns wether the structure is nil value or not.
func (v *CodeActionRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeAction) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTitle:
		return dec.String(&v.Title)
	case keyKind:
		return dec.String((*string)(&v.Kind))
	case keyDiagnostics:
		return dec.Array((*Diagnostics)(&v.Diagnostics))
	case keyEdit:
		if v.Edit == nil {
			v.Edit = &WorkspaceEdit{}
		}
		return dec.Object(v.Edit)
	case keyCommand:
		if v.Command == nil {
			v.Command = &Command{}
		}
		return dec.Object(v.Command)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeAction) NKeys() int { return 5 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeAction) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyTitle, v.Title)
	enc.StringKeyOmitEmpty(keyKind, string(v.Kind))
	enc.ArrayKeyOmitEmpty(keyDiagnostics, Diagnostics(v.Diagnostics))
	enc.ObjectKeyOmitEmpty(keyEdit, v.Edit)
	enc.ObjectKeyOmitEmpty(keyCommand, v.Command)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeAction) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeActionRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		return dec.Array(&v.DocumentSelector)
	case keyCodeActionKinds:
		return dec.Array((*CodeActionKinds)(&v.CodeActionKinds))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeActionRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.ArrayKeyOmitEmpty(keyCodeActionKinds, CodeActionKinds(v.CodeActionKinds))
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeLensParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTextDocument {
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeLensParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeLensParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeLensParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeLens) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRange:
		return dec.Object(&v.Range)
	case keyCommand:
		if v.Command == nil {
			v.Command = &Command{}
		}
		return dec.Object(v.Command)
	case keyData:
		return dec.Interface(&v.Data)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeLens) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeLens) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyRange, &v.Range)
	enc.ObjectKeyOmitEmpty(keyCommand, v.Command)
	enc.AddInterfaceKeyOmitEmpty(keyData, &v.Data)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeLens) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeLensRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		return dec.Array(&v.DocumentSelector)
	case keyResolveProvider:
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeLensRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeLensRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyResolveProvider, v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeLensRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentLinkParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTextDocument {
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentLinkParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentLinkParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentLinkParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentLink) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRange:
		return dec.Object(&v.Range)
	case keyTarget:
		return dec.String((*string)(&v.Target))
	case keyData:
		return dec.Interface(&v.Data)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentLink) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentLink) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyRange, &v.Range)
	enc.StringKeyOmitEmpty(keyTarget, string(v.Target))
	enc.AddInterfaceKeyOmitEmpty(keyData, &v.Data)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentLink) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentColorParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTextDocument {
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentColorParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentColorParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentColorParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ColorInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyRange:
		return dec.Object(&v.Range)
	case keyColor:
		return dec.Object(&v.Color)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ColorInformation) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ColorInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyRange, &v.Range)
	enc.ObjectKey(keyColor, &v.Color)
}

// IsNil returns wether the structure is nil value or not.
func (v *ColorInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Color) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyAlpha:
		return dec.Float64(&v.Alpha)
	case keyBlue:
		return dec.Float64(&v.Blue)
	case keyGreen:
		return dec.Float64(&v.Green)
	case keyRed:
		return dec.Float64(&v.Red)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Color) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Color) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key(keyAlpha, v.Alpha)
	enc.Float64Key(keyBlue, v.Blue)
	enc.Float64Key(keyGreen, v.Green)
	enc.Float64Key(keyRed, v.Red)
}

// IsNil returns wether the structure is nil value or not.
func (v *Color) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ColorPresentationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyColor:
		return dec.Object(&v.Color)
	case keyRange:
		return dec.Object(&v.Range)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ColorPresentationParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ColorPresentationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ObjectKey(keyColor, &v.Color)
	enc.ObjectKey(keyRange, &v.Range)
}

// IsNil returns wether the structure is nil value or not.
func (v *ColorPresentationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ColorPresentation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyLabel:
		return dec.String(&v.Label)
	case keyTextEdit:
		if v.TextEdit == nil {
			v.TextEdit = &TextEdit{}
		}
		return dec.Object(v.TextEdit)
	case keyAdditionalTextEdits:
		return dec.Array((*TextEdits)(&v.AdditionalTextEdits))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ColorPresentation) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ColorPresentation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyLabel, v.Label)
	enc.ObjectKey(keyTextEdit, v.TextEdit)
	enc.AddArrayKeyOmitEmpty(keyAdditionalTextEdits, (*TextEdits)(&v.AdditionalTextEdits))
}

// IsNil returns wether the structure is nil value or not.
func (v *ColorPresentation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FormattingOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyInsertSpaces:
		return dec.Bool(&v.InsertSpaces)
	case keyTabSize:
		return dec.Float64(&v.TabSize)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FormattingOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FormattingOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey(keyInsertSpaces, v.InsertSpaces)
	enc.Float64Key(keyTabSize, v.TabSize)
}

// IsNil returns wether the structure is nil value or not.
func (v *FormattingOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentRangeFormattingParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyRange:
		return dec.Object(&v.Range)
	case keyOptions:
		return dec.Object(&v.Options)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentRangeFormattingParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentRangeFormattingParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ObjectKey(keyRange, &v.Range)
	enc.ObjectKey(keyOptions, &v.Options)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentRangeFormattingParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentOnTypeFormattingParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyCh:
		return dec.String(&v.Ch)
	case keyOptions:
		return dec.Object(&v.Options)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentOnTypeFormattingParams) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentOnTypeFormattingParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ObjectKey(keyPosition, &v.Position)
	enc.StringKey(keyCh, v.Ch)
	enc.ObjectKey(keyOptions, &v.Options)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentOnTypeFormattingParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentOnTypeFormattingRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		return dec.Array(&v.DocumentSelector)
	case keyFirstTriggerCharacter:
		return dec.String(&v.FirstTriggerCharacter)
	case keyMoreTriggerCharacter:
		return dec.Array((*Strings)(&v.MoreTriggerCharacter))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentOnTypeFormattingRegistrationOptions) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentOnTypeFormattingRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.StringKey(keyFirstTriggerCharacter, v.FirstTriggerCharacter)
	enc.ArrayKeyOmitEmpty(keyMoreTriggerCharacter, (*Strings)(&v.MoreTriggerCharacter))
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentOnTypeFormattingRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RenameParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyTextDocument:
		return dec.Object(&v.TextDocument)
	case keyPosition:
		return dec.Object(&v.Position)
	case keyNewName:
		return dec.String(&v.NewName)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RenameParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *RenameParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
	enc.ObjectKey(keyPosition, &v.Position)
	enc.StringKey(keyNewName, v.NewName)
}

// IsNil returns wether the structure is nil value or not.
func (v *RenameParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RenameRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyDocumentSelector:
		return dec.Array(&v.DocumentSelector)
	case keyPrepareProvider:
		return dec.Bool(&v.PrepareProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RenameRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *RenameRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyDocumentSelector, &v.DocumentSelector)
	enc.BoolKeyOmitEmpty(keyPrepareProvider, v.PrepareProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *RenameRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FoldingRangeParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyTextDocument {
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FoldingRangeParams) NKeys() int { return 1 }

// IsNil returns wether the structure is nil value or not.
func (v *FoldingRangeParams) IsNil() bool { return v == nil }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FoldingRangeParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyTextDocument, &v.TextDocument)
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FoldingRange) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyStartLine:
		return dec.Float64(&v.StartLine)
	case keyStartCharacter:
		return dec.Float64(&v.StartCharacter)
	case keyEndLine:
		return dec.Float64(&v.EndLine)
	case keyEndCharacter:
		return dec.Float64(&v.EndCharacter)
	case keyKind:
		return dec.String((*string)(&v.Kind))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FoldingRange) NKeys() int { return 5 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FoldingRange) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key(keyStartLine, v.StartLine)
	enc.Float64KeyOmitEmpty(keyStartCharacter, v.StartCharacter)
	enc.Float64Key(keyEndLine, v.EndLine)
	enc.Float64KeyOmitEmpty(keyEndCharacter, v.EndCharacter)
	enc.StringKeyOmitEmpty(keyKind, string(v.Kind))
}

// IsNil returns wether the structure is nil value or not.
func (v *FoldingRange) IsNil() bool { return v == nil }
