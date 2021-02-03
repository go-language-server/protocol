// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strconv"

	"go.lsp.dev/uri"
)

// DocumentURI represents the URI of a document.
//
// Many of the interfaces contain fields that correspond to the URI of a document.
// For clarity, the type of such a field is declared as a DocumentURI.
// Over the wire, it will still be transferred as a string, but this guarantees
// that the contents of that string can be parsed as a valid URI.
type DocumentURI string

// EOL denotes represents the character offset.
var EOL = []string{"\n", "\r\n", "\r"}

// Position represents a text document expressed as zero-based line and zero-based character offset.
//
// A position is between two characters like an "insert" cursor in a editor.
type Position struct {
	// Line position in a document (zero-based).
	Line float64 `json:"line"`

	// Character offset on a line in a document (zero-based). Assuming that the line is
	// represented as a string, the `character` value represents the gap between the
	// `character` and `character + 1`.
	// If the character value is greater than the line length it defaults back to the
	// line length.
	Character float64 `json:"character"`
}

// Range represents a text document expressed as (zero-based) start and end positions.
//
// A range is comparable to a selection in an editor. Therefore the end position is exclusive.
// If you want to specify a range that contains a line including the line ending character(s) then use an end position denoting the start of the next line.
type Range struct {
	// Start is the range's start position.
	Start Position `json:"start"`

	// End is the range's end position.
	End Position `json:"end"`
}

// Location represents a location inside a resource, such as a line inside a text file.
type Location struct {
	URI   uri.URI `json:"uri"`
	Range Range   `json:"range"`
}

// LocationLink represents a link between a source and a target location.
type LocationLink struct {
	// OriginSelectionRange span of the origin of this link.
	// Used as the underlined span for mouse interaction. Defaults to the word range at
	// the mouse position.
	OriginSelectionRange *Range `json:"originSelectionRange,omitempty"`

	// TargetURI is the target resource identifier of this link.
	TargetURI uri.URI `json:"targetUri"`

	// TargetRange is the full target range of this link. If the target for example is a symbol then target range is the
	// range enclosing this symbol not including leading/trailing whitespace but everything else
	// like comments. This information is typically used to highlight the range in the editor.
	TargetRange Range `json:"targetRange"`

	// TargetSelectionRange is the range that should be selected and revealed when this link is being followed, e.g the name of a function.
	// Must be contained by the the `targetRange`. See also `DocumentSymbol#range`
	TargetSelectionRange Range `json:"targetSelectionRange"`
}

