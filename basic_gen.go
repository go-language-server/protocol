// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"

	"github.com/go-language-server/protocol/internal/pkg/errors"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Position) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "line":
		return dec.Float64(&v.Line)
	case "character":
		return dec.Float64(&v.Character)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Position) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Position) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key("line", v.Line)
	enc.Float64Key("character", v.Character)
}

// IsNil returns wether the structure is nil value or not
func (v *Position) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Range) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "start":
		if &v.Start == nil {
			return errors.ErrorInvalidParams("Range.Start field must be not empty")
		}
		return dec.Object(&v.Start)
	case "end":
		if &v.End == nil {
			return errors.ErrorInvalidParams("Range.End field must be not empty")
		}
		return dec.Object(&v.End)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Range) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Range) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("start", &v.Start)
	enc.ObjectKey("end", &v.End)
}

// IsNil returns wether the structure is nil value or not
func (v *Range) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Location) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "uri":
		return dec.String((*string)(&v.URI))
	case "range":
		if &v.Range == nil {
			return errors.ErrorInvalidParams("Location.Range field must be non-nil")
		}
		return dec.Object(&v.Range)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Location) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Location) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("uri", string(v.URI))
	enc.ObjectKey("range", &v.Range)
}

// IsNil returns wether the structure is nil value or not
func (v *Location) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *LocationLink) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "originSelectionRange":
		if &v.OriginSelectionRange == nil {
			return errors.ErrorInvalidParams("LocationLink.OriginSelectionRange field must be non-nil")
		}
		return dec.Object(v.OriginSelectionRange)
	case "targetURI":
		return dec.String(&v.TargetURI)
	case "targetRange":
		if &v.TargetRange == nil {
			return errors.ErrorInvalidParams("LocationLink.TargetRange field must be non-nil")
		}
		return dec.Object(&v.TargetRange)
	case "targetSelectionRange":
		if &v.TargetSelectionRange == nil {
			return errors.ErrorInvalidParams("LocationLink.TargetSelectionRange field must be non-nil")
		}
		return dec.Object(&v.TargetSelectionRange)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *LocationLink) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *LocationLink) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("originSelectionRange", v.OriginSelectionRange)
	enc.StringKey("targetURI", v.TargetURI)
	enc.ObjectKey("targetRange", &v.TargetRange)
	enc.ObjectKey("targetSelectionRange", &v.TargetSelectionRange)
}

// IsNil returns wether the structure is nil value or not
func (v *LocationLink) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Diagnostic) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "range":
		if &v.Range == nil {
			return errors.ErrorInvalidParams("Diagnostic.Range field must be non-nil")
		}
		return dec.Object(&v.Range)
	case "severity":
		return dec.Float64((*float64)(&v.Severity))
	case "source":
		return dec.String(&v.Source)
	case "message":
		return dec.String(&v.Message)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Diagnostic) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Diagnostic) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("range", &v.Range)
	enc.Float64Key("severity", float64(v.Severity))
	enc.StringKey("source", v.Source)
	enc.StringKey("message", v.Message)
}

// IsNil returns wether the structure is nil value or not
func (v *Diagnostic) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DiagnosticRelatedInformation) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "location":
		if &v.Location == nil {
			return errors.ErrorInvalidParams("DiagnosticRelatedInformation.Location field must be non-nil")
		}
		return dec.Object(&v.Location)
	case "message":
		return dec.String(&v.Message)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DiagnosticRelatedInformation) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DiagnosticRelatedInformation) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("location", &v.Location)
	enc.StringKey("message", v.Message)
}

// IsNil returns wether the structure is nil value or not
func (v *DiagnosticRelatedInformation) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Command) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "title":
		return dec.String(&v.Title)
	case "command":
		return dec.String(&v.Command)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Command) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Command) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("title", v.Title)
	enc.StringKey("command", v.Command)
}

// IsNil returns wether the structure is nil value or not
func (v *Command) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextEdit) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "range":
		if &v.Range == nil {
			return errors.ErrorInvalidParams("TextEdit.Range field must be non-nil")
		}
		return dec.Object(&v.Range)
	case "newText":
		return dec.String(&v.NewText)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextEdit) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextEdit) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("range", &v.Range)
	enc.StringKey("newText", v.NewText)
}

// IsNil returns wether the structure is nil value or not
func (v *TextEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentEdit) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		if &v.TextDocument == nil {
			return errors.ErrorInvalidParams("TextDocumentEdit.TextDocument field must be non-nil")
		}
		return dec.Object(&v.TextDocument)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentEdit) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentEdit) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", &v.TextDocument)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *CreateFileOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "overwrite":
		return dec.Bool(&v.Overwrite)
	case "ignoreIfExists":
		return dec.Bool(&v.IgnoreIfExists)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *CreateFileOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *CreateFileOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("overwrite", v.Overwrite)
	enc.BoolKey("ignoreIfExists", v.IgnoreIfExists)
}

// IsNil returns wether the structure is nil value or not
func (v *CreateFileOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *CreateFile) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "kind":
		return dec.String(&v.Kind)
	case "uri":
		return dec.String(&v.URI)
	case "options":
		if v.Options == nil {
			v.Options = &CreateFileOptions{}
		}
		return dec.Object(v.Options)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *CreateFile) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *CreateFile) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("kind", v.Kind)
	enc.StringKey("uRI", v.URI)
	enc.ObjectKey("options", v.Options)
}

