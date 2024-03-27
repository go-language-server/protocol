package protocol

import (
	"go.lsp.dev/uri"
)

// SemanticTokenTypes a set of predefined token types. This set is not fixed an clients can specify additional token types via the corresponding client capabilities.
//
// @since 3.16.0
type SemanticTokenTypes string

const (
	NamespaceSemanticTokenTypes SemanticTokenTypes = "namespace"

	// TypeSemanticTokenTypes represents a generic type. Acts as a fallback for types which can't be mapped to a specific type like class or enum.
	TypeSemanticTokenTypes SemanticTokenTypes = "type"

	ClassSemanticTokenTypes SemanticTokenTypes = "class"

	EnumSemanticTokenTypes SemanticTokenTypes = "enum"

	InterfaceSemanticTokenTypes SemanticTokenTypes = "interface"

	StructSemanticTokenTypes SemanticTokenTypes = "struct"

	TypeParameterSemanticTokenTypes SemanticTokenTypes = "typeParameter"

	ParameterSemanticTokenTypes SemanticTokenTypes = "parameter"

	VariableSemanticTokenTypes SemanticTokenTypes = "variable"

	PropertySemanticTokenTypes SemanticTokenTypes = "property"

	EnumMemberSemanticTokenTypes SemanticTokenTypes = "enumMember"

	EventSemanticTokenTypes SemanticTokenTypes = "event"

	FunctionSemanticTokenTypes SemanticTokenTypes = "function"

	MethodSemanticTokenTypes SemanticTokenTypes = "method"

	MacroSemanticTokenTypes SemanticTokenTypes = "macro"

	KeywordSemanticTokenTypes SemanticTokenTypes = "keyword"

	ModifierSemanticTokenTypes SemanticTokenTypes = "modifier"

	CommentSemanticTokenTypes SemanticTokenTypes = "comment"

	StringSemanticTokenTypes SemanticTokenTypes = "string"

	NumberSemanticTokenTypes SemanticTokenTypes = "number"

	RegexpSemanticTokenTypes SemanticTokenTypes = "regexp"

	OperatorSemanticTokenTypes SemanticTokenTypes = "operator"

	// DecoratorSemanticTokenTypes.
	//
	// @since 3.17.0
	DecoratorSemanticTokenTypes SemanticTokenTypes = "decorator"
)

// SemanticTokenModifiers a set of predefined token modifiers. This set is not fixed an clients can specify additional token types via the corresponding client capabilities.
//
// @since 3.16.0
type SemanticTokenModifiers string

const (
	DeclarationSemanticTokenModifiers SemanticTokenModifiers = "declaration"

	DefinitionSemanticTokenModifiers SemanticTokenModifiers = "definition"

	ReadonlySemanticTokenModifiers SemanticTokenModifiers = "readonly"

	StaticSemanticTokenModifiers SemanticTokenModifiers = "static"

	DeprecatedSemanticTokenModifiers SemanticTokenModifiers = "deprecated"

	AbstractSemanticTokenModifiers SemanticTokenModifiers = "abstract"

	AsyncSemanticTokenModifiers SemanticTokenModifiers = "async"

	ModificationSemanticTokenModifiers SemanticTokenModifiers = "modification"

	DocumentationSemanticTokenModifiers SemanticTokenModifiers = "documentation"

	DefaultLibrarySemanticTokenModifiers SemanticTokenModifiers = "defaultLibrary"
)

// DocumentDiagnosticReportKind the document diagnostic report kinds.
//
// @since 3.17.0
type DocumentDiagnosticReportKind string

const (
	// FullDocumentDiagnosticReportKind a diagnostic report with a full set of problems.
	FullDocumentDiagnosticReportKind DocumentDiagnosticReportKind = "full"

	// UnchangedDocumentDiagnosticReportKind a report indicating that the last returned report is still accurate.
	UnchangedDocumentDiagnosticReportKind DocumentDiagnosticReportKind = "unchanged"
)

// ErrorCodes predefined error codes.
type ErrorCodes int32

const (
	ParseErrorErrorCodes ErrorCodes = -32700

	InvalidRequestErrorCodes ErrorCodes = -32600

	MethodNotFoundErrorCodes ErrorCodes = -32601

	InvalidParamsErrorCodes ErrorCodes = -32602

	InternalErrorErrorCodes ErrorCodes = -32603

	// ServerNotInitializedErrorCodes error code indicating that a server received a notification or request before the server has received the `initialize` request.
	ServerNotInitializedErrorCodes ErrorCodes = -32002

	UnknownErrorCodeErrorCodes ErrorCodes = -32001
)

