// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateForwardingParameters creates a new UpdateForwardingParameters IE.
func NewUpdateForwardingParameters(ies ...*IE) *IE {
	return newGroupedIE(UpdateForwardingParameters, 0, ies...)
}

// UpdateForwardingParameters returns the IEs above UpdateForwardingParameters if the type of IE matches.
func (i *IE) UpdateForwardingParameters() (*UpdateForwardingParametersFields, error) {
	switch i.Type {
	case UpdateForwardingParameters:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &UpdateForwardingParametersFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseUpdateForwardingParametersFields(i.Payload)
	case UpdateFAR:
		ies, err := i.UpdateFAR()
		if err != nil {
			return nil, err
		}
		if ies.UpdateForwardingParameters != nil {
			return ies.UpdateForwardingParameters, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// UpdateForwardingParametersFields is a set of fields in UpdateeForwardingParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateForwardingParametersFields struct {
	DestinationInterface        uint8
	NetworkInstance             string
	RedirectInformation         *RedirectInformationFields
	OuterHeaderCreation         *OuterHeaderCreationFields
	TransportLevelMarking       uint16
	ForwardingPolicy            []byte
	ForwardingPolicyIdentifier  string
	HeaderEnrichment            *HeaderEnrichmentFields
	PFCPSMReqFlags              uint8
	LinkedTrafficEndpointID     uint8
	DestinationInterfaceType    uint8
	DataNetworkAccessIdentifier string
}

// ParseUpdateForwardingParametersFields returns the IEs above Update Forwarding Parameters
func ParseUpdateForwardingParametersFields(b []byte) (*UpdateForwardingParametersFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	u := &UpdateForwardingParametersFields{}
	if err := u.ParseIEs(ies...); err != nil {
		return u, err
	}
	return u, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (u *UpdateForwardingParametersFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case DestinationInterface:
			dest, err := ie.DestinationInterface()
			if err != nil {
				return err
			}
			u.DestinationInterface = dest
		case NetworkInstance:
			network, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			u.NetworkInstance = network
		case RedirectInformation:
			redirect, err := ie.RedirectInformation()
			if err != nil {
				return err
			}
			u.RedirectInformation = redirect
		case OuterHeaderCreation:
			creation, err := ie.OuterHeaderCreation()
			if err != nil {
				return err
			}
			u.OuterHeaderCreation = creation
		case TransportLevelMarking:
			transport, err := ie.TransportLevelMarking()
			if err != nil {
				return err
			}
			u.TransportLevelMarking = transport
		case ForwardingPolicy:
			policy, err := ie.ForwardingPolicy()
			if err != nil {
				return err
			}
			u.ForwardingPolicy = policy
			identifier, err := ie.ForwardingPolicyIdentifier()
			if err != nil {
				return err
			}
			u.ForwardingPolicyIdentifier = identifier
		case PFCPSMReqFlags:
			v, err := ie.PFCPSMReqFlags()
			if err != nil {
				return err
			}
			u.PFCPSMReqFlags = v
		case HeaderEnrichment:
			header, err := ie.HeaderEnrichment()
			if err != nil {
				return err
			}
			u.HeaderEnrichment = header
		case TrafficEndpointID:
			traficID, err := ie.TrafficEndpointID()
			if err != nil {
				return err
			}
			u.LinkedTrafficEndpointID = traficID
		case TGPPInterfaceType:
			tgppinterface, err := ie.TGPPInterfaceType()
			if err != nil {
				return err
			}
			u.DestinationInterfaceType = tgppinterface
		case DataNetworkAccessIdentifier:
			v, err := ie.DataNetworkAccessIdentifier()
			if err != nil {
				return err
			}
			u.DataNetworkAccessIdentifier = v
		}
	}
	return nil
}
