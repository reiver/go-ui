package ui_driver

import (
	"errors"
)

var (
	errNilReceiver = errors.New("Nil Receiver")
	errNotFound    = errors.New("Not Found")
)