// Diagnostic represents a diagnostic, such as a compiler error or warning.
//
// Diagnostic objects are only valid in the scope of a resource.
type Diagnostic struct {
	// Range is the range at which the message applies.
	Range Range `json:"range"`

	// Severity is the diagnostic's severity. Can be omitted. If omitted it is up to the
	// client to interpret diagnostics as error, warning, info or hint.
	Severity DiagnosticSeverity `json:"severity,omitempty"`

	// Code is the diagnostic's code, which might appear in the user interface.
	Code interface{} `json:"code,omitempty"`

	// Source a human-readable string describing the source of this
	// diagnostic, e.g. 'typescript' or 'super lint'.
	Source string `json:"source,omitempty"`

	// Message is the diagnostic's message.
	Message string `json:"message"`

	// RelatedInformation an array of related diagnostic information, e.g. when symbol-names within
	// a scope collide all definitions can be marked via this property.
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`
}

// DiagnosticSeverity indicates the severity of a Diagnostic message.
type DiagnosticSeverity float64

const (
	// SeverityError reports an error.
	SeverityError DiagnosticSeverity = 1

	// SeverityWarning reports a warning.
	SeverityWarning DiagnosticSeverity = 2

	// SeverityInformation reports an information.
	SeverityInformation DiagnosticSeverity = 3

	// SeverityHint reports a hint.
	SeverityHint DiagnosticSeverity = 4
)

// String implements fmt.Stringer.
func (d DiagnosticSeverity) String() string {
	switch d {
	case SeverityError:
		return "Error"
	case SeverityWarning:
		return "Warning"
	case SeverityInformation:
		return "Information"
	case SeverityHint:
		return "Hint"
	default:
		return strconv.FormatFloat(float64(d), 'f', -10, 64)
	}
}

// DiagnosticRelatedInformation represents a related message and source code location for a diagnostic.
//
// This should be used to point to code locations that cause or related to a diagnostics, e.g when duplicating
// a symbol in a scope.
type DiagnosticRelatedInformation struct {
	// Location is the location of this related diagnostic information.
	Location Location `json:"location"`

	// Message is the message of this related diagnostic information.
	Message string `json:"message"`
}

// Command represents a reference to a command. Provides a title which will be used to represent a command in the UI.
//
// Commands are identified by a string identifier.
// The recommended way to handle commands is to implement their execution on the server side if the client and server provides the corresponding capabilities.
// Alternatively the tool extension code could handle the command. The protocol currently doesn't specify a set of well-known commands.
type Command struct {
	// Title of the command, like `save`.
	Title string `json:"title"`

	// Command is the identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments that the command handler should be invoked with.
	Arguments []interface{} `json:"arguments,omitempty"`
}

// TextEdit is a textual edit applicable to a text document.
type TextEdit struct {
	// Range is the range of the text document to be manipulated.
	// To insert text into a document create a range where start === end.
	Range Range `json:"range"`

	// NewText is the string to be inserted. For delete operations use an
	// empty string.
	NewText string `json:"newText"`
}

// TextDocumentEdit describes textual changes on a single text document.
//
// The text document is referred to as a VersionedTextDocumentIdentifier to allow clients to check the text document version before an edit is applied.
// A TextDocumentEdit describes all changes on a version Si and after they are applied move the document to version Si+1.
// So the creator of a TextDocumentEdit doesn't need to sort the array or do any kind of ordering. However the edits must be non overlapping.
type TextDocumentEdit struct {
	// TextDocument is the text document to change.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// Edits is the edits to be applied.
	Edits []TextEdit `json:"edits"`
}

// ResourceOperationKind is the file event type.
type ResourceOperationKind string

const (
	// CreateResourceOperation supports creating new files and folders.
	CreateResourceOperation ResourceOperationKind = "create"

	// RenameResourceOperation supports renaming existing files and folders.
	RenameResourceOperation ResourceOperationKind = "rename"

	// DeleteResourceOperation supports deleting existing files and folders.
	DeleteResourceOperation ResourceOperationKind = "delete"
)

// CreateFileOptions represents an options to create a file.
type CreateFileOptions struct {
	// Overwrite existing file. Overwrite wins over `ignoreIfExists`.
	Overwrite bool `json:"overwrite,omitempty"`

	// IgnoreIfExists ignore if exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

// CreateFile represents a create file operation.
type CreateFile struct {
	// Kind a create.
	Kind ResourceOperationKind `json:"kind"` // should be `create`

	// URI is the resource to create.
	URI uri.URI `json:"uri"`

	// Options additional options.
	Options *CreateFileOptions `json:"options,omitempty"`
}

// RenameFileOptions represents a rename file options.
type RenameFileOptions struct {
	// Overwrite target if existing. Overwrite wins over `ignoreIfExists`.
	Overwrite bool `json:"overwrite,omitempty"`

	// IgnoreIfExists ignores if target exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

// RenameFile represents a rename file operation.
type RenameFile struct {
	// Kind a rename.
	Kind ResourceOperationKind `json:"kind"` // should be `rename`

	// OldURI is the old (existing) location.
	OldURI uri.URI `json:"oldUri"`

	// NewURI is the new location.
	NewURI uri.URI `json:"newUri"`

	// Options rename options.
	Options *RenameFileOptions `json:"options,omitempty"`
}

// DeleteFileOptions represents a delete file options.
type DeleteFileOptions struct {
	// Recursive delete the content recursively if a folder is denoted.
	Recursive bool `json:"recursive,omitempty"`

	// IgnoreIfNotExists ignore the operation if the file doesn't exist.
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
}

// DeleteFile represents a delete file operation.
type DeleteFile struct {
	// Kind is a delete.
	Kind ResourceOperationKind `json:"kind"` // should be `delete`

	// URI is the file to delete.
	URI uri.URI `json:"uri"`

	// Options delete options.
	Options *DeleteFileOptions `json:"options,omitempty"`
}

// WorkspaceEdit represent a changes to many resources managed in the workspace.
//
// The edit should either provide changes or documentChanges.
// If the client can handle versioned document edits and if documentChanges are present, the latter are preferred over changes.
type WorkspaceEdit struct {
	// Changes holds changes to existing resources.
	Changes map[uri.URI][]TextEdit `json:"changes,omitempty"`

	// DocumentChanges depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes
	// are either an array of `TextDocumentEdit`s to express changes to n different text documents
	// where each text document edit addresses a specific version of a text document. Or it can contain
	// above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations.
	//
	// Whether a client supports versioned document edits is expressed via
	// `workspace.workspaceEdit.documentChanges` client capability.
	//
	// If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then
	// only plain `TextEdit`s using the `changes` property are supported.
	DocumentChanges []TextDocumentEdit `json:"documentChanges,omitempty"`
}

// TextDocumentIdentifier indicates the using a URI. On the protocol level, URIs are passed as strings.
type TextDocumentIdentifier struct {
	// URI is the text document's URI.
	URI uri.URI `json:"uri"`
}

// TextDocumentItem represent an item to transfer a text document from the client to the server.
type TextDocumentItem struct {
	// URI is the text document's URI.
	URI uri.URI `json:"uri"`

	// LanguageID is the text document's language identifier.
	LanguageID LanguageIdentifier `json:"languageId"`

	// Version is the version number of this document (it will increase after each
	// change, including undo/redo).
	Version float64 `json:"version"`

	// Text is the content of the opened text document.
	Text string `json:"text"`
}

// LanguageIdentifier represent a text document's language identifier.
type LanguageIdentifier string

const (
	// BatLanguage Windows Bat Language.
	BatLanguage LanguageIdentifier = "bat"

	// BibtexLanguage BibTeX Language.
	BibtexLanguage LanguageIdentifier = "bibtex"

	// ClojureLanguage Clojure Language.
	ClojureLanguage LanguageIdentifier = "clojure"

	// CoffeescriptLanguage Coffeescript Language.
	CoffeescriptLanguage LanguageIdentifier = "coffeescript"

	// CLanguage C Language.
	CLanguage LanguageIdentifier = "c"

	// CppLanguage C++ Language.
	CppLanguage LanguageIdentifier = "cpp"

	// CsharpLanguage C# Language.
	CsharpLanguage LanguageIdentifier = "csharp"

	// CSSLanguage CSS Language.
	CSSLanguage LanguageIdentifier = "css"

	// DiffLanguage Diff Language.
	DiffLanguage LanguageIdentifier = "diff"

	// DartLanguage Dart Language.
	DartLanguage LanguageIdentifier = "dart"

	// DockerfileLanguage Dockerfile Language.
	DockerfileLanguage LanguageIdentifier = "dockerfile"

	// FsharpLanguage F# Language.
	FsharpLanguage LanguageIdentifier = "fsharp"

	// GitCommitLanguage Git Language.
	GitCommitLanguage LanguageIdentifier = "git-commit"

	// GitRebaseLanguage Git Language.
	GitRebaseLanguage LanguageIdentifier = "git-rebase"

	// GoLanguage Go Language.
	GoLanguage LanguageIdentifier = "go"

	// GroovyLanguage Groovy Language.
	GroovyLanguage LanguageIdentifier = "groovy"

	// HandlebarsLanguage Handlebars Language.
	HandlebarsLanguage LanguageIdentifier = "handlebars"

	// HTMLLanguage HTML Language.
	HTMLLanguage LanguageIdentifier = "html"

	// IniLanguage Ini Language.
	IniLanguage LanguageIdentifier = "ini"

	// JavaLanguage Java Language.
	JavaLanguage LanguageIdentifier = "java"

	// JavaScriptLanguage JavaScript Language.
	JavaScriptLanguage LanguageIdentifier = "javascript"

	// JSONLanguage JSON Language.
	JSONLanguage LanguageIdentifier = "json"

	// LatexLanguage LaTeX Language.
	LatexLanguage LanguageIdentifier = "latex"

	// LessLanguage Less Language.
	LessLanguage LanguageIdentifier = "less"

	// LuaLanguage Lua Language.
	LuaLanguage LanguageIdentifier = "lua"

	// MakefileLanguage Makefile Language.
	MakefileLanguage LanguageIdentifier = "makefile"

	// MarkdownLanguage Markdown Language.
	MarkdownLanguage LanguageIdentifier = "markdown"

	// ObjectiveCLanguage Objective-C Language.
	ObjectiveCLanguage LanguageIdentifier = "objective-c"

	// ObjectiveCppLanguage Objective-C++ Language.
	ObjectiveCppLanguage LanguageIdentifier = "objective-cpp"

	// PerlLanguage Perl Language.
	PerlLanguage LanguageIdentifier = "perl"

	// Perl6Language Perl Language.
	Perl6Language LanguageIdentifier = "perl6"

	// PHPLanguage PHP Language.
	PHPLanguage LanguageIdentifier = "php"

	// PowershellLanguage Powershell Language.
	PowershellLanguage LanguageIdentifier = "powershell"

	// JadeLanguage Pug Language.
	JadeLanguage LanguageIdentifier = "jade"

	// PythonLanguage Python Language.
	PythonLanguage LanguageIdentifier = "python"

	// RLanguage R Language.
	RLanguage LanguageIdentifier = "r"

	// RazorLanguage Razor(cshtml) Language.
	RazorLanguage LanguageIdentifier = "razor"

	// RubyLanguage Ruby Language.
	RubyLanguage LanguageIdentifier = "ruby"

	// RustLanguage Rust Language.
	RustLanguage LanguageIdentifier = "rust"

	// ScssLanguage Sass Language.
	ScssLanguage LanguageIdentifier = "scss"

	// SassLanguage Sass Language.
	SassLanguage LanguageIdentifier = "sass"

	// ScalaLanguage Scala Language.
	ScalaLanguage LanguageIdentifier = "scala"

	// ShaderlabLanguage ShaderLab Language.
	ShaderlabLanguage LanguageIdentifier = "shaderlab"

	// ShellscriptLanguage Shell Script (Bash) Language.
	ShellscriptLanguage LanguageIdentifier = "shellscript"

	// SQLLanguage SQL Language.
	SQLLanguage LanguageIdentifier = "sql"

	// SwiftLanguage Swift Language.
	SwiftLanguage LanguageIdentifier = "swift"

	// TypeScriptLanguage TypeScript Language.
	TypeScriptLanguage LanguageIdentifier = "typescript"

	// TexLanguage TeX Language.
	TexLanguage LanguageIdentifier = "tex"

	// VBLanguage Visual Basic Language.
	VBLanguage LanguageIdentifier = "vb"

	// XMLLanguage XML Language.
	XMLLanguage LanguageIdentifier = "xml"

	// XslLanguage XSL Language.
	XslLanguage LanguageIdentifier = "xsl"

	// YamlLanguage YAML Language.
	YamlLanguage LanguageIdentifier = "yaml"
)

// languageIdentifierMap map of LanguageIdentifiers.
var languageIdentifierMap = map[string]LanguageIdentifier{
	"bat":           BatLanguage,
	"bibtex":        BibtexLanguage,
	"clojure":       ClojureLanguage,
	"coffeescript":  CoffeescriptLanguage,
	"c":             CLanguage,
	"cpp":           CppLanguage,
	"csharp":        CsharpLanguage,
	"css":           CSSLanguage,
	"diff":          DiffLanguage,
	"dart":          DartLanguage,
	"dockerfile":    DockerfileLanguage,
	"fsharp":        FsharpLanguage,
	"git-commit":    GitCommitLanguage,
	"git-rebase":    GitRebaseLanguage,
	"go":            GoLanguage,
	"groovy":        GroovyLanguage,
	"handlebars":    HandlebarsLanguage,
	"html":          HTMLLanguage,
	"ini":           IniLanguage,
	"java":          JavaLanguage,
	"javascript":    JavaScriptLanguage,
	"json":          JSONLanguage,
	"latex":         LatexLanguage,
	"less":          LessLanguage,
	"lua":           LuaLanguage,
	"makefile":      MakefileLanguage,
	"markdown":      MarkdownLanguage,
	"objective-c":   ObjectiveCLanguage,
	"objective-cpp": ObjectiveCppLanguage,
	"perl":          PerlLanguage,
	"perl6":         Perl6Language,
	"php":           PHPLanguage,
	"powershell":    PowershellLanguage,
	"jade":          JadeLanguage,
	"python":        PythonLanguage,
	"r":             RLanguage,
	"razor":         RazorLanguage,
	"ruby":          RubyLanguage,
	"rust":          RustLanguage,
	"scss":          ScssLanguage,
	"sass":          SassLanguage,
	"scala":         ScalaLanguage,
	"shaderlab":     ShaderlabLanguage,
	"shellscript":   ShellscriptLanguage,
	"sql":           SQLLanguage,
	"swift":         SwiftLanguage,
	"typescript":    TypeScriptLanguage,
	"tex":           TexLanguage,
	"vb":            VBLanguage,
	"xml":           XMLLanguage,
	"xsl":           XslLanguage,
	"yaml":          YamlLanguage,
}

// ToLanguageIdentifier converts ft to LanguageIdentifier.
func ToLanguageIdentifier(ft string) LanguageIdentifier {
	langID, ok := languageIdentifierMap[ft]
	if !ok {
		return LanguageIdentifier(ft)
	}

	return langID
}

// VersionedTextDocumentIdentifier represents an identifier to denote a specific version of a text document.
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	// Version is the version number of this document.
	//
	// If a versioned text document identifier is sent from the server to the client and the file is not open in the editor
	// (the server has not received an open notification before) the server can send
	// `null` to indicate that the version is known and the content on disk is the
	// truth (as speced with document content ownership).
	//
	// The version number of a document will increase after each change, including
	// undo/redo. The number doesn't need to be consecutive.
	Version *uint64 `json:"version"`
}

// TextDocumentPositionParams is a parameter literal used in requests to pass a text document and a position inside that document.
type TextDocumentPositionParams struct {
	// TextDocument is the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position is the position inside the text document.
	Position Position `json:"position"`
}

// DocumentFilter is a document filter denotes a document through properties like language, scheme or pattern.
//
// An example is a filter that applies to TypeScript files on disk.
type DocumentFilter struct {
	// Language a language id, like `typescript`.
	Language string `json:"language,omitempty"`

	// Scheme a URI scheme, like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`

	// Pattern a glob pattern, like `*.{ts,js}`.
	//
	// Glob patterns can have the following syntax:
	// - `*` to match one or more characters in a path segment
	// - `?` to match on one character in a path segment
	// - `**` to match any number of path segments, including none
	// - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
	// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
	Pattern string `json:"pattern,omitempty"`
}

