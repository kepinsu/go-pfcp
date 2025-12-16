// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewLoadControlInformation creates a new LoadControlInformation IE.
func NewLoadControlInformation(ies ...*IE) *IE {
	return newGroupedIE(LoadControlInformation, 0, ies...)
}

// LoadControlInformation returns the IEs above LoadControlInformation if the type of IE matches.
func (i *IE) LoadControlInformation() (*LoadControlInformationFields, error) {
	if i.Type != LoadControlInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// Check if the ie.Parse have called or not
	if len(i.ChildIEs) > 0 {
		p := &LoadControlInformationFields{}
		if err := p.ParseIEs(i.ChildIEs...); err != nil {
			return p, err
		}
		return p, nil
	}
	// If the ChildIEs not already parsed
	return ParseLoadControlInformationFields(i.Payload)
}

// LoadControlInformationFields represents a fields contained in LoadControlInformation IE.
type LoadControlInformationFields struct {
	Sequence uint32
	Metric   uint8
}

// NewLoadControlInformationFields creates a new NewLoadControlInformationFields.
func NewLoadControlInformationFields(seq uint32, metrics uint8) *LoadControlInformationFields {
	return &LoadControlInformationFields{
		Sequence: seq,
		Metric:   metrics,
	}
}

// Marshal returns the serialized bytes of LoadControlInformationFields.
func (l *LoadControlInformationFields) Marshal() ([]byte, error) {
	seq := NewSequenceNumber(l.Sequence)
	s, err := seq.Marshal()
	if err != nil {
		return nil, err
	}
	metrics := NewMetric(l.Metric)
	m, err := metrics.Marshal()
	if err != nil {
		return nil, err
	}
	b := make([]byte, 0, len(s)+len(m))
	b = append(b, s...)
	b = append(b, m...)
	return b, nil
}

// NewLoadControlInformationFields creates a new NewLoadControlInformationFields.
func ParseLoadControlInformationFields(b []byte) (*LoadControlInformationFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	load := &LoadControlInformationFields{}
	if err := load.ParseIEs(ies...); err != nil {
		return load, err
	}
	return load, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (l *LoadControlInformationFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case SequenceNumber:
			v, err := ie.SequenceNumber()
			if err != nil {
				return err
			}
			l.Sequence = v
		case Metric:
			v, err := ie.Metric()
			if err != nil {
				return err
			}
			l.Metric = v
		}
	}
	return nil
}
