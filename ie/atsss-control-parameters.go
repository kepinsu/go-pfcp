// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewATSSSControlParameters creates a new ATSSSControlParameters IE.
func NewATSSSControlParameters(ies ...*IE) *IE {
	return newGroupedIE(ATSSSControlParameters, 0, ies...)
}

// ATSSSControlParameters returns the IEs above ATSSSControlParameters if the type of IE matches.
func (i *IE) ATSSSControlParameters() (*ATSSSControlParametersFields, error) {
	if i.Type != ATSSSControlParameters {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseATSSSControlParametersFields(i.Payload)
}

// ATSSSControlParametersFields is a set of felds in ATSSSControlParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ATSSSControlParametersFields struct {
	MPTCPParameters   *MPTCPParametersFields
	ATSSSLLParameters *ATSSSLLParametersFields
	PMFParameters     *PMFParametersFields
}

// ParseATSSSControlParametersFields returns the IEs above ProvideATSSSControlInformation IE
func ParseATSSSControlParametersFields(b []byte) (*ATSSSControlParametersFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &ATSSSControlParametersFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case MPTCPParameters:
			v, err := ie.MPTCPParameters()
			if err != nil {
				return p, err
			}
			p.MPTCPParameters = v
		case ATSSSLLParameters:
			v, err := ie.ATSSSLLParameters()
			if err != nil {
				return p, err
			}
			p.ATSSSLLParameters = v
		case PMFParameters:
			v, err := ie.PMFParameters()
			if err != nil {
				return p, err
			}
			p.PMFParameters = v
		}
	}
	return p, nil
}
