// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewApplicationInstanceID creates a new ApplicationInstanceID IE.
func NewApplicationInstanceID(id string) *IE {
	return newStringIE(ApplicationInstanceID, id)
}

// ApplicationInstanceID returns ApplicationInstanceID in string if the type of IE matches.
func (i *IE) ApplicationInstanceID() (string, error) {
	switch i.Type {
	case ApplicationInstanceID:
		return i.ValueAsString()
	case ApplicationDetectionInformation:
		ies, err := i.ApplicationDetectionInformation()
		if err != nil {
			return "", err
		}
		if len(ies.ApplicationInstanceID) > 0 {
			return ies.ApplicationInstanceID, nil
		}
		return "", ErrIENotFound
	case UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return "", err
		}
		if ies.ApplicationDetectionInformation != nil && len(ies.ApplicationDetectionInformation.ApplicationInstanceID) > 0 {
			return ies.ApplicationDetectionInformation.ApplicationInstanceID, nil
		}
		return "", ErrIENotFound
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
