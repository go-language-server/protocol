// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"go.lsp.dev/uri"
)

func testDidOpenTextDocumentParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/basic.go","languageId":"go","version":10,"text":"Go Language"}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/basic_gen.go","languageId":"cpp","version":10,"text":"C++ Language"}}`
	)
	wantType := DidOpenTextDocumentParams{
		TextDocument: TextDocumentItem{
			URI:        uri.File("/path/to/basic.go"),
			LanguageID: GoLanguage,
			Version:    float64(10),
			Text:       "Go Language",
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DidOpenTextDocumentParams
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
			want             DidOpenTextDocumentParams
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

				var got DidOpenTextDocumentParams
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

func testDidChangeTextDocumentParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go","version":10},"contentChanges":[{"range":{"start":{"line":25,"character":1},"end":{"line":25,"character":3}},"rangeLength":2,"text":"testText"}]}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/test.go","version":10},"contentChanges":[{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":4}},"rangeLength":3,"text":"invalidText"}]}`
	)
	wantType := DidChangeTextDocumentParams{
		TextDocument: VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: TextDocumentIdentifier{
				URI: uri.File("/path/to/test.go"),
			},
			Version: NewVersion(10),
		},
		ContentChanges: []TextDocumentContentChangeEvent{
			{
				Range: &Range{
					Start: Position{
						Line:      25,
						Character: 1,
					},
					End: Position{
						Line:      25,
						Character: 3,
					},
				},
				RangeLength: 2,
				Text:        "testText",
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DidChangeTextDocumentParams
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
			want             DidChangeTextDocumentParams
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

				var got DidChangeTextDocumentParams
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

func testTextDocument(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"uri":"file:///path/to/test.go","languageId":"go","version":1,"lineCount":50}`
		wantInvalid = `{"uri":"file:///path/to/invalid.go","languageId":"typescript","version":2,"lineCount":100}`
	)
	wantType := TextDocument{
		URI:        uri.File("/path/to/test.go"),
		LanguageID: "go",
		Version:    1,
		LineCount:  50,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          TextDocument
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
			want             TextDocument
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

				var got TextDocument
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

func testTextDocumentChangeEvent(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"document":{"uri":"file:///path/to/test.go","languageId":"go","version":1,"lineCount":50}}`
		wantInvalid = `{"document":{"uri":"file:///path/to/invalid.go","languageId":"typescript","version":2,"lineCount":100}}`
	)
	wantType := TextDocumentChangeEvent{
		Document: TextDocument{
			URI:        uri.File("/path/to/test.go"),
			LanguageID: "go",
			Version:    1,
			LineCount:  50,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          TextDocumentChangeEvent
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
			want             TextDocumentChangeEvent
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

				var got TextDocumentChangeEvent
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

func TestTextDocumentSaveReason_String(t *testing.T) {
	tests := []struct {
		name string
		k    TextDocumentSaveReason
		want string
	}{
		{
			name: "Manual",
			k:    Manual,
			want: "Manual",
		},
		{
			name: "AfterDelay",
			k:    AfterDelay,
			want: "AfterDelay",
		},
		{
			name: "FocusOut",
			k:    FocusOut,
			want: "FocusOut",
		},
		{
			name: "Unknown",
			k:    TextDocumentSaveReason(0),
			want: "0",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.k.String(); got != tt.want {
				t.Errorf("TextDocumentSaveReason.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testTextDocumentContentChangeEvent(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"range":{"start":{"line":25,"character":1},"end":{"line":25,"character":3}},"rangeLength":2,"text":"testText"}`
		wantInvalid = `{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":4}},"rangeLength":3,"text":"invalidText"}`
	)
	wantType := TextDocumentContentChangeEvent{
		Range: &Range{
			Start: Position{
				Line:      25,
				Character: 1,
			},
			End: Position{
				Line:      25,
				Character: 3,
			},
		},
		RangeLength: 2,
		Text:        "testText",
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          TextDocumentContentChangeEvent
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
			want             TextDocumentContentChangeEvent
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

				var got TextDocumentContentChangeEvent
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

func testTextDocumentChangeRegistrationOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"}],"syncKind":2}`
		wantInvalid = `{"documentSelector":[{"language":"typescript","scheme":"file","pattern":"*.{ts,js}"}],"syncKind":1}`
	)
	wantType := TextDocumentChangeRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
			},
		},
		SyncKind: 2,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          TextDocumentChangeRegistrationOptions
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
			want             TextDocumentChangeRegistrationOptions
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

				var got TextDocumentChangeRegistrationOptions
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

func testWillSaveTextDocumentParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"},"reason":3}`
		wantNilAll  = `{"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/invalid.go"},"reason":1}`
	)
	wantType := WillSaveTextDocumentParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
		Reason: FocusOut,
	}
	wantTypeNilAll := WillSaveTextDocumentParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          WillSaveTextDocumentParams
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
			want             WillSaveTextDocumentParams
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

				var got WillSaveTextDocumentParams
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

func testDidSaveTextDocumentParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"text":"testText","textDocument":{"uri":"file:///path/to/test.go"}}`
		wantNilAll  = `{"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantInvalid = `{"text":"invalidText","textDocument":{"uri":"file:///path/to/invalid.go"}}`
	)
	wantType := DidSaveTextDocumentParams{
		Text: "testText",
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}
	wantTypeNilAll := DidSaveTextDocumentParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DidSaveTextDocumentParams
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
			want             DidSaveTextDocumentParams
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

				var got DidSaveTextDocumentParams
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

func testTextDocumentSaveRegistrationOptions(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"}],"includeText":true}`
		wantNilAll  = `{"documentSelector":[{"language":"go","scheme":"file","pattern":"*.go"}]}`
		wantInvalid = `{"documentSelector":[{"language":"typescript","scheme":"file","pattern":"*.{ts,js}"}],"includeText":false}`
	)
	wantType := TextDocumentSaveRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
			},
		},
		IncludeText: true,
	}
	wantTypeNilAll := TextDocumentSaveRegistrationOptions{
		TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
			DocumentSelector: DocumentSelector{
				{
					Language: "go",
					Scheme:   "file",
					Pattern:  "*.go",
				},
			},
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          TextDocumentSaveRegistrationOptions
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
			want             TextDocumentSaveRegistrationOptions
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

				var got TextDocumentSaveRegistrationOptions
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

func testDidCloseTextDocumentParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"textDocument":{"uri":"file:///path/to/test.go"}}`
		wantInvalid = `{"textDocument":{"uri":"file:///path/to/invalid.go"}}`
	)
	wantType := DidCloseTextDocumentParams{
		TextDocument: TextDocumentIdentifier{
			URI: uri.File("/path/to/test.go"),
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          DidCloseTextDocumentParams
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
			want             DidCloseTextDocumentParams
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

				var got DidCloseTextDocumentParams
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
