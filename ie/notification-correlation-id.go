package ie

// NewNotificationCorrelationID creates a new NotificationCorrelationID IE.
func NewNotificationCorrelationID(uri string) *IE {
	return newStringIE(NotificationCorrelationID, uri)
}

// NotificationCorrelationID returns NotificationCorrelationID in string if the type of IE matches.
func (i *IE) NotificationCorrelationID() (string, error) {
	switch i.Type {
	case NotificationCorrelationID:
		return i.ValueAsString()
	case DirectReportingInformation:
		v, err := i.DirectReportingInformation()
		if err != nil {
			return "", err
		}
		return v.NotificationCorrelationID, nil
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
