// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"go.lsp.dev/uri"
)

// SemanticTokenTypes a set of predefined token types. This set is not fixed an clients can specify additional token types
// via the corresponding client capabilities.
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

	// LabelSemanticTokenTypes.
	//
	// @since 3.18.0
	LabelSemanticTokenTypes SemanticTokenTypes = "label"
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

	// AdjustIndentationInsertTextMode the editor adjusts leading whitespace of new lines so that they match the indentation up to the cursor of the line for which the item is accepted. Consider a line like this: <2tabs><cursor><3tabs>foo.
	// Accepting a multi line completion item is indented using 2 tabs and all following lines inserted will be indented using 2 tabs as well.
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

	// RefactorMoveCodeActionKind base kind for refactoring move actions: `refactor.move` Example move actions: - Move a function to a
	// new file - Move a property between classes - Move method to base class - ... 3.18.0 @proposed.
	//
	// @since 3.18.0 proposed
	RefactorMoveCodeActionKind CodeActionKind = "refactor.move"

	// RefactorRewriteCodeActionKind base kind for refactoring rewrite actions: 'refactor.rewrite' Example rewrite actions: - Convert JavaScript function to class - Add or remove parameter - Encapsulate field - Make method static - Move method to base class - .
	RefactorRewriteCodeActionKind CodeActionKind = "refactor.rewrite"

	// SourceCodeActionKind base kind for source actions: `source` Source code actions apply to the entire file.
	SourceCodeActionKind CodeActionKind = "source"

	// SourceOrganizeImportsCodeActionKind base kind for an organize imports source action: `source.organizeImports`.
	SourceOrganizeImportsCodeActionKind CodeActionKind = "source.organizeImports"

	// SourceFixAllCodeActionKind base kind for auto-fix source actions: `source.fixAll`. Fix all actions automatically fix errors that have a clear fix that do not require user input. They should not suppress errors or perform unsafe
	// fixes such as generating new types or classes.
	//
	// @since 3.15.0
	SourceFixAllCodeActionKind CodeActionKind = "source.fixAll"

	// NotebookCodeActionKind base kind for all code actions applying to the entire notebook's scope. CodeActionKinds using this should always begin with `notebook.`
	//
	// @since 3.18.0
	NotebookCodeActionKind CodeActionKind = "notebook"
)

// CodeActionTag code action tags are extra annotations that tweak the behavior of a code action.  3.18.0 - proposed.
//
// @since 3.18.0 - proposed
type CodeActionTag uint32

const (
	// LlmgeneratedCodeActionTag marks the code action as LLM-generated.
	LlmgeneratedCodeActionTag CodeActionTag = 1
)

// InlineCompletionTriggerKind describes how an InlineCompletionItemProvider inline completion provider was triggered. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionTriggerKind uint32

const (
	// InvokedInlineCompletionTriggerKind completion was triggered explicitly by a user gesture.
	InvokedInlineCompletionTriggerKind InlineCompletionTriggerKind = 1

	// AutomaticInlineCompletionTriggerKind completion was triggered automatically while editing.
	AutomaticInlineCompletionTriggerKind InlineCompletionTriggerKind = 2
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

// ApplyKind defines how values from a set of defaults and an individual item will be merged.
//
// @since 3.18.0
type ApplyKind string

const (
	// ReplaceApplyKind the value from the individual item (if provided and not `null`) will be used instead of the default.
	ReplaceApplyKind ApplyKind = "replace"

	// MergeApplyKind the value from the item will be merged with the default. The specific rules for mergeing values are defined against each field that supports merging.
	MergeApplyKind ApplyKind = "merge"
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

type PrepareSupportDefaultBehavior uint32

const (
	// IdentifierPrepareSupportDefaultBehavior the client's default behavior is to select the identifier according the to language's syntax rule.
	IdentifierPrepareSupportDefaultBehavior PrepareSupportDefaultBehavior = 1
)

type TokenFormat string

const (
	RelativeTokenFormat TokenFormat = "relative"
)

type ImplementationParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
	PartialResultParams
}

type ImplementationOptions struct {
	// mixins
	WorkDoneProgressOptions
}

type ImplementationRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	ImplementationOptions
	// mixins
	StaticRegistrationOptions
}

type TypeDefinitionParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
	PartialResultParams
}

type TypeDefinitionOptions struct {
	// mixins
	WorkDoneProgressOptions
}

type TypeDefinitionRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	TypeDefinitionOptions
	// mixins
	StaticRegistrationOptions
}

// DocumentColorParams parameters for a DocumentColorRequest.
type DocumentColorParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// Color represents a color in RGBA space.
type Color struct {
	// Red the red component of this color in the range [0-1].
	Red float64 `json:"red"`

	// Green the green component of this color in the range [0-1].
	Green float64 `json:"green"`

	// Blue the blue component of this color in the range [0-1].
	Blue float64 `json:"blue"`

	// Alpha the alpha component of this color in the range [0-1].
	Alpha float64 `json:"alpha"`
}

// ColorInformation represents a color range from a document.
type ColorInformation struct {
	// Range the range in the document where this color appears.
	Range Range `json:"range"`

	// Color the actual color value for this color range.
	Color Color `json:"color"`
}

type DocumentColorOptions struct {
	// mixins
	WorkDoneProgressOptions
}

type DocumentColorRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DocumentColorOptions
	// mixins
	StaticRegistrationOptions
}

// ColorPresentationParams parameters for a ColorPresentationRequest.
type ColorPresentationParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Color the color to request presentations for.
	Color Color `json:"color"`

	// Range the range where the color would be inserted. Serves as a context.
	Range Range `json:"range"`
}

type ColorPresentation struct {
	// Label the label of this color presentation. It will be shown on the color picker header. By default this is also the text that is inserted when selecting this color presentation.
	Label string `json:"label"`

	// TextEdit an TextEdit edit which is applied to a document when selecting this presentation for the color. When
	// `falsy` the ColorPresentation.label label is used.
	TextEdit *TextEdit `json:"textEdit,omitempty"`

	// AdditionalTextEdits an optional array of additional TextEdit text edits that are applied when selecting this color presentation. Edits must not overlap with the main ColorPresentation.textEdit edit nor with themselves.
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
}

// FoldingRangeParams parameters for a FoldingRangeRequest.
type FoldingRangeParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// FoldingRange represents a folding range. To be valid, start and end line must be bigger than zero and smaller than the number of lines in the document. Clients are free to ignore invalid ranges.
type FoldingRange struct {
	// StartLine the zero-based start line of the range to fold. The folded area starts after the line's last character. To be valid, the end must be zero or larger and smaller than the number of lines in the document.
	StartLine uint32 `json:"startLine"`

	// StartCharacter the zero-based character offset from where the folded range starts. If not defined, defaults to the length of the start line.
	StartCharacter uint32 `json:"startCharacter,omitempty"`

	// EndLine the zero-based end line of the range to fold. The folded area ends with the line's last character. To be valid, the end must be zero or larger and smaller than the number of lines in the document.
	EndLine uint32 `json:"endLine"`

	// EndCharacter the zero-based character offset before the folded range ends. If not defined, defaults to the length
	// of the end line.
	EndCharacter uint32 `json:"endCharacter,omitempty"`

	// Kind describes the kind of the folding range such as 'comment' or 'region'. The kind is used to categorize folding ranges and used by commands like 'Fold all comments'. See FoldingRangeKind for an enumeration of standardized kinds.
	Kind FoldingRangeKind `json:"kind,omitempty"`

	// CollapsedText the text that the client should show when the specified range is collapsed. If not defined or not supported by the client, a default will be chosen by the client.
	CollapsedText string `json:"collapsedText,omitempty"`
}

type FoldingRangeOptions struct {
	// mixins
	WorkDoneProgressOptions
}

type FoldingRangeRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	FoldingRangeOptions
	// mixins
	StaticRegistrationOptions
}

type DeclarationParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
	PartialResultParams
}

type DeclarationOptions struct {
	// mixins
	WorkDoneProgressOptions
}

type DeclarationRegistrationOptions struct {
	// extends
	DeclarationOptions
	TextDocumentRegistrationOptions
	// mixins
	StaticRegistrationOptions
}

// SelectionRangeParams a parameter literal used in selection range requests.
type SelectionRangeParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Positions the positions inside the text document.
	Positions []Position `json:"positions"`
}

// SelectionRange a selection range represents a part of a selection hierarchy. A selection range may have a parent selection range that contains it.
type SelectionRange struct {
	// Range the Range range of this selection range.
	Range Range `json:"range"`

	// Parent the parent selection range containing this range. Therefore `parent.range` must contain `this.range`.
	Parent *SelectionRange `json:"parent,omitempty"`
}

type SelectionRangeOptions struct {
	// mixins
	WorkDoneProgressOptions
}

type SelectionRangeRegistrationOptions struct {
	// extends
	SelectionRangeOptions
	TextDocumentRegistrationOptions
	// mixins
	StaticRegistrationOptions
}

// CallHierarchyPrepareParams the parameter of a `textDocument/prepareCallHierarchy` request.
//
// @since 3.16.0
type CallHierarchyPrepareParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
}

