package step

import (
	"os/exec"
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/service"
)

// StartDriver
// Convenience method for tests
// Wraps service.Start and driver.New functions
func StartDriver(t *testing.T) (*exec.Cmd, driver.WebDriver) {
	s, err := service.Start()
	if err != nil {
		defer s.Process.Kill()
		t.Fatalf("Error start service: %+v", err)
	}

	d, err := driver.New(capabilities.ImplicitWait(300))
	if err != nil {
		defer d.Quit()
		t.Fatalf("Error start session: %+v", err)
	}

	return s, d
}
