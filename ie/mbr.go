// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"

	"github.com/wmnsk/go-pfcp/internal/utils"
)

// NewMBR creates a new MBR IE.
func NewMBR(ul, dl uint64) *IE {
	mbr := NewMBRFields(ul, dl)
	payload, err := mbr.Marshal()
	if err != nil {
		return nil
	}
	return New(MBR, payload)
}

// MBR returns MBR in []byte if the type of IE matches.
func (i *IE) MBR() (*MBRFields, error) {
	if len(i.Payload) < 10 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case MBR:
		return ParseMBRFields(i.Payload)
	case CreateQER:
		ies, err := i.CreateQER()
		if err != nil {
			return nil, err
		}
		if ies.MBR != nil {
			return ies.MBR, nil
		}
		return nil, ErrIENotFound
	case UpdateQER:
		ies, err := i.UpdateQER()
		if err != nil {
			return nil, err
		}
		if ies.MBR != nil {
			return ies.MBR, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MBRFields is a set of fields in MBR IE.
type MBRFields struct {
	Uplink   uint64
	Downlink uint64
}

// NewMBRFields creates a new MBRFields.
func NewMBRFields(up, down uint64) *MBRFields {
	return &MBRFields{
		Uplink:   up,
		Downlink: down,
	}
}

// Marshal serializes MBRFields.
func (f *MBRFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes MBRFields.
func (f *MBRFields) MarshalTo(b []byte) error {
	if len(b) < 10 {
		return io.ErrUnexpectedEOF
	}

	copy(b[0:5], utils.Uint64To40(f.Uplink))
	copy(b[5:10], utils.Uint64To40(f.Downlink))
	return nil
}

// ParseMBRFields decodes MBRFields.
func ParseMBRFields(b []byte) (*MBRFields, error) {
	f := &MBRFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary decodes given bytes into MBRFields.
func (f *MBRFields) UnmarshalBinary(b []byte) error {
	if len(b) < 10 {
		return io.ErrUnexpectedEOF
	}
	f.Uplink = utils.Uint40To64(b[0:5])
	f.Downlink = utils.Uint40To64(b[5:10])
	return nil
}

// MarshalLen returns the serial length of MBRFields in int.
func (f *MBRFields) MarshalLen() int {
	return 10
}

// MBRUL returns MBRUL in uint64 if the type of IE matches.
func (i *IE) MBRUL() (uint64, error) {
	v, err := i.MBR()
	if err != nil {
		return 0, err
	}
	return v.Uplink, nil
}

// MBRDL returns MBRDL in uint64 if the type of IE matches.
func (i *IE) MBRDL() (uint64, error) {
	v, err := i.MBR()
	if err != nil {
		return 0, err
	}
	return v.Downlink, nil
}
