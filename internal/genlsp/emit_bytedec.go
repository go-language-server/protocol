// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"fmt"
	"sort"
	"strings"
)

// This file emits the byte-level decoders: for every eligible generated struct
// a `unmarshalLSP(raw []byte, i int) (int, error)` walker over the raw bytes,
// plus an `unmarshalLSPValue` whole-value entry and an `UnmarshalJSONFrom`
// shim so reflection contexts route into the same walker. Coverage starts from
// all generated structs and named slices; unsupported field shapes fall back to
// decodeWith on the raw sub-value, which keeps correctness independent of
// direct byte-walker coverage.

// byteDecodeExclude is reserved for generated structs that are proven unsafe
// for direct byte walking. It is intentionally empty: broad generated-type
// coverage is the default and code-size tradeoffs are benchmarked instead of
// preemptively excluding capability trees.
var byteDecodeExclude = map[string]bool{}

// byteDecCtx carries the resolved type information the byte-decoder emitters
// need.
type byteDecCtx struct {
	structs      map[string]*renderedStruct
	enumBase     map[string]string // enum name -> string|int32|uint32
	aliasType    map[string]string // alias name -> lowered Go type
	covered      map[string]bool   // struct names with emitted walkers
	coveredSlice map[string]string // named slice wrapper -> element type
	sliceElems   []string          // element types needing a slice helper
	sliceElemSet map[string]bool
	unions       map[string]*unionDecl
	unionCanon   map[string]string // union name (incl. aliases) -> canonical decl name
	unionDecls   []string          // canonical union decl names, emission order
	wrapByName   map[string]string // named slice wrapper -> element type name
}

// unionValueFn names the whole-value decode function for a union type name,
// resolving role-name aliases to the canonical decl that carries the decoder.
func (c *byteDecCtx) unionValueFn(t string) string {
	if canon, ok := c.unionCanon[t]; ok {
		t = canon
	}
	return "unmarshal" + t + "Value"
}

// scalarWrapperBase maps the union scalar wrapper types to their underlying
// Go scalar.
var scalarWrapperBase = map[string]string{
	"String":   "string",
	"Integer":  "int32",
	"Uinteger": "uint32",
	"Decimal":  "float64",
	"Boolean":  "bool",
}

// buildByteDecCtx resolves coverage and type maps for the byte decoders.
func (g *Generator) buildByteDecCtx(structs []*renderedStruct, enums []*renderedEnum, aliases []*renderedAlias) *byteDecCtx {
	c := newByteDecCtx(g, structs, enums, aliases)
	c.computeClosure()
	c.collectSliceElems()
	return c
}

// newByteDecCtx indexes the lowered model declarations.
func newByteDecCtx(g *Generator, structs []*renderedStruct, enums []*renderedEnum, aliases []*renderedAlias) *byteDecCtx {
	c := &byteDecCtx{
		structs:      make(map[string]*renderedStruct, len(structs)),
		enumBase:     make(map[string]string, len(enums)),
		aliasType:    make(map[string]string, len(aliases)),
		covered:      map[string]bool{},
		coveredSlice: map[string]string{},
		sliceElemSet: map[string]bool{},
		unions:       g.unionByName,
		unionCanon:   map[string]string{},
		wrapByName:   make(map[string]string, len(g.arrayWrap)),
	}
	for _, s := range structs {
		c.structs[s.Name] = s
	}
	for _, e := range enums {
		c.enumBase[e.Name] = e.Underlying
	}
	for _, a := range aliases {
		c.aliasType[a.Name] = a.Type
	}
	// URI types are hand-written string types.
	c.aliasType[unionURIWrapperType] = "string"
	c.aliasType[legacyURIRef] = "string"
	c.aliasType[generatedURIType] = "string"
	for _, sig := range g.unionOrder {
		n := g.unions[sig].Name
		c.unionCanon[n] = n
		c.unionDecls = append(c.unionDecls, n)
	}
	for _, a := range g.unionAliases {
		c.unionCanon[a.Name] = a.Target
	}
	for _, w := range g.arrayWrap {
		if elem, ok := strings.CutPrefix(w.Element, "[]"); ok {
			c.wrapByName[w.Name] = elem
		}
	}
	for _, a := range aliases {
		if elem, ok := strings.CutPrefix(a.Type, "[]"); ok {
			c.wrapByName[a.Name] = elem
		}
	}
	return c
}

