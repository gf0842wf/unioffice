// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/gf0842wf/unioffice/schema/soo/pml"
)

func TestPresentationConstructor(t *testing.T) {
	v := pml.NewPresentation()
	if v == nil {
		t.Errorf("pml.NewPresentation must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.Presentation should validate: %s", err)
	}
}

func TestPresentationMarshalUnmarshal(t *testing.T) {
	v := pml.NewPresentation()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewPresentation()
	xml.Unmarshal(buf, v2)
}
