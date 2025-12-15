// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRemoveQER creates a new RemoveQER IE.
func NewRemoveQER(qer *IE) *IE {
	return newGroupedIE(RemoveQER, 0, qer)
}

// RemoveQER returns the IEs above RemoveQER if the type of IE matches.
func (i *IE) RemoveQER() (*RemoveQERFields, error) {
	if i.Type != RemoveQER {
		return nil, &InvalidTypeError{Type: i.Type}
	}
	return ParseRemoveQERFields(i.Payload)
}

// RemoveQERFields is a set of fields in Remove QER IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RemoveQERFields struct {
	QERID uint32
}

// ParseRemoveQERFields returns the IEs above RemoveMBSUnicastParameters.
func ParseRemoveQERFields(b []byte) (*RemoveQERFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &RemoveQERFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case QERID:
			v, err := ie.QERID()
			if err != nil {
				return f, err
			}
			f.QERID = v
		}
	}
	return f, nil
}
