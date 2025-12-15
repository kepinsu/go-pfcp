package ie

import (
	"net"
	"net/netip"
)

// NewLNSAddress creates a new LNSAddress IE.
func NewLNSAddress(addr string) *IE {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return nil
	}
	return New(LNSAddress, ip.AsSlice())
}

// LNSAddress returns LNSAddress in uint8 if the type of IE matches.
func (i *IE) LNSAddress() (net.IP, error) {
	switch i.Type {
	case LNSAddress:
		return net.IP(i.Payload), nil
	case L2TPTunnelInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return nil, err
		}
		for _, x := range ies {
			if x.Type == LNSAddress {
				return x.LNSAddress()
			}
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
