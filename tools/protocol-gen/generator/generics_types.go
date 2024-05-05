// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package generator

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"go.lsp.dev/protocol/tools/protocol-gen/protocol"
)

// GenericsTypes generates Generics Go type from the metaModel schema definition.
func (gen *Generator) GenericsTypes() error {
	g := NewPrinter("types_generics")

	sorted := make([]genericsType, len(gen.genericsTypes))
	i := 0
	for generics := range gen.genericsTypes {
		sorted[i] = generics
		i++
	}
	slices.SortFunc(sorted, func(a, b genericsType) int {
		return cmp.Compare(a.Name, b.Name)
	})

	for _, generics := range sorted {
		types := gen.genericsTypes[generics]

		// write Documentation
		if generics.Documentation != "" {
			g.PP(`// `, generics.Name, normalizeDocumentation(generics.Documentation))
		}
		if generics.Since != "" {
			if generics.Documentation != "" {
				g.PP(`//`)
			}
			g.P(`// @since `, strings.ReplaceAll(generics.Since, "\n", " "))
			if generics.Proposed {
				g.P(` proposed`)
			}
			g.P("\n")
		}

		g.PP(`type `, generics.Name, ` struct {`)
		g.PP(`	Value any `, "`json:\"value\"`")
		g.PP(`}`)

		g.P("\n")

		g.P(`func New`, generics.Name)
		g.P(`[T `)
		var typs []string
		seem := map[string]bool{}
		for i, gt := range types {
			switch gt := gt.(type) {
			case protocol.BaseType:
				t := gt.String()
				if !seem[t] {
					seem[t] = true
					typs = append(typs, t)
					g.P(t)
				}

			case *protocol.NullType:
				// nothing to do

			case *protocol.ReferenceType:
				t := normalizeLSPTypes(strings.ReplaceAll(gt.String(), "Uri", "URI"))
				if !seem[t] {
					seem[t] = true
					typs = append(typs, t)
					g.P(t)
				}

			case *protocol.ArrayType:
				elem := gt.Element
				switch elem := elem.(type) {
				case protocol.BaseType:
					t := `[]` + elem.String()
					if !seem[t] {
						seem[t] = true
						typs = append(typs, t)
						g.P(t)
					}
				case *protocol.ReferenceType:
					t := `[]` + normalizeLSPTypes(elem.String())
					if !seem[t] {
						seem[t] = true
						typs = append(typs, t)
						g.P(t)
					}
				default:
					panic(fmt.Sprintf("GenericsTypes.Array: %#v\n", elem))
				}

			case *protocol.TupleType:
				for _, item := range gt.Items {
					switch item := item.(type) {
					case protocol.BaseType:
						name := item.String()
						if !seem[name] {
							seem[name] = true
							t := name
							typs = append(typs, t)
							g.P(t)
						}
					default:
						panic(fmt.Sprintf("GenericsTypes.TupleType: %[1]T = %#[1]v\n", item))
					}
				}

			default:
				panic(fmt.Sprintf("GenericsTypes: %[1]T = %#[1]v\n", gt))
			}

			if i < len(types)-1 && !isNull(types[i+1]) {
				g.P(` | `)
			}
		}
		g.PP(`](x T) `, generics.Name, ` {`)
		g.PP(`	return `, generics.Name, `{`)
		g.PP(`		Value: x,`)
		g.PP(`	}`)
		g.PP(`}`, "\n")

		g.PP(`func (t `, generics.Name, `) MarshalJSON() ([]byte, error) {`)
		g.PP(`	switch x := t.Value.(type) {`)
		for _, gt := range types {
			switch gt := gt.(type) {
			case protocol.BaseType:
				g.PP(`	case `, gt.String(), `:`)
				g.PP(`		return marshal(x)`)

			case *protocol.NullType:
				// nothing to do

			case *protocol.ReferenceType:
				g.PP(`	case `, normalizeLSPTypes(strings.ReplaceAll(gt.String(), "Uri", "URI")), `:`)
				g.PP(`		return marshal(x)`)

			case *protocol.ArrayType:
				elem := gt.Element
				switch elem := elem.(type) {
				case protocol.BaseType:
					g.PP(`	case `, `[]`+elem.String(), `:`)
				case *protocol.ReferenceType:
					g.PP(`	case `, `[]`+normalizeLSPTypes(elem.String()), `:`)
				default:
					panic(fmt.Sprintf("GenericsTypes.Array: %#v\n", elem))
				}
				g.PP(`		return marshal(x)`)

			case *protocol.TupleType:
				seem := map[string]bool{}
				for i, item := range gt.Items {
					switch item := item.(type) {
					case protocol.BaseType:
						name := item.String()
						if !seem[name] {
							g.PP(`	case `, name, `:`)
							seem[name] = true
							continue
						}
					default:
						panic(fmt.Sprintf("GenericsTypes.TupleType: %[1]T = %#[1]v\n", item))
					}
					if i < len(gt.Items) {
						g.PP(`		return marshal(x)`)
					}
				}

			default:
				panic(fmt.Sprintf("GenericsTypes: %[1]T = %#[1]v\n", gt))
			}
		}
		g.PP(`	case nil:`)
		g.PP(`		return []byte("null"), nil`)
		g.PP(`	}`)
		g.PP(`	return nil, fmt.Errorf("unknown type: %T", t)`)
		g.PP(`}`)

		g.PP(`func (t *`, generics.Name, `) UnmarshalJSON(x []byte) error {`)
		g.PP(`if string(x) == "null" {`)
		g.PP(`	t.Value = nil`)
		g.PP(`	return nil`)
		g.PP(`}`)

		for i, gt := range types {
			switch gt := gt.(type) {
			case protocol.BaseType:
				g.PP(`var h`, i, ` `, gt.String())
				g.PP(`if err := unmarshal(x, &h`, i, `); err == nil {`)
				g.PP(`	t.Value = h`, i)
				g.PP(`	return nil`)
				g.PP(`}`)

			case *protocol.NullType:
				// nothing to do

			case *protocol.ReferenceType:
				g.PP(`var h`, i, ` `, normalizeLSPTypes(strings.ReplaceAll(gt.String(), "Uri", "URI")))
				g.PP(`if err := unmarshal(x, &h`, i, `); err == nil {`)
				g.PP(`	t.Value = h`, i)
				g.PP(`	return nil`)
				g.PP(`}`)

			case *protocol.ArrayType:
				g.P(`var h`, i, ` `)
				elem := gt.Element
				switch elem := elem.(type) {
				case protocol.BaseType:
					g.PP(`[]` + elem.String())
				case *protocol.ReferenceType:
					g.PP(`[]` + normalizeLSPTypes(elem.String()))
				default:
					panic(fmt.Sprintf("GenericsTypes.Array: %#v\n", elem))
				}
				g.PP(`if err := unmarshal(x, &h`, i, `); err == nil {`)
				g.PP(`	t.Value = h`, i)
				g.PP(`	return nil`)
				g.PP(`}`)

			case *protocol.TupleType:
				seem := map[string]bool{}
				for j, item := range gt.Items {
					switch item := item.(type) {
					case protocol.BaseType:
						name := item.String()
						if !seem[name] {
							g.PP(`var h`, i+j, ` `, name)
							seem[name] = true
						}
					default:
						panic(fmt.Sprintf("GenericsTypes.TupleType: %[1]T = %#[1]v\n", item))
					}
					if j < len(gt.Items)-1 {
						g.PP(`if err := unmarshal(x, &h`, i+j, `); err == nil {`)
						g.PP(`	t.Value = h`, i+j)
						g.PP(`	return nil`)
						g.PP(`}`)
					}
				}

			default:
				panic(fmt.Sprintf("GenericsTypes: %[1]T = %#[1]v\n", gt))
			}
		}
		g.P(`	return &UnmarshalError{"unmarshal failed to match one of [`)
		for i, t := range typs {
			g.P(t)
			if i < len(typs)-1 {
				g.P(` `)
			}
		}
		g.PP(`]"}`)
		g.PP(`}`)
	}

	if err := g.WriteTo(); err != nil {
		return err
	}

	return nil
}
