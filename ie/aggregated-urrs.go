// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewAggregatedURRs creates a new AggregatedURRs IE.
func NewAggregatedURRs(ies ...*IE) *IE {
	return newGroupedIE(AggregatedURRs, 0, ies...)
}

// AggregatedURRs returns the IEs above AggregatedURRs if the type of IE matches.
func (i *IE) AggregatedURRs() (*AggregatedURRsField, error) {
	switch i.Type {
	case AggregatedURRs:
		return ParseAggregatedURRsField(i.Payload)
	case CreateURR:
		ies, err := i.CreateURR()
		if err != nil {
			return nil, err
		}
		if len(ies.AggregatedURRs) > 0 {
			return ies.AggregatedURRs[0], nil
		}
		return nil, ErrIENotFound
	case UpdateURR:
		ies, err := i.UpdateURR()
		if err != nil {
			return nil, err
		}
		if len(ies.AggregatedURRs) > 0 {
			return ies.AggregatedURRs[0], nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// AggregatedURRsField is a set of fields in AggregatedURRs IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type AggregatedURRsField struct {
	AggregatedURRID uint32
	Multiplier      *IE
}

func ParseAggregatedURRsField(b []byte) (*AggregatedURRsField, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	a := &AggregatedURRsField{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case AggregatedURRID:
			urrid, err := ie.AggregatedURRID()
			if err != nil {
				return a, err
			}
			a.AggregatedURRID = urrid
		case Multiplier:
			a.Multiplier = ie
		}
	}
	return a, nil
}
