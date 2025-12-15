// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPFCPSessionRetentionInformation creates a new PFCPSessionRetentionInformation IE.
func NewPFCPSessionRetentionInformation(cpIP *IE) *IE {
	return newGroupedIE(PFCPSessionRetentionInformation, 0, cpIP)
}

// PFCPSessionRetentionInformation returns the IEs above PFCPSessionRetentionInformation if the type of IE matches.
func (i *IE) PFCPSessionRetentionInformation() (*PFCPSessionRetentionInformationFields, error) {
	if i.Type != PFCPSessionRetentionInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return ParsePFCPSessionRetentionInformationFields(i.Payload)
}

// PFCPSessionRetentionInformationFields represents a fields contained in PFCPSessionRetentionInformation IE.
type PFCPSessionRetentionInformationFields struct {
	CPPFCPEntityIPAddress *CPPFCPEntityIPAddressFields
}

// ParsePFCPSessionRetentionInformationFields returns the IEs above Update FAR
func ParsePFCPSessionRetentionInformationFields(b []byte) (*PFCPSessionRetentionInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &PFCPSessionRetentionInformationFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *PFCPSessionRetentionInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case CPPFCPEntityIPAddress:
			v, err := ie.CPPFCPEntityIPAddress()
			if err != nil {
				return err
			}
			far.CPPFCPEntityIPAddress = v
		}
	}
	return nil
}
