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

type CT_ConditionalFormats struct {
	// Conditional Format Count
	CountAttr *uint32
	// Conditional Formatting
	ConditionalFormat []*CT_ConditionalFormat
}

func NewCT_ConditionalFormats() *CT_ConditionalFormats {
	ret := &CT_ConditionalFormats{}
	return ret
}

func (m *CT_ConditionalFormats) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	seconditionalFormat := xml.StartElement{Name: xml.Name{Local: "ma:conditionalFormat"}}
	for _, c := range m.ConditionalFormat {
		e.EncodeElement(c, seconditionalFormat)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConditionalFormats) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_ConditionalFormats:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "conditionalFormat"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "conditionalFormat"}:
				tmp := NewCT_ConditionalFormat()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ConditionalFormat = append(m.ConditionalFormat, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_ConditionalFormats %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConditionalFormats
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConditionalFormats and its children
func (m *CT_ConditionalFormats) Validate() error {
	return m.ValidateWithPath("CT_ConditionalFormats")
}

// ValidateWithPath validates the CT_ConditionalFormats and its children, prefixing error messages with path
func (m *CT_ConditionalFormats) ValidateWithPath(path string) error {
	for i, v := range m.ConditionalFormat {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ConditionalFormat[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
