// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPartialFailureInformation creates a new PartialFailureInformation IE.
func NewPartialFailureInformation(ies ...*IE) *IE {
	return newGroupedIE(PartialFailureInformation, 0, ies...)
}

// PartialFailureInformation returns the IEs above PartialFailureInformation if the type of IE matches.
func (i *IE) PartialFailureInformation() ([]*IE, error) {
	switch i.Type {
	case PartialFailureInformation:
		return ParseMultiIEs(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// PartialFailureInformationFields is a set of felds in PartialFailureInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type PartialFailureInformationFields struct {
	FailedRuleID           uint32
	FailureCause           uint8
	OffendingIEInformation *OffendingIEInformationFields
}

// PartialFailureInformationFields returns the IEs above PartialFailureInformation IE
func ParsePartialFailureInformationFields(b []byte) (*PartialFailureInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &PartialFailureInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case FailedRuleID:
			v, err := ie.FailedRuleID()
			if err != nil {
				return p, err
			}
			p.FailedRuleID = v
		case Cause:
			v, err := ie.Cause()
			if err != nil {
				return p, err
			}
			p.FailureCause = v
		case OffendingIEInformation:
			v, err := ie.OffendingIEInformation()
			if err != nil {
				return p, err
			}
			p.OffendingIEInformation = v
		}
	}
	return p, nil
}
