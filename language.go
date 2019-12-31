// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strconv"

	"github.com/go-language-server/uri"
)

// CompletionParams params of Completion Request.
type CompletionParams struct {
	TextDocumentPositionParams

	// Context is the completion context. This is only available if the client specifies
	// to send this using `ClientCapabilities.textDocument.completion.contextSupport === true`
	Context *CompletionContext `json:"context,omitempty"`
}

// CompletionTriggerKind how a completion was triggered.
type CompletionTriggerKind float64

const (
	// Invoked completion was triggered by typing an identifier (24x7 code
	// complete), manual invocation (e.g Ctrl+Space) or via API.
	Invoked CompletionTriggerKind = 1

	// TriggerCharacter completion was triggered by a trigger character specified by
	// the `triggerCharacters` properties of the `CompletionRegistrationOptions`.
	TriggerCharacter CompletionTriggerKind = 2

	// TriggerForIncompleteCompletions completion was re-triggered as the current completion list is incomplete.
	TriggerForIncompleteCompletions CompletionTriggerKind = 3
)

// String implements fmt.Stringer.
func (k CompletionTriggerKind) String() string {
	switch k {
	case Invoked:
		return "Invoked"
	case TriggerCharacter:
		return "TriggerCharacter"
	case TriggerForIncompleteCompletions:
		return "TriggerForIncompleteCompletions"
	default:
		return strconv.FormatFloat(float64(k), 'f', -10, 64)
	}
}

// CompletionContext contains additional information about the context in which a completion request is triggered.
type CompletionContext struct {
	// TriggerCharacter is the trigger character (a single character) that has trigger code complete.
	// Is undefined if `triggerKind !== CompletionTriggerKind.TriggerCharacter`
	TriggerCharacter string `json:"triggerCharacter,omitempty"`

	// TriggerKind how the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"triggerKind"`
}

// CompletionList represents a collection of [completion items](#CompletionItem) to be presented
// in the editor.
type CompletionList struct {
	// IsIncomplete this list it not complete. Further typing should result in recomputing
	// this list.
	IsIncomplete bool `json:"isIncomplete"`

	// Items is the completion items.
	Items []CompletionItem `json:"items"`
}

// InsertTextFormat defines whether the insert text in a completion item should be interpreted as
// plain text or a snippet.
type InsertTextFormat float64

const (
	// TextFormatPlainText is the primary text to be inserted is treated as a plain string.
	TextFormatPlainText InsertTextFormat = 1

	// TextFormatSnippet is the primary text to be inserted is treated as a snippet.
	//
	// A snippet can define tab stops and placeholders with `$1`, `$2`
	// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	// the end of the snippet. Placeholders with equal identifiers are linked,
	// that is typing in one will update others too.
	TextFormatSnippet InsertTextFormat = 2
)

// String implements fmt.Stringer.
func (tf InsertTextFormat) String() string {
	switch tf {
	case TextFormatPlainText:
		return "PlainText"
	case TextFormatSnippet:
		return "Snippet"
	default:
		return strconv.FormatFloat(float64(tf), 'f', -10, 64)
	}
}

