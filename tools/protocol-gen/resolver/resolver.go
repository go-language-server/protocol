// Copyright 2024 The Go Language Server Authors
// SPDX-License-Identifier: Apache-2.0

package resolver

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gobuffalo/flect"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"go.lsp.dev/protocol/tools/protocol-gen/protocol"
	"go.lsp.dev/protocol/tools/protocol-gen/schema"
)

// Resolve resolves the JSON MetaModel to a Protocol, which can be consumed by the templates
func Resolve(model schema.MetaModel) (*protocol.Protocol, error) {
	r := resolver{}
	out := r.model(model)
	if r.err != nil {
		return &protocol.Protocol{}, r.err
	}
	return out, nil
}

type resolver struct {
	stack                   []string
	err                     error
	newStructureLiteralType func(*schema.StructureLiteralType) protocol.Type
	allReferenceTypes       []*protocol.ReferenceType
	typeDecls               map[string]protocol.TypeDecl
	typeDeclsSeen           map[string]bool
}

func (r *resolver) pushScope(msg string, args ...any) (pop func()) {
	r.stack = append(r.stack, fmt.Sprintf(msg, args...))
	return func() { r.stack = r.stack[:len(r.stack)-1] }
}

func (r *resolver) error(msg string, args ...any) {
	if r.err != nil {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(msg, args...))
	for i := len(r.stack) - 1; i >= 0; i-- {
		sb.WriteString("\nwhile ")
		sb.WriteString(r.stack[i])
	}
	r.err = fmt.Errorf("%v", sb.String())
}

func (r *resolver) model(in schema.MetaModel) *protocol.Protocol {
	defer r.pushScope("resolving model")()
	out := &protocol.Protocol{
		Enumerations:                transform(in.Enumerations, r.enumeration),
		MetaData:                    r.metadata(in.MetaData),
		Notifications:               transform(in.Notifications, r.notification),
		ClientToServerNotifications: []*protocol.Notification{},
		ServerToClientNotifications: []*protocol.Notification{},
		BidirectionalNotifications:  []*protocol.Notification{},
		Requests:                    transform(in.Requests, r.request),
		Structures:                  transform(in.Structures, r.structure),
		TypeAliases:                 transform(in.TypeAliases, r.typeAlias),
		ServerToClientRequests:      []*protocol.Request{},
		ClientToServerRequests:      []*protocol.Request{},
		BidirectionalRequests:       []*protocol.Request{},
	}

	r.resolveTypes(out)

	for _, notification := range out.Notifications {
		switch notification.MessageDirection {
		case protocol.MessageDirectionClientToServer:
			out.ClientToServerNotifications = append(out.ClientToServerNotifications, notification)
		case protocol.MessageDirectionServerToClient:
			out.ServerToClientNotifications = append(out.ServerToClientNotifications, notification)
		case protocol.MessageDirectionBidirectional:
			out.BidirectionalNotifications = append(out.BidirectionalNotifications, notification)
		}
	}
	for _, request := range out.Requests {
		switch request.MessageDirection {
		case protocol.MessageDirectionClientToServer:
			out.ClientToServerRequests = append(out.ClientToServerRequests, request)
		case protocol.MessageDirectionServerToClient:
			out.ServerToClientRequests = append(out.ServerToClientRequests, request)
		case protocol.MessageDirectionBidirectional:
			out.BidirectionalRequests = append(out.BidirectionalRequests, request)
		}
	}
	out.TypeAliases = r.sortTypeAliases(out.TypeAliases)
	out.Structures = r.sortStructures(out.Structures)

	return out
}

func (r *resolver) enumeration(in schema.Enumeration) *protocol.Enumeration {
	defer r.pushScope("resolving enumeration '%v'", in.Name)()
	return &protocol.Enumeration{
		Deprecated:           in.Deprecated,
		Documentation:        r.documentation(in.Documentation),
		Name:                 r.className(in.Name),
		Proposed:             in.Proposed,
		Since:                in.Since,
		SupportsCustomValues: in.SupportsCustomValues,
		Type:                 r.type_(in.Type),
		Values:               transform(in.Values, r.enumerationEntry),
	}
}

