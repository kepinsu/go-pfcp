// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewEthernetPacketFilter creates a new EthernetPacketFilter IE.
func NewEthernetPacketFilter(ies ...*IE) *IE {
	return newGroupedIE(EthernetPacketFilter, 0, ies...)
}

// EthernetPacketFilter returns the IEs above EthernetPacketFilter if the type of IE matches.
func (i *IE) EthernetPacketFilter() (*EthernetPacketFilterFields, error) {
	switch i.Type {
	case EthernetPacketFilter:
		return ParseEthernetPacketFilter(i.Payload)
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return nil, err
		}
		if ies.PDI != nil && ies.PDI.EthernetPacketFilter != nil {
			return ies.PDI.EthernetPacketFilter, nil
		}
		return nil, ErrIENotFound
	case PDI:
		ies, err := i.PDI()
		if err != nil {
			return nil, err
		}
		if ies.EthernetPacketFilter != nil {
			return ies.EthernetPacketFilter, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// EthernetPacketFilterFields is a set of fields in EthernetPacketFilter IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type EthernetPacketFilterFields struct {
	EthernetFilterID         uint32
	EthernetFilterProperties uint8
	MacAddress               *MACAddressFields
	EtherType                uint16
	CTag                     *CTAGFields
	STAG                     *STAGFields
	SDFFilter                []*SDFFilterFields `tlv:"23"  json:"sdf_filter,omitempty"`
}

func ParseEthernetPacketFilter(b []byte) (*EthernetPacketFilterFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &EthernetPacketFilterFields{}
	err = f.ParseIEs(ies)
	return f, err
}

// parseUserLocationInformationFields decodes parseUserLocationInformationFields.
func (f *EthernetPacketFilterFields) ParseIEs(ies []*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case EthernetFilterID:
			v, err := ie.EthernetFilterID()
			if err != nil {
				return err
			}
			f.EthernetFilterID = v
		case EthernetFilterProperties:
			v, err := ie.EthernetFilterProperties()
			if err != nil {
				return err
			}
			f.EthernetFilterProperties = v
		case MACAddress:
			v, err := ie.MACAddress()
			if err != nil {
				return err
			}
			f.MacAddress = v
		case Ethertype:
			v, err := ie.Ethertype()
			if err != nil {
				return err
			}
			f.EtherType = v
		case CTAG:
			v, err := ie.CTAG()
			if err != nil {
				return err
			}
			f.CTag = v
		case STAG:
			v, err := ie.STAG()
			if err != nil {
				return err
			}
			f.STAG = v
		case SDFFilter:
			field, err := ParseSDFFilterFields(ie.Payload)
			if err != nil {
				return err
			}
			f.SDFFilter = append(f.SDFFilter, field)
		}
	}
	return nil
}
