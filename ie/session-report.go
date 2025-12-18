// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewSessionReport creates a new SessionReport IE.
func NewSessionReport(ies ...*IE) *IE {
	return newGroupedIE(SessionReport, 0, ies...)
}

// SessionReport returns the IEs above SessionReport if the type of IE matches.
func (i *IE) SessionReport() (*SessionReportFields, error) {
	switch i.Type {
	case SessionReport:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &SessionReportFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseSessionReportFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// SessionReportFields is a set of fields in SessionReport IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type SessionReportFields struct {
	SRRID                    uint8
	AccessAvailabilityReport *AccessAvailabilityReportFields
	QoSMonitoringReport      *QoSMonitoringReportFields
}

// ParseSessionReportFields returns the IEs above SessionReport
func ParseSessionReportFields(b []byte) (*SessionReportFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	c := &SessionReportFields{}
	if err := c.ParseIEs(ies...); err != nil {
		return c, err
	}

	return c, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (s *SessionReportFields) ParseIEs(ies ...*IE) error {
	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case SRRID:
			v, err := i.SRRID()
			if err != nil {
				return nil
			}
			s.SRRID = v
		case AccessAvailabilityReport:
			v, err := i.AccessAvailabilityReport()
			if err != nil {
				return nil
			}
			s.AccessAvailabilityReport = v
		case QoSMonitoringReport:
			v, err := i.QoSMonitoringReport()
			if err != nil {
				return nil
			}
			s.QoSMonitoringReport = v
		}
	}
	return nil
}
