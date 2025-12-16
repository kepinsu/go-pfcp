// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPeerUPRestartReport creates a new PeerUPRestartReport IE.
func NewPeerUPRestartReport(peer *IE) *IE {
	return newGroupedIE(PeerUPRestartReport, 0, peer)
}

// PeerUPRestartReport returns the IEs above PeerUPRestartReport if the type of IE matches.
func (i *IE) PeerUPRestartReport() (*PeerUPRestartReportFields, error) {
	if i.Type != PeerUPRestartReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParsePeerUPRestartReportFields(i.Payload)
}

// PeerUPRestartReportFields is a set of fields in User Plane Path Failure Report IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type PeerUPRestartReportFields struct {
	RemoteGTPUPeer *RemoteGTPUPeerFields
}

// ParsePeerUPRestartReportFields returns the IEs above User Plane Path Failure Report
func ParsePeerUPRestartReportFields(b []byte) (*PeerUPRestartReportFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &PeerUPRestartReportFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *PeerUPRestartReportFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case RemoteGTPUPeer:
			v, err := ie.RemoteGTPUPeer()
			if err != nil {
				return err
			}
			far.RemoteGTPUPeer = v
		}
	}
	return nil
}
