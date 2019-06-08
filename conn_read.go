package ui

// Read makes ui.Conn fit the io.Reader interface.
func (receiver *internalConn) Read(p []byte) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	driver := receiver.driver
	if nil == driver {
		return 0, errNilDriver
	}

	if nil == p {
		return 0, nil
	}

	return driver.Read(p)
}
