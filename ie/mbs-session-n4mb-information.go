package ie

// NewMBSSessionN4mbInformation creates a new MulticastParameters IE.
func NewMBSSessionN4mbInformation(ies ...*IE) *IE {
	return newGroupedIE(MBSSessionN4mbInformation, 0, ies...)
}

// MBSSessionN4mbInformation returns the IEs above MBSSessionN4mbInformation if the type of IE matches.
func (i *IE) MBSSessionN4mbInformation() (*MBSSessionN4mbInformationFields, error) {
	switch i.Type {
	case MBSSessionN4mbInformation:
		return ParseMBSSessionN4mbInformationFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MBSUnicastParametersFields is a set of fields in MBSUnicastParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type MBSSessionN4mbInformationFields struct {
	MulticastTransportInformation *MulticastTransportInformationFields
}

// ParseMBSUnicastParametersFields returns the IEs above MBSUnicastParameters.
func ParseMBSSessionN4mbInformationFields(b []byte) (*MBSSessionN4mbInformationFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &MBSSessionN4mbInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		if ie.Type == MulticastTransportInformation {
			transport, err := ie.MulticastTransportInformation()
			if err != nil {
				return f, err
			}
			f.MulticastTransportInformation = transport
		}
	}
	return f, nil
}
