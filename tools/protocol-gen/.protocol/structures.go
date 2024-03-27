package protocol

import (
	"go.lsp.dev/uri"
)

// TextDocumentIdentifier a literal to identify a text document in the client.
type TextDocumentIdentifier struct {
	// URI the text document's uri.
	URI DocumentURI `json:"uri"`
}

// Position position in a text document expressed as zero-based line and character offset. Prior to 3.17 the offsets were always based on a UTF-16 string representation. So a string of the form `a𐐀b` the character offset of the character `a` is 0, the character offset of `𐐀` is 1 and the character offset of b is 3 since `𐐀` is represented using two code units in UTF-16. Since 3.17 clients and servers can agree on a different string encoding representation (e.g. UTF-8). The client announces it's supported encoding via the client capability [`general.positionEncodings`](https://microsoft.github.io/language-server-protocol/specifications/specification-current/#clientCapabilities). The value is an array of position encodings the client supports, with decreasing preference (e.g. the encoding at index `0` is the most preferred one). To stay backwards compatible the only mandatory encoding is UTF-16 represented via the string `utf-16`. The server can pick one of the encodings offered by the client and signals that encoding back to the client via the initialize result's property [`capabilities.positionEncoding`](https://microsoft.github.io/language-server-protocol/specifications/specification-current/#serverCapabilities). If the string value `utf-16` is missing from the client's capability `general.positionEncodings` servers can safely assume that the client supports UTF-16. If the server omits the position encoding in its initialize result the encoding defaults to the string value `utf-16`. Implementation considerations: since the conversion from one encoding into another requires the content of the file / line the conversion is best done where the file is read which is usually on the server side. Positions are line end character agnostic. So you can not specify a position that denotes `\r|\n` or `\n|` where `|` represents the character offset.
//
//	3.17.0 - support for negotiated position encoding.
//
// @since 3.17.0 - support for negotiated position encoding.
type Position struct {
	// Line line position in a document (zero-based). If a line number is greater than the number of lines in a document, it defaults back to the number of lines in the document. If a line number is negative, it defaults to .
	//
	// @since 3.17.0 - support for negotiated position encoding.
	Line uint32 `json:"line"`

	// Character character offset on a line in a document (zero-based). The meaning of this offset is determined by the negotiated `PositionEncodingKind`. If the character value is greater than the line length it defaults back to the line length.
	//
	// @since 3.17.0 - support for negotiated position encoding.
	Character uint32 `json:"character"`
}

// TextDocumentPositionParams a parameter literal used in requests to pass a text document and a position inside that document.
type TextDocumentPositionParams struct {
	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// Position the position inside the text document.
	Position Position `json:"position"`
}

type ImplementationParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT
}

// Range a range in a text document expressed as (zero-based) start and end positions. If you want to specify a range that contains a line including the line ending character(s) then use an end position denoting the start of the next line. For example: ```ts {   start: { line: 5, character: 23 }   end : { line 6, character : 0 } } ```.
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

// TextDocumentRegistrationOptions general text document registration options.
type TextDocumentRegistrationOptions struct {
	// DocumentSelector a document selector to identify the scope of the registration. If set to null the document selector provided on the client side will be used.
	DocumentSelector any /* or */ `json:"document_selector"`
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

type TypeDefinitionParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT
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
	ScopeURI uri.URI `json:"scope_uri,omitempty"`

	// Section the configuration section asked for.
	Section string `json:"section,omitempty"`
}

// ConfigurationParams the parameters of a configuration request.
type ConfigurationParams struct {
	Items []ConfigurationItem `json:"items"`
}

// DocumentColorParams parameters for a DocumentColorRequest.
type DocumentColorParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"text_document"`
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
type ColorPresentationParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// Color the color to request presentations for.
	Color Color `json:"color"`

	// Range the range where the color would be inserted. Serves as a context.
	Range Range `json:"range"`
}

// TextEdit a text edit applicable to a text document.
type TextEdit struct {
	// Range the range of the text document to be manipulated. To insert text into a document create a range where start === end.
	Range Range `json:"range"`

	// NewText the string to be inserted. For delete operations use an empty string.
	NewText string `json:"new_text"`
}

type ColorPresentation struct {
	// Label the label of this color presentation. It will be shown on the color picker header. By default this is also the text that is inserted when selecting this color presentation.
	Label string `json:"label"`

	// TextEdit an TextEdit edit which is applied to a document when selecting this presentation for the color. When `falsy` the ColorPresentation.label label is used.
	TextEdit *TextEdit `json:"text_edit,omitempty"`

	// AdditionalTextEdits an optional array of additional TextEdit text edits that are applied when selecting this color presentation. Edits must not overlap with the main ColorPresentation.textEdit edit nor with themselves.
	AdditionalTextEdits []TextEdit `json:"additional_text_edits,omitempty"`
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"work_done_progress,omitempty"`
}

// FoldingRangeParams parameters for a FoldingRangeRequest.
type FoldingRangeParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"text_document"`
}

// FoldingRange represents a folding range. To be valid, start and end line must be bigger than zero and smaller than the number of lines in the document. Clients are free to ignore invalid ranges.
type FoldingRange struct {
	// StartLine the zero-based start line of the range to fold. The folded area starts after the line's last character. To be valid, the end must be zero or larger and smaller than the number of lines in the document.
	StartLine uint32 `json:"start_line"`

	// StartCharacter the zero-based character offset from where the folded range starts. If not defined, defaults to the length of the start line.
	StartCharacter uint32 `json:"start_character,omitempty"`

	// EndLine the zero-based end line of the range to fold. The folded area ends with the line's last character. To be valid, the end must be zero or larger and smaller than the number of lines in the document.
	EndLine uint32 `json:"end_line"`

	// EndCharacter the zero-based character offset before the folded range ends. If not defined, defaults to the length of the end line.
	EndCharacter uint32 `json:"end_character,omitempty"`

	// Kind describes the kind of the folding range such as 'comment' or 'region'. The kind is used to categorize folding ranges and used by commands like 'Fold all comments'. See FoldingRangeKind for an enumeration of standardized kinds.
	Kind *FoldingRangeKind `json:"kind,omitempty"`

	// CollapsedText the text that the client should show when the specified range is collapsed. If not defined or not supported by the client, a default will be chosen by the client.
	//
	//
	CollapsedText string `json:"collapsed_text,omitempty"`
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

type DeclarationParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT
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
type SelectionRangeParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"text_document"`

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

type WorkDoneProgressCreateParams[T ProgressToken] struct {
	// Token the token to be used to report progress.
	Token T `json:"token"`
}

type WorkDoneProgressCancelParams[T ProgressToken] struct {
	// Token the token to be used to report progress.
	Token T `json:"token"`
}

// CallHierarchyPrepareParams the parameter of a `textDocument/prepareCallHierarchy` request.
//
// @since 3.16.0
type CallHierarchyPrepareParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
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
	SelectionRange Range `json:"selection_range"`

	// Data a data entry field that is preserved between a call hierarchy prepare and incoming calls or outgoing calls requests.
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
type CallHierarchyIncomingCallsParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

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
	FromRanges []Range `json:"from_ranges"`
}

// CallHierarchyOutgoingCallsParams the parameter of a `callHierarchy/outgoingCalls` request.
//
// @since 3.16.0
type CallHierarchyOutgoingCallsParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

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

	// FromRanges the range at which this item is called. This is the range relative to the caller, e.g the item passed to CallHierarchyItemProvider.provideCallHierarchyOutgoingCalls `provideCallHierarchyOutgoingCalls` and not CallHierarchyOutgoingCall.to `this.to`.
	//
	// @since 3.16.0
	FromRanges []Range `json:"from_ranges"`
}

// SemanticTokensParams.
//
// @since 3.16.0
type SemanticTokensParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the text document.
	//
	// @since 3.16.0
	TextDocument TextDocumentIdentifier `json:"text_document"`
}

// SemanticTokens.
//
// @since 3.16.0
type SemanticTokens struct {
	// ResultID an optional result id. If provided and clients support delta updating the client will include the result id in the next semantic token request. A server can then instead of computing all semantic tokens again simply send a delta.
	//
	// @since 3.16.0
	ResultID string `json:"result_id,omitempty"`

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
	TokenTypes []string `json:"token_types"`

	// TokenModifiers the token modifiers a server uses.
	//
	// @since 3.16.0
	TokenModifiers []string `json:"token_modifiers"`
}

// SemanticTokensFullDelta semantic tokens options to support deltas for full documents
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type SemanticTokensFullDelta struct {
	// Delta the server supports deltas for full documents.
	//
	// @since 3.18.0 proposed
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
	Range any/* or */ `json:"range,omitempty"`

	// Full server supports providing semantic tokens for a full document.
	//
	// @since 3.16.0
	Full any/* or */ `json:"full,omitempty"`
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
type SemanticTokensDeltaParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the text document.
	//
	// @since 3.16.0
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// PreviousResultID the result id of a previous response. The result Id can either point to a full response or a delta response depending on what was received last.
	//
	// @since 3.16.0
	PreviousResultID string `json:"previous_result_id"`
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
	DeleteCount uint32 `json:"delete_count"`

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
	ResultID string `json:"result_id,omitempty"`

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
type SemanticTokensRangeParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the text document.
	//
	// @since 3.16.0
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// Range the range the semantic tokens are requested for.
	//
	// @since 3.16.0
	Range Range `json:"range"`
}

// ShowDocumentParams params to show a resource in the UI.
//
// @since 3.16.0
type ShowDocumentParams struct {
	// URI the uri to show.
	//
	// @since 3.16.0
	URI uri.URI `json:"uri"`

	// External indicates to show the resource in an external program. To show, for example, `https://code.visualstudio.com/` in the default WEB browser set `external` to `true`.
	//
	// @since 3.16.0
	External bool `json:"external,omitempty"`

	// TakeFocus an optional property to indicate whether the editor showing the document should take focus or not. Clients might ignore this property if an external program is started.
	//
	// @since 3.16.0
	TakeFocus bool `json:"take_focus,omitempty"`

	// Selection an optional selection range if the document is a text document. Clients might ignore the property if an external program is started or the file is not a text file.
	//
	// @since 3.16.0
	Selection *Range `json:"selection,omitempty"`
}

// ShowDocumentResult the result of a showDocument request.
//
// @since 3.16.0
type ShowDocumentResult struct {
	// Success a boolean indicating if the show was successful.
	//
	// @since 3.16.0
	Success bool `json:"success"`
}

type LinkedEditingRangeParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
}

// LinkedEditingRanges the result of a linked editing range request.
//
// @since 3.16.0
type LinkedEditingRanges struct {
	// Ranges a list of ranges that can be edited together. The ranges must have identical length and contain identical text content. The ranges cannot overlap.
	//
	// @since 3.16.0
	Ranges []Range `json:"ranges"`

	// WordPattern an optional word pattern (regular expression) that describes valid contents for the given ranges. If no pattern is provided, the client configuration's word pattern will be used.
	//
	// @since 3.16.0
	WordPattern string `json:"word_pattern,omitempty"`
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
	//
	//
	AnnotationID *ChangeAnnotationIdentifier `json:"annotation_id,omitempty"`
}

// DeleteFileOptions delete file options.
type DeleteFileOptions struct {
	// Recursive delete the content recursively if a folder is denoted.
	Recursive bool `json:"recursive,omitempty"`

	// IgnoreIfNotExists ignore the operation if the file doesn't exist.
	IgnoreIfNotExists bool `json:"ignore_if_not_exists,omitempty"`
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
	IgnoreIfExists bool `json:"ignore_if_exists,omitempty"`
}

// RenameFile rename file operation.
type RenameFile struct {
	// extends
	ResourceOperation

	// OldURI the old (existing) location.
	OldURI DocumentURI `json:"old_uri"`

	// NewURI the new location.
	NewURI DocumentURI `json:"new_uri"`

	// Options rename options.
	Options *RenameFileOptions `json:"options,omitempty"`
}

// CreateFileOptions options to create a file.
type CreateFileOptions struct {
	// Overwrite overwrite existing file. Overwrite wins over `ignoreIfExists`.
	Overwrite bool `json:"overwrite,omitempty"`

	// IgnoreIfExists ignore if exists.
	IgnoreIfExists bool `json:"ignore_if_exists,omitempty"`
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

// OptionalVersionedTextDocumentIdentifier a text document identifier to optionally denote a specific version of a text document.
type OptionalVersionedTextDocumentIdentifier struct {
	// extends
	TextDocumentIdentifier

	// Version the version number of this document. If a versioned text document identifier is sent from the server to the client and the file is not open in the editor (the server has not received an open notification before) the server can send `null` to indicate that the version is unknown and the content on disk is the truth (as specified with document content ownership).
	Version any/* or */ `json:"version"`
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
	AnnotationID ChangeAnnotationIdentifier `json:"annotation_id"`
}

// TextDocumentEdit describes textual changes on a text document. A TextDocumentEdit describes all changes on a document version Si and after they are applied move the document to version Si+1. So the creator of a TextDocumentEdit doesn't need to sort the array of edits or do any kind of ordering. However the edits must be non overlapping.
type TextDocumentEdit struct {
	// TextDocument the text document to change.
	TextDocument OptionalVersionedTextDocumentIdentifier `json:"text_document"`

	// Edits the edits to be applied.
	//
	//  3.16.0 - support for AnnotatedTextEdit. This is guarded using a client capability.
	Edits any/* or */ `json:"edits"`
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
	NeedsConfirmation bool `json:"needs_confirmation,omitempty"`

	// Description a human-readable string which is rendered less prominent in the user interface.
	//
	// @since 3.16.0
	Description string `json:"description,omitempty"`
}

// WorkspaceEdit a workspace edit represents changes to many resources managed in the workspace. The edit should either provide `changes` or `documentChanges`. If documentChanges are present they are preferred over `changes` if the client can handle versioned document edits. Since version 3.13.0 a workspace edit can contain resource operations as well. If resource operations are present clients need to execute the operations in the order in which they are provided. So a workspace edit for example can consist of the following two changes: (1) a create file a.txt and (2) a text document edit which insert text into file a.txt. An invalid sequence (e.g. (1) delete file a.txt and (2) insert text into file a.txt) will cause failure of the operation. How the client recovers from the failure is described by the client capability: `workspace.workspaceEdit.failureHandling`.
type WorkspaceEdit struct {
	// Changes holds changes to existing resources.
	Changes map[string]any `json:"changes,omitempty"`

	// DocumentChanges depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes are either an array of `TextDocumentEdit`s to express changes to n different text documents where each text document edit addresses a specific version of a text document. Or it can contain above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations. Whether a client supports versioned document edits is expressed via `workspace.workspaceEdit.documentChanges` client capability. If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then only plain `TextEdit`s using the `changes` property are supported.
	DocumentChanges any/* or */ `json:"document_changes,omitempty"`

	// ChangeAnnotations a map of change annotations that can be referenced in `AnnotatedTextEdit`s or create, rename and delete file / folder operations. Whether clients honor this property depends on the client capability `workspace.changeAnnotationSupport`.
	//
	//
	ChangeAnnotations map[string]any `json:"change_annotations,omitempty"`
}

// FileOperationPatternOptions matching options for the file operation pattern.
//
// @since 3.16.0
type FileOperationPatternOptions struct {
	// IgnoreCase the pattern should be matched ignoring casing.
	//
	// @since 3.16.0
	IgnoreCase bool `json:"ignore_case,omitempty"`
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
	Matches *FileOperationPatternKind `json:"matches,omitempty"`

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
	OldURI string `json:"old_uri"`

	// NewURI a file:// URI for the new location of the file/folder being renamed.
	//
	// @since 3.16.0
	NewURI string `json:"new_uri"`
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

type MonikerParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT
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
	Kind *MonikerKind `json:"kind,omitempty"`
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
type TypeHierarchyPrepareParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
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
	SelectionRange Range `json:"selection_range"`

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
type TypeHierarchySupertypesParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// @since 3.17.0
	Item TypeHierarchyItem `json:"item"`
}

// TypeHierarchySubtypesParams the parameter of a `typeHierarchy/subtypes` request.
//
// @since 3.17.0
type TypeHierarchySubtypesParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

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
	FrameID int32 `json:"frame_id"`

	// StoppedLocation the document range where execution has stopped. Typically the end position of the range denotes the line where the inline values are shown.
	//
	// @since 3.17.0
	StoppedLocation Range `json:"stopped_location"`
}

// InlineValueParams a parameter literal used in inline value requests.
//
// @since 3.17.0
type InlineValueParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT

	// TextDocument the text document.
	//
	// @since 3.17.0
	TextDocument TextDocumentIdentifier `json:"text_document"`

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
type InlayHintParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT

	// TextDocument the text document.
	//
	// @since 3.17.0
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// Range the document range for which inlay hints should be computed.
	//
	// @since 3.17.0
	Range Range `json:"range"`
}

