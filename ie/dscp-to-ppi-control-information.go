// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewDSCPToPPIControlInformation creates a new DSCPToPPIControlInformation IE.
func NewDSCPToPPIControlInformation(ies ...*IE) *IE {
	return newGroupedIE(DSCPToPPIControlInformation, 0, ies...)
}

// DSCPToPPIControlInformation returns the IEs above DSCPToPPIControlInformation if the type of IE matches.
func (i *IE) DSCPToPPIControlInformation() (*DSCPToPPIControlInformationFields, error) {
	if i.Type != DSCPToPPIControlInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseDSCPToPPIControlInformation(i.Payload)
}

// DSCPToPPIControlInformationFields represents a fields contained in DSCPToPPIControlInformation IE.
type DSCPToPPIControlInformationFields struct {
	DSCPToPPIMappingInformation []*DSCPToPPIMappingInformationFields
	QFI                         []uint8
}

// ParseDSCPToPPIControlInformation parses b into DSCPToPPIControlInformation.
func ParseDSCPToPPIControlInformation(b []byte) (*DSCPToPPIControlInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &DSCPToPPIControlInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case DSCPToPPIMappingInformation:
			v, err := ie.DSCPToPPIMappingInformation()
			if err != nil {
				return p, err
			}
			p.DSCPToPPIMappingInformation = append(p.DSCPToPPIMappingInformation, v)
		case QFI:
			v, err := ie.QFI()
			if err != nil {
				return p, err
			}
			p.QFI = append(p.QFI, v)
		}
	}
	return p, nil
}
