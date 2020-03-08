// Copyright 2020 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.lsp.dev/uri"
)

func testCompletionParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"},"position":{"line":25,"character":1},"context":{"triggerCharacter":".","triggerKind":1}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/test.go"},"position":{"line":2,"character":0},"context":{"triggerCharacter":".","triggerKind":1}}`
	)
	wantType := CompletionParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: uri.File("/path/to/test.go"),
			},
			Position: Position{
				Line:      25,
				Character: 1,
			},
		},
		Context: &CompletionContext{
			TriggerCharacter: ".",
			TriggerKind:      Invoked,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CompletionParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CompletionParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CompletionParams
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

func TestCompletionTriggerKind_String(t *testing.T) {
	tests := []struct {
		name string
		k    CompletionTriggerKind
		want string
	}{
		{
			name: "Invoked",
			k:    Invoked,
			want: "Invoked",
		},
		{
			name: "TriggerCharacter",
			k:    TriggerCharacter,
			want: "TriggerCharacter",
		},
		{
			name: "TriggerForIncompleteCompletions",
			k:    TriggerForIncompleteCompletions,
			want: "TriggerForIncompleteCompletions",
		},
		{
			name: "Unknown",
			k:    CompletionTriggerKind(0),
			want: "0",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.k.String(); got != tt.want {
				t.Errorf("CompletionTriggerKind.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testCompletionContext(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"triggerCharacter":".","triggerKind":1}`
		wantInvalid = `{"triggerCharacter":" ","triggerKind":0}`
	)
	wantType := CompletionContext{
		TriggerCharacter: ".",
		TriggerKind:      Invoked,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CompletionContext
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CompletionContext
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CompletionContext
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

func testCompletionList(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"isIncomplete":true,"items":[{"detail":"string","documentation":"Detail a human-readable string with additional information about this item, like type or symbol information.","filterText":"Detail","insertTextFormat":2,"kind":5,"label":"Detail","preselect":true,"sortText":"00000","textEdit":{"range":{"start":{"line":255,"character":4},"end":{"line":255,"character":10}},"newText":"Detail: ${1:},"}}]}`
		wantInvalid = `{"isIncomplete":false,"items":[]}`
	)
	wantType := CompletionList{
		IsIncomplete: true,
		Items: []CompletionItem{
			{
				AdditionalTextEdits: nil,
				Command:             nil,
				CommitCharacters:    nil,
				Deprecated:          false,
				Detail:              "string",
				Documentation:       "Detail a human-readable string with additional information about this item, like type or symbol information.",
				FilterText:          "Detail",
				InsertText:          "",
				InsertTextFormat:    TextFormatSnippet,
				Kind:                FieldCompletion,
				Label:               "Detail",
				Preselect:           true,
				SortText:            "00000",
				TextEdit: &TextEdit{
					Range: Range{
						Start: Position{
							Line:      255,
							Character: 4,
						},
						End: Position{
							Line:      255,
							Character: 10,
						},
					},
					NewText: "Detail: ${1:},",
				},
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CompletionList
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CompletionList
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CompletionList
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

func TestInsertTextFormat_String(t *testing.T) {
	tests := []struct {
		name string
		k    InsertTextFormat
		want string
	}{
		{
			name: "PlainText",
			k:    TextFormatPlainText,
			want: "PlainText",
		},
		{
			name: "Snippet",
			k:    TextFormatSnippet,
			want: "Snippet",
		},
		{
			name: "Unknown",
			k:    InsertTextFormat(0),
			want: "0",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.k.String(); got != tt.want {
				t.Errorf("InsertTextFormat.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testCompletionItem(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"additionalTextEdits":[{"range":{"start":{"line":255,"character":4},"end":{"line":255,"character":10}},"newText":"Detail: ${1:},"}],"command":{"title":"exec echo","command":"echo","arguments":["hello"]},"commitCharacters":["a"],"data":"testData","deprecated":true,"detail":"string","documentation":"Detail a human-readable string with additional information about this item, like type or symbol information.","filterText":"Detail","insertText":"testInsert","insertTextFormat":2,"kind":5,"label":"Detail","preselect":true,"sortText":"00000","textEdit":{"range":{"start":{"line":255,"character":4},"end":{"line":255,"character":10}},"newText":"Detail: ${1:},"}}`
		wantNilAll  = `{"label":"Detail"}`
		wantInvalid = `{"items":[]}`
	)
	wantType := CompletionItem{
		AdditionalTextEdits: []TextEdit{
			{
				Range: Range{
					Start: Position{
						Line:      255,
						Character: 4,
					},
					End: Position{
						Line:      255,
						Character: 10,
					},
				},
				NewText: "Detail: ${1:},",
			},
		},
		Command: &Command{
			Title:     "exec echo",
			Command:   "echo",
			Arguments: []interface{}{"hello"},
		},
		CommitCharacters: []string{"a"},
		Data:             "testData",
		Deprecated:       true,
		Detail:           "string",
		Documentation:    "Detail a human-readable string with additional information about this item, like type or symbol information.",
		FilterText:       "Detail",
		InsertText:       "testInsert",
		InsertTextFormat: TextFormatSnippet,
		Kind:             FieldCompletion,
		Label:            "Detail",
		Preselect:        true,
		SortText:         "00000",
		TextEdit: &TextEdit{
			Range: Range{
				Start: Position{
					Line:      255,
					Character: 4,
				},
				End: Position{
					Line:      255,
					Character: 10,
				},
			},
			NewText: "Detail: ${1:},",
		},
	}
	wantTypeNilAll := CompletionItem{
		Label: "Detail",
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CompletionItem
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
				field:          wantTypeNilAll,
				want:           wantNilAll,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CompletionItem
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
				field:            wantNilAll,
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CompletionItem
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

func TestCompletionItemKind_String(t *testing.T) {
	tests := []struct {
		name string
		k    CompletionItemKind
		want string
	}{
		{
			name: "Text",
			k:    TextCompletion,
			want: "Text",
		},
		{
			name: "Method",
			k:    MethodCompletion,
			want: "Method",
		},
		{
			name: "Function",
			k:    FunctionCompletion,
			want: "Function",
		},
		{
			name: "Constructor",
			k:    ConstructorCompletion,
			want: "Constructor",
		},
		{
			name: "Field",
			k:    FieldCompletion,
			want: "Field",
		},
		{
			name: "Variable",
			k:    VariableCompletion,
			want: "Variable",
		},
		{
			name: "Class",
			k:    ClassCompletion,
			want: "Class",
		},
		{
			name: "Interface",
			k:    InterfaceCompletion,
			want: "Interface",
		},
		{
			name: "Module",
			k:    ModuleCompletion,
			want: "Module",
		},
		{
			name: "Property",
			k:    PropertyCompletion,
			want: "Property",
		},
		{
			name: "Unit",
			k:    UnitCompletion,
			want: "Unit",
		},
		{
			name: "Value",
			k:    ValueCompletion,
			want: "Value",
		},
		{
			name: "Enum",
			k:    EnumCompletion,
			want: "Enum",
		},
		{
			name: "Keyword",
			k:    KeywordCompletion,
			want: "Keyword",
		},
		{
			name: "Snippet",
			k:    SnippetCompletion,
			want: "Snippet",
		},
		{
			name: "Color",
			k:    ColorCompletion,
			want: "Color",
		},
		{
			name: "File",
			k:    FileCompletion,
			want: "File",
		},
		{
			name: "Reference",
			k:    ReferenceCompletion,
			want: "Reference",
		},
		{
			name: "Folder",
			k:    FolderCompletion,
			want: "Folder",
		},
		{
			name: "EnumMember",
			k:    EnumMemberCompletion,
			want: "EnumMember",
		},
		{
			name: "Constant",
			k:    ConstantCompletion,
			want: "Constant",
		},
		{
			name: "Struct",
			k:    StructCompletion,
			want: "Struct",
		},
		{
			name: "Event",
			k:    EventCompletion,
			want: "Event",
		},
		{
			name: "Operator",
			k:    OperatorCompletion,
			want: "Operator",
		},
		{
			name: "TypeParameter",
			k:    TypeParameterCompletion,
			want: "TypeParameter",
		},
		{
			name: "Unknown",
			k:    CompletionItemKind(0),
			want: "0",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.k.String(); got != tt.want {
				t.Errorf("CompletionItemKind.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testCompletionRegistrationOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"}],"triggerCharacters":["."],"resolveProvider":true}`
		wantInvalid = `{"documentSelector":[{"language":"typescript","scheme":"file","pattern":"*.{ts,js}"}],"triggerCharacters":[" "],"resolveProvider":true}`
	)
	wantType := CompletionRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
			},
		},
		TriggerCharacters: []string{"."},
		ResolveProvider:   true,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CompletionRegistrationOptions
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CompletionRegistrationOptions
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CompletionRegistrationOptions
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

func testHover(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"contents":{"kind":"markdown","value":"example value"},"range":{"start":{"line":255,"character":4},"end":{"line":255,"character":10}}}`
		wantInvalid = `{"contents":{"kind":"markdown","value":"example value"},"range":{"start":{"line":25,"character":2},"end":{"line":25,"character":5}}}`
	)
	wantType := Hover{
		Contents: MarkupContent{
			Kind:  Markdown,
			Value: "example value",
		},
		Range: Range{
			Start: Position{
				Line:      255,
				Character: 4,
			},
			End: Position{
				Line:      255,
				Character: 10,
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          Hover
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             Hover
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got Hover
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

func testSignatureHelp(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"signatures":[{"documentationFormat":["markdown"],"parameterInformation":{"label":"test label","documentation":"test documentation"}}],"activeParameter":10,"activeSignature":5}`
		wantNilAll  = `{"signatures":[]}`
		wantInvalid = `{"signatures":[{"documentationFormat":["markdown"],"parameterInformation":{"label":"test label","documentation":"test documentation"}}],"activeParameter":1,"activeSignature":0}`
	)
	wantType := SignatureHelp{
		Signatures: []SignatureInformation{
			{
				DocumentationFormat: []MarkupKind{
					Markdown,
				},
				ParameterInformation: &ParameterInformation{
					Label:         "test label",
					Documentation: "test documentation",
				},
			},
		},
		ActiveParameter: 10,
		ActiveSignature: 5,
	}
	wantTypeNilAll := SignatureHelp{
		Signatures: []SignatureInformation{},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          SignatureHelp
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
				field:          wantTypeNilAll,
				want:           wantNilAll,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             SignatureHelp
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
				field:            wantNilAll,
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got SignatureHelp
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

func testSignatureInformation(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"documentationFormat":["markdown"],"parameterInformation":{"label":"test label","documentation":"test documentation"}}`
		wantInvalid = `{"documentationFormat":["markdown","plaintext"],"parameterInformation":{"label":"test label","documentation":"test documentation"}}`
	)
	wantType := SignatureInformation{
		DocumentationFormat: []MarkupKind{
			Markdown,
		},
		ParameterInformation: &ParameterInformation{
			Label:         "test label",
			Documentation: "test documentation",
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          SignatureInformation
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             SignatureInformation
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got SignatureInformation
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

func testParameterInformation(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"label":"test label","documentation":"test documentation"}`
		wantInvalid = `{"label":"invalid","documentation":"invalid"}`
	)
	wantType := ParameterInformation{
		Label:         "test label",
		Documentation: "test documentation",
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          ParameterInformation
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             ParameterInformation
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ParameterInformation
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

func testSignatureHelpRegistrationOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"}],"triggerCharacters":["."]}`
		wantInvalid = `{"documentSelector":[{"language":"typescript","scheme":"file","pattern":"*.{ts,js}"}],"triggerCharacters":[" "]}`
	)
	wantType := SignatureHelpRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
			},
		},
		TriggerCharacters: []string{"."},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          SignatureHelpRegistrationOptions
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             SignatureHelpRegistrationOptions
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got SignatureHelpRegistrationOptions
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

func testReferenceParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"},"position":{"line":25,"character":1},"context":{"includeDeclaration":true}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/test.go"},"position":{"line":2,"character":0},"context":{"includeDeclaration":false}}`
	)
	wantType := ReferenceParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: uri.File("/path/to/test.go"),
			},
			Position: Position{
				Line:      25,
				Character: 1,
			},
		},
		Context: ReferenceContext{
			IncludeDeclaration: true,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          ReferenceParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             ReferenceParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ReferenceParams
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

func testReferenceContext(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"includeDeclaration":true}`
		wantInvalid = `{"includeDeclaration":false}`
	)
	wantType := ReferenceContext{
		IncludeDeclaration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          ReferenceContext
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             ReferenceContext
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ReferenceContext
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

func testDocumentHighlight(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"range":{"start":{"line":255,"character":4},"end":{"line":255,"character":10}},"kind":1}`
		wantInvalid = `{"range":{"start":{"line":25,"character":2},"end":{"line":25,"character":5}},"kind":1}`
	)
	wantType := DocumentHighlight{
		Range: Range{
			Start: Position{
				Line:      255,
				Character: 4,
			},
			End: Position{
				Line:      255,
				Character: 10,
			},
		},
		Kind: Text,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentHighlight
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentHighlight
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentHighlight
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

func TestDocumentHighlightKind_String(t *testing.T) {
	tests := []struct {
		name string
		k    DocumentHighlightKind
		want string
	}{
		{
			name: "Text",
			k:    Text,
			want: "Text",
		},
		{
			name: "Read",
			k:    Read,
			want: "Read",
		},
		{
			name: "Write",
			k:    Write,
			want: "Write",
		},
		{
			name: "Unknown",
			k:    DocumentHighlightKind(0),
			want: "0",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.k.String(); got != tt.want {
				t.Errorf("DocumentHighlightKind.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testDocumentSymbolParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/nottest.go"}}`
	)
	wantType := DocumentSymbolParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentSymbolParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentSymbolParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentSymbolParams
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

