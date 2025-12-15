// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPFDContext creates a new PFDContext IE.
func NewPFDContext(contents ...*IE) *IE {
	return newGroupedIE(PFDContext, 0, contents...)
}

// PFDContext returns the IEs above PFDContext if the type of IE matches.
func (i *IE) PFDContext() (*PFDContextFields, error) {
	switch i.Type {
	case PFDContext:
		// If Parse is called
		if len(i.ChildIEs) > 0 {
			f := &PFDContextFields{}
			if err := f.ParseIEs(i.ChildIEs...); err != nil {
				return f, err
			}
			return f, nil
		}
		// Else parse with payload
		return ParsePFDContextFields(i.Payload)

	case ApplicationIDsPFDs:
		ies, err := i.ApplicationIDsPFDs()
		if err != nil {
			return nil, err
		}
		for _, i := range ies.PFDContexts {
			if i == nil {
				continue
			}
			return i, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// PFDContextFields returns PFDContents in structured format if the type of IE matches.
//
// This IE has a complex payload that costs much when parsing.
type PFDContextFields struct {
	PFDContents []*PFDContentsFields
}

// ParsePFDContextFields parses b into PFDContextFields.
func ParsePFDContextFields(b []byte) (*PFDContextFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &PFDContextFields{}
	if err := f.ParseIEs(ies...); err != nil {
		return f, err
	}
	return f, nil
}

// ParseIEs parses ies into IE.
func (f *PFDContextFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case PFDContents:
			v, err := ie.PFDContents()
			if err != nil {
				return err
			}
			f.PFDContents = append(f.PFDContents, v)
		}
	}
	return nil
}
