// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"testing"

	"github.com/go-json-experiment/json"
)

type diagnosticPointerShape struct {
	Range              Range                          `json:"range"`
	Severity           DiagnosticSeverity             `json:"severity,omitzero"`
	Code               ProgressToken                  `json:"code,omitzero"`
	CodeDescription    CodeDescription                `json:"codeDescription,omitzero"`
	Source             *string                        `json:"source,omitzero"`
	Message            InlayHintTooltip               `json:"message"`
	Tags               []DiagnosticTag                `json:"tags,omitzero"`
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitzero"`
	Data               LSPAny                         `json:"data,omitzero"`
}

type publishDiagnosticsPointerShape struct {
	URI         DocumentURI              `json:"uri"`
	Version     *int32                   `json:"version,omitzero"`
	Diagnostics []diagnosticPointerShape `json:"diagnostics"`
}

func TestPublishDiagnosticsGeneratedCodecsMatchLegacyWireRepresentation(t *testing.T) {
	payloads := map[string][]byte{
		"present_zero_optionals": []byte(`{
			"uri":"file:///tmp/main.go",
			"version":0,
			"diagnostics":[{
				"range":{"start":{"line":0,"character":1},"end":{"line":0,"character":5}},
				"severity":1,
				"code":"E0001",
				"codeDescription":{"href":"https://example.invalid/E0001"},
				"source":"",
				"message":"plain diagnostic",
				"tags":[1],
				"relatedInformation":[{"location":{"uri":"file:///tmp/other.go","range":{"start":{"line":2,"character":3},"end":{"line":2,"character":4}}},"message":"related"}],
				"data":{"rule":"E0001"}
			}]
		}`),
		"null_optionals": []byte(`{
			"uri":"file:///tmp/main.go",
			"version":null,
			"diagnostics":[{
				"range":{"start":{"line":0,"character":1},"end":{"line":0,"character":5}},
				"source":null,
				"message":{"kind":"markdown","value":"**diagnostic**"}
			}]
		}`),
	}
	for name, data := range payloads {
		t.Run(name, func(t *testing.T) {
			assertPublishDiagnosticsMatchesLegacyWire(t, data)
		})
	}
}

func TestPublishDiagnosticsGeneratedCodecsPreservePresentZeroOptionals(t *testing.T) {
	data := []byte(`{
		"uri":"file:///tmp/main.go",
		"version":0,
		"diagnostics":[{
			"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}},
			"source":"",
			"message":"m"
		}]
	}`)
	var got PublishDiagnosticsParams
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("decode publish diagnostics: %v", err)
	}
	if v, ok := got.Version.Get(); !ok || v != 0 {
		t.Fatalf("version = %d, %v; want 0, true", v, ok)
	}
	if len(got.Diagnostics) != 1 {
		t.Fatalf("diagnostic count = %d, want 1", len(got.Diagnostics))
	}
	if v, ok := got.Diagnostics[0].Source.Get(); !ok || v != "" {
		t.Fatalf("source = %q, %v; want empty, true", v, ok)
	}
}

func assertPublishDiagnosticsMatchesLegacyWire(t *testing.T, data []byte) {
	t.Helper()

	var got PublishDiagnosticsParams
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("generated decode: %v", err)
	}
	var want publishDiagnosticsPointerShape
	if err := json.Unmarshal(data, &want, json.WithUnmarshalers(unionUnmarshalers)); err != nil {
		t.Fatalf("legacy decode: %v", err)
	}
	gotJSON, err := Marshal(got)
	if err != nil {
		t.Fatalf("marshal generated: %v", err)
	}
	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal legacy: %v", err)
	}
	if !bytes.Equal(gotJSON, wantJSON) {
		t.Fatalf("wire mismatch\ngot:  %s\nwant: %s\ngot value:  %#v\nwant value: %#v", gotJSON, wantJSON, got, want)
	}
}
