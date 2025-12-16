// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewLeaveIPMulticastInformationWithinUsageReport creates a new LeaveIPMulticastInformationWithinUsageReport IE.
func NewLeaveIPMulticastInformationWithinUsageReport(ies ...*IE) *IE {
	return newGroupedIE(LeaveIPMulticastInformationWithinUsageReport, 0, ies...)
}

// LeaveIPMulticastInformationWithinUsageReport returns the IEs above LeaveIPMulticastInformationWithinUsageReport if the type of IE matches.
func (i *IE) LeaveIPMulticastInformationWithinUsageReport() (*LeaveIPMulticastInformationFields, error) {
	switch i.Type {
	case LeaveIPMulticastInformationWithinUsageReport:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &LeaveIPMulticastInformationFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseLeaveIPMulticastInformationFields(i.Payload)
	case UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return nil, err
		}
		if ies.LeaveIPMulticastInformation != nil {
			return ies.LeaveIPMulticastInformation, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// LeaveIPMulticastInformationFields is a set of fields in LeaveIPMulticastInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type LeaveIPMulticastInformationFields struct {
	IPMulticastAddress *IPMulticastAddressFields
	SourceIPAddress    *SourceIPAddressFields
}

// ParseLeaveIPMulticastInformationFields returns the IEs above Update FAR
func ParseLeaveIPMulticastInformationFields(b []byte) (*LeaveIPMulticastInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &LeaveIPMulticastInformationFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (j *LeaveIPMulticastInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case IPMulticastAddress:
			v, err := ie.IPMulticastAddress()
			if err != nil {
				return err
			}
			j.IPMulticastAddress = v
		case SourceIPAddress:
			v, err := ie.SourceIPAddress()
			if err != nil {
				return err
			}
			j.SourceIPAddress = v
		}
	}
	return nil
}
