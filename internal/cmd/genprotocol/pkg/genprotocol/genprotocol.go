// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package genprotocol generates the TypeScript codes from Microsoft/language-server-protocol/specification.md markdown.
//
// The "gitlab.com/golang-commonmark/markdown" AST are:
//
//  *markdown.HeadingOpen
//    *markdown.Inline
//  *markdown.HeadingClose
//
//  *markdown.ParagraphOpen
//    *markdown.Inline
//  *markdown.ParagraphClose
//
//  *markdown.BulletListOpen
//    *markdown.ListItemOpen
//      *markdown.ParagraphOpen
//        *markdown.Inline
//      *markdown.ParagraphClose
//    *markdown.ListItemClose
//  *markdown.BulletListClose
//
//  *markdown.BlockquoteOpen
//    *markdown.ParagraphOpen
//      *markdown.Inline
//    *markdown.ParagraphClose
//  *markdown.BlockquoteClose
//
//  *markdown.Fence
//  *markdown.CodeBlock
//
//  *markdown.OrderedListOpen
//    *markdown.ListItemOpen
//      *markdown.ParagraphOpen
//        *markdown.Inline
//      *markdown.ParagraphClose
//    *markdown.ListItemClose
//  *markdown.OrderedListClose
//
//  *markdown.TableOpen
//    *markdown.TheadOpen
//      *markdown.TrOpen
//        *markdown.ThOpen
//          *markdown.Inline
//        *markdown.ThClose
//      *markdown.TrClose
//    *markdown.TheadClose
//
//    *markdown.TbodyOpen
//      *markdown.TrOpen
//        *markdown.TdOpen
//          *markdown.ThClose
//            *markdown.Inline
//          *markdown.TdClose
//        *markdown.TrClose
//      *markdown.TbodyClose
//  *markdown.TableClose
//
//  *markdown.Hr
package genprotocol

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gitlab.com/golang-commonmark/markdown"
	"golang.org/x/net/html"
)

const (
	// SpecificationMarkdownURI is the Microsoft/language-server-protocol/specification.md markdown uri.
	SpecificationMarkdownURI = "https://github.com/Microsoft/language-server-protocol/raw/gh-pages/specification.md"
)

// IgnoreHeaders ignores header comment map.
var IgnoreHeaders = map[string]bool{
	"$ Notifications and Requests": true,
	"Goto Declaration Request":     true,
	"Goto Definition Request":      true,
	"Goto Type Definition Request": true,
	"Goto Implementation Request":  true,
}

// InlineReplacer replacer of inline code block.
var InlineReplacer = strings.NewReplacer(
	"`", "",
	"\n", " ",
)

// CommonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
//
// This code copied from golang.org/x/lint/lint.go
//  https://github.com/golang/lint/blob/5614ed5bae6fb75893070bdc0996a68765fdd275/lint.go#L768-L810
var CommonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

// IgnoreFences map of ignore Fences codes map.
var IgnoreFences = map[string]bool{
	`{
    start: { line: 5, character: 23 },
    end : { line 6, character : 0 }
}
`: true,
	`{ language: 'typescript', scheme: 'file' }
{ language: 'json', pattern: '**/package.json' }
`: true,
}

// GetSpec gets the specification.md.
func Getspecification(uri string) ([]byte, error) {
	hc := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return data, nil
}

// TypeScriptFunc represents a typescript function.
type TypeScriptFunc struct {
	Comment string
	Src     string
}

// Parser represents a specification markdown parser.
type Parser struct {
	Tokens    []markdown.Token
	TokenPos  int
	PrevToken markdown.Token

	TypeScriptFuncs   []*TypeScriptFunc
	TypeScriptFuncPos int
}

// Parse parses the specification.md.
func (p *Parser) Parse(tokens []markdown.Token) {
	p.Tokens = tokens
	p.TypeScriptFuncs = make([]*TypeScriptFunc, 0, len(tokens))
	p.TypeScriptFuncPos = 0

	for p.TokenPos = 0; p.TokenPos < len(tokens); p.TokenPos++ {
		switch tok := tokens[p.TokenPos].(type) {
		case *markdown.HeadingOpen:
			p.ParseHeading(tok)
		case *markdown.Inline:
			p.ParseInline(tok)
		case *markdown.ParagraphOpen:
			p.ParseParagraph(tok)
		case *markdown.Fence:
			p.ParseFence(tok)
		}
	}
}

// HeaderLevel needs headers comment heading level.
const HeaderLevel = 3