// walkStructFields invokes fn for every field of s, flattening embedded
// structs the way JSON promotion flattens their members.
func (c *byteDecCtx) walkStructFields(s *renderedStruct, fn func(f *renderedField)) {
	for _, e := range s.Embeds {
		if es, ok := c.structs[e]; ok {
			c.walkStructFields(es, fn)
		}
	}
	for i := range s.Fields {
		fn(&s.Fields[i])
	}
}

// closureWalk carries the in-progress state of the coverage closure.
type closureWalk struct {
	seenUnion map[string]bool
	queue     []string
}

// visitType records coverage references reachable from a field type
// expression, queueing struct names for the closure walk.
func (c *byteDecCtx) visitType(t string, w *closureWalk) {
	switch {
	case t == "":
		return
	case strings.HasPrefix(t, "Optional["):
		c.visitType(strings.TrimSuffix(strings.TrimPrefix(t, "Optional["), "]"), w)
	case strings.HasPrefix(t, "Nullable["):
		c.visitType(strings.TrimSuffix(strings.TrimPrefix(t, "Nullable["), "]"), w)
	case strings.HasPrefix(t, "*"):
		c.visitType(t[1:], w)
	case strings.HasPrefix(t, "[]"):
		c.visitType(t[2:], w)
	case strings.HasPrefix(t, "map["):
		return // fallback territory
	default:
		if elem, ok := c.wrapByName[t]; ok {
			c.coveredSlice[t] = elem
			c.visitType(elem, w)
			return
		}
		if u, ok := c.unions[t]; ok {
			if w.seenUnion[t] {
				return
			}
			w.seenUnion[t] = true
			for _, m := range u.Members {
				c.visitType(strings.TrimPrefix(m.Receiver, "*"), w)
			}
			return
		}
		if _, ok := c.structs[t]; ok {
			w.queue = append(w.queue, t)
		}
	}
}

// computeClosure marks every generated struct and named slice as covered,
// walking referenced structs/unions so generated helpers can call each other
// instead of falling back unnecessarily.
func (c *byteDecCtx) computeClosure() {
	w := &closureWalk{seenUnion: map[string]bool{}}
	structNames := make([]string, 0, len(c.structs))
	for name := range c.structs {
		structNames = append(structNames, name)
	}
	sort.Strings(structNames)
	w.queue = append(w.queue, structNames...)

	wrapNames := make([]string, 0, len(c.wrapByName))
	for name := range c.wrapByName {
		wrapNames = append(wrapNames, name)
	}
	sort.Strings(wrapNames)
	for _, name := range wrapNames {
		elem := c.wrapByName[name]
		c.coveredSlice[name] = elem
		c.visitType(elem, w)
	}
	for len(w.queue) > 0 {
		t := w.queue[0]
		w.queue = w.queue[1:]
		if c.covered[t] || byteDecodeExclude[t] {
			continue
		}
		s, ok := c.structs[t]
		if !ok {
			continue
		}
		c.covered[t] = true
		c.walkStructFields(s, func(f *renderedField) { c.visitType(f.Type, w) })
	}
}

// collectSliceElems records every covered element type that needs a generated
// []T decode helper.
func (c *byteDecCtx) collectSliceElems() {
	need := func(elem string) {
		if c.sliceElemSet[elem] {
			return
		}
		c.sliceElemSet[elem] = true
		c.sliceElems = append(c.sliceElems, elem)
	}
	for _, elem := range c.coveredSlice {
		if c.covered[elem] || c.unions[elem] != nil {
			need(elem)
		}
	}
	var scan func(t string)
	scan = func(t string) {
		switch {
		case strings.HasPrefix(t, "Optional["):
			scan(strings.TrimSuffix(strings.TrimPrefix(t, "Optional["), "]"))
		case strings.HasPrefix(t, "Nullable["):
			scan(strings.TrimSuffix(strings.TrimPrefix(t, "Nullable["), "]"))
		case strings.HasPrefix(t, "[]"):
			elem := t[2:]
			if c.covered[elem] || c.unions[elem] != nil {
				need(elem)
			}
		}
	}
	for name := range c.covered {
		c.walkStructFields(c.structs[name], func(f *renderedField) { scan(f.Type) })
	}
	sort.Strings(c.sliceElems)
}

