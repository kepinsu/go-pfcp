// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewQueryPacketRateStatusWithinSessionModificationRequest creates a new QueryPacketRateStatusWithinSessionModificationRequest IE.
func NewQueryPacketRateStatusWithinSessionModificationRequest(ies ...*IE) *IE {
	return newGroupedIE(QueryPacketRateStatusWithinSessionModificationRequest, 0, ies...)
}

// QueryPacketRateStatus returns the IEs above QueryPacketRateStatus if the type of IE matches.
func (i *IE) QueryPacketRateStatus() (*QueryPacketRateStatusFields, error) {
	if i.Type != QueryPacketRateStatusWithinSessionModificationRequest {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseQueryPacketRateStatusFields(i.Payload)
}

// QueryPacketRateStatusFields represents a fields contained in QueryPacketRateStatus IE.
type QueryPacketRateStatusFields struct {
	QERID uint32
}

// NewQueryPacketRateStatusFields creates a new QueryPacketRateStatusFields.
func NewQueryPacketRateStatusFields(qerid uint32) *QueryPacketRateStatusFields {
	f := &QueryPacketRateStatusFields{QERID: qerid}
	return f
}

// ParseQueryPacketRateStatusFields parses b into QueryPacketRateStatusFields.
func ParseQueryPacketRateStatusFields(b []byte) (*QueryPacketRateStatusFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &QueryPacketRateStatusFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == QERID {
			v, err := ie.QERID()
			if err != nil {
				return p, err
			}
			p.QERID = v
		}
	}
	return p, nil
}
