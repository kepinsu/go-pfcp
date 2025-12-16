// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewProvideATSSSControlInformation creates a new ProvideATSSSControlInformation IE.
func NewProvideATSSSControlInformation(ies ...*IE) *IE {
	return newGroupedIE(ProvideATSSSControlInformation, 0, ies...)
}

// ProvideATSSSControlInformation returns the IEs above ProvideATSSSControlInformation if the type of IE matches.
func (i *IE) ProvideATSSSControlInformation() (*ProvideATSSSControlInformationFields, error) {
	if i.Type != ProvideATSSSControlInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseProvideATSSSControlInformationFields(i.Payload)
}

// ProvideATSSSControlInformationFields is a set of felds in ProvideATSSSControlInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ProvideATSSSControlInformationFields struct {
	MPTCPControlInformation   uint8
	ATSSSLLControlInformation uint8
	PMFControlInformation     uint8
}

// ParseProvideATSSSControlInformationFields returns the IEs above ProvideATSSSControlInformation IE
func ParseProvideATSSSControlInformationFields(b []byte) (*ProvideATSSSControlInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &ProvideATSSSControlInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case MPTCPControlInformation:
			v, err := ie.MPTCPControlInformation()
			if err != nil {
				return p, err
			}
			p.MPTCPControlInformation = v
		case ATSSSLLControlInformation:
			v, err := ie.ATSSSLLControlInformation()
			if err != nil {
				return p, err
			}
			p.ATSSSLLControlInformation = v
		case PMFControlInformation:
			v, err := ie.PMFControlInformation()
			if err != nil {
				return p, err
			}
			p.PMFControlInformation = v
		}
	}
	return p, nil
}
