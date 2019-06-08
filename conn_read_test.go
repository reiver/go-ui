package ui

import (
	"testing"
)

func TestInternalConnReadNilReceiver(t *testing.T) {

	var conn *internalConn = nil

	var p []byte = nil
	actualN, actualErr := conn.Read(p)

	if expected, actual := 0, actualN; expected != actual {
		t.Errorf("Expected %d, but actually got %d", expected, actual)
		return
	}

	if expected, actual := errNilReceiver, actualErr; expected != actual {
		t.Errorf("Expected (%T) %q, but actually got (%T) %q", expected, expected, actual, actual)
		return
	}
}

func TestInternalConnReadNilDriver(t *testing.T) {

	var conn *internalConn = new(internalConn)

	var p []byte = nil
	actualN, actualErr := conn.Read(p)

	if expected, actual := 0, actualN; expected != actual {
		t.Errorf("Expected %d, but actually got %d", expected, actual)
		return
	}

	if expected, actual := errNilDriver, actualErr; expected != actual {
		t.Errorf("Expected (%T) %q, but actually got (%T) %q", expected, expected, actual, actual)
		return
	}
}
