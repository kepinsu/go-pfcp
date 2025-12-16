// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewFramedRoute creates a new FramedRoute IE.
func NewFramedRoute(name string) *IE {
	return newStringIE(FramedRoute, name)
}

// FramedRoute returns FramedRoute in string if the type of IE matches.
func (i *IE) FramedRoute() (string, error) {
	switch i.Type {
	case FramedRoute:
		return i.ValueAsString()
	case CreateTrafficEndpoint:
		for _, x := range i.ChildIEs {
			if x.Type == FramedIPv6Route {
				return x.FramedIPv6Route()
			}
		}
		return "", ErrIENotFound
	case UpdateTrafficEndpoint:
		for _, x := range i.ChildIEs {
			if x.Type == FramedIPv6Route {
				return x.FramedIPv6Route()
			}
		}
		return "", ErrIENotFound
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
