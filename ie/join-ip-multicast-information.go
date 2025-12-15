// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewJoinIPMulticastInformationWithinUsageReport creates a new JoinIPMulticastInformationWithinUsageReport IE.
func NewJoinIPMulticastInformationWithinUsageReport(ies ...*IE) *IE {
	return newGroupedIE(JoinIPMulticastInformationWithinUsageReport, 0, ies...)
}

// JoinIPMulticastInformationWithinUsageReport returns the IEs above JoinIPMulticastInformationWithinUsageReport if the type of IE matches.
func (i *IE) JoinIPMulticastInformationWithinUsageReport() (*JoinIPMulticastInformationFields, error) {
	switch i.Type {
	case JoinIPMulticastInformationWithinUsageReport:

		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			p := &JoinIPMulticastInformationFields{}
			if err := p.ParseIEs(i.ChildIEs...); err != nil {
				return p, err
			}
			return p, nil
		}
		// If the ChildIEs not already parsed
		return ParseJoinIPMulticastInformationFields(i.Payload)
	case UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return nil, err
		}
		if ies.JoinIPMulticastInformation != nil {
			return ies.JoinIPMulticastInformation, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// JoinIPMulticastInformationFields is a set of fields in JoinIPMulticastInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type JoinIPMulticastInformationFields struct {
	IPMulticastAddress *IPMulticastAddressFields
	SourceIPAddress    *SourceIPAddressFields
}

// ParseJoinIPMulticastInformationFields returns the IEs above Update FAR
func ParseJoinIPMulticastInformationFields(b []byte) (*JoinIPMulticastInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &JoinIPMulticastInformationFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (j *JoinIPMulticastInformationFields) ParseIEs(ies ...*IE) error {
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
