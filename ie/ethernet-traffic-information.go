// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewEthernetTrafficInformation creates a new  IE.
func NewEthernetTrafficInformation(ies ...*IE) *IE {
	return newGroupedIE(EthernetTrafficInformation, 0, ies...)
}

// EthernetTrafficInformation returns the IEs above EthernetTrafficInformation if the type of IE matches.
func (i *IE) EthernetTrafficInformation() (*EthernetTrafficInformationFields, error) {
	switch i.Type {
	case EthernetTrafficInformation:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &EthernetTrafficInformationFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseEthernetTrafficInformationFields(i.Payload)
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return nil, err
		}
		if ies.EthernetTrafficInformation != nil {
			return ies.EthernetTrafficInformation, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// EthernetTrafficInformationFields is a set of fields in Ethernet Traffic Information IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type EthernetTrafficInformationFields struct {
	MACAddressesDetected *MACAddressesDetectedFields
	MACAddressesRemoved  *MACAddressesRemovedFields
}

// ParseEthernetTrafficInformationFields returns the IEs above Update FAR
func ParseEthernetTrafficInformationFields(b []byte) (*EthernetTrafficInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &EthernetTrafficInformationFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *EthernetTrafficInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case MACAddressesDetected:
			v, err := ie.MACAddressesDetected()
			if err != nil {
				return err
			}
			far.MACAddressesDetected = v

		case MACAddressesRemoved:
			v, err := ie.MACAddressesRemoved()
			if err != nil {
				return err
			}
			far.MACAddressesRemoved = v
		}
	}
	return nil
}
