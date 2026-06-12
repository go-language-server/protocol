// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"fmt"
	"sort"
	"strings"
)

// This file emits the byte-append encoders: for every covered generated struct
// an `appendLSP(dst []byte) ([]byte, error)` writer appending canonical JSON
// directly, plus an `appendLSPJSON` root entry (the appendMarshaler interface)
// that pre-sizes the buffer. Coverage reuses the byte-decoder closure
// (byteDecCtx) so encode and decode stay structurally symmetric, and field
// order/omission reuse flattenJSONFields and fieldEncodeCondition so the
// output is byte-identical to the streaming MarshalJSONTo oracle. Field shapes
// outside the closure fall back to appendJSONMarshal on that field only, which
// keeps correctness independent of direct coverage.

// byteEncCtx carries the per-run lookup tables the append-encoder emitters
// need beyond the shared decode closure.
type byteEncCtx struct {
	dec      *byteDecCtx
	enc      *encoderCtx
	tupleOf  map[string]string // tuple wrapper name -> underlying expr
	estimate map[string]int    // struct name -> shallow output size estimate
}

func (g *Generator) newByteEncCtx(c *byteDecCtx, structs []*renderedStruct) *byteEncCtx {
	enc := &encoderCtx{
		g:            g,
		structByName: make(map[string]*renderedStruct, len(structs)),
		zeroHelpers:  map[string]bool{},
		collecting:   map[string]bool{},
	}
	for _, s := range structs {
		enc.structByName[s.Name] = s
	}
	for _, s := range structs {
		enc.collectZeroHelpersForStruct(s)
	}
	e := &byteEncCtx{
		dec:      c,
		enc:      enc,
		tupleOf:  map[string]string{},
		estimate: map[string]int{},
	}
	for expr, d := range g.tupleWrap {
		e.tupleOf[d.Name] = expr
	}
	return e
}

// renderByteEncoders emits the append writers, slice helpers, union arm
// switches, and the union root dispatcher.
func (g *Generator) renderByteEncoders(c *byteDecCtx, structs []*renderedStruct) string {
	e := g.newByteEncCtx(c, structs)
	var b strings.Builder

	names := make([]string, 0, len(c.covered))
	for n := range c.covered {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		g.renderByteEncWriter(&b, e, c.structs[n])
	}

	for _, elem := range c.sliceElems {
		renderByteEncSliceHelper(&b, e, elem)
	}

	wrapNames := make([]string, 0, len(c.coveredSlice))
	for n := range c.coveredSlice {
		wrapNames = append(wrapNames, n)
	}
	sort.Strings(wrapNames)
	for _, n := range wrapNames {
		renderByteEncSliceWrapper(&b, e, n, c.coveredSlice[n])
	}

	tupleNames := make([]string, 0, len(e.tupleOf))
	for n := range e.tupleOf {
		tupleNames = append(tupleNames, n)
	}
	sort.Strings(tupleNames)
	for _, n := range tupleNames {
		renderByteEncTuple(&b, e, n)
	}

	unionNames := append([]string(nil), c.unionDecls...)
	sort.Strings(unionNames)
	for _, n := range unionNames {
		g.renderByteEncUnion(&b, e, n)
	}

	renderByteEncUnionRoot(&b, unionNames)
	return b.String()
}

// ---- size estimation ----

const (
	encEstStringHint = 24
	encEstMaxDepth   = 3
	encEstCap        = 2048
)

// encEstimateStruct returns a shallow canonical-output size estimate for a
// covered struct, used to pre-size the root buffer in one Grow.
func (e *byteEncCtx) encEstimateStruct(name string, depth int) int {
	if est, ok := e.estimate[name]; ok && depth == 0 {
		return est
	}
	s := e.dec.structs[name]
	if s == nil || depth >= encEstMaxDepth {
		return 64
	}
	est := 2
	for _, f := range flattenJSONFields(e.dec.structs, s, map[string]bool{}) {
		est += len(f.JSONName) + 3 + e.encEstimateType(f.Type, depth+1)
		if est >= encEstCap {
			est = encEstCap
			break
		}
	}
	if depth == 0 {
		e.estimate[name] = est
	}
	return est
}

