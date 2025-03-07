// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/gf0842wf/unioffice/schema/soo/dml/diagram"
)

func TestCT_PresentationOfConstructor(t *testing.T) {
	v := diagram.NewCT_PresentationOf()
	if v == nil {
		t.Errorf("diagram.NewCT_PresentationOf must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_PresentationOf should validate: %s", err)
	}
}

func TestCT_PresentationOfMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_PresentationOf()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_PresentationOf()
	xml.Unmarshal(buf, v2)
}
