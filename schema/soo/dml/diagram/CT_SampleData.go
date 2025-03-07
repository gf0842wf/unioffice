// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gf0842wf/unioffice"
)

type CT_SampleData struct {
	UseDefAttr *bool
	DataModel  *CT_DataModel
}

func NewCT_SampleData() *CT_SampleData {
	ret := &CT_SampleData{}
	return ret
}

func (m *CT_SampleData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UseDefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "useDef"},
			Value: fmt.Sprintf("%d", b2i(*m.UseDefAttr))})
	}
	e.EncodeToken(start)
	if m.DataModel != nil {
		sedataModel := xml.StartElement{Name: xml.Name{Local: "dataModel"}}
		e.EncodeElement(m.DataModel, sedataModel)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SampleData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "useDef" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseDefAttr = &parsed
			continue
		}
	}
lCT_SampleData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "dataModel"}:
				m.DataModel = NewCT_DataModel()
				if err := d.DecodeElement(m.DataModel, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SampleData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SampleData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SampleData and its children
func (m *CT_SampleData) Validate() error {
	return m.ValidateWithPath("CT_SampleData")
}

// ValidateWithPath validates the CT_SampleData and its children, prefixing error messages with path
func (m *CT_SampleData) ValidateWithPath(path string) error {
	if m.DataModel != nil {
		if err := m.DataModel.ValidateWithPath(path + "/DataModel"); err != nil {
			return err
		}
	}
	return nil
}
