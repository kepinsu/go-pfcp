// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewAdditionalMonitoringTime creates a new AdditionalMonitoringTime IE.
func NewAdditionalMonitoringTime(ies ...*IE) *IE {
	return newGroupedIE(AdditionalMonitoringTime, 0, ies...)
}

// AdditionalMonitoringTime returns the IEs above AdditionalMonitoringTime if the type of IE matches.
func (i *IE) AdditionalMonitoringTime() (*AdditionalMonitoringTimeFields, error) {
	switch i.Type {
	case AdditionalMonitoringTime:
		return ParseAdditionalMonitoringTimeFields(i.Payload)
	case CreateURR:
		ies, err := i.CreateURR()
		if err != nil {
			return nil, err
		}
		if ies.AdditionalMonitoringTime != nil {
			return ies.AdditionalMonitoringTime, nil
		}
		return nil, ErrIENotFound
	case UpdateURR:
		ies, err := i.UpdateURR()
		if err != nil {
			return nil, err
		}
		if ies.AdditionalMonitoringTime != nil {
			return ies.AdditionalMonitoringTime, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// AdditionalMonitoringTimeFields is a set of felds in AdditionalMonitoringTime IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type AdditionalMonitoringTimeFields struct {
	MonitoringTime            time.Time
	SubsequentVolumeThreshold *SubsequentVolumeThresholdFields
	SubsequentTimeThreshold   time.Duration
	SubsequentVolumeQuota     *SubsequentVolumeQuotaFields
	SubsequentTimeQuota       time.Duration
	SubsequentEventThreshold  uint32
	SubsequentEventQuota      uint32
}

func ParseAdditionalMonitoringTimeFields(b []byte) (*AdditionalMonitoringTimeFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	a := &AdditionalMonitoringTimeFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case MonitoringTime:
			monitoringTime, err := ie.MonitoringTime()
			if err != nil {
				return a, err
			}
			a.MonitoringTime = monitoringTime
		case SubsequentVolumeThreshold:
			volume, err := ie.SubsequentVolumeThreshold()
			if err != nil {
				return a, err
			}
			a.SubsequentVolumeThreshold = volume
		case SubsequentTimeThreshold:
			duration, err := ie.SubsequentTimeThreshold()
			if err != nil {
				return a, err
			}
			a.SubsequentTimeThreshold = duration
		case SubsequentVolumeQuota:
			quota, err := ie.SubsequentVolumeQuota()
			if err != nil {
				return a, err
			}
			a.SubsequentVolumeQuota = quota
		case SubsequentTimeQuota:
			quota, err := ie.SubsequentTimeQuota()
			if err != nil {
				return a, err
			}
			a.SubsequentTimeQuota = quota
		case SubsequentEventThreshold:
			event, err := ie.SubsequentEventThreshold()
			if err != nil {
				return a, err
			}
			a.SubsequentEventThreshold = event
		case SubsequentEventQuota:
			event, err := ie.SubsequentEventQuota()
			if err != nil {
				return a, err
			}
			a.SubsequentEventQuota = event
		}
	}
	return a, nil
}
