// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"strings"
	"testing"
)

// referenceScanQuoteBackslash is the byte-at-a-time oracle for
// dvScanQuoteBackslash.
func referenceScanQuoteBackslash(raw []byte, i int) int {
	for ; i < len(raw); i++ {
		if c := raw[i]; c == '"' || c == '\\' {
			break
		}
	}
	return i
}

// referenceScanStringSpecial is the byte-at-a-time oracle for
// dvScanStringSpecial.
func referenceScanStringSpecial(raw []byte, i int) int {
	for ; i < len(raw); i++ {
		c := raw[i]
		if c == '"' || c == '\\' || c < 0x20 || c >= 0x80 {
			break
		}
	}
	return i
}

// TestSWARScannersMatchReference pins the word-boundary behavior the SWAR
// rewrite must preserve: every special byte class at every lane offset,
// straddling loads, borrow-adjacent lanes, and truncated tails.
func TestSWARScannersMatchReference(t *testing.T) {
	t.Parallel()

	classes := []byte{'"', '\\', 0x00, 0x01, 0x1F, 0x7F + 1, 0xFF, 0xE3}
	pad := strings.Repeat("a", 40)

	inputs := make([][]byte, 0, len(classes)*18+10)
	// Each special class at each offset 0..17, with clean tails.
	for _, c := range classes {
		for off := range 18 {
			b := []byte(pad[:off])
			b = append(b, c)
			b = append(b, pad[:11]...)
			inputs = append(inputs, b)
		}
	}
	inputs = append(
		inputs,
		nil,
		[]byte(""),
		[]byte(pad),                      // clean run, no special byte
		[]byte(pad[:8]),                  // exactly one word
		[]byte(pad[:7]),                  // tail only
		[]byte("\x1f\x20"),               // borrow-adjacent: control then 0x20
		[]byte("aaaaaaa\x1f\x20aaaaaaa"), // borrow lane straddles a word edge
		[]byte("aaaaaaaa\"after"),        // special exactly at lane 0 of word 2
		[]byte("aaaaaaa\\\"x"),           // backslash lane 7, quote lane 0
		[]byte("\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"), // multibyte-only input (UTF-8 CJK)
	)

	for _, in := range inputs {
		for start := 0; start <= len(in) && start <= 3; start++ {
			if got, want := dvScanQuoteBackslash(in, start), referenceScanQuoteBackslash(in, start); got != want {
				t.Fatalf("dvScanQuoteBackslash(%q, %d) = %d, want %d", in, start, got, want)
			}
			if got, want := dvScanStringSpecial(in, start), referenceScanStringSpecial(in, start); got != want {
				t.Fatalf("dvScanStringSpecial(%q, %d) = %d, want %d", in, start, got, want)
			}
		}
	}
}

// FuzzSWARScanners proves both SWAR classifiers agree with their
// byte-at-a-time references for arbitrary inputs and start offsets.
func FuzzSWARScanners(f *testing.F) {
	f.Add([]byte(`{"uri":"file:///a/b.go","range":{}}`), 0)
	f.Add([]byte("aaaaaaa\x1f\x20"), 0)
	f.Add([]byte("plain ascii with \\\" escapes"), 3)

	f.Fuzz(func(t *testing.T, raw []byte, start int) {
		if start < 0 || start > len(raw) {
			return
		}
		if got, want := dvScanQuoteBackslash(raw, start), referenceScanQuoteBackslash(raw, start); got != want {
			t.Fatalf("dvScanQuoteBackslash(%q, %d) = %d, want %d", raw, start, got, want)
		}
		if got, want := dvScanStringSpecial(raw, start), referenceScanStringSpecial(raw, start); got != want {
			t.Fatalf("dvScanStringSpecial(%q, %d) = %d, want %d", raw, start, got, want)
		}
	})
}
