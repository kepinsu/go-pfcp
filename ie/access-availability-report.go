// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewAccessAvailabilityReport creates a new AccessAvailabilityReport IE.
func NewAccessAvailabilityReport(info *IE) *IE {
	return newGroupedIE(AccessAvailabilityReport, 0, info)
}

// AccessAvailabilityReport returns the IEs above AccessAvailabilityReport if the type of IE matches.
func (i *IE) AccessAvailabilityReport() (*AccessAvailabilityReportFields, error) {
	if i.Type != AccessAvailabilityReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &AccessAvailabilityReportFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseAccessAvailabilityReportFields(i.Payload)
}

// AccessAvailabilityReportFields is a set of fields in AccessAvailabilityReport IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type AccessAvailabilityReportFields struct {
	AccessAvailabilityInformation uint8
}

// ParseAccessAvailabilityReportFields returns the IEs above AccessAvailabilityReport
func ParseAccessAvailabilityReportFields(b []byte) (*AccessAvailabilityReportFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	c := &AccessAvailabilityReportFields{}
	if err := c.ParseIEs(ies...); err != nil {
		return c, err
	}

	return c, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (s *AccessAvailabilityReportFields) ParseIEs(ies ...*IE) error {
	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case AccessAvailabilityInformation:
			v, err := i.AccessAvailabilityInformation()
			if err != nil {
				return nil
			}
			s.AccessAvailabilityInformation = v
		}
	}
	return nil
}
