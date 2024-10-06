// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"go.lsp.dev/uri"
)

type TraceValue string

const (
	// OffTraceValue turn tracing off.
	OffTraceValue TraceValue = "off"

	// MessagesTraceValue trace messages only.
	MessagesTraceValue TraceValue = "messages"

	// VerboseTraceValue verbose message tracing.
	VerboseTraceValue TraceValue = "verbose"
)

// MarkupKind describes the content type that a client supports in various result literals like `Hover`, `ParameterInfo` or `CompletionItem`. Please note that `MarkupKinds` must not start with a `$`. This kinds are
// reserved for internal usage.
type MarkupKind string

const (
	// PlainTextMarkupKind plain text is supported as a content format.
	PlainTextMarkupKind MarkupKind = "plaintext"

	// MarkdownMarkupKind markdown is supported as a content format.
	MarkdownMarkupKind MarkupKind = "markdown"
)

// LanguageKind predefined Language kinds  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type LanguageKind string

const (
	AbapLanguageKind LanguageKind = "abap"

	WindowsBatLanguageKind LanguageKind = "bat"

	BibTeXLanguageKind LanguageKind = "bibtex"

	ClojureLanguageKind LanguageKind = "clojure"

	CoffeescriptLanguageKind LanguageKind = "coffeescript"

	CLanguageKind LanguageKind = "c"

	CppLanguageKind LanguageKind = "cpp"

	CsharpLanguageKind LanguageKind = "csharp"

	CssLanguageKind LanguageKind = "css"

	// DLanguageKind.
	//
	// @since 3.18.0 proposed
	DLanguageKind LanguageKind = "d"

	// DelphiLanguageKind.
	//
	// @since 3.18.0 proposed
	DelphiLanguageKind LanguageKind = "pascal"

	DiffLanguageKind LanguageKind = "diff"

	DartLanguageKind LanguageKind = "dart"

	DockerfileLanguageKind LanguageKind = "dockerfile"

	ElixirLanguageKind LanguageKind = "elixir"

	ErlangLanguageKind LanguageKind = "erlang"

	FsharpLanguageKind LanguageKind = "fsharp"

	GitCommitLanguageKind LanguageKind = "git-commit"

	GitRebaseLanguageKind LanguageKind = "rebase"

	GoLanguageKind LanguageKind = "go"

	GroovyLanguageKind LanguageKind = "groovy"

	HandlebarsLanguageKind LanguageKind = "handlebars"

	HaskellLanguageKind LanguageKind = "haskell"

	HTMLLanguageKind LanguageKind = "html"

	IniLanguageKind LanguageKind = "ini"

	JavaLanguageKind LanguageKind = "java"

	JavaScriptLanguageKind LanguageKind = "javascript"

	JavaScriptReactLanguageKind LanguageKind = "javascriptreact"

	JSONLanguageKind LanguageKind = "json"

	LaTeXLanguageKind LanguageKind = "latex"

	LessLanguageKind LanguageKind = "less"

	LuaLanguageKind LanguageKind = "lua"

	MakefileLanguageKind LanguageKind = "makefile"

	MarkdownLanguageKind LanguageKind = "markdown"

	ObjectiveCLanguageKind LanguageKind = "objective-c"

	ObjectiveCPPLanguageKind LanguageKind = "objective-cpp"

	// PascalLanguageKind.
	//
	// @since 3.18.0 proposed
	PascalLanguageKind LanguageKind = "pascal"

	PerlLanguageKind LanguageKind = "perl"

	Perl6LanguageKind LanguageKind = "perl6"

	PhpLanguageKind LanguageKind = "php"

	PowershellLanguageKind LanguageKind = "powershell"

	PugLanguageKind LanguageKind = "jade"

	PythonLanguageKind LanguageKind = "python"

	RLanguageKind LanguageKind = "r"

	RazorLanguageKind LanguageKind = "razor"

	RubyLanguageKind LanguageKind = "ruby"

	RustLanguageKind LanguageKind = "rust"

	ScssLanguageKind LanguageKind = "scss"

	SassLanguageKind LanguageKind = "sass"

	ScalaLanguageKind LanguageKind = "scala"

	ShaderLabLanguageKind LanguageKind = "shaderlab"

	ShellScriptLanguageKind LanguageKind = "shellscript"

	SQLLanguageKind LanguageKind = "sql"

	SwiftLanguageKind LanguageKind = "swift"

	TypeScriptLanguageKind LanguageKind = "typescript"

	TypeScriptReactLanguageKind LanguageKind = "typescriptreact"

	TeXLanguageKind LanguageKind = "tex"

	VisualBasicLanguageKind LanguageKind = "vb"

	XmlLanguageKind LanguageKind = "xml"

	XslLanguageKind LanguageKind = "xsl"

	YamlLanguageKind LanguageKind = "yaml"
)

