// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package reference

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gf0842wf/unioffice/spreadsheet/update"
)

// CellReference is a parsed reference to a cell.  Input is of the form 'A1',
// '$C$2', etc.
type CellReference struct {
	RowIdx         uint32
	ColumnIdx      uint32
	Column         string
	AbsoluteColumn bool
	AbsoluteRow    bool
	SheetName      string
}

// String returns a string representation of CellReference.
func (c CellReference) String() string {
	buf := make([]byte, 0, 4)
	if c.AbsoluteColumn {
		buf = append(buf, '$')
	}
	buf = append(buf, c.Column...)
	if c.AbsoluteRow {
		buf = append(buf, '$')
	}
	buf = strconv.AppendInt(buf, int64(c.RowIdx), 10)
	return string(buf)
}

// ParseCellReference parses a cell reference of the form 'A10' and splits it
// into column/row segments.
func ParseCellReference(s string) (CellReference, error) {
	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return CellReference{}, errors.New("cell reference must have at least two characters")
	}

	r := CellReference{}
	sl := strings.Split(s, "!")
	if len(sl) == 2 {
		r.SheetName = sl[0]
		s = sl[1]
	}
	// check for absolute column
	if s[0] == '$' {
		r.AbsoluteColumn = true
		s = s[1:]
	}

	split := -1
lfor:
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9' || s[i] == '$':
			split = i
			break lfor
		}
	}
	switch split {
	case 0:
		return CellReference{}, fmt.Errorf("no letter prefix in %s", s)
	case -1:
		return CellReference{}, fmt.Errorf("no digits in %s", s)
	}

	r.Column = s[0:split]

	if s[split] == '$' {
		r.AbsoluteRow = true
		split++
	}

	r.ColumnIdx = ColumnToIndex(r.Column)
	r64, err := strconv.ParseUint(s[split:], 10, 32)
	if err != nil {
		return CellReference{}, fmt.Errorf("error parsing row: %s", err)
	}
	r.RowIdx = uint32(r64)

	return r, nil
}

// Update updates reference to point one of the neighboring cells with respect to the update type after removing a row/column.
func (ref *CellReference) Update(updateType update.UpdateAction) *CellReference {
	switch updateType {
	case update.UpdateActionRemoveColumn:
		newRef := ref
		newRef.ColumnIdx = ref.ColumnIdx - 1
		newRef.Column = IndexToColumn(newRef.ColumnIdx)
		return newRef
	default:
		return ref
	}
}
