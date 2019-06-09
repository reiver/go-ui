package ui_driver

import (
	"io"
)

// Conn represents a connection between the UI server, and UI client.
type Conn interface {
	io.Closer
	io.Reader
	io.Writer
}
