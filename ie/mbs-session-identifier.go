// Copyright go-pfcp authors. All rights reserved.
// Use of this TMGIce code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"

	"github.com/wmnsk/go-pfcp/internal/utils"
)

// NewMBSSessionIdentifier creates a new MBSSessionIdentifier IE.
func NewMBSSessionIdentifier(flags uint8, tmgi uint64, src *MulticastAddress, nid []byte) *IE {
	fields := NewMBSSessionIdentifierFields(flags, tmgi, src, nid)
	b, err := fields.Marshal()
	if err != nil {
		return nil
	}
	return New(MBSSessionIdentifier, b)
}

// MBSSessionIdentifier returns MBSSessionIdentifier in structured format if the type of IE matches.
func (i *IE) MBSSessionIdentifier() (*MBSSessionIdentifierFields, error) {
	switch i.Type {
	case MBSSessionIdentifier:
		return ParseMBSSessionIdentifierFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MBSSessionIdentifierFields represents a fields contained in MBSSessionIdentifier IE.
type MBSSessionIdentifierFields struct {
	Flags uint8
	// TMGI in 5 octets
	TMGI uint64

	// Source Specific IP Multicast Address
	Src *MulticastAddress

	// 11 digits
	NID []byte
}

// NewMBSSessionIdentifierFields creates a new NewMBSSessionIdentifierFields.
func NewMBSSessionIdentifierFields(flags uint8, tmgi uint64, src *MulticastAddress, nid []byte) *MBSSessionIdentifierFields {
	f := &MBSSessionIdentifierFields{
		Flags: flags,
		TMGI:  tmgi,
		Src:   src,
		NID:   nid,
	}

	return f
}

// HasNIDI reports whether NIDI flag is set.
func (f *MBSSessionIdentifierFields) HasNIDI() bool {
	return has3rdBit(f.Flags)
}

// SetNIDIFlag sets NIDI flag in MBSSessionIdentifier.
func (f *MBSSessionIdentifierFields) SetNIDIFlag() {
	f.Flags |= 0x04
}

// HasSSMI reports whether SSMI flag is set.
func (f *MBSSessionIdentifierFields) HasSSMI() bool {
	return has2ndBit(f.Flags)
}

// SetSSMIFlag sets SSMI flag in MBSSessionIdentifier.
func (f *MBSSessionIdentifierFields) SetSSMIFlag() {
	f.Flags |= 0x02
}

// HasTMGI reports whether TMGI flag is set.
func (f *MBSSessionIdentifierFields) HasTMGI() bool {
	return has1stBit(f.Flags)
}

// SetTMGIFlag sets TMGI flag in MBSSessionIdentifier.
func (f *MBSSessionIdentifierFields) SetTMGIFlag() {
	f.Flags |= 0x01
}

// ParseMBSSessionIdentifierFields parses b into MBSSessionIdentifierFields.
func ParseMBSSessionIdentifierFields(b []byte) (*MBSSessionIdentifierFields, error) {
	f := &MBSSessionIdentifierFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *MBSSessionIdentifierFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}

	f.Flags = b[0]
	offset := 1

	if f.HasTMGI() {
		if l < offset+5 {
			return io.ErrUnexpectedEOF
		}
		f.TMGI = utils.Uint40To64(b[offset : offset+5])
		offset += 5
	}

	if f.HasSSMI() {
		if l < offset+6 {
			return io.ErrUnexpectedEOF
		}
		offset += 6
	}

	if f.HasNIDI() {
		if l < offset+6 {
			return io.ErrUnexpectedEOF
		}
		copy(f.NID, b[offset:offset+6])
		offset += 6
	}

	return nil
}

// Marshal returns the serialized bytes of MBSSessionIdentifierFields.
func (f *MBSSessionIdentifierFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *MBSSessionIdentifierFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.Flags
	offset := 1

	if f.HasTMGI() {
		if l < offset+5 {
			return io.ErrUnexpectedEOF
		}
		copy(b[offset:offset+5], utils.Uint64To40(f.TMGI))
		offset += 5
	}

	if f.HasSSMI() && f.Src != nil {
		srcl := f.Src.MarshalLen()
		if l < offset+srcl {
			return io.ErrUnexpectedEOF
		}
		if err := f.Src.MarshalTo(b[offset : offset+srcl]); err != nil {
			return err
		}
		offset += srcl
	}

	if f.HasNIDI() {
		if l < offset+6 {
			return io.ErrUnexpectedEOF
		}
		copy(b[offset:offset+6], f.NID)
	}

	return nil
}

// MarshalLen returns field length in integer.
func (f *MBSSessionIdentifierFields) MarshalLen() int {
	l := 1
	if f.HasTMGI() {
		l += 5
	}
	if f.HasSSMI() && f.Src != nil {
		l += f.Src.MarshalLen()
	}
	if f.HasNIDI() {
		l += 6
	}
	return l
}