// CompletionItem item of CompletionList.
type CompletionItem struct {
	// AdditionalTextEdits an optional array of additional text edits that are applied when
	// selecting this completion. Edits must not overlap (including the same insert position)
	// with the main edit nor with themselves.
	//
	// Additional text edits should be used to change text unrelated to the current cursor position
	// (for example adding an import statement at the top of the file if the completion item will
	// insert an unqualified type).
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`

	// Command an optional command that is executed *after* inserting this completion. *Note* that
	// additional modifications to the current document should be described with the
	// additionalTextEdits-property.
	Command *Command `json:"command,omitempty"`

	// CommitCharacters an optional set of characters that when pressed while this completion is active will accept it first and
	// then type that character. *Note* that all commit characters should have `length=1` and that superfluous
	// characters will be ignored.
	CommitCharacters []string `json:"commitCharacters,omitempty"`

	// Data an data entry field that is preserved on a completion item between
	// a completion and a completion resolve request.
	Data interface{} `json:"data"`

	// Deprecated indicates if this item is deprecated.
	Deprecated bool `json:"deprecated,omitempty"`

	// Detail a human-readable string with additional information
	// about this item, like type or symbol information.
	Detail string `json:"detail,omitempty"`

	// Documentation a human-readable string that represents a doc-comment.
	Documentation interface{} `json:"documentation,omitempty"`

	// FilterText a string that should be used when filtering a set of
	// completion items. When `falsy` the label is used.
	FilterText string `json:"filterText,omitempty"`

	// InsertText a string that should be inserted into a document when selecting
	// this completion. When `falsy` the label is used.
	//
	// The `insertText` is subject to interpretation by the client side.
	// Some tools might not take the string literally. For example
	// VS Code when code complete is requested in this example `con<cursor position>`
	// and a completion item with an `insertText` of `console` is provided it
	// will only insert `sole`. Therefore it is recommended to use `textEdit` instead
	// since it avoids additional client side interpretation.
	InsertText string `json:"insertText,omitempty"`

	// InsertTextFormat is the format of the insert text. The format applies to both the `insertText` property
	// and the `newText` property of a provided `textEdit`.
	InsertTextFormat InsertTextFormat `json:"insertTextFormat,omitempty"`

	// Kind is the kind of this completion item. Based of the kind
	// an icon is chosen by the editor.
	Kind CompletionItemKind `json:"kind,omitempty"`

	// Label is the label of this completion item. By default
	// also the text that is inserted when selecting
	// this completion.
	Label string `json:"label"`

	// Preselect select this item when showing.
	//
	// *Note* that only one completion item can be selected and that the
	// tool / client decides which item that is. The rule is that the *first*
	// item of those that match best is selected.
	Preselect bool `json:"preselect,omitempty"`

	// SortText a string that should be used when comparing this item
	// with other items. When `falsy` the label is used.
	SortText string `json:"sortText,omitempty"`

	// TextEdit an edit which is applied to a document when selecting this completion. When an edit is provided the value of
	// `insertText` is ignored.
	//
	// *Note:* The range of the edit must be a single line range and it must contain the position at which completion
	// has been requested.
	TextEdit *TextEdit `json:"textEdit,omitempty"`
}

// CompletionItemKind is the completion item kind values the client supports. When this
// property exists the client also guarantees that it will
// handle values outside its set gracefully and falls back
// to a default value when unknown.
//
// If this property is not present the client only supports
// the completion items kinds from `Text` to `Reference` as defined in
// the initial version of the protocol.
type CompletionItemKind int

const (
	// TextCompletion text completion kind.
	TextCompletion CompletionItemKind = 1
	// MethodCompletion method completion kind.
	MethodCompletion CompletionItemKind = 2
	// FunctionCompletion function completion kind.
	FunctionCompletion CompletionItemKind = 3
	// ConstructorCompletion constructor completion kind.
	ConstructorCompletion CompletionItemKind = 4
	// FieldCompletion field completion kind.
	FieldCompletion CompletionItemKind = 5
	// VariableCompletion variable completion kind.
	VariableCompletion CompletionItemKind = 6
	// ClassCompletion class completion kind.
	ClassCompletion CompletionItemKind = 7
	// InterfaceCompletion interface completion kind.
	InterfaceCompletion CompletionItemKind = 8
	// ModuleCompletion module completion kind.
	ModuleCompletion CompletionItemKind = 9
	// PropertyCompletion property completion kind.
	PropertyCompletion CompletionItemKind = 10
	// UnitCompletion unit completion kind.
	UnitCompletion CompletionItemKind = 11
	// ValueCompletion value completion kind.
	ValueCompletion CompletionItemKind = 12
	// EnumCompletion enum completion kind.
	EnumCompletion CompletionItemKind = 13
	// KeywordCompletion keyword completion kind.
	KeywordCompletion CompletionItemKind = 14
	// SnippetCompletion snippet completion kind.
	SnippetCompletion CompletionItemKind = 15
	// ColorCompletion color completion kind.
	ColorCompletion CompletionItemKind = 16
	// FileCompletion file completion kind.
	FileCompletion CompletionItemKind = 17
	// ReferenceCompletion reference completion kind.
	ReferenceCompletion CompletionItemKind = 18
	// FolderCompletion folder completion kind.
	FolderCompletion CompletionItemKind = 19
	// EnumMemberCompletion enum member completion kind.
	EnumMemberCompletion CompletionItemKind = 20
	// ConstantCompletion constant completion kind.
	ConstantCompletion CompletionItemKind = 21
	// StructCompletion struct completion kind.
	StructCompletion CompletionItemKind = 22
	// EventCompletion event completion kind.
	EventCompletion CompletionItemKind = 23
	// OperatorCompletion operator completion kind.
	OperatorCompletion CompletionItemKind = 24
	// TypeParameterCompletion type parameter completion kind.
	TypeParameterCompletion CompletionItemKind = 25
)

// String implements fmt.Stringer.
func (k CompletionItemKind) String() string {
	switch k {
	case TextCompletion:
		return "Text"
	case MethodCompletion:
		return "Method"
	case FunctionCompletion:
		return "Function"
	case ConstructorCompletion:
		return "Constructor"
	case FieldCompletion:
		return "Field"
	case VariableCompletion:
		return "Variable"
	case ClassCompletion:
		return "Class"
	case InterfaceCompletion:
		return "Interface"
	case ModuleCompletion:
		return "Module"
	case PropertyCompletion:
		return "Property"
	case UnitCompletion:
		return "Unit"
	case ValueCompletion:
		return "Value"
	case EnumCompletion:
		return "Enum"
	case KeywordCompletion:
		return "Keyword"
	case SnippetCompletion:
		return "Snippet"
	case ColorCompletion:
		return "Color"
	case FileCompletion:
		return "File"
	case ReferenceCompletion:
		return "Reference"
	case FolderCompletion:
		return "Folder"
	case EnumMemberCompletion:
		return "EnumMember"
	case ConstantCompletion:
		return "Constant"
	case StructCompletion:
		return "Struct"
	case EventCompletion:
		return "Event"
	case OperatorCompletion:
		return "Operator"
	case TypeParameterCompletion:
		return "TypeParameter"
	default:
		return strconv.FormatInt(int64(k), 64)
	}
}

// CompletionRegistrationOptions CompletionRegistration options.
type CompletionRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// TriggerCharacters most tools trigger completion request automatically without explicitly requesting
	// it using a keyboard shortcut (e.g. Ctrl+Space). Typically they do so when the user
	// starts to type an identifier. For example if the user types `c` in a JavaScript file
	// code complete will automatically pop up present `console` besides others as a
	// completion item. Characters that make up identifiers don't need to be listed here.
	//
	// If code complete should automatically be trigger on characters not being valid inside
	// an identifier (for example `.` in JavaScript) list them in `triggerCharacters`.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// ResolveProvider is the server provides support to resolve additional
	// information for a completion item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// Hover is the result of a hover request.
type Hover struct {
	// Contents is the hover's content
	Contents MarkupContent `json:"contents"`

	// Range an optional range is a range inside a text document
	// that is used to visualize a hover, e.g. by changing the background color.
	Range Range `json:"range,omitempty"`
}

// SignatureHelp signature help represents the signature of something
// callable. There can be multiple signature but only one
// active and only one active parameter.
type SignatureHelp struct {
	// Signatures one or more signatures.
	Signatures []SignatureInformation `json:"signatures"`

	// ActiveParameter is the active parameter of the active signature. If omitted or the value
	// lies outside the range of `signatures[activeSignature].parameters`
	// defaults to 0 if the active signature has parameters. If
	// the active signature has no parameters it is ignored.
	// In future version of the protocol this property might become
	// mandatory to better express the active parameter if the
	// active signature does have any.
	ActiveParameter float64 `json:"activeParameter,omitempty"`

	// ActiveSignature is the active signature. If omitted or the value lies outside the
	// range of `signatures` the value defaults to zero or is ignored if
	// `signatures.length === 0`. Whenever possible implementors should
	// make an active decision about the active signature and shouldn't
	// rely on a default value.
	// In future version of the protocol this property might become
	// mandatory to better express this.
	ActiveSignature float64 `json:"activeSignature,omitempty"`
}

// SignatureInformation is the client supports the following `SignatureInformation`
// specific properties.
type SignatureInformation struct {
	// DocumentationFormat is the client supports the follow content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

	// ParameterInformation client capabilities specific to parameter information.
	ParameterInformation *ParameterInformation `json:"parameterInformation,omitempty"`
}

// ParameterInformation represents a parameter of a callable-signature. A parameter can
// have a label and a doc-comment.
type ParameterInformation struct {
	// Label is the label of this parameter information.
	//
	// Either a string or an inclusive start and exclusive end offsets within its containing
	// signature label. (see SignatureInformation.label). The offsets are based on a UTF-16
	// string representation as `Position` and `Range` does.
	//
	// *Note*: a label of type string should be a substring of its containing signature label.
	// Its intended use case is to highlight the parameter label part in the `SignatureInformation.label`.
	Label string `json:"label"`

	// Documentation is the human-readable doc-comment of this parameter. Will be shown
	// in the UI but can be omitted.
	Documentation interface{} `json:"documentation,omitempty"`
}

// SignatureHelpRegistrationOptions SignatureHelp Registration options.
type SignatureHelpRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// TriggerCharacters is the characters that trigger signature help
	// automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
}

// ReferenceParams params of Find References Request
type ReferenceParams struct {
	TextDocumentPositionParams

	Context ReferenceContext `json:"context"`
}

// ReferenceContext context of ReferenceParams.
type ReferenceContext struct {
	// IncludeDeclaration include the declaration of the current symbol.
	IncludeDeclaration bool `json:"includeDeclaration"`
}

// DocumentHighlight a document highlight is a range inside a text document which deserves
// special attention. Usually a document highlight is visualized by changing
// the background color of its range.
type DocumentHighlight struct {
	// Range is the range this highlight applies to.
	Range Range `json:"range"`

	// Kind is the highlight kind, default is DocumentHighlightKind.Text.
	Kind DocumentHighlightKind `json:"kind,omitempty"`
}

// DocumentHighlightKind a document highlight kind.
type DocumentHighlightKind int

const (
	// Text a textual occurrence.
	Text DocumentHighlightKind = 1

	// Read read-access of a symbol, like reading a variable.
	Read DocumentHighlightKind = 2

	// Write write-access of a symbol, like writing to a variable.
	Write DocumentHighlightKind = 3
)

// String implements fmt.Stringer.
func (k DocumentHighlightKind) String() string {
	switch k {
	case Text:
		return "Text"
	case Read:
		return "Read"
	case Write:
		return "Write"
	default:
		return strconv.FormatInt(int64(k), 64)
	}
}

// DocumentSymbolParams params of Document Symbols Request.
type DocumentSymbolParams struct {
	// TextDocument is the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// SymbolKind specific capabilities for the `SymbolKind`.
// The symbol kind values the client supports. When this
// property exists the client also guarantees that it will
// handle values outside its set gracefully and falls back
// to a default value when unknown.
//
// If this property is not present the client only supports
// the symbol kinds from `File` to `Array` as defined in
// the initial version of the protocol.
type SymbolKind float64

const (
	// FileSymbol symbol of file.
	FileSymbol SymbolKind = 1
	// ModuleSymbol symbol of module.
	ModuleSymbol SymbolKind = 2
	// NamespaceSymbol symbol of namespace.
	NamespaceSymbol SymbolKind = 3
	// PackageSymbol symbol of package.
	PackageSymbol SymbolKind = 4
	// ClassSymbol symbol of class.
	ClassSymbol SymbolKind = 5
	// MethodSymbol symbol of method.
	MethodSymbol SymbolKind = 6
	// PropertySymbol symbol of property.
	PropertySymbol SymbolKind = 7
	// FieldSymbol symbol of field.
	FieldSymbol SymbolKind = 8
	// ConstructorSymbol symbol of constructor.
	ConstructorSymbol SymbolKind = 9
	// EnumSymbol symbol of enum.
	EnumSymbol SymbolKind = 10
	// InterfaceSymbol symbol of interface.
	InterfaceSymbol SymbolKind = 11
	// FunctionSymbol symbol of function.
	FunctionSymbol SymbolKind = 12
	// VariableSymbol symbol of variable.
	VariableSymbol SymbolKind = 13
	// ConstantSymbol symbol of constant.
	ConstantSymbol SymbolKind = 14
	// StringSymbol symbol of string.
	StringSymbol SymbolKind = 15
	// NumberSymbol symbol of number.
	NumberSymbol SymbolKind = 16
	// BooleanSymbol symbol of boolean.
	BooleanSymbol SymbolKind = 17
	// ArraySymbol symbol of array.
	ArraySymbol SymbolKind = 18
	// ObjectSymbol symbol of object.
	ObjectSymbol SymbolKind = 19
	// KeySymbol symbol of key.
	KeySymbol SymbolKind = 20
	// NullSymbol symbol of null.
	NullSymbol SymbolKind = 21
	// EnumMemberSymbol symbol of enum member.
	EnumMemberSymbol SymbolKind = 22
	// StructSymbol symbol of struct.
	StructSymbol SymbolKind = 23
	// EventSymbol symbol of event.
	EventSymbol SymbolKind = 24
	// OperatorSymbol symbol of operator.
	OperatorSymbol SymbolKind = 25
	// TypeParameterSymbol symbol of type parameter.
	TypeParameterSymbol SymbolKind = 26
)

// String implements fmt.Stringer.
func (k SymbolKind) String() string {
	switch k {
	case FileSymbol:
		return "File"
	case ModuleSymbol:
		return "Module"
	case NamespaceSymbol:
		return "Namespace"
	case PackageSymbol:
		return "Package"
	case ClassSymbol:
		return "Class"
	case MethodSymbol:
		return "Method"
	case PropertySymbol:
		return "Property"
	case FieldSymbol:
		return "Field"
	case ConstructorSymbol:
		return "Constructor"
	case EnumSymbol:
		return "Enum"
	case InterfaceSymbol:
		return "Interface"
	case FunctionSymbol:
		return "Function"
	case VariableSymbol:
		return "Variable"
	case ConstantSymbol:
		return "Constant"
	case StringSymbol:
		return "String"
	case NumberSymbol:
		return "Number"
	case BooleanSymbol:
		return "Boolean"
	case ArraySymbol:
		return "Array"
	case ObjectSymbol:
		return "Object"
	case KeySymbol:
		return "Key"
	case NullSymbol:
		return "Null"
	case EnumMemberSymbol:
		return "EnumMember"
	case StructSymbol:
		return "Struct"
	case EventSymbol:
		return "Event"
	case OperatorSymbol:
		return "Operator"
	case TypeParameterSymbol:
		return "TypeParameter"
	default:
		return strconv.FormatFloat(float64(k), 'f', -10, 64)
	}
}

// DocumentSymbol represents programming constructs like variables, classes, interfaces etc. that appear in a document. Document symbols can be
// hierarchical and they have two ranges: one that encloses its definition and one that points to its most interesting range,
// e.g. the range of an identifier.
type DocumentSymbol struct {
	// Name is the name of this symbol. Will be displayed in the user interface and therefore must not be
	// an empty string or a string only consisting of white spaces.
	Name string `json:"name"`

	// Detail is the more detail for this symbol, e.g the signature of a function.
	Detail string `json:"detail,omitempty"`

	// Kind is the kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Deprecated indicates if this symbol is deprecated.
	Deprecated bool `json:"deprecated,omitempty"`

	// Range is the range enclosing this symbol not including leading/trailing whitespace but everything else
	// like comments. This information is typically used to determine if the clients cursor is
	// inside the symbol to reveal in the symbol in the UI.
	Range Range `json:"range"`

	// SelectionRange is the range that should be selected and revealed when this symbol is being picked, e.g the name of a function.
	// Must be contained by the `range`.
	SelectionRange Range `json:"selectionRange"`

	// Children children of this symbol, e.g. properties of a class.
	Children []DocumentSymbol `json:"children,omitempty"`
}

// SymbolInformation represents information about programming constructs like variables, classes,
// interfaces etc.
type SymbolInformation struct {
	// Name is the name of this symbol.
	Name string `json:"name"`

	// Kind is the kind of this symbol.
	Kind float64 `json:"kind"`

	// Deprecated indicates if this symbol is deprecated.
	Deprecated bool `json:"deprecated,omitempty"`

	// Location is the location of this symbol. The location's range is used by a tool
	// to reveal the location in the editor. If the symbol is selected in the
	// tool the range's start information is used to position the cursor. So
	// the range usually spans more then the actual symbol's name and does
	// normally include things like visibility modifiers.
	//
	// The range doesn't have to denote a node range in the sense of a abstract
	// syntax tree. It can therefore not be used to re-construct a hierarchy of
	// the symbols.
	Location Location `json:"location"`

	// ContainerName is the name of the symbol containing this symbol. This information is for
	// user interface purposes (e.g. to render a qualifier in the user interface
	// if necessary). It can't be used to re-infer a hierarchy for the document
	// symbols.
	ContainerName string `json:"containerName,omitempty"`
}

// CodeActionParams params for the CodeActionRequest.
type CodeActionParams struct {
	// TextDocument is the document in which the command was invoked.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Context carrying additional information.
	Context CodeActionContext `json:"context"`

	// Range is the range for which the command was invoked.
	Range Range `json:"range"`
}

// CodeActionKind is the code action kind values the client supports. When this
// property exists the client also guarantees that it will
// handle values outside its set gracefully and falls back
// to a default value when unknown.
type CodeActionKind string

// A set of predefined code action kinds
const (
	// QuickFix base kind for quickfix actions: 'quickfix'
	QuickFix CodeActionKind = "quickfix"

	// Refactor base kind for refactoring actions: 'refactor'
	Refactor CodeActionKind = "refactor"

	// RefactorExtract base kind for refactoring extraction actions: 'refactor.extract'
	//
	// Example extract actions:
	//
	// - Extract method
	// - Extract function
	// - Extract variable
	// - Extract interface from class
	// - ...
	RefactorExtract CodeActionKind = "refactor.extract"

	// RefactorInline base kind for refactoring inline actions: 'refactor.inline'
	//
	// Example inline actions:
	//
	// - Inline function
	// - Inline variable
	// - Inline constant
	// - ...
	RefactorInline CodeActionKind = "refactor.inline"

	// RefactorRewrite base kind for refactoring rewrite actions: 'refactor.rewrite'
	//
	// Example rewrite actions:
	//
	// - Convert JavaScript function to class
	// - Add or remove parameter
	// - Encapsulate field
	// - Make method static
	// - Move method to base class
	// - ...
	RefactorRewrite CodeActionKind = "refactor.rewrite"

	// Source base kind for source actions: `source`
	//
	// Source code actions apply to the entire file.
	Source CodeActionKind = "source"

	// SourceOrganizeImports base kind for an organize imports source action: `source.organizeImports`
	SourceOrganizeImports CodeActionKind = "source.organizeImports"
)

// CodeActionContext contains additional diagnostic information about the context in which
// a code action is run.
type CodeActionContext struct {
	// Diagnostics is an array of diagnostics.
	Diagnostics []Diagnostic `json:"diagnostics"`

	// Only requested kind of actions to return.
	//
	// Actions not of this kind are filtered out by the client before being shown. So servers
	// can omit computing them.
	Only []CodeActionKind `json:"only,omitempty"`
}

// CodeAction capabilities specific to the `textDocument/codeAction`.
type CodeAction struct {
	// Title is a short, human-readable, title for this code action.
	Title string `json:"title"`

	// Kind is the kind of the code action.
	//
	// Used to filter code actions.
	Kind CodeActionKind `json:"kind,omitempty"`

	// Diagnostics is the diagnostics that this code action resolves.
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`

	// Edit is the workspace edit this code action performs.
	Edit *WorkspaceEdit `json:"edit,omitempty"`

	// Command is a command this code action executes. If a code action
	// provides an edit and a command, first the edit is
	// executed and then the command.
	Command *Command `json:"command,omitempty"`
}