type LsperrorCodes int32

const (
	// RequestFailedLsperrorCodes a request failed but it was syntactically correct, e.g the method name was known and the parameters were valid. The error message should contain human readable information about why the request failed.
	//
	//
	//
	// @since 3.17.0
	RequestFailedLsperrorCodes LsperrorCodes = -32803

	// ServerCancelledLsperrorCodes the server cancelled the request. This error code should only be used for requests that explicitly support being server cancellable.
	//
	//
	//
	// @since 3.17.0
	ServerCancelledLsperrorCodes LsperrorCodes = -32802

	// ContentModifiedLsperrorCodes the server detected that the content of a document got modified outside normal conditions. A server should NOT send this error code if it detects a content change in it unprocessed messages. The result even computed on an older state might still be useful for the client. If a client decides that a result is not of any use anymore the client should cancel the request.
	ContentModifiedLsperrorCodes LsperrorCodes = -32801

	// RequestCancelledLsperrorCodes the client has canceled a request and a server as detected the cancel.
	RequestCancelledLsperrorCodes LsperrorCodes = -32800
)

// FoldingRangeKind a set of predefined range kinds.
type FoldingRangeKind string

const (
	// CommentFoldingRangeKind folding range for a comment.
	CommentFoldingRangeKind FoldingRangeKind = "comment"

	// ImportsFoldingRangeKind folding range for an import or include.
	ImportsFoldingRangeKind FoldingRangeKind = "imports"

	// RegionFoldingRangeKind folding range for a region (e.g. `#region`).
	RegionFoldingRangeKind FoldingRangeKind = "region"
)

// SymbolKind a symbol kind.
type SymbolKind uint32

const (
	FileSymbolKind SymbolKind = 1

	ModuleSymbolKind SymbolKind = 2

	NamespaceSymbolKind SymbolKind = 3

	PackageSymbolKind SymbolKind = 4

	ClassSymbolKind SymbolKind = 5

	MethodSymbolKind SymbolKind = 6

	PropertySymbolKind SymbolKind = 7

	FieldSymbolKind SymbolKind = 8

	ConstructorSymbolKind SymbolKind = 9

	EnumSymbolKind SymbolKind = 10

	InterfaceSymbolKind SymbolKind = 11

	FunctionSymbolKind SymbolKind = 12

	VariableSymbolKind SymbolKind = 13

	ConstantSymbolKind SymbolKind = 14

	StringSymbolKind SymbolKind = 15

	NumberSymbolKind SymbolKind = 16

	BooleanSymbolKind SymbolKind = 17

	ArraySymbolKind SymbolKind = 18

	ObjectSymbolKind SymbolKind = 19

	KeySymbolKind SymbolKind = 20

	NullSymbolKind SymbolKind = 21

	EnumMemberSymbolKind SymbolKind = 22

	StructSymbolKind SymbolKind = 23

	EventSymbolKind SymbolKind = 24

	OperatorSymbolKind SymbolKind = 25

	TypeParameterSymbolKind SymbolKind = 26
)

// SymbolTag symbol tags are extra annotations that tweak the rendering of a symbol.
//
// @since 3.16
type SymbolTag uint32

const (
	// DeprecatedSymbolTag render a symbol as obsolete, usually using a strike-out.
	DeprecatedSymbolTag SymbolTag = 1
)

// UniquenessLevel moniker uniqueness level to define scope of the moniker.
//
// @since 3.16.0
type UniquenessLevel string

const (
	// DocumentUniquenessLevel the moniker is only unique inside a document.
	DocumentUniquenessLevel UniquenessLevel = "document"

	// ProjectUniquenessLevel the moniker is unique inside a project for which a dump got created.
	ProjectUniquenessLevel UniquenessLevel = "project"

	// GroupUniquenessLevel the moniker is unique inside the group to which a project belongs.
	GroupUniquenessLevel UniquenessLevel = "group"

	// SchemeUniquenessLevel the moniker is unique inside the moniker scheme.
	SchemeUniquenessLevel UniquenessLevel = "scheme"

	// GlobalUniquenessLevel the moniker is globally unique.
	GlobalUniquenessLevel UniquenessLevel = "global"
)

// MonikerKind the moniker kind.
//
// @since 3.16.0
type MonikerKind string

const (
	// ImportMonikerKind the moniker represent a symbol that is imported into a project.
	ImportMonikerKind MonikerKind = "import"

	// ExportMonikerKind the moniker represents a symbol that is exported from a project.
	ExportMonikerKind MonikerKind = "export"

	// LocalMonikerKind the moniker represents a symbol that is local to a project (e.g. a local variable of a function, a class not visible outside the project, ...).
	LocalMonikerKind MonikerKind = "local"
)

