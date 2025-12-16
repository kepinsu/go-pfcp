// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewCreatedPDR creates a new CreatedPDR IE.
func NewCreatedPDR(ies ...*IE) *IE {
	return newGroupedIE(CreatedPDR, 0, ies...)
}

// CreatedPDR returns the IEs above CreatedPDR if the type of IE matches.
func (i *IE) CreatedPDR() (*CreatedPDRFields, error) {
	if i.Type != CreatedPDR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &CreatedPDRFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseCreatedPDRFields(i.Payload)
}

// CreatedPDRFields is a set of fields in CreatedPDR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type CreatedPDRFields struct {
	// PDR ID
	ID uint16
	// Cannot determine if the Local TEID is a Local F-TEID or
	// Local F-TEID For Redundant Transmission
	LocalTeid          []*FTEIDFields
	UEIPAddress        []*UEIPAddressFields
	LocalIngressTunnel *LocalIngressTunnelFields
}

// ParseCreatedPDRFields returns the IEs above CreatedPDR
func ParseCreatedPDRFields(b []byte) (*CreatedPDRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	c := &CreatedPDRFields{}
	if err := c.ParseIEs(ies...); err != nil {
		return c, err
	}

	return c, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (c *CreatedPDRFields) ParseIEs(ies ...*IE) error {
	for _, i := range ies {
		if i == nil {
			continue
		}
		switch i.Type {
		case PDRID:
			v, err := i.PDRID()
			if err != nil {
				return nil
			}
			c.ID = v
		case FTEID:
			v, err := i.FTEID()
			if err != nil {
				return nil
			}
			c.LocalTeid = append(c.LocalTeid, v)
		case UEIPAddress:
			v, err := i.UEIPAddress()
			if err != nil {
				return nil
			}
			c.UEIPAddress = append(c.UEIPAddress, v)
		case LocalIngressTunnel:
			v, err := i.LocalIngressTunnel()
			if err != nil {
				return nil
			}
			c.LocalIngressTunnel = v
		}
	}
	return nil
}