func (e *byteEncCtx) encEstimateType(t string, depth int) int {
	switch {
	case t == "", depth >= encEstMaxDepth:
		return 16
	case strings.HasPrefix(t, "Optional["):
		return e.encEstimateType(strings.TrimSuffix(strings.TrimPrefix(t, "Optional["), "]"), depth)
	case strings.HasPrefix(t, "Nullable["):
		return e.encEstimateType(strings.TrimSuffix(strings.TrimPrefix(t, "Nullable["), "]"), depth)
	case strings.HasPrefix(t, "*"):
		return e.encEstimateType(t[1:], depth)
	case strings.HasPrefix(t, "[]"):
		return 16 + 2*e.encEstimateType(t[2:], depth+1)
	case strings.HasPrefix(t, "map["):
		return 64
	case isURIStringType(t):
		return encEstStringHint
	case t == "LSPAny" || t == "jsontext.Value":
		return 32
	case t == "LSPObject" || t == "LSPArray":
		return 64
	case t == "DiagnosticTags":
		return 8
	}
	if _, ok := e.tupleOf[t]; ok {
		return 16
	}
	if e.dec.unions[t] != nil || e.dec.unionCanon[t] != "" {
		return 48
	}
	if e.dec.covered[t] {
		return e.encEstimateStruct(t, depth)
	}
	switch resolveScalar(e.dec, t) {
	case "string":
		return encEstStringHint
	case "bool":
		return 5
	case "uint32", "int32":
		return 8
	case "float64":
		return 12
	}
	return 32
}

// ---- struct writers ----

func (g *Generator) renderByteEncWriter(b *strings.Builder, e *byteEncCtx, s *renderedStruct) {
	fmt.Fprintf(b, "func (x *%s) appendLSP(dst []byte) ([]byte, error) {\n", s.Name)
	b.WriteString("\tif x == nil {\n\t\treturn append(dst, nullLiteral...), nil\n\t}\n")
	b.WriteString("\tvar err error\n\t_ = err\n")
	b.WriteString("\tdst = append(dst, '{')\n")
	b.WriteString("\tfirst := true\n\t_ = first\n")
	for _, f := range flattenJSONFields(e.dec.structs, s, map[string]bool{}) {
		g.renderByteEncField(b, e, &f)
	}
	b.WriteString("\treturn append(dst, '}'), nil\n}\n\n")

	fmt.Fprintf(b, "// appendLSPJSON implements appendMarshaler with a pre-sized buffer.\n")
	fmt.Fprintf(b, "func (x *%s) appendLSPJSON(dst []byte) ([]byte, error) {\n", s.Name)
	b.WriteString("\tif x == nil {\n\t\treturn append(dst, nullLiteral...), nil\n\t}\n")
	fmt.Fprintf(b, "\tdst = slices.Grow(dst, %d)\n", e.encEstimateStruct(s.Name, 0))
	b.WriteString("\treturn x.appendLSP(dst)\n}\n\n")
}

// renderByteEncField emits the append statements for one JSON member,
// including its omission guard.
func (g *Generator) renderByteEncField(b *strings.Builder, e *byteEncCtx, f *renderedField) {
	cond := e.enc.fieldEncodeCondition(f)
	ind := "\t"
	if cond != "" {
		fmt.Fprintf(b, "\tif %s {\n", cond)
		ind = "\t\t"
	}
	fmt.Fprintf(b, "%sdst = appendObjectName(dst, &first, %s)\n", ind, goJSONNameLiteral(f.JSONName))
	g.renderByteEncValue(b, e, f, ind, cond != "")
	if cond != "" {
		b.WriteString("\t}\n")
	}
}

// renderByteEncValue emits the value append for x.<f.Name>.
func (g *Generator) renderByteEncValue(b *strings.Builder, e *byteEncCtx, f *renderedField, ind string, guarded bool) {
	c := e.dec
	t := f.Type
	src := "x." + f.Name

	switch {
	case strings.HasPrefix(t, "Optional["):
		base := strings.TrimSuffix(strings.TrimPrefix(t, "Optional["), "]")
		if !guarded {
			// An unguarded Optional encodes null when absent; not generated
			// today, so route through the oracle to stay total.
			emitAppendFallback(b, ind, src)
			return
		}
		fmt.Fprintf(b, "%sov, _ := %s.Get()\n", ind, src)
		g.renderByteEncScalarOrValue(b, e, base, "ov", ind)
	case strings.HasPrefix(t, "Nullable["):
		base := strings.TrimSuffix(strings.TrimPrefix(t, "Nullable["), "]")
		fmt.Fprintf(b, "%sif %s.IsNull() {\n", ind, src)
		fmt.Fprintf(b, "%s\tdst = append(dst, nullLiteral...)\n", ind)
		fmt.Fprintf(b, "%s} else {\n", ind)
		fmt.Fprintf(b, "%s\tnv, _ := %s.Get()\n", ind, src)
		g.renderByteEncScalarOrValue(b, e, base, "nv", ind+"\t")
		fmt.Fprintf(b, "%s}\n", ind)
	case strings.HasPrefix(t, "*"):
		base := t[1:]
		fmt.Fprintf(b, "%sif %s == nil {\n", ind, src)
		fmt.Fprintf(b, "%s\tdst = append(dst, nullLiteral...)\n", ind)
		fmt.Fprintf(b, "%s} else {\n", ind)
		switch {
		case c.covered[base]:
			fmt.Fprintf(b, "%s\tif dst, err = %s.appendLSP(dst); err != nil {\n%s\t\treturn nil, err\n%s\t}\n", ind, src, ind, ind)
		default:
			g.renderByteEncScalarOrValue(b, e, base, "*"+src, ind+"\t")
		}
		fmt.Fprintf(b, "%s}\n", ind)
	default:
		g.renderByteEncScalarOrValue(b, e, t, src, ind)
	}
}

