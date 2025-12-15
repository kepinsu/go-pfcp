// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTGPPAccessForwardingActionInformation creates a new TGPPAccessForwardingActionInformation IE.
func NewTGPPAccessForwardingActionInformation(ies ...*IE) *IE {
	return newGroupedIE(TGPPAccessForwardingActionInformation, 0, ies...)
}

// TGPPAccessForwardingActionInformation returns the IEs above TGPPAccessForwardingActionInformation if the type of IE matches.
func (i *IE) TGPPAccessForwardingActionInformation() (*TGPPAccessForwardingActionInformationFields, error) {
	switch i.Type {
	case TGPPAccessForwardingActionInformation:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			t := &TGPPAccessForwardingActionInformationFields{}
			if err := t.ParseIEs(i.ChildIEs...); err != nil {
				return t, err
			}
			return t, nil
		}
		// If the ChildIEs not already parsed
		return ParseTGPPAccessForwardingActionInformationFields(i.Payload)
	case CreateMAR:
		ies, err := i.CreateMAR()
		if err != nil {
			return nil, err
		}
		if ies.TGPPAccessForwardingActionInformation != nil {
			return ies.TGPPAccessForwardingActionInformation, nil
		}
		return nil, ErrIENotFound
	case UpdateMAR:
		ies, err := i.UpdateMAR()
		if err != nil {
			return nil, err
		}
		if ies.TGPPAccessForwardingActionInformation != nil {
			return ies.TGPPAccessForwardingActionInformation, nil
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
type TGPPAccessForwardingActionInformationFields struct {
	FARID    uint32
	Weight   uint8
	Priority uint8
	URRID    uint32
	RATType  uint8
}

// TGPPAccessForwardingActionInformation returns the IEs above TGPPAccessForwardingActionInformation.
func ParseTGPPAccessForwardingActionInformationFields(b []byte) (*TGPPAccessForwardingActionInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	t := &TGPPAccessForwardingActionInformationFields{}
	if err := t.ParseIEs(ies...); err != nil {
		return t, err
	}
	return t, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (t *TGPPAccessForwardingActionInformationFields) ParseIEs(ies ...*IE) error {

	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case FARID:
			v, err := ie.FARID()
			if err != nil {
				return err
			}
			t.FARID = v
		case Weight:
			v, err := ie.Weight()
			if err != nil {
				return err
			}
			t.Weight = v
		case Priority:
			v, err := ie.Priority()
			if err != nil {
				return err
			}
			t.Priority = v
		case URRID:
			v, err := ie.URRID()
			if err != nil {
				return err
			}
			t.URRID = v
		case RATType:
			v, err := ie.RATType()
			if err != nil {
				return err
			}
			t.RATType = v
		}
	}
	return nil
}
