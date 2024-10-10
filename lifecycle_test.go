package protocol

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"go.lsp.dev/uri"
)

func TestInitializeParams(t *testing.T) {
	t.Parallel()

	const wantWorkDoneToken = "156edea9-9d8d-422f-b7ee-81a84594afbb"
	const (
		want    = `{"workDoneToken":"` + wantWorkDoneToken + `","processId":25556,"clientInfo":{"name":"testClient","version":"v0.0.0"},"locale":"en-US","initializationOptions":"testdata","capabilities":{},"trace":"on","workspaceFolders":[{"uri":"file:///Users/zchee/go/src/go.lsp.dev/protocol","name":"protocol"},{"uri":"file:///Users/zchee/go/src/go.lsp.dev/jsonrpc2","name":"jsonrpc2"}]}`
		wantNil = `{"processId":25556,"capabilities":{}}`
	)
	ptoken := NewProgressToken(wantWorkDoneToken)
	wantType := InitializeParams{
		InitializeParamsBase: InitializeParamsBase{
			WorkDoneProgressParams: WorkDoneProgressParams{
				WorkDoneToken: &ptoken,
			},
			ProcessID: 25556,
			ClientInfo: &ClientInfo{
				Name:    "testClient",
				Version: "v0.0.0",
			},
			Locale:                "en-US",
			InitializationOptions: "testdata",
			Capabilities:          ClientCapabilities{},
			Trace:                 "on",
		},
		WorkspaceFoldersInitializeParams: WorkspaceFoldersInitializeParams{
			WorkspaceFolders: []WorkspaceFolder{
				{
					Name: filepath.Base("/Users/zchee/go/src/go.lsp.dev/protocol"),
					URI:  uri.File("/Users/zchee/go/src/go.lsp.dev/protocol"),
				},
				{
					Name: filepath.Base("/Users/zchee/go/src/go.lsp.dev/jsonrpc2"),
					URI:  uri.File("/Users/zchee/go/src/go.lsp.dev/jsonrpc2"),
				},
			},
		},
	}
	wantTypeNilAll := InitializeParams{
		InitializeParamsBase: InitializeParamsBase{
			ProcessID:    25556,
			Capabilities: ClientCapabilities{},
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

				if diff := cmp.Diff(tt.want, string(got)); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-want +got)\n%s", tt.name, tt.wantErr, diff)
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

				if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreTypes(WorkDoneProgressParams{})); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-want +got)\n%s", tt.name, tt.wantErr, diff)
				}

				if diff := cmp.Diff(got.WorkDoneToken, wantWorkDoneToken); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-want +got)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}
