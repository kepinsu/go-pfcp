// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewUEIPAddressUsageInformation creates a new UEIPAddressUsageInformation IE.
func NewUEIPAddressUsageInformation(ies ...*IE) *IE {
	return newGroupedIE(UEIPAddressUsageInformation, 0, ies...)
}

// UEIPAddressUsageInformation returns the IEs above UEIPAddressUsageInformation if the type of IE matches.
func (i *IE) UEIPAddressUsageInformation() (*UEIPAddressUsageInformationFields, error) {
	switch i.Type {
	case UEIPAddressUsageInformation:
		return ParseUEIPAddressUsageInformationFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// UEIPAddressUsageInformationFields is a set of fields in UEIPAddressUsageInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UEIPAddressUsageInformationFields struct {
	SequenceNumber        uint32
	Metric                uint8
	ValidityTimer         time.Duration
	NumberOfUEIPAddresses *NumberOfUEIPAddressesFields
	NetworkInstance       string
	UEIPAddressPoolId     []byte
	SNSSAI                []byte
}

// ParseUEIPAddressUsageInformationFields returns the IEs above UEIPAddressUsageInformation
func ParseUEIPAddressUsageInformationFields(b []byte) (*UEIPAddressUsageInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &UEIPAddressUsageInformationFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *UEIPAddressUsageInformationFields) ParseIEs(ies ...*IE) error {
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
			far.SequenceNumber = v
		case Metric:
			v, err := ie.Metric()
			if err != nil {
				return err
			}
			far.Metric = v
		case ValidityTimer:
			v, err := ie.ValidityTimer()
			if err != nil {
				return err
			}
			far.ValidityTimer = v
		case NumberOfUEIPAddresses:
			v, err := ie.NumberOfUEIPAddresses()
			if err != nil {
				return err
			}
			far.NumberOfUEIPAddresses = v
		case NetworkInstance:
			v, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			far.NetworkInstance = v
		case UEIPAddressPoolIdentity:
			v, err := ie.UEIPAddressPoolIdentity()
			if err != nil {
				return err
			}
			far.UEIPAddressPoolId = v
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
