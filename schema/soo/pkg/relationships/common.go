// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package relationships

import (
	"encoding/xml"
	"fmt"

	"github.com/gf0842wf/unioffice"
)

type ST_TargetMode byte

const (
	ST_TargetModeUnset    ST_TargetMode = 0
	ST_TargetModeExternal ST_TargetMode = 1
	ST_TargetModeInternal ST_TargetMode = 2
)

func (e ST_TargetMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TargetModeUnset:
		attr.Value = ""
	case ST_TargetModeExternal:
		attr.Value = "External"
	case ST_TargetModeInternal:
		attr.Value = "Internal"
	}
	return attr, nil
}

func (e *ST_TargetMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "External":
		*e = 1
	case "Internal":
		*e = 2
	}
	return nil
}

func (m ST_TargetMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TargetMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "External":
			*m = 1
		case "Internal":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TargetMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "External"
	case 2:
		return "Internal"
	}
	return ""
}

func (m ST_TargetMode) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TargetMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/relationships", "CT_Relationships", NewCT_Relationships)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/relationships", "CT_Relationship", NewCT_Relationship)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/relationships", "Relationships", NewRelationships)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/relationships", "Relationship", NewRelationship)
}
