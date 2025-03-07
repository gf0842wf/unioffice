// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"

	"github.com/gf0842wf/unioffice"
)

type CT_DLbl struct {
	Idx    *CT_UnsignedInt
	Choice *CT_DLblChoice
	ExtLst *CT_ExtensionList
}

func NewCT_DLbl() *CT_DLbl {
	ret := &CT_DLbl{}
	ret.Idx = NewCT_UnsignedInt()
	return ret
}

func (m *CT_DLbl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seidx := xml.StartElement{Name: xml.Name{Local: "c:idx"}}
	e.EncodeElement(m.Idx, seidx)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DLbl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
lCT_DLbl:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "idx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "idx"}:
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "delete"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "delete"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.Delete, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "layout"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "layout"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.Layout, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "tx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "tx"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.Tx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numFmt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numFmt"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.NumFmt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "spPr"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "txPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "txPr"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.TxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLblPos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dLblPos"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.DLblPos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showLegendKey"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "showLegendKey"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowLegendKey, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showVal"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "showVal"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowVal, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showCatName"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "showCatName"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowCatName, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showSerName"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "showSerName"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowSerName, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showPercent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "showPercent"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowPercent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showBubbleSize"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "showBubbleSize"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowBubbleSize, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "separator"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "separator"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblChoice()
				}
				if err := d.DecodeElement(&m.Choice.Separator, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_DLbl %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DLbl
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DLbl and its children
func (m *CT_DLbl) Validate() error {
	return m.ValidateWithPath("CT_DLbl")
}

// ValidateWithPath validates the CT_DLbl and its children, prefixing error messages with path
func (m *CT_DLbl) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
