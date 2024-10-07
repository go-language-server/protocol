// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package generator

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/flect"

	"go.lsp.dev/protocol/tools/protocol-gen/protocol"
)

var structureNames = map[string]string{
	"TextDocumentIdentifier":                      "basic",
	"Position":                                    "basic",
	"TextDocumentPositionParams":                  "basic",
	"ImplementationParams":                        "language",
	"Range":                                       "basic",
	"Location":                                    "basic",
	"TextDocumentRegistrationOptions":             "lifecycle",
	"ImplementationOptions":                       "language",
	"ImplementationRegistrationOptions":           "language",
	"TypeDefinitionParams":                        "language",
	"TypeDefinitionOptions":                       "language",
	"TypeDefinitionRegistrationOptions":           "language",
	"WorkspaceFolder":                             "workspace",
	"WorkspaceFoldersChangeEvent":                 "workspace",
	"DidChangeWorkspaceFoldersParams":             "workspace",
	"ConfigurationItem":                           "workspace",
	"ConfigurationParams":                         "workspace",
	"DocumentColorParams":                         "language",
	"Color":                                       "language",
	"ColorInformation":                            "language",
	"DocumentColorOptions":                        "language",
	"DocumentColorRegistrationOptions":            "language",
	"ColorPresentationParams":                     "language",
	"TextEdit":                                    "basic",
	"ColorPresentation":                           "language",
	"WorkDoneProgressOptions":                     "basic",
	"FoldingRangeParams":                          "language",
	"FoldingRange":                                "language",
	"FoldingRangeOptions":                         "language",
	"FoldingRangeRegistrationOptions":             "language",
	"DeclarationParams":                           "language",
	"DeclarationOptions":                          "language",
	"DeclarationRegistrationOptions":              "language",
	"SelectionRangeParams":                        "language",
	"SelectionRange":                              "language",
	"SelectionRangeOptions":                       "language",
	"SelectionRangeRegistrationOptions":           "language",
	"WorkDoneProgressCreateParams":                "window",
	"WorkDoneProgressCancelParams":                "window",
	"CallHierarchyPrepareParams":                  "language",
	"CallHierarchyItem":                           "language",
	"CallHierarchyOptions":                        "language",
	"CallHierarchyRegistrationOptions":            "language",
	"CallHierarchyIncomingCallsParams":            "language",
	"CallHierarchyIncomingCall":                   "language",
	"CallHierarchyOutgoingCallsParams":            "language",
	"CallHierarchyOutgoingCall":                   "language",
	"SemanticTokensParams":                        "language",
	"SemanticTokens":                              "language",
	"SemanticTokensPartialResult":                 "language",
	"SemanticTokensLegend":                        "language",
	"SemanticTokensFullDelta":                     "language",
	"SemanticTokensOptions":                       "language",
	"SemanticTokensRegistrationOptions":           "language",
	"SemanticTokensDeltaParams":                   "language",
	"SemanticTokensEdit":                          "language",
	"SemanticTokensDelta":                         "language",
	"SemanticTokensDeltaPartialResult":            "language",
	"SemanticTokensRangeParams":                   "language",
	"ShowDocumentParams":                          "window",
	"ShowDocumentResult":                          "window",
	"LinkedEditingRangeParams":                    "language",
	"LinkedEditingRanges":                         "language",
	"LinkedEditingRangeOptions":                   "language",
	"LinkedEditingRangeRegistrationOptions":       "language",
	"FileCreate":                                  "workspace",
	"CreateFilesParams":                           "workspace",
	"ResourceOperation":                           "workspace",
	"DeleteFileOptions":                           "workspace",
	"DeleteFile":                                  "workspace",
	"RenameFileOptions":                           "workspace",
	"RenameFile":                                  "workspace",
	"CreateFileOptions":                           "workspace",
	"CreateFile":                                  "workspace",
	"OptionalVersionedTextDocumentIdentifier":     "basic",
	"AnnotatedTextEdit":                           "basic",
	"TextDocumentEdit":                            "basic",
	"ChangeAnnotation":                            "basic",
	"WorkspaceEdit":                               "basic",
	"FileOperationPatternOptions":                 "workspace",
	"FileOperationPattern":                        "workspace",
	"FileOperationFilter":                         "workspace",
	"FileOperationRegistrationOptions":            "workspace",
	"FileRename":                                  "workspace",
	"RenameFilesParams":                           "workspace",
	"FileDelete":                                  "workspace",
	"DeleteFilesParams":                           "workspace",
	"MonikerParams":                               "language",
	"Moniker":                                     "language",
	"MonikerOptions":                              "language",
	"MonikerRegistrationOptions":                  "language",
	"TypeHierarchyPrepareParams":                  "language",
	"TypeHierarchyItem":                           "language",
	"TypeHierarchyOptions":                        "language",
	"TypeHierarchyRegistrationOptions":            "language",
	"TypeHierarchySupertypesParams":               "language",
	"TypeHierarchySubtypesParams":                 "language",
	"InlineValueContext":                          "language",
	"InlineValueParams":                           "language",
	"InlineValueOptions":                          "language",
	"InlineValueRegistrationOptions":              "language",
	"InlayHintParams":                             "language",
	"MarkupContent":                               "basic",
	"Command":                                     "basic",
	"InlayHintLabelPart":                          "language",
	"InlayHint":                                   "language",
	"InlayHintOptions":                            "language",
	"InlayHintRegistrationOptions":                "language",
	"DocumentDiagnosticParams":                    "language",
	"UnchangedDocumentDiagnosticReport":           "language",
	"CodeDescription":                             "basic",
	"DiagnosticRelatedInformation":                "basic",
	"Diagnostic":                                  "basic",
	"FullDocumentDiagnosticReport":                "language",
	"DocumentDiagnosticReportPartialResult":       "language",
	"DiagnosticServerCancellationData":            "language",
	"DiagnosticOptions":                           "language",
	"DiagnosticRegistrationOptions":               "language",
	"PreviousResultID":                            "language",
	"WorkspaceDiagnosticParams":                   "language",
	"WorkspaceDiagnosticReport":                   "language",
	"WorkspaceDiagnosticReportPartialResult":      "language",
	"ExecutionSummary":                            "document",
	"NotebookCell":                                "document",
	"NotebookDocument":                            "document",
	"TextDocumentItem":                            "basic",
	"DidOpenNotebookDocumentParams":               "document",
	"NotebookCellLanguage":                        "document",
	"NotebookDocumentFilterWithCells":             "document",
	"NotebookDocumentFilterWithNotebook":          "document",
	"NotebookDocumentSyncOptions":                 "document",
	"NotebookDocumentSyncRegistrationOptions":     "document",
	"VersionedNotebookDocumentIdentifier":         "document",
	"NotebookCellArrayChange":                     "document",
	"NotebookDocumentCellChangeStructure":         "document",
	"VersionedTextDocumentIdentifier":             "basic",
	"NotebookDocumentCellContentChanges":          "document",
	"NotebookDocumentCellChanges":                 "document",
	"NotebookDocumentChangeEvent":                 "document",
	"DidChangeNotebookDocumentParams":             "document",
	"NotebookDocumentIdentifier":                  "document",
	"DidSaveNotebookDocumentParams":               "document",
	"DidCloseNotebookDocumentParams":              "document",
	"SelectedCompletionInfo":                      "language",
	"InlineCompletionContext":                     "language",
	"InlineCompletionParams":                      "language",
	"StringValue":                                 "basic",
	"InlineCompletionItem":                        "language",
	"InlineCompletionList":                        "language",
	"InlineCompletionOptions":                     "language",
	"InlineCompletionRegistrationOptions":         "language",
	"Registration":                                "lifecycle",
	"RegistrationParams":                          "lifecycle",
	"Unregistration":                              "lifecycle",
	"UnregistrationParams":                        "lifecycle",
	"ClientInfo":                                  "lifecycle",
	"ChangeAnnotationsSupportOptions":             "basic",
	"WorkspaceEditClientCapabilities":             "basic",
	"DidChangeConfigurationClientCapabilities":    "workspace",
	"DidChangeWatchedFilesClientCapabilities":     "workspace",
	"ClientSymbolKindOptions":                     "workspace",
	"ClientSymbolTagOptions":                      "workspace",
	"ClientSymbolResolveOptions":                  "workspace",
	"WorkspaceSymbolClientCapabilities":           "workspace",
	"ExecuteCommandClientCapabilities":            "workspace",
	"SemanticTokensWorkspaceClientCapabilities":   "lifecycle",
	"CodeLensWorkspaceClientCapabilities":         "lifecycle",
	"FileOperationClientCapabilities":             "lifecycle",
	"InlineValueWorkspaceClientCapabilities":      "lifecycle",
	"InlayHintWorkspaceClientCapabilities":        "lifecycle",
	"DiagnosticWorkspaceClientCapabilities":       "lifecycle",
	"FoldingRangeWorkspaceClientCapabilities":     "lifecycle",
	"WorkspaceClientCapabilities":                 "lifecycle",
	"TextDocumentSyncClientCapabilities":          "lifecycle",
	"CompletionItemTagOptions":                    "lifecycle",
	"ClientCompletionItemResolveOptions":          "lifecycle",
	"ClientCompletionItemInsertTextModeOptions":   "lifecycle",
	"ClientCompletionItemOptions":                 "lifecycle",
	"ClientCompletionItemOptionsKind":             "lifecycle",
	"CompletionListCapabilities":                  "lifecycle",
	"CompletionClientCapabilities":                "lifecycle",
	"HoverClientCapabilities":                     "lifecycle",
	"ClientSignatureParameterInformationOptions":  "lifecycle",
	"ClientSignatureInformationOptions":           "lifecycle",
	"SignatureHelpClientCapabilities":             "lifecycle",
	"DeclarationClientCapabilities":               "lifecycle",
	"DefinitionClientCapabilities":                "lifecycle",
	"TypeDefinitionClientCapabilities":            "lifecycle",
	"ImplementationClientCapabilities":            "lifecycle",
	"ReferenceClientCapabilities":                 "lifecycle",
	"DocumentHighlightClientCapabilities":         "lifecycle",
	"DocumentSymbolClientCapabilities":            "lifecycle",
	"ClientCodeActionKindOptions":                 "lifecycle",
	"ClientCodeActionLiteralOptions":              "lifecycle",
	"ClientCodeActionResolveOptions":              "lifecycle",
	"CodeActionClientCapabilities":                "lifecycle",
	"CodeLensClientCapabilities":                  "lifecycle",
	"DocumentLinkClientCapabilities":              "lifecycle",
	"DocumentColorClientCapabilities":             "lifecycle",
	"DocumentFormattingClientCapabilities":        "lifecycle",
	"DocumentRangeFormattingClientCapabilities":   "lifecycle",
	"DocumentOnTypeFormattingClientCapabilities":  "lifecycle",
	"RenameClientCapabilities":                    "lifecycle",
	"ClientFoldingRangeKindOptions":               "lifecycle",
	"ClientFoldingRangeOptions":                   "lifecycle",
	"FoldingRangeClientCapabilities":              "lifecycle",
	"SelectionRangeClientCapabilities":            "lifecycle",
	"ClientDiagnosticsTagOptions":                 "lifecycle",
	"PublishDiagnosticsClientCapabilities":        "lifecycle",
	"CallHierarchyClientCapabilities":             "lifecycle",
	"ClientSemanticTokensRequestFullDelta":        "lifecycle",
	"ClientSemanticTokensRequestOptions":          "lifecycle",
	"SemanticTokensClientCapabilities":            "lifecycle",
	"LinkedEditingRangeClientCapabilities":        "lifecycle",
	"MonikerClientCapabilities":                   "lifecycle",
	"TypeHierarchyClientCapabilities":             "lifecycle",
	"InlineValueClientCapabilities":               "lifecycle",
	"ClientInlayHintResolveOptions":               "lifecycle",
	"InlayHintClientCapabilities":                 "lifecycle",
	"DiagnosticClientCapabilities":                "lifecycle",
	"InlineCompletionClientCapabilities":          "lifecycle",
	"TextDocumentClientCapabilities":              "lifecycle",
	"NotebookDocumentSyncClientCapabilities":      "lifecycle",
	"NotebookDocumentClientCapabilities":          "lifecycle",
	"ClientShowMessageActionItemOptions":          "lifecycle",
	"ShowMessageRequestClientCapabilities":        "lifecycle",
	"ShowDocumentClientCapabilities":              "lifecycle",
	"WindowClientCapabilities":                    "lifecycle",
	"StaleRequestSupportOptions":                  "lifecycle",
	"RegularExpressionsClientCapabilities":        "lifecycle",
	"MarkdownClientCapabilities":                  "lifecycle",
	"GeneralClientCapabilities":                   "lifecycle",
	"ClientCapabilities":                          "lifecycle",
	"InitializeParamsBase":                        "lifecycle",
	"WorkspaceFoldersInitializeParams":            "lifecycle",
	"InitializeParams":                            "lifecycle",
	"SaveOptions":                                 "document",
	"TextDocumentSyncOptions":                     "document",
	"ServerCompletionItemOptions":                 "language",
	"CompletionOptions":                           "language",
	"HoverOptions":                                "language",
	"SignatureHelpOptions":                        "language",
	"DefinitionOptions":                           "language",
	"ReferenceOptions":                            "language",
	"DocumentHighlightOptions":                    "language",
	"DocumentSymbolOptions":                       "language",
	"CodeActionKindDocumentation":                 "language",
	"CodeActionOptions":                           "language",
	"CodeLensOptions":                             "language",
	"DocumentLinkOptions":                         "language",
	"WorkspaceSymbolOptions":                      "language",
	"DocumentFormattingOptions":                   "language",
	"DocumentRangeFormattingOptions":              "language",
	"DocumentOnTypeFormattingOptions":             "language",
	"RenameOptions":                               "language",
	"ExecuteCommandOptions":                       "language",
	"WorkspaceFoldersServerCapabilities":          "workspace",
	"FileOperationOptions":                        "lifecycle",
	"WorkspaceOptions":                            "lifecycle",
	"ServerCapabilities":                          "lifecycle",
	"ServerInfo":                                  "lifecycle",
	"InitializeResult":                            "lifecycle",
	"InitializeError":                             "lifecycle",
	"InitializedParams":                           "lifecycle",
	"DidChangeConfigurationParams":                "workspace",
	"DidChangeConfigurationRegistrationOptions":   "workspace",
	"ShowMessageParams":                           "window",
	"MessageActionItem":                           "window",
	"ShowMessageRequestParams":                    "window",
	"LogMessageParams":                            "window",
	"DidOpenTextDocumentParams":                   "document",
	"DidChangeTextDocumentParams":                 "document",
	"TextDocumentChangeRegistrationOptions":       "document",
	"DidCloseTextDocumentParams":                  "document",
	"DidSaveTextDocumentParams":                   "document",
	"TextDocumentSaveRegistrationOptions":         "document",
	"WillSaveTextDocumentParams":                  "document",
	"FileEvent":                                   "workspace",
	"DidChangeWatchedFilesParams":                 "workspace",
	"FileSystemWatcher":                           "workspace",
	"DidChangeWatchedFilesRegistrationOptions":    "workspace",
	"PublishDiagnosticsParams":                    "language",
	"CompletionContext":                           "language",
	"CompletionParams":                            "language",
	"CompletionItemLabelDetails":                  "language",
	"InsertReplaceEdit":                           "language",
	"CompletionItem":                              "language",
	"EditRangeWithInsertReplace":                  "language",
	"CompletionItemDefaults":                      "language",
	"CompletionList":                              "language",
	"CompletionRegistrationOptions":               "language",
	"HoverParams":                                 "language",
	"Hover":                                       "language",
	"HoverRegistrationOptions":                    "language",
	"ParameterInformation":                        "language",
	"SignatureInformation":                        "language",
	"SignatureHelp":                               "language",
	"SignatureHelpContext":                        "language",
	"SignatureHelpParams":                         "language",
	"SignatureHelpRegistrationOptions":            "language",
	"DefinitionParams":                            "language",
	"DefinitionRegistrationOptions":               "language",
	"ReferenceContext":                            "language",
	"ReferenceParams":                             "language",
	"ReferenceRegistrationOptions":                "language",
	"DocumentHighlightParams":                     "language",
	"DocumentHighlight":                           "language",
	"DocumentHighlightRegistrationOptions":        "language",
	"DocumentSymbolParams":                        "language",
	"BaseSymbolInformation":                       "language",
	"SymbolInformation":                           "language",
	"DocumentSymbol":                              "language",
	"DocumentSymbolRegistrationOptions":           "language",
	"CodeActionContext":                           "language",
	"CodeActionParams":                            "language",
	"CodeActionDisabled":                          "language",
	"CodeAction":                                  "language",
	"CodeActionRegistrationOptions":               "language",
	"WorkspaceSymbolParams":                       "workspace",
	"LocationURIOnly":                             "language",
	"WorkspaceSymbol":                             "workspace",
	"WorkspaceSymbolRegistrationOptions":          "workspace",
	"CodeLensParams":                              "language",
	"CodeLens":                                    "language",
	"CodeLensRegistrationOptions":                 "language",
	"DocumentLinkParams":                          "language",
	"DocumentLink":                                "language",
	"DocumentLinkRegistrationOptions":             "language",
	"FormattingOptions":                           "language",
	"DocumentFormattingParams":                    "language",
	"DocumentFormattingRegistrationOptions":       "language",
	"DocumentRangeFormattingParams":               "language",
	"DocumentRangeFormattingRegistrationOptions":  "language",
	"DocumentRangesFormattingParams":              "language",
	"DocumentOnTypeFormattingParams":              "language",
	"DocumentOnTypeFormattingRegistrationOptions": "language",
	"RenameParams":                                "language",
	"RenameRegistrationOptions":                   "language",
	"PrepareRenameParams":                         "language",
	"ExecuteCommandParams":                        "language",
	"ExecuteCommandRegistrationOptions":           "language",
	"ApplyWorkspaceEditParams":                    "workspace",
	"ApplyWorkspaceEditResult":                    "workspace",
	"WorkDoneProgressBegin":                       "basic",
	"WorkDoneProgressReport":                      "basic",
	"WorkDoneProgressEnd":                         "basic",
	"SetTraceParams":                              "lifecycle",
	"LogTraceParams":                              "lifecycle",
	"CancelParams":                                "base",
	"ProgressParams":                              "base",
	"WorkDoneProgressParams":                      "basic",
	"PartialResultParams":                         "basic",
	"LocationLink":                                "basic",
	"StaticRegistrationOptions":                   "lifecycle",
	"InlineValueText":                             "language",
	"InlineValueVariableLookup":                   "language",
	"InlineValueEvaluatableExpression":            "language",
	"RelatedFullDocumentDiagnosticReport":         "language",
	"RelatedUnchangedDocumentDiagnosticReport":    "language",
	"PrepareRenamePlaceholder":                    "language",
	"PrepareRenameDefaultBehavior":                "language",
	"WorkspaceFullDocumentDiagnosticReport":       "language",
	"WorkspaceUnchangedDocumentDiagnosticReport":  "language",
	"TextDocumentContentChangePartial":            "language",
	"TextDocumentContentChangeWholeDocument":      "language",
	"MarkedStringWithLanguage":                    "language",
	"NotebookCellTextDocumentFilter":              "language",
	"RelativePattern":                             "workspace",
	"TextDocumentFilterLanguage":                  "language",
	"TextDocumentFilterScheme":                    "language",
	"TextDocumentFilterPattern":                   "language",
	"NotebookDocumentFilterNotebookType":          "document",
	"NotebookDocumentFilterScheme":                "document",
	"NotebookDocumentFilterPattern":               "document",

	"SnippetTextEdit":                        "basic",
	"TextDocumentContentParams":              "basic",
	"TextDocumentContentResult":              "basic",
	"TextDocumentContentOptions":             "basic",
	"TextDocumentContentRegistrationOptions": "basic",
	"TextDocumentContentRefreshParams":       "basic",
	"TextDocumentContentClientCapabilities":  "lifecycle",
	"ClientCodeLensResolveOptions":           "lifecycle",
	"DiagnosticsCapabilities":                "lifecycle",
	"WorkspaceEditMetadata":                  "workspace",
	"CodeActionTagOptions":                   "language",
	"CompletionItemApplyKinds":               "language",
}

