// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"runtime"
	"strings"
	"testing"
	"unsafe"

	gocmp "github.com/google/go-cmp/cmp"
)

func TestUnmarshalOwnsInputForAliasedStrings(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		data       []byte
		mutate     string
		wantLabel  string
		decodeFunc func(*testing.T, []byte) string
	}{
		"success: byte walker root": {
			data:      []byte(`{"label":"owned"}`),
			mutate:    "owned",
			wantLabel: "owned",
			decodeFunc: func(t *testing.T, data []byte) string {
				t.Helper()

				var got CompletionItem
				if err := Unmarshal(data, &got); err != nil {
					t.Fatalf("Unmarshal CompletionItem: %v", err)
				}
				return got.Label
			},
		},
		"success: union root byte walker arm": {
			data:      []byte(`[{"label":"owned"}]`),
			mutate:    "owned",
			wantLabel: "owned",
			decodeFunc: func(t *testing.T, data []byte) string {
				t.Helper()

				var got CompletionResult
				if err := Unmarshal(data, &got); err != nil {
					t.Fatalf("Unmarshal CompletionResult: %v", err)
				}
				items, ok := got.(CompletionItemSlice)
				if !ok {
					t.Fatalf("CompletionResult = %T, want CompletionItemSlice", got)
				}
				if len(items) != 1 {
					t.Fatalf("len(items) = %d, want 1", len(items))
				}
				return items[0].Label
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotLabel := tt.decodeFunc(t, tt.data)
			mutateCallerInput(t, tt.data, tt.mutate, strings.Repeat("X", len(tt.mutate)))
			if diff := gocmp.Diff(tt.wantLabel, gotLabel); diff != "" {
				t.Errorf("decoded label changed after caller input mutation (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUnmarshalOwnsInputForRawJSONValues(t *testing.T) {
	t.Parallel()

	data := []byte(`[{"name":"Symbol0","kind":1,"location":{"uri":"file:///x.go"},"data":{"index":1}}]`)
	var got WorkspaceSymbolResult
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal WorkspaceSymbolResult: %v", err)
	}
	symbols, ok := got.(WorkspaceSymbolSlice)
	if !ok {
		t.Fatalf("WorkspaceSymbolResult = %T, want WorkspaceSymbolSlice", got)
	}
	if len(symbols) != 1 {
		t.Fatalf("len(symbols) = %d, want 1", len(symbols))
	}
	if got, want := string(symbols[0].Data), `{"index":1}`; got != want {
		t.Fatalf("symbol data = %s, want %s", got, want)
	}
	if pointsIntoBytes(unsafe.SliceData(symbols[0].Data), data) {
		t.Fatal("decoded raw JSON value aliases caller-owned input; want per-message owned copy")
	}

	mutateCallerInput(t, data, "index", "muted")
	if got, want := string(symbols[0].Data), `{"index":1}`; got != want {
		t.Fatalf("symbol data changed after caller input mutation: got %s, want %s", got, want)
	}
}

func TestBoxedScalarUnionDecodePreservesStringDynamicType(t *testing.T) {
	t.Parallel()

	var got PublishDiagnosticsParams
	if err := Unmarshal([]byte(`{"uri":"file:///x.go","diagnostics":[{"range":{"start":{"line":1,"character":2},"end":{"line":1,"character":3}},"code":"E1","message":"msg"}]}`), &got); err != nil {
		t.Fatalf("Unmarshal PublishDiagnosticsParams: %v", err)
	}
	if len(got.Diagnostics) != 1 {
		t.Fatalf("len(diagnostics) = %d, want 1", len(got.Diagnostics))
	}
	if code, ok := got.Diagnostics[0].Code.(String); !ok || code != "E1" {
		t.Fatalf("diagnostic code = %#v (%T), want protocol.String(%q)", got.Diagnostics[0].Code, got.Diagnostics[0].Code, "E1")
	}
	if message, ok := got.Diagnostics[0].Message.(String); !ok || message != "msg" {
		t.Fatalf("diagnostic message = %#v (%T), want protocol.String(%q)", got.Diagnostics[0].Message, got.Diagnostics[0].Message, "msg")
	}
}

func TestZeroCopyStringOwnershipContract(t *testing.T) {
	t.Parallel()

	var got CompletionItem
	var originalLabelData *byte
	{
		data := []byte(`{"label":"pin","xPad":"` + strings.Repeat("x", 1<<16) + `"}`)
		if err := Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal CompletionItem: %v", err)
		}
		if diff := gocmp.Diff("pin", got.Label); diff != "" {
			t.Fatalf("label mismatch (-want +got):\n%s", diff)
		}

		labelData := unsafe.StringData(got.Label)
		if pointsIntoBytes(labelData, data) {
			t.Fatal("decoded label aliases caller-owned input; want per-message owned copy")
		}
		originalLabelData = labelData
	}

	for range 3 {
		runtime.GC()
	}
	if diff := gocmp.Diff("pin", got.Label); diff != "" {
		t.Fatalf("label mismatch after dropping caller input and GC (-want +got):\n%s", diff)
	}
	if unsafe.StringData(got.Label) != originalLabelData {
		t.Fatal("decoded label moved after GC; want stable owned string data")
	}
}

func TestCloneDetachesAliasedStrings(t *testing.T) {
	t.Parallel()

	var decoded CompletionItem
	if err := Unmarshal([]byte(`{"label":"detach","detail":"kept"}`), &decoded); err != nil {
		t.Fatalf("Unmarshal CompletionItem: %v", err)
	}
	cloned, err := Clone(decoded)
	if err != nil {
		t.Fatalf("Clone CompletionItem: %v", err)
	}

	if diff := gocmp.Diff(decoded.Label, cloned.Label); diff != "" {
		t.Fatalf("cloned label mismatch (-want +got):\n%s", diff)
	}
	if unsafe.StringData(decoded.Label) == unsafe.StringData(cloned.Label) {
		t.Fatal("Clone retained original aliased string storage; want detached storage")
	}
	gotDetail, ok := cloned.Detail.Get()
	if !ok {
		t.Fatal("cloned detail is absent, want present")
	}
	if got, want := gotDetail, "kept"; got != want {
		t.Fatalf("cloned detail = %q, want %q", got, want)
	}
}

func mutateCallerInput(t *testing.T, data []byte, old, replacement string) {
	t.Helper()

	if len(old) != len(replacement) {
		t.Fatalf("replacement length %d does not match original length %d", len(replacement), len(old))
	}
	i := bytes.Index(data, []byte(old))
	if i < 0 {
		t.Fatalf("input %q does not contain %q", data, old)
	}
	copy(data[i:i+len(old)], replacement)
}

func pointsIntoBytes(p *byte, data []byte) bool {
	if p == nil || len(data) == 0 {
		return false
	}
	start := uintptr(unsafe.Pointer(unsafe.SliceData(data)))
	end := start + uintptr(len(data))
	addr := uintptr(unsafe.Pointer(p))
	return start <= addr && addr < end
}
