// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewQoSInformationInGTPUPathQoSReport creates a new QoSInformationInGTPUPathQoSReport IE.
func NewQoSInformationInGTPUPathQoSReport(ies ...*IE) *IE {
	return newGroupedIE(QoSInformationInGTPUPathQoSReport, 0, ies...)
}

// QoSInformationInGTPUPathQoSReport returns the IEs above QoSInformationInGTPUPathQoSReport if the type of IE matches.
func (i *IE) QoSInformationInGTPUPathQoSReport() (*QoSInformationInGTPUPathQoSReportFields, error) {
	switch i.Type {
	case QoSInformationInGTPUPathQoSReport:
		return ParseQoSInformationInGTPUPathQoSReportFields(i.Payload)
	case GTPUPathQoSReport:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == QoSInformationInGTPUPathQoSReport {
				return x.QoSInformationInGTPUPathQoSReport()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// QoSInformationInGTPUPathQoSReportFields is a set of fields in QoSInformationInGTPUPathQoSReport IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type QoSInformationInGTPUPathQoSReportFields struct {
	AveragePacketDelayThreshold time.Duration
	MinimumPacketDelayThreshold time.Duration
	MaximumPacketDelayThreshold time.Duration
	DSCP                        uint16
}

// ParseQoSInformationInGTPUPathQoSReportFields returns the IEs above QoSInformationInGTPUPathQoSReport
func ParseQoSInformationInGTPUPathQoSReportFields(b []byte) (*QoSInformationInGTPUPathQoSReportFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &QoSInformationInGTPUPathQoSReportFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *QoSInformationInGTPUPathQoSReportFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case TransportLevelMarking:
			v, err := ie.TransportLevelMarking()
			if err != nil {
				return err
			}
			far.DSCP = v
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
		}
	}
	return nil
}
