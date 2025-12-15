// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewDSCPToPPIMappingInformation creates a new DSCPToPPIMappingInformation IE.
func NewDSCPToPPIMappingInformation(ppi uint8, dscp ...uint8) *IE {
	fields := NewDSCPToPPIMappingInformationFields(ppi, dscp...)
	v, err := fields.Marshal()
	if err != nil {
		return nil
	}
	return New(DSCPToPPIMappingInformation, v)
}

// DSCPToPPIMappingInformation returns the IEs above DSCPToPPIMappingInformation if the type of IE matches.
func (i *IE) DSCPToPPIMappingInformation() (*DSCPToPPIMappingInformationFields, error) {
	switch i.Type {
	case DSCPToPPIMappingInformation:
		return ParseDSCPToPPIMappingInformation(i.Payload)
	case DSCPToPPIControlInformation:
		ies, err := i.DSCPToPPIControlInformation()
		if err != nil {
			return nil, err
		}
		if len(ies.DSCPToPPIMappingInformation) > 0 {
			return ies.DSCPToPPIMappingInformation[0], nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

func NewDSCPToPPIMappingInformationFields(ppi uint8, dscp ...uint8) *DSCPToPPIMappingInformationFields {
	return &DSCPToPPIMappingInformationFields{PPIValue: ppi & 0x05, DSCP: dscp}
}

// DSCPToPPIMappingInformationFields represents a fields contained in DSCPToPPIMappingInformation IE.
type DSCPToPPIMappingInformationFields struct {
	PPIValue uint8
	DSCP     []uint8
}

// ParseDSCPToPPIMappingInformation parses b into DSCPToPPIMappingInformation.
func ParseDSCPToPPIMappingInformation(b []byte) (*DSCPToPPIMappingInformationFields, error) {
	f := &DSCPToPPIMappingInformationFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *DSCPToPPIMappingInformationFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}
	f.PPIValue = b[0] & 0x05
	if len(b[1:]) > 0 {
		f.DSCP = b[1:]
	}
	return nil
}

// Marshal returns the serialized bytes of DSCPToPPIMappingInformationFields.
func (f *DSCPToPPIMappingInformationFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *DSCPToPPIMappingInformationFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}
	b[0] = f.PPIValue
	copy(b[1:], f.DSCP)
	return nil
}

// MarshalLen returns field length in integer.
func (f *DSCPToPPIMappingInformationFields) MarshalLen() int {
	// PPI Value
	l := 1
	// DSCP Value
	l += len(f.DSCP)
	return l
}
