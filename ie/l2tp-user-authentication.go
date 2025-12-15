// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
	"net"
)

// NewL2TPUserAuthentication creates a new L2TPUserAuthentication IE.
// TODO: implement the Marshal and Unmarshal method
func NewL2TPUserAuthentication(flags uint8, v4, v6 string, v6d, v6pl uint8) *IE {
	fields := NewL2TPUserAuthenticationFields(flags, v4, v6, v6d, v6pl)
	b, err := fields.Marshal()
	if err != nil {
		return nil
	}

	return New(L2TPUserAuthentication, b)
}

// L2TPUserAuthentication returns L2TPUserAuthentication in *L2TPUserAuthenticationFields if the type of IE matches.
func (i *IE) L2TPUserAuthentication() (*L2TPUserAuthenticationFields, error) {
	switch i.Type {
	case L2TPUserAuthentication:
		fields, err := ParseL2TPUserAuthenticationFields(i.Payload)
		if err != nil {
			return nil, err
		}

		return fields, nil
	case L2TPSessionInformation:
		ies, err := i.L2TPSessionInformation()
		if err != nil {
			return nil, err
		}
		if ies.L2TPUserAuthentication != nil {
			return ies.L2TPUserAuthentication, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// L2TPUserAuthenticationFields represents a fields contained in L2TPUserAuthentication IE.
type L2TPUserAuthenticationFields struct {
	Flags                    uint8
	IPv4Address              net.IP
	IPv6Address              net.IP
	IPv6PrefixDelegationBits uint8
	IPv6PrefixLength         uint8
}

// NewL2TPUserAuthenticationFields creates a new L2TPUserAuthenticationFields.
func NewL2TPUserAuthenticationFields(flags uint8, v4, v6 string, v6d, v6pl uint8) *L2TPUserAuthenticationFields {
	f := &L2TPUserAuthenticationFields{Flags: flags}

	if has2ndBit(flags) && !has5thBit(flags) {
		f.IPv4Address = net.ParseIP(v4).To4()
	}

	if has1stBit(flags) && !has6thBit(flags) {
		f.IPv6Address = net.ParseIP(v6).To16()
	}

	if has4thBit(flags) {
		f.IPv6PrefixDelegationBits = v6d
	}

	if has7thBit(flags) {
		f.IPv6PrefixLength = v6pl
	}

	return f
}

// ParseL2TPUserAuthenticationFields parses b into L2TPUserAuthenticationFields.
func ParseL2TPUserAuthenticationFields(b []byte) (*L2TPUserAuthenticationFields, error) {
	f := &L2TPUserAuthenticationFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *L2TPUserAuthenticationFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	f.Flags = b[0]
	offset := 1

	if has2ndBit(f.Flags) && !has5thBit(f.Flags) {
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		f.IPv4Address = net.IP(b[offset : offset+4]).To4()
		offset += 4
	}

	if has1stBit(f.Flags) && !has6thBit(f.Flags) {
		if l < offset+16 {
			return io.ErrUnexpectedEOF
		}
		f.IPv6Address = net.IP(b[offset : offset+16]).To16()
		offset += 16
	}

	if has4thBit(f.Flags) {
		if l < offset+1 {
			return io.ErrUnexpectedEOF
		}
		f.IPv6PrefixDelegationBits = b[offset]
	}

	if has7thBit(f.Flags) {
		if l < offset+1 {
			return io.ErrUnexpectedEOF
		}
		f.IPv6PrefixLength = b[offset]
	}

	return nil
}

// Marshal returns the serialized bytes of L2TPUserAuthenticationFields.
func (f *L2TPUserAuthenticationFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *L2TPUserAuthenticationFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.Flags
	offset := 1

	if has2ndBit(f.Flags) && !has5thBit(f.Flags) {
		copy(b[offset:offset+4], f.IPv4Address)
		offset += 4
	}

	if has1stBit(f.Flags) && !has6thBit(f.Flags) {
		copy(b[offset:offset+16], f.IPv6Address)
		offset += 16
	}

	if has4thBit(f.Flags) {
		b[offset] = f.IPv6PrefixDelegationBits
		offset++
	}

	if has7thBit(f.Flags) {
		b[offset] = f.IPv6PrefixLength
	}

	return nil
}

// MarshalLen returns field length in integer.
func (f *L2TPUserAuthenticationFields) MarshalLen() int {
	l := 1
	if has2ndBit(f.Flags) && !has5thBit(f.Flags) {
		l += 4
	}
	if has1stBit(f.Flags) && !has6thBit(f.Flags) {
		l += 16
	}
	if has4thBit(f.Flags) {
		l++
	}
	if has7thBit(f.Flags) {
		l++
	}

	return l
}
