// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/go-json-experiment/json"
)

type completionItemPointerShape struct {
	Label               string                      `json:"label"`
	LabelDetails        *CompletionItemLabelDetails `json:"labelDetails,omitzero"`
	Kind                CompletionItemKind          `json:"kind,omitzero"`
	Tags                []CompletionItemTag         `json:"tags,omitzero"`
	Detail              *string                     `json:"detail,omitzero"`
	Documentation       InlayHintTooltip            `json:"documentation,omitzero"`
	Deprecated          *bool                       `json:"deprecated,omitzero"`
	Preselect           *bool                       `json:"preselect,omitzero"`
	SortText            *string                     `json:"sortText,omitzero"`
	FilterText          *string                     `json:"filterText,omitzero"`
	InsertText          *string                     `json:"insertText,omitzero"`
	InsertTextFormat    InsertTextFormat            `json:"insertTextFormat,omitzero"`
	InsertTextMode      InsertTextMode              `json:"insertTextMode,omitzero"`
	TextEdit            CompletionItemTextEdit      `json:"textEdit,omitzero"`
	TextEditText        *string                     `json:"textEditText,omitzero"`
	AdditionalTextEdits []TextEdit                  `json:"additionalTextEdits,omitzero"`
	CommitCharacters    []string                    `json:"commitCharacters,omitzero"`
	Command             Command                     `json:"command,omitzero"`
	Data                LSPAny                      `json:"data,omitzero"`
}

func TestCompletionItemGeneratedDecoderMatchesLegacyWireRepresentation(t *testing.T) {
	payloads := map[string][]byte{
		"minimal":      []byte(`{"label":"fmt.Println"}`),
		"with_unknown": []byte(`{"label":"fmt.Println","unknown":{"nested":[1,true,null]}}`),
		"null_optionals": []byte(`{
			"label":"fmt.Println",
			"detail":null,
			"deprecated":null,
			"preselect":null,
			"sortText":null,
			"filterText":null,
			"insertText":null,
			"textEditText":null
		}`),
		"all_common_fields": []byte(`{
			"label":"fmt.Println",
			"labelDetails":{"detail":"(a ...any)","description":"builtin"},
			"kind":3,
			"tags":[1],
			"detail":"func Println(a ...any) (n int, err error)",
			"documentation":{"kind":"markdown","value":"Prints with a trailing newline."},
			"deprecated":false,
			"preselect":true,
			"sortText":"0001",
			"filterText":"fmt.Println",
			"insertText":"Println($1)",
			"insertTextFormat":2,
			"insertTextMode":2,
			"textEdit":{"range":{"start":{"line":1,"character":2},"end":{"line":1,"character":5}},"newText":"Println"},
			"textEditText":"Println",
			"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"import \"fmt\"\n"}],
			"commitCharacters":[".","("],
			"command":{"title":"after","command":"cmd.after","arguments":[{"source":"completion"}]},
			"data":{"score":7,"source":"test"}
		}`),
	}
	for name, data := range payloads {
		t.Run(name, func(t *testing.T) {
			assertCompletionItemMatchesLegacyWire(t, data)
		})
	}
}

func TestCompletionItemGeneratedDecoderMatchesCorpusArray(t *testing.T) {
	data := benchCorpus(t, "completion_result_array")

	var got []CompletionItem
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("generated completion item array decode: %v", err)
	}
	var want []completionItemPointerShape
	if err := json.Unmarshal(data, &want, json.WithUnmarshalers(unionUnmarshalers)); err != nil {
		t.Fatalf("legacy completion item array decode: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("decoded item count = %d, want %d", len(got), len(want))
	}
	for i := range got {
		assertCompletionItemWireEqual(t, &got[i], &want[i])
	}
}

