package ui_driver

import (
	"bytes"
	"math/rand"
	"time"

	"testing"
)

func TestInternalRegistrarFetchNilReceiver(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	tests := []struct{
		Name string
	}{
		{
			Name: "",
		},



		{
			Name: " ",
		},
		{
			Name: "  ",
		},
		{
			Name: "   ",
		},



		{
			Name: "\n",
		},
		{
			Name: "\r",
		},
		{
			Name: "\t",
		},



		{
			Name: "apple",
		},
		{
			Name: "banana",
		},
		{
			Name: "cherry",
		},



		{
			Name: "Apple",
		},
		{
			Name: "Banana",
		},
		{
			Name: "Cherry",
		},



		{
			Name: "APPLE",
		},
		{
			Name: "BANANA",
		},
		{
			Name: "CHERRY",
		},



		{
			Name: "aPpLe",
		},
		{
			Name: "bAnAnA",
		},
		{
			Name: "cHeRrY",
		},



		{
			Name: "Hello world!",
		},
	}

	{
		var characters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz _./\\")

		for i:=0; i<100; i++ {
			var buffer bytes.Buffer

			{
				length := randomness.Intn(50)
				for ii:=0; ii<length; ii++ {
					buffer.WriteRune( characters[randomness.Intn(len(characters))] )
				}
			}


			test := struct{
				Name string
			}{
				Name: buffer.String(),
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var registrar *internalRegistrar = nil

		actualDriver, actualErr := registrar.Fetch(test.Name)

		if expected, actual := errNilReceiver, actualErr; expected != actual {
			t.Errorf("For test #%d (name=%q), expected (%T) %q, but actually got (%T) %q", testNumber, test.Name, expected, expected, actual, actual)
			continue
		}

		if expected, actual := Driver(nil), actualDriver; expected != actual {
			t.Errorf("For test #%d (name=%q), expected %#v, but actually got %#v", testNumber, test.Name, expected, actual)
			continue
		}
	}
}

func TestInternalRegistrarFetchNilDrivers(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	tests := []struct{
		Name string
	}{
		{
			Name: "",
		},



		{
			Name: " ",
		},
		{
			Name: "  ",
		},
		{
			Name: "   ",
		},



		{
			Name: "\n",
		},
		{
			Name: "\r",
		},
		{
			Name: "\t",
		},



		{
			Name: "apple",
		},
		{
			Name: "banana",
		},
		{
			Name: "cherry",
		},



		{
			Name: "Apple",
		},
		{
			Name: "Banana",
		},
		{
			Name: "Cherry",
		},



		{
			Name: "APPLE",
		},
		{
			Name: "BANANA",
		},
		{
			Name: "CHERRY",
		},



		{
			Name: "aPpLe",
		},
		{
			Name: "bAnAnA",
		},
		{
			Name: "cHeRrY",
		},



		{
			Name: "Hello world!",
		},
	}

	{
		var characters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz _./\\")

		for i:=0; i<100; i++ {
			var buffer bytes.Buffer

			{
				length := randomness.Intn(50)
				for ii:=0; ii<length; ii++ {
					buffer.WriteRune( characters[randomness.Intn(len(characters))] )
				}
			}


			test := struct{
				Name string
			}{
				Name: buffer.String(),
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var registrar *internalRegistrar = new(internalRegistrar)
		registrar.drivers = nil

		actualDriver, actualErr := registrar.Fetch(test.Name)

		if expected, actual := errNotFound, actualErr; expected != actual {
			t.Errorf("For test #%d (name=%q), expected (%T) %q, but actually got (%T) %q", testNumber, test.Name, expected, expected, actual, actual)
			continue
		}

		if expected, actual := Driver(nil), actualDriver; expected != actual {
			t.Errorf("For test #%d (name=%q), expected %#v, but actually got %#v", testNumber, test.Name, expected, actual)
			continue
		}
	}
}
