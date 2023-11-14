package hooks

import (
	"log"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/session"
)

func StartDriver() (session.WebDriver, func()) {
	// Setup code here
	d, cmd := session.NewDriver(
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