// renderByteEncScalarOrValue emits the append statements for a non-pointer,
// non-wrapper value expression src of type t.
func (g *Generator) renderByteEncScalarOrValue(b *strings.Builder, e *byteEncCtx, t, src, ind string) {
	c := e.dec
	switch {
	case t == "DiagnosticTags":
		fmt.Fprintf(b, "%sdst = appendDiagnosticTagsJSON(dst, %s)\n", ind, src)
		return
	case t == "LSPAny" || t == "jsontext.Value":
		fmt.Fprintf(b, "%sif dst, err = appendRawJSONValue(dst, %s); err != nil {\n%s\treturn nil, err\n%s}\n", ind, src, ind, ind)
		return
	case t == "LSPObject" || t == "LSPArray" || strings.HasPrefix(t, "map["):
		emitAppendFallback(b, ind, src)
		return
	case isURIStringType(t):
		fmt.Fprintf(b, "%sdst = appendJSONString(dst, string(%s))\n", ind, src)
		return
	case strings.HasPrefix(t, "[]"):
		g.renderByteEncSliceValue(b, e, t[2:], src, ind)
		return
	case c.covered[t]:
		if e.emitInlineStruct(b, t, src, ind, 0) {
			return
		}
		fmt.Fprintf(b, "%sif dst, err = %s.appendLSP(dst); err != nil {\n%s\treturn nil, err\n%s}\n", ind, src, ind, ind)
		return
	case c.unions[t] != nil || c.unionCanon[t] != "":
		fmt.Fprintf(b, "%sif dst, err = %s(dst, %s); err != nil {\n%s\treturn nil, err\n%s}\n", ind, unionAppendFn(c, t), src, ind, ind)
		return
	}
	if expr, ok := e.tupleOf[t]; ok {
		if elem, _, ok := tupleElemCount(expr); ok && resolveScalar(c, elem) == "uint32" {
			fmt.Fprintf(b, "%sdst = %s.appendLSPTuple(dst)\n", ind, src)
			return
		}
	}
	switch resolveScalar(c, t) {
	case "string":
		fmt.Fprintf(b, "%sdst = appendJSONString(dst, string(%s))\n", ind, src)
	case "uint32":
		fmt.Fprintf(b, "%sdst = appendUint32JSON(dst, uint32(%s))\n", ind, src)
	case "int32":
		fmt.Fprintf(b, "%sdst = appendInt32JSON(dst, int32(%s))\n", ind, src)
	case "float64":
		fmt.Fprintf(b, "%sif dst, err = appendFloat64JSON(dst, float64(%s)); err != nil {\n%s\treturn nil, err\n%s}\n", ind, src, ind, ind)
	case "bool":
		fmt.Fprintf(b, "%sdst = appendBoolJSON(dst, bool(%s))\n", ind, src)
	default:
		emitAppendFallback(b, ind, src)
	}
}

// renderByteEncSliceValue emits the append statements for a []elem value.
func (g *Generator) renderByteEncSliceValue(b *strings.Builder, e *byteEncCtx, elem, src, ind string) {
	c := e.dec
	switch {
	case elem == "uint32":
		// Exact-size single-allocation path shared with the semantic-tokens
		// codec; dominant for SemanticTokens.Data.
		fmt.Fprintf(b, "%sdst = appendUint32JSONArray(dst, %s)\n", ind, src)
	case resolveScalar(c, elem) == "uint32":
		fmt.Fprintf(b, "%sdst = appendUint32SliceJSON(dst, %s)\n", ind, src)
	case resolveScalar(c, elem) == "string" || isURIStringType(elem):
		fmt.Fprintf(b, "%sdst = appendStringSliceJSON(dst, %s)\n", ind, src)
	case c.sliceElemSet[elem]:
		fmt.Fprintf(b, "%sif dst, err = appendSlice%sJSON(dst, %s); err != nil {\n%s\treturn nil, err\n%s}\n", ind, exportName(elem), src, ind, ind)
	default:
		emitAppendFallback(b, ind, src)
	}
}

