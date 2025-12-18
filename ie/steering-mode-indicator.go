// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewSteeringModeIndicator creates a new SteeringModeIndicator IE.
func NewSteeringModeIndicator(ueai, rtt bool) *IE {

	fields := NewSteeringModeIndicatorFields(ueai, rtt)
	b, err := fields.Marshal()
	if err != nil {
		return nil
	}

	return New(SteeringModeIndicator, b)
}

// SteeringModeIndicatorFields is a set of fields in SteeringModeIndicator IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type SteeringModeIndicatorFields struct {
	Flags uint8
}

// NewSteeringModeIndicatorFields creates a new NewSteeringModeIndicatorFields.
func NewSteeringModeIndicatorFields(ueai, rtt bool) *SteeringModeIndicatorFields {
	var flags uint8
	if ueai {
		flags |= 0x02
	}
	if rtt {
		flags |= 0x01
	}

	return &SteeringModeIndicatorFields{
		Flags: flags,
	}
}

// HasAlbi reports whether Albi flag is set.
func (f *SteeringModeIndicatorFields) HasAlbi() bool {
	return has1stBit(f.Flags)
}

// SetAlbiFlag sets Albi flag in Thresholds.
func (f *SteeringModeIndicatorFields) SetAlbiFlag() {
	f.Flags |= 0x01
}

// HasUEAI reports whether UEAI flag is set.
func (f *SteeringModeIndicatorFields) HasUEAI() bool {
	return has2ndBit(f.Flags)
}

// SetUEAIFlag sets UEAI flag in Thresholds.
func (f *SteeringModeIndicatorFields) SetUEAIFlag() {
	f.Flags |= 0x02
}

// SteeringModeIndicator returns the IEs above SteeringModeIndicator if the type of IE matches.
func (i *IE) SteeringModeIndicator() (*SteeringModeIndicatorFields, error) {
	switch i.Type {
	case SteeringModeIndicator:
		return ParseSteeringModeIndicatorFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// ParseSteeringModeIndicatorFields parses b into SteeringModeIndicatorFields.
func ParseSteeringModeIndicatorFields(b []byte) (*SteeringModeIndicatorFields, error) {
	f := &SteeringModeIndicatorFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *SteeringModeIndicatorFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	f.Flags = b[0]
	return nil
}

// Marshal returns the serialized bytes of SteeringModeIndicatorFields.
func (f *SteeringModeIndicatorFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *SteeringModeIndicatorFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.Flags
	return nil
}

// MarshalLen returns field length in integer.
func (f *SteeringModeIndicatorFields) MarshalLen() int {
	return 1
}
