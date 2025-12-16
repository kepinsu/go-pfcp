// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateSRR Updates a new UpdateSRR IE.
func NewUpdateSRR(ies ...*IE) *IE {
	return newGroupedIE(UpdateSRR, 0, ies...)
}

// UpdateSRR returns the IEs above UpdateSRR if the type of IE matches.
func (i *IE) UpdateSRR() (*UpdateSRRFields, error) {
	switch i.Type {
	case UpdateSRR:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &UpdateSRRFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseUpdateSRRFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// UpdateSRRFields is a set of fields in UpdateSSR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateSRRFields struct {
	SSRID                                     uint8
	AccessAvailabilityControlInformation      *AccessAvailabilityControlInformationFields
	QoSMonitoringPerQoSFlowControlInformation *QoSMonitoringPerQoSFlowControlInformationFields
	DirectReportingInformation                *DirectReportingInformationFields
}

// ParseUpdateSRRFields returns the IEs above UpdateSRR
func ParseUpdateSRRFields(b []byte) (*UpdateSRRFields, error) {

	// Parse all IES here
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	srr := &UpdateSRRFields{}
	if err := srr.ParseIEs(ies...); err != nil {
		return srr, nil
	}
	return srr, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (srr *UpdateSRRFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case SRRID:
			v, err := ie.SRRID()
			if err != nil {
				return err
			}
			srr.SSRID = v
		case AccessAvailabilityControlInformation:
			v, err := ie.AccessAvailabilityControlInformation()
			if err != nil {
				return err
			}
			srr.AccessAvailabilityControlInformation = v
		case QoSMonitoringPerQoSFlowControlInformation:
			v, err := ie.QoSMonitoringPerQoSFlowControlInformation()
			if err != nil {
				return err
			}
			srr.QoSMonitoringPerQoSFlowControlInformation = v
		case DirectReportingInformation:
			v, err := ie.DirectReportingInformation()
			if err != nil {
				return err
			}
			srr.DirectReportingInformation = v
		}
	}
	return nil
}