// InlayHintKind inlay hint kinds.
//
// @since 3.17.0
type InlayHintKind uint32

const (
	// TypeInlayHintKind an inlay hint that for a type annotation.
	TypeInlayHintKind InlayHintKind = 1

	// ParameterInlayHintKind an inlay hint that is for a parameter.
	ParameterInlayHintKind InlayHintKind = 2
)

// MessageType the message type.
type MessageType uint32

const (
	// ErrorMessageType an error message.
	ErrorMessageType MessageType = 1

	// WarningMessageType a warning message.
	WarningMessageType MessageType = 2

	// InfoMessageType an information message.
	InfoMessageType MessageType = 3

	// LogMessageType a log message.
	LogMessageType MessageType = 4

	// DebugMessageType a debug message.
	//
	//  3.18.0
	//
	// Proposed in:.
	//
	// @since 3.18.0 proposed
	DebugMessageType MessageType = 5
)

// TextDocumentSyncKind defines how the host (editor) should sync document changes to the language server.
type TextDocumentSyncKind uint32

const (
	// NoneTextDocumentSyncKind documents should not be synced at all.
	NoneTextDocumentSyncKind TextDocumentSyncKind = 0

	// FullTextDocumentSyncKind documents are synced by always sending the full content of the document.
	FullTextDocumentSyncKind TextDocumentSyncKind = 1

	// IncrementalTextDocumentSyncKind documents are synced by sending the full content on open. After that only incremental updates to the document are send.
	IncrementalTextDocumentSyncKind TextDocumentSyncKind = 2
)

// TextDocumentSaveReason represents reasons why a text document is saved.
type TextDocumentSaveReason uint32

const (
	// ManualTextDocumentSaveReason manually triggered, e.g. by the user pressing save, by starting debugging, or by an API call.
	ManualTextDocumentSaveReason TextDocumentSaveReason = 1

	// AfterDelayTextDocumentSaveReason automatic after a delay.
	AfterDelayTextDocumentSaveReason TextDocumentSaveReason = 2

	// FocusOutTextDocumentSaveReason when the editor lost focus.
	FocusOutTextDocumentSaveReason TextDocumentSaveReason = 3
)

// CompletionItemKind the kind of a completion entry.
type CompletionItemKind uint32

const (
	TextCompletionItemKind CompletionItemKind = 1

	MethodCompletionItemKind CompletionItemKind = 2

	FunctionCompletionItemKind CompletionItemKind = 3

	ConstructorCompletionItemKind CompletionItemKind = 4

	FieldCompletionItemKind CompletionItemKind = 5

	VariableCompletionItemKind CompletionItemKind = 6

	ClassCompletionItemKind CompletionItemKind = 7

	InterfaceCompletionItemKind CompletionItemKind = 8

	ModuleCompletionItemKind CompletionItemKind = 9

	PropertyCompletionItemKind CompletionItemKind = 10

	UnitCompletionItemKind CompletionItemKind = 11

	ValueCompletionItemKind CompletionItemKind = 12

	EnumCompletionItemKind CompletionItemKind = 13

	KeywordCompletionItemKind CompletionItemKind = 14

	SnippetCompletionItemKind CompletionItemKind = 15

	ColorCompletionItemKind CompletionItemKind = 16

	FileCompletionItemKind CompletionItemKind = 17

	ReferenceCompletionItemKind CompletionItemKind = 18

	FolderCompletionItemKind CompletionItemKind = 19

	EnumMemberCompletionItemKind CompletionItemKind = 20

	ConstantCompletionItemKind CompletionItemKind = 21

	StructCompletionItemKind CompletionItemKind = 22

	EventCompletionItemKind CompletionItemKind = 23

	OperatorCompletionItemKind CompletionItemKind = 24

	TypeParameterCompletionItemKind CompletionItemKind = 25
)

// CompletionItemTag completion item tags are extra annotations that tweak the rendering of a completion item.
//
// @since 3.15.0
type CompletionItemTag uint32

const (
	// DeprecatedCompletionItemTag render a completion as obsolete, usually using a strike-out.
	DeprecatedCompletionItemTag CompletionItemTag = 1
)

// InsertTextFormat defines whether the insert text in a completion item should be interpreted as plain text or a snippet.
type InsertTextFormat uint32

