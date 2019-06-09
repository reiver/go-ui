package ui_driver

import (
	"sync"
)

type internalRegistrar struct {
	mutex   sync.RWMutex
	drivers map[string]Driver
}

func (receiver *internalRegistrar) Fetch(name string) (Driver, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	drivers := receiver.drivers
	if nil == drivers {
		return nil, errNotFound
	}

	driver, loaded := drivers[name]
	if !loaded {
		return nil, errNotFound
	}

	return driver, nil
}