// armByteCovered reports whether a union arm receiver decodes through an
// emitted byte walker.
func (c *byteDecCtx) armByteCovered(receiver string) bool {
	base := strings.TrimPrefix(receiver, "*")
	if c.covered[base] {
		return true
	}
	elem, ok := c.coveredSlice[base]
	return ok && c.sliceElemSet[elem]
}

// renderByteDecoders emits the walkers, slice helpers, whole-value entries,
// UnmarshalJSONFrom shims, and the union root dispatcher.
func (g *Generator) renderByteDecoders(c *byteDecCtx) string {
	var b strings.Builder

	names := make([]string, 0, len(c.covered))
	for n := range c.covered {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		g.renderByteWalker(&b, c, c.structs[n])
		renderByteValueEntry(&b, n)
	}
	g.renderBoxedByteWalkerVariants(&b, c)

	for _, elem := range c.sliceElems {
		renderByteSliceHelper(&b, c, elem)
	}

	wrapNames := make([]string, 0, len(c.coveredSlice))
	for n := range c.coveredSlice {
		wrapNames = append(wrapNames, n)
	}
	sort.Strings(wrapNames)
	for _, n := range wrapNames {
		elem := c.coveredSlice[n]
		renderByteSliceValueEntry(&b, n, elem, c.sliceElemSet[elem])
	}

	renderUnionRootDispatch(&b, c)
	return b.String()
}

// renderByteValueEntry emits the whole-value entry and reflection shim for a
// covered struct.
func renderByteValueEntry(b *strings.Builder, name string) {
	fmt.Fprintf(b, "func (x *%s) unmarshalLSPValue(raw jsontext.Value) error {\n", name)
	b.WriteString("\ti, err := x.unmarshalLSP(raw, skipSpace(raw, 0))\n")
	b.WriteString("\tif err != nil {\n\t\treturn err\n\t}\n")
	b.WriteString("\treturn dvEnd(raw, i)\n}\n\n")

	fmt.Fprintf(b, "// UnmarshalJSONFrom implements the v2 UnmarshalerFrom interface via the byte walker.\n")
	fmt.Fprintf(b, "func (x *%s) UnmarshalJSONFrom(dec *jsontext.Decoder) error {\n", name)
	b.WriteString("\traw, err := dec.ReadValue()\n\tif err != nil {\n\t\treturn err\n\t}\n")
	b.WriteString("\treturn x.unmarshalLSPValue(slices.Clone(raw))\n}\n\n")
}

// renderByteSliceValueEntry emits the whole-value entry and reflection shim
// for a covered named slice wrapper.
func renderByteSliceValueEntry(b *strings.Builder, name, elem string, direct bool) {
	fmt.Fprintf(b, "func (x *%s) unmarshalLSPValue(raw jsontext.Value) error {\n", name)
	if direct {
		fmt.Fprintf(b, "\tv, i, err := unmarshalSlice%s(raw, skipSpace(raw, 0), []%s(*x))\n", exportName(elem), elem)
		b.WriteString("\tif err != nil {\n\t\treturn err\n\t}\n")
		b.WriteString("\tif err := dvEnd(raw, i); err != nil {\n\t\treturn err\n\t}\n")
		fmt.Fprintf(b, "\t*x = %s(v)\n\treturn nil\n}\n\n", name)
	} else {
		fmt.Fprintf(b, "\tvar v []%s\n", elem)
		b.WriteString("\tif err := decodeWith(raw, &v); err != nil {\n\t\treturn err\n\t}\n")
		fmt.Fprintf(b, "\t*x = %s(v)\n\treturn nil\n}\n\n", name)
	}

	fmt.Fprintf(b, "// UnmarshalJSONFrom implements the v2 UnmarshalerFrom interface via the byte walker.\n")
	fmt.Fprintf(b, "func (x *%s) UnmarshalJSONFrom(dec *jsontext.Decoder) error {\n", name)
	b.WriteString("\traw, err := dec.ReadValue()\n\tif err != nil {\n\t\treturn err\n\t}\n")
	b.WriteString("\treturn x.unmarshalLSPValue(slices.Clone(raw))\n}\n\n")
}