// MarkupContent a `MarkupContent` literal represents a string value which content is interpreted base on its kind flag. Currently the protocol supports `plaintext` and `markdown` as markup kinds. If the kind is `markdown` then the value can contain fenced code blocks like in GitHub issues. See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting Here is an example how such a string can be constructed using JavaScript / TypeScript: ```ts let markdown: MarkdownContent = { kind: MarkupKind.Markdown, value: [  '# Header',  'Some text',  '```typescript',  'someCode();',  '```' ].join('\n') }; ``` *Please Note* that clients might sanitize the return markdown. A client could decide to remove HTML from the markdown to avoid script execution.
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

	// Tooltip an optional tooltip.
	//
	//  3.18.0
	//
	// Proposed in:.
	Tooltip string `json:"tooltip,omitempty"`

	// Command the identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments arguments that the command handler should be invoked with.
	Arguments []any `json:"arguments,omitempty"`
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
	Tooltip any/* or */ `json:"tooltip,omitempty"`

	// Location an optional source code location that represents this label part. The editor will use this location for the hover and for code navigation features: This part will become a clickable link that resolves to the definition of the symbol at the given location (not necessarily the location itself), it shows the hover that shows at the given location, and it shows a context menu with further code navigation commands. Depending on the client capability `inlayHint.resolveSupport` clients might resolve this property late using the resolve request.
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
	// Position the position of this hint. If multiple hints have the same position, they will be shown in the order they appear in the response.
	//
	// @since 3.17.0
	Position Position `json:"position"`

	// Label the label of this hint. A human readable string or an array of InlayHintLabelPart label parts. *Note* that neither the string nor the label part can be empty.
	//
	// @since 3.17.0
	Label any/* or */ `json:"label"`

	// Kind the kind of this hint. Can be omitted in which case the client should fall back to a reasonable default.
	//
	// @since 3.17.0
	Kind *InlayHintKind `json:"kind,omitempty"`

	// TextEdits optional text edits that are performed when accepting this inlay hint. *Note* that edits are expected to change the document so that the inlay hint (or its nearest variant) is now part of the document and the inlay hint itself is now obsolete.
	//
	// @since 3.17.0
	TextEdits []TextEdit `json:"text_edits,omitempty"`

	// Tooltip the tooltip text when you hover over this item.
	//
	// @since 3.17.0
	Tooltip any/* or */ `json:"tooltip,omitempty"`

	// PaddingLeft render padding before the hint. Note: Padding should use the editor's background color, not the background color of the hint itself. That means padding can be used to visually align/separate an inlay hint.
	//
	// @since 3.17.0
	PaddingLeft bool `json:"padding_left,omitempty"`

	// PaddingRight render padding after the hint. Note: Padding should use the editor's background color, not the background color of the hint itself. That means padding can be used to visually align/separate an inlay hint.
	//
	// @since 3.17.0
	PaddingRight bool `json:"padding_right,omitempty"`

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
	ResolveProvider bool `json:"resolve_provider,omitempty"`
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
type DocumentDiagnosticParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the text document.
	//
	// @since 3.17.0
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// Identifier the additional identifier provided during registration.
	//
	// @since 3.17.0
	Identifier string `json:"identifier,omitempty"`

	// PreviousResultID the result id of a previous response if provided.
	//
	// @since 3.17.0
	PreviousResultID string `json:"previous_result_id,omitempty"`
}

// UnchangedDocumentDiagnosticReport a diagnostic report indicating that the last returned report is still accurate.
//
// @since 3.17.0
type UnchangedDocumentDiagnosticReport struct {
	// ResultID a result id which will be sent on the next diagnostic request for the same document.
	//
	// @since 3.17.0
	ResultID string `json:"result_id"`
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

// DiagnosticRelatedInformation represents a related message and source code location for a diagnostic. This should be used to point to code locations that cause or related to a diagnostics, e.g when duplicating a symbol in a scope.
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

	// Severity the diagnostic's severity. Can be omitted. If omitted it is up to the client to interpret diagnostics as error, warning, info or hint.
	Severity *DiagnosticSeverity `json:"severity,omitempty"`

	// Code the diagnostic's code, which usually appear in the user interface.
	Code any/* or */ `json:"code,omitempty"`

	// CodeDescription an optional property to describe the error code. Requires the code field (above) to be present/not null.
	//
	//
	CodeDescription *CodeDescription `json:"code_description,omitempty"`

	// Source a human-readable string describing the source of this diagnostic, e.g. 'typescript' or 'super lint'. It usually appears in the user interface.
	Source string `json:"source,omitempty"`

	// Message the diagnostic's message. It usually appears in the user interface.
	Message string `json:"message"`

	// Tags additional metadata about the diagnostic.
	//
	//
	Tags []DiagnosticTag `json:"tags,omitempty"`

	// RelatedInformation an array of related diagnostic information, e.g. when symbol-names within a scope collide all definitions can be marked via this property.
	RelatedInformation []DiagnosticRelatedInformation `json:"related_information,omitempty"`

	// Data a data entry field that is preserved between a `textDocument/publishDiagnostics` notification and `textDocument/codeAction` request.
	//
	//
	Data any `json:"data,omitempty"`
}

// FullDocumentDiagnosticReport a diagnostic report with a full set of problems.
//
// @since 3.17.0
type FullDocumentDiagnosticReport struct {
	// ResultID an optional result id. If provided it will be sent on the next diagnostic request for the same document.
	//
	// @since 3.17.0
	ResultID string `json:"result_id,omitempty"`

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
	RelatedDocuments map[string]any `json:"related_documents"`
}

// DiagnosticServerCancellationData cancellation data returned from a diagnostic request.
//
// @since 3.17.0
type DiagnosticServerCancellationData struct {
	// @since 3.17.0
	RetriggerRequest bool `json:"retrigger_request"`
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

	// InterFileDependencies whether the language has inter file dependencies meaning that editing code in one file can result in a different diagnostic set in another file. Inter file dependencies are common for most programming languages and typically uncommon for linters.
	//
	// @since 3.17.0
	InterFileDependencies bool `json:"inter_file_dependencies"`

	// WorkspaceDiagnostics the server provides support for workspace diagnostics as well.
	//
	// @since 3.17.0
	WorkspaceDiagnostics bool `json:"workspace_diagnostics"`
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
type WorkspaceDiagnosticParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// Identifier the additional identifier provided during registration.
	//
	// @since 3.17.0
	Identifier string `json:"identifier,omitempty"`

	// PreviousResultIDS the currently known diagnostic reports with their previous result ids.
	//
	// @since 3.17.0
	PreviousResultIDS []PreviousResultID `json:"previous_result_ids"`
}

// WorkspaceDiagnosticReport a workspace diagnostic report.
//
// @since 3.17.0
type WorkspaceDiagnosticReport[T WorkspaceDocumentDiagnosticReport] struct {
	// @since 3.17.0
	Items []T `json:"items"`
}

// WorkspaceDiagnosticReportPartialResult a partial result for a workspace diagnostic report.
//
// @since 3.17.0
type WorkspaceDiagnosticReportPartialResult[T WorkspaceDocumentDiagnosticReport] struct {
	// @since 3.17.0
	Items []T `json:"items"`
}

type ExecutionSummary struct {
	// ExecutionOrder a strict monotonically increasing value indicating the execution order of a cell inside a notebook.
	ExecutionOrder uint32 `json:"execution_order"`

	// Success whether the execution was successful or not if known by the client.
	Success bool `json:"success,omitempty"`
}

// NotebookCell a notebook cell. A cell's document URI must be unique across ALL notebook cells and can therefore be used to uniquely identify a notebook cell or the cell's text document.
//
// @since 3.17.0
type NotebookCell struct {
	// Kind the cell's kind.
	//
	// @since 3.17.0
	Kind NotebookCellKind `json:"kind"`

	// Document the URI of the cell's text document content.
	//
	// @since 3.17.0
	Document DocumentURI `json:"document"`

	// Metadata additional metadata stored with the cell. Note: should always be an object literal (e.g. LSPObject).
	//
	// @since 3.17.0
	Metadata any `json:"metadata,omitempty"`

	// ExecutionSummary additional execution summary information if supported by the client.
	//
	// @since 3.17.0
	ExecutionSummary *ExecutionSummary `json:"execution_summary,omitempty"`
}

// NotebookDocument a notebook document.
//
// @since 3.17.0
type NotebookDocument struct {
	// URI the notebook document's uri.
	//
	// @since 3.17.0
	URI uri.URI `json:"uri"`

	// NotebookType the type of the notebook.
	//
	// @since 3.17.0
	NotebookType string `json:"notebook_type"`

	// Version the version number of this document (it will increase after each change, including undo/redo).
	//
	// @since 3.17.0
	Version int32 `json:"version"`

	// Metadata additional metadata stored with the notebook document. Note: should always be an object literal (e.g. LSPObject).
	//
	// @since 3.17.0
	Metadata any `json:"metadata,omitempty"`

	// Cells the cells of a notebook.
	//
	// @since 3.17.0
	Cells []NotebookCell `json:"cells"`
}

// TextDocumentItem an item to transfer a text document from the client to the server.
type TextDocumentItem struct {
	// URI the text document's uri.
	URI DocumentURI `json:"uri"`

	// LanguageID the text document's language identifier.
	LanguageID LanguageKind `json:"language_id"`

	// Version the version number of this document (it will increase after each change, including undo/redo).
	Version int32 `json:"version"`

	// Text the content of the opened text document.
	Text string `json:"text"`
}

// DidOpenNotebookDocumentParams the params sent in an open notebook document notification.
//
// @since 3.17.0
type DidOpenNotebookDocumentParams struct {
	// NotebookDocument the notebook document that got opened.
	//
	// @since 3.17.0
	NotebookDocument NotebookDocument `json:"notebook_document"`

	// CellTextDocuments the text documents that represent the content of a notebook cell.
	//
	// @since 3.17.0
	CellTextDocuments []TextDocumentItem `json:"cell_text_documents"`
}

// NotebookCellLanguage.
//
// @since 3.18.0 proposed
type NotebookCellLanguage struct {
	// @since 3.18.0 proposed
	Language string `json:"language"`
}

// NotebookDocumentFilterWithCells.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterWithCells struct {
	// Notebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
	//
	// @since 3.18.0 proposed
	Notebook any/* or */ `json:"notebook,omitempty"`

	// Cells the cells of the matching notebook to be synced.
	//
	// @since 3.18.0 proposed
	Cells []NotebookCellLanguage `json:"cells"`
}

// NotebookDocumentFilterWithNotebook.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterWithNotebook struct {
	// Notebook the notebook to be synced If a string value is provided it matches against the notebook type. '*' matches every notebook.
	//
	// @since 3.18.0 proposed
	Notebook any/* or */ `json:"notebook"`

	// Cells the cells of the matching notebook to be synced.
	//
	// @since 3.18.0 proposed
	Cells []NotebookCellLanguage `json:"cells,omitempty"`
}

// NotebookDocumentSyncOptions options specific to a notebook plus its cells to be synced to the server. If a selector provides a notebook document filter but no cell selector all cells of a matching notebook document will be synced. If a selector provides no notebook document filter but only a cell selector all notebook document that contain at least one matching cell will be synced.
//
// @since 3.17.0
type NotebookDocumentSyncOptions struct {
	// NotebookSelector the notebooks to be synced.
	//
	// @since 3.17.0
	NotebookSelector any/* or */ `json:"notebook_selector"`

	// Save whether save notification should be forwarded to the server. Will only be honored if mode === `notebook`.
	//
	// @since 3.17.0
	Save bool `json:"save,omitempty"`
}

// NotebookDocumentSyncRegistrationOptions registration options specific to a notebook.
//
// @since 3.17.0
type NotebookDocumentSyncRegistrationOptions struct {
	// extends
	NotebookDocumentSyncOptions
	// mixins
	StaticRegistrationOptions
}

// VersionedNotebookDocumentIdentifier a versioned notebook document identifier.
//
// @since 3.17.0
type VersionedNotebookDocumentIdentifier struct {
	// Version the version number of this notebook document.
	//
	// @since 3.17.0
	Version int32 `json:"version"`

	// URI the notebook document's uri.
	//
	// @since 3.17.0
	URI uri.URI `json:"uri"`
}

// NotebookCellArrayChange a change describing how to move a `NotebookCell` array from state S to S'.
//
// @since 3.17.0
type NotebookCellArrayChange struct {
	// Start the start oftest of the cell that changed.
	//
	// @since 3.17.0
	Start uint32 `json:"start"`

	// DeleteCount the deleted cells.
	//
	// @since 3.17.0
	DeleteCount uint32 `json:"delete_count"`

	// Cells the new cells, if any.
	//
	// @since 3.17.0
	Cells []NotebookCell `json:"cells,omitempty"`
}

// NotebookDocumentCellChangeStructure structural changes to cells in a notebook document.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type NotebookDocumentCellChangeStructure struct {
	// Array the change to the cell array.
	//
	// @since 3.18.0 proposed
	Array NotebookCellArrayChange `json:"array"`

	// DidOpen additional opened cell text documents.
	//
	// @since 3.18.0 proposed
	DidOpen []TextDocumentItem `json:"did_open,omitempty"`

	// DidClose additional closed cell text documents.
	//
	// @since 3.18.0 proposed
	DidClose []TextDocumentIdentifier `json:"did_close,omitempty"`
}

// VersionedTextDocumentIdentifier a text document identifier to denote a specific version of a text document.
type VersionedTextDocumentIdentifier struct {
	// extends
	TextDocumentIdentifier

	// Version the version number of this document.
	Version int32 `json:"version"`
}

// NotebookDocumentCellContentChanges content changes to a cell in a notebook document.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type NotebookDocumentCellContentChanges[T TextDocumentContentChangeEvent] struct {
	// @since 3.18.0 proposed
	Document VersionedTextDocumentIdentifier `json:"document"`

	// @since 3.18.0 proposed
	Changes []T `json:"changes"`
}

// NotebookDocumentCellChanges cell changes to a notebook document.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type NotebookDocumentCellChanges[T TextDocumentContentChangeEvent, TT NotebookDocumentCellContentChanges[T]] struct {
	// Structure changes to the cell structure to add or remove cells.
	//
	// @since 3.18.0 proposed
	Structure *NotebookDocumentCellChangeStructure `json:"structure,omitempty"`

	// Data changes to notebook cells properties like its kind, execution summary or metadata.
	//
	// @since 3.18.0 proposed
	Data []NotebookCell `json:"data,omitempty"`

	// TextContent changes to the text content of notebook cells.
	//
	// @since 3.18.0 proposed
	TextContent []TT `json:"text_content,omitempty"`
}

// NotebookDocumentChangeEvent a change event for a notebook document.
//
// @since 3.17.0
type NotebookDocumentChangeEvent[T TextDocumentContentChangeEvent, TT NotebookDocumentCellContentChanges[T]] struct {
	// Metadata the changed meta data if any. Note: should always be an object literal (e.g. LSPObject).
	//
	// @since 3.17.0
	Metadata any `json:"metadata,omitempty"`

	// Cells changes to cells.
	//
	// @since 3.17.0
	Cells *NotebookDocumentCellChanges[T, TT] `json:"cells,omitempty"`
}

// DidChangeNotebookDocumentParams the params sent in a change notebook document notification.
//
// @since 3.17.0
type DidChangeNotebookDocumentParams[T TextDocumentContentChangeEvent, TT NotebookDocumentCellContentChanges[T]] struct {
	// NotebookDocument the notebook document that did change. The version number points to the version after all provided changes have been applied. If only the text document content of a cell changes the notebook version doesn't necessarily have to change.
	//
	// @since 3.17.0
	NotebookDocument VersionedNotebookDocumentIdentifier `json:"notebook_document"`

	// Change the actual changes to the notebook document. The changes describe single state changes to the notebook document. So if there are two changes c1 (at array index 0) and c2 (at array index 1) for a notebook in state S then c1 moves the notebook from S to S' and c2 from S' to S''. So c1 is computed on the state S and c2 is computed on the state S'. To mirror the content of a notebook using change events use the following approach: - start with the same initial content - apply the 'notebookDocument/didChange' notifications in the order you receive them. - apply the `NotebookChangeEvent`s in a single notification in the order  you receive them.
	//
	// @since 3.17.0
	Change NotebookDocumentChangeEvent[T, TT] `json:"change"`
}

// NotebookDocumentIdentifier a literal to identify a notebook document in the client.
//
// @since 3.17.0
type NotebookDocumentIdentifier struct {
	// URI the notebook document's uri.
	//
	// @since 3.17.0
	URI uri.URI `json:"uri"`
}

// DidSaveNotebookDocumentParams the params sent in a save notebook document notification.
//
// @since 3.17.0
type DidSaveNotebookDocumentParams struct {
	// NotebookDocument the notebook document that got saved.
	//
	// @since 3.17.0
	NotebookDocument NotebookDocumentIdentifier `json:"notebook_document"`
}

