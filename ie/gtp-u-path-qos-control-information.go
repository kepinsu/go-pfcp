// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewGTPUPathQoSControlInformation creates a new GTPUPathQoSControlInformation IE.
func NewGTPUPathQoSControlInformation(ies ...*IE) *IE {
	return newGroupedIE(GTPUPathQoSControlInformation, 0, ies...)
}

// GTPUPathQoSControlInformation returns the IEs above GTPUPathQoSControlInformation if the type of IE matches.
func (i *IE) GTPUPathQoSControlInformation() (*GTPUPathQoSControlInformationFields, error) {
	if i.Type != GTPUPathQoSControlInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseGTPUPathQoSControlInformationFields(i.Payload)
}

// GTPUPathQoSControlInformationFields is a set of fields in GTPUPathQoSControlInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type GTPUPathQoSControlInformationFields struct {
	RemoteGTPUPeer              *RemoteGTPUPeerFields
	GTPUPathInterfaceType       uint8
	QoSReportTrigger            uint8
	DSCP                        uint16
	MeasurementPeriod           time.Duration
	AveragePacketDelayThreshold time.Duration
	MinimumPacketDelayThreshold time.Duration
	MaximumPacketDelayThreshold time.Duration
	MinimumWaitingTime          time.Duration
}

// ParseGTPUPathQoSControlInformationFields returns the IEs above GTPUPathQoSControlInformation
func ParseGTPUPathQoSControlInformationFields(b []byte) (*GTPUPathQoSControlInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &GTPUPathQoSControlInformationFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *GTPUPathQoSControlInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case RemoteGTPUPeer:
			v, err := ie.RemoteGTPUPeer()
			if err != nil {
				return err
			}
			far.RemoteGTPUPeer = v
		case GTPUPathInterfaceType:
			v, err := ie.GTPUPathInterfaceType()
			if err != nil {
				return err
			}
			far.GTPUPathInterfaceType = v
		case QoSReportTrigger:
			v, err := ie.QoSReportTrigger()
			if err != nil {
				return err
			}
			far.QoSReportTrigger = v
		case TransportLevelMarking:
			v, err := ie.TransportLevelMarking()
			if err != nil {
				return err
			}
			far.DSCP = v
		case MeasurementPeriod:
			v, err := ie.MeasurementPeriod()
			if err != nil {
				return err
			}
			far.MeasurementPeriod = v
		case AveragePacketDelay:
			v, err := ie.AveragePacketDelay()
			if err != nil {
				return err
			}
			far.AveragePacketDelayThreshold = v
		case MinimumPacketDelay:
			v, err := ie.MinimumPacketDelay()
			if err != nil {
				return err
			}
			far.MinimumPacketDelayThreshold = v
		case MaximumPacketDelay:
			v, err := ie.MaximumPacketDelay()
			if err != nil {
				return err
			}
			far.MaximumPacketDelayThreshold = v
		case Timer:
			v, err := ie.Timer()
			if err != nil {
				return err
			}
			far.MinimumWaitingTime = v
		}
	}
	return nil
}
