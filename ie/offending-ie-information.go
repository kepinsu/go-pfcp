package ie

// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

import (
	"encoding/binary"
	"io"
)

// NewOffendingIEInformation creates a new OffendingIEInformation IE.
func NewOffendingIEInformation(ie IEType, value []byte) *IE {
	fields := NewOffendingIEInformationFields(ie, value)
	b, err := fields.Marshal()
	if err != nil {
		return nil
	}
	return New(OffendingIEInformation, b)
}

// OffendingIEInformation returns OffendingIEInformation in *OffendingIEInformationFields if the type of IE matches.
func (i *IE) OffendingIEInformation() (*OffendingIEInformationFields, error) {
	switch i.Type {
	case OffendingIEInformation:
		fields, err := ParseOffendingIEInformationFields(i.Payload)
		if err != nil {
			return nil, err
		}

		return fields, nil
	case PartialFailureInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return nil, err
		}
		for _, ie := range ies {
			if ie == nil {
				continue
			}
			if ie.Type == OffendingIEInformation {
				return i.OffendingIEInformation()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// OffendingIEInformationFields represents a fields contained in OffendingIEInformation IE.
type OffendingIEInformationFields struct {
	IE    IEType
	Value []byte
}

// NewOffendingIEInformationFields creates a new OffendingIEInformationFields.
func NewOffendingIEInformationFields(ie IEType, value []byte) *OffendingIEInformationFields {
	f := &OffendingIEInformationFields{IE: ie, Value: value}
	return f
}

// ParseOffendingIEInformationFields parses b into OffendingIEInformationFields.
func ParseOffendingIEInformationFields(b []byte) (*OffendingIEInformationFields, error) {
	f := &OffendingIEInformationFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *OffendingIEInformationFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}
	f.IE = IEType(binary.BigEndian.Uint16(b[0:2]))
	if l > 2 {
		copy(f.Value, b[2:])
	}
	return nil
}

// Marshal returns the serialized bytes of OffendingIEInformationFields.
func (f *OffendingIEInformationFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *OffendingIEInformationFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}
	// Insert Type of offending IE
	binary.BigEndian.PutUint16(b[0:2], uint16(f.IE))
	// Insert the value of the offending IE
	copy(b[2:], f.Value)
	return nil
}

// MarshalLen returns field length in integer.
func (f *OffendingIEInformationFields) MarshalLen() int {
	// Length of Type of IE
	l := 2
	l += len(f.Value)
	return l
}
