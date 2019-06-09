package ui_driver

// Driver can dial open a connection between the UI server, and UI client.
type Driver interface {
	Dial() (Conn, error)
}
