package ie

// NewMBSSessionN4ControlInformation creates a new MBSSessionN4ControlInformation IE.
func NewMBSSessionN4ControlInformation(ies ...*IE) *IE {
	return newGroupedIE(MBSSessionN4ControlInformation, 0, ies...)
}

// MBSSessionN4ControlInformation returns the IEs above MBSSessionN4ControlInformation if the type of IE matches.
func (i *IE) MBSSessionN4ControlInformation() (*MBSSessionN4ControlInformationFields, error) {
	switch i.Type {
	case MBSSessionN4ControlInformation:
		return ParseMBSSessionN4ControlInformationFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MBSUnicastParametersFields is a set of fields in MBSUnicastParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type MBSSessionN4ControlInformationFields struct {
	MBSSessionIdentifier                           *MBSSessionIdentifierFields
	AreaSessionID                                  uint16
	MulticastTransportInformationForN3mbAndOrN19mb *MulticastTransportInformationFields
}

// ParseMBSUnicastParametersFields returns the IEs above MBSUnicastParameters.
func ParseMBSSessionN4ControlInformationFields(b []byte) (*MBSSessionN4ControlInformationFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &MBSSessionN4ControlInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case MBSSessionIdentifier:
			v, err := ie.MBSSessionIdentifier()
			if err != nil {
				return f, err
			}
			f.MBSSessionIdentifier = v
		case AreaSessionID:
			v, err := ie.AreaSessionID()
			if err != nil {
				return f, err
			}
			f.AreaSessionID = v
		case MulticastTransportInformation:
			transport, err := ie.MulticastTransportInformation()
			if err != nil {
				return f, err
			}
			f.MulticastTransportInformationForN3mbAndOrN19mb = transport
		}
	}
	return f, nil
}
