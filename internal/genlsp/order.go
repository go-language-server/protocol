// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import "strings"

// specFeature is one entry of the LSP specification's table of contents. Types
// are assigned to a feature by longest-matching name prefix; methods by exact
// membership. The slice order below is the spec's document order, so a stable
// sort on feature rank reorders feature blocks into spec order while preserving
// each block's original (reference) order from the meta-model. file is the
// generated source file the feature's declarations are written to: major
// sections share one file, while each language feature gets its own file.
type specFeature struct {
	file     string
	methods  []string
	prefixes []string
}

// basicStructuresFile holds shared/foundational declarations (the spec's "Basic
// JSON Structures"), i.e. every type matching no feature prefix (rank 0).
const basicStructuresFile = "basic_structures.go"

// specFeatures lists the spec sections in document order.
var specFeatures = []specFeature{
	// ---- Lifecycle Messages ----
	{
		file:    "lifecycle.go",
		methods: []string{"initialize", "initialized"},
		prefixes: []string{
			"InitializeParams", "InitializeResult", "InitializeError", "InitializedParams", "_Initialize",
			"ServerCapabilities", "ClientCapabilities", "ServerInfo", "ClientInfo",
			"TextDocumentClientCapabilities", "NotebookDocumentClientCapabilities",
			"WorkspaceClientCapabilities", "WindowClientCapabilities", "GeneralClientCapabilities",
			"WorkspaceFoldersServerCapabilities", "WorkspaceOptions", "StaleRequestSupportOptions",
		},
	},
	{file: "lifecycle.go", methods: []string{"client/registerCapability", "client/unregisterCapability"}, prefixes: []string{"Registration", "Unregistration"}},
	{file: "lifecycle.go", methods: []string{"$/setTrace", "$/logTrace"}, prefixes: []string{"SetTrace", "LogTrace"}},
	{file: "lifecycle.go", methods: []string{"shutdown"}, prefixes: nil},
	{file: "lifecycle.go", methods: []string{"exit"}, prefixes: nil},

	// ---- Document Synchronization ----
	{file: "document_synchronization.go", methods: []string{"textDocument/didOpen"}, prefixes: []string{"DidOpenTextDocument"}},
	{file: "document_synchronization.go", methods: []string{"textDocument/didChange"}, prefixes: []string{"DidChangeTextDocument", "TextDocumentChangeRegistrationOptions"}},
	{file: "document_synchronization.go", methods: []string{"textDocument/willSave"}, prefixes: []string{"WillSaveTextDocument"}},
	{file: "document_synchronization.go", methods: []string{"textDocument/willSaveWaitUntil"}, prefixes: nil},
	{file: "document_synchronization.go", methods: []string{"textDocument/didSave"}, prefixes: []string{"DidSaveTextDocument", "SaveOptions", "TextDocumentSaveRegistrationOptions", "TextDocumentSaveReason"}},
	{file: "document_synchronization.go", methods: []string{"textDocument/didClose"}, prefixes: []string{"DidCloseTextDocument"}},
	{file: "document_synchronization.go", methods: nil, prefixes: []string{"TextDocumentSync"}},
	{
		file:    "document_synchronization.go",
		methods: []string{"notebookDocument/didOpen", "notebookDocument/didChange", "notebookDocument/didSave", "notebookDocument/didClose"},
		prefixes: []string{
			"NotebookDocument", "NotebookCell", "DidOpenNotebook", "DidChangeNotebook",
			"DidSaveNotebook", "DidCloseNotebook", "ExecutionSummary", "VersionedNotebookDocumentIdentifier",
		},
	},

	// ---- Language Features ----
	{file: "declaration.go", methods: []string{"textDocument/declaration"}, prefixes: []string{"Declaration"}},
	{file: "definition.go", methods: []string{"textDocument/definition"}, prefixes: []string{"Definition"}},
	{file: "type_definition.go", methods: []string{"textDocument/typeDefinition"}, prefixes: []string{"TypeDefinition"}},
	{file: "implementation.go", methods: []string{"textDocument/implementation"}, prefixes: []string{"Implementation"}},
	{file: "references.go", methods: []string{"textDocument/references"}, prefixes: []string{"Reference"}},
	{file: "call_hierarchy.go", methods: []string{"textDocument/prepareCallHierarchy", "callHierarchy/incomingCalls", "callHierarchy/outgoingCalls"}, prefixes: []string{"CallHierarchy"}},
	{file: "type_hierarchy.go", methods: []string{"textDocument/prepareTypeHierarchy", "typeHierarchy/supertypes", "typeHierarchy/subtypes"}, prefixes: []string{"TypeHierarchy"}},
	{file: "document_highlight.go", methods: []string{"textDocument/documentHighlight"}, prefixes: []string{"DocumentHighlight"}},
	{file: "document_link.go", methods: []string{"textDocument/documentLink", "documentLink/resolve"}, prefixes: []string{"DocumentLink"}},
	{file: "hover.go", methods: []string{"textDocument/hover"}, prefixes: []string{"Hover"}},
	{file: "code_lens.go", methods: []string{"textDocument/codeLens", "codeLens/resolve", "workspace/codeLens/refresh"}, prefixes: []string{"CodeLens", "ClientCodeLens"}},
	{file: "folding_range.go", methods: []string{"textDocument/foldingRange", "workspace/foldingRange/refresh"}, prefixes: []string{"FoldingRange", "ClientFoldingRange"}},
	{file: "selection_range.go", methods: []string{"textDocument/selectionRange"}, prefixes: []string{"SelectionRange"}},
	{file: "document_symbol.go", methods: []string{"textDocument/documentSymbol"}, prefixes: []string{"DocumentSymbol"}},
	{file: "semantic_tokens.go", methods: []string{"textDocument/semanticTokens/full", "textDocument/semanticTokens/full/delta", "textDocument/semanticTokens/range", "workspace/semanticTokens/refresh"}, prefixes: []string{"SemanticTokens", "ClientSemanticTokens", "SemanticToken"}},
	{file: "inline_value.go", methods: []string{"textDocument/inlineValue", "workspace/inlineValue/refresh"}, prefixes: []string{"InlineValue"}},
	{file: "inlay_hint.go", methods: []string{"textDocument/inlayHint", "inlayHint/resolve", "workspace/inlayHint/refresh"}, prefixes: []string{"InlayHint", "ClientInlayHint"}},
	{file: "moniker.go", methods: []string{"textDocument/moniker"}, prefixes: []string{"Moniker", "UniquenessLevel"}},
	{file: "completion.go", methods: []string{"textDocument/completion", "completionItem/resolve"}, prefixes: []string{"Completion", "SelectedCompletionInfo", "ClientCompletion", "ServerCompletionItemOptions", "EditRangeWithInsertReplace", "InsertTextFormat", "InsertTextMode"}},
	{
		file:    "diagnostic.go",
		methods: []string{"textDocument/diagnostic", "workspace/diagnostic", "workspace/diagnostic/refresh"},
		prefixes: []string{
			"DocumentDiagnostic", "WorkspaceDiagnostic", "DiagnosticOptions", "DiagnosticRegistrationOptions",
			"DiagnosticServerCancellationData", "DiagnosticWorkspaceClientCapabilities", "DiagnosticClientCapabilities",
			"FullDocumentDiagnosticReport", "UnchangedDocumentDiagnosticReport", "RelatedFullDocumentDiagnosticReport",
			"RelatedUnchangedDocumentDiagnosticReport", "ClientDiagnostics", "DiagnosticsCapabilities",
			"PreviousResultId", "WorkspaceFullDocumentDiagnosticReport", "WorkspaceUnchangedDocumentDiagnosticReport",
		},
	},
	{file: "publish_diagnostics.go", methods: []string{"textDocument/publishDiagnostics"}, prefixes: []string{"PublishDiagnostics"}},
	{file: "signature_help.go", methods: []string{"textDocument/signatureHelp"}, prefixes: []string{"SignatureHelp", "SignatureInformation", "ParameterInformation", "ClientSignature"}},
	{file: "code_action.go", methods: []string{"textDocument/codeAction", "codeAction/resolve"}, prefixes: []string{"CodeAction", "ClientCodeAction"}},
	{file: "document_color.go", methods: []string{"textDocument/documentColor"}, prefixes: []string{"DocumentColor", "ColorInformation", "Color"}},
	{file: "color_presentation.go", methods: []string{"textDocument/colorPresentation"}, prefixes: []string{"ColorPresentation"}},
	{file: "formatting.go", methods: []string{"textDocument/formatting"}, prefixes: []string{"DocumentFormatting", "FormattingOptions"}},
	{file: "range_formatting.go", methods: []string{"textDocument/rangeFormatting", "textDocument/rangesFormatting"}, prefixes: []string{"DocumentRangeFormatting", "DocumentRangesFormatting"}},
	{file: "on_type_formatting.go", methods: []string{"textDocument/onTypeFormatting"}, prefixes: []string{"DocumentOnTypeFormatting"}},
	{file: "rename.go", methods: []string{"textDocument/rename", "textDocument/prepareRename"}, prefixes: []string{"Rename", "PrepareRename", "PrepareSupportDefaultBehavior"}},
	{file: "linked_editing_range.go", methods: []string{"textDocument/linkedEditingRange"}, prefixes: []string{"LinkedEditing"}},
	{file: "inline_completion.go", methods: []string{"textDocument/inlineCompletion"}, prefixes: []string{"InlineCompletion", "SelectedCompletionInfo"}},

	// ---- Workspace Features ----
	{file: "workspace_features.go", methods: []string{"workspace/symbol", "workspaceSymbol/resolve"}, prefixes: []string{"WorkspaceSymbol", "ClientSymbol"}},
	{file: "workspace_features.go", methods: []string{"workspace/configuration"}, prefixes: []string{"Configuration"}},
	{file: "workspace_features.go", methods: []string{"workspace/didChangeConfiguration"}, prefixes: []string{"DidChangeConfiguration"}},
	{file: "workspace_features.go", methods: []string{"workspace/workspaceFolders"}, prefixes: []string{"WorkspaceFolder"}},
	{file: "workspace_features.go", methods: []string{"workspace/didChangeWorkspaceFolders"}, prefixes: []string{"DidChangeWorkspaceFolders"}},
	{
		file:     "workspace_features.go",
		methods:  []string{"workspace/willCreateFiles", "workspace/didCreateFiles", "workspace/willRenameFiles", "workspace/didRenameFiles", "workspace/willDeleteFiles", "workspace/didDeleteFiles"},
		prefixes: []string{"FileOperation", "CreateFilesParams", "RenameFilesParams", "DeleteFilesParams", "FileCreate", "FileRename", "FileDelete"},
	},
	{file: "workspace_features.go", methods: []string{"workspace/didChangeWatchedFiles"}, prefixes: []string{"DidChangeWatchedFiles", "FileSystemWatcher", "FileEvent", "RelativePattern"}},
	{file: "workspace_features.go", methods: []string{"workspace/executeCommand"}, prefixes: []string{"ExecuteCommand"}},
	{file: "workspace_features.go", methods: []string{"workspace/applyEdit"}, prefixes: []string{"ApplyWorkspaceEdit", "WorkspaceEditMetadata"}},
	{file: "workspace_features.go", methods: []string{"workspace/textDocumentContent", "workspace/textDocumentContent/refresh"}, prefixes: []string{"TextDocumentContent"}},

	// ---- Window Features ----
	{file: "window_features.go", methods: []string{"window/showMessage"}, prefixes: []string{"ShowMessage", "MessageType"}},
	{file: "window_features.go", methods: []string{"window/showMessageRequest"}, prefixes: []string{"ShowMessageRequest", "MessageActionItem", "ClientShowMessageActionItem"}},
	{file: "window_features.go", methods: []string{"window/logMessage"}, prefixes: []string{"LogMessage"}},
	{file: "window_features.go", methods: []string{"window/showDocument"}, prefixes: []string{"ShowDocument"}},
	{file: "window_features.go", methods: []string{"window/workDoneProgress/create", "window/workDoneProgress/cancel"}, prefixes: []string{"WorkDoneProgressCreate", "WorkDoneProgressCancel"}},
	{file: "window_features.go", methods: []string{"telemetry/event"}, prefixes: nil},
	{file: "window_features.go", methods: []string{"$/progress"}, prefixes: []string{"ProgressParams"}},
	{file: "window_features.go", methods: []string{"$/cancelRequest"}, prefixes: nil},
}

