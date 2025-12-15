// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "time"

// NewCreateBAR creates a new CreateBAR IE.
func NewCreateBAR(ies ...*IE) *IE {
	return newGroupedIE(CreateBAR, 0, ies...)
}

// CreateBAR returns the IEs above CreateBAR if the type of IE matches.
func (i *IE) CreateBAR() (*CreateBARFields, error) {
	switch i.Type {
	case CreateBAR:
		return ParseCreateBAR(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// CreateBARFields is a set of fields in CreateBAR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type CreateBARFields struct {
	BarID                          uint8
	DownlinkDataNotificationDelay  time.Duration
	SuggestedBufferingPacketsCount uint8
	MTEDTControlInformation        uint8
}

// ParseCreateBAR returns the IEs above CreateSRR
func ParseCreateBAR(b []byte) (*CreateBARFields, error) {

	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	bar := &CreateBARFields{}

	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case BARID:
			barid, err := ie.BARID()
			if err != nil {
				return bar, err
			}
			bar.BarID = barid
		case DownlinkDataNotificationDelay:
			v, err := ie.DownlinkDataNotificationDelay()
			if err != nil {
				return bar, err
			}
			bar.DownlinkDataNotificationDelay = v
		case SuggestedBufferingPacketsCount:
			v, err := ie.SuggestedBufferingPacketsCount()
			if err != nil {
				return bar, err
			}
			bar.SuggestedBufferingPacketsCount = v
		case MTEDTControlInformation:
			v, err := ie.MTEDTControlInformation()
			if err != nil {
				return bar, err
			}
			bar.MTEDTControlInformation = v
		}
	}

	return bar, nil
}
