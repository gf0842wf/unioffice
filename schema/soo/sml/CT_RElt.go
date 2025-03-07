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

	"github.com/gf0842wf/unioffice"
)

type CT_RElt struct {
	// Run Properties
	RPr *CT_RPrElt
	// Text
	T string
}

func NewCT_RElt() *CT_RElt {
	ret := &CT_RElt{}
	return ret
}

func (m *CT_RElt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "ma:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	set := xml.StartElement{Name: xml.Name{Local: "ma:t"}}
	unioffice.AddPreserveSpaceAttr(&set, m.T)
	e.EncodeElement(m.T, set)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RElt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RElt:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rPr"}:
				m.RPr = NewCT_RPrElt()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "t"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "t"}:
				if err := d.DecodeElement(&m.T, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_RElt %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RElt
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RElt and its children
func (m *CT_RElt) Validate() error {
	return m.ValidateWithPath("CT_RElt")
}

// ValidateWithPath validates the CT_RElt and its children, prefixing error messages with path
func (m *CT_RElt) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
