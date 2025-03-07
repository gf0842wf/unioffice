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

type HandoutMaster struct {
	CT_HandoutMaster
}

func NewHandoutMaster() *HandoutMaster {
	ret := &HandoutMaster{}
	ret.CT_HandoutMaster = *NewCT_HandoutMaster()
	return ret
}

func (m *HandoutMaster) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:handoutMaster"
	return m.CT_HandoutMaster.MarshalXML(e, start)
}

func (m *HandoutMaster) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_HandoutMaster = *NewCT_HandoutMaster()
lHandoutMaster:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cSld"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cSld"}:
				if err := d.DecodeElement(m.CSld, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "clrMap"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "clrMap"}:
				if err := d.DecodeElement(m.ClrMap, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "hf"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "hf"}:
				m.Hf = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.Hf, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on HandoutMaster %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lHandoutMaster
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the HandoutMaster and its children
func (m *HandoutMaster) Validate() error {
	return m.ValidateWithPath("HandoutMaster")
}

// ValidateWithPath validates the HandoutMaster and its children, prefixing error messages with path
func (m *HandoutMaster) ValidateWithPath(path string) error {
	if err := m.CT_HandoutMaster.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
