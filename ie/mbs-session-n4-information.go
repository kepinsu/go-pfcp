package ie

// NewMBSSessionN4Information creates a new MulticastParameters IE.
func NewMBSSessionN4Information(ies ...*IE) *IE {
	return newGroupedIE(MBSSessionN4mbInformation, 0, ies...)
}

// MBSSessionN4mbInformation returns the IEs above MBSSessionN4mbInformation if the type of IE matches.
func (i *IE) MBSSessionN4Information() (*MBSSessionN4InformationFields, error) {
	switch i.Type {
	case MBSSessionN4mbInformation:
		return ParseMBSSessionN4InformationFields(i.Payload)
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// MBSUnicastParametersFields is a set of fields in MBSUnicastParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type MBSSessionN4InformationFields struct {
	MBSSessionIdentifier *MBSSessionIdentifierFields
	AreaSessionID        uint16
	N19mbDLTunnelID      *FTEIDFields
	MBSN4RespFlags       uint8
}

// ParseMBSUnicastParametersFields returns the IEs above MBSUnicastParameters.
func ParseMBSSessionN4InformationFields(b []byte) (*MBSSessionN4InformationFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &MBSSessionN4InformationFields{}
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
		case FTEID:
			v, err := ie.FTEID()
			if err != nil {
				return f, err
			}
			f.N19mbDLTunnelID = v
		case MBSN4RespFlags:
			v, err := ie.MBSN4RespFlags()
			if err != nil {
				return f, err
			}
			f.MBSN4RespFlags = v
		}
	}
	return f, nil
}