// PositionEncodingKind a set of predefined position encoding kinds.
//
// @since 3.17.0
type PositionEncodingKind string

const (
	// UTF8PositionEncodingKind character offsets count UTF-8 code units (e.g. bytes).
	UTF8PositionEncodingKind PositionEncodingKind = "utf-8"

	// Utf16PositionEncodingKind character offsets count UTF-16 code units. This is the default and must always be supported by servers.
	Utf16PositionEncodingKind PositionEncodingKind = "utf-16"

	// Utf32PositionEncodingKind character offsets count UTF-32 code units. Implementation note: these are the same as Unicode codepoints, so this `PositionEncodingKind` may also be used for an encoding-agnostic representation of character offsets.
	Utf32PositionEncodingKind PositionEncodingKind = "utf-32"
)

// DiagnosticSeverity the diagnostic's severity.
type DiagnosticSeverity uint32

const (
	// ErrorDiagnosticSeverity reports an error.
	ErrorDiagnosticSeverity DiagnosticSeverity = 1

	// WarningDiagnosticSeverity reports a warning.
	WarningDiagnosticSeverity DiagnosticSeverity = 2

	// InformationDiagnosticSeverity reports an information.
	InformationDiagnosticSeverity DiagnosticSeverity = 3

	// HintDiagnosticSeverity reports a hint.
	HintDiagnosticSeverity DiagnosticSeverity = 4
)

// DiagnosticTag the diagnostic tags.
//
// @since 3.15.0
type DiagnosticTag uint32

const (
	// UnnecessaryDiagnosticTag unused or unnecessary code. Clients are allowed to render diagnostics with this tag faded out instead of having an error squiggle.
	UnnecessaryDiagnosticTag DiagnosticTag = 1

	// DeprecatedDiagnosticTag deprecated or obsolete code. Clients are allowed to rendered diagnostics with this tag strike through.
	DeprecatedDiagnosticTag DiagnosticTag = 2
)

type ResourceOperationKind string

const (
	// CreateResourceOperationKind supports creating new files and folders.
	CreateResourceOperationKind ResourceOperationKind = "create"

	// RenameResourceOperationKind supports renaming existing files and folders.
	RenameResourceOperationKind ResourceOperationKind = "rename"

	// DeleteResourceOperationKind supports deleting existing files and folders.
	DeleteResourceOperationKind ResourceOperationKind = "delete"
)

type FailureHandlingKind string

const (
	// AbortFailureHandlingKind applying the workspace change is simply aborted if one of the changes provided fails. All operations
	// executed before the failing operation stay executed.
	AbortFailureHandlingKind FailureHandlingKind = "abort"

	// TransactionalFailureHandlingKind all operations are executed transactional. That means they either all succeed or no changes at all are applied to the workspace.
	TransactionalFailureHandlingKind FailureHandlingKind = "transactional"

	// TextOnlyTransactionalFailureHandlingKind if the workspace edit contains only textual file changes they are executed transactional. If resource changes (create, rename or delete file) are part of the change the failure handling strategy is abort.
	TextOnlyTransactionalFailureHandlingKind FailureHandlingKind = "textOnlyTransactional"

	// UndoFailureHandlingKind the client tries to undo the operations already executed. But there is no guarantee that this is succeeding.
	UndoFailureHandlingKind FailureHandlingKind = "undo"
)

// TextDocumentIdentifier a literal to identify a text document in the client.
type TextDocumentIdentifier struct {
	// URI the text document's uri.
	URI DocumentURI `json:"uri"`
}

