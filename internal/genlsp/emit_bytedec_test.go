// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"strings"
	"testing"
)

func TestByteDecodeCoverageFromModel(t *testing.T) {
	m, err := Load("testdata/metaModel.json")
	if err != nil {
		t.Fatalf("load model: %v", err)
	}
	g := NewGenerator(m, "protocol")
	if _, err := g.Emit(); err != nil {
		t.Fatalf("emit: %v", err)
	}
	c := g.byteCtx

	for _, want := range byteDecodeRoots {
		if _, isStruct := c.structs[want]; isStruct && !c.covered[want] {
			t.Errorf("root struct %s not covered by byte decoders", want)
		}
	}
	for _, excluded := range []string{"ClientCapabilities", "ServerCapabilities"} {
		if c.covered[excluded] {
			t.Errorf("excluded capability tree %s gained a byte walker", excluded)
		}
	}
	// The bench-spine slice arms must route through byte walkers, not through
	// a per-element reflection shim.
	for _, wrap := range []string{"WorkspaceSymbolSlice", "SymbolInformationSlice", "CompletionItemSlice"} {
		if !c.armByteCovered(wrap) {
			t.Errorf("union arm %s does not route through the byte walkers", wrap)
		}
	}
	for _, leaf := range []string{"Position", "Range", "Diagnostic", "TextEdit", "WorkspaceSymbol"} {
		if !c.covered[leaf] {
			t.Errorf("expected covered leaf %s", leaf)
		}
	}
}

func TestRenderByteWalkerEmission(t *testing.T) {
	c := &byteDecCtx{
		structs:      map[string]*renderedStruct{},
		enumBase:     map[string]string{"CompletionItemKind": "uint32"},
		aliasType:    map[string]string{},
		covered:      map[string]bool{},
		coveredSlice: map[string]string{},
		sliceElemSet: map[string]bool{},
		unions:       map[string]*unionDecl{"InlayHintTooltip": {Name: "InlayHintTooltip"}},
		unionCanon:   map[string]string{"InlayHintTooltip": "InlayHintTooltip"},
	}
	g := &Generator{byteCtx: c}
	var b strings.Builder
	g.renderByteWalker(&b, c, &renderedStruct{
		Name: "CompletionItem",
		Fields: []renderedField{
			{Name: "Label", Type: "string", JSONName: "label", Tag: "label"},
			{Name: "Kind", Type: "CompletionItemKind", JSONName: "kind", Tag: "kind,omitzero"},
			{Name: "Detail", Type: "Optional[string]", JSONName: "detail", Tag: "detail,omitzero"},
			{Name: "Documentation", Type: "InlayHintTooltip", JSONName: "documentation", Tag: "documentation,omitzero"},
		},
	})
	got := b.String()

	for _, want := range []string{
		"func (x *CompletionItem) unmarshalLSP(raw []byte, i int) (int, error)",
		"keyEquals(key, \"label\")",
		"keyEquals(key, \"kind\")",
		"keyEquals(key, \"detail\")",
		"keyEquals(key, \"documentation\")",
		"dvString(raw, i)",
		"dvUint32(raw, i)",
		"x.Detail.Clear()",
		"unmarshalInlayHintTooltipValue(val, &x.Documentation)",
		"_, n, err := dvValue(raw, i)",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("renderByteWalker() missing %q:\n%s", want, got)
		}
	}
}
