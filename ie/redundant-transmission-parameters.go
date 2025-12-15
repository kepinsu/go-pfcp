// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRedundantTransmissionParameters creates a new RedundantTransmissionParameters IE.
func NewRedundantTransmissionParameters(ies ...*IE) *IE {
	return newGroupedIE(RedundantTransmissionParameters, 0, ies...)
}

// NewRedundantTransmissionParametersInPDI creates a new RedundantTransmissionParameters IE.
func NewRedundantTransmissionParametersInPDI(fteid, ni *IE) *IE {
	return newGroupedIE(RedundantTransmissionParameters, 0, fteid, ni)
}

// NewRedundantTransmissionParametersInFAR creates a new RedundantTransmissionParameters IE.
func NewRedundantTransmissionParametersInFAR(ohc, ni *IE) *IE {
	return newGroupedIE(RedundantTransmissionParameters, 0, ohc, ni)
}

// RedundantTransmissionParameters returns the IEs above RedundantTransmissionParameters if the type of IE matches.
func (i *IE) RedundantTransmissionParameters() (*RedundantTransmissionParametersField, error) {
	switch i.Type {
	case RedundantTransmissionParameters:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			r := &RedundantTransmissionParametersField{}
			if err := r.ParseIEs(i.ChildIEs...); err != nil {
				return r, err
			}
			return r, nil
		}
		// If the ChildIEs not already parsed
		return ParseRedundantTransmissionParametersField(i.Payload)
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return nil, err
		}
		if ies.PDI != nil && ies.PDI.RedundantTransmissionParameters != nil {
			return ies.PDI.RedundantTransmissionParameters, nil
		}
		return nil, ErrIENotFound
	case PDI:
		ies, err := i.PDI()
		if err != nil {
			return nil, err
		}
		if ies.RedundantTransmissionParameters != nil {
			return ies.RedundantTransmissionParameters, nil
		}
		return nil, ErrIENotFound
	case CreateFAR:
		ies, err := i.CreateFAR()
		if err != nil {
			return nil, err
		}
		if ies.RedundantTransmissionParameters != nil {
			return ies.RedundantTransmissionParameters, err
		}
		return nil, ErrIENotFound
	case UpdateFAR:
		ies, err := i.UpdateFAR()
		if err != nil {
			return nil, err
		}
		if ies.RedundantTransmissionParameters != nil {
			return ies.RedundantTransmissionParameters, err
		}
		return nil, ErrIENotFound
	case CreateTrafficEndpoint:
		ies, err := i.CreateTrafficEndpoint()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == RedundantTransmissionParameters {
				return x.RedundantTransmissionParameters()
			}
		}
		return nil, ErrIENotFound
	case UpdateTrafficEndpoint:
		ies, err := i.UpdateTrafficEndpoint()
		if err != nil {
			return nil, err
		}
		if ies.RedundantTransmissionDetectionParameters != nil {
			return ies.RedundantTransmissionDetectionParameters, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// ParseRedundantTransmissionParametersField returns the IEs above RedundantTransmissionParameters
func ParseRedundantTransmissionParametersField(b []byte) (*RedundantTransmissionParametersField, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	r := &RedundantTransmissionParametersField{}
	if err := r.ParseIEs(ies...); err != nil {
		return r, nil
	}
	return r, nil
}

// RedundantTransmissionParametersField is a set of fields in RedundantTransmissionParametersField IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RedundantTransmissionParametersField struct {
	// Fields only present in PDI IE
	LocalTeID *FTEIDFields

	// Fields only present in FAR IE
	OuterHeaderCreation *OuterHeaderCreationFields

	// Fields present in FAR IE and PDI IE
	NetworkInstance string
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (r *RedundantTransmissionParametersField) ParseIEs(ies ...*IE) error {
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
