// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"

	"github.com/gf0842wf/unioffice"
)

type CT_TextLineBreak struct {
	RPr *CT_TextCharacterProperties
}

func NewCT_TextLineBreak() *CT_TextLineBreak {
	ret := &CT_TextLineBreak{}
	return ret
}

func (m *CT_TextLineBreak) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "a:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextLineBreak) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TextLineBreak:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "rPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "rPr"}:
				m.RPr = NewCT_TextCharacterProperties()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TextLineBreak %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextLineBreak
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextLineBreak and its children
func (m *CT_TextLineBreak) Validate() error {
	return m.ValidateWithPath("CT_TextLineBreak")
}

// ValidateWithPath validates the CT_TextLineBreak and its children, prefixing error messages with path
func (m *CT_TextLineBreak) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