// renderByteSliceHelper emits the []T decode loop for a covered element type.
func renderByteSliceHelper(b *strings.Builder, c *byteDecCtx, elem string) {
	fn := "unmarshalSlice" + exportName(elem)
	fmt.Fprintf(b, "func %s(raw []byte, i int, dst []%s) ([]%s, int, error) {\n", fn, elem, elem)
	b.WriteString("\tif n, ok := dvNull(raw, i); ok {\n\t\treturn nil, n, nil\n\t}\n")
	b.WriteString("\tif i >= len(raw) || raw[i] != '[' {\n\t\treturn dst, i, dvSyntaxError(i, \"array\")\n\t}\n")
	b.WriteString("\ti = skipSpace(raw, i+1)\n\tout := dst[:0]\n")
	b.WriteString("\tif i < len(raw) && raw[i] == ']' {\n")
	fmt.Fprintf(b, "\t\tif out == nil {\n\t\t\tout = []%s{}\n\t\t}\n", elem)
	b.WriteString("\t\treturn out, i + 1, nil\n\t}\n")
	zero := elem + "{}"
	if _, isUnion := c.unions[elem]; isUnion {
		zero = "nil" // union elements are interfaces
	}
	switch elem {
	case "Diagnostic":
		b.WriteString("\tvar scalarBoxes []String\n")
	case "WorkspaceSymbol":
		b.WriteString("\tvar locationBoxes []Location\n")
		b.WriteString("\tvar locationURIOnlyBoxes []LocationUriOnly\n")
	}
	b.WriteString("\tfor {\n")
	b.WriteString("\t\tif len(out) < cap(out) {\n\t\t\tout = out[:len(out)+1]\n")
	fmt.Fprintf(b, "\t\t\tout[len(out)-1] = %s\n", zero)
	b.WriteString("\t\t} else {\n")
	fmt.Fprintf(b, "\t\t\tout = append(out, %s)\n", zero)
	b.WriteString("\t\t}\n")
	switch {
	case elem == "Diagnostic":
		b.WriteString("\t\tn, err := out[len(out)-1].unmarshalLSPWithScalarBoxes(raw, i, &scalarBoxes)\n")
		b.WriteString("\t\tif err != nil {\n\t\t\treturn dst, n, err\n\t\t}\n")
	case elem == "WorkspaceSymbol":
		b.WriteString("\t\tn, err := out[len(out)-1].unmarshalLSPWithLocationBoxes(raw, i, &locationBoxes, &locationURIOnlyBoxes)\n")
		b.WriteString("\t\tif err != nil {\n\t\t\treturn dst, n, err\n\t\t}\n")
	case c.unions[elem] != nil:
		b.WriteString("\t\tval, n, err := dvValue(raw, i)\n")
		b.WriteString("\t\tif err != nil {\n\t\t\treturn dst, n, err\n\t\t}\n")
		fmt.Fprintf(b, "\t\tif err := %s(val, &out[len(out)-1]); err != nil {\n\t\t\treturn dst, i, err\n\t\t}\n", c.unionValueFn(elem))
	default:
		b.WriteString("\t\tn, err := out[len(out)-1].unmarshalLSP(raw, i)\n")
		b.WriteString("\t\tif err != nil {\n\t\t\treturn dst, n, err\n\t\t}\n")
	}
	b.WriteString("\t\tvar done bool\n\t\tvar err2 error\n\t\ti, done, err2 = dvArrayNext(raw, n)\n")
	b.WriteString("\t\tif err2 != nil {\n\t\t\treturn dst, i, err2\n\t\t}\n")
	b.WriteString("\t\tif done {\n\t\t\treturn out, i, nil\n\t\t}\n")
	b.WriteString("\t}\n}\n\n")
}

