// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strconv"
)

// TraceValue represents a InitializeParams Trace mode.
type TraceValue string

// list of TraceValue.
const (
	// TraceOff disable tracing.
	TraceOff TraceValue = "off"

	// TraceMessage normal tracing mode.
	TraceMessage TraceValue = "message"

	// TraceVerbose verbose tracing mode.
	TraceVerbose TraceValue = "verbose"
)

// ClientInfo information about the client.
//
// @since 3.15.0.
type ClientInfo struct {
	// Name is the name of the client as defined by the client.
	Name string `json:"name"`

	// Version is the client's version as defined by the client.
	Version string `json:"version,omitempty"`
}

// InitializeParams params of Initialize Request.
type InitializeParams struct {
	WorkDoneProgressParams

	// ProcessID is the process Id of the parent process that started
	// the server. Is null if the process has not been started by another process.
	// If the parent process is not alive then the server should exit (see exit notification) its process.
	ProcessID int32 `json:"processId"`

	// ClientInfo is the information about the client.
	//
	// @since 3.15.0
	ClientInfo *ClientInfo `json:"clientInfo,omitempty"`

	// Locale is the locale the client is currently showing the user interface
	// in. This must not necessarily be the locale of the operating
	// system.
	//
	// Uses IETF language tags as the value's syntax
	// (See https://en.wikipedia.org/wiki/IETF_language_tag)
	//
	// @since 3.16.0.
	Locale string `json:"locale,omitempty"`

	// RootPath is the rootPath of the workspace. Is null
	// if no folder is open.
	//
	// Deprecated: Use RootURI instead.
	RootPath string `json:"rootPath,omitempty"`

	// RootURI is the rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and "rootUri" are set
	// "rootUri" wins.
	//
	// Deprecated: Use WorkspaceFolders instead.
	RootURI DocumentURI `json:"rootUri,omitempty"`

	// InitializationOptions user provided initialization options.
	InitializationOptions interface{} `json:"initializationOptions,omitempty"`

	// Capabilities is the capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`

	// Trace is the initial trace setting. If omitted trace is disabled ('off').
	Trace TraceValue `json:"trace,omitempty"`

	// WorkspaceFolders is the workspace folders configured in the client when the server starts.
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	//
	// @since 3.6.0.
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

// LogTraceParams params of LogTrace notification.
//
// @since 3.16.0.
type LogTraceParams struct {
	// Message is the message to be logged.
	Message string `json:"message"`

	// Verbose is the additional information that can be computed if the "trace" configuration
	// is set to "verbose".
	Verbose TraceValue `json:"verbose,omitempty"`
}

// SetTraceParams params of SetTrace notification.
//
// @since 3.16.0.
type SetTraceParams struct {
	// Value is the new value that should be assigned to the trace setting.
	Value TraceValue `json:"value"`
}

// FileOperationRegistrationOptions is the options to register for file operations.
//
// @since 3.16.0.
type FileOperationRegistrationOptions struct {
	// filters is the actual filters.
	Filters []FileOperationFilter `json:"filters"`
}

// FileOperationPatternKind is a pattern kind describing if a glob pattern matches a file a folder or
// both.
//
// @since 3.16.0.
type FileOperationPatternKind string

// list of FileOperationPatternKind.
const (
	// FileOperationPatternKindFile is the pattern matches a file only.
	FileOperationPatternKindFile FileOperationPatternKind = "file"

	// FileOperationPatternKindFolder is the pattern matches a folder only.
	FileOperationPatternKindFolder FileOperationPatternKind = "folder"
)

// FileOperationPatternOptions matching options for the file operation pattern.
//
// @since 3.16.0.
type FileOperationPatternOptions struct {
	// IgnoreCase is The pattern should be matched ignoring casing.
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}

// FileOperationPattern a pattern to describe in which file operation requests or notifications
// the server is interested in.
//
// @since 3.16.0.
type FileOperationPattern struct {
	// The glob pattern to match. Glob patterns can have the following syntax:
	//  - `*` to match one or more characters in a path segment
	//  - `?` to match on one character in a path segment
	//  - `**` to match any number of path segments, including none
	//  - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript
	//    and JavaScript files)
	//  - `[]` to declare a range of characters to match in a path segment
	//    (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	//  - `[!...]` to negate a range of characters to match in a path segment
	//    (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but
	//    not `example.0`)
	Glob string `json:"glob"`

	// Matches whether to match files or folders with this pattern.
	//
	// Matches both if undefined.
	Matches FileOperationPatternKind `json:"matches,omitempty"`

	// Options additional options used during matching.
	Options FileOperationPatternOptions `json:"options,omitempty"`
}

// FileOperationFilter is a filter to describe in which file operation requests or notifications
// the server is interested in.
//
// @since 3.16.0.
type FileOperationFilter struct {
	// Scheme is a URI like "file" or "untitled".
	Scheme string `json:"scheme,omitempty"`

	// Pattern is the actual file operation pattern.
	Pattern FileOperationPattern `json:"pattern"`
}

// CreateFilesParams is the parameters sent in notifications/requests for user-initiated creation
// of files.
//
// @since 3.16.0.
type CreateFilesParams struct {
	// Files an array of all files/folders created in this operation.
	Files []FileCreate `json:"files"`
}

// FileCreate nepresents information on a file/folder create.
//
// @since 3.16.0.
type FileCreate struct {
	// URI is a file:// URI for the location of the file/folder being created.
	URI string `json:"uri"`
}

// RenameFilesParams is the parameters sent in notifications/requests for user-initiated renames
// of files.
//
// @since 3.16.0.
type RenameFilesParams struct {
	// Files an array of all files/folders renamed in this operation. When a folder
	// is renamed, only the folder will be included, and not its children.
	Files []FileRename `json:"files"`
}

// FileRename represents information on a file/folder rename.
//
// @since 3.16.0.
type FileRename struct {
	// OldURI is a file:// URI for the original location of the file/folder being renamed.
	OldURI string `json:"oldUri"`

	// NewURI is a file:// URI for the new location of the file/folder being renamed.
	NewURI string `json:"newUri"`
}

// DeleteFilesParams is the parameters sent in notifications/requests for user-initiated deletes
// of files.
//
// @since 3.16.0.
type DeleteFilesParams struct {
	// Files an array of all files/folders deleted in this operation.
	Files []FileDelete `json:"files"`
}

// FileDelete represents information on a file/folder delete.
//
// @since 3.16.0.
type FileDelete struct {
	// URI is a file:// URI for the location of the file/folder being deleted.
	URI string `json:"uri"`
}

// ReferencesParams params of References Request.
//
// @since 3.15.0.
type ReferencesParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	// Context is the ReferenceParams context.
	Context ReferenceContext `json:"context"`
}

