// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPacketRateStatusReport creates a new PacketRateStatusReport IE.
func NewPacketRateStatusReport(ies ...*IE) *IE {
	return newGroupedIE(PacketRateStatusReport, 0, ies...)
}

// NewPacketRateStatusReportWithinSessionModificationResponse creates a new PacketRateStatusReportWithinSessionModificationResponse IE.
func NewPacketRateStatusReportWithinSessionModificationResponse(ies ...*IE) *IE {
	return newGroupedIE(PacketRateStatusReportWithinSessionModificationResponse, 0, ies...)
}

// PacketRateStatusReport returns the IEs above PacketRateStatusReport if the type of IE matches.
func (i *IE) PacketRateStatusReport() ([]*IE, error) {
	switch i.Type {
	case PacketRateStatusReport, PacketRateStatusReportWithinSessionModificationResponse:
		return ParseMultiIEs(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// PacketRateStatusReportFields is a set of felds in PacketRateStatusReport IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type PacketRateStatusReportFields struct {
	QERID            uint32
	PacketRateStatus *PacketRateStatusFields
}

// ParsePacketRateStatusReportFields returns the IEs above PacketRateStatusReport IE
func ParsePacketRateStatusReportFields(b []byte) (*PacketRateStatusReportFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &PacketRateStatusReportFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case QERID:
			v, err := ie.QERID()
			if err != nil {
				return p, err
			}
			p.QERID = v
		case PacketRateStatus:
			v, err := ie.PacketRateStatus()
			if err != nil {
				return p, err
			}
			p.PacketRateStatus = v
		}
	}
	return p, nil
}
