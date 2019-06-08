package ui

import (
	"testing"
)

func TestInternalConnCloseNilReceiver(t *testing.T) {

	var conn *internalConn = nil

	if expected, actual := errNilReceiver, conn.Close(); expected != actual {
		t.Errorf("Expected (%T) %q, but actually got (%T) %q", expected, expected, actual, actual)
		return
	}
}
