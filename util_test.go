// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"testing"
)

func TestNewVersion(t *testing.T) {
	tests := []struct {
		name string
		i    int32
	}{
		{
			name: "Valid",
			i:    5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := NewVersion(tt.i)
			if got := NewVersion(tt.i); *got != *want {
				t.Errorf("NewVersion(%v) = %v, want %v", tt.i, *got, *want)
			}
		})
	}
}
