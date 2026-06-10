// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build protocol_backend_probe

package protocol

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestThirdPartyJSONBackendSemanticGate(t *testing.T) {
	protocolDir, err := filepath.Abs(".")
	if err != nil {
		t.Fatalf("resolve protocol dir: %v", err)
	}
	jsonrpc2Dir := filepath.Clean(filepath.Join(protocolDir, "..", "jsonrpc2"))
	for _, backend := range []struct {
		name    string
		module  string
		version string
		import_ string
	}{
		{
			name:    "segmentio",
			module:  "github.com/segmentio/encoding",
			version: "v0.5.4",
			import_: "github.com/segmentio/encoding/json",
		},
		{
			name:    "sonic",
			module:  "github.com/bytedance/sonic",
			version: "v1.15.2",
			import_: "github.com/bytedance/sonic",
		},
		{
			name:    "goccy",
			module:  "github.com/goccy/go-json",
			version: "v0.10.6",
			import_: "github.com/goccy/go-json",
		},
	} {
		t.Run(backend.name, func(t *testing.T) {
			out := runThirdPartyBackendProbe(t, protocolDir, jsonrpc2Dir, backend.module, backend.version, backend.import_)
			t.Logf("%s probe output:\n%s", backend.name, out)
			if !strings.Contains(out, "semantic_ok=false") {
				t.Fatalf("%s unexpectedly passed the semantic gate; review whether it should be promoted:\n%s", backend.name, out)
			}
		})
	}
}

func runThirdPartyBackendProbe(t *testing.T, protocolDir, jsonrpc2Dir, module, version, importPath string) string {
	t.Helper()

	dir := t.TempDir()
	gomod := fmt.Sprintf(`module protocol-backend-probe

go 1.26

require go.lsp.dev/protocol v0.0.0

replace go.lsp.dev/protocol => %s
replace go.lsp.dev/jsonrpc2 => %s
`, protocolDir, jsonrpc2Dir)
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(gomod), 0o600); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
	mainSrc := fmt.Sprintf(`package main

import (
	"fmt"
	backendjson %q
	"go.lsp.dev/protocol"
)

func main() {
	data := []byte(`+"`"+`{"label":"x","documentation":"doc","detail":"d","kind":3}`+"`"+`)
	var got protocol.CompletionItem
	err := backendjson.Unmarshal(data, &got)
	out, marshalErr := backendjson.Marshal(got)
	var want protocol.CompletionItem
	protocolErr := protocol.Unmarshal(data, &want)
	wantOut, wantMarshalErr := protocol.Marshal(want)
	detail, detailOK := got.Detail.Get()
	_, docOK := got.Documentation.(protocol.String)
	semanticOK := err == nil && marshalErr == nil && protocolErr == nil && wantMarshalErr == nil &&
		got.Label == "x" && detailOK && detail == "d" && docOK && got.Kind == protocol.CompletionItemKindFunction &&
		string(out) == string(wantOut)
	fmt.Printf("err=%%T %%v\n", err, err)
	fmt.Printf("marshal_err=%%T %%v\n", marshalErr, marshalErr)
	fmt.Printf("protocol_err=%%T %%v want_marshal_err=%%T %%v\n", protocolErr, protocolErr, wantMarshalErr, wantMarshalErr)
	fmt.Printf("label=%%q detail=(%%q,%%v) doc_ok=%%v kind=%%d\n", got.Label, detail, detailOK, docOK, got.Kind)
	fmt.Printf("backend_json=%%s\n", out)
	fmt.Printf("protocol_json=%%s\n", wantOut)
	fmt.Printf("semantic_ok=%%v\n", semanticOK)
}
`, importPath)
	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(mainSrc), 0o600); err != nil {
		t.Fatalf("write main.go: %v", err)
	}
	run := func(args ...string) string {
		t.Helper()
		cmd := exec.Command("go", args...)
		cmd.Dir = dir
		cmd.Env = append(os.Environ(), "GOWORK=off")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("go %s failed: %v\n%s", strings.Join(args, " "), err, out)
		}
		return string(out)
	}
	run("get", module+"@"+version)
	run("mod", "tidy")
	return run("run", ".")
}
