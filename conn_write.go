package ui

// Write makes ui.Conn fit the io.Writer interface.
func (receiver *internalConn) Write(p []byte) (int, error) {
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

	return driver.Write(p)
}
