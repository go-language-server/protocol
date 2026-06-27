// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"testing"

	gocmp "github.com/google/go-cmp/cmp"
)

// TestNewNullable verifies the three constructor paths for Nullable[T]:
// NewNullable (value state), NullNullable (null state), and the zero value
// (absent state). It also checks JSON round-trips through a real struct field.
func TestNewNullable(t *testing.T) {
	t.Run("state predicates", func(t *testing.T) {
		wf := WorkspaceFolder{URI: "file:///a", Name: "a"}
		var zeroWF WorkspaceFolder

		tests := map[string]struct {
			n          Nullable[WorkspaceFolder]
			wantIsZero bool
			wantIsNull bool
			wantVal    WorkspaceFolder
			wantOK     bool
		}{
			"success: NewNullable sets value state": {
				n:          NewNullable(wf),
				wantIsZero: false,
				wantIsNull: false,
				wantVal:    wf,
				wantOK:     true,
			},
			"success: NullNullable sets null state": {
				n:          NullNullable[WorkspaceFolder](),
				wantIsZero: false,
				wantIsNull: true,
				wantVal:    zeroWF,
				wantOK:     false,
			},
			"success: zero Nullable is absent state": {
				n:          Nullable[WorkspaceFolder]{},
				wantIsZero: true,
				wantIsNull: false,
				wantVal:    zeroWF,
				wantOK:     false,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				if got := tt.n.IsZero(); got != tt.wantIsZero {
					t.Errorf("IsZero = %v, want %v", got, tt.wantIsZero)
				}
				if got := tt.n.IsNull(); got != tt.wantIsNull {
					t.Errorf("IsNull = %v, want %v", got, tt.wantIsNull)
				}
				got, ok := tt.n.Get()
				if ok != tt.wantOK {
					t.Errorf("Get() ok = %v, want %v", ok, tt.wantOK)
				}
				if diff := gocmp.Diff(tt.wantVal, got); diff != "" {
					t.Errorf("Get() value mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})

	t.Run("JSON round-trip via WorkspaceFoldersInitializeParams", func(t *testing.T) {
		wf := []WorkspaceFolder{{URI: "file:///w", Name: "w"}}

		tests := map[string]struct {
			build    func() WorkspaceFoldersInitializeParams
			wantJSON string
		}{
			"success: NewNullable marshals to value": {
				build: func() WorkspaceFoldersInitializeParams {
					return WorkspaceFoldersInitializeParams{WorkspaceFolders: NewNullable(wf)}
				},
				wantJSON: `{"workspaceFolders":[{"uri":"file:///w","name":"w"}]}`,
			},
			"success: NullNullable marshals to null": {
				build: func() WorkspaceFoldersInitializeParams {
					return WorkspaceFoldersInitializeParams{WorkspaceFolders: NullNullable[[]WorkspaceFolder]()}
				},
				wantJSON: `{"workspaceFolders":null}`,
			},
			"success: zero Nullable is omitted": {
				build:    func() WorkspaceFoldersInitializeParams { return WorkspaceFoldersInitializeParams{} },
				wantJSON: `{}`,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				out, err := Marshal(tt.build())
				if err != nil {
					t.Fatalf("Marshal: %v", err)
				}
				if got, want := canon(t, out), canon(t, []byte(tt.wantJSON)); got != want {
					t.Errorf("JSON mismatch: got=%s want=%s", got, want)
				}
			})
		}
	})
}

// TestNullableTriState verifies that an optional-AND-nullable field
// distinguishes absent, explicit null, and a value across a JSON round-trip.
func TestNullableTriState(t *testing.T) {
	tests := map[string]struct {
		json       string
		wantAbsent bool
		wantNull   bool
		wantLen    int
	}{
		"success: absent": {
			json:       `{}`,
			wantAbsent: true,
			wantNull:   false,
			wantLen:    0,
		},
		"success: null": {
			json:       `{"workspaceFolders":null}`,
			wantAbsent: false,
			wantNull:   true,
			wantLen:    0,
		},
		"success: value": {
			json:       `{"workspaceFolders":[{"uri":"file:///w","name":"w"}]}`,
			wantAbsent: false,
			wantNull:   false,
			wantLen:    1,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var p WorkspaceFoldersInitializeParams
			if err := Unmarshal([]byte(tt.json), &p); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got := p.WorkspaceFolders.IsZero(); got != tt.wantAbsent {
				t.Errorf("IsZero = %v, want %v", got, tt.wantAbsent)
			}
			if got := p.WorkspaceFolders.IsNull(); got != tt.wantNull {
				t.Errorf("IsNull = %v, want %v", got, tt.wantNull)
			}
			if v, ok := p.WorkspaceFolders.Get(); ok && len(v) != tt.wantLen {
				t.Errorf("len(value) = %d, want %d", len(v), tt.wantLen)
			}
			out, err := Marshal(p)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if got, want := canon(t, out), canon(t, []byte(tt.json)); got != want {
				t.Errorf("round-trip: got=%s want=%s", got, want)
			}
		})
	}
}
