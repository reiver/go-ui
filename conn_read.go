package ui

// Read makes ui.Conn fit the io.Reader interface.
func (receiver *internalConn) Read(p []byte) (int, error) {
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

	return driverConn.Read(p)
}
