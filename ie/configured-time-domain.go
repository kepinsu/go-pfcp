// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewConfiguredTimeDomain creates a new ConfiguredTimeDomain IE.
func NewConfiguredTimeDomain(flags uint8) *IE {
	return newUint8ValIE(ConfiguredTimeDomain, flags)
}

// ConfiguredTimeDomain returns ConfiguredTimeDomain in uint8 if the type of IE matches.
func (i *IE) ConfiguredTimeDomain() (uint8, error) {

	switch i.Type {
	case ConfiguredTimeDomain:
		return i.ValueAsUint8()
	case ClockDriftControlInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, i := range ies {
			if i == nil {
				continue
			}
			switch i.Type {
			case ConfiguredTimeDomain:
				return i.ConfiguredTimeDomain()
			}
		}
		return 0, ErrElementNotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
