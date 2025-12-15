package ie

// NewMBSSessionN4mbControlInformation creates a new MulticastParameters IE.
func NewMBSSessionN4mbControlInformation(ies ...*IE) *IE {
	return newGroupedIE(MBSSessionN4mbControlInformation, 0, ies...)
}

// MBSSessionN4mbControlInformation returns the IEs above MBSSessionN4mbControlInformation if the type of IE matches.
func (i *IE) MBSSessionN4mbControlInformation() (*MBSSessionN4mbControlInformationFields, error) {
	switch i.Type {
	case MBSSessionN4mbControlInformation:
		return ParseMBSSessionN4mbControlInformationFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MBSUnicastParametersFields is a set of fields in MBSUnicastParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type MBSSessionN4mbControlInformationFields struct {
	MBSSessionIdentifier                           *MBSSessionIdentifierFields
	AreaSessionID                                  uint16
	MBSN4mbReqFlags                                uint8
	MulticastTransportInformationForN3mbAndOrN19mb *MulticastTransportInformationFields
}

// ParseMBSUnicastParametersFields returns the IEs above MBSUnicastParameters.
func ParseMBSSessionN4mbControlInformationFields(b []byte) (*MBSSessionN4mbControlInformationFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &MBSSessionN4mbControlInformationFields{}
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
		case MBSN4mbReqFlags:
			v, err := ie.MBSN4mbReqFlags()
			if err != nil {
				return f, err
			}
			f.MBSN4mbReqFlags = v
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