// DidCloseNotebookDocumentParams the params sent in a close notebook document notification.
//
// @since 3.17.0
type DidCloseNotebookDocumentParams struct {
	// NotebookDocument the notebook document that got closed.
	//
	// @since 3.17.0
	NotebookDocument NotebookDocumentIdentifier `json:"notebook_document"`

	// CellTextDocuments the text documents that represent the content of a notebook cell that got closed.
	//
	// @since 3.17.0
	CellTextDocuments []TextDocumentIdentifier `json:"cell_text_documents"`
}

// SelectedCompletionInfo describes the currently selected completion item.
//
//	3.18.0
//
// Proposed in:.
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

// InlineCompletionContext provides information about the context in which an inline completion was requested.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type InlineCompletionContext struct {
	// TriggerKind describes how the inline completion was triggered.
	//
	// @since 3.18.0 proposed
	TriggerKind InlineCompletionTriggerKind `json:"trigger_kind"`

	// SelectedCompletionInfo provides information about the currently selected item in the autocomplete widget if it is visible.
	//
	// @since 3.18.0 proposed
	SelectedCompletionInfo *SelectedCompletionInfo `json:"selected_completion_info,omitempty"`
}

// InlineCompletionParams a parameter literal used in inline completion requests.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type InlineCompletionParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT

	// Context additional information about the context in which inline completions were requested.
	//
	// @since 3.18.0 proposed
	Context InlineCompletionContext `json:"context"`
}

// StringValue a string value used as a snippet is a template which allows to insert text and to control the editor cursor when insertion happens. A snippet can define tab stops and placeholders with `$1`, `$2` and `${3:foo}`. `$0` defines the final tab stop, it defaults to the end of the snippet. Variables are defined with `$name` and `${name:default value}`.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type StringValue struct {
	// Value the snippet string.
	//
	// @since 3.18.0 proposed
	Value string `json:"value"`
}

// InlineCompletionItem an inline completion item represents a text snippet that is proposed inline to complete text that is being typed.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type InlineCompletionItem struct {
	// InsertText the text to replace the range with. Must be set.
	//
	// @since 3.18.0 proposed
	InsertText any/* or */ `json:"insert_text"`

	// FilterText a text that is used to decide if this inline completion should be shown. When `falsy` the InlineCompletionItem.insertText is used.
	//
	// @since 3.18.0 proposed
	FilterText string `json:"filter_text,omitempty"`

	// Range the range to replace. Must begin and end on the same line.
	//
	// @since 3.18.0 proposed
	Range *Range `json:"range,omitempty"`

	// Command an optional Command that is executed *after* inserting this completion.
	//
	// @since 3.18.0 proposed
	Command *Command `json:"command,omitempty"`
}

// InlineCompletionList represents a collection of InlineCompletionItem inline completion items to be presented in the editor.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type InlineCompletionList struct {
	// Items the inline completion items.
	//
	// @since 3.18.0 proposed
	Items []InlineCompletionItem `json:"items"`
}

// InlineCompletionOptions inline completion options used during static registration.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type InlineCompletionOptions struct {
	// mixins
	WorkDoneProgressOptions
}

// InlineCompletionRegistrationOptions inline completion options used during static or dynamic registration.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type InlineCompletionRegistrationOptions struct {
	// extends
	InlineCompletionOptions
	TextDocumentRegistrationOptions
	// mixins
	StaticRegistrationOptions
}

// Registration general parameters to register for a notification or to register a provider.
type Registration struct {
	// ID the id used to register the request. The id can be used to deregister the request again.
	ID string `json:"id"`

	// Method the method / capability to register for.
	Method string `json:"method"`

	// RegisterOptions options necessary for the registration.
	RegisterOptions any `json:"register_options,omitempty"`
}

type RegistrationParams struct {
	Registrations []Registration `json:"registrations"`
}

// Unregistration general parameters to unregister a request or notification.
type Unregistration struct {
	// ID the id used to unregister the request or notification. Usually an id provided during the register request.
	ID string `json:"id"`

	// Method the method to unregister for.
	Method string `json:"method"`
}

type UnregistrationParams struct {
	Unregisterations []Unregistration `json:"unregisterations"`
}

// ClientInfo information about the client
//
//	3.15.0
//
//	3.18.0 ClientInfo type name added.
//
// Proposed in:.
//
// @since 3.18.0 ClientInfo type name added. proposed
type ClientInfo struct {
	// Name the name of the client as defined by the client.
	//
	// @since 3.18.0 ClientInfo type name added. proposed
	Name string `json:"name"`

	// Version the client's version as defined by the client.
	//
	// @since 3.18.0 ClientInfo type name added. proposed
	Version string `json:"version,omitempty"`
}

// ChangeAnnotationsSupportOptions.
//
// @since 3.18.0 proposed
type ChangeAnnotationsSupportOptions struct {
	// GroupsOnLabel whether the client groups edits with equal labels into tree nodes, for instance all edits labelled with "Changes in Strings" would be a tree node.
	//
	// @since 3.18.0 proposed
	GroupsOnLabel bool `json:"groups_on_label,omitempty"`
}

type WorkspaceEditClientCapabilities struct {
	// DocumentChanges the client supports versioned document changes in `WorkspaceEdit`s.
	DocumentChanges bool `json:"document_changes,omitempty"`

	// ResourceOperations the resource operations the client supports. Clients should at least support 'create', 'rename' and 'delete' files and folders.
	//
	//
	ResourceOperations []ResourceOperationKind `json:"resource_operations,omitempty"`

	// FailureHandling the failure handling strategy of a client if applying the workspace edit fails.
	//
	//
	FailureHandling *FailureHandlingKind `json:"failure_handling,omitempty"`

	// NormalizesLineEndings whether the client normalizes line endings to the client specific setting. If set to `true` the client will normalize line ending characters in a workspace edit to the client-specified new line character.
	//
	//
	NormalizesLineEndings bool `json:"normalizes_line_endings,omitempty"`

	// ChangeAnnotationSupport whether the client in general supports change annotations on text edits, create file, rename file and delete file changes.
	//
	//
	ChangeAnnotationSupport *ChangeAnnotationsSupportOptions `json:"change_annotation_support,omitempty"`
}

type DidChangeConfigurationClientCapabilities struct {
	// DynamicRegistration did change configuration notification supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

type DidChangeWatchedFilesClientCapabilities struct {
	// DynamicRegistration did change watched files notification supports dynamic registration. Please note that the current protocol doesn't support static configuration for file changes from the server side.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// RelativePatternSupport whether the client has support for RelativePattern relative pattern or not.
	//
	//
	RelativePatternSupport bool `json:"relative_pattern_support,omitempty"`
}

// ClientSymbolKindOptions.
//
// @since 3.18.0 proposed
type ClientSymbolKindOptions struct {
	// ValueSet the symbol kind values the client supports. When this property exists the client also guarantees that it will handle values outside its set gracefully and falls back to a default value when unknown. If this property is not present the client only supports the symbol kinds from `File` to `Array` as defined in the initial version of the protocol.
	//
	// @since 3.18.0 proposed
	ValueSet []SymbolKind `json:"value_set,omitempty"`
}

// ClientSymbolTagOptions.
//
// @since 3.18.0 proposed
type ClientSymbolTagOptions struct {
	// ValueSet the tags supported by the client.
	//
	// @since 3.18.0 proposed
	ValueSet []SymbolTag `json:"value_set"`
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
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// SymbolKind specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
	SymbolKind *ClientSymbolKindOptions `json:"symbol_kind,omitempty"`

	// TagSupport the client supports tags on `SymbolInformation`. Clients supporting tags have to handle unknown tags gracefully.
	//
	//
	TagSupport *ClientSymbolTagOptions `json:"tag_support,omitempty"`

	// ResolveSupport the client support partial workspace symbols. The client will send the request `workspaceSymbol/resolve` to the server to resolve additional properties.
	//
	//
	ResolveSupport *ClientSymbolResolveOptions `json:"resolve_support,omitempty"`
}

// ExecuteCommandClientCapabilities the client capabilities of a ExecuteCommandRequest.
type ExecuteCommandClientCapabilities struct {
	// DynamicRegistration execute command supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// SemanticTokensWorkspaceClientCapabilities.
//
// @since 3.16.0
type SemanticTokensWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all semantic tokens currently shown. It should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation.
	//
	// @since 3.16.0
	RefreshSupport bool `json:"refresh_support,omitempty"`
}

// CodeLensWorkspaceClientCapabilities.
//
// @since 3.16.0
type CodeLensWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all code lenses currently shown. It should be used with absolute care and is useful for situation where a server for example detect a project wide change that requires such a calculation.
	//
	// @since 3.16.0
	RefreshSupport bool `json:"refresh_support,omitempty"`
}

// FileOperationClientCapabilities capabilities relating to events from file operations by the user in the client. These events do not come from the file system, they come from user operations like renaming a file in the UI.
//
// @since 3.16.0
type FileOperationClientCapabilities struct {
	// DynamicRegistration whether the client supports dynamic registration for file requests/notifications.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// DidCreate the client has support for sending didCreateFiles notifications.
	//
	// @since 3.16.0
	DidCreate bool `json:"did_create,omitempty"`

	// WillCreate the client has support for sending willCreateFiles requests.
	//
	// @since 3.16.0
	WillCreate bool `json:"will_create,omitempty"`

	// DidRename the client has support for sending didRenameFiles notifications.
	//
	// @since 3.16.0
	DidRename bool `json:"did_rename,omitempty"`

	// WillRename the client has support for sending willRenameFiles requests.
	//
	// @since 3.16.0
	WillRename bool `json:"will_rename,omitempty"`

	// DidDelete the client has support for sending didDeleteFiles notifications.
	//
	// @since 3.16.0
	DidDelete bool `json:"did_delete,omitempty"`

	// WillDelete the client has support for sending willDeleteFiles requests.
	//
	// @since 3.16.0
	WillDelete bool `json:"will_delete,omitempty"`
}

// InlineValueWorkspaceClientCapabilities client workspace capabilities specific to inline values.
//
// @since 3.17.0
type InlineValueWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all inline values currently shown. It should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation.
	//
	// @since 3.17.0
	RefreshSupport bool `json:"refresh_support,omitempty"`
}

// InlayHintWorkspaceClientCapabilities client workspace capabilities specific to inlay hints.
//
// @since 3.17.0
type InlayHintWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all inlay hints currently shown. It should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation.
	//
	// @since 3.17.0
	RefreshSupport bool `json:"refresh_support,omitempty"`
}

// DiagnosticWorkspaceClientCapabilities workspace client capabilities specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all pulled diagnostics currently shown. It should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation.
	//
	// @since 3.17.0
	RefreshSupport bool `json:"refresh_support,omitempty"`
}

// FoldingRangeWorkspaceClientCapabilities client workspace capabilities specific to folding ranges
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type FoldingRangeWorkspaceClientCapabilities struct {
	// RefreshSupport whether the client implementation supports a refresh request sent from the server to the client. Note that this event is global and will force the client to refresh all folding ranges currently shown. It should be used with absolute care and is useful for situation where a server for example detects a project wide change that requires such a calculation.
	//
	//  3.18.0
	//
	// Proposed in:.
	// @since 3.18.0 proposed
	RefreshSupport bool `json:"refresh_support,omitempty"`
}

// WorkspaceClientCapabilities workspace specific client capabilities.
type WorkspaceClientCapabilities struct {
	// ApplyEdit the client supports applying batch edits to the workspace by supporting the request 'workspace/applyEdit'.
	ApplyEdit bool `json:"apply_edit,omitempty"`

	// WorkspaceEdit capabilities specific to `WorkspaceEdit`s.
	WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspace_edit,omitempty"`

	// DidChangeConfiguration capabilities specific to the `workspace/didChangeConfiguration` notification.
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"did_change_configuration,omitempty"`

	// DidChangeWatchedFiles capabilities specific to the `workspace/didChangeWatchedFiles` notification.
	DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"did_change_watched_files,omitempty"`

	// Symbol capabilities specific to the `workspace/symbol` request.
	Symbol *WorkspaceSymbolClientCapabilities `json:"symbol,omitempty"`

	// ExecuteCommand capabilities specific to the `workspace/executeCommand` request.
	ExecuteCommand *ExecuteCommandClientCapabilities `json:"execute_command,omitempty"`

	// WorkspaceFolders the client has support for workspace folders.
	//
	//
	WorkspaceFolders bool `json:"workspace_folders,omitempty"`

	// Configuration the client supports `workspace/configuration` requests.
	//
	//
	Configuration bool `json:"configuration,omitempty"`

	// SemanticTokens capabilities specific to the semantic token requests scoped to the workspace.
	//
	//
	SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semantic_tokens,omitempty"`

	// CodeLens capabilities specific to the code lens requests scoped to the workspace.
	//
	//
	CodeLens *CodeLensWorkspaceClientCapabilities `json:"code_lens,omitempty"`

	// FileOperations the client has support for file notifications/requests for user operations on files. Since .
	FileOperations *FileOperationClientCapabilities `json:"file_operations,omitempty"`

	// InlineValue capabilities specific to the inline values requests scoped to the workspace.
	//
	//
	InlineValue *InlineValueWorkspaceClientCapabilities `json:"inline_value,omitempty"`

	// InlayHint capabilities specific to the inlay hint requests scoped to the workspace.
	//
	//
	InlayHint *InlayHintWorkspaceClientCapabilities `json:"inlay_hint,omitempty"`

	// Diagnostics capabilities specific to the diagnostic requests scoped to the workspace.
	//
	//
	Diagnostics *DiagnosticWorkspaceClientCapabilities `json:"diagnostics,omitempty"`

	// FoldingRange capabilities specific to the folding range requests scoped to the workspace.
	//
	//  3.18.0
	//
	// Proposed in:.
	FoldingRange *FoldingRangeWorkspaceClientCapabilities `json:"folding_range,omitempty"`
}

type TextDocumentSyncClientCapabilities struct {
	// DynamicRegistration whether text document synchronization supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// WillSave the client supports sending will save notifications.
	WillSave bool `json:"will_save,omitempty"`

	// WillSaveWaitUntil the client supports sending a will save request and waits for a response providing text edits which will be applied to the document before it is saved.
	WillSaveWaitUntil bool `json:"will_save_wait_until,omitempty"`

	// DidSave the client supports did save notifications.
	DidSave bool `json:"did_save,omitempty"`
}

// CompletionItemTagOptions.
//
// @since 3.18.0 proposed
type CompletionItemTagOptions struct {
	// ValueSet the tags supported by the client.
	//
	// @since 3.18.0 proposed
	ValueSet []CompletionItemTag `json:"value_set"`
}

// ClientCompletionItemResolveOptions.
//
// @since 3.18.0 proposed
type ClientCompletionItemResolveOptions struct {
	// Properties the properties that a client can resolve lazily.
	//
	// @since 3.18.0 proposed
	Properties []string `json:"properties"`
}

// ClientCompletionItemInsertTextModeOptions.
//
// @since 3.18.0 proposed
type ClientCompletionItemInsertTextModeOptions struct {
	// @since 3.18.0 proposed
	ValueSet []InsertTextMode `json:"value_set"`
}

// ClientCompletionItemOptions.
//
// @since 3.18.0 proposed
type ClientCompletionItemOptions struct {
	// SnippetSupport client supports snippets as insert text. A snippet can define tab stops and placeholders with `$1`, `$2` and `${3:foo}`. `$0` defines the final tab stop, it defaults to the end of the snippet. Placeholders with equal identifiers are linked, that is typing in one will update others too.
	//
	// @since 3.18.0 proposed
	SnippetSupport bool `json:"snippet_support,omitempty"`

	// CommitCharactersSupport client supports commit characters on a completion item.
	//
	// @since 3.18.0 proposed
	CommitCharactersSupport bool `json:"commit_characters_support,omitempty"`

	// DocumentationFormat client supports the following content formats for the documentation property. The order describes the preferred format of the client.
	//
	// @since 3.18.0 proposed
	DocumentationFormat []MarkupKind `json:"documentation_format,omitempty"`

	// DeprecatedSupport client supports the deprecated property on a completion item.
	//
	// @since 3.18.0 proposed
	DeprecatedSupport bool `json:"deprecated_support,omitempty"`

	// PreselectSupport client supports the preselect property on a completion item.
	//
	// @since 3.18.0 proposed
	PreselectSupport bool `json:"preselect_support,omitempty"`

	// TagSupport client supports the tag property on a completion item. Clients supporting tags have to handle unknown tags gracefully. Clients especially need to preserve unknown tags when sending a completion item back to the server in a resolve call.
	//
	//
	// @since 3.18.0 proposed
	TagSupport *CompletionItemTagOptions `json:"tag_support,omitempty"`

	// InsertReplaceSupport client support insert replace edit to control different behavior if a completion item is inserted in the text or should replace text.
	//
	//
	// @since 3.18.0 proposed
	InsertReplaceSupport bool `json:"insert_replace_support,omitempty"`

	// ResolveSupport indicates which properties a client can resolve lazily on a completion item. Before version 3.16.0 only the predefined properties `documentation` and `details` could be resolved lazily.
	//
	//
	// @since 3.18.0 proposed
	ResolveSupport *ClientCompletionItemResolveOptions `json:"resolve_support,omitempty"`

	// InsertTextModeSupport the client supports the `insertTextMode` property on a completion item to override the whitespace handling mode as defined by the client (see `insertTextMode`).
	//
	//
	// @since 3.18.0 proposed
	InsertTextModeSupport *ClientCompletionItemInsertTextModeOptions `json:"insert_text_mode_support,omitempty"`

	// LabelDetailsSupport the client has support for completion item label details (see also `CompletionItemLabelDetails`).
	//
	//
	// @since 3.18.0 proposed
	LabelDetailsSupport bool `json:"label_details_support,omitempty"`
}

