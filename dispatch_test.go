// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"fmt"
	"strings"
	"testing"
)

// TestCommandOrCodeActionDispatch covers the textDocument/codeAction result
// union, where a CodeAction may carry an OPTIONAL `command` object whose key
// collides with Command's required `command` string. Required-key probing alone
// mis-routes; the decode-error gate must reject Command for an object command.
func TestCommandOrCodeActionDispatch(t *testing.T) {
	tests := map[string]struct {
		json     string
		wantType string // "*protocol.Command" or "*protocol.CodeAction"
	}{
		"success: plain command (command is a string)": {
			json:     `{"title":"x","command":"do.it"}`,
			wantType: "*protocol.Command",
		},
		"success: code action carrying a command object": {
			json:     `{"title":"Fix","command":{"title":"y","command":"do.it"}}`,
			wantType: "*protocol.CodeAction",
		},
		"success: code action with edit but no command": {
			json:     `{"title":"Fix","kind":"quickfix"}`,
			wantType: "*protocol.CodeAction",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var v CommandOrCodeAction
			if err := Unmarshal([]byte(tt.json), &v); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got := typeName(v); got != tt.wantType {
				t.Fatalf("got %s, want %s", got, tt.wantType)
			}
			out, err := Marshal(v)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if got, want := canon(t, out), canon(t, []byte(tt.json)); got != want {
				t.Errorf("round-trip mismatch:\n got=%s\nwant=%s", got, want)
			}
		})
	}

	// The actual response shape: an array mixing both arms.
	t.Run("success: []CommandOrCodeAction mixed", func(t *testing.T) {
		in := `[{"title":"a","command":"run"},{"title":"Fix","command":{"title":"y","command":"run"}}]`
		var v []CommandOrCodeAction
		if err := Unmarshal([]byte(in), &v); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(v) != 2 {
			t.Fatalf("len = %d, want 2", len(v))
		}
		if _, ok := v[0].(*Command); !ok {
			t.Errorf("v[0] = %T, want *Command", v[0])
		}
		if _, ok := v[1].(*CodeAction); !ok {
			t.Errorf("v[1] = %T, want *CodeAction", v[1])
		}
	})
}

// TestRegistrationOptionsSupersetNotLost covers Options|RegistrationOptions
// unions where the RegistrationOptions arm adds an optional `id`
// (StaticRegistrationOptions). The superset arm must win when `id` is present so
// it is not dropped.
func TestRegistrationOptionsSupersetNotLost(t *testing.T) {
	t.Run("success: notebook sync keeps id", func(t *testing.T) {
		in := `{"notebookSelector":[],"id":"reg-1"}`
		var v NotebookDocumentSync
		if err := Unmarshal([]byte(in), &v); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := v.(*NotebookDocumentSyncRegistrationOptions); !ok {
			t.Fatalf("got %T, want *NotebookDocumentSyncRegistrationOptions", v)
		}
		out, _ := Marshal(v)
		if !strings.Contains(string(out), `"id":"reg-1"`) {
			t.Errorf("id dropped: %s", out)
		}
	})

	t.Run("success: text document content keeps id", func(t *testing.T) {
		in := `{"schemes":["file"],"id":"reg-123"}`
		var v WorkspaceOptionsTextDocumentContent
		if err := Unmarshal([]byte(in), &v); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := v.(*TextDocumentContentRegistrationOptions); !ok {
			t.Fatalf("got %T, want *TextDocumentContentRegistrationOptions", v)
		}
		out, _ := Marshal(v)
		if !strings.Contains(string(out), `"id":"reg-123"`) {
			t.Errorf("id dropped: %s", out)
		}
	})
}

// TestDocumentFilterNotebookNotShadowed ensures a NotebookCellTextDocumentFilter
// is not captured by the flattened TextDocumentFilter language probe.
func TestDocumentFilterNotebookNotShadowed(t *testing.T) {
	in := `{"notebook":"jupyter-notebook","language":"python"}`
	var v DocumentFilter
	if err := Unmarshal([]byte(in), &v); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if _, ok := v.(*NotebookCellTextDocumentFilter); !ok {
		t.Fatalf("got %T, want *NotebookCellTextDocumentFilter", v)
	}
	out, _ := Marshal(v)
	if !strings.Contains(string(out), `"notebook"`) {
		t.Errorf("notebook dropped: %s", out)
	}
}

// TestTextDocumentFilterByRequiredKey ensures arms with identical full key sets
// but different required keys still dispatch correctly.
func TestTextDocumentFilterByRequiredKey(t *testing.T) {
	tests := map[string]struct {
		json     string
		wantType string
	}{
		"success: language-required arm": {`{"language":"go"}`, "*protocol.TextDocumentFilterLanguage"},
		"success: scheme-required arm":   {`{"scheme":"file"}`, "*protocol.TextDocumentFilterScheme"},
		"success: pattern-required arm":  {`{"pattern":"**/*.go"}`, "*protocol.TextDocumentFilterPattern"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var v TextDocumentFilter
			if err := Unmarshal([]byte(tt.json), &v); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got := typeName(v); got != tt.wantType {
				t.Errorf("got %s, want %s", got, tt.wantType)
			}
		})
	}
}

// TestSymbolArrayDistinguishedByData ensures a WorkspaceSymbol[] carrying the
// WorkspaceSymbol-only `data` field decodes to WorkspaceSymbolSlice (not
// SymbolInformationSlice) and does not lose `data`.
func TestSymbolArrayDistinguishedByData(t *testing.T) {
	in := `[{"name":"Foo","kind":5,"location":{"uri":"file:///x.go"},"data":{"k":"v"}}]`
	var v WorkspaceSymbolResult
	if err := Unmarshal([]byte(in), &v); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if _, ok := v.(WorkspaceSymbolSlice); !ok {
		t.Fatalf("got %T, want WorkspaceSymbolSlice", v)
	}
	out, _ := Marshal(v)
	if !strings.Contains(string(out), `"data"`) {
		t.Errorf("data dropped: %s", out)
	}
}

func typeName(v any) string { return fmt.Sprintf("%T", v) }
