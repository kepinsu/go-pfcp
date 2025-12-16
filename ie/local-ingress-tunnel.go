// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
	"net"
)

// NewLocalIngressTunnel creates a new LocalIngressTunnel IE.
func NewLocalIngressTunnel(flags uint8, port uint16, v4, v6 string) *IE {
	fields := NewLocalIngressTunnelFields(flags, port, v4, v6)
	b, err := fields.Marshal()
	if err != nil {
		return nil
	}

	return New(LocalIngressTunnel, b)
}

// LocalIngressTunnel returns LocalIngressTunnel in *LocalIngressTunnelFields if the type of IE matches.
func (i *IE) LocalIngressTunnel() (*LocalIngressTunnelFields, error) {
	switch i.Type {
	case LocalIngressTunnel:
		fields, err := ParseLocalIngressTunnelFields(i.Payload)
		if err != nil {
			return nil, err
		}

		return fields, nil
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return nil, err
		}
		if ies.PDI != nil && ies.PDI.LocalIngressTunnel != nil {
			return ies.PDI.LocalIngressTunnel, nil
		}

		return nil, ErrIENotFound
	case PDI:
		ies, err := i.PDI()
		if err != nil {
			return nil, err
		}
		if ies.LocalIngressTunnel != nil {
			return ies.LocalIngressTunnel, nil
		}
		return nil, ErrIENotFound
	case CreatedPDR:
		ies, err := i.CreatedPDR()
		if err != nil {
			return nil, err
		}
		if ies.LocalIngressTunnel != nil {
			return ies.LocalIngressTunnel, nil
		}
		return nil, ErrIENotFound
	case CreateTrafficEndpoint:
		ies, err := i.CreateTrafficEndpoint()
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == LocalIngressTunnel {
				return x.LocalIngressTunnel()
			}
		}
		return nil, ErrIENotFound
	case CreatedTrafficEndpoint:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == LocalIngressTunnel {
				return x.LocalIngressTunnel()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// LocalIngressTunnelFields represents a fields contained in LocalIngressTunnel IE.
type LocalIngressTunnelFields struct {
	Flags       uint8
	Port        uint16
	IPv4Address net.IP
	IPv6Address net.IP
}

// NewLocalIngressTunnelFields creates a new LocalIngressTunnelFields.
func NewLocalIngressTunnelFields(flags uint8, port uint16, v4, v6 string) *LocalIngressTunnelFields {
	f := &LocalIngressTunnelFields{Flags: flags, Port: port}

	if has2ndBit(flags) {
		f.IPv4Address = net.ParseIP(v4).To4()
	}
	if has1stBit(flags) {
		f.IPv6Address = net.ParseIP(v6).To16()
	}

	return f
}

// ParseLocalIngressTunnelFields parses b into LocalIngressTunnelFields.
func ParseLocalIngressTunnelFields(b []byte) (*LocalIngressTunnelFields, error) {
	f := &LocalIngressTunnelFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *LocalIngressTunnelFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 3 {
		return io.ErrUnexpectedEOF
	}

	f.Flags = b[0]
	f.Port = binary.BigEndian.Uint16(b[1:3])

	offset := 3
	if has2ndBit(f.Flags) {
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		f.IPv4Address = net.IP(b[offset : offset+4]).To4()
		offset += 4
	}

	if has1stBit(f.Flags) {
		if l < offset+16 {
			return io.ErrUnexpectedEOF
		}
		f.IPv6Address = net.IP(b[offset : offset+16]).To16()
		offset += 16
	}

	return nil
}

// Marshal returns the serialized bytes of LocalIngressTunnelFields.
func (f *LocalIngressTunnelFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *LocalIngressTunnelFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.Flags
	binary.BigEndian.PutUint16(b[1:3], f.Port)

	// See TS 29.244 clause 8.2.209
	// Only one of the V4 and V6 flags shall be set to "1".
	if has1stBit(f.Flags) && has2ndBit(f.Flags) {
		return &InvalidTypeError{Type: LocalIngressTunnel}
	}

	offset := 3
	if has2ndBit(f.Flags) {
		copy(b[offset:offset+4], f.IPv4Address)
		offset += 4
	}

	if has1stBit(f.Flags) {
		copy(b[offset:offset+16], f.IPv6Address)
		offset += 16
	}

	return nil
}

// MarshalLen returns field length in integer.
func (f *LocalIngressTunnelFields) MarshalLen() int {
	// 1 for the Flag and 2 for the port
	l := 3
	if has2ndBit(f.Flags) {
		l += 4
	}
	if has1stBit(f.Flags) {
		l += 16
	}

	return l
}
