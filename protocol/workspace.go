// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"go.lsp.dev/uri"
)

// FileChangeType the file event type.
type FileChangeType uint32

const (
	// CreatedFileChangeType the file got created.
	CreatedFileChangeType FileChangeType = 1

	// ChangedFileChangeType the file got changed.
	ChangedFileChangeType FileChangeType = 2

	// DeletedFileChangeType the file got deleted.
	DeletedFileChangeType FileChangeType = 3
)

type WatchKind uint32

const (
	// CreateWatchKind interested in create events.
	CreateWatchKind WatchKind = 1

	// ChangeWatchKind interested in change events.
	ChangeWatchKind WatchKind = 2

	// DeleteWatchKind interested in delete events.
	DeleteWatchKind WatchKind = 4
)

// FileOperationPatternKind a pattern kind describing if a glob pattern matches a file a folder or both.
//
// @since 3.16.0
type FileOperationPatternKind string

const (
	// FileFileOperationPatternKind the pattern matches a file only.
	FileFileOperationPatternKind FileOperationPatternKind = "file"

	// FolderFileOperationPatternKind the pattern matches a folder only.
	FolderFileOperationPatternKind FileOperationPatternKind = "folder"
)

// WorkspaceFolder a workspace folder inside a client.
type WorkspaceFolder struct {
	// URI the associated URI for this workspace folder.
	URI uri.URI `json:"uri"`

	// Name the name of the workspace folder. Used to refer to this workspace folder in the user interface.
	Name string `json:"name"`
}

// WorkspaceFoldersChangeEvent the workspace folder change event.
type WorkspaceFoldersChangeEvent struct {
	// Added the array of added workspace folders.
	Added []WorkspaceFolder `json:"added"`

	// Removed the array of the removed workspace folders.
	Removed []WorkspaceFolder `json:"removed"`
}

// DidChangeWorkspaceFoldersParams the parameters of a `workspace/didChangeWorkspaceFolders` notification.
type DidChangeWorkspaceFoldersParams struct {
	// Event the actual workspace folder change event.
	Event WorkspaceFoldersChangeEvent `json:"event"`
}

type ConfigurationItem struct {
	// ScopeURI the scope to get the configuration section for.
	ScopeURI uri.URI `json:"scopeUri,omitempty"`

	// Section the configuration section asked for.
	Section string `json:"section,omitempty"`
}

// ConfigurationParams the parameters of a configuration request.
type ConfigurationParams struct {
	Items []ConfigurationItem `json:"items"`
}

// FileCreate represents information on a file/folder create.
//
// @since 3.16.0
type FileCreate struct {
	// URI a file:// URI for the location of the file/folder being created.
	//
	// @since 3.16.0
	URI string `json:"uri"`
}

// CreateFilesParams the parameters sent in notifications/requests for user-initiated creation of files.
//
// @since 3.16.0
type CreateFilesParams struct {
	// Files an array of all files/folders created in this operation.
	//
	// @since 3.16.0
	Files []FileCreate `json:"files"`
}

// ResourceOperation a generic resource operation.
type ResourceOperation struct {
	// Kind the resource operation kind.
	Kind string `json:"kind"`

	// AnnotationID an optional annotation identifier describing the operation.
	AnnotationID *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
}

// DeleteFileOptions delete file options.
type DeleteFileOptions struct {
	// Recursive delete the content recursively if a folder is denoted.
	Recursive bool `json:"recursive,omitempty"`

	// IgnoreIfNotExists ignore the operation if the file doesn't exist.
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
}

// DeleteFile delete file operation.
type DeleteFile struct {
	// extends
	ResourceOperation

	// URI the file to delete.
	URI DocumentURI `json:"uri"`

	// Options delete options.
	Options *DeleteFileOptions `json:"options,omitempty"`
}

