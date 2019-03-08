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
			name     string
			position Position
			want     string
			wantErr  bool
		}{
			{
				name:     "Valid",
				position: Position{Line: 25, Character: 1},
				want:     `{"line":25,"character":1}`,
				wantErr:  false,
			},
			{
				name:     "Invalid",
				position: Position{Line: 25, Character: 1},
				want:     `{"line":2,"character":0}`,
				wantErr:  true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.MarshalJSONObject(&tt.position)
				if err != nil {
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
			name     string
			position string
			want     Position
			wantErr  bool
		}{
			{
				name:     "Valid",
				position: `{"line":25, "character":1}`,
				want:     Position{Line: 25, Character: 1},
				wantErr:  false,
			},
			{
				name:     "Invalid",
				position: `{"line":2, "character":0}`,
				want:     Position{Line: 25, Character: 1},
				wantErr:  true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got Position
				dec := gojay.NewDecoder(strings.NewReader(tt.position))
				defer dec.Release()
				if err := dec.Decode(&got); err != nil {
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
			name    string
			Range   Range
			want    string
			wantErr bool
		}{
			{
				name:    "Valid",
				Range:   Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}},
				want:    `{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}`,
				wantErr: false,
			},
			{
				name:    "Invalid",
				Range:   Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}},
				want:    `{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}`,
				wantErr: true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := gojay.MarshalJSONObject(&tt.Range)
				if err != nil {
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
			name    string
			Range   string
			want    Range
			wantErr bool
		}{
			{
				name:    "Valid",
				Range:   `{"start":{"line":25,"character":1},"end":{"line":27,"character":3}}`,
				want:    Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}},
				wantErr: false,
			},
			{
				name:    "Invalid",
				Range:   `{"start":{"line":2,"character":1},"end":{"line":3,"character":2}}`,
				want:    Range{Start: Position{Line: 25, Character: 1}, End: Position{Line: 27, Character: 3}},
				wantErr: true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got Range
				dec := gojay.NewDecoder(strings.NewReader(tt.Range))
				defer dec.Release()
				if err := dec.Decode(&got); err != nil {
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
