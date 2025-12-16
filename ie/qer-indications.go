package ie

// NewQERIndications creates a new QERIndications IE.
func NewQERIndications(flag uint8) *IE {
	return newUint8ValIE(QERIndications, flag)
}

// QERIndications returns QERIndications in uint8 if the type of IE matches.
func (i *IE) QERIndications() (uint8, error) {
	if i.Type != QERIndications {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.ValueAsUint8()
}
