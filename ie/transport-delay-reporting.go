// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTransportDelayReporting creates a new TransportDelayReporting IE.
func NewTransportDelayReporting(ies ...*IE) *IE {
	return newGroupedIE(TransportDelayReporting, 0, ies...)
}

// TransportDelayReporting returns the IEs above TransportDelayReporting if the type of IE matches.
func (i *IE) TransportDelayReporting() ([]*IE, error) {
	switch i.Type {
	case TransportDelayReporting:
		return ParseMultiIEs(i.Payload)
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return nil, err
		}
		if ies.TransportDelayReporting != nil {
			return ies.TransportDelayReporting.TransportDelayReporting()
		}
		return nil, ErrIENotFound
	case UpdatePDR:
		ies, err := i.UpdatePDR()
		if err != nil {
			return nil, err
		}
		if ies.TransportDelayReporting != nil {
			return ies.TransportDelayReporting.TransportDelayReporting()
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// TransportDelayReportingFields is a set of felds in TransportDelayReporting IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type TransportDelayReportingFields struct {
	PrecedingULGTPUPeer *RemoteGTPUPeerFields
	DSCP                uint16
}

// TransportDelayReportingFields returns the IEs above TransportDelayReporting IE
func ParseTransportDelayReportingFields(b []byte) (*TransportDelayReportingFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &TransportDelayReportingFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case RemoteGTPUPeer:
			v, err := ie.RemoteGTPUPeer()
			if err != nil {
				return p, err
			}
			p.PrecedingULGTPUPeer = v
		case TransportLevelMarking:
			v, err := ie.TransportLevelMarking()
			if err != nil {
				return p, err
			}
			p.DSCP = v
		}
	}
	return p, nil
}
