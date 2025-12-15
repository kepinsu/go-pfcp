// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPMFParameters creates a new PMFParameters IE.
func NewPMFParameters(info *IE) *IE {
	return newGroupedIE(PMFParameters, 0, info)
}

// PMFParameters returns the IEs above PMFParameters if the type of IE matches.
func (i *IE) PMFParameters() (*PMFParametersFields, error) {
	switch i.Type {
	case PMFParameters:
		return ParsePMFParametersFields(i.Payload)
	case ATSSSControlParameters:
		ies, err := i.ATSSSControlParameters()
		if err != nil {
			return nil, err
		}
		if ies.PMFParameters != nil {
			return ies.PMFParameters, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// PMFParametersFields is a set of felds in PMFParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type PMFParametersFields struct {
	PMFAddressInformation *PMFAddressInformationFields
	QFI                   uint8
}

// ParsePMFParametersFields returns the IEs above ProvideATSSSControlInformation IE
func ParsePMFParametersFields(b []byte) (*PMFParametersFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &PMFParametersFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case PMFAddressInformation:
			v, err := ie.PMFAddressInformation()
			if err != nil {
				return p, err
			}
			p.PMFAddressInformation = v
		case QFI:
			v, err := ie.QFI()
			if err != nil {
				return p, err
			}
			p.QFI = v
		}
	}
	return p, nil
}
