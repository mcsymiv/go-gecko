package drivertest

import (
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/hooks"
)

// TestDriver
// Opens firefox as default driver
// With default capabilities for driver with host, url: http://localhost:4444
func TestDriver(t *testing.T) {

	// Connect to the WebDriver instance running locally
	d, tear := hooks.Driver(
		// chrome caps commented out by design
		capabilities.BrowserName("firefox"),
	// capabilities.Port(":4444"),
	)
	defer tear()

	d.Open("https://google.com")
	time.Sleep(5 * time.Second)
}
