// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewMPTCPParameters creates a new MPTCPParameters IE.
func NewMPTCPParameters(ies ...*IE) *IE {
	return newGroupedIE(MPTCPParameters, 0, ies...)
}

// MPTCPParameters returns the IEs above MPTCPParameters if the type of IE matches.
func (i *IE) MPTCPParameters() (*MPTCPParametersFields, error) {
	switch i.Type {
	case MPTCPParameters:
		return ParseMPTCPParametersFields(i.Payload)
	case ATSSSControlParameters:
		ies, err := i.ATSSSControlParameters()
		if err != nil {
			return nil, err
		}
		if ies.MPTCPParameters != nil {
			return ies.MPTCPParameters, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MPTCPParametersFields is a set of felds in MPTCPParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type MPTCPParametersFields struct {
	MPTCPAddressInformation *MPTCPAddressInformationFields
	UELinkSpecificIPAddress *UELinkSpecificIPAddressFields
}

// ParseMPTCPParametersFields returns the IEs above ProvideATSSSControlInformation IE
func ParseMPTCPParametersFields(b []byte) (*MPTCPParametersFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &MPTCPParametersFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case MPTCPAddressInformation:
			v, err := ie.MPTCPAddressInformation()
			if err != nil {
				return p, err
			}
			p.MPTCPAddressInformation = v
		case UELinkSpecificIPAddress:
			v, err := ie.UELinkSpecificIPAddress()
			if err != nil {
				return p, err
			}
			p.UELinkSpecificIPAddress = v
		}
	}
	return p, nil
}
