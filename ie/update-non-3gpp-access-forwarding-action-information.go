// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateNonTGPPAccessForwardingActionInformation creates a new UpdateNonTGPPAccessForwardingActionInformation IE.
func NewUpdateNonTGPPAccessForwardingActionInformation(ies ...*IE) *IE {
	return newGroupedIE(UpdateNonTGPPAccessForwardingActionInformation, 0, ies...)
}

// UpdateNonTGPPAccessForwardingActionInformation returns the IEs above UpdateNonTGPPAccessForwardingActionInformation if the type of IE matches.
func (i *IE) UpdateNonTGPPAccessForwardingActionInformation() (*UpdateNonTGPPAccessForwardingActionInformationFields, error) {
	switch i.Type {
	case UpdateNonTGPPAccessForwardingActionInformation:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &UpdateNonTGPPAccessForwardingActionInformationFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseUpdateNonTGPPAccessForwardingActionInformationFields(i.Payload)
	case UpdateMAR:
		ies, err := i.UpdateMAR()
		if err != nil {
			return nil, err
		}
		if ies.UpdateNonTGPPAccessForwardingActionInformation != nil {
			return ies.UpdateNonTGPPAccessForwardingActionInformation, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// UpdateNonTGPPAccessForwardingActionInformationFields is a set of fields in UpdateNonTGPPAccessForwardingActionInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateNonTGPPAccessForwardingActionInformationFields struct {
	FARID    uint32
	Weight   uint8
	Priority uint8
	URRID    uint32
	RATType  uint8
}

// UpdateNonTGPPAccessForwardingActionInformation returns the IEs above UpdateNonTGPPAccessForwardingActionInformation.
func ParseUpdateNonTGPPAccessForwardingActionInformationFields(b []byte) (*UpdateNonTGPPAccessForwardingActionInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	t := &UpdateNonTGPPAccessForwardingActionInformationFields{}
	if err := t.ParseIEs(ies...); err != nil {
		return t, err
	}
	return t, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (t *UpdateNonTGPPAccessForwardingActionInformationFields) ParseIEs(ies ...*IE) error {

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
