// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewDroppedDLTrafficThreshold creates a new DroppedDLTrafficThreshold IE.
func NewDroppedDLTrafficThreshold(dlpa, dlby bool, packets, bytes uint64) *IE {
	fields := NewDroppedDLTrafficThresholdFields(dlpa, dlby, packets, bytes)
	b, err := fields.Marshal()
	if err != nil {
		return nil
	}
	return New(DroppedDLTrafficThreshold, b)
}

// DroppedDLTrafficThreshold returns DroppedDLTrafficThreshold in uint8 if the type of IE matches.
func (i *IE) DroppedDLTrafficThreshold() (*DroppedDLTrafficThresholdFields, error) {
	switch i.Type {
	case DroppedDLTrafficThreshold:
		return ParseDroppedDLTrafficThreshold(i.Payload)
	case CreateURR:
		ies, err := i.CreateURR()
		if err != nil {
			return nil, err
		}
		if ies.DroppedDLTrafficThreshold != nil {
			return ies.DroppedDLTrafficThreshold, nil
		}
		return nil, ErrIENotFound
	case UpdateURR:
		ies, err := i.UpdateURR()
		if err != nil {
			return nil, err
		}
		if ies.DroppedDLTrafficThreshold != nil {
			return ies.DroppedDLTrafficThreshold, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// DroppedDLTrafficThresholdFields is a set of fields in DroppedDLTrafficThreshold IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type DroppedDLTrafficThresholdFields struct {
	Flags   uint8
	Packets uint64
	Bytes   uint64
}

// NewDroppedDLTrafficThresholdFields creates a new NewDroppedDLTrafficThresholdFields.
func NewDroppedDLTrafficThresholdFields(dlpa, dlby bool, packets, bytes uint64) *DroppedDLTrafficThresholdFields {
	var flags uint8
	if dlpa {
		flags |= 0x01
	}
	if dlby {
		flags |= 0x02
	}
	return &DroppedDLTrafficThresholdFields{
		Flags:   flags,
		Packets: packets,
		Bytes:   bytes,
	}
}

// DroppedDLTrafficThreshold returns DroppedDLTrafficThreshold in uint8 if the type of IE matches.
func ParseDroppedDLTrafficThreshold(b []byte) (*DroppedDLTrafficThresholdFields, error) {
	d := &DroppedDLTrafficThresholdFields{}
	if err := d.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DroppedDLTrafficThresholdFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}
	d.Flags = b[0]

	offset := 1
	if d.HasDLPA() {
		if l < 1+offset {
			return io.ErrUnexpectedEOF
		}
		d.Packets = binary.BigEndian.Uint64(b[offset : offset+8])
		offset = +8
	}
	if d.HasDLBY() {
		if l < 1+offset {
			return io.ErrUnexpectedEOF
		}
		d.Bytes = binary.BigEndian.Uint64(b[offset : offset+8])
	}

	return nil
}

// HasDLBY reports whether an IE has DLBY bit.
func (d *DroppedDLTrafficThresholdFields) HasDLBY() bool {
	return has2ndBit(d.Flags)
}

// HasDLPA reports whether an IE has DLPA bit.
func (d *DroppedDLTrafficThresholdFields) HasDLPA() bool {
	return has1stBit(d.Flags)
}

// Marshal returns the serialized bytes of CTAGFields.
func (d *DroppedDLTrafficThresholdFields) Marshal() ([]byte, error) {
	b := make([]byte, d.MarshalLen())
	if err := d.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (d *DroppedDLTrafficThresholdFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}
	b[0] = d.Flags
	// Had Both
	if d.HasDLBY() && d.HasDLPA() {
		binary.BigEndian.PutUint64(b[1:9], d.Packets)
		binary.BigEndian.PutUint64(b[9:17], d.Bytes)
		// Has DLBY Only
	} else if d.HasDLBY() {
		binary.BigEndian.PutUint64(b[1:9], d.Bytes)
		// Has DLPA Only
	} else if d.HasDLPA() {
		binary.BigEndian.PutUint64(b[1:9], d.Packets)
	}
	return nil
}

// MarshalLen returns field length in integer.
func (d *DroppedDLTrafficThresholdFields) MarshalLen() int {
	if d.HasDLBY() && d.HasDLPA() {
		return 17
	} else if d.HasDLBY() {
		return 9
	} else if d.HasDLPA() {
		return 9
	}
	return 1
}
