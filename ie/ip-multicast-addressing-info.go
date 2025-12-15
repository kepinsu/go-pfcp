// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewIPMulticastAddressingInfo creates a new IPMulticastAddressingInfo IE.
func NewIPMulticastAddressingInfo(ies ...*IE) *IE {
	return newGroupedIE(IPMulticastAddressingInfo, 0, ies...)
}

// IPMulticastAddressingInfo returns the IEs above IPMulticastAddressingInfo if the type of IE matches.
func (i *IE) IPMulticastAddressingInfo() (*IPMulticastAddressingInfoField, error) {
	switch i.Type {
	case IPMulticastAddressingInfo:
		return ParseIPMulticastAddressingInfo(i.Payload)
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return nil, err
		}
		if len(ies.IPMulticastAddressingInfo) > 0 {
			return ies.IPMulticastAddressingInfo[0], nil
		}
		return nil, ErrIENotFound
	case UpdatePDR:
		ies, err := i.UpdatePDR()
		if err != nil {
			return nil, err
		}
		if len(ies.IPMulticastAddressingInfo) > 0 {
			return ies.IPMulticastAddressingInfo[0], nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

type IPMulticastAddressingInfoField struct {
	IPMulticastAddress *IPMulticastAddressFields
	SourceIPAddress    *SourceIPAddressFields
}

func ParseIPMulticastAddressingInfo(b []byte) (*IPMulticastAddressingInfoField, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	fields := &IPMulticastAddressingInfoField{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case IPMulticastAddress:
			field, err := ParseIPMulticastAddressFields(ie.Payload)
			if err != nil {
				return nil, err
			}
			fields.IPMulticastAddress = field
		case SourceIPAddress:
			field, err := ParseSourceIPAddressFields(ie.Payload)
			if err != nil {
				return nil, err
			}
			fields.SourceIPAddress = field
		}
	}

	return fields, nil
}