// CallHierarchyItem represents programming constructs like functions or constructors in the context of call hierarchy.
//
// @since 3.16.0
type CallHierarchyItem struct {
	// Name the name of this item.
	//
	// @since 3.16.0
	Name string `json:"name"`

	// Kind the kind of this item.
	//
	// @since 3.16.0
	Kind SymbolKind `json:"kind"`

	// Tags tags for this item.
	//
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitempty"`

	// Detail more detail for this item, e.g. the signature of a function.
	//
	// @since 3.16.0
	Detail string `json:"detail,omitempty"`

	// URI the resource identifier of this item.
	//
	// @since 3.16.0
	URI DocumentURI `json:"uri"`

	// Range the range enclosing this symbol not including leading/trailing whitespace but everything else, e.g. comments and code.
	//
	// @since 3.16.0
	Range Range `json:"range"`

	// SelectionRange the range that should be selected and revealed when this symbol is being picked, e.g. the name of a function. Must be contained by the CallHierarchyItem.range `range`.
	//
	// @since 3.16.0
	SelectionRange Range `json:"selectionRange"`

	// Data a data entry field that is preserved between a call hierarchy prepare and incoming calls or outgoing
	// calls requests.
	//
	// @since 3.16.0
	Data any `json:"data,omitempty"`
}

// CallHierarchyOptions call hierarchy options used during static registration.
//
// @since 3.16.0
type CallHierarchyOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// CallHierarchyRegistrationOptions call hierarchy options used during static or dynamic registration.
//
// @since 3.16.0
type CallHierarchyRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	CallHierarchyOptions
	// mixins
	StaticRegistrationOptions
}

// CallHierarchyIncomingCallsParams the parameter of a `callHierarchy/incomingCalls` request.
//
// @since 3.16.0
type CallHierarchyIncomingCallsParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// @since 3.16.0
	Item CallHierarchyItem `json:"item"`
}

// CallHierarchyIncomingCall represents an incoming call, e.g. a caller of a method or constructor.
//
// @since 3.16.0
type CallHierarchyIncomingCall struct {
	// From the item that makes the call.
	//
	// @since 3.16.0
	From CallHierarchyItem `json:"from"`

	// FromRanges the ranges at which the calls appear. This is relative to the caller denoted by CallHierarchyIncomingCall.from `this.from`.
	//
	// @since 3.16.0
	FromRanges []Range `json:"fromRanges"`
}

// CallHierarchyOutgoingCallsParams the parameter of a `callHierarchy/outgoingCalls` request.
//
// @since 3.16.0
type CallHierarchyOutgoingCallsParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// @since 3.16.0
	Item CallHierarchyItem `json:"item"`
}

// CallHierarchyOutgoingCall represents an outgoing call, e.g. calling a getter from a method or a method from a constructor etc.
//
// @since 3.16.0
type CallHierarchyOutgoingCall struct {
	// To the item that is called.
	//
	// @since 3.16.0
	To CallHierarchyItem `json:"to"`

	// FromRanges the range at which this item is called. This is the range relative to the caller, e.g the item passed to CallHierarchyItemProvider.provideCallHierarchyOutgoingCalls `provideCallHierarchyOutgoingCalls`
	// and not CallHierarchyOutgoingCall.to `this.to`.
	//
	// @since 3.16.0
	FromRanges []Range `json:"fromRanges"`
}

// SemanticTokensParams.
//
// @since 3.16.0
type SemanticTokensParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the text document.
	//
	// @since 3.16.0
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// SemanticTokens.
//
// @since 3.16.0
type SemanticTokens struct {
	// ResultID an optional result id. If provided and clients support delta updating the client will include the result id in the next semantic token request. A server can then instead of computing all semantic tokens again simply send a delta.
	//
	// @since 3.16.0
	ResultID string `json:"resultId,omitempty"`

	// Data the actual tokens.
	//
	// @since 3.16.0
	Data []uint32 `json:"data"`
}

// SemanticTokensPartialResult.
//
// @since 3.16.0
type SemanticTokensPartialResult struct {
	// @since 3.16.0
	Data []uint32 `json:"data"`
}

// SemanticTokensLegend.
//
// @since 3.16.0
type SemanticTokensLegend struct {
	// TokenTypes the token types a server uses.
	//
	// @since 3.16.0
	TokenTypes []string `json:"tokenTypes"`

	// TokenModifiers the token modifiers a server uses.
	//
	// @since 3.16.0
	TokenModifiers []string `json:"tokenModifiers"`
}

// SemanticTokensFullDelta semantic tokens options to support deltas for full documents
//
// @since 3.18.0
type SemanticTokensFullDelta struct {
	// Delta the server supports deltas for full documents.
	//
	// @since 3.18.0
	Delta bool `json:"delta,omitempty"`
}

// SemanticTokensOptions.
//
// @since 3.16.0
type SemanticTokensOptions struct {
	// mixins
	WorkDoneProgressOptions

	// Legend the legend used by the server.
	//
	// @since 3.16.0
	Legend SemanticTokensLegend `json:"legend"`

	// Range server supports providing semantic tokens for a specific range of a document.
	//
	// @since 3.16.0
	Range SemanticTokensOptionsRange `json:"range,omitempty"`

	// Full server supports providing semantic tokens for a full document.
	//
	// @since 3.16.0
	Full SemanticTokensOptionsFull `json:"full,omitempty"`
}

// SemanticTokensRegistrationOptions.
//
// @since 3.16.0
type SemanticTokensRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	SemanticTokensOptions
	// mixins
	StaticRegistrationOptions
}

// SemanticTokensDeltaParams.
//
// @since 3.16.0
type SemanticTokensDeltaParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the text document.
	//
	// @since 3.16.0
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// PreviousResultID the result id of a previous response. The result Id can either point to a full response or a delta response depending on what was received last.
	//
	// @since 3.16.0
	PreviousResultID string `json:"previousResultId"`
}

// SemanticTokensEdit.
//
// @since 3.16.0
type SemanticTokensEdit struct {
	// Start the start offset of the edit.
	//
	// @since 3.16.0
	Start uint32 `json:"start"`

	// DeleteCount the count of elements to remove.
	//
	// @since 3.16.0
	DeleteCount uint32 `json:"deleteCount"`

	// Data the elements to insert.
	//
	// @since 3.16.0
	Data []uint32 `json:"data,omitempty"`
}

// SemanticTokensDelta.
//
// @since 3.16.0
type SemanticTokensDelta struct {
	// @since 3.16.0
	ResultID string `json:"resultId,omitempty"`

	// Edits the semantic token edits to transform a previous result into a new result.
	//
	// @since 3.16.0
	Edits []SemanticTokensEdit `json:"edits"`
}

// SemanticTokensDeltaPartialResult.
//
// @since 3.16.0
type SemanticTokensDeltaPartialResult struct {
	// @since 3.16.0
	Edits []SemanticTokensEdit `json:"edits"`
}

// SemanticTokensRangeParams.
//
// @since 3.16.0
type SemanticTokensRangeParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the text document.
	//
	// @since 3.16.0
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Range the range the semantic tokens are requested for.
	//
	// @since 3.16.0
	Range Range `json:"range"`
}

type LinkedEditingRangeParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
}

// LinkedEditingRanges the result of a linked editing range request.
//
// @since 3.16.0
type LinkedEditingRanges struct {
	// Ranges a list of ranges that can be edited together. The ranges must have identical length and contain identical text content. The ranges cannot overlap.
	//
	// @since 3.16.0
	Ranges []Range `json:"ranges"`

	// WordPattern an optional word pattern (regular expression) that describes valid contents for the given ranges. If
	// no pattern is provided, the client configuration's word pattern will be used.
	//
	// @since 3.16.0
	WordPattern string `json:"wordPattern,omitempty"`
}

type LinkedEditingRangeOptions struct {
	// mixins
	WorkDoneProgressOptions
}

type LinkedEditingRangeRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	LinkedEditingRangeOptions
	// mixins
	StaticRegistrationOptions
}

type MonikerParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
	PartialResultParams
}

// Moniker moniker definition to match LSIF 0.5 moniker definition.
//
// @since 3.16.0
type Moniker struct {
	// Scheme the scheme of the moniker. For example tsc or .Net.
	//
	// @since 3.16.0
	Scheme string `json:"scheme"`

	// Identifier the identifier of the moniker. The value is opaque in LSIF however schema owners are allowed to define the structure if they want.
	//
	// @since 3.16.0
	Identifier string `json:"identifier"`

	// Unique the scope in which the moniker is unique.
	//
	// @since 3.16.0
	Unique UniquenessLevel `json:"unique"`

	// Kind the moniker kind if known.
	//
	// @since 3.16.0
	Kind MonikerKind `json:"kind,omitempty"`
}

type MonikerOptions struct {
	// mixins
	WorkDoneProgressOptions
}

type MonikerRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	MonikerOptions
}

// TypeHierarchyPrepareParams the parameter of a `textDocument/prepareTypeHierarchy` request.
//
// @since 3.17.0
type TypeHierarchyPrepareParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
}

// TypeHierarchyItem.
//
// @since 3.17.0
type TypeHierarchyItem struct {
	// Name the name of this item.
	//
	// @since 3.17.0
	Name string `json:"name"`

	// Kind the kind of this item.
	//
	// @since 3.17.0
	Kind SymbolKind `json:"kind"`

	// Tags tags for this item.
	//
	// @since 3.17.0
	Tags []SymbolTag `json:"tags,omitempty"`

	// Detail more detail for this item, e.g. the signature of a function.
	//
	// @since 3.17.0
	Detail string `json:"detail,omitempty"`

	// URI the resource identifier of this item.
	//
	// @since 3.17.0
	URI DocumentURI `json:"uri"`

	// Range the range enclosing this symbol not including leading/trailing whitespace but everything else, e.g. comments and code.
	//
	// @since 3.17.0
	Range Range `json:"range"`

	// SelectionRange the range that should be selected and revealed when this symbol is being picked, e.g. the name of a function. Must be contained by the TypeHierarchyItem.range `range`.
	//
	// @since 3.17.0
	SelectionRange Range `json:"selectionRange"`

	// Data a data entry field that is preserved between a type hierarchy prepare and supertypes or subtypes requests. It could also be used to identify the type hierarchy in the server, helping improve the performance on resolving supertypes and subtypes.
	//
	// @since 3.17.0
	Data any `json:"data,omitempty"`
}

