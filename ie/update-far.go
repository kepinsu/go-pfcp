// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateFAR creates a new UpdateFAR IE.
func NewUpdateFAR(ies ...*IE) *IE {
	return newGroupedIE(UpdateFAR, 0, ies...)
}

// UpdateFAR returns the IEs above UpdateFAR if the type of IE matches.
func (i *IE) UpdateFAR() (*UpdateFARFields, error) {
	if i.Type != UpdateFAR {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &UpdateFARFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseUpdateFARFields(i.Payload)
}

// UpdateFARFields is a set of fields in Update FAR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateFARFields struct {
	FARID                                     uint32
	ApplyAction                               *ApplyActionFields
	UpdateForwardingParameters                *UpdateForwardingParametersFields
	UpdateDuplicatingParameters               []*UpdateDuplicatingParametersFields
	BARID                                     uint8
	RedundantTransmissionParameters           *RedundantTransmissionParametersField
	RedundantTransmissionForwardingParameters *RedundantTransmissionForwardingParametersField
	AddMBSUnicastParameters                   []*AddMBSUnicastParametersFields
	RemoveMBSUnicastParameters                []*RemoveMBSUnicastParametersFields
}

// ParseUpdateFARFields returns the IEs above Update FAR
func ParseUpdateFARFields(b []byte) (*UpdateFARFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &UpdateFARFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *UpdateFARFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case FARID:
			farid, err := ie.FARID()
			if err != nil {
				return err
			}
			far.FARID = farid
		case ApplyAction:
			apply, err := ie.ApplyAction()
			if err != nil {
				return err
			}
			far.ApplyAction = apply
		case ForwardingParameters:
			apply, err := ie.ApplyAction()
			if err != nil {
				return err
			}
			far.ApplyAction = apply
		case UpdateDuplicatingParameters:
			duplicating, err := ie.UpdateDuplicatingParameters()
			if err != nil {
				return err
			}
			far.UpdateDuplicatingParameters = append(far.UpdateDuplicatingParameters, duplicating)
		case BARID:
			barID, err := ie.BARID()
			if err != nil {
				return err
			}
			far.BARID = barID
		case RedundantTransmissionForwardingParameters:
			forward, err := ie.RedundantTransmissionForwardingParameters()
			if err != nil {
				return err
			}
			far.RedundantTransmissionForwardingParameters = forward
		case RedundantTransmissionParameters:
			forward, err := ie.RedundantTransmissionParameters()
			if err != nil {
				return err
			}
			far.RedundantTransmissionParameters = forward
		case RemoveMBSUnicastParameters:
			v, err := ie.RemoveMBSUnicastParameters()
			if err != nil {
				return err
			}
			far.RemoveMBSUnicastParameters = append(far.RemoveMBSUnicastParameters, v)
		case AddMBSUnicastParameters:
			v, err := ie.AddMBSUnicastParameters()
			if err != nil {
				return err
			}
			far.AddMBSUnicastParameters = append(far.AddMBSUnicastParameters, v)
		}
	}
	return nil
}
