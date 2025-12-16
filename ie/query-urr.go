// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewQueryURR creates a new QueryURR IE.
func NewQueryURR(urrID *IE) *IE {
	return newGroupedIE(QueryURR, 0, urrID)
}

// QueryURR returns the IEs above QueryURR if the type of IE matches.
func (i *IE) QueryURR() (*QueryURRFields, error) {
	if i.Type != QueryURR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseQueryURRFields(i.Payload)
}

// QueryURRFields is a set of fields in Query URR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type QueryURRFields struct {
	URRID uint32
}

// ParseQueryURRFields returns the IEs above Query URR .
func ParseQueryURRFields(b []byte) (*QueryURRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &QueryURRFields{}
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
