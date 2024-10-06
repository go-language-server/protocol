// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package generator

import (
	"fmt"

	"github.com/gobuffalo/flect"

	"go.lsp.dev/protocol/tools/protocol-gen/protocol"
)

var enumerationNames = map[string]string{
	"SemanticTokenTypes":            "language",
	"SemanticTokenModifiers":        "language",
	"DocumentDiagnosticReportKind":  "language",
	"ErrorCodes":                    "base",
	"LSPErrorCodes":                 "base",
	"FoldingRangeKind":              "language",
	"SymbolKind":                    "language",
	"SymbolTag":                     "language",
	"UniquenessLevel":               "language",
	"MonikerKind":                   "language",
	"InlayHintKind":                 "language",
	"MessageType":                   "window",
	"TextDocumentSyncKind":          "lifecycle",
	"TextDocumentSaveReason":        "document",
	"CompletionItemKind":            "language",
	"CompletionItemTag":             "language",
	"InsertTextFormat":              "language",
	"InsertTextMode":                "language",
	"DocumentHighlightKind":         "language",
	"CodeActionKind":                "language",
	"TraceValue":                    "basic",
	"MarkupKind":                    "basic",
	"LanguageKind":                  "basic",
	"InlineCompletionTriggerKind":   "language",
	"PositionEncodingKind":          "basic",
	"FileChangeType":                "workspace",
	"WatchKind":                     "workspace",
	"DiagnosticSeverity":            "basic",
	"DiagnosticTag":                 "basic",
	"CompletionTriggerKind":         "language",
	"SignatureHelpTriggerKind":      "language",
	"CodeActionTriggerKind":         "language",
	"FileOperationPatternKind":      "workspace",
	"NotebookCellKind":              "document",
	"ResourceOperationKind":         "basic",
	"FailureHandlingKind":           "basic",
	"PrepareSupportDefaultBehavior": "language",
	"TokenFormat":                   "language",

	"CodeActionTag": "language",
	"ApplyKind":     "language",
}

// Enumerations generates Enumerations Go type from the metaModel schema definition.
func (gen *Generator) Enumerations(enumerations []*protocol.Enumeration) error {
	for _, enum := range enumerations {
		enumName := flect.Pascalize(enum.Name)
		filename, ok := enumerationNames[enumName]
		if !ok {
			panic(fmt.Sprintf("not found %s enumerations file", enumName))
		}

		// Init filename printers
		g := NewPrinter(filename)
		gen.enumerations = append(gen.enumerations, g)

		// write Documentation
		if enum.Documentation != "" {
			g.PP(`// `, enumName, normalizeDocumentation(enum.Documentation))
		}
		if enum.Since != "" {
			if enum.Documentation != "" {
				g.PP(`//`)
			}
			g.P(`// @since `, enum.Since)
			if enum.Proposed {
				g.P(` proposed`)
			}
			g.P("\n")
		}

		g.P(`type `, enumName)
		switch enum.Type.Name {
		case protocol.EnumerationNameInteger:
			g.P(` int32`)
		case protocol.EnumerationNameString:
			g.P(` string`)
		case protocol.EnumerationNameUinteger:
			g.P(` uint32`)
		default:
			panic(fmt.Sprintf("enum: %#v\n", enum))
		}
		g.PP("\n")

		g.PP(`const (`)
		for i, val := range enum.Values {
			// write Documentation
			if val.Documentation != "" {
				g.PP(`	// `, flect.Pascalize(val.Name), enumName, normalizeDocumentation(val.Documentation))
			}
			if val.Since != "" {
				if val.Documentation != "" {
					g.PP(`	//`)
				}
				g.P(`	// @since `, val.Since)
				if val.Proposed {
					g.P(` proposed`)
				}
				g.P("\n")
			}

			g.P(`	`, flect.Pascalize(val.Name), enumName)
			g.P(` `, enumName, ` = `, val.Value)
			g.P("\n")

			// Add newline per fields
			if i < len(enum.Values)-1 {
				g.P("\n")
			}
		}
		g.PP(`)`, "\n")
	}

	return nil
}
