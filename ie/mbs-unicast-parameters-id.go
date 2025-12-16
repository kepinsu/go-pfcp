package ie

// NewMBSUnicastParametersID creates a new MBSUnicastParametersID IE.
func NewMBSUnicastParametersID(id uint16) *IE {
	return newUint16ValIE(MBSUnicastParametersID, id)
}

// MBSUnicastParametersID returns MBSUnicastParametersID in uint8 if the type of IE matches.
func (i *IE) MBSUnicastParametersID() (uint16, error) {
	switch i.Type {
	case MBSUnicastParametersID:
		return i.ValueAsUint16()
	case AddMBSUnicastParameters:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == MBSUnicastParametersID {
				return x.MBSUnicastParametersID()
			}
		}
		return 0, ErrIENotFound
	case RemoveMBSUnicastParameters:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == MBSUnicastParametersID {
				return x.MBSUnicastParametersID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