// ClientCompletionItemOptionsKind.
//
// @since 3.18.0 proposed
type ClientCompletionItemOptionsKind struct {
	// ValueSet the completion item kind values the client supports. When this property exists the client also guarantees that it will handle values outside its set gracefully and falls back to a default value when unknown. If this property is not present the client only supports the completion items kinds from `Text` to `Reference` as defined in the initial version of the protocol.
	//
	// @since 3.18.0 proposed
	ValueSet []CompletionItemKind `json:"value_set,omitempty"`
}

// CompletionListCapabilities the client supports the following `CompletionList` specific capabilities.
//
// @since 3.17.0
type CompletionListCapabilities struct {
	// ItemDefaults the client supports the following itemDefaults on a completion list. The value lists the supported property names of the `CompletionList.itemDefaults` object. If omitted no properties are supported.
	//
	//
	// @since 3.17.0
	ItemDefaults []string `json:"item_defaults,omitempty"`
}

// CompletionClientCapabilities completion client capabilities.
type CompletionClientCapabilities struct {
	// DynamicRegistration whether completion supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// CompletionItem the client supports the following `CompletionItem` specific capabilities.
	CompletionItem *ClientCompletionItemOptions `json:"completion_item,omitempty"`

	CompletionItemKind *ClientCompletionItemOptionsKind `json:"completion_item_kind,omitempty"`

	// InsertTextMode defines how the client handles whitespace and indentation when accepting a completion item that uses multi line text in either `insertText` or `textEdit`.
	//
	//
	InsertTextMode *InsertTextMode `json:"insert_text_mode,omitempty"`

	// ContextSupport the client supports to send additional context information for a `textDocument/completion` request.
	ContextSupport bool `json:"context_support,omitempty"`

	// CompletionList the client supports the following `CompletionList` specific capabilities.
	//
	//
	CompletionList *CompletionListCapabilities `json:"completion_list,omitempty"`
}

type HoverClientCapabilities struct {
	// DynamicRegistration whether hover supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// ContentFormat client supports the following content formats for the content property. The order describes the preferred format of the client.
	ContentFormat []MarkupKind `json:"content_format,omitempty"`
}

// ClientSignatureParameterInformationOptions.
//
// @since 3.18.0 proposed
type ClientSignatureParameterInformationOptions struct {
	// LabelOffsetSupport the client supports processing label offsets instead of a simple label string.
	//
	//
	// @since 3.18.0 proposed
	LabelOffsetSupport bool `json:"label_offset_support,omitempty"`
}

// ClientSignatureInformationOptions.
//
// @since 3.18.0 proposed
type ClientSignatureInformationOptions struct {
	// DocumentationFormat client supports the following content formats for the documentation property. The order describes the preferred format of the client.
	//
	// @since 3.18.0 proposed
	DocumentationFormat []MarkupKind `json:"documentation_format,omitempty"`

	// ParameterInformation client capabilities specific to parameter information.
	//
	// @since 3.18.0 proposed
	ParameterInformation *ClientSignatureParameterInformationOptions `json:"parameter_information,omitempty"`

	// ActiveParameterSupport the client supports the `activeParameter` property on `SignatureInformation` literal.
	//
	//
	// @since 3.18.0 proposed
	ActiveParameterSupport bool `json:"active_parameter_support,omitempty"`

	// NoActiveParameterSupport the client supports the `activeParameter` property on `SignatureHelp`/`SignatureInformation` being set to `null` to indicate that no parameter should be active.
	//
	//  3.18.0
	//
	// Proposed in:.
	// @since 3.18.0 proposed
	NoActiveParameterSupport bool `json:"no_active_parameter_support,omitempty"`
}

// SignatureHelpClientCapabilities client Capabilities for a SignatureHelpRequest.
type SignatureHelpClientCapabilities struct {
	// DynamicRegistration whether signature help supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// SignatureInformation the client supports the following `SignatureInformation` specific properties.
	SignatureInformation *ClientSignatureInformationOptions `json:"signature_information,omitempty"`

	// ContextSupport the client supports to send additional context information for a `textDocument/signatureHelp` request. A client that opts into contextSupport will also support the `retriggerCharacters` on `SignatureHelpOptions`.
	//
	//
	ContextSupport bool `json:"context_support,omitempty"`
}

// DeclarationClientCapabilities.
//
// @since 3.14.0
type DeclarationClientCapabilities struct {
	// DynamicRegistration whether declaration supports dynamic registration. If this is set to `true` the client supports the new `DeclarationRegistrationOptions` return value for the corresponding server capability as well.
	//
	// @since 3.14.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// LinkSupport the client supports additional metadata in the form of declaration links.
	//
	// @since 3.14.0
	LinkSupport bool `json:"link_support,omitempty"`
}

// DefinitionClientCapabilities client Capabilities for a DefinitionRequest.
type DefinitionClientCapabilities struct {
	// DynamicRegistration whether definition supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// LinkSupport the client supports additional metadata in the form of definition links.
	//
	//
	LinkSupport bool `json:"link_support,omitempty"`
}

// TypeDefinitionClientCapabilities since .
type TypeDefinitionClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `TypeDefinitionRegistrationOptions` return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// LinkSupport the client supports additional metadata in the form of definition links. Since .
	LinkSupport bool `json:"link_support,omitempty"`
}

// ImplementationClientCapabilities.
//
// @since 3.6.0
type ImplementationClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `ImplementationRegistrationOptions` return value for the corresponding server capability as well.
	//
	// @since 3.6.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// LinkSupport the client supports additional metadata in the form of definition links.
	//
	//
	// @since 3.6.0
	LinkSupport bool `json:"link_support,omitempty"`
}

