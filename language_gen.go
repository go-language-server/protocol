// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(&v.TextDocument)
	case "position":
		return dec.Object(&v.Position)
	case "context":
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
	enc.ObjectKey("textDocument", &v.TextDocument)
	enc.ObjectKey("position", &v.Position)
	enc.ObjectKey("context", v.Context)
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionContext) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "triggerCharacter":
		return dec.String(&v.TriggerCharacter)
	case "triggerKind":
		return dec.Float64((*float64)(v.TriggerKind))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionContext) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionContext) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("triggerCharacter", v.TriggerCharacter)
	enc.Float64Key("triggerKind", float64(*v.TriggerKind))
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionContext) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionList) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "isIncomplete":
		return dec.Bool(&v.IsIncomplete)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionList) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionList) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("isIncomplete", v.IsIncomplete)
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionList) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "command":
		return dec.Object(v.Command)
	case "deprecated":
		return dec.Bool(&v.Deprecated)
	case "detail":
		return dec.String(&v.Detail)
	case "filterText":
		return dec.String(&v.FilterText)
	case "insertText":
		return dec.String(&v.InsertText)
	case "insertTextFormat":
		return dec.Float64((*float64)(&v.InsertTextFormat))
	case "kind":
		return dec.Float64(&v.Kind)
	case "label":
		return dec.String(&v.Label)
	case "preselect":
		return dec.Bool(&v.Preselect)
	case "sortText":
		return dec.String(&v.SortText)
	case "textEdit":
		return dec.Object(v.TextEdit)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionItem) NKeys() int { return 11 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("command", v.Command)
	enc.BoolKey("deprecated", v.Deprecated)
	enc.StringKey("detail", v.Detail)
	enc.StringKey("filterText", v.FilterText)
	enc.StringKey("insertText", v.InsertText)
	enc.Float64Key("insertTextFormat", float64(v.InsertTextFormat))
	enc.Float64Key("kind", v.Kind)
	enc.StringKey("label", v.Label)
	enc.BoolKey("preselect", v.Preselect)
	enc.StringKey("sortText", v.SortText)
	enc.ObjectKey("textEdit", v.TextEdit)
}

// IsNil returns wether the structure is nil value or not.
func (v *CompletionItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CompletionRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	case "resolveProvider":
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CompletionRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CompletionRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("documentSelector", v.DocumentSelector)
	enc.BoolKey("resolveProvider", v.ResolveProvider)
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Hover) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "contents":
		return dec.Object(&v.Contents)
	case "range":
		return dec.Object(&v.Range)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Hover) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Hover) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("contents", &v.Contents)
	enc.ObjectKey("range", &v.Range)
}

// IsNil returns wether the structure is nil value or not.
func (v *Hover) IsNil() bool { return v == nil }

// IsNil returns wether the structure is nil value or not.
func (v *CompletionRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SignatureHelp) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "activeParameter":
		return dec.Float64(&v.ActiveParameter)
	case "activeSignature":
		return dec.Float64(&v.ActiveSignature)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SignatureHelp) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SignatureHelp) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key("activeParameter", v.ActiveParameter)
	enc.Float64Key("activeSignature", v.ActiveSignature)
}

// IsNil returns wether the structure is nil value or not.
func (v *SignatureHelp) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SignatureInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "parameterInformation":
		if v.ParameterInformation == nil {
			v.ParameterInformation = &ParameterInformation{}
		}
		return dec.Object(v.ParameterInformation)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SignatureInformation) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SignatureInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("parameterInformation", v.ParameterInformation)
}

// IsNil returns wether the structure is nil value or not.
func (v *SignatureInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ParameterInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "label":
		return dec.String(&v.Label)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ParameterInformation) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ParameterInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("label", v.Label)
}

// IsNil returns wether the structure is nil value or not.
func (v *ParameterInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SignatureHelpRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SignatureHelpRegistrationOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SignatureHelpRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("documentSelector", v.DocumentSelector)
}

// IsNil returns wether the structure is nil value or not.
func (v *SignatureHelpRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ReferenceContext) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "includeDeclaration":
		return dec.Bool(&v.IncludeDeclaration)
	}
	return nil
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ReferenceParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(&v.TextDocument)
	case "position":
		return dec.Object(&v.Position)
	case "context":
		if v.Context == nil {
			v.Context = &ReferenceContext{}
		}
		return dec.Object(v.Context)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ReferenceParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ReferenceParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", &v.TextDocument)
	enc.ObjectKey("position", &v.Position)
	enc.ObjectKey("context", v.Context)
}

// IsNil returns wether the structure is nil value or not.
func (v *ReferenceParams) IsNil() bool { return v == nil }

// NKeys returns the number of keys to unmarshal.
func (v *ReferenceContext) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ReferenceContext) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("includeDeclaration", v.IncludeDeclaration)
}

