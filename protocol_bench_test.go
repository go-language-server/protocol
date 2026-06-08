// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"os"
	"path/filepath"
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
	corpus string
	newDst func() any
}

// benchCases enumerates the real LSP payload categories exercised by the
// decode/encode benchmarks. Union-bearing result types (CompletionResult,
// WorkspaceSymbolResult) and union-bearing fields (didChange contentChanges,
// the dispatch collision shapes) drive the union-dispatch hot path that the
// optimization targets.
var benchCases = map[string]benchCase{
	"initialize_request":    {"initialize_request", func() any { return new(InitializeParams) }},
	"initialize_result":     {"initialize_result", func() any { return new(InitializeResult) }},
	"completion_list":       {"completion_result", func() any { return new(CompletionResult) }},
	"completion_array":      {"completion_result_array", func() any { return new(CompletionResult) }},
	"didchange":             {"didchange", func() any { return new(DidChangeTextDocumentParams) }},
	"semantic_tokens":       {"semantic_tokens", func() any { return new(SemanticTokens) }},
	"publish_diagnostics":   {"publish_diagnostics", func() any { return new(PublishDiagnosticsParams) }},
	"workspace_symbol":      {"workspace_symbol_result", func() any { return new(WorkspaceSymbolResult) }},
	"workspace_symbol_info": {"workspace_symbol_result_info", func() any { return new(WorkspaceSymbolResult) }},
}

// BenchmarkDecode measures Unmarshal throughput and allocations per payload
// category. The union-bearing categories dominate the dispatch cost the
// optimization addresses.
func BenchmarkDecode(b *testing.B) {
	for name, bc := range benchCases {
		data := benchCorpus(b, bc.corpus)
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(data)))
			for b.Loop() {
				dst := bc.newDst()
				if err := Unmarshal(data, dst); err != nil {
					b.Fatalf("unmarshal %s: %v", name, err)
				}
			}
		})
	}
}

// BenchmarkEncode measures Marshal throughput and allocations per payload
// category. Each payload is decoded once outside the timed loop so the loop
// measures encode only.
func BenchmarkEncode(b *testing.B) {
	for name, bc := range benchCases {
		data := benchCorpus(b, bc.corpus)
		dst := bc.newDst()
		if err := Unmarshal(data, dst); err != nil {
			b.Fatalf("setup unmarshal %s: %v", name, err)
		}
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				out, err := Marshal(dst)
				if err != nil {
					b.Fatalf("marshal %s: %v", name, err)
				}
				b.SetBytes(int64(len(out)))
			}
		})
	}
}
