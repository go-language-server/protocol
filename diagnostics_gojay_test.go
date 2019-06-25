// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"strings"
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/go-language-server/uri"
	"github.com/google/go-cmp/cmp"
)

func TestPublishDiagnosticsParams(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          PublishDiagnosticsParams
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: PublishDiagnosticsParams{
					URI: uri.File("/path/to/diagnostics.go"),
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
										URI: uri.File("/path/to/diagnostics.go"),
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
									Message: "diagnostics.go",
								},
							},
						},
					},
				},
				want:           `{"uri":"file:///path/to/diagnostics.go","diagnostics":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/diagnostics.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"diagnostics.go"}]}]}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: PublishDiagnosticsParams{
					URI: uri.File("/path/to/diagnostics.go"),
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
										URI: uri.File("/path/to/diagnostics.go"),
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
									Message: "diagnostics.go",
								},
							},
						},
					},
				},
				want:           `{"uri":"file:///path/to/diagnostics_gen.go","diagnostics":[{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/diagnostics_gen.go","range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}},"message":"diagnostics_gen.go"}]}]}`,
				wantMarshalErr: false,
				wantErr:        true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.MarshalJSONObject(&tt.field)
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
			want             PublishDiagnosticsParams
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"uri":"file:///path/to/diagnostics.go","diagnostics":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/diagnostics.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"diagnostics.go"}]}]}`,
				want: PublishDiagnosticsParams{
					URI: uri.File("/path/to/diagnostics.go"),
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
										URI: uri.File("/path/to/diagnostics.go"),
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
									Message: "diagnostics.go",
								},
							},
						},
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"uri":"file:///path/to/diagnostics.go","diagnostics":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/diagnostics.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"diagnostics.go"}]}]}`,
				want: PublishDiagnosticsParams{
					URI: uri.File("file:///path/to/diagnostics_gen.go"),
					Diagnostics: []Diagnostic{
						{
							Range: Range{
								Start: Position{
									Line:      2,
									Character: 1,
								},
								End: Position{
									Line:      3,
									Character: 2,
								},
							},
							Severity: SeverityError,
							Code:     "foo/bar",
							Source:   "test foo bar",
							Message:  "foo bar",
							RelatedInformation: []DiagnosticRelatedInformation{
								{
									Location: Location{
										URI: uri.File("file:///path/to/diagnostics_gen.go"),
										Range: Range{
											Start: Position{
												Line:      2,
												Character: 1,
											},
											End: Position{
												Line:      3,
												Character: 2,
											},
										},
									},
									Message: "diagnostics_gen.go",
								},
							},
						},
					},
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got PublishDiagnosticsParams
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
