package ie

import (
	"net"
	"net/netip"
)

// NewDNSServerAddress creates a new DNSServerAddress IE.
func NewDNSServerAddress(addr string) *IE {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return nil
	}
	// The TS 29.244 specify this IE contain IPv4 address
	// See Section 8.2.193
	if !ip.Is4() {
		return nil
	}

	return New(DNSServerAddress, ip.AsSlice())
}

// DNSServerAddress returns DNSServerAddress in uint8 if the type of IE matches.
func (i *IE) DNSServerAddress() (net.IP, error) {
	switch i.Type {
	case DNSServerAddress:
		return net.IP(i.Payload), nil
	case CreatedL2TPSession:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == DNSServerAddress {
				return x.DNSServerAddress()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
