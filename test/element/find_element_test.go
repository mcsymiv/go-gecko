package element

import (
	"testing"

	"github.com/mcsymiv/go-gecko/driver"
)

func TestFindElement(t *testing.T) {

	// Starts firefox browser
	s := driver.New()
	defer s.Quit()

	s.Open("https://www.google.com")
	s.FindElement()
}
