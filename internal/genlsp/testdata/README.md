# genlsp testdata

These files are the upstream LSP meta-model and its JSON Schema. The `genlsp`
generator lowers `metaModel.json` into the `go.lsp.dev/protocol/lsp` package.

## Provenance

| File                   | Source |
| ---------------------- | ------ |
| `metaModel.json`       | <https://raw.githubusercontent.com/microsoft/vscode-languageserver-node/release/protocol/3.18.0/protocol/metaModel.json> |
| `metaModel.schema.json`| <https://raw.githubusercontent.com/microsoft/vscode-languageserver-node/release/protocol/3.18.0/protocol/metaModel.schema.json> |

- Upstream repository: <https://github.com/microsoft/vscode-languageserver-node>
- Branch: `release/protocol/3.18.0`
- Retrieved: 2026-06-04

## Caveat: `metaData.version` is unreliable

`metaModel.json` reports `metaData.version` as **`3.17.0`** even though it is the
3.18.0 branch content (387 structures, post-3.17 additions such as
`InlineCompletion*`, `TextDocumentContent*`). Never key generation decisions on
`metaData.version`; treat the branch/commit above as the authoritative version.

## Refreshing

Re-download both files from the URLs above (pin to a tag or commit SHA when
reproducibility matters), then run `make generate` from the repository root.
