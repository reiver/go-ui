package ui

import (
	"testing"
)

func TestConnWriteNilReceiver(t *testing.T) {

	var conn *Conn = nil

	var p []byte = nil
	actualN, actualErr := conn.Write(p)

	if expected, actual := 0, actualN; expected != actual {
		t.Errorf("Expected %d, but actually got %d", expected, actual)
		return
	}

	if expected, actual := errNilReceiver, actualErr; expected != actual {
		t.Errorf("Expected (%T) %q, but actually got (%T) %q", expected, expected, actual, actual)
		return
	}
}

func TestConnWriteNilDriver(t *testing.T) {

	var conn *Conn = new(Conn)

	var p []byte = nil
	actualN, actualErr := conn.Write(p)

	if expected, actual := 0, actualN; expected != actual {
		t.Errorf("Expected %d, but actually got %d", expected, actual)
		return
	}

	if expected, actual := errNilDriver, actualErr; expected != actual {
		t.Errorf("Expected (%T) %q, but actually got (%T) %q", expected, expected, actual, actual)
		return
	}
}
