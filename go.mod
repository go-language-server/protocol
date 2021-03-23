module go.lsp.dev/protocol

go 1.15

require (
	github.com/francoispqt/gojay v1.2.13
	github.com/google/go-cmp v0.5.4
	go.lsp.dev/jsonrpc2 v0.9.0
	go.lsp.dev/pkg v0.0.0-20210323044036-f7deec69b52e
	go.lsp.dev/uri v0.3.0
	go.uber.org/zap v1.16.0
)

replace go.lsp.dev/jsonrpc2 => ../jsonrpc2