// CodeActionRegistrationOptions CodeAction Registrationi options.
type CodeActionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeActionOptions
}

// CodeLensParams params of Code Lens Request.
type CodeLensParams struct {
	// TextDocument is the document to request code lens for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// CodeLens is a code lens represents a command that should be shown along with
// source text, like the number of references, a way to run tests, etc.
//
// A code lens is _unresolved_ when no command is associated to it. For performance
// reasons the creation of a code lens and resolving should be done in two stages.
type CodeLens struct {
	// Range is the range in which this code lens is valid. Should only span a single line.
	Range Range `json:"range"`

	// Command is the command this code lens represents.
	Command *Command `json:"command,omitempty"`

	// Data is a data entry field that is preserved on a code lens item between
	// a code lens and a code lens resolve request.
	Data interface{} `json:"data,omitempty"`
}

// CodeLensRegistrationOptions CodeLens Registration options.
type CodeLensRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// ResolveProvider code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DocumentLinkParams params of Document Link Request.
type DocumentLinkParams struct {
	// TextDocument is the document to provide document links for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// DocumentLink is a document link is a range in a text document that links to an internal or external resource, like another
// text document or a web site.
type DocumentLink struct {
	// Range is the range this link applies to.
	Range Range `json:"range"`

	// Target is the uri this link points to. If missing a resolve request is sent later.
	Target uri.URI `json:"target,omitempty"`

	// Data is a data entry field that is preserved on a document link between a
	// DocumentLinkRequest and a DocumentLinkResolveRequest.
	Data interface{} `json:"data,omitempty"`
}

// DocumentColorParams params of Document Color Request.
type DocumentColorParams struct {
	// TextDocument is the document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// ColorInformation response of Document Color Request.
type ColorInformation struct {
	// Range is the range in the document where this color appears.
	Range Range `json:"range"`

	// Color is the actual color value for this color range.
	Color Color `json:"color"`
}

// Color represents a color in RGBA space.
type Color struct {
	// Alpha is the alpha component of this color in the range [0-1].
	Alpha float64 `json:"alpha"`

	// Blue is the blue component of this color in the range [0-1].
	Blue float64 `json:"blue"`

	// Green is the green component of this color in the range [0-1].
	Green float64 `json:"green"`

	// Red is the red component of this color in the range [0-1].
	Red float64 `json:"red"`
}

// ColorPresentationParams params of Color Presentation Request.
type ColorPresentationParams struct {
	// TextDocument is the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Color is the color information to request presentations for.
	Color Color `json:"color"`

	// Range is the range where the color would be inserted. Serves as a context.
	Range Range `json:"range"`
}

// ColorPresentation response of Color Presentation Request.
type ColorPresentation struct {
	// Label is the label of this color presentation. It will be shown on the color
	// picker header. By default this is also the text that is inserted when selecting
	// this color presentation.
	Label string `json:"label"`

	// TextEdit an edit which is applied to a document when selecting
	// this presentation for the color.  When `falsy` the label is used.
	TextEdit *TextEdit `json:"textEdit,omitempty"`

	// AdditionalTextEdits an optional array of additional [text edits](#TextEdit) that are applied when
	// selecting this color presentation. Edits must not overlap with the main [edit](#ColorPresentation.textEdit) nor with themselves.
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
}

// DocumentFormattingParams params of Document Formatting Request.
type DocumentFormattingParams struct {
	// Options is the format options.
	Options FormattingOptions `json:"options"`

	// TextDocument is the document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// FormattingOptions value-object describing what options formatting should use.
type FormattingOptions struct {
	// InsertSpaces prefer spaces over tabs.
	InsertSpaces bool `json:"insertSpaces"`

	// TabSize size of a tab in spaces.
	TabSize float64 `json:"tabSize"`
}

// DocumentRangeFormattingParams params of Document Range Formatting Request.
type DocumentRangeFormattingParams struct {
	// TextDocument is the document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Range is the range to format
	Range Range `json:"range"`

	// Options is the format options.
	Options FormattingOptions `json:"options"`
}

// DocumentOnTypeFormattingParams params of Document on Type Formatting Request.
type DocumentOnTypeFormattingParams struct {
	// TextDocument is the document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position is the position at which this request was sent.
	Position Position `json:"position"`

	// Ch is the character that has been typed.
	Ch string `json:"ch"`

	// Options is the format options.
	Options FormattingOptions `json:"options"`
}

// DocumentOnTypeFormattingRegistrationOptions DocumentOnTypeFormatting Registration options.
type DocumentOnTypeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// FirstTriggerCharacter a character on which formatting should be triggered, like `}`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

	// MoreTriggerCharacter a More trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter"`
}

// RenameParams params of Rename Request.
type RenameParams struct {
	// TextDocument is the document to rename.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position is the position at which this request was sent.
	Position Position `json:"position"`

	// NewName is the new name of the symbol. If the given name is not valid the
	// request must return a [ResponseError](#ResponseError) with an
	// appropriate message set.
	NewName string `json:"newName"`
}

// RenameRegistrationOptions Rename Registration options.
type RenameRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// PrepareProvider is the renames should be checked and tested for validity before being executed.
	PrepareProvider bool `json:"prepareProvider,omitempty"`
}

// FoldingRangeParams params of Folding Range Request.
type FoldingRangeParams struct {
	// TextDocument is the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// FoldingRangeKind is the enum of known range kinds.
type FoldingRangeKind string

const (
	// CommentFoldingRange is the folding range for a comment.
	CommentFoldingRange FoldingRangeKind = "comment"

	// ImportsFoldingRange is the folding range for a imports or includes.
	ImportsFoldingRange FoldingRangeKind = "imports"

	// RegionFoldingRange is the folding range for a region (e.g. `#region`).
	RegionFoldingRange FoldingRangeKind = "region"
)

// FoldingRange capabilities specific to `textDocument/foldingRange` requests.
//
// Since 3.10.0
type FoldingRange struct {
	// StartLine is the zero-based line number from where the folded range starts.
	StartLine float64 `json:"startLine"`

	// StartCharacter is the zero-based character offset from where the folded range starts. If not defined, defaults to the length of the start line.
	StartCharacter float64 `json:"startCharacter,omitempty"`

	// EndLine is the zero-based line number where the folded range ends.
	EndLine float64 `json:"endLine"`

	// EndCharacter is the zero-based character offset before the folded range ends. If not defined, defaults to the length of the end line.
	EndCharacter float64 `json:"endCharacter,omitempty"`

	// Kind describes the kind of the folding range such as `comment' or 'region'. The kind
	// is used to categorize folding ranges and used by commands like 'Fold all comments'.
	// See FoldingRangeKind for an enumeration of standardized kinds.
	Kind FoldingRangeKind `json:"kind,omitempty"`
}