// TypeHierarchyOptions type hierarchy options used during static registration.
//
// @since 3.17.0
type TypeHierarchyOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// TypeHierarchyRegistrationOptions type hierarchy options used during static or dynamic registration.
//
// @since 3.17.0
type TypeHierarchyRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	TypeHierarchyOptions
	// mixins
	StaticRegistrationOptions
}

// TypeHierarchySupertypesParams the parameter of a `typeHierarchy/supertypes` request.
//
// @since 3.17.0
type TypeHierarchySupertypesParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// @since 3.17.0
	Item TypeHierarchyItem `json:"item"`
}

// TypeHierarchySubtypesParams the parameter of a `typeHierarchy/subtypes` request.
//
// @since 3.17.0
type TypeHierarchySubtypesParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// @since 3.17.0
	Item TypeHierarchyItem `json:"item"`
}

// InlineValueContext.
//
// @since 3.17.0
type InlineValueContext struct {
	// FrameID the stack frame (as a DAP Id) where the execution has stopped.
	//
	// @since 3.17.0
	FrameID int32 `json:"frameId"`

	// StoppedLocation the document range where execution has stopped. Typically the end position of the range denotes the line where the inline values are shown.
	//
	// @since 3.17.0
	StoppedLocation Range `json:"stoppedLocation"`
}

// InlineValueParams a parameter literal used in inline value requests.
//
// @since 3.17.0
type InlineValueParams struct {
	// mixins
	WorkDoneProgressParams

	// TextDocument the text document.
	//
	// @since 3.17.0
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Range the document range for which inline values should be computed.
	//
	// @since 3.17.0
	Range Range `json:"range"`

	// Context additional information about the context in which inline values were requested.
	//
	// @since 3.17.0
	Context InlineValueContext `json:"context"`
}

// InlineValueOptions inline value options used during static registration.
//
// @since 3.17.0
type InlineValueOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// InlineValueRegistrationOptions inline value options used during static or dynamic registration.
//
// @since 3.17.0
type InlineValueRegistrationOptions struct {
	// extends
	InlineValueOptions
	TextDocumentRegistrationOptions
	// mixins
	StaticRegistrationOptions
}

// InlayHintParams a parameter literal used in inlay hint requests.
//
// @since 3.17.0
type InlayHintParams struct {
	// mixins
	WorkDoneProgressParams

	// TextDocument the text document.
	//
	// @since 3.17.0
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Range the document range for which inlay hints should be computed.
	//
	// @since 3.17.0
	Range Range `json:"range"`
}

// InlayHintLabelPart an inlay hint label part allows for interactive and composite labels of inlay hints.
//
// @since 3.17.0
type InlayHintLabelPart struct {
	// Value the value of this label part.
	//
	// @since 3.17.0
	Value string `json:"value"`

	// Tooltip the tooltip text when you hover over this label part. Depending on the client capability `inlayHint.resolveSupport` clients might resolve this property late using the resolve request.
	//
	// @since 3.17.0
	Tooltip InlayHintLabelPartTooltip `json:"tooltip,omitempty"`

	// Location an optional source code location that represents this label part. The editor will use this location for the hover and for code navigation features: This part will become a clickable link that resolves
	// to the definition of the symbol at the given location (not necessarily the location itself), it shows the hover that shows at the given location, and it shows a context menu with further code navigation commands. Depending on the client capability `inlayHint.resolveSupport` clients might resolve this property late using the resolve request.
	//
	// @since 3.17.0
	Location *Location `json:"location,omitempty"`

	// Command an optional command for this label part. Depending on the client capability `inlayHint.resolveSupport` clients might resolve this property late using the resolve request.
	//
	// @since 3.17.0
	Command *Command `json:"command,omitempty"`
}

// InlayHint inlay hint information.
//
// @since 3.17.0
type InlayHint struct {
	// Position the position of this hint. If multiple hints have the same position, they will be shown in the order
	// they appear in the response.
	//
	// @since 3.17.0
	Position Position `json:"position"`

	// Label the label of this hint. A human readable string or an array of InlayHintLabelPart label parts. *Note* that neither the string nor the label part can be empty.
	//
	// @since 3.17.0
	Label InlayHintLabel `json:"label"`

	// Kind the kind of this hint. Can be omitted in which case the client should fall back to a reasonable default.
	//
	// @since 3.17.0
	Kind InlayHintKind `json:"kind,omitempty"`

	// TextEdits optional text edits that are performed when accepting this inlay hint. *Note* that edits are expected to change the document so that the inlay hint (or its nearest variant) is now part of the document
	// and the inlay hint itself is now obsolete.
	//
	// @since 3.17.0
	TextEdits []TextEdit `json:"textEdits,omitempty"`

	// Tooltip the tooltip text when you hover over this item.
	//
	// @since 3.17.0
	Tooltip InlayHintTooltip `json:"tooltip,omitempty"`

	// PaddingLeft render padding before the hint. Note: Padding should use the editor's background color, not the background color of the hint itself. That means padding can be used to visually align/separate an inlay hint.
	//
	// @since 3.17.0
	PaddingLeft bool `json:"paddingLeft,omitempty"`

	// PaddingRight render padding after the hint. Note: Padding should use the editor's background color, not the background color of the hint itself. That means padding can be used to visually align/separate an inlay hint.
	//
	// @since 3.17.0
	PaddingRight bool `json:"paddingRight,omitempty"`

	// Data a data entry field that is preserved on an inlay hint between a `textDocument/inlayHint` and a `inlayHint/resolve` request.
	//
	// @since 3.17.0
	Data any `json:"data,omitempty"`
}

// InlayHintOptions inlay hint options used during static registration.
//
// @since 3.17.0
type InlayHintOptions struct {
	// mixins
	WorkDoneProgressOptions

	// ResolveProvider the server provides support to resolve additional information for an inlay hint item.
	//
	// @since 3.17.0
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// InlayHintRegistrationOptions inlay hint options used during static or dynamic registration.
//
// @since 3.17.0
type InlayHintRegistrationOptions struct {
	// extends
	InlayHintOptions
	TextDocumentRegistrationOptions
	// mixins
	StaticRegistrationOptions
}

// DocumentDiagnosticParams parameters of the document diagnostic request.
//
// @since 3.17.0
type DocumentDiagnosticParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the text document.
	//
	// @since 3.17.0
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Identifier the additional identifier provided during registration.
	//
	// @since 3.17.0
	Identifier string `json:"identifier,omitempty"`

	// PreviousResultID the result id of a previous response if provided.
	//
	// @since 3.17.0
	PreviousResultID string `json:"previousResultId,omitempty"`
}

// UnchangedDocumentDiagnosticReport a diagnostic report indicating that the last returned report is still accurate.
//
// @since 3.17.0
type UnchangedDocumentDiagnosticReport struct {
	// ResultID a result id which will be sent on the next diagnostic request for the same document.
	//
	// @since 3.17.0
	ResultID string `json:"resultId"`
}

// FullDocumentDiagnosticReport a diagnostic report with a full set of problems.
//
// @since 3.17.0
type FullDocumentDiagnosticReport struct {
	// ResultID an optional result id. If provided it will be sent on the next diagnostic request for the same document.
	//
	// @since 3.17.0
	ResultID string `json:"resultId,omitempty"`

	// Items the actual items.
	//
	// @since 3.17.0
	Items []Diagnostic `json:"items"`
}

// DocumentDiagnosticReportPartialResult a partial result for a document diagnostic report.
//
// @since 3.17.0
type DocumentDiagnosticReportPartialResult struct {
	// @since 3.17.0
	RelatedDocuments map[DocumentURI]DocumentDiagnosticReportPartialResultRelatedDocuments `json:"relatedDocuments"`
}

// DiagnosticServerCancellationData cancellation data returned from a diagnostic request.
//
// @since 3.17.0
type DiagnosticServerCancellationData struct {
	// @since 3.17.0
	RetriggerRequest bool `json:"retriggerRequest"`
}

// DiagnosticOptions diagnostic options.
//
// @since 3.17.0
type DiagnosticOptions struct {
	// mixins
	WorkDoneProgressOptions

	// Identifier an optional identifier under which the diagnostics are managed by the client.
	//
	// @since 3.17.0
	Identifier string `json:"identifier,omitempty"`

	// InterFileDependencies whether the language has inter file dependencies meaning that editing code in one file can result in
	// a different diagnostic set in another file. Inter file dependencies are common for most programming languages and typically uncommon for linters.
	//
	// @since 3.17.0
	InterFileDependencies bool `json:"interFileDependencies"`

	// WorkspaceDiagnostics the server provides support for workspace diagnostics as well.
	//
	// @since 3.17.0
	WorkspaceDiagnostics bool `json:"workspaceDiagnostics"`
}

// DiagnosticRegistrationOptions diagnostic registration options.
//
// @since 3.17.0
type DiagnosticRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DiagnosticOptions
	// mixins
	StaticRegistrationOptions
}