// renderUnionRootDispatch emits the top-level fast path for decoding directly
// into a union interface variable.
func renderUnionRootDispatch(b *strings.Builder, c *byteDecCtx) {
	names := append([]string(nil), c.unionDecls...)
	sort.Strings(names)
	b.WriteString("// unmarshalUnionRoot decodes data directly into a union interface\n")
	b.WriteString("// pointer without constructing a decoder, reporting whether v was one.\n")
	b.WriteString("func unmarshalUnionRoot(data []byte, v any) (bool, error) {\n")
	b.WriteString("\tvar owned []byte\n")
	b.WriteString("\town := func() []byte {\n")
	b.WriteString("\t\tif owned == nil {\n")
	b.WriteString("\t\t\towned = slices.Clone(data)\n")
	b.WriteString("\t\t}\n")
	b.WriteString("\t\treturn owned\n")
	b.WriteString("\t}\n")
	b.WriteString("\tswitch p := v.(type) {\n")
	for _, n := range names {
		fmt.Fprintf(b, "\tcase *%s:\n\t\treturn true, unmarshal%sValue(own(), p)\n", n, n)
	}
	b.WriteString("\t}\n\treturn false, nil\n}\n\n")
}

// renderByteWalker emits the unmarshalLSP byte walker for s.
func (g *Generator) renderByteWalker(b *strings.Builder, c *byteDecCtx, s *renderedStruct) {
	g.renderByteWalkerMethod(b, c, s, "unmarshalLSP", "", nil)
}

type byteFieldOverride func(b *strings.Builder, c *byteDecCtx, f *renderedField) bool

func (g *Generator) renderBoxedByteWalkerVariants(b *strings.Builder, c *byteDecCtx) {
	if s := c.structs["Diagnostic"]; s != nil && c.covered["Diagnostic"] {
		g.renderByteWalkerMethod(b, c, s, "unmarshalLSPWithScalarBoxes", ", scalarBoxes *[]String", renderDiagnosticBoxedFieldCase)
	}
	if s := c.structs["WorkspaceSymbol"]; s != nil && c.covered["WorkspaceSymbol"] {
		g.renderByteWalkerMethod(b, c, s, "unmarshalLSPWithLocationBoxes", ", locationBoxes *[]Location, locationURIOnlyBoxes *[]LocationUriOnly", renderWorkspaceSymbolBoxedFieldCase)
	}
}

func (g *Generator) renderByteWalkerMethod(b *strings.Builder, c *byteDecCtx, s *renderedStruct, method, extraParams string, override byteFieldOverride) {
	fmt.Fprintf(b, "func (x *%s) %s(raw []byte, i int%s) (int, error) {\n", s.Name, method, extraParams)
	b.WriteString("\tif n, ok := dvNull(raw, i); ok {\n")
	fmt.Fprintf(b, "\t\t*x = %s{}\n\t\treturn n, nil\n\t}\n", s.Name)
	b.WriteString("\tif i >= len(raw) || raw[i] != '{' {\n\t\treturn i, dvSyntaxError(i, \"object\")\n\t}\n")
	b.WriteString("\ti = skipSpace(raw, i+1)\n")
	b.WriteString("\tif i < len(raw) && raw[i] == '}' {\n\t\treturn i + 1, nil\n\t}\n")
	b.WriteString("\tfor {\n")
	b.WriteString("\t\tkey, n, err := dvMemberKey(raw, i)\n")
	b.WriteString("\t\tif err != nil {\n\t\t\treturn n, err\n\t\t}\n")
	b.WriteString("\t\ti = n\n")
	b.WriteString("\t\t_ = key\n")
	b.WriteString("\t\tswitch {\n")
	fields := flattenJSONFields(c.structs, s, map[string]bool{})
	for fi := range fields {
		f := &fields[fi]
		fmt.Fprintf(b, "\t\tcase keyEquals(key, %q):\n", f.JSONName)
		if override != nil && override(b, c, f) {
			continue
		}
		g.renderByteFieldCase(b, c, f)
	}
	b.WriteString("\t\tdefault:\n\t\t\t_, n, err := dvValue(raw, i)\n")
	b.WriteString("\t\t\tif err != nil {\n\t\t\t\treturn n, err\n\t\t\t}\n")
	b.WriteString("\t\t\ti = n\n")
	b.WriteString("\t\t}\n")
	b.WriteString("\t\tvar done bool\n\t\ti, done, err = dvObjectNext(raw, i)\n")
	b.WriteString("\t\tif err != nil {\n\t\t\treturn i, err\n\t\t}\n")
	b.WriteString("\t\tif done {\n\t\t\treturn i, nil\n\t\t}\n")
	b.WriteString("\t}\n}\n\n")
}

