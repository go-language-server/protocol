// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"testing"
)

func TestObjectGuardEmission(t *testing.T) {
	tests := []struct {
		name     string
		required []string
		all      []string
		wantObj  string
		wantArr  string
	}{
		{
			name:     "required and known",
			required: []string{"uri", "range"},
			all:      []string{"uri", "range", "data"},
			wantObj:  `objectHasAndKnownGuard(raw, []string{"uri", "range"}, []string{"uri", "range", "data"})`,
			wantArr:  `arrayFirstHasAndKnown(raw, []string{"uri", "range"}, []string{"uri", "range", "data"})`,
		},
		{
			name:     "required only",
			required: []string{"kind"},
			wantObj:  `objectHasKeys(raw, "kind")`,
			wantArr:  `arrayFirstHasKeys(raw, "kind")`,
		},
		{
			name:    "known only",
			all:     []string{"title", "command"},
			wantObj: `objectKeysKnown(raw, "title", "command")`,
			wantArr: `arrayFirstKeysKnown(raw, "title", "command")`,
		},
		{
			name: "no signal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := objectGuard(tt.required, tt.all); got != tt.wantObj {
				t.Fatalf("objectGuard() = %q, want %q", got, tt.wantObj)
			}
			if got := arrayGuard(tt.required, tt.all); got != tt.wantArr {
				t.Fatalf("arrayGuard() = %q, want %q", got, tt.wantArr)
			}
		})
	}
}

func TestRenderMarshalersSortsUnionNames(t *testing.T) {
	g := &Generator{
		unions: map[string]*unionDecl{
			"z": {Name: "Zeta"},
			"a": {Name: "Alpha"},
		},
		unionOrder: []string{"z", "a"},
	}

	got := g.renderMarshalers()
	alpha := strings.Index(got, "unmarshalAlpha")
	zeta := strings.Index(got, "unmarshalZeta")
	if alpha < 0 || zeta < 0 {
		t.Fatalf("renderMarshalers() missing expected union hooks:\n%s", got)
	}
	if alpha > zeta {
		t.Fatalf("renderMarshalers() is not sorted by union name:\n%s", got)
	}
}

