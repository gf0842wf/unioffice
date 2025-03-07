// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/gf0842wf/unioffice"
	"github.com/gf0842wf/unioffice/schema/soo/ofc/sharedTypes"
)

type OfcCT_OLEObject struct {
	TypeAttr       OfcST_OLEType
	ProgIDAttr     *string
	ShapeIDAttr    *string
	DrawAspectAttr OfcST_OLEDrawAspect
	ObjectIDAttr   *string
	IdAttr         *string
	UpdateModeAttr OfcST_OLEUpdateMode
	LinkType       *string
	LockedField    sharedTypes.ST_TrueFalseBlank
	FieldCodes     *string
}

func NewOfcCT_OLEObject() *OfcCT_OLEObject {
	ret := &OfcCT_OLEObject{}
	return ret
}

func (m *OfcCT_OLEObject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != OfcST_OLETypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "Type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ProgIDAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ProgID"},
			Value: fmt.Sprintf("%v", *m.ProgIDAttr)})
	}
	if m.ShapeIDAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ShapeID"},
			Value: fmt.Sprintf("%v", *m.ShapeIDAttr)})
	}
	if m.DrawAspectAttr != OfcST_OLEDrawAspectUnset {
		attr, err := m.DrawAspectAttr.MarshalXMLAttr(xml.Name{Local: "DrawAspect"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ObjectIDAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ObjectID"},
			Value: fmt.Sprintf("%v", *m.ObjectIDAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.UpdateModeAttr != OfcST_OLEUpdateModeUnset {
		attr, err := m.UpdateModeAttr.MarshalXMLAttr(xml.Name{Local: "UpdateMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.LinkType != nil {
		seLinkType := xml.StartElement{Name: xml.Name{Local: "o:LinkType"}}
		unioffice.AddPreserveSpaceAttr(&seLinkType, *m.LinkType)
		e.EncodeElement(m.LinkType, seLinkType)
	}
	if m.LockedField != sharedTypes.ST_TrueFalseBlankUnset {
		seLockedField := xml.StartElement{Name: xml.Name{Local: "o:LockedField"}}
		e.EncodeElement(m.LockedField, seLockedField)
	}
	if m.FieldCodes != nil {
		seFieldCodes := xml.StartElement{Name: xml.Name{Local: "o:FieldCodes"}}
		unioffice.AddPreserveSpaceAttr(&seFieldCodes, *m.FieldCodes)
		e.EncodeElement(m.FieldCodes, seFieldCodes)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_OLEObject) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "Type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ProgID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProgIDAttr = &parsed
			continue
		}
		if attr.Name.Local == "ShapeID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ShapeIDAttr = &parsed
			continue
		}
		if attr.Name.Local == "DrawAspect" {
			m.DrawAspectAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ObjectID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ObjectIDAttr = &parsed
			continue
		}
		if attr.Name.Local == "UpdateMode" {
			m.UpdateModeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lOfcCT_OLEObject:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "LinkType"}:
				m.LinkType = new(string)
				if err := d.DecodeElement(m.LinkType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "LockedField"}:
				m.LockedField = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.LockedField, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "FieldCodes"}:
				m.FieldCodes = new(string)
				if err := d.DecodeElement(m.FieldCodes, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on OfcCT_OLEObject %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcCT_OLEObject
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcCT_OLEObject and its children
func (m *OfcCT_OLEObject) Validate() error {
	return m.ValidateWithPath("OfcCT_OLEObject")
}

// ValidateWithPath validates the OfcCT_OLEObject and its children, prefixing error messages with path
func (m *OfcCT_OLEObject) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.DrawAspectAttr.ValidateWithPath(path + "/DrawAspectAttr"); err != nil {
		return err
	}
	if err := m.UpdateModeAttr.ValidateWithPath(path + "/UpdateModeAttr"); err != nil {
		return err
	}
	if err := m.LockedField.ValidateWithPath(path + "/LockedField"); err != nil {
		return err
	}
	return nil
}
