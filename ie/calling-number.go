package ie

// NewCallingNumber creates a new CallingNumber IE.
func NewCallingNumber(number string) *IE {
	return newStringIE(CallingNumber, number)
}

// CallingNumber returns CallingNumber in uint8 if the type of IE matches.
func (i *IE) CallingNumber() (string, error) {
	switch i.Type {
	case CallingNumber:
		return i.ValueAsString()
	case L2TPSessionInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return "", err
		}
		for _, x := range ies {
			if x.Type == CallingNumber {
				return x.CallingNumber()
			}
		}
		return "", ErrIENotFound
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