// Position position in a text document expressed as zero-based line and character offset. Prior to 3.17 the offsets were always based on a UTF-16 string representation. So a string of the form `a𐐀b` the character offset of the character `a` is 0, the character offset of `𐐀` is 1 and the character offset of b is 3 since `𐐀` is represented using two code units in UTF-16. Since 3.17 clients and servers
// can agree on a different string encoding representation (e.g. UTF-8). The client announces it's supported encoding via the client capability [`general.positionEncodings`](https://microsoft.github.io/language-server-protocol/specifications/specification-current/#clientCapabilities). The value is an array of position encodings the client supports, with decreasing preference (e.g. the encoding at index `0` is the most preferred one). To stay backwards compatible the only mandatory encoding is
// UTF-16 represented via the string `utf-16`. The server can pick one of the encodings offered by the client and signals that encoding back to the client via the initialize result's property [`capabilities.positionEncoding`](https://microsoft.github.io/language-server-protocol/specifications/specification-current/#serverCapabilities). If the string value `utf-16` is missing from the client's capability `general.positionEncodings` servers can safely assume that the client supports UTF-16. If the server omits the position encoding in its initialize result the encoding defaults to the string value `utf-16`. Implementation considerations: since the conversion from one encoding into another requires the content of the file / line the conversion is best done where the file is read which is usually on the server side. Positions are line end character agnostic. So you can not specify a position that denotes `\r|\n` or `\n|` where `|` represents the character offset. 3.17.0 - support for negotiated position encoding.
//
// @since 3.17.0 - support for negotiated position encoding.
type Position struct {
	// Line line position in a document (zero-based).
	//
	// @since 3.17.0 - support for negotiated position encoding.
	Line uint32 `json:"line"`

	// Character character offset on a line in a document (zero-based). The meaning of this offset is determined by the negotiated `PositionEncodingKind`.
	//
	// @since 3.17.0 - support for negotiated position encoding.
	Character uint32 `json:"character"`
}

// TextDocumentPositionParams a parameter literal used in requests to pass a text document and a position inside that document.
type TextDocumentPositionParams struct {
	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position the position inside the text document.
	Position Position `json:"position"`
}

// Range a range in a text document expressed as (zero-based) start and end positions. If you want to specify
// a range that contains a line including the line ending character(s) then use an end position denoting the start of the next line. For example: ```ts { start: { line: 5, character: 23 } end : { line 6, character : 0 } } ```.
type Range struct {
	// Start the range's start position.
	Start Position `json:"start"`

	// End the range's end position.
	End Position `json:"end"`
}

// Location represents a location inside a resource, such as a line inside a text file.
type Location struct {
	URI DocumentURI `json:"uri"`

	Range Range `json:"range"`
}

// TextEdit a text edit applicable to a text document.
type TextEdit struct {
	// Range the range of the text document to be manipulated. To insert text into a document create a range where start === end.
	Range Range `json:"range"`

	// NewText the string to be inserted. For delete operations use an empty string.
	NewText string `json:"newText"`
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

// OptionalVersionedTextDocumentIdentifier a text document identifier to optionally denote a specific version of a text document.
type OptionalVersionedTextDocumentIdentifier struct {
	// extends
	TextDocumentIdentifier

	// Version the version number of this document. If a versioned text document identifier is sent from the server
	// to the client and the file is not open in the editor (the server has not received an open notification before) the server can send `null` to indicate that the version is unknown and the content on disk is the truth (as specified with document content ownership).
	Version int32 `json:"version,omitempty"`
}

// StringValue a string value used as a snippet is a template which allows to insert text and to control the editor
// cursor when insertion happens. A snippet can define tab stops and placeholders with `$1`, `$2`
// and `${3:foo}`. `$0` defines the final tab stop, it defaults to the end of the snippet. Variables are defined with `$name` and `${name:default value}`. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
type StringValue struct {
	// Value the snippet string.
	//
	// @since 3.18.0 proposed
	Value string `json:"value"`
}

// SnippetTextEdit an interactive text edit.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type SnippetTextEdit struct {
	// Range the range of the text document to be manipulated.
	//
	// @since 3.18.0 proposed
	Range Range `json:"range"`

	// Snippet the snippet to be inserted.
	//
	// @since 3.18.0 proposed
	Snippet StringValue `json:"snippet"`

	// AnnotationID the actual identifier of the snippet edit.
	//
	// @since 3.18.0 proposed
	AnnotationID *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
}

// AnnotatedTextEdit a special text edit with an additional change annotation.
//
// @since 3.16.0.
type AnnotatedTextEdit struct {
	// extends
	TextEdit

	// AnnotationID the actual identifier of the change annotation.
	//
	// @since 3.16.0.
	AnnotationID ChangeAnnotationIdentifier `json:"annotationId"`
}

