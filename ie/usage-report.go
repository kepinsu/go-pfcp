// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewUsageReport creates a new UsageReport IE.
func NewUsageReport(typ IEType, ies ...*IE) *IE {
	return newGroupedIE(typ, 0, ies...)
}

// NewUsageReportWithinSessionModificationResponse creates a new UsageReportWithinSessionModificationResponse IE.
func NewUsageReportWithinSessionModificationResponse(ies ...*IE) *IE {
	return NewUsageReport(UsageReportWithinSessionModificationResponse, ies...)
}

// NewUsageReportWithinSessionDeletionResponse creates a new UsageReportWithinSessionDeletionResponse IE.
func NewUsageReportWithinSessionDeletionResponse(ies ...*IE) *IE {
	return NewUsageReport(UsageReportWithinSessionDeletionResponse, ies...)
}

// NewUsageReportWithinSessionReportRequest creates a new UsageReportWithinSessionReportRequest IE.
func NewUsageReportWithinSessionReportRequest(ies ...*IE) *IE {
	return NewUsageReport(UsageReportWithinSessionReportRequest, ies...)
}

// UsageReport returns the IEs above UsageReport if the type of IE matches.
func (i *IE) UsageReport() (*UsageReportFields, error) {
	switch i.Type {
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &UsageReportFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseUsageReportFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// UsageReportFields is a set of fields in Usage Report IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UsageReportFields struct {
	// Communs Fields for Session Report/Modification/Deletion

	URRID                      uint32
	URSEQN                     uint32
	UsageReportTrigger         []byte
	StartTime                  time.Time
	EndTime                    time.Time
	VolumeMeasurement          *VolumeMeasurementFields
	DurationMeasurement        time.Duration
	TimeOfFirstPacket          time.Time
	TimeOfLastPacket           time.Time
	UsageInformation           uint8
	QueryURRReference          uint32
	EthernetTrafficInformation *EthernetTrafficInformationFields
	// Only for Session Report
	ApplicationDetectionInformation *ApplicationDetectionInformationFields
	UEIPAddress                     *UEIPAddressFields
	NetworkInstance                 string
	EventTimeStamp                  time.Time
	JoinIPMulticastInformation      *JoinIPMulticastInformationFields
	LeaveIPMulticastInformation     *LeaveIPMulticastInformationFields
	PredefinedRulesName             string
}

// ParseUsageReportFields returns the IEs above Update FAR
func ParseUsageReportFields(b []byte) (*UsageReportFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &UsageReportFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *UsageReportFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case URRID:
			v, err := ie.URRID()
			if err != nil {
				return err
			}
			far.URRID = v
		case URSEQN:
			v, err := ie.URSEQN()
			if err != nil {
				return err
			}
			far.URSEQN = v
		case UsageReportTrigger:
			v, err := ie.UsageReportTrigger()
			if err != nil {
				return err
			}
			far.UsageReportTrigger = v
		case StartTime:
			v, err := ie.StartTime()
			if err != nil {
				return err
			}
			far.StartTime = v
		case EndTime:
			v, err := ie.EndTime()
			if err != nil {
				return err
			}
			far.EndTime = v
		case VolumeMeasurement:
			v, err := ie.VolumeMeasurement()
			if err != nil {
				return err
			}
			far.VolumeMeasurement = v
		case DurationMeasurement:
			v, err := ie.DurationMeasurement()
			if err != nil {
				return err
			}
			far.DurationMeasurement = v
		case TimeOfFirstPacket:
			v, err := ie.TimeOfFirstPacket()
			if err != nil {
				return err
			}
			far.TimeOfFirstPacket = v
		case TimeOfLastPacket:
			v, err := ie.TimeOfLastPacket()
			if err != nil {
				return err
			}
			far.TimeOfLastPacket = v
		case UsageInformation:
			v, err := ie.UsageInformation()
			if err != nil {
				return err
			}
			far.UsageInformation = v
		case QueryURRReference:
			v, err := ie.QueryURRReference()
			if err != nil {
				return err
			}
			far.QueryURRReference = v
		case EthernetTrafficInformation:
			v, err := ie.EthernetTrafficInformation()
			if err != nil {
				return err
			}
			far.EthernetTrafficInformation = v
		case ApplicationDetectionInformation:
			v, err := ie.ApplicationDetectionInformation()
			if err != nil {
				return err
			}
			far.ApplicationDetectionInformation = v
		case UEIPAddress:
			v, err := ie.UEIPAddress()
			if err != nil {
				return err
			}
			far.UEIPAddress = v
		case NetworkInstance:
			v, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			far.NetworkInstance = v
		case EventTimeStamp:
			v, err := ie.EventTimeStamp()
			if err != nil {
				return err
			}
			far.EventTimeStamp = v
		case JoinIPMulticastInformationWithinUsageReport:
			v, err := ie.JoinIPMulticastInformationWithinUsageReport()
			if err != nil {
				return err
			}
			far.JoinIPMulticastInformation = v
		case LeaveIPMulticastInformationWithinUsageReport:
			v, err := ie.LeaveIPMulticastInformationWithinUsageReport()
			if err != nil {
				return err
			}
			far.LeaveIPMulticastInformation = v
		case PredefinedRulesName:
			v, err := ie.PredefinedRulesName()
			if err != nil {
				return err
			}
			far.PredefinedRulesName = v
		}
	}
	return nil
}
