package ui

import (
	"errors"
)

var (
	errNilDriver   = errors.New("Nil Driver")
	errNilReceiver = errors.New("Nil Receiver")
	errNilRegistry = errors.New("Nil Registry")
)
