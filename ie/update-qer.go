// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateQER Updates a new UpdateQER IE.
func NewUpdateQER(ies ...*IE) *IE {
	return newGroupedIE(UpdateQER, 0, ies...)
}

// UpdateQER returns the IEs above UpdateQER if the type of IE matches.
func (i *IE) UpdateQER() (*UpdateQERFields, error) {
	if i.Type != UpdateQER {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &UpdateQERFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseUpdateQERFields(i.Payload)
}

// UpdateQERFields is a set of fields in UpdateQER IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateQERFields struct {
	QERID                 uint32
	QERCorrelationID      uint32
	GateStatus            Gate
	MBR                   *MBRFields
	GBR                   *GBRFields
	PacketRate            *PacketRateFields
	PacketRateStatus      *PacketRateStatusFields
	DLFlowLevelMarking    *DLFlowLevelMarkingFields
	QFI                   uint8
	ReflectiveQoS         uint8
	PagingPolicyIndicator uint8
	AveragingWindows      uint32
	QERControlIndications uint8
}

// ParseUpdateQERFields returns the IEs above UpdateURR if the type of IE matches.
func ParseUpdateQERFields(b []byte) (*UpdateQERFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	q := &UpdateQERFields{}
	if err := q.ParseIEs(ies...); err != nil {
		return q, err
	}
	return q, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (u *UpdateQERFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case QERID:
			a, err := ie.QERID()
			if err != nil {
				return err
			}
			u.QERID = a
		case QERCorrelationID:
			a, err := ie.QERCorrelationID()
			if err != nil {
				return err
			}
			u.QERCorrelationID = a
		case GateStatus:
			a, err := ie.GateStatus()
			if err != nil {
				return err
			}
			u.GateStatus = a
		case MBR:
			m, err := ie.MBR()
			if err != nil {
				return err
			}
			u.MBR = m
		case GBR:
			m, err := ie.GBR()
			if err != nil {
				return err
			}
			u.GBR = m
		case PacketRate:
			m, err := ie.PacketRate()
			if err != nil {
				return err
			}
			u.PacketRate = m
		case PacketRateStatus:
			m, err := ie.PacketRateStatus()
			if err != nil {
				return err
			}
			u.PacketRateStatus = m
		case DLFlowLevelMarking:
			m, err := ie.DLFlowLevelMarking()
			if err != nil {
				return err
			}
			u.DLFlowLevelMarking = m
		case QFI:
			m, err := ie.QFI()
			if err != nil {
				return err
			}
			u.QFI = m
		case RQI:
			m, err := ie.RQI()
			if err != nil {
				return err
			}
			u.ReflectiveQoS = m
		case PagingPolicyIndicator:
			m, err := ie.PagingPolicyIndicator()
			if err != nil {
				return err
			}
			u.PagingPolicyIndicator = m
		case AveragingWindow:
			m, err := ie.AveragingWindow()
			if err != nil {
				return err
			}
			u.AveragingWindows = m
		case QERControlIndications:
			m, err := ie.QERControlIndications()
			if err != nil {
				return err
			}
			u.QERControlIndications = m
		}
	}
	return nil
}
