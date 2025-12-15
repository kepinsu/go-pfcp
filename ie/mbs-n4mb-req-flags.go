package ie

// NewMBSN4mbReqFlags creates a new MBSN4mbReqFlags IE.
func NewMBSN4mbReqFlags(flag uint8) *IE {
	return newUint8ValIE(MBSN4mbReqFlags, flag)
}

// MBSN4mbReqFlags returns MBSN4mbReqFlags in uint8 if the type of IE matches.
func (i *IE) MBSN4mbReqFlags() (uint8, error) {
	if i.Type != MBSN4mbReqFlags {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.ValueAsUint8()
}
