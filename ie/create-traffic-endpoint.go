// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewCreateTrafficEndpoint creates a new CreateTrafficEndpoint IE.
func NewCreateTrafficEndpoint(ies ...*IE) *IE {
	return newGroupedIE(CreateTrafficEndpoint, 0, ies...)
}

// CreateTrafficEndpoint returns the IEs above CreateTrafficEndpoint if the type of IE matches.
func (i *IE) CreateTrafficEndpoint() ([]*IE, error) {
	if i.Type != CreateTrafficEndpoint {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseMultiIEs(i.Payload)
}

// CreateTrafficEndpointFields is a set of fields in CreateTrafficEndpoint IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type CreateTrafficEndpointFields struct {
	TrafficEndpointID                        uint8
	LocalFTEID                               *FTEIDFields
	NetworkInstance                          string
	RedundantTransmissionDetectionParameters *RedundantTransmissionParametersField
	UEIPAddress                              []*UEIPAddressFields
	EthernetPDUSessionInformation            uint8
	FramedRoute                              []string
	FramedRouting                            uint32
	FramedIPv6Route                          string
	QFI                                      []uint8
	SourceInterfaceType                      uint8
	LocalIngressTunnel                       *IE
	IpMulticastAddressingInfo                []*IE
	MBSSession                               *IE
	AreaSessionID                            *IE
	RatType                                  uint8
}

// ParseTrafficEndpointFields returns the IEs above CreateTrafficEndpoint
func ParseCreateTrafficEndpointFields(b []byte) (*CreateTrafficEndpointFields, error) {

	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	traffic := &CreateTrafficEndpointFields{}
	err = traffic.ParseIEs(ies...)
	return traffic, err
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (c *CreateTrafficEndpointFields) ParseIEs(ies ...*IE) error {
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
			c.LocalFTEID = v
		case NetworkInstance:
			v, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			c.NetworkInstance = v
		case RedundantTransmissionParameters:
			v, err := ie.RedundantTransmissionParameters()
			if err != nil {
				return err
			}
			c.RedundantTransmissionDetectionParameters = v
		case UEIPAddress:
			v, err := ie.UEIPAddress()
			if err != nil {
				return err
			}
			c.UEIPAddress = append(c.UEIPAddress, v)
		case EthernetPDUSessionInformation:
			v, err := ie.EthernetPDUSessionInformation()
			if err != nil {
				return err
			}
			c.EthernetPDUSessionInformation = v
		case FramedRoute:
			v, err := ie.FramedRoute()
			if err != nil {
				return err
			}
			c.FramedRoute = append(c.FramedRoute, v)
		case FramedRouting:
			v, err := ie.FramedRouting()
			if err != nil {
				return err
			}
			c.FramedRouting = v
		case FramedIPv6Route:
			v, err := ie.FramedIPv6Route()
			if err != nil {
				return err
			}
			c.FramedIPv6Route = v
		case QFI:
			v, err := ie.QFI()
			if err != nil {
				return err
			}
			c.QFI = append(c.QFI, v)
		case TGPPInterfaceType:
			v, err := ie.TGPPInterfaceType()
			if err != nil {
				return err
			}
			c.SourceInterfaceType = v
		case LocalIngressTunnel:
			c.LocalIngressTunnel = ie
		case IPMulticastAddress:
			c.IpMulticastAddressingInfo = append(c.IpMulticastAddressingInfo, ie)
		case MBSSessionIdentifier:
			c.MBSSession = ie
		case AreaSessionID:
			c.AreaSessionID = ie
		case RATType:
			v, err := ie.RATType()
			if err != nil {
				return err
			}
			c.RatType = v
		}
	}
	return nil
}
