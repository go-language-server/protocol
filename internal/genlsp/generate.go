// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"fmt"
	"sort"
	"strings"
)

const (
	generatedURIType    = "uri.URI"
	uriImportPath       = "go.lsp.dev/uri"
	uriPackageQualifier = "uri"
	unionURIWrapperType = "URI"
	legacyURIRef        = "URI"
)

func isURIStringType(t string) bool {
	return t == generatedURIType || t == unionURIWrapperType || t == legacyURIRef
}

// Generator lowers a [MetaModel] into Go declarations for the lsp package.
//
// Synthesized declarations (sealed-interface unions, scalar wrappers, inline
// struct literals, and-merges, named tuple/array wrappers) are accumulated in
// deterministic registries keyed by a structural signature so that identical
// shapes are emitted exactly once and regeneration is idempotent.
type Generator struct {
	model *MetaModel
	pkg   string

	structures map[string]*Structure
	enums      map[string]*Enumeration
	aliases    map[string]*TypeAlias

	unions       map[string]*unionDecl // keyed by structural signature
	unionByName  map[string]*unionDecl // name collision guard
	literals     map[string]*literalDecl
	ands         map[string]*andDecl
	arrayWrap    map[string]*arrayWrapDecl // element-expr -> wrapper
	tupleWrap    map[string]*tupleWrapDecl
	byteCtx      *byteDecCtx           // byte-decoder coverage, resolved during Emit
	scalarUsed   map[BaseTypeName]bool // scalar wrappers referenced by unions
	literalOrder []string              // signature insertion order
	andOrder     []string
	arrayOrder   []string
	tupleOrder   []string
	unionOrder   []string

	unionAliases []unionAlias    // role-name aliases to a canonical union shape
	aliasNames   map[string]bool // names reserved by unionAliases

	warnings []string
}

// unionAlias is a Go type alias to a canonical union, emitted when several use
// sites share one structural shape but each reads better under its own name
// (e.g. ImplementationResult = DefinitionResult).
type unionAlias struct{ Name, Target string }

// unionDecl is a synthesized sealed-interface union.
type unionDecl struct {
	Name    string
	Members []*unionMember
	Doc     string
}

// unionMember is one arm of a [unionDecl].
type unionMember struct {
	GoType       string   // type expression used at the call site (e.g. "*Location", "Integer")
	Receiver     string   // receiver type for the marker method (e.g. "*Location", "Integer")
	Token        byte     // primary JSON token: '{' '[' '"' '0' 't' 'n'
	IsObject     bool     // resolved to a struct/literal
	KindConst    string   // discriminator value when the object has a const "kind" property
	Required     []string // required JSON field names when IsObject
	AllKeys      []string // all JSON field names (required+optional, incl. inherited) when IsObject
	ElemRequired []string // required JSON field names of the element when Token=='['
	ElemAllKeys  []string // all JSON field names of the element when Token=='['
}

// literalDecl is a synthesized struct from an inline object literal.
type literalDecl struct {
	Name string
	Lit  *StructureLiteral
}

// andDecl is a synthesized struct merging the operands of an "and" type.
type andDecl struct {
	Name     string
	Operands []string // embedded Go type names
}

// arrayWrapDecl is a named slice type so an array can be a union arm.
type arrayWrapDecl struct {
	Name    string
	Element string
}

// tupleWrapDecl is a named array/struct type so a tuple can be a union arm.
type tupleWrapDecl struct {
	Name string
	Expr string // underlying type expression, e.g. "[2]uint32"
}

// NewGenerator indexes the model and returns a ready generator.
func NewGenerator(m *MetaModel, pkg string) *Generator {
	g := &Generator{
		model:       m,
		pkg:         pkg,
		structures:  make(map[string]*Structure, len(m.Structures)),
		enums:       make(map[string]*Enumeration, len(m.Enumerations)),
		aliases:     make(map[string]*TypeAlias, len(m.TypeAliases)),
		unions:      make(map[string]*unionDecl),
		unionByName: make(map[string]*unionDecl),
		literals:    make(map[string]*literalDecl),
		ands:        make(map[string]*andDecl),
		arrayWrap:   make(map[string]*arrayWrapDecl),
		tupleWrap:   make(map[string]*tupleWrapDecl),
		scalarUsed:  make(map[BaseTypeName]bool),
		aliasNames:  make(map[string]bool),
	}
	for _, s := range m.Structures {
		g.structures[s.Name] = s
	}
	for _, e := range m.Enumerations {
		g.enums[e.Name] = e
	}
	for _, a := range m.TypeAliases {
		g.aliases[a.Name] = a
	}
	return g
}

