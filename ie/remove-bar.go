// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRemoveBAR creates a new RemoveBAR IE.
func NewRemoveBAR(barID *IE) *IE {
	return newGroupedIE(RemoveBAR, 0, barID)
}

// RemoveBAR returns the IEs above RemoveBAR if the type of IE matches.
func (i *IE) RemoveBAR() (*RemoveBARFields, error) {
	if i.Type != RemoveBAR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseRemoveBARFields(i.Payload)
}

// RemoveBARFields is a set of fields in Remove BAR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RemoveBARFields struct {
	BARID uint8
}

// ParseRemoveBARFields returns the IEs above RemoveMBSUnicastParameters.
func ParseRemoveBARFields(b []byte) (*RemoveBARFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &RemoveBARFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == BARID {
			v, err := ie.BARID()
			if err != nil {
				return f, err
			}
			f.BARID = v
		}
	}
	return f, nil
}