func (r *resolver) enumerationEntry(in schema.EnumerationEntry) *protocol.EnumerationEntry {
	defer r.pushScope("resolving enumeration entry '%v'", in.Name)()
	return &protocol.EnumerationEntry{
		Deprecated:    in.Deprecated,
		Documentation: r.documentation(in.Documentation),
		Name:          r.className(in.Name),
		Proposed:      in.Proposed,
		Since:         in.Since,
		Value:         r.value(in.Value),
	}
}

func (r *resolver) metadata(in schema.MetaData) *protocol.MetaData {
	defer r.pushScope("resolving metadata")()
	return &protocol.MetaData{Version: in.Version}
}

func (r *resolver) notification(in schema.Notification) *protocol.Notification {
	defer r.pushScope("resolving notification '%v'", in.Method)()
	return &protocol.Notification{
		Deprecated:          in.Deprecated,
		Documentation:       r.documentation(in.Documentation),
		MessageDirection:    r.messageDirection(in.MessageDirection),
		Method:              in.Method,
		Params:              r.types(in.Params),
		Proposed:            in.Proposed,
		RegistrationMethod:  in.RegistrationMethod,
		RegistrationOptions: r.type_(in.RegistrationOptions),
		Since:               in.Since,
	}
}

func (r *resolver) request(in schema.Request) *protocol.Request {
	defer r.pushScope("resolving request '%v'", in.Method)()
	return &protocol.Request{
		Deprecated:          in.Deprecated,
		Documentation:       r.documentation(in.Documentation),
		ErrorData:           r.type_(in.ErrorData),
		MessageDirection:    r.messageDirection(in.MessageDirection),
		Method:              in.Method,
		Params:              r.types(in.Params),
		PartialResult:       r.type_(in.PartialResult),
		Proposed:            in.Proposed,
		RegistrationMethod:  in.RegistrationMethod,
		RegistrationOptions: r.type_(in.RegistrationOptions),
		Result:              r.type_(in.Result),
		Since:               in.Since,
	}
}

func (r *resolver) structure(in schema.Structure) *protocol.Structure {
	defer r.pushScope("resolving structure '%v'", in.Name)()
	name := r.className(in.Name)
	out := &protocol.Structure{
		Deprecated:    in.Deprecated,
		Documentation: r.documentation(in.Documentation),
		Extends:       transform(in.Extends, r.type_),
		Mixins:        transform(in.Mixins, r.type_),
		Name:          name,
		Proposed:      in.Proposed,
		Since:         in.Since,
		NestedNames:   []string{name},
	}
	for _, propertyIn := range in.Properties {
		defer scopedAssignment(&r.newStructureLiteralType, func(in *schema.StructureLiteralType) protocol.Type {
			name := cases.Title(language.Und, cases.NoLower).String(propertyIn.Name)
			out.NestedStructures = append(out.NestedStructures,
				&protocol.Structure{
					Deprecated:    in.Value.Deprecated,
					Documentation: r.documentation(in.Value.Documentation),
					Properties:    transform(in.Value.Properties, r.property),
					Name:          name,
					Proposed:      in.Value.Proposed,
					Since:         in.Value.Since,
					NestedNames:   append(append([]string{}, out.NestedNames...), name),
				},
			)
			ref := &protocol.ReferenceType{Name: name}
			r.allReferenceTypes = append(r.allReferenceTypes, ref)
			return ref
		})()
		propertyOut := r.property(propertyIn)
		if propertyOut.JSONName == "kind" {
			if lit, ok := propertyOut.Type.(*protocol.StringLiteralType); ok {
				out.Kind = lit.Value
				continue
			}
		}
		out.Properties = append(out.Properties, propertyOut)
	}
	return out
}

