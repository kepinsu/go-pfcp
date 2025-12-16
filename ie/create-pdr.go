// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"time"
)

// NewCreatePDR creates a new CreatePDR IE.
func NewCreatePDR(ies ...*IE) *IE {
	return newGroupedIE(CreatePDR, 0, ies...)
}

// CreatePDR returns the IEs above CreatePDR if the type of IE matches.
func (i *IE) CreatePDR() (*CreatePDRFields, error) {
	if i.Type != CreatePDR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &CreatePDRFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseCreatePDRFields(i.Payload)
}

// CreatePDRFields is a set of fields in CreatePDR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type CreatePDRFields struct {
	// PDR ID
	ID uint16
	// Precedence
	Precedence uint32
	// PDI
	PDI                                             *PDIFields
	OuterHeaderRemoval                              []byte
	FARID                                           uint32
	URRID                                           []uint32
	QERID                                           []uint32
	ActivatePredefinedRules                         string
	ActivationTime                                  time.Time
	DeactivationTime                                time.Time
	MarID                                           uint16
	PacketReplicationAndDetectionCarryOnInformation uint8
	IPMulticastAddressingInfo                       []*IPMulticastAddressingInfoField
	// Max 2 times for IPv4 and IPv6
	UEIPAddressPoolIdentity     []*IE
	MPTCPApplicationIndiciation uint8
	TransportDelayReporting     *IE
	RatType                     uint8
}

// ParseCreatePDRFields returns the IEs above CreatePDR
func ParseCreatePDRFields(b []byte) (*CreatePDRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	createpdr := &CreatePDRFields{}
	if err := createpdr.ParseIEs(ies...); err != nil {
		return createpdr, nil
	}
	return createpdr, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (c *CreatePDRFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case PDRID:
			a, err := ie.PDRID()
			if err != nil {
				return err
			}
			c.ID = a
		case Precedence:
			a, err := ie.Precedence()
			if err != nil {
				return err
			}
			c.Precedence = a
		case PDI:
			a, err := ie.PDI()
			if err != nil {
				return err
			}
			c.PDI = a
		case OuterHeaderRemoval:
			a, err := ie.OuterHeaderRemoval()
			if err != nil {
				return err
			}
			c.OuterHeaderRemoval = a
		case FARID:
			a, err := ie.FARID()
			if err != nil {
				return err
			}
			c.FARID = a
		case URRID:
			a, err := ie.URRID()
			if err != nil {
				return err
			}
			c.URRID = append(c.URRID, a)
		case QERID:
			a, err := ie.QERID()
			if err != nil {
				return err
			}
			c.QERID = append(c.QERID, a)
		case ActivatePredefinedRules:
			a, err := ie.ActivatePredefinedRules()
			if err != nil {
				return err
			}
			c.ActivatePredefinedRules = a
		case ActivationTime:
			a, err := ie.ActivationTime()
			if err != nil {
				return err
			}
			c.ActivationTime = a
		case DeactivationTime:
			a, err := ie.DeactivationTime()
			if err != nil {
				return err
			}
			c.DeactivationTime = a
		case MARID:
			a, err := ie.MARID()
			if err != nil {
				return err
			}
			c.MarID = a
		case PacketReplicationAndDetectionCarryOnInformation:
			a, err := ie.PacketReplicationAndDetectionCarryOnInformation()
			if err != nil {
				return err
			}
			c.PacketReplicationAndDetectionCarryOnInformation = a
		case IPMulticastAddressingInfo:
			v, err := ie.IPMulticastAddressingInfo()
			if err != nil {
				return nil
			}
			c.IPMulticastAddressingInfo = append(c.IPMulticastAddressingInfo, v)
		case UEIPAddressPoolIdentity:
			c.UEIPAddressPoolIdentity = append(c.UEIPAddressPoolIdentity, ie)
		case MPTCPApplicableIndication:
			a, err := ie.MPTCPApplicableIndication()
			if err != nil {
				return err
			}
			c.MPTCPApplicationIndiciation = a
		case TransportDelayReporting:
			c.TransportDelayReporting = ie
		case RATType:
			a, err := ie.RATType()
			if err != nil {
				return err
			}
			c.RatType = a
		}
	}
	return nil
}
