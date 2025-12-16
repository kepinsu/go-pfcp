// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRemovePDR creates a new RemovePDR IE.
func NewRemovePDR(pdr *IE) *IE {
	return newGroupedIE(RemovePDR, 0, pdr)
}

// RemovePDR returns the IEs above RemovePDR if the type of IE matches.
func (i *IE) RemovePDR() (*RemovePDRFields, error) {
	if i.Type != RemovePDR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseRemovePDRFields(i.Payload)
}

// RemovePDRFields is a set of fields in Remove PDR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RemovePDRFields struct {
	PDRID uint16
}

// ParseRemovePDRFields returns the IEs above RemoveMBSUnicastParameters.
func ParseRemovePDRFields(b []byte) (*RemovePDRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &RemovePDRFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case PDRID:
			v, err := ie.PDRID()
			if err != nil {
				return f, err
			}
			f.PDRID = v
		}
	}
	return f, nil
}
