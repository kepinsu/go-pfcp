// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPortManagementInformationContainer creates a new PortManagementInformationContainer IE.
func NewPortManagementInformationContainer(info string) *IE {
	return newStringIE(PortManagementInformationContainer, info)
}

// PortManagementInformationContainer returns PortManagementInformationContainer in string if the type of IE matches.
func (i *IE) PortManagementInformationContainer() (string, error) {
	switch i.Type {
	case PortManagementInformationContainer:
		return i.ValueAsString()
	case TSCManagementInformationWithinSessionModificationRequest,
		TSCManagementInformationWithinSessionModificationResponse,
		TSCManagementInformationWithinSessionReportRequest:
		ies, err := i.TSCManagementInformation()
		if err != nil {
			return "", err
		}
		if len(ies.PortManagementInformationContainer) > 0 {
			return ies.PortManagementInformationContainer, nil
		}
		return "", ErrIENotFound
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
