// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"

	"github.com/wmnsk/go-pfcp/internal/utils"
)

// NewGBR creates a new GBR IE.
func NewGBR(ul, dl uint64) *IE {
	i := New(GBR, make([]byte, 10))
	copy(i.Payload[0:5], utils.Uint64To40(ul))
	copy(i.Payload[5:10], utils.Uint64To40(dl))
	return i
}

// GBR returns GBR in []byte if the type of IE matches.
func (i *IE) GBR() (*GBRFields, error) {
	if len(i.Payload) < 10 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case GBR:
		return ParseGBRFields(i.Payload)
	case CreateQER:
		ies, err := i.CreateQER()
		if err != nil {
			return nil, err
		}
		if ies.GBR != nil {
			return ies.GBR, nil
		}
		return nil, ErrIENotFound
	case UpdateQER:
		ies, err := i.UpdateQER()
		if err != nil {
			return nil, err
		}
		if ies.GBR != nil {
			return ies.GBR, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// GBRFields is a set of fields in GBR IE.
type GBRFields struct {
	Uplink   uint64
	Downlink uint64
}

// NewBBRFields creates a new BBRFields.
func NewBBRFields(up, down uint64) *GBRFields {
	return &GBRFields{
		Uplink:   up,
		Downlink: down,
	}
}

// Marshal serializes BBRFields.
func (f *GBRFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo serializes BBRFields.
func (f *GBRFields) MarshalTo(b []byte) error {
	if len(b) < 10 {
		return io.ErrUnexpectedEOF
	}

	copy(b[0:5], utils.Uint64To40(f.Uplink))
	copy(b[5:10], utils.Uint64To40(f.Downlink))
	return nil
}

// ParseGBRFields decodes GBRFields.
func ParseGBRFields(b []byte) (*GBRFields, error) {
	f := &GBRFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalBinary decodes given bytes into BBRFields.
func (f *GBRFields) UnmarshalBinary(b []byte) error {
	if len(b) < 10 {
		return io.ErrUnexpectedEOF
	}
	f.Uplink = utils.Uint40To64(b[0:5])
	f.Downlink = utils.Uint40To64(b[5:10])
	return nil
}

// MarshalLen returns the serial length of BBRFields in int.
func (f *GBRFields) MarshalLen() int {
	return 10
}

// GBRUL returns GBRUL in uint64 if the type of IE matches.
func (i *IE) GBRUL() (uint64, error) {
	v, err := i.GBR()
	if err != nil {
		return 0, err
	}
	return v.Uplink, nil
}

// GBRDL returns GBRDL in uint64 if the type of IE matches.
func (i *IE) GBRDL() (uint64, error) {
	v, err := i.GBR()
	if err != nil {
		return 0, err
	}
	return v.Downlink, nil
}
