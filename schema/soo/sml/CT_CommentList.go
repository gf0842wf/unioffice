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

type CT_CommentList struct {
	// Comment
	Comment []*CT_Comment
}

func NewCT_CommentList() *CT_CommentList {
	ret := &CT_CommentList{}
	return ret
}

func (m *CT_CommentList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Comment != nil {
		secomment := xml.StartElement{Name: xml.Name{Local: "ma:comment"}}
		for _, c := range m.Comment {
			e.EncodeElement(c, secomment)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CommentList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CommentList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "comment"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "comment"}:
				tmp := NewCT_Comment()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Comment = append(m.Comment, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_CommentList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CommentList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CommentList and its children
func (m *CT_CommentList) Validate() error {
	return m.ValidateWithPath("CT_CommentList")
}

// ValidateWithPath validates the CT_CommentList and its children, prefixing error messages with path
func (m *CT_CommentList) ValidateWithPath(path string) error {
	for i, v := range m.Comment {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Comment[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
