// Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "go.lsp.dev/uri"

// URI scheme constants.
const (
	// FileScheme is the "file" URI scheme for filesystem paths.
	FileScheme = uri.FileScheme
	// HTTPScheme is the "http" URI scheme.
	HTTPScheme = uri.HTTPScheme
	// HTTPSScheme is the "https" URI scheme.
	HTTPSScheme = uri.HTTPSScheme
)

// Filename returns the filesystem path for a file URI.
//
// It panics if u is not a valid file URI; callers that cannot guarantee the
// scheme should parse u with [Parse] instead.
func (u URI) Filename() string {
	return uri.URI(u).Filename()
}

// New parses s into a go.lsp.dev/uri URI: a "file://" string is treated as
// already encoded, any other string is treated as a filesystem path and
// converted with [File]. Convert the result with URI(u) only when assigning a
// generated sealed-union URI arm such as [RelativePatternBaseURI].
func New(s string) uri.URI {
	return uri.New(s)
}

// File returns a go.lsp.dev/uri file URI for the filesystem path. A "$GOROOT"
// prefix is expanded to the current GOROOT, and relative paths are made
// absolute.
func File(path string) uri.URI {
	return uri.File(path)
}

// Parse parses s into a go.lsp.dev/uri URI, supporting the file, http, and https
// schemes. Unlike [New] it reports an error rather than panicking on an
// unsupported scheme.
func Parse(s string) (uri.URI, error) {
	return uri.Parse(s)
}

// From builds a go.lsp.dev/uri URI from its components for the file, http, and
// https schemes. It panics on an unknown scheme.
func From(scheme, authority, path, query, fragment string) uri.URI {
	return uri.From(scheme, authority, path, query, fragment)
}
