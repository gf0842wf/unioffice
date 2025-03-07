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
	"strconv"

	"github.com/gf0842wf/unioffice"
)

type CT_SlideTransition struct {
	// Transition Speed
	SpdAttr ST_TransitionSpeed
	// Advance on Click
	AdvClickAttr *bool
	// Advance after time
	AdvTmAttr *uint32
	Choice    *CT_SlideTransitionChoice
	// Sound Action
	SndAc  *CT_TransitionSoundAction
	ExtLst *CT_ExtensionListModify
}

func NewCT_SlideTransition() *CT_SlideTransition {
	ret := &CT_SlideTransition{}
	return ret
}

func (m *CT_SlideTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SpdAttr != ST_TransitionSpeedUnset {
		attr, err := m.SpdAttr.MarshalXMLAttr(xml.Name{Local: "spd"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AdvClickAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "advClick"},
			Value: fmt.Sprintf("%d", b2i(*m.AdvClickAttr))})
	}
	if m.AdvTmAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "advTm"},
			Value: fmt.Sprintf("%v", *m.AdvTmAttr)})
	}
	e.EncodeToken(start)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	if m.SndAc != nil {
		sesndAc := xml.StartElement{Name: xml.Name{Local: "p:sndAc"}}
		e.EncodeElement(m.SndAc, sesndAc)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spd" {
			m.SpdAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "advClick" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AdvClickAttr = &parsed
			continue
		}
		if attr.Name.Local == "advTm" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.AdvTmAttr = &pt
			continue
		}
	}
lCT_SlideTransition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "blinds"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "blinds"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Blinds, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "checker"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "checker"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Checker, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "circle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "circle"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Circle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "dissolve"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "dissolve"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Dissolve, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "comb"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "comb"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Comb, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cover"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cover"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Cover, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cut"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cut"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Cut, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "diamond"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "diamond"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Diamond, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "fade"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "fade"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Fade, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "newsflash"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "newsflash"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Newsflash, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "plus"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "plus"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Plus, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "pull"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "pull"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Pull, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "push"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "push"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Push, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "random"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "random"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Random, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "randomBar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "randomBar"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.RandomBar, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "split"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "split"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Split, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "strips"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "strips"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Strips, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "wedge"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "wedge"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Wedge, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "wheel"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "wheel"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Wheel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "wipe"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "wipe"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Wipe, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "zoom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "zoom"}:
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Zoom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sndAc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sndAc"}:
				m.SndAc = NewCT_TransitionSoundAction()
				if err := d.DecodeElement(m.SndAc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SlideTransition %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideTransition
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideTransition and its children
func (m *CT_SlideTransition) Validate() error {
	return m.ValidateWithPath("CT_SlideTransition")
}

// ValidateWithPath validates the CT_SlideTransition and its children, prefixing error messages with path
func (m *CT_SlideTransition) ValidateWithPath(path string) error {
	if err := m.SpdAttr.ValidateWithPath(path + "/SpdAttr"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.SndAc != nil {
		if err := m.SndAc.ValidateWithPath(path + "/SndAc"); err != nil {
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
