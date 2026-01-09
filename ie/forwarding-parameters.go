// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewForwardingParameters creates a new ForwardingParameters IE.
func NewForwardingParameters(ies ...*IE) *IE {
	return newGroupedIE(ForwardingParameters, 0, ies...)
}

// ForwardingParameters returns the IEs above ForwardingParameters if the type of IE matches.
func (i *IE) ForwardingParameters() (*ForwardingParametersFields, error) {
	switch i.Type {
	case ForwardingParameters:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &ForwardingParametersFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		return ParseForwardingParametersFields(i.Payload)
	case CreateFAR:
		ies, err := i.CreateFAR()
		if err != nil {
			return nil, err
		}
		if ies.ForwardingParameters != nil {
			return ies.ForwardingParameters, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// ForwardingParametersFields is a set of fields in ForwardingParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ForwardingParametersFields struct {
	DestinationInterface        uint8
	NetworkInstance             string
	RedirectInformation         *RedirectInformationFields
	OuterHeaderCreation         *OuterHeaderCreationFields
	TransportLevelMarking       uint16
	ForwardingPolicy            []byte
	ForwardingPolicyIdentifier  string
	HeaderEnrichment            *HeaderEnrichmentFields
	LinkedTrafficEndpointID     uint8
	Proxying                    uint8
	DestinationInterfaceType    uint8
	DataNetworkAccessIdentifier string
}

// ParseForwardingParametersFields returns the IEs above ParseForwardingParameters.
func ParseForwardingParametersFields(b []byte) (*ForwardingParametersFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &ForwardingParametersFields{}
	if err := f.ParseIEs(ies...); err != nil {
		return f, nil
	}
	return f, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (f *ForwardingParametersFields) ParseIEs(ies ...*IE) error {
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
			f.DestinationInterface = dest
		case NetworkInstance:
			network, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			f.NetworkInstance = network
		case RedirectInformation:
			redirect, err := ie.RedirectInformation()
			if err != nil {
				return err
			}
			f.RedirectInformation = redirect
		case OuterHeaderCreation:
			creation, err := ie.OuterHeaderCreation()
			if err != nil {
				return err
			}
			f.OuterHeaderCreation = creation
		case TransportLevelMarking:
			transport, err := ie.TransportLevelMarking()
			if err != nil {
				return err
			}
			f.TransportLevelMarking = transport
		case ForwardingPolicy:
			policy, err := ie.ForwardingPolicy()
			if err != nil {
				return err
			}
			f.ForwardingPolicy = policy
			identifier, err := ie.ForwardingPolicyIdentifier()
			if err != nil {
				return err
			}
			f.ForwardingPolicyIdentifier = identifier
		case HeaderEnrichment:
			header, err := ie.HeaderEnrichment()
			if err != nil {
				return err
			}
			f.HeaderEnrichment = header
		case TrafficEndpointID:
			traficID, err := ie.TrafficEndpointID()
			if err != nil {
				return err
			}
			f.LinkedTrafficEndpointID = traficID
		case Proxying:
			proxying, err := ie.Proxying()
			if err != nil {
				return err
			}
			f.Proxying = proxying
		case TGPPInterfaceType:
			tgppinterface, err := ie.TGPPInterfaceType()
			if err != nil {
				return err
			}
			f.DestinationInterfaceType = tgppinterface
		case DataNetworkAccessIdentifier:
			v, err := ie.DataNetworkAccessIdentifier()
			if err != nil {
				return err
			}
			f.DataNetworkAccessIdentifier = v
		}
	}
	return nil
}
