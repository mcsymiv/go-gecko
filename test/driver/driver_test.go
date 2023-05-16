package driver

import (
	"testing"

	"github.com/mcsymiv/go-gecko/driver"
)

func TestDriver(t *testing.T) {

	// Starts firefox browser
	s := driver.New()
	defer s.Quit()
	s.Status()
}