// IsNil returns wether the structure is nil value or not
func (v *CreateFile) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *RenameFileOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "overwrite":
		return dec.Bool(&v.Overwrite)
	case "ignoreIfExists":
		return dec.Bool(&v.IgnoreIfExists)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *RenameFileOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *RenameFileOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("overwrite", v.Overwrite)
	enc.BoolKey("ignoreIfExists", v.IgnoreIfExists)
}

// IsNil returns wether the structure is nil value or not
func (v *RenameFileOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *RenameFile) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "kind":
		return dec.String(&v.Kind)
	case "oldURI":
		return dec.String(&v.OldURI)
	case "newURI":
		return dec.String(&v.NewURI)
	case "options":
		if v.Options == nil {
			v.Options = &RenameFileOptions{}
		}
		return dec.Object(v.Options)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *RenameFile) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *RenameFile) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("kind", v.Kind)
	enc.StringKey("oldURI", v.OldURI)
	enc.StringKey("newURI", v.NewURI)
	enc.ObjectKey("options", v.Options)
}

// IsNil returns wether the structure is nil value or not
func (v *RenameFile) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DeleteFileOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "recursive":
		return dec.Bool(&v.Recursive)
	case "ignoreIfNotExists":
		return dec.Bool(&v.IgnoreIfNotExists)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DeleteFileOptions) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DeleteFileOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey("recursive", v.Recursive)
	enc.BoolKey("ignoreIfNotExists", v.IgnoreIfNotExists)
}

// IsNil returns wether the structure is nil value or not
func (v *DeleteFileOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DeleteFile) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "kind":
		return dec.String(&v.Kind)
	case "uRI":
		return dec.String(&v.URI)
	case "options":
		if v.Options == nil {
			v.Options = &DeleteFileOptions{}
		}
		return dec.Object(v.Options)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DeleteFile) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DeleteFile) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("kind", v.Kind)
	enc.StringKey("uRI", v.URI)
	enc.ObjectKey("options", v.Options)
}

// IsNil returns wether the structure is nil value or not
func (v *DeleteFile) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *WorkspaceEdit) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *WorkspaceEdit) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *WorkspaceEdit) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *WorkspaceEdit) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentIdentifier) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "uri":
		return dec.String((*string)(&v.URI))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentIdentifier) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentIdentifier) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("uri", string(v.URI))
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentIdentifier) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "uri":
		return dec.String((*string)(&v.URI))
	case "languageID":
		return dec.String((*string)(&v.LanguageID))
	case "version":
		return dec.Float64(&v.Version)
	case "text":
		return dec.String(&v.Text)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentItem) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("uri", string(v.URI))
	enc.StringKey("languageID", string(v.LanguageID))
	enc.Float64Key("version", v.Version)
	enc.StringKey("text", v.Text)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentItem) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *VersionedTextDocumentIdentifier) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "uri":
		return dec.String((*string)(&v.URI))
	case "version":
		return dec.Uint64(v.Version)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *VersionedTextDocumentIdentifier) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *VersionedTextDocumentIdentifier) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("uri", string(v.URI))
	enc.Uint64Key("version", *v.Version)
}

// IsNil returns wether the structure is nil value or not
func (v *VersionedTextDocumentIdentifier) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *TextDocumentPositionParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "textDocument":
		if &v.TextDocument == nil {
			return errors.ErrorInvalidParams("TextDocumentPositionParams.TextDocument field must be non-nil")
		}
		return dec.Object(&v.TextDocument)
	case "position":
		if &v.Position == nil {
			return errors.ErrorInvalidParams("TextDocumentPositionParams.Position field must be non-nil")
		}
		return dec.Object(&v.Position)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *TextDocumentPositionParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *TextDocumentPositionParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("textDocument", &v.TextDocument)
	enc.ObjectKey("position", &v.Position)
}

// IsNil returns wether the structure is nil value or not
func (v *TextDocumentPositionParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *DocumentFilter) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "language":
		return dec.String(&v.Language)
	case "scheme":
		return dec.String(&v.Scheme)
	case "pattern":
		return dec.String(&v.Pattern)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DocumentFilter) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *DocumentFilter) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("language", v.Language)
	enc.StringKey("scheme", v.Scheme)
	enc.StringKey("pattern", v.Pattern)
}

// IsNil returns wether the structure is nil value or not
func (v *DocumentFilter) IsNil() bool { return v == nil }

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *DocumentSelector) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var s = &DocumentFilter{}
	if err := dec.Object(s); err != nil {
		return err
	}
	*v = append(*v, s)
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *DocumentSelector) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *DocumentSelector) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Object(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *DocumentSelector) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *MarkupContent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "kind":
		return dec.String((*string)(&v.Kind))
	case "value":
		return dec.String(&v.Value)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *MarkupContent) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *MarkupContent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("kind", string(v.Kind))
	enc.StringKey("value", v.Value)
}

// IsNil returns wether the structure is nil value or not
func (v *MarkupContent) IsNil() bool { return v == nil }
