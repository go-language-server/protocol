// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testShowMessageParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"message":"error message","type":1}`
		wantUnknown = `{"message":"unknown message","type":0}`
	)
	wantType := ShowMessageParams{
		Message: "error message",
		Type:    Error,
	}
	wantTypeUnkonwn := ShowMessageParams{
		Message: "unknown message",
		Type:    MessageType(0),
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

				var got ShowMessageParams
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

func testShowMessageRequestParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"actions":[{"title":"Retry"}],"message":"error message","type":1}`
		wantUnknown = `{"actions":[{"title":"Retry"}],"message":"unknown message","type":0}`
	)
	wantType := ShowMessageRequestParams{
		Actions: []MessageActionItem{
			{
				Title: "Retry",
			},
		},
		Message: "error message",
		Type:    Error,
	}
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

func testMessageActionItem(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"title":"Retry"}`
		wantOpenLog = `{"title":"Open Log"}`
	)
	wantType := MessageActionItem{
		Title: "Retry",
	}
	wantTypeOpenLog := MessageActionItem{
		Title: "Open Log",
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          MessageActionItem
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
				field:          wantTypeOpenLog,
				want:           wantOpenLog,
				wantMarshalErr: false,
				wantErr:        false,
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
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             MessageActionItem
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
				field:            wantOpenLog,
				want:             wantTypeOpenLog,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got MessageActionItem
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

func testLogMessageParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want        = `{"message":"error message","type":1}`
		wantUnknown = `{"message":"unknown message","type":0}`
	)
	wantType := LogMessageParams{
		Message: "error message",
		Type:    Error,
	}
	wantTypeUnknown := LogMessageParams{
		Message: "unknown message",
		Type:    MessageType(0),
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          LogMessageParams
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
				field:          wantTypeUnknown,
				want:           wantUnknown,
				wantMarshalErr: false,
				wantErr:        false,
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
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             LogMessageParams
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
				want:             wantTypeUnknown,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got LogMessageParams
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

func testWorkDoneProgressCreateParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		wantToken    = int64(1569)
		invalidToken = int64(1348)
	)
	var (
		want        = `{"token":` + strconv.FormatInt(wantToken, 10) + `}`
		wantInvalid = `{"token":` + strconv.FormatInt(invalidToken, 10) + `}`
	)
	token := NewNumberProgressToken(wantToken)
	wantType := WorkDoneProgressCreateParams{
		Token: *token,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          WorkDoneProgressCreateParams
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
			want             WorkDoneProgressCreateParams
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

				var got WorkDoneProgressCreateParams
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(fmt.Sprint(got.Token), strconv.FormatInt(wantToken, 10)); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testWorkDoneProgressCancelParams(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		wantToken    = int64(1569)
		invalidToken = int64(1348)
	)
	var (
		want        = `{"token":` + strconv.FormatInt(wantToken, 10) + `}`
		wantInvalid = `{"token":` + strconv.FormatInt(invalidToken, 10) + `}`
	)
	token := NewNumberProgressToken(wantToken)
	wantType := WorkDoneProgressCancelParams{
		Token: *token,
	}

	t.Run("Marshal", func(t *testing.T) {
		tests := []struct {
			name           string
			field          WorkDoneProgressCancelParams
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
			want             WorkDoneProgressCancelParams
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

				var got WorkDoneProgressCancelParams
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(fmt.Sprint(got.Token), strconv.FormatInt(wantToken, 10)); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func TestMessageType_String(t *testing.T) {
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
