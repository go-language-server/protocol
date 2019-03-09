// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"strings"
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/google/go-cmp/cmp"
)

func TestInitializeParams(t *testing.T) {
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
				name: "Valid",
				field: InitializeParams{
					ProcessID:             25556,
					RootPath:              "~/go/src/github.com/go-language-server/protocol",
					RootURI:               "file:///Users/zchee/go/src/github.com/go-language-server/protocol",
					InitializationOptions: "testdata",
					Capabilities:          ClientCapabilities{},
				},
				want:           `{"processId":25556,"rootPath":"~/go/src/github.com/go-language-server/protocol","rootUri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","initializationOptions":"testdata","capabilities":{}}`,
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
				want:           `{"processId":25556,"rootUri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","capabilities":{}}`,
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
			field            string
			want             InitializeParams
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"processId":25556,"rootPath":"~/go/src/github.com/go-language-server/protocol","rootUri":"file:///Users/zchee/go/src/github.com/go-language-server/protocol","initializationOptions":"testdata","capabilities":{}}`,
				want: InitializeParams{
					ProcessID:             25556,
					RootPath:              "~/go/src/github.com/go-language-server/protocol",
					RootURI:               "file:///Users/zchee/go/src/github.com/go-language-server/protocol",
					InitializationOptions: "testdata",
					Capabilities:          ClientCapabilities{},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got InitializeParams
				dec := gojay.BorrowDecoder(strings.NewReader(tt.field))
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
				name: "Valid",
				field: WorkspaceClientCapabilities{
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
				want:           `{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"]},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[,"valueSet":1,"valueSet":2,"valueSet":3,"valueSet":4,"valueSet":5,"valueSet":6]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true}`,
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
			field            string
			want             WorkspaceClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"]},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true}`,
				want: WorkspaceClientCapabilities{
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
					},
					ExecuteCommand: &WorkspaceClientCapabilitiesExecuteCommand{
						DynamicRegistration: true,
					},
					WorkspaceFolders: true,
					Configuration:    true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got WorkspaceClientCapabilities
				dec := gojay.BorrowDecoder(strings.NewReader(tt.field))
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
