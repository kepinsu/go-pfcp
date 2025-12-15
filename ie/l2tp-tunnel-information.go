// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "net"

// NewL2TPTunnelInformation creates a new L2TPTunnelInformation IE.
func NewL2TPTunnelInformation(ies ...*IE) *IE {
	return newGroupedIE(L2TPTunnelInformation, 0, ies...)
}

// L2TPTunnelInformation returns the IEs above L2TPTunnelInformation if the type of IE matches.
func (i *IE) L2TPTunnelInformation() (*L2TPTunnelInformationFields, error) {
	if i.Type != L2TPTunnelInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseL2TPTunnelInformationFields(i.Payload)
}

// L2TPTunnelInformationFields represents a fields contained in L2TPTunnelInformation IE.
type L2TPTunnelInformationFields struct {
	LNSAddress       net.IP
	TunnelPassword   string
	TunnelPreference uint32
}

// ParseL2TPTunnelInformationFields parses b into L2TPTunnelInformationFields.
func ParseL2TPTunnelInformationFields(b []byte) (*L2TPTunnelInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &L2TPTunnelInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case LNSAddress:
			v, err := ie.LNSAddress()
			if err != nil {
				return p, err
			}
			p.LNSAddress = v
		case TunnelPassword:
			v, err := ie.TunnelPassword()
			if err != nil {
				return p, err
			}
			p.TunnelPassword = v
		case TunnelPreference:
			v, err := ie.TunnelPreference()
			if err != nil {
				return p, err
			}
			p.TunnelPreference = v
		}
	}
	return p, nil
}
