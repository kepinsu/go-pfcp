// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import "github.com/wmnsk/go-pfcp/internal/utils"

// NewSMFSetID creates a new SMFSetID IE.
func NewSMFSetID(id string) *IE {

	// Warning in the SMF-SET-ID we had the first octet like a spare
	//   Octets    Bits
	//   ---------------------------------------------
	//   1 - 2     Option Type = 180 (decimal)
	//   3 - 4     Length = n
	//   5         Spare
	//   6 - m     FQDN
	fqdn := utils.EncodeFQDN(id)
	value := make([]byte, len(fqdn)+1)
	copy(value[1:], fqdn)
	return New(SMFSetID, value)
}

// SMFSetID returns SMFSetID in string if the type of IE matches.
func (i *IE) SMFSetID() (string, error) {
	if i.Type != SMFSetID {
		return "", &InvalidTypeError{Type: i.Type}
	}

	// Warning in the SMF-SET-ID we had the first octet like a spare
	//   Octets    Bits
	//   ---------------------------------------------
	//   1 - 2     Option Type = 180 (decimal)
	//   3 - 4     Length = n
	//   5         Spare
	//   6 - m     FQDN
	if len(i.Payload) == 0 || len(i.Payload) == 1 {
		return "", nil
	}
	// If the IE is grouped
	if i.IsGrouped() {
		return "", &InvalidTypeError{Type: i.Type}
	}
	return utils.DecodeFQDN(i.Payload[1:]), nil
}
