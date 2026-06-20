// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"math"
	"testing"

	"github.com/go-json-experiment/json"
	"go.lsp.dev/uri"
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
				Location:              &LocationUriOnly{URI: uri.URI("file:///x.go")},
				Data:                  invalid,
			}}
			return &result
		}(),
		"publish diagnostics params": &PublishDiagnosticsParams{
			URI: uri.URI("file:///x.go"),
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

// appendParityTargets enumerates the decode destinations the append-encode
// parity fuzzer drives. It mirrors the bench spine plus the structurally
// ambiguous dispatch unions so every generated writer family (struct, named
// slice, boxed union, scalar arm) is exercised against the streaming oracle.
func appendParityTargets() []func() any {
	return []func() any{
		func() any { return new(InitializeParams) },
		func() any { return new(InitializeResult) },
		func() any { return new(CompletionResult) },
		func() any { return new(DidChangeTextDocumentParams) },
		func() any { return new(SemanticTokens) },
		func() any { return new(PublishDiagnosticsParams) },
		func() any { return new(WorkspaceSymbolResult) },
		func() any { return new(CommandOrCodeAction) },
		func() any { return new(DocumentFilter) },
		func() any { return new(DocumentChange) },
		func() any { return new(NotebookDocumentSync) },
		func() any { return new(TextDocumentFilter) },
	}
}

// TestBoxedUnionArmMatchesSinglePath pins the cross-path consistency the
// boxed slab deciders must preserve: a union payload decoded as a lone field
// and as a one-element slice element must select the same dynamic arm, even
// for spec-violating objects carrying both shapes' keys (the case that caught
// a tier-order divergence in review).
func TestBoxedUnionArmMatchesSinglePath(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		textEdit string
	}{
		"success: plain text edit":          {textEdit: `{"range":{"start":{"line":1,"character":2},"end":{"line":1,"character":3}},"newText":"x"}`},
		"success: insert replace edit":      {textEdit: `{"newText":"x","insert":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"replace":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}}}`},
		"success: dual-shape both arm keys": {textEdit: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"x","insert":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"replace":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}}}`},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			single := []byte(`{"label":"a","textEdit":` + tt.textEdit + `}`)
			slice := []byte(`[{"label":"a","textEdit":` + tt.textEdit + `}]`)

			var one CompletionItem
			if err := Unmarshal(single, &one); err != nil {
				t.Fatalf("single decode: %v", err)
			}
			var many CompletionItemSlice
			if err := Unmarshal(slice, &many); err != nil {
				t.Fatalf("slice decode: %v", err)
			}
			if len(many) != 1 {
				t.Fatalf("slice decode produced %d items", len(many))
			}
			gotSingle := dynamicTypeName(one.TextEdit)
			gotSlice := dynamicTypeName(many[0].TextEdit)
			if gotSingle != gotSlice {
				t.Fatalf("arm divergence: single=%s slice=%s", gotSingle, gotSlice)
			}
		})
	}
}

func dynamicTypeName(v any) string {
	switch v.(type) {
	case nil:
		return "nil"
	case *TextEdit:
		return "*TextEdit"
	case *InsertReplaceEdit:
		return "*InsertReplaceEdit"
	default:
		return "other"
	}
}

// TestAppendFloat64JSONValueClasses pins the jsonwire-compatible formatting
// boundaries (f/e switchover, exponent cleanup, -0 sign) as table cases so a
// regression is named without running the fuzzer.
func TestAppendFloat64JSONValueClasses(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in float64
	}{
		"success: zero":                 {in: 0},
		"success: negative zero":        {in: math.Copysign(0, -1)},
		"success: one":                  {in: 1},
		"success: fraction":             {in: 0.5},
		"success: smallest f boundary":  {in: 1e-6},
		"success: below f boundary":     {in: 9.999999e-7},
		"success: largest f boundary":   {in: 1e20},
		"success: e boundary":           {in: 1e21},
		"success: max float":            {in: math.MaxFloat64},
		"success: smallest subnormal":   {in: math.SmallestNonzeroFloat64},
		"success: negative exponent":    {in: -123456.789e-12},
		"success: shortest round-trip":  {in: 0.1},
		"success: large negative":       {in: -1e21},
		"success: exponent two digits":  {in: 1e-9},
		"success: exponent three digit": {in: 1e-300},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := appendFloat64JSON(nil, tt.in)
			if err != nil {
				t.Fatalf("appendFloat64JSON(%v): %v", tt.in, err)
			}
			want, err := json.Marshal(tt.in, wireOptions)
			if err != nil {
				t.Fatalf("oracle marshal(%v): %v", tt.in, err)
			}
			if !bytes.Equal(got, want) {
				t.Fatalf("appendFloat64JSON(%v) = %s, want %s", tt.in, got, want)
			}
		})
	}

	for name, in := range map[string]float64{
		"error: NaN":  math.NaN(),
		"error: +Inf": math.Inf(1),
		"error: -Inf": math.Inf(-1),
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := appendFloat64JSON(nil, in); err == nil {
				t.Fatalf("appendFloat64JSON(%v): expected error", in)
			}
			if _, err := json.Marshal(in, wireOptions); err == nil {
				t.Fatalf("oracle marshal(%v): expected error (contract drift)", in)
			}
		})
	}
}