// Structures generates Structure Go type from the metaModel schema definition.
func (gen *Generator) Structures(structures []*protocol.Structure) error {
	for _, structure := range structures {
		structuresName := flect.Pascalize(structure.Name)

		filename, ok := structureNames[structuresName]
		if !ok {
			panic(fmt.Sprintf("not found %s structures file", structuresName))
		}

		// Init structures printers
		g := NewPrinter(filename)
		gen.structures = append(gen.structures, g)

		if structure.Documentation != "" {
			g.PP(`// `, structuresName, normalizeDocumentation(structure.Documentation))
		}
		if structure.Since != "" {
			if structure.Documentation != "" {
				g.PP(`//`)
			}
			g.P(`// @since `, structure.Since)
			if structure.Proposed {
				g.P(` proposed`)
			}
			g.P("\n")
		}

		g.PP(`type `, structuresName, ` struct {`)

		var needNewline bool
		if len(structure.Extends) > 0 {
			g.PP(`	// extends`)
			for i, extend := range structure.Extends {
				switch extend := extend.(type) {
				case *protocol.ReferenceType:
					g.PP(`	`, normalizeLSPTypes(extend.Name))
				default:
					fmt.Printf("mixin[%d]: %#[2]v %[2]T\n", i, extend)
				}
				if ns := extend.SubTypes(); ns != nil {
					for i, n := range ns {
						switch n := n.(type) {
						default:
							fmt.Printf("extend[%d]: %#[2]v %[2]T\n", i, n)
						}
						g.PP(`	`, n)
					}
				}
			}
			needNewline = true
		}

		if len(structure.Mixins) > 0 {
			g.PP(`	// mixins`)
			for i, mixin := range structure.Mixins {
				switch mixin := mixin.(type) {
				case *protocol.ReferenceType:
					g.PP(`	`, normalizeLSPTypes(mixin.Name))
				default:
					fmt.Printf("mixin[%d]: %#[2]v %[2]T\n", i, mixin)
				}
				if ns := mixin.SubTypes(); ns != nil {
					for i, n := range ns {
						switch n := n.(type) {
						default:
							fmt.Printf("mixin.SubTypes[%d]: %#[2]v %[2]T\n", i, n)
						}
						g.PP(`	`, n)
					}
				}
			}
			needNewline = true
		}

		if len(structure.Properties) > 0 {
			if needNewline {
				g.P("\n")
			}

			for i, prop := range structure.Properties {
				propName := flect.Pascalize(prop.Name)

				// write Documentation
				if prop.Documentation != "" {
					g.PP(`	// `, propName, normalizeDocumentation(prop.Documentation))
				}
				if structure.Since != "" {
					if prop.Documentation != "" && !strings.Contains(prop.Documentation, "since") {
						g.PP(`	//`)
					}
					g.P(`	// @since `, structure.Since)
					if structure.Proposed {
						g.P(` proposed`)
					}
					g.P("\n")
				}

				g.P(`	`, propName)

				propType := prop.Type
				switch node := propType.(type) {
				case protocol.BaseType:
					g.P(` `, node.String())

				case *protocol.ReferenceType:
					g.P(` `)
					gen.renderStructuresReferenceType(g, prop, node)

				case *protocol.ArrayType:
					genericsProp := &protocol.Property{
						Name:          structuresName + propName,
						Documentation: prop.Documentation,
						Since:         structure.Since,
						Proposed:      structure.Proposed,
					}
					g.P(` `)
					gen.renderStructuresArrayTypeGeneric(g, node, genericsProp)

				case *protocol.MapType:
					genericsProp := &protocol.Property{
						Name:          structuresName + propName,
						Documentation: prop.Documentation,
						Since:         structure.Since,
						Proposed:      structure.Proposed,
					}
					gen.renderStructuresMapType(g, node, genericsProp)

				case *protocol.OrType:
					switch {
					case len(node.Items) == 2 && (isNull(node.Items[0], node.Items[1])):
						prop.Optional = true
						for _, item := range node.Items {
							if !isNull(item) {
								switch item := item.(type) {
								case protocol.BaseType:
									g.P(` `, item.String())
								case *protocol.ReferenceType:
									g.P(` `)
									gen.renderStructuresReferenceType(g, prop, item)
								case *protocol.ArrayType:
									g.P(` `)
									gen.renderStructuresArrayType(g, item)
								default:
									panic(fmt.Sprintf("structures.OrType: %#v\n", item))
								}
							}
						}
					default:
						if isNull(node.Items...) {
							prop.Optional = true
						}
						genericsProp := &protocol.Property{
							Name:          structuresName + propName,
							Optional:      prop.Optional,
							Documentation: prop.Documentation,
							Proposed:      prop.Proposed,
							Since:         prop.Since,
							Deprecated:    prop.Deprecated,
						}
						gen.renderStructuresOrType(g, node, genericsProp)
					}

				default:
					// *protocol.AndType and *protocol.TupleType are nothing to do in structures
					panic(fmt.Sprintf("structures: %#v\n", node))
				}

				g.P(" `json:\"", prop.JSONName)
				if prop.Optional {
					g.P(`,omitempty`)
				}
				g.PP("\"`")

				// Add newline per fields
				if i < len(structure.Properties)-1 {
					g.P("\n")
				}
			}
		}
		g.PP(`}`)
		g.P("\n")
	}

	return nil
}

