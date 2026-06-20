// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// initialisms is the set of word segments that should be fully upper-cased in
// exported Go identifiers, following common Go style.
var initialisms = map[string]bool{
	"ID": true, "URI": true, "URL": true, "JSON": true, "HTTP": true,
	"HTTPS": true, "API": true, "UTF": true, "LSP": true, "EOL": true,
	"RPC": true, "TLS": true, "ASCII": true, "UUID": true, "CPU": true,
}

// splitWords breaks an identifier into word segments on case transitions and
// any non-alphanumeric separators (camelCase, snake_case, slash, dot, etc.).
func splitWords(s string) []string {
	var words []string
	var cur strings.Builder
	flush := func() {
		if cur.Len() > 0 {
			words = append(words, cur.String())
			cur.Reset()
		}
	}
	runes := []rune(s)
	for i, r := range runes {
		switch {
		case !unicode.IsLetter(r) && !unicode.IsDigit(r):
			flush()
			continue
		case i > 0 && unicode.IsUpper(r) && (unicode.IsLower(runes[i-1]) || unicode.IsDigit(runes[i-1])):
			// lower|digit -> Upper boundary (camelCase hump).
			flush()
		case i > 0 && unicode.IsUpper(r) && i+1 < len(runes) && unicode.IsLower(runes[i+1]) && unicode.IsUpper(runes[i-1]):
			// UPPER -> Upper+lower boundary (e.g. "JSONValue" -> "JSON","Value").
			flush()
		}
		cur.WriteRune(r)
	}
	flush()
	return words
}

// exportName converts an arbitrary meta-model name into an exported Go
// identifier, applying initialism rules per word segment.
func exportName(s string) string {
	words := splitWords(s)
	var b strings.Builder
	for _, w := range words {
		up := strings.ToUpper(w)
		switch {
		case initialisms[up]:
			b.WriteString(up)
		default:
			r := []rune(w)
			r[0] = unicode.ToUpper(r[0])
			b.WriteString(string(r))
		}
	}
	out := b.String()
	if out == "" {
		return "Empty"
	}
	// Leading digit is illegal in a Go identifier.
	if r, _ := utf8.DecodeRuneInString(out); unicode.IsDigit(r) {
		out = "N" + out
	}
	return out
}

// methodConstName returns the exported constant name for an LSP method string,
// e.g. "textDocument/implementation" -> "MethodTextDocumentImplementation".
func methodConstName(method string) string {
	return "Method" + exportName(method)
}
