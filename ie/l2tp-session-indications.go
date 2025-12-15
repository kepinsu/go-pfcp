package ie

// NewL2TPSessionIndications creates a new L2TpNewL2TPSessionIndications IE.
func NewL2TPSessionIndications(flags uint8) *IE {
	return newUint8ValIE(L2TPSessionIndications, flags)
}

// L2TPSessionIndications returns L2TPSessionIndications in uint8 if the type of IE matches.
func (i *IE) L2TPSessionIndications() (uint8, error) {
	switch i.Type {
	case L2TPSessionIndications:
		return i.ValueAsUint8()
	case L2TPTunnelInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == L2TPSessionIndications {
				return x.L2TPSessionIndications()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
