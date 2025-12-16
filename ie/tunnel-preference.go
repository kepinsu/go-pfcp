// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "github.com/wmnsk/go-pfcp/internal/utils"

// NewTunnelPreference creates a new TunnelPreference IE.
func NewTunnelPreference(v uint32) *IE {
	i := New(TunnelPreference, make([]byte, 3))
	copy(i.Payload[0:3], utils.Uint32To24(v))
	return i
}

// TunnelPreference returns TunnelPreference in *TunnelPreferenceFields if the type of IE matches.
func (i *IE) TunnelPreference() (uint32, error) {
	switch i.Type {
	case TunnelPreference:
		return utils.Uint24To32(i.Payload), nil
	case L2TPTunnelInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, ie := range ies {
			if ie == nil {
				continue
			}
			if ie.Type == TunnelPreference {
				return ie.TunnelPreference()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
