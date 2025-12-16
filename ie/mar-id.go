// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewMARID creates a new MARID IE.
func NewMARID(id uint16) *IE {
	return newUint16ValIE(MARID, id)
}

// MARID returns MARID in uint16 if the type of IE matches.
func (i *IE) MARID() (uint16, error) {
	switch i.Type {
	case MARID:
		return i.ValueAsUint16()
	case CreatePDR:
		for _, x := range i.ChildIEs {
			if x.Type == MARID {
				return x.MARID()
			}
		}
		return 0, ErrIENotFound
	case CreateMAR:
		for _, x := range i.ChildIEs {
			if x.Type == MARID {
				return x.MARID()
			}
		}
		return 0, ErrIENotFound
	case RemoveMAR:
		for _, x := range i.ChildIEs {
			if x.Type == MARID {
				return x.MARID()
			}
		}
		return 0, ErrIENotFound
	case UpdateMAR:
		for _, x := range i.ChildIEs {
			if x.Type == MARID {
				return x.MARID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
