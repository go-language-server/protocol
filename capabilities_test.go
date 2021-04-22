// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testWorkspaceClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const want = `{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"],"normalizesLineEndings":true,"changeAnnotationSupport":{"groupsOnLabel":true}},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true,"semanticTokens":{"refreshSupport":true},"codeLens":{"refreshSupport":true},"fileOperations":{"dynamicRegistration":true,"didCreate":true,"willCreate":true,"didRename":true,"willRename":true,"didDelete":true,"willDelete":true}}`
	wantType := WorkspaceClientCapabilities{
		ApplyEdit: true,
		WorkspaceEdit: &WorkspaceClientCapabilitiesWorkspaceEdit{
			DocumentChanges:       true,
			FailureHandling:       "FailureHandling",
			ResourceOperations:    []string{"ResourceOperations"},
			NormalizesLineEndings: true,
			ChangeAnnotationSupport: &WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport{
				GroupsOnLabel: true,
			},
		},
		DidChangeConfiguration: &WorkspaceClientCapabilitiesDidChangeConfiguration{
			DynamicRegistration: true,
		},
		DidChangeWatchedFiles: &WorkspaceClientCapabilitiesDidChangeWatchedFiles{
			DynamicRegistration: true,
		},
		Symbol: &WorkspaceClientCapabilitiesSymbol{
			DynamicRegistration: true,
			SymbolKind: &WorkspaceClientCapabilitiesSymbolKind{
				ValueSet: []SymbolKind{
					SymbolKindFile,
					SymbolKindModule,
					SymbolKindNamespace,
					SymbolKindPackage,
					SymbolKindClass,
					SymbolKindMethod,
				},
			},
		},
		ExecuteCommand: &WorkspaceClientCapabilitiesExecuteCommand{
			DynamicRegistration: true,
		},
		WorkspaceFolders: true,
		Configuration:    true,
		SemanticTokens: &WorkspaceClientCapabilitiesSemanticTokens{
			RefreshSupport: true,
		},
		CodeLens: &WorkspaceClientCapabilitiesCodeLens{
			RefreshSupport: true,
		},
		FileOperations: &WorkspaceClientCapabilitiesFileOperations{
			DynamicRegistration: true,
			DidCreate:           true,
			WillCreate:          true,
			DidRename:           true,
			WillRename:          true,
			DidDelete:           true,
			WillDelete:          true,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          WorkspaceClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             WorkspaceClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got WorkspaceClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesSynchronization(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"didSave":true,"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesSynchronization{
		DidSave:             true,
		DynamicRegistration: true,
		WillSave:            true,
		WillSaveWaitUntil:   true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesSynchronization
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesSynchronization{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesSynchronization
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesSynchronization{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesSynchronization
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesCompletion(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true,"tagSupport":{"valueSet":[1]},"insertReplaceSupport":true,"resolveSupport":{"properties":["test","properties"]},"insertTextModeSupport":{"valueSet":[1,2]}},"completionItemKind":{"valueSet":[1]},"contextSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesCompletion{
		DynamicRegistration: true,
		CompletionItem: &TextDocumentClientCapabilitiesCompletionItem{
			SnippetSupport:          true,
			CommitCharactersSupport: true,
			DocumentationFormat: []MarkupKind{
				PlainText,
				Markdown,
			},
			DeprecatedSupport: true,
			PreselectSupport:  true,
			TagSupport: &TextDocumentClientCapabilitiesCompletionItemTagSupport{
				ValueSet: []CompletionItemTag{
					CompletionItemTagDeprecated,
				},
			},
			InsertReplaceSupport: true,
			ResolveSupport: &TextDocumentClientCapabilitiesCompletionItemResolveSupport{
				Properties: []string{"test", "properties"},
			},
			InsertTextModeSupport: &TextDocumentClientCapabilitiesCompletionItemInsertTextModeSupport{
				ValueSet: []InsertTextMode{
					InsertTextModeAsIs,
					InsertTextModeAdjustIndentation,
				},
			},
		},
		CompletionItemKind: &TextDocumentClientCapabilitiesCompletionItemKind{
			ValueSet: []CompletionItemKind{CompletionItemKindText},
		},
		ContextSupport: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesCompletion
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesCompletion{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesCompletion
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesCompletion{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesCompletion
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesHover(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesHover{
		DynamicRegistration: true,
		ContentFormat: []MarkupKind{
			PlainText,
			Markdown,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesHover
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesHover{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesHover
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesHover{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesHover
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesSignatureHelp(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"],"parameterInformation":{"labelOffsetSupport":true},"activeParameterSupport":true},"contextSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesSignatureHelp{
		DynamicRegistration: true,
		SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
			DocumentationFormat: []MarkupKind{
				PlainText,
				Markdown,
			},
			ParameterInformation: &TextDocumentClientCapabilitiesParameterInformation{
				LabelOffsetSupport: true,
			},
			ActiveParameterSupport: true,
		},
		ContextSupport: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesSignatureHelp
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesSignatureHelp{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesSignatureHelp
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesSignatureHelp{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesSignatureHelp
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesReferences(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesReferences{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesReferences
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesReferences{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesReferences
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesReferences{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesReferences
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesDocumentHighlight(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesDocumentHighlight{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesDocumentHighlight
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDocumentHighlight{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesDocumentHighlight
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesDocumentHighlight{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesDocumentHighlight
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesDocumentSymbol(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true,"tagSupport":{"valueSet":[1]},"labelSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesDocumentSymbol{
		DynamicRegistration: true,
		SymbolKind: &WorkspaceClientCapabilitiesSymbolKind{
			ValueSet: []SymbolKind{
				SymbolKindFile,
				SymbolKindModule,
				SymbolKindNamespace,
				SymbolKindPackage,
				SymbolKindClass,
				SymbolKindMethod,
			},
		},
		HierarchicalDocumentSymbolSupport: true,
		TagSupport: &TextDocumentClientCapabilitiesDocumentSymbolTagSupport{
			ValueSet: []SymbolTag{
				SymbolTagDeprecated,
			},
		},
		LabelSupport: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesDocumentSymbol
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDocumentSymbol{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesDocumentSymbol
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesDocumentSymbol{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesDocumentSymbol
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesFormatting(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesFormatting{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesFormatting
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesFormatting{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesFormatting
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesFormatting{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesFormatting
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesRangeFormatting(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesRangeFormatting{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesRangeFormatting
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesRangeFormatting{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesRangeFormatting
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesRangeFormatting{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesRangeFormatting
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesOnTypeFormatting(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesOnTypeFormatting{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesOnTypeFormatting
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesOnTypeFormatting{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesOnTypeFormatting
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesOnTypeFormatting{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesOnTypeFormatting
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesDeclaration(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"linkSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesDeclaration{
		DynamicRegistration: true,
		LinkSupport:         true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesDeclaration
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDeclaration{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesDeclaration
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesDeclaration{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesDeclaration
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesDefinition(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"linkSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesDefinition{
		DynamicRegistration: true,
		LinkSupport:         true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesDefinition
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDefinition{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesDefinition
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesDefinition{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesDefinition
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}
func testTextDocumentClientCapabilitiesTypeDefinition(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"linkSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesTypeDefinition{
		DynamicRegistration: true,
		LinkSupport:         true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesTypeDefinition
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesTypeDefinition{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesTypeDefinition
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesTypeDefinition{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesTypeDefinition
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesImplementation(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"linkSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesImplementation{
		DynamicRegistration: true,
		LinkSupport:         true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesImplementation
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesImplementation{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesImplementation
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesImplementation{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesImplementation
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesCodeAction(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}},"isPreferredSupport":true,"disabledSupport":true,"dataSupport":true,"resolveSupport":{"properties":["testProperties"]},"honorsChangeAnnotations":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesCodeAction{
		DynamicRegistration: true,
		CodeActionLiteralSupport: &TextDocumentClientCapabilitiesCodeActionLiteralSupport{
			CodeActionKind: &TextDocumentClientCapabilitiesCodeActionKind{
				ValueSet: []CodeActionKind{
					QuickFix,
					Refactor,
					RefactorExtract,
					RefactorRewrite,
					Source,
					SourceOrganizeImports,
				},
			},
		},
		IsPreferredSupport: true,
		DisabledSupport:    true,
		DataSupport:        true,
		ResolveSupport: &TextDocumentClientCapabilitiesCodeActionResolveSupport{
			Properties: []string{"testProperties"},
		},
		HonorsChangeAnnotations: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesCodeAction
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesCodeAction{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesCodeAction
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesCodeAction{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesCodeAction
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesCodeLens(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"tooltipSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesCodeLens{
		DynamicRegistration: true,
		TooltipSupport:      true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesCodeLens
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesCodeLens{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesCodeLens
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesCodeLens{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesCodeLens
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesDocumentLink(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"tooltipSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesDocumentLink{
		DynamicRegistration: true,
		TooltipSupport:      true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesDocumentLink
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDocumentLink{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesDocumentLink
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesDocumentLink{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesDocumentLink
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesColorProvider(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesColorProvider{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesColorProvider
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesColorProvider{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesColorProvider
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesColorProvider{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesColorProvider
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesRename(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"prepareSupport":true,"prepareSupportDefaultBehavior":1,"honorsChangeAnnotations":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesRename{
		DynamicRegistration:           true,
		PrepareSupport:                true,
		PrepareSupportDefaultBehavior: PrepareSupportDefaultBehaviorIdentifier,
		HonorsChangeAnnotations:       true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesRename
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesRename{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesRename
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesRename{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesRename
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesPublishDiagnostics(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"relatedInformation":true,"tagSupport":{"valueSet":[2,1]},"versionSupport":true,"codeDescriptionSupport":true,"dataSupport":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesPublishDiagnostics{
		RelatedInformation: true,
		TagSupport: &TextDocumentClientCapabilitiesPublishDiagnosticsTagSupport{
			ValueSet: []DiagnosticTag{
				DiagnosticTagDeprecated,
				DiagnosticTagUnnecessary,
			},
		},
		VersionSupport:         true,
		CodeDescriptionSupport: true,
		DataSupport:            true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesPublishDiagnostics
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesPublishDiagnostics{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesPublishDiagnostics
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesPublishDiagnostics{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesPublishDiagnostics
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesFoldingRange(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"rangeLimit":5,"lineFoldingOnly":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesFoldingRange{
		DynamicRegistration: true,
		RangeLimit:          uint32(5),
		LineFoldingOnly:     true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesFoldingRange
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesFoldingRange{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesFoldingRange
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesFoldingRange{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesFoldingRange
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"synchronization":{"didSave":true,"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true},"completion":{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":{"valueSet":[1]},"contextSupport":true},"hover":{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]},"signatureHelp":{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}},"references":{"dynamicRegistration":true},"documentHighlight":{"dynamicRegistration":true},"documentSymbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true},"formatting":{"dynamicRegistration":true},"rangeFormatting":{"dynamicRegistration":true},"onTypeFormatting":{"dynamicRegistration":true},"declaration":{"dynamicRegistration":true,"linkSupport":true},"definition":{"dynamicRegistration":true,"linkSupport":true},"typeDefinition":{"dynamicRegistration":true,"linkSupport":true},"implementation":{"dynamicRegistration":true,"linkSupport":true},"codeAction":{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"documentLink":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":true},"rename":{"dynamicRegistration":true,"prepareSupport":true},"publishDiagnostics":{"relatedInformation":true},"foldingRange":{"dynamicRegistration":true,"rangeLimit":5,"lineFoldingOnly":true},"selectionRange":{"dynamicRegistration":true},"linkedEditingRange":{"dynamicRegistration":true},"callHierarchy":{"dynamicRegistration":true},"semanticTokens":{"dynamicRegistration":true,"requests":{"range":true,"full":true},"tokenTypes":["test","tokenTypes"],"tokenModifiers":["test","tokenModifiers"],"formats":["relative"],"overlappingTokenSupport":true,"multilineTokenSupport":true},"moniker":{"dynamicRegistration":true}}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilities{
		Synchronization: &TextDocumentClientCapabilitiesSynchronization{
			DidSave:             true,
			DynamicRegistration: true,
			WillSave:            true,
			WillSaveWaitUntil:   true,
		},
		Completion: &TextDocumentClientCapabilitiesCompletion{
			DynamicRegistration: true,
			CompletionItem: &TextDocumentClientCapabilitiesCompletionItem{
				SnippetSupport:          true,
				CommitCharactersSupport: true,
				DocumentationFormat: []MarkupKind{
					PlainText,
					Markdown,
				},
				DeprecatedSupport: true,
				PreselectSupport:  true,
			},
			CompletionItemKind: &TextDocumentClientCapabilitiesCompletionItemKind{
				ValueSet: []CompletionItemKind{CompletionItemKindText},
			},
			ContextSupport: true,
		},
		Hover: &TextDocumentClientCapabilitiesHover{
			DynamicRegistration: true,
			ContentFormat: []MarkupKind{
				PlainText,
				Markdown,
			},
		},
		SignatureHelp: &TextDocumentClientCapabilitiesSignatureHelp{
			DynamicRegistration: true,
			SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
				DocumentationFormat: []MarkupKind{
					PlainText,
					Markdown,
				},
			},
		},
		References: &TextDocumentClientCapabilitiesReferences{
			DynamicRegistration: true,
		},
		DocumentHighlight: &TextDocumentClientCapabilitiesDocumentHighlight{
			DynamicRegistration: true,
		},
		DocumentSymbol: &TextDocumentClientCapabilitiesDocumentSymbol{
			DynamicRegistration: true,
			SymbolKind: &WorkspaceClientCapabilitiesSymbolKind{
				ValueSet: []SymbolKind{
					SymbolKindFile,
					SymbolKindModule,
					SymbolKindNamespace,
					SymbolKindPackage,
					SymbolKindClass,
					SymbolKindMethod,
				},
			},
			HierarchicalDocumentSymbolSupport: true,
		},
		Formatting: &TextDocumentClientCapabilitiesFormatting{
			DynamicRegistration: true,
		},
		RangeFormatting: &TextDocumentClientCapabilitiesRangeFormatting{
			DynamicRegistration: true,
		},
		OnTypeFormatting: &TextDocumentClientCapabilitiesOnTypeFormatting{
			DynamicRegistration: true,
		},
		Declaration: &TextDocumentClientCapabilitiesDeclaration{
			DynamicRegistration: true,
			LinkSupport:         true,
		},
		Definition: &TextDocumentClientCapabilitiesDefinition{
			DynamicRegistration: true,
			LinkSupport:         true,
		},
		TypeDefinition: &TextDocumentClientCapabilitiesTypeDefinition{
			DynamicRegistration: true,
			LinkSupport:         true,
		},
		Implementation: &TextDocumentClientCapabilitiesImplementation{
			DynamicRegistration: true,
			LinkSupport:         true,
		},
		CodeAction: &TextDocumentClientCapabilitiesCodeAction{
			DynamicRegistration: true,
			CodeActionLiteralSupport: &TextDocumentClientCapabilitiesCodeActionLiteralSupport{
				CodeActionKind: &TextDocumentClientCapabilitiesCodeActionKind{
					ValueSet: []CodeActionKind{
						QuickFix,
						Refactor,
						RefactorExtract,
						RefactorRewrite,
						Source,
						SourceOrganizeImports,
					},
				},
			},
		},
		CodeLens: &TextDocumentClientCapabilitiesCodeLens{
			DynamicRegistration: true,
		},
		DocumentLink: &TextDocumentClientCapabilitiesDocumentLink{
			DynamicRegistration: true,
		},
		ColorProvider: &TextDocumentClientCapabilitiesColorProvider{
			DynamicRegistration: true,
		},
		Rename: &TextDocumentClientCapabilitiesRename{
			DynamicRegistration: true,
			PrepareSupport:      true,
		},
		PublishDiagnostics: &TextDocumentClientCapabilitiesPublishDiagnostics{
			RelatedInformation: true,
		},
		FoldingRange: &TextDocumentClientCapabilitiesFoldingRange{
			DynamicRegistration: true,
			RangeLimit:          uint32(5),
			LineFoldingOnly:     true,
		},
		SelectionRange: &TextDocumentClientCapabilitiesSelectionRange{
			DynamicRegistration: true,
		},
		LinkedEditingRange: &TextDocumentClientCapabilitiesLinkedEditingRange{
			DynamicRegistration: true,
		},
		CallHierarchy: &TextDocumentClientCapabilitiesCallHierarchy{
			DynamicRegistration: true,
		},
		SemanticTokens: &TextDocumentClientCapabilitiesSemanticTokens{
			DynamicRegistration: true,
			Requests: WorkspaceClientCapabilitiesSemanticTokensRequests{
				Range: true,
				Full:  true,
			},
			TokenTypes:     []string{"test", "tokenTypes"},
			TokenModifiers: []string{"test", "tokenModifiers"},
			Formats: []TokenFormat{
				TokenFormatRelative,
			},
			OverlappingTokenSupport: true,
			MultilineTokenSupport:   true,
		},
		Moniker: &TextDocumentClientCapabilitiesMoniker{
			DynamicRegistration: true,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"workspace":{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"]},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true},"textDocument":{"synchronization":{"didSave":true,"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true},"completion":{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":{"valueSet":[1]},"contextSupport":true},"hover":{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]},"signatureHelp":{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}},"references":{"dynamicRegistration":true},"documentHighlight":{"dynamicRegistration":true},"documentSymbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true},"formatting":{"dynamicRegistration":true},"rangeFormatting":{"dynamicRegistration":true},"onTypeFormatting":{"dynamicRegistration":true},"declaration":{"dynamicRegistration":true,"linkSupport":true},"definition":{"dynamicRegistration":true,"linkSupport":true},"typeDefinition":{"dynamicRegistration":true,"linkSupport":true},"implementation":{"dynamicRegistration":true,"linkSupport":true},"codeAction":{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"documentLink":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":true},"rename":{"dynamicRegistration":true,"prepareSupport":true},"publishDiagnostics":{"relatedInformation":true},"foldingRange":{"dynamicRegistration":true,"rangeLimit":5,"lineFoldingOnly":true},"selectionRange":{"dynamicRegistration":true}},"window":{"workDoneProgress":true,"showMessage":{"messageActionItem":{"additionalPropertiesSupport":true}},"showDocument":{"support":true}},"general":{"regularExpressions":{"engine":"ECMAScript","version":"ES2020"},"markdown":{"parser":"marked","version":"1.1.0"}},"experimental":"testExperimental"}`
		wantNil = `{}`
	)
	wantType := ClientCapabilities{
		Workspace: &WorkspaceClientCapabilities{
			ApplyEdit: true,
			WorkspaceEdit: &WorkspaceClientCapabilitiesWorkspaceEdit{
				DocumentChanges:    true,
				FailureHandling:    "FailureHandling",
				ResourceOperations: []string{"ResourceOperations"},
			},
			DidChangeConfiguration: &WorkspaceClientCapabilitiesDidChangeConfiguration{
				DynamicRegistration: true,
			},
			DidChangeWatchedFiles: &WorkspaceClientCapabilitiesDidChangeWatchedFiles{
				DynamicRegistration: true,
			},
			Symbol: &WorkspaceClientCapabilitiesSymbol{
				DynamicRegistration: true,
				SymbolKind: &WorkspaceClientCapabilitiesSymbolKind{
					ValueSet: []SymbolKind{
						SymbolKindFile,
						SymbolKindModule,
						SymbolKindNamespace,
						SymbolKindPackage,
						SymbolKindClass,
						SymbolKindMethod,
					},
				},
			},
			ExecuteCommand: &WorkspaceClientCapabilitiesExecuteCommand{
				DynamicRegistration: true,
			},
			WorkspaceFolders: true,
			Configuration:    true,
		},
		TextDocument: &TextDocumentClientCapabilities{
			Synchronization: &TextDocumentClientCapabilitiesSynchronization{
				DidSave:             true,
				DynamicRegistration: true,
				WillSave:            true,
				WillSaveWaitUntil:   true,
			},
			Completion: &TextDocumentClientCapabilitiesCompletion{
				DynamicRegistration: true,
				CompletionItem: &TextDocumentClientCapabilitiesCompletionItem{
					SnippetSupport:          true,
					CommitCharactersSupport: true,
					DocumentationFormat: []MarkupKind{
						PlainText,
						Markdown,
					},
					DeprecatedSupport: true,
					PreselectSupport:  true,
				},
				CompletionItemKind: &TextDocumentClientCapabilitiesCompletionItemKind{
					ValueSet: []CompletionItemKind{CompletionItemKindText},
				},
				ContextSupport: true,
			},
			Hover: &TextDocumentClientCapabilitiesHover{
				DynamicRegistration: true,
				ContentFormat: []MarkupKind{
					PlainText,
					Markdown,
				},
			},
			SignatureHelp: &TextDocumentClientCapabilitiesSignatureHelp{
				DynamicRegistration: true,
				SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
					DocumentationFormat: []MarkupKind{
						PlainText,
						Markdown,
					},
				},
			},
			References: &TextDocumentClientCapabilitiesReferences{
				DynamicRegistration: true,
			},
			DocumentHighlight: &TextDocumentClientCapabilitiesDocumentHighlight{
				DynamicRegistration: true,
			},
			DocumentSymbol: &TextDocumentClientCapabilitiesDocumentSymbol{
				DynamicRegistration: true,
				SymbolKind: &WorkspaceClientCapabilitiesSymbolKind{
					ValueSet: []SymbolKind{
						SymbolKindFile,
						SymbolKindModule,
						SymbolKindNamespace,
						SymbolKindPackage,
						SymbolKindClass,
						SymbolKindMethod,
					},
				},
				HierarchicalDocumentSymbolSupport: true,
			},
			Formatting: &TextDocumentClientCapabilitiesFormatting{
				DynamicRegistration: true,
			},
			RangeFormatting: &TextDocumentClientCapabilitiesRangeFormatting{
				DynamicRegistration: true,
			},
			OnTypeFormatting: &TextDocumentClientCapabilitiesOnTypeFormatting{
				DynamicRegistration: true,
			},
			Declaration: &TextDocumentClientCapabilitiesDeclaration{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			Definition: &TextDocumentClientCapabilitiesDefinition{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			TypeDefinition: &TextDocumentClientCapabilitiesTypeDefinition{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			Implementation: &TextDocumentClientCapabilitiesImplementation{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			CodeAction: &TextDocumentClientCapabilitiesCodeAction{
				DynamicRegistration: true,
				CodeActionLiteralSupport: &TextDocumentClientCapabilitiesCodeActionLiteralSupport{
					CodeActionKind: &TextDocumentClientCapabilitiesCodeActionKind{
						ValueSet: []CodeActionKind{
							QuickFix,
							Refactor,
							RefactorExtract,
							RefactorRewrite,
							Source,
							SourceOrganizeImports,
						},
					},
				},
			},
			CodeLens: &TextDocumentClientCapabilitiesCodeLens{
				DynamicRegistration: true,
			},
			DocumentLink: &TextDocumentClientCapabilitiesDocumentLink{
				DynamicRegistration: true,
			},
			ColorProvider: &TextDocumentClientCapabilitiesColorProvider{
				DynamicRegistration: true,
			},
			Rename: &TextDocumentClientCapabilitiesRename{
				DynamicRegistration: true,
				PrepareSupport:      true,
			},
			PublishDiagnostics: &TextDocumentClientCapabilitiesPublishDiagnostics{
				RelatedInformation: true,
			},
			FoldingRange: &TextDocumentClientCapabilitiesFoldingRange{
				DynamicRegistration: true,
				RangeLimit:          uint32(5),
				LineFoldingOnly:     true,
			},
			SelectionRange: &TextDocumentClientCapabilitiesSelectionRange{
				DynamicRegistration: true,
			},
		},
		Window: &WindowClientCapabilities{
			WorkDoneProgress: true,
			ShowMessage: &ClientCapabilitiesShowMessageRequest{
				MessageActionItem: &ClientCapabilitiesShowMessageRequestMessageActionItem{
					AdditionalPropertiesSupport: true,
				},
			},
			ShowDocument: &ClientCapabilitiesShowDocument{
				Support: true,
			},
		},
		General: &GeneralClientCapabilities{
			RegularExpressions: &RegularExpressionsClientCapabilities{
				Engine:  "ECMAScript",
				Version: "ES2020",
			},
			Markdown: &MarkdownClientCapabilities{
				Parser:  "marked",
				Version: "1.1.0",
			},
		},
		Experimental: "testExperimental",
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          ClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          ClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             ClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             ClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}
