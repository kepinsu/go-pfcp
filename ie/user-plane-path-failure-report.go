// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUserPlanePathFailureReport creates a new UserPlanePathFailureReport IE.
func NewUserPlanePathFailureReport(peer *IE) *IE {
	return newGroupedIE(UserPlanePathFailureReport, 0, peer)
}

// UserPlanePathFailureReport returns the IEs above UserPlanePathFailureReport if the type of IE matches.
func (i *IE) UserPlanePathFailureReport() (*UserPlanePathFailureReportFields, error) {
	if i.Type != UserPlanePathFailureReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseUserPlanePathFailureReportFields(i.Payload)
}

// UserPlanePathFailureReportFields is a set of fields in User Plane Path Failure Report IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UserPlanePathFailureReportFields struct {
	RemoteGTPUPeer []*RemoteGTPUPeerFields
}

// ParseUserPlanePathFailureReportFields returns the IEs above User Plane Path Failure Report
func ParseUserPlanePathFailureReportFields(b []byte) (*UserPlanePathFailureReportFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &UserPlanePathFailureReportFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *UserPlanePathFailureReportFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == RemoteGTPUPeer {
			v, err := ie.RemoteGTPUPeer()
			if err != nil {
				return err
			}
			far.RemoteGTPUPeer = append(far.RemoteGTPUPeer, v)
		}
	}
	return nil
}
