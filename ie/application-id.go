// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewApplicationID creates a new ApplicationID IE.
func NewApplicationID(instance string) *IE {
	return newStringIE(ApplicationID, instance)
}

// ApplicationID returns ApplicationID in string if the type of IE matches.
func (i *IE) ApplicationID() (string, error) {
	switch i.Type {
	case ApplicationID:
		return i.ValueAsString()
	case CreatePDR:
		ies, err := i.CreatePDR()
		if err != nil {
			return "", err
		}
		if ies.PDI != nil && len(ies.PDI.ApplicationID) > 0 {
			return ies.PDI.ApplicationID, nil
		}
		return "", ErrIENotFound
	case PDI:
		ies, err := i.PDI()
		if err != nil {
			return "", err
		}
		return ies.ApplicationID, nil
	case ApplicationIDsPFDs:
		ies, err := i.ApplicationIDsPFDs()
		if err != nil {
			return "", err
		}
		if len(ies.ApplicationID) > 0 {
			return ies.ApplicationID, nil
		}
		return "", ErrIENotFound
	case ApplicationDetectionInformation:
		ies, err := i.ApplicationDetectionInformation()
		if err != nil {
			return "", err
		}
		if len(ies.ApplicationID) > 0 {
			return ies.ApplicationID, nil
		}
		return "", ErrIENotFound
	case UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return "", err
		}
		if ies.ApplicationDetectionInformation != nil && len(ies.ApplicationDetectionInformation.ApplicationID) > 0 {
			return ies.ApplicationDetectionInformation.ApplicationID, nil
		}
		return "", ErrIENotFound
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