// IsNil returns wether the structure is nil value or not.
func (v *ReferenceContext) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentHighlight) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "range":
		return dec.Object(&v.Range)
	case "kind":
		return dec.Int((*int)(&v.Kind))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentHighlight) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentHighlight) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("range", &v.Range)
	enc.IntKey("kind", int(v.Kind))
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentHighlight) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentSymbolParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentSymbolParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentSymbolParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentSymbolParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentSymbol) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "name":
		return dec.String(&v.Name)
	case "detail":
		return dec.String(&v.Detail)
	case "kind":
		return dec.Float64((*float64)(&v.Kind))
	case "deprecated":
		return dec.Bool(&v.Deprecated)
	case "range":
		return dec.Object(v.Range)
	case "selectionRange":
		return dec.Object(v.SelectionRange)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentSymbol) NKeys() int { return 6 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentSymbol) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("name", v.Name)
	enc.StringKey("detail", v.Detail)
	enc.Float64Key("kind", float64(v.Kind))
	enc.BoolKey("deprecated", v.Deprecated)
	enc.ObjectKey("range", v.Range)
	enc.ObjectKey("selectionRange", v.SelectionRange)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentSymbol) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentFormattingParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "options":
		if v.Options == nil {
			v.Options = &FormattingOptions{}
		}
		return dec.Object(v.Options)
	case "textDocument":
		return dec.Object(v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentFormattingParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentFormattingParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("options", v.Options)
	enc.ObjectKey("textDocument", v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentFormattingParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *SymbolInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "name":
		return dec.String(&v.Name)
	case "kind":
		return dec.Float64(&v.Kind)
	case "deprecated":
		return dec.Bool(&v.Deprecated)
	case "location":
		return dec.Object(v.Location)
	case "containerName":
		return dec.String(&v.ContainerName)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *SymbolInformation) NKeys() int { return 5 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *SymbolInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("name", v.Name)
	enc.Float64Key("kind", v.Kind)
	enc.BoolKey("deprecated", v.Deprecated)
	enc.ObjectKey("location", v.Location)
	enc.StringKey("containerName", v.ContainerName)
}

// IsNil returns wether the structure is nil value or not.
func (v *SymbolInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeActionParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(v.TextDocument)
	case "context":
		if v.Context == nil {
			v.Context = &CodeActionContext{}
		}
		return dec.Object(v.Context)
	case "range":
		return dec.Object(v.Range)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeActionParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", v.TextDocument)
	enc.ObjectKey("context", v.Context)
	enc.ObjectKey("range", v.Range)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeActionParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeActionContext) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionContext) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeActionContext) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeActionContext) IsNil() bool { return v == nil }

// IsNil returns wether the structure is nil value or not.
func (v *CodeActionRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeAction) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "title":
		return dec.String(&v.Title)
	case "kind":
		return dec.String((*string)(&v.Kind))
	case "edit":
		return dec.Object(v.Edit)
	case "command":
		return dec.Object(v.Command)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeAction) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeAction) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("title", v.Title)
	enc.StringKey("kind", string(v.Kind))
	enc.ObjectKey("edit", v.Edit)
	enc.ObjectKey("command", v.Command)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeAction) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeActionRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeActionRegistrationOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeActionRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("documentSelector", v.DocumentSelector)
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeLensParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeLensParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeLensParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeLensParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeLens) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "range":
		return dec.Object(v.Range)
	case "command":
		return dec.Object(v.Command)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeLens) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeLens) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("range", v.Range)
	enc.ObjectKey("command", v.Command)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeLens) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *CodeLensRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	case "resolveProvider":
		return dec.Bool(&v.ResolveProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *CodeLensRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *CodeLensRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("documentSelector", v.DocumentSelector)
	enc.BoolKey("resolveProvider", v.ResolveProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *CodeLensRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentLinkParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentLinkParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentLinkParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentLinkParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentLink) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "range":
		return dec.Object(&v.Range)
	case "target":
		return dec.String((*string)(&v.Target))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentLink) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentLink) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("range", &v.Range)
	enc.StringKey("target", string(v.Target))
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentLink) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentColorParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentColorParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentColorParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentColorParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ColorInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "range":
		return dec.Object(v.Range)
	case "color":
		if v.Color == nil {
			v.Color = &Color{}
		}
		return dec.Object(v.Color)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ColorInformation) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ColorInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("range", v.Range)
	enc.ObjectKey("color", v.Color)
}

// IsNil returns wether the structure is nil value or not.
func (v *ColorInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *Color) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "alpha":
		return dec.Float64(&v.Alpha)
	case "blue":
		return dec.Float64(&v.Blue)
	case "green":
		return dec.Float64(&v.Green)
	case "red":
		return dec.Float64(&v.Red)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *Color) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *Color) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key("alpha", v.Alpha)
	enc.Float64Key("blue", v.Blue)
	enc.Float64Key("green", v.Green)
	enc.Float64Key("red", v.Red)
}

