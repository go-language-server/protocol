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
					Source:   "",
					Message:  "foo bar",
					RelatedInformation: []DiagnosticRelatedInformation{
						{
							Location: Location{URI: "file:///path/to/basic.go", Range: Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}}},
							Message:  "basic_gen.go",
						},
					},
				},
				want:           `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
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
					Source:   "",
					Message:  "foo bar",
					RelatedInformation: []DiagnosticRelatedInformation{
						{
							Location: Location{URI: "file:///path/to/basic.go", Range: Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}}},
							Message:  "basic_gen.go",
						},
					},
				},
				want:           `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
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
			want             Diagnostic
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:  "Valid",
				field: `{"range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}},"severity":1,"code":"foo/bar","source":"","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}},"message":"basic_gen.go"}]}`,
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
					Source:   "",
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
				name:  "Invalid",
				field: `{"range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}},"severity":1,"code":"foo/bar","source":"","message":"foo bar","relatedInformation":[{"location":{"uri":"file:///path/to/basic.go","range":{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}},"message":"basic_gen.go"}]}`,
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
					Source:   "",
					Message:  "foo bar",
					RelatedInformation: []DiagnosticRelatedInformation{
						{
							Location: Location{URI: "file:///path/to/basic.go", Range: Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}}},
							Message:  "basic_gen.go",
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
