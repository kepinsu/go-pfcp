// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewCreatedTrafficEndpoint creates a new CreatedTrafficEndpoint IE.
func NewCreatedTrafficEndpoint(ies ...*IE) *IE {
	return newGroupedIE(CreatedTrafficEndpoint, 0, ies...)
}

// CreatedTrafficEndpoint returns the IEs above CreatedTrafficEndpoint if the type of IE matches.
func (i *IE) CreatedTrafficEndpoint() (*CreatedTrafficEndpointFields, error) {
	if i.Type != CreatedTrafficEndpoint {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &CreatedTrafficEndpointFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseCreatedTrafficEndpointFields(i.Payload)
}

// CreatedTrafficEndpointFields is a set of fields in CreatedTrafficEndpoint IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type CreatedTrafficEndpointFields struct {
	TrafficEndpointID  uint8
	LocalFTEID         []*FTEIDFields
	UEIPAddress        []*UEIPAddressFields
	LocalIngressTunnel *IE
}

// ParseCreatedTrafficEndpointFields returns the IEs above CreatedTrafficEndpoint
func ParseCreatedTrafficEndpointFields(b []byte) (*CreatedTrafficEndpointFields, error) {

	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	traffic := &CreatedTrafficEndpointFields{}
	err = traffic.ParseIEs(ies...)
	return traffic, err
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (c *CreatedTrafficEndpointFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case TrafficEndpointID:
			id, err := ie.TrafficEndpointID()
			if err != nil {
				return err
			}
			c.TrafficEndpointID = id
		case FTEID:
			v, err := ie.FTEID()
			if err != nil {
				return err
			}
			c.LocalFTEID = append(c.LocalFTEID, v)
		case UEIPAddress:
			v, err := ie.UEIPAddress()
			if err != nil {
				return err
			}
			c.UEIPAddress = append(c.UEIPAddress, v)
		case LocalIngressTunnel:
			c.LocalIngressTunnel = ie
		}
	}
	return nil
}

// LocalFTEID returns FTEID that is found first in a grouped IE in structured format
// if the type of IE matches.
//
// This can only be used on the grouped IEs that may have multiple Local F-TEID IEs.
func (i *IE) LocalFTEID() (*FTEIDFields, error) {
	switch i.Type {
	case CreatedTrafficEndpoint:
		ies, err := i.CreatedTrafficEndpoint()
		if err != nil {
			return nil, err
		}
		for _, x := range ies.LocalFTEID {
			return x, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// LocalFTEIDN returns FTEID that is found Nth in a grouped IE in structured format
// if the type of IE matches.
//
// This can only be used on the grouped IEs that may have multiple Local F-TEID IEs.
func (i *IE) LocalFTEIDN(n int) (*FTEIDFields, error) {
	if n < 1 {
		return nil, ErrIENotFound
	}

	switch i.Type {
	case CreatedTrafficEndpoint: // has two F-TEID
		ies, err := i.CreatedTrafficEndpoint()
		if err != nil {
			return nil, err
		}
		c := 0
		for _, x := range ies.LocalFTEID {
			c++
			if c == n {
				return x, nil
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
