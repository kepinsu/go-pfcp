// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewNonTGPPAccessForwardingActionInformation creates a new NonTGPPAccessForwardingActionInformation IE.
func NewNonTGPPAccessForwardingActionInformation(ies ...*IE) *IE {
	return newGroupedIE(NonTGPPAccessForwardingActionInformation, 0, ies...)
}

// NonTGPPAccessForwardingActionInformation returns the IEs above NonTGPPAccessForwardingActionInformation if the type of IE matches.
func (i *IE) NonTGPPAccessForwardingActionInformation() (*NonTGPPAccessForwardingActionInformationFields, error) {
	switch i.Type {
	case NonTGPPAccessForwardingActionInformation:
		return ParseNonTGPPAccessForwardingActionInformationFields(i.Payload)
	case CreateMAR:
		ies, err := i.CreateMAR()
		if err != nil {
			return nil, err
		}
		if ies.NonTGPPAccessForwardingActionInformation != nil {
			return ies.NonTGPPAccessForwardingActionInformation, nil
		}
		return nil, ErrIENotFound
	case UpdateMAR:
		ies, err := i.UpdateMAR()
		if err != nil {
			return nil, err
		}
		if ies.NonTGPPAccessForwardingActionInformation != nil {
			return ies.NonTGPPAccessForwardingActionInformation, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// TGPPAccessForwardingActionInformationFields is a set of fields in TGPPAccessForwardingActionInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type NonTGPPAccessForwardingActionInformationFields struct {
	FARID    uint32
	Weight   uint8
	Priority uint8
	URRID    uint32
	RATType  uint8
}

// TGPPNonAccessForwardingActionInformation returns the IEs above NonTGPPAccessForwardingActionInformation.
func ParseNonTGPPAccessForwardingActionInformationFields(b []byte) (*NonTGPPAccessForwardingActionInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	t := &NonTGPPAccessForwardingActionInformationFields{}

	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case FARID:
			v, err := ie.FARID()
			if err != nil {
				return t, err
			}
			t.FARID = v
		case Weight:
			v, err := ie.Weight()
			if err != nil {
				return t, err
			}
			t.Weight = v
		case Priority:
			v, err := ie.Priority()
			if err != nil {
				return t, err
			}
			t.Priority = v
		case URRID:
			v, err := ie.URRID()
			if err != nil {
				return t, err
			}
			t.URRID = v
		case RATType:
			v, err := ie.RATType()
			if err != nil {
				return t, err
			}
			t.RATType = v
		}
	}

	return t, nil
}
