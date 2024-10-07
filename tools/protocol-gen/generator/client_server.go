// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package generator

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"go.lsp.dev/protocol/tools/protocol-gen/protocol"
)

func (gen *Generator) ClientToServer(clientNotifications, bidiNotifications []*protocol.Notification, clientRequests, bidiRequests []*protocol.Request) error {
	g := NewPrinter("client_interface")
	gen.client = append(gen.client, g)

	g.PP(`const (`)
	if len(bidiNotifications) > 0 {
		slices.SortFunc(bidiNotifications, func(a, b *protocol.Notification) int {
			return cmp.Compare(strings.ToLower(a.TypeName), strings.ToLower(b.TypeName))
		})
		for _, meth := range bidiNotifications {
			g.PP(`	`, `MethodClient`+normalizeMethodName(meth.Method), ` ClientMethod `, ` = `, strconv.Quote(meth.Method), ` // bidirect client notification`)
		}
	}
	slices.SortFunc(clientNotifications, func(a, b *protocol.Notification) int {
		return cmp.Compare(strings.ToLower(a.TypeName), strings.ToLower(b.TypeName))
	})
	for _, meth := range clientNotifications {
		g.PP(`	`, `Method`+normalizeMethodName(meth.Method), ` ClientMethod `, ` = `, strconv.Quote(meth.Method), ` // client notification`)
	}
	if len(bidiRequests) > 0 {
		slices.SortFunc(bidiRequests, func(a, b *protocol.Request) int {
			return cmp.Compare(strings.ToLower(a.TypeName), strings.ToLower(b.TypeName))
		})
		for _, meth := range bidiRequests {
			g.PP(`	`, `MethodClient`+normalizeMethodName(meth.Method), ` ClientMethod `, ` = `, strconv.Quote(meth.Method), ` // bidirect client request`)
		}
	}
	slices.SortFunc(clientRequests, func(a, b *protocol.Request) int {
		return cmp.Compare(strings.ToLower(a.TypeName), strings.ToLower(b.TypeName))
	})
	for _, meth := range clientRequests {
		g.PP(`	`, `Method`+normalizeMethodName(meth.Method), ` ClientMethod `, ` = `, strconv.Quote(meth.Method), ` // client request`)
	}
	g.PP(`)`)

	g.PP(`type Client interface {`)

	notifications := append(bidiNotifications, clientNotifications...)
	reqests := append(bidiRequests, clientRequests...)

	for i, notify := range notifications {
		meth := normalizeMethodName(notify.TypeName)
		meth = strings.TrimSuffix(meth, "Notification")
		// write Documentation
		if notify.Documentation != "" {
			g.PP(`// `, meth, normalizeDocumentation(notify.Documentation))
		}
		if notify.Since != "" {
			if notify.Documentation != "" {
				g.PP(`//`)
			}
			g.P(`// @since `, strings.ReplaceAll(notify.Since, "\n", " "))
			if notify.Proposed {
				g.P(` proposed`)
			}
			g.P("\n")
		}
		if err := gen.notification(g, meth, notify); err != nil {
			return err
		}
		g.P("\n")
		// add newline per clientNotifications and bidiNotifications
		if i < len(notifications)-1 {
			g.P("\n")
		}
	}
	for i, req := range reqests {
		meth := normalizeMethodName(req.TypeName)
		if meth != "ShowMessageRequest" {
			meth = strings.TrimSuffix(meth, "Request")
		}
		// write Documentation
		if req.Documentation != "" {
			g.PP(`// `, meth, normalizeDocumentation(req.Documentation))
		}
		if req.Since != "" {
			if req.Documentation != "" {
				g.PP(`//`)
			}
			g.P(`// @since `, strings.ReplaceAll(req.Since, "\n", " "))
			if req.Proposed {
				g.P(` proposed`)
			}
			g.P("\n")
		}
		if _, err := gen.request(g, meth, req); err != nil {
			return err
		}
		g.P("\n")
		// add newline per clientRequests and bidiRequests
		if i < len(reqests)-1 {
			g.P("\n")
		}
	}
	g.PP(`}`)
	g.P("\n")

	g.PP(`// UnimplementedClient should be embedded to have forward compatible implementations.`)
	g.PP(`type UnimplementedClient struct {}`)
	g.P("\n")
	for i, notify := range notifications {
		meth := normalizeMethodName(notify.TypeName)
		meth = strings.TrimSuffix(meth, "Notification")
		g.P(`func (UnimplementedClient) `)
		if err := gen.notification(g, meth, notify); err != nil {
			return err
		}
		g.PP(` {`)
		g.PP(`	return jsonrpc2.ErrInternal`)
		g.PP(`}`)
		// add newline per clientNotifications and bidiNotifications
		if i < len(notifications)-1 {
			g.P("\n")
		}
	}
	for i, req := range reqests {
		meth := normalizeMethodName(req.TypeName)
		if meth != "ShowMessageRequest" {
			meth = strings.TrimSuffix(meth, "Request")
		}
		if meth == "" {
			continue
		}
		g.P(`func (UnimplementedClient) `)
		n, err := gen.request(g, meth, req)
		if err != nil {
			return err
		}
		g.PP(` {`)
		g.P(`	return `)
		if n > 0 {
			g.P(`	nil, `)
		}
		g.PP(`jsonrpc2.ErrInternal`)
		g.PP(`}`)
		// add newline per clientRequests and bidiRequests
		if i < len(reqests)-1 {
			g.P("\n")
		}
	}

	if err := g.WriteTo(); err != nil {
		return err
	}

	return nil
}

