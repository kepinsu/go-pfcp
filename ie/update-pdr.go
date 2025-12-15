// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewUpdatePDR creates a new UpdatePDR IE.
func NewUpdatePDR(ies ...*IE) *IE {
	return newGroupedIE(UpdatePDR, 0, ies...)
}

// UpdatePDR returns the IEs above UpdatePDR if the type of IE matches.
func (i *IE) UpdatePDR() (*UpdatePDRFields, error) {
	if i.Type != UpdatePDR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &UpdatePDRFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseUpdatePDRFields(i.Payload)
}

// UpdatePDRFields is a set of fields in Update PDR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdatePDRFields struct {
	// PDR ID
	ID                 uint16
	OuterHeaderRemoval []byte
	// Precedence
	Precedence uint32
	// PDI
	PDI                       *PDIFields
	FARID                     uint32
	URRID                     []uint32
	QERID                     []uint32
	ActivatePredefinedRules   string
	DeactivatePredefinedRules string
	ActivationTime            time.Time
	DeactivationTime          time.Time
	IPMulticastAddressingInfo []*IPMulticastAddressingInfoField
	TransportDelayReporting   *IE
	RatType                   uint8
}

// ParseUpdatePDRFields returns the IEs above Update PDR
func ParseUpdatePDRFields(b []byte) (*UpdatePDRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	u := &UpdatePDRFields{}
	if err := u.ParseIEs(ies...); err != nil {
		return u, err
	}
	return u, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (u *UpdatePDRFields) ParseIEs(ies ...*IE) error {

	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case PDRID:
			a, err := ie.PDRID()
			if err != nil {
				return err
			}
			u.ID = a
		case Precedence:
			a, err := ie.Precedence()
			if err != nil {
				return err
			}
			u.Precedence = a
		case PDI:
			a, err := ie.PDI()
			if err != nil {
				return err
			}
			u.PDI = a
		case OuterHeaderRemoval:
			a, err := ie.OuterHeaderRemoval()
			if err != nil {
				return err
			}
			u.OuterHeaderRemoval = a
		case FARID:
			a, err := ie.FARID()
			if err != nil {
				return err
			}
			u.FARID = a
		case URRID:
			a, err := ie.URRID()
			if err != nil {
				return err
			}
			u.URRID = append(u.URRID, a)
		case QERID:
			a, err := ie.QERID()
			if err != nil {
				return err
			}
			u.QERID = append(u.QERID, a)
		case ActivatePredefinedRules:
			a, err := ie.ActivatePredefinedRules()
			if err != nil {
				return err
			}
			u.ActivatePredefinedRules = a
		case DeactivatePredefinedRules:
			a, err := ie.DeactivatePredefinedRules()
			if err != nil {
				return err
			}
			u.DeactivatePredefinedRules = a
		case ActivationTime:
			a, err := ie.ActivationTime()
			if err != nil {
				return err
			}
			u.ActivationTime = a
		case DeactivationTime:
			a, err := ie.DeactivationTime()
			if err != nil {
				return err
			}
			u.DeactivationTime = a
		case IPMulticastAddressingInfo:
			v, err := ie.IPMulticastAddressingInfo()
			if err != nil {
				return nil
			}
			u.IPMulticastAddressingInfo = append(u.IPMulticastAddressingInfo, v)
		case TransportDelayReporting:
			u.TransportDelayReporting = ie
		case RATType:
			a, err := ie.RATType()
			if err != nil {
				return err
			}
			u.RatType = a
		}
	}
	return nil
}
