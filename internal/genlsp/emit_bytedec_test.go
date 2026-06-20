// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"strings"
	"testing"
)

func TestByteDecodeCoverageFromModel(t *testing.T) {
	m, err := Load(t.Context(), "testdata/metaModel.json")
	if err != nil {
		t.Fatalf("load model: %v", err)
	}
	g := NewGenerator(m, "protocol")
	if _, err := g.Emit(); err != nil {
		t.Fatalf("emit: %v", err)
	}
	c := g.byteCtx

	for name := range c.structs {
		if !c.covered[name] {
			t.Errorf("generated struct %s not covered by byte decoders", name)
		}
	}
	for _, want := range []string{"ClientCapabilities", "ServerCapabilities", "SemanticTokensOptionsRange"} {
		if !c.covered[want] {
			t.Errorf("expected generated struct %s to be covered", want)
		}
	}
	for _, want := range []string{"DocumentSelector", "StringSlice", "WorkspaceSymbolSlice"} {
		if _, ok := c.coveredSlice[want]; !ok {
			t.Errorf("expected generated named slice %s to be covered", want)
		}
	}
	// Struct and union slice arms route through byte walkers, not through a
	// per-element reflection shim.
	for _, wrap := range []string{"DocumentSelector", "WorkspaceSymbolSlice", "SymbolInformationSlice", "CompletionItemSlice"} {
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
		"keyEquals(key, `label`)",
		"keyEquals(key, `kind`)",
		"keyEquals(key, `detail`)",
		"keyEquals(key, `documentation`)",
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

func TestRenderByteWalkerCollapsesDuplicateEmbeddedJSONNames(t *testing.T) {
	base := &renderedStruct{
		Name: "Base",
		Fields: []renderedField{
			{Name: "BaseName", Type: "string", JSONName: "baseName", Tag: "baseName"},
			{Name: "Shadow", Type: "string", JSONName: "dup", Tag: "dup"},
		},
	}
	child := &renderedStruct{
		Name:   "Child",
		Embeds: []string{"Base"},
		Fields: []renderedField{
			{Name: "ShadowLocal", Type: "string", JSONName: "dup", Tag: "dup"},
			{Name: "ChildName", Type: "string", JSONName: "childName", Tag: "childName"},
		},
	}
	c := &byteDecCtx{
		structs: map[string]*renderedStruct{
			"Base":  base,
			"Child": child,
		},
		enumBase:     map[string]string{},
		aliasType:    map[string]string{},
		covered:      map[string]bool{},
		coveredSlice: map[string]string{},
		sliceElemSet: map[string]bool{},
		unions:       map[string]*unionDecl{},
		unionCanon:   map[string]string{},
	}
	g := &Generator{byteCtx: c}
	var b strings.Builder
	g.renderByteWalker(&b, c, child)
	got := b.String()

	if count := strings.Count(got, "case keyEquals(key, `dup`):"); count != 1 {
		t.Fatalf("duplicate embedded JSON field cases = %d, want 1:\n%s", count, got)
	}
	if strings.Contains(got, "x.Shadow =") {
		t.Fatalf("duplicate embedded JSON field used embedded field, want local field:\n%s", got)
	}
	if !strings.Contains(got, "x.ShadowLocal =") {
		t.Fatalf("duplicate embedded JSON field did not use local field:\n%s", got)
	}
}

func TestGeneratedByteWalkersCollapseCurrentDuplicateJSONNames(t *testing.T) {
	files, err := NewGenerator(loadTestModel(t), "protocol").Emit()
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}
	decoderFile := string(files["decoders.gen.go"])

	for _, name := range []string{"CreateFile", "DeleteFile", "RenameFile"} {
		body := extractGeneratedFunction(t, decoderFile, "func (x *"+name+") unmarshalLSP(")
		if count := strings.Count(body, "case keyEquals(key, `kind`):"); count != 1 {
			t.Fatalf("%s kind decode cases = %d, want 1:\n%s", name, count, body)
		}
	}
}

func extractGeneratedFunction(t *testing.T, src, signature string) string {
	t.Helper()

	start := strings.Index(src, signature)
	if start < 0 {
		t.Fatalf("generated source missing signature %q", signature)
	}
	rest := src[start:]
	if end := strings.Index(rest[len(signature):], "\n}\n\nfunc "); end >= 0 {
		return rest[:len(signature)+end+len("\n}\n")]
	}
	return rest
}

func TestByteDecodeScalarRecognizesURIType(t *testing.T) {
	c := &byteDecCtx{aliasType: map[string]string{generatedURIType: "string"}}
	if got := resolveScalar(c, generatedURIType); got != "string" {
		t.Fatalf("resolveScalar(uri.URI) = %q, want string", got)
	}
	if got := castV(c, generatedURIType); got != "v" {
		t.Fatalf("castV(uri.URI) = %q, want v", got)
	}
}

func TestRenderPlainFieldDecodesURIInline(t *testing.T) {
	c := &byteDecCtx{
		aliasType:    map[string]string{generatedURIType: "string"},
		covered:      map[string]bool{},
		unions:       map[string]*unionDecl{},
		sliceElemSet: map[string]bool{},
	}
	g := &Generator{}
	var b strings.Builder
	g.renderPlainField(&b, c, generatedURIType, "x.URI")
	got := b.String()
	for _, want := range []string{"dvURI(raw, i)", "x.URI = v"} {
		if !strings.Contains(got, want) {
			t.Fatalf("renderPlainField(uri.URI) missing %q:\n%s", want, got)
		}
	}
}