func (gen *Generator) ServerToClient(serverNotifications, bidiNotifications []*protocol.Notification, serverNequests, bidiRequests []*protocol.Request) error {
	g := NewPrinter("server_interface")
	gen.server = append(gen.server, g)

	g.PP(`const (`)
	if len(bidiNotifications) > 0 {
		slices.SortFunc(bidiNotifications, func(a, b *protocol.Notification) int {
			return cmp.Compare(a.TypeName, b.TypeName)
		})
		for _, meth := range bidiNotifications {
			g.PP(`	`, `MethodServer`+normalizeMethodName(meth.Method), ` ServerMethod `, ` = `, strconv.Quote(meth.Method), ` // bidirect server notification`)
		}
	}
	slices.SortFunc(serverNotifications, func(a, b *protocol.Notification) int {
		return cmp.Compare(a.TypeName, b.TypeName)
	})
	for _, meth := range serverNotifications {
		g.PP(`	`, `Method`+normalizeMethodName(meth.Method), ` ServerMethod `, ` = `, strconv.Quote(meth.Method), ` // server notification`)
	}
	if len(bidiRequests) > 0 {
		slices.SortFunc(bidiRequests, func(a, b *protocol.Request) int {
			return cmp.Compare(a.TypeName, b.TypeName)
		})
		for _, meth := range bidiRequests {
			g.PP(`	`, `MethodServer`+normalizeMethodName(meth.Method), ` ServerMethod `, ` = `, strconv.Quote(meth.Method), ` // bidirect server request`)
		}
	}
	slices.SortFunc(serverNequests, func(a, b *protocol.Request) int {
		return cmp.Compare(a.TypeName, b.TypeName)
	})
	for _, meth := range serverNequests {
		g.PP(`	`, `Method`+normalizeMethodName(meth.Method), ` ServerMethod `, ` = `, strconv.Quote(meth.Method), ` // server request`)
	}
	g.PP(`)`)

	g.PP(`type Server interface {`)

	notifications := append(slices.Clip(bidiNotifications), slices.Clip(serverNotifications)...)
	reqests := slices.Clip(serverNequests)

	for i, notify := range notifications {
		meth := normalizeMethodName(notify.TypeName)
		meth = strings.TrimSuffix(meth, "Notification")
		// write Documentation
		if notify.Documentation != "" {
			g.PP(`// `, meth, normalizeDocumentation(notify.Documentation))
		}
		if notify.Since != "" {
			if notify.Documentation != "" {
				g.PP(`//`)
			}
			g.P(`// @since `, strings.ReplaceAll(notify.Since, "\n", " "))
			if notify.Proposed {
				g.P(` proposed`)
			}
			g.P("\n")
		}
		if err := gen.notification(g, meth, notify); err != nil {
			return err
		}
		g.P("\n")
		// add newline per serverNotifications and bidiNotifications
		if i < len(notifications)-1 {
			g.P("\n")
		}
	}
	for i, req := range reqests {
		meth := normalizeMethodName(req.TypeName)
		if meth != "ShowMessageRequest" {
			meth = strings.TrimSuffix(meth, "Request")
		}
		// write Documentation
		if req.Documentation != "" {
			g.PP(`// `, meth, normalizeDocumentation(req.Documentation))
		}
		if req.Since != "" {
			if req.Documentation != "" {
				g.PP(`//`)
			}
			g.P(`// @since `, strings.ReplaceAll(req.Since, "\n", " "))
			if req.Proposed {
				g.P(` proposed`)
			}
			g.P("\n")
		}
		if _, err := gen.request(g, meth, req); err != nil {
			return err
		}
		g.P("\n")
		// add newline per serverNequests and bidiRequests
		if i < len(reqests)-1 {
			g.P("\n")
		}
	}
	g.P("\n")
	g.PP(`Request(ctx context.Context, method string, params any) (any, error)`)
	g.PP(`}`)
	g.P("\n")

	g.PP(`// UnimplementedServer should be embedded to have forward compatible implementations.`)
	g.PP(`type UnimplementedServer struct {}`)
	g.P("\n")
	for i, notify := range notifications {
		meth := normalizeMethodName(notify.TypeName)
		meth = strings.TrimSuffix(meth, "Notification")
		g.P(`func (UnimplementedServer) `)
		if err := gen.notification(g, meth, notify); err != nil {
			return err
		}
		g.PP(` {`)
		g.PP(`	return jsonrpc2.ErrInternal`)
		g.PP(`}`)
		// add newline per clientNotifications and bidiNotifications
		if i < len(notifications)-1 {
			g.P("\n")
		}
	}
	for i, req := range reqests {
		meth := normalizeMethodName(req.TypeName)
		if meth != "ShowMessageRequest" {
			meth = strings.TrimSuffix(meth, "Request")
		}
		g.P(`func (UnimplementedServer) `)
		n, err := gen.request(g, meth, req)
		if err != nil {
			return err
		}
		g.PP(` {`)
		g.P(`	return `)
		if n > 0 {
			g.P(`	nil, `)
		}
		g.PP(`jsonrpc2.ErrInternal`)
		g.PP(`}`)
		// add newline per clientRequests and bidiRequests
		if i < len(reqests)-1 {
			g.P("\n")
		}
	}

	if err := g.WriteTo(); err != nil {
		return err
	}

	return nil
}

