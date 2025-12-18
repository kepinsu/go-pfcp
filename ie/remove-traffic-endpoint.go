// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewRemoveTrafficEndpoint creates a new RemoveTrafficEndpoint IE.
func NewRemoveTrafficEndpoint(id *IE) *IE {
	return newGroupedIE(RemoveTrafficEndpoint, 0, id)
}

// RemoveTrafficEndpoint returns the IEs above RemoveTrafficEndpoint if the type of IE matches.
func (i *IE) RemoveTrafficEndpoint() (*RemoveTrafficEndpointFields, error) {
	if i.Type != RemoveTrafficEndpoint {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseRemoveTrafficEndpointFields(i.Payload)
}

// RemoveTrafficEndpointFields is a set of fields in RemoveTrafficEndpoint IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RemoveTrafficEndpointFields struct {
	TrafficEndpointID uint8
}

// ParseRemoveTrafficEndpointFields returns the IEs above RemoveTrafficEndpoint.
func ParseRemoveTrafficEndpointFields(b []byte) (*RemoveTrafficEndpointFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &RemoveTrafficEndpointFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == TrafficEndpointID {
			v, err := ie.TrafficEndpointID()
			if err != nil {
				return f, err
			}
			f.TrafficEndpointID = v
		}
	}
	return f, nil
}
