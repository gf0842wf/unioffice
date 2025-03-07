// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/gf0842wf/unioffice/schema/soo/sml"
)

func TestCT_VolTypeConstructor(t *testing.T) {
	v := sml.NewCT_VolType()
	if v == nil {
		t.Errorf("sml.NewCT_VolType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_VolType should validate: %s", err)
	}
}

func TestCT_VolTypeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_VolType()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_VolType()
	xml.Unmarshal(buf, v2)
}
