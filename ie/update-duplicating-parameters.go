// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateDuplicatingParameters creates a new UpdateDuplicatingParameters IE.
func NewUpdateDuplicatingParameters(ies ...*IE) *IE {
	return newGroupedIE(UpdateDuplicatingParameters, 0, ies...)
}

// UpdateDuplicatingParameters returns the IEs above UpdateDuplicatingParameters if the type of IE matches.
func (i *IE) UpdateDuplicatingParameters() (*UpdateDuplicatingParametersFields, error) {
	switch i.Type {
	case UpdateDuplicatingParameters:
		return ParseUpdateDuplicatingParametersFields(i.Payload)
	case UpdateFAR:
		ies, err := i.UpdateFAR()
		if err != nil {
			return nil, err
		}
		// Return the first IE
		for _, i := range ies.UpdateDuplicatingParameters {
			if i == nil {
				continue
			}
			return i, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// DuplicatingParametersFields is a set of fields in DuplicatingParametersFields IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateDuplicatingParametersFields struct {
	DestinationInterface       uint8
	OuterHeaderCreation        *OuterHeaderCreationFields
	TransportLevelMarking      uint16
	ForwardingPolicy           []byte
	ForwardingPolicyIdentifier string
}

func ParseUpdateDuplicatingParametersFields(b []byte) (*UpdateDuplicatingParametersFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	d := &UpdateDuplicatingParametersFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case DestinationInterface:
			dest, err := ie.DestinationInterface()
			if err != nil {
				return d, err
			}
			d.DestinationInterface = dest
		case OuterHeaderCreation:
			creation, err := ie.OuterHeaderCreation()
			if err != nil {
				return d, err
			}
			d.OuterHeaderCreation = creation
		case TransportLevelMarking:
			transport, err := ie.TransportLevelMarking()
			if err != nil {
				return d, err
			}
			d.TransportLevelMarking = transport
		case ForwardingPolicy:
			policy, err := ie.ForwardingPolicy()
			if err != nil {
				return d, err
			}
			d.ForwardingPolicy = policy
			identifier, err := ie.ForwardingPolicyIdentifier()
			if err != nil {
				return d, err
			}
			d.ForwardingPolicyIdentifier = identifier
		}
	}
	return d, nil
}