// notification generates notification Go type from the metaModel schema definition.
func (gen *Generator) notification(g Printer, meth string, notify *protocol.Notification) error {
	g.P(`	`, meth, `(ctx context.Context, `)

	if len(notify.Params) == 0 {
		g.P(`) error`)
		return nil
	}
	for _, param := range notify.Params {
		switch p := param.(type) {
		case *protocol.ReferenceType:
			g.P(`params `)
			pt, ok := normalizeHasLSPTypes(p.String())
			if !ok {
				g.P(`*`)
			}
			g.P(pt)

		default:
			panic(fmt.Sprintf("notification: %#v\n", p))
		}
	}

	g.P(`) error`)

	return nil
}

// request generates request Go type from the metaModel schema definition.
func (gen *Generator) request(g Printer, meth string, req *protocol.Request) (nResult int, err error) {
	g.P(`	`, meth, `(ctx context.Context, `)

	for _, param := range req.Params {
		switch p := param.(type) {
		case *protocol.ReferenceType:
			g.P(`params *`, normalizeLSPTypes(p.String()))

		default:
			panic(fmt.Sprintf("requests: %#v\n", p))
		}
	}

	g.P(`) (`)

	switch r := req.Result.(type) {
	case *protocol.NullType:
		// nothing to do

	case *protocol.ReferenceType:
		nResult++
		g.P(`*`)
		g.P(normalizeLSPTypes(r.String()))
		g.P(`, `)

	case *protocol.ArrayType:
		nResult++
		gen.renderRequestsArrayType(g, req, r)
		g.P(`, `)

	case *protocol.OrType:
		nResult++
		switch {
		case len(r.Items) == 2 && (isNull(r.Items[0], r.Items[1])):
			gen.renderRequestssOrTypeNull(g, req, r)
			g.P(`, `)
		default:
			genericsProp := &protocol.Property{
				Name:          meth + "Result",
				Documentation: req.Documentation,
				Since:         req.Since,
				Proposed:      req.Proposed,
			}
			gen.renderRequestssOrType(g, r, genericsProp)
			g.P(`, `)
		}

	default:
		panic(fmt.Sprintf("requests: %#v\n", r))
	}

	g.P(`error)`)

	return nResult, nil
}

func (gen *Generator) renderRequestsArrayType(g Printer, req *protocol.Request, array *protocol.ArrayType) {
	elem := array.Element
	switch elem := elem.(type) {
	case *protocol.ReferenceType:
		g.P(`[]`)
		pt, ok := normalizeHasLSPTypes(elem.String())
		if !ok {
			g.P(`*`)
		}
		g.P(pt)

	case *protocol.OrType:
		switch {
		case len(elem.Items) == 2 && (isNull(elem.Items[0], elem.Items[1])):
			g.P(`[]*`)
			gen.renderRequestssOrTypeNull(g, req, elem)
		default:
			genericsProp := &protocol.Property{
				Name:          normalizeMethodName(req.TypeName) + "Result",
				Documentation: req.Documentation,
				Since:         req.Since,
				Proposed:      req.Proposed,
			}
			gen.renderRequestssOrType(g, elem, genericsProp)
		}

	default:
		panic(fmt.Sprintf("request.Array: %#v\n", elem))
	}
}

func (gen *Generator) renderRequestssOrTypeNull(g Printer, req *protocol.Request, or *protocol.OrType) {
	for _, item := range or.Items {
		if isNull(item) {
			continue
		}

		switch item := item.(type) {
		case *protocol.ReferenceType:
			pt, ok := normalizeHasLSPTypes(item.String())
			if !ok {
				g.P(`*`)
			}
			g.P(pt)

		case *protocol.ArrayType:
			gen.renderRequestsArrayType(g, req, item)

		default:
			panic(fmt.Sprintf("request.OrType: %[1]T = %#[1]v\n", item))
		}
	}
}

func (gen *Generator) renderRequestssOrType(g Printer, or *protocol.OrType, genericsProp *protocol.Property) {
	g.P(` *`, genericsProp.Name)
	gen.genericsTypes[genericsProp] = or.Items
}
