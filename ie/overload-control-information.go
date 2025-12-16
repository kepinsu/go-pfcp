// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewOverloadControlInformation creates a new OverloadControlInformation IE.
func NewOverloadControlInformation(ies ...*IE) *IE {
	return newGroupedIE(OverloadControlInformation, 0, ies...)
}

// OverloadControlInformation returns the IEs above OverloadControlInformation if the type of IE matches.
func (i *IE) OverloadControlInformation() (*OverLoadControlInformationFields, error) {
	if i.Type != OverloadControlInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		o := &OverLoadControlInformationFields{}
		if err := o.ParseIEs(i.ChildIEs...); err != nil {
			return o, err
		}
		return o, nil
	}
	// If the ChildIEs not already parsed
	return ParseOverLoadControlInformationFields(i.Payload)
}

// OverLoadControlInformationFields represents a fields contained in OverLoadControlInformation IE.
type OverLoadControlInformationFields struct {
	Sequence                         uint32
	Metric                           uint8
	PeriodOfValidity                 time.Duration
	OverloadControlInformationsFlags uint8
}

// ParseOverLoadControlInformationFields creates a new OverLoadControlInformation IE.
func ParseOverLoadControlInformationFields(b []byte) (*OverLoadControlInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	load := &OverLoadControlInformationFields{}
	if err := load.ParseIEs(ies...); err != nil {
		return load, err
	}
	return load, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (l *OverLoadControlInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case SequenceNumber:
			v, err := ie.SequenceNumber()
			if err != nil {
				return err
			}
			l.Sequence = v
		case Metric:
			v, err := ie.Metric()
			if err != nil {
				return err
			}
			l.Metric = v
		case Timer:
			v, err := ie.Timer()
			if err != nil {
				return err
			}
			l.PeriodOfValidity = v
		case OCIFlags:
			v, err := ie.OCIFlags()
			if err != nil {
				return err
			}
			l.OverloadControlInformationsFlags = v
		}
	}
	return nil
}
