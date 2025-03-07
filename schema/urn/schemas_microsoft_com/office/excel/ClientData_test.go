// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package excel_test

import (
	"encoding/xml"
	"testing"

	"github.com/gf0842wf/unioffice/schema/urn/schemas_microsoft_com/office/excel"
)

func TestClientDataConstructor(t *testing.T) {
	v := excel.NewClientData()
	if v == nil {
		t.Errorf("excel.NewClientData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed excel.ClientData should validate: %s", err)
	}
}

func TestClientDataMarshalUnmarshal(t *testing.T) {
	v := excel.NewClientData()
	buf, _ := xml.Marshal(v)
	v2 := excel.NewClientData()
	xml.Unmarshal(buf, v2)
}
