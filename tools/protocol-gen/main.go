// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/kortschak/utter"

	"go.lsp.dev/protocol/tools/protocol-gen/generator"
	"go.lsp.dev/protocol/tools/protocol-gen/resolver"
	"go.lsp.dev/protocol/tools/protocol-gen/schema"
)

const (
	LSPSchemaURI  = "https://github.com/microsoft/lsprotocol/raw/%s/generator/lsp.json"
	schemaVersion = "v2024.0.0b1"
)

func init() {
	utter.Config = utter.ConfigState{
		Indent:          " ",
		NumericWidth:    1,
		StringWidth:     1,
		Quoting:         utter.AvoidEscapes,
		BytesWidth:      16,
		CommentBytes:    true,
		AddressBytes:    true,
		CommentPointers: true,
		ElideType:       true,
		SortKeys:        true,
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := fetchLSPSchema(ctx)
	if err != nil {
		return err
	}

	model, err := schema.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	protocol, err := resolver.Resolve(model)
	if err != nil {
		return err
	}

	g := &generator.Generator{}
	g.Init()
	if err := g.TypeAliases(protocol.TypeAliases); err != nil {
		return fmt.Errorf("unable to generate type aliases: %w", err)
	}
	if err := g.Enumerations(protocol.Enumerations); err != nil {
		return fmt.Errorf("unable to generate enumerations: %w", err)
	}
	if err := g.Structures(protocol.Structures); err != nil {
		return fmt.Errorf("unable to generate structures: %w", err)
	}
	if err := g.ClientToServer(protocol.ServerToClientNotifications, protocol.BidirectionalNotifications, protocol.ServerToClientRequests, protocol.BidirectionalRequests); err != nil {
		return fmt.Errorf("unable to generate ClientToServer: %w", err)
	}
	if err := g.ServerToClient(protocol.ClientToServerNotifications, protocol.BidirectionalNotifications, protocol.ClientToServerRequests, protocol.BidirectionalRequests); err != nil {
		return fmt.Errorf("unable to generate ServerToClient: %w", err)
	}
	if err := g.GenericsTypes(); err != nil {
		return fmt.Errorf("unable to generate GenericsTypes: %w", err)
	}

	if err := g.WriteTo(); err != nil {
		return err
	}

	return nil
}

func fetchLSPSchema(ctx context.Context) ([]byte, error) {
	uri := fmt.Sprintf(LSPSchemaURI, schemaVersion)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get lsp.json http request: %w", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("response body is empty: %w", err)
		}
		return nil, fmt.Errorf("failed to get lsp.json http request: %w", err)
	}

	return b, nil
}

func debug(model schema.MetaModel) {
	sb := struct {
		MetaData      strings.Builder
		Requests      strings.Builder
		Notifications strings.Builder
		Structures    strings.Builder
		Enumerations  strings.Builder
		TypeAliases   strings.Builder
	}{}

	fmt.Fprintf(&sb.MetaData, "model.MetaData: %s\n", model.MetaData)
	for i, request := range model.Requests {
		fmt.Fprintf(&sb.Requests, "model.Requests[%d]: %#v\n", i, request)
	}
	for i, notification := range model.Notifications {
		fmt.Fprintf(&sb.Notifications, "model.Notifications[%d]: %#v\n", i, notification)
	}
	for i, structure := range model.Structures {
		fmt.Fprintf(&sb.Structures, "model.Structures[%d]: %#v\n", i, structure)
	}
	for i, enumeration := range model.Enumerations {
		fmt.Fprintf(&sb.Enumerations, "model.Enumerations[%d]: %#v\n", i, enumeration)
	}
	for i, typeAliase := range model.TypeAliases {
		fmt.Fprintf(&sb.TypeAliases, "model.TypeAliases[%d]: %#v\n", i, typeAliase)
	}

	os.Stdout.WriteString("MetaData:\n")
	os.Stdout.WriteString(sb.MetaData.String())
	os.Stdout.WriteString("\nRequests:\n")
	os.Stdout.WriteString(sb.Requests.String())
	os.Stdout.WriteString("\nNotifications:\n")
	os.Stdout.WriteString(sb.Notifications.String())
	os.Stdout.WriteString("\nStructures:\n")
	os.Stdout.WriteString(sb.Structures.String())
	os.Stdout.WriteString("\nEnumerations:\n")
	os.Stdout.WriteString(sb.Enumerations.String())
	os.Stdout.WriteString("\nTypeAliases:\n")
	os.Stdout.WriteString(sb.TypeAliases.String())
}
