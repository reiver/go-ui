package ui

// Close makes ui.Conn fit the io.Closer interface.
//
// Close closes the connection between the UI client, and the UI server.
func (receiver *internalConn) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	driver := receiver.driver
	if nil == driver {
		return errNilDriver
	}

	return driver.Close()
}