func TestCompletionItemOptionalFieldsPreservePresentZeroOnWire(t *testing.T) {
	data := []byte(`{
		"label":"zeroes",
		"detail":"",
		"deprecated":false,
		"preselect":false,
		"sortText":"",
		"filterText":"",
		"insertText":"",
		"textEditText":""
	}`)
	var got CompletionItem
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("decode present zero optionals: %v", err)
	}
	assertOptionalPresentEmpty(t, got.Detail)
	assertOptionalBool(t, got.Deprecated, false)
	assertOptionalBool(t, got.Preselect, false)
	assertOptionalPresentEmpty(t, got.SortText)
	assertOptionalPresentEmpty(t, got.FilterText)
	assertOptionalPresentEmpty(t, got.InsertText)
	assertOptionalPresentEmpty(t, got.TextEditText)

	empty := ""
	falseValue := false
	want := completionItemPointerShape{
		Label:        "zeroes",
		Detail:       &empty,
		Deprecated:   &falseValue,
		Preselect:    &falseValue,
		SortText:     &empty,
		FilterText:   &empty,
		InsertText:   &empty,
		TextEditText: &empty,
	}
	assertCompletionItemWireEqual(t, &got, &want)
}

func TestCompletionItemGeneratedDecoderNullZerosExistingValue(t *testing.T) {
	got := CompletionItem{Label: "stale", Detail: NewOptional("stale detail")}
	if err := Unmarshal([]byte(`null`), &got); err != nil {
		t.Fatalf("decode null: %v", err)
	}
	if !reflect.DeepEqual(got, CompletionItem{}) {
		t.Fatalf("decode null = %#v, want zero CompletionItem", got)
	}
}

func TestCompletionItemGeneratedDecoderMergesObjectIntoExistingValue(t *testing.T) {
	detail := "kept detail"
	got := CompletionItem{Label: "old", Detail: NewOptional(detail)}
	if err := Unmarshal([]byte(`{"label":"new"}`), &got); err != nil {
		t.Fatalf("merge object: %v", err)
	}
	if got.Label != "new" {
		t.Fatalf("label = %q, want new", got.Label)
	}
	if v, ok := got.Detail.Get(); !ok || v != detail {
		t.Fatalf("detail = %q, %v; want existing detail preserved", v, ok)
	}
}

func TestCompletionItemGeneratedDecoderDuplicateMembersLastWins(t *testing.T) {
	var got CompletionItem
	if err := Unmarshal([]byte(`{"label":"a","label":"b"}`), &got); err != nil {
		t.Fatalf("duplicate label decode: %v", err)
	}
	if got.Label != "b" {
		t.Fatalf("duplicate label = %q, want last-wins %q", got.Label, "b")
	}
}

func assertCompletionItemMatchesLegacyWire(t *testing.T, data []byte) {
	t.Helper()

	var got CompletionItem
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("generated decode: %v", err)
	}
	var want completionItemPointerShape
	if err := json.Unmarshal(data, &want, json.WithUnmarshalers(unionUnmarshalers)); err != nil {
		t.Fatalf("legacy decode: %v", err)
	}
	assertCompletionItemWireEqual(t, &got, &want)
}

func assertCompletionItemWireEqual(t *testing.T, got *CompletionItem, want *completionItemPointerShape) {
	t.Helper()

	gotJSON, err := Marshal(got)
	if err != nil {
		t.Fatalf("marshal generated completion item: %v", err)
	}
	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal legacy completion item: %v", err)
	}
	if !bytes.Equal(gotJSON, wantJSON) {
		t.Fatalf("wire mismatch\ngot:  %s\nwant: %s\ngot value:  %#v\nwant value: %#v", gotJSON, wantJSON, got, want)
	}
}

func assertOptionalPresentEmpty(t *testing.T, got Optional[string]) {
	t.Helper()
	v, ok := got.Get()
	if !ok || v != "" {
		t.Fatalf("optional string = %q, %v; want empty string present", v, ok)
	}
}

func assertOptionalBool(t *testing.T, got Optional[bool], want bool) {
	t.Helper()
	v, ok := got.Get()
	if !ok || v != want {
		t.Fatalf("optional bool = %v, %v; want %v, true", v, ok, want)
	}
}