// ReferenceClientCapabilities client Capabilities for a ReferencesRequest.
type ReferenceClientCapabilities struct {
	// DynamicRegistration whether references supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// DocumentHighlightClientCapabilities client Capabilities for a DocumentHighlightRequest.
type DocumentHighlightClientCapabilities struct {
	// DynamicRegistration whether document highlight supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// DocumentSymbolClientCapabilities client Capabilities for a DocumentSymbolRequest.
type DocumentSymbolClientCapabilities struct {
	// DynamicRegistration whether document symbol supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// SymbolKind specific capabilities for the `SymbolKind` in the `textDocument/documentSymbol` request.
	SymbolKind *ClientSymbolKindOptions `json:"symbol_kind,omitempty"`

	// HierarchicalDocumentSymbolSupport the client supports hierarchical document symbols.
	HierarchicalDocumentSymbolSupport bool `json:"hierarchical_document_symbol_support,omitempty"`

	// TagSupport the client supports tags on `SymbolInformation`. Tags are supported on `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true. Clients supporting tags have to handle unknown tags gracefully.
	//
	//
	TagSupport *ClientSymbolTagOptions `json:"tag_support,omitempty"`

	// LabelSupport the client supports an additional label presented in the UI when registering a document symbol provider.
	//
	//
	LabelSupport bool `json:"label_support,omitempty"`
}

// ClientCodeActionKindOptions.
//
// @since 3.18.0 proposed
type ClientCodeActionKindOptions struct {
	// ValueSet the code action kind values the client supports. When this property exists the client also guarantees that it will handle values outside its set gracefully and falls back to a default value when unknown.
	//
	// @since 3.18.0 proposed
	ValueSet []CodeActionKind `json:"value_set"`
}

// ClientCodeActionLiteralOptions.
//
// @since 3.18.0 proposed
type ClientCodeActionLiteralOptions struct {
	// CodeActionKind the code action kind is support with the following value set.
	//
	// @since 3.18.0 proposed
	CodeActionKind ClientCodeActionKindOptions `json:"code_action_kind"`
}

// ClientCodeActionResolveOptions.
//
// @since 3.18.0 proposed
type ClientCodeActionResolveOptions struct {
	// Properties the properties that a client can resolve lazily.
	//
	// @since 3.18.0 proposed
	Properties []string `json:"properties"`
}

// CodeActionClientCapabilities the Client Capabilities of a CodeActionRequest.
type CodeActionClientCapabilities struct {
	// DynamicRegistration whether code action supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// CodeActionLiteralSupport the client support code action literals of type `CodeAction` as a valid response of the `textDocument/codeAction` request. If the property is not set the request can only return `Command` literals.
	//
	//
	CodeActionLiteralSupport *ClientCodeActionLiteralOptions `json:"code_action_literal_support,omitempty"`

	// IsPreferredSupport whether code action supports the `isPreferred` property.
	//
	//
	IsPreferredSupport bool `json:"is_preferred_support,omitempty"`

	// DisabledSupport whether code action supports the `disabled` property.
	//
	//
	DisabledSupport bool `json:"disabled_support,omitempty"`

	// DataSupport whether code action supports the `data` property which is preserved between a `textDocument/codeAction` and a `codeAction/resolve` request.
	//
	//
	DataSupport bool `json:"data_support,omitempty"`

	// ResolveSupport whether the client supports resolving additional code action properties via a separate `codeAction/resolve` request.
	//
	//
	ResolveSupport *ClientCodeActionResolveOptions `json:"resolve_support,omitempty"`

	// HonorsChangeAnnotations whether the client honors the change annotations in text edits and resource operations returned via the `CodeAction#edit` property by for example presenting the workspace edit in the user interface and asking for confirmation.
	//
	//
	HonorsChangeAnnotations bool `json:"honors_change_annotations,omitempty"`

	// DocumentationSupport whether the client supports documentation for a class of code actions.
	//
	//  3.18.0
	//
	// Proposed in:.
	DocumentationSupport bool `json:"documentation_support,omitempty"`
}

// CodeLensClientCapabilities the client capabilities of a CodeLensRequest.
type CodeLensClientCapabilities struct {
	// DynamicRegistration whether code lens supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// DocumentLinkClientCapabilities the client capabilities of a DocumentLinkRequest.
type DocumentLinkClientCapabilities struct {
	// DynamicRegistration whether document link supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// TooltipSupport whether the client supports the `tooltip` property on `DocumentLink`.
	//
	//
	TooltipSupport bool `json:"tooltip_support,omitempty"`
}

type DocumentColorClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `DocumentColorRegistrationOptions` return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// DocumentFormattingClientCapabilities client capabilities of a DocumentFormattingRequest.
type DocumentFormattingClientCapabilities struct {
	// DynamicRegistration whether formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// DocumentRangeFormattingClientCapabilities client capabilities of a DocumentRangeFormattingRequest.
type DocumentRangeFormattingClientCapabilities struct {
	// DynamicRegistration whether range formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// RangesSupport whether the client supports formatting multiple ranges at once.
	//
	//  3.18.0
	//
	// Proposed in:.
	RangesSupport bool `json:"ranges_support,omitempty"`
}

// DocumentOnTypeFormattingClientCapabilities client capabilities of a DocumentOnTypeFormattingRequest.
type DocumentOnTypeFormattingClientCapabilities struct {
	// DynamicRegistration whether on type formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

type RenameClientCapabilities struct {
	// DynamicRegistration whether rename supports dynamic registration.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// PrepareSupport client supports testing for validity of rename operations before execution.
	//
	//
	PrepareSupport bool `json:"prepare_support,omitempty"`

	// PrepareSupportDefaultBehavior client supports the default behavior result. The value indicates the default behavior used by the client.
	//
	//
	PrepareSupportDefaultBehavior *PrepareSupportDefaultBehavior `json:"prepare_support_default_behavior,omitempty"`

	// HonorsChangeAnnotations whether the client honors the change annotations in text edits and resource operations returned via the rename request's workspace edit by for example presenting the workspace edit in the user interface and asking for confirmation.
	//
	//
	HonorsChangeAnnotations bool `json:"honors_change_annotations,omitempty"`
}

// ClientFoldingRangeKindOptions.
//
// @since 3.18.0 proposed
type ClientFoldingRangeKindOptions struct {
	// ValueSet the folding range kind values the client supports. When this property exists the client also guarantees that it will handle values outside its set gracefully and falls back to a default value when unknown.
	//
	// @since 3.18.0 proposed
	ValueSet []FoldingRangeKind `json:"value_set,omitempty"`
}

// ClientFoldingRangeOptions.
//
// @since 3.18.0 proposed
type ClientFoldingRangeOptions struct {
	// CollapsedText if set, the client signals that it supports setting collapsedText on folding ranges to display custom labels instead of the default text.
	//
	//
	// @since 3.18.0 proposed
	CollapsedText bool `json:"collapsed_text,omitempty"`
}

type FoldingRangeClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration for folding range providers. If this is set to `true` the client supports the new `FoldingRangeRegistrationOptions` return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// RangeLimit the maximum number of folding ranges that the client prefers to receive per document. The value serves as a hint, servers are free to follow the limit.
	RangeLimit uint32 `json:"range_limit,omitempty"`

	// LineFoldingOnly if set, the client signals that it only supports folding complete lines. If set, client will ignore specified `startCharacter` and `endCharacter` properties in a FoldingRange.
	LineFoldingOnly bool `json:"line_folding_only,omitempty"`

	// FoldingRangeKind specific options for the folding range kind.
	//
	//
	FoldingRangeKind *ClientFoldingRangeKindOptions `json:"folding_range_kind,omitempty"`

	// FoldingRange specific options for the folding range.
	//
	//
	FoldingRange *ClientFoldingRangeOptions `json:"folding_range,omitempty"`
}

type SelectionRangeClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration for selection range providers. If this is set to `true` the client supports the new `SelectionRangeRegistrationOptions` return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// ClientDiagnosticsTagOptions.
//
// @since 3.18.0 proposed
type ClientDiagnosticsTagOptions struct {
	// ValueSet the tags supported by the client.
	//
	// @since 3.18.0 proposed
	ValueSet []DiagnosticTag `json:"value_set"`
}

// PublishDiagnosticsClientCapabilities the publish diagnostic client capabilities.
type PublishDiagnosticsClientCapabilities struct {
	// RelatedInformation whether the clients accepts diagnostics with related information.
	RelatedInformation bool `json:"related_information,omitempty"`

	// TagSupport client supports the tag property to provide meta data about a diagnostic. Clients supporting tags have to handle unknown tags gracefully.
	//
	//
	TagSupport *ClientDiagnosticsTagOptions `json:"tag_support,omitempty"`

	// VersionSupport whether the client interprets the version property of the `textDocument/publishDiagnostics` notification's parameter.
	//
	//
	VersionSupport bool `json:"version_support,omitempty"`

	// CodeDescriptionSupport client supports a codeDescription property
	//
	//
	CodeDescriptionSupport bool `json:"code_description_support,omitempty"`

	// DataSupport whether code action supports the `data` property which is preserved between a `textDocument/publishDiagnostics` and `textDocument/codeAction` request.
	//
	//
	DataSupport bool `json:"data_support,omitempty"`
}

// CallHierarchyClientCapabilities.
//
// @since 3.16.0
type CallHierarchyClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// ClientSemanticTokensRequestFullDelta.
//
// @since 3.18.0 proposed
type ClientSemanticTokensRequestFullDelta struct {
	// Delta the client will send the `textDocument/semanticTokens/full/delta` request if the server provides a corresponding handler.
	//
	// @since 3.18.0 proposed
	Delta bool `json:"delta,omitempty"`
}

// ClientSemanticTokensRequestOptions.
//
// @since 3.18.0 proposed
type ClientSemanticTokensRequestOptions struct {
	// Range the client will send the `textDocument/semanticTokens/range` request if the server provides a corresponding handler.
	//
	// @since 3.18.0 proposed
	Range any/* or */ `json:"range,omitempty"`

	// Full the client will send the `textDocument/semanticTokens/full` request if the server provides a corresponding handler.
	//
	// @since 3.18.0 proposed
	Full any/* or */ `json:"full,omitempty"`
}

// SemanticTokensClientCapabilities.
//
// @since 3.16.0
type SemanticTokensClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// Requests which requests the client supports and might send to the server depending on the server's capability. Please note that clients might not show semantic tokens or degrade some of the user experience if a range or full request is advertised by the client but not provided by the server. If for example the client capability `requests.full` and `request.range` are both set to true but the server only provides a range provider the client might not render a minimap correctly or might even decide to not show any semantic tokens at all.
	//
	// @since 3.16.0
	Requests ClientSemanticTokensRequestOptions `json:"requests"`

	// TokenTypes the token types that the client supports.
	//
	// @since 3.16.0
	TokenTypes []string `json:"token_types"`

	// TokenModifiers the token modifiers that the client supports.
	//
	// @since 3.16.0
	TokenModifiers []string `json:"token_modifiers"`

	// Formats the token formats the clients supports.
	//
	// @since 3.16.0
	Formats []TokenFormat `json:"formats"`

	// OverlappingTokenSupport whether the client supports tokens that can overlap each other.
	//
	// @since 3.16.0
	OverlappingTokenSupport bool `json:"overlapping_token_support,omitempty"`

	// MultilineTokenSupport whether the client supports tokens that can span multiple lines.
	//
	// @since 3.16.0
	MultilineTokenSupport bool `json:"multiline_token_support,omitempty"`

	// ServerCancelSupport whether the client allows the server to actively cancel a semantic token request, e.g. supports returning LSPErrorCodes.ServerCancelled. If a server does the client needs to retrigger the request.
	//
	//
	// @since 3.16.0
	ServerCancelSupport bool `json:"server_cancel_support,omitempty"`

	// AugmentsSyntaxTokens whether the client uses semantic tokens to augment existing syntax tokens. If set to `true` client side created syntax tokens and semantic tokens are both used for colorization. If set to `false` the client only uses the returned semantic tokens for colorization. If the value is `undefined` then the client behavior is not specified.
	//
	//
	// @since 3.16.0
	AugmentsSyntaxTokens bool `json:"augments_syntax_tokens,omitempty"`
}

// LinkedEditingRangeClientCapabilities client capabilities for the linked editing range request.
//
// @since 3.16.0
type LinkedEditingRangeClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// MonikerClientCapabilities client capabilities specific to the moniker request.
//
// @since 3.16.0
type MonikerClientCapabilities struct {
	// DynamicRegistration whether moniker supports dynamic registration. If this is set to `true` the client supports the new `MonikerRegistrationOptions` return value for the corresponding server capability as well.
	//
	// @since 3.16.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// TypeHierarchyClientCapabilities.
//
// @since 3.17.0
type TypeHierarchyClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// InlineValueClientCapabilities client capabilities specific to inline values.
//
// @since 3.17.0
type InlineValueClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration for inline value providers.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// ClientInlayHintResolveOptions.
//
// @since 3.18.0 proposed
type ClientInlayHintResolveOptions struct {
	// Properties the properties that a client can resolve lazily.
	//
	// @since 3.18.0 proposed
	Properties []string `json:"properties"`
}

// InlayHintClientCapabilities inlay hint client capabilities.
//
// @since 3.17.0
type InlayHintClientCapabilities struct {
	// DynamicRegistration whether inlay hints support dynamic registration.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// ResolveSupport indicates which properties a client can resolve lazily on an inlay hint.
	//
	// @since 3.17.0
	ResolveSupport *ClientInlayHintResolveOptions `json:"resolve_support,omitempty"`
}

// DiagnosticClientCapabilities client capabilities specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// RelatedDocumentSupport whether the clients supports related documents for document diagnostic pulls.
	//
	// @since 3.17.0
	RelatedDocumentSupport bool `json:"related_document_support,omitempty"`
}

// InlineCompletionClientCapabilities client capabilities specific to inline completions.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type InlineCompletionClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration for inline completion providers.
	//
	// @since 3.18.0 proposed
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`
}

// TextDocumentClientCapabilities text document specific client capabilities.
type TextDocumentClientCapabilities struct {
	// Synchronization defines which synchronization capabilities the client supports.
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitempty"`

	// Completion capabilities specific to the `textDocument/completion` request.
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`

	// Hover capabilities specific to the `textDocument/hover` request.
	Hover *HoverClientCapabilities `json:"hover,omitempty"`

	// SignatureHelp capabilities specific to the `textDocument/signatureHelp` request.
	SignatureHelp *SignatureHelpClientCapabilities `json:"signature_help,omitempty"`

	// Declaration capabilities specific to the `textDocument/declaration` request.
	//
	//
	Declaration *DeclarationClientCapabilities `json:"declaration,omitempty"`

	// Definition capabilities specific to the `textDocument/definition` request.
	Definition *DefinitionClientCapabilities `json:"definition,omitempty"`

	// TypeDefinition capabilities specific to the `textDocument/typeDefinition` request.
	//
	//
	TypeDefinition *TypeDefinitionClientCapabilities `json:"type_definition,omitempty"`

	// Implementation capabilities specific to the `textDocument/implementation` request.
	//
	//
	Implementation *ImplementationClientCapabilities `json:"implementation,omitempty"`

	// References capabilities specific to the `textDocument/references` request.
	References *ReferenceClientCapabilities `json:"references,omitempty"`

	// DocumentHighlight capabilities specific to the `textDocument/documentHighlight` request.
	DocumentHighlight *DocumentHighlightClientCapabilities `json:"document_highlight,omitempty"`

	// DocumentSymbol capabilities specific to the `textDocument/documentSymbol` request.
	DocumentSymbol *DocumentSymbolClientCapabilities `json:"document_symbol,omitempty"`

	// CodeAction capabilities specific to the `textDocument/codeAction` request.
	CodeAction *CodeActionClientCapabilities `json:"code_action,omitempty"`

	// CodeLens capabilities specific to the `textDocument/codeLens` request.
	CodeLens *CodeLensClientCapabilities `json:"code_lens,omitempty"`

	// DocumentLink capabilities specific to the `textDocument/documentLink` request.
	DocumentLink *DocumentLinkClientCapabilities `json:"document_link,omitempty"`

	// ColorProvider capabilities specific to the `textDocument/documentColor` and the `textDocument/colorPresentation` request.
	//
	//
	ColorProvider *DocumentColorClientCapabilities `json:"color_provider,omitempty"`

	// Formatting capabilities specific to the `textDocument/formatting` request.
	Formatting *DocumentFormattingClientCapabilities `json:"formatting,omitempty"`

	// RangeFormatting capabilities specific to the `textDocument/rangeFormatting` request.
	RangeFormatting *DocumentRangeFormattingClientCapabilities `json:"range_formatting,omitempty"`

	// OnTypeFormatting capabilities specific to the `textDocument/onTypeFormatting` request.
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities `json:"on_type_formatting,omitempty"`

	// Rename capabilities specific to the `textDocument/rename` request.
	Rename *RenameClientCapabilities `json:"rename,omitempty"`

	// FoldingRange capabilities specific to the `textDocument/foldingRange` request.
	//
	//
	FoldingRange *FoldingRangeClientCapabilities `json:"folding_range,omitempty"`

	// SelectionRange capabilities specific to the `textDocument/selectionRange` request.
	//
	//
	SelectionRange *SelectionRangeClientCapabilities `json:"selection_range,omitempty"`

	// PublishDiagnostics capabilities specific to the `textDocument/publishDiagnostics` notification.
	PublishDiagnostics *PublishDiagnosticsClientCapabilities `json:"publish_diagnostics,omitempty"`

	// CallHierarchy capabilities specific to the various call hierarchy requests.
	//
	//
	CallHierarchy *CallHierarchyClientCapabilities `json:"call_hierarchy,omitempty"`

	// SemanticTokens capabilities specific to the various semantic token request.
	//
	//
	SemanticTokens *SemanticTokensClientCapabilities `json:"semantic_tokens,omitempty"`

	// LinkedEditingRange capabilities specific to the `textDocument/linkedEditingRange` request.
	//
	//
	LinkedEditingRange *LinkedEditingRangeClientCapabilities `json:"linked_editing_range,omitempty"`

	// Moniker client capabilities specific to the `textDocument/moniker` request.
	//
	//
	Moniker *MonikerClientCapabilities `json:"moniker,omitempty"`

	// TypeHierarchy capabilities specific to the various type hierarchy requests.
	//
	//
	TypeHierarchy *TypeHierarchyClientCapabilities `json:"type_hierarchy,omitempty"`

	// InlineValue capabilities specific to the `textDocument/inlineValue` request.
	//
	//
	InlineValue *InlineValueClientCapabilities `json:"inline_value,omitempty"`

	// InlayHint capabilities specific to the `textDocument/inlayHint` request.
	//
	//
	InlayHint *InlayHintClientCapabilities `json:"inlay_hint,omitempty"`

	// Diagnostic capabilities specific to the diagnostic pull model.
	//
	//
	Diagnostic *DiagnosticClientCapabilities `json:"diagnostic,omitempty"`

	// InlineCompletion client capabilities specific to inline completions.
	//
	//  3.18.0
	//
	// Proposed in:.
	InlineCompletion *InlineCompletionClientCapabilities `json:"inline_completion,omitempty"`
}

// NotebookDocumentSyncClientCapabilities notebook specific client capabilities.
//
// @since 3.17.0
type NotebookDocumentSyncClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration. If this is set to `true` the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)` return value for the corresponding server capability as well.
	//
	// @since 3.17.0
	DynamicRegistration bool `json:"dynamic_registration,omitempty"`

	// ExecutionSummarySupport the client supports sending execution summary data per cell.
	//
	// @since 3.17.0
	ExecutionSummarySupport bool `json:"execution_summary_support,omitempty"`
}

// NotebookDocumentClientCapabilities capabilities specific to the notebook document support.
//
// @since 3.17.0
type NotebookDocumentClientCapabilities struct {
	// Synchronization capabilities specific to notebook document synchronization
	//
	//
	// @since 3.17.0
	Synchronization NotebookDocumentSyncClientCapabilities `json:"synchronization"`
}

// ClientShowMessageActionItemOptions.
//
// @since 3.18.0 proposed
type ClientShowMessageActionItemOptions struct {
	// AdditionalPropertiesSupport whether the client supports additional attributes which are preserved and send back to the server in the request's response.
	//
	// @since 3.18.0 proposed
	AdditionalPropertiesSupport bool `json:"additional_properties_support,omitempty"`
}

// ShowMessageRequestClientCapabilities show message request client capabilities.
type ShowMessageRequestClientCapabilities struct {
	// MessageActionItem capabilities specific to the `MessageActionItem` type.
	MessageActionItem *ClientShowMessageActionItemOptions `json:"message_action_item,omitempty"`
}

// ShowDocumentClientCapabilities client capabilities for the showDocument request.
//
// @since 3.16.0
type ShowDocumentClientCapabilities struct {
	// Support the client has support for the showDocument request.
	//
	// @since 3.16.0
	Support bool `json:"support"`
}

type WindowClientCapabilities struct {
	// WorkDoneProgress it indicates whether the client supports server initiated progress using the `window/workDoneProgress/create` request. The capability also controls Whether client supports handling of progress notifications. If set servers are allowed to report a `workDoneProgress` property in the request specific server capabilities.
	//
	//
	WorkDoneProgress bool `json:"work_done_progress,omitempty"`

	// ShowMessage capabilities specific to the showMessage request.
	//
	//
	ShowMessage *ShowMessageRequestClientCapabilities `json:"show_message,omitempty"`

	// ShowDocument capabilities specific to the showDocument request.
	//
	//
	ShowDocument *ShowDocumentClientCapabilities `json:"show_document,omitempty"`
}

// StaleRequestSupportOptions.
//
// @since 3.18.0 proposed
type StaleRequestSupportOptions struct {
	// Cancel the client will actively cancel the request.
	//
	// @since 3.18.0 proposed
	Cancel bool `json:"cancel"`

	// RetryOnContentModified the list of requests for which the client will retry the request if it receives a response with error code `ContentModified`.
	//
	// @since 3.18.0 proposed
	RetryOnContentModified []string `json:"retry_on_content_modified"`
}

// RegularExpressionsClientCapabilities client capabilities specific to regular expressions.
//
// @since 3.16.0
type RegularExpressionsClientCapabilities struct {
	// Engine the engine's name.
	//
	// @since 3.16.0
	Engine RegularExpressionEngineKind `json:"engine"`

	// Version the engine's version.
	//
	// @since 3.16.0
	Version string `json:"version,omitempty"`
}

// MarkdownClientCapabilities client capabilities specific to the used markdown parser.
//
// @since 3.16.0
type MarkdownClientCapabilities struct {
	// Parser the name of the parser.
	//
	// @since 3.16.0
	Parser string `json:"parser"`

	// Version the version of the parser.
	//
	// @since 3.16.0
	Version string `json:"version,omitempty"`

	// AllowedTags a list of HTML tags that the client allows / supports in Markdown.
	//
	//
	// @since 3.16.0
	AllowedTags []string `json:"allowed_tags,omitempty"`
}

// GeneralClientCapabilities general client capabilities.
//
// @since 3.16.0
type GeneralClientCapabilities struct {
	// StaleRequestSupport client capability that signals how the client handles stale requests (e.g. a request for which the client will not process the response anymore since the information is outdated).
	//
	//
	// @since 3.16.0
	StaleRequestSupport *StaleRequestSupportOptions `json:"stale_request_support,omitempty"`

	// RegularExpressions client capabilities specific to regular expressions.
	//
	//
	// @since 3.16.0
	RegularExpressions *RegularExpressionsClientCapabilities `json:"regular_expressions,omitempty"`

	// Markdown client capabilities specific to the client's markdown parser.
	//
	//
	// @since 3.16.0
	Markdown *MarkdownClientCapabilities `json:"markdown,omitempty"`

	// PositionEncodings the position encodings supported by the client. Client and server have to agree on the same position encoding to ensure that offsets (e.g. character position in a line) are interpreted the same on both sides. To keep the protocol backwards compatible the following applies: if the value 'utf-16' is missing from the array of position encodings servers can assume that the client supports UTF-16. UTF-16 is therefore a mandatory encoding. If omitted it defaults to ['utf-16']. Implementation considerations: since the conversion from one encoding into another requires the content of the file / line the conversion is best done where the file is read which is usually on the server side.
	//
	//
	// @since 3.16.0
	PositionEncodings []PositionEncodingKind `json:"position_encodings,omitempty"`
}

// ClientCapabilities defines the capabilities provided by the client.
type ClientCapabilities struct {
	// Workspace workspace specific client capabilities.
	Workspace *WorkspaceClientCapabilities `json:"workspace,omitempty"`

	// TextDocument text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"text_document,omitempty"`

	// NotebookDocument capabilities specific to the notebook document support.
	//
	//
	NotebookDocument *NotebookDocumentClientCapabilities `json:"notebook_document,omitempty"`

	// Window window specific client capabilities.
	Window *WindowClientCapabilities `json:"window,omitempty"`

	// General general client capabilities.
	//
	//
	General *GeneralClientCapabilities `json:"general,omitempty"`

	// Experimental experimental client capabilities.
	Experimental any `json:"experimental,omitempty"`
}

// InitializeParamsBase the initialize parameters.
type InitializeParamsBase[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT

	// ProcessID the process Id of the parent process that started the server. Is `null` if the process has not been started by another process. If the parent process is not alive then the server should exit.
	ProcessID any/* or */ `json:"process_id"`

	// ClientInfo information about the client
	//
	//
	ClientInfo *ClientInfo `json:"client_info,omitempty"`

	// Locale the locale the client is currently showing the user interface in. This must not necessarily be the locale of the operating system. Uses IETF language tags as the value's syntax (See https://en.wikipedia.org/wiki/IETF_language_tag)
	//
	//
	Locale string `json:"locale,omitempty"`

	// RootPath the rootPath of the workspace. Is null if no folder is open.
	//
	// Deprecated: in favour of rootUri.
	RootPath any/* or */ `json:"root_path,omitempty"`

	// RootURI the rootUri of the workspace. Is null if no folder is open. If both `rootPath` and `rootUri` are set `rootUri` wins.
	//
	// Deprecated: in favour of workspaceFolders.
	RootURI any/* or */ `json:"root_uri"`

	// Capabilities the capabilities provided by the client (editor or tool).
	Capabilities ClientCapabilities `json:"capabilities"`

	// InitializationOptions user provided initialization options.
	InitializationOptions any `json:"initialization_options,omitempty"`

	// Trace the initial trace setting. If omitted trace is disabled ('off').
	Trace *TraceValue `json:"trace,omitempty"`
}

type WorkspaceFoldersInitializeParams struct {
	// WorkspaceFolders the workspace folders configured in the client when the server starts. This property is only available if the client supports workspace folders. It can be `null` if the client supports workspace folders but none are configured.
	//
	//
	WorkspaceFolders any /* or */ `json:"workspace_folders,omitempty"`
}

type InitializeParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// extends
	InitializeParamsBase[T, TT]
	WorkspaceFoldersInitializeParams
}

// SaveOptions save options.
type SaveOptions struct {
	// IncludeText the client is supposed to include the content on save.
	IncludeText bool `json:"include_text,omitempty"`
}

type TextDocumentSyncOptions struct {
	// OpenClose open and close notifications are sent to the server. If omitted open close notification should not be sent.
	OpenClose bool `json:"open_close,omitempty"`

	// Change change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
	Change *TextDocumentSyncKind `json:"change,omitempty"`

	// WillSave if present will save notifications are sent to the server. If omitted the notification should not be sent.
	WillSave bool `json:"will_save,omitempty"`

	// WillSaveWaitUntil if present will save wait until requests are sent to the server. If omitted the request should not be sent.
	WillSaveWaitUntil bool `json:"will_save_wait_until,omitempty"`

	// Save if present save notifications are sent to the server. If omitted the notification should not be sent.
	Save any/* or */ `json:"save,omitempty"`
}

// ServerCompletionItemOptions.
//
// @since 3.18.0 proposed
type ServerCompletionItemOptions struct {
	// LabelDetailsSupport the server has support for completion item label details (see also `CompletionItemLabelDetails`) when receiving a completion item in a resolve call.
	//
	//
	// @since 3.18.0 proposed
	LabelDetailsSupport bool `json:"label_details_support,omitempty"`
}

// CompletionOptions completion options.
type CompletionOptions struct {
	// mixins
	WorkDoneProgressOptions

	// TriggerCharacters most tools trigger completion request automatically without explicitly requesting it using a keyboard shortcut (e.g. Ctrl+Space). Typically they do so when the user starts to type an identifier. For example if the user types `c` in a JavaScript file code complete will automatically pop up present `console` besides others as a completion item. Characters that make up identifiers don't need to be listed here. If code complete should automatically be trigger on characters not being valid inside an identifier (for example `.` in JavaScript) list them in `triggerCharacters`.
	TriggerCharacters []string `json:"trigger_characters,omitempty"`

	// AllCommitCharacters the list of all possible characters that commit a completion. This field can be used if clients don't support individual commit characters per completion item. See `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport` If a server provides both `allCommitCharacters` and commit characters on an individual completion item the ones on the completion item win.
	//
	//
	AllCommitCharacters []string `json:"all_commit_characters,omitempty"`

	// ResolveProvider the server provides support to resolve additional information for a completion item.
	ResolveProvider bool `json:"resolve_provider,omitempty"`

	// CompletionItem the server supports the following `CompletionItem` specific capabilities.
	//
	//
	CompletionItem *ServerCompletionItemOptions `json:"completion_item,omitempty"`
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
	TriggerCharacters []string `json:"trigger_characters,omitempty"`

	// RetriggerCharacters list of characters that re-trigger signature help. These trigger characters are only active when signature help is already showing. All trigger characters are also counted as re-trigger characters.
	//
	//
	RetriggerCharacters []string `json:"retrigger_characters,omitempty"`
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
	//
	//
	Label string `json:"label,omitempty"`
}

// CodeActionKindDocumentation documentation for a class of code actions.
//
//	3.18.0
//
// Proposed in:.
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
	CodeActionKinds []CodeActionKind `json:"code_action_kinds,omitempty"`

	// Documentation static documentation for a class of code actions. Documentation from the provider should be shown in the code actions menu if either: - Code actions of `kind` are requested by the editor. In this case, the editor will show the documentation that  most closely matches the requested code action kind. For example, if a provider has documentation for  both `Refactor` and `RefactorExtract`, when the user requests code actions for `RefactorExtract`,  the editor will use the documentation for `RefactorExtract` instead of the documentation for `Refactor`. - Any code actions of `kind` are returned by the provider. At most one documentation entry should be shown per provider.
	//
	//  3.18.0
	//
	// Proposed in:.
	Documentation []CodeActionKindDocumentation `json:"documentation,omitempty"`

	// ResolveProvider the server provides support to resolve additional information for a code action.
	//
	//
	ResolveProvider bool `json:"resolve_provider,omitempty"`
}

