// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	stdjson "encoding/json"
	"fmt"
	"testing"
)

func TestWorkspaceSymbolResultGeneratedEncodersMatchCorpus(t *testing.T) {
	for _, name := range []string{"workspace_symbol_result", "workspace_symbol_result_info"} {
		t.Run(name, func(t *testing.T) {
			data := benchCorpus(t, name)
			var got WorkspaceSymbolResult
			if err := Unmarshal(data, &got); err != nil {
				t.Fatalf("decode %s: %v", name, err)
			}
			gotJSON, err := Marshal(got)
			if err != nil {
				t.Fatalf("marshal %s: %v", name, err)
			}
			if want := compactJSON(t, data); string(gotJSON) != want {
				t.Fatalf("workspace symbol wire mismatch\ngot:  %s\nwant: %s", gotJSON, want)
			}
		})
	}
}

func TestWorkspaceSymbolGeneratedEncodersPreserveOptionalFields(t *testing.T) {
	container := "pkg"
	workspaceSymbols := WorkspaceSymbolSlice{
		{
			BaseSymbolInformation: BaseSymbolInformation{Name: "WS", Kind: SymbolKindClass},
			Location:              &LocationUriOnly{URI: DocumentURI("file:///ws.go")},
			Data:                  LSPAny(`{"index":1}`),
		},
		{
			BaseSymbolInformation: BaseSymbolInformation{
				Name:          "Sym",
				Kind:          SymbolKindFunction,
				Tags:          []SymbolTag{SymbolTagDeprecated},
				ContainerName: &container,
			},
			Location: &Location{
				URI: DocumentURI("file:///x.go"),
				Range: Range{
					Start: Position{Line: 1, Character: 2},
					End:   Position{Line: 3, Character: 4},
				},
			},
		},
	}
	got, err := Marshal(workspaceSymbols)
	if err != nil {
		t.Fatalf("marshal workspace symbol slice: %v", err)
	}
	want := fmt.Sprintf(`[{"name":"WS","kind":%d,"location":{"uri":"file:///ws.go"},"data":{"index":1}},{"name":"Sym","kind":%d,"tags":[%d],"containerName":"pkg","location":{"uri":"file:///x.go","range":{"start":{"line":1,"character":2},"end":{"line":3,"character":4}}}}]`, SymbolKindClass, SymbolKindFunction, SymbolTagDeprecated)
	if string(got) != want {
		t.Fatalf("workspace symbol wire mismatch\ngot:  %s\nwant: %s", got, want)
	}
}

func compactJSON(t *testing.T, data []byte) string {
	t.Helper()
	var out bytes.Buffer
	if err := stdjson.Compact(&out, data); err != nil {
		t.Fatalf("compact json: %v", err)
	}
	return out.String()
}