// TextDocumentEdit describes textual changes on a text document. A TextDocumentEdit describes all changes on a document
// version Si and after they are applied move the document to version Si+1. So the creator of a TextDocumentEdit doesn't need to sort the array of edits or do any kind of ordering. However the edits must be non overlapping.
type TextDocumentEdit struct {
	// TextDocument the text document to change.
	TextDocument OptionalVersionedTextDocumentIdentifier `json:"textDocument"`

	// Edits the edits to be applied. 3.16.0 - support for AnnotatedTextEdit. This is guarded using a client capability. 3.18.0 - support for SnippetTextEdit. This is guarded using a client capability.
	Edits TextDocumentEditEdits `json:"edits"`
}

// ChangeAnnotation additional information that describes document changes.
//
// @since 3.16.0
type ChangeAnnotation struct {
	// Label a human-readable string describing the actual change. The string is rendered prominent in the user interface.
	//
	// @since 3.16.0
	Label string `json:"label"`

	// NeedsConfirmation a flag which indicates that user confirmation is needed before applying the change.
	//
	// @since 3.16.0
	NeedsConfirmation bool `json:"needsConfirmation,omitempty"`

	// Description a human-readable string which is rendered less prominent in the user interface.
	//
	// @since 3.16.0
	Description string `json:"description,omitempty"`
}

// WorkspaceEdit a workspace edit represents changes to many resources managed in the workspace. The edit should either provide `changes` or `documentChanges`. If documentChanges are present they are preferred over `changes` if the client can handle versioned document edits. Since version 3.13.0 a workspace edit can
// contain resource operations as well. If resource operations are present clients need to execute the operations in the order in which they are provided. So a workspace edit for example can consist of the following two changes: (1) a create file a.txt and (2) a text document edit which insert text into file a.txt. An invalid sequence (e.g. (1) delete file a.txt and (2) insert text into file a.txt) will cause failure of the operation. How the client recovers from the failure is described by the client capability: `workspace.workspaceEdit.failureHandling`.
type WorkspaceEdit struct {
	// Changes holds changes to existing resources.
	Changes map[DocumentURI][]TextEdit `json:"changes,omitempty"`

	// DocumentChanges depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes are
	// either an array of `TextDocumentEdit`s to express changes to n different text documents where each text document edit addresses a specific version of a text document. Or it can contain above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations. Whether a client supports versioned document edits is expressed via `workspace.workspaceEdit.documentChanges` client capability. If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then only plain `TextEdit`s using the `changes` property are supported.
	DocumentChanges WorkspaceEditDocumentChanges `json:"documentChanges,omitempty"`

	// ChangeAnnotations a map of change annotations that can be referenced in `AnnotatedTextEdit`s or create, rename and delete file / folder operations. Whether clients honor this property depends on the client capability `workspace.changeAnnotationSupport`.
	ChangeAnnotations map[ChangeAnnotationIdentifier]ChangeAnnotation `json:"changeAnnotations,omitempty"`
}

// MarkupContent a `MarkupContent` literal represents a string value which content is interpreted base on its kind flag. Currently the protocol supports `plaintext` and `markdown` as markup kinds. If the kind is `markdown` then the value can contain fenced code blocks like in GitHub issues. See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting Here is an example how such a
// string can be constructed using JavaScript / TypeScript: ```ts let markdown: MarkdownContent =
// { kind: MarkupKind.Markdown, value: [ '# Header', 'Some text', '```typescript', 'someCode();',
// '```' ].join('\n') }; ``` *Please Note* that clients might sanitize the return markdown. A client could decide to remove HTML from the markdown to avoid script execution.
type MarkupContent struct {
	// Kind the type of the Markup.
	Kind MarkupKind `json:"kind"`

	// Value the content itself.
	Value string `json:"value"`
}

// Command represents a reference to a command. Provides a title which will be used to represent a command in the UI and, optionally, an array of arguments which will be passed to the command handler function when invoked.
type Command struct {
	// Title title of the command, like `save`.
	Title string `json:"title"`

	// Tooltip an optional tooltip.  3.18.0 @proposed.
	Tooltip string `json:"tooltip,omitempty"`

	// Command the identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments arguments that the command handler should be invoked with.
	Arguments []any `json:"arguments,omitempty"`
}

// CodeDescription structure to capture a description for an error code.
//
// @since 3.16.0
type CodeDescription struct {
	// Href an URI to open with more information about the diagnostic error.
	//
	// @since 3.16.0
	Href uri.URI `json:"href"`
}