// ParseHeading parses the markdown.Heading.
func (p *Parser) ParseHeading(tok *markdown.HeadingOpen) {
	if tok.HLevel < HeaderLevel {
		p.TokenPos++ // skip next Inline
		p.TokenPos++ // skip next HeadingClose
	}
	p.PrevToken = &markdown.HeadingOpen{}
}

// ParseInline parses inline markdown text.
func (p *Parser) ParseInline(tok *markdown.Inline) {
	switch p.PrevToken.(type) {
	case *markdown.HeadingOpen:
		if comment := tok.Content; comment != "" {
			if comment[0] == '<' { // trim HTML text
				comment = ExpandHTML(comment)
			}

			if IgnoreHeaders[comment] {
				return
			}

			if idx := strings.Index(comment, " "); idx > -1 {
				comment = comment[:idx] + comment[idx+1:]
			}

			tsfunc := &TypeScriptFunc{
				Comment: "// " + strings.TrimSpace(comment),
			}
			p.TypeScriptFuncs = append(p.TypeScriptFuncs, tsfunc)
			p.PrevToken = &markdown.HeadingClose{}
		}

	case *markdown.ParagraphOpen:
		if comment := tok.Content; comment != "" {
			if comment[0] == '<' { // trim HTML text
				comment = ExpandHTML(comment)
			}

			if IgnoreHeaders[comment] {
				return
			}

			if strings.ContainsRune(comment, '`') {
				comment = InlineReplacer.Replace(comment)
			}

			if len(p.TypeScriptFuncs)-1 < p.TypeScriptFuncPos {
				p.PrevToken = &markdown.ParagraphClose{}
				return
			}

			if idx := strings.Index(comment, ","); idx > -1 {
				if !CommonInitialisms[comment[:idx]] {
					comment = ToLowerCamelCase(comment[:idx])
				}
			} else {
				comment = ToLowerCamelCase(comment)
			}

			p.TypeScriptFuncs[p.TypeScriptFuncPos].Comment += " " + strings.TrimSpace(comment)
			p.TypeScriptFuncs[p.TypeScriptFuncPos].Comment = p.TypeScriptFuncs[p.TypeScriptFuncPos].Comment

			p.PrevToken = &markdown.ParagraphClose{}
		}
	default:
		p.PrevToken = &markdown.Inline{}
	}
}

// ParseParagraph parses markdown Paragraph text.
func (p *Parser) ParseParagraph(tok *markdown.ParagraphOpen) {
	switch p.PrevToken.(type) {
	case *markdown.HeadingOpen, *markdown.Inline:
		p.PrevToken = &markdown.ParagraphOpen{}
	default:
		// noting to do
	}
}

// ParseFence parses markdown Fence codes.
func (p *Parser) ParseFence(tok *markdown.Fence) {
	if tok.Params != "typescript" {
		return
	}

	if IgnoreFences[tok.Content] {
		return
	}

	if len(p.TypeScriptFuncs) == p.TypeScriptFuncPos {
		tsfunc := &TypeScriptFunc{
			Src: strings.TrimSpace(tok.Content),
		}
		p.TypeScriptFuncs = append(p.TypeScriptFuncs, tsfunc)
		p.TypeScriptFuncPos++
		return
	}
	p.TypeScriptFuncs[p.TypeScriptFuncPos].Src = tok.Content
	p.TypeScriptFuncPos++

	p.PrevToken = &markdown.Fence{}
}

// ExpandHTML expands the HTML links from s.
func ExpandHTML(s string) string {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			b = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b == nil {
		return s
	}

	s = b.LastChild.Data

	// trim `↩` emoji
	if idx := strings.Index(s, " (:leftwards_arrow_with_hook:)"); idx > -1 {
		s = s[:idx]
	}

	return s
}

// ToLowerCamelCase converts s to LowerCamelCase.
func ToLowerCamelCase(s string) string {
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = string(r^32) + s[1:]
	}

	return s
}

func (p *Parser) ParseComment(src string) {
	var comment string
	var i int
	if src[0] == '/' {
	parent:
		for j := 1; ; j++ {
			s := src[j]
			switch s {
			case '*':
				continue
			case '/':
				i = j + 1
				break parent
			default:
				comment += string(s)
			}
		}
	}
	_ = i
}

func SplitComment(comment string) string {
	if len(comment) > 120 {
		cs := strings.SplitN(comment, " ", len(comment)/120)

		for _, s := range cs {
			idx := strings.LastIndex(s, " ")
			if idx == -1 {
				break
			}
			comment = comment[idx+1:] + "\n// " + comment[:idx+2]
		}
	}

	return comment
}
