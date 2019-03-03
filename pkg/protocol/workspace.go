// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

// WorkspaceFolder response of Workspace folders request.
type WorkspaceFolder struct {

	// Name is the name of the workspace folder. Used to refer to this
	// workspace folder in the user interface.
	Name string `json:"name"`

	// URI is the associated URI for this workspace folder.
	URI string `json:"uri"`
}

// DidChangeWorkspaceFoldersParams params of DidChangeWorkspaceFolders Notification.
type DidChangeWorkspaceFoldersParams struct {

	// Event is the actual workspace folder change event.
	Event WorkspaceFoldersChangeEvent `json:"event"`
}

// WorkspaceFoldersChangeEvent is the workspace folder change event.
type WorkspaceFoldersChangeEvent struct {

	// Added is the array of added workspace folders
	Added []WorkspaceFolder `json:"added"`

	// Removed is the array of the removed workspace folders
	Removed []WorkspaceFolder `json:"removed"`
}

// DidChangeConfigurationParams params of DidChangeConfiguration Notification.
type DidChangeConfigurationParams struct {

	// Settings is the actual changed settings
	Settings interface{} `json:"settings,omitempty"`
}

// ConfigurationParams params of Configuration Request.
type ConfigurationParams struct {
	Items []ConfigurationItem `json:"items"`
}

// ConfigurationItem a ConfigurationItem consists of the configuration section to ask for and an additional scope URI.
// The configuration section ask for is defined by the server and doesn’t necessarily need to correspond to the configuration store used be the client.
// So a server might ask for a configuration cpp.formatterOptions but the client stores the configuration in a XML store layout differently.
// It is up to the client to do the necessary conversion. If a scope URI is provided the client should return the setting scoped to the provided resource.
// If the client for example uses EditorConfig to manage its settings the configuration should be returned for the passed resource URI. If the client can’t provide a configuration setting for a given scope then null need to be present in the returned array.
type ConfigurationItem struct {

	// ScopeURI is the scope to get the configuration section for.
	ScopeURI string `json:"scopeUri,omitempty"`

	// Section is the configuration section asked for.
	Section string `json:"section,omitempty"`
}

// DidChangeWatchedFilesParams params of DidChangeWatchedFiles Notification.
type DidChangeWatchedFilesParams struct {

	// Changes is the actual file events.
	Changes []*FileEvent `json:"changes,omitempty"`
}

// FileEvent an event describing a file change.
type FileEvent struct {

	// Type is the change type.
	Type float64 `json:"type"`

	// URI is the file's URI.
	URI DocumentURI `json:"uri"`
}

// FileChangeType is the file event type.
type FileChangeType float64

const (
	// Created is the file got created.
	Created FileChangeType = 1
	// Changed is the file got changed.
	Changed FileChangeType = 2
	// Deleted is the file got deleted.
	Deleted FileChangeType = 3
)

// DidChangeWatchedFilesRegistrationOptions describe options to be used when registering for file system change events.
type DidChangeWatchedFilesRegistrationOptions struct {

	// Watchers is the watchers to register.
	Watchers []FileSystemWatcher `json:"watchers"`
}

// FileSystemWatcher watchers of DidChangeWatchedFiles Registration options.
type FileSystemWatcher struct {

	// GlobPattern is the glob pattern to watch.
	//
	// Glob patterns can have the following syntax:
	// - `*` to match one or more characters in a path segment
	// - `?` to match on one character in a path segment
	// - `**` to match any number of path segments, including none
	// - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
	// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
	GlobPattern string `json:"globPattern"`

	// Kind is the kind of events of interest. If omitted it defaults
	// to WatchKind.Create | WatchKind.Change | WatchKind.Delete
	// which is 7.
	Kind WatchKind `json:"kind,omitempty"`
}

// WatchKind kind of FileSystemWatcher kind.
type WatchKind float64

const (
	// CreateWatch interested in create events.
	CreateWatch WatchKind = 1

	// ChangeWatch interested in change events
	ChangeWatch WatchKind = 2

	// DeleteWatch interested in delete events
	DeleteWatch WatchKind = 4
)

// WorkspaceSymbolParams is the parameters of a Workspace Symbol Request.
type WorkspaceSymbolParams struct {

	// Query a non-empty query string
	Query string `json:"query"`
}

// ExecuteCommandParams params of Execute a command.
type ExecuteCommandParams struct {

	// Command is the identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments that the command should be invoked with.
	Arguments []interface{} `json:"arguments,omitempty"`
}

// ExecuteCommandRegistrationOptions execute command registration options.
type ExecuteCommandRegistrationOptions struct {

	// Commands is the commands to be executed on the server
	Commands []string `json:"commands"`
}

// ApplyWorkspaceEditParams params of Applies a WorkspaceEdit.
type ApplyWorkspaceEditParams struct {

	// Label an optional label of the workspace edit. This label is
	// presented in the user interface for example on an undo
	// stack to undo the workspace edit.
	Label string `json:"label,omitempty"`

	// Edit is the edits to apply.
	Edit WorkspaceEdit `json:"edit"`
}

// ApplyWorkspaceEditResponse response of Applies a WorkspaceEdit.
type ApplyWorkspaceEditResponse struct {

	// Applied indicates whether the edit was applied or not.
	Applied bool `json:"applied"`
}
