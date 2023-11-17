package drivertest

import (
	"github.com/mcsymiv/go-gecko/driver"
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
)

// TestNewDriver
// Tests gecko driver and firefox instance
func TestNewDriver(t *testing.T) {

	// Connect to the WebDriver instance running locally
	d := driver.NewDriver(capabilities.ImplicitWait(3000))

	defer d.Service().Process.Kill()
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
	d := driver.NewDriver(
		capabilities.ImplicitWait(3000),
		capabilities.Firefox(moz),
	)

	// LIFO defer stack to quit firefox, and then kill driver proc
	defer d.Service().Process.Kill()
	defer d.Quit()
}
