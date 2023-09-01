package driver

import (
	"log"
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/session"
	"github.com/mcsymiv/go-gecko/service"
)

// TestDriver
// Tests gecko driver and firefox instance
func TestDriver(t *testing.T) {

  // Starts gecko process
	cmd, err := service.Start()
	if err != nil {
		log.Fatal("start gecko", err)
	}

  // Connect to the WebDriver instance running locally
	d, err := session.New(capabilities.ImplicitWait(3000))
	if err != nil {
		log.Fatal("session start err", err)
	}

  defer cmd.Process.Kill()
	defer d.Quit()
}

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
