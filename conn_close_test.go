package ui

import (
	"testing"
)

func TestConnCloseNilReceiver(t *testing.T) {

	var conn *Conn = nil

	if expected, actual := errNilReceiver, conn.Close(); expected != actual {
		t.Errorf("Expected (%T) %q, but actually got (%T) %q", expected, expected, actual, actual)
		return
	}
}
