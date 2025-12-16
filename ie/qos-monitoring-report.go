// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewQoSMonitoringReport creates a new QoSMonitoringReport IE.
func NewQoSMonitoringReport(ies ...*IE) *IE {
	return newGroupedIE(QoSMonitoringReport, 0, ies...)
}

// QoSMonitoringReport returns the IEs above QoSMonitoringReport if the type of IE matches.
func (i *IE) QoSMonitoringReport() (*QoSMonitoringReportFields, error) {
	switch i.Type {
	case QoSMonitoringReport:

		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &QoSMonitoringReportFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseQoSMonitoringReportFields(i.Payload)
	case SessionReport:
		ies, err := i.SessionReport()
		if err != nil {
			return nil, err
		}
		if ies.QoSMonitoringReport != nil {
			return ies.QoSMonitoringReport, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// QoSMonitoringReportFields is a set of fields in QoSMonitoringReport IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type QoSMonitoringReportFields struct {
	QFI                      uint8
	QoSMonitoringMeasurement *QoSMonitoringMeasurementFields
	Timestamp                time.Time
	StartTime                time.Time
}

// ParseQoSMonitoringReportFields returns the IEs above QoSMonitoringReport
func ParseQoSMonitoringReportFields(b []byte) (*QoSMonitoringReportFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	c := &QoSMonitoringReportFields{}
	if err := c.ParseIEs(ies...); err != nil {
		return c, err
	}

	return c, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (s *QoSMonitoringReportFields) ParseIEs(ies ...*IE) error {
	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case QFI:
			v, err := i.QFI()
			if err != nil {
				return nil
			}
			s.QFI = v
		case QoSMonitoringMeasurement:
			v, err := i.QoSMonitoringMeasurement()
			if err != nil {
				return nil
			}
			s.QoSMonitoringMeasurement = v
		case EventTimeStamp:
			v, err := i.EventTimeStamp()
			if err != nil {
				return nil
			}
			s.Timestamp = v
		case StartTime:
			v, err := i.StartTime()
			if err != nil {
				return nil
			}
			s.StartTime = v
		}
	}
	return nil
}
