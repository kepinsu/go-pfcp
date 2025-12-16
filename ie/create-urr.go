// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewCreateURR creates a new CreateURR IE.
func NewCreateURR(ies ...*IE) *IE {
	return newGroupedIE(CreateURR, 0, ies...)
}

// CreateURR returns the IEs above CreateURR if the type of IE matches.
func (i *IE) CreateURR() (*CreateURRFields, error) {
	if i.Type != CreateURR {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return ParseCreateURRFields(i.Payload)
}

// CreateURRFields is a set of fields in CreateURR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type CreateURRFields struct {
	// For Measurement Method and ReporingTrigger is more simplier to use direct the IE
	// structure

	URRID                                uint32
	MeasurementMethod                    *IE
	ReportingTrigger                     *IE
	MeasurementPeriod                    time.Duration
	VolumeThreshold                      *VolumeThresholdFields
	VolumeQuota                          *VolumeQuotaFields
	EventThreshold                       uint32
	EventQuota                           uint32
	TimeThreshold                        time.Duration
	TimeQuota                            time.Duration
	QuotaHoldingTime                     time.Duration
	DroppedDLTrafficThreshold            *DroppedDLTrafficThresholdFields
	QuotaValidityTime                    time.Duration
	MonitoringTime                       time.Time
	SubsequentVolumeThreshold            *SubsequentVolumeThresholdFields
	SubsequentTimeThreshold              time.Duration
	SubsequentVolumeQuota                *SubsequentVolumeQuotaFields
	SubsequentTimeQuota                  time.Duration
	SubsequentEventThreshold             uint32
	SubsequentEventQuota                 uint32
	InactivityDetectionTime              uint32
	LinkedURRIDs                         []uint32
	MeasurementInformation               *IE
	AggregatedURRs                       []*AggregatedURRsField
	TimeQuotaMechanism                   []byte
	FarIdForQuotaAction                  uint32
	EthernetInactivityTimer              time.Duration
	AdditionalMonitoringTime             *AdditionalMonitoringTimeFields
	NumberOfReports                      uint16
	ExempltedApplicationIdForQuotaAction string
	ExempltedSdfFilterForQuotaAction     []*SDFFilterFields
	UserPlaneInactivityTimer             time.Duration
}

// ParseCreateURRFields returns the IEs above CreateURR if the type of IE matches.
func ParseCreateURRFields(b []byte) (*CreateURRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	u := &CreateURRFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case URRID:
			id, err := ie.URRID()
			if err != nil {
				return u, err
			}
			u.URRID = id
		case MeasurementMethod:
			u.MeasurementMethod = ie
		case ReportingTriggers:
			u.ReportingTrigger = ie
		case MeasurementPeriod:
			period, err := ie.MeasurementPeriod()
			if err != nil {
				return u, err
			}
			u.MeasurementPeriod = period
		case VolumeThreshold:
			volume, err := ie.VolumeThreshold()
			if err != nil {
				return u, err
			}
			u.VolumeThreshold = volume
		case VolumeQuota:
			volume, err := ie.VolumeQuota()
			if err != nil {
				return u, err
			}
			u.VolumeQuota = volume
		case EventThreshold:
			event, err := ie.EventThreshold()
			if err != nil {
				return u, err
			}
			u.EventThreshold = event
		case EventQuota:
			event, err := ie.EventQuota()
			if err != nil {
				return u, err
			}
			u.EventQuota = event
		case TimeThreshold:
			threshold, err := ie.TimeThreshold()
			if err != nil {
				return u, err
			}
			u.TimeThreshold = threshold
		case TimeQuota:
			quota, err := ie.TimeQuota()
			if err != nil {
				return u, err
			}
			u.TimeQuota = quota
		case QuotaHoldingTime:
			quota, err := ie.QuotaHoldingTime()
			if err != nil {
				return u, err
			}
			u.QuotaHoldingTime = quota
		case QuotaValidityTime:
			quota, err := ie.QuotaValidityTime()
			if err != nil {
				return u, err
			}
			u.QuotaValidityTime = quota
		case MonitoringTime:
			monitoringTime, err := ie.MonitoringTime()
			if err != nil {
				return u, err
			}
			u.MonitoringTime = monitoringTime
		case SubsequentVolumeThreshold:
			volume, err := ie.SubsequentVolumeThreshold()
			if err != nil {
				return u, err
			}
			u.SubsequentVolumeThreshold = volume
		case SubsequentTimeThreshold:
			duration, err := ie.SubsequentTimeThreshold()
			if err != nil {
				return u, err
			}
			u.SubsequentTimeThreshold = duration
		case SubsequentVolumeQuota:
			quota, err := ie.SubsequentVolumeQuota()
			if err != nil {
				return u, err
			}
			u.SubsequentVolumeQuota = quota
		case SubsequentTimeQuota:
			quota, err := ie.SubsequentTimeQuota()
			if err != nil {
				return u, err
			}
			u.SubsequentTimeQuota = quota
		case SubsequentEventThreshold:
			event, err := ie.SubsequentEventThreshold()
			if err != nil {
				return u, err
			}
			u.SubsequentEventThreshold = event
		case SubsequentEventQuota:
			event, err := ie.SubsequentEventQuota()
			if err != nil {
				return u, err
			}
			u.SubsequentEventQuota = event
		case InactivityDetectionTime:
			event, err := ie.InactivityDetectionTime()
			if err != nil {
				return u, err
			}
			u.InactivityDetectionTime = event
		case LinkedURRID:
			id, err := ie.LinkedURRID()
			if err != nil {
				return u, err
			}
			u.LinkedURRIDs = append(u.LinkedURRIDs, id)
		case MeasurementInformation:
			u.MeasurementInformation = ie
		case TimeQuotaMechanism:
			quota, err := ie.TimeQuotaMechanism()
			if err != nil {
				return u, err
			}
			u.TimeQuotaMechanism = quota
		case AggregatedURRs:
			id, err := ie.AggregatedURRs()
			if err != nil {
				return u, err
			}
			u.AggregatedURRs = append(u.AggregatedURRs, id)
		case FARID:
			id, err := ie.FARID()
			if err != nil {
				return u, err
			}
			u.FarIdForQuotaAction = id
		case EthernetInactivityTimer:
			timer, err := ie.EthernetInactivityTimer()
			if err != nil {
				return u, err
			}
			u.EthernetInactivityTimer = timer
		case AdditionalMonitoringTime:
			timer, err := ie.AdditionalMonitoringTime()
			if err != nil {
				return u, err
			}
			u.AdditionalMonitoringTime = timer
		case NumberOfReports:
			reports, err := ie.NumberOfReports()
			if err != nil {
				return u, err
			}
			u.NumberOfReports = reports
		case ApplicationID:
			applicationId, err := ie.ApplicationID()
			if err != nil {
				return u, err
			}
			u.ExempltedApplicationIdForQuotaAction = applicationId
		case SDFFilter:
			filter, err := ie.SDFFilter()
			if err != nil {
				return u, err
			}
			u.ExempltedSdfFilterForQuotaAction = append(u.ExempltedSdfFilterForQuotaAction, filter)
		case UserPlaneInactivityTimer:
			timer, err := ie.UserPlaneInactivityTimer()
			if err != nil {
				return u, err
			}
			u.UserPlaneInactivityTimer = timer
		}
	}
	return u, nil
}
