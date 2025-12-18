// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewProvideRDSConfigurationInformation creates a new ProvideRDSConfigurationInformation IE.
func NewProvideRDSConfigurationInformation(ies ...*IE) *IE {
	return newGroupedIE(ProvideRDSConfigurationInformation, 0, ies...)
}

// ProvideRDSConfigurationInformation returns the IEs above ProvideRDSConfigurationInformation if the type of IE matches.
func (i *IE) ProvideRDSConfigurationInformation() (*ProvideRDSConfigurationInformationFields, error) {
	if i.Type != ProvideRDSConfigurationInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseProvideRDSConfigurationInformationFields(i.Payload)
}

// ProvideRDSConfigurationInformationFields is a set of felds in ProvideRDSConfigurationInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ProvideRDSConfigurationInformationFields struct {
	RDSConfigurationInformation uint8
}

// ParseProvideRDSConfigurationInformationFields returns the IEs above ProvideATSSSControlInformation IE
func ParseProvideRDSConfigurationInformationFields(b []byte) (*ProvideRDSConfigurationInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &ProvideRDSConfigurationInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == RDSConfigurationInformation {
			v, err := ie.RDSConfigurationInformation()
			if err != nil {
				return p, err
			}
			p.RDSConfigurationInformation = v
		}
	}
	return p, nil
}
