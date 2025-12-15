// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRemoveMAR creates a new RemoveMAR IE.
func NewRemoveMAR(marID *IE) *IE {
	return newGroupedIE(RemoveMAR, 0, marID)
}

// RemoveMAR returns the IEs above RemoveMAR if the type of IE matches.
func (i *IE) RemoveMAR() (*RemoveMARFields, error) {
	if i.Type != RemoveMAR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseRemoveMARFields(i.Payload)
}

// RemoveMARFields is a set of fields in Remove MAR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RemoveMARFields struct {
	MARID uint16
}

// ParseRemoveMARFields returns the IEs above RemoveMBSUnicastParameters.
func ParseRemoveMARFields(b []byte) (*RemoveMARFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &RemoveMARFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case MARID:
			v, err := ie.MARID()
			if err != nil {
				return f, err
			}
			f.MARID = v
		}
	}
	return f, nil
}