func (gen *Generator) renderStructuresReferenceType(g Printer, prop *protocol.Property, ref *protocol.ReferenceType) {
	name := ref.String()

	// don't use normalizeLSPTypes for the avoid pointer to `map[string]any` and `[]any`
	switch name {
	case `LSPAny`:
		name = `any`
	case `LSPObject`:
		name = `map[string]any`
	case `LSPArray`:
		name = `[]any`
	default:
		if _, ok := enumerationNames[name]; !ok && prop.Optional {
			g.P(`*`)
		}
	}

	g.P(name)
}

func (gen *Generator) renderStructuresArrayType(g Printer, array *protocol.ArrayType) {
	elem := array.Element
	switch elem := elem.(type) {
	case *protocol.ReferenceType:
		g.P(`[]` + normalizeLSPTypes(elem.String()))

	default:
		panic(fmt.Sprintf("structures.ArrayKind: %#v\n", elem))
	}
}

func (gen *Generator) renderStructuresArrayTypeGeneric(g Printer, array *protocol.ArrayType, genericsProp *protocol.Property) {
	elem := array.Element
	switch elem := elem.(type) {
	case protocol.BaseType:
		g.P(`[]` + elem.String())

	case *protocol.ReferenceType:
		g.P(`[]` + normalizeLSPTypes(elem.String()))

	case *protocol.OrType:
		gen.renderStructuresOrType(g, elem, genericsProp)

	default:
		panic(fmt.Sprintf("structures.ArrayKind: %#v\n", elem))
	}
}

func (gen *Generator) renderStructuresMapType(g Printer, m *protocol.MapType, genericsProp *protocol.Property) {
	g.P(` map`)

	// write map key
	switch key := m.Key.(type) {
	case *protocol.DocumentUriType:
		g.P(`[`, key.String(), `]`)

	case *protocol.ReferenceType:
		g.P(`[`, normalizeLSPTypes(key.String()), `]`)

	default:
		panic(fmt.Sprintf("structures.MapType.Key: %[1]T = %#[1]v\n", m.Key))
	}

	// write map value
	switch val := m.Value.(type) {
	case *protocol.ReferenceType:
		g.P(normalizeLSPTypes(val.String()))

	case *protocol.ArrayType:
		gen.renderStructuresArrayTypeGeneric(g, val, genericsProp)

	case *protocol.OrType:
		gen.renderStructuresOrType(g, val, genericsProp)

	default:
		panic(fmt.Sprintf("structures.MapType.Value: %[1]T = %#[1]v\n", m.Value))
	}
}

func (gen *Generator) renderStructuresOrType(g Printer, or *protocol.OrType, genericsProp *protocol.Property) {
	g.P(` `, genericsProp.Name)
	gen.genericsTypes[genericsProp] = or.Items
}