// DocumentHighlightOptions registration option of DocumentHighlight server capability.
//
// @since 3.15.0.
type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

// DocumentHighlightParams params of DocumentHighlight Request.
//
// @since 3.15.0.
type DocumentHighlightParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// DocumentSymbolOptions registration option of DocumentSymbol server capability.
//
// @since 3.15.0.
type DocumentSymbolOptions struct {
	WorkDoneProgressOptions

	// Label a human-readable string that is shown when multiple outlines trees
	// are shown for the same document.
	//
	// @since 3.16.0.
	Label string `json:"label,omitempty"`
}

// WorkspaceSymbolOptions registration option of WorkspaceSymbol server capability.
//
// @since 3.15.0.
type WorkspaceSymbolOptions struct {
	WorkDoneProgressOptions
}

// DocumentFormattingOptions registration option of DocumentFormatting server capability.
//
// @since 3.15.0.
type DocumentFormattingOptions struct {
	WorkDoneProgressOptions
}

// DocumentRangeFormattingOptions registration option of DocumentRangeFormatting server capability.
//
// @since 3.15.0.
type DocumentRangeFormattingOptions struct {
	WorkDoneProgressOptions
}

// DeclarationOptions registration option of Declaration server capability.
//
// @since 3.15.0.
type DeclarationOptions struct {
	WorkDoneProgressOptions
}

// DeclarationRegistrationOptions registration option of Declaration server capability.
//
// @since 3.15.0.
type DeclarationRegistrationOptions struct {
	DeclarationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// DeclarationParams params of Declaration Request.
//
// @since 3.15.0.
type DeclarationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// DefinitionOptions registration option of Definition server capability.
//
// @since 3.15.0.
type DefinitionOptions struct {
	WorkDoneProgressOptions
}

// DefinitionParams params of Definition Request.
//
// @since 3.15.0.
type DefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// TypeDefinitionOptions registration option of TypeDefinition server capability.
//
// @since 3.15.0.
type TypeDefinitionOptions struct {
	WorkDoneProgressOptions
}

// TypeDefinitionRegistrationOptions registration option of TypeDefinition server capability.
//
// @since 3.15.0.
type TypeDefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	TypeDefinitionOptions
	StaticRegistrationOptions
}

