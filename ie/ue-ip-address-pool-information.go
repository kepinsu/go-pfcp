// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUEIPAddressPoolInformation creates a new UEIPAddressPoolInformation IE.
func NewUEIPAddressPoolInformation(ies ...*IE) *IE {
	return newGroupedIE(UEIPAddressPoolInformation, 0, ies...)
}

// UEIPAddressPoolInformation returns the IEs above UEIPAddressPoolInformation if the type of IE matches.
func (i *IE) UEIPAddressPoolInformation() ([]*IE, error) {
	if i.Type != UEIPAddressPoolInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}

// UEIPAddressPoolInformationFields represents a fields contained in UEIPAddressPoolInformation IE.
type UEIPAddressPoolInformationFields struct {
	UEIPAddressPoolIdentity []*IE
	NetworkInstance         string
	SNSSAI                  []byte
	IPVersion               uint8
}

// ParseUEIPAddressPoolInformationFields parses b into UEIPAddressPoolInformationFields.
func ParseUEIPAddressPoolInformationFields(b []byte) (*UEIPAddressPoolInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &UEIPAddressPoolInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case UEIPAddressPoolIdentity:
			p.UEIPAddressPoolIdentity = append(p.UEIPAddressPoolIdentity, ie)
		case NetworkInstance:
			v, err := ie.NetworkInstance()
			if err != nil {
				return p, err
			}
			p.NetworkInstance = v
		case SNSSAI:
			v, err := ie.SNSSAI()
			if err != nil {
				return p, err
			}
			p.SNSSAI = v
		case IPVersion:
			v, err := ie.IPVersion()
			if err != nil {
				return p, err
			}
			p.IPVersion = v
		}
	}
	return p, nil
}
