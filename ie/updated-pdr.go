// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdatedPDR creates a new UpdatedPDR IE.
func NewUpdatedPDR(ies ...*IE) *IE {
	return newGroupedIE(UpdatedPDR, 0, ies...)
}

// UpdatedPDR returns the IEs above UpdatedPDR if the type of IE matches.
func (i *IE) UpdatedPDR() (*UpdatedPDRFields, error) {
	if i.Type != UpdatedPDR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &UpdatedPDRFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseUpdatedPDRFields(i.Payload)
}

// UpdatedPDRFields is a set of fields in UpdatedPDR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdatedPDRFields struct {
	PDRID uint16
	// Cannot determine if the Local TEID is a Local F-TEID or
	// Local F-TEID For Redundant Transmission
	LocalFTEID  []*FTEIDFields
	UEIPAddress []*UEIPAddressFields
}

// ParseUpdatedPDRFields returns the IEs above UpdatedPDR if the type of IE matches.
func ParseUpdatedPDRFields(b []byte) (*UpdatedPDRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	u := &UpdatedPDRFields{}
	if err := u.ParseIEs(ies...); err != nil {
		return u, err
	}
	return u, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (u *UpdatedPDRFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case PDRID:
			id, err := ie.PDRID()
			if err != nil {
				return err
			}
			u.PDRID = id
		case FTEID:
			v, err := ie.FTEID()
			if err != nil {
				return err
			}
			u.LocalFTEID = append(u.LocalFTEID, v)
		case UEIPAddress:
			ip, err := ie.UEIPAddress()
			if err != nil {
				return err
			}
			u.UEIPAddress = append(u.UEIPAddress, ip)
		}
	}
	return nil
}
