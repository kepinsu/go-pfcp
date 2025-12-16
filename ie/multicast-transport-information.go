// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"net"
)

// NewMulticastTransportInformation creates a new MulticastTransportInformation IE.
func NewMulticastTransportInformation(id uint32, ipv4, ipv6 string) *IE {
	fields := NewMulticastTransportInformationFields(id, ipv4, ipv6)

	b, err := fields.Marshal()
	if err != nil {
		return nil
	}

	return New(MulticastTransportInformation, b)
}

// MulticastTransportInformation returns MulticastTransportInformation in structured format if the type of IE matches.
func (i *IE) MulticastTransportInformation() (*MulticastTransportInformationFields, error) {
	switch i.Type {
	case MulticastTransportInformation:
		fields, err := ParseMulticastTransportInformationFields(i.Payload)
		if err != nil {
			return nil, err
		}

		return fields, nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MulticastTransportInformationFields represents a fields contained in MulticastTransportInformation IE.
type MulticastTransportInformationFields struct {
	CommonTunnelEndpointIdentifier uint32
	MulticastAddress               *MulticastAddress
}

// NewMulticastTransportInformationFields creates a new NewMulticastTransportInformationFields.
func NewMulticastTransportInformationFields(id uint32, ipv4, ipv6 string) *MulticastTransportInformationFields {
	f := &MulticastTransportInformationFields{
		CommonTunnelEndpointIdentifier: id,
	}
	if len(ipv4) > 0 {
		f.MulticastAddress = new(MulticastAddress)
		f.MulticastAddress.IPv4Address = net.ParseIP(ipv4).To4()
	}
	if len(ipv6) > 0 {
		if f.MulticastAddress == nil {
			f.MulticastAddress = new(MulticastAddress)
		}
		f.MulticastAddress.IPv6Address = net.ParseIP(ipv6).To16()
	}
	return f
}

// ParseMulticastTransportInformationFields parses b into MulticastTransportInformationFields.
func ParseMulticastTransportInformationFields(b []byte) (*MulticastTransportInformationFields, error) {
	f := &MulticastTransportInformationFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *MulticastTransportInformationFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 5 {
		return io.ErrUnexpectedEOF
	}
	f.CommonTunnelEndpointIdentifier = binary.BigEndian.Uint32(b[1:5])
	offset := 5
	// If the Multicast Transport Information don't had a IP adrress
	if offset == l {
		return nil
	}
	// Unmarshal the rest
	return f.MulticastAddress.UnmarshalBinary(b[offset:])
}

// Marshal returns the serialized bytes of MulticastTransportInformationFields.
func (f *MulticastTransportInformationFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *MulticastTransportInformationFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}
	// Spare Octet
	b[0] = 0x00
	binary.BigEndian.PutUint32(b[1:5], f.CommonTunnelEndpointIdentifier)
	offset := 5
	if f.MulticastAddress != nil {
		err := f.MulticastAddress.MarshalTo(b[offset : offset+f.MulticastAddress.MarshalLen()])
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalLen returns field length in integer.
func (f *MulticastTransportInformationFields) MarshalLen() int {
	// Spare + Commo Tunnel Endpoint Idenfier
	l := 5
	if f.MulticastAddress != nil {
		l += f.MulticastAddress.MarshalLen()
	}
	return l
}

// MulticastAddress represents a fields contained in MulticastTransportInformation and MBSSessionIdentifierFields IE.
type MulticastAddress struct {
	IPv4Address net.IP
	IPv6Address net.IP
}

// UnmarshalBinary parses b into IE.
// TODO: implement this Method
func (f *MulticastAddress) UnmarshalBinary(b []byte) error {
	// Ignore the Source IP
	if len(b) == 0 {
		return nil
	}
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	return nil
}

// Marshal returns the serialized bytes of MBSSessionIdentifierFields.
func (f *MulticastAddress) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *MulticastAddress) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}
	offset := 0
	if f.IPv4Address != nil {
		// Address Type 0 and Address Length 4 shall be used when Address is an IPv4 address.
		b[offset] = 0x04
		offset += 1
		copy(b[offset:offset+4], f.IPv4Address)
		offset += 4
	}
	if f.IPv6Address != nil {
		// Address Type 1 and Address Length 16 shall be used when Address is an IPv6 address.
		b[offset] = 0x80 | 0x10
		offset += 1
		copy(b[offset:offset+16], f.IPv6Address)
		offset += 16
	}
	return nil
}

// MarshalLen returns field length in integer.
func (f *MulticastAddress) MarshalLen() int {
	l := 0
	if f.IPv4Address != nil {
		// type + length + ip
		l += 5
	}
	if f.IPv6Address != nil {
		// type + length + ip
		l += 17
	}
	return l
}
