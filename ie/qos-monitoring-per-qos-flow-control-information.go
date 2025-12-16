// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewQoSMonitoringPerQoSFlowControlInformation creates a new QoSMonitoringPerQoSFlowControlInformation IE.
func NewQoSMonitoringPerQoSFlowControlInformation(ies ...*IE) *IE {
	return newGroupedIE(QoSMonitoringPerQoSFlowControlInformation, 0, ies...)
}

// QoSMonitoringPerQoSFlowControlInformation returns the IEs above QoSMonitoringPerQoSFlowControlInformation if the type of IE matches.
func (i *IE) QoSMonitoringPerQoSFlowControlInformation() (*QoSMonitoringPerQoSFlowControlInformationFields, error) {
	switch i.Type {
	case QoSMonitoringPerQoSFlowControlInformation:
		return ParseQoSMonitoringPerQoSFlowControlInformationFields(i.Payload)
	case CreateSRR:
		ies, err := i.CreateSRR()
		if err != nil {
			return nil, err
		}
		if ies.QoSMonitoringPerQoSFlowControlInformation != nil {
			return ies.QoSMonitoringPerQoSFlowControlInformation, nil
		}
		return nil, ErrIENotFound
	case UpdateSRR:
		ies, err := i.UpdateSRR()
		if err != nil {
			return nil, err
		}
		if ies.QoSMonitoringPerQoSFlowControlInformation != nil {
			return ies.QoSMonitoringPerQoSFlowControlInformation, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// QoSMonitoringPerQoSFlowControlInformationFields is a set of fields in QoSMonitoringPerQoSFlowControlInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type QoSMonitoringPerQoSFlowControlInformationFields struct {
	QFI                    uint8
	RequestedQoSMonitoring uint8
	ReportingFrequency     uint8
	PacketDelayThresholds  *PacketDelayThresholdsFields
	MinimumWaitTime        time.Duration
	MeasurementPeriod      time.Duration
}

// ParseQoSMonitoringPerQoSFlowControlInformationFields returns the IEs above QoSMonitoringPerQoSFlowControlInformation
func ParseQoSMonitoringPerQoSFlowControlInformationFields(b []byte) (*QoSMonitoringPerQoSFlowControlInformationFields, error) {

	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	bar := &QoSMonitoringPerQoSFlowControlInformationFields{}

	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case QFI:
			v, err := ie.QFI()
			if err != nil {
				return bar, err
			}
			bar.QFI = v
		case RequestedQoSMonitoring:
			v, err := ie.RequestedQoSMonitoring()
			if err != nil {
				return bar, err
			}
			bar.RequestedQoSMonitoring = v
		case ReportingFrequency:
			v, err := ie.ReportingFrequency()
			if err != nil {
				return bar, err
			}
			bar.ReportingFrequency = v
		case PacketDelayThresholds:
			v, err := ie.PacketDelayThresholds()
			if err != nil {
				return bar, err
			}
			bar.PacketDelayThresholds = v
		case MinimumWaitTime:
			v, err := ie.MinimumWaitTime()
			if err != nil {
				return bar, err
			}
			bar.MinimumWaitTime = v
		case MeasurementPeriod:
			v, err := ie.MeasurementPeriod()
			if err != nil {
				return bar, err
			}
			bar.MinimumWaitTime = v
		}
	}

	return bar, nil
}
