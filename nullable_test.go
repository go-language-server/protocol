// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "testing"

// TestNullableTriState verifies that an optional-AND-nullable field
// distinguishes absent, explicit null, and a value across a JSON round-trip.
func TestNullableTriState(t *testing.T) {
	tests := map[string]struct {
		json       string
		wantAbsent bool
		wantNull   bool
		wantLen    int
	}{
		"success: absent": {`{}`, true, false, 0},
		"success: null":   {`{"workspaceFolders":null}`, false, true, 0},
		"success: value":  {`{"workspaceFolders":[{"uri":"file:///w","name":"w"}]}`, false, false, 1},
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