func emitAppendFallback(b *strings.Builder, ind, src string) {
	fmt.Fprintf(b, "%sif dst, err = appendJSONMarshal(dst, %s); err != nil {\n%s\treturn nil, err\n%s}\n", ind, src, ind, ind)
}

const (
	encInlineMaxDepth  = 2
	encInlineMaxFields = 2
)

// inlinableLeafFields returns the flattened fields of a struct whose encoding
// can be inlined at the call site: at most two always-present fields that are
// either uint32-scalar or themselves inlinable structs. This captures the
// Position/Range leaves that dominate range-bearing payloads without paying a
// call, a nil check, and a first-flag per leaf.
func (e *byteEncCtx) inlinableLeafFields(name string, depth int) ([]renderedField, bool) {
	if depth > encInlineMaxDepth {
		return nil, false
	}
	s := e.dec.structs[name]
	if s == nil || !e.dec.covered[name] {
		return nil, false
	}
	fields := flattenJSONFields(e.dec.structs, s, map[string]bool{})
	if len(fields) == 0 || len(fields) > encInlineMaxFields {
		return nil, false
	}
	for i := range fields {
		f := &fields[i]
		if strings.Contains(f.Tag, ",omitzero") {
			return nil, false
		}
		if resolveScalar(e.dec, f.Type) == "uint32" {
			continue
		}
		if _, ok := e.inlinableLeafFields(f.Type, depth+1); ok {
			continue
		}
		return nil, false
	}
	return fields, true
}

// emitInlineStruct appends an inlinable leaf struct as literal member-name
// bytes plus scalar appends, reporting whether the type qualified.
func (e *byteEncCtx) emitInlineStruct(b *strings.Builder, name, src, ind string, depth int) bool {
	fields, ok := e.inlinableLeafFields(name, depth)
	if !ok {
		return false
	}
	for i := range fields {
		f := &fields[i]
		prefix := ","
		if i == 0 {
			prefix = "{"
		}
		fmt.Fprintf(b, "%sdst = append(dst, `%s%s:`...)\n", ind, prefix, jsonNameQuoted(f.JSONName))
		if resolveScalar(e.dec, f.Type) == "uint32" {
			fmt.Fprintf(b, "%sdst = appendUint32JSON(dst, uint32(%s.%s))\n", ind, src, f.Name)
			continue
		}
		e.emitInlineStruct(b, f.Type, src+"."+f.Name, ind, depth+1)
	}
	fmt.Fprintf(b, "%sdst = append(dst, '}')\n", ind)
	return true
}

// jsonNameQuoted returns the JSON-quoted member name for embedding inside a
// backquoted Go literal; generated member names never contain backquotes.
func jsonNameQuoted(name string) string {
	return `"` + name + `"`
}

// unionAppendFn names the append function for a union type name, resolving
// role-name aliases to the canonical decl.
func unionAppendFn(c *byteDecCtx, t string) string {
	if canon, ok := c.unionCanon[t]; ok {
		t = canon
	}
	return "appendUnion" + t + "JSON"
}

// tupleElemCount parses a "[N]elem" expression.
func tupleElemCount(expr string) (elem string, n int, ok bool) {
	rest, found := strings.CutPrefix(expr, "[")
	if !found {
		return "", 0, false
	}
	num, elem, found := strings.Cut(rest, "]")
	if !found {
		return "", 0, false
	}
	count := 0
	for _, ch := range num {
		if ch < '0' || ch > '9' {
			return "", 0, false
		}
		count = count*10 + int(ch-'0')
	}
	return elem, count, count > 0
}

// ---- slice helpers and wrappers ----

// encSliceGrowHint sizes the per-element Grow constant of generated slice
// helpers. The explicit entries are corpus-calibrated (carried over from the
// retired hand-written encoders, where they kept the bench-spine roots at one
// output allocation); everything else uses a bounded shallow estimate so a
// deep struct estimate cannot balloon the reservation.
var encSliceGrowHint = map[string]int{
	"CompletionItem":                 160,
	"Diagnostic":                     200,
	"WorkspaceSymbol":                112,
	"SymbolInformation":              176,
	"TextDocumentContentChangeEvent": 128,
}

