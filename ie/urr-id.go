// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewURRID creates a new URRID IE.
func NewURRID(id uint32) *IE {
	return newUint32ValIE(URRID, id)
}

// URRID returns URRID in uint32 if the type of IE matches.
func (i *IE) URRID() (uint32, error) {
	switch i.Type {
	case URRID:
		return i.ValueAsUint32()
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return 0, err
		}
		for _, id := range ies.URRID {
			return id, nil
		}
		return 0, ErrIENotFound
	case CreateURR:
		ies, err := i.CreateURR()
		if err != nil {
			return 0, err
		}
		return ies.URRID, nil
	case UpdateURR:
		for _, x := range i.ChildIEs {
			if x.Type == URRID {
				return x.URRID()
			}
		}
		return 0, ErrIENotFound
	case UpdatePDR:
		ies, err := i.UpdatePDR()
		if err != nil {
			return 0, err
		}
		for _, id := range ies.URRID {
			return id, nil
		}
		return 0, ErrIENotFound
	case RemoveURR:
		for _, x := range i.ChildIEs {
			if x.Type == URRID {
				return x.URRID()
			}
		}
		return 0, ErrIENotFound
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest:
		for _, x := range i.ChildIEs {
			if x.Type == URRID {
				return x.URRID()
			}
		}
		return 0, ErrIENotFound
	case CreateMAR:
		ies, err := i.CreateMAR()
		if err != nil {
			return 0, err
		}
		if ies.TGPPAccessForwardingActionInformation != nil {
			return ies.TGPPAccessForwardingActionInformation.URRID, nil
		}
		if ies.NonTGPPAccessForwardingActionInformation != nil {
			return ies.NonTGPPAccessForwardingActionInformation.URRID, nil
		}
		return 0, ErrIENotFound
	case UpdateMAR:
		for _, x := range i.ChildIEs {
			switch x.Type {
			case TGPPAccessForwardingActionInformation, NonTGPPAccessForwardingActionInformation:
				return x.URRID()
			}
		}
		return 0, ErrIENotFound
	case QueryURR:
		for _, x := range i.ChildIEs {
			if x.Type == URRID {
				return x.URRID()
			}
		}
		return 0, ErrIENotFound
	case TGPPAccessForwardingActionInformation:
		ies, err := i.TGPPAccessForwardingActionInformation()
		if err != nil {
			return 0, err
		}
		return ies.URRID, nil
	case NonTGPPAccessForwardingActionInformation:
		ies, err := i.NonTGPPAccessForwardingActionInformation()
		if err != nil {
			return 0, err
		}
		return ies.URRID, nil
	case UpdateTGPPAccessForwardingActionInformation:
		for _, x := range i.ChildIEs {
			if x.Type == URRID {
				return x.URRID()
			}
		}
		return 0, ErrIENotFound
	case UpdateNonTGPPAccessForwardingActionInformation:
		for _, x := range i.ChildIEs {
			if x.Type == URRID {
				return x.URRID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// IsAllocatedByCPFunction reports whether URRID is allocated by CP Function.
func (i *IE) IsAllocatedByCPFunction() bool {
	if i.Type != URRID {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return (i.Payload[0]>>7)&0x01 == 1
}

// IsAllocatedByUPFunction reports whether URRID is allocated by UP Function.
func (i *IE) IsAllocatedByUPFunction() bool {
	if i.Type != URRID {
		return false
	}
	if len(i.Payload) < 1 {
		return false
	}

	return (i.Payload[0]>>7)&0x01 != 1
}