func TestGeneratedFileName(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"success: section file receives gen suffix": {
			input: "basic_structures.go",
			want:  "basic_structures.gen.go",
		},
		"success: legacy generated suffix is normalized": {
			input: "decoders_generated.go",
			want:  "decoders.gen.go",
		},
		"success: already suffixed name is stable": {
			input: "types_unions.gen.go",
			want:  "types_unions.gen.go",
		},
		"success: non-Go artifact receives literal gen suffix": {
			input: "metaModel.json",
			want:  "metaModel.json.gen",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := generatedFileName(tt.input); got != tt.want {
				t.Fatalf("generatedFileName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestHotOptionalField(t *testing.T) {
	tests := []struct {
		name      string
		owner     string
		fieldName string
		base      string
		optional  bool
		nullable  bool
		want      bool
	}{
		{
			name:      "completion item hot string",
			owner:     "CompletionItem",
			fieldName: "Detail",
			base:      "string",
			optional:  true,
			want:      true,
		},
		{
			name:      "completion item hot bool",
			owner:     "CompletionItem",
			fieldName: "Deprecated",
			base:      "bool",
			optional:  true,
			want:      true,
		},
		{
			name:      "publish diagnostics hot int32",
			owner:     "PublishDiagnosticsParams",
			fieldName: "Version",
			base:      "int32",
			optional:  true,
			want:      true,
		},
		{
			name:      "not allowlisted field",
			owner:     "CompletionItem",
			fieldName: "Label",
			base:      "string",
			optional:  true,
		},
		{
			name:      "not optional",
			owner:     "CompletionItem",
			fieldName: "Detail",
			base:      "string",
		},
		{
			name:      "nullable stays nullable",
			owner:     "CompletionItem",
			fieldName: "Detail",
			base:      "string",
			optional:  true,
			nullable:  true,
		},
		{
			name:      "unsupported base type",
			owner:     "CompletionItem",
			fieldName: "Detail",
			base:      "MarkupContent",
			optional:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hotOptionalField(tt.owner, tt.fieldName, tt.base, tt.optional, tt.nullable)
			if got != tt.want {
				t.Fatalf("hotOptionalField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRenderEncodersEmitsEligibleStructsAndNamedSlices(t *testing.T) {
	g := &Generator{}
	got := g.renderEncoders([]*renderedStruct{
		{
			Name: "Base",
			Fields: []renderedField{
				{Name: "BaseName", Type: "string", JSONName: "baseName", Tag: "baseName"},
				{Name: "Shadow", Type: "string", JSONName: "dup", Tag: "dup"},
			},
		},
		{
			Name: "CompletionItem",
			Fields: []renderedField{
				{Name: "Label", Type: "string", JSONName: "label", Tag: "label"},
				{Name: "Kind", Type: "CompletionItemKind", JSONName: "kind", Tag: "kind,omitzero"},
				{Name: "Detail", Type: "Optional[string]", JSONName: "detail", Tag: "detail,omitzero"},
				{Name: "Documentation", Type: "InlayHintTooltip", JSONName: "documentation", Tag: "documentation,omitzero"},
				{Name: "Command", Type: "Command", JSONName: "command", Tag: "command,omitzero"},
			},
		},
		{
			Name:   "Child",
			Embeds: []string{"Base"},
			Fields: []renderedField{
				{Name: "ShadowLocal", Type: "string", JSONName: "dup", Tag: "dup"},
				{Name: "ChildName", Type: "string", JSONName: "childName", Tag: "childName"},
			},
		},
		{
			Name: "CompletionList",
			Fields: []renderedField{
				{Name: "Items", Type: "[]CompletionItem", JSONName: "items", Tag: "items"},
			},
		},
	}, []*renderedAlias{
		{Name: "DocumentSelector", Type: "[]DocumentFilter"},
	})

	for _, want := range []string{
		"func (x CompletionItem) MarshalJSONTo(enc *jsontext.Encoder) error",
		"func (x CompletionList) MarshalJSONTo(enc *jsontext.Encoder) error",
		"func (x Child) MarshalJSONTo(enc *jsontext.Encoder) error",
		"func (x DocumentSelector) MarshalJSONTo(enc *jsontext.Encoder) error",
		"enc.WriteToken(jsontext.BeginObject)",
		"enc.WriteToken(jsontext.String(`label`))",
		"enc.WriteToken(jsontext.String(x.Label))",
		"enc.WriteToken(jsontext.Uint(uint64(x.Kind)))",
		"if v, ok := x.Detail.Get(); ok",
		"enc.WriteToken(jsontext.String(v))",
		"if v, ok := x.Documentation.(String); ok",
		"!isZeroCommand(x.Command)",
		"json.MarshalEncode(enc, x.Items)",
		"json.MarshalEncode(enc, v)",
		"enc.WriteToken(jsontext.EndObject)",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("renderEncoders() missing %q:\n%s", want, got)
		}
	}
	childStart := strings.Index(got, "func (x Child) MarshalJSONTo")
	if childStart < 0 {
		t.Fatalf("renderEncoders() missing Child encoder:\n%s", got)
	}
	childBody := got[childStart:]
	baseIdx := strings.Index(childBody, "enc.WriteToken(jsontext.String(`baseName`))")
	childIdx := strings.Index(childBody, "enc.WriteToken(jsontext.String(`childName`))")
	if baseIdx < 0 || childIdx < 0 || baseIdx > childIdx {
		t.Fatalf("embedded fields not emitted before local fields:\n%s", childBody)
	}
	if strings.Count(childBody, "enc.WriteToken(jsontext.String(`dup`))") != 1 {
		t.Fatalf("duplicate embedded JSON field was not collapsed:\n%s", childBody)
	}
	if !strings.Contains(childBody, "enc.WriteToken(jsontext.String(x.ShadowLocal))") {
		t.Fatalf("duplicate embedded JSON field did not use local field:\n%s", childBody)
	}
}

func TestGoJSONNameLiteral(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"success: raw identifier": {
			input: "label",
			want:  "`label`",
		},
		"success: raw camel case": {
			input: "selectionRange",
			want:  "`selectionRange`",
		},
		"success: raw tab": {
			input: "tab\tname",
			want:  "`tab\tname`",
		},
		"success: fallback for backtick": {
			input: "bad`name",
			want:  "\"bad`name\"",
		},
		"success: fallback for carriage return": {
			input: "bad\rname",
			want:  "\"bad\\rname\"",
		},
		"success: fallback for newline": {
			input: "bad\nname",
			want:  "\"bad\\nname\"",
		},
		"success: fallback for nul": {
			input: "bad\x00name",
			want:  "\"bad\\x00name\"",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := goJSONNameLiteral(tt.input); got != tt.want {
				t.Fatalf("goJSONNameLiteral(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestEmitProducesDeterministicFileSet(t *testing.T) {
	m := loadTestModel(t)

	first, err := NewGenerator(m, "protocol").Emit()
	if err != nil {
		t.Fatalf("first Emit: %v", err)
	}
	second, err := NewGenerator(m, "protocol").Emit()
	if err != nil {
		t.Fatalf("second Emit: %v", err)
	}

	firstNames := sortedFileNames(first)
	secondNames := sortedFileNames(second)
	if strings.Join(firstNames, "\x00") != strings.Join(secondNames, "\x00") {
		t.Fatalf("generated file names differ:\nfirst:  %q\nsecond: %q", firstNames, secondNames)
	}
	for _, name := range firstNames {
		if len(first[name]) == 0 {
			t.Fatalf("generated file %s is empty", name)
		}
		if !bytes.Equal(first[name], second[name]) {
			t.Fatalf("generated file %s differs between Emit runs", name)
		}
	}
}

func TestEmitIncludesGeneratedMarkerAndExpectedFiles(t *testing.T) {
	files, err := NewGenerator(loadTestModel(t), "protocol").Emit()
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}

	for _, name := range []string{"basic_structures.gen.go", "types_unions.gen.go", "metamodel_messages.gen.go", "marshalers.gen.go", "decoders.gen.go", "encoders.gen.go"} {
		if _, ok := files[name]; !ok {
			t.Fatalf("generated files missing %s", name)
		}
	}
	for name, data := range files {
		if !strings.HasSuffix(name, ".gen.go") {
			t.Fatalf("generated filename %s does not end with .gen.go", name)
		}
		if !bytes.Contains(data, []byte("Code generated by internal/genlsp from metaModel.json; DO NOT EDIT.")) {
			t.Fatalf("%s missing generated marker", name)
		}
		if !bytes.Contains(data, []byte("package protocol")) {
			t.Fatalf("%s generated with wrong package", name)
		}
	}
	unionFile := string(files["types_unions.gen.go"])
	for _, want := range []string{"objectHasAndKnownGuard(", "arrayFirstHasAndKnown(", "objectKind("} {
		if !strings.Contains(unionFile, want) {
			t.Fatalf("types_unions.gen.go missing generated scanner call %q", want)
		}
	}
	decoderFile := string(files["decoders.gen.go"])
	for _, want := range []string{
		"func (x *CompletionItem) unmarshalLSP(raw []byte, i int) (int, error)",
		"func (x *CompletionItem) UnmarshalJSONFrom",
		"func (x *CompletionList) unmarshalLSP(raw []byte, i int) (int, error)",
		"func (x *ClientCapabilities) unmarshalLSP(raw []byte, i int) (int, error)",
		"func (x *DocumentSelector) UnmarshalJSONFrom",
		"func (x *SemanticTokensOptionsRange) UnmarshalJSONFrom",
		"func unmarshalUnionRoot(data []byte, v any) (bool, error)",
		"func unmarshalSliceCompletionItem(",
	} {
		if !strings.Contains(decoderFile, want) {
			t.Fatalf("decoders.gen.go missing byte decoder piece %q", want)
		}
	}
	encoderFile := string(files["encoders.gen.go"])
	for _, want := range []string{
		"func (x CompletionItem) MarshalJSONTo",
		"func (x CompletionList) MarshalJSONTo",
		"func (x ClientCapabilities) MarshalJSONTo",
		"func (x DocumentSelector) MarshalJSONTo",
		"func (x SemanticTokensOptionsRange) MarshalJSONTo",
		"func (x WorkspaceSymbolSlice) MarshalJSONTo",
	} {
		if !strings.Contains(encoderFile, want) {
			t.Fatalf("encoders.gen.go missing encoder piece %q:\n%s", want, encoderFile)
		}
	}
}

func sortedFileNames(files map[string][]byte) []string {
	names := make([]string, 0, len(files))
	for name := range files {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func TestGenLSPCommandTempDirAndStaleCleanup(t *testing.T) {
	tmp := t.TempDir()

	stale := filepath.Join(tmp, "stale_generated.go")
	if err := os.WriteFile(stale, []byte(`// Code generated by internal/genlsp from metaModel.json; DO NOT EDIT.

package protocol

const staleGeneratedShouldBeRemoved = true
`), 0o600); err != nil {
		t.Fatalf("write stale generated file: %v", err)
	}

	handWritten := filepath.Join(tmp, "handwritten.go")
	if err := os.WriteFile(handWritten, []byte("package protocol\n\nconst handwrittenShouldRemain = true\n"), 0o600); err != nil {
		t.Fatalf("write handwritten file: %v", err)
	}

	cmd := exec.CommandContext(t.Context(), "go", "run", "./cmd/genlsp", "-input", "testdata/metaModel.json", "-output", tmp, "-pkg", "protocol")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run ./cmd/genlsp failed: %v\n%s", err, out)
	}
	if bytes.Contains(out, []byte("stale_generated.go")) {
		t.Fatalf("generator wrote stale file instead of removing it:\n%s", out)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatalf("stale generated file stat err = %v, want not exist", err)
	}
	if _, err := os.Stat(handWritten); err != nil {
		t.Fatalf("handwritten file was not preserved: %v", err)
	}

	for _, name := range []string{"types_unions.gen.go", "metamodel_messages.gen.go", "marshalers.gen.go", "decoders.gen.go", "encoders.gen.go"} {
		path := filepath.Join(tmp, name)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read generated %s: %v\noutput:\n%s", name, err, out)
		}
		if !bytes.Contains(data, []byte("Code generated by internal/genlsp")) {
			t.Fatalf("%s missing generated marker", name)
		}
		if !bytes.Contains(data, []byte("package protocol")) {
			t.Fatalf("%s generated with wrong package", name)
		}
	}
}

func TestSanitizeDoc(t *testing.T) {
	const dashProseIn = `@since 3.18.0 - support for relative patterns. Whether clients support
relative patterns depends on the client capability.`
	const dashProseWant = `support for relative patterns. Whether clients support
relative patterns depends on the client capability.`
	const blankRunIn = "Some text.\n\n@since 3.18.0\n\nMore text."
	const blankRunWant = "Some text.\n\nMore text."
	const deprLineIn = "Leading.\n@deprecated use range instead.\nTrailing."
	const deprLineWant = "Leading.\nTrailing."
	tests := map[string]struct {
		input string
		want  string
	}{
		"success: standalone since dropped": {
			input: "@since 3.17.0",
			want:  "",
		},
		"success: since with version word dropped": {
			input: "@since version 3.12.0",
			want:  "",
		},
		"success: since proposed marker dropped": {
			input: "@since 3.18.0 - proposed",
			want:  "",
		},
		"success: trailing period dropped": {
			input: "@since 3.16.0.",
			want:  "",
		},
		"success: since changelog prose preserved": {
			input: "@since 3.17 renamed from ApplyWorkspaceEditResponse",
			want:  "renamed from ApplyWorkspaceEditResponse",
		},
		"success: since additional type prose preserved": {
			input: "@since 3.16.0 additional type InsertReplaceEdit",
			want:  "additional type InsertReplaceEdit",
		},
		"success: since dash prose preserved across lines": {
			input: dashProseIn,
			want:  dashProseWant,
		},
		"success: bare link bracketed as doc link": {
			input: "The result of a {@link CodeLensRequest}.",
			want:  "The result of a [CodeLensRequest].",
		},
		"success: member link brackets target drops display text": {
			input: "A Uri {@link Uri.scheme scheme}, like file.",
			want:  "A Uri [Uri.scheme], like file.",
		},
		"success: member link without display bracketed": {
			input: "See {@link CompletionItem.detail}.",
			want:  "See [CompletionItem.detail].",
		},
		"success: linkcode bracketed as doc link": {
			input: "taken from {@linkcode Command.title}",
			want:  "taken from [Command.title]",
		},
		"success: link with array display text brackets target": {
			input: "returns {@link ColorInformation ColorInformation[]} array",
			want:  "returns [ColorInformation] array",
		},
		"success: link with double space brackets target": {
			input: "a {@link  RelativePattern relative pattern}",
			want:  "a [RelativePattern]",
		},
		"success: deprecated prose dropped": {
			input: "@deprecated Use tags instead.",
			want:  "",
		},
		"success: embedded since collapses blank run": {
			input: blankRunIn,
			want:  blankRunWant,
		},
		"success: embedded deprecated line removed": {
			input: deprLineIn,
			want:  deprLineWant,
		},
		"success: inline since keeps terminating period": {
			input: "The use of a string as a document filter is deprecated @since 3.16.0.",
			want:  "The use of a string as a document filter is deprecated.",
		},
		"success: inline since mid-sentence collapses spacing": {
			input: "Foo bar @since 3.17.0 baz.",
			want:  "Foo bar baz.",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := sanitizeDoc(tt.input); got != tt.want {
				t.Errorf("sanitizeDoc(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDetectImportsIncludesURI(t *testing.T) {
	got := detectImports("type X struct { URI uri.URI }\n")
	if !slices.Contains(got, uriImportPath) {
		t.Fatalf("detectImports(uri.URI) = %v, want %s", got, uriImportPath)
	}
}

func TestEncoderScalarBaseTreatsURIAsString(t *testing.T) {
	ctx := &encoderCtx{}
	for _, typ := range []string{unionURIWrapperType, legacyURIRef, generatedURIType} {
		if got := ctx.scalarBase(typ); got != "string" {
			t.Fatalf("scalarBase(%q) = %q, want string", typ, got)
		}
	}
}

func TestWriteEncoderValueEmitsURIAsString(t *testing.T) {
	ctx := &encoderCtx{}
	var b strings.Builder
	writeEncoderValue(ctx, &b, &renderedField{Name: "URI", Type: generatedURIType}, "\t", false)
	got := b.String()
	if !strings.Contains(got, "jsontext.String(string(x.URI))") {
		t.Fatalf("writeEncoderValue(uri.URI) =\n%s\nwant string conversion", got)
	}
}