// wellKnownAny reports the three meta names lowered to the raw jsontext.Value
// family rather than generated declarations.
func wellKnownAny(name string) (string, bool) {
	switch name {
	case "LSPAny":
		return "LSPAny", true
	case "LSPObject":
		return "LSPObject", true
	case "LSPArray":
		return "LSPArray", true
	}
	return "", false
}

// baseGoType maps a base type name to its Go type in field context.
func baseGoType(name BaseTypeName) string {
	switch name {
	case BaseString, BaseRegExp:
		return "string"
	case BaseBoolean:
		return "bool"
	case BaseInteger:
		return "int32"
	case BaseUinteger:
		return "uint32"
	case BaseDecimal:
		return "float64"
	case BaseURI, BaseDocumentURI:
		return generatedURIType
	case BaseNull:
		// A standalone null (e.g. a void result type) is a raw JSON value that
		// is always null; LSPAny round-trips it without an "any" field.
		return "LSPAny"
	default:
		return "LSPAny"
	}
}

// lower returns the Go type expression for a meta type in field context.
// hint seeds names for synthesized declarations (literals, ands, tuples).
func (g *Generator) lower(t *Type, hint string) string {
	switch t.Kind {
	case KindBase:
		return baseGoType(BaseTypeName(t.Name))
	case KindReference:
		if w, ok := wellKnownAny(t.Name); ok {
			return w
		}
		return t.Name
	case KindArray:
		return "[]" + g.lower(t.Element, hint+"Elem")
	case KindMap:
		return "map[" + g.mapKey(t.Key) + "]" + g.lower(t.Value, hint+"Value")
	case KindOr:
		return g.lowerOr(t, hint)
	case KindAnd:
		return g.andType(t, hint)
	case KindTuple:
		return g.tupleType(t, hint, false)
	case KindLiteral:
		return g.literalType(t.Literal, hint)
	case KindStringLiteral:
		return "string"
	case KindIntegerLiteral:
		return "int32"
	case KindBooleanLiteral:
		return "bool"
	default:
		g.warnf("lower: unhandled kind %q", t.Kind)
		return "LSPAny"
	}
}

// mapKey lowers a map key type (base or reference resolving to string/integer).
func (g *Generator) mapKey(t *Type) string {
	switch t.Kind {
	case KindBase:
		return baseGoType(BaseTypeName(t.Name))
	case KindReference:
		if w, ok := wellKnownAny(t.Name); ok {
			return w
		}
		return t.Name
	default:
		g.warnf("mapKey: unexpected kind %q", t.Kind)
		return "string"
	}
}

// isNullable reports whether an "or" type includes a null arm at top level.
func isNullable(t *Type) bool {
	if t.Kind != KindOr {
		return false
	}
	for _, it := range t.Items {
		if it.Kind == KindBase && BaseTypeName(it.Name) == BaseNull {
			return true
		}
	}
	return false
}

// nonNullItems returns the non-null arms of an or type, recursively flattening
// nested or arms (or-of-or) and nested union type-alias references.
func (g *Generator) nonNullItems(items []*Type) []*Type {
	var out []*Type
	for _, it := range items {
		switch {
		case it.Kind == KindBase && BaseTypeName(it.Name) == BaseNull:
			continue
		case it.Kind == KindOr:
			out = append(out, g.nonNullItems(it.Items)...)
		case it.Kind == KindReference:
			if a, ok := g.aliases[it.Name]; ok && a.Type.Kind == KindOr {
				if _, isAny := wellKnownAny(it.Name); !isAny {
					out = append(out, g.nonNullItems(a.Type.Items)...)
					continue
				}
			}
			out = append(out, it)
		default:
			out = append(out, it)
		}
	}
	return out
}

// lowerOr lowers an or type: a single non-null arm collapses to that arm's Go
// type (Tier 0); two or more arms become a sealed-interface union (Tier 1).
//
// A union containing LSPAny/LSPObject/LSPArray accepts arbitrary JSON, so it
// degrades to the raw LSPAny representation: those names alias map/slice types
// that cannot carry the marker method a sealed interface requires.
func (g *Generator) lowerOr(t *Type, hint string) string {
	items := g.nonNullItems(t.Items)
	for _, it := range items {
		if it.Kind == KindReference {
			if _, ok := wellKnownAny(it.Name); ok {
				return "LSPAny"
			}
		}
	}
	switch len(items) {
	case 0:
		return "LSPAny"
	case 1:
		return g.lower(items[0], hint)
	default:
		return g.union(items, hint).Name
	}
}

