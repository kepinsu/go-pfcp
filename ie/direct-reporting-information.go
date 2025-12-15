package ie

// NewDirectReportingInformation creates a new DirectReportingInformation IE.
func NewDirectReportingInformation(ies ...*IE) *IE {
	return newGroupedIE(DirectReportingInformation, 0, ies...)
}

// DirectReportingInformation returns the IEs above DirectReportingInformation if the type of IE matches.
func (i *IE) DirectReportingInformation() (*DirectReportingInformationFields, error) {
	switch i.Type {
	case DirectReportingInformation:
		return ParseDirectReportingInformationFields(i.Payload)
	case CreateSRR:
		ies, err := i.CreateSRR()
		if err != nil {
			return nil, err
		}
		if ies.DirectReportingInformation != nil {
			return ies.DirectReportingInformation, nil
		}
		return nil, ErrIENotFound
	case UpdateSRR:
		ies, err := i.UpdateSRR()
		if err != nil {
			return nil, err
		}
		if ies.DirectReportingInformation != nil {
			return ies.DirectReportingInformation, nil
		}
		return nil, ErrIENotFound
	default:
		return nil, &InvalidTypeError{Type: i.Type}
	}
}

// DirectReportingInformationFields is a set of fields in DirectReportingInformation IE.
//
// The contained fields are of type struct, as they are too complex to handle with
// existing (standard) types in Go.
type DirectReportingInformationFields struct {
	EventNotificationURI      string
	NotificationCorrelationID string
	ReportingFlags            uint8
}

// ParseDirectReportingInformationFields returns the IEs above DirectReportingInformation
func ParseDirectReportingInformationFields(b []byte) (*DirectReportingInformationFields, error) {

	// Parse all IES heres
	ies, err := ParseMultiIEs(b)
	if err != nil {
		return nil, err
	}
	bar := &DirectReportingInformationFields{}

	for _, ie := range ies {
		if ie == nil {
			continue
		}
		switch ie.Type {
		case EventNotificationURI:
			v, err := ie.EventNotificationURI()
			if err != nil {
				return bar, err
			}
			bar.EventNotificationURI = v
		case NotificationCorrelationID:
			v, err := ie.NotificationCorrelationID()
			if err != nil {
				return bar, err
			}
			bar.NotificationCorrelationID = v
		case ReportingFlags:
			v, err := ie.ReportingFlags()
			if err != nil {
				return bar, err
			}
			bar.ReportingFlags = v
		}
	}

	return bar, nil
}
