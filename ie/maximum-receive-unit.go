package ie

// NewMaximumReceiveUnit creates a new MaximumReceiveUnit IE.
func NewMaximumReceiveUnit(v uint16) *IE {
	return newUint16ValIE(MaximumReceiveUnit, v)
}

// MaximumReceiveUnit returns MaximumReceiveUnit in uint8 if the type of IE matches.
func (i *IE) MaximumReceiveUnit() (uint16, error) {
	switch i.Type {
	case MaximumReceiveUnit:
		return i.ValueAsUint16()
	case L2TPTunnelInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == MaximumReceiveUnit {
				return x.MaximumReceiveUnit()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
