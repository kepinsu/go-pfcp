// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewCreateFAR creates a new CreateFAR IE.
func NewCreateFAR(ies ...*IE) *IE {
	return newGroupedIE(CreateFAR, 0, ies...)
}

// CreateFAR returns the IEs above CreateFAR if the type of IE matches.
func (i *IE) CreateFAR() (*CreateFARFields, error) {
	if i.Type != CreateFAR {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return ParseCreateFARFields(i.Payload)
}

// CreateFARFields is a set of fields in CreateFAR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type CreateFARFields struct {
	FARID                                     uint32
	ApplyAction                               *ApplyActionFields
	ForwardingParameters                      *ForwardingParametersFields
	DuplicatingParameters                     []*DuplicatingParametersFields
	BARID                                     uint8
	RedundantTransmissionParameters           *RedundantTransmissionParametersField
	RedundantTransmissionForwardingParameters *RedundantTransmissionForwardingParametersField
	MBSMulticastParameters                    *MBSMulticastParametersFields
	AddMBSUnicastParameters                   []*AddMBSUnicastParametersFields
}

// ParseCreateFARFields returns the IEs above CreateFARFields if the type of IE matches.
func ParseCreateFARFields(b []byte) (*CreateFARFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &CreateFARFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case FARID:
			farid, err := ie.FARID()
			if err != nil {
				return far, err
			}
			far.FARID = farid
		case ApplyAction:
			apply, err := ie.ApplyAction()
			if err != nil {
				return far, err
			}
			far.ApplyAction = apply
		case ForwardingParameters:
			apply, err := ie.ApplyAction()
			if err != nil {
				return far, err
			}
			far.ApplyAction = apply
		case DuplicatingParameters:
			duplicating, err := ie.DuplicatingParameters()
			if err != nil {
				return far, err
			}
			far.DuplicatingParameters = append(far.DuplicatingParameters, duplicating)
		case BARID:
			barID, err := ie.BARID()
			if err != nil {
				return far, err
			}
			far.BARID = barID
		case RedundantTransmissionParameters:
			forward, err := ie.RedundantTransmissionParameters()
			if err != nil {
				return far, err
			}
			far.RedundantTransmissionParameters = forward
		case RedundantTransmissionForwardingParameters:
			forward, err := ie.RedundantTransmissionForwardingParameters()
			if err != nil {
				return far, err
			}
			far.RedundantTransmissionForwardingParameters = forward
		case MBSMulticastParameters:
			v, err := ie.MBSMulticastParameters()
			if err != nil {
				return far, err
			}
			far.MBSMulticastParameters = v
		case AddMBSUnicastParameters:
			v, err := ie.AddMBSUnicastParameters()
			if err != nil {
				return far, err
			}
			far.AddMBSUnicastParameters = append(far.AddMBSUnicastParameters, v)
		}
	}
	return far, nil
}