// DocumentSelector is a document selector is the combination of one or more document filters.
type DocumentSelector []*DocumentFilter

// MarkupKind describes the content type that a client supports in various
// result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
//
// Please note that `MarkupKinds` must not start with a `$`. This kinds
// are reserved for internal usage.
type MarkupKind string

const (
	// PlainText is supported as a content format.
	PlainText MarkupKind = "plaintext"

	// Markdown is supported as a content format.
	Markdown MarkupKind = "markdown"
)

// MarkupContent a `MarkupContent` literal represents a string value which content is interpreted base on its
// kind flag.
//
// Currently the protocol supports `plaintext` and `markdown` as markup kinds.
//
// If the kind is `markdown` then the value can contain fenced code blocks like in GitHub issues.
// See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting
//
// Here is an example how such a string can be constructed using JavaScript / TypeScript:
// ```typescript
//  * let markdown: MarkdownContent = {
//  *  kind: MarkupKind.Markdown,
//  *	value: [
//  *		'# Header',
//  *		'Some text',
//  *		'```typescript',
// 		'someCode();',
// 		'```'
//  *	].join('\n')
//  * };
//  * ```
//
// NOTE: clients might sanitize the return markdown. A client could decide to
// remove HTML from the markdown to avoid script execution.
type MarkupContent struct {
	// Kind is the type of the Markup
	Kind MarkupKind `json:"kind"`

	// Value is the content itself
	Value string `json:"value"`
}
