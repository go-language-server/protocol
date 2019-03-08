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
				name:           "Valid",
				field:          Position{Line: 25, Character: 1},
				want:           `{"line":25,"character":1}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          Position{Line: 25, Character: 1},
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
				name:             "Valid",
				field:            `{"line":25, "character":1}`,
				want:             Position{Line: 25, Character: 1},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            `{"line":2, "character":0}`,
				want:             Position{Line: 25, Character: 1},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got Position
				dec := gojay.NewDecoder(strings.NewReader(tt.field))
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
				name:           "Valid",
				field:          Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}},
				want:           `{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}`,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "Invalid",
				field:          Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}},
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
				name:             "Valid",
				field:            `{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}`,
				want:             Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "Invalid",
				field:            `{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}`,
				want:             Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}},
				wantUnmarshalErr: false,
				wantErr:          true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := Range{}
				dec := gojay.NewDecoder(strings.NewReader(tt.field))
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