// Warnings returns non-fatal diagnostics accumulated during lowering.
func (g *Generator) Warnings() []string { return g.warnings }

// orIsRaw reports whether an or's arms include LSPAny/LSPObject/LSPArray, in
// which case the union degrades to the raw LSPAny representation.
func (g *Generator) orIsRaw(items []*Type) bool {
	for _, it := range items {
		if it.Kind == KindReference {
			if _, ok := wellKnownAny(it.Name); ok {
				return true
			}
		}
	}
	return false
}

// orLowersToUnion reports whether an or type becomes a sealed-interface union
// (two or more non-null, non-raw arms).
func (g *Generator) orLowersToUnion(t *Type) bool {
	if t.Kind != KindOr {
		return false
	}
	items := g.nonNullItems(t.Items)
	return len(items) >= 2 && !g.orIsRaw(items)
}

// union finds or synthesizes the sealed-interface union for the given arms.
func (g *Generator) union(items []*Type, hint string) *unionDecl {
	members := make([]*unionMember, 0, len(items))
	for _, it := range items {
		members = append(members, g.member(it, hint))
	}
	// Structural signature for dedup: sorted member receiver expressions.
	sig := unionSignature(members)
	if u, ok := g.unions[sig]; ok {
		return u
	}
	name := g.uniqueUnionName(members)
	u := &unionDecl{Name: name, Members: members}
	g.unions[sig] = u
	g.unionByName[name] = u
	g.unionOrder = append(g.unionOrder, sig)
	return u
}

func unionSignature(members []*unionMember) string {
	parts := make([]string, len(members))
	for i, m := range members {
		parts[i] = m.Receiver
	}
	sort.Strings(parts)
	return strings.Join(parts, "|")
}

// shortName returns the segment used to build a union's interface name.
func shortName(m *unionMember) string {
	return strings.TrimPrefix(m.Receiver, "*")
}

func (g *Generator) uniqueUnionName(members []*unionMember) string {
	parts := make([]string, len(members))
	for i, m := range members {
		parts[i] = shortName(m)
	}
	return g.uniqueName(strings.Join(parts, "Or"))
}

// preferUnionName registers name as the interface name for the union formed by
// items, unless that structural shape is already named, in which case the
// existing (higher-priority or earlier) name is kept. It returns the canonical
// name now in effect for the shape so callers can record aliases. A role name
// that collides with an existing declaration receives a numeric suffix.
func (g *Generator) preferUnionName(items []*Type, name, doc string) string {
	members := make([]*unionMember, 0, len(items))
	for _, it := range items {
		members = append(members, g.member(it, name))
	}
	sig := unionSignature(members)
	if u, ok := g.unions[sig]; ok {
		return u.Name
	}
	if g.reserved(name) {
		// The role name collides with an existing declaration (often a sibling
		// arm). Rather than mint name+"2", leave the shape unnamed so the main
		// pass falls back to its structural AOrB name.
		return ""
	}
	u := &unionDecl{Name: name, Members: members, Doc: doc}
	g.unions[sig] = u
	g.unionByName[name] = u
	g.unionOrder = append(g.unionOrder, sig)
	return name
}

// registerResultUnionNames names request-result unions after their method
// (e.g. textDocument/definition -> DefinitionResult) and records aliases when
// several methods share one result shape. Requests are processed in method
// order so the canonical name is stable regardless of model ordering.
func (g *Generator) registerResultUnionNames() {
	reqs := append([]*Request(nil), g.model.Requests...)
	sort.SliceStable(reqs, func(i, j int) bool { return reqs[i].Method < reqs[j].Method })
	for _, r := range reqs {
		if r.Result == nil || !g.orLowersToUnion(r.Result) {
			continue
		}
		want := resultUnionName(r)
		got := g.preferUnionName(g.nonNullItems(r.Result.Items), want, "")
		if got != "" && got != want && !g.reserved(want) {
			g.unionAliases = append(g.unionAliases, unionAlias{Name: want, Target: got})
			g.aliasNames[want] = true
		}
	}
}

