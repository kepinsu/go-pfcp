// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewCreateQER creates a new CreateQER IE.
func NewCreateQER(ies ...*IE) *IE {
	return newGroupedIE(CreateQER, 0, ies...)
}

// CreateQER returns the IEs above CreateQER if the type of IE matches.
func (i *IE) CreateQER() (*CreateQERFields, error) {
	if i.Type != CreateQER {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		c := &CreateQERFields{}
		if err := c.ParseIEs(i.ChildIEs...); err != nil {
			return c, err
		}
		return c, nil
	}
	// If the ChildIEs not already parsed
	return ParseCreateQERFields(i.Payload)
}

// CreateQERFields is a set of fields in CreateQER IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type CreateQERFields struct {
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

// ParseCreateQERFields returns the IEs above UpdateURR if the type of IE matches.
func ParseCreateQERFields(b []byte) (*CreateQERFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	q := &CreateQERFields{}
	if err := q.ParseIEs(ies...); err != nil {
		return q, nil
	}
	return q, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (q *CreateQERFields) ParseIEs(ies ...*IE) error {
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
			q.QERID = a
		case QERCorrelationID:
			a, err := ie.QERCorrelationID()
			if err != nil {
				return err
			}
			q.QERCorrelationID = a
		case GateStatus:
			a, err := ie.GateStatus()
			if err != nil {
				return err
			}
			q.GateStatus = a
		case MBR:
			m, err := ie.MBR()
			if err != nil {
				return err
			}
			q.MBR = m
		case GBR:
			m, err := ie.GBR()
			if err != nil {
				return err
			}
			q.GBR = m
		case PacketRate:
			m, err := ie.PacketRate()
			if err != nil {
				return err
			}
			q.PacketRate = m
		case PacketRateStatus:
			m, err := ie.PacketRateStatus()
			if err != nil {
				return err
			}
			q.PacketRateStatus = m
		case DLFlowLevelMarking:
			m, err := ie.DLFlowLevelMarking()
			if err != nil {
				return err
			}
			q.DLFlowLevelMarking = m
		case QFI:
			m, err := ie.QFI()
			if err != nil {
				return err
			}
			q.QFI = m
		case RQI:
			m, err := ie.RQI()
			if err != nil {
				return err
			}
			q.ReflectiveQoS = m
		case PagingPolicyIndicator:
			m, err := ie.PagingPolicyIndicator()
			if err != nil {
				return err
			}
			q.PagingPolicyIndicator = m
		case AveragingWindow:
			m, err := ie.AveragingWindow()
			if err != nil {
				return err
			}
			q.AveragingWindows = m
		case QERControlIndications:
			m, err := ie.QERControlIndications()
			if err != nil {
				return err
			}
			q.QERControlIndications = m
		}
	}
	return nil
}
