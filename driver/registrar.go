package ui_driver

type Registrar interface {
	Fetch(string) (Driver, error)
}
