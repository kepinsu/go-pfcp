// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"time"
)

// NewClockDriftControlInformation creates a new ClockDriftControlInformation IE.
func NewClockDriftControlInformation(ies ...*IE) *IE {
	return newGroupedIE(ClockDriftControlInformation, 0, ies...)
}

// ClockDriftControlInformation returns the IEs above ClockDriftControlInformation if the type of IE matches.
func (i *IE) ClockDriftControlInformation() (*ClockDriftControlInformationFields, error) {
	if i.Type != ClockDriftControlInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseClockDriftControlInformationFields(i.Payload)
}

// ClockDriftControlInformationFields is a set of fields in ClockDriftControlInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ClockDriftControlInformationFields struct {
	RequestedClockDriftInformation uint8
	TimeDomainNumber               uint8
	ConfiguredTimeDomain           uint8
	TimeOffsetThreshold            time.Duration
	CumulativeRateRatioThreshold   uint32
}

// ParseClockDriftControlInformationFields returns the IEs above ClockDriftControlInformation
func ParseClockDriftControlInformationFields(b []byte) (*ClockDriftControlInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &ClockDriftControlInformationFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *ClockDriftControlInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case RequestedClockDriftInformation:
			v, err := ie.RequestedClockDriftInformation()
			if err != nil {
				return err
			}
			far.RequestedClockDriftInformation = v
		case TSNTimeDomainNumber:
			v, err := ie.TSNTimeDomainNumber()
			if err != nil {
				return err
			}
			far.TimeDomainNumber = v
		case ConfiguredTimeDomain:
			v, err := ie.ConfiguredTimeDomain()
			if err != nil {
				return err
			}
			far.ConfiguredTimeDomain = v
		case TimeOffsetThreshold:
			v, err := ie.TimeOffsetThreshold()
			if err != nil {
				return err
			}
			far.TimeOffsetThreshold = v
		case CumulativeRateRatioThreshold:
			v, err := ie.CumulativeRateRatioThreshold()
			if err != nil {
				return err
			}
			far.CumulativeRateRatioThreshold = v
		}
	}
	return nil
}
