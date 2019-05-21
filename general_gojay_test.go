// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/google/go-cmp/cmp"
)

func Test_WorkspaceFolders(t *testing.T) {
	const want = `[{"uri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","name":"protocol"},{"uri":"file:///Users/zchee/go/src/github.com/go-language-server/jsonrpc2","name":"jsonrpc2"}]`
	var wantType = WorkspaceFolders{
		{
			URI:  string(ToDocumentURI("/Users/zchee/go/src/github.com/go-language-server/protocol")),
			Name: "protocol",
		},
		{
			URI:  string(ToDocumentURI("/Users/zchee/go/src/github.com/go-language-server/jsonrpc2")),
			Name: "jsonrpc2",
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          WorkspaceFolders
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

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            []byte
			want             WorkspaceFolders
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            []byte(want),
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got WorkspaceFolders
				if err := gojay.Unmarshal(tt.field, &got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestInitializeParams(t *testing.T) {
	const want = `{"processId":25556,"rootPath":"~/go/src/github.com/go-language-server/protocol","rootUri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","initializationOptions":"testdata","capabilities":{},"trace":"on","workspaceFolders":[{"uri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","name":"protocol"},{"uri":"file:///Users/zchee/go/src/github.com/go-language-server/jsonrpc2","name":"jsonrpc2"}]}`
	const wantNil = `{"processId":25556,"rootUri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","capabilities":{}}`
	var wantType = InitializeParams{
		ProcessID:             25556,
		RootPath:              "~/go/src/github.com/go-language-server/protocol",
		RootURI:               "file:///Users/zchee/go/src/github.com/go-language-server/protocol",
		InitializationOptions: "testdata",
		Capabilities:          ClientCapabilities{},
		Trace:                 "on",
		WorkspaceFolders: []WorkspaceFolder{
			{
				Name: filepath.Base("/Users/zchee/go/src/github.com/go-language-server/protocol"),
				URI:  string(ToDocumentURI("/Users/zchee/go/src/github.com/go-language-server/protocol")),
			},
			{
				Name: filepath.Base("/Users/zchee/go/src/github.com/go-language-server/jsonrpc2"),
				URI:  string(ToDocumentURI("/Users/zchee/go/src/github.com/go-language-server/jsonrpc2")),
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          InitializeParams
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
				name: "ValidNilAll",
				field: InitializeParams{
					ProcessID:    25556,
					RootURI:      "file:///Users/zchee/go/src/github.com/go-language-server/protocol",
					Capabilities: ClientCapabilities{},
				},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             InitializeParams
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            strings.NewReader(want),
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got InitializeParams
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestWorkspaceClientCapabilities(t *testing.T) {
	const want = `{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"]},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true}`
	var wantType = WorkspaceClientCapabilities{
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
					FileSymbol,
					ModuleSymbol,
					NamespaceSymbol,
					PackageSymbol,
					ClassSymbol,
					MethodSymbol,
				},
			},
		},
		ExecuteCommand: &WorkspaceClientCapabilitiesExecuteCommand{
			DynamicRegistration: true,
		},
		WorkspaceFolders: true,
		Configuration:    true,
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

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             WorkspaceClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            strings.NewReader(want),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesSynchronization(t *testing.T) {
	const want = `{"didSave":true,"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true}`
	var wantType = TextDocumentClientCapabilitiesSynchronization{
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
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesSynchronization
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            strings.NewReader(want),
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesCompletion(t *testing.T) {
	const want = `{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":1,"contextSupport":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesCompletion{
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
					CompletionItemKind: TextCompletion,
					ContextSupport:     true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesCompletion{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesCompletion
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesCompletion{
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
					CompletionItemKind: TextCompletion,
					ContextSupport:     true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesHover(t *testing.T) {
	const want = `{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesHover{
					DynamicRegistration: true,
					ContentFormat: []MarkupKind{
						PlainText,
						Markdown,
					},
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesHover{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesHover
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesHover{
					DynamicRegistration: true,
					ContentFormat: []MarkupKind{
						PlainText,
						Markdown,
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesSignatureHelp(t *testing.T) {
	const want = `{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesSignatureHelp{
					DynamicRegistration: true,
					SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
						DocumentationFormat: []MarkupKind{
							PlainText,
							Markdown,
						},
					},
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesSignatureHelp{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesSignatureHelp
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesSignatureHelp{
					DynamicRegistration: true,
					SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
						DocumentationFormat: []MarkupKind{
							PlainText,
							Markdown,
						},
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesReferences(t *testing.T) {
	const want = `{"dynamicRegistration":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesReferences{
					DynamicRegistration: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesReferences{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesReferences
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesReferences{
					DynamicRegistration: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesDocumentHighlight(t *testing.T) {
	const want = `{"dynamicRegistration":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesDocumentHighlight{
					DynamicRegistration: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDocumentHighlight{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesDocumentHighlight
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesDocumentHighlight{
					DynamicRegistration: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesDocumentSymbol(t *testing.T) {
	const want = `{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesDocumentSymbol{
					DynamicRegistration: true,
					SymbolKind: &WorkspaceClientCapabilitiesSymbolKind{
						ValueSet: []SymbolKind{
							FileSymbol,
							ModuleSymbol,
							NamespaceSymbol,
							PackageSymbol,
							ClassSymbol,
							MethodSymbol,
						},
					},
					HierarchicalDocumentSymbolSupport: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDocumentSymbol{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesDocumentSymbol
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesDocumentSymbol{
					DynamicRegistration: true,
					SymbolKind: &WorkspaceClientCapabilitiesSymbolKind{
						ValueSet: []SymbolKind{
							FileSymbol,
							ModuleSymbol,
							NamespaceSymbol,
							PackageSymbol,
							ClassSymbol,
							MethodSymbol,
						},
					},
					HierarchicalDocumentSymbolSupport: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesFormatting(t *testing.T) {
	const want = `{"dynamicRegistration":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesFormatting{
					DynamicRegistration: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesFormatting{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesFormatting
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesFormatting{
					DynamicRegistration: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesRangeFormatting(t *testing.T) {
	const want = `{"dynamicRegistration":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesRangeFormatting{
					DynamicRegistration: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesRangeFormatting{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesRangeFormatting
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesRangeFormatting{
					DynamicRegistration: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesOnTypeFormatting(t *testing.T) {
	const want = `{"dynamicRegistration":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesOnTypeFormatting{
					DynamicRegistration: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesOnTypeFormatting{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesOnTypeFormatting
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesOnTypeFormatting{
					DynamicRegistration: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesDeclaration(t *testing.T) {
	const want = `{"dynamicRegistration":true,"linkSupport":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesDeclaration{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDeclaration{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesDeclaration
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesDeclaration{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesDefinition(t *testing.T) {
	const want = `{"dynamicRegistration":true,"linkSupport":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesDefinition{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDefinition{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesDefinition
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesDefinition{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesTypeDefinition(t *testing.T) {
	const want = `{"dynamicRegistration":true,"linkSupport":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesTypeDefinition{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesTypeDefinition{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesTypeDefinition
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesTypeDefinition{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesImplementation(t *testing.T) {
	const want = `{"dynamicRegistration":true,"linkSupport":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesImplementation{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesImplementation{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesImplementation
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesImplementation{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesCodeAction(t *testing.T) {
	const want = `{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesCodeAction{
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
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesCodeAction{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesCodeAction
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesCodeAction{
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
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesCodeLens(t *testing.T) {
	const want = `{"dynamicRegistration":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesCodeLens{
					DynamicRegistration: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesCodeLens{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesCodeLens
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesCodeLens{
					DynamicRegistration: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesDocumentLink(t *testing.T) {
	const want = `{"dynamicRegistration":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesDocumentLink{
					DynamicRegistration: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesDocumentLink{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesDocumentLink
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesDocumentLink{
					DynamicRegistration: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesColorProvider(t *testing.T) {
	const want = `{"dynamicRegistration":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesColorProvider{
					DynamicRegistration: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesColorProvider{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesColorProvider
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesColorProvider{
					DynamicRegistration: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesRename(t *testing.T) {
	const want = `{"dynamicRegistration":true,"prepareSupport":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesRename{
					DynamicRegistration: true,
					PrepareSupport:      true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesRename{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesRename
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesRename{
					DynamicRegistration: true,
					PrepareSupport:      true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesPublishDiagnostics(t *testing.T) {
	const want = `{"relatedInformation":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesPublishDiagnostics{
					RelatedInformation: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesPublishDiagnostics{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesPublishDiagnostics
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesPublishDiagnostics{
					RelatedInformation: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilitiesFoldingRange(t *testing.T) {
	const want = `{"dynamicRegistration":true,"rangeLimit":0.5,"lineFoldingOnly":true}`

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
				name: "Valid",
				field: TextDocumentClientCapabilitiesFoldingRange{
					DynamicRegistration: true,
					RangeLimit:          float64(0.5),
					LineFoldingOnly:     true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesFoldingRange{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilitiesFoldingRange
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: TextDocumentClientCapabilitiesFoldingRange{
					DynamicRegistration: true,
					RangeLimit:          float64(0.5),
					LineFoldingOnly:     true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestTextDocumentClientCapabilities(t *testing.T) {
	const want = `{"synchronization":{"didSave":true,"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true},"completion":{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":1,"contextSupport":true},"hover":{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]},"signatureHelp":{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}},"references":{"dynamicRegistration":true},"documentHighlight":{"dynamicRegistration":true},"documentSymbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true},"formatting":{"dynamicRegistration":true},"rangeFormatting":{"dynamicRegistration":true},"onTypeFormatting":{"dynamicRegistration":true},"declaration":{"dynamicRegistration":true,"linkSupport":true},"definition":{"dynamicRegistration":true,"linkSupport":true},"typeDefinition":{"dynamicRegistration":true,"linkSupport":true},"implementation":{"dynamicRegistration":true,"linkSupport":true},"codeAction":{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"documentLink":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":true},"rename":{"dynamicRegistration":true,"prepareSupport":true},"publishDiagnostics":{"relatedInformation":true},"foldingRange":{"dynamicRegistration":true,"rangeLimit":0.5,"lineFoldingOnly":true}}`
	var wantType = TextDocumentClientCapabilities{
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
			CompletionItemKind: TextCompletion,
			ContextSupport:     true,
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
					FileSymbol,
					ModuleSymbol,
					NamespaceSymbol,
					PackageSymbol,
					ClassSymbol,
					MethodSymbol,
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
			RangeLimit:          float64(0.5),
			LineFoldingOnly:     true,
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
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             TextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            strings.NewReader(want),
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(emptyData),
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
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestClientCapabilities(t *testing.T) {
	const want = `{"workspace":{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"]},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true},"textDocument":{"synchronization":{"didSave":true,"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true},"completion":{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":1,"contextSupport":true},"hover":{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]},"signatureHelp":{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}},"references":{"dynamicRegistration":true},"documentHighlight":{"dynamicRegistration":true},"documentSymbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true},"formatting":{"dynamicRegistration":true},"rangeFormatting":{"dynamicRegistration":true},"onTypeFormatting":{"dynamicRegistration":true},"declaration":{"dynamicRegistration":true,"linkSupport":true},"definition":{"dynamicRegistration":true,"linkSupport":true},"typeDefinition":{"dynamicRegistration":true,"linkSupport":true},"implementation":{"dynamicRegistration":true,"linkSupport":true},"codeAction":{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"documentLink":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":true},"rename":{"dynamicRegistration":true,"prepareSupport":true},"publishDiagnostics":{"relatedInformation":true},"foldingRange":{"dynamicRegistration":true,"rangeLimit":0.5,"lineFoldingOnly":true}}}`
	var wantType = ClientCapabilities{
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
						FileSymbol,
						ModuleSymbol,
						NamespaceSymbol,
						PackageSymbol,
						ClassSymbol,
						MethodSymbol,
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
				CompletionItemKind: TextCompletion,
				ContextSupport:     true,
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
						FileSymbol,
						ModuleSymbol,
						NamespaceSymbol,
						PackageSymbol,
						ClassSymbol,
						MethodSymbol,
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
				RangeLimit:          float64(0.5),
				LineFoldingOnly:     true,
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name               string
			marshalFunc        marshalFunc
			compareMarshalFunc marshalFunc
			field              ClientCapabilities
			want               string
			wantMarshalErr     bool
			wantErr            bool
		}{
			{
				name:           "Valid",
				marshalFunc:    gojay.Marshal,
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				marshalFunc:    gojay.Marshal,
				field:          ClientCapabilities{},
				want:           emptyData,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "jsonValid",
				marshalFunc:    json.Marshal,
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:               "Compare",
				marshalFunc:        gojay.Marshal,
				compareMarshalFunc: json.Marshal,
				field:              wantType,
				want:               want,
				wantMarshalErr:     false,
				wantErr:            false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := tt.marshalFunc(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
				}
				if tt.compareMarshalFunc != nil {
					got2, err := tt.compareMarshalFunc(&tt.field)
					if (err != nil) != tt.wantMarshalErr {
						t.Error(err)
						return
					}
					if diff := cmp.Diff(string(got), string(got2)); (diff != "") != tt.wantErr {
						t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
					}
					return
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
			name                 string
			unmarshalFunc        unmarshalFunc
			compareUnmarshalFunc unmarshalFunc
			field                []byte
			want                 ClientCapabilities
			wantUnmarshalErr     bool
			wantErr              bool
		}{
			{
				name:             "Valid",
				unmarshalFunc:    gojay.Unsafe.Unmarshal,
				field:            []byte(want),
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				unmarshalFunc:    gojay.Unsafe.Unmarshal,
				field:            []byte(emptyData),
				want:             ClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "jsonValid",
				unmarshalFunc:    json.Unmarshal,
				field:            []byte(want),
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:                 "compare",
				unmarshalFunc:        gojay.Unsafe.Unmarshal,
				compareUnmarshalFunc: json.Unmarshal,
				field:                []byte(want),
				want:                 wantType,
				wantUnmarshalErr:     false,
				wantErr:              false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ClientCapabilities
				if err := tt.unmarshalFunc(tt.field, &got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}
				if tt.compareUnmarshalFunc != nil {
					var got2 ClientCapabilities
					if err := tt.compareUnmarshalFunc(tt.field, &got2); (err != nil) != tt.wantUnmarshalErr {
						t.Error(err)
						return
					}
					if diff := cmp.Diff(got, got2); (diff != "") != tt.wantErr {
						t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
					}
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestInitializeResult(t *testing.T) {
	const want = `{"capabilities":{"hoverProvider":true,"completionProvider":{"resolveProvider":true,"triggerCharacters":["Tab"]},"signatureHelpProvider":{"triggerCharacters":["C-K"]},"definitionProvider":true,"referencesProvider":true,"documentHighlightProvider":true,"documentSymbolProvider":true,"workspaceSymbolProvider":true,"codeActionProvider":true,"codeLensProvider":{"resolveProvider":true},"documentFormattingProvider":true,"documentRangeFormattingProvider":true,"documentOnTypeFormattingProvider":{"firstTriggerCharacter":"<Space>","moreTriggerCharacter":["f"]},"renameProvider":true,"documentLinkProvider":{"resolveProvider":true},"executeCommandProvider":{"commands":["test","command"]},"workspace":{"workspaceFolders":{"supported":true}}}}`
	const wantNil = `{"capabilities":{}}`

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          InitializeResult
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: InitializeResult{
					Capabilities: ServerCapabilities{
						TextDocumentSync: nil,
						HoverProvider:    true,
						CompletionProvider: &CompletionOptions{
							ResolveProvider:   true,
							TriggerCharacters: []string{"Tab"},
						},
						SignatureHelpProvider: &SignatureHelpOptions{
							TriggerCharacters: []string{"C-K"},
						},
						DefinitionProvider:        true,
						TypeDefinitionProvider:    nil,
						ImplementationProvider:    nil,
						ReferencesProvider:        true,
						DocumentHighlightProvider: true,
						DocumentSymbolProvider:    true,
						WorkspaceSymbolProvider:   true,
						CodeActionProvider:        true,
						CodeLensProvider: &CodeLensOptions{
							ResolveProvider: true,
						},
						DocumentFormattingProvider:      true,
						DocumentRangeFormattingProvider: true,
						DocumentOnTypeFormattingProvider: &DocumentOnTypeFormattingOptions{
							FirstTriggerCharacter: "<Space>",
							MoreTriggerCharacter:  []string{"f"},
						},
						RenameProvider: true,
						DocumentLinkProvider: &DocumentLinkOptions{
							ResolveProvider: true,
						},
						ColorProvider:        nil,
						FoldingRangeProvider: nil,
						ExecuteCommandProvider: &ExecuteCommandOptions{
							Commands: []string{"test", "command"},
						},
						Workspace: &ServerCapabilitiesWorkspace{
							WorkspaceFolders: &ServerCapabilitiesWorkspaceFolders{
								Supported:           true,
								ChangeNotifications: nil,
							},
						},
						Experimental: nil,
					},
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          InitializeResult{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Logf("got: %s", string(got))
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            *strings.Reader
			want             InitializeResult
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: InitializeResult{
					Capabilities: ServerCapabilities{
						TextDocumentSync: nil,
						HoverProvider:    true,
						CompletionProvider: &CompletionOptions{
							ResolveProvider:   true,
							TriggerCharacters: []string{"Tab"},
						},
						SignatureHelpProvider: &SignatureHelpOptions{
							TriggerCharacters: []string{"C-K"},
						},
						DefinitionProvider:        true,
						TypeDefinitionProvider:    nil,
						ImplementationProvider:    nil,
						ReferencesProvider:        true,
						DocumentHighlightProvider: true,
						DocumentSymbolProvider:    true,
						WorkspaceSymbolProvider:   true,
						CodeActionProvider:        true,
						CodeLensProvider: &CodeLensOptions{
							ResolveProvider: true,
						},
						DocumentFormattingProvider:      true,
						DocumentRangeFormattingProvider: true,
						DocumentOnTypeFormattingProvider: &DocumentOnTypeFormattingOptions{
							FirstTriggerCharacter: "<Space>",
							MoreTriggerCharacter:  []string{"f"},
						},
						RenameProvider: true,
						DocumentLinkProvider: &DocumentLinkOptions{
							ResolveProvider: true,
						},
						ColorProvider:        nil,
						FoldingRangeProvider: nil,
						ExecuteCommandProvider: &ExecuteCommandOptions{
							Commands: []string{"test", "command"},
						},
						Workspace: &ServerCapabilitiesWorkspace{
							WorkspaceFolders: &ServerCapabilitiesWorkspaceFolders{
								Supported:           true,
								ChangeNotifications: nil,
							},
						},
						Experimental: nil,
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(wantNil),
				want:             InitializeResult{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got InitializeResult
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestDocumentLinkRegistrationOptions(t *testing.T) {
	const want = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*"}],"resolveProvider":true}`
	const wantNil = `{"documentSelector":[]}`

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DocumentLinkRegistrationOptions
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: DocumentLinkRegistrationOptions{
					TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
						DocumentSelector: DocumentSelector{
							{
								Language: "go",
								Scheme:   "file",
								Pattern:  `*`,
							},
						},
					},
					ResolveProvider: true,
				},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          DocumentLinkRegistrationOptions{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             DocumentLinkRegistrationOptions
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: strings.NewReader(want),
				want: DocumentLinkRegistrationOptions{
					TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
						DocumentSelector: DocumentSelector{
							{
								Language: "go",
								Scheme:   "file",
								Pattern:  `*`,
							},
						},
					},
					ResolveProvider: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            strings.NewReader(wantNil),
				want:             DocumentLinkRegistrationOptions{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentLinkRegistrationOptions
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestInitializedParams(t *testing.T) {
	const want = `{}`

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          InitializedParams
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          InitializedParams{},
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.Marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Error(err)
					return
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
			field            *strings.Reader
			want             InitializedParams
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            strings.NewReader(want),
				want:             InitializedParams{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got InitializedParams
				dec := gojay.BorrowDecoder(tt.field)
				defer dec.Release()
				if err := dec.Decode(&got); (err != nil) != tt.wantUnmarshalErr {
					t.Error(err)
					return
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}
