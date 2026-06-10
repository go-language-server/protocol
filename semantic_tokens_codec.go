// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"slices"
	"strconv"

	"github.com/go-json-experiment/json/jsontext"
)

func appendSemanticTokensJSON(dst []byte, resultID *string, data []uint32) []byte {
	dst = slices.Grow(dst, semanticTokensObjectLenHint(resultID, data))
	dst = append(dst, '{')
	if resultID != nil {
		dst = append(dst, `"resultId":`...)
		// jsontext.AppendQuote already replaces invalid UTF-8 with U+FFFD.
		// Its only error for string input is reporting that invalid UTF-8 was
		// seen, which is explicitly allowed by this package's wireOptions.
		dst, _ = jsontext.AppendQuote(dst, *resultID)
		dst = append(dst, ',')
	}
	dst = append(dst, `"data":`...)
	dst = appendUint32JSONArray(dst, data)
	dst = append(dst, '}')
	return dst
}

func appendSemanticTokensDataObject(dst []byte, data []uint32) []byte {
	dst = slices.Grow(dst, len(`{"data":}`)+uint32JSONArrayLen(data))
	dst = append(dst, `{"data":`...)
	dst = appendUint32JSONArray(dst, data)
	dst = append(dst, '}')
	return dst
}

func appendUint32JSONArray(dst []byte, data []uint32) []byte {
	dst = slices.Grow(dst, uint32JSONArrayLen(data))
	dst = append(dst, '[')
	for i, v := range data {
		if i > 0 {
			dst = append(dst, ',')
		}
		dst = strconv.AppendUint(dst, uint64(v), 10)
	}
	dst = append(dst, ']')
	return dst
}

func semanticTokensObjectLenHint(resultID *string, data []uint32) int {
	n := len(`{"data":}`) + uint32JSONArrayLen(data)
	if resultID != nil {
		n += len(`"resultId":,`) + len(*resultID) + len(`""`)
	}
	return n
}

func uint32JSONArrayLen(data []uint32) int {
	if len(data) == 0 {
		return len(`[]`)
	}
	n := len(`[]`) + len(data) - 1
	for _, v := range data {
		n += uint32DecimalLen(v)
	}
	return n
}

func uint32DecimalLen(v uint32) int {
	switch {
	case v < 10:
		return 1
	case v < 100:
		return 2
	case v < 1000:
		return 3
	case v < 10_000:
		return 4
	case v < 100_000:
		return 5
	case v < 1_000_000:
		return 6
	case v < 10_000_000:
		return 7
	case v < 100_000_000:
		return 8
	case v < 1_000_000_000:
		return 9
	default:
		return 10
	}
}
