package drivertest

import (
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/hooks"
	"log"
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
)

// TestNewDriver
// Tests chromedriver, driver capabilities
func TestNewDriver(t *testing.T) {

	// Connect to the WebDriver instance running locally
	d, tear := hooks.Driver(
		capabilities.ImplicitWait(3000),
		capabilities.BrowserName("chrome"),
		capabilities.Port(":9515"),
	)
	defer tear()

	d.Open("https://google.com")
	time.Sleep(2 * time.Second)
}

// TestDriver
// Opens firefox as default driver
// With default capabilities for driver with host, url: http://localhost:4444
func TestDriver(t *testing.T) {

	// Connect to the WebDriver instance running locally
	d, tear := hooks.Driver(
		capabilities.ImplicitWait(3000),
		// chrome caps commented out by design
		// capabilities.BrowserName("chrome"),
		// capabilities.Port(":9515"),
	)
	defer tear()

	d.Open("https://google.com")
	time.Sleep(2 * time.Second)
}

// TestDriver
// Opens firefox as default driver
// With default capabilities for driver with host, url: http://localhost:4444
func TestDriverWithExternalApiCall(t *testing.T) {

	// Connect to the WebDriver instance running locally
	d, tear := hooks.Driver(
		capabilities.ImplicitWait(3000),
	)
	defer tear()

	// Updates driver client url to call external API
	res, _ := d.MakeRequest(
		func(ro *driver.RequestOptions) {
			ro.Url = "https://pokeapi.co/api/v2/pokemon/ditto"
		},
	)

	log.Print(string(res))

	d.Open("https://google.com")
	time.Sleep(2 * time.Second)
}

// TestNewHeadlessDriver
// Tests gecko driver and firefox instance
//func TestNewHeadlessDriver(t *testing.T) {
//	// moz:options
//	// Uses headless arg
//	moz := &capabilities.MozOptions{
//		Args: []string{"-headless"},
//	}
//
//	// Connect to the WebDriver instance running locally
//	d := driver.NewDriver(
//		capabilities.ImplicitWait(3000),
//		capabilities.Firefox(moz),
//	)
//
//	// LIFO defer stack to quit firefox, and then kill driver proc
//	defer d.Service().Process.Kill()
//	defer d.Quit()
//}
