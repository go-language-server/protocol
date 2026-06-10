// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"testing"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	gocmp "github.com/google/go-cmp/cmp"
)

// The shadow types below mirror generated structs field-for-field but carry
// no methods anywhere in their field graph, so decoding into them exercises
// the pure go-json-experiment reflection path. Union-typed fields are held as
// raw jsontext.Value. Comparing the re-marshaled wire forms of a byte-walker
// decode against a shadow decode pins the walkers to the reflection
// semantics.

type positionShadow struct {
	Line      uint32 `json:"line"`
	Character uint32 `json:"character"`
}

type rangeShadow struct {
	Start positionShadow `json:"start"`
	End   positionShadow `json:"end"`
}

type locationShadow struct {
	URI   string      `json:"uri"`
	Range rangeShadow `json:"range"`
}

type workspaceSymbolShadow struct {
	Name          string         `json:"name"`
	Kind          SymbolKind     `json:"kind"`
	Tags          []SymbolTag    `json:"tags,omitzero"`
	ContainerName *string        `json:"containerName,omitzero"`
	Location      jsontext.Value `json:"location"`
	Data          jsontext.Value `json:"data,omitzero"`
}

type symbolInformationShadow struct {
	Name          string         `json:"name"`
	Kind          SymbolKind     `json:"kind"`
	Tags          []SymbolTag    `json:"tags,omitzero"`
	ContainerName *string        `json:"containerName,omitzero"`
	Deprecated    *bool          `json:"deprecated,omitzero"`
	Location      locationShadow `json:"location"`
}

type didChangeShadow struct {
	TextDocument struct {
		URI     string `json:"uri"`
		Version int32  `json:"version"`
	} `json:"textDocument"`
	ContentChanges []jsontext.Value `json:"contentChanges"`
}

type semanticTokensShadow struct {
	ResultID *string  `json:"resultId,omitzero"`
	Data     []uint32 `json:"data"`
}

type publishDiagnosticsShadow struct {
	URI         string           `json:"uri"`
	Version     *int32           `json:"version,omitzero"`
	Diagnostics []jsontext.Value `json:"diagnostics"`
}

// decodeWireAny decodes JSON into the generic representation used for
// order-insensitive wire comparison.
func decodeWireAny(t *testing.T, b []byte) any {
	t.Helper()
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatalf("decode wire form %s: %v", b, err)
	}
	return v
}

// assertWireEqual asserts two JSON encodings denote the same value modulo
// object member order.
func assertWireEqual(t *testing.T, got, want []byte) {
	t.Helper()
	if diff := gocmp.Diff(decodeWireAny(t, want), decodeWireAny(t, got)); diff != "" {
		t.Errorf("wire mismatch (-reflection +byte):\n%s\ngot:  %s\nwant: %s", diff, got, want)
	}
}