// PreviousResultID a previous result id in a workspace pull request.
//
// @since 3.17.0
type PreviousResultID struct {
	// URI the URI for which the client knowns a result id.
	//
	// @since 3.17.0
	URI DocumentURI `json:"uri"`

	// Value the value of the previous result id.
	//
	// @since 3.17.0
	Value string `json:"value"`
}

// WorkspaceDiagnosticParams parameters of the workspace diagnostic request.
//
// @since 3.17.0
type WorkspaceDiagnosticParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// Identifier the additional identifier provided during registration.
	//
	// @since 3.17.0
	Identifier string `json:"identifier,omitempty"`

	// PreviousResultIDS the currently known diagnostic reports with their previous result ids.
	//
	// @since 3.17.0
	PreviousResultIDS []PreviousResultID `json:"previousResultIds"`
}

// WorkspaceDiagnosticReport a workspace diagnostic report.
//
// @since 3.17.0
type WorkspaceDiagnosticReport struct {
	// @since 3.17.0
	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
}

// WorkspaceDiagnosticReportPartialResult a partial result for a workspace diagnostic report.
//
// @since 3.17.0
type WorkspaceDiagnosticReportPartialResult struct {
	// @since 3.17.0
	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
}

// SelectedCompletionInfo describes the currently selected completion item.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type SelectedCompletionInfo struct {
	// Range the range that will be replaced if this completion item is accepted.
	//
	// @since 3.18.0 proposed
	Range Range `json:"range"`

	// Text the text the range will be replaced with if this completion is accepted.
	//
	// @since 3.18.0 proposed
	Text string `json:"text"`
}

// InlineCompletionContext provides information about the context in which an inline completion was requested. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionContext struct {
	// TriggerKind describes how the inline completion was triggered.
	//
	// @since 3.18.0 proposed
	TriggerKind InlineCompletionTriggerKind `json:"triggerKind"`

	// SelectedCompletionInfo provides information about the currently selected item in the autocomplete widget if it is visible.
	//
	// @since 3.18.0 proposed
	SelectedCompletionInfo *SelectedCompletionInfo `json:"selectedCompletionInfo,omitempty"`
}

// InlineCompletionParams a parameter literal used in inline completion requests.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams

	// Context additional information about the context in which inline completions were requested.
	//
	// @since 3.18.0 proposed
	Context InlineCompletionContext `json:"context"`
}

// InlineCompletionItem an inline completion item represents a text snippet that is proposed inline to complete text that is
// being typed. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionItem struct {
	// InsertText the text to replace the range with. Must be set.
	//
	// @since 3.18.0 proposed
	InsertText InlineCompletionItemInsertText `json:"insertText"`

	// FilterText a text that is used to decide if this inline completion should be shown. When `falsy` the InlineCompletionItem.insertText is used.
	//
	// @since 3.18.0 proposed
	FilterText string `json:"filterText,omitempty"`

	// Range the range to replace. Must begin and end on the same line.
	//
	// @since 3.18.0 proposed
	Range *Range `json:"range,omitempty"`

	// Command an optional Command that is executed *after* inserting this completion.
	//
	// @since 3.18.0 proposed
	Command *Command `json:"command,omitempty"`
}

// InlineCompletionList represents a collection of InlineCompletionItem inline completion items to be presented in the editor. 3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionList struct {
	// Items the inline completion items.
	//
	// @since 3.18.0 proposed
	Items []InlineCompletionItem `json:"items"`
}

// InlineCompletionOptions inline completion options used during static registration.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// InlineCompletionRegistrationOptions inline completion options used during static or dynamic registration.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type InlineCompletionRegistrationOptions struct {
	// extends
	InlineCompletionOptions
	TextDocumentRegistrationOptions
	// mixins
	StaticRegistrationOptions
}

// CodeActionTagOptions.
//
// @since 3.18.0 - proposed
type CodeActionTagOptions struct {
	// ValueSet the tags supported by the client.
	//
	// @since 3.18.0 - proposed
	ValueSet []CodeActionTag `json:"valueSet"`
}

// ServerCompletionItemOptions.
//
// @since 3.18.0
type ServerCompletionItemOptions struct {
	// LabelDetailsSupport the server has support for completion item label details (see also `CompletionItemLabelDetails`) when receiving a completion item in a resolve call.
	// @since 3.18.0
	LabelDetailsSupport bool `json:"labelDetailsSupport,omitempty"`
}

