package step

import (
	"os/exec"
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/service"
	"github.com/mcsymiv/go-gecko/session"
)

// StartDriver
// Convenience method for tests
// Wraps service.Start and driver.New functions
func StartDriver(t *testing.T) (*exec.Cmd, session.WebDriver) {
	d, err := service.Start()
	if err != nil {
		defer d.Process.Kill()
		t.Fatalf("Error start service: %+v", err)
	}

	s, err := session.New(capabilities.ImplicitWait(300))
	if err != nil {
		defer s.Quit()
		t.Fatalf("Error start session: %+v", err)
	}

	return d, s
}
