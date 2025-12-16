// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewEthernetContextInformation creates a new EthernetContextInformation IE.
func NewEthernetContextInformation(mac *IE) *IE {
	return newGroupedIE(EthernetContextInformation, 0, mac)
}

// EthernetContextInformation returns the IEs above EthernetContextInformation if the type of IE matches.
func (i *IE) EthernetContextInformation() (*EthernetContextInformationFields, error) {
	if i.Type != EthernetContextInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		e := &EthernetContextInformationFields{}
		if err := e.ParseIEs(i.ChildIEs...); err != nil {
			return e, err
		}
		return e, nil
	}
	// If the ChildIEs not already parsed
	return ParseEthernetContextInformationFields(i.Payload)
}

// EthernetContextInformationFields is a set of fields in EthernetContextInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type EthernetContextInformationFields struct {
	MACAddressesDetected *MACAddressesDetectedFields
}

// ParseEthernetContextInformationFields returns the IEs above UpdateURR if the type of IE matches.
func ParseEthernetContextInformationFields(b []byte) (*EthernetContextInformationFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	q := &EthernetContextInformationFields{}
	if err := q.ParseIEs(ies...); err != nil {
		return q, nil
	}
	return q, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (e *EthernetContextInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case MACAddressesDetected:
			a, err := ie.MACAddressesDetected()
			if err != nil {
				return err
			}
			e.MACAddressesDetected = a
		}
	}
	return nil
}
