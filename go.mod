module go.lsp.dev/protocol

go 1.18

replace go.lsp.dev/jsonrpc2 => go.lsp.dev/jsonrpc2 v0.10.1-0.20220407150525-abf64a1b10d1

replace github.com/bytedance/sonic => github.com/zchee/sonic v0.0.0-20220407142055-d3441945585f

require (
	github.com/bytedance/sonic v1.2.0
	github.com/google/go-cmp v0.5.6
	go.lsp.dev/jsonrpc2 v0.10.0
	go.lsp.dev/pkg v0.0.0-20210717090340-384b27a52fb2
	go.lsp.dev/uri v0.3.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/chenzhuoyu/base64x v0.0.0-20211019084208-fb5309c8db06 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
