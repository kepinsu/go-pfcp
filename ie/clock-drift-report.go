// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewClockDriftReport creates a new ClockDriftReport IE.
func NewClockDriftReport(ies ...*IE) *IE {
	return newGroupedIE(ClockDriftReport, 0, ies...)
}

// ClockDriftReport returns the IEs above ClockDriftReport if the type of IE matches.
func (i *IE) ClockDriftReport() (*ClockDriftReportFields, error) {
	if i.Type != ClockDriftReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseClockDriftReportFields(i.Payload)
}

// ClockDriftReportFields is a set of fields in ClockDriftReport IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ClockDriftReportFields struct {
	TSNTimeDomainNumber          uint8
	TimeOffsetMeasurement        time.Duration
	CumulativeRateRatioThreshold uint32
	TimeStamp                    time.Time
	NetworkInstance              string
	APNDNN                       string
	SNSSAI                       []byte
}

// ParseClockDriftReportFields returns the IEs above ClockDriftReport
func ParseClockDriftReportFields(b []byte) (*ClockDriftReportFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &ClockDriftReportFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *ClockDriftReportFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case TSNTimeDomainNumber:
			v, err := ie.TSNTimeDomainNumber()
			if err != nil {
				return err
			}
			far.TSNTimeDomainNumber = v
		case TimeOffsetMeasurement:
			v, err := ie.TimeOffsetMeasurement()
			if err != nil {
				return err
			}
			far.TimeOffsetMeasurement = v
		case CumulativeRateRatioThreshold:
			v, err := ie.CumulativeRateRatioThreshold()
			if err != nil {
				return err
			}
			far.CumulativeRateRatioThreshold = v
		case EventTimeStamp:
			v, err := ie.EventTimeStamp()
			if err != nil {
				return err
			}
			far.TimeStamp = v
		case NetworkInstance:
			v, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			far.NetworkInstance = v
		case APNDNN:
			v, err := ie.APNDNN()
			if err != nil {
				return err
			}
			far.APNDNN = v
		case SNSSAI:
			v, err := ie.SNSSAI()
			if err != nil {
				return err
			}
			far.SNSSAI = v
		}
	}
	return nil
}