func renderDiagnosticBoxedFieldCase(b *strings.Builder, _ *byteDecCtx, f *renderedField) bool {
	const ind = "\t\t\t"
	switch f.Name {
	case "Code":
		emitBoxedUnionField(b, ind, "unmarshalProgressTokenValueBoxed", "x.Code", "scalarBoxes")
		return true
	case "Message":
		emitBoxedUnionField(b, ind, "unmarshalInlayHintTooltipValueBoxed", "x.Message", "scalarBoxes")
		return true
	default:
		return false
	}
}

func renderWorkspaceSymbolBoxedFieldCase(b *strings.Builder, _ *byteDecCtx, f *renderedField) bool {
	const ind = "\t\t\t"
	if f.Name != "Location" {
		return false
	}
	fmt.Fprintf(b, "%sval, n, err := dvValue(raw, i)\n", ind)
	fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n", ind, ind, ind)
	fmt.Fprintf(b, "%sif err := unmarshalWorkspaceSymbolLocationValueBoxed(val, &x.Location, locationBoxes, locationURIOnlyBoxes); err != nil {\n%s\treturn i, err\n%s}\n%si = n\n", ind, ind, ind, ind)
	return true
}

func emitBoxedUnionField(b *strings.Builder, ind, fn, dst, boxes string) {
	fmt.Fprintf(b, "%sval, n, err := dvValue(raw, i)\n", ind)
	fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n", ind, ind, ind)
	fmt.Fprintf(b, "%sif err := %s(val, &%s, %s); err != nil {\n%s\treturn i, err\n%s}\n%si = n\n", ind, fn, dst, boxes, ind, ind, ind)
}

// renderByteFieldCase emits the decode statements for one member case. The
// emitted code may use `i` (current position, already past the colon), `raw`,
// and must leave `i` just past the consumed value or return an error.
func (g *Generator) renderByteFieldCase(b *strings.Builder, c *byteDecCtx, f *renderedField) {
	t := f.Type
	dst := "x." + f.Name
	switch {
	case strings.HasPrefix(t, "Optional["):
		base := strings.TrimSuffix(strings.TrimPrefix(t, "Optional["), "]")
		b.WriteString("\t\t\tif n, ok := dvNull(raw, i); ok {\n")
		fmt.Fprintf(b, "\t\t\t\t%s.Clear()\n\t\t\t\ti = n\n\t\t\t} else {\n", dst)
		g.renderScalarInto(b, c, "\t\t\t\t", base, func(expr string) string {
			return fmt.Sprintf("%s.Set(%s)", dst, expr)
		})
		b.WriteString("\t\t\t}\n")
	case strings.HasPrefix(t, "Nullable["):
		base := strings.TrimSuffix(strings.TrimPrefix(t, "Nullable["), "]")
		b.WriteString("\t\t\tif n, ok := dvNull(raw, i); ok {\n")
		fmt.Fprintf(b, "\t\t\t\t%s = %s{set: true, null: true}\n\t\t\t\ti = n\n\t\t\t} else {\n", dst, t)
		if expr, ok := g.byteInlineDecode(c, base); ok {
			fmt.Fprintf(b, "\t\t\t\tv, n, err := %s\n", expr)
			b.WriteString("\t\t\t\tif err != nil {\n\t\t\t\t\treturn n, err\n\t\t\t\t}\n")
			fmt.Fprintf(b, "\t\t\t\t%s = %s{set: true, value: %s}\n\t\t\t\ti = n\n", dst, t, castV(c, base))
		} else {
			b.WriteString("\t\t\t\tval, n, err := dvValue(raw, i)\n")
			b.WriteString("\t\t\t\tif err != nil {\n\t\t\t\t\treturn n, err\n\t\t\t\t}\n")
			fmt.Fprintf(b, "\t\t\t\tvar v %s\n", base)
			fmt.Fprintf(b, "\t\t\t\tif err := %s; err != nil {\n\t\t\t\t\treturn i, err\n\t\t\t\t}\n", byteValueDecode(c, base, "val", "&v"))
			fmt.Fprintf(b, "\t\t\t\t%s = %s{set: true, value: v}\n\t\t\t\ti = n\n", dst, t)
		}
		b.WriteString("\t\t\t}\n")
	default:
		g.renderPlainField(b, c, t, dst)
	}
}

