// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"strings"
	"testing"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

func TestOptionalAccessors(t *testing.T) {
	var opt Optional[string]
	if !opt.IsZero() {
		t.Fatal("zero Optional is present")
	}
	if got, ok := opt.Get(); ok || got != "" {
		t.Fatalf("zero Get() = %q, %v; want empty, false", got, ok)
	}

	opt.Set("")
	if opt.IsZero() {
		t.Fatal("present empty string reported zero")
	}
	if got, ok := opt.Get(); !ok || got != "" {
		t.Fatalf("present empty Get() = %q, %v; want empty, true", got, ok)
	}

	opt.Set("value")
	if got, ok := opt.Get(); !ok || got != "value" {
		t.Fatalf("present value Get() = %q, %v; want value, true", got, ok)
	}
	opt.Clear()
	if !opt.IsZero() {
		t.Fatal("cleared Optional is present")
	}
	if got, ok := opt.Get(); ok || got != "" {
		t.Fatalf("cleared Get() = %q, %v; want empty, false", got, ok)
	}
}

func TestOptionalJSON(t *testing.T) {
	t.Run("standalone absent marshals null", func(t *testing.T) {
		got, err := json.Marshal(Optional[string]{})
		if err != nil {
			t.Fatalf("marshal absent Optional: %v", err)
		}
		if string(got) != "null" {
			t.Fatalf("marshal absent Optional = %s, want null", got)
		}
	})

	t.Run("present zero survives omitzero", func(t *testing.T) {
		type holder struct {
			S Optional[string] `json:"s,omitzero"`
			B Optional[bool]   `json:"b,omitzero"`
		}
		got, err := json.Marshal(holder{S: NewOptional(""), B: NewOptional(false)})
		if err != nil {
			t.Fatalf("marshal holder: %v", err)
		}
		if string(got) != `{"s":"","b":false}` {
			t.Fatalf("marshal holder = %s, want present zero fields", got)
		}
	})

	t.Run("absent omits", func(t *testing.T) {
		type holder struct {
			S Optional[string] `json:"s,omitzero"`
		}
		got, err := json.Marshal(holder{})
		if err != nil {
			t.Fatalf("marshal holder: %v", err)
		}
		if string(got) != `{}` {
			t.Fatalf("marshal absent holder = %s, want {}", got)
		}
	})

	t.Run("null clears", func(t *testing.T) {
		opt := NewOptional("old")
		if err := json.Unmarshal([]byte(`null`), &opt); err != nil {
			t.Fatalf("unmarshal null: %v", err)
		}
		if !opt.IsZero() {
			t.Fatalf("unmarshal null left value present: %#v", opt)
		}
	})
}

func TestDecodeStringFrom(t *testing.T) {
	var got string
	if err := decodeStringFrom(jsontext.NewDecoder(strings.NewReader(`"hello"`)), &got); err != nil {
		t.Fatalf("decode string: %v", err)
	}
	if got != "hello" {
		t.Fatalf("decoded string = %q, want hello", got)
	}
	if err := decodeStringFrom(jsontext.NewDecoder(strings.NewReader(`null`)), &got); err != nil {
		t.Fatalf("decode null string: %v", err)
	}
	if got != "" {
		t.Fatalf("decoded null string = %q, want empty", got)
	}
	if err := decodeStringFrom(jsontext.NewDecoder(strings.NewReader(`123`)), &got); err == nil {
		t.Fatal("decode numeric string succeeded")
	}
}

func TestDecodeOptionalStringFrom(t *testing.T) {
	var got Optional[string]
	if err := decodeOptionalStringFrom(jsontext.NewDecoder(strings.NewReader(`""`)), &got); err != nil {
		t.Fatalf("decode empty optional string: %v", err)
	}
	if v, ok := got.Get(); !ok || v != "" {
		t.Fatalf("decoded empty optional = %q, %v; want empty, true", v, ok)
	}
	if err := decodeOptionalStringFrom(jsontext.NewDecoder(strings.NewReader(`null`)), &got); err != nil {
		t.Fatalf("decode null optional string: %v", err)
	}
	if !got.IsZero() {
		t.Fatalf("decoded null optional = %#v, want absent", got)
	}
	if err := decodeOptionalStringFrom(jsontext.NewDecoder(strings.NewReader(`false`)), &got); err == nil {
		t.Fatal("decode bool optional string succeeded")
	}
}

func TestDecodeOptionalBoolFrom(t *testing.T) {
	var got Optional[bool]
	if err := decodeOptionalBoolFrom(jsontext.NewDecoder(strings.NewReader(`false`)), &got); err != nil {
		t.Fatalf("decode false optional bool: %v", err)
	}
	if v, ok := got.Get(); !ok || v {
		t.Fatalf("decoded false optional = %v, %v; want false, true", v, ok)
	}
	if err := decodeOptionalBoolFrom(jsontext.NewDecoder(strings.NewReader(`null`)), &got); err != nil {
		t.Fatalf("decode null optional bool: %v", err)
	}
	if !got.IsZero() {
		t.Fatalf("decoded null optional bool = %#v, want absent", got)
	}
	if err := decodeOptionalBoolFrom(jsontext.NewDecoder(strings.NewReader(`"false"`)), &got); err == nil {
		t.Fatal("decode string optional bool succeeded")
	}
}

