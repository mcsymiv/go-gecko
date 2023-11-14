package hooks

import (
	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/session"
)

func StartDriver() (session.WebDriver, func()) {
	// Setup code here
	d, cmd := session.NewDriver(
		capabilities.ImplicitWait(3000),
	)

	// tear down later
	return d, func() {
		// tear-down code here
		d.Quit()
		cmd.Process.Kill()
	}
}