func (r *resolver) property(in schema.Property) *protocol.Property {
	defer r.pushScope("resolving property '%v'", in.Name)()
	return &protocol.Property{
		Deprecated:    in.Deprecated,
		Documentation: r.documentation(in.Documentation),
		JSONName:      in.Name,
		Name:          flect.Underscore(in.Name),
		Optional:      in.Optional,
		Proposed:      in.Proposed,
		Since:         in.Since,
		Type:          r.type_(in.Type),
	}
}

func (r *resolver) typeAlias(in schema.TypeAlias) *protocol.TypeAlias {
	defer r.pushScope("resolving type alias '%v'", in.Name)()
	return &protocol.TypeAlias{
		Deprecated:    in.Deprecated,
		Documentation: r.documentation(in.Documentation),
		Name:          r.className(in.Name),
		Proposed:      in.Proposed,
		Since:         in.Since,
		Type:          r.type_(in.Type),
	}
}

func (r *resolver) types(in schema.Nodes) []protocol.Type {
	return transform(in.Nodes, r.typeImpl)
}

func (r *resolver) type_(in schema.Type) protocol.Type {
	return r.typeImpl(in.Node)
}

func (r *resolver) value(in any) any {
	switch in := in.(type) {
	case string:
		return fmt.Sprintf(`"%v"`, in)
	default:
		return in
	}
}

func (r *resolver) typeImpl(in schema.Node) protocol.Type {
	switch in := in.(type) {
	case nil:
		return nil
	case *schema.BaseType:
		switch in.Name {
		case schema.URI:
			return &protocol.URIType{}
		case schema.DocumentUri:
			return &protocol.DocumentUriType{}
		case schema.Integer:
			return &protocol.IntegerType{}
		case schema.Uinteger:
			return &protocol.UintegerType{}
		case schema.Decimal:
			return &protocol.DecimalType{}
		case schema.RegExp:
			return &protocol.RegExpType{}
		case schema.String:
			return &protocol.StringType{}
		case schema.Boolean:
			return &protocol.BooleanType{}
		case schema.Null:
			return &protocol.NullType{}
		}

	case *schema.ArrayType:
		return &protocol.ArrayType{Element: r.type_(in.Element)}

	case *schema.ReferenceType:
		out := &protocol.ReferenceType{Name: r.className(in.Name)}
		r.allReferenceTypes = append(r.allReferenceTypes, out)
		return out

	case *schema.AndType:
		return &protocol.AndType{Items: transform(in.Items, r.type_)}

	case *schema.OrType:
		return &protocol.OrType{Items: transform(in.Items, r.type_)}

	case *schema.MapType:
		return &protocol.MapType{Key: r.type_(in.Key), Value: r.type_(in.Value)}

	case *schema.StringLiteralType:
		return &protocol.StringLiteralType{Value: in.Value}

	case *schema.StructureLiteralType:
		return r.newStructureLiteralType(in)

	case *schema.TupleType:
		return &protocol.TupleType{Items: transform(in.Items, r.type_)}
	}
	r.error("invalid type %T %+v", in, in)
	return nil
}

func (r *resolver) structureLiteral(in schema.StructureLiteral) *protocol.StructureLiteral {
	defer r.pushScope("resolving structure literal")()
	return &protocol.StructureLiteral{
		Deprecated:    in.Deprecated,
		Documentation: r.documentation(in.Documentation),
		Properties:    transform(in.Properties, r.property),
		Proposed:      in.Proposed,
		Since:         in.Since,
	}
}

func (r *resolver) messageDirection(in schema.MessageDirection) protocol.MessageDirection {
	switch in {
	case schema.MessageDirectionServerToClient:
		return protocol.MessageDirectionServerToClient
	case schema.MessageDirectionClientToServer:
		return protocol.MessageDirectionClientToServer
	case schema.MessageDirectionBoth:
		return protocol.MessageDirectionBidirectional
	}
	r.error("invalid message direction %+v", in)
	return ""
}

func (r *resolver) className(name string) string {
	if strings.HasPrefix(name, "_") {
		name = strings.TrimLeft(name, "_") + "Base"
	}
	name = strings.TrimLeft(name, "$")
	name = strings.ReplaceAll(name, "/", "_")
	name = flect.Capitalize(name)
	return name
}

