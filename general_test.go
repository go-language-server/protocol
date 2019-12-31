// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"path/filepath"
	"testing"

	"github.com/go-language-server/uri"
	"github.com/google/go-cmp/cmp"
)

func testWorkspaceFolders(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const want = `[{"uri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","name":"protocol"},{"uri":"file:///Users/zchee/go/src/github.com/go-language-server/jsonrpc2","name":"jsonrpc2"}]`
	wantType := WorkspaceFolders{
		{
			URI:  string(uri.File("/Users/zchee/go/src/github.com/go-language-server/protocol")),
			Name: "protocol",
		},
		{
			URI:  string(uri.File("/Users/zchee/go/src/github.com/go-language-server/jsonrpc2")),
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

func testInitializeParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"processId":25556,"rootPath":"~/go/src/github.com/go-language-server/protocol","rootUri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","initializationOptions":"testdata","capabilities":{},"trace":"on","workspaceFolders":[{"uri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","name":"protocol"},{"uri":"file:///Users/zchee/go/src/github.com/go-language-server/jsonrpc2","name":"jsonrpc2"}]}`
		wantNil = `{"processId":25556,"rootUri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","capabilities":{}}`
	)
	wantType := InitializeParams{
		ProcessID:             25556,
		RootPath:              "~/go/src/github.com/go-language-server/protocol",
		RootURI:               uri.File("/Users/zchee/go/src/github.com/go-language-server/protocol"),
		InitializationOptions: "testdata",
		Capabilities:          ClientCapabilities{},
		Trace:                 "on",
		WorkspaceFolders: []WorkspaceFolder{
			{
				Name: filepath.Base("/Users/zchee/go/src/github.com/go-language-server/protocol"),
				URI:  string(uri.File("/Users/zchee/go/src/github.com/go-language-server/protocol")),
			},
			{
				Name: filepath.Base("/Users/zchee/go/src/github.com/go-language-server/jsonrpc2"),
				URI:  string(uri.File("/Users/zchee/go/src/github.com/go-language-server/jsonrpc2")),
			},
		},
	}
	wantTypeNilAll := InitializeParams{
		ProcessID:    25556,
		RootURI:      uri.File("//Users/zchee/go/src/github.com/go-language-server/protocol"),
		Capabilities: ClientCapabilities{},
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
				name:           "ValidNilAll",
				field:          wantTypeNilAll,
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
			want             InitializeParams
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
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got InitializeParams
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

func testWorkspaceClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const want = `{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"]},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true}`
	wantType := WorkspaceClientCapabilities{
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
	const want = `{"didSave":true,"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":{"valueSet":[1]},"contextSupport":true}`
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
		},
		CompletionItemKind: &TextDocumentClientCapabilitiesCompletionItemKind{
			ValueSet: []CompletionItemKind{TextCompletion},
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}}`
	wantType := TextDocumentClientCapabilitiesSignatureHelp{
		DynamicRegistration: true,
		SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
			DocumentationFormat: []MarkupKind{
				PlainText,
				Markdown,
			},
		},
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true}`
	wantType := TextDocumentClientCapabilitiesDocumentSymbol{
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"linkSupport":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"linkSupport":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"linkSupport":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"linkSupport":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true}`
	wantType := TextDocumentClientCapabilitiesCodeLens{
		DynamicRegistration: true,
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true}`
	wantType := TextDocumentClientCapabilitiesDocumentLink{
		DynamicRegistration: true,
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true}`
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"prepareSupport":true}`
	wantType := TextDocumentClientCapabilitiesRename{
		DynamicRegistration: true,
		PrepareSupport:      true,
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"relatedInformation":true,"tagSupport":true}`
	wantType := TextDocumentClientCapabilitiesPublishDiagnostics{
		RelatedInformation: true,
		TagSupport:         true,
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"dynamicRegistration":true,"rangeLimit":0.5,"lineFoldingOnly":true}`
	wantType := TextDocumentClientCapabilitiesFoldingRange{
		DynamicRegistration: true,
		RangeLimit:          float64(0.5),
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"synchronization":{"didSave":true,"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true},"completion":{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":{"valueSet":[1]},"contextSupport":true},"hover":{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]},"signatureHelp":{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}},"references":{"dynamicRegistration":true},"documentHighlight":{"dynamicRegistration":true},"documentSymbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true},"formatting":{"dynamicRegistration":true},"rangeFormatting":{"dynamicRegistration":true},"onTypeFormatting":{"dynamicRegistration":true},"declaration":{"dynamicRegistration":true,"linkSupport":true},"definition":{"dynamicRegistration":true,"linkSupport":true},"typeDefinition":{"dynamicRegistration":true,"linkSupport":true},"implementation":{"dynamicRegistration":true,"linkSupport":true},"codeAction":{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"documentLink":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":true},"rename":{"dynamicRegistration":true,"prepareSupport":true},"publishDiagnostics":{"relatedInformation":true},"foldingRange":{"dynamicRegistration":true,"rangeLimit":0.5,"lineFoldingOnly":true},"selectionRange":{"dynamicRegistration":true}}`
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
				ValueSet: []CompletionItemKind{TextCompletion},
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
		SelectionRange: &TextDocumentClientCapabilitiesSelectionRange{
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
				want:           emptyData,
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
				field:            emptyData,
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
	const want = `{"workspace":{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"]},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true},"textDocument":{"synchronization":{"didSave":true,"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true},"completion":{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":{"valueSet":[1]},"contextSupport":true},"hover":{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]},"signatureHelp":{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}},"references":{"dynamicRegistration":true},"documentHighlight":{"dynamicRegistration":true},"documentSymbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true},"formatting":{"dynamicRegistration":true},"rangeFormatting":{"dynamicRegistration":true},"onTypeFormatting":{"dynamicRegistration":true},"declaration":{"dynamicRegistration":true,"linkSupport":true},"definition":{"dynamicRegistration":true,"linkSupport":true},"typeDefinition":{"dynamicRegistration":true,"linkSupport":true},"implementation":{"dynamicRegistration":true,"linkSupport":true},"codeAction":{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"documentLink":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":true},"rename":{"dynamicRegistration":true,"prepareSupport":true},"publishDiagnostics":{"relatedInformation":true},"foldingRange":{"dynamicRegistration":true,"rangeLimit":0.5,"lineFoldingOnly":true},"selectionRange":{"dynamicRegistration":true}}}`
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
				CompletionItemKind: &TextDocumentClientCapabilitiesCompletionItemKind{
					ValueSet: []CompletionItemKind{TextCompletion},
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
			SelectionRange: &TextDocumentClientCapabilitiesSelectionRange{
				DynamicRegistration: true,
			},
		},
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
				want:           emptyData,
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
			field            []byte
			want             ClientCapabilities
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
			{
				name:             "ValidNilAll",
				field:            []byte(emptyData),
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

func testInitializeResult(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"capabilities":{"hoverProvider":true,"completionProvider":{"resolveProvider":true,"triggerCharacters":["Tab"]},"signatureHelpProvider":{"triggerCharacters":["C-K"]},"definitionProvider":true,"referencesProvider":true,"documentHighlightProvider":true,"documentSymbolProvider":true,"workspaceSymbolProvider":true,"codeActionProvider":true,"codeLensProvider":{"resolveProvider":true},"documentFormattingProvider":true,"documentRangeFormattingProvider":true,"documentOnTypeFormattingProvider":{"firstTriggerCharacter":".","moreTriggerCharacter":["f"]},"renameProvider":true,"documentLinkProvider":{"resolveProvider":true},"executeCommandProvider":{"commands":["test","command"]},"workspace":{"workspaceFolders":{"supported":true}}}}`
		wantNil = `{"capabilities":{}}`
	)
	wantType := InitializeResult{
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
				FirstTriggerCharacter: ".",
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
	}

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
				name:           "Valid",
				field:          wantType,
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

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
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
			field            string
			want             InitializeResult
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

func TestTextDocumentSyncKind_String(t *testing.T) {
	tests := []struct {
		name string
		k    TextDocumentSyncKind
		want string
	}{
		{
			name: "NoneKind",
			k:    None,
			want: "None",
		},
		{
			name: "FullKind",
			k:    Full,
			want: "Full",
		},
		{
			name: "IncrementalKind",
			k:    Incremental,
			want: "Incremental",
		},
		{
			name: "UnknownKind",
			k:    TextDocumentSyncKind(99),
			want: "99",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.k.String(); got != tt.want {
				t.Errorf("TextDocumentSyncKind.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testDocumentLinkRegistrationOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*"}],"resolveProvider":true}`
		wantNil = `{"documentSelector":[]}`
	)
	wantType := DocumentLinkRegistrationOptions{
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
	}
	wantTypeNilAll := DocumentLinkRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{},
		},
	}

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
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          wantTypeNilAll,
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
			want             DocumentLinkRegistrationOptions
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
				want:             wantTypeNilAll,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentLinkRegistrationOptions
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

func testInitializedParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
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
			want             InitializedParams
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
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
