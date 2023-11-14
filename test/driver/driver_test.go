package driver

import (
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/session"
)

// TestNewDriver
// Tests gecko driver and firefox instance
func TestNewDriver(t *testing.T) {

	// Connect to the WebDriver instance running locally
	d, cmd := session.NewDriver(capabilities.ImplicitWait(3000))

	defer cmd.Process.Kill()
	defer d.Quit()
}

// TestNewHeadlessDriver
// Tests gecko driver and firefox instance
func TestNewHeadlessDriver(t *testing.T) {
	// moz:options
	// Uses headless arg
	moz := &capabilities.MozOptions{
		Args: []string{"-headless"},
	}

	// Connect to the WebDriver instance running locally
	d, cmd := session.NewDriver(
		capabilities.ImplicitWait(3000),
		capabilities.Firefox(moz),
	)

	// LIFO defer stack to quit firefox, and then kill driver proc
	defer cmd.Process.Kill()
	defer d.Quit()
}
