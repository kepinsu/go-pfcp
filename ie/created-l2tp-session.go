// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "net"

// NewCreatedL2TPSession creates a new CreatedL2TPSession IE.
func NewCreatedL2TPSession(ies ...*IE) *IE {
	return newGroupedIE(CreatedL2TPSession, 0, ies...)
}

// CreatedL2TPSession returns the IEs above CreatedL2TPSession if the type of IE matches.
func (i *IE) CreatedL2TPSession() (*CreatedL2TPSessionFields, error) {
	if i.Type != CreatedL2TPSession {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseCreatedL2TPSession(i.Payload)
}

// CreatedL2TPSessionFields represents a fields contained in CreatedL2TPSession IE.
type CreatedL2TPSessionFields struct {
	DNSServerAddress  []net.IP
	NBNSServerAddress []net.IP
	LNSAddress        net.IP
}

// ParseCreatedL2TPSession parses b into CreatedL2TPSession.
func ParseCreatedL2TPSession(b []byte) (*CreatedL2TPSessionFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &CreatedL2TPSessionFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case DNSServerAddress:
			v, err := ie.DNSServerAddress()
			if err != nil {
				return p, err
			}
			p.DNSServerAddress = append(p.DNSServerAddress, v)
		case NBNSServerAddress:
			v, err := ie.NBNSServerAddress()
			if err != nil {
				return p, err
			}
			p.NBNSServerAddress = append(p.NBNSServerAddress, v)
		case LNSAddress:
			v, err := ie.LNSAddress()
			if err != nil {
				return p, err
			}
			p.LNSAddress = v
		}
	}
	return p, nil
}
