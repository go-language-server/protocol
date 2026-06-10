// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// benchCorpus loads a corpus payload by base name (without extension) from
// testdata/corpus. It fails the benchmark if the file is missing so a renamed
// or absent payload is never silently benchmarked as an empty input.
func benchCorpus(tb testing.TB, name string) []byte {
	tb.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", "corpus", name+".json"))
	if err != nil {
		tb.Fatalf("load corpus %s: %v", name, err)
	}
	return b
}

// benchCase pairs a corpus payload with a factory for the concrete type it
// decodes into. The factory returns a fresh pointer each call so decode never
// reuses state across iterations.
type benchCase struct {
	name   string
	corpus string
	newDst func() any
}

// benchCases enumerates the real LSP payload categories exercised by the
// decode/encode benchmarks. Union-bearing result types (CompletionResult,
// WorkspaceSymbolResult) and union-bearing fields (didChange contentChanges,
// the dispatch collision shapes) drive the union-dispatch hot path that the
// optimization targets.
var benchCases = []benchCase{
	{"initialize_request", "initialize_request", func() any { return new(InitializeParams) }},
	{"initialize_result", "initialize_result", func() any { return new(InitializeResult) }},
	{"completion_list", "completion_result", func() any { return new(CompletionResult) }},
	{"completion_array", "completion_result_array", func() any { return new(CompletionResult) }},
	{"didchange", "didchange", func() any { return new(DidChangeTextDocumentParams) }},
	{"semantic_tokens", "semantic_tokens", func() any { return new(SemanticTokens) }},
	{"publish_diagnostics", "publish_diagnostics", func() any { return new(PublishDiagnosticsParams) }},
	{"workspace_symbol", "workspace_symbol_result", func() any { return new(WorkspaceSymbolResult) }},
	{"workspace_symbol_info", "workspace_symbol_result_info", func() any { return new(WorkspaceSymbolResult) }},
}

// dispatchBenchCases pin the corpus fixtures that exercise structurally
// ambiguous union arms. They are decode-only: most fixtures are a single union
// value rather than a complete request/response payload, so adding them to the
// encode benchmark would mix dispatch cost with tiny-fixture marshal noise.
var dispatchBenchCases = []benchCase{
	{"command_or_codeaction_command", "dispatch_command_string", func() any { return new(CommandOrCodeAction) }},
	{"command_or_codeaction_codeaction_command", "dispatch_codeaction_command_obj", func() any { return new(CommandOrCodeAction) }},
	{"command_or_codeaction_codeaction_kind", "dispatch_codeaction_kind", func() any { return new(CommandOrCodeAction) }},
	{"document_filter_notebook", "dispatch_docfilter_notebook", func() any { return new(DocumentFilter) }},
	{"document_change_create", "dispatch_documentchange_create", func() any { return new(DocumentChange) }},
	{"document_change_textedit", "dispatch_documentchange_textedit", func() any { return new(DocumentChange) }},
	{"notebook_sync_registration", "dispatch_notebook_sync_id", func() any { return new(NotebookDocumentSync) }},
	{"text_document_filter_language", "dispatch_textdocfilter_lang", func() any { return new(TextDocumentFilter) }},
	{"text_document_filter_scheme", "dispatch_textdocfilter_scheme", func() any { return new(TextDocumentFilter) }},
	{"text_document_filter_pattern", "dispatch_textdocfilter_pattern", func() any { return new(TextDocumentFilter) }},
}

func TestBenchmarkCasesHaveStableUniqueNamesAndCorpus(t *testing.T) {
	wantOrder := []string{
		"initialize_request",
		"initialize_result",
		"completion_list",
		"completion_array",
		"didchange",
		"semantic_tokens",
		"publish_diagnostics",
		"workspace_symbol",
		"workspace_symbol_info",
	}
	assertBenchCases(t, benchCases, wantOrder)
}

func TestDispatchBenchmarkCasesHaveStableUniqueNamesAndCorpus(t *testing.T) {
	wantOrder := []string{
		"command_or_codeaction_command",
		"command_or_codeaction_codeaction_command",
		"command_or_codeaction_codeaction_kind",
		"document_filter_notebook",
		"document_change_create",
		"document_change_textedit",
		"notebook_sync_registration",
		"text_document_filter_language",
		"text_document_filter_scheme",
		"text_document_filter_pattern",
	}
	assertBenchCases(t, dispatchBenchCases, wantOrder)
}

func assertBenchCases(t *testing.T, cases []benchCase, wantOrder []string) {
	t.Helper()

	gotOrder := make([]string, 0, len(cases))
	seen := make(map[string]struct{}, len(cases))
	for _, bc := range cases {
		gotOrder = append(gotOrder, bc.name)
		if bc.name == "" {
			t.Fatal("benchmark case has empty name")
		}
		if _, ok := seen[bc.name]; ok {
			t.Fatalf("duplicate benchmark case name %q", bc.name)
		}
		seen[bc.name] = struct{}{}
		if bc.corpus == "" {
			t.Fatalf("benchmark case %q has empty corpus", bc.name)
		}
		data := benchCorpus(t, bc.corpus)
		dst := bc.newDst()
		if dst == nil {
			t.Fatalf("benchmark case %q returned nil destination", bc.name)
		}
		if err := Unmarshal(data, dst); err != nil {
			t.Fatalf("benchmark case %q corpus %q does not decode: %v", bc.name, bc.corpus, err)
		}
	}
	if !slices.Equal(gotOrder, wantOrder) {
		t.Fatalf("benchmark case order = %q, want %q", gotOrder, wantOrder)
	}
}

// BenchmarkDecode measures Unmarshal throughput and allocations per payload
// category. The union-bearing categories dominate the dispatch cost the
// optimization addresses.
func BenchmarkDecode(b *testing.B) {
	for _, bc := range benchCases {
		data := benchCorpus(b, bc.corpus)
		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(data)))
			for b.Loop() {
				dst := bc.newDst()
				if err := Unmarshal(data, dst); err != nil {
					b.Fatalf("unmarshal %s: %v", bc.name, err)
				}
			}
		})
	}
}

// BenchmarkDecodeDispatchCorpus isolates the union dispatch collision corpus
// from the larger LSP payload suite. It provides a stable, low-noise benchmark
// for scanner/dispatch work without waiting for the full decode corpus.
func BenchmarkDecodeDispatchCorpus(b *testing.B) {
	for _, bc := range dispatchBenchCases {
		data := benchCorpus(b, bc.corpus)
		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(data)))
			for b.Loop() {
				dst := bc.newDst()
				if err := Unmarshal(data, dst); err != nil {
					b.Fatalf("unmarshal %s: %v", bc.name, err)
				}
			}
		})
	}
}

// BenchmarkEncode measures Marshal throughput and allocations per payload
// category. Each payload is decoded once outside the timed loop so the loop
// measures encode only.
func BenchmarkEncode(b *testing.B) {
	for _, bc := range benchCases {
		data := benchCorpus(b, bc.corpus)
		dst := bc.newDst()
		if err := Unmarshal(data, dst); err != nil {
			b.Fatalf("setup unmarshal %s: %v", bc.name, err)
		}
		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				out, err := Marshal(dst)
				if err != nil {
					b.Fatalf("marshal %s: %v", bc.name, err)
				}
				b.SetBytes(int64(len(out)))
			}
		})
	}
}
