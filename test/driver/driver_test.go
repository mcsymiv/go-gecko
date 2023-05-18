package driver

import (
	"testing"

	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

func TestDriver(t *testing.T) {

	// Starts firefox browser
	s := driver.New()
	defer s.Quit()
	driver.GetStatus()
	s.Open("https://www.google.com")
	el := s.FindElement(element.ByCssSelector, "#APjFqb")
	el.Click()
	el.GetAttribute("value")
	el.SendKeys("hello")
	el.GetAttribute("value")
}
