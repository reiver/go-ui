package ui

import (
	"github.com/reiver/go-ui/driver"
)

// Dial makes a UI client connection to the UI server service specified by ‘name’.
func Dial(name string) (Conn, error) {

	registry := ui_driver.Registry
	if nil == registry {
		return nil, errNilRegistry
	}

	driver, err := registry.Fetch(name)
	if nil != err {
		return nil, err
	}
	if nil == driver {
		return nil, errNilDriver
	}

	driverConn, err := driver.Dial()
	if nil != err {
		return nil, err
	}

	var conn internalConn
	conn.driverConn = driverConn

	return &conn, nil
}
