// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPFCPSessionChangeInfo creates a new PFCPSessionChangeInfo IE.
func NewPFCPSessionChangeInfo(ies ...*IE) *IE {
	return newGroupedIE(PFCPSessionChangeInfo, 0, ies...)
}

// PFCPSessionChangeInfo returns the IEs above PFCPSessionChangeInfo if the type of IE matches.
func (i *IE) PFCPSessionChangeInfo() (*PFCPSessionChangeInfoFields, error) {
	if i.Type != PFCPSessionChangeInfo {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParsePFCPSessionChangeInfo(i.Payload)
}

// PFCPSessionChangeInfoFields represents a fields contained in PFCPSessionChangeInfo IE.
type PFCPSessionChangeInfoFields struct {
	FQCSID                  [][]byte
	GroupID                 [][]byte
	CPIPAddress             []*CPIPAddressFields
	AlternativeSMFIPAddress *AlternativeSMFIPAddressFields
}

// ParsePFCPSessionChangeInfo parses b into PFCPSessionChangeInfo.
func ParsePFCPSessionChangeInfo(b []byte) (*PFCPSessionChangeInfoFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &PFCPSessionChangeInfoFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case FQCSID:
			v, err := ie.FQCSID()
			if err != nil {
				return p, err
			}
			p.FQCSID = append(p.FQCSID, v)
		case GroupID:
			v, err := ie.GroupID()
			if err != nil {
				return p, err
			}
			p.GroupID = append(p.GroupID, v)
		case CPIPAddress:
			v, err := ie.CPIPAddress()
			if err != nil {
				return p, err
			}
			p.CPIPAddress = append(p.CPIPAddress, v)
		case AlternativeSMFIPAddress:
			v, err := ie.AlternativeSMFIPAddress()
			if err != nil {
				return p, err
			}
			p.AlternativeSMFIPAddress = v
		}
	}
	return p, nil
}