const (
	// PlainTextInsertTextFormat the primary text to be inserted is treated as a plain string.
	PlainTextInsertTextFormat InsertTextFormat = 1

	// SnippetInsertTextFormat the primary text to be inserted is treated as a snippet. A snippet can define tab stops and placeholders with `$1`, `$2` and `${3:foo}`. `$0` defines the final tab stop, it defaults to the end of the snippet. Placeholders with equal identifiers are linked, that is typing in one will update others too. See also: https://microsoft.github.io/language-server-protocol/specifications/specification-current/#snippet_syntax.
	SnippetInsertTextFormat InsertTextFormat = 2
)

// InsertTextMode how whitespace and indentation is handled during completion item insertion.
//
// @since 3.16.0
type InsertTextMode uint32

const (
	// AsIsInsertTextMode the insertion or replace strings is taken as it is. If the value is multi line the lines below the cursor will be inserted using the indentation defined in the string value. The client will not apply any kind of adjustments to the string.
	AsIsInsertTextMode InsertTextMode = 1

	// AdjustIndentationInsertTextMode the editor adjusts leading whitespace of new lines so that they match the indentation up to the cursor of the line for which the item is accepted. Consider a line like this: <2tabs><cursor><3tabs>foo. Accepting a multi line completion item is indented using 2 tabs and all following lines inserted will be indented using 2 tabs as well.
	AdjustIndentationInsertTextMode InsertTextMode = 2
)

// DocumentHighlightKind a document highlight kind.
type DocumentHighlightKind uint32

const (
	// TextDocumentHighlightKind a textual occurrence.
	TextDocumentHighlightKind DocumentHighlightKind = 1

	// ReadDocumentHighlightKind read-access of a symbol, like reading a variable.
	ReadDocumentHighlightKind DocumentHighlightKind = 2

	// WriteDocumentHighlightKind write-access of a symbol, like writing to a variable.
	WriteDocumentHighlightKind DocumentHighlightKind = 3
)

// CodeActionKind a set of predefined code action kinds.
type CodeActionKind string

