package driver

import (
	"os"
	"testing"

  "github.com/mcsymiv/go-gecko/session"
  "github.com/mcsymiv/go-gecko/capabilities"
)

func SetupTest() (session.WebDriver, func()) {
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

func TestMain(m *testing.M) {
	t := m.Run()
	os.Exit(t)
}
