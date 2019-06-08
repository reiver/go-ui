package ui

import (
	"github.com/reiver/go-ui/driver"
)

// Conn represents the connection between the UI client, and the UI server.
type Conn struct {
	driver ui_driver.Driver
}
