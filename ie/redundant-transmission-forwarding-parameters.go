// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRedundantTransmissionForwardingParameters creates a new RedundantTransmissionForwardingParameters IE.
func NewRedundantTransmissionForwardingParameters(ies ...*IE) *IE {
	return newGroupedIE(RedundantTransmissionForwardingParameters, 0, ies...)
}

// RedundantTransmissionForwardingParameters returns the IEs above RedundantTransmissionForwardingParameters if the type of IE matches.
func (i *IE) RedundantTransmissionForwardingParameters() (*RedundantTransmissionForwardingParametersField, error) {
	switch i.Type {
	case RedundantTransmissionForwardingParameters:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &RedundantTransmissionForwardingParametersField{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseRedundantTransmissionForwardingParametersField(i.Payload)
	case CreateFAR:
		ies, err := i.CreateFAR()
		if err != nil {
			return nil, err
		}
		if ies.RedundantTransmissionForwardingParameters != nil {
			return ies.RedundantTransmissionForwardingParameters, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// RedundantTransmissionForwardingParametersField is a set of fields in RedundantTransmissionForwardingParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RedundantTransmissionForwardingParametersField struct {
	// Fields only present in PDI IE
	LocalTeID *FTEIDFields

	// Fields only present in FAR IE
	OuterHeaderCreation *OuterHeaderCreationFields

	// Fields present in FAR IE and PDI IE
	NetworkInstance string
}

func ParseRedundantTransmissionForwardingParametersField(b []byte) (*RedundantTransmissionForwardingParametersField, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	r := &RedundantTransmissionForwardingParametersField{}
	if err := r.ParseIEs(ies...); err != nil {
		return r, nil
	}

	return r, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (r *RedundantTransmissionForwardingParametersField) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case FTEID:
			fteid, err := ie.FTEID()
			if err != nil {
				return err
			}
			r.LocalTeID = fteid
		case OuterHeaderCreation:
			header, err := ie.OuterHeaderCreation()
			if err != nil {
				return err
			}
			r.OuterHeaderCreation = header
		case NetworkInstance:
			network, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			r.NetworkInstance = network
		}
	}
	return nil
}
