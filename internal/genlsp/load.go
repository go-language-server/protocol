// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"fmt"
	"os"

	"github.com/go-json-experiment/json"
)

// Load reads and decodes a metaModel.json file from path.
func Load(path string) (*MetaModel, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read meta model: %w", err)
	}
	return Decode(data)
}

// Decode decodes a metaModel.json payload.
func Decode(data []byte) (*MetaModel, error) {
	var m MetaModel
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("decode meta model: %w", err)
	}
	return &m, nil
}