func (e *byteEncCtx) sliceGrowHint(elem string) int {
	if v, ok := encSliceGrowHint[elem]; ok {
		return v
	}
	return min(e.encEstimateType(elem, 1), 96)
}

func renderByteEncSliceHelper(b *strings.Builder, e *byteEncCtx, elem string) {
	c := e.dec
	fmt.Fprintf(b, "func appendSlice%sJSON(dst []byte, x []%s) ([]byte, error) {\n", exportName(elem), elem)
	fmt.Fprintf(b, "\tdst = slices.Grow(dst, 2+len(x)*%d)\n", e.sliceGrowHint(elem))
	b.WriteString("\tdst = append(dst, '[')\n")
	b.WriteString("\tfor i := range x {\n")
	b.WriteString("\t\tif i > 0 {\n\t\t\tdst = append(dst, ',')\n\t\t}\n")
	b.WriteString("\t\tvar err error\n")
	switch {
	case c.unions[elem] != nil:
		fmt.Fprintf(b, "\t\tif dst, err = %s(dst, x[i]); err != nil {\n\t\t\treturn nil, err\n\t\t}\n", unionAppendFn(c, elem))
	default:
		b.WriteString("\t\tif dst, err = (&x[i]).appendLSP(dst); err != nil {\n\t\t\treturn nil, err\n\t\t}\n")
	}
	b.WriteString("\t}\n\treturn append(dst, ']'), nil\n}\n\n")
}

func renderByteEncSliceWrapper(b *strings.Builder, e *byteEncCtx, name, elem string) {
	c := e.dec
	fmt.Fprintf(b, "// appendLSPJSON implements appendMarshaler for the named slice wrapper.\n")
	fmt.Fprintf(b, "func (x %s) appendLSPJSON(dst []byte) ([]byte, error) {\n", name)
	switch {
	case c.sliceElemSet[elem]:
		fmt.Fprintf(b, "\treturn appendSlice%sJSON(dst, []%s(x))\n}\n\n", exportName(elem), elem)
	case resolveScalar(c, elem) == "string":
		fmt.Fprintf(b, "\treturn appendStringSliceJSON(dst, []%s(x)), nil\n}\n\n", elem)
	case resolveScalar(c, elem) == "uint32":
		fmt.Fprintf(b, "\treturn appendUint32SliceJSON(dst, []%s(x)), nil\n}\n\n", elem)
	default:
		b.WriteString("\treturn appendJSONMarshal(dst, x)\n}\n\n")
	}
}

func renderByteEncTuple(b *strings.Builder, e *byteEncCtx, name string) {
	expr := e.tupleOf[name]
	elem, n, ok := tupleElemCount(expr)
	if !ok || resolveScalar(e.dec, elem) != "uint32" {
		return
	}
	fmt.Fprintf(b, "// appendLSPTuple appends the fixed-size tuple as a JSON array.\n")
	fmt.Fprintf(b, "func (x %s) appendLSPTuple(dst []byte) []byte {\n", name)
	b.WriteString("\tdst = append(dst, '[')\n")
	for i := range n {
		if i > 0 {
			b.WriteString("\tdst = append(dst, ',')\n")
		}
		fmt.Fprintf(b, "\tdst = appendUint32JSON(dst, uint32(x[%d]))\n", i)
	}
	b.WriteString("\treturn append(dst, ']')\n}\n\n")
	fmt.Fprintf(b, "// appendLSPJSON implements appendMarshaler for the tuple wrapper.\n")
	fmt.Fprintf(b, "func (x %s) appendLSPJSON(dst []byte) ([]byte, error) {\n", name)
	b.WriteString("\treturn x.appendLSPTuple(dst), nil\n}\n\n")
}

// ---- unions ----

// renderByteEncUnion emits the dynamic-arm append switch for one canonical
// union decl.
func (g *Generator) renderByteEncUnion(b *strings.Builder, e *byteEncCtx, name string) {
	c := e.dec
	u := c.unions[name]
	if u == nil {
		return
	}
	fmt.Fprintf(b, "func appendUnion%sJSON(dst []byte, x %s) ([]byte, error) {\n", name, name)
	b.WriteString("\tswitch v := x.(type) {\n")
	b.WriteString("\tcase nil:\n\t\treturn append(dst, nullLiteral...), nil\n")
	seen := map[string]bool{}
	for _, m := range u.Members {
		if seen[m.Receiver] {
			continue
		}
		seen[m.Receiver] = true
		g.renderByteEncUnionArm(b, e, m)
	}
	b.WriteString("\tdefault:\n\t\treturn appendJSONMarshal(dst, x)\n\t}\n}\n\n")
}

