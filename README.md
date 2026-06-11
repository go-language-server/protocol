# protocol

[![test][test-badge]][test]
[![pkg.go.dev][pkg.go.dev-badge]][pkg.go.dev]
[![Go module][module-badge]][module]
[![codecov.io][codecov-badge]][codecov]

Package protocol implements Language Server Protocol specification in Go.

## URI types

Generated URI and URI fields use `go.lsp.dev/uri.URI` directly. Use
`go.lsp.dev/uri` constructors such as `uri.Parse`, `uri.File`, and
`uri.From` for new URI values.

`protocol.URI` remains as a package-local named type for compatibility and for
sealed union arms that require a local marker-method receiver. In ordinary code,
prefer `go.lsp.dev/uri.URI`; when assigning a URI string arm to
`RelativePatternBaseURI`, convert explicitly with `protocol.URI(u)`.

## Decode ownership and zero-copy strings

Generated decoders copy typed JSON input once before walking it. The caller may
reuse or mutate the original `[]byte` after `protocol.Unmarshal` or the LSP
codec returns. To reduce allocation cost on hot paths, unescaped decoded strings
and raw JSON value fields may alias that owned per-message copy. Retaining a
small decoded string or raw value can therefore keep the whole JSON message copy
live. Use `protocol.Clone` when a decoded protocol value must be detached from a
much larger input message before long-term retention.


<!-- badge links -->
[test]: https://github.com/go-language-server/protocol/actions/workflows/test.yaml
[pkg.go.dev]: https://pkg.go.dev/go.lsp.dev/protocol
[module]: https://github.com/go-language-server/protocol/releases/latest
[codecov]: https://app.codecov.io/gh/go-language-server/protocol

[test-badge]: https://img.shields.io/github/actions/workflow/status/go-language-server/protocol/test.yaml?branch=main&style=for-the-badge&label=TEST&logo=github
[pkg.go.dev-badge]: https://img.shields.io/badge/pkg.go.dev-doc-00add8?style=for-the-badge&logo=go
[module-badge]: https://img.shields.io/github/release/go-language-server/protocol.svg?color=00add8&label=MODULE&style=for-the-badge&logo=go
[codecov-badge]: https://img.shields.io/codecov/c/github/go-language-server/protocol/main?logo=codecov&style=for-the-badge
