package announcer

import (
	"fmt"
	"github.com/oofpgDLD/goSkill/config-center/announcer/driver"
	"sync"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]driver.Driver)
)

func Register(name string, driver driver.Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("announcer: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("announcer: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

func Create(driverName string) (driver.Driver, error) {
	driversMu.RLock()
	driveri, ok := drivers[driverName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("announcer: unknown driver %q (forgotten import?)", driverName)
	}
	return driveri, nil
}