// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewL2TPSessionInformation creates a new L2TPSessionInformation IE.
func NewL2TPSessionInformation(ies ...*IE) *IE {
	return newGroupedIE(L2TPSessionInformation, 0, ies...)
}

// L2TPSessionInformation returns the IEs above L2TPSessionInformation if the type of IE matches.
func (i *IE) L2TPSessionInformation() (*L2TPSessionInformationFields, error) {
	if i.Type != L2TPSessionInformation {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseL2TPSessionInformationFields(i.Payload)
}

// L2TPSessionInformationFields represents a fields contained in L2TPSessionInformation IE.
type L2TPSessionInformationFields struct {
	CallingNumber          string
	CalledNumber           string
	MaximumReceiveUnit     uint16
	L2TPSessionIndications uint8
	L2TPUserAuthentication *L2TPUserAuthenticationFields
}

// ParseL2TPSessionInformationFields parses b into L2TPSessionInformationFields.
func ParseL2TPSessionInformationFields(b []byte) (*L2TPSessionInformationFields, error) {

	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	p := &L2TPSessionInformationFields{}
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case CallingNumber:
			v, err := ie.CallingNumber()
			if err != nil {
				return p, err
			}
			p.CallingNumber = v
		case CalledNumber:
			v, err := ie.CalledNumber()
			if err != nil {
				return p, err
			}
			p.CalledNumber = v
		case MaximumReceiveUnit:
			v, err := ie.MaximumReceiveUnit()
			if err != nil {
				return p, err
			}
			p.MaximumReceiveUnit = v
		case L2TPSessionIndications:
			v, err := ie.L2TPSessionIndications()
			if err != nil {
				return p, err
			}
			p.L2TPSessionIndications = v
		case L2TPUserAuthentication:
			v, err := ie.L2TPUserAuthentication()
			if err != nil {
				return p, err
			}
			p.L2TPUserAuthentication = v
		}
	}
	return p, nil
}