var reLinkTag = regexp.MustCompile(`{@link[\s]+([^}]+)}`)

func (r *resolver) documentation(in string) string {
	s := reLinkTag.ReplaceAllString(in, "$1")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.ReplaceAll(s, "   ", " ")
	// s = strings.ReplaceAll(s, "@proposed", "\n\nProposed in:")
	s = strings.ReplaceAll(s, "@sample", "\n\nExample:")
	// s = strings.ReplaceAll(s, "@since", "\n\n@since")
	s = strings.ReplaceAll(s, "@deprecated", "\n\nDeprecated:")

	return strings.TrimSpace(s)
}

func (r *resolver) resolveTypes(p *protocol.Protocol) {
	r.typeDecls = map[string]protocol.TypeDecl{}
	register := func(name string, ty protocol.TypeDecl) {
		if existing, found := r.typeDecls[name]; found {
			r.error("duplicate definition for '%v'. %T and %T", name, ty, existing)
			return
		}
		r.typeDecls[name] = ty
	}
	lookup := func(name string) protocol.TypeDecl {
		typeDecl, found := r.typeDecls[name]
		if !found {
			r.error("referenced type '%v' not found", name)
		}
		return typeDecl
	}

	for _, a := range p.TypeAliases {
		register(a.Name, a)
	}
	for _, e := range p.Enumerations {
		register(e.Name, e)
	}
	for _, s := range p.Structures {
		register(s.Name, s)
		for _, n := range s.NestedStructures {
			register(s.Name+"::"+n.Name, s)
		}
	}

	for _, a := range r.allReferenceTypes {
		a.TypeDecl = lookup(a.Name)
	}
}

func (r *resolver) sortTypeAliases(in []*protocol.TypeAlias) []*protocol.TypeAlias {
	aliases := map[string]*protocol.TypeAlias{}
	for _, a := range in {
		aliases[a.Name] = a
	}

	sorted := make([]*protocol.TypeAlias, 0, len(in))
	seen := map[string]struct{}{}
	stack := append([]*protocol.TypeAlias{}, in...)

	for len(stack) > 0 {
		alias := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, found := seen[alias.Name]; found {
			continue
		}
		seen[alias.Name] = struct{}{}
		for _, ty := range protocol.RecursiveTypesOf(alias.Type) {
			if ref, ok := ty.(*protocol.ReferenceType); ok {
				if dep, ok := aliases[ref.Name]; ok {
					stack = append(stack, dep)
				}
			}
		}
		sorted = append(sorted, alias)
	}

	return sorted
}

func (r *resolver) sortStructures(in []*protocol.Structure) []*protocol.Structure {
	structures := map[string]*protocol.Structure{}
	for _, s := range in {
		structures[s.Name] = s
	}

	sorted := []*protocol.Structure{}
	seen := map[string]struct{}{}

	var visit func(s *protocol.Structure)
	visit = func(s *protocol.Structure) {
		if _, found := seen[s.Name]; found {
			return
		}
		seen[s.Name] = struct{}{}
		for _, ext := range s.Extends {
			if ref, ok := ext.(*protocol.ReferenceType); ok {
				if dep, ok := structures[ref.Name]; ok {
					visit(dep)
				}
			}
		}
		for _, property := range s.Properties {
			for _, ty := range protocol.RecursiveTypesOf(property.Type) {
				if ref, ok := ty.(*protocol.ReferenceType); ok {
					if dep, ok := structures[ref.Name]; ok {
						visit(dep)
					}
				}
			}
		}
		sorted = append(sorted, s)
	}

	for _, s := range in {
		visit(s)
	}
	return sorted
}

func scopedAssignment[T any](p *T, val T) func() {
	old := *p
	*p = val
	return func() { *p = old }
}

// transform returns a new slice by transforming each element with the function fn
func transform[IN, OUT any](in []IN, fn func(in IN) OUT) []OUT {
	out := make([]OUT, len(in))
	for i, el := range in {
		out[i] = fn(el)
	}
	return out
}
