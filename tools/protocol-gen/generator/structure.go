// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package generator

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/flect"

	"go.lsp.dev/protocol/tools/protocol-gen/protocol"
)

var structureNames = [...]string{
	"structures",
}

// Structures generates Go type from the metaModel Structure schema definition.
func (gen *Generator) Structures(structures []*protocol.Structure) error {
	// Init structures printers
	for _, name := range structureNames {
		// NOTE(zchee): Special case for underlying type of InitializeParams
		if strings.EqualFold(name, "_InitializeParams") {
			continue
		}
		g := NewPrinter(name)
		gen.structures = append(gen.structures, g)

		for _, s := range structures {
			if s.Documentation != "" {
				g.PP(`// `, flect.Pascalize(s.Name), normalizeDocumentation(s.Documentation))
			}
			if s.Since != "" {
				if s.Documentation != "" {
					g.PP(`//`)
				}
				g.P(`// @since `, s.Since)
				if s.Proposed {
					g.P(` proposed`)
				}
				g.P("\n")
			}

			g.PP(`type `, flect.Pascalize(s.Name), ` struct {`)
			var needNewline bool

			// NOTE(zchee): Special case for InitializeParams extends `WorkspaceFoldersInitializeParams`
			if strings.EqualFold(name, "InitializeParams") {
				s.Extends = append(s.Extends, &protocol.ReferenceType{
					Name: "WorkspaceFoldersInitializeParams",
				})
			}
			if len(s.Extends) > 0 {
				g.PP(`	// extends`)
				for i, extend := range s.Extends {
					switch extend := extend.(type) {
					case *protocol.ReferenceType:
						g.PP(`	`, extend.Name)
					default:
						fmt.Printf("mixin[%d]: %#[2]v %[2]T\n", i, extend)
					}
					if ns := extend.SubTypes(); ns != nil {
						for i, n := range ns {
							switch n := n.(type) {
							default:
								fmt.Printf("extend[%d]: %#[2]v %[2]T\n", i, n)
							}
							g.PP(`	`, n)
						}
					}
				}
				needNewline = true
			}

			if len(s.Mixins) > 0 {
				g.PP(`	// mixins`)
				for i, mixin := range s.Mixins {
					switch mixin := mixin.(type) {
					case *protocol.ReferenceType:
						g.PP(`	`, mixin.Name)
					default:
						fmt.Printf("mixin[%d]: %#[2]v %[2]T\n", i, mixin)
					}
					if ns := mixin.SubTypes(); ns != nil {
						for i, n := range ns {
							switch n := n.(type) {
							default:
								fmt.Printf("mixin.SubTypes[%d]: %#[2]v %[2]T\n", i, n)
							}
							g.PP(`	`, n)
						}
					}
				}
				needNewline = true
			}

			if len(s.Properties) > 0 {
				if needNewline {
					g.PP(``)
				}
				for i, prop := range s.Properties {
					if prop.Documentation != "" {
						g.PP(`	// `, flect.Pascalize(prop.Name), normalizeDocumentation(prop.Documentation))
					}
					if s.Since != "" {
						if prop.Documentation != "" && !strings.Contains(prop.Documentation, "since") {
							g.PP(`	//`)
						}
						g.P(`	// @since `, s.Since)
						if s.Proposed {
							g.P(` proposed`)
						}
						g.P("\n")
					}
					g.P(`	`, flect.Pascalize(prop.Name))

					propType := prop.Type
					switch node := propType.(type) {
					case *protocol.URIType:
						g.P(` `, `uri.URI`)
					case *protocol.DocumentUriType:
						g.P(` `, flect.Pascalize("DocumentUri"))
					case *protocol.IntegerType:
						g.P(` `, "int32")
					case *protocol.UintegerType:
						g.P(` `, "uint32")
					case *protocol.DecimalType:
						g.P(` `, "float64") // TODO(zchee): use [image.Color]?
					case *protocol.RegExpType:
						g.P(` `, "*regexp.Regexp")
					case *protocol.StringType:
						g.P(` `, "string")
					case *protocol.BooleanType:
						g.P(` `, "bool")

					case *protocol.ReferenceType:
						g.P(` `)
						name := string(node.Name)
						if name == "LSPAny" {
							name = "any"
						}
						if prop.Optional && name != "any" {
							g.P(`*`)
						}
						g.P(name)

					case *protocol.ArrayType:
						elem := node.Element
						switch elem := elem.(type) {
						case *protocol.URIType:
							g.P(` `, `[]uri.URI`)
						case *protocol.DocumentUriType:
							g.P(` `, `[]`+flect.Pascalize("DocumentUri"))
						case *protocol.StringType:
							g.P(` `, `[]string`)
						case *protocol.IntegerType:
							g.P(` `, `[]`+"int32")
						case *protocol.UintegerType:
							g.P(` `, `[]`+"uint32")
						case *protocol.RegExpType:
							g.P(` `, `[]`+"*regexp.Regexp")
						case *protocol.BooleanType:
							g.P(` `, `[]`+"bool")

						case *protocol.ReferenceType:
							name := elem.Name
							if name == "LSPAny" {
								name = "any"
							}
							g.P(` `, `[]`+name)

						case *protocol.OrType:
							g.P(` `, `any /* or */`)

						default:
							panic(fmt.Sprintf("ArrayKind: %#v\n", elem))
						}

					case *protocol.MapType:
						val := node.Value
						_ = val
						// fmt.Printf("MapType: val: %#[1]v %[1]T\n", val)
						g.P(` `, `map[string]any`)
						// switch val := val.(type) {
						// case *protocol.MapType:
						// 	// TODO(zchee): implements correctly
						// 	fmt.Printf("MapType: val: %#v\n", val)
						// 	g.P(` `, `map[string]any`)
						// 	// key := propType["key"].(map[string]any)
						// 	// switch schema.TypeKind(key["kind"].(string)) {
						// 	// case schema.BaseKind:
						// 	// 	g.P(` `, `map[`, key["name"].(string), `]`)
						// 	// case schema.ReferenceKind:
						// 	// 	g.P(` `, `map[`, key["name"].(string), `]`)
						// 	// }
						// 	//
						// 	// value := propType["value"].(map[string]any)
						// 	// switch schema.TypeKind(value["kind"].(string)) {
						// 	// case schema.BaseKind:
						// 	// 	g.P(value["name"])
						// 	// case schema.ReferenceKind:
						// 	// 	g.P(value["name"])
						// 	// default:
						// 	// 	fmt.Printf("ArrayKind: %#v\n", value)
						// 	// 	// panic(fmt.Sprintf("ArrayKind: %#v\n", elem))
						// 	// }
						// default:
						// 	panic(fmt.Sprintf("MapKind: %#v\n", propType))
						// }

					case *protocol.AndType, *protocol.TupleType:
						// AndKind and TupleKind are nothing to do on structures

					case *protocol.OrType:
						// tt := schema.ToOrItems(propType["items"])
						// _ = tt
						g.P(` `, `any /* or */`)

					case *protocol.StructureLiteralType:
						lit := node.Value
						_ = lit
						fmt.Printf("StructureLiteralType: val: %#v\n", lit)
						g.P(` `)
						if prop.Optional {
							g.P(`*`)
						}
						g.P(flect.Pascalize(s.Name)+flect.Pascalize(prop.Name), `/* LiteralKind */`)

					case *protocol.StringLiteralType:
						g.P(` `, node.Value)

					case *protocol.IntegerLiteralType:
						g.P(` `, node.Value)

					case *protocol.BooleanLiteralType:
						g.P(` `, node.Value)

					default:
						panic(fmt.Sprintf("prop: %#v\n", propType))
					}

					g.P(" `json:\"")
					g.P(prop.Name)
					if prop.Optional {
						g.P(`,omitempty`)
					}
					g.PP("\"`")

					// Add newline per fields
					if i <= len(s.Properties)-1 {
						g.P("\n")
					}
				}
			}
			g.PP(`}`)
			g.PP(``)
		}

		if err := g.WriteTo(); err != nil {
			return err
		}
	}

	return nil
}