// CodeLensOptions code Lens provider options of a CodeLensRequest.
type CodeLensOptions struct {
	// mixins
	WorkDoneProgressOptions

	// ResolveProvider code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolve_provider,omitempty"`
}

// DocumentLinkOptions provider options for a DocumentLinkRequest.
type DocumentLinkOptions struct {
	// mixins
	WorkDoneProgressOptions

	// ResolveProvider document links have a resolve provider as well.
	ResolveProvider bool `json:"resolve_provider,omitempty"`
}

// WorkspaceSymbolOptions server capabilities for a WorkspaceSymbolRequest.
type WorkspaceSymbolOptions struct {
	// mixins
	WorkDoneProgressOptions

	// ResolveProvider the server provides support to resolve additional information for a workspace symbol.
	//
	//
	ResolveProvider bool `json:"resolve_provider,omitempty"`
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

	// RangesSupport whether the server supports formatting multiple ranges at once.
	//
	//  3.18.0
	//
	// Proposed in:.
	RangesSupport bool `json:"ranges_support,omitempty"`
}

// DocumentOnTypeFormattingOptions provider options for a DocumentOnTypeFormattingRequest.
type DocumentOnTypeFormattingOptions struct {
	// FirstTriggerCharacter a character on which formatting should be triggered, like `{`.
	FirstTriggerCharacter string `json:"first_trigger_character"`

	// MoreTriggerCharacter more trigger characters.
	MoreTriggerCharacter []string `json:"more_trigger_character,omitempty"`
}

// RenameOptions provider options for a RenameRequest.
type RenameOptions struct {
	// mixins
	WorkDoneProgressOptions

	// PrepareProvider renames should be checked and tested before being executed.
	//
	//  version .
	PrepareProvider bool `json:"prepare_provider,omitempty"`
}

// ExecuteCommandOptions the server capabilities of a ExecuteCommandRequest.
type ExecuteCommandOptions struct {
	// mixins
	WorkDoneProgressOptions

	// Commands the commands to be executed on the server.
	Commands []string `json:"commands"`
}

type WorkspaceFoldersServerCapabilities struct {
	// Supported the server has support for workspace folders.
	Supported bool `json:"supported,omitempty"`

	// ChangeNotifications whether the server wants to receive workspace folder change notifications. If a string is provided the string is treated as an ID under which the notification is registered on the client side. The ID can be used to unregister for these events using the `client/unregisterCapability` request.
	ChangeNotifications any/* or */ `json:"change_notifications,omitempty"`
}

// FileOperationOptions options for notifications/requests for user operations on files.
//
// @since 3.16.0
type FileOperationOptions struct {
	// DidCreate the server is interested in receiving didCreateFiles notifications.
	//
	// @since 3.16.0
	DidCreate *FileOperationRegistrationOptions `json:"did_create,omitempty"`

	// WillCreate the server is interested in receiving willCreateFiles requests.
	//
	// @since 3.16.0
	WillCreate *FileOperationRegistrationOptions `json:"will_create,omitempty"`

	// DidRename the server is interested in receiving didRenameFiles notifications.
	//
	// @since 3.16.0
	DidRename *FileOperationRegistrationOptions `json:"did_rename,omitempty"`

	// WillRename the server is interested in receiving willRenameFiles requests.
	//
	// @since 3.16.0
	WillRename *FileOperationRegistrationOptions `json:"will_rename,omitempty"`

	// DidDelete the server is interested in receiving didDeleteFiles file notifications.
	//
	// @since 3.16.0
	DidDelete *FileOperationRegistrationOptions `json:"did_delete,omitempty"`

	// WillDelete the server is interested in receiving willDeleteFiles file requests.
	//
	// @since 3.16.0
	WillDelete *FileOperationRegistrationOptions `json:"will_delete,omitempty"`
}

// WorkspaceOptions defines workspace specific capabilities of the server.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type WorkspaceOptions struct {
	// WorkspaceFolders the server supports workspace folder.
	//
	//
	// @since 3.18.0 proposed
	WorkspaceFolders *WorkspaceFoldersServerCapabilities `json:"workspace_folders,omitempty"`

	// FileOperations the server is interested in notifications/requests for operations on files.
	//
	//
	// @since 3.18.0 proposed
	FileOperations *FileOperationOptions `json:"file_operations,omitempty"`
}

// ServerCapabilities defines the capabilities provided by a language server.
type ServerCapabilities struct {
	// PositionEncoding the position encoding the server picked from the encodings offered by the client via the client capability `general.positionEncodings`. If the client didn't provide any position encodings the only valid value that a server can return is 'utf-16'. If omitted it defaults to 'utf-16'.
	//
	//
	PositionEncoding *PositionEncodingKind `json:"position_encoding,omitempty"`

	// TextDocumentSync defines how text documents are synced. Is either a detailed structure defining each notification or for backwards compatibility the TextDocumentSyncKind number.
	TextDocumentSync any/* or */ `json:"text_document_sync,omitempty"`

	// NotebookDocumentSync defines how notebook documents are synced.
	//
	//
	NotebookDocumentSync any/* or */ `json:"notebook_document_sync,omitempty"`

	// CompletionProvider the server provides completion support.
	CompletionProvider *CompletionOptions `json:"completion_provider,omitempty"`

	// HoverProvider the server provides hover support.
	HoverProvider any/* or */ `json:"hover_provider,omitempty"`

	// SignatureHelpProvider the server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signature_help_provider,omitempty"`

	// DeclarationProvider the server provides Goto Declaration support.
	DeclarationProvider any/* or */ `json:"declaration_provider,omitempty"`

	// DefinitionProvider the server provides goto definition support.
	DefinitionProvider any/* or */ `json:"definition_provider,omitempty"`

	// TypeDefinitionProvider the server provides Goto Type Definition support.
	TypeDefinitionProvider any/* or */ `json:"type_definition_provider,omitempty"`

	// ImplementationProvider the server provides Goto Implementation support.
	ImplementationProvider any/* or */ `json:"implementation_provider,omitempty"`

	// ReferencesProvider the server provides find references support.
	ReferencesProvider any/* or */ `json:"references_provider,omitempty"`

	// DocumentHighlightProvider the server provides document highlight support.
	DocumentHighlightProvider any/* or */ `json:"document_highlight_provider,omitempty"`

	// DocumentSymbolProvider the server provides document symbol support.
	DocumentSymbolProvider any/* or */ `json:"document_symbol_provider,omitempty"`

	// CodeActionProvider the server provides code actions. CodeActionOptions may only be specified if the client states that it supports `codeActionLiteralSupport` in its initial `initialize` request.
	CodeActionProvider any/* or */ `json:"code_action_provider,omitempty"`

	// CodeLensProvider the server provides code lens.
	CodeLensProvider *CodeLensOptions `json:"code_lens_provider,omitempty"`

	// DocumentLinkProvider the server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"document_link_provider,omitempty"`

	// ColorProvider the server provides color provider support.
	ColorProvider any/* or */ `json:"color_provider,omitempty"`

	// WorkspaceSymbolProvider the server provides workspace symbol support.
	WorkspaceSymbolProvider any/* or */ `json:"workspace_symbol_provider,omitempty"`

	// DocumentFormattingProvider the server provides document formatting.
	DocumentFormattingProvider any/* or */ `json:"document_formatting_provider,omitempty"`

	// DocumentRangeFormattingProvider the server provides document range formatting.
	DocumentRangeFormattingProvider any/* or */ `json:"document_range_formatting_provider,omitempty"`

	// DocumentOnTypeFormattingProvider the server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"document_on_type_formatting_provider,omitempty"`

	// RenameProvider the server provides rename support. RenameOptions may only be specified if the client states that it supports `prepareSupport` in its initial `initialize` request.
	RenameProvider any/* or */ `json:"rename_provider,omitempty"`

	// FoldingRangeProvider the server provides folding provider support.
	FoldingRangeProvider any/* or */ `json:"folding_range_provider,omitempty"`

	// SelectionRangeProvider the server provides selection range support.
	SelectionRangeProvider any/* or */ `json:"selection_range_provider,omitempty"`

	// ExecuteCommandProvider the server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"execute_command_provider,omitempty"`

	// CallHierarchyProvider the server provides call hierarchy support.
	//
	//
	CallHierarchyProvider any/* or */ `json:"call_hierarchy_provider,omitempty"`

	// LinkedEditingRangeProvider the server provides linked editing range support.
	//
	//
	LinkedEditingRangeProvider any/* or */ `json:"linked_editing_range_provider,omitempty"`

	// SemanticTokensProvider the server provides semantic tokens support.
	//
	//
	SemanticTokensProvider any/* or */ `json:"semantic_tokens_provider,omitempty"`

	// MonikerProvider the server provides moniker support.
	//
	//
	MonikerProvider any/* or */ `json:"moniker_provider,omitempty"`

	// TypeHierarchyProvider the server provides type hierarchy support.
	//
	//
	TypeHierarchyProvider any/* or */ `json:"type_hierarchy_provider,omitempty"`

	// InlineValueProvider the server provides inline values.
	//
	//
	InlineValueProvider any/* or */ `json:"inline_value_provider,omitempty"`

	// InlayHintProvider the server provides inlay hints.
	//
	//
	InlayHintProvider any/* or */ `json:"inlay_hint_provider,omitempty"`

	// DiagnosticProvider the server has support for pull model diagnostics.
	//
	//
	DiagnosticProvider any/* or */ `json:"diagnostic_provider,omitempty"`

	// InlineCompletionProvider inline completion options used during static registration.
	//
	//  3.18.0
	//
	// Proposed in:.
	InlineCompletionProvider any/* or */ `json:"inline_completion_provider,omitempty"`

	// Workspace workspace specific server capabilities.
	Workspace *WorkspaceOptions `json:"workspace,omitempty"`

	// Experimental experimental server capabilities.
	Experimental any `json:"experimental,omitempty"`
}

// ServerInfo information about the server
//
//	3.15.0
//
//	3.18.0 ServerInfo type name added.
//
// Proposed in:.
//
// @since 3.18.0 ServerInfo type name added. proposed
type ServerInfo struct {
	// Name the name of the server as defined by the server.
	//
	// @since 3.18.0 ServerInfo type name added. proposed
	Name string `json:"name"`

	// Version the server's version as defined by the server.
	//
	// @since 3.18.0 ServerInfo type name added. proposed
	Version string `json:"version,omitempty"`
}

// InitializeResult the result returned from an initialize request.
type InitializeResult struct {
	// Capabilities the capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`

	// ServerInfo information about the server.
	//
	//
	ServerInfo *ServerInfo `json:"server_info,omitempty"`
}

// InitializeError the data type of the ResponseError if the initialize request fails.
type InitializeError struct {
	// Retry indicates whether the client execute the following retry logic: (1) show the message provided by the ResponseError to the user (2) user selects retry or cancel (3) if user selected retry the initialize method is sent again.
	Retry bool `json:"retry"`
}

type InitializedParams struct{}

// DidChangeConfigurationParams the parameters of a change configuration notification.
type DidChangeConfigurationParams struct {
	// Settings the actual changed settings.
	Settings any `json:"settings"`
}

type DidChangeConfigurationRegistrationOptions struct {
	Section any /* or */ `json:"section,omitempty"`
}

// ShowMessageParams the parameters of a notification message.
type ShowMessageParams struct {
	// Type the message type. See MessageType.
	Type MessageType `json:"type"`

	// Message the actual message.
	Message string `json:"message"`
}

type MessageActionItem struct {
	// Title a short title like 'Retry', 'Open Log' etc.
	Title string `json:"title"`
}

type ShowMessageRequestParams struct {
	// Type the message type. See MessageType.
	Type MessageType `json:"type"`

	// Message the actual message.
	Message string `json:"message"`

	// Actions the message action items to present.
	Actions []MessageActionItem `json:"actions,omitempty"`
}

// LogMessageParams the log message parameters.
type LogMessageParams struct {
	// Type the message type. See MessageType.
	Type MessageType `json:"type"`

	// Message the actual message.
	Message string `json:"message"`
}

// DidOpenTextDocumentParams the parameters sent in an open text document notification.
type DidOpenTextDocumentParams struct {
	// TextDocument the document that was opened.
	TextDocument TextDocumentItem `json:"text_document"`
}

// DidChangeTextDocumentParams the change text document notification's parameters.
type DidChangeTextDocumentParams[T TextDocumentContentChangeEvent] struct {
	// TextDocument the document that did change. The version number points to the version after all provided content changes have been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"text_document"`

	// ContentChanges the actual content changes. The content changes describe single state changes to the document. So if there are two content changes c1 (at array index 0) and c2 (at array index 1) for a document in state S then c1 moves the document from S to S' and c2 from S' to S''. So c1 is computed on the state S and c2 is computed on the state S'. To mirror the content of a document using change events use the following approach: - start with the same initial content - apply the 'textDocument/didChange' notifications in the order you receive them. - apply the `TextDocumentContentChangeEvent`s in a single notification in the order  you receive them.
	ContentChanges []T `json:"content_changes"`
}

