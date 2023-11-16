package hooks

import (
	"log"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
)

func StartDriver() (driver.WebDriver, func()) {
	// Setup code here
	d, cmd := driver.NewDriver(
		capabilities.ImplicitWait(5000),
	)
	if d == nil {
		log.Fatal("Unable to start driver")
	}

	// tear down later
	return d, func() {
		// tear-down code here
		d.Quit()
		cmd.Process.Kill()
	}
}

func Driver(caps ...capabilities.CapabilitiesFunc) (driver.WebDriver, func()) {
	d, cmd := driver.NewDriver(caps...)
	if d == nil {
		log.Fatal("Unable to start driver")
	}
	// tear down later
	return d, func() {
		// tear-down code here
		d.Quit()
		cmd.Process.Kill()
	}
}
