package driver

import (
	"testing"

	"github.com/mcsymiv/go-gecko/driver"
)

func TestSession(t *testing.T) {

	// Starts firefox browser
	s := driver.New()
	defer s.Quit()
	s.Open("https://www.google.com")
	s.Status()
}