// TextDocumentChangeRegistrationOptions describe options to be used when registered for text document change events.
type TextDocumentChangeRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions

	// SyncKind how documents are synced to the server.
	SyncKind TextDocumentSyncKind `json:"sync_kind"`
}

// DidCloseTextDocumentParams the parameters sent in a close text document notification.
type DidCloseTextDocumentParams struct {
	// TextDocument the document that was closed.
	TextDocument TextDocumentIdentifier `json:"text_document"`
}

// DidSaveTextDocumentParams the parameters sent in a save text document notification.
type DidSaveTextDocumentParams struct {
	// TextDocument the document that was saved.
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// Text optional the content when saved. Depends on the includeText value when the save notification was requested.
	Text string `json:"text,omitempty"`
}

// TextDocumentSaveRegistrationOptions save registration options.
type TextDocumentSaveRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	SaveOptions
}

// WillSaveTextDocumentParams the parameters sent in a will save text document notification.
type WillSaveTextDocumentParams struct {
	// TextDocument the document that will be saved.
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// Reason the 'TextDocumentSaveReason'.
	Reason TextDocumentSaveReason `json:"reason"`
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

type FileSystemWatcher[T GlobPattern] struct {
	// GlobPattern the glob pattern to watch. See GlobPattern glob pattern for more detail.
	//
	//  3.17.0 support for relative patterns.
	GlobPattern T `json:"glob_pattern"`

	// Kind the kind of events of interest. If omitted it defaults to WatchKind.Create | WatchKind.Change | WatchKind.Delete which is .
	Kind *WatchKind `json:"kind,omitempty"`
}

// DidChangeWatchedFilesRegistrationOptions describe options to be used when registered for text document change events.
type DidChangeWatchedFilesRegistrationOptions[T GlobPattern] struct {
	// Watchers the watchers to register.
	Watchers []FileSystemWatcher[T] `json:"watchers"`
}

// PublishDiagnosticsParams the publish diagnostic notification's parameters.
type PublishDiagnosticsParams struct {
	// URI the URI for which diagnostic information is reported.
	URI DocumentURI `json:"uri"`

	// Version optional the version number of the document the diagnostics are published for.
	//
	//
	Version int32 `json:"version,omitempty"`

	// Diagnostics an array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
}

// CompletionContext contains additional information about the context in which a completion request is triggered.
type CompletionContext struct {
	// TriggerKind how the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"trigger_kind"`

	// TriggerCharacter the trigger character (a single character) that has trigger code complete. Is undefined if `triggerKind !== CompletionTriggerKind.TriggerCharacter`.
	TriggerCharacter string `json:"trigger_character,omitempty"`
}

// CompletionParams completion parameters.
type CompletionParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// Context the completion context. This is only available it the client specifies to send this using the client capability `textDocument.completion.contextSupport === true`.
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
	NewText string `json:"new_text"`

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
	//
	//
	LabelDetails *CompletionItemLabelDetails `json:"label_details,omitempty"`

	// Kind the kind of this completion item. Based of the kind an icon is chosen by the editor.
	Kind *CompletionItemKind `json:"kind,omitempty"`

	// Tags tags for this completion item.
	//
	//
	Tags []CompletionItemTag `json:"tags,omitempty"`

	// Detail a human-readable string with additional information about this item, like type or symbol information.
	Detail string `json:"detail,omitempty"`

	// Documentation a human-readable string that represents a doc-comment.
	Documentation any/* or */ `json:"documentation,omitempty"`

	// Deprecated indicates if this item is deprecated.
	//
	// Deprecated: Use `tags` instead.
	Deprecated bool `json:"deprecated,omitempty"`

	// Preselect select this item when showing. *Note* that only one completion item can be selected and that the tool / client decides which item that is. The rule is that the *first* item of those that match best is selected.
	Preselect bool `json:"preselect,omitempty"`

	// SortText a string that should be used when comparing this item with other items. When `falsy` the CompletionItem.label label is used.
	SortText string `json:"sort_text,omitempty"`

	// FilterText a string that should be used when filtering a set of completion items. When `falsy` the CompletionItem.label label is used.
	FilterText string `json:"filter_text,omitempty"`

	// InsertText a string that should be inserted into a document when selecting this completion. When `falsy` the CompletionItem.label label is used. The `insertText` is subject to interpretation by the client side. Some tools might not take the string literally. For example VS Code when code complete is requested in this example `con<cursor position>` and a completion item with an `insertText` of `console` is provided it will only insert `sole`. Therefore it is recommended to use `textEdit` instead since it avoids additional client side interpretation.
	InsertText string `json:"insert_text,omitempty"`

	// InsertTextFormat the format of the insert text. The format applies to both the `insertText` property and the `newText` property of a provided `textEdit`. If omitted defaults to `InsertTextFormat.PlainText`. Please note that the insertTextFormat doesn't apply to `additionalTextEdits`.
	InsertTextFormat *InsertTextFormat `json:"insert_text_format,omitempty"`

	// InsertTextMode how whitespace and indentation is handled during completion item insertion. If not provided the clients default value depends on the `textDocument.completion.insertTextMode` client capability.
	//
	//
	InsertTextMode *InsertTextMode `json:"insert_text_mode,omitempty"`

	// TextEdit an TextEdit edit which is applied to a document when selecting this completion. When an edit is provided the value of CompletionItem.insertText insertText is ignored. Most editors support two different operations when accepting a completion item. One is to insert a completion text and the other is to replace an existing text with a completion text. Since this can usually not be predetermined by a server it can report both ranges. Clients need to signal support for `InsertReplaceEdits` via the `textDocument.completion.insertReplaceSupport` client capability property. *Note 1:* The text edit's range as well as both ranges from an insert replace edit must be a [single line] and they must contain the position at which completion has been requested. *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range must be a prefix of the edit's replace range, that means it must be contained and starting at the same position.
	//
	//  3.16.0 additional type `InsertReplaceEdit`.
	TextEdit any/* or */ `json:"text_edit,omitempty"`

	// TextEditText the edit text used if the completion item is part of a CompletionList and CompletionList defines an item default for the text edit range. Clients will only honor this property if they opt into completion list item defaults using the capability `completionList.itemDefaults`. If not provided and a list's default range is provided the label property is used as a text.
	//
	//
	TextEditText string `json:"text_edit_text,omitempty"`

	// AdditionalTextEdits an optional array of additional TextEdit text edits that are applied when selecting this completion. Edits must not overlap (including the same insert position) with the main CompletionItem.textEdit edit nor with themselves. Additional text edits should be used to change text unrelated to the current cursor position (for example adding an import statement at the top of the file if the completion item will insert an unqualified type).
	AdditionalTextEdits []TextEdit `json:"additional_text_edits,omitempty"`

	// CommitCharacters an optional set of characters that when pressed while this completion is active will accept it first and then type that character. *Note* that all commit characters should have `length=1` and that superfluous characters will be ignored.
	CommitCharacters []string `json:"commit_characters,omitempty"`

	// Command an optional Command command that is executed *after* inserting this completion. *Note* that additional modifications to the current document should be described with the CompletionItem.additionalTextEdits additionalTextEdits-property.
	Command *Command `json:"command,omitempty"`

	// Data a data entry field that is preserved on a completion item between a CompletionRequest and a CompletionResolveRequest.
	Data any `json:"data,omitempty"`
}

// EditRangeWithInsertReplace edit range variant that includes ranges for insert and replace operations.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type EditRangeWithInsertReplace struct {
	// @since 3.18.0 proposed
	Insert Range `json:"insert"`

	// @since 3.18.0 proposed
	Replace Range `json:"replace"`
}

// CompletionItemDefaults in many cases the items of an actual completion result share the same value for properties like `commitCharacters` or the range of a text edit. A completion list can therefore define item defaults which will be used if a completion item itself doesn't specify the value. If a completion list specifies a default value and a completion item also specifies a corresponding value the one from the item is used. Servers are only allowed to return default values if the client signals support for this via the `completionList.itemDefaults` capability.
//
// @since 3.17.0
type CompletionItemDefaults struct {
	// CommitCharacters a default commit character set.
	//
	//
	// @since 3.17.0
	CommitCharacters []string `json:"commit_characters,omitempty"`

	// EditRange a default edit range.
	//
	//
	// @since 3.17.0
	EditRange any/* or */ `json:"edit_range,omitempty"`

	// InsertTextFormat a default insert text format.
	//
	//
	// @since 3.17.0
	InsertTextFormat *InsertTextFormat `json:"insert_text_format,omitempty"`

	// InsertTextMode a default insert text mode.
	//
	//
	// @since 3.17.0
	InsertTextMode *InsertTextMode `json:"insert_text_mode,omitempty"`

	// Data a default data value.
	//
	//
	// @since 3.17.0
	Data any `json:"data,omitempty"`
}

// CompletionList represents a collection of CompletionItem completion items to be presented in the editor.
type CompletionList struct {
	// IsIncomplete this list it not complete. Further typing results in recomputing this list. Recomputed lists have all their items replaced (not appended) in the incomplete completion sessions.
	IsIncomplete bool `json:"is_incomplete"`

	// ItemDefaults in many cases the items of an actual completion result share the same value for properties like `commitCharacters` or the range of a text edit. A completion list can therefore define item defaults which will be used if a completion item itself doesn't specify the value. If a completion list specifies a default value and a completion item also specifies a corresponding value the one from the item is used. Servers are only allowed to return default values if the client signals support for this via the `completionList.itemDefaults` capability.
	//
	//
	ItemDefaults *CompletionItemDefaults `json:"item_defaults,omitempty"`

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
type HoverParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
}

// Hover the result of a hover request.
type Hover struct {
	// Contents the hover's content.
	Contents any/* or */ `json:"contents"`

	// Range an optional range inside the text document that is used to visualize the hover, e.g. by changing the background color.
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
	// Label the label of this parameter information. Either a string or an inclusive start and exclusive end offsets within its containing signature label. (see SignatureInformation.label). The offsets are based on a UTF-16 string representation as `Position` and `Range` does. To avoid ambiguities a server should use the [start, end] offset value instead of using a substring. Whether a client support this is controlled via `labelOffsetSupport` client capability. *Note*: a label of type string should be a substring of its containing signature label. Its intended use case is to highlight the parameter label part in the `SignatureInformation.label`.
	Label any/* or */ `json:"label"`

	// Documentation the human-readable doc-comment of this parameter. Will be shown in the UI but can be omitted.
	Documentation any/* or */ `json:"documentation,omitempty"`
}

// SignatureInformation represents the signature of something callable. A signature can have a label, like a function-name, a doc-comment, and a set of parameters.
type SignatureInformation struct {
	// Label the label of this signature. Will be shown in the UI.
	Label string `json:"label"`

	// Documentation the human-readable doc-comment of this signature. Will be shown in the UI but can be omitted.
	Documentation any/* or */ `json:"documentation,omitempty"`

	// Parameters the parameters of this signature.
	Parameters []ParameterInformation `json:"parameters,omitempty"`

	// ActiveParameter the index of the active parameter. If `null`, no parameter of the signature is active (for example a named argument that does not match any declared parameters). This is only valid if the client specifies the client capability `textDocument.signatureHelp.noActiveParameterSupport === true` If provided (or `null`), this is used in place of `SignatureHelp.activeParameter`.
	//
	//
	ActiveParameter any/* or */ `json:"active_parameter,omitempty"`
}

// SignatureHelp signature help represents the signature of something callable. There can be multiple signature but only one active and only one active parameter.
type SignatureHelp struct {
	// Signatures one or more signatures.
	Signatures []SignatureInformation `json:"signatures"`

	// ActiveSignature the active signature. If omitted or the value lies outside the range of `signatures` the value defaults to zero or is ignored if the `SignatureHelp` has no signatures. Whenever possible implementors should make an active decision about the active signature and shouldn't rely on a default value. In future version of the protocol this property might become mandatory to better express this.
	ActiveSignature uint32 `json:"active_signature,omitempty"`

	// ActiveParameter the active parameter of the active signature. If `null`, no parameter of the signature is active (for example a named argument that does not match any declared parameters). This is only valid if the client specifies the client capability `textDocument.signatureHelp.noActiveParameterSupport === true` If omitted or the value lies outside the range of `signatures[activeSignature].parameters` defaults to 0 if the active signature has parameters. If the active signature has no parameters it is ignored. In future version of the protocol this property might become mandatory (but still nullable) to better express the active parameter if the active signature does have any.
	ActiveParameter any/* or */ `json:"active_parameter,omitempty"`
}

// SignatureHelpContext additional information about the context in which a signature help request was triggered.
//
// @since 3.15.0
type SignatureHelpContext struct {
	// TriggerKind action that caused signature help to be triggered.
	//
	// @since 3.15.0
	TriggerKind SignatureHelpTriggerKind `json:"trigger_kind"`

	// TriggerCharacter character that caused signature help to be triggered. This is undefined when `triggerKind !== SignatureHelpTriggerKind.TriggerCharacter`.
	//
	// @since 3.15.0
	TriggerCharacter string `json:"trigger_character,omitempty"`

	// IsRetrigger `true` if signature help was already showing when it was triggered. Retriggers occurs when the signature help is already active and can be caused by actions such as typing a trigger character, a cursor move, or document content changes.
	//
	// @since 3.15.0
	IsRetrigger bool `json:"is_retrigger"`

	// ActiveSignatureHelp the currently active `SignatureHelp`. The `activeSignatureHelp` has its `SignatureHelp.activeSignature` field updated based on the user navigating through available signatures.
	//
	// @since 3.15.0
	ActiveSignatureHelp *SignatureHelp `json:"active_signature_help,omitempty"`
}

// SignatureHelpParams parameters for a SignatureHelpRequest.
type SignatureHelpParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT

	// Context the signature help context. This is only available if the client specifies to send this using the client capability `textDocument.signatureHelp.contextSupport === true`
	//
	//
	Context *SignatureHelpContext `json:"context,omitempty"`
}

// SignatureHelpRegistrationOptions registration options for a SignatureHelpRequest.
type SignatureHelpRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	SignatureHelpOptions
}

// DefinitionParams parameters for a DefinitionRequest.
type DefinitionParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT
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
	IncludeDeclaration bool `json:"include_declaration"`
}

// ReferenceParams parameters for a ReferencesRequest.
type ReferenceParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	Context ReferenceContext `json:"context"`
}

// ReferenceRegistrationOptions registration options for a ReferencesRequest.
type ReferenceRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	ReferenceOptions
}

// DocumentHighlightParams parameters for a DocumentHighlightRequest.
type DocumentHighlightParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT
}

// DocumentHighlight a document highlight is a range inside a text document which deserves special attention. Usually a document highlight is visualized by changing the background color of its range.
type DocumentHighlight struct {
	// Range the range this highlight applies to.
	Range Range `json:"range"`

	// Kind the highlight kind, default is DocumentHighlightKind.Text text.
	Kind *DocumentHighlightKind `json:"kind,omitempty"`
}

// DocumentHighlightRegistrationOptions registration options for a DocumentHighlightRequest.
type DocumentHighlightRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	DocumentHighlightOptions
}

// DocumentSymbolParams parameters for a DocumentSymbolRequest.
type DocumentSymbolParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the text document.
	TextDocument TextDocumentIdentifier `json:"text_document"`
}

// BaseSymbolInformation a base for all symbol information.
type BaseSymbolInformation struct {
	// Name the name of this symbol.
	Name string `json:"name"`

	// Kind the kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags tags for this symbol.
	//
	//
	Tags []SymbolTag `json:"tags,omitempty"`

	// ContainerName the name of the symbol containing this symbol. This information is for user interface purposes (e.g. to render a qualifier in the user interface if necessary). It can't be used to re-infer a hierarchy for the document symbols.
	ContainerName string `json:"container_name,omitempty"`
}

// SymbolInformation represents information about programming constructs like variables, classes, interfaces etc.
type SymbolInformation struct {
	// extends
	BaseSymbolInformation

	// Deprecated indicates if this symbol is deprecated.
	//
	// Deprecated: Use tags instead.
	Deprecated bool `json:"deprecated,omitempty"`

	// Location the location of this symbol. The location's range is used by a tool to reveal the location in the editor. If the symbol is selected in the tool the range's start information is used to position the cursor. So the range usually spans more than the actual symbol's name and does normally include things like visibility modifiers. The range doesn't have to denote a node range in the sense of an abstract syntax tree. It can therefore not be used to re-construct a hierarchy of the symbols.
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
	//
	//
	Tags []SymbolTag `json:"tags,omitempty"`

	// Deprecated indicates if this symbol is deprecated.
	//
	// Deprecated: Use tags instead.
	Deprecated bool `json:"deprecated,omitempty"`

	// Range the range enclosing this symbol not including leading/trailing whitespace but everything else like comments. This information is typically used to determine if the clients cursor is inside the symbol to reveal in the symbol in the UI.
	Range Range `json:"range"`

	// SelectionRange the range that should be selected and revealed when this symbol is being picked, e.g the name of a function. Must be contained by the `range`.
	SelectionRange Range `json:"selection_range"`

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
	//
	//
	TriggerKind *CodeActionTriggerKind `json:"trigger_kind,omitempty"`
}

