// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package genlsp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-json-experiment/json"
)

const maxHTTPSchemaBytes = 32 << 20

// Load reads and decodes the metaModel.json from path, which may be a local
// file path or an http(s)/file URL.
func Load(ctx context.Context, path string) (*MetaModel, error) {
	if strings.TrimSpace(path) == "" {
		return nil, errors.New("require non-empty meta-model path or URL")
	}

	data, err := readSchemaSource(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("read meta model: %w", err)
	}

	return Decode(data)
}

func readSchemaSource(ctx context.Context, source string) ([]byte, error) {
	parsed, err := url.Parse(source)
	if err != nil {
		return nil, fmt.Errorf("parse source: %w", err)
	}

	switch parsed.Scheme {
	case "http", "https":
		if parsed.Host == "" {
			return nil, fmt.Errorf("invalid %s URL: missing host", parsed.Scheme)
		}
		return readHTTPSchemaSourceWithLimit(ctx, source)

	case "file":
		source = strings.TrimPrefix(source, "file://")
		return os.ReadFile(source)

	default:
		if parsed.Host != "" {
			return nil, fmt.Errorf("unsupported schema URL scheme %q", parsed.Scheme)
		}
		return os.ReadFile(source)
	}
}

func readHTTPSchemaSourceWithLimit(ctx context.Context, source string) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, source, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := client.Do(request)
	if err != nil {
		if urlError, ok := errors.AsType[*url.Error](err); ok {
			err = urlError.Err
		}
		return nil, fmt.Errorf("fetch %s: %w", source, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("fetch %s: unexpected HTTP status %s", source, resp.Status)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, maxHTTPSchemaBytes))
	if err != nil {
		return nil, fmt.Errorf("read %s response: %w", source, err)
	}

	if len(body) > maxHTTPSchemaBytes {
		return nil, fmt.Errorf("read %s response: schema exceeds %d bytes", source, maxHTTPSchemaBytes)
	}

	return body, nil
}

// Decode decodes a metaModel.json payload.
func Decode(data []byte) (*MetaModel, error) {
	var m MetaModel
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("decode meta model: %w", err)
	}
	return &m, nil
}
