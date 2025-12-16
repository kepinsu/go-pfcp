package ie

import (
	"net"
	"net/netip"
)

// NewNBNSServerAddress creates a new NBNSServerAddress IE.
func NewNBNSServerAddress(addr string) *IE {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return nil
	}
	// The TS 29.244 specify this IE contain IPv4 address
	// See Section 8.2.194
	if !ip.Is4() {
		return nil
	}

	return New(NBNSServerAddress, ip.AsSlice())
}

// NBNSServerAddress returns NBNSServerAddress in uint8 if the type of IE matches.
func (i *IE) NBNSServerAddress() (net.IP, error) {
	switch i.Type {
	case NBNSServerAddress:
		return net.IP(i.Payload), nil
	case CreatedL2TPSession:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == NBNSServerAddress {
				return x.NBNSServerAddress()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
