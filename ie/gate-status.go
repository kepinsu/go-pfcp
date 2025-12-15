// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// GateStatus definitions.
const (
	GateStatusOpen   uint8 = 0
	GateStatusClosed uint8 = 1
)

// NewGateStatus creates a new GateStatus IE.
func NewGateStatus(ul, dl uint8) *IE {
	return newUint8ValIE(GateStatus, (ul<<2)|dl)
}

// GateStatus returns GateStatus in uint8 if the type of IE matches.
func (i *IE) GateStatus() (Gate, error) {
	switch i.Type {
	case GateStatus:
		gateStatus, err := i.ValueAsUint8()
		return Gate(gateStatus), err
	case CreateQER:
		ies, err := i.CreateQER()
		if err != nil {
			return 0, err
		}
		return ies.GateStatus, nil
	case UpdateQER:
		ies, err := i.UpdateQER()
		if err != nil {
			return 0, err
		}
		return ies.GateStatus, nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// Gate is alias of uint8 to determine UL and DL status
type Gate uint8

// GateStatusUL returns GateStatusUL in uint8
func (g Gate) GateStatusUL() uint8 {
	return (uint8(g >> 2)) & 0x03
}

// GateStatusDL returns GateStatusDL in uint8
func (g Gate) GateStatusDL() uint8 {
	return uint8(g) & 0x03
}

// GateStatusULDL returns GateStatusUL and GateStatusDL in uint8
func (g Gate) GateStatusULDL() (uint8, uint8, error) {
	return g.GateStatusUL(), g.GateStatusDL(), nil
}
