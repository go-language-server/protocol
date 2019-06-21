// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

import (
	"strings"
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/google/go-cmp/cmp"
)

func TestShowMessageParams(t *testing.T) {
	const want = `{"message":"error message","type":1}`
	wantType := ShowMessageParams{
		Message: "error message",
		Type:    Error,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          ShowMessageParams
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
				name: "Unknown",
				field: ShowMessageParams{
					Message: "unknown message",
					Type:    MessageType(0),
				},
				want:           `{"message":"unknown message","type":0}`,
				wantMarshalErr: false,
				wantErr:        false,
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
			want             ShowMessageParams
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
				name:  "Unknown",
				field: `{"message":"unknown message","type":0}`,
				want: ShowMessageParams{
					Message: "unknown message",
					Type:    MessageType(0),
				},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ShowMessageParams
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

func TestShowMessageRequestParams(t *testing.T) {
	const want = `{"actions":[{"title":"Retry"}],"message":"error message","type":1}`
	wantType := ShowMessageRequestParams{
		Actions: []MessageActionItem{
			{
				Title: "Retry",
			},
		},
		Message: "error message",
		Type:    Error,
	}
	const wantUnknown = `{"actions":[{"title":"Retry"}],"message":"unknown message","type":0}`
	wantTypeUnkonwn := ShowMessageRequestParams{
		Actions: []MessageActionItem{
			{
				Title: "Retry",
			},
		},
		Message: "unknown message",
		Type:    MessageType(0),
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          ShowMessageRequestParams
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
				name:           "Unknown",
				field:          wantTypeUnkonwn,
				want:           wantUnknown,
				wantMarshalErr: false,
				wantErr:        false,
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
			want             ShowMessageRequestParams
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
				name:             "Unknown",
				field:            wantUnknown,
				want:             wantTypeUnkonwn,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ShowMessageRequestParams
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

func TestMessageTypeString(t *testing.T) {
	tests := []struct {
		name string
		m    MessageType
		want string
	}{
		{
			name: "Error",
			m:    Error,
			want: "error",
		},
		{
			name: "Warning",
			m:    Warning,
			want: "warning",
		},
		{
			name: "Info",
			m:    Info,
			want: "info",
		},
		{
			name: "Log",
			m:    Log,
			want: "log",
		},
		{
			name: "Unknown",
			m:    MessageType(0),
			want: "0",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.String(); got != tt.want {
				t.Errorf("MessageType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageType_Enabled(t *testing.T) {
	tests := []struct {
		name  string
		m     MessageType
		level MessageType
		want  bool
	}{
		{
			name:  "ErrorError",
			m:     Error,
			level: Error,
			want:  true,
		},
		{
			name:  "ErrorInfo",
			m:     Error,
			level: Info,
			want:  false,
		},
		{
			name:  "ErrorUnknown",
			m:     Error,
			level: MessageType(0),
			want:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.Enabled(tt.level); got != tt.want {
				t.Errorf("MessageType.Enabled(%v) = %v, want %v", tt.level, got, tt.want)
			}
		})
	}
}

func TestToMessageType(t *testing.T) {
	tests := []struct {
		name  string
		level string
		want  MessageType
	}{
		{
			name:  "Error",
			level: "error",
			want:  Error,
		},
		{
			name:  "Warning",
			level: "warning",
			want:  Warning,
		},
		{
			name:  "Info",
			level: "info",
			want:  Info,
		},
		{
			name:  "Log",
			level: "log",
			want:  Log,
		},
		{
			name:  "Unknown",
			level: "0",
			want:  MessageType(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ToMessageType(tt.level); got != tt.want {
				t.Errorf("ToMessageType(%v) = %v, want %v", tt.level, got, tt.want)
			}
		})
	}
}
