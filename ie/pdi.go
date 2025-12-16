// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPDI creates a new PDI IE.
func NewPDI(ies ...*IE) *IE {
	return newGroupedIE(PDI, 0, ies...)
}

// PDI returns the IEs above PDI if the type of IE matches.
func (i *IE) PDI() (*PDIFields, error) {
	switch i.Type {
	case PDI:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &PDIFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParsePDIFields(i.Payload)
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return nil, err
		}
		if ies.PDI != nil {
			return ies.PDI, nil
		}
		return nil, ErrIENotFound
	case UpdatePDR:
		ies, err := i.UpdatePDR()
		if err != nil {
			return nil, err
		}
		if ies.PDI != nil {
			return ies.PDI, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// PDIFields is a set of fields in PDI IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type PDIFields struct {
	SourceInterface                 uint8
	LocalFTEID                      *FTEIDFields
	LocalIngressTunnel              *LocalIngressTunnelFields
	NetworkInstance                 string
	RedundantTransmissionParameters *RedundantTransmissionParametersField
	UEIPAddress                     []*UEIPAddressFields
	TrafficEndpointID               []uint8
	SDFFilter                       []*SDFFilterFields
	ApplicationID                   string
	EthernetPDUSessionInformation   uint8
	EthernetPacketFilter            *EthernetPacketFilterFields
	QFI                             []uint8
	FramedRoute                     []string
	FramedRouting                   uint32
	FramedIPv6Route                 string
	SourceInterfaceType             uint8
	IpMulticastAddressingInfo       []*IE
	DNSQueryFilter                  *IE
	MBSSession                      *IE
	AreaSessionID                   *IE
}

// ParsePDIFields returns the IEs above PDI
func ParsePDIFields(payload []byte) (*PDIFields, error) {

	// Parse all IES heres
	ies, err := ParseMultiIEs(payload)
	if err != nil {
		return nil, err
	}
	pdi := &PDIFields{}
	if err := pdi.ParseIEs(ies...); err != nil {
		return pdi, nil
	}
	return pdi, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (p *PDIFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case SourceInterface:
			source, err := ie.SourceInterface()
			if err != nil {
				return err
			}
			p.SourceInterface = source
		case FTEID:
			tfeid, err := ie.FTEID()
			if err != nil {
				return err
			}
			p.LocalFTEID = tfeid
		case LocalIngressTunnel:
			v, err := ie.LocalIngressTunnel()
			if err != nil {
				return err
			}
			p.LocalIngressTunnel = v
		case RedundantTransmissionParameters:
			ies, err := ParseRedundantTransmissionParametersField(ie.Payload)
			if err != nil {
				return err
			}
			p.RedundantTransmissionParameters = ies
		case NetworkInstance:
			network, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			p.NetworkInstance = network
		case UEIPAddress:
			ip, err := ie.UEIPAddress()
			if err != nil {
				return err
			}
			p.UEIPAddress = append(p.UEIPAddress, ip)
		case TrafficEndpointID:
			trafic, err := ie.TrafficEndpointID()
			if err != nil {
				return err
			}
			p.TrafficEndpointID = append(p.TrafficEndpointID, trafic)
		case SDFFilter:
			filter, err := ie.SDFFilter()
			if err != nil {
				return err
			}
			p.SDFFilter = append(p.SDFFilter, filter)
		case ApplicationID:
			appID, err := ie.ApplicationID()
			if err != nil {
				return err
			}
			p.ApplicationID = appID
		case EthernetPDUSessionInformation:
			ethernetPDU, err := ie.EthernetPDUSessionInformation()
			if err != nil {
				return err
			}
			p.EthernetPDUSessionInformation = ethernetPDU
		case EthernetPacketFilter:
			etherFilter, err := ie.EthernetPacketFilter()
			if err != nil {
				return err
			}
			p.EthernetPacketFilter = etherFilter
		case QFI:
			qfi, err := ie.QFI()
			if err != nil {
				return err
			}
			p.QFI = append(p.QFI, qfi)
		case FramedRoute:
			i, err := ie.FramedRoute()
			if err != nil {
				return err
			}
			p.FramedRoute = append(p.FramedRoute, i)
		case FramedRouting:
			i, err := ie.FramedRouting()
			if err != nil {
				return err
			}
			p.FramedRouting = i
		case FramedIPv6Route:
			i, err := ie.FramedIPv6Route()
			if err != nil {
				return err
			}
			p.FramedIPv6Route = i
		case TGPPInterfaceType:
			i, err := ie.TGPPInterfaceType()
			if err != nil {
				return err
			}
			p.SourceInterface = i

		case IPMulticastAddressingInfo:
			ies, err := ParseMultiIEs(ie.Payload)
			if err != nil {
				return err
			}
			p.IpMulticastAddressingInfo = ies
		}
	}
	return nil
}
