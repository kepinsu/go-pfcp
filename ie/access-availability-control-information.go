// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewAccessAvailabilityControlInformation creates a new AccessAvailabilityControlInformation IE.
func NewAccessAvailabilityControlInformation(info *IE) *IE {
	return newGroupedIE(AccessAvailabilityControlInformation, 0, info)
}

// AccessAvailabilityControlInformation returns the IEs above AccessAvailabilityControlInformation if the type of IE matches.
func (i *IE) AccessAvailabilityControlInformation() (*AccessAvailabilityControlInformationFields, error) {
	switch i.Type {
	case AccessAvailabilityControlInformation:
		return ParseAccessAvailabilityControlInformationFields(i.Payload)
	case CreateSRR:
		ies, err := i.CreateSRR()
		if err != nil {
			return nil, err
		}
		if ies.AccessAvailabilityControlInformation != nil {
			return ies.AccessAvailabilityControlInformation, nil
		}
		return nil, ErrIENotFound
	case UpdateSRR:
		ies, err := i.UpdateSRR()
		if err != nil {
			return nil, err
		}
		if ies.AccessAvailabilityControlInformation != nil {
			return ies.AccessAvailabilityControlInformation, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// AccessAvailabilityControlInformationFields is a set of fields in CreateSSR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type AccessAvailabilityControlInformationFields struct {
	RequestedAccessAvailabilityInformation uint8
}

func ParseAccessAvailabilityControlInformationFields(b []byte) (*AccessAvailabilityControlInformationFields, error) {

	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	a := &AccessAvailabilityControlInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == RequestedAccessAvailabilityInformation {
			v, err := ie.RequestedAccessAvailabilityInformation()
			if err != nil {
				return a, err
			}
			a.RequestedAccessAvailabilityInformation = v
		}
	}
	return a, nil
}
