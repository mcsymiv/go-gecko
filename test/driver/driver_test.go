package driver

import (
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

func TestDriver(t *testing.T) {

	s, err := driver.New(
		capabilities.ImplicitWait(3000),
	)
	if err != nil {
		t.Errorf("Creating new driver session: %+v", err)
	}
	defer s.Quit()

	_, err = s.GetStatus()
	if err != nil {
		t.Errorf("Status driver: %+v", err)
	}

	_, err = s.Open("https://www.google.com")
	if err != nil {
		t.Errorf("Url: %+v", err)
	}

	el, err := s.FindElement(element.ByCssSelector, "#APjFqb")
	if err != nil {
		t.Errorf("Element not found: %+v", err)
	}

	el.Click()

	attr, err := el.Attribute("id")
	if err != nil && attr == "" {
		t.Errorf("No attribute: %+v", err)
	}

	el.SendKeys("hello")
}
