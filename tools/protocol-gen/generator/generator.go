// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package generator

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"go.lsp.dev/protocol/tools/protocol-gen/schema"
)

func init() {
	spew.Config = spew.ConfigState{
		Indent:           " ",
		ContinueOnMethod: true,
		SortKeys:         true,
	}
}

const (
	pkgURI = `"go.lsp.dev/uri"`
)

type Generator struct {
	structures     []Printer
	enumerations   []Printer
	typeAliases    []Printer
	structureProps map[string]schema.Property
}

type Printer interface {
	P(a ...any)
	PP(a ...any)
	WriteTo() error
}

type printer struct {
	filename string
	buf      bytes.Buffer
}

func NewPrinter(filename string) Printer {
	return &printer{
		filename: filename + ".go",
	}
}

func (p *printer) p(w io.Writer, a ...any) {
	fmt.Fprint(w, a...)
}

// P prints a line to the generated output.
//
// It converts each parameter to a string following the same rules as [fmt.Print].
// It never inserts spaces between parameters.
func (p *printer) P(a ...any) {
	p.p(&p.buf, a...)
}

// PP prints a line to the generated output with newline.
//
// It converts each parameter to a string following the same rules as [fmt.Print].
// It never inserts spaces between parameters.
func (p *printer) PP(a ...any) {
	p.p(&p.buf, a...)
	p.p(&p.buf, "\n")
}

// WriteTo creates a p.filename file and writes the generated code.
func (p *printer) WriteTo() (err error) {
	// Creates file for writing target
	f, err := os.Create(filepath.Join(".protocol", p.filename))
	if err != nil {
		return fmt.Errorf("unable to open %s file: %w", p.filename, err)
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()

	// Writes header content to the file at first
	p.p(f, `package protocol`+"\n")
	p.p(f, "\n")
	p.p(f, `import (`+"\n")
	p.p(f, `	`+pkgURI+"\n")
	p.p(f, `)`+"\n\n")

	if _, err := f.Write(p.buf.Bytes()); err != nil {
		return fmt.Errorf("unable to write to %s file: %w", p.filename, err)
	}

	return nil
}

// transform returns a new slice by transforming each element with the function fn.
func transform[IN, OUT any](in []IN, fn func(in IN) OUT) []OUT {
	out := make([]OUT, len(in))
	for i, el := range in {
		out[i] = fn(el)
	}
	return out
}

func normalizeDocumentation(s string) string {
	if strings.HasPrefix(s, "@since") {
		return "."
	}

	s = strings.ReplaceAll(s, "\n", "\n// ")
	s = strings.ReplaceAll(s, "@since", "")
	s = strings.TrimRight(s, "1234567890.")

	if s[len(s)-1] != '.' {
		s = s + "."
	}
	s = strings.ReplaceAll(s, "  .", "")

	s = strings.ToLower(string(s[0])) + s[1:]

	return " " + s
}
