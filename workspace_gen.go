// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceFolder) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyURI:
		return dec.String(&v.URI)
	case keyName:
		return dec.String(&v.Name)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceFolder) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceFolder) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyURI, v.URI)
	enc.StringKey(keyName, v.Name)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceFolder) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidChangeWorkspaceFoldersParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyEvent {
		return dec.Object(&v.Event)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidChangeWorkspaceFoldersParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidChangeWorkspaceFoldersParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey(keyEvent, &v.Event)
}

// IsNil returns wether the structure is nil value or not.
func (v *DidChangeWorkspaceFoldersParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceFoldersChangeEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyAdded:
		dec.Array((*workspaceFolders)(&v.Added))
	case keyRemoved:
		dec.Array((*workspaceFolders)(&v.Removed))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceFoldersChangeEvent) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceFoldersChangeEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyAdded, (*workspaceFolders)(&v.Added))
	enc.ArrayKey(keyRemoved, (*workspaceFolders)(&v.Removed))
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceFoldersChangeEvent) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidChangeConfigurationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keySettings {
		dec.Interface(&v.Settings)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidChangeConfigurationParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidChangeConfigurationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddInterfaceKeyOmitEmpty(keySettings, &v.Settings)
}

// IsNil returns wether the structure is nil value or not.
func (v *DidChangeConfigurationParams) IsNil() bool { return v == nil }

type configurationItem []ConfigurationItem

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *configurationItem) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := ConfigurationItem{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *configurationItem) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *configurationItem) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.ObjectOmitEmpty(&t)
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *configurationItem) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ConfigurationParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyItems {
		dec.Array((*configurationItem)(&v.Items))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ConfigurationParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ConfigurationParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyItems, (*configurationItem)(&v.Items))
}

// IsNil returns wether the structure is nil value or not.
func (v *ConfigurationParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ConfigurationItem) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyScopeURI:
		return dec.String(&v.ScopeURI)
	case keySection:
		return dec.String(&v.Section)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ConfigurationItem) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ConfigurationItem) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty(keyScopeURI, v.ScopeURI)
	enc.StringKeyOmitEmpty(keySection, v.Section)
}

// IsNil returns wether the structure is nil value or not.
func (v *ConfigurationItem) IsNil() bool { return v == nil }

type fileEvents []*FileEvent

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *fileEvents) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := &FileEvent{}
	if err := dec.Object(t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *fileEvents) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *fileEvents) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.ObjectOmitEmpty(t)
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *fileEvents) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidChangeWatchedFilesParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyChanges {
		if v.Changes == nil {
			v.Changes = []*FileEvent{}
		}
		dec.Array((*fileEvents)(&v.Changes))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidChangeWatchedFilesParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidChangeWatchedFilesParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKeyOmitEmpty(keyChanges, (*fileEvents)(&v.Changes))
}

// IsNil returns wether the structure is nil value or not.
func (v *DidChangeWatchedFilesParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FileEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyType:
		return dec.Float64(&v.Type)
	case keyURI:
		return dec.String((*string)(&v.URI))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FileEvent) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FileEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Float64Key(keyType, v.Type)
	enc.StringKey(keyURI, string(v.URI))
}

// IsNil returns wether the structure is nil value or not.
func (v *FileEvent) IsNil() bool { return v == nil }

type fileSystemWatcher []FileSystemWatcher

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray.
func (v *fileSystemWatcher) UnmarshalJSONArray(dec *gojay.Decoder) error {
	t := FileSystemWatcher{}
	if err := dec.Object(&t); err != nil {
		return err
	}
	*v = append(*v, t)
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *fileSystemWatcher) NKeys() int { return 1 }

// MarshalJSONArray implements gojay's MarshalerJSONArray.
func (v *fileSystemWatcher) MarshalJSONArray(enc *gojay.Encoder) {
	for _, t := range *v {
		enc.ObjectOmitEmpty(&t)
	}
}

// IsNil implements gojay's MarshalerJSONArray.
func (v *fileSystemWatcher) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *DidChangeWatchedFilesRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyWatchers {
		dec.Array((*fileSystemWatcher)(&v.Watchers))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *DidChangeWatchedFilesRegistrationOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *DidChangeWatchedFilesRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyWatchers, (*fileSystemWatcher)(&v.Watchers))
}

// IsNil returns wether the structure is nil value or not.
func (v *DidChangeWatchedFilesRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *FileSystemWatcher) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyGlobPattern:
		return dec.String(&v.GlobPattern)
	case keyKind:
		return dec.Float64((*float64)(&v.Kind))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *FileSystemWatcher) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *FileSystemWatcher) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyGlobPattern, v.GlobPattern)
	enc.Float64KeyOmitEmpty(keyKind, float64(v.Kind))
}

// IsNil returns wether the structure is nil value or not.
func (v *FileSystemWatcher) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *WorkspaceSymbolParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyQuery {
		return dec.String(&v.Query)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *WorkspaceSymbolParams) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *WorkspaceSymbolParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyQuery, v.Query)
}

// IsNil returns wether the structure is nil value or not.
func (v *WorkspaceSymbolParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ExecuteCommandParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyCommand:
		return dec.String(&v.Command)
	case keyArguments:
		return dec.Array((*interfaces)(&v.Arguments))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ExecuteCommandParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ExecuteCommandParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(keyCommand, v.Command)
	enc.ArrayKeyOmitEmpty(keyArguments, (*interfaces)(&v.Arguments))
}

// IsNil returns wether the structure is nil value or not.
func (v *ExecuteCommandParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ExecuteCommandRegistrationOptions) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyCommands {
		dec.Array((*stringSlice)(&v.Commands))
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ExecuteCommandRegistrationOptions) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ExecuteCommandRegistrationOptions) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey(keyCommands, (*stringSlice)(&v.Commands))
}

// IsNil returns wether the structure is nil value or not.
func (v *ExecuteCommandRegistrationOptions) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ApplyWorkspaceEditParams) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case keyLabel:
		return dec.String(&v.Label)
	case keyEdit:
		return dec.Object(&v.Edit)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ApplyWorkspaceEditParams) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ApplyWorkspaceEditParams) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKeyOmitEmpty(keyLabel, v.Label)
	enc.ObjectKey(keyEdit, &v.Edit)
}

// IsNil returns wether the structure is nil value or not.
func (v *ApplyWorkspaceEditParams) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject.
func (v *ApplyWorkspaceEditResponse) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == keyApplied {
		return dec.Bool(&v.Applied)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal.
func (v *ApplyWorkspaceEditResponse) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject.
func (v *ApplyWorkspaceEditResponse) MarshalJSONObject(enc *gojay.Encoder) {
	enc.BoolKey(keyApplied, v.Applied)
}

// IsNil returns wether the structure is nil value or not.
func (v *ApplyWorkspaceEditResponse) IsNil() bool { return v == nil }
