// Copyright 2019 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
	"unicode"
)

// URI scheme constants.
const (
	// FileScheme is the "file" URI scheme for filesystem paths.
	FileScheme = "file"
	// HTTPScheme is the "http" URI scheme.
	HTTPScheme = "http"
	// HTTPSScheme is the "https" URI scheme.
	HTTPSScheme = "https"
)

const hierPart = "://"

// Filename returns the filesystem path for a file URI.
//
// It panics if u is not a valid file URI; callers that cannot guarantee the
// scheme should parse u with [Parse] instead.
func (u URI) Filename() string {
	fn, err := filenameFromURI(u)
	if err != nil {
		panic(err)
	}
	return filepath.FromSlash(fn)
}

func filenameFromURI(u URI) (string, error) {
	parsed, err := url.ParseRequestURI(string(u))
	if err != nil {
		return "", fmt.Errorf("parse request URI: %w", err)
	}
	if parsed.Scheme != FileScheme {
		return "", fmt.Errorf("only file URIs are supported, got %q", parsed.Scheme)
	}
	// url.Parse does not special-case Windows drive paths (golang.org/issue/6027);
	// a "/C:/..." path keeps a leading slash that must be trimmed.
	if isWindowsDriveURI(parsed.Path) {
		parsed.Path = parsed.Path[1:]
	}
	return parsed.Path, nil
}

// New parses s into a URI: a "file://" string is treated as already encoded, any
// other string is treated as a filesystem path and converted with [File].
func New(s string) URI {
	if unescaped, err := url.PathUnescape(s); err == nil {
		s = unescaped
	}
	if strings.HasPrefix(s, FileScheme+hierPart) {
		return URI(s)
	}
	return File(s)
}

// File returns a file URI for the filesystem path. A "$GOROOT" prefix is expanded
// to the current GOROOT, and relative paths are made absolute.
func File(path string) URI {
	if !isWindowsDrivePath(path) {
		if abs, err := filepath.Abs(path); err == nil {
			path = abs
		}
	}
	if isWindowsDrivePath(path) {
		path = "/" + path
	}
	u := url.URL{Scheme: FileScheme, Path: filepath.ToSlash(path)}
	return URI(u.String())
}

// Parse parses s into a URI, supporting the file, http, and https schemes. Unlike
// [New] it reports an error rather than panicking on an unsupported scheme.
func Parse(s string) (URI, error) {
	parsed, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("url.Parse: %w", err)
	}
	switch parsed.Scheme {
	case FileScheme:
		u := url.URL{Scheme: FileScheme, Path: parsed.Path, RawPath: filepath.FromSlash(parsed.Path)}
		return URI(u.String()), nil
	case HTTPScheme, HTTPSScheme:
		u := url.URL{
			Scheme:   parsed.Scheme,
			Host:     parsed.Host,
			Path:     parsed.Path,
			RawQuery: parsed.Query().Encode(),
			Fragment: parsed.Fragment,
		}
		return URI(u.String()), nil
	default:
		return "", errors.New("unknown scheme")
	}
}

// From builds a URI from its components for the file, http, and https schemes. It
// panics on an unknown scheme.
func From(scheme, authority, path, query, fragment string) URI {
	switch scheme {
	case FileScheme:
		u := url.URL{Scheme: FileScheme, Path: path, RawPath: filepath.FromSlash(path)}
		return URI(u.String())
	case HTTPScheme, HTTPSScheme:
		u := url.URL{
			Scheme:   scheme,
			Host:     authority,
			Path:     path,
			RawQuery: url.QueryEscape(query),
			Fragment: fragment,
		}
		return URI(u.String())
	default:
		panic(fmt.Sprintf("unknown scheme: %s", scheme))
	}
}

// isWindowsDrivePath reports whether path begins with a Windows drive letter
// followed by ":" (e.g. "C:").
func isWindowsDrivePath(path string) bool {
	if len(path) < 4 {
		return false
	}
	return unicode.IsLetter(rune(path[0])) && path[1] == ':'
}

// isWindowsDriveURI reports whether a URI path has a Windows drive prefix such as
// "/C:".
func isWindowsDriveURI(uri string) bool {
	if len(uri) < 4 {
		return false
	}
	return uri[0] == '/' && unicode.IsLetter(rune(uri[1])) && uri[2] == ':'
}
