// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewUpdateURR Updates a new UpdateURR IE.
func NewUpdateURR(ies ...*IE) *IE {
	return newGroupedIE(UpdateURR, 0, ies...)
}

// UpdateURR returns the IEs above UpdateURR if the type of IE matches.
func (i *IE) UpdateURR() (*UpdateURRFields, error) {
	if i.Type != UpdateURR {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &UpdateURRFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseUpdateURRFields(i.Payload)
}

// UpdateURRFields is a set of fields in UpdateURR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateURRFields struct {
	// For Measurement Method and ReporingTrigger is more simpler to use direct the IE
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

// ParseUpdateURRFields returns the IEs above UpdateURR if the type of IE matches.
func ParseUpdateURRFields(b []byte) (*UpdateURRFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	u := &UpdateURRFields{}
	if err := u.ParseIEs(ies...); err != nil {
		return u, err
	}
	return u, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (u *UpdateURRFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case URRID:
			id, err := ie.URRID()
			if err != nil {
				return err
			}
			u.URRID = id
		case MeasurementMethod:
			u.MeasurementMethod = ie
		case ReportingTriggers:
			u.ReportingTrigger = ie
		case MeasurementPeriod:
			period, err := ie.MeasurementPeriod()
			if err != nil {
				return err
			}
			u.MeasurementPeriod = period
		case VolumeThreshold:
			volume, err := ie.VolumeThreshold()
			if err != nil {
				return err
			}
			u.VolumeThreshold = volume
		case VolumeQuota:
			volume, err := ie.VolumeQuota()
			if err != nil {
				return err
			}
			u.VolumeQuota = volume
		case EventThreshold:
			event, err := ie.EventThreshold()
			if err != nil {
				return err
			}
			u.EventThreshold = event
		case EventQuota:
			event, err := ie.EventQuota()
			if err != nil {
				return err
			}
			u.EventQuota = event
		case TimeThreshold:
			threshold, err := ie.TimeThreshold()
			if err != nil {
				return err
			}
			u.TimeThreshold = threshold
		case TimeQuota:
			quota, err := ie.TimeQuota()
			if err != nil {
				return err
			}
			u.TimeQuota = quota
		case QuotaHoldingTime:
			quota, err := ie.QuotaHoldingTime()
			if err != nil {
				return err
			}
			u.QuotaHoldingTime = quota
		case QuotaValidityTime:
			quota, err := ie.QuotaValidityTime()
			if err != nil {
				return err
			}
			u.QuotaValidityTime = quota
		case MonitoringTime:
			monitoringTime, err := ie.MonitoringTime()
			if err != nil {
				return err
			}
			u.MonitoringTime = monitoringTime
		case SubsequentVolumeThreshold:
			volume, err := ie.SubsequentVolumeThreshold()
			if err != nil {
				return err
			}
			u.SubsequentVolumeThreshold = volume
		case SubsequentTimeThreshold:
			duration, err := ie.SubsequentTimeThreshold()
			if err != nil {
				return err
			}
			u.SubsequentTimeThreshold = duration
		case SubsequentVolumeQuota:
			quota, err := ie.SubsequentVolumeQuota()
			if err != nil {
				return err
			}
			u.SubsequentVolumeQuota = quota
		case SubsequentTimeQuota:
			quota, err := ie.SubsequentTimeQuota()
			if err != nil {
				return err
			}
			u.SubsequentTimeQuota = quota
		case SubsequentEventThreshold:
			event, err := ie.SubsequentEventThreshold()
			if err != nil {
				return err
			}
			u.SubsequentEventThreshold = event
		case SubsequentEventQuota:
			event, err := ie.SubsequentEventQuota()
			if err != nil {
				return err
			}
			u.SubsequentEventQuota = event
		case InactivityDetectionTime:
			event, err := ie.InactivityDetectionTime()
			if err != nil {
				return err
			}
			u.InactivityDetectionTime = event
		case LinkedURRID:
			id, err := ie.LinkedURRID()
			if err != nil {
				return err
			}
			u.LinkedURRIDs = append(u.LinkedURRIDs, id)
		case MeasurementInformation:
			u.MeasurementInformation = ie
		case TimeQuotaMechanism:
			quota, err := ie.TimeQuotaMechanism()
			if err != nil {
				return err
			}
			u.TimeQuotaMechanism = quota
		case AggregatedURRs:
			id, err := ie.AggregatedURRs()
			if err != nil {
				return err
			}
			u.AggregatedURRs = append(u.AggregatedURRs, id)
		case FARID:
			id, err := ie.FARID()
			if err != nil {
				return err
			}
			u.FarIdForQuotaAction = id
		case EthernetInactivityTimer:
			timer, err := ie.EthernetInactivityTimer()
			if err != nil {
				return err
			}
			u.EthernetInactivityTimer = timer
		case AdditionalMonitoringTime:
			timer, err := ie.AdditionalMonitoringTime()
			if err != nil {
				return err
			}
			u.AdditionalMonitoringTime = timer
		case NumberOfReports:
			reports, err := ie.NumberOfReports()
			if err != nil {
				return err
			}
			u.NumberOfReports = reports
		case ApplicationID:
			applicationId, err := ie.ApplicationID()
			if err != nil {
				return err
			}
			u.ExempltedApplicationIdForQuotaAction = applicationId
		case SDFFilter:
			filter, err := ie.SDFFilter()
			if err != nil {
				return err
			}
			u.ExempltedSdfFilterForQuotaAction = append(u.ExempltedSdfFilterForQuotaAction, filter)
		case UserPlaneInactivityTimer:
			timer, err := ie.UserPlaneInactivityTimer()
			if err != nil {
				return err
			}
			u.UserPlaneInactivityTimer = timer
		}
	}
	return nil
}