// renderPlainField handles non-wrapper field types.
func (g *Generator) renderPlainField(b *strings.Builder, c *byteDecCtx, t, dst string) {
	const ind = "\t\t\t"
	switch {
	case t == "DiagnosticTags":
		fmt.Fprintf(b, "%sn, err := dvDiagnosticTags(raw, i, &%s)\n", ind, dst)
		fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n%si = n\n", ind, ind, ind, ind)
	case t == "LSPAny" || t == "jsontext.Value":
		fmt.Fprintf(b, "%sval, n, err := dvValue(raw, i)\n", ind)
		fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n", ind, ind, ind)
		fmt.Fprintf(b, "%s%s = jsontext.Value(val)\n%si = n\n", ind, dst, ind)
	case strings.HasPrefix(t, "*"):
		base := t[1:]
		if c.covered[base] {
			b.WriteString(ind + "if n, ok := dvNull(raw, i); ok {\n")
			fmt.Fprintf(b, "%s\t%s = nil\n%s\ti = n\n%s} else {\n", ind, dst, ind, ind)
			fmt.Fprintf(b, "%s\tif %s == nil {\n%s\t\t%s = new(%s)\n%s\t}\n", ind, dst, ind, dst, base, ind)
			fmt.Fprintf(b, "%s\tn, err := %s.unmarshalLSP(raw, i)\n", ind, dst)
			fmt.Fprintf(b, "%s\tif err != nil {\n%s\t\treturn n, err\n%s\t}\n%s\ti = n\n%s}\n", ind, ind, ind, ind, ind)
			return
		}
		if expr, ok := g.byteInlineDecode(c, base); ok {
			b.WriteString(ind + "if n, ok := dvNull(raw, i); ok {\n")
			fmt.Fprintf(b, "%s\t%s = nil\n%s\ti = n\n%s} else {\n", ind, dst, ind, ind)
			fmt.Fprintf(b, "%s\tv, n, err := %s\n", ind, expr)
			fmt.Fprintf(b, "%s\tif err != nil {\n%s\t\treturn n, err\n%s\t}\n", ind, ind, ind)
			fmt.Fprintf(b, "%s\tif %s == nil {\n%s\t\t%s = new(%s)\n%s\t}\n", ind, dst, ind, dst, base, ind)
			fmt.Fprintf(b, "%s\t*%s = %s\n%s\ti = n\n%s}\n", ind, dst, castV(c, base), ind, ind)
			return
		}
		g.renderFallbackField(b, dst)
	case c.covered[t]:
		fmt.Fprintf(b, "%sn, err := %s.unmarshalLSP(raw, i)\n", ind, dst)
		fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n%si = n\n", ind, ind, ind, ind)
	case c.unions[t] != nil:
		fmt.Fprintf(b, "%sval, n, err := dvValue(raw, i)\n", ind)
		fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n", ind, ind, ind)
		fmt.Fprintf(b, "%sif err := %s(val, &%s); err != nil {\n%s\treturn i, err\n%s}\n%si = n\n", ind, c.unionValueFn(t), dst, ind, ind, ind)
	case strings.HasPrefix(t, "[]"):
		g.renderSliceField(b, c, t[2:], dst)
	default:
		if expr, ok := g.byteInlineDecode(c, t); ok {
			fmt.Fprintf(b, "%sv, n, err := %s\n", ind, expr)
			fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n", ind, ind, ind)
			fmt.Fprintf(b, "%s%s = %s\n%si = n\n", ind, dst, castV(c, t), ind)
			return
		}
		g.renderFallbackField(b, dst)
	}
}

// renderSliceField emits the decode statements for a []T member.
func (g *Generator) renderSliceField(b *strings.Builder, c *byteDecCtx, elem, dst string) {
	const ind = "\t\t\t"
	switch {
	case resolveScalar(c, elem) == "uint32":
		fmt.Fprintf(b, "%sv, n, err := dvUint32Slice(raw, i, %s)\n", ind, dst)
		fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n%s%s = v\n%si = n\n", ind, ind, ind, ind, dst, ind)
	case resolveScalar(c, elem) == "string":
		fmt.Fprintf(b, "%sv, n, err := dvStringSlice(raw, i, %s)\n", ind, dst)
		fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n%s%s = v\n%si = n\n", ind, ind, ind, ind, dst, ind)
	case c.sliceElemSet[elem]:
		fmt.Fprintf(b, "%sv, n, err := unmarshalSlice%s(raw, i, %s)\n", ind, exportName(elem), dst)
		fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n%s%s = v\n%si = n\n", ind, ind, ind, ind, dst, ind)
	default:
		g.renderFallbackField(b, dst)
	}
}

