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
	"fmt"
	"strconv"

	"github.com/gf0842wf/unioffice"
)

type CT_Camera struct {
	PrstAttr ST_PresetCameraType
	FovAttr  *int32
	ZoomAttr *ST_PositivePercentage
	Rot      *CT_SphereCoords
}

func NewCT_Camera() *CT_Camera {
	ret := &CT_Camera{}
	ret.PrstAttr = ST_PresetCameraType(1)
	return ret
}

func (m *CT_Camera) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.PrstAttr.MarshalXMLAttr(xml.Name{Local: "prst"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.FovAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fov"},
			Value: fmt.Sprintf("%v", *m.FovAttr)})
	}
	if m.ZoomAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoom"},
			Value: fmt.Sprintf("%v", *m.ZoomAttr)})
	}
	e.EncodeToken(start)
	if m.Rot != nil {
		serot := xml.StartElement{Name: xml.Name{Local: "a:rot"}}
		e.EncodeElement(m.Rot, serot)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Camera) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PrstAttr = ST_PresetCameraType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "prst" {
			m.PrstAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "fov" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.FovAttr = &pt
			continue
		}
		if attr.Name.Local == "zoom" {
			parsed, err := ParseUnionST_PositivePercentage(attr.Value)
			if err != nil {
				return err
			}
			m.ZoomAttr = &parsed
			continue
		}
	}
lCT_Camera:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "rot"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "rot"}:
				m.Rot = NewCT_SphereCoords()
				if err := d.DecodeElement(m.Rot, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Camera %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Camera
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Camera and its children
func (m *CT_Camera) Validate() error {
	return m.ValidateWithPath("CT_Camera")
}

// ValidateWithPath validates the CT_Camera and its children, prefixing error messages with path
func (m *CT_Camera) ValidateWithPath(path string) error {
	if m.PrstAttr == ST_PresetCameraTypeUnset {
		return fmt.Errorf("%s/PrstAttr is a mandatory field", path)
	}
	if err := m.PrstAttr.ValidateWithPath(path + "/PrstAttr"); err != nil {
		return err
	}
	if m.FovAttr != nil {
		if *m.FovAttr < 0 {
			return fmt.Errorf("%s/m.FovAttr must be >= 0 (have %v)", path, *m.FovAttr)
		}
		if *m.FovAttr > 10800000 {
			return fmt.Errorf("%s/m.FovAttr must be <= 10800000 (have %v)", path, *m.FovAttr)
		}
	}
	if m.ZoomAttr != nil {
		if err := m.ZoomAttr.ValidateWithPath(path + "/ZoomAttr"); err != nil {
			return err
		}
	}
	if m.Rot != nil {
		if err := m.Rot.ValidateWithPath(path + "/Rot"); err != nil {
			return err
		}
	}
	return nil
}
