package ie

// NewReportingFlags creates a new ReportingFlags IE.
func NewReportingFlags(dupl bool) *IE {
	var flags uint8
	if dupl {
		flags |= 0x01
	}
	return newUint8ValIE(ReportingFlags, flags)
}

// ReportingFlags returns the IEs above ReportingFlags if the type of IE matches.
func (i *IE) ReportingFlags() (uint8, error) {
	switch i.Type {
	case ReportingFlags:
		return i.ValueAsUint8()
	case CreateSRR:
		ies, err := i.CreateSRR()
		if err != nil {
			return 0, err
		}
		if ies.QoSMonitoringPerQoSFlowControlInformation != nil {
			return ies.DirectReportingInformation.ReportingFlags, nil
		}
		return 0, ErrIENotFound
	case UpdateSRR:
		ies, err := i.UpdateSRR()
		if err != nil {
			return 0, err
		}
		if ies.QoSMonitoringPerQoSFlowControlInformation != nil {
			return ies.DirectReportingInformation.ReportingFlags, nil
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
