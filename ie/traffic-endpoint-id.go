// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTrafficEndpointID creates a new TrafficEndpointID IE.
func NewTrafficEndpointID(id uint8) *IE {
	return newUint8ValIE(TrafficEndpointID, id)
}

// TrafficEndpointID returns TrafficEndpointID in uint8 if the type of IE matches.
func (i *IE) TrafficEndpointID() (uint8, error) {
	switch i.Type {
	case TrafficEndpointID:
		return i.ValueAsUint8()
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return 0, err
		}
		if ies.PDI != nil {
			if len(ies.PDI.TrafficEndpointID) > 0 {
				return ies.PDI.TrafficEndpointID[0], nil
			}
		}
		return 0, ErrIENotFound
	case PDI:
		ies, err := i.PDI()
		if err != nil {
			return 0, err
		}
		if len(ies.TrafficEndpointID) > 0 {
			return ies.TrafficEndpointID[0], nil
		}
		return 0, ErrIENotFound
	case ForwardingParameters:
		ies, err := i.ForwardingParameters()
		if err != nil {
			return 0, err
		}
		return ies.LinkedTrafficEndpointID, nil
	case UpdateForwardingParameters:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TrafficEndpointID {
				return x.TrafficEndpointID()
			}
		}
		return 0, ErrIENotFound
	case CreateTrafficEndpoint:
		ies, err := i.CreateTrafficEndpoint()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TrafficEndpointID {
				return x.TrafficEndpointID()
			}
		}
		return 0, ErrIENotFound
	case CreatedTrafficEndpoint:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TrafficEndpointID {
				return x.TrafficEndpointID()
			}
		}
		return 0, ErrIENotFound
	case UpdateTrafficEndpoint:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TrafficEndpointID {
				return x.TrafficEndpointID()
			}
		}
		return 0, ErrIENotFound
	case RemoveTrafficEndpoint:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TrafficEndpointID {
				return x.TrafficEndpointID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
