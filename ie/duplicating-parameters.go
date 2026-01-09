// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewDuplicatingParameters creates a new DuplicatingParameters IE.
func NewDuplicatingParameters(ies ...*IE) *IE {
	return newGroupedIE(DuplicatingParameters, 0, ies...)
}

// DuplicatingParameters returns the IEs above DuplicatingParameters if the type of IE matches.
func (i *IE) DuplicatingParameters() (*DuplicatingParametersFields, error) {
	switch i.Type {
	case DuplicatingParameters:
		return ParseDuplicatingParametersFields(i.Payload)
	case CreateFAR:
		ies, err := i.CreateFAR()
		if err != nil {
			return nil, err
		}
		for _, i := range ies.DuplicatingParameters {
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
type DuplicatingParametersFields struct {
	DestinationInterface       uint8
	OuterHeaderCreation        *OuterHeaderCreationFields
	TransportLevelMarking      uint16
	ForwardingPolicy           []byte
	ForwardingPolicyIdentifier string
}

func ParseDuplicatingParametersFields(b []byte) (*DuplicatingParametersFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	d := &DuplicatingParametersFields{}
	if err := d.ParseIEs(ies...); err != nil {
		return d, err
	}
	return d, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (d *DuplicatingParametersFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case DestinationInterface:
			dest, err := ie.DestinationInterface()
			if err != nil {
				return err
			}
			d.DestinationInterface = dest
		case OuterHeaderCreation:
			creation, err := ie.OuterHeaderCreation()
			if err != nil {
				return err
			}
			d.OuterHeaderCreation = creation
		case TransportLevelMarking:
			transport, err := ie.TransportLevelMarking()
			if err != nil {
				return err
			}
			d.TransportLevelMarking = transport
		case ForwardingPolicy:
			policy, err := ie.ForwardingPolicy()
			if err != nil {
				return err
			}
			d.ForwardingPolicy = policy
			identifier, err := ie.ForwardingPolicyIdentifier()
			if err != nil {
				return err
			}
			d.ForwardingPolicyIdentifier = identifier
		}
	}
	return nil
}