func (g *Generator) renderByteEncUnionArm(b *strings.Builder, e *byteEncCtx, m *unionMember) {
	c := e.dec
	recv := m.Receiver
	base := strings.TrimPrefix(recv, "*")
	switch {
	case strings.HasPrefix(recv, "*") && c.covered[base]:
		fmt.Fprintf(b, "\tcase %s:\n", recv)
		b.WriteString("\t\tif v == nil {\n\t\t\treturn append(dst, nullLiteral...), nil\n\t\t}\n")
		// appendLSPJSON (not appendLSP) so a union value reached at the root
		// still gets its pre-sized buffer; in field context the reservation is
		// already covered and Grow degenerates to a capacity check.
		b.WriteString("\t\treturn v.appendLSPJSON(dst)\n")
	case c.covered[recv]:
		fmt.Fprintf(b, "\tcase %s:\n\t\treturn (&v).appendLSPJSON(dst)\n", recv)
	case c.coveredSlice[recv] != "":
		fmt.Fprintf(b, "\tcase %s:\n\t\treturn v.appendLSPJSON(dst)\n", recv)
	case e.tupleOf[recv] != "":
		fmt.Fprintf(b, "\tcase %s:\n\t\treturn v.appendLSPJSON(dst)\n", recv)
	case recv == "LSPAny":
		fmt.Fprintf(b, "\tcase LSPAny:\n\t\treturn appendRawJSONValue(dst, v)\n")
	case recv == "LSPObject" || recv == "LSPArray":
		fmt.Fprintf(b, "\tcase %s:\n\t\treturn appendJSONMarshal(dst, v)\n", recv)
	case isURIStringType(recv):
		fmt.Fprintf(b, "\tcase %s:\n\t\treturn appendJSONString(dst, string(v)), nil\n", recv)
	default:
		switch resolveScalar(c, recv) {
		case "string":
			fmt.Fprintf(b, "\tcase %s:\n\t\treturn appendJSONString(dst, string(v)), nil\n", recv)
			fmt.Fprintf(b, "\tcase *%s:\n", recv)
			b.WriteString("\t\tif v == nil {\n\t\t\treturn append(dst, nullLiteral...), nil\n\t\t}\n")
			fmt.Fprintf(b, "\t\treturn appendJSONString(dst, string(*v)), nil\n")
		case "uint32":
			fmt.Fprintf(b, "\tcase %s:\n\t\treturn appendUint32JSON(dst, uint32(v)), nil\n", recv)
		case "int32":
			fmt.Fprintf(b, "\tcase %s:\n\t\treturn appendInt32JSON(dst, int32(v)), nil\n", recv)
		case "float64":
			fmt.Fprintf(b, "\tcase %s:\n\t\treturn appendFloat64JSON(dst, float64(v))\n", recv)
		case "bool":
			fmt.Fprintf(b, "\tcase %s:\n\t\treturn appendBoolJSON(dst, bool(v)), nil\n", recv)
		}
	}
}

// renderByteEncUnionRoot emits the root dispatcher for Marshal callers holding
// a pointer to a union interface variable, mirroring unmarshalUnionRoot. Union
// values passed directly resolve through their dynamic arm: struct/slice arms
// implement appendMarshaler themselves, and scalar wrapper arms are handled by
// AppendMarshal's concrete-type switch, so only the pointer-to-interface form
// needs a per-union case here.
func renderByteEncUnionRoot(b *strings.Builder, names []string) {
	b.WriteString("// appendUnionRootJSON appends a union value reached through a pointer to\n")
	b.WriteString("// its interface type, reporting whether v was one.\n")
	b.WriteString("func appendUnionRootJSON(dst []byte, v any) ([]byte, bool, error) {\n")
	b.WriteString("\tswitch p := v.(type) {\n")
	for _, n := range names {
		fmt.Fprintf(b, "\tcase *%s:\n", n)
		fmt.Fprintf(b, "\t\tout, err := appendUnion%sJSON(dst, *p)\n\t\treturn out, true, err\n", n)
	}
	b.WriteString("\t}\n\treturn dst, false, nil\n}\n")
}
