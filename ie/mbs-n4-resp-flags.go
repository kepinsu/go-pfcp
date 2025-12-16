package ie

// NewMBSN4mbReqFlags creates a new MBSN4mbReqFlags IE.
func NewMBSN4RespFlags(flag uint8) *IE {
	return newUint8ValIE(MBSN4RespFlags, flag)
}

// MBSN4mbReqFlags returns MBSN4mbReqFlags in uint8 if the type of IE matches.
func (i *IE) MBSN4RespFlags() (uint8, error) {
	if i.Type != MBSN4RespFlags {
		return 0, &InvalidTypeError{Type: i.Type}
	}

	return i.ValueAsUint8()
}