// DiagnosticRelatedInformation represents a related message and source code location for a diagnostic. This should be used to point
// to code locations that cause or related to a diagnostics, e.g when duplicating a symbol in a scope.
type DiagnosticRelatedInformation struct {
	// Location the location of this related diagnostic information.
	Location Location `json:"location"`

	// Message the message of this related diagnostic information.
	Message string `json:"message"`
}

// Diagnostic represents a diagnostic, such as a compiler error or warning. Diagnostic objects are only valid in the scope of a resource.
type Diagnostic struct {
	// Range the range at which the message applies.
	Range Range `json:"range"`

	// Severity the diagnostic's severity. To avoid interpretation mismatches when a server is used with different clients it is highly recommended that servers always provide a severity value.
	Severity DiagnosticSeverity `json:"severity,omitempty"`

	// Code the diagnostic's code, which usually appear in the user interface.
	Code DiagnosticCode `json:"code,omitempty"`

	// CodeDescription an optional property to describe the error code. Requires the code field (above) to be present/not null.
	CodeDescription *CodeDescription `json:"codeDescription,omitempty"`

	// Source a human-readable string describing the source of this diagnostic, e.g. 'typescript' or 'super lint'.
	// It usually appears in the user interface.
	Source string `json:"source,omitempty"`

	// Message the diagnostic's message. It usually appears in the user interface.
	Message string `json:"message"`

	// Tags additional metadata about the diagnostic.
	Tags []DiagnosticTag `json:"tags,omitempty"`

	// RelatedInformation an array of related diagnostic information, e.g. when symbol-names within a scope collide all definitions can be marked via this property.
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`

	// Data a data entry field that is preserved between a `textDocument/publishDiagnostics` notification and `textDocument/codeAction` request.
	Data any `json:"data,omitempty"`
}

// TextDocumentItem an item to transfer a text document from the client to the server.
type TextDocumentItem struct {
	// URI the text document's uri.
	URI DocumentURI `json:"uri"`

	// LanguageID the text document's language identifier.
	LanguageID LanguageKind `json:"languageId"`

	// Version the version number of this document (it will increase after each change, including undo/redo).
	Version int32 `json:"version"`

	// Text the content of the opened text document.
	Text string `json:"text"`
}

// VersionedTextDocumentIdentifier a text document identifier to denote a specific version of a text document.
type VersionedTextDocumentIdentifier struct {
	// extends
	TextDocumentIdentifier

	// Version the version number of this document.
	Version int32 `json:"version"`
}

// TextDocumentContentParams parameters for the `workspace/textDocumentContent` request.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type TextDocumentContentParams struct {
	// URI the uri of the text document.
	//
	// @since 3.18.0 proposed
	URI DocumentURI `json:"uri"`
}

// TextDocumentContentResult result of the `workspace/textDocumentContent` request.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type TextDocumentContentResult struct {
	// Text the text content of the text document. Please note, that the content of any subsequent open notifications for the text document might differ from the returned content due to whitespace and line ending
	// normalizations done on the client.
	//
	// @since 3.18.0 proposed
	Text string `json:"text"`
}

// TextDocumentContentOptions text document content provider options.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type TextDocumentContentOptions struct {
	// Schemes the schemes for which the server provides content.
	//
	// @since 3.18.0 proposed
	Schemes []string `json:"schemes"`
}

// TextDocumentContentRegistrationOptions text document content provider registration options.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type TextDocumentContentRegistrationOptions struct {
	// extends
	TextDocumentContentOptions
	// mixins
	StaticRegistrationOptions
}

// TextDocumentContentRefreshParams parameters for the `workspace/textDocumentContent/refresh` request.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type TextDocumentContentRefreshParams struct {
	// URI the uri of the text document to refresh.
	//
	// @since 3.18.0 proposed
	URI DocumentURI `json:"uri"`
}

// ChangeAnnotationsSupportOptions.
//
// @since 3.18.0
type ChangeAnnotationsSupportOptions struct {
	// GroupsOnLabel whether the client groups edits with equal labels into tree nodes, for instance all edits labelled with "Changes in Strings" would be a tree node.
	//
	// @since 3.18.0
	GroupsOnLabel bool `json:"groupsOnLabel,omitempty"`
}

type WorkspaceEditClientCapabilities struct {
	// DocumentChanges the client supports versioned document changes in `WorkspaceEdit`s.
	DocumentChanges bool `json:"documentChanges,omitempty"`

	// ResourceOperations the resource operations the client supports. Clients should at least support 'create', 'rename' and 'delete' files and folders.
	ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitempty"`

	// FailureHandling the failure handling strategy of a client if applying the workspace edit fails.
	FailureHandling FailureHandlingKind `json:"failureHandling,omitempty"`

	// NormalizesLineEndings whether the client normalizes line endings to the client specific setting. If set to `true` the client will normalize line ending characters in a workspace edit to the client-specified new line character.
	NormalizesLineEndings bool `json:"normalizesLineEndings,omitempty"`

	// ChangeAnnotationSupport whether the client in general supports change annotations on text edits, create file, rename file and delete file changes.
	ChangeAnnotationSupport *ChangeAnnotationsSupportOptions `json:"changeAnnotationSupport,omitempty"`

	// MetadataSupport whether the client supports `WorkspaceEditMetadata` in `WorkspaceEdit`s.  3.18.0 @proposed.
	MetadataSupport bool `json:"metadataSupport,omitempty"`

	// SnippetEditSupport whether the client supports snippets as text edits.  3.18.0 @proposed.
	SnippetEditSupport bool `json:"snippetEditSupport,omitempty"`
}

