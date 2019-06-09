package ui

import (
	"errors"
)

var (
	errNilDriver     = errors.New("Nil Driver")
	errNilDriverConn = errors.New("Nil Driver Conn")
	errNilReceiver   = errors.New("Nil Receiver")
	errNilRegistry   = errors.New("Nil Registry")
)
