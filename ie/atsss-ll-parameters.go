// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewATSSSLLParameters creates a new ATSSSLLParameters IE.
func NewATSSSLLParameters(info *IE) *IE {
	return newGroupedIE(ATSSSLLParameters, 0, info)
}

// ATSSSLLParameters returns the IEs above ATSSSLLParameters if the type of IE matches.
func (i *IE) ATSSSLLParameters() (*ATSSSLLParametersFields, error) {
	switch i.Type {
	case ATSSSLLParameters:
		return ParseATSSSLLParametersFields(i.Payload)
	case ATSSSControlParameters:
		ies, err := i.ATSSSControlParameters()
		if err != nil {
			return nil, err
		}
		if ies.ATSSSLLParameters != nil {
			return ies.ATSSSLLParameters, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// ATSSSLLParametersFields is a set of felds in ATSSSLLParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ATSSSLLParametersFields struct {
	ATSSSLLInformation uint8
}

// ParseATSSSLLParametersFields returns the IEs above ProvideATSSSControlInformation IE
func ParseATSSSLLParametersFields(b []byte) (*ATSSSLLParametersFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &ATSSSLLParametersFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == ATSSSLLInformation {
			v, err := ie.ATSSSLLInformation()
			if err != nil {
				return p, err
			}
			p.ATSSSLLInformation = v
		}
	}
	return p, nil
}