// methodRank maps a method to its 1-based feature rank (0 = uncategorized).
var methodRank = func() map[string]int {
	m := make(map[string]int)
	for i, f := range specFeatures {
		for _, meth := range f.methods {
			m[meth] = i + 1
		}
	}
	return m
}()

// featureRank returns the 1-based spec-feature rank for a type name via the
// longest matching feature prefix, or 0 when the type is foundational/shared
// (the Basic JSON Structures bucket, emitted first).
func featureRank(name string) int {
	best, bestLen := 0, -1
	for i, f := range specFeatures {
		for _, p := range f.prefixes {
			if len(p) > bestLen && strings.HasPrefix(name, p) {
				best, bestLen = i+1, len(p)
			}
		}
	}
	return best
}

// fileForType returns the generated source file a declaration is written to,
// derived from its feature (rank 0 = the Basic JSON Structures file).
func fileForType(name string) string {
	r := featureRank(name)
	if r == 0 {
		return basicStructuresFile
	}
	return specFeatures[r-1].file
}

// orderedTypeFiles lists every generated type file in spec-document order
// (Basic Structures first), deduplicated, for deterministic emission.
func orderedTypeFiles() []string {
	out := []string{basicStructuresFile}
	seen := map[string]bool{basicStructuresFile: true}
	for _, f := range specFeatures {
		if !seen[f.file] {
			seen[f.file] = true
			out = append(out, f.file)
		}
	}
	return out
}
