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
	"strconv"

	"github.com/gf0842wf/unioffice"
)

type CT_ExternalRow struct {
	// Row
	RAttr uint32
	// External Cell Data
	Cell []*CT_ExternalCell
}

func NewCT_ExternalRow() *CT_ExternalRow {
	ret := &CT_ExternalRow{}
	return ret
}

func (m *CT_ExternalRow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	e.EncodeToken(start)
	if m.Cell != nil {
		secell := xml.StartElement{Name: xml.Name{Local: "ma:cell"}}
		for _, c := range m.Cell {
			e.EncodeElement(c, secell)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExternalRow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.RAttr = uint32(parsed)
			continue
		}
	}
lCT_ExternalRow:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cell"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cell"}:
				tmp := NewCT_ExternalCell()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cell = append(m.Cell, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_ExternalRow %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalRow
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExternalRow and its children
func (m *CT_ExternalRow) Validate() error {
	return m.ValidateWithPath("CT_ExternalRow")
}

// ValidateWithPath validates the CT_ExternalRow and its children, prefixing error messages with path
func (m *CT_ExternalRow) ValidateWithPath(path string) error {
	for i, v := range m.Cell {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cell[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