// registerPropertyUnionNames names struct-property unions after the property
// (and, for an array-of-union property, after the singularized property). It
// runs after results so result names win shared shapes. Structures are visited
// in model order so that any nested literal/tuple a union arm synthesizes keeps
// the same first-use-site name it had before role naming was introduced.
func (g *Generator) registerPropertyUnionNames() {
	for _, s := range g.model.Structures {
		for _, p := range s.Properties {
			switch {
			case g.orLowersToUnion(p.Type):
				g.preferUnionName(g.nonNullItems(p.Type.Items),
					g.propertyUnionName(s.Name, p.Name, false), "")
			case p.Type.Kind == KindArray && g.orLowersToUnion(p.Type.Element):
				g.preferUnionName(g.nonNullItems(p.Type.Element.Items),
					g.propertyUnionName(s.Name, p.Name, true), "")
			}
		}
	}
}

// resultUnionName derives a request's result-union name from its typeName
// (DefinitionRequest -> DefinitionResult), falling back to the method's final
// path segment when typeName is absent.
func resultUnionName(r *Request) string {
	base := strings.TrimSuffix(r.TypeName, "Request")
	if base == "" {
		seg := r.Method
		if i := strings.LastIndex(seg, "/"); i >= 0 {
			seg = seg[i+1:]
		}
		base = seg
	}
	return exportName(base) + "Result"
}

// propertyUnionNameOverrides hand-names the few property unions whose
// mechanical owner/field name is awkward: TextDocumentEdit.edits would
// singularize to the too-generic "Edit", and NotebookDocumentFilterWithNotebook
// .notebook would stutter to "NotebookDocumentFilterWithNotebookNotebook".
//
// Keyed by "Owner.property"; an upstream rename of either silently reverts that
// union to its mechanical name, so revisit this table when metaModel.json moves.
var propertyUnionNameOverrides = map[string]string{
	"TextDocumentEdit.edits":                      "TextDocumentEditElement",
	"NotebookDocumentFilterWithNotebook.notebook": "NotebookDocumentFilterNotebook",
}

// propertyUnionName derives a property union's name. An explicit override wins;
// otherwise fields of *Capabilities structs and array-element unions drop the
// owner prefix (HoverProvider, DocumentChange) while every other property is
// owner-qualified (ParameterInformationLabel). A bare name that collides with an
// existing declaration is owner-qualified to disambiguate.
func (g *Generator) propertyUnionName(owner, field string, arrayElem bool) string {
	if n, ok := propertyUnionNameOverrides[owner+"."+field]; ok {
		return n
	}
	f := exportName(field)
	var base string
	switch {
	case arrayElem:
		base = singularize(f)
	case strings.HasSuffix(owner, "Capabilities"):
		base = f
	default:
		base = exportName(owner) + f
	}
	if g.reserved(base) && !strings.HasPrefix(base, exportName(owner)) {
		base = exportName(owner) + base
	}
	return base
}

// singularize converts a plural identifier to its singular form for naming an
// array-element union after its (plural) property (documentChanges ->
// DocumentChange). It handles the common English endings present in the model.
func singularize(s string) string {
	switch {
	case strings.HasSuffix(s, "ies") && len(s) > 3:
		return s[:len(s)-3] + "y"
	case strings.HasSuffix(s, "ches") || strings.HasSuffix(s, "shes") ||
		strings.HasSuffix(s, "ses") || strings.HasSuffix(s, "xes"):
		return s[:len(s)-2]
	case strings.HasSuffix(s, "s") && !strings.HasSuffix(s, "ss") && len(s) > 1:
		return s[:len(s)-1]
	default:
		return s
	}
}