func TestSymbolKind_String(t *testing.T) {
	tests := []struct {
		name string
		k    SymbolKind
		want string
	}{
		{
			name: "File",
			k:    FileSymbol,
			want: "File",
		},
		{
			name: "Module",
			k:    ModuleSymbol,
			want: "Module",
		},
		{
			name: "Namespace",
			k:    NamespaceSymbol,
			want: "Namespace",
		},
		{
			name: "Package",
			k:    PackageSymbol,
			want: "Package",
		},
		{
			name: "Class",
			k:    ClassSymbol,
			want: "Class",
		},
		{
			name: "Method",
			k:    MethodSymbol,
			want: "Method",
		},
		{
			name: "Property",
			k:    PropertySymbol,
			want: "Property",
		},
		{
			name: "Field",
			k:    FieldSymbol,
			want: "Field",
		},
		{
			name: "Constructor",
			k:    ConstructorSymbol,
			want: "Constructor",
		},
		{
			name: "Enum",
			k:    EnumSymbol,
			want: "Enum",
		},
		{
			name: "Interface",
			k:    InterfaceSymbol,
			want: "Interface",
		},
		{
			name: "Function",
			k:    FunctionSymbol,
			want: "Function",
		},
		{
			name: "Variable",
			k:    VariableSymbol,
			want: "Variable",
		},
		{
			name: "Constant",
			k:    ConstantSymbol,
			want: "Constant",
		},
		{
			name: "String",
			k:    StringSymbol,
			want: "String",
		},
		{
			name: "Number",
			k:    NumberSymbol,
			want: "Number",
		},
		{
			name: "Boolean",
			k:    BooleanSymbol,
			want: "Boolean",
		},
		{
			name: "Array",
			k:    ArraySymbol,
			want: "Array",
		},
		{
			name: "Object",
			k:    ObjectSymbol,
			want: "Object",
		},
		{
			name: "Key",
			k:    KeySymbol,
			want: "Key",
		},
		{
			name: "Null",
			k:    NullSymbol,
			want: "Null",
		},
		{
			name: "EnumMember",
			k:    EnumMemberSymbol,
			want: "EnumMember",
		},
		{
			name: "Struct",
			k:    StructSymbol,
			want: "Struct",
		},
		{
			name: "Event",
			k:    EventSymbol,
			want: "Event",
		},
		{
			name: "Operator",
			k:    OperatorSymbol,
			want: "Operator",
		},
		{
			name: "TypeParameter",
			k:    TypeParameterSymbol,
			want: "TypeParameter",
		},
		{
			name: "Unknown",
			k:    SymbolKind(0),
			want: "0",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.k.String(); got != tt.want {
				t.Errorf("SymbolKind.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testDocumentSymbol(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"name":"test symbol","detail":"test detail","kind":1,"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":6}},"selectionRange":{"start":{"line":25,"character":3},"end":{"line":26,"character":10}},"children":[{"name":"child symbol","detail":"child detail","kind":11,"deprecated":true,"range":{"start":{"line":255,"character":4},"end":{"line":255,"character":10}},"selectionRange":{"start":{"line":255,"character":5},"end":{"line":255,"character":8}}}]}`
		wantInvalid = `{"name":"invalid symbol","detail":"invalid detail","kind":1,"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"selectionRange":{"start":{"line":2,"character":5},"end":{"line":3,"character":1}},"children":[{"name":"invalid child symbol","kind":1,"detail":"invalid child detail","range":{"start":{"line":255,"character":4},"end":{"line":255,"character":10}},"selectionRange":{"start":{"line":255,"character":5},"end":{"line":255,"character":8}}}]}`
	)
	wantType := DocumentSymbol{
		Name:       "test symbol",
		Detail:     "test detail",
		Kind:       FileSymbol,
		Deprecated: false,
		Range: Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      27,
				Character: 6,
			},
		},
		SelectionRange: Range{
			Start: Position{
				Line:      25,
				Character: 3,
			},
			End: Position{
				Line:      26,
				Character: 10,
			},
		},
		Children: []DocumentSymbol{
			{
				Name:       "child symbol",
				Detail:     "child detail",
				Kind:       InterfaceSymbol,
				Deprecated: true,
				Range: Range{
					Start: Position{
						Line:      255,
						Character: 4,
					},
					End: Position{
						Line:      255,
						Character: 10,
					},
				},
				SelectionRange: Range{
					Start: Position{
						Line:      255,
						Character: 5,
					},
					End: Position{
						Line:      255,
						Character: 8,
					},
				},
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentSymbol
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentSymbol
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentSymbol
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

func testSymbolInformation(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"name":"test symbol","kind":1,"deprecated":true,"location":{"uri":"file:///path/to/test.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"containerName":"testContainerName"}`
		wantNilAll  = `{"name":"test symbol","kind":1,"location":{"uri":"file:///path/to/test.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}}`
		wantInvalid = `{"name":"invalid symbol","kind":1,"deprecated":false,"location":{"uri":"file:///path/to/test_test.go","range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}},"containerName":"invalidContainerName"}`
	)
	wantType := SymbolInformation{
		Name:       "test symbol",
		Kind:       1,
		Deprecated: true,
		Location: Location{
			URI: uri.File("/path/to/test.go"),
			Range: Range{
				Start: Position{
					Line:      25,
					Character: 1,
				},
				End: Position{
					Line:      27,
					Character: 3,
				},
			},
		},
		ContainerName: "testContainerName",
	}
	wantTypeNilAll := SymbolInformation{
		Name:       "test symbol",
		Kind:       1,
		Deprecated: false,
		Location: Location{
			URI: uri.File("/path/to/test.go"),
			Range: Range{
				Start: Position{
					Line:      25,
					Character: 1,
				},
				End: Position{
					Line:      27,
					Character: 3,
				},
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          SymbolInformation
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
				field:          wantTypeNilAll,
				want:           wantNilAll,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             SymbolInformation
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
				field:            wantNilAll,
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got SymbolInformation
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

func testCodeActionParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"},"context":{"diagnostics":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/test.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"test.go"}]}],"only":["quickfix"]},"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":6}}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/test.go"},"context":{"diagnostics":[{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/test.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"test.go"}]}],"only":["quickfix"]},"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}}`
	)
	wantType := CodeActionParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
		Context: CodeActionContext{
			Diagnostics: []Diagnostic{
				{
					Range: Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					Severity: SeverityError,
					Code:     "foo/bar",
					Source:   "test foo bar",
					Message:  "foo bar",
					RelatedInformation: []DiagnosticRelatedInformation{
						{
							Location: Location{
								URI: uri.File("/path/to/test.go"),
								Range: Range{
									Start: Position{
										Line:      25,
										Character: 1,
									},
									End: Position{
										Line:      27,
										Character: 3,
									},
								},
							},
							Message: "test.go",
						},
					},
				},
			},
			Only: []CodeActionKind{
				QuickFix,
			},
		},
		Range: Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      27,
				Character: 6,
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CodeActionParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CodeActionParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CodeActionParams
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

func TestCodeActionKind_String(t *testing.T) {
	tests := []struct {
		name string
		k    CodeActionKind
		want string
	}{
		{
			name: "QuickFix",
			k:    QuickFix,
			want: "quickfix",
		},
		{
			name: "Refactor",
			k:    Refactor,
			want: "refactor",
		},
		{
			name: "RefactorExtract",
			k:    RefactorExtract,
			want: "refactor.extract",
		},
		{
			name: "RefactorInline",
			k:    RefactorInline,
			want: "refactor.inline",
		},
		{
			name: "RefactorRewrite",
			k:    RefactorRewrite,
			want: "refactor.rewrite",
		},
		{
			name: "Source",
			k:    Source,
			want: "source",
		},
		{
			name: "SourceOrganizeImports",
			k:    SourceOrganizeImports,
			want: "source.organizeImports",
		},
		{
			name: "Unknown",
			k:    CodeActionKind(""),
			want: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.k; got != CodeActionKind(tt.want) {
				t.Errorf("CodeActionKind = %v, want %v", got, tt.want)
			}
		})
	}
}

func testCodeActionContext(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"diagnostics":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/test.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"test.go"}]}],"only":["quickfix"]}`
		wantInvalid = `{"diagnostics":[{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/test.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"test.go"}]}],"only":["quickfix"]}`
	)
	wantType := CodeActionContext{
		Diagnostics: []Diagnostic{
			{
				Range: Range{
					Start: Position{
						Line:      25,
						Character: 1,
					},
					End: Position{
						Line:      27,
						Character: 3,
					},
				},
				Severity: SeverityError,
				Code:     "foo/bar",
				Source:   "test foo bar",
				Message:  "foo bar",
				RelatedInformation: []DiagnosticRelatedInformation{
					{
						Location: Location{
							URI: uri.File("/path/to/test.go"),
							Range: Range{
								Start: Position{
									Line:      25,
									Character: 1,
								},
								End: Position{
									Line:      27,
									Character: 3,
								},
							},
						},
						Message: "test.go",
					},
				},
			},
		},
		Only: []CodeActionKind{
			QuickFix,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CodeActionContext
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CodeActionContext
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CodeActionContext
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

func testCodeAction(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"title":"Refactoring","kind":"refactor.rewrite","diagnostics":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/test.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"test.go"}]}],"edit":{"changes":{"file:///path/to/test.go":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]},"documentChanges":[{"textDocument":{"uri":"file:///path/to/test.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}]},"command":{"title":"rewrite","command":"rewriter","arguments":["-w"]}}`
		wantInvalid = `{"title":"Refactoring","kind":"refactor","diagnostics":[{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/test.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"test.go"}]}],"edit":{"changes":{"file:///path/to/test.go":[{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"newText":"foo bar"}]},"documentChanges":[{"textDocument":{"uri":"file:///path/to/test.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}]},"command":{"title":"rewrite","command":"rewriter","arguments":["-w"]}}`
	)
	wantType := CodeAction{
		Title: "Refactoring",
		Kind:  RefactorRewrite,
		Diagnostics: []Diagnostic{
			{
				Range: Range{
					Start: Position{
						Line:      25,
						Character: 1,
					},
					End: Position{
						Line:      27,
						Character: 3,
					},
				},
				Severity: SeverityError,
				Code:     "foo/bar",
				Source:   "test foo bar",
				Message:  "foo bar",
				RelatedInformation: []DiagnosticRelatedInformation{
					{
						Location: Location{
							URI: uri.File("/path/to/test.go"),
							Range: Range{
								Start: Position{
									Line:      25,
									Character: 1,
								},
								End: Position{
									Line:      27,
									Character: 3,
								},
							},
						},
						Message: "test.go",
					},
				},
			},
		},
		Edit: &WorkspaceEdit{
			Changes: map[uri.URI][]TextEdit{
				uri.File("/path/to/test.go"): {
					{
						Range: Range{
							Start: Position{
								Line:      25,
								Character: 1,
							},
							End: Position{
								Line:      27,
								Character: 3,
							},
						},
						NewText: "foo bar",
					},
				},
			},
			DocumentChanges: []TextDocumentEdit{
				{
					TextDocument: VersionedTextDocumentIdentifier{
						TextDocumentIdentifier: TextDocumentIdentifier{
							URI: uri.File("/path/to/test.go"),
						},
						Version: Uint64Ptr(10),
					},
					Edits: []TextEdit{
						{
							Range: Range{
								Start: Position{
									Line:      25,
									Character: 1,
								},
								End: Position{
									Line:      27,
									Character: 3,
								},
							},
							NewText: "foo bar",
						},
					},
				},
			},
		},
		Command: &Command{
			Title:     "rewrite",
			Command:   "rewriter",
			Arguments: []interface{}{"-w"},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CodeAction
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CodeAction
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CodeAction
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

func testCodeActionRegistrationOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"},{"language":"cpp","scheme":"untitled","pattern":"*.{cpp,hpp}"}],"codeActionKinds":["quickfix","refactor"]}`
		wantInvalid = `{"documentSelector":[{"language":"typescript","scheme":"file","pattern":"*.{ts,js}"},{"language":"c","scheme":"untitled","pattern":"*.{c,h}"}],"codeActionKinds":["quickfix","refactor"]}`
	)
	wantType := CodeActionRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
				{
					Language: "cpp",
					Scheme:   "untitled",
					Pattern:  "*.{cpp,hpp}",
				},
			},
		},
		CodeActionOptions: CodeActionOptions{
			CodeActionKinds: []CodeActionKind{
				QuickFix,
				Refactor,
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CodeActionRegistrationOptions
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CodeActionRegistrationOptions
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CodeActionRegistrationOptions
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

func testCodeLensParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/invalid.go"}}`
	)
	wantType := CodeLensParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CodeLensParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CodeLensParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CodeLensParams
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

func testCodeLens(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"command":{"title":"rewrite","command":"rewriter","arguments":["-w"]},"data":"testData"}`
		wantNilAll  = `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}`
		wantInvalid = `{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"command":{"title":"rewrite","command":"rewriter","arguments":["-w"]},"data":"invalidData"}`
	)
	wantType := CodeLens{
		Range: Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      27,
				Character: 3,
			},
		},
		Command: &Command{
			Title:     "rewrite",
			Command:   "rewriter",
			Arguments: []interface{}{"-w"},
		},
		Data: "testData",
	}
	wantTypeNilAll := CodeLens{
		Range: Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      27,
				Character: 3,
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CodeLens
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
				field:          wantTypeNilAll,
				want:           wantNilAll,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CodeLens
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
				field:            wantNilAll,
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CodeLens
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

func testCodeLensRegistrationOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"},{"language":"cpp","scheme":"untitled","pattern":"*.{cpp,hpp}"}],"resolveProvider":true}`
		wantNilAll  = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"},{"language":"cpp","scheme":"untitled","pattern":"*.{cpp,hpp}"}]}`
		wantInvalid = `{"documentSelector":[{"language":"typescript","scheme":"file","pattern":"*.{ts,js}"},{"language":"c","scheme":"untitled","pattern":"*.{c,h}"}],"resolveProvider":false}`
	)
	wantType := CodeLensRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
				{
					Language: "cpp",
					Scheme:   "untitled",
					Pattern:  "*.{cpp,hpp}",
				},
			},
		},
		ResolveProvider: true,
	}
	wantTypeNilAll := CodeLensRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
				{
					Language: "cpp",
					Scheme:   "untitled",
					Pattern:  "*.{cpp,hpp}",
				},
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          CodeLensRegistrationOptions
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
				field:          wantTypeNilAll,
				want:           wantNilAll,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             CodeLensRegistrationOptions
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
				field:            wantNilAll,
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CodeLensRegistrationOptions
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

func testDocumentLinkParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantNilAll  = `{"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/invalid.go"}}`
	)
	wantType := DocumentLinkParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentLinkParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentLinkParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentLinkParams
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

func testDocumentLink(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"target":"file:///path/to/test.go","data":"testData"}`
		wantNilAll  = `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}`
		wantInvalid = `{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"target":"file:///path/to/test.go","data":"testData"}`
	)
	wantType := DocumentLink{
		Range: Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      27,
				Character: 3,
			},
		},
		Target: uri.File("/path/to/test.go"),
		Data:   "testData",
	}
	wantTypeNilAll := DocumentLink{
		Range: Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      27,
				Character: 3,
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentLink
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
				field:          wantTypeNilAll,
				want:           wantNilAll,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentLink
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
				field:            wantNilAll,
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentLink
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

func testDocumentColorParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/invalid.go"}}`
	)
	wantType := DocumentColorParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentColorParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentColorParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentColorParams
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

func testColorInformation(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"color":{"alpha":1,"blue":0.2,"green":0.3,"red":0.4}}`
		wantInvalid = `{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"color":{"alpha":0,"blue":0.4,"green":0.3,"red":0.2}}`
	)
	wantType := ColorInformation{
		Range: Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      27,
				Character: 3,
			},
		},
		Color: Color{
			Alpha: 1,
			Blue:  0.2,
			Green: 0.3,
			Red:   0.4,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          ColorInformation
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             ColorInformation
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ColorInformation
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

func testColor(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"alpha":1,"blue":0.2,"green":0.3,"red":0.4}`
		wantInvalid = `{"alpha":0,"blue":0.4,"green":0.3,"red":0.2}`
	)
	wantType := Color{
		Alpha: 1,
		Blue:  0.2,
		Green: 0.3,
		Red:   0.4,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          Color
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             Color
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got Color
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

func testColorPresentationParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"},"color":{"alpha":1,"blue":0.2,"green":0.3,"red":0.4},"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/test.go"},"color":{"alpha":0,"blue":0.4,"green":0.3,"red":0.2},"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}}`
	)
	wantType := ColorPresentationParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
		Color: Color{
			Alpha: 1,
			Blue:  0.2,
			Green: 0.3,
			Red:   0.4,
		},
		Range: Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      27,
				Character: 3,
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          ColorPresentationParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             ColorPresentationParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ColorPresentationParams
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

func testColorPresentation(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"label":"testLabel","textEdit":{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"},"additionalTextEdits":[{"range":{"start":{"line":100,"character":10},"end":{"line":102,"character":15}},"newText":"baz qux"}]}`
		wantNilAll  = `{"label":"testLabel"}`
		wantInvalid = `{"label":"invalidLabel","textEdit":{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"newText":"quux quuz"},"additionalTextEdits":[{"range":{"start":{"line":105,"character":15},"end":{"line":107,"character":20}},"newText":"corge grault"}]}`
	)
	wantType := ColorPresentation{
		Label: "testLabel",
		TextEdit: &TextEdit{
			Range: Range{
				Start: Position{
					Line:      25,
					Character: 1,
				},
				End: Position{
					Line:      27,
					Character: 3,
				},
			},
			NewText: "foo bar",
		},
		AdditionalTextEdits: []TextEdit{
			{
				Range: Range{
					Start: Position{
						Line:      100,
						Character: 10,
					},
					End: Position{
						Line:      102,
						Character: 15,
					},
				},
				NewText: "baz qux",
			},
		},
	}
	wantTypeNilAll := ColorPresentation{
		Label: "testLabel",
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          ColorPresentation
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
				field:          wantTypeNilAll,
				want:           wantNilAll,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             ColorPresentation
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
				field:            wantNilAll,
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ColorPresentation
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

func testDocumentFormattingParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"options":{"insertSpaces":true,"tabSize":4},"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantInvalid = `{"options":{"insertSpaces":false,"tabSize":2},"textDocument":{"uri":"file:///path/to/invalid.go"}}`
	)
	wantType := DocumentFormattingParams{
		Options: FormattingOptions{
			InsertSpaces: true,
			TabSize:      4,
		},
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentFormattingParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentFormattingParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentFormattingParams
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

func testFormattingOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"insertSpaces":true,"tabSize":4}`
		wantInvalid = `{"insertSpaces":false,"tabSize":2}`
	)
	wantType := FormattingOptions{
		InsertSpaces: true,
		TabSize:      4,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          FormattingOptions
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             FormattingOptions
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got FormattingOptions
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

func testDocumentRangeFormattingParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"},"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"options":{"insertSpaces":true,"tabSize":4}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/invalid.go"},"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"options":{"insertSpaces":false,"tabSize":2}}`
	)
	wantType := DocumentRangeFormattingParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
		Range: Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      27,
				Character: 3,
			},
		},
		Options: FormattingOptions{
			InsertSpaces: true,
			TabSize:      4,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentRangeFormattingParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentRangeFormattingParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentRangeFormattingParams
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

func testDocumentOnTypeFormattingParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"},"position":{"line":25,"character":1},"ch":"character","options":{"insertSpaces":true,"tabSize":4}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/invalid.go"},"position":{"line":2,"character":1},"ch":"invalidChar","options":{"insertSpaces":false,"tabSize":2}}`
	)
	wantType := DocumentOnTypeFormattingParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
		Position: Position{
			Line:      25,
			Character: 1,
		},
		Ch: "character",
		Options: FormattingOptions{
			InsertSpaces: true,
			TabSize:      4,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentOnTypeFormattingParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentOnTypeFormattingParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentOnTypeFormattingParams
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

func testDocumentOnTypeFormattingRegistrationOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"},{"language":"cpp","scheme":"untitled","pattern":"*.{cpp,hpp}"}],"firstTriggerCharacter":"}","moreTriggerCharacter":[".","{"]}`
		wantInvalid = `{"documentSelector":[{"language":"typescript","scheme":"file","pattern":"*.{ts,js}"},{"language":"c","scheme":"untitled","pattern":"*.{c,h}"}],"firstTriggerCharacter":"{","moreTriggerCharacter":[" ","("]}`
	)
	wantType := DocumentOnTypeFormattingRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
				{
					Language: "cpp",
					Scheme:   "untitled",
					Pattern:  "*.{cpp,hpp}",
				},
			},
		},
		FirstTriggerCharacter: "}",
		MoreTriggerCharacter:  []string{".", "{"},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DocumentOnTypeFormattingRegistrationOptions
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             DocumentOnTypeFormattingRegistrationOptions
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentOnTypeFormattingRegistrationOptions
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

func testRenameParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"},"position":{"line":25,"character":1},"newName":"newNameSymbol"}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/invalid.go"},"position":{"line":2,"character":1},"newName":"invalidSymbol"}`
	)
	wantType := RenameParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
		Position: Position{
			Line:      25,
			Character: 1,
		},
		NewName: "newNameSymbol",
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          RenameParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             RenameParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got RenameParams
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

func testRenameRegistrationOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"},{"language":"cpp","scheme":"untitled","pattern":"*.{cpp,hpp}"}],"prepareProvider":true}`
		wantNilAll  = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"},{"language":"cpp","scheme":"untitled","pattern":"*.{cpp,hpp}"}]}`
		wantInvalid = `{"documentSelector":[{"language":"typescript","scheme":"file","pattern":"*.{ts,js}"},{"language":"c","scheme":"untitled","pattern":"*.{c,h}"}],"prepareProvider":false}`
	)
	wantType := RenameRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
				{
					Language: "cpp",
					Scheme:   "untitled",
					Pattern:  "*.{cpp,hpp}",
				},
			},
		},
		PrepareProvider: true,
	}
	wantTypeNilAll := RenameRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
				{
					Language: "cpp",
					Scheme:   "untitled",
					Pattern:  "*.{cpp,hpp}",
				},
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          RenameRegistrationOptions
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
				field:          wantTypeNilAll,
				want:           wantNilAll,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             RenameRegistrationOptions
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
				field:            wantNilAll,
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got RenameRegistrationOptions
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

func testFoldingRangeParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/invalid.go"}}`
	)
	wantType := FoldingRangeParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          FoldingRangeParams
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
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             FoldingRangeParams
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
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got FoldingRangeParams
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

func TestFoldingRangeKind_String(t *testing.T) {
	tests := []struct {
		name string
		s    FoldingRangeKind
		want string
	}{
		{
			name: "Comment",
			s:    CommentFoldingRange,
			want: "comment",
		},
		{
			name: "Imports",
			s:    ImportsFoldingRange,
			want: "imports",
		},
		{
			name: "Region",
			s:    RegionFoldingRange,
			want: "region",
		},
		{
			name: "Unknown",
			s:    FoldingRangeKind(""),
			want: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.s; !strings.EqualFold(string(got), tt.want) {
				t.Errorf("FoldingRangeKind(%v), want %v", got, tt.want)
			}
		})
	}
}

func testFoldingRange(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"startLine":10,"startCharacter":1,"endLine":10,"endCharacter":8,"kind":"imports"}`
		wantNilAll  = `{"startLine":10,"endLine":10}`
		wantInvalid = `{"startLine":0,"startCharacter":1,"endLine":0,"endCharacter":8,"kind":"comment"}`
	)
	wantType := FoldingRange{
		StartLine:      10,
		StartCharacter: 1,
		EndLine:        10,
		EndCharacter:   8,
		Kind:           ImportsFoldingRange,
	}
	wantTypeNilAll := FoldingRange{
		StartLine: 10,
		EndLine:   10,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          FoldingRange
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
				field:          wantTypeNilAll,
				want:           wantNilAll,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          wantType,
				want:           wantInvalid,
				wantMarshalErr: false,
				wantErr:        true,
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
		tests := []struct {
			name             string
			field            string
			want             FoldingRange
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
				field:            wantNilAll,
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            wantInvalid,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got FoldingRange
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