// CompletionOptions completion options.
type CompletionOptions struct {
	// mixins
	WorkDoneProgressOptions

	// TriggerCharacters most tools trigger completion request automatically without explicitly requesting it using a keyboard shortcut (e.g. Ctrl+Space). Typically they do so when the user starts to type an identifier. For example if the user types `c` in a JavaScript file code complete will automatically pop up present `console` besides others as a completion item. Characters that make up identifiers don't need to be listed here. If code complete should automatically be trigger on characters not being valid inside an identifier (for example `.` in JavaScript) list them in `triggerCharacters`.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// AllCommitCharacters the list of all possible characters that commit a completion. This field can be used if clients don't support individual commit characters per completion item. See `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport` If a server provides both `allCommitCharacters` and commit characters on an individual completion item the ones on the completion item win.
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`

	// ResolveProvider the server provides support to resolve additional information for a completion item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	// CompletionItem the server supports the following `CompletionItem` specific capabilities.
	CompletionItem *ServerCompletionItemOptions `json:"completionItem,omitempty"`
}

// HoverOptions hover options.
type HoverOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// SignatureHelpOptions server Capabilities for a SignatureHelpRequest.
type SignatureHelpOptions struct {
	// mixins
	WorkDoneProgressOptions

	// TriggerCharacters list of characters that trigger signature help automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// RetriggerCharacters list of characters that re-trigger signature help. These trigger characters are only active when signature help is already showing. All trigger characters are also counted as re-trigger characters.
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

// DefinitionOptions server Capabilities for a DefinitionRequest.
type DefinitionOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// ReferenceOptions reference options.
type ReferenceOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// DocumentHighlightOptions provider options for a DocumentHighlightRequest.
type DocumentHighlightOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// DocumentSymbolOptions provider options for a DocumentSymbolRequest.
type DocumentSymbolOptions struct {
	// mixins
	WorkDoneProgressOptions

	// Label a human-readable string that is shown when multiple outlines trees are shown for the same document.
	Label string `json:"label,omitempty"`
}

// CodeActionKindDocumentation documentation for a class of code actions.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type CodeActionKindDocumentation struct {
	// Kind the kind of the code action being documented. If the kind is generic, such as `CodeActionKind.Refactor`, the documentation will be shown whenever any refactorings are returned. If the kind if more specific, such as `CodeActionKind.RefactorExtract`, the documentation will only be shown when extract refactoring code actions are returned.
	//
	// @since 3.18.0 proposed
	Kind CodeActionKind `json:"kind"`

	// Command command that is ued to display the documentation to the user. The title of this documentation code action is taken from {@linkcode Command.title}.
	//
	// @since 3.18.0 proposed
	Command Command `json:"command"`
}

// CodeActionOptions provider options for a CodeActionRequest.
type CodeActionOptions struct {
	// mixins
	WorkDoneProgressOptions

	// CodeActionKinds codeActionKinds that this server may return. The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`

	// Documentation static documentation for a class of code actions. Documentation from the provider should be shown in
	// the code actions menu if either: - Code actions of `kind` are requested by the editor. In this
	// case, the editor will show the documentation that most closely matches the requested code action kind. For example, if a provider has documentation for both `Refactor` and `RefactorExtract`, when the user requests code actions for `RefactorExtract`, the editor will use the documentation for `RefactorExtract` instead of the documentation for `Refactor`. - Any code actions of `kind` are returned by the provider. At most one documentation entry should be shown per provider. 3.18.0 @proposed.
	Documentation []CodeActionKindDocumentation `json:"documentation,omitempty"`

	// ResolveProvider the server provides support to resolve additional information for a code action.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// CodeLensOptions code Lens provider options of a CodeLensRequest.
type CodeLensOptions struct {
	// mixins
	WorkDoneProgressOptions

	// ResolveProvider code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DocumentLinkOptions provider options for a DocumentLinkRequest.
type DocumentLinkOptions struct {
	// mixins
	WorkDoneProgressOptions

	// ResolveProvider document links have a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// WorkspaceSymbolOptions server capabilities for a WorkspaceSymbolRequest.
type WorkspaceSymbolOptions struct {
	// mixins
	WorkDoneProgressOptions

	// ResolveProvider the server provides support to resolve additional information for a workspace symbol.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DocumentFormattingOptions provider options for a DocumentFormattingRequest.
type DocumentFormattingOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// DocumentRangeFormattingOptions provider options for a DocumentRangeFormattingRequest.
type DocumentRangeFormattingOptions struct {
	// mixins
	WorkDoneProgressOptions

	// RangesSupport whether the server supports formatting multiple ranges at once.  3.18.0 @proposed.
	RangesSupport bool `json:"rangesSupport,omitempty"`
}

// DocumentOnTypeFormattingOptions provider options for a DocumentOnTypeFormattingRequest.
type DocumentOnTypeFormattingOptions struct {
	// FirstTriggerCharacter a character on which formatting should be triggered, like `{`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

	// MoreTriggerCharacter more trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
}

// RenameOptions provider options for a RenameRequest.
type RenameOptions struct {
	// mixins
	WorkDoneProgressOptions

	// PrepareProvider renames should be checked and tested before being executed.  version .
	PrepareProvider bool `json:"prepareProvider,omitempty"`
}

// ExecuteCommandOptions the server capabilities of a ExecuteCommandRequest.
type ExecuteCommandOptions struct {
	// mixins
	WorkDoneProgressOptions

	// Commands the commands to be executed on the server.
	Commands []string `json:"commands"`
}

// PublishDiagnosticsParams the publish diagnostic notification's parameters.
type PublishDiagnosticsParams struct {
	// URI the URI for which diagnostic information is reported.
	URI DocumentURI `json:"uri"`

	// Version optional the version number of the document the diagnostics are published for.
	Version int32 `json:"version,omitempty"`

	// Diagnostics an array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
}

// CompletionContext contains additional information about the context in which a completion request is triggered.
type CompletionContext struct {
	// TriggerKind how the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"triggerKind"`

	// TriggerCharacter the trigger character (a single character) that has trigger code complete. Is undefined if `triggerKind !== CompletionTriggerKind.TriggerCharacter`.
	TriggerCharacter string `json:"triggerCharacter,omitempty"`
}

// CompletionParams completion parameters.
type CompletionParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// Context the completion context. This is only available it the client specifies to send this using the client
	// capability `textDocument.completion.contextSupport === true`.
	Context *CompletionContext `json:"context,omitempty"`
}

// CompletionItemLabelDetails additional details for a completion item label.
//
// @since 3.17.0
type CompletionItemLabelDetails struct {
	// Detail an optional string which is rendered less prominently directly after CompletionItem.label label, without any spacing. Should be used for function signatures and type annotations.
	//
	// @since 3.17.0
	Detail string `json:"detail,omitempty"`

	// Description an optional string which is rendered less prominently after CompletionItem.detail. Should be used for fully qualified names and file paths.
	//
	// @since 3.17.0
	Description string `json:"description,omitempty"`
}

// InsertReplaceEdit a special text edit to provide an insert and a replace operation.
//
// @since 3.16.0
type InsertReplaceEdit struct {
	// NewText the string to be inserted.
	//
	// @since 3.16.0
	NewText string `json:"newText"`

	// Insert the range if the insert is requested.
	//
	// @since 3.16.0
	Insert Range `json:"insert"`

	// Replace the range if the replace is requested.
	//
	// @since 3.16.0
	Replace Range `json:"replace"`
}

// CompletionItem a completion item represents a text snippet that is proposed to complete text that is being typed.
type CompletionItem struct {
	// Label the label of this completion item. The label property is also by default the text that is inserted when selecting this completion. If label details are provided the label itself should be an unqualified name of the completion item.
	Label string `json:"label"`

	// LabelDetails additional details for the label
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails,omitempty"`

	// Kind the kind of this completion item. Based of the kind an icon is chosen by the editor.
	Kind CompletionItemKind `json:"kind,omitempty"`

	// Tags tags for this completion item.
	Tags []CompletionItemTag `json:"tags,omitempty"`

	// Detail a human-readable string with additional information about this item, like type or symbol information.
	Detail string `json:"detail,omitempty"`

	// Documentation a human-readable string that represents a doc-comment.
	Documentation CompletionItemDocumentation `json:"documentation,omitempty"`

	// Deprecated indicates if this item is deprecated.
	//
	// Deprecated: Use `tags` instead.
	Deprecated bool `json:"deprecated,omitempty"`

	// Preselect select this item when showing. *Note* that only one completion item can be selected and that the tool / client decides which item that is. The rule is that the *first* item of those that match best is
	// selected.
	Preselect bool `json:"preselect,omitempty"`

	// SortText a string that should be used when comparing this item with other items. When `falsy` the CompletionItem.label label is used.
	SortText string `json:"sortText,omitempty"`

	// FilterText a string that should be used when filtering a set of completion items. When `falsy` the CompletionItem.label label is used.
	FilterText string `json:"filterText,omitempty"`

	// InsertText a string that should be inserted into a document when selecting this completion. When `falsy` the CompletionItem.label label is used. The `insertText` is subject to interpretation by the client side. Some tools might not take the string literally. For example VS Code when code complete is requested in this example `con<cursor position>` and a completion item with an `insertText` of `console` is provided it will only insert `sole`. Therefore it is recommended to use `textEdit` instead since it avoids additional client side interpretation.
	InsertText string `json:"insertText,omitempty"`

	// InsertTextFormat the format of the insert text. The format applies to both the `insertText` property and the `newText` property of a provided `textEdit`. If omitted defaults to `InsertTextFormat.PlainText`. Please note that the insertTextFormat doesn't apply to `additionalTextEdits`.
	InsertTextFormat InsertTextFormat `json:"insertTextFormat,omitempty"`

	// InsertTextMode how whitespace and indentation is handled during completion item insertion. If not provided the clients default value depends on the `textDocument.completion.insertTextMode` client capability.
	InsertTextMode InsertTextMode `json:"insertTextMode,omitempty"`

	// TextEdit an TextEdit edit which is applied to a document when selecting this completion. When an edit is provided the value of CompletionItem.insertText insertText is ignored. Most editors support two different operations when accepting a completion item. One is to insert a completion text and the other is to replace an existing text with a completion text. Since this can usually not be predetermined by a server it can report both ranges. Clients need to signal support for `InsertReplaceEdits` via the `textDocument.completion.insertReplaceSupport` client capability property. *Note 1:* The text edit's range as well as both ranges from an insert replace edit must be a [single line] and they must contain the position at which completion has been requested. *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range must be a prefix of the edit's replace range, that means it must be contained and starting at the same position. 3.16.0 additional type `InsertReplaceEdit`.
	TextEdit CompletionItemTextEdit `json:"textEdit,omitempty"`

	// TextEditText the edit text used if the completion item is part of a CompletionList and CompletionList defines an item default for the text edit range. Clients will only honor this property if they opt into completion list item defaults using the capability `completionList.itemDefaults`. If not provided and a list's default range is provided the label property is used as a text.
	TextEditText string `json:"textEditText,omitempty"`

	// AdditionalTextEdits an optional array of additional TextEdit text edits that are applied when selecting this completion.
	// Edits must not overlap (including the same insert position) with the main CompletionItem.textEdit edit nor with themselves. Additional text edits should be used to change text unrelated to the current cursor position (for example adding an import statement at the top of the file if the completion item will insert an unqualified type).
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`

	// CommitCharacters an optional set of characters that when pressed while this completion is active will accept it first
	// and then type that character. *Note* that all commit characters should have `length=1` and that superfluous characters will be ignored.
	CommitCharacters []string `json:"commitCharacters,omitempty"`

	// Command an optional Command command that is executed *after* inserting this completion. *Note* that additional modifications to the current document should be described with the CompletionItem.additionalTextEdits additionalTextEdits-property.
	Command *Command `json:"command,omitempty"`

	// Data a data entry field that is preserved on a completion item between a CompletionRequest and a CompletionResolveRequest.
	Data any `json:"data,omitempty"`
}

// EditRangeWithInsertReplace edit range variant that includes ranges for insert and replace operations.
//
// @since 3.18.0
type EditRangeWithInsertReplace struct {
	// @since 3.18.0
	Insert Range `json:"insert"`

	// @since 3.18.0
	Replace Range `json:"replace"`
}

// CompletionItemDefaults in many cases the items of an actual completion result share the same value for properties like `commitCharacters` or the range of a text edit. A completion list can therefore define item defaults which will be used if a completion item itself doesn't specify the value. If a completion list specifies a default value and a completion item also specifies a corresponding value, the rules for combining these are defined by `applyKinds` (if the client supports it), defaulting to "replace". Servers are only allowed to return default values if the client signals support for this via the `completionList.itemDefaults` capability.
//
// @since 3.17.0
type CompletionItemDefaults struct {
	// CommitCharacters a default commit character set.
	// @since 3.17.0
	CommitCharacters []string `json:"commitCharacters,omitempty"`

	// EditRange a default edit range.
	// @since 3.17.0
	EditRange CompletionItemDefaultsEditRange `json:"editRange,omitempty"`

	// InsertTextFormat a default insert text format.
	// @since 3.17.0
	InsertTextFormat InsertTextFormat `json:"insertTextFormat,omitempty"`

	// InsertTextMode a default insert text mode.
	// @since 3.17.0
	InsertTextMode InsertTextMode `json:"insertTextMode,omitempty"`

	// Data a default data value.
	// @since 3.17.0
	Data any `json:"data,omitempty"`
}

// CompletionItemApplyKinds specifies how fields from a completion item should be combined with those from `completionList.itemDefaults`. If unspecified, all fields will be treated as "replace". If a field's value is "replace", the value from a completion item (if provided and not `null`) will always be used instead of the value from `completionItem.itemDefaults`. If a field's value is "merge", the values will be merged using the rules defined against each field below. Servers are only allowed to return `applyKind` if the client signals support for this via the `completionList.applyKindSupport` capability.
//
// @since 3.18.0
type CompletionItemApplyKinds struct {
	// CommitCharacters specifies whether commitCharacters on a completion will replace or be merged with those in `completionList.itemDefaults.commitCharacters`. If "replace", the commit characters from the completion item will always be used unless not provided, in which case those from `completionList.itemDefaults.commitCharacters` will be used. An empty list can be used if a completion item does not have any commit characters and also should not use those from `completionList.itemDefaults.commitCharacters`. If "merge" the commitCharacters for the completion will be the union of all values in both `completionList.itemDefaults.commitCharacters` and the completion's own `commitCharacters`.
	// @since 3.18.0
	CommitCharacters ApplyKind `json:"commitCharacters,omitempty"`

	// Data specifies whether the `data` field on a completion will replace or be merged with data from `completionList.itemDefaults.data`. If "replace", the data from the completion item will be used if provided
	// (and not `null`), otherwise `completionList.itemDefaults.data` will be used. An empty object can be used if a completion item does not have any data but also should not use the value from `completionList.itemDefaults.data`. If "merge", a shallow merge will be performed between `completionList.itemDefaults.data` and the completion's own data using the following rules: - If a completion's `data` field is not provided (or `null`), the entire `data` field from `completionList.itemDefaults.data` will be used as-is. - If a completion's `data` field is provided, each field will overwrite the field of the same name in `completionList.itemDefaults.data` but no merging of nested fields within that value will occur.
	// @since 3.18.0
	Data ApplyKind `json:"data,omitempty"`
}

// CompletionList represents a collection of CompletionItem completion items to be presented in the editor.
type CompletionList struct {
	// IsIncomplete this list it not complete. Further typing results in recomputing this list. Recomputed lists have all their items replaced (not appended) in the incomplete completion sessions.
	IsIncomplete bool `json:"isIncomplete"`

	// ItemDefaults in many cases the items of an actual completion result share the same value for properties like `commitCharacters` or the range of a text edit. A completion list can therefore define item defaults which will be used if a completion item itself doesn't specify the value. If a completion list specifies a default value and a completion item also specifies a corresponding value, the rules for combining these are defined by `applyKinds` (if the client supports it), defaulting to "replace". Servers are only allowed to return default values if the client signals support for this via the `completionList.itemDefaults` capability.
	ItemDefaults *CompletionItemDefaults `json:"itemDefaults,omitempty"`

	// ApplyKind specifies how fields from a completion item should be combined with those from `completionList.itemDefaults`. If unspecified, all fields will be treated as "replace". If a field's value is "replace", the value from a completion item (if provided and not `null`) will always be used instead of the value from `completionItem.itemDefaults`. If a field's value is "merge", the values will be merged using the rules defined against each field below. Servers are only allowed to return `applyKind` if the client signals support for this via the `completionList.applyKindSupport` capability.
	ApplyKind *CompletionItemApplyKinds `json:"applyKind,omitempty"`

	// Items the completion items.
	Items []CompletionItem `json:"items"`
}

// CompletionRegistrationOptions registration options for a CompletionRequest.
type CompletionRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	CompletionOptions
}

// HoverParams parameters for a HoverRequest.
type HoverParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
}

// Hover the result of a hover request.
type Hover struct {
	// Contents the hover's content.
	Contents HoverContents `json:"contents"`

	// Range an optional range inside the text document that is used to visualize the hover, e.g. by changing the
	// background color.
	Range *Range `json:"range,omitempty"`
}

// HoverRegistrationOptions registration options for a HoverRequest.
type HoverRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	HoverOptions
}

// ParameterInformation represents a parameter of a callable-signature. A parameter can have a label and a doc-comment.
type ParameterInformation struct {
	// Label the label of this parameter information. Either a string or an inclusive start and exclusive end offsets within its containing signature label. (see SignatureInformation.label). The offsets are based on a UTF-16 string representation as `Position` and `Range` does. To avoid ambiguities a server should use the [start, end] offset value instead of using a substring. Whether a client support this is controlled via `labelOffsetSupport` client capability. *Note*: a label of type string should be a substring of its containing signature label. Its intended use case is to highlight the parameter label
	// part in the `SignatureInformation.label`.
	Label ParameterInformationLabel `json:"label"`

	// Documentation the human-readable doc-comment of this parameter. Will be shown in the UI but can be omitted.
	Documentation ParameterInformationDocumentation `json:"documentation,omitempty"`
}

// SignatureInformation represents the signature of something callable. A signature can have a label, like a function-name, a doc-comment, and a set of parameters.
type SignatureInformation struct {
	// Label the label of this signature. Will be shown in the UI.
	Label string `json:"label"`

	// Documentation the human-readable doc-comment of this signature. Will be shown in the UI but can be omitted.
	Documentation SignatureInformationDocumentation `json:"documentation,omitempty"`

	// Parameters the parameters of this signature.
	Parameters []ParameterInformation `json:"parameters,omitempty"`

	// ActiveParameter the index of the active parameter. If `null`, no parameter of the signature is active (for example a
	// named argument that does not match any declared parameters). This is only valid if the client specifies the client capability `textDocument.signatureHelp.noActiveParameterSupport === true` If provided (or `null`), this is used in place of `SignatureHelp.activeParameter`.
	ActiveParameter uint32 `json:"activeParameter,omitempty"`
}

// SignatureHelp signature help represents the signature of something callable. There can be multiple signature but only one active and only one active parameter.
type SignatureHelp struct {
	// Signatures one or more signatures.
	Signatures []SignatureInformation `json:"signatures"`

	// ActiveSignature the active signature. If omitted or the value lies outside the range of `signatures` the value defaults to zero or is ignored if the `SignatureHelp` has no signatures. Whenever possible implementors should make an active decision about the active signature and shouldn't rely on a default value. In future version of the protocol this property might become mandatory to better express this.
	ActiveSignature uint32 `json:"activeSignature,omitempty"`

	// ActiveParameter the active parameter of the active signature. If `null`, no parameter of the signature is active (for example a named argument that does not match any declared parameters). This is only valid if the client specifies the client capability `textDocument.signatureHelp.noActiveParameterSupport === true`
	// If omitted or the value lies outside the range of `signatures[activeSignature].parameters` defaults to 0 if the active signature has parameters. If the active signature has no parameters it is ignored. In future version of the protocol this property might become mandatory (but still nullable) to better express the active parameter if the active signature does have any.
	ActiveParameter uint32 `json:"activeParameter,omitempty"`
}

// SignatureHelpContext additional information about the context in which a signature help request was triggered.
//
// @since 3.15.0
type SignatureHelpContext struct {
	// TriggerKind action that caused signature help to be triggered.
	//
	// @since 3.15.0
	TriggerKind SignatureHelpTriggerKind `json:"triggerKind"`

	// TriggerCharacter character that caused signature help to be triggered. This is undefined when `triggerKind !== SignatureHelpTriggerKind.TriggerCharacter`.
	//
	// @since 3.15.0
	TriggerCharacter string `json:"triggerCharacter,omitempty"`

	// IsRetrigger `true` if signature help was already showing when it was triggered. Retriggers occurs when the signature help is already active and can be caused by actions such as typing a trigger character, a cursor move, or document content changes.
	//
	// @since 3.15.0
	IsRetrigger bool `json:"isRetrigger"`

	// ActiveSignatureHelp the currently active `SignatureHelp`. The `activeSignatureHelp` has its `SignatureHelp.activeSignature` field updated based on the user navigating through available signatures.
	//
	// @since 3.15.0
	ActiveSignatureHelp *SignatureHelp `json:"activeSignatureHelp,omitempty"`
}

// SignatureHelpParams parameters for a SignatureHelpRequest.
type SignatureHelpParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams

	// Context the signature help context. This is only available if the client specifies to send this using the client capability `textDocument.signatureHelp.contextSupport === true`
	Context *SignatureHelpContext `json:"context,omitempty"`
}

// SignatureHelpRegistrationOptions registration options for a SignatureHelpRequest.
type SignatureHelpRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	SignatureHelpOptions
}

// DefinitionParams parameters for a DefinitionRequest.
type DefinitionParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
	PartialResultParams
}

// DefinitionRegistrationOptions registration options for a DefinitionRequest.
type DefinitionRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DefinitionOptions
}

// ReferenceContext value-object that contains additional information when requesting references.
type ReferenceContext struct {
	// IncludeDeclaration include the declaration of the current symbol.
	IncludeDeclaration bool `json:"includeDeclaration"`
}

// ReferenceParams parameters for a ReferencesRequest.
type ReferenceParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	Context ReferenceContext `json:"context"`
}

// ReferenceRegistrationOptions registration options for a ReferencesRequest.
type ReferenceRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	ReferenceOptions
}

// DocumentHighlightParams parameters for a DocumentHighlightRequest.
type DocumentHighlightParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
	PartialResultParams
}

// DocumentHighlight a document highlight is a range inside a text document which deserves special attention. Usually a document highlight is visualized by changing the background color of its range.
type DocumentHighlight struct {
	// Range the range this highlight applies to.
	Range Range `json:"range"`

	// Kind the highlight kind, default is DocumentHighlightKind.Text text.
	Kind DocumentHighlightKind `json:"kind,omitempty"`
}

// DocumentHighlightRegistrationOptions registration options for a DocumentHighlightRequest.
type DocumentHighlightRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DocumentHighlightOptions
}

// DocumentSymbolParams parameters for a DocumentSymbolRequest.
type DocumentSymbolParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// BaseSymbolInformation a base for all symbol information.
type BaseSymbolInformation struct {
	// Name the name of this symbol.
	Name string `json:"name"`

	// Kind the kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags tags for this symbol.
	Tags []SymbolTag `json:"tags,omitempty"`

	// ContainerName the name of the symbol containing this symbol. This information is for user interface purposes (e.g.
	// to render a qualifier in the user interface if necessary). It can't be used to re-infer a hierarchy for the document symbols.
	ContainerName string `json:"containerName,omitempty"`
}

// SymbolInformation represents information about programming constructs like variables, classes, interfaces etc.
type SymbolInformation struct {
	// extends
	BaseSymbolInformation

	// Deprecated indicates if this symbol is deprecated.
	//
	// Deprecated: Use tags instead.
	Deprecated bool `json:"deprecated,omitempty"`

	// Location the location of this symbol. The location's range is used by a tool to reveal the location in the editor. If the symbol is selected in the tool the range's start information is used to position the cursor. So the range usually spans more than the actual symbol's name and does normally include things
	// like visibility modifiers. The range doesn't have to denote a node range in the sense of an abstract syntax tree. It can therefore not be used to re-construct a hierarchy of the symbols.
	Location Location `json:"location"`
}

// DocumentSymbol represents programming constructs like variables, classes, interfaces etc. that appear in a document. Document symbols can be hierarchical and they have two ranges: one that encloses its definition and one that points to its most interesting range, e.g. the range of an identifier.
type DocumentSymbol struct {
	// Name the name of this symbol. Will be displayed in the user interface and therefore must not be an empty string or a string only consisting of white spaces.
	Name string `json:"name"`

	// Detail more detail for this symbol, e.g the signature of a function.
	Detail string `json:"detail,omitempty"`

	// Kind the kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags tags for this document symbol.
	Tags []SymbolTag `json:"tags,omitempty"`

	// Deprecated indicates if this symbol is deprecated.
	//
	// Deprecated: Use tags instead.
	Deprecated bool `json:"deprecated,omitempty"`

	// Range the range enclosing this symbol not including leading/trailing whitespace but everything else like comments. This information is typically used to determine if the clients cursor is inside the symbol to reveal in the symbol in the UI.
	Range Range `json:"range"`

	// SelectionRange the range that should be selected and revealed when this symbol is being picked, e.g the name of a function. Must be contained by the `range`.
	SelectionRange Range `json:"selectionRange"`

	// Children children of this symbol, e.g. properties of a class.
	Children []DocumentSymbol `json:"children,omitempty"`
}

// DocumentSymbolRegistrationOptions registration options for a DocumentSymbolRequest.
type DocumentSymbolRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DocumentSymbolOptions
}

// CodeActionContext contains additional diagnostic information about the context in which a CodeActionProvider.provideCodeActions code action is run.
type CodeActionContext struct {
	// Diagnostics an array of diagnostics known on the client side overlapping the range provided to the `textDocument/codeAction` request. They are provided so that the server knows which errors are currently presented to the user for the given range. There is no guarantee that these accurately reflect the error state of the resource. The primary parameter to compute code actions is the provided range.
	Diagnostics []Diagnostic `json:"diagnostics"`

	// Only requested kind of actions to return. Actions not of this kind are filtered out by the client before being shown. So servers can omit computing them.
	Only []CodeActionKind `json:"only,omitempty"`

	// TriggerKind the reason why code actions were requested.
	TriggerKind CodeActionTriggerKind `json:"triggerKind,omitempty"`
}

// CodeActionParams the parameters of a CodeActionRequest.
type CodeActionParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the document in which the command was invoked.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Range the range for which the command was invoked.
	Range Range `json:"range"`

	// Context context carrying additional information.
	Context CodeActionContext `json:"context"`
}

// CodeActionDisabled captures why the code action is currently disabled.
//
// @since 3.18.0
type CodeActionDisabled struct {
	// Reason human readable description of why the code action is currently disabled. This is displayed in the code actions UI.
	//
	// @since 3.18.0
	Reason string `json:"reason"`
}

// CodeAction a code action represents a change that can be performed in code, e.g. to fix a problem or to refactor code. A CodeAction must set either `edit` and/or a `command`. If both are supplied, the `edit` is applied first, then the `command` is executed.
type CodeAction struct {
	// Title a short, human-readable, title for this code action.
	Title string `json:"title"`

	// Kind the kind of the code action. Used to filter code actions.
	Kind CodeActionKind `json:"kind,omitempty"`

	// Diagnostics the diagnostics that this code action resolves.
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`

	// IsPreferred marks this as a preferred action. Preferred actions are used by the `auto fix` command and can be targeted by keybindings. A quick fix should be marked preferred if it properly addresses the underlying error. A refactoring should be marked preferred if it is the most reasonable choice of actions to take.
	IsPreferred bool `json:"isPreferred,omitempty"`

	// Disabled marks that the code action cannot currently be applied. Clients should follow the following guidelines regarding disabled code actions: - Disabled code actions are not shown in automatic [lightbulbs](https://code.visualstudio.com/docs/editor/editingevolved#_code-action) code action menus. - Disabled
	// actions are shown as faded out in the code action menu when the user requests a more specific type of code action, such as refactorings. - If the user has a [keybinding](https://code.visualstudio.com/docs/editor/refactoring#_keybindings-for-code-actions) that auto applies a code action and only disabled code actions are returned, the client should show the user an error message with `reason`
	// in the editor.
	Disabled *CodeActionDisabled `json:"disabled,omitempty"`

	// Edit the workspace edit this code action performs.
	Edit *WorkspaceEdit `json:"edit,omitempty"`

	// Command a command this code action executes. If a code action provides an edit and a command, first the edit
	// is executed and then the command.
	Command *Command `json:"command,omitempty"`

	// Data a data entry field that is preserved on a code action between a `textDocument/codeAction` and a `codeAction/resolve` request.
	Data any `json:"data,omitempty"`

	// Tags tags for this code action.  3.18.0 - proposed.
	Tags []CodeActionTag `json:"tags,omitempty"`
}

// CodeActionRegistrationOptions registration options for a CodeActionRequest.
type CodeActionRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	CodeActionOptions
}

// LocationURIOnly location with only uri and does not include range.
//
// @since 3.18.0
type LocationURIOnly struct {
	// @since 3.18.0
	URI DocumentURI `json:"uri"`
}

// CodeLensParams the parameters of a CodeLensRequest.
type CodeLensParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the document to request code lens for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// CodeLens a code lens represents a Command command that should be shown along with source text, like the number of references, a way to run tests, etc. A code lens is _unresolved_ when no command is associated to it. For performance reasons the creation of a code lens and resolving should be done in two stages.
type CodeLens struct {
	// Range the range in which this code lens is valid. Should only span a single line.
	Range Range `json:"range"`

	// Command the command this code lens represents.
	Command *Command `json:"command,omitempty"`

	// Data a data entry field that is preserved on a code lens item between a CodeLensRequest and a CodeLensResolveRequest.
	Data any `json:"data,omitempty"`
}

// CodeLensRegistrationOptions registration options for a CodeLensRequest.
type CodeLensRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	CodeLensOptions
}

// DocumentLinkParams the parameters of a DocumentLinkRequest.
type DocumentLinkParams struct {
	// mixins
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument the document to provide document links for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// DocumentLink a document link is a range in a text document that links to an internal or external resource, like another text document or a web site.
type DocumentLink struct {
	// Range the range this link applies to.
	Range Range `json:"range"`

	// Target the uri this link points to. If missing a resolve request is sent later.
	Target uri.URI `json:"target,omitempty"`

	// Tooltip the tooltip text when you hover over this link. If a tooltip is provided, is will be displayed in a string that includes instructions on how to trigger the link, such as `{0} (ctrl + click)`. The specific instructions vary depending on OS, user settings, and localization.
	Tooltip string `json:"tooltip,omitempty"`

	// Data a data entry field that is preserved on a document link between a DocumentLinkRequest and a DocumentLinkResolveRequest.
	Data any `json:"data,omitempty"`
}

// DocumentLinkRegistrationOptions registration options for a DocumentLinkRequest.
type DocumentLinkRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DocumentLinkOptions
}

// FormattingOptions value-object describing what options formatting should use.
type FormattingOptions struct {
	// TabSize size of a tab in spaces.
	TabSize uint32 `json:"tabSize"`

	// InsertSpaces prefer spaces over tabs.
	InsertSpaces bool `json:"insertSpaces"`

	// TrimTrailingWhitespace trim trailing whitespace on a line.
	TrimTrailingWhitespace bool `json:"trimTrailingWhitespace,omitempty"`

	// InsertFinalNewline insert a newline character at the end of the file if one does not exist.
	InsertFinalNewline bool `json:"insertFinalNewline,omitempty"`

	// TrimFinalNewlines trim all newlines after the final newline at the end of the file.
	TrimFinalNewlines bool `json:"trimFinalNewlines,omitempty"`
}

// DocumentFormattingParams the parameters of a DocumentFormattingRequest.
type DocumentFormattingParams struct {
	// mixins
	WorkDoneProgressParams

	// TextDocument the document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Options the format options.
	Options FormattingOptions `json:"options"`
}

// DocumentFormattingRegistrationOptions registration options for a DocumentFormattingRequest.
type DocumentFormattingRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DocumentFormattingOptions
}

