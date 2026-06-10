// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

// TestEveryMethodRouted asserts that every LSP method constant declared in the
// generated registry (metamodel_messages.gen.go) is handled by a dispatch case in
// the hand-written transport (server.go or client.go), except $/cancelRequest
// which is special-cased in CancelHandler.
//
// This locks the registry<->dispatch invariant: when the meta-model adds a
// method and the generated registry grows, this test fails until the hand-port
// adds a matching dispatch case, so a spec method cannot be silently unrouted.
func TestEveryMethodRouted(t *testing.T) {
	t.Parallel()

	registry, err := os.ReadFile("metamodel_messages.gen.go")
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}
	// Match only constant definitions (Method... = "literal"), not the table
	// entries which use `Method: MethodX,`.
	constRe := regexp.MustCompile(`(?m)^\s*(Method[A-Za-z]+)\s*=\s*"`)
	matches := constRe.FindAllStringSubmatch(string(registry), -1)
	if len(matches) == 0 {
		t.Fatal("parsed zero method constants from metamodel_messages.gen.go")
	}

	var dispatch strings.Builder
	for _, f := range []string{"server.go", "client.go", "handler.go"} {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("read %s: %v", f, err)
		}
		dispatch.Write(b)
	}
	src := dispatch.String()

	// $/cancelRequest is intercepted by CancelHandler before dispatch.
	const cancelSpecialCase = "MethodCancelRequest"

	routed := 0
	for _, m := range matches {
		name := m[1]
		if name == cancelSpecialCase {
			continue
		}
		if !strings.Contains(src, "case "+name+":") {
			t.Errorf("method constant %s has no dispatch case in server.go/client.go", name)
			continue
		}
		routed++
	}
	t.Logf("routed %d/%d method constants (%s handled by CancelHandler)", routed, len(matches), cancelSpecialCase)
}
