// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package presentation

import (
	"github.com/gf0842wf/unioffice/drawing"
	"github.com/gf0842wf/unioffice/schema/soo/dml"
	"github.com/gf0842wf/unioffice/schema/soo/pml"
)

// Image is an image within a slide.
type Image struct {
	x *pml.CT_Picture
}

// Properties returns the properties of the TextBox.
func (i Image) Properties() drawing.ShapeProperties {
	if i.x.SpPr == nil {
		i.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(i.x.SpPr)
}
