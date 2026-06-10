// Copyright 2026 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"unsafe"

	"github.com/go-json-experiment/json/jsontext"
)

type ifaceWords struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}

var (
	progressTokenStringTab = func() unsafe.Pointer {
		var x ProgressToken = String("")
		return (*ifaceWords)(unsafe.Pointer(&x)).tab
	}()
	inlayHintTooltipStringTab = func() unsafe.Pointer {
		var x InlayHintTooltip = String("")
		return (*ifaceWords)(unsafe.Pointer(&x)).tab
	}()
)

func appendStringBox(boxes *[]String, v string) *String {
	*boxes = append(*boxes, String(v))
	return &(*boxes)[len(*boxes)-1]
}

func appendLocationBox(boxes *[]Location) *Location {
	*boxes = append(*boxes, Location{})
	return &(*boxes)[len(*boxes)-1]
}

func appendLocationURIOnlyBox(boxes *[]LocationUriOnly) *LocationUriOnly {
	*boxes = append(*boxes, LocationUriOnly{})
	return &(*boxes)[len(*boxes)-1]
}

func boxedStringProgressToken(p *String) ProgressToken {
	var x ProgressToken
	w := (*ifaceWords)(unsafe.Pointer(&x))
	// The slice backing array owns *p for the lifetime of every returned
	// interface value. Installing the cached itab preserves the public dynamic
	// type as String while avoiding one heap box per scalar union arm.
	w.tab = progressTokenStringTab
	w.data = unsafe.Pointer(p)
	return x
}

func boxedStringInlayHintTooltip(p *String) InlayHintTooltip {
	var x InlayHintTooltip
	w := (*ifaceWords)(unsafe.Pointer(&x))
	// The slice backing array owns *p for the lifetime of every returned
	// interface value. Installing the cached itab preserves the public dynamic
	// type as String while avoiding one heap box per scalar union arm.
	w.tab = inlayHintTooltipStringTab
	w.data = unsafe.Pointer(p)
	return x
}

func unmarshalProgressTokenValueBoxed(raw jsontext.Value, val *ProgressToken, scalarBoxes *[]String) error {
	switch raw.Kind() {
	case 'n':
		*val = nil
		return dvNullValue(raw)
	case '0':
		v, err := dvScalarInt32(raw)
		if err != nil {
			return err
		}
		*val = Integer(v)
		return nil
	case '"':
		v, err := dvScalarString(raw)
		if err != nil {
			return err
		}
		*val = boxedStringProgressToken(appendStringBox(scalarBoxes, v))
		return nil
	}
	return unmarshalProgressTokenValue(raw, val)
}

func unmarshalInlayHintTooltipValueBoxed(raw jsontext.Value, val *InlayHintTooltip, scalarBoxes *[]String) error {
	switch raw.Kind() {
	case 'n':
		*val = nil
		return dvNullValue(raw)
	case '"':
		v, err := dvScalarString(raw)
		if err != nil {
			return err
		}
		*val = boxedStringInlayHintTooltip(appendStringBox(scalarBoxes, v))
		return nil
	}
	return unmarshalInlayHintTooltipValue(raw, val)
}

func unmarshalWorkspaceSymbolLocationValueBoxed(raw jsontext.Value, val *WorkspaceSymbolLocation, locationBoxes *[]Location, locationURIOnlyBoxes *[]LocationUriOnly) error {
	switch raw.Kind() {
	case 'n':
		*val = nil
		return dvNullValue(raw)
	case '{':
		if objectHasAndKnownGuard(raw, []string{"uri", "range"}, []string{"uri", "range"}) {
			n := len(*locationBoxes)
			v := appendLocationBox(locationBoxes)
			if v.unmarshalLSPValue(raw) == nil {
				*val = v
				return nil
			}
			*locationBoxes = (*locationBoxes)[:n]
		}
		if objectHasAndKnownGuard(raw, []string{"uri"}, []string{"uri"}) {
			n := len(*locationURIOnlyBoxes)
			v := appendLocationURIOnlyBox(locationURIOnlyBoxes)
			if v.unmarshalLSPValue(raw) == nil {
				*val = v
				return nil
			}
			*locationURIOnlyBoxes = (*locationURIOnlyBoxes)[:n]
		}
		if objectHasKeys(raw, "uri", "range") {
			n := len(*locationBoxes)
			v := appendLocationBox(locationBoxes)
			if v.unmarshalLSPValue(raw) == nil {
				*val = v
				return nil
			}
			*locationBoxes = (*locationBoxes)[:n]
		}
		if objectHasKeys(raw, "uri") {
			n := len(*locationURIOnlyBoxes)
			v := appendLocationURIOnlyBox(locationURIOnlyBoxes)
			if v.unmarshalLSPValue(raw) == nil {
				*val = v
				return nil
			}
			*locationURIOnlyBoxes = (*locationURIOnlyBoxes)[:n]
		}
		{
			n := len(*locationBoxes)
			v := appendLocationBox(locationBoxes)
			if v.unmarshalLSPValue(raw) == nil {
				*val = v
				return nil
			}
			*locationBoxes = (*locationBoxes)[:n]
		}
		{
			n := len(*locationURIOnlyBoxes)
			v := appendLocationURIOnlyBox(locationURIOnlyBoxes)
			if v.unmarshalLSPValue(raw) == nil {
				*val = v
				return nil
			}
			*locationURIOnlyBoxes = (*locationURIOnlyBoxes)[:n]
		}
	}
	return unmarshalWorkspaceSymbolLocationValue(raw, val)
}