// DocumentRangeFormattingParams the parameters of a DocumentRangeFormattingRequest.
type DocumentRangeFormattingParams struct {
	// mixins
	WorkDoneProgressParams

	// TextDocument the document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Range the range to format.
	Range Range `json:"range"`

	// Options the format options.
	Options FormattingOptions `json:"options"`
}

// DocumentRangeFormattingRegistrationOptions registration options for a DocumentRangeFormattingRequest.
type DocumentRangeFormattingRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DocumentRangeFormattingOptions
}

// DocumentRangesFormattingParams the parameters of a DocumentRangesFormattingRequest.  3.18.0 @proposed.
//
// @since 3.18.0 proposed
type DocumentRangesFormattingParams struct {
	// mixins
	WorkDoneProgressParams

	// TextDocument the document to format.
	//
	// @since 3.18.0 proposed
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Ranges the ranges to format.
	//
	// @since 3.18.0 proposed
	Ranges []Range `json:"ranges"`

	// Options the format options.
	//
	// @since 3.18.0 proposed
	Options FormattingOptions `json:"options"`
}

// DocumentOnTypeFormattingParams the parameters of a DocumentOnTypeFormattingRequest.
type DocumentOnTypeFormattingParams struct {
	// TextDocument the document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position the position around which the on type formatting should happen. This is not necessarily the exact position where the character denoted by the property `ch` got typed.
	Position Position `json:"position"`

	// Ch the character that has been typed that triggered the formatting on type request. That is not necessarily the last character that got inserted into the document since the client could auto insert characters as well (e.g. like automatic brace completion).
	Ch string `json:"ch"`

	// Options the formatting options.
	Options FormattingOptions `json:"options"`
}