// RenameFileOptions rename file options.
type RenameFileOptions struct {
	// Overwrite overwrite target if existing. Overwrite wins over `ignoreIfExists`.
	Overwrite bool `json:"overwrite,omitempty"`

	// IgnoreIfExists ignores if target exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

// RenameFile rename file operation.
type RenameFile struct {
	// extends
	ResourceOperation

	// OldURI the old (existing) location.
	OldURI DocumentURI `json:"oldUri"`

	// NewURI the new location.
	NewURI DocumentURI `json:"newUri"`

	// Options rename options.
	Options *RenameFileOptions `json:"options,omitempty"`
}

// CreateFileOptions options to create a file.
type CreateFileOptions struct {
	// Overwrite overwrite existing file. Overwrite wins over `ignoreIfExists`.
	Overwrite bool `json:"overwrite,omitempty"`

	// IgnoreIfExists ignore if exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

// CreateFile create file operation.
type CreateFile struct {
	// extends
	ResourceOperation

	// URI the resource to create.
	URI DocumentURI `json:"uri"`

	// Options additional options.
	Options *CreateFileOptions `json:"options,omitempty"`
}

// FileOperationPatternOptions matching options for the file operation pattern.
//
// @since 3.16.0
type FileOperationPatternOptions struct {
	// IgnoreCase the pattern should be matched ignoring casing.
	//
	// @since 3.16.0
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}

// FileOperationPattern a pattern to describe in which file operation requests or notifications the server is interested in receiving.
//
// @since 3.16.0
type FileOperationPattern struct {
	// Glob the glob pattern to match. Glob patterns can have the following syntax: - `*` to match one or more characters in a path segment - `?` to match on one character in a path segment - `**` to match any number of path segments, including none - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files) - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …) - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`).
	//
	// @since 3.16.0
	Glob string `json:"glob"`

	// Matches whether to match files or folders with this pattern. Matches both if undefined.
	//
	// @since 3.16.0
	Matches FileOperationPatternKind `json:"matches,omitempty"`

	// Options additional options used during matching.
	//
	// @since 3.16.0
	Options *FileOperationPatternOptions `json:"options,omitempty"`
}

// FileOperationFilter a filter to describe in which file operation requests or notifications the server is interested in receiving.
//
// @since 3.16.0
type FileOperationFilter struct {
	// Scheme a Uri scheme like `file` or `untitled`.
	//
	// @since 3.16.0
	Scheme string `json:"scheme,omitempty"`

	// Pattern the actual file operation pattern.
	//
	// @since 3.16.0
	Pattern FileOperationPattern `json:"pattern"`
}

// FileOperationRegistrationOptions the options to register for file operations.
//
// @since 3.16.0
type FileOperationRegistrationOptions struct {
	// Filters the actual filters.
	//
	// @since 3.16.0
	Filters []FileOperationFilter `json:"filters"`
}

// FileRename represents information on a file/folder rename.
//
// @since 3.16.0
type FileRename struct {
	// OldURI a file:// URI for the original location of the file/folder being renamed.
	//
	// @since 3.16.0
	OldURI string `json:"oldUri"`

	// NewURI a file:// URI for the new location of the file/folder being renamed.
	//
	// @since 3.16.0
	NewURI string `json:"newUri"`
}

// RenameFilesParams the parameters sent in notifications/requests for user-initiated renames of files.
//
// @since 3.16.0
type RenameFilesParams struct {
	// Files an array of all files/folders renamed in this operation. When a folder is renamed, only the folder will be included, and not its children.
	//
	// @since 3.16.0
	Files []FileRename `json:"files"`
}

// FileDelete represents information on a file/folder delete.
//
// @since 3.16.0
type FileDelete struct {
	// URI a file:// URI for the location of the file/folder being deleted.
	//
	// @since 3.16.0
	URI string `json:"uri"`
}

// DeleteFilesParams the parameters sent in notifications/requests for user-initiated deletes of files.
//
// @since 3.16.0
type DeleteFilesParams struct {
	// Files an array of all files/folders deleted in this operation.
	//
	// @since 3.16.0
	Files []FileDelete `json:"files"`
}

type DidChangeConfigurationClientCapabilities struct {
	// DynamicRegistration did change configuration notification supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type DidChangeWatchedFilesClientCapabilities struct {
	// DynamicRegistration did change watched files notification supports dynamic registration. Please note that the current protocol doesn't support static configuration for file changes from the server side.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// RelativePatternSupport whether the client has support for RelativePattern relative pattern or not.
	RelativePatternSupport bool `json:"relativePatternSupport,omitempty"`
}

// ClientSymbolKindOptions.
//
// @since 3.18.0 proposed
type ClientSymbolKindOptions struct {
	// ValueSet the symbol kind values the client supports. When this property exists the client also guarantees that it will handle values outside its set gracefully and falls back to a default value when unknown. If this property is not present the client only supports the symbol kinds from `File` to `Array` as defined in the initial version of the protocol.
	//
	// @since 3.18.0 proposed
	ValueSet []SymbolKind `json:"valueSet,omitempty"`
}

// ClientSymbolTagOptions.
//
// @since 3.18.0 proposed
type ClientSymbolTagOptions struct {
	// ValueSet the tags supported by the client.
	//
	// @since 3.18.0 proposed
	ValueSet []SymbolTag `json:"valueSet"`
}

// ClientSymbolResolveOptions.
//
// @since 3.18.0 proposed
type ClientSymbolResolveOptions struct {
	// Properties the properties that a client can resolve lazily. Usually `location.range`.
	//
	// @since 3.18.0 proposed
	Properties []string `json:"properties"`
}

// WorkspaceSymbolClientCapabilities client capabilities for a WorkspaceSymbolRequest.
type WorkspaceSymbolClientCapabilities struct {
	// DynamicRegistration symbol request supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// SymbolKind specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
	SymbolKind *ClientSymbolKindOptions `json:"symbolKind,omitempty"`

	// TagSupport the client supports tags on `SymbolInformation`. Clients supporting tags have to handle unknown tags
	// gracefully.
	TagSupport *ClientSymbolTagOptions `json:"tagSupport,omitempty"`

	// ResolveSupport the client support partial workspace symbols. The client will send the request `workspaceSymbol/resolve` to the server to resolve additional properties.
	ResolveSupport *ClientSymbolResolveOptions `json:"resolveSupport,omitempty"`
}

// ExecuteCommandClientCapabilities the client capabilities of a ExecuteCommandRequest.
type ExecuteCommandClientCapabilities struct {
	// DynamicRegistration execute command supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type WorkspaceFoldersServerCapabilities struct {
	// Supported the server has support for workspace folders.
	Supported bool `json:"supported,omitempty"`

	// ChangeNotifications whether the server wants to receive workspace folder change notifications. If a string is provided the string is treated as an ID under which the notification is registered on the client side. The ID can be used to unregister for these events using the `client/unregisterCapability` request.
	ChangeNotifications WorkspaceFoldersServerCapabilitiesChangeNotifications `json:"changeNotifications,omitempty"`
}

// DidChangeConfigurationParams the parameters of a change configuration notification.
type DidChangeConfigurationParams struct {
	// Settings the actual changed settings.
	Settings any `json:"settings"`
}

type DidChangeConfigurationRegistrationOptions struct {
	Section DidChangeConfigurationRegistrationOptionsSection `json:"section,omitempty"`
}

// FileEvent an event describing a file change.
type FileEvent struct {
	// URI the file's uri.
	URI DocumentURI `json:"uri"`

	// Type the change type.
	Type FileChangeType `json:"type"`
}

// DidChangeWatchedFilesParams the watched files change notification's parameters.
type DidChangeWatchedFilesParams struct {
	// Changes the actual file events.
	Changes []FileEvent `json:"changes"`
}

type FileSystemWatcher struct {
	// GlobPattern the glob pattern to watch. See GlobPattern glob pattern for more detail. 3.17.0 support for relative
	// patterns.
	GlobPattern GlobPattern `json:"globPattern"`

	// Kind the kind of events of interest. If omitted it defaults to WatchKind.Create | WatchKind.Change | WatchKind.Delete which is .
	Kind WatchKind `json:"kind,omitempty"`
}

// DidChangeWatchedFilesRegistrationOptions describe options to be used when registered for text document change events.
type DidChangeWatchedFilesRegistrationOptions struct {
	// Watchers the watchers to register.
	Watchers []FileSystemWatcher `json:"watchers"`
}

// WorkspaceSymbolParams the parameters of a WorkspaceSymbolRequest.
type WorkspaceSymbolParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// Query a query string to filter symbols by. Clients may send an empty string here to request all symbols.
	Query string `json:"query"`
}

// WorkspaceSymbol a special workspace symbol that supports locations without a range. See also SymbolInformation.
//
// @since 3.17.0
type WorkspaceSymbol struct {
	// extends
	BaseSymbolInformation

	// Location the location of the symbol. Whether a server is allowed to return a location without a range depends
	// on the client capability `workspace.symbol.resolveSupport`. See SymbolInformation#location for
	// more details.
	//
	// @since 3.17.0
	Location WorkspaceSymbolLocation `json:"location"`

	// Data a data entry field that is preserved on a workspace symbol between a workspace symbol request and a workspace symbol resolve request.
	//
	// @since 3.17.0
	Data any `json:"data,omitempty"`
}

// WorkspaceSymbolRegistrationOptions registration options for a WorkspaceSymbolRequest.
type WorkspaceSymbolRegistrationOptions struct {
	// extends
	WorkspaceSymbolOptions
}

// ApplyWorkspaceEditParams the parameters passed via an apply workspace edit request.
type ApplyWorkspaceEditParams struct {
	// Label an optional label of the workspace edit. This label is presented in the user interface for example on an undo stack to undo the workspace edit.
	Label string `json:"label,omitempty"`

	// Edit the edits to apply.
	Edit WorkspaceEdit `json:"edit"`
}

// ApplyWorkspaceEditResult the result returned from the apply workspace edit request. 3.17 renamed from ApplyWorkspaceEditResponse.
//
// @since 3.17 renamed from ApplyWorkspaceEditResponse
type ApplyWorkspaceEditResult struct {
	// Applied indicates whether the edit was applied or not.
	//
	// @since 3.17 renamed from ApplyWorkspaceEditResponse
	Applied bool `json:"applied"`

	// FailureReason an optional textual description for why the edit was not applied. This may be used by the server for
	// diagnostic logging or to provide a suitable error for a request that triggered the edit.
	//
	// @since 3.17 renamed from ApplyWorkspaceEditResponse
	FailureReason string `json:"failureReason,omitempty"`

	// FailedChange depending on the client's failure handling strategy `failedChange` might contain the index of the change that failed. This property is only available if the client signals a `failureHandlingStrategy` in its client capabilities.
	//
	// @since 3.17 renamed from ApplyWorkspaceEditResponse
	FailedChange uint32 `json:"failedChange,omitempty"`
}

// RelativePattern a relative pattern is a helper to construct glob patterns that are matched relatively to a base URI.
// The common value for a `baseUri` is a workspace folder root, but it can be another absolute URI as well.
//
// @since 3.17.0
type RelativePattern struct {
	// BaseURI a workspace folder or a base URI to which this pattern will be matched against relatively.
	//
	// @since 3.17.0
	BaseURI RelativePatternBaseURI `json:"baseUri"`

	// Pattern the actual glob pattern;.
	//
	// @since 3.17.0
	Pattern Pattern `json:"pattern"`
}
