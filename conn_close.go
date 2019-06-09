package ui

// Close makes ui.Conn fit the io.Closer interface.
//
// Close closes the connection between the UI client, and the UI server.
func (receiver *internalConn) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	driverConn := receiver.driverConn
	if nil == driverConn {
		return errNilDriverConn
	}

	return driverConn.Close()
}