// DocumentOnTypeFormattingRegistrationOptions registration options for a DocumentOnTypeFormattingRequest.
type DocumentOnTypeFormattingRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DocumentOnTypeFormattingOptions
}

// RenameParams the parameters of a RenameRequest.
type RenameParams struct {
	// mixins
	WorkDoneProgressParams

	// TextDocument the document to rename.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position the position at which this request was sent.
	Position Position `json:"position"`

	// NewName the new name of the symbol. If the given name is not valid the request must return a ResponseError with an appropriate message set.
	NewName string `json:"newName"`
}

// RenameRegistrationOptions registration options for a RenameRequest.
type RenameRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	RenameOptions
}

type PrepareRenameParams struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams
}

// ExecuteCommandParams the parameters of a ExecuteCommandRequest.
type ExecuteCommandParams struct {
	// mixins
	WorkDoneProgressParams

	// Command the identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments arguments that the command should be invoked with.
	Arguments []any `json:"arguments,omitempty"`
}

// ExecuteCommandRegistrationOptions registration options for a ExecuteCommandRequest.
type ExecuteCommandRegistrationOptions struct {
	// extends
	ExecuteCommandOptions
}

// InlineValueText provide inline value as text.
//
// @since 3.17.0
type InlineValueText struct {
	// Range the document range for which the inline value applies.
	//
	// @since 3.17.0
	Range Range `json:"range"`

	// Text the text of the inline value.
	//
	// @since 3.17.0
	Text string `json:"text"`
}

