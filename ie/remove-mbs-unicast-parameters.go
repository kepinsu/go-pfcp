package ie

// NewRemoveMBSUnicastParameters creates a new AbsMBSUnicastParameters IE.
func NewRemoveMBSUnicastParameters(ies ...*IE) *IE {
	return newGroupedIE(RemoveMBSUnicastParameters, 0, ies...)
}

// RemoveMBSUnicastParameters returns the IEs above RemoveMBSUnicastParameters if the type of IE matches.
func (i *IE) RemoveMBSUnicastParameters() (*RemoveMBSUnicastParametersFields, error) {
	switch i.Type {
	case RemoveMBSUnicastParameters:
		return ParseRemoveMBSUnicastParametersFields(i.Payload)
	case UpdateFAR:
		ies, err := i.UpdateFAR()
		if err != nil {
			return nil, err
		}
		for _, i := range ies.RemoveMBSUnicastParameters {
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

// RemoveMBSUnicastParametersFields is a set of fields in RemoveMBSUnicastParameters IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type RemoveMBSUnicastParametersFields struct {
	MBSUnicastParametersID uint16
}

// ParseRemoveMBSUnicastParametersFields returns the IEs above RemoveMBSUnicastParameters.
func ParseRemoveMBSUnicastParametersFields(b []byte) (*RemoveMBSUnicastParametersFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &RemoveMBSUnicastParametersFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}

		switch ie.Type {
		case MBSUnicastParametersID:
			v, err := ie.MBSUnicastParametersID()
			if err != nil {
				return f, err
			}
			f.MBSUnicastParametersID = v
		}
	}
	return f, nil
}
