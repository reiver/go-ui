package ui_driver

import (
	"io"
)

// Driver represents a connection between the UI server, and UI client.
type Driver interface {
	io.Closer
	io.Reader
	io.Writer
}
