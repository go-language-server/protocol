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

// TypeAliases generates Go type from the metaModel TypeAliases schema definition.
func (gen *Generator) TypeAliases(typeAliases []*protocol.TypeAlias) error {
	// Init typeAliases printers
	for _, name := range typeAliasNames {
		g := NewPrinter(name)
		gen.typeAliases = append(gen.typeAliases, g)

		for _, alias := range typeAliases {
			switch alias.Name {
			case "LSPAny", "LSPObject", "LSPArray":
				continue
			}

			if alias.Documentation != "" {
				g.PP(`// `, flect.Pascalize(alias.Name), normalizeDocumentation(alias.Documentation))
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

			g.P(`type `, flect.Pascalize(alias.Name))
			switch a := alias.Type.(type) {
			case *protocol.StringType:
				g.P(` `, `string`)

			case *protocol.ArrayType:
				elem := a.Element
				switch elem := elem.(type) {
				case *protocol.URIType:
					g.P(` `, `[]uri.URI`)
				case *protocol.DocumentUriType:
					g.P(` `, `[]`+flect.Pascalize("DocumentUri"))
				case *protocol.StringType:
					g.P(` `, `[]string`)
				case *protocol.IntegerType:
					g.P(` `, `[]int32`)
				case *protocol.UintegerType:
					g.P(` `, `[]uint32`)
				case *protocol.RegExpType:
					g.P(` `, `[]*regexp.Regexp`)
				case *protocol.BooleanType:
					g.P(` `, `[]bool`)
				case *protocol.ReferenceType:
					name := elem.Name
					if name == "LSPAny" {
						name = "any"
					}
					g.P(` `, `[]`+name)
				case *protocol.OrType:
					g.P(` `, `any /* or */`)
				}

			case *protocol.MapType:
				val := a.Value
				_ = val
				g.P(` `, `map[string]any`)

			case *protocol.OrType:
				g.PP(` interface {`)
				g.P(`	`)
				for i, item := range a.Items {
					switch item := item.(type) {
					default:
						fmt.Printf("%T\n", item)
					case *protocol.ArrayType:
						elem := item.Element
						switch elem := elem.(type) {
						case *protocol.URIType:
							g.P(`[]uri.URI`)
						case *protocol.DocumentUriType:
							g.P(`[]` + flect.Pascalize("DocumentUri"))
						case *protocol.StringType:
							g.P(`[]string`)
						case *protocol.IntegerType:
							g.P(`[]int32`)
						case *protocol.UintegerType:
							g.P(`[]uint32`)
						case *protocol.RegExpType:
							g.P(`[]*regexp.Regexp`)
						case *protocol.BooleanType:
							g.P(`[]bool`)
						case *protocol.ReferenceType:
							name := elem.Name
							if name == "LSPAny" {
								name = "any"
							}
							g.P(`[]` + name)
						case *protocol.OrType:
							g.P(`any /* or */`)
						}
					case *protocol.BooleanType:
						g.P(`bool`)
					case *protocol.DecimalType:
						g.P(`float64`)
					case *protocol.IntegerType:
						g.P(`int32`)
					case *protocol.ReferenceType:
						g.P(item.Name)
					case *protocol.StringType:
						g.P(`string`)
					case *protocol.UintegerType:
						g.P(`uint32`)
					}
					if i < len(a.Items)-1 {
						g.P(" | ")
					}
				}
				g.P("\n", `}`)

			case *protocol.ReferenceType:
				g.P(` `)
				name := a.Name
				if name == "LSPAny" {
					name = "any"
				}
				g.P(name)
			}
			g.P("\n\n")
		}

		if err := g.WriteTo(); err != nil {
			return err
		}
	}

	return nil
}
