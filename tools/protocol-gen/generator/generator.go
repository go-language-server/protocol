// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package generator

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gobuffalo/flect"

	"go.lsp.dev/protocol/tools/protocol-gen/protocol"
)

var acronyms = [...]string{
	"LSP",
}

func init() {
	spew.Config = spew.ConfigState{
		Indent:           " ",
		ContinueOnMethod: true,
		SortKeys:         true,
	}

	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, acronym := range acronyms {
		buf.WriteString(strconv.Quote(acronym))
		if i < len(acronyms)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')
	if err := flect.LoadAcronyms(&buf); err != nil {
		panic(err)
	}
}

const (
	pkgContext = `"context"`
	pkgURI     = `"go.lsp.dev/uri"`
	pkgJSONRPC = `"go.lsp.dev/jsonrpc2"`
)

type genericsType struct {
	Name          string
	Documentation string
	Since         string
	Proposed      bool
}

type Generator struct {
	enumerations  []Printer
	typeAliases   []Printer
	structures    []Printer
	client        []Printer
	server        []Printer
	generics      map[string]bool
	genericsTypes map[genericsType][]protocol.Type
	files         map[string]*os.File
}

func (gen *Generator) Init() {
	gen.generics = make(map[string]bool)
	gen.genericsTypes = make(map[genericsType][]protocol.Type)
	gen.files = make(map[string]*os.File)
}

func (gen *Generator) writeTo(pp []Printer) (err error) {
	for _, p := range pp {
		filename := p.Filename()

		f, ok := gen.files[filename]
		if !ok {
			// creates file for writing target
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			root := filepath.Join(filepath.Dir(filepath.Dir(cwd)), "protocol")
			f, err = os.Create(filepath.Join(root, filename))
			if err != nil {
				return fmt.Errorf("unable to open %s file: %w", filename, err)
			}
			gen.files[filename] = f

			// Writes header content to the file at first
			f.WriteString(`// Copyright 2024 The Go Language Server Authors` + "\n")
			f.WriteString(`// SPDX-License-Identifier: BSD-3-Clause` + "\n")
			f.WriteString("\n")
			f.WriteString(`package protocol` + "\n")
			f.WriteString("\n")
			f.WriteString(`import (` + "\n")
			f.WriteString(`	` + pkgContext + "\n\n")
			f.WriteString(`	` + pkgURI + "\n")
			f.WriteString(`	` + pkgJSONRPC + "\n")
			f.WriteString(`)` + "\n\n")
		}

		f.Write(p.Bytes())
	}

	return nil
}

func (gen *Generator) WriteTo() (err error) {
	if err := gen.writeTo(gen.enumerations); err != nil {
		return err
	}
	if err := gen.writeTo(gen.structures); err != nil {
		return err
	}

	for _, f := range gen.files {
		if err := f.Close(); err != nil {
			return err
		}
	}

	return nil
}

type Printer interface {
	P(a ...any)
	PP(a ...any)
	Filename() string
	Bytes() []byte
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

func (p *printer) Filename() string {
	return p.filename
}

func (p *printer) Bytes() []byte {
	return p.buf.Bytes()
}

// WriteTo creates a p.filename file and writes the generated code.
func (p *printer) WriteTo() (err error) {
	// Creates file for writing target
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	root := filepath.Join(filepath.Dir(filepath.Dir(cwd)), "protocol")
	f, err := os.Create(filepath.Join(root, p.filename))
	if err != nil {
		return fmt.Errorf("unable to open %s file: %w", p.filename, err)
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()

	// Writes header content to the file at first
	p.p(f, `// Copyright 2024 The Go Language Server Authors`+"\n")
	p.p(f, `// SPDX-License-Identifier: BSD-3-Clause`+"\n")
	p.p(f, "\n")
	p.p(f, `package protocol`+"\n")
	p.p(f, "\n")
	p.p(f, `import (`+"\n")
	p.p(f, `        `+pkgContext+"\n\n")
	p.p(f, `        `+pkgURI+"\n\n")
	p.p(f, `        `+pkgJSONRPC+"\n\n")
	p.p(f, `)`+"\n\n")

	if _, err := f.Write(p.buf.Bytes()); err != nil {
		return fmt.Errorf("unable to write to %s file: %w", p.filename, err)
	}

	return nil
}

func normalizeMethodName(methName string) (methIdent string) {
	pairs := strings.Split(methName, "/")
	for _, s := range pairs {
		methIdent += flect.Pascalize(s)
	}

	return methIdent
}

func normalizeDocumentation(s string) string {
	if s == "." {
		return "."
	}

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

	if len(s) > 100 {
		var sb strings.Builder
		sc := bufio.NewScanner(strings.NewReader(s))
		sc.Split(bufio.ScanWords)
		for sc.Scan() {
			s := sc.Text() + " "
			if s != ". " && (sb.Len()%100 == 1 || sb.Len()%200 == 1 || sb.Len()%300 == 1) {
				s = "\n // " + s
			}
			sb.WriteString(s)
		}
		s = sb.String()
	}

	// s = strings.ReplaceAll(s, " . ", ". ")
	// s = strings.ReplaceAll(s, " , ", ", ")
	// s = strings.ReplaceAll(s, " '", "'")
	// s = strings.ReplaceAll(s, "' ", "'")
	// s = strings.ReplaceAll(s, " - ", "-")
	// s = strings.ReplaceAll(s, " x", "x")
	// s = strings.ReplaceAll(s, "( ", "(")
	// s = strings.ReplaceAll(s, " )", ")")
	// s = strings.ReplaceAll(s, "e. g", "e.g")

	return " " + s
}

func writeDocumentation(g Printer, typeName, docs, since string, proposed bool) {
	if docs != "" {
		g.PP(`// `, typeName, normalizeDocumentation(docs))
	}
	if since != "" && !strings.Contains(docs, "since") {
		if docs != "" {
			g.PP(`//`)
		}
		g.P(`// @since `, docs)
		if proposed {
			g.P(` proposed`)
		}
		g.P("\n")
	}
}

func normalizeLSPTypes(name string) string {
	switch name {
	case "LSPAny":
		name = "any"
	case "LSPObject":
		name = "map[string]any"
	case "LSPArray":
		name = "[]any"
	default:
		name = flect.Pascalize(name)
	}
	return name
}

func normalizeHasLSPTypes(name string) (string, bool) {
	isLSPType := false
	switch name {
	case "LSPAny":
		name = "any"
		isLSPType = true
	case "LSPObject":
		name = "map[string]any"
		isLSPType = true
	case "LSPArray":
		name = "[]any"
		isLSPType = true
	}
	return name, isLSPType
}

func isNull(tt ...protocol.Type) bool {
	for _, t := range tt {
		if _, ok := t.(*protocol.NullType); ok {
			return true
		}
	}
	return false
}
