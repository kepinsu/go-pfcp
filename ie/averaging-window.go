// Copyright go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewAveragingWindow creates a new AveragingWindow IE.
func NewAveragingWindow(window uint32) *IE {
	return newUint32ValIE(AveragingWindow, window)
}

// AveragingWindow returns AveragingWindow in uint32 if the type of IE matches.
func (i *IE) AveragingWindow() (uint32, error) {
	switch i.Type {
	case AveragingWindow:
		return i.ValueAsUint32()
	case CreateQER:
		ies, err := i.CreateQER()
		if err != nil {
			return 0, err
		}
		return ies.AveragingWindows, nil
	case UpdateQER:
		ies, err := i.UpdateQER()
		if err != nil {
			return 0, err
		}
		return ies.AveragingWindows, nil
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
