// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command genprotocol generates the TypeScript codes from Microsoft/language-server-protocol/specification.md markdown.
package main

import (
	"fmt"
	"log"
	"strings"

	"gitlab.com/golang-commonmark/markdown"

	"github.com/go-language-server/protocol/internal/cmd/genprotocol/pkg/genprotocol"
)

func main() {
	spec, err := genprotocol.Getspecification(genprotocol.SpecificationMarkdownURI)
	if err != nil {
		log.Fatal(err)
	}
	md := markdown.New()
	tokens := md.Parse(spec)

	p := &genprotocol.Parser{}
	p.Parse(tokens)

	builder := new(strings.Builder)
	for _, fn := range p.TypeScriptFuncs {
		if fn.Src == "" {
			continue
		}
		fmt.Fprintf(builder, "%s\n%s\n", fn.Comment, fn.Src)
	}

	fmt.Println(strings.TrimSpace(builder.String()))
}
