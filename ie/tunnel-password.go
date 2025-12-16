// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewTunnelPassword creates a new TunnelPassword IE.
func NewTunnelPassword(pwd string) *IE {
	return newStringIE(TunnelPassword, pwd)
}

// TunnelPassword returns TunnelPassword in *TunnelPasswordFields if the type of IE matches.
func (i *IE) TunnelPassword() (string, error) {
	switch i.Type {
	case TunnelPassword:
		return i.ValueAsString()
	case L2TPTunnelInformation:
		ies, err := i.L2TPTunnelInformation()
		if err != nil {
			return "", err
		}
		if len(ies.TunnelPassword) > 0 {
			return ies.TunnelPassword, nil
		}
		return "", ErrIENotFound
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
