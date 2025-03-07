// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/gf0842wf/unioffice/schema/soo/dml/chartDrawing"
)

func TestCT_DrawingConstructor(t *testing.T) {
	v := chartDrawing.NewCT_Drawing()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_Drawing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_Drawing should validate: %s", err)
	}
}

func TestCT_DrawingMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_Drawing()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_Drawing()
	xml.Unmarshal(buf, v2)
}
