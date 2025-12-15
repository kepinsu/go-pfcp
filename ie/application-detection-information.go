// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewApplicationDetectionInformation creates a new ApplicationDetectionInformation IE.
func NewApplicationDetectionInformation(ies ...*IE) *IE {
	return newGroupedIE(ApplicationDetectionInformation, 0, ies...)
}

// ApplicationDetectionInformation returns the IEs above ApplicationDetectionInformation if the type of IE matches.
func (i *IE) ApplicationDetectionInformation() (*ApplicationDetectionInformationFields, error) {
	switch i.Type {
	case ApplicationDetectionInformation:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &ApplicationDetectionInformationFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseApplicationDetectionInformationFields(i.Payload)
	case UsageReportWithinSessionModificationResponse,
		UsageReportWithinSessionDeletionResponse,
		UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return nil, err
		}
		if ies.ApplicationDetectionInformation != nil {
			return ies.ApplicationDetectionInformation, err
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// ApplicationDetectionInformationFields is a set of fields in Usage Report IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ApplicationDetectionInformationFields struct {
	ApplicationID         string
	ApplicationInstanceID string
	FlowInformation       []byte
	PDRID                 uint16
}

// ParseApplicationDetectionInformationFields returns the IEs above Update FAR
func ParseApplicationDetectionInformationFields(b []byte) (*ApplicationDetectionInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &ApplicationDetectionInformationFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (a *ApplicationDetectionInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case ApplicationID:
			v, err := ie.ApplicationID()
			if err != nil {
				return err
			}
			a.ApplicationID = v
		case ApplicationInstanceID:
			v, err := ie.ApplicationInstanceID()
			if err != nil {
				return err
			}
			a.ApplicationInstanceID = v
		case FlowInformation:
			v, err := ie.FlowInformation()
			if err != nil {
				return err
			}
			a.FlowInformation = v
		case PDRID:
			v, err := ie.PDRID()
			if err != nil {
				return err
			}
			a.PDRID = v
		}
	}
	return nil
}
