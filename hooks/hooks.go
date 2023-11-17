package hooks

import (
	"log"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
)

func Driver(caps ...capabilities.CapabilitiesFunc) (driver.WebDriver, func()) {
	d := driver.NewDriver(caps...)
	if d == nil {
		log.Fatal("Unable to start driver")
	}
	// tear down later
	return d, func() {
		// tear-down code here
		d.Quit()
		d.Service().Process.Kill()
	}
}
