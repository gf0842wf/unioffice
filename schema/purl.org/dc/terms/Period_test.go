// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/gf0842wf/unioffice/schema/purl.org/dc/terms"
)

func TestPeriodConstructor(t *testing.T) {
	v := terms.NewPeriod()
	if v == nil {
		t.Errorf("terms.NewPeriod must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.Period should validate: %s", err)
	}
}

func TestPeriodMarshalUnmarshal(t *testing.T) {
	v := terms.NewPeriod()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewPeriod()
	xml.Unmarshal(buf, v2)
}