// TypeDefinitionParams params of TypeDefinition Request.
//
// @since 3.15.0.
type TypeDefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// ImplementationOptions registration option of Implementation server capability.
//
// @since 3.15.0.
type ImplementationOptions struct {
	WorkDoneProgressOptions
}

// ImplementationRegistrationOptions registration option of Implementation server capability.
//
// @since 3.15.0.
type ImplementationRegistrationOptions struct {
	TextDocumentRegistrationOptions
	ImplementationOptions
	StaticRegistrationOptions
}

// ImplementationParams params of Implementation Request.
//
// @since 3.15.0.
type ImplementationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// DocumentColorOptions registration option of DocumentColor server capability.
//
// @since 3.15.0.
type DocumentColorOptions struct {
	WorkDoneProgressOptions
}

// DocumentColorRegistrationOptions registration option of DocumentColor server capability.
//
// @since 3.15.0.
type DocumentColorRegistrationOptions struct {
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
	DocumentColorOptions
}

// PrepareSupportDefaultBehavior default behavior of PrepareSupport.
//
// @since 3.16.0.
type PrepareSupportDefaultBehavior float64

// list of PrepareSupportDefaultBehavior.
const (
	// PrepareSupportDefaultBehaviorIdentifier is the client's default behavior is to select the identifier
	// according the to language's syntax rule.
	PrepareSupportDefaultBehaviorIdentifier PrepareSupportDefaultBehavior = 1
)

// String returns a string representation of the PrepareSupportDefaultBehavior.
func (k PrepareSupportDefaultBehavior) String() string {
	switch k {
	case PrepareSupportDefaultBehaviorIdentifier:
		return "Identifier"
	default:
		return strconv.FormatFloat(float64(k), 'f', -10, 64)
	}
}

// FoldingRangeOptions registration option of FoldingRange server capability.
//
// @since 3.15.0.
type FoldingRangeOptions struct {
	WorkDoneProgressOptions
}

// FoldingRangeRegistrationOptions registration option of FoldingRange server capability.
//
// @since 3.15.0.
type FoldingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	FoldingRangeOptions
	StaticRegistrationOptions
}

// ShowDocumentParams params to show a document.
//
// @since 3.16.0.
type ShowDocumentParams struct {
	// URI is the document uri to show.
	URI URI `json:"uri"`

	// External indicates to show the resource in an external program.
	// To show for example `https://code.visualstudio.com/`
	// in the default WEB browser set `external` to `true`.
	External bool `json:"external,omitempty"`

	// TakeFocus an optional property to indicate whether the editor
	// showing the document should take focus or not.
	// Clients might ignore this property if an external
	// program is started.
	TakeFocus bool `json:"takeFocus,omitempty"`

	// Selection an optional selection range if the document is a text
	// document. Clients might ignore the property if an
	// external program is started or the file is not a text
	// file.
	Selection *Range `json:"selection,omitempty"`
}

// ShowDocumentResult is the result of an show document request.
//
// @since 3.16.0.
type ShowDocumentResult struct {
	// Success a boolean indicating if the show was successful.
	Success bool `json:"success"`
}

// ServerInfo Information about the server.
//
// @since 3.15.0.
type ServerInfo struct {
	// Name is the name of the server as defined by the server.
	Name string `json:"name"`

	// Version is the server's version as defined by the server.
	Version string `json:"version,omitempty"`
}

// InitializeError known error codes for an "InitializeError".
type InitializeError struct {
	// Retry indicates whether the client execute the following retry logic:
	// (1) show the message provided by the ResponseError to the user
	// (2) user selects retry or cancel
	// (3) if user selected retry the initialize method is sent again.
	Retry bool `json:"retry,omitempty"`
}

// TextDocumentSyncKind defines how the host (editor) should sync document changes to the language server.
type TextDocumentSyncKind float64

const (
	// TextDocumentSyncKindNone documents should not be synced at all.
	TextDocumentSyncKindNone TextDocumentSyncKind = 0

	// TextDocumentSyncKindFull documents are synced by always sending the full content
	// of the document.
	TextDocumentSyncKindFull TextDocumentSyncKind = 1

	// TextDocumentSyncKindIncremental documents are synced by sending the full content on open.
	// After that only incremental updates to the document are
	// send.
	TextDocumentSyncKindIncremental TextDocumentSyncKind = 2
)

