// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

// RegularExpressionsClientCapabilities represents a client capabilities specific to regular expressions.
//
// The following features from the ECMAScript 2020 regular expression specification are NOT mandatory for a client:
//
//  Assertions
// Lookahead assertion, Negative lookahead assertion, lookbehind assertion, negative lookbehind assertion.
//  Character classes
// Matching control characters using caret notation (e.g. "\cX") and matching UTF-16 code units (e.g. "\uhhhh").
//  Group and ranges
// Named capturing groups.
//  Unicode property escapes
// None of the features needs to be supported.
//
// The only regular expression flag that a client needs to support is "i" to specify a case insensitive search.
//
// @since 3.16.0.
type RegularExpressionsClientCapabilities struct {
	// Engine is the engine's name.
	//
	// Well known engine name is "ECMAScript".
	//  https://tc39.es/ecma262/#sec-regexp-regular-expression-objects
	//  https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions
	Engine string `json:"engine"`

	// Version is the engine's version.
	//
	// Well known engine version is "ES2020".
	//  https://tc39.es/ecma262/#sec-regexp-regular-expression-objects
	//  https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions
	Version string `json:"version,omitempty"`
}
