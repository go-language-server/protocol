// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package generator

import (
	"github.com/gobuffalo/flect"

	"go.lsp.dev/protocol/tools/protocol-gen/protocol"
)

var enumerationNames = [...]string{
	"enumerations",
}

// Enumerations generates Go type from the metaModel Enumerations schema definition.
func (gen *Generator) Enumerations(enumerations []*protocol.Enumeration) error {
	// Init enumerations printers
	for _, name := range enumerationNames {
		g := NewPrinter(name)
		gen.enumerations = append(gen.enumerations, g)

		for _, enum := range enumerations {
			if enum.Documentation != "" {
				g.PP(`// `, flect.Pascalize(enum.Name), normalizeDocumentation(enum.Documentation))
			}
			if enum.Since != "" {
				if enum.Documentation != "" {
					g.PP(`//`)
				}
				g.P(`// @since `, enum.Since)
				if enum.Proposed {
					g.P(` proposed`)
				}
				g.P("\n")
			}

			g.P(`type `, flect.Pascalize(enum.Name))
			switch e := enum.Type.(type) {
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
				name := string(e.Name)
				if name == "LSPAny" {
					name = "any"
				}
				g.P(name)
			}
			g.P("\n")

			g.PP(`const (`)
			for _, val := range enum.Values {
				if val.Documentation != "" {
					g.PP(`// `, flect.Pascalize(val.Name), flect.Pascalize(enum.Name), normalizeDocumentation(val.Documentation))
				}
				if val.Since != "" {
					if val.Documentation != "" {
						g.PP(`//`)
					}
					g.P(`// @since `, val.Since)
					if val.Proposed {
						g.P(` proposed`)
					}
					g.P("\n")
				}
				g.PP(`	`, flect.Pascalize(val.Name), flect.Pascalize(enum.Name), ` `, flect.Pascalize(enum.Name), ` = `, val.Value)
				g.P("\n")
			}
			g.PP(`)`)
		}

		if err := g.WriteTo(); err != nil {
			return err
		}
	}

	return nil
}
