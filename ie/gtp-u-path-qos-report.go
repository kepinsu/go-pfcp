// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewGTPUPathQoSReport creates a new GTPUPathQoSReport IE.
func NewGTPUPathQoSReport(ies ...*IE) *IE {
	return newGroupedIE(GTPUPathQoSReport, 0, ies...)
}

// GTPUPathQoSReport returns the IEs above GTPUPathQoSReport if the type of IE matches.
func (i *IE) GTPUPathQoSReport() (*GTPUPathQoSReportFields, error) {
	if i.Type != GTPUPathQoSReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return ParseGTPUPathQoSReportFields(i.Payload)
}

// GTPUPathQoSReportFields is a set of fields in GTPUPathQoSReport IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type GTPUPathQoSReportFields struct {
	RemoteGTPUPeer        *RemoteGTPUPeerFields
	GTPUPathInterfaceType uint8
	QoSReportTrigger      uint8
	TimeStamp             time.Time
	StartTime             time.Time
	QoSInformation        *QoSInformationInGTPUPathQoSReportFields
}

// ParseGTPUPathQoSReportFields returns the IEs above GTPUPathQoSReport
func ParseGTPUPathQoSReportFields(b []byte) (*GTPUPathQoSReportFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &GTPUPathQoSReportFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *GTPUPathQoSReportFields) ParseIEs(ies ...*IE) error {
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
		case EventTimeStamp:
			v, err := ie.EventTimeStamp()
			if err != nil {
				return err
			}
			far.TimeStamp = v
		case StartTime:
			v, err := ie.StartTime()
			if err != nil {
				return err
			}
			far.StartTime = v
		case QoSInformationInGTPUPathQoSReport:
			v, err := ie.QoSInformationInGTPUPathQoSReport()
			if err != nil {
				return err
			}
			far.QoSInformation = v
		}
	}
	return nil
}