func TestDecodeOptionalInt32From(t *testing.T) {
	var got Optional[int32]
	if err := decodeOptionalInt32From(jsontext.NewDecoder(strings.NewReader(`0`)), &got); err != nil {
		t.Fatalf("decode zero optional int32: %v", err)
	}
	if v, ok := got.Get(); !ok || v != 0 {
		t.Fatalf("decoded zero optional int32 = %d, %v; want 0, true", v, ok)
	}
	if err := decodeOptionalInt32From(jsontext.NewDecoder(strings.NewReader(`null`)), &got); err != nil {
		t.Fatalf("decode null optional int32: %v", err)
	}
	if !got.IsZero() {
		t.Fatalf("decoded null optional int32 = %#v, want absent", got)
	}
	if err := decodeOptionalInt32From(jsontext.NewDecoder(strings.NewReader(`\"1\"`)), &got); err == nil {
		t.Fatal("decode string optional int32 succeeded")
	}
}

func TestDecodeInt32From(t *testing.T) {
	var got int32
	if err := decodeInt32From(jsontext.NewDecoder(strings.NewReader(`-42`)), &got); err != nil {
		t.Fatalf("decode int32: %v", err)
	}
	if got != -42 {
		t.Fatalf("decoded int32 = %d, want -42", got)
	}
	if err := decodeInt32From(jsontext.NewDecoder(strings.NewReader(`2147483648`)), &got); err == nil {
		t.Fatal("decode overflowing int32 succeeded")
	}
	if err := decodeInt32From(jsontext.NewDecoder(strings.NewReader(`null`)), &got); err != nil {
		t.Fatalf("decode null int32: %v", err)
	}
	if got != 0 {
		t.Fatalf("decoded null int32 = %d, want 0", got)
	}
}

func TestDecodeUint32From(t *testing.T) {
	var got CompletionItemKind
	if err := decodeUint32From(jsontext.NewDecoder(strings.NewReader(`3`)), &got); err != nil {
		t.Fatalf("decode uint32 enum: %v", err)
	}
	if got != CompletionItemKindFunction {
		t.Fatalf("decoded enum = %d, want %d", got, CompletionItemKindFunction)
	}
	if err := decodeUint32From(jsontext.NewDecoder(strings.NewReader(`null`)), &got); err != nil {
		t.Fatalf("decode null uint32 enum: %v", err)
	}
	if got != 0 {
		t.Fatalf("decoded null enum = %d, want 0", got)
	}
	if err := decodeUint32From(jsontext.NewDecoder(strings.NewReader(`"3"`)), &got); err == nil {
		t.Fatal("decode string uint32 enum succeeded")
	}
}

func TestDecodeInlayHintTooltipFrom(t *testing.T) {
	var got InlayHintTooltip
	if err := decodeInlayHintTooltipFrom(jsontext.NewDecoder(strings.NewReader(`"plain"`)), &got); err != nil {
		t.Fatalf("decode string tooltip: %v", err)
	}
	if v, ok := got.(String); !ok || string(v) != "plain" {
		t.Fatalf("decoded string tooltip = %#v, want String plain", got)
	}
	if err := decodeInlayHintTooltipFrom(jsontext.NewDecoder(strings.NewReader(`null`)), &got); err != nil {
		t.Fatalf("decode null tooltip: %v", err)
	}
	if got != nil {
		t.Fatalf("decoded null tooltip = %#v, want nil", got)
	}
	if err := decodeInlayHintTooltipFrom(jsontext.NewDecoder(strings.NewReader(`{"kind":"markdown","value":"doc"}`)), &got); err != nil {
		t.Fatalf("decode object tooltip: %v", err)
	}
	markup, ok := got.(*MarkupContent)
	if !ok || markup.Kind != "markdown" || markup.Value != "doc" {
		t.Fatalf("decoded object tooltip = %#v, want markup content", got)
	}
	if err := decodeInlayHintTooltipFrom(jsontext.NewDecoder(strings.NewReader(`false`)), &got); err == nil {
		t.Fatal("decode bool tooltip succeeded")
	}
}

func TestIsZeroCommand(t *testing.T) {
	if !isZeroCommand(Command{}) {
		t.Fatal("zero Command reported non-zero")
	}
	tooltip := ""
	tests := []Command{
		{Title: "title"},
		{Tooltip: &tooltip},
		{Command: "cmd"},
		{Arguments: []LSPAny{[]byte(`1`)}},
	}
	for _, tc := range tests {
		if isZeroCommand(tc) {
			t.Fatalf("non-zero Command reported zero: %#v", tc)
		}
	}
}
