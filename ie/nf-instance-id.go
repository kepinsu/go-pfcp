// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"io"
)

// NewNFInstanceID creates a new NFInstanceID IE.
func NewNFInstanceID(id []byte) *IE {
	return New(NFInstanceID, id[:16])
}

// NFInstanceID returns NFInstanceID in []byte if the type of IE matches.
func (i *IE) NFInstanceID() ([]byte, error) {
	if len(i.Payload) < 16 {
		return nil, io.ErrUnexpectedEOF
	}

	switch i.Type {
	case NFInstanceID:
		return i.Payload[:16], nil
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}
