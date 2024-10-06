// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package generator

import (
	"fmt"

	"github.com/gobuffalo/flect"

	"go.lsp.dev/protocol/tools/protocol-gen/protocol"
)

var typeAliasNames = [...]string{
	"typealias",
}

// TypeAliases generates TypeAliases Go type from the metaModel schema definition.
func (gen *Generator) TypeAliases(typeAliases []*protocol.TypeAlias) error {
	// Init typeAliases printers
	g := NewPrinter("typealias")
	gen.typeAliases = append(gen.typeAliases, g)

	for _, alias := range typeAliases {
		switch alias.Name {
		case "LSPAny", "LSPObject", "LSPArray":
			continue
		}

		aliasName := flect.Pascalize(alias.Name)
		gen.generics[aliasName] = true

		// write documentation
		if alias.Documentation != "" {
			g.PP(`// `, aliasName, normalizeDocumentation(alias.Documentation))
		}
		if alias.Since != "" {
			if alias.Documentation != "" {
				g.PP(`//`)
			}
			g.P(`// @since `, alias.Since)
			if alias.Proposed {
				g.P(` proposed`)
			}
			g.P("\n")
		}

		switch a := alias.Type.(type) {
		case protocol.BaseType:
			g.PP(`type `, aliasName, ` `, a.String())

		case *protocol.ReferenceType:
			g.PP(`type `, aliasName, ` `, normalizeLSPTypes(a.Name))

		case *protocol.ArrayType:
			g.P(`type `, aliasName, ` `)
			elem := a.Element
			switch elem := elem.(type) {
			case *protocol.ReferenceType:
				g.PP(`[]` + normalizeLSPTypes(elem.String()))
			default:
				panic(fmt.Sprintf("typealias: %T\n", elem))
			}

		case *protocol.MapType:
			g.P(`type `, aliasName, ` `, `map`)
			switch key := a.Key.(type) {
			case protocol.BaseType:
				g.P(`[`, key.String(), `]`)
			default:
				panic(fmt.Sprintf("structures.MapType.Key: %[1]T = %#[1]v\n", a.Key))
			}
			switch a.Value.(type) {
			case *protocol.OrType:
				g.PP(aliasName)
			default:
				panic(fmt.Sprintf("typealias.MapType.Value: %[1]T = %#[1]v\n", a.Value))
			}

		case *protocol.OrType:
			g.PP(`type `, aliasName, ` struct {`)
			g.PP(`	Value any `, "`json:\"value\"`")
			g.PP(`}`)
		}

		g.P("\n")

		seem := map[string]bool{}
		switch a := alias.Type.(type) {
		case *protocol.OrType:
			g.P(`func New`, aliasName)
			g.P(`[T `)
			for i, item := range a.Items {
				switch item := item.(type) {
				case protocol.BaseType:
					t := item.String()
					if !seem[t] {
						seem[t] = true
						g.P(t)
					}
				case *protocol.ReferenceType:
					t := normalizeLSPTypes(item.String())
					if !seem[t] {
						seem[t] = true
						g.P(t)
					}
				case *protocol.ArrayType:
					elem := item.Element
					switch elem := elem.(type) {
					case protocol.BaseType:
						t := `[]` + elem.String()
						if !seem[t] {
							seem[t] = true
							g.P(t)
						}
					case *protocol.ReferenceType:
						t := `[]` + normalizeLSPTypes(elem.String())
						if !seem[t] {
							seem[t] = true
							g.P(t)
						}
					default:
						panic(fmt.Sprintf("GenericsTypes.Array: %#v\n", elem))
					}
				default:
					panic(fmt.Sprintf("typealias.OrType: %#v\n", item))
				}

				if i < len(a.Items)-1 {
					g.P(` | `)
				}
			}
			g.PP(`](val T) `, aliasName, ` {`)
			g.PP(`	return `, aliasName, `{`)
			g.PP(`		Value: val,`)
			g.PP(`	}`)
			g.PP(`}`, "\n")
		}

		switch a := alias.Type.(type) {
		case *protocol.OrType:
			g.PP(`func (t `, aliasName, `) MarshalJSON() ([]byte, error) {`)
			g.PP(`	switch val := t.Value.(type) {`)
			for i, item := range a.Items {
				switch item := item.(type) {
				case protocol.BaseType:
					g.PP(`	case `, item.String(), `:`)
				case *protocol.ArrayType:
					elem := item.Element
					switch elem := elem.(type) {
					case *protocol.ReferenceType:
						g.PP(`	case `, `[]`+normalizeLSPTypes(elem.String()), `:`)
					default:
						panic(fmt.Sprintf("GenericsTypes.Array: %#v\n", elem))
					}
				case *protocol.ReferenceType:
					g.PP(`	case `, normalizeLSPTypes(item.String()), `:`)
				default:
					panic(fmt.Sprintf("typealias.OrType: %#v\n", item))
				}
				if i <= len(a.Items)-1 {
					g.PP(`		return marshal(val)`)
				}
			}
			g.PP(`	case nil:`)
			g.PP(`		return []byte("null"), nil`)
			g.PP(`	}`)
			g.PP(`	return nil, fmt.Errorf("unkonwn type: %T", t)`)
			g.PP(`}`)
		}

		switch a := alias.Type.(type) {
		case *protocol.OrType:
			g.PP(`func (t *`, aliasName, `) UnmarshalJSON(val []byte) error {`)
			g.PP(`if string(val) == "null" {`)
			g.PP(`	t.Value = nil`)
			g.PP(`	return nil`)
			g.PP(`}`)
			for i, item := range a.Items {
				g.P(`var h`, i, ` `)
				switch item := item.(type) {
				case protocol.BaseType:
					g.PP(item.String())
				case *protocol.ArrayType:
					elem := item.Element
					switch elem := elem.(type) {
					case *protocol.ReferenceType:
						g.PP(`[]` + normalizeLSPTypes(elem.String()))
					default:
						panic(fmt.Sprintf("GenericsTypes.Array: %#v\n", elem))
					}
				case *protocol.ReferenceType:
					g.PP(normalizeLSPTypes(item.String()))
				default:
					panic(fmt.Sprintf("typealias.OrType: %#v\n", item))
				}
				g.PP(`if err := unmarshal(val, &h`, i, `); err == nil {`)
				g.PP(`	t.Value = h`, i)
				g.PP(`	return nil`)
				g.PP(`}`)
			}
			g.PP(`	return &UnmarshalError{fmt.Sprintf("failed to unmarshal %T", t)}`)
			g.PP(`}`)
		}
	}

	if err := g.WriteTo(); err != nil {
		return err
	}

	return nil
}