// IsNil returns wether the structure is nil value or not.
func (v *Color) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ColorPresentationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(v.TextDocument)
	case "color":
		if v.Color == nil {
			v.Color = &Color{}
		}
		return dec.Object(v.Color)
	case "range":
		return dec.Object(v.Range)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ColorPresentationParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ColorPresentationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", v.TextDocument)
	enc.ObjectKey("color", v.Color)
	enc.ObjectKey("range", v.Range)
}

// IsNil returns wether the structure is nil value or not.
func (v *ColorPresentationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ColorPresentation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "label":
		return dec.String(&v.Label)
	case "textEdit":
		return dec.Object(v.TextEdit)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ColorPresentation) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ColorPresentation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("label", v.Label)
	enc.ObjectKey("textEdit", v.TextEdit)
}

// IsNil returns wether the structure is nil value or not.
func (v *ColorPresentation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FormattingOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "insertSpaces":
		return dec.Bool(&v.InsertSpaces)
	case "tabSize":
		return dec.Float64(&v.TabSize)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FormattingOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FormattingOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("insertSpaces", v.InsertSpaces)
	enc.Float64Key("tabSize", v.TabSize)
}

// IsNil returns wether the structure is nil value or not.
func (v *FormattingOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentRangeFormattingParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(v.TextDocument)
	case "range":
		return dec.Object(v.Range)
	case "options":
		if v.Options == nil {
			v.Options = &FormattingOptions{}
		}
		return dec.Object(v.Options)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentRangeFormattingParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentRangeFormattingParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", v.TextDocument)
	enc.ObjectKey("range", v.Range)
	enc.ObjectKey("options", v.Options)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentRangeFormattingParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentOnTypeFormattingParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(v.TextDocument)
	case "position":
		return dec.Object(v.Position)
	case "ch":
		return dec.String(&v.Ch)
	case "options":
		if v.Options == nil {
			v.Options = &FormattingOptions{}
		}
		return dec.Object(v.Options)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentOnTypeFormattingParams) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentOnTypeFormattingParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", v.TextDocument)
	enc.ObjectKey("position", v.Position)
	enc.StringKey("ch", v.Ch)
	enc.ObjectKey("options", v.Options)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentOnTypeFormattingParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DocumentOnTypeFormattingRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	case "firstTriggerCharacter":
		return dec.String(&v.FirstTriggerCharacter)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DocumentOnTypeFormattingRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DocumentOnTypeFormattingRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("documentSelector", v.DocumentSelector)
	enc.StringKey("firstTriggerCharacter", v.FirstTriggerCharacter)
}

// IsNil returns wether the structure is nil value or not.
func (v *DocumentOnTypeFormattingRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RenameParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(v.TextDocument)
	case "position":
		return dec.Object(v.Position)
	case "newName":
		return dec.String(&v.NewName)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RenameParams) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *RenameParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", v.TextDocument)
	enc.ObjectKey("position", v.Position)
	enc.StringKey("newName", v.NewName)
}

// IsNil returns wether the structure is nil value or not.
func (v *RenameParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *RenameRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "documentSelector":
		return dec.Array(v.DocumentSelector)
	case "prepareProvider":
		return dec.Bool(&v.PrepareProvider)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *RenameRegistrationOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *RenameRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("documentSelector", v.DocumentSelector)
	enc.BoolKey("prepareProvider", v.PrepareProvider)
}

// IsNil returns wether the structure is nil value or not.
func (v *RenameRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FoldingRangeParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		return dec.Object(v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FoldingRangeParams) NKeys() int { return 1 }

// IsNil returns wether the structure is nil value or not.
func (v *FoldingRangeParams) IsNil() bool { return v == nil }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FoldingRangeParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", v.TextDocument)
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FoldingRange) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "startLine":
		return dec.Float64(&v.StartLine)
	case "startCharacter":
		return dec.Float64(&v.StartCharacter)
	case "endLine":
		return dec.Float64(&v.EndLine)
	case "endCharacter":
		return dec.Float64(&v.EndCharacter)
	case "kind":
		return dec.String((*string)(&v.Kind))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FoldingRange) NKeys() int { return 5 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FoldingRange) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key("startLine", v.StartLine)
	enc.Float64Key("startCharacter", v.StartCharacter)
	enc.Float64Key("endLine", v.EndLine)
	enc.Float64Key("endCharacter", v.EndCharacter)
	enc.StringKey("kind", string(v.Kind))
}

// IsNil returns wether the structure is nil value or not.
func (v *FoldingRange) IsNil() bool { return v == nil }
