// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"testing"
)

// loadTestModel decodes the vendored metaModel.json fixture.
func loadTestModel(t *testing.T) *MetaModel {
	t.Helper()
	m, err := Load("testdata/metaModel.json")
	if err != nil {
		t.Fatalf("load meta model: %v", err)
	}
	return m
}

func TestLoadCounts(t *testing.T) {
	m := loadTestModel(t)

	tests := map[string]struct {
		got  int
		want int
	}{
		"success: requests":      {len(m.Requests), 69},
		"success: notifications": {len(m.Notifications), 26},
		"success: structures":    {len(m.Structures), 387},
		"success: enumerations":  {len(m.Enumerations), 40},
		"success: typeAliases":   {len(m.TypeAliases), 22},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("count = %d, want %d", tt.got, tt.want)
			}
		})
	}
}

func TestDecodeTypeKinds(t *testing.T) {
	m := loadTestModel(t)

	// LSPAny is the canonical 8-arity recursive union; assert it decodes as an or.
	var lspAny *TypeAlias
	for _, ta := range m.TypeAliases {
		if ta.Name == "LSPAny" {
			lspAny = ta
			break
		}
	}
	if lspAny == nil {
		t.Fatal("LSPAny type alias not found")
	}
	if lspAny.Type.Kind != KindOr {
		t.Fatalf("LSPAny kind = %q, want %q", lspAny.Type.Kind, KindOr)
	}
	if got, want := len(lspAny.Type.Items), 8; got != want {
		t.Fatalf("LSPAny arity = %d, want %d", got, want)
	}

	// WorkspaceEdit.documentChanges is the 4-arity reference union.
	var we *Structure
	for _, s := range m.Structures {
		if s.Name == "WorkspaceEdit" {
			we = s
			break
		}
	}
	if we == nil {
		t.Fatal("WorkspaceEdit structure not found")
	}
	var docChanges *Property
	for _, p := range we.Properties {
		if p.Name == "documentChanges" {
			docChanges = p
			break
		}
	}
	if docChanges == nil {
		t.Fatal("WorkspaceEdit.documentChanges not found")
	}
	// documentChanges is `(TextDocumentEdit | (CreateFile|RenameFile|DeleteFile)[]) | null`-ish;
	// in the 3.18 model it is an array of an or. Assert the array element is an or of 4.
	if docChanges.Type.Kind != KindArray {
		t.Fatalf("documentChanges kind = %q, want %q", docChanges.Type.Kind, KindArray)
	}
	if el := docChanges.Type.Element; el == nil || el.Kind != KindOr || len(el.Items) != 4 {
		t.Fatalf("documentChanges element = %+v, want or of 4", docChanges.Type.Element)
	}
}

func TestURIBaseTypeLoweringUsesExternalFieldAndLocalUnionWrapper(t *testing.T) {
	if got := baseGoType(BaseURI); got != generatedURIType {
		t.Fatalf("baseGoType(BaseURI) = %q, want %s", got, generatedURIType)
	}
	if got := baseGoType(BaseDocumentURI); got != generatedURIType {
		t.Fatalf("baseGoType(BaseDocumentURI) = %q, want %s", got, generatedURIType)
	}
	if got, token := scalarWrapper(BaseURI); got != unionURIWrapperType || token != '"' {
		t.Fatalf("scalarWrapper(BaseURI) = %q, %q, want URI, quote", got, token)
	}
	if got, token := scalarWrapper(BaseDocumentURI); got != unionURIWrapperType || token != '"' {
		t.Fatalf("scalarWrapper(BaseDocumentURI) = %q, %q, want URI, quote", got, token)
	}
}