// renderFallbackField decodes one member value through the reflection path.
func (g *Generator) renderFallbackField(b *strings.Builder, dst string) {
	const ind = "\t\t\t"
	fmt.Fprintf(b, "%sval, n, err := dvValue(raw, i)\n", ind)
	fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n", ind, ind, ind)
	fmt.Fprintf(b, "%sif err := decodeWith(val, &%s); err != nil {\n%s\treturn i, err\n%s}\n%si = n\n", ind, dst, ind, ind, ind)
}

// byteInlineDecode returns the scalar decode call expression `v, n, err := …`
// for type t, when t lowers to a scalar.
func (g *Generator) byteInlineDecode(c *byteDecCtx, t string) (string, bool) {
	if t == generatedURIType {
		return "dvURI(raw, i)", true
	}
	switch resolveScalar(c, t) {
	case "string":
		return "dvString(raw, i)", true
	case "uint32":
		return "dvUint32(raw, i)", true
	case "int32":
		return "dvInt32(raw, i)", true
	case "float64":
		return "dvFloat64(raw, i)", true
	case "bool":
		return "dvBool(raw, i)", true
	}
	return "", false
}

// resolveScalar resolves enums, scalar wrappers, and string aliases down to a
// base scalar type name, or "".
func resolveScalar(c *byteDecCtx, t string) string {
	for range 8 { // bounded alias chains
		switch t {
		case "string", "uint32", "int32", "float64", "bool":
			return t
		}
		if base, ok := scalarWrapperBase[t]; ok {
			t = base
			continue
		}
		if base, ok := c.enumBase[t]; ok {
			t = base
			continue
		}
		if next, ok := c.aliasType[t]; ok {
			t = next
			continue
		}
		return ""
	}
	return ""
}

// castV wraps the scalar decode result variable v in a conversion when the
// field type is not the base scalar itself.
func castV(c *byteDecCtx, t string) string {
	if t == generatedURIType {
		return "v"
	}
	if base := resolveScalar(c, t); base != "" && base != t {
		return t + "(v)"
	}
	return "v"
}

// renderScalarInto emits a scalar decode whose result feeds assign (used for
// Optional fields). Falls back to decodeWith when the base is not scalar.
func (g *Generator) renderScalarInto(b *strings.Builder, c *byteDecCtx, ind, base string, assign func(expr string) string) {
	if expr, ok := g.byteInlineDecode(c, base); ok {
		fmt.Fprintf(b, "%sv, n, err := %s\n", ind, expr)
		fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n", ind, ind, ind)
		fmt.Fprintf(b, "%s%s\n%si = n\n", ind, assign(castV(c, base)), ind)
		return
	}
	fmt.Fprintf(b, "%sval, n, err := dvValue(raw, i)\n", ind)
	fmt.Fprintf(b, "%sif err != nil {\n%s\treturn n, err\n%s}\n", ind, ind, ind)
	fmt.Fprintf(b, "%svar v %s\n", ind, base)
	fmt.Fprintf(b, "%sif err := %s; err != nil {\n%s\treturn i, err\n%s}\n", ind, byteValueDecode(c, base, "val", "&v"), ind, ind)
	fmt.Fprintf(b, "%s%s\n%si = n\n", ind, assign("v"), ind)
}

// byteValueDecode returns the whole-value decode expression for a complete
// raw value: byte walkers for covered types, union value functions for
// unions, decodeWith otherwise.
func byteValueDecode(c *byteDecCtx, t, rawExpr, dstExpr string) string {
	switch {
	case c.covered[t]:
		return fmt.Sprintf("(%s).unmarshalLSPValue(%s)", dstExpr, rawExpr)
	case c.unions[t] != nil:
		return fmt.Sprintf("%s(%s, %s)", c.unionValueFn(t), rawExpr, dstExpr)
	default:
		return fmt.Sprintf("decodeWith(%s, %s)", rawExpr, dstExpr)
	}
}
