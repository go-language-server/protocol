// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"strconv"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type appendMarshaler interface {
	appendLSPJSON([]byte) ([]byte, error)
}

func appendObjectName(dst []byte, first *bool, name string) []byte {
	if *first {
		*first = false
	} else {
		dst = append(dst, ',')
	}
	// Field names are generated/static ASCII identifiers, so direct quoting is
	// equivalent to json string quoting without the per-field scan.
	dst = append(dst, '"')
	dst = append(dst, name...)
	dst = append(dst, '"', ':')
	return dst
}

func appendJSONString(dst []byte, s string) []byte {
	// jsontext.AppendQuote already replaces invalid UTF-8 with U+FFFD. Its
	// only error for string input is reporting invalid UTF-8, which is allowed
	// by wireOptions, so ignoring the error preserves this package's wire
	// contract while keeping the direct append path allocation-free.
	dst, _ = jsontext.AppendQuote(dst, s)
	return dst
}

func appendRawJSONValue(dst []byte, v LSPAny) ([]byte, error) {
	if v == nil {
		return append(dst, "null"...), nil
	}
	_, n, err := dvValue(v, 0)
	if err != nil {
		return nil, err
	}
	if err := dvEnd(v, n); err != nil {
		return nil, err
	}
	return append(dst, v...), nil
}

func appendJSONMarshal(dst []byte, v any) ([]byte, error) {
	b, err := json.Marshal(v, wireOptions)
	if err != nil {
		return nil, err
	}
	return append(dst, b...), nil
}

func appendInt32JSON(dst []byte, v int32) []byte {
	return strconv.AppendInt(dst, int64(v), 10)
}

func appendUint32JSON(dst []byte, v uint32) []byte {
	return strconv.AppendUint(dst, uint64(v), 10)
}

func appendBoolJSON(dst []byte, v bool) []byte {
	if v {
		return append(dst, "true"...)
	}
	return append(dst, "false"...)
}
