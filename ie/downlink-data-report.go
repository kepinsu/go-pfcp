// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewDownlinkDataReport creates a new DownlinkDataReport IE.
func NewDownlinkDataReport(ies ...*IE) *IE {
	return newGroupedIE(DownlinkDataReport, 0, ies...)
}

// DownlinkDataReport returns the IEs above DownlinkDataReport if the type of IE matches.
func (i *IE) DownlinkDataReport() (*DownlinkDataReportFields, error) {
	if i.Type != DownlinkDataReport {
		return nil, &InvalidTypeError{Type: i.Type}
	}

	return ParseDownlinkDataReport(i.Payload)
}

// DownlinkDataReportFields is a set of fields in DownlinkDataReport IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type DownlinkDataReportFields struct {
	PDRID                          uint16
	DownlinkDataServiceInformation []byte
	DLDataPacketsSize              uint16
	DLDataStatus                   uint8
}

func ParseDownlinkDataReport(b []byte) (*DownlinkDataReportFields, error) {
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	f := &DownlinkDataReportFields{}
	err = f.ParseIEs(ies)
	return f, err
}

// parseUserLocationInformationFields decodes parseUserLocationInformationFields.
func (f *DownlinkDataReportFields) ParseIEs(ies []*IE) error {
	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case PDRID:
			v, err := ie.PDRID()
			if err != nil {
				return err
			}
			f.PDRID = v
		case DownlinkDataServiceInformation:
			v, err := ie.DownlinkDataServiceInformation()
			if err != nil {
				return err
			}
			f.DownlinkDataServiceInformation = v
		case DLDataPacketsSize:
			v, err := ie.DLDataPacketsSize()
			if err != nil {
				return err
			}
			f.DLDataPacketsSize = v
		case DataStatus:
			v, err := ie.DataStatus()
			if err != nil {
				return err
			}
			f.DLDataStatus = v
		}
	}
	return nil
}
