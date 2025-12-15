// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTransportLevelMarking creates a new TransportLevelMarking IE.
func NewTransportLevelMarking(tos uint16) *IE {
	return newUint16ValIE(TransportLevelMarking, tos)
}

// TransportLevelMarking returns TransportLevelMarking in uint16 if the type of IE matches.
func (i *IE) TransportLevelMarking() (uint16, error) {
	switch i.Type {
	case TransportLevelMarking:
		return i.ValueAsUint16()
	case ForwardingParameters:
		ies, err := i.ForwardingParameters()
		if err != nil {
			return 0, err
		}
		return ies.TransportLevelMarking, nil
	case UpdateForwardingParameters:
		ies, err := i.UpdateForwardingParameters()
		if err != nil {
			return 0, err
		}
		return ies.TransportLevelMarking, nil
	case DuplicatingParameters:
		ies, err := i.DuplicatingParameters()
		if err != nil {
			return 0, err
		}
		return ies.TransportLevelMarking, nil
	case UpdateDuplicatingParameters:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TransportLevelMarking {
				return x.TransportLevelMarking()
			}
		}
		return 0, ErrIENotFound
	case GTPUPathQoSControlInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TransportLevelMarking {
				return x.TransportLevelMarking()
			}
		}
		return 0, ErrIENotFound
	case GTPUPathQoSReport:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TransportLevelMarking {
				return x.TransportLevelMarking()
			}
		}
		return 0, ErrIENotFound
	case QoSInformationInGTPUPathQoSReport:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == TransportLevelMarking {
				return x.TransportLevelMarking()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}

}
