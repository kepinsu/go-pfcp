package ie

// NewEventNotificationURI creates a new EventNotificationURI IE.
func NewEventNotificationURI(uri string) *IE {
	return newStringIE(EventNotificationURI, uri)
}

// EventNotificationURI returns EventNotificationURI in string if the type of IE matches.
func (i *IE) EventNotificationURI() (string, error) {
	switch i.Type {
	case EventNotificationURI:
		return i.ValueAsString()
	case DirectReportingInformation:
		v, err := i.DirectReportingInformation()
		if err != nil {
			return "", err
		}
		return v.EventNotificationURI, nil
	default:
		return "", &InvalidTypeError{Type: i.Type}
	}
}
