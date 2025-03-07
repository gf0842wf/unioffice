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
	"fmt"

	"github.com/gf0842wf/unioffice/schema/soo/dml"
)

type CT_Guide struct {
	// Guide Orientation
	OrientAttr ST_Direction
	// Guide Position
	PosAttr *dml.ST_Coordinate32
}

func NewCT_Guide() *CT_Guide {
	ret := &CT_Guide{}
	return ret
}

func (m *CT_Guide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OrientAttr != ST_DirectionUnset {
		attr, err := m.OrientAttr.MarshalXMLAttr(xml.Name{Local: "orient"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.PosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pos"},
			Value: fmt.Sprintf("%v", *m.PosAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Guide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "orient" {
			m.OrientAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "pos" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.PosAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Guide: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Guide and its children
func (m *CT_Guide) Validate() error {
	return m.ValidateWithPath("CT_Guide")
}

// ValidateWithPath validates the CT_Guide and its children, prefixing error messages with path
func (m *CT_Guide) ValidateWithPath(path string) error {
	if err := m.OrientAttr.ValidateWithPath(path + "/OrientAttr"); err != nil {
		return err
	}
	if m.PosAttr != nil {
		if err := m.PosAttr.ValidateWithPath(path + "/PosAttr"); err != nil {
			return err
		}
	}
	return nil
}
