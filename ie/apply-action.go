// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "io"

// NewApplyAction creates a new ApplyAction IE.
func NewApplyAction(flagsOctets ...uint8) *IE {
	return New(ApplyAction, flagsOctets)
}

// ApplyAction returns ApplyAction in []byte if the type of IE matches.
func (i *IE) ApplyAction() (*ApplyActionFields, error) {
	switch i.Type {
	case ApplyAction:
		if len(i.Payload) < 1 {
			return nil, io.ErrUnexpectedEOF
		}
		return ParseApplyAction(i.Payload)
	case CreateFAR:
		ies, err := i.CreateFAR()
		if err != nil {
			return nil, err
		}
		if ies.ApplyAction != nil {
			return ies.ApplyAction, nil
		}
		return nil, ErrIENotFound
	case UpdateFAR:
		ies, err := i.UpdateFAR()
		if err != nil {
			return nil, err
		}
		if ies.ApplyAction != nil {
			return ies.ApplyAction, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// ApplyActionFields is a set of fields in CreateFAR IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type ApplyActionFields struct {
	// First octet

	// this indicates a request to drop the packets.
	Drop bool
	// this indicates a request to forward the packets.
	Forw bool
	// this indicates a request to buffer the packets.
	Buff bool
	// this indicates a request to notify the CP function about
	// the arrival of a first downlink packet being buffered.
	Nocp bool
	// this indicates a request to duplicate the packet
	Dupl bool
	// this indicates a request to accept UE Requests to join an IP multicast
	Ipma bool
	// this indicates a request to deny UE Requests to join an IP multicast
	Ipmd bool
	// this indicates a request to duplicate the packets for redundant transmission
	Drft bool

	// Second octet

	// when set to "1", this indicates a
	// request to eliminate duplicate packets used for redundant transmission
	Edrt bool
	// when set to "1",  this indicates a request to notify the CP
	// function about the first buffered DL packet for downlink data delivery status notification
	Bdpn bool

	// when set to "1",  this indicates a request to notify the CP
	// function about the first buffered DL packet for downlink data delivery status notification
	Ddpn bool

	// when set to "1", this indicates a request to the MB-UPF to
	// forward MBS session data towards a low layer SSM address allocated by the MB-UPF using multicast transport
	Fssm bool
	// when set to "1", this indicates a
	// request to forward and replicate MBS session data towards multiple remote GTP-U peers using unicast transport
	Mbsu bool
}

// ParseApplyAction will parse the IE
func ParseApplyAction(b []byte) (*ApplyActionFields, error) {
	a := &ApplyActionFields{}
	if err := a.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return a, nil
}

// UnmarshalBinary parses b into IE.
func (a *ApplyActionFields) UnmarshalBinary(b []byte) error {
	l := len(b)
	if l < 1 {
		return io.ErrUnexpectedEOF
	}
	firstFlags := b[0]
	a.Drop = has1stBit(firstFlags)
	a.Forw = has2ndBit(firstFlags)
	a.Buff = has3rdBit(firstFlags)
	a.Nocp = has4thBit(firstFlags)
	a.Dupl = has5thBit(firstFlags)
	a.Ipma = has6thBit(firstFlags)
	a.Ipmd = has7thBit(firstFlags)
	a.Drft = has8thBit(firstFlags)
	if l > 1 {
		secondFlags := b[1]
		a.Edrt = has1stBit(secondFlags)
		a.Bdpn = has2ndBit(secondFlags)
		a.Ddpn = has3rdBit(secondFlags)
		a.Fssm = has4thBit(secondFlags)
		a.Mbsu = has5thBit(secondFlags)
	}

	return nil
}

// ValidateApplyAction can be used to facilitate the detection of some inconsistencies in Apply Action flags.
// Its use is optional because validation could also be done on upper layers, or completely skipped for testing purposes.
func (a *ApplyActionFields) ValidateApplyAction() error {
	// One and only one of the DROP, FORW, BUFF, IPMA and IPMD flags shall be set to "1".
	flags := []bool{a.Drop, a.Forw, a.Buff, a.Ipma, a.Ipmd}
	counter := 0
	for _, v := range flags {
		if v {
			counter++
		}
	}
	if counter != 1 {
		return ErrMalformed
	}
	// The NOCP flag and BDPN flag may only be set if the BUFF flag is set.
	if (a.Nocp || a.Bdpn) && !a.Buff {
		return ErrMalformed
	}
	// The DUPL flag may be set with any of the DROP, FORW, BUFF and NOCP flags.
	if a.Dupl && !(a.Drop || a.Forw || a.Buff || a.Nocp) {
		return ErrMalformed
	}
	// The DFRT flag may only be set if the FORW flag is set.
	// Note: in TS 29.244 V18.0.1 (most recent as of writing), there is a typo and DFRN is stated instead of DFRT
	if a.Drft && !a.Forw {
		return ErrMalformed
	}
	// The DDPN flag may be set with any of the DROP and BUFF flags.
	if a.Ddpn && !a.Drop || a.Buff {
		return ErrMalformed
	}

	// Note: The following is also stated in TS 29.244 section 8.2.26 V18.0.1,
	// but since "may" is used and not "may only"
	// it cannot be used to check for inconsistent IEs:
	// - The EDRT flag may be set if the FORW flag is set.
	// - Both the MBSU flag and the FSSM flag may be set […]
	return nil
}

// Marshal returns the serialized bytes of ApplyActionFields.
func (a *ApplyActionFields) Marshal() ([]byte, error) {
	b := make([]byte, a.MarshalLen())
	if err := a.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (a *ApplyActionFields) MarshalTo(b []byte) error {
	l := len(b)
	if l < 2 {
		return io.ErrUnexpectedEOF
	}
	var (
		firstOctet  uint8
		secondOctet uint8
	)
	if a.Drop {
		firstOctet |= firstBit
	}
	if a.Forw {
		firstOctet |= secondBit
	}
	if a.Buff {
		firstOctet |= thirdBit
	}
	if a.Nocp {
		firstOctet |= fourthBit
	}
	if a.Dupl {
		firstOctet |= fivethBit
	}
	if a.Ipma {
		firstOctet |= sixthBit
	}
	if a.Ipmd {
		firstOctet |= seventhBit
	}
	if a.Drft {
		firstOctet |= seventhBit
	}
	if a.Edrt {
		secondOctet |= firstBit
	}
	if a.Bdpn {
		secondOctet |= secondBit
	}
	if a.Ddpn {
		secondOctet |= thirdBit
	}
	if a.Fssm {
		secondOctet |= fourthBit
	}
	if a.Mbsu {
		secondOctet |= fivethBit
	}
	b[0] = firstOctet
	b[1] = secondOctet
	return nil
}

// MarshalLen returns field length in integer.
func (a *ApplyActionFields) MarshalLen() int {
	// In Release 17, the minimal size is 2 octets
	return 2
}
