// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRemoveSRR creates a new RemoveSRR IE.
func NewRemoveSRR(srr *IE) *IE {
	return newGroupedIE(RemoveSRR, 0, srr)
}

// RemoveSRR returns the IEs above RemoveSRR if the type of IE matches.
func (i *IE) RemoveSRR() (*RemoveSRRFields, error) {
	if i.Type != RemoveSRR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseRemoveSRRFields(i.Payload)
}

// RemoveSRRFields is a set of fields in Remove QER IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RemoveSRRFields struct {
	SRRID uint8
}

// ParseRemoveSRRFields returns the IEs above RemoveMBSUnicastParameters.
func ParseRemoveSRRFields(b []byte) (*RemoveSRRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &RemoveSRRFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == SRRID {
			v, err := ie.SRRID()
			if err != nil {
				return f, err
			}
			f.SRRID = v
		}
	}
	return f, nil
}
