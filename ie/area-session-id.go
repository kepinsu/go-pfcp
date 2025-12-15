package ie

// NewAreaSessionID creates a new AreaSessionID IE.
func NewAreaSessionID(id uint16) *IE {
	return newUint16ValIE(AreaSessionID, id)
}

// AreaSessionID returns AreaSessionID in uint16 if the type of IE matches.
func (i *IE) AreaSessionID() (uint16, error) {
	switch i.Type {
	case AreaSessionID:
		return i.ValueAsUint16()
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
