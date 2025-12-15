// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewPredefinedRulesName creates a new PredefinedRulesName IE.
func NewPredefinedRulesName(name string) *IE {
	return newStringIE(PredefinedRulesName, name)
}

// PredefinedRulesName returns PredefinedRulesName in string if the type of IE matches.
func (i *IE) PredefinedRulesName() (string, error) {
	switch i.Type {
	case PredefinedRulesName:
		return i.ValueAsString()
	case UsageReportWithinSessionReportRequest:
		ies, err := i.UsageReport()
		if err != nil {
			return "", err
		}
		if len(ies.PredefinedRulesName) > 0 {
			return ies.PredefinedRulesName, nil
		}
		return "", ErrIENotFound
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
