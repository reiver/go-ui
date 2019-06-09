package ui

// Write makes ui.Conn fit the io.Writer interface.
func (receiver *internalConn) Write(p []byte) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	driverConn := receiver.driverConn
	if nil == driverConn {
		return 0, errNilDriverConn
	}

	if nil == p {
		return 0, nil
	}

	return driverConn.Write(p)
}
