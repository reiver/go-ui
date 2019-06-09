package ui_driver

func init() {
	Registry = new(internalRegistrar)
}

var (
	Registry Registrar
)
