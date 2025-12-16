// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewWeight creates a new Weight IE.
func NewWeight(weight uint8) *IE {
	return newUint8ValIE(Weight, weight)
}

// Weight returns Weight in uint8 if the type of IE matches.
func (i *IE) Weight() (uint8, error) {
	switch i.Type {
	case Weight:
		return i.ValueAsUint8()
	case CreateMAR:
		ies, err := i.CreateMAR()
		if err != nil {
			return 0, err
		}
		if ies.TGPPAccessForwardingActionInformation != nil {
			return ies.TGPPAccessForwardingActionInformation.Weight, nil
		}
		if ies.NonTGPPAccessForwardingActionInformation != nil {
			return ies.NonTGPPAccessForwardingActionInformation.Weight, nil
		}
		return 0, ErrIENotFound
	case UpdateMAR:
		ies, err := i.UpdateMAR()
		if err != nil {
			return 0, err
		}
		if ies.TGPPAccessForwardingActionInformation != nil {
			return ies.TGPPAccessForwardingActionInformation.Weight, nil
		}
		if ies.NonTGPPAccessForwardingActionInformation != nil {
			return ies.NonTGPPAccessForwardingActionInformation.Weight, nil
		}
		return 0, ErrIENotFound
	case TGPPAccessForwardingActionInformation:
		ies, err := i.TGPPAccessForwardingActionInformation()
		if err != nil {
			return 0, err
		}
		return ies.Weight, nil
	case NonTGPPAccessForwardingActionInformation:
		ies, err := i.NonTGPPAccessForwardingActionInformation()
		if err != nil {
			return 0, err
		}

		return ies.Weight, nil
	case UpdateTGPPAccessForwardingActionInformation:
		ies, err := i.UpdateTGPPAccessForwardingActionInformation()
		if err != nil {
			return 0, err
		}
		return ies.Weight, nil
	case UpdateNonTGPPAccessForwardingActionInformation:
		ies, err := i.UpdateNonTGPPAccessForwardingActionInformation()
		if err != nil {
			return 0, err
		}
		return ies.Weight, nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
