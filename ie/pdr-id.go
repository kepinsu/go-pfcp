// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPDRID creates a new PDRID IE.
func NewPDRID(id uint16) *IE {
	return newUint16ValIE(PDRID, id)
}

// PDRID returns PDRID in uint16 if the type of IE matches.
func (i *IE) PDRID() (uint16, error) {
	switch i.Type {
	case PDRID:
		return i.ValueAsUint16()
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return 0, err
		}
		if ies.ID > 0 {
			return ies.ID, nil
		}
		return 0, ErrIENotFound
	case UpdatePDR:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PDRID {
				return x.PDRID()
			}
		}
		return 0, ErrIENotFound
	case RemovePDR:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PDRID {
				return x.PDRID()
			}
		}
		return 0, ErrIENotFound
	case CreatedPDR:
		ies, err := i.CreatedPDR()
		if err != nil {
			return 0, err
		}
		if ies.ID > 0 {
			return ies.ID, nil
		}
		return 0, ErrIENotFound
	case ApplicationDetectionInformation:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PDRID {
				return x.PDRID()
			}
		}
		return 0, ErrIENotFound
	case UsageReportWithinSessionReportRequest:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == ApplicationDetectionInformation {
				return x.PDRID()
			}
		}
		return 0, ErrIENotFound
	case DownlinkDataReport:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PDRID {
				return x.PDRID()
			}
		}
		return 0, ErrIENotFound
	case UpdatedPDR:
		ies, err := ParseMultiIEs(i.Payload)
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == PDRID {
				return x.PDRID()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