// CodeActionParams the parameters of a CodeActionRequest.
type CodeActionParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the document in which the command was invoked.
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// Range the range for which the command was invoked.
	Range Range `json:"range"`

	// Context context carrying additional information.
	Context CodeActionContext `json:"context"`
}

// CodeActionDisabled captures why the code action is currently disabled.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type CodeActionDisabled struct {
	// Reason human readable description of why the code action is currently disabled. This is displayed in the code actions UI.
	//
	// @since 3.18.0 proposed
	Reason string `json:"reason"`
}

// CodeAction a code action represents a change that can be performed in code, e.g. to fix a problem or to refactor code. A CodeAction must set either `edit` and/or a `command`. If both are supplied, the `edit` is applied first, then the `command` is executed.
type CodeAction struct {
	// Title a short, human-readable, title for this code action.
	Title string `json:"title"`

	// Kind the kind of the code action. Used to filter code actions.
	Kind *CodeActionKind `json:"kind,omitempty"`

	// Diagnostics the diagnostics that this code action resolves.
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`

	// IsPreferred marks this as a preferred action. Preferred actions are used by the `auto fix` command and can be targeted by keybindings. A quick fix should be marked preferred if it properly addresses the underlying error. A refactoring should be marked preferred if it is the most reasonable choice of actions to take.
	//
	//
	IsPreferred bool `json:"is_preferred,omitempty"`

	// Disabled marks that the code action cannot currently be applied. Clients should follow the following guidelines regarding disabled code actions:  - Disabled code actions are not shown in automatic [lightbulbs](https://code.visualstudio.com/docs/editor/editingevolved#_code-action)   code action menus.  - Disabled actions are shown as faded out in the code action menu when the user requests a more specific type   of code action, such as refactorings.  - If the user has a [keybinding](https://code.visualstudio.com/docs/editor/refactoring#_keybindings-for-code-actions)   that auto applies a code action and only disabled code actions are returned, the client should show the user an   error message with `reason` in the editor.
	//
	//
	Disabled *CodeActionDisabled `json:"disabled,omitempty"`

	// Edit the workspace edit this code action performs.
	Edit *WorkspaceEdit `json:"edit,omitempty"`

	// Command a command this code action executes. If a code action provides an edit and a command, first the edit is executed and then the command.
	Command *Command `json:"command,omitempty"`

	// Data a data entry field that is preserved on a code action between a `textDocument/codeAction` and a `codeAction/resolve` request.
	//
	//
	Data any `json:"data,omitempty"`
}

// CodeActionRegistrationOptions registration options for a CodeActionRequest.
type CodeActionRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	CodeActionOptions
}

// WorkspaceSymbolParams the parameters of a WorkspaceSymbolRequest.
type WorkspaceSymbolParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// Query a query string to filter symbols by. Clients may send an empty string here to request all symbols.
	Query string `json:"query"`
}

// LocationURIOnly location with only uri and does not include range.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type LocationURIOnly struct {
	// @since 3.18.0 proposed
	URI DocumentURI `json:"uri"`
}

// WorkspaceSymbol a special workspace symbol that supports locations without a range. See also SymbolInformation.
//
// @since 3.17.0
type WorkspaceSymbol struct {
	// extends
	BaseSymbolInformation

	// Location the location of the symbol. Whether a server is allowed to return a location without a range depends on the client capability `workspace.symbol.resolveSupport`. See SymbolInformation#location for more details.
	//
	// @since 3.17.0
	Location any/* or */ `json:"location"`

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

// CodeLensParams the parameters of a CodeLensRequest.
type CodeLensParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the document to request code lens for.
	TextDocument TextDocumentIdentifier `json:"text_document"`
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
type DocumentLinkParams[T ProgressToken, TT WorkDoneProgressParams[T], TTT PartialResultParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT
	PartialResultParams    TTT

	// TextDocument the document to provide document links for.
	TextDocument TextDocumentIdentifier `json:"text_document"`
}

// DocumentLink a document link is a range in a text document that links to an internal or external resource, like another text document or a web site.
type DocumentLink struct {
	// Range the range this link applies to.
	Range Range `json:"range"`

	// Target the uri this link points to. If missing a resolve request is sent later.
	Target uri.URI `json:"target,omitempty"`

	// Tooltip the tooltip text when you hover over this link. If a tooltip is provided, is will be displayed in a string that includes instructions on how to trigger the link, such as `{0} (ctrl + click)`. The specific instructions vary depending on OS, user settings, and localization.
	//
	//
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
	TabSize uint32 `json:"tab_size"`

	// InsertSpaces prefer spaces over tabs.
	InsertSpaces bool `json:"insert_spaces"`

	// TrimTrailingWhitespace trim trailing whitespace on a line.
	//
	//
	TrimTrailingWhitespace bool `json:"trim_trailing_whitespace,omitempty"`

	// InsertFinalNewline insert a newline character at the end of the file if one does not exist.
	//
	//
	InsertFinalNewline bool `json:"insert_final_newline,omitempty"`

	// TrimFinalNewlines trim all newlines after the final newline at the end of the file.
	//
	//
	TrimFinalNewlines bool `json:"trim_final_newlines,omitempty"`
}

// DocumentFormattingParams the parameters of a DocumentFormattingRequest.
type DocumentFormattingParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT

	// TextDocument the document to format.
	TextDocument TextDocumentIdentifier `json:"text_document"`

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
type DocumentRangeFormattingParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT

	// TextDocument the document to format.
	TextDocument TextDocumentIdentifier `json:"text_document"`

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

// DocumentRangesFormattingParams the parameters of a DocumentRangesFormattingRequest.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type DocumentRangesFormattingParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT

	// TextDocument the document to format.
	//
	// @since 3.18.0 proposed
	TextDocument TextDocumentIdentifier `json:"text_document"`

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
	TextDocument TextDocumentIdentifier `json:"text_document"`

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
type RenameParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT

	// TextDocument the document to rename.
	TextDocument TextDocumentIdentifier `json:"text_document"`

	// Position the position at which this request was sent.
	Position Position `json:"position"`

	// NewName the new name of the symbol. If the given name is not valid the request must return a ResponseError with an appropriate message set.
	NewName string `json:"new_name"`
}

// RenameRegistrationOptions registration options for a RenameRequest.
type RenameRegistrationOptions struct {
	// extends
	TextDocumentRegistrationOptions
	RenameOptions
}

type PrepareRenameParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// extends
	TextDocumentPositionParams
	// mixins
	WorkDoneProgressParams TT
}

// ExecuteCommandParams the parameters of a ExecuteCommandRequest.
type ExecuteCommandParams[T ProgressToken, TT WorkDoneProgressParams[T]] struct {
	// mixins
	WorkDoneProgressParams TT

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

// ApplyWorkspaceEditParams the parameters passed via an apply workspace edit request.
type ApplyWorkspaceEditParams struct {
	// Label an optional label of the workspace edit. This label is presented in the user interface for example on an undo stack to undo the workspace edit.
	Label string `json:"label,omitempty"`

	// Edit the edits to apply.
	Edit WorkspaceEdit `json:"edit"`
}

// ApplyWorkspaceEditResult the result returned from the apply workspace edit request.
//
//	3.17 renamed from ApplyWorkspaceEditResponse.
//
// @since 3.17 renamed from ApplyWorkspaceEditResponse
type ApplyWorkspaceEditResult struct {
	// Applied indicates whether the edit was applied or not.
	//
	// @since 3.17 renamed from ApplyWorkspaceEditResponse
	Applied bool `json:"applied"`

	// FailureReason an optional textual description for why the edit was not applied. This may be used by the server for diagnostic logging or to provide a suitable error for a request that triggered the edit.
	//
	// @since 3.17 renamed from ApplyWorkspaceEditResponse
	FailureReason string `json:"failure_reason,omitempty"`

	// FailedChange depending on the client's failure handling strategy `failedChange` might contain the index of the change that failed. This property is only available if the client signals a `failureHandlingStrategy` in its client capabilities.
	//
	// @since 3.17 renamed from ApplyWorkspaceEditResponse
	FailedChange uint32 `json:"failed_change,omitempty"`
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

type SetTraceParams struct {
	Value TraceValue `json:"value"`
}

type LogTraceParams struct {
	Message string `json:"message"`

	Verbose string `json:"verbose,omitempty"`
}

type CancelParams struct {
	// ID the request id to cancel.
	ID any /* or */ `json:"id"`
}

type ProgressParams[T ProgressToken] struct {
	// Token the progress token provided by the client or server.
	Token T `json:"token"`

	// Value the progress data.
	Value any `json:"value"`
}

type WorkDoneProgressParams[T ProgressToken] struct {
	// WorkDoneToken an optional token that a server can use to report work done progress.
	WorkDoneToken *T `json:"work_done_token,omitempty"`
}

type PartialResultParams[T ProgressToken] struct {
	// PartialResultToken an optional token that a server can use to report partial results (e.g. streaming) to the client.
	PartialResultToken *T `json:"partial_result_token,omitempty"`
}

// LocationLink represents the connection of two locations. Provides additional metadata over normal Location locations, including an origin range.
type LocationLink struct {
	// OriginSelectionRange span of the origin of this link. Used as the underlined span for mouse interaction. Defaults to the word range at the definition position.
	OriginSelectionRange *Range `json:"origin_selection_range,omitempty"`

	// TargetURI the target resource identifier of this link.
	TargetURI DocumentURI `json:"target_uri"`

	// TargetRange the full target range of this link. If the target for example is a symbol then target range is the range enclosing this symbol not including leading/trailing whitespace but everything else like comments. This information is typically used to highlight the range in the editor.
	TargetRange Range `json:"target_range"`

	// TargetSelectionRange the range that should be selected and revealed when this link is being followed, e.g the name of a function. Must be contained by the `targetRange`. See also `DocumentSymbol#range`.
	TargetSelectionRange Range `json:"target_selection_range"`
}

// StaticRegistrationOptions static registration options to be returned in the initialize request.
type StaticRegistrationOptions struct {
	// ID the id used to register the request. The id can be used to deregister the request again. See also Registration#id.
	ID string `json:"id,omitempty"`
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

// InlineValueVariableLookup provide inline value through a variable lookup. If only a range is specified, the variable name will be extracted from the underlying document. An optional variable name can be used to override the extracted name.
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
	VariableName string `json:"variable_name,omitempty"`

	// CaseSensitiveLookup how to perform the lookup.
	//
	// @since 3.17.0
	CaseSensitiveLookup bool `json:"case_sensitive_lookup"`
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
	//
	//
	// @since 3.17.0
	RelatedDocuments map[string]any `json:"related_documents,omitempty"`
}

// RelatedUnchangedDocumentDiagnosticReport an unchanged diagnostic report with a set of related documents.
//
// @since 3.17.0
type RelatedUnchangedDocumentDiagnosticReport struct {
	// extends
	UnchangedDocumentDiagnosticReport

	// RelatedDocuments diagnostics of related documents. This information is useful in programming languages where code in a file A can generate diagnostics in a file B which A depends on. An example of such a language is C/C++ where marco definitions in a file a.cpp and result in errors in a header file b.hpp.
	//
	//
	// @since 3.17.0
	RelatedDocuments map[string]any `json:"related_documents,omitempty"`
}

// PrepareRenamePlaceholder.
//
// @since 3.18.0 proposed
type PrepareRenamePlaceholder struct {
	// @since 3.18.0 proposed
	Range Range `json:"range"`

	// @since 3.18.0 proposed
	Placeholder string `json:"placeholder"`
}

// PrepareRenameDefaultBehavior.
//
// @since 3.18.0 proposed
type PrepareRenameDefaultBehavior struct {
	// @since 3.18.0 proposed
	DefaultBehavior bool `json:"default_behavior"`
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
	Version any/* or */ `json:"version"`
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
	Version any/* or */ `json:"version"`
}

// TextDocumentContentChangePartial.
//
// @since 3.18.0 proposed
type TextDocumentContentChangePartial struct {
	// Range the range of the document that changed.
	//
	// @since 3.18.0 proposed
	Range Range `json:"range"`

	// RangeLength the optional length of the range that got replaced.
	//
	// Deprecated: use range instead.
	//
	// @since 3.18.0 proposed
	RangeLength uint32 `json:"range_length,omitempty"`

	// Text the new text for the provided range.
	//
	// @since 3.18.0 proposed
	Text string `json:"text"`
}

// TextDocumentContentChangeWholeDocument.
//
// @since 3.18.0 proposed
type TextDocumentContentChangeWholeDocument struct {
	// Text the new text of the whole document.
	//
	// @since 3.18.0 proposed
	Text string `json:"text"`
}

// MarkedStringWithLanguage.
//
// @since 3.18.0 proposed
type MarkedStringWithLanguage struct {
	// @since 3.18.0 proposed
	Language string `json:"language"`

	// @since 3.18.0 proposed
	Value string `json:"value"`
}

// NotebookCellTextDocumentFilter a notebook cell text document filter denotes a cell text document by different properties.
//
// @since 3.17.0
type NotebookCellTextDocumentFilter struct {
	// Notebook a filter that matches against the notebook containing the notebook cell. If a string value is provided it matches against the notebook type. '*' matches every notebook.
	//
	// @since 3.17.0
	Notebook any/* or */ `json:"notebook"`

	// Language a language id like `python`. Will be matched against the language id of the notebook cell document. '*' matches every language.
	//
	// @since 3.17.0
	Language string `json:"language,omitempty"`
}

// RelativePattern a relative pattern is a helper to construct glob patterns that are matched relatively to a base URI. The common value for a `baseUri` is a workspace folder root, but it can be another absolute URI as well.
//
// @since 3.17.0
type RelativePattern struct {
	// BaseURI a workspace folder or a base URI to which this pattern will be matched against relatively.
	//
	// @since 3.17.0
	BaseURI any/* or */ `json:"base_uri"`

	// Pattern the actual glob pattern;.
	//
	// @since 3.17.0
	Pattern Pattern `json:"pattern"`
}

// TextDocumentFilterLanguage a document filter where `language` is required field.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type TextDocumentFilterLanguage struct {
	// Language a language id, like `typescript`.
	//
	// @since 3.18.0 proposed
	Language string `json:"language"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0 proposed
	Scheme string `json:"scheme,omitempty"`

	// Pattern a glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples.
	//
	// @since 3.18.0 proposed
	Pattern string `json:"pattern,omitempty"`
}

// TextDocumentFilterScheme a document filter where `scheme` is required field.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type TextDocumentFilterScheme struct {
	// Language a language id, like `typescript`.
	//
	// @since 3.18.0 proposed
	Language string `json:"language,omitempty"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0 proposed
	Scheme string `json:"scheme"`

	// Pattern a glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples.
	//
	// @since 3.18.0 proposed
	Pattern string `json:"pattern,omitempty"`
}

// TextDocumentFilterPattern a document filter where `pattern` is required field.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type TextDocumentFilterPattern struct {
	// Language a language id, like `typescript`.
	//
	// @since 3.18.0 proposed
	Language string `json:"language,omitempty"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0 proposed
	Scheme string `json:"scheme,omitempty"`

	// Pattern a glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples.
	//
	// @since 3.18.0 proposed
	Pattern string `json:"pattern"`
}

// NotebookDocumentFilterNotebookType a notebook document filter where `notebookType` is required field.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterNotebookType struct {
	// NotebookType the type of the enclosing notebook.
	//
	// @since 3.18.0 proposed
	NotebookType string `json:"notebook_type"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0 proposed
	Scheme string `json:"scheme,omitempty"`

	// Pattern a glob pattern.
	//
	// @since 3.18.0 proposed
	Pattern string `json:"pattern,omitempty"`
}

// NotebookDocumentFilterScheme a notebook document filter where `scheme` is required field.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterScheme struct {
	// NotebookType the type of the enclosing notebook.
	//
	// @since 3.18.0 proposed
	NotebookType string `json:"notebook_type,omitempty"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0 proposed
	Scheme string `json:"scheme"`

	// Pattern a glob pattern.
	//
	// @since 3.18.0 proposed
	Pattern string `json:"pattern,omitempty"`
}

// NotebookDocumentFilterPattern a notebook document filter where `pattern` is required field.
//
//	3.18.0
//
// Proposed in:.
//
// @since 3.18.0 proposed
type NotebookDocumentFilterPattern struct {
	// NotebookType the type of the enclosing notebook.
	//
	// @since 3.18.0 proposed
	NotebookType string `json:"notebook_type,omitempty"`

	// Scheme a Uri Uri.scheme scheme, like `file` or `untitled`.
	//
	// @since 3.18.0 proposed
	Scheme string `json:"scheme,omitempty"`

	// Pattern a glob pattern.
	//
	// @since 3.18.0 proposed
	Pattern string `json:"pattern"`
}
