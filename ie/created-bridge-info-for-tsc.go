// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "net"

// NewCreatedBridgeInfoForTSC creates a new CreatedBridgeInfoForTSC IE.
func NewCreatedBridgeInfoForTSC(ies ...*IE) *IE {
	return newGroupedIE(CreatedBridgeInfoForTSC, 0, ies...)
}

// CreatedBridgeInfoForTSC returns the IEs above CreatedBridgeInfoForTSC if the type of IE matches.
func (i *IE) CreatedBridgeInfoForTSC() (*CreatedBridgeInfoForTSCFields, error) {
	if i.Type != CreatedBridgeInfoForTSC {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseCreatedBridgeInfoForTSCFields(i.Payload)
}

// CreatedBridgeInfoForTSCFields is a set of fields in CreatedBridgeInfoForTSC IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type CreatedBridgeInfoForTSCFields struct {
	DSTTPortNumber uint32
	TSNBridgeID    net.HardwareAddr
}

// ParseCreatedBridgeInfoForTSCFields returns the IEs above CreatedBridgeInfoForTSC
func ParseCreatedBridgeInfoForTSCFields(b []byte) (*CreatedBridgeInfoForTSCFields, error) {
	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	far := &CreatedBridgeInfoForTSCFields{}
	if err := far.ParseIEs(ies...); err != nil {
		return far, err
	}
	return far, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (far *CreatedBridgeInfoForTSCFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case DSTTPortNumber:
			v, err := ie.DSTTPortNumber()
			if err != nil {
				return err
			}
			far.DSTTPortNumber = v

		case TSNBridgeID:
			v, err := ie.TSNBridgeID()
			if err != nil {
				return err
			}
			far.TSNBridgeID = v
		}
	}
	return nil
}
