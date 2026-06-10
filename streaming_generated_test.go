// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"testing"

	"github.com/go-json-experiment/json"
)

var (
	_ json.MarshalerTo     = CompletionList{}
	_ json.MarshalerTo     = DocumentSelector{}
	_ json.MarshalerTo     = SemanticTokensOptionsRange{}
	_ json.UnmarshalerFrom = (*CompletionList)(nil)
	_ json.UnmarshalerFrom = (*DocumentSelector)(nil)
	_ json.UnmarshalerFrom = (*SemanticTokensOptionsRange)(nil)
)

func TestGeneratedStreamingMarshalEmbeddedStructPromotesFields(t *testing.T) {
	sym := WorkspaceSymbol{
		BaseSymbolInformation: BaseSymbolInformation{
			Name: "handler",
			Kind: SymbolKindFunction,
		},
		Location: &Location{
			URI: "file:///tmp/main.go",
			Range: Range{
				Start: Position{Line: 1, Character: 2},
				End:   Position{Line: 3, Character: 4},
			},
		},
	}
	got, err := json.Marshal(sym, wireOptions)
	if err != nil {
		t.Fatalf("json.Marshal(WorkspaceSymbol): %v", err)
	}
	const want = `{"name":"handler","kind":12,"location":{"uri":"file:///tmp/main.go","range":{"start":{"line":1,"character":2},"end":{"line":3,"character":4}}}}`
	if string(got) != want {
		t.Fatalf("WorkspaceSymbol JSON mismatch:\ngot:  %s\nwant: %s", got, want)
	}
}

func TestGeneratedStreamingNamedSliceUnmarshalDispatchesUnionElements(t *testing.T) {
	data := []byte(`[{"language":"go"}]`)
	var selector DocumentSelector
	if err := Unmarshal(data, &selector); err != nil {
		t.Fatalf("Unmarshal(DocumentSelector): %v", err)
	}
	if len(selector) != 1 {
		t.Fatalf("DocumentSelector length = %d, want 1", len(selector))
	}
	filter, ok := selector[0].(*TextDocumentFilterLanguage)
	if !ok {
		t.Fatalf("DocumentSelector[0] = %T, want *TextDocumentFilterLanguage", selector[0])
	}
	if filter.Language != "go" {
		t.Fatalf("DocumentSelector[0].Language = %q, want go", filter.Language)
	}
}

func TestGeneratedStreamingSyntheticLiteralRoundTrip(t *testing.T) {
	data := []byte(`{}`)
	var rng SemanticTokensOptionsRange
	if err := Unmarshal(data, &rng); err != nil {
		t.Fatalf("Unmarshal(SemanticTokensOptionsRange): %v", err)
	}
	got, err := json.Marshal(rng, wireOptions)
	if err != nil {
		t.Fatalf("json.Marshal(SemanticTokensOptionsRange): %v", err)
	}
	if string(got) != `{}` {
		t.Fatalf("SemanticTokensOptionsRange JSON = %s, want {}", got)
	}
}
