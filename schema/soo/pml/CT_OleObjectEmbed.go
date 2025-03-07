// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"

	"github.com/gf0842wf/unioffice"
)

type CT_OleObjectEmbed struct {
	// Color Scheme Properties for Embedded object
	FollowColorSchemeAttr ST_OleObjectFollowColorScheme
	ExtLst                *CT_ExtensionList
}

func NewCT_OleObjectEmbed() *CT_OleObjectEmbed {
	ret := &CT_OleObjectEmbed{}
	return ret
}

func (m *CT_OleObjectEmbed) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FollowColorSchemeAttr != ST_OleObjectFollowColorSchemeUnset {
		attr, err := m.FollowColorSchemeAttr.MarshalXMLAttr(xml.Name{Local: "followColorScheme"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleObjectEmbed) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "followColorScheme" {
			m.FollowColorSchemeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_OleObjectEmbed:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_OleObjectEmbed %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleObjectEmbed
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleObjectEmbed and its children
func (m *CT_OleObjectEmbed) Validate() error {
	return m.ValidateWithPath("CT_OleObjectEmbed")
}

// ValidateWithPath validates the CT_OleObjectEmbed and its children, prefixing error messages with path
func (m *CT_OleObjectEmbed) ValidateWithPath(path string) error {
	if err := m.FollowColorSchemeAttr.ValidateWithPath(path + "/FollowColorSchemeAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
