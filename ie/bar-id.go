// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewBARID creates a new BARID IE.
func NewBARID(id uint8) *IE {
	return newUint8ValIE(BARID, id)
}

// BARID returns BARID in uint8 if the type of IE matches.
func (i *IE) BARID() (uint8, error) {
	switch i.Type {
	case BARID:
		return i.ValueAsUint8()
	case CreateFAR:
		for _, x := range i.ChildIEs {
			if x.Type == BARID {
				return x.BARID()
			}
		}
		return 0, ErrIENotFound
	case UpdateFAR:
		for _, x := range i.ChildIEs {
			if x.Type == BARID {
				return x.BARID()
			}
		}
		return 0, ErrIENotFound
	case CreateBAR:
		ies, err := i.CreateBAR()
		if err != nil {
			return 0, err
		}
		return ies.BarID, nil
	case UpdateBARWithinSessionReportResponse,
		UpdateBARWithinSessionModificationRequest:
		for _, x := range i.ChildIEs {
			if x.Type == BARID {
				return x.BARID()
			}
		}
		return 0, ErrIENotFound
	case RemoveBAR:
		for _, x := range i.ChildIEs {
			if x.Type == BARID {
				return x.BARID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
