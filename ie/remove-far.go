// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRemoveFAR creates a new RemoveFAR IE.
func NewRemoveFAR(far *IE) *IE {
	return newGroupedIE(RemoveFAR, 0, far)
}

// RemoveFAR returns the IEs above RemoveFAR if the type of IE matches.
func (i *IE) RemoveFAR() (*RemoveFARFields, error) {
	if i.Type != RemoveFAR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseRemoveFARFields(i.Payload)
}

// RemoveFARFields is a set of fields in Remove FAR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RemoveFARFields struct {
	FARID uint32
}

// ParseRemoveFARFields returns the IEs above RemoveMBSUnicastParameters.
func ParseRemoveFARFields(b []byte) (*RemoveFARFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &RemoveFARFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case FARID:
			v, err := ie.FARID()
			if err != nil {
				return f, err
			}
			f.FARID = v
		}
	}
	return f, nil
}
