// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

// URI is a Uniform Resource Identifier as defined by RFC 3986 and used by the
// LSP base protocol. It is transported as a JSON string.
//
// Construct a URI with [New], [File], [Parse], or [From]; recover a filesystem
// path from a file URI with [URI.Filename].
type URI string

// DocumentURI identifies a text document by URI.
//
// The LSP meta-model names DocumentURI and [URI] separately, but both are file
// URIs carried as JSON strings and are used interchangeably, so DocumentURI is
// an alias of URI: a value produced by [File] or [New] satisfies either, exactly
// as the pre-3.18 package's uri.URI alias did.
type DocumentURI = URI