// String implements fmt.Stringer.
func (k TextDocumentSyncKind) String() string {
	switch k {
	case TextDocumentSyncKindNone:
		return "None"
	case TextDocumentSyncKindFull:
		return "Full"
	case TextDocumentSyncKindIncremental:
		return "Incremental"
	default:
		return strconv.FormatFloat(float64(k), 'f', -10, 64)
	}
}

// CompletionOptions Completion options.
type CompletionOptions struct {
	// The server provides support to resolve additional
	// information for a completion item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	// The characters that trigger completion automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
}

// SignatureHelpOptions SignatureHelp options.
type SignatureHelpOptions struct {
	// The characters that trigger signature help
	// automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// RetriggerCharacters is the slist of characters that re-trigger signature help.
	//
	// These trigger characters are only active when signature help is already
	// showing.
	// All trigger characters are also counted as re-trigger characters.
	//
	// @since 3.15.0.
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

// ReferencesOptions ReferencesProvider options.
//
// @since 3.15.0.
type ReferencesOptions struct {
	WorkDoneProgressOptions
}

// WorkDoneProgressOptions WorkDoneProgress options.
//
// @since 3.15.0.
type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

// CodeActionOptions CodeAction options.
type CodeActionOptions struct {
	// CodeActionKinds that this server may return.
	//
	// The list of kinds may be generic, such as "CodeActionKind.Refactor", or the server
	// may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`

	// ResolveProvider is the server provides support to resolve additional
	// information for a code action.
	//
	// @since 3.16.0.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// CodeLensOptions CodeLens options.
type CodeLensOptions struct {
	// Code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DocumentOnTypeFormattingOptions format document on type options.
type DocumentOnTypeFormattingOptions struct {
	// FirstTriggerCharacter a character on which formatting should be triggered, like "}".
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

	// MoreTriggerCharacter more trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
}

// RenameOptions rename options.
type RenameOptions struct {
	// PrepareProvider renames should be checked and tested before being executed.
	PrepareProvider bool `json:"prepareProvider,omitempty"`
}

// DocumentLinkOptions document link options.
type DocumentLinkOptions struct {
	// ResolveProvider document links have a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// ExecuteCommandOptions execute command options.
type ExecuteCommandOptions struct {
	// Commands is the commands to be executed on the server
	Commands []string `json:"commands"`
}

// SaveOptions save options.
type SaveOptions struct {
	// IncludeText is the client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
}

// TextDocumentSyncOptions TextDocumentSync options.
type TextDocumentSyncOptions struct {
	// OpenClose open and close notifications are sent to the server.
	OpenClose bool `json:"openClose,omitempty"`

	// Change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full
	// and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
	Change TextDocumentSyncKind `json:"change,omitempty"`

	// WillSave notifications are sent to the server.
	WillSave bool `json:"willSave,omitempty"`

	// WillSaveWaitUntil will save wait until requests are sent to the server.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

	// Save notifications are sent to the server.
	Save *SaveOptions `json:"save,omitempty"`
}

// HoverOptions option of hover provider server capabilities.
type HoverOptions struct {
	WorkDoneProgressOptions
}

// SemanticTokensOptions option of semantic tokens provider server capabilities.
//
// @since 3.16.0.
type SemanticTokensOptions struct {
	WorkDoneProgressOptions
}

// SemanticTokensRegistrationOptions registration option of semantic tokens provider server capabilities.
//
// @since 3.16.0.
type SemanticTokensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SemanticTokensOptions
	StaticRegistrationOptions
}

// LinkedEditingRangeOptions option of linked editing range provider server capabilities.
//
// @since 3.16.0.
type LinkedEditingRangeOptions struct {
	WorkDoneProgressOptions
}

// LinkedEditingRangeRegistrationOptions registration option of linked editing range provider server capabilities.
//
// @since 3.16.0.
type LinkedEditingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	LinkedEditingRangeOptions
	StaticRegistrationOptions
}

// LinkedEditingRangeParams params for the LinkedEditingRange request.
//
// @since 3.16.0.
type LinkedEditingRangeParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// LinkedEditingRanges result of LinkedEditingRange request.
//
// @since 3.16.0.
type LinkedEditingRanges struct {
	// Ranges a list of ranges that can be renamed together.
	//
	// The ranges must have identical length and contain identical text content.
	//
	// The ranges cannot overlap.
	Ranges []Range `json:"ranges"`

	// WordPattern an optional word pattern (regular expression) that describes valid contents for
	// the given ranges.
	//
	// If no pattern is provided, the client configuration's word pattern will be used.
	WordPattern string `json:"wordPattern,omitempty"`
}

// MonikerOptions option of moniker provider server capabilities.
//
// @since 3.16.0.
type MonikerOptions struct {
	WorkDoneProgressOptions
}

// MonikerRegistrationOptions registration option of moniker provider server capabilities.
//
// @since 3.16.0.
type MonikerRegistrationOptions struct {
	TextDocumentRegistrationOptions
	MonikerOptions
}

// MonikerParams params for the Moniker request.
//
// @since 3.16.0.
type MonikerParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// UniquenessLevel is the Moniker uniqueness level to define scope of the moniker.
//
// @since 3.16.0.
type UniquenessLevel string

// list of UniquenessLevel.
const (
	// UniquenessLevelDocument is the moniker is only unique inside a document.
	UniquenessLevelDocument UniquenessLevel = "document"

	// UniquenessLevelProject is the moniker is unique inside a project for which a dump got created.
	UniquenessLevelProject UniquenessLevel = "project"

	// UniquenessLevelGroup is the moniker is unique inside the group to which a project belongs.
	UniquenessLevelGroup UniquenessLevel = "group"

	// UniquenessLevelScheme is the moniker is unique inside the moniker scheme.
	UniquenessLevelScheme UniquenessLevel = "scheme"

	// UniquenessLevelGlobal is the moniker is globally unique.
	UniquenessLevelGlobal UniquenessLevel = "global"
)

// MonikerKind is the moniker kind.
//
// @since 3.16.0.
type MonikerKind string

// list of MonikerKind.
const (
	// MonikerKindImport is the moniker represent a symbol that is imported into a project.
	MonikerKindImport MonikerKind = "import"

	// MonikerKindExport is the moniker represents a symbol that is exported from a project.
	MonikerKindExport MonikerKind = "export"

	// MonikerKindLocal is the moniker represents a symbol that is local to a project (e.g. a local
	// variable of a function, a class not visible outside the project, ...).
	MonikerKindLocal MonikerKind = "local"
)

// Moniker definition to match LSIF 0.5 moniker definition.
//
// @since 3.16.0.
type Moniker struct {
	// Scheme is the scheme of the moniker. For example tsc or .Net.
	Scheme string `json:"scheme"`

	// Identifier is the identifier of the moniker.
	//
	// The value is opaque in LSIF however schema owners are allowed to define the structure if they want.
	Identifier string `json:"identifier"`

	// Unique is the scope in which the moniker is unique.
	Unique UniquenessLevel `json:"unique"`

	// Kind is the moniker kind if known.
	Kind MonikerKind `json:"kind,omitempty"`
}

// StaticRegistrationOptions staticRegistration options to be returned in the initialize request.
type StaticRegistrationOptions struct {
	// ID is the id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	ID string `json:"id,omitempty"`
}

// DocumentLinkRegistrationOptions DocumentLinkRegistration options.
type DocumentLinkRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// ResolveProvider document links have a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// InitializedParams params of Initialized Notification.
type InitializedParams struct{}

// FailureHandlingKind is the kind of failure handling .
type FailureHandlingKind string

const (
	// Abort applying the workspace change is simply aborted if one of the changes provided
	// fails. All operations executed before the failing operation stay executed.
	Abort FailureHandlingKind = "abort"

	// Transactional all operations are executed transactional. That means they either all
	// succeed or no changes at all are applied to the workspace.
	Transactional FailureHandlingKind = "transactional"

	// TextOnlyTransactional if the workspace edit contains only textual file changes they are executed transactional.
	// If resource changes (create, rename or delete file) are part of the change the failure
	// handling strategy is abort.
	TextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"

	// Undo the client tries to undo the operations already executed. But there is no
	// guarantee that this is succeeding.
	Undo FailureHandlingKind = "undo"
)

// WorkspaceFolders represents a slice of WorkspaceFolder.
type WorkspaceFolders []WorkspaceFolder
