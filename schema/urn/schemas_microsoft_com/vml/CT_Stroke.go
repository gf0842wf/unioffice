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
	"strconv"

	"github.com/gf0842wf/unioffice"
	"github.com/gf0842wf/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_Stroke struct {
	Left                 *OfcLeft
	Top                  *OfcTop
	Right                *OfcRight
	Bottom               *OfcBottom
	Column               *OfcColumn
	IdAttr               *string
	OnAttr               sharedTypes.ST_TrueFalse
	WeightAttr           *string
	ColorAttr            *string
	OpacityAttr          *string
	LinestyleAttr        ST_StrokeLineStyle
	MiterlimitAttr       *float64
	JoinstyleAttr        ST_StrokeJoinStyle
	EndcapAttr           ST_StrokeEndCap
	DashstyleAttr        *string
	FilltypeAttr         ST_FillType
	SrcAttr              *string
	ImageaspectAttr      ST_ImageAspect
	ImagesizeAttr        *string
	ImagealignshapeAttr  sharedTypes.ST_TrueFalse
	Color2Attr           *string
	StartarrowAttr       ST_StrokeArrowType
	StartarrowwidthAttr  ST_StrokeArrowWidth
	StartarrowlengthAttr ST_StrokeArrowLength
	EndarrowAttr         ST_StrokeArrowType
	EndarrowwidthAttr    ST_StrokeArrowWidth
	EndarrowlengthAttr   ST_StrokeArrowLength
	HrefAttr             *string
	AlthrefAttr          *string
	TitleAttr            *string
	ForcedashAttr        sharedTypes.ST_TrueFalse
	RIdAttr              *string
	InsetpenAttr         sharedTypes.ST_TrueFalse
	RelidAttr            *string
}

func NewCT_Stroke() *CT_Stroke {
	ret := &CT_Stroke{}
	return ret
}