type WorkDoneProgressBegin struct {
	// Title mandatory title of the progress operation. Used to briefly inform about the kind of operation being performed. Examples: "Indexing" or "Linking dependencies".
	Title string `json:"title"`

	// Cancellable controls if a cancel button should show to allow the user to cancel the long running operation. Clients that don't support cancellation are allowed to ignore the setting.
	Cancellable bool `json:"cancellable,omitempty"`

	// Message optional, more detailed associated progress message. Contains complementary information to the `title`. Examples: "3/25 files", "project/src/module2", "node_modules/some_dep". If unset, the previous progress message (if any) is still valid.
	Message string `json:"message,omitempty"`

	// Percentage optional progress percentage to display (value 100 is considered 100%). If not provided infinite progress is assumed and clients are allowed to ignore the `percentage` value in subsequent in report notifications. The value should be steadily rising. Clients are free to ignore values that are not following this rule. The value range is [0, 100].
	Percentage uint32 `json:"percentage,omitempty"`
}

type WorkDoneProgressReport struct {
	// Cancellable controls enablement state of a cancel button. Clients that don't support cancellation or don't support controlling the button's enablement state are allowed to ignore the property.
	Cancellable bool `json:"cancellable,omitempty"`

	// Message optional, more detailed associated progress message. Contains complementary information to the `title`. Examples: "3/25 files", "project/src/module2", "node_modules/some_dep". If unset, the previous progress message (if any) is still valid.
	Message string `json:"message,omitempty"`

	// Percentage optional progress percentage to display (value 100 is considered 100%). If not provided infinite progress is assumed and clients are allowed to ignore the `percentage` value in subsequent in report notifications. The value should be steadily rising. Clients are free to ignore values that are not following this rule. The value range is [0, 100].
	Percentage uint32 `json:"percentage,omitempty"`
}

type WorkDoneProgressEnd struct {
	// Message optional, a final message indicating to for example indicate the outcome of the operation.
	Message string `json:"message,omitempty"`
}

type WorkDoneProgressParams struct {
	// WorkDoneToken an optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}

type PartialResultParams struct {
	// PartialResultToken an optional token that a server can use to report partial results (e.g. streaming) to the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
}

// LocationLink represents the connection of two locations. Provides additional metadata over normal Location locations, including an origin range.
type LocationLink struct {
	// OriginSelectionRange span of the origin of this link. Used as the underlined span for mouse interaction. Defaults to the word range at the definition position.
	OriginSelectionRange *Range `json:"originSelectionRange,omitempty"`

	// TargetURI the target resource identifier of this link.
	TargetURI DocumentURI `json:"targetUri"`

	// TargetRange the full target range of this link. If the target for example is a symbol then target range is the range enclosing this symbol not including leading/trailing whitespace but everything else like comments. This information is typically used to highlight the range in the editor.
	TargetRange Range `json:"targetRange"`

	// TargetSelectionRange the range that should be selected and revealed when this link is being followed, e.g the name of a function. Must be contained by the `targetRange`. See also `DocumentSymbol#range`.
	TargetSelectionRange Range `json:"targetSelectionRange"`
}