const (
	// EmptyCodeActionKind empty kind.
	EmptyCodeActionKind CodeActionKind = ""

	// QuickFixCodeActionKind base kind for quickfix actions: 'quickfix'.
	QuickFixCodeActionKind CodeActionKind = "quickfix"

	// RefactorCodeActionKind base kind for refactoring actions: 'refactor'.
	RefactorCodeActionKind CodeActionKind = "refactor"

	// RefactorExtractCodeActionKind base kind for refactoring extraction actions: 'refactor.extract' Example extract actions: - Extract method - Extract function - Extract variable - Extract interface from class - .
	RefactorExtractCodeActionKind CodeActionKind = "refactor.extract"

	// RefactorInlineCodeActionKind base kind for refactoring inline actions: 'refactor.inline' Example inline actions: - Inline function - Inline variable - Inline constant - .
	RefactorInlineCodeActionKind CodeActionKind = "refactor.inline"

	// RefactorMoveCodeActionKind base kind for refactoring move actions: `refactor.move` Example move actions: - Move a function to a new file - Move a property between classes - Move method to base class - ...
	//
	//  3.18.0
	//
	// Proposed in:.
	//
	// @since 3.18.0 proposed
	RefactorMoveCodeActionKind CodeActionKind = "refactor.move"

	// RefactorRewriteCodeActionKind base kind for refactoring rewrite actions: 'refactor.rewrite' Example rewrite actions: - Convert JavaScript function to class - Add or remove parameter - Encapsulate field - Make method static - Move method to base class - .
	RefactorRewriteCodeActionKind CodeActionKind = "refactor.rewrite"

	// SourceCodeActionKind base kind for source actions: `source` Source code actions apply to the entire file.
	SourceCodeActionKind CodeActionKind = "source"

	// SourceOrganizeImportsCodeActionKind base kind for an organize imports source action: `source.organizeImports`.
	SourceOrganizeImportsCodeActionKind CodeActionKind = "source.organizeImports"

	// SourceFixAllCodeActionKind base kind for auto-fix source actions: `source.fixAll`. Fix all actions automatically fix errors that have a clear fix that do not require user input. They should not suppress errors or perform unsafe fixes such as generating new types or classes.
	//
	//
	//
	// @since 3.15.0
	SourceFixAllCodeActionKind CodeActionKind = "source.fixAll"

	// NotebookCodeActionKind base kind for all code actions applying to the entire notebook's scope. CodeActionKinds using this should always begin with `notebook.`
	//
	//
	//
	// @since 3.18.0
	NotebookCodeActionKind CodeActionKind = "notebook"
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

// MarkupKind describes the content type that a client supports in various result literals like `Hover`, `ParameterInfo` or `CompletionItem`. Please note that `MarkupKinds` must not start with a `$`. This kinds are reserved for internal usage.
type MarkupKind string

const (
	// PlainTextMarkupKind plain text is supported as a content format.
	PlainTextMarkupKind MarkupKind = "plaintext"

	// MarkdownMarkupKind markdown is supported as a content format.
	MarkdownMarkupKind MarkupKind = "markdown"
)

// LanguageKind predefined Language kinds
//
//	3.18.0
//
// Proposed in:.
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

// InlineCompletionTriggerKind describes how an InlineCompletionItemProvider inline completion provider was triggered.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type InlineCompletionTriggerKind uint32

const (
	// InvokedInlineCompletionTriggerKind completion was triggered explicitly by a user gesture.
	InvokedInlineCompletionTriggerKind InlineCompletionTriggerKind = 1

	// AutomaticInlineCompletionTriggerKind completion was triggered automatically while editing.
	AutomaticInlineCompletionTriggerKind InlineCompletionTriggerKind = 2
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

// CompletionTriggerKind how a completion was triggered.
type CompletionTriggerKind uint32

const (
	// InvokedCompletionTriggerKind completion was triggered by typing an identifier (24x7 code complete), manual invocation (e.g Ctrl+Space) or via API.
	InvokedCompletionTriggerKind CompletionTriggerKind = 1

	// TriggerCharacterCompletionTriggerKind completion was triggered by a trigger character specified by the `triggerCharacters` properties of the `CompletionRegistrationOptions`.
	TriggerCharacterCompletionTriggerKind CompletionTriggerKind = 2

	// TriggerForIncompleteCompletionsCompletionTriggerKind completion was re-triggered as current completion list is incomplete.
	TriggerForIncompleteCompletionsCompletionTriggerKind CompletionTriggerKind = 3
)

// SignatureHelpTriggerKind how a signature help was triggered.
//
// @since 3.15.0
type SignatureHelpTriggerKind uint32

const (
	// InvokedSignatureHelpTriggerKind signature help was invoked manually by the user or by a command.
	InvokedSignatureHelpTriggerKind SignatureHelpTriggerKind = 1

	// TriggerCharacterSignatureHelpTriggerKind signature help was triggered by a trigger character.
	TriggerCharacterSignatureHelpTriggerKind SignatureHelpTriggerKind = 2

	// ContentChangeSignatureHelpTriggerKind signature help was triggered by the cursor moving or by the document content changing.
	ContentChangeSignatureHelpTriggerKind SignatureHelpTriggerKind = 3
)

// CodeActionTriggerKind the reason why code actions were requested.
//
// @since 3.17.0
type CodeActionTriggerKind uint32

const (
	// InvokedCodeActionTriggerKind code actions were explicitly requested by the user or by an extension.
	InvokedCodeActionTriggerKind CodeActionTriggerKind = 1

	// AutomaticCodeActionTriggerKind code actions were requested automatically. This typically happens when current selection in a file changes, but can also be triggered when file content changes.
	AutomaticCodeActionTriggerKind CodeActionTriggerKind = 2
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

// NotebookCellKind a notebook cell kind.
//
// @since 3.17.0
type NotebookCellKind uint32

const (
	// MarkupNotebookCellKind a markup-cell is formatted source that is used for display.
	MarkupNotebookCellKind NotebookCellKind = 1

	// CodeNotebookCellKind a code-cell is source code.
	CodeNotebookCellKind NotebookCellKind = 2
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
	// AbortFailureHandlingKind applying the workspace change is simply aborted if one of the changes provided fails. All operations executed before the failing operation stay executed.
	AbortFailureHandlingKind FailureHandlingKind = "abort"

	// TransactionalFailureHandlingKind all operations are executed transactional. That means they either all succeed or no changes at all are applied to the workspace.
	TransactionalFailureHandlingKind FailureHandlingKind = "transactional"

	// TextOnlyTransactionalFailureHandlingKind if the workspace edit contains only textual file changes they are executed transactional. If resource changes (create, rename or delete file) are part of the change the failure handling strategy is abort.
	TextOnlyTransactionalFailureHandlingKind FailureHandlingKind = "textOnlyTransactional"

	// UndoFailureHandlingKind the client tries to undo the operations already executed. But there is no guarantee that this is succeeding.
	UndoFailureHandlingKind FailureHandlingKind = "undo"
)

type PrepareSupportDefaultBehavior uint32

const (
	// IdentifierPrepareSupportDefaultBehavior the client's default behavior is to select the identifier according the to language's syntax rule.
	IdentifierPrepareSupportDefaultBehavior PrepareSupportDefaultBehavior = 1
)

type TokenFormat string

const (
	RelativeTokenFormat TokenFormat = "relative"
)