// member resolves one union arm into its Go representation and the metadata
// required to generate a discriminating decoder.
func (g *Generator) member(t *Type, hint string) *unionMember {
	switch t.Kind {
	case KindBase:
		return g.scalarMember(BaseTypeName(t.Name))
	case KindReference:
		return g.referenceMember(t, hint)
	case KindArray:
		elem := g.lower(t.Element, hint+"Elem")
		w := g.arrayWrapper(elem)
		m := &unionMember{GoType: w, Receiver: w, Token: '['}
		if t.Element.Kind == KindReference {
			if s := g.resolveStruct(t.Element.Name); s != nil {
				m.ElemRequired = g.structRequired(s)
				m.ElemAllKeys = g.structAllKeys(s)
			}
		}
		return m
	case KindTuple:
		w := g.tupleType(t, hint, true)
		return &unionMember{GoType: w, Receiver: w, Token: '['}
	case KindLiteral:
		name := g.literalType(t.Literal, hint)
		return &unionMember{
			GoType: "*" + name, Receiver: "*" + name, Token: '{', IsObject: true,
			Required: requiredFieldNames(t.Literal.Properties),
			AllKeys:  allFieldNames(t.Literal.Properties),
		}
	case KindMap:
		gt := g.lower(t, hint)
		w := g.mapWrapper(gt)
		return &unionMember{GoType: w, Receiver: w, Token: '{', IsObject: true}
	case KindStringLiteral:
		return g.scalarMember(BaseString)
	case KindIntegerLiteral:
		return g.scalarMember(BaseInteger)
	case KindBooleanLiteral:
		return g.scalarMember(BaseBoolean)
	case KindAnd:
		name := g.andType(t, hint)
		return &unionMember{GoType: "*" + name, Receiver: "*" + name, Token: '{', IsObject: true}
	default:
		g.warnf("union member: unhandled kind %q; treating as raw object", t.Kind)
		w := g.mapWrapper("map[string]jsontext.Value")
		return &unionMember{GoType: w, Receiver: w, Token: '{', IsObject: true}
	}
}

// scalarMember returns the canonical wrapper member for a base scalar.
func (g *Generator) scalarMember(name BaseTypeName) *unionMember {
	g.scalarUsed[name] = true
	wrap, token := scalarWrapper(name)
	return &unionMember{GoType: wrap, Receiver: wrap, Token: token}
}

// scalarWrapper returns the named wrapper type and JSON token for a base scalar.
func scalarWrapper(name BaseTypeName) (wrap string, token byte) {
	switch name {
	case BaseString, BaseRegExp:
		return "String", '"'
	case BaseBoolean:
		return "Boolean", 't'
	case BaseInteger:
		return "Integer", '0'
	case BaseUinteger:
		return "Uinteger", '0'
	case BaseDecimal:
		return "Decimal", '0'
	case BaseURI, BaseDocumentURI:
		return unionURIWrapperType, '"'
	default:
		return "String", '"'
	}
}

// referenceMember resolves a reference arm to a struct/enum/alias member.
func (g *Generator) referenceMember(t *Type, _ string) *unionMember {
	name := t.Name
	if w, ok := wellKnownAny(name); ok {
		// LSPAny/LSPObject/LSPArray as a union arm: treat as raw value member.
		token := byte('{')
		if name == "LSPArray" {
			token = '['
		}
		return &unionMember{GoType: w, Receiver: w, Token: token}
	}
	if s, ok := g.structures[name]; ok {
		return &unionMember{
			GoType: "*" + name, Receiver: "*" + name, Token: '{', IsObject: true,
			KindConst: g.structKindConst(s),
			Required:  g.structRequired(s),
			AllKeys:   g.structAllKeys(s),
		}
	}
	if e, ok := g.enums[name]; ok {
		token := byte('"')
		if e.Type.Name == BaseInteger || e.Type.Name == BaseUinteger {
			token = '0'
		}
		return &unionMember{GoType: name, Receiver: name, Token: token}
	}
	if a, ok := g.aliases[name]; ok {
		// Alias to a struct: treat like that struct for disambiguation, but keep
		// the alias's own (defined) Go type as the arm.
		if s := g.resolveStruct(name); s != nil {
			return &unionMember{
				GoType: name, Receiver: name, Token: '{', IsObject: true,
				KindConst: g.structKindConst(s), Required: g.structRequired(s),
				AllKeys: g.structAllKeys(s),
			}
		}
		// Other non-union alias: token from the aliased type; the alias is a
		// defined Go type that carries its own marker method.
		token := g.aliasToken(a.Type)
		return &unionMember{GoType: name, Receiver: name, Token: token, IsObject: token == '{'}
	}
	g.warnf("union member: unresolved reference %q", name)
	return &unionMember{GoType: name, Receiver: name, Token: '{'}
}

