// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"testing"

	"go.lsp.dev/jsonrpc2"
)

// Invalid-UTF-8 and duplicate-member behavior is pinned by the wireOptions
// tests (wire_options_test.go): both are deliberately relaxed — U+FFFD
// mangling and last-wins respectively — instead of rejecting the message.

func TestBackendMemberNamesAreCaseSensitive(t *testing.T) {
	t.Parallel()

	var v Position
	if err := Unmarshal([]byte(`{"Line":99,"line":1,"Character":88,"character":2}`), &v); err != nil {
		t.Fatalf("Unmarshal exact-case fixture: %v", err)
	}
	if v.Line != 1 || v.Character != 2 {
		t.Fatalf("decoded Position = %+v, want exact lowercase members line=1 character=2", v)
	}

	var onlyWrongCase Position
	if err := Unmarshal([]byte(`{"Line":99,"Character":88}`), &onlyWrongCase); err != nil {
		t.Fatalf("Unmarshal wrong-case unknown fields should preserve forward compatibility: %v", err)
	}
	if onlyWrongCase.Line != 0 || onlyWrongCase.Character != 0 {
		t.Fatalf("wrong-case members populated Position = %+v, want zero value", onlyWrongCase)
	}
}

func TestLSPCodecRawMessageLifetime(t *testing.T) {
	t.Parallel()

	data := []byte(`{"params":{"x":1}}`)
	want := append([]byte(nil), data...)

	var raw jsonrpc2.RawMessage
	if err := (lspCodec{}).Unmarshal(data, &raw); err != nil {
		t.Fatalf("Unmarshal RawMessage: %v", err)
	}
	for i := range data {
		data[i] = 'x'
	}
	if !bytes.Equal(raw, want) {
		t.Fatalf("RawMessage aliases input buffer after mutation: got %q, want %q", raw, want)
	}
}

func TestLSPCodecRawMessagePassthrough(t *testing.T) {
	t.Parallel()

	raw := jsonrpc2.RawMessage(`{"k":<unescaped>,"n":1}`)
	got, err := (lspCodec{}).Marshal(raw)
	if err != nil {
		t.Fatalf("Marshal RawMessage: %v", err)
	}
	if !bytes.Equal(got, raw) {
		t.Fatalf("Marshal RawMessage = %q, want verbatim %q", got, raw)
	}

	got, err = (lspCodec{}).Marshal(jsonrpc2.RawMessage(nil))
	if err != nil {
		t.Fatalf("Marshal nil RawMessage: %v", err)
	}
	if string(got) != "null" {
		t.Fatalf("Marshal nil RawMessage = %q, want null", got)
	}
}
