// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewGroupID creates a new GroupID IE.
func NewGroupID(groupID []byte) *IE {
	return New(GroupID, groupID)
}

// GroupID returns GroupID in string if the type of IE matches.
func (i *IE) GroupID() ([]byte, error) {
	switch i.Type {
	case GroupID:
		return i.Payload, nil
	case PFCPSessionChangeInfo:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return nil, err
		}
		for _, ie := range ies {
			if ie.Type == GroupID {
				return ie.GroupID()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