// TestByteDecodeMatchesReflectionShadow decodes each corpus payload through
// the byte walkers and through a methodless shadow type (pure reflection),
// then compares the re-marshaled wire forms.
func TestByteDecodeMatchesReflectionShadow(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		corpus string
		decode func(t *testing.T, data []byte) ([]byte, []byte)
	}{
		"workspace_symbol_result": {
			corpus: "workspace_symbol_result",
			decode: func(t *testing.T, data []byte) ([]byte, []byte) {
				t.Helper()
				var got WorkspaceSymbolResult
				if err := Unmarshal(data, &got); err != nil {
					t.Fatalf("byte decode: %v", err)
				}
				var want []workspaceSymbolShadow
				if err := json.Unmarshal(data, &want, wireOptions); err != nil {
					t.Fatalf("shadow decode: %v", err)
				}
				return marshalBoth(t, got, want)
			},
		},
		"workspace_symbol_result_info": {
			corpus: "workspace_symbol_result_info",
			decode: func(t *testing.T, data []byte) ([]byte, []byte) {
				t.Helper()
				var got WorkspaceSymbolResult
				if err := Unmarshal(data, &got); err != nil {
					t.Fatalf("byte decode: %v", err)
				}
				var want []symbolInformationShadow
				if err := json.Unmarshal(data, &want, wireOptions); err != nil {
					t.Fatalf("shadow decode: %v", err)
				}
				return marshalBoth(t, got, want)
			},
		},
		"didchange": {
			corpus: "didchange",
			decode: func(t *testing.T, data []byte) ([]byte, []byte) {
				t.Helper()
				var got DidChangeTextDocumentParams
				if err := Unmarshal(data, &got); err != nil {
					t.Fatalf("byte decode: %v", err)
				}
				var want didChangeShadow
				if err := json.Unmarshal(data, &want, wireOptions); err != nil {
					t.Fatalf("shadow decode: %v", err)
				}
				return marshalBoth(t, got, want)
			},
		},
		"semantic_tokens": {
			corpus: "semantic_tokens",
			decode: func(t *testing.T, data []byte) ([]byte, []byte) {
				t.Helper()
				var got SemanticTokens
				if err := Unmarshal(data, &got); err != nil {
					t.Fatalf("byte decode: %v", err)
				}
				var want semanticTokensShadow
				if err := json.Unmarshal(data, &want, wireOptions); err != nil {
					t.Fatalf("shadow decode: %v", err)
				}
				return marshalBoth(t, got, want)
			},
		},
		"publish_diagnostics": {
			corpus: "publish_diagnostics",
			decode: func(t *testing.T, data []byte) ([]byte, []byte) {
				t.Helper()
				var got PublishDiagnosticsParams
				if err := Unmarshal(data, &got); err != nil {
					t.Fatalf("byte decode: %v", err)
				}
				var want publishDiagnosticsShadow
				if err := json.Unmarshal(data, &want, wireOptions); err != nil {
					t.Fatalf("shadow decode: %v", err)
				}
				return marshalBoth(t, got, want)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			data := benchCorpus(t, tt.corpus)
			gotJSON, wantJSON := tt.decode(t, data)
			assertWireEqual(t, gotJSON, wantJSON)
		})
	}
}

func marshalBoth(t *testing.T, got, want any) (gotJSON, wantJSON []byte) {
	t.Helper()
	gotJSON, err := Marshal(got)
	if err != nil {
		t.Fatalf("marshal byte-decoded value: %v", err)
	}
	wantJSON, err = json.Marshal(want, wireOptions)
	if err != nil {
		t.Fatalf("marshal shadow value: %v", err)
	}
	return gotJSON, wantJSON
}

// TestByteDecodeCorpusFixpoint asserts decode→marshal reaches a fixpoint for
// every bench category: re-decoding the marshaled form and marshaling again
// must reproduce identical bytes.
func TestByteDecodeCorpusFixpoint(t *testing.T) {
	t.Parallel()

	for _, bc := range append(append([]benchCase(nil), benchCases...), dispatchBenchCases...) {
		t.Run(bc.name, func(t *testing.T) {
			t.Parallel()

			data := benchCorpus(t, bc.corpus)
			v1 := bc.newDst()
			if err := Unmarshal(data, v1); err != nil {
				t.Fatalf("decode corpus: %v", err)
			}
			m1, err := Marshal(v1)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			v2 := bc.newDst()
			if err := Unmarshal(m1, v2); err != nil {
				t.Fatalf("re-decode marshaled form: %v\n%s", err, m1)
			}
			m2, err := Marshal(v2)
			if err != nil {
				t.Fatalf("re-marshal: %v", err)
			}
			if !bytes.Equal(m1, m2) {
				t.Errorf("marshal fixpoint diverged\nfirst:  %s\nsecond: %s", m1, m2)
			}
		})
	}
}

func TestByteDecodeRejectsMalformedSkippedValues(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		dst   any
	}{
		"unknown field invalid scalar": {
			input: `{"line":1,"x":<bad>,"character":2}`,
			dst:   new(Position),
		},
		"unknown field trailing comma array": {
			input: `{"line":1,"x":[1,],"character":2}`,
			dst:   new(Position),
		},
		"unknown field invalid member name": {
			input: `{"\q":1,"line":1,"character":2}`,
			dst:   new(Position),
		},
		"raw LSPAny field invalid payload": {
			input: `{"name":"s","kind":12,"location":{"uri":"file:///a"},"data":{"broken":[1,]}}`,
			dst:   new(WorkspaceSymbol),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if err := Unmarshal([]byte(tt.input), tt.dst); err == nil {
				t.Fatalf("Unmarshal(%s) succeeded; malformed skipped/raw value must be rejected", tt.input)
			}
		})
	}
}