// InlineValueVariableLookup provide inline value through a variable lookup. If only a range is specified, the variable name will
// be extracted from the underlying document. An optional variable name can be used to override the extracted name.
//
// @since 3.17.0
type InlineValueVariableLookup struct {
	// Range the document range for which the inline value applies. The range is used to extract the variable name from the underlying document.
	//
	// @since 3.17.0
	Range Range `json:"range"`

	// VariableName if specified the name of the variable to look up.
	//
	// @since 3.17.0
	VariableName string `json:"variableName,omitempty"`

	// CaseSensitiveLookup how to perform the lookup.
	//
	// @since 3.17.0
	CaseSensitiveLookup bool `json:"caseSensitiveLookup"`
}

// InlineValueEvaluatableExpression provide an inline value through an expression evaluation. If only a range is specified, the expression will be extracted from the underlying document. An optional expression can be used to override the extracted expression.
//
// @since 3.17.0
type InlineValueEvaluatableExpression struct {
	// Range the document range for which the inline value applies. The range is used to extract the evaluatable expression from the underlying document.
	//
	// @since 3.17.0
	Range Range `json:"range"`

	// Expression if specified the expression overrides the extracted expression.
	//
	// @since 3.17.0
	Expression string `json:"expression,omitempty"`
}

// RelatedFullDocumentDiagnosticReport a full diagnostic report with a set of related documents.
//
// @since 3.17.0
type RelatedFullDocumentDiagnosticReport struct {
	// extends
	FullDocumentDiagnosticReport

	// RelatedDocuments diagnostics of related documents. This information is useful in programming languages where code in a file A can generate diagnostics in a file B which A depends on. An example of such a language is C/C++ where marco definitions in a file a.cpp and result in errors in a header file b.hpp.
	// @since 3.17.0
	RelatedDocuments map[DocumentURI]RelatedFullDocumentDiagnosticReportRelatedDocuments `json:"relatedDocuments,omitempty"`
}

// RelatedUnchangedDocumentDiagnosticReport an unchanged diagnostic report with a set of related documents.
//
// @since 3.17.0
type RelatedUnchangedDocumentDiagnosticReport struct {
	// extends
	UnchangedDocumentDiagnosticReport

	// RelatedDocuments diagnostics of related documents. This information is useful in programming languages where code in a file A can generate diagnostics in a file B which A depends on. An example of such a language is C/C++ where marco definitions in a file a.cpp and result in errors in a header file b.hpp.
	// @since 3.17.0
	RelatedDocuments map[DocumentURI]RelatedUnchangedDocumentDiagnosticReportRelatedDocuments `json:"relatedDocuments,omitempty"`
}

// PrepareRenamePlaceholder.
//
// @since 3.18.0
type PrepareRenamePlaceholder struct {
	// @since 3.18.0
	Range Range `json:"range"`

	// @since 3.18.0
	Placeholder string `json:"placeholder"`
}

// PrepareRenameDefaultBehavior.
//
// @since 3.18.0
type PrepareRenameDefaultBehavior struct {
	// @since 3.18.0
	DefaultBehavior bool `json:"defaultBehavior"`
}

// WorkspaceFullDocumentDiagnosticReport a full document diagnostic report for a workspace diagnostic result.
//
// @since 3.17.0
type WorkspaceFullDocumentDiagnosticReport struct {
	// extends
	FullDocumentDiagnosticReport

	// URI the URI for which diagnostic information is reported.
	//
	// @since 3.17.0
	URI DocumentURI `json:"uri"`

	// Version the version number for which the diagnostics are reported. If the document is not marked as open `null` can be provided.
	//
	// @since 3.17.0
	Version int32 `json:"version,omitempty"`
}

// WorkspaceUnchangedDocumentDiagnosticReport an unchanged document diagnostic report for a workspace diagnostic result.
//
// @since 3.17.0
type WorkspaceUnchangedDocumentDiagnosticReport struct {
	// extends
	UnchangedDocumentDiagnosticReport

	// URI the URI for which diagnostic information is reported.
	//
	// @since 3.17.0
	URI DocumentURI `json:"uri"`

	// Version the version number for which the diagnostics are reported. If the document is not marked as open `null` can be provided.
	//
	// @since 3.17.0
	Version int32 `json:"version,omitempty"`
}

// TextDocumentContentChangePartial.
//
// @since 3.18.0
type TextDocumentContentChangePartial struct {
	// Range the range of the document that changed.
	//
	// @since 3.18.0
	Range Range `json:"range"`

	// RangeLength the optional length of the range that got replaced.
	//
	// Deprecated: use range instead.
	//
	// @since 3.18.0
	RangeLength uint32 `json:"rangeLength,omitempty"`

	// Text the new text for the provided range.
	//
	// @since 3.18.0
	Text string `json:"text"`
}

// TextDocumentContentChangeWholeDocument.
//
// @since 3.18.0
type TextDocumentContentChangeWholeDocument struct {
	// Text the new text of the whole document.
	//
	// @since 3.18.0
	Text string `json:"text"`
}

// MarkedStringWithLanguage.
//
// @since 3.18.0
type MarkedStringWithLanguage struct {
	// @since 3.18.0
	Language string `json:"language"`

	// @since 3.18.0
	Value string `json:"value"`
}

// NotebookCellTextDocumentFilter a notebook cell text document filter denotes a cell text document by different properties.
//
// @since 3.17.0
type NotebookCellTextDocumentFilter struct {
	// Notebook a filter that matches against the notebook containing the notebook cell. If a string value is provided it matches against the notebook type. '*' matches every notebook.
	//
	// @since 3.17.0
	Notebook NotebookCellTextDocumentFilterNotebook `json:"notebook"`

	// Language a language id like `python`. Will be matched against the language id of the notebook cell document. '*' matches every language.
	//
	// @since 3.17.0
	Language string `json:"language,omitempty"`
}

// TextDocumentFilterLanguage a document filter where `language` is required field.
//
// @since 3.18.0
type TextDocumentFilterLanguage struct {
	// Language a language id, like `typescript`.
	//
	// @since 3.18.0
	Language string `json:"language"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0
	Scheme string `json:"scheme,omitempty"`

	// Pattern a glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples. 3.18.0 - support for relative patterns.
	// @since 3.18.0
	Pattern *GlobPattern `json:"pattern,omitempty"`
}

// TextDocumentFilterScheme a document filter where `scheme` is required field.
//
// @since 3.18.0
type TextDocumentFilterScheme struct {
	// Language a language id, like `typescript`.
	//
	// @since 3.18.0
	Language string `json:"language,omitempty"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0
	Scheme string `json:"scheme"`

	// Pattern a glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples. 3.18.0 - support for relative patterns.
	// @since 3.18.0
	Pattern *GlobPattern `json:"pattern,omitempty"`
}

// TextDocumentFilterPattern a document filter where `pattern` is required field.
//
// @since 3.18.0
type TextDocumentFilterPattern struct {
	// Language a language id, like `typescript`.
	//
	// @since 3.18.0
	Language string `json:"language,omitempty"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0
	Scheme string `json:"scheme,omitempty"`

	// Pattern a glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples. 3.18.0 - support for relative patterns.
	// @since 3.18.0
	Pattern GlobPattern `json:"pattern"`
}
