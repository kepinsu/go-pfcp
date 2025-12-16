// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewErrorIndicationReport creates a new ErrorIndicationReport IE.
func NewErrorIndicationReport(fteid *IE) *IE {
	return newGroupedIE(ErrorIndicationReport, 0, fteid)
}

// ErrorIndicationReport returns the IEs above ErrorIndicationReport if the type of IE matches.
func (i *IE) ErrorIndicationReport() (*ErrorIndicationReportFields, error) {
	if i.Type != ErrorIndicationReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &ErrorIndicationReportFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseErrorIndicationReportFields(i.Payload)
}

// ErrorIndicationReportFields is a set of fields in ErrorIndicationReport IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ErrorIndicationReportFields struct {
	RemoteFTEID *FTEIDFields
}

// ParseErrorIndicationReportFields returns the IEs above ErrorIndicationReport
func ParseErrorIndicationReportFields(b []byte) (*ErrorIndicationReportFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	c := &ErrorIndicationReportFields{}
	if err := c.ParseIEs(ies...); err != nil {
		return c, err
	}

	return c, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (e *ErrorIndicationReportFields) ParseIEs(ies ...*IE) error {
	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case FTEID:
			v, err := i.FTEID()
			if err != nil {
				return nil
			}
			e.RemoteFTEID = v
		}
	}
	return nil
}
