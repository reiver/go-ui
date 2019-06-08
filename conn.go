package ui

import (
	"github.com/reiver/go-ui/driver"

	"io"
)

// Conn represents the connection between the UI client, and the UI server.
type Conn interface {
	io.Closer
	io.Reader
	io.Writer
}

type internalConn struct {
	driver ui_driver.Driver
}
