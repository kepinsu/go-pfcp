// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTSCManagementInformation creates a new TSCManagementInformation IE.
func NewTSCManagementInformation(typ IEType, ies ...*IE) *IE {
	return newGroupedIE(typ, 0, ies...)
}

// NewTSCManagementInformationWithinSessionModificationRequest creates a new TSCManagementInformationWithinSessionModificationRequest IE.
func NewTSCManagementInformationWithinSessionModificationRequest(ies ...*IE) *IE {
	return newGroupedIE(TSCManagementInformationWithinSessionModificationRequest, 0, ies...)
}

// NewTSCManagementInformationWithinSessionModificationResponse creates a new TSCManagementInformationWithinSessionModificationResponse IE.
func NewTSCManagementInformationWithinSessionModificationResponse(ies ...*IE) *IE {
	return newGroupedIE(TSCManagementInformationWithinSessionModificationResponse, 0, ies...)
}

// NewTSCManagementInformationWithinSessionReportRequest creates a new TSCManagementInformationWithinSessionReportRequest IE.
func NewTSCManagementInformationWithinSessionReportRequest(ies ...*IE) *IE {
	return newGroupedIE(TSCManagementInformationWithinSessionReportRequest, 0, ies...)
}

// TSCManagementInformation returns the IEs above TSCManagementInformation if the type of IE matches.
func (i *IE) TSCManagementInformation() (*TSCManagementInformationFields, error) {
	switch i.Type {
	case TSCManagementInformationWithinSessionModificationRequest,
		TSCManagementInformationWithinSessionModificationResponse,
		TSCManagementInformationWithinSessionReportRequest:
		// Check if the ie.Parse have called or not
		if len(i.ChildIEs) > 0 {
			t := &TSCManagementInformationFields{}
			if err := t.ParseIEs(i.ChildIEs...); err != nil {
				return t, err
			}
		}
		// If the ChildIEs not already parsed
		return ParseTSCManagementInformationFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// TSCManagementInformationFields is a set of fields in TSCManagementInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type TSCManagementInformationFields struct {
	PortManagementInformationContainer         string
	UserPlanNodeManagementInformationContainer string
	NWTTPortNumber                             uint32
}

// ParseTSCManagementInformationFields returns the IEs above TSCManagementInformation.
func ParseTSCManagementInformationFields(b []byte) (*TSCManagementInformationFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &TSCManagementInformationFields{}
	if err := f.ParseIEs(ies...); err != nil {
		return f, err
	}
	return f, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (t *TSCManagementInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case PortManagementInformationContainer:
			dest, err := ie.PortManagementInformationContainer()
			if err != nil {
				return err
			}
			t.PortManagementInformationContainer = dest
		case BridgeManagementInformationContainer:
			v, err := ie.BridgeManagementInformationContainer()
			if err != nil {
				return err
			}
			t.UserPlanNodeManagementInformationContainer = v
		case NWTTPortNumber:
			v, err := ie.NWTTPortNumber()
			if err != nil {
				return err
			}
			t.NWTTPortNumber = v
		}
	}
	return nil
}
