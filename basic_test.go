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

func TestPosition(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          Position
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: Position{
					Line:      25,
					Character: 1,
				},
				want:           `{"line":25,"character":1}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: Position{
					Line:      25,
					Character: 1,
				},
				want:           `{"line":2,"character":0}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             Position
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"line":25, "character":1}`,
				want: Position{
					Line:      25,
					Character: 1,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"line":2, "character":0}`,
				want: Position{
					Line:      25,
					Character: 1,
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got Position
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

func TestRange(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          Range
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: Range{
					Start: Position{
						Line:      25,
						Character: 1,
					},
					End: Position{
						Line:      27,
						Character: 3,
					},
				},
				want:           `{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: Range{
					Start: Position{
						Line:      25,
						Character: 1,
					},
					End: Position{
						Line:      27,
						Character: 3,
					},
				},
				want:           `{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             Range
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}`,
				want: Range{
					Start: Position{
						Line:      25,
						Character: 1,
					},
					End: Position{
						Line:      27,
						Character: 3,
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}`,
				want: Range{
					Start: Position{
						Line:      25,
						Character: 1,
					},
					End: Position{
						Line:      27,
						Character: 3,
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

				got := Range{}
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

func TestLocation(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          Location
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: Location{
					URI: "file:///Users/gopher/go/src/github.com/go-language-server/protocol/basic_test.go",
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
				want:           `{"uri":"file:///Users/gopher/go/src/github.com/go-language-server/protocol/basic_test.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: Location{
					URI: "file:///Users/gopher/go/src/github.com/go-language-server/protocol/basic_test.go",
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
				want:           `{"uri":"file:///Users/gopher/go/src/github.com/go-language-server/protocol/basic_test.go","range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             Location
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}`,
				want: Location{
					URI: "file:///path/to/basic.go",
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
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"uri":"file:///path/to/basic.go","range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}}`,
				want: Location{
					URI: "file:///path/to/basic.go",
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
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := Location{}
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

func TestLocationLink(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          LocationLink
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: LocationLink{
					OriginSelectionRange: &Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetURI: "file:///path/to/test.go",
					TargetRange: Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetSelectionRange: Range{
						Start: Position{
							Line: 25, Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
				},
				want:           `{"originSelectionRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"targetUri":"file:///path/to/test.go","targetRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"targetSelectionRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilOriginSelectionRange",
				field: LocationLink{
					TargetURI: "file:///path/to/test.go",
					TargetRange: Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetSelectionRange: Range{
						Start: Position{
							Line: 25, Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
				},
				want:           `{"targetUri":"file:///path/to/test.go","targetRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"targetSelectionRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: LocationLink{
					OriginSelectionRange: &Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetURI: "file:///path/to/test.go",
					TargetRange: Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetSelectionRange: Range{
						Start: Position{
							Line: 25, Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
				},
				want:           `{"originSelectionRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"targetUri":"file:///path/to/test.go","targetRange":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"targetSelectionRange":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             LocationLink
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"originSelectionRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"targetUri":"file:///path/to/test.go","targetRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"targetSelectionRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}`,
				want: LocationLink{
					OriginSelectionRange: &Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetURI: "file:///path/to/test.go",
					TargetRange: Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetSelectionRange: Range{
						Start: Position{
							Line: 25, Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilOriginSelectionRange",
				field: `{"targetUri":"file:///path/to/test.go","targetRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"targetSelectionRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}}`,
				want: LocationLink{
					TargetURI: "file:///path/to/test.go",
					TargetRange: Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetSelectionRange: Range{
						Start: Position{
							Line: 25, Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"originSelectionRange":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"targetUri":"file:///path/to/test.go","targetRange":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"targetSelectionRange":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}}`,
				want: LocationLink{
					OriginSelectionRange: &Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetURI: "file:///path/to/test.go",
					TargetRange: Range{
						Start: Position{
							Line:      25,
							Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
						},
					},
					TargetSelectionRange: Range{
						Start: Position{
							Line: 25, Character: 1,
						},
						End: Position{
							Line:      27,
							Character: 3,
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

				got := LocationLink{}
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

func TestDiagnostic(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          Diagnostic
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: Diagnostic{
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
								URI: "file:///path/to/basic.go",
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
							Message: "basic_gen.go",
						},
					},
				},
				want:           `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilSeverity",
				field: Diagnostic{
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
					Code:    "foo/bar",
					Source:  "test foo bar",
					Message: "foo bar",
					RelatedInformation: []DiagnosticRelatedInformation{
						{
							Location: Location{
								URI: "file:///path/to/basic.go",
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
							Message: "basic_gen.go",
						},
					},
				},
				want:           `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilCode",
				field: Diagnostic{
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
					Source:   "test foo bar",
					Message:  "foo bar",
					RelatedInformation: []DiagnosticRelatedInformation{
						{
							Location: Location{
								URI: "file:///path/to/basic.go",
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
							Message: "basic_gen.go",
						},
					},
				},
				want:           `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilRelatedInformation",
				field: Diagnostic{
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
				},
				want:           `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilAll",
				field: Diagnostic{
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
					Message: "foo bar",
				},
				want:           `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"message":"foo bar"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: Diagnostic{
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
								URI: "file:///path/to/basic.go",
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
							Message: "basic_gen.go",
						},
					},
				},
				want:           `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             Diagnostic
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
				want: Diagnostic{
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
							Location: Location{URI: "file:///path/to/basic.go", Range: Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}}},
							Message:  "basic_gen.go",
						},
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilSeverity",
				field: `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
				want: Diagnostic{
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
					Code:    "foo/bar",
					Source:  "test foo bar",
					Message: "foo bar",
					RelatedInformation: []DiagnosticRelatedInformation{
						{
							Location: Location{
								URI: "file:///path/to/basic.go",
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
							Message: "basic_gen.go",
						},
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilCode",
				field: `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
				want: Diagnostic{
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
					Source:   "test foo bar",
					Message:  "foo bar",
					RelatedInformation: []DiagnosticRelatedInformation{
						{
							Location: Location{
								URI: "file:///path/to/basic.go",
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
							Message: "basic_gen.go",
						},
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilRelatedInformation",
				field: `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar"}`,
				want: Diagnostic{
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
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilAll",
				field: `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"message":"foo bar"}`,
				want: Diagnostic{
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
					Message: "foo bar",
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"severity":1,"code":"foo/bar","source":"test foo bar","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}},"message":"basic_gen.go"}]}`,
				want: Diagnostic{
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
								URI: "file:///path/to/basic.go",
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
							Message: "basic_gen.go",
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

				got := Diagnostic{}
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

func TestDiagnosticRelatedInformation(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DiagnosticRelatedInformation
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: DiagnosticRelatedInformation{
					Location: Location{
						URI: "file:///path/to/basic.go",
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
					Message: "basic_gen.go",
				},
				want:           `{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: DiagnosticRelatedInformation{
					Location: Location{
						URI: "file:///path/to/basic.go",
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
					Message: "basic_gen.go",
				},
				want:           `{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}},"message":"basic_gen.go"}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             DiagnosticRelatedInformation
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}`,
				want: DiagnosticRelatedInformation{
					Location: Location{
						URI: "file:///path/to/basic.go",
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
					Message: "basic_gen.go",
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}},"message":"basic_gen.go"}`,
				want: DiagnosticRelatedInformation{
					Location: Location{
						URI: "file:///path/to/basic.go",
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
					Message: "basic_gen.go",
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := DiagnosticRelatedInformation{}
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

func TestCommand(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          Command
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: Command{
					Title:     "exec echo",
					Command:   "echo",
					Arguments: []interface{}{"hello"},
				},
				want:           `{"title":"exec echo","command":"echo","arguments":["hello"]}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilArguments",
				field: Command{
					Title:   "exec echo",
					Command: "echo",
				},
				want:           `{"title":"exec echo","command":"echo"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: Command{
					Title:     "exec echo",
					Command:   "echo",
					Arguments: []interface{}{"hello"},
				},
				want:           `{"title":"exec echo","command":"true","arguments":["hello"]}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             Command
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"title":"exec echo","command":"echo","arguments":["hello"]}`,
				want: Command{
					Title:     "exec echo",
					Command:   "echo",
					Arguments: []interface{}{"hello"},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilArguments",
				field: `{"title":"exec echo","command":"echo"}`,
				want: Command{
					Title:   "exec echo",
					Command: "echo",
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"title":"exec echo","command":"echo","arguments":["hello"]}`,
				want: Command{
					Title:     "exec echo",
					Command:   "true",
					Arguments: []interface{}{"hello"},
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := Command{}
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

func TestTextEdit(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextEdit
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: TextEdit{
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
				want:           `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: TextEdit{
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
				want:           `{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"newText":"foo bar"}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             TextEdit
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}`,
				want: TextEdit{
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
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}`,
				want: TextEdit{
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
					NewText: "foo bar",
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := TextEdit{}
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

func TestTextDocumentEdit(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentEdit
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: TextDocumentEdit{
					TextDocument: VersionedTextDocumentIdentifier{
						TextDocumentIdentifier: TextDocumentIdentifier{
							URI: "file:///path/to/basic.go",
						},
						Version: Uint64(10),
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
				want:           `{"textDocument":{"uri":"file:///path/to/basic.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: TextDocumentEdit{
					TextDocument: VersionedTextDocumentIdentifier{
						TextDocumentIdentifier: TextDocumentIdentifier{
							URI: "file:///path/to/basic.go",
						},
						Version: Uint64(10),
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
				want:           `{"textDocument":{"uri":"file:///path/to/basic_gen.go","version":10},"edits":[{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"newText":"foo bar"}]}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             TextDocumentEdit
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"textDocument":{"uri":"file:///path/to/basic.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}`,
				want: TextDocumentEdit{
					TextDocument: VersionedTextDocumentIdentifier{
						TextDocumentIdentifier: TextDocumentIdentifier{
							URI: "file:///path/to/basic.go",
						},
						Version: Uint64(10),
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
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"textDocument":{"uri":"file:///path/to/basic.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}`,
				want: TextDocumentEdit{
					TextDocument: VersionedTextDocumentIdentifier{
						TextDocumentIdentifier: TextDocumentIdentifier{
							URI: "file:///path/to/basic.go",
						},
						Version: Uint64(10),
					},
					Edits: []TextEdit{
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
							NewText: "foo bar",
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

				got := TextDocumentEdit{}
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

func TestCreateFileOptions(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          CreateFileOptions
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: CreateFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
				want:           `{"overwrite":true,"ignoreIfExists":true}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilOverwrite",
				field: CreateFileOptions{
					IgnoreIfExists: true,
				},
				want:           `{"ignoreIfExists":true}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilIgnoreIfExists",
				field: CreateFileOptions{
					Overwrite: true,
				},
				want:           `{"overwrite":true}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          CreateFileOptions{},
				want:           `{}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: CreateFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
				want:           `{"overwrite":false,"ignoreIfExists":false}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             CreateFileOptions
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"overwrite":true,"ignoreIfExists":true}`,
				want: CreateFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilOverwrite",
				field: `{"ignoreIfExists":true}`,
				want: CreateFileOptions{
					IgnoreIfExists: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilIgnoreIfExists",
				field: `{"overwrite":true}`,
				want: CreateFileOptions{
					Overwrite: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            `{}`,
				want:             CreateFileOptions{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"overwrite":true,"ignoreIfExists":true}`,
				want: CreateFileOptions{
					Overwrite:      false,
					IgnoreIfExists: false,
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := CreateFileOptions{}
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

func TestCreateFile(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          CreateFile
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: CreateFile{
					Kind: "create",
					URI:  "file:///path/to/basic.go",
					Options: &CreateFileOptions{
						Overwrite:      true,
						IgnoreIfExists: true,
					},
				},
				want:           `{"kind":"create","uri":"file:///path/to/basic.go","options":{"overwrite":true,"ignoreIfExists":true}}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilOptions",
				field: CreateFile{
					Kind: "create",
					URI:  "file:///path/to/basic.go",
				},
				want:           `{"kind":"create","uri":"file:///path/to/basic.go"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: CreateFile{
					Kind: "create",
					URI:  "file:///path/to/basic.go",
					Options: &CreateFileOptions{
						Overwrite:      true,
						IgnoreIfExists: true,
					},
				},
				want:           `{"kind":"create","uri":"file:///path/to/basic_gen.go","options":{"overwrite":false,"ignoreIfExists":false}}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             CreateFile
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"kind":"create","uri":"file:///path/to/basic.go","options":{"overwrite":true,"ignoreIfExists":true}}`,
				want: CreateFile{
					Kind: "create",
					URI:  "file:///path/to/basic.go",
					Options: &CreateFileOptions{
						Overwrite:      true,
						IgnoreIfExists: true,
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilOptions",
				field: `{"kind":"create","uri":"file:///path/to/basic.go"}`,
				want: CreateFile{
					Kind: "create",
					URI:  "file:///path/to/basic.go",
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"kind":"create","uri":"file:///path/to/basic.go","options":{"overwrite":true,"ignoreIfExists":true}}`,
				want: CreateFile{
					Kind: "create",
					URI:  "file:///path/to/basic_gen.go",
					Options: &CreateFileOptions{
						Overwrite:      false,
						IgnoreIfExists: false,
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

				got := CreateFile{}
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

func TestRenameFileOptions(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          RenameFileOptions
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: RenameFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
				want:           `{"overwrite":true,"ignoreIfExists":true}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilOverwrite",
				field: RenameFileOptions{
					IgnoreIfExists: true,
				},
				want:           `{"ignoreIfExists":true}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilIgnoreIfExists",
				field: RenameFileOptions{
					Overwrite: true,
				},
				want:           `{"overwrite":true}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          RenameFileOptions{},
				want:           `{}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: RenameFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
				want:           `{"overwrite":false,"ignoreIfExists":false}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             RenameFileOptions
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"overwrite":true,"ignoreIfExists":true}`,
				want: RenameFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilOverwrite",
				field: `{"ignoreIfExists":true}`,
				want: RenameFileOptions{
					IgnoreIfExists: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilIgnoreIfExists",
				field: `{"overwrite":true}`,
				want: RenameFileOptions{
					Overwrite: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            `{}`,
				want:             RenameFileOptions{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"overwrite":true,"ignoreIfExists":true}`,
				want: RenameFileOptions{
					Overwrite:      false,
					IgnoreIfExists: false,
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := RenameFileOptions{}
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

func TestRenameFile(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          RenameFile
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: RenameFile{
					Kind:   "rename",
					OldURI: "file:///path/to/old.go",
					NewURI: "file:///path/to/new.go",
					Options: &RenameFileOptions{
						Overwrite:      true,
						IgnoreIfExists: true,
					},
				},
				want:           `{"kind":"rename","oldUri":"file:///path/to/old.go","newUri":"file:///path/to/new.go","options":{"overwrite":true,"ignoreIfExists":true}}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilOptions",
				field: RenameFile{
					Kind:   "rename",
					OldURI: "file:///path/to/old.go",
					NewURI: "file:///path/to/new.go",
				},
				want:           `{"kind":"rename","oldUri":"file:///path/to/old.go","newUri":"file:///path/to/new.go"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: RenameFile{
					Kind:   "rename",
					OldURI: "file:///path/to/old.go",
					NewURI: "file:///path/to/new.go",
					Options: &RenameFileOptions{
						Overwrite:      true,
						IgnoreIfExists: true,
					},
				},
				want:           `{"kind":"rename","oldUri":"file:///path/to/old2.go","newUri":"file:///path/to/new2.go","options":{"overwrite":false,"ignoreIfExists":false}}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             RenameFile
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"kind":"rename","oldUri":"file:///path/to/old.go","newUri":"file:///path/to/new.go","options":{"overwrite":true,"ignoreIfExists":true}}`,
				want: RenameFile{
					Kind:   "rename",
					OldURI: "file:///path/to/old.go",
					NewURI: "file:///path/to/new.go",
					Options: &RenameFileOptions{
						Overwrite:      true,
						IgnoreIfExists: true,
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilOptions",
				field: `{"kind":"rename","oldUri":"file:///path/to/old.go","newUri":"file:///path/to/new.go"}`,
				want: RenameFile{
					Kind:   "rename",
					OldURI: "file:///path/to/old.go",
					NewURI: "file:///path/to/new.go",
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"kind":"rename","oldUri":"file:///path/to/old.go","newUri":"file:///path/to/new.go","options":{"overwrite":true,"ignoreIfExists":true}}`,
				want: RenameFile{
					Kind:   "rename",
					OldURI: "file:///path/to/old2.go",
					NewURI: "file:///path/to/new2.go",
					Options: &RenameFileOptions{
						Overwrite:      false,
						IgnoreIfExists: false,
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

				got := RenameFile{}
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

func TestDeleteFileOptions(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DeleteFileOptions
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: DeleteFileOptions{
					Recursive:         true,
					IgnoreIfNotExists: true,
				},
				want:           `{"recursive":true,"ignoreIfNotExists":true}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilRecursive",
				field: DeleteFileOptions{
					IgnoreIfNotExists: true,
				},
				want:           `{"ignoreIfNotExists":true}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNiIgnoreIfNotExists",
				field: DeleteFileOptions{
					Recursive: true,
				},
				want:           `{"recursive":true}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          DeleteFileOptions{},
				want:           `{}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: DeleteFileOptions{
					Recursive:         true,
					IgnoreIfNotExists: true,
				},
				want:           `{"recursive":false,"ignoreIfNotExists":false}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             DeleteFileOptions
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"recursive":true,"ignoreIfNotExists":true}`,
				want: DeleteFileOptions{
					Recursive:         true,
					IgnoreIfNotExists: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilRecursive",
				field: `{"ignoreIfNotExists":true}`,
				want: DeleteFileOptions{
					IgnoreIfNotExists: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilIgnoreIfNotExists",
				field: `{"recursive":true}`,
				want: DeleteFileOptions{
					Recursive: true,
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            `{}`,
				want:             DeleteFileOptions{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"recursive":true,"ignoreIfNotExists":true}`,
				want: DeleteFileOptions{
					Recursive:         false,
					IgnoreIfNotExists: false,
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := DeleteFileOptions{}
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

func TestDeleteFile(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DeleteFile
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: DeleteFile{
					Kind: "delete",
					URI:  "file:///path/to/delete.go",
					Options: &DeleteFileOptions{
						Recursive:         true,
						IgnoreIfNotExists: true,
					},
				},
				want:           `{"kind":"delete","uri":"file:///path/to/delete.go","options":{"recursive":true,"ignoreIfNotExists":true}}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilOptions",
				field: DeleteFile{
					Kind: "delete",
					URI:  "file:///path/to/delete.go",
				},
				want:           `{"kind":"delete","uri":"file:///path/to/delete.go"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: DeleteFile{
					Kind: "delete",
					URI:  "file:///path/to/delete.go",
					Options: &DeleteFileOptions{
						Recursive:         true,
						IgnoreIfNotExists: true,
					},
				},
				want:           `{"kind":"delete","uri":"file:///path/to/delete2.go","options":{"recursive":false,"ignoreIfNotExists":false}}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             DeleteFile
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"kind":"delete","uri":"file:///path/to/delete.go","options":{"recursive":true,"ignoreIfNotExists":true}}`,
				want: DeleteFile{
					Kind: "delete",
					URI:  "file:///path/to/delete.go",
					Options: &DeleteFileOptions{
						Recursive:         true,
						IgnoreIfNotExists: true,
					},
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilOptions",
				field: `{"kind":"delete","uri":"file:///path/to/delete.go"}`,
				want: DeleteFile{
					Kind: "delete",
					URI:  "file:///path/to/delete.go",
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"kind":"rename","uri":"file:///path/to/delete.go","options":{"overwrite":true,"ignoreIfExists":true}}`,
				want: DeleteFile{
					Kind: "delete",
					URI:  "file:///path/to/delete2.go",
					Options: &DeleteFileOptions{
						Recursive:         false,
						IgnoreIfNotExists: false,
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

				got := DeleteFile{}
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

func TestWorkspaceEdit(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          WorkspaceEdit
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: WorkspaceEdit{
					Changes: map[DocumentURI][]TextEdit{
						"file:///path/to/basic.go": {
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
									URI: "file:///path/to/basic.go",
								},
								Version: Uint64(10),
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
				want:           `{"changes":{"file:///path/to/basic.go":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]},"documentChanges":[{"textDocument":{"uri":"file:///path/to/basic.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}]}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilChanges",
				field: WorkspaceEdit{
					DocumentChanges: []TextDocumentEdit{
						{
							TextDocument: VersionedTextDocumentIdentifier{
								TextDocumentIdentifier: TextDocumentIdentifier{
									URI: "file:///path/to/basic.go",
								},
								Version: Uint64(10),
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
				want:           `{"documentChanges":[{"textDocument":{"uri":"file:///path/to/basic.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}]}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "ValidNilDocumentChanges",
				field: WorkspaceEdit{
					Changes: map[DocumentURI][]TextEdit{
						"file:///path/to/basic.go": {
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
				want:           `{"changes":{"file:///path/to/basic.go":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: WorkspaceEdit{
					Changes: map[DocumentURI][]TextEdit{
						"file:///path/to/basic.go": {
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
									URI: "file:///path/to/basic.go",
								},
								Version: Uint64(10),
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
				want:           `{"changes":{"file:///path/to/basic_gen.go":[{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"newText":"foo bar"}]},"documentChanges":[{"textDocument":{"uri":"file:///path/to/basic_gen.go","version":10},"edits":[{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"newText":"foo bar"}]}]}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             WorkspaceEdit
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"changes":{"file:///path/to/basic.go":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]},"documentChanges":[{"textDocument":{"uri":"file:///path/to/basic.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}]}`,
				want: WorkspaceEdit{
					Changes: map[DocumentURI][]TextEdit{
						"file:///path/to/basic.go": {
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
									URI: "file:///path/to/basic.go",
								},
								Version: Uint64(10),
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
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilChanges",
				field: `{"documentChanges":[{"textDocument":{"uri":"file:///path/to/basic.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}]}`,
				want: WorkspaceEdit{
					DocumentChanges: []TextDocumentEdit{
						{
							TextDocument: VersionedTextDocumentIdentifier{
								TextDocumentIdentifier: TextDocumentIdentifier{
									URI: "file:///path/to/basic.go",
								},
								Version: Uint64(10),
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
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "ValidNilDocumentChanges",
				field: `{"changes":{"file:///path/to/basic.go":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}}`,
				want: WorkspaceEdit{
					Changes: map[DocumentURI][]TextEdit{
						"file:///path/to/basic.go": {
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
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"changes":{"file:///path/to/basic.go":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]},"documentChanges":[{"textDocument":{"uri":"file:///path/to/basic.go","version":10},"edits":[{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"newText":"foo bar"}]}]}`,
				want: WorkspaceEdit{
					Changes: map[DocumentURI][]TextEdit{
						"file:///path/to/basic.go": {
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
								NewText: "foo bar",
							},
						},
					},
					DocumentChanges: []TextDocumentEdit{
						{
							TextDocument: VersionedTextDocumentIdentifier{
								TextDocumentIdentifier: TextDocumentIdentifier{
									URI: "file:///path/to/basic_gen.go",
								},
								Version: Uint64(10),
							},
							Edits: []TextEdit{
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
									NewText: "foo bar",
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

				got := WorkspaceEdit{}
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

func TestTextDocumentIdentifier(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentIdentifier
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: TextDocumentIdentifier{
					URI: "file:///path/to/basic.go",
				},
				want:           `{"uri":"file:///path/to/basic.go"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: TextDocumentIdentifier{
					URI: "file:///path/to/basic.go",
				},
				want:           `{"uri":"file:///path/to/unknown.go"}`,
				wantMarshalErr: false,
				wantErr:        true,
			},
			{
				name:           "InvalidEmpty",
				field:          TextDocumentIdentifier{},
				want:           `{}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             TextDocumentIdentifier
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"uri":"file:///path/to/basic.go"}`,
				want: TextDocumentIdentifier{
					URI: "file:///path/to/basic.go",
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Invalid",
				field: `{"uri":"file:///path/to/basic.go"}`,
				want: TextDocumentIdentifier{
					URI: "file:///path/to/unknown.go",
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := TextDocumentIdentifier{}
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

func TestTextDocumentItem(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentItem
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name: "Valid",
				field: TextDocumentItem{
					URI:        "file:///path/to/basic.go",
					LanguageID: GoLanguage,
					Version:    float64(10),
					Text:       "Go Language",
				},
				want:           `{"uri":"file:///path/to/basic.go","languageId":"go","version":10,"text":"Go Language"}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name: "Invalid",
				field: TextDocumentItem{
					URI:        "file:///path/to/basic.go",
					LanguageID: GoLanguage,
					Version:    float64(10),
					Text:       "Go Language",
				},
				want:           `{"uri":"file:///path/to/basic_gen.go","languageId":"cpp","version":10,"text":"C++ Language"}`,
				wantMarshalErr: false,
				wantErr:        true,
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
			want             TextDocumentItem
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"uri":"file:///path/to/basic.go","languageId":"go","version":10,"text":"Go Language"}`,
				want: TextDocumentItem{
					URI:        "file:///path/to/basic.go",
					LanguageID: GoLanguage,
					Version:    float64(10),
					Text:       "Go Language",
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:  "Valid",
				field: `{"uri":"file:///path/to/basic.go","languageId":"go","version":10,"text":"Go Language"}`,
				want: TextDocumentItem{
					URI:        "file:///path/to/basic_gen.go",
					LanguageID: CppLanguage,
					Version:    float64(10),
					Text:       "C++ Language",
				},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := TextDocumentItem{}
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