// aliasToken returns the primary JSON token of an aliased type.
func (g *Generator) aliasToken(t *Type) byte {
	switch t.Kind {
	case KindBase:
		if BaseTypeName(t.Name) == BaseNull {
			return 'n'
		}
		_, tok := scalarWrapper(BaseTypeName(t.Name))
		return tok
	case KindArray, KindTuple:
		return '['
	case KindMap, KindLiteral:
		return '{'
	case KindReference:
		if _, ok := g.structures[t.Name]; ok {
			return '{'
		}
		if e, ok := g.enums[t.Name]; ok {
			if e.Type.Name == BaseInteger || e.Type.Name == BaseUinteger {
				return '0'
			}
			return '"'
		}
		if a, ok := g.aliases[t.Name]; ok {
			return g.aliasToken(a.Type)
		}
	}
	return '{'
}

// resolveStruct follows reference type-aliases until it reaches a structure,
// returning nil if the name does not ultimately denote one.
func (g *Generator) resolveStruct(name string) *Structure {
	seen := map[string]bool{}
	for name != "" && !seen[name] {
		seen[name] = true
		if s, ok := g.structures[name]; ok {
			return s
		}
		if a, ok := g.aliases[name]; ok && a.Type.Kind == KindReference {
			name = a.Type.Name
			continue
		}
		return nil
	}
	return nil
}

// structKindConst returns the const value of a struct's "kind" string-literal
// discriminator property, if any.
func (g *Generator) structKindConst(s *Structure) string {
	for _, p := range g.allProperties(s) {
		if p.Name == "kind" && p.Type.Kind == KindStringLiteral {
			return p.Type.StringValue
		}
	}
	return ""
}

// structRequired returns the JSON names of a struct's required properties,
// including those inherited via extends and mixins.
func (g *Generator) structRequired(s *Structure) []string {
	var out []string
	for _, p := range g.allProperties(s) {
		if !p.Optional {
			out = append(out, p.Name)
		}
	}
	return out
}

// optionalDropsPointer reports whether an optional field of type goType may be
// emitted as a plain value with ",omitzero" instead of a pointer. This holds when
// the type's zero value can never be a legitimate present value, so omitting it
// via omitzero is indistinguishable from absent:
//
//   - structs whose zero is not meaningful (a required field has an invalid zero);
//   - enums with no zero member (their zero, 0 or "", is never a valid value);
//   - named scalar aliases whose underlying type drops the pointer (string aliases).
//
// Pointers are kept for raw base primitives (bool/int/uint/decimal/string, where
// an explicit zero must stay distinguishable from absent), enums that declare a
// zero member, numeric aliases, and zero-meaningful structs (Range, Position,
// all-optional capability/options containers) so a present-zero or present-{}
// value round-trips faithfully.
func (g *Generator) optionalDropsPointer(goType string) bool {
	if s := g.resolveStruct(goType); s != nil {
		return !g.zeroStructMeaningful(s, map[string]bool{})
	}
	if e, ok := g.enums[goType]; ok {
		return !enumHasZeroMember(e)
	}
	if a, ok := g.aliases[goType]; ok {
		return !g.zeroValueValid(a.Type, map[string]bool{})
	}
	return false
}

// zeroStructMeaningful reports whether a struct's zero value (all fields zero)
// is itself a legitimate present value: true when it has no required fields (an
// explicit empty {} differs from absent) or every required field has a valid
// zero (e.g. Range's required numeric Positions). seen breaks reference cycles.
func (g *Generator) zeroStructMeaningful(s *Structure, seen map[string]bool) bool {
	if seen[s.Name] {
		return true // cyclic reference: assume meaningful (keep the pointer)
	}
	seen[s.Name] = true
	defer delete(seen, s.Name)
	for _, p := range g.allProperties(s) {
		if p.Optional {
			continue
		}
		if !g.zeroValueValid(p.Type, seen) {
			return false
		}
	}
	return true
}

// zeroValueValid reports whether the zero value of t is a legitimate value that
// must remain distinguishable from an absent field. Numbers and booleans have a
// valid zero; strings/URIs do not (""). A struct is judged by
// [Generator.zeroStructMeaningful]; an enum by whether it declares a zero member.
func (g *Generator) zeroValueValid(t *Type, seen map[string]bool) bool {
	switch t.Kind {
	case KindBase:
		switch BaseTypeName(t.Name) {
		case BaseInteger, BaseUinteger, BaseDecimal, BaseBoolean:
			return true
		default: // string, URI, DocumentUri, RegExp, null
			return false
		}
	case KindReference:
		if s, ok := g.structures[t.Name]; ok {
			return g.zeroStructMeaningful(s, seen)
		}
		if e, ok := g.enums[t.Name]; ok {
			return enumHasZeroMember(e)
		}
		if a, ok := g.aliases[t.Name]; ok {
			return g.zeroValueValid(a.Type, seen)
		}
		return false
	default:
		// arrays/maps/tuples/unions/literals: the zero value is nil/empty, not a
		// meaningful present value for this purpose.
		return false
	}
}

