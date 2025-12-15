// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUpdateTrafficEndpoint Updates a new UpdateTrafficEndpoint IE.
func NewUpdateTrafficEndpoint(ies ...*IE) *IE {
	return newGroupedIE(UpdateTrafficEndpoint, 0, ies...)
}

// UpdateTrafficEndpoint returns the IEs above UpdateTrafficEndpoint if the type of IE matches.
func (i *IE) UpdateTrafficEndpoint() (*UpdateTrafficEndpointFields, error) {
	if i.Type != UpdateTrafficEndpoint {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &UpdateTrafficEndpointFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}

	// If the ChildIEs not already parsed
	return ParseUpdateTrafficEndpointFields(i.Payload)
}

// UpdateTrafficEndpointFields is a set of fields in UpdateTrafficEndpoint IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateTrafficEndpointFields struct {
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
	IpMulticastAddressingInfo                []*IE
	MBSSession                               *IE
	AreaSessionID                            *IE
	RatType                                  uint8
}

// ParseTrafficEndpointFields returns the IEs above UpdateTrafficEndpoint
func ParseUpdateTrafficEndpointFields(b []byte) (*UpdateTrafficEndpointFields, error) {

	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	traffic := &UpdateTrafficEndpointFields{}
	if err := traffic.ParseIEs(ies...); err != nil {
		return traffic, nil
	}
	return traffic, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (u *UpdateTrafficEndpointFields) ParseIEs(ies ...*IE) error {
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
			u.TrafficEndpointID = id
		case FTEID:
			v, err := ie.FTEID()
			if err != nil {
				return err
			}
			u.LocalFTEID = v
		case NetworkInstance:
			v, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			u.NetworkInstance = v
		case RedundantTransmissionParameters:
			v, err := ie.RedundantTransmissionParameters()
			if err != nil {
				return err
			}
			u.RedundantTransmissionDetectionParameters = v
		case UEIPAddress:
			v, err := ie.UEIPAddress()
			if err != nil {
				return err
			}
			u.UEIPAddress = append(u.UEIPAddress, v)
		case EthernetPDUSessionInformation:
			v, err := ie.EthernetPDUSessionInformation()
			if err != nil {
				return err
			}
			u.EthernetPDUSessionInformation = v
		case FramedRoute:
			v, err := ie.FramedRoute()
			if err != nil {
				return err
			}
			u.FramedRoute = append(u.FramedRoute, v)
		case FramedRouting:
			v, err := ie.FramedRouting()
			if err != nil {
				return err
			}
			u.FramedRouting = v
		case FramedIPv6Route:
			v, err := ie.FramedIPv6Route()
			if err != nil {
				return err
			}
			u.FramedIPv6Route = v
		case QFI:
			v, err := ie.QFI()
			if err != nil {
				return err
			}
			u.QFI = append(u.QFI, v)
		case TGPPInterfaceType:
			v, err := ie.TGPPInterfaceType()
			if err != nil {
				return err
			}
			u.SourceInterfaceType = v
		case IPMulticastAddress:
			u.IpMulticastAddressingInfo = append(u.IpMulticastAddressingInfo, ie)
		case MBSSessionIdentifier:
			u.MBSSession = ie
		case AreaSessionID:
			u.AreaSessionID = ie
		case RATType:
			v, err := ie.RATType()
			if err != nil {
				return err
			}
			u.RatType = v
		}
	}
	return nil
}