func (m *CT_Stroke) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.OnAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OnAttr.MarshalXMLAttr(xml.Name{Local: "on"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.WeightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "weight"},
			Value: fmt.Sprintf("%v", *m.WeightAttr)})
	}
	if m.ColorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "color"},
			Value: fmt.Sprintf("%v", *m.ColorAttr)})
	}
	if m.OpacityAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "opacity"},
			Value: fmt.Sprintf("%v", *m.OpacityAttr)})
	}
	if m.LinestyleAttr != ST_StrokeLineStyleUnset {
		attr, err := m.LinestyleAttr.MarshalXMLAttr(xml.Name{Local: "linestyle"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.MiterlimitAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "miterlimit"},
			Value: fmt.Sprintf("%v", *m.MiterlimitAttr)})
	}
	if m.JoinstyleAttr != ST_StrokeJoinStyleUnset {
		attr, err := m.JoinstyleAttr.MarshalXMLAttr(xml.Name{Local: "joinstyle"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EndcapAttr != ST_StrokeEndCapUnset {
		attr, err := m.EndcapAttr.MarshalXMLAttr(xml.Name{Local: "endcap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DashstyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dashstyle"},
			Value: fmt.Sprintf("%v", *m.DashstyleAttr)})
	}
	if m.FilltypeAttr != ST_FillTypeUnset {
		attr, err := m.FilltypeAttr.MarshalXMLAttr(xml.Name{Local: "filltype"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SrcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "src"},
			Value: fmt.Sprintf("%v", *m.SrcAttr)})
	}
	if m.ImageaspectAttr != ST_ImageAspectUnset {
		attr, err := m.ImageaspectAttr.MarshalXMLAttr(xml.Name{Local: "imageaspect"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ImagesizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "imagesize"},
			Value: fmt.Sprintf("%v", *m.ImagesizeAttr)})
	}
	if m.ImagealignshapeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ImagealignshapeAttr.MarshalXMLAttr(xml.Name{Local: "imagealignshape"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.Color2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "color2"},
			Value: fmt.Sprintf("%v", *m.Color2Attr)})
	}
	if m.StartarrowAttr != ST_StrokeArrowTypeUnset {
		attr, err := m.StartarrowAttr.MarshalXMLAttr(xml.Name{Local: "startarrow"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StartarrowwidthAttr != ST_StrokeArrowWidthUnset {
		attr, err := m.StartarrowwidthAttr.MarshalXMLAttr(xml.Name{Local: "startarrowwidth"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StartarrowlengthAttr != ST_StrokeArrowLengthUnset {
		attr, err := m.StartarrowlengthAttr.MarshalXMLAttr(xml.Name{Local: "startarrowlength"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EndarrowAttr != ST_StrokeArrowTypeUnset {
		attr, err := m.EndarrowAttr.MarshalXMLAttr(xml.Name{Local: "endarrow"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EndarrowwidthAttr != ST_StrokeArrowWidthUnset {
		attr, err := m.EndarrowwidthAttr.MarshalXMLAttr(xml.Name{Local: "endarrowwidth"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EndarrowlengthAttr != ST_StrokeArrowLengthUnset {
		attr, err := m.EndarrowlengthAttr.MarshalXMLAttr(xml.Name{Local: "endarrowlength"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:href"},
			Value: fmt.Sprintf("%v", *m.HrefAttr)})
	}
	if m.AlthrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:althref"},
			Value: fmt.Sprintf("%v", *m.AlthrefAttr)})
	}
	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}
	if m.ForcedashAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ForcedashAttr.MarshalXMLAttr(xml.Name{Local: "forcedash"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.RIdAttr)})
	}
	if m.InsetpenAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.InsetpenAttr.MarshalXMLAttr(xml.Name{Local: "insetpen"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RelidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:relid"},
			Value: fmt.Sprintf("%v", *m.RelidAttr)})
	}
	e.EncodeToken(start)
	if m.Left != nil {
		seleft := xml.StartElement{Name: xml.Name{Local: "o:left"}}
		e.EncodeElement(m.Left, seleft)
	}
	if m.Top != nil {
		setop := xml.StartElement{Name: xml.Name{Local: "o:top"}}
		e.EncodeElement(m.Top, setop)
	}
	if m.Right != nil {
		seright := xml.StartElement{Name: xml.Name{Local: "o:right"}}
		e.EncodeElement(m.Right, seright)
	}
	if m.Bottom != nil {
		sebottom := xml.StartElement{Name: xml.Name{Local: "o:bottom"}}
		e.EncodeElement(m.Bottom, sebottom)
	}
	if m.Column != nil {
		secolumn := xml.StartElement{Name: xml.Name{Local: "o:column"}}
		e.EncodeElement(m.Column, secolumn)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Stroke) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "relid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RelidAttr = &parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RIdAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "forcedash" {
			m.ForcedashAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "althref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlthrefAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
			continue
		}
		if attr.Name.Local == "imagesize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ImagesizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "startarrowlength" {
			m.StartarrowlengthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "joinstyle" {
			m.JoinstyleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "endarrow" {
			m.EndarrowAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dashstyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DashstyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "endarrowwidth" {
			m.EndarrowwidthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "src" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SrcAttr = &parsed
			continue
		}
		if attr.Name.Local == "endarrowlength" {
			m.EndarrowlengthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "linestyle" {
			m.LinestyleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "opacity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OpacityAttr = &parsed
			continue
		}
		if attr.Name.Local == "weight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WeightAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "miterlimit" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MiterlimitAttr = &parsed
			continue
		}
		if attr.Name.Local == "color" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
			continue
		}
		if attr.Name.Local == "imageaspect" {
			m.ImageaspectAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "filltype" {
			m.FilltypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "startarrowwidth" {
			m.StartarrowwidthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "insetpen" {
			m.InsetpenAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "imagealignshape" {
			m.ImagealignshapeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "startarrow" {
			m.StartarrowAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "endcap" {
			m.EndcapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "color2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Color2Attr = &parsed
			continue
		}
	}
lCT_Stroke:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "left"}:
				m.Left = NewOfcLeft()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "top"}:
				m.Top = NewOfcTop()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "right"}:
				m.Right = NewOfcRight()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "bottom"}:
				m.Bottom = NewOfcBottom()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "column"}:
				m.Column = NewOfcColumn()
				if err := d.DecodeElement(m.Column, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Stroke %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Stroke
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Stroke and its children
func (m *CT_Stroke) Validate() error {
	return m.ValidateWithPath("CT_Stroke")
}

// ValidateWithPath validates the CT_Stroke and its children, prefixing error messages with path
func (m *CT_Stroke) ValidateWithPath(path string) error {
	if m.Left != nil {
		if err := m.Left.ValidateWithPath(path + "/Left"); err != nil {
			return err
		}
	}
	if m.Top != nil {
		if err := m.Top.ValidateWithPath(path + "/Top"); err != nil {
			return err
		}
	}
	if m.Right != nil {
		if err := m.Right.ValidateWithPath(path + "/Right"); err != nil {
			return err
		}
	}
	if m.Bottom != nil {
		if err := m.Bottom.ValidateWithPath(path + "/Bottom"); err != nil {
			return err
		}
	}
	if m.Column != nil {
		if err := m.Column.ValidateWithPath(path + "/Column"); err != nil {
			return err
		}
	}
	if err := m.OnAttr.ValidateWithPath(path + "/OnAttr"); err != nil {
		return err
	}
	if err := m.LinestyleAttr.ValidateWithPath(path + "/LinestyleAttr"); err != nil {
		return err
	}
	if err := m.JoinstyleAttr.ValidateWithPath(path + "/JoinstyleAttr"); err != nil {
		return err
	}
	if err := m.EndcapAttr.ValidateWithPath(path + "/EndcapAttr"); err != nil {
		return err
	}
	if err := m.FilltypeAttr.ValidateWithPath(path + "/FilltypeAttr"); err != nil {
		return err
	}
	if err := m.ImageaspectAttr.ValidateWithPath(path + "/ImageaspectAttr"); err != nil {
		return err
	}
	if err := m.ImagealignshapeAttr.ValidateWithPath(path + "/ImagealignshapeAttr"); err != nil {
		return err
	}
	if err := m.StartarrowAttr.ValidateWithPath(path + "/StartarrowAttr"); err != nil {
		return err
	}
	if err := m.StartarrowwidthAttr.ValidateWithPath(path + "/StartarrowwidthAttr"); err != nil {
		return err
	}
	if err := m.StartarrowlengthAttr.ValidateWithPath(path + "/StartarrowlengthAttr"); err != nil {
		return err
	}
	if err := m.EndarrowAttr.ValidateWithPath(path + "/EndarrowAttr"); err != nil {
		return err
	}
	if err := m.EndarrowwidthAttr.ValidateWithPath(path + "/EndarrowwidthAttr"); err != nil {
		return err
	}
	if err := m.EndarrowlengthAttr.ValidateWithPath(path + "/EndarrowlengthAttr"); err != nil {
		return err
	}
	if err := m.ForcedashAttr.ValidateWithPath(path + "/ForcedashAttr"); err != nil {
		return err
	}
	if err := m.InsetpenAttr.ValidateWithPath(path + "/InsetpenAttr"); err != nil {
		return err
	}
	return nil
}