// FuzzByteDecodeOracle mutates LSP-shaped payloads and cross-checks the byte
// walker against the pure-reflection shadow decode. The byte walkers must
// accept everything reflection accepts and produce the same value; inputs both
// paths reject are ignored.
func FuzzByteDecodeOracle(f *testing.F) {
	f.Add([]byte(`{"name":"s","kind":12,"location":{"uri":"file:///a"},"data":{"i":1}}`))
	f.Add([]byte(`{"name":"s","kind":5,"tags":[1],"containerName":"c","location":{"uri":"file:///a","range":{"start":{"line":1,"character":2},"end":{"line":3,"character":4}}}}`))
	f.Add([]byte(`{"label":"item","kind":1,"detail":"","documentation":{"kind":"markdown","value":"d"},"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}},"newText":"x"}}`))
	f.Add([]byte(`{"resultId":"r","data":[0,1,2,4294967295]}`))
	f.Add([]byte("{\"name\":\"\xff\",\"kind\":1,\"location\":{\"uri\":\"u\"},\"name\":\"dup\"}"))

	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzOracleOne[WorkspaceSymbol, workspaceSymbolShadow](t, data)
		fuzzOracleOne[CompletionItem, completionItemPointerShape](t, data)
		fuzzOracleOne[SemanticTokens, semanticTokensShadow](t, data)
	})
}

func fuzzOracleOne[T, S any](t *testing.T, data []byte) {
	t.Helper()

	var shadow S
	shadowErr := json.Unmarshal(data, &shadow, wireOptions)

	var got T
	byteErr := Unmarshal(data, &got)

	if shadowErr == nil && byteErr != nil {
		// The shadow is a relaxation of the typed schema (union fields are
		// raw), so a byte-side rejection is only a bug when the shadow's
		// re-marshaled wire form — schema-shaped by construction modulo the
		// union arms — decodes cleanly. Inputs violating a union arm fail
		// both ways and are correct rejections.
		wantJSON, err := json.Marshal(shadow, wireOptions)
		if err == nil {
			var probe T
			if Unmarshal(wantJSON, &probe) == nil {
				t.Errorf("byte walker rejected input reflection accepts: %v\ninput: %q", byteErr, data)
			}
		}
		return
	}
	if shadowErr != nil || byteErr != nil {
		return // documented leniency class or agreement on rejection
	}
	gotJSON, err := Marshal(got)
	if err != nil {
		t.Fatalf("marshal byte-decoded value: %v", err)
	}
	wantJSON, err := json.Marshal(shadow, wireOptions)
	if err != nil {
		t.Fatalf("marshal shadow value: %v", err)
	}
	// Union fields are raw in the shadow, while the typed decode normalizes
	// them through arm dispatch (e.g. the lenient tier materializes a zero
	// arm). Run the shadow wire form through one typed decode so both sides
	// carry the same normalization; the strict, unnormalized differential is
	// pinned by TestByteDecodeMatchesReflectionShadow on the corpus.
	var norm T
	if err := Unmarshal(wantJSON, &norm); err != nil {
		return // shadow wire form is not typed-schema-valid (union arm mismatch)
	}
	wantNorm, err := Marshal(norm)
	if err != nil {
		t.Fatalf("re-marshal normalized shadow value: %v", err)
	}
	var gotAny, wantAny any
	gotErr := json.Unmarshal(gotJSON, &gotAny)
	wantErr := json.Unmarshal(wantNorm, &wantAny)
	if gotErr != nil || wantErr != nil {
		// Raw passthrough fields can carry numbers outside the generic
		// float64 model (e.g. 1e700 in an LSPAny payload); compare the wire
		// bytes directly when the generic reparse cannot represent them.
		if !bytes.Equal(gotJSON, wantNorm) {
			t.Errorf("wire mismatch beyond generic model\nbyte:   %s\nshadow: %s\ninput: %q", gotJSON, wantNorm, data)
		}
		return
	}
	if diff := gocmp.Diff(wantAny, gotAny); diff != "" {
		t.Errorf("wire mismatch (-reflection +byte):\n%s\ninput: %q", diff, data)
	}
}
