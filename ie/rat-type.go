// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// RATType definitions.
const (
	RATTypeUtran        uint8 = 1
	RATTypeGeran        uint8 = 2
	RATTypeWlan         uint8 = 3
	RATTypeGan          uint8 = 4
	RATTypeHslavolution uint8 = 5
	RATTypeEutran       uint8 = 6
	RATTypeVirtual      uint8 = 7
	RATTypeEutranNbIot  uint8 = 8
	RATTypeLTEM         uint8 = 9
	RATTypeNR           uint8 = 10
)

// NewRATType creates a new RATType IE.
func NewRATType(rat uint8) *IE {
	return NewUint8IE(RATType, rat)
}

// RATType returns RATType in uint8 if the type of IE matches.
func (i *IE) RATType() (uint8, error) {
	switch i.Type {
	case RATType:
		return i.ValueAsUint8()
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
