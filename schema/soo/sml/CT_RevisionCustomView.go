// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/gf0842wf/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_RevisionCustomView struct {
	// GUID
	GuidAttr string
	// User Action
	ActionAttr ST_RevisionAction
}

func NewCT_RevisionCustomView() *CT_RevisionCustomView {
	ret := &CT_RevisionCustomView{}
	ret.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	ret.ActionAttr = ST_RevisionAction(1)
	return ret
}

func (m *CT_RevisionCustomView) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "guid"},
		Value: fmt.Sprintf("%v", m.GuidAttr)})
	attr, err := m.ActionAttr.MarshalXMLAttr(xml.Name{Local: "action"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionCustomView) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	m.ActionAttr = ST_RevisionAction(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "guid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GuidAttr = parsed
			continue
		}
		if attr.Name.Local == "action" {
			m.ActionAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RevisionCustomView: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_RevisionCustomView and its children
func (m *CT_RevisionCustomView) Validate() error {
	return m.ValidateWithPath("CT_RevisionCustomView")
}

// ValidateWithPath validates the CT_RevisionCustomView and its children, prefixing error messages with path
func (m *CT_RevisionCustomView) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.GuidAttr) {
		return fmt.Errorf(`%s/m.GuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.GuidAttr)
	}
	if m.ActionAttr == ST_RevisionActionUnset {
		return fmt.Errorf("%s/ActionAttr is a mandatory field", path)
	}
	if err := m.ActionAttr.ValidateWithPath(path + "/ActionAttr"); err != nil {
		return err
	}
	return nil
}