// enumHasZeroMember reports whether an enumeration declares a member whose value
// is the zero value of its base type (0 for numeric, "" for string).
func enumHasZeroMember(e *Enumeration) bool {
	zero := "0"
	if e.Type.Name == BaseString {
		zero = `""`
	}
	for _, v := range e.Values {
		if strings.TrimSpace(string(v.Value)) == zero {
			return true
		}
	}
	return false
}

func requiredFieldNames(props []*Property) []string {
	var out []string
	for _, p := range props {
		if !p.Optional {
			out = append(out, p.Name)
		}
	}
	return out
}

func allFieldNames(props []*Property) []string {
	out := make([]string, len(props))
	for i, p := range props {
		out[i] = p.Name
	}
	return out
}

// structAllKeys returns the JSON names of every property of a struct, including
// those inherited via extends and mixins.
func (g *Generator) structAllKeys(s *Structure) []string {
	return allFieldNames(g.allProperties(s))
}

// allProperties returns a struct's own properties plus those merged from its
// extends and mixins, resolved transitively. Later definitions do not override
// earlier ones (first occurrence wins) to keep ordering stable.
func (g *Generator) allProperties(s *Structure) []*Property {
	seen := make(map[string]bool)
	var out []*Property
	add := func(props []*Property) {
		for _, p := range props {
			if !seen[p.Name] {
				seen[p.Name] = true
				out = append(out, p)
			}
		}
	}
	collect := func(refs []*Type) {
		for _, r := range refs {
			if r.Kind == KindReference {
				if parent, ok := g.structures[r.Name]; ok {
					add(g.allProperties(parent))
				}
			}
		}
	}
	add(s.Properties)
	collect(s.Extends)
	collect(s.Mixins)
	return out
}

// cleanIdent strips composite-type punctuation so a type expression can seed a
// Go identifier.
func cleanIdent(expr string) string {
	return exportName(strings.NewReplacer(
		"[]", "", "*", "", ".", "", "[", "", "]", "", " ", "", "map", "Map",
	).Replace(expr))
}

// namedWrapper synthesizes (or reuses) a named defined type whose underlying
// type is the given expression, so an otherwise-unnamed composite (slice, map)
// can carry the marker method a sealed-interface arm requires. Keyed by the
// underlying expression for deduplication.
func (g *Generator) namedWrapper(underlying, base string) string {
	if d, ok := g.arrayWrap[underlying]; ok {
		return d.Name
	}
	name := base
	for i := 2; g.nameTaken(name) || g.unionByName[name] != nil; i++ {
		name = fmt.Sprintf("%s%d", base, i)
	}
	d := &arrayWrapDecl{Name: name, Element: underlying}
	g.arrayWrap[underlying] = d
	g.arrayOrder = append(g.arrayOrder, underlying)
	return name
}

// arrayWrapper synthesizes a named slice type for an array union arm.
func (g *Generator) arrayWrapper(elem string) string {
	return g.namedWrapper("[]"+elem, cleanIdent(elem)+"Slice")
}

// mapWrapper synthesizes a named map type so a map can be a union arm.
func (g *Generator) mapWrapper(expr string) string {
	return g.namedWrapper(expr, cleanIdent(expr))
}

// tupleType lowers a tuple. Homogeneous tuples become a fixed-size array; in
// union context they are wrapped in a named type so they can carry a marker.
func (g *Generator) tupleType(t *Type, hint string, named bool) string {
	homogeneous := true
	var elem string
	for i, it := range t.Items {
		e := g.lower(it, hint)
		if i == 0 {
			elem = e
		} else if e != elem {
			homogeneous = false
		}
	}
	var expr string
	if homogeneous {
		expr = fmt.Sprintf("[%d]%s", len(t.Items), elem)
	} else {
		// Heterogeneous tuples are not present in the model; fall back to a
		// slice of raw values to remain total.
		expr = "[]jsontext.Value"
		g.warnf("heterogeneous tuple lowered to []jsontext.Value (hint %q)", hint)
	}
	if !named {
		return expr
	}
	if d, ok := g.tupleWrap[expr]; ok {
		return d.Name
	}
	name := g.uniqueName(exportName(hint) + "Tuple")
	d := &tupleWrapDecl{Name: name, Expr: expr}
	g.tupleWrap[expr] = d
	g.tupleOrder = append(g.tupleOrder, expr)
	return name
}

