package ie

import (
	"encoding/binary"
	"io"
)

// NewThresholds creates a new Thresholds IE.
func NewThresholds(plr, rtt bool, rttvalue uint16, packetlossrate uint8) *IE {
	fields := NewThresholdsFields(plr, rtt, rttvalue, packetlossrate)

	b, err := fields.Marshal()
	if err != nil {
		return nil
	}

	return New(Thresholds, b)
}

// ThresholdsFields represents a fields contained in Thresholds IE.
type ThresholdsFields struct {
	Flags          uint8
	RRT            uint16
	PacketLossRate uint8
}

// NewThresholdsFields creates a new NewThresholdsFields.
func NewThresholdsFields(plr, rtt bool, rttvalue uint16, packetlossrate uint8) *ThresholdsFields {
	var flags uint8
	if plr {
		flags |= 0x02
	}
	if rtt {
		flags |= 0x01
	}

	return &ThresholdsFields{
		Flags:          flags,
		RRT:            rttvalue,
		PacketLossRate: packetlossrate,
	}
}

// HasRRT reports whether RRT flag is set.
func (f *ThresholdsFields) HasRRT() bool {
	return has1stBit(f.Flags)
}

// SetRRTFlag sets RRT flag in Thresholds.
func (f *ThresholdsFields) SetRRTFlag() {
	f.Flags |= 0x01
}

// HasPLR reports whether PLR flag is set.
func (f *ThresholdsFields) HasPLR() bool {
	return has2ndBit(f.Flags)
}

// SetPLRFlag sets PLR flag in Thresholds.
func (f *ThresholdsFields) SetPLRFlag() {
	f.Flags |= 0x02
}

// Thresholds returns the IEs above Thresholds if the type of IE matches.
func (i *IE) Thresholds() (*ThresholdsFields, error) {
	switch i.Type {
	case Thresholds:
		return ParseThresholdsFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// ParseThresholdsFields parses b into ThresholdsFields.
func ParseThresholdsFields(b []byte) (*ThresholdsFields, error) {
	f := &ThresholdsFields{}
	if err := f.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return f, nil
}

// UnmarshalBinary parses b into IE.
func (f *ThresholdsFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	f.Flags = b[0]
	offset := 1

	if f.HasRRT() {
		if l < offset+1 {
			return io.ErrUnexpectedEOF
		}
		f.RRT = binary.BigEndian.Uint16(b[offset : offset+1])
		offset += 2
	}

	if f.HasPLR() {
		if l < offset+4 {
			return io.ErrUnexpectedEOF
		}
		f.PacketLossRate = b[offset]
	}

	return nil
}

// Marshal returns the serialized bytes of ThresholdsFields.
func (f *ThresholdsFields) Marshal() ([]byte, error) {
	b := make([]byte, f.MarshalLen())
	if err := f.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (f *ThresholdsFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}

	b[0] = f.Flags
	offset := 1

	if f.HasRRT() {
		if l < offset+1 {
			return io.ErrUnexpectedEOF
		}
		binary.BigEndian.PutUint16(b[offset:offset+1], f.RRT)
		offset += 2
	}

	if f.HasPLR() {
		if l < offset+1 {
			return io.ErrUnexpectedEOF
		}
		b[offset] = f.PacketLossRate
	}
	return nil
}

// MarshalLen returns field length in integer.
func (f *ThresholdsFields) MarshalLen() int {
	l := 1

	if f.HasRRT() {
		l += 2
	}

	if f.HasPLR() {
		l += 1
	}
	return l
}
