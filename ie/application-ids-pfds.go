// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewApplicationIDsPFDs creates a new ApplicationIDsPFDs IE.
func NewApplicationIDsPFDs(ies ...*IE) *IE {
	return newGroupedIE(ApplicationIDsPFDs, 0, ies...)
}

// ApplicationIDsPFDs returns the IEs above ApplicationIDsPFDs if the type of IE matches.
func (i *IE) ApplicationIDsPFDs() (*ApplicationIDsPFDsFields, error) {
	if i.Type != ApplicationIDsPFDs {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	// If Parse is called
	if len(i.ChildIEs) > 0 {
		f := &ApplicationIDsPFDsFields{}
		if err := f.ParseIEs(i.ChildIEs...); err != nil {
			return f, err
		}
		return f, nil
	}
	// Else parse with payload
	return ParseApplicationIDsPFDsFields(i.Payload)
}

// ApplicationIDsPFDsFields returns PFDContents in structured format if the type of IE matches.
//
// This IE has a complex payload that costs much when parsing.
type ApplicationIDsPFDsFields struct {
	ApplicationID string
	PFDContexts   []*PFDContextFields
}

// ParseApplicationIDsPFDsFields parses b into ApplicationIDsPFDsFields.
func ParseApplicationIDsPFDsFields(b []byte) (*ApplicationIDsPFDsFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &ApplicationIDsPFDsFields{}
	if err := f.ParseIEs(ies...); err != nil {
		return f, err
	}
	return f, nil
}

// ParseIEs parses ies into IE.
func (f *ApplicationIDsPFDsFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case ApplicationID:
			v, err := ie.ApplicationID()
			if err != nil {
				return err
			}
			f.ApplicationID = v
		case PFDContext:
			v, err := ie.PFDContext()
			if err != nil {
				return err
			}
			f.PFDContexts = append(f.PFDContexts, v)
		}
	}
	return nil
}