// andType synthesizes a struct embedding the operands of an "and" type.
func (g *Generator) andType(t *Type, hint string) string {
	operands := make([]string, 0, len(t.Items))
	for _, it := range t.Items {
		if it.Kind == KindReference {
			operands = append(operands, it.Name)
		} else {
			g.warnf("and operand is not a reference (hint %q)", hint)
		}
	}
	sig := strings.Join(operands, "&")
	if d, ok := g.ands[sig]; ok {
		return d.Name
	}
	name := g.uniqueName(strings.Join(operands, "And"))
	d := &andDecl{Name: name, Operands: operands}
	g.ands[sig] = d
	g.andOrder = append(g.andOrder, sig)
	return name
}

// literalType synthesizes a named struct for an inline object literal,
// deduplicated by structural signature.
func (g *Generator) literalType(lit *StructureLiteral, hint string) string {
	sig := literalSignature(lit)
	if d, ok := g.literals[sig]; ok {
		return d.Name
	}
	name := exportName(hint)
	if name == "" {
		name = "Literal"
	}
	// Avoid collisions with real declarations.
	orig := name
	for i := 2; g.nameTaken(name); i++ {
		name = fmt.Sprintf("%s%d", orig, i)
	}
	d := &literalDecl{Name: name, Lit: lit}
	g.literals[sig] = d
	g.literalOrder = append(g.literalOrder, sig)
	return name
}

func (g *Generator) nameTaken(name string) bool {
	if _, ok := g.structures[name]; ok {
		return true
	}
	if _, ok := g.enums[name]; ok {
		return true
	}
	if _, ok := g.aliases[name]; ok {
		return true
	}
	for _, d := range g.literals {
		if d.Name == name {
			return true
		}
	}
	return false
}

// reserved reports whether a Go identifier is already used by any declared or
// synthesized type, so that minted names can be made unique.
func (g *Generator) reserved(name string) bool {
	if g.nameTaken(name) || g.unionByName[name] != nil || g.aliasNames[name] {
		return true
	}
	for _, d := range g.arrayWrap {
		if d.Name == name {
			return true
		}
	}
	for _, d := range g.tupleWrap {
		if d.Name == name {
			return true
		}
	}
	for _, d := range g.ands {
		if d.Name == name {
			return true
		}
	}
	return false
}

// uniqueName returns base, or base with the smallest numeric suffix that is not
// reserved.
func (g *Generator) uniqueName(base string) string {
	name := base
	for i := 2; g.reserved(name); i++ {
		name = fmt.Sprintf("%s%d", base, i)
	}
	return name
}

func literalSignature(lit *StructureLiteral) string {
	parts := make([]string, len(lit.Properties))
	for i, p := range lit.Properties {
		opt := ""
		if p.Optional {
			opt = "?"
		}
		parts[i] = p.Name + opt + ":" + typeSignature(p.Type)
	}
	sort.Strings(parts)
	return "{" + strings.Join(parts, ",") + "}"
}

// typeSignature returns a stable structural signature for a meta type, used to
// deduplicate synthesized literals.
func typeSignature(t *Type) string {
	if t == nil {
		return "nil"
	}
	switch t.Kind {
	case KindBase, KindReference:
		return string(t.Kind) + ":" + t.Name
	case KindArray:
		return "[]" + typeSignature(t.Element)
	case KindMap:
		return "map[" + typeSignature(t.Key) + "]" + typeSignature(t.Value)
	case KindOr, KindAnd, KindTuple:
		parts := make([]string, len(t.Items))
		for i, it := range t.Items {
			parts[i] = typeSignature(it)
		}
		return string(t.Kind) + "(" + strings.Join(parts, ",") + ")"
	case KindLiteral:
		return "lit" + literalSignature(t.Literal)
	case KindStringLiteral:
		return "str:" + t.StringValue
	case KindIntegerLiteral:
		return fmt.Sprintf("int:%d", t.IntegerValue)
	case KindBooleanLiteral:
		return fmt.Sprintf("bool:%t", t.BooleanValue)
	default:
		return "?"
	}
}

func (g *Generator) warnf(format string, args ...any) {
	g.warnings = append(g.warnings, fmt.Sprintf(format, args...))
}
