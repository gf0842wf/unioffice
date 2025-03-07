// Copyright 2018 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document

import (
	"github.com/gf0842wf/unioffice"
	"github.com/gf0842wf/unioffice/measurement"
	"github.com/gf0842wf/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/gf0842wf/unioffice/schema/soo/wml"
)

// ParagraphSpacing controls the spacing for a paragraph and its lines.
type ParagraphSpacing struct {
	x *wml.CT_Spacing
}

// SetBefore sets the spacing that comes before the paragraph.
func (p ParagraphSpacing) SetBefore(before measurement.Distance) {
	p.x.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.BeforeAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(before / measurement.Twips))
}

// SetAfter sets the spacing that comes after the paragraph.
func (p ParagraphSpacing) SetAfter(after measurement.Distance) {
	p.x.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.AfterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(after / measurement.Twips))
}

// SetLineSpacing sets the spacing between lines in a paragraph.
func (p ParagraphSpacing) SetLineSpacing(d measurement.Distance, rule wml.ST_LineSpacingRule) {
	if rule == wml.ST_LineSpacingRuleUnset {
		p.x.LineRuleAttr = wml.ST_LineSpacingRuleUnset
		p.x.LineAttr = nil
	} else {
		p.x.LineRuleAttr = rule
		p.x.LineAttr = &wml.ST_SignedTwipsMeasure{}
		p.x.LineAttr.Int64 = unioffice.Int64(int64(d / measurement.Twips))
	}
}

// SetBeforeAuto controls if spacing before a paragraph is automatically determined.
func (p ParagraphSpacing) SetBeforeAuto(b bool) {
	if b {
		p.x.BeforeAutospacingAttr = &sharedTypes.ST_OnOff{}
		p.x.BeforeAutospacingAttr.Bool = unioffice.Bool(true)
	} else {
		p.x.BeforeAutospacingAttr = nil
	}
}

// SetAfterAuto controls if spacing after a paragraph is automatically determined.
func (p ParagraphSpacing) SetAfterAuto(b bool) {
	if b {
		p.x.AfterAutospacingAttr = &sharedTypes.ST_OnOff{}
		p.x.AfterAutospacingAttr.Bool = unioffice.Bool(true)
	} else {
		p.x.AfterAutospacingAttr = nil
	}
}
