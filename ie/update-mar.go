// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateMAR Updates a new UpdateMAR IE.
func NewUpdateMAR(ies ...*IE) *IE {
	return newGroupedIE(UpdateMAR, 0, ies...)
}

// UpdateMAR returns the IEs above UpdateMAR if the type of IE matches.
func (i *IE) UpdateMAR() (*UpdateMARFields, error) {
	if i.Type != UpdateMAR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &UpdateMARFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseUpdateMarFields(i.Payload)
}

// UpdateMARFields is a set of fields in UpdateMAR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateMARFields struct {
	// MAR ID
	ID                                             uint16
	SteeringFunctionality                          uint8
	SteeringMode                                   uint8
	UpdateTGPPAccessForwardingActionInformation    *UpdateTGPPAccessForwardingActionInformationFields
	UpdateNonTGPPAccessForwardingActionInformation *UpdateNonTGPPAccessForwardingActionInformationFields
	TGPPAccessForwardingActionInformation          *TGPPAccessForwardingActionInformationFields
	NonTGPPAccessForwardingActionInformation       *NonTGPPAccessForwardingActionInformationFields
	Thresholds                                     *ThresholdsFields
	SteeringModeIndicator                          *SteeringModeIndicatorFields
}

// UpdateMAR returns the IEs above UpdateMAR.
func ParseUpdateMarFields(b []byte) (*UpdateMARFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	mar := &UpdateMARFields{}
	if err := mar.ParseIEs(ies...); err != nil {
		return mar, err
	}
	return mar, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (c *UpdateMARFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case MARID:
			a, err := ie.MARID()
			if err != nil {
				return err
			}
			c.ID = a
		case SteeringFunctionality:
			a, err := ie.SteeringFunctionality()
			if err != nil {
				return err
			}
			c.SteeringFunctionality = a

		case SteeringMode:
			a, err := ie.SteeringMode()
			if err != nil {
				return err
			}
			c.SteeringMode = a
		case UpdateTGPPAccessForwardingActionInformation:
			a, err := ie.UpdateTGPPAccessForwardingActionInformation()
			if err != nil {
				return err
			}
			c.UpdateTGPPAccessForwardingActionInformation = a
		case UpdateNonTGPPAccessForwardingActionInformation:
			a, err := ie.UpdateNonTGPPAccessForwardingActionInformation()
			if err != nil {
				return err
			}
			c.UpdateNonTGPPAccessForwardingActionInformation = a
		case TGPPAccessForwardingActionInformation:
			a, err := ie.TGPPAccessForwardingActionInformation()
			if err != nil {
				return err
			}
			c.TGPPAccessForwardingActionInformation = a
		case NonTGPPAccessForwardingActionInformation:
			a, err := ie.NonTGPPAccessForwardingActionInformation()
			if err != nil {
				return err
			}
			c.NonTGPPAccessForwardingActionInformation = a
		case Thresholds:
			a, err := ie.Thresholds()
			if err != nil {
				return err
			}
			c.Thresholds = a
		case SteeringModeIndicator:
			a, err := ie.SteeringModeIndicator()
			if err != nil {
				return err
			}
			c.SteeringModeIndicator = a
		}
	}
	return nil
}
