// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"testing"

	"github.com/go-json-experiment/json"
)

func TestAppendEncodeBenchmarkRootsMatchCanonicalJSON(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		corpus string
		newDst func() any
	}{
		"success: initialize request": {
			corpus: "initialize_request",
			newDst: func() any { return new(InitializeParams) },
		},
		"success: initialize result": {
			corpus: "initialize_result",
			newDst: func() any { return new(InitializeResult) },
		},
		"success: completion list": {
			corpus: "completion_result",
			newDst: func() any { return new(CompletionResult) },
		},
		"success: completion array": {
			corpus: "completion_result_array",
			newDst: func() any { return new(CompletionResult) },
		},
		"success: did change": {
			corpus: "didchange",
			newDst: func() any { return new(DidChangeTextDocumentParams) },
		},
		"success: semantic tokens": {
			corpus: "semantic_tokens",
			newDst: func() any { return new(SemanticTokens) },
		},
		"success: publish diagnostics": {
			corpus: "publish_diagnostics",
			newDst: func() any { return new(PublishDiagnosticsParams) },
		},
		"success: workspace symbol": {
			corpus: "workspace_symbol_result",
			newDst: func() any { return new(WorkspaceSymbolResult) },
		},
		"success: workspace symbol information": {
			corpus: "workspace_symbol_result_info",
			newDst: func() any { return new(WorkspaceSymbolResult) },
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			data := benchCorpus(t, tt.corpus)
			dst := tt.newDst()
			if err := Unmarshal(data, dst); err != nil {
				t.Fatalf("Unmarshal(%s): %v", tt.corpus, err)
			}
			got, err := Marshal(dst)
			if err != nil {
				t.Fatalf("Marshal(%T): %v", dst, err)
			}
			want, err := json.Marshal(dst, wireOptions)
			if err != nil {
				t.Fatalf("canonical json.Marshal(%T): %v", dst, err)
			}
			if !bytes.Equal(got, want) {
				t.Errorf("append encode mismatch for %s\ngot:  %s\nwant: %s", tt.corpus, got, want)
			}
		})
	}
}

func TestAppendEncodeRejectsInvalidRawData(t *testing.T) {
	t.Parallel()

	invalid := LSPAny(`{"broken":[1,]}`)
	tests := map[string]any{
		"completion item": &CompletionItem{
			Label: "x",
			Data:  invalid,
		},
		"completion result": func() *CompletionResult {
			var result CompletionResult = CompletionItemSlice{{Label: "x", Data: invalid}}
			return &result
		}(),
		"workspace symbol result": func() *WorkspaceSymbolResult {
			var result WorkspaceSymbolResult = WorkspaceSymbolSlice{{
				BaseSymbolInformation: BaseSymbolInformation{Name: "s", Kind: SymbolKindClass},
				Location:              &LocationUriOnly{URI: DocumentURI("file:///x.go")},
				Data:                  invalid,
			}}
			return &result
		}(),
		"publish diagnostics params": &PublishDiagnosticsParams{
			URI: DocumentURI("file:///x.go"),
			Diagnostics: []Diagnostic{{
				Range: Range{
					Start: Position{Line: 1, Character: 2},
					End:   Position{Line: 3, Character: 4},
				},
				Message: String("broken"),
				Data:    invalid,
			}},
		},
	}

	for name, value := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if out, err := Marshal(value); err == nil {
				t.Fatalf("Marshal(%T) succeeded with invalid JSON: %s", value, out)
			}
		})
	}
}
