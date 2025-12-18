// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewUserPlanePathRecoveryReport creates a new UserPlanePathRecoveryReport IE.
func NewUserPlanePathRecoveryReport(peer *IE) *IE {
	return newGroupedIE(UserPlanePathRecoveryReport, 0, peer)
}

// UserPlanePathRecoveryReport returns the IEs above UserPlanePathRecoveryReport if the type of IE matches.
func (i *IE) UserPlanePathRecoveryReport() (*UserPlanePathRecoveryReportFields, error) {
	if i.Type != UserPlanePathRecoveryReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseUserPlanePathRecoveryReportFields(i.Payload)
}

// UserPlanePathRecoveryReportFields is a set of fields in User Plane Path Recovery Report IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type UserPlanePathRecoveryReportFields struct {
	RemoteGTPUPeer *RemoteGTPUPeerFields
}

// ParseUserPlanePathRecoveryReportFields returns the IEs above User Plane Path Recovery Report
func ParseUserPlanePathRecoveryReportFields(b []byte) (*UserPlanePathRecoveryReportFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &UserPlanePathRecoveryReportFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *UserPlanePathRecoveryReportFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == RemoteGTPUPeer {
			v, err := ie.RemoteGTPUPeer()
			if err != nil {
				return err
			}
			far.RemoteGTPUPeer = v
		}
	}
	return nil
}
