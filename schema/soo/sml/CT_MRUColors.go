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

	"github.com/gf0842wf/unioffice"
)

type CT_MRUColors struct {
	// Color
	Color []*CT_Color
}

func NewCT_MRUColors() *CT_MRUColors {
	ret := &CT_MRUColors{}
	return ret
}

func (m *CT_MRUColors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secolor := xml.StartElement{Name: xml.Name{Local: "ma:color"}}
	for _, c := range m.Color {
		e.EncodeElement(c, secolor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MRUColors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MRUColors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "color"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "color"}:
				tmp := NewCT_Color()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Color = append(m.Color, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_MRUColors %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MRUColors
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MRUColors and its children
func (m *CT_MRUColors) Validate() error {
	return m.ValidateWithPath("CT_MRUColors")
}

// ValidateWithPath validates the CT_MRUColors and its children, prefixing error messages with path
func (m *CT_MRUColors) ValidateWithPath(path string) error {
	for i, v := range m.Color {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Color[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
