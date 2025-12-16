package ie

// NewAddMBSUnicastParameters creates a new AbsMBSUnicastParameters IE.
func NewAddMBSUnicastParameters(ies ...*IE) *IE {
	return newGroupedIE(AddMBSUnicastParameters, 0, ies...)
}

// AddMBSUnicastParameters returns the IEs above AddMBSUnicastParameters if the type of IE matches.
func (i *IE) AddMBSUnicastParameters() (*AddMBSUnicastParametersFields, error) {
	switch i.Type {
	case AddMBSUnicastParameters:
		return ParseAddMBSUnicastParametersFields(i.Payload)
	case CreateFAR:
		ies, err := i.CreateFAR()
		if err != nil {
			return nil, err
		}
		for _, i := range ies.AddMBSUnicastParameters {
			if i == nil {
				continue
			}
			return i, nil
		}
		return nil, ErrIENotFound
	case UpdateFAR:
		ies, err := i.UpdateFAR()
		if err != nil {
			return nil, err
		}
		for _, i := range ies.AddMBSUnicastParameters {
			if i == nil {
				continue
			}
			return i, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// AddMBSUnicastParametersFields is a set of fields in AddMBSUnicastParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type AddMBSUnicastParametersFields struct {
	DestinationInterface     uint8
	MBSUnicastParametersID   uint16
	NetworkInstance          string
	OuterHeaderCreation      *OuterHeaderCreationFields
	TransportLevelMarking    uint16
	DestinationInterfaceType uint8
}

// ParseAddMBSUnicastParametersFields returns the IEs above AddMBSUnicastParameters.
func ParseAddMBSUnicastParametersFields(b []byte) (*AddMBSUnicastParametersFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &AddMBSUnicastParametersFields{}
	if err := f.ParseIEs(ies...); err != nil {
		return f, err
	}
	return f, nil
}

// ParseIEs will iterator over all childs IE to avoid to use Parse or ParseMultiIEs any time we iterate in IE
func (a *AddMBSUnicastParametersFields) ParseIEs(ies ...*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case DestinationInterface:
			dest, err := ie.DestinationInterface()
			if err != nil {
				return err
			}
			a.DestinationInterface = dest
		case MBSUnicastParametersID:
			v, err := ie.MBSUnicastParametersID()
			if err != nil {
				return err
			}
			a.MBSUnicastParametersID = v
		case NetworkInstance:
			v, err := ie.NetworkInstance()
			if err != nil {
				return err
			}
			a.NetworkInstance = v
		case OuterHeaderCreation:
			creation, err := ie.OuterHeaderCreation()
			if err != nil {
				return err
			}
			a.OuterHeaderCreation = creation
		case TransportLevelMarking:
			transport, err := ie.TransportLevelMarking()
			if err != nil {
				return err
			}
			a.TransportLevelMarking = transport
		case TGPPInterfaceType:
			tgppinterface, err := ie.TGPPInterfaceType()
			if err != nil {
				return err
			}
			a.DestinationInterfaceType = tgppinterface
		}
	}
	return nil
}
