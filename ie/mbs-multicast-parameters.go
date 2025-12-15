package ie

// NewMBSMulticastParameters creates a new MulticastParameters IE.
func NewMBSMulticastParameters(ies ...*IE) *IE {
	return newGroupedIE(MBSMulticastParameters, 0, ies...)
}

// MBSMulticastParameters returns the IEs above MBSMulticastParameters if the type of IE matches.
func (i *IE) MBSMulticastParameters() (*MBSMulticastParametersFields, error) {
	switch i.Type {
	case MBSMulticastParameters:
		return ParseMBSMulticastParametersFields(i.Payload)
	case CreateFAR:
		ies, err := i.CreateFAR()
		if err != nil {
			return nil, err
		}
		if ies.MBSMulticastParameters != nil {
			return ies.MBSMulticastParameters, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MBSUnicastParametersFields is a set of fields in MBSUnicastParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type MBSMulticastParametersFields struct {
	DestinationInterface     uint8
	NetworkInstance          string
	OuterHeaderCreation      *OuterHeaderCreationFields
	TransportLevelMarking    uint16
	DestinationInterfaceType uint8
}

// ParseMBSUnicastParametersFields returns the IEs above MBSUnicastParameters.
func ParseMBSMulticastParametersFields(b []byte) (*MBSMulticastParametersFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &MBSMulticastParametersFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case DestinationInterface:
			dest, err := ie.DestinationInterface()
			if err != nil {
				return f, err
			}
			f.DestinationInterface = dest
		case NetworkInstance:
			v, err := ie.NetworkInstance()
			if err != nil {
				return f, err
			}
			f.NetworkInstance = v
		case OuterHeaderCreation:
			creation, err := ie.OuterHeaderCreation()
			if err != nil {
				return f, err
			}
			f.OuterHeaderCreation = creation
		case TransportLevelMarking:
			transport, err := ie.TransportLevelMarking()
			if err != nil {
				return f, err
			}
			f.TransportLevelMarking = transport
		case TGPPInterfaceType:
			tgppinterface, err := ie.TGPPInterfaceType()
			if err != nil {
				return f, err
			}
			f.DestinationInterfaceType = tgppinterface
		}
	}
	return f, nil
}
