// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"strings"
	"testing"
)

// TestEnumPointerPolicy verifies the generated optional-enum pointer policy: an
// enum with no zero member is emitted as a plain value with ",omitzero" (its zero
// is never a valid value, so omitting on zero is indistinguishable from unset),
// while an enum that declares a zero member keeps its pointer so a present zero
// stays distinguishable from absent.
func TestEnumPointerPolicy(t *testing.T) {
	t.Parallel()

	t.Run("no-zero enum is a value omitted when unset", func(t *testing.T) {
		t.Parallel()

		// Diagnostic.Severity is DiagnosticSeverity (values 1..4, no zero member),
		// so it is emitted as a value field, not a pointer.
		out, err := Marshal(Diagnostic{})
		if err != nil {
			t.Fatalf("marshal unset: %v", err)
		}
		if strings.Contains(string(out), `"severity"`) {
			t.Errorf("unset Severity should be omitted, got %s", out)
		}

		out, err = Marshal(Diagnostic{Severity: DiagnosticSeverityError})
		if err != nil {
			t.Fatalf("marshal set: %v", err)
		}
		if !strings.Contains(string(out), `"severity":1`) {
			t.Errorf("set Severity should be present as 1, got %s", out)
		}
	})

	t.Run("zero-member enum keeps its pointer to preserve a present zero", func(t *testing.T) {
		t.Parallel()

		// CodeAction.Kind is *CodeActionKind because CodeActionKind declares
		// Empty == "" (a meaningful "all kinds" value): a nil pointer is absent, a
		// non-nil "" is a present empty that must survive the round-trip.
		out, err := Marshal(CodeAction{})
		if err != nil {
			t.Fatalf("marshal absent: %v", err)
		}
		if strings.Contains(string(out), `"kind"`) {
			t.Errorf("absent Kind should be omitted, got %s", out)
		}

		empty := CodeActionKindEmpty
		out, err = Marshal(CodeAction{Kind: &empty})
		if err != nil {
			t.Fatalf("marshal present empty: %v", err)
		}
		if !strings.Contains(string(out), `"kind":""`) {
			t.Errorf("present empty Kind must serialize as \"kind\":\"\", got %s", out)
		}
	})
}
