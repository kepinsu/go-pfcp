// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRemoveURR creates a new RemoveURR IE.
func NewRemoveURR(urr *IE) *IE {
	return newGroupedIE(RemoveURR, 0, urr)
}

// RemoveURR returns the IEs above RemoveURR if the type of IE matches.
func (i *IE) RemoveURR() (*RemoveURRFields, error) {
	if i.Type != RemoveURR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseRemoveURRFields(i.Payload)
}

// RemoveURRFields is a set of fields in Remove URR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RemoveURRFields struct {
	URRID uint32
}

// ParseRemoveURRFields returns the IEs above Remove URR.
func ParseRemoveURRFields(b []byte) (*RemoveURRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &RemoveURRFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case URRID:
			v, err := ie.URRID()
			if err != nil {
				return f, err
			}
			f.URRID = v
		}
	}
	return f, nil
}
