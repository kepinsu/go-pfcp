package ie

// NewCalledNumber creates a new CalledNumber IE.
func NewCalledNumber(number string) *IE {
	return newStringIE(CalledNumber, number)
}

// CalledNumber returns CalledNumber in uint8 if the type of IE matches.
func (i *IE) CalledNumber() (string, error) {
	switch i.Type {
	case CalledNumber:
		return i.ValueAsString()
	case L2TPSessionInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return "", err
		}
		for _, x := range ies {
			if x.Type == CalledNumber {
				return x.CalledNumber()
			}
		}
		return "", ErrIENotFound
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
