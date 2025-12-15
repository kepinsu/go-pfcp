// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewUpdateBAR Updates a new UpdateBAR IE.
func NewUpdateBAR(typ IEType, ies ...*IE) *IE {
	return newGroupedIE(typ, 0, ies...)
}

// NewUpdateBARWithinSessionModificationRequest Updates a new UpdateBARWithinSessionModificationRequest IE.
func NewUpdateBARWithinSessionModificationRequest(ies ...*IE) *IE {
	return NewUpdateBAR(UpdateBARWithinSessionModificationRequest, ies...)
}

// NewUpdateBARWithinSessionReportResponse Updates a new UpdateBARWithinSessionReportResponse IE.
func NewUpdateBARWithinSessionReportResponse(ies ...*IE) *IE {
	return NewUpdateBAR(UpdateBARWithinSessionReportResponse, ies...)
}

// UpdateBAR returns the IEs above UpdateBAR if the type of IE matches.
func (i *IE) UpdateBAR() (*UpdateBARFields, error) {
	switch i.Type {
	case UpdateBARWithinSessionModificationRequest,
		UpdateBARWithinSessionReportResponse:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &UpdateBARFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseUpdateBAR(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// UpdateBARFields is a set of fields in UpdateBAR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UpdateBARFields struct {
	BarID                          uint8
	DownlinkDataNotificationDelay  time.Duration
	SuggestedBufferingPacketsCount uint8
	MTEDTControlInformation        uint8
	DLBufferingDuration            time.Duration
}

// ParseUpdateBAR returns the IEs above UpdateSRR
func ParseUpdateBAR(b []byte) (*UpdateBARFields, error) {

	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	bar := &UpdateBARFields{}
	if err := bar.ParseIEs(ies...); err != nil {
		return bar, err
	}
	return bar, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (bar *UpdateBARFields) ParseIEs(ies ...*IE) error {

	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case BARID:
			barid, err := ie.BARID()
			if err != nil {
				return err
			}
			bar.BarID = barid
		case DownlinkDataNotificationDelay:
			v, err := ie.DownlinkDataNotificationDelay()
			if err != nil {
				return err
			}
			bar.DownlinkDataNotificationDelay = v
		case SuggestedBufferingPacketsCount:
			v, err := ie.SuggestedBufferingPacketsCount()
			if err != nil {
				return err
			}
			bar.SuggestedBufferingPacketsCount = v
		case MTEDTControlInformation:
			v, err := ie.MTEDTControlInformation()
			if err != nil {
				return err
			}
			bar.MTEDTControlInformation = v
		case DLBufferingDuration:
			v, err := ie.DLBufferingDuration()
			if err != nil {
				return err
			}
			bar.DLBufferingDuration = v
		}
	}
	return nil
}
